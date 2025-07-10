package executor

import (
	"context"
	"log/slog"
	"reflect"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
	applog "github.com/eragon-mdi/ksu/pkg/log"
)

type TaskState interface {
	Advanced(context.Context, entity.Task) entity.Task
	Failed(context.Context, entity.Task) entity.Task
	Result(entity.Task) (entity.Task, error)
	Duration(entity.Task) entity.Task
}

// Дроп таски "мягкий", так как HardIOBoundWork запускается в этом же процессе и не имеет контекса (по ТЗ явно прописано не было)
// Принцип реализации, в  executor.cancels хранятся функции отмены по контексту для каждой таски такая функция и вызывается
func (e executor) DropTask(ctx context.Context, key string) {
	l := applog.GetCtxLogger(ctx).With("key", key)

	cancelCtxFunc, ok := e.cancels.Get(key)
	if !ok {
		l.Warn(ErrInvalidKey.Error())
		return
	}

	cancelCtxFunc()
	// e.cancels.Delete(key) - вызовется в отложенной функции контекста (AfterFunc), смотреть ниже
}

// Со стартом задачи запускается 2 горутины, одна отвечает за синхронизацию и и обновление информации о состоянии задачи,
// во второй запуск задачи (она же ограничена семафором)
//
// taskSyncCh - сигнал, уведомляющий что в хранилище была добавлена запись, с которой
// по мере выполнения задачи executor будет работать (обновлять статус задачи)
//
// ПРИЧИНА его использования - хочется не зная реализации и проблем репозитория
// сначала заупстить задачу, и потом уже ждать ответа сохранения задачи от репозитория
func (e executor) StartNewTask(c context.Context, syncCh chan struct{}, task entity.Task) context.CancelFunc {
	l := applog.GetCtxLogger(c).With("task", task)

	key := task.ID

	// Контекст задаёт сигнал отмены по дедлайну, таким образом используется как механизм дропа задач
	// вызывается либо по запросу в DropTask либо после окончания расчётов через defer
	ctx, cancel := context.WithCancel(c)
	e.cancels.Set(key, cancel)
	// после сигнала отмены по контексу удаляется cancel() из мапы
	_ = context.AfterFunc(ctx, func() {
		e.cancels.Delete(key)
		l.Debug("executor: afterFunc: deleting map elem func by key")
	})

	go func() {
		defer cancel() // каждый return = дроп задачи и удаление из мапы

		if err := e.sem.AcquireCtx(ctx); err != nil {
			l.Info("executor: go-control: task dropped before starting")
			return // задачу дропнули, пока горутина ждала очереди
		}

		task.SetStartedAtTime()

		// запсук задачи
		data := make(chan entity.ResultType, 1)
		go func() {
			defer apperrors.HandlePanic(l) // e.sem.Release() может запаниковать
			defer close(data)
			defer e.sem.Release()

			l.Debug("executor: go-io-task: staring do task")

			// В данной реализации это простейшая I/O bound задача, которая не поддерживает контекст
			// и не запускается как отдельный процесс.
			// Поэтому в любом случае при удалении задачи во время выполнения придётся ждать окончания I/O расчётов
			//
			// Тут есть ещё 2 варианта:
			// - если задача поддерживает ctx, то можно его прокинуть, тогда и ждать ненужного результата нет необходимости
			// - можно вынести HardIOBoundWork в отдельную компилируемую программу и вызывать как exec.CommandContext()
			r, err := HardIOBoundWork(nil)
			if err != nil { //cancel()
				l.Error("executor: go-io-task: task completed with err", slog.Any("cause", err))
				return
			}
			l.Debug("executor: go-io-task: task completed", slog.Any("result", r))

			result, ok := r.(entity.ResultType)
			if !ok {
				l.Warn("executor: go-io-task: task result type no assertion entity.ResultType",
					slog.Any("result", r),
					slog.String("actual_type", reflect.TypeOf(r).String()))
			}

			data <- result
		}()

		// дождаться пока в репозитории будет сущность,
		// только после этого можно писать результаты
		select {
		case <-syncCh:
		case <-ctx.Done():
			l.Info("executor: go-control: task dropped in running state")
			return
		}

		task = e.taskState.Advanced(ctx, task) // status running

		// сохранение результата
		select {
		case <-ctx.Done():
			l.Info("executor: go-control: task dropped in running state")
			return
		case res, ok := <-data:
			task = e.taskState.Duration(task)

			if !ok {
				e.taskState.Failed(ctx, task) // status failed
				return
			}

			e.taskState.Advanced(ctx, task.SetResult(res)) // status complete
		}
	}()

	return cancel
}

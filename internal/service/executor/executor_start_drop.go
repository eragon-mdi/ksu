package executor

import (
	"context"
	"log/slog"
	"time"

	"github.com/eragon-mdi/ksu/internal/entity"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
	applog "github.com/eragon-mdi/ksu/pkg/log"
)

type Executer interface {
	StartNewTask(context.Context, chan struct{}, entity.Task) context.CancelFunc
	DropTask(string) error
}

// Дроп таски "мягкий", так как HardIOBoundWork запускается в этом же процессе и не имеет контекса (по ТЗ явно прописано не было)
// Принцип реализации, в  executor.cancels хранятся функции отмены по контексту для каждой таски такая функция и вызывается
func (e executor) DropTask(key string) error {
	cancelCtxFunc, ok := e.cancels.Get(key)
	if !ok {
		return ErrInvalidKey
	}

	cancelCtxFunc()
	// e.cancels.Delete(key) - вызовется в отложенной функции контекста (AfterFunc), смотреть ниже

	return nil
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
	l := applog.GetCtxLogger(c)

	key := task.ID

	// Контекст задаёт сигнал отмены по дедлайну, таким образом используется как механизм дропа задач
	// вызывается либо по запросу в DropTask либо после окончания расчётов через defer
	ctx, cancel := context.WithCancel(c)
	e.cancels.Set(key, cancel)
	// после сигнала отмены по контексу удаляется cancel() из мапы
	_ = context.AfterFunc(ctx, func() {
		e.cancels.Delete(key)
		l.Debug("executor: afterFunc: deleting map key")
	})

	go func() {
		defer cancel() // каждый return = дроп задачи и удаление из мапы

		if err := e.sem.AcquireCtx(ctx); err != nil {
			l.Debug("executor: go-control: task dropped before starting")
			return // задачу дропнули, пока горутина ждала очереди
		}

		task.StartedAt = time.Now()

		// запсук задачи
		data := make(chan entity.ResultType, 1)
		errCh := make(chan error, 1)
		go func() {
			defer apperrors.HandlePanic(l) // e.sem.Release() может запаниковать
			defer close(errCh)
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
			result, err := HardIOBoundWork(nil)
			if err != nil {
				//cancel()
				errCh <- err
				l.Error("executor: go-io-task: task completed with err", slog.Any("cause", err))
				return
			}

			l.Debug("executor: go-io-task: task completed", slog.Any("result", result))

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

		if err := e.repository.UpdateTaskInfo(task.Update(entity.STATUS_RUNNING)); err != nil {
			l.Error("executor: go-control: err updated task status to running", slog.Any("cause", err))
			return
		}

		// сохранение результата
		select {
		case <-ctx.Done():
			l.Info("executor: go-control: task dropped in running state")
			return
		case res, ok := <-data:
			task.Duration = time.Duration(time.Since(task.StartedAt).Seconds())

			if !ok {
				if err := e.repository.UpdateTaskInfo(task.Update(entity.STATUS_FAILED)); err != nil {
					l.Error("executor: go-control: err updated task status to failed", slog.Any("cause", err))
					return
				}
				err := <-errCh
				l.Error("executor: go-control: task failed, no result", slog.Any("internal-task-err", err)) // канал закрыли, но данные не были переданы
				return
			}

			if err := e.repository.UpdateTaskInfo(task.Update(entity.STATUS_COMPLETED, res)); err != nil {
				l.Error("executor: go-control: err updated task status to completed", slog.Any("cause", err))
				return
			}
		}
	}()

	return cancel
}

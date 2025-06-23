package executor

import (
	"context"
	"time"

	"github.com/eragon-mdi/ksu/internal/entity"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
)

type Executer interface {
	StartNewTask(chan struct{}, entity.Task) context.CancelFunc
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
func (e executor) StartNewTask(syncCh chan struct{}, task entity.Task) context.CancelFunc {
	key := task.ID

	// Контекст задаёт сигнал отмены по дедлайну, таким образом используется как механизм дропа задач
	// вызывается либо по запросу в DropTask либо после окончания расчётов через defer
	ctx, cancel := context.WithCancel(context.Background())
	e.cancels.Set(key, cancel)
	// после сигнала отмены по контексу удаляется cancel() из мапы
	_ = context.AfterFunc(ctx, func() {
		// TODO log - логи уже не было возможности добавить, поэтому просто комментарии
		e.cancels.Delete(key)
	})

	go func() {
		defer cancel() // каждый return = дроп задачи и удаление из мапы

		if err := e.sem.AcquireCtx(ctx); err != nil {
			// TODO log
			return // задачу дропнули, пока горутина ждала очереди
		}

		task.StartedAt = time.Now()

		// запсук задачи
		data := make(chan entity.ResultType, 1)
		go func() {
			defer apperrors.HandlePanic() // e.sem.Release() может запаниковать
			defer close(data)
			defer e.sem.Release()

			// TODO log

			// В данной реализации это простейшая I/O bound задача, которая не поддерживает контекст
			// и не запускается как отдельный процесс.
			// Поэтому в любом случае при удалении задачи во время выполнения придётся ждать окончания I/O расчётов
			//
			// Тут есть ещё 2 варианта:
			// - если задача поддерживает ctx, то можно его прокинуть, тогда и ждать ненужного результата не нужно
			// - можно вынести HardIOBoundWork в отдельную компилируемую программу и вызывать как exec.CommandContext()
			result, err := HardIOBoundWork(nil)
			if err != nil {
				cancel()
				task.Duration = time.Duration(time.Since(task.StartedAt).Seconds())
				if err := e.repository.UpdateTaskInfo(task.Update(entity.STATUS_FAILED)); err != nil {
					// TODO log
					return
				}
				// TODO log
				return
			}

			data <- result
		}()

		// дождаться пока в репозитории будет сущность,
		// только после этого можно писать результаты
		select {
		case <-syncCh:
		case <-ctx.Done():
			// TODO log
			return
		}

		if err := e.repository.UpdateTaskInfo(task.Update(entity.STATUS_RUNNING)); err != nil {
			// TODO log
			return
		}

		// сохранение результата
		select {
		case <-ctx.Done():
			// TODO log
			return
		case res, ok := <-data:
			if !ok {
				// TODO log
				return
			}

			task.Duration = time.Duration(time.Since(task.StartedAt).Seconds())
			if err := e.repository.UpdateTaskInfo(task.Update(entity.STATUS_COMPLETED, res)); err != nil {
				// TODO log
				return
			}
		}

		// TODO log
	}()

	return cancel
}

package executor

import (
	"context"

	"github.com/eragon-mdi/ksu/internal/repository"
	taskstate "github.com/eragon-mdi/ksu/internal/service/task_state"
	"github.com/eragon-mdi/ksu/pkg/config"
	mapwithmutex "github.com/eragon-mdi/ksu/pkg/map_with_mutex"
	"github.com/eragon-mdi/ksu/pkg/semaphor"
)

type TaskExecutor interface {
	Executer
}

type executor struct {
	repository repository.Repositorier
	taskState  taskstate.TaskState

	sem     semaphor.Semaphorer
	cancels mapwithmutex.MaperWithMutex[context.CancelFunc]
}

// семафор - ограничивает кол-во задач (так как I/O bound work оевидно, что syscall)
func New(cfg config.Config, r repository.Repositorier) TaskExecutor {
	maxSemaphoreCount := cfg.App().Semaphore()

	return executor{
		repository: r,
		taskState:  taskstate.New(r),
		sem:        semaphor.New(maxSemaphoreCount),
		cancels:    mapwithmutex.New[context.CancelFunc](maxSemaphoreCount),
	}
}

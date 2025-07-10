package executor

import (
	"context"

	"github.com/eragon-mdi/ksu/internal/service"
	"github.com/eragon-mdi/ksu/pkg/config"
	mapwithmutex "github.com/eragon-mdi/ksu/pkg/map_with_mutex"
	"github.com/eragon-mdi/ksu/pkg/semaphor"
)

type executorImplement interface {
	service.Executer
}

type executor struct {
	taskState TaskState

	sem     semaphor.Semaphor
	cancels mapwithmutex.MaperWithMutex[context.CancelFunc]
}

// семафор - ограничивает кол-во задач (так как I/O bound work оевидно, что syscall)
// func New(cfg config.Config, r repository.Repositorier) TaskExecutor {
func New(cfg config.Config, ts TaskState) executorImplement {
	maxSemaphoreCount := cfg.App().Semaphore()

	return &executor{
		taskState: ts,
		sem:       semaphor.New(maxSemaphoreCount),
		cancels:   mapwithmutex.New[context.CancelFunc](maxSemaphoreCount),
	}
}

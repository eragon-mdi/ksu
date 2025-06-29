package executor

import (
	"context"
	"fmt"
	"time"

	"github.com/eragon-mdi/ksu/internal/repository"
	mapwithmutex "github.com/eragon-mdi/ksu/pkg/map_with_mutex"
	"github.com/eragon-mdi/ksu/pkg/semaphor"
)

type TaskExecutor interface {
	Executer
}

type executor struct {
	repository repository.Repositorier

	sem     semaphor.Semaphorer
	cancels mapwithmutex.MaperWithMutex[context.CancelFunc]
}

// семафор - ограничивает кол-во задач (так как I/O bound work оевидно, что syscall)
func New(r repository.Repositorier) TaskExecutor {
	return executor{
		repository: r,
		sem:        semaphor.New(MAX_SEMAPHORE),
		cancels:    mapwithmutex.New[context.CancelFunc](MAX_SEMAPHORE),
	}
}

const (
	TASK_TIMEOUT  = 5 * time.Minute
	MAX_SEMAPHORE = 16
)

var (
	prefix        = "executor: "
	ErrInvalidKey = fmt.Errorf("%sinvalid key", prefix)
)

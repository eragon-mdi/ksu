package service

import (
	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/internal/service/executor"
	taskstate "github.com/eragon-mdi/ksu/internal/service/task_state"
	"github.com/eragon-mdi/ksu/pkg/config"
)

type Servicer interface {
	Tasker
}

type service struct {
	repository repository.Repositorier

	executor  executor.TaskExecutor
	taskState taskstate.TaskState
}

func New(cfg config.Config, r repository.Repositorier) Servicer {
	return service{
		repository: r,
		executor:   executor.New(cfg, r),
		taskState:  taskstate.New(r),
	}
}

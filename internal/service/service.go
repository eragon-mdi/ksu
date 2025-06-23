package service

import (
	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/internal/service/executor"
)

type Servicer interface {
	Tasker
}

type service struct {
	repository repository.Repositorier
	executor   executor.TaskExecutor
}

func New(r repository.Repositorier) Servicer {
	return service{
		repository: r,
		executor:   executor.New(r),
	}
}

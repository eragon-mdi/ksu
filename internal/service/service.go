package service

import (
	"github.com/eragon-mdi/ksu/internal/handlers"
	"github.com/eragon-mdi/ksu/pkg/config"
)

type serviceImplement interface {
	handlers.Service
}

type service struct {
	repository Repository

	executor  Executer
	taskState TaskState
}

func New(cfg config.Config, r Repository, e Executer, tu TaskState) serviceImplement {
	return &service{
		repository: r,
		executor:   e,
		taskState:  tu,
	}
}

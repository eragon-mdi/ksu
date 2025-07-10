package taskstate

import (
	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	"github.com/eragon-mdi/ksu/internal/service"
	"github.com/eragon-mdi/ksu/internal/service/executor"
)

type taskStateImplement interface {
	service.TaskState
	executor.TaskState
}

type taskState struct {
	repository  Repository
	transitions map[entity.TaskStatusType]entity.TaskStatusType
}

func New(r Repository) taskStateImplement {
	return &taskState{
		repository: r,
		transitions: map[entity.TaskStatusType]entity.TaskStatusType{
			entity.STATUS_NULL:    entity.STATUS_PENDING,
			entity.STATUS_PENDING: entity.STATUS_RUNNING,
			entity.STATUS_RUNNING: entity.STATUS_COMPLETED,
		},
	}
}

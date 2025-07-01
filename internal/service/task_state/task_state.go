package taskstate

import (
	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	"github.com/eragon-mdi/ksu/internal/repository"
)

type TaskState interface {
	StateTransitioner
	TaskUtils
}

type taskState struct {
	repository  repository.Repositorier
	transitions map[entity.TaskStatusType]entity.TaskStatusType
}

func New(r repository.Repositorier) TaskState {
	return &taskState{
		repository: r,
		transitions: map[entity.TaskStatusType]entity.TaskStatusType{
			0:                     entity.STATUS_PENDING,
			entity.STATUS_PENDING: entity.STATUS_RUNNING,
			entity.STATUS_RUNNING: entity.STATUS_COMPLETED,
		},
	}
}

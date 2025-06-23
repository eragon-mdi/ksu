package service

import (
	"time"

	"github.com/eragon-mdi/ksu/internal/entity"
	"github.com/google/uuid"
)

func initTask() entity.Task {
	return entity.Task{
		ID: uuid.NewString(),
		TaskStatus: entity.TaskStatus{
			Status:    entity.STATUS_PENDING,
			CreatedAt: time.Now(),
			Duration:  0,
		},
	}
}

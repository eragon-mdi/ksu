package entity

import (
	"time"
)

const (
	STATUS_NULL      = "null"
	STATUS_PENDING   = "pending"
	STATUS_RUNNING   = "running"
	STATUS_COMPLETED = "completed"
	STATUS_FAILED    = "failed"
)

type TaskStatusType = string

// model: service - repository
type Task struct {
	ID string
	TaskResult
	TaskStatus
	StartedAt time.Time
}

type ResultType = string
type TaskResult struct {
	Result ResultType `json:"result"`
}

type TaskStatus struct {
	Status    TaskStatusType `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	Duration  time.Duration  `json:"duration"`
}

// DTO: handler - service
type TaskResultResponse = TaskResult

type TaskStatusResponse struct {
	StatusString string `json:"status"`
	TaskStatus
}

type TaskCreateResponse struct {
	ID string `json:"id"`
	TaskStatus
}

// .
type TaskResponse struct {
	ID        string        `json:"id"`
	Result    ResultType    `json:"result"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	Duration  time.Duration `json:"duration"`
}

func New(id string) Task {
	return Task{
		ID: id,
		TaskStatus: TaskStatus{
			Status:    STATUS_PENDING,
			CreatedAt: time.Now(),
			Duration:  0,
		},
	}
}

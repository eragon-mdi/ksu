package entity

import (
	"time"
)

type TaskStatusType = int

const (
	_ = iota
	STATUS_PENDING
	STATUS_RUNNING
	STATUS_COMPLETED
	STATUS_FAILED
)

// model: service - repository
type Task struct {
	ID string
	TaskResult
	TaskStatus
	StartedAt time.Time
}

type ResultType = any
type TaskResult struct {
	Result ResultType `json:"result"`
}

type TaskStatus struct {
	Status    int           `json:"-"`
	CreatedAt time.Time     `json:"created_at"`
	Duration  time.Duration `json:"duration"`
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

package service

import (
	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	"github.com/google/uuid"
)

func newTask() entity.Task {
	return entity.New(uuid.NewString())
}

func (s service) mapTasksToResponse(tasks []entity.Task) ([]entity.TaskResponse, error) {
	if len(tasks) < 1 {
		return []entity.TaskResponse{}, ErrGetTasks
	}

	res := make([]entity.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		task = s.taskState.Duration(task)
		res = append(res, task.Response())
	}

	return res, nil
}

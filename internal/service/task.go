package service

import (
	"context"
	"errors"

	"github.com/eragon-mdi/ksu/internal/entity"
)

type Tasker interface {
	CreateTask(context.Context) (entity.TaskCreateResponse, error)
	DropTask(context.Context, string) error
	GetTaskResult(string) (entity.TaskResultResponse, bool, error)
	GetTaskStatus(string) (entity.TaskStatusResponse, error)
	GetTasksAll() ([]entity.TaskResponse, error)
}

func (s service) CreateTask(c context.Context) (entity.TaskCreateResponse, error) {
	task := initTask()

	// запсук IO bound задачи
	taskSyncCh := make(chan struct{})
	defer close(taskSyncCh)
	dropTask := s.executor.StartNewTask(c, taskSyncCh, task)

	task, err := s.repository.SaveTask(task)
	if err != nil {
		dropTask()
		return entity.TaskCreateResponse{},
			errors.Join(errors.New("service: failed to save task"), err)
	}

	return task.ResponseCreate(), nil
}

func (s service) DropTask(c context.Context, id string) error {
	s.executor.DropTask(id)

	if err := s.repository.DeleteTask(id); err != nil {
		return errors.Join(errors.New("service: failed to delete task"), err)
	}

	return nil
}

// данные берём из хранилища, так как executor пишет в него же
func (s service) GetTaskResult(id string) (entity.TaskResultResponse, bool, error) {
	taskResult, err, taskNoCompleted := s.repository.GetTaskResultById(id)
	if taskNoCompleted {
		return entity.TaskResultResponse{},
			true,
			nil
	}

	if err != nil {
		return entity.TaskResultResponse{},
			false,
			errors.Join(errors.New("service: failed to get task result"), err)
	}

	return taskResult, false, nil
}

// данные берём из хранилища, так как executor пишет в него же
func (s service) GetTaskStatus(id string) (entity.TaskStatusResponse, error) {
	taskStatus, err := s.repository.GetTaskStatusById(id)
	if err != nil {
		return entity.TaskStatusResponse{},
			errors.Join(errors.New("service: failed to get task status"), err)
	}

	return taskStatus.Response(), nil
}

// .
func (s service) GetTasksAll() ([]entity.TaskResponse, error) {
	tasks, err := s.repository.GetAllTasks()
	if err != nil {
		return []entity.TaskResponse{},
			errors.Join(errors.New("service: failed to get task status"), err)
	}

	tasksResponse := make([]entity.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, task.Response())
	}

	return tasksResponse, nil
}

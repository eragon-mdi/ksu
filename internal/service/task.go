package service

import (
	"context"
	"errors"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

type Tasker interface {
	CreateTask(context.Context) (entity.TaskCreateResponse, error)
	DropTask(context.Context, string) error
	GetTaskResult(string) (entity.TaskResultResponse, error)
	GetTaskStatus(string) (entity.TaskStatusResponse, error)
	GetTasksAll() ([]entity.TaskResponse, error)
}

func (s service) CreateTask(c context.Context) (entity.TaskCreateResponse, error) {
	task := newTask()

	// запсук IO bound задачи
	taskSyncCh := make(chan struct{})
	defer close(taskSyncCh)
	dropTask := s.executor.StartNewTask(c, taskSyncCh, task)

	task, err := s.repository.SaveTask(task)
	if err != nil {
		dropTask()
		return entity.TaskCreateResponse{}, errors.Join(ErrBySave, err)
	}

	return task.CreateResponse(), nil
}

func (s service) DropTask(c context.Context, id string) error {
	s.executor.DropTask(id)

	if err := s.repository.DeleteTask(id); err != nil {
		return errors.Join(ErrByDelete, err)
	}

	return nil
}

func (s service) GetTaskResult(id string) (entity.TaskResultResponse, error) {
	task, err := s.repository.GetTaskResultById(id)
	if err != nil {
		return entity.TaskResultResponse{}, errors.Join(ErrGetTaskResult, err)
	}

	task, err = s.taskState.Result(task)

	return task.ResultResponse(), err
}

func (s service) GetTaskStatus(id string) (entity.TaskStatusResponse, error) {
	task, err := s.repository.GetTaskStatusById(id)
	if err != nil {
		return entity.TaskStatusResponse{}, errors.Join(ErrGetTaskStatus, err)
	}

	return task.StatusResponse(), nil
}

// .
func (s service) GetTasksAll() ([]entity.TaskResponse, error) {
	tasks, err := s.repository.GetAllTasks()
	if err != nil {
		return []entity.TaskResponse{}, errors.Join(ErrGetAllTask, err)
	}

	return s.mapTasksToResponse(tasks)
}

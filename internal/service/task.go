package service

import (
	"context"
	"errors"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
)

type Executer interface {
	StartNewTask(context.Context, chan struct{}, entity.Task) context.CancelFunc
	DropTask(context.Context, string)
}

type TaskState interface {
	Result(entity.Task) (entity.Task, error)
	Duration(entity.Task) entity.Task
}

type Repository interface {
	SaveTask(entity.Task) (entity.Task, error)
	DeleteTask(string) error
	GetTaskResultById(string) (entity.Task, error)
	GetTaskStatusById(string) (entity.Task, error)

	GetAllTasks() ([]entity.Task, error)
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
	s.executor.DropTask(c, id)

	if err := s.repository.DeleteTask(id); err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return errors.Join(apperrors.NotFound, err)
		}
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

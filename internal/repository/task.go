package repository

import (
	"errors"
	"fmt"

	"github.com/eragon-mdi/ksu/internal/entity"
)

type Tasker interface {
	SaveTask(entity.Task) (entity.Task, error)
	DeleteTask(string) error
	GetTaskResultById(string) (entity.TaskResult, error, bool)
	GetTaskStatusById(string) (entity.TaskStatus, error)

	UpdateTaskInfo(entity.Task) error

	GetAllTasks() ([]entity.Task, error)
}

func (r repository) SaveTask(task entity.Task) (entity.Task, error) {
	task, err := r.storage.InsertTaskWithReturns(task.ID, task)
	if err != nil {
		return task,
			errors.Join(errors.New("repository: failed insert task"), err)
	}

	return task, nil
}

func (r repository) DeleteTask(id string) error {
	err := r.storage.DeleteTaskById(id)
	if err != nil {
		return errors.Join(errors.New("repository: failed delete task"), err)
	}
	return err
}

func (r repository) GetTaskResultById(id string) (entity.TaskResult, error, bool) {
	task, err := r.storage.SelectAllInfoById(id)
	if task.Status == entity.STATUS_PENDING ||
		task.Status == entity.STATUS_RUNNING {
		return entity.TaskResult{}, nil, true
	}

	if err != nil {
		return entity.TaskResult{},
			errors.Join(errors.New("repository: failed get task reuslt by id"), err),
			false
	}

	return entity.TaskResult{Result: task.Result}, nil, false
}

func (r repository) GetTaskStatusById(id string) (entity.TaskStatus, error) {
	task, err := r.storage.SelectAllInfoById(id)
	if err != nil {
		return entity.TaskStatus{},
			errors.Join(errors.New("repository: failed get task status by id"), err)
	}

	return task.EntityTaskStatus(), nil
}

func (r repository) UpdateTaskInfo(task entity.Task) error {
	err := r.storage.UpdateTask(task)
	if err != nil {
		return errors.Join(fmt.Errorf("repository: failed update task info (%d %v)\n%w", task.Status, task.Result, err))
	}
	return nil
}

// .
func (r repository) GetAllTasks() ([]entity.Task, error) {
	data := r.storage.SelectAll()

	if len(data) == 0 {
		return []entity.Task{}, errors.New("repository: empty")
	}

	tasks := append(make([]entity.Task, 0, len(data)), data...)

	return tasks, nil
}

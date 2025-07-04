package repository

import (
	"errors"
	"fmt"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

type Tasker interface {
	SaveTask(entity.Task) (entity.Task, error)
	DeleteTask(string) error
	GetTaskResultById(string) (entity.Task, error)
	GetTaskStatusById(string) (entity.Task, error)

	UpdateTaskInfo(entity.Task) error

	GetAllTasks() ([]entity.Task, error)
}

func (r repository) SaveTask(task entity.Task) (entity.Task, error) {
	task, err := r.storage.InsertTaskWithReturns(task.ID, task)
	if err != nil {
		return task, errors.Join(ErrInsertTask, err)
	}

	return task, nil
}

func (r repository) DeleteTask(id string) error {
	err := r.storage.DeleteTaskById(id)
	if err != nil {
		return errors.Join(ErrDeleteTask, NotFound, err)
	}
	return err
}

func (r repository) GetTaskResultById(id string) (entity.Task, error) {
	task, err := r.storage.SelectAllInfoById(id)
	if err != nil {
		return entity.Task{}, errors.Join(ErrGetResultByID, err)
	}

	return task, nil
}

func (r repository) GetTaskStatusById(id string) (entity.Task, error) {
	task, err := r.storage.SelectAllInfoById(id)
	if err != nil {
		return entity.Task{}, errors.Join(ErrGetStatusByID, err)
	}

	return task, nil
}

func (r repository) UpdateTaskInfo(task entity.Task) error {
	err := r.storage.UpdateTask(task)
	if err != nil {
		return errors.Join(fmt.Errorf("%w (%d %v)\n%w", ErrUpdateTaskInfo, task.Status, task.Result, err))
	}
	return nil
}

// .
func (r repository) GetAllTasks() ([]entity.Task, error) {
	return r.storage.SelectAll(), nil
}

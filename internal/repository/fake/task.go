package fakerepo

import (
	"errors"
	"fmt"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

type FakeStorage interface {
	InsertTaskWithReturns(string, entity.Task) (entity.Task, error)
	DeleteTaskById(string) error
	SelectAllInfoById(string) (entity.Task, error)
	UpdateTask(entity.Task) error
	SelectAll() []entity.Task
}

func (r fakeRepository) SaveTask(task entity.Task) (entity.Task, error) {
	task, err := r.storage.InsertTaskWithReturns(task.ID, task)
	if err != nil {
		return task, errors.Join(ErrInsertTask, err)
	}

	return task, nil
}

func (r fakeRepository) DeleteTask(id string) error {
	err := r.storage.DeleteTaskById(id)
	if err != nil {
		return errors.Join(ErrDeleteTask, NotFound, err)
	}
	return err
}

func (r fakeRepository) GetTaskResultById(id string) (entity.Task, error) {
	task, err := r.storage.SelectAllInfoById(id)
	if err != nil {
		return entity.Task{}, errors.Join(ErrGetResultByID, err)
	}

	return task, nil
}

func (r fakeRepository) GetTaskStatusById(id string) (entity.Task, error) {
	task, err := r.storage.SelectAllInfoById(id)
	if err != nil {
		return entity.Task{}, errors.Join(ErrGetStatusByID, err)
	}

	return task, nil
}

func (r fakeRepository) UpdateTaskInfo(task entity.Task) error {
	err := r.storage.UpdateTask(task)
	if err != nil {
		return errors.Join(fmt.Errorf("%w (%s %v)\n%w", ErrUpdateTaskInfo, task.Status, task.Result, err))
	}
	return nil
}

// .
func (r fakeRepository) GetAllTasks() ([]entity.Task, error) {
	return r.storage.SelectAll(), nil
}

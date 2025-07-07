package sqlrepo

import (
	"context"
	"database/sql"
	"errors"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

/*
Exec - ответ не нужен, Insert Update Delete
Query - Select
QueryRow - тоже но одна строка
*/
const QUERY_SAVE_TASK_RETURN = `
INSERT INTO task
	(id, result, status, created_at, duration, started_at)
VALUES
	($1, $2, $3, $4, $5, $6)
RETURNING
	(id, result, status, created_at, duration, started_at)
`

func (r sqlRepository) SaveTask(t entity.Task) (task entity.Task, err error) {
	ctx := context.TODO()

	err = r.storage.defaultTx(ctx).begin(
		func(tx *sql.Tx) error {
			row := tx.QueryRowContext(ctx, QUERY_SAVE_TASK_RETURN,
				t.ID,
				t.Result,
				t.Status,
				t.CreatedAt,
				t.Duration,
				t.StartedAt,
			)

			// sql.Null...Types можно добавить с проверкой за транзакцией
			err := row.Scan(&task.ID, &task.Result, &task.Status, &task.CreatedAt, &task.Duration, &task.StartedAt)

			return err
		},
	)

	if err != nil {
		return task, errors.Join(ErrInsertTask, err)
	}

	return task, err
}

func (r sqlRepository) DeleteTask(id string) error {
	//err := r.storage.DeleteTaskById(id)
	//if err != nil {
	//	return errors.Join(ErrDeleteTask, NotFound, err)
	//}
	return nil
}

func (r sqlRepository) GetTaskResultById(id string) (entity.Task, error) {
	//task, err := r.storage.SelectAllInfoById(id)
	//if err != nil {
	//	return entity.Task{}, errors.Join(ErrGetResultByID, err)
	//}

	return entity.Task{}, nil
}

func (r sqlRepository) GetTaskStatusById(id string) (entity.Task, error) {
	//task, err := r.storage.SelectAllInfoById(id)
	//if err != nil {
	//	return entity.Task{}, errors.Join(ErrGetStatusByID, err)
	//}

	return entity.Task{}, nil
}

func (r sqlRepository) UpdateTaskInfo(task entity.Task) error {
	//err := r.storage.UpdateTask(task)
	//if err != nil {
	//	return errors.Join(fmt.Errorf("%w (%d %v)\n%w", ErrUpdateTaskInfo, task.Status, task.Result, err))
	//}

	return nil
}

// .
func (r sqlRepository) GetAllTasks() ([]entity.Task, error) {
	//return r.storage.SelectAll(), nil
	return nil, nil
}

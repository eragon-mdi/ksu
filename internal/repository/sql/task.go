package sqlrepo

import (
	"context"
	"database/sql"
	"errors"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

func (r sqlRepository) SaveTask(t entity.Task) (task entity.Task, err error) {
	ctx := context.TODO()

	err = r.defaultTx(ctx).begin(
		func(tx *sql.Tx) error {
			row := tx.QueryRowContext(ctx, QUERY_SAVE_TASK_RETURN,
				t.ID,
				t.Result,
				t.Status,
				t.CreatedAt,
				t.Duration,
				t.StartedAt,
			)

			// TODO sql.Null...Types можно добавить с проверкой за транзакцией
			err := row.Scan(&task.ID, &task.Result, &task.Status, &task.CreatedAt, &task.Duration, &task.StartedAt)

			return err
		},
	)

	if err != nil {
		return entity.Task{}, errors.Join(ErrInsertTask, err)
	}

	return task, err
}

func (r sqlRepository) DeleteTask(id string) (err error) {
	ctx := context.TODO()

	err = r.defaultTx(ctx).begin(
		func(tx *sql.Tx) error {
			res, err := tx.ExecContext(ctx, QUERY_DELETE_TASK, id)
			if err != nil {
				return err
			}

			c, err := res.RowsAffected()
			if err != nil {
				return err
			}
			if c != 1 {
				return NotFound
			}

			return err
		})

	if err != nil {
		return errors.Join(ErrDeleteTask, err)
	}

	return err
}

func (r sqlRepository) GetTaskResultById(id string) (task entity.Task, err error) {
	ctx := context.TODO()

	err = r.defaultTx(ctx).begin(func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, QUERY_GET_TASK_RESULT, id)

		err := row.Scan(&task.Status, &task.CreatedAt, &task.Duration)

		return err
	})

	if err != nil {
		return entity.Task{}, errors.Join(ErrGetResultByID, err)
	}

	return entity.Task{}, nil
}

func (r sqlRepository) GetTaskStatusById(id string) (task entity.Task, err error) {
	ctx := context.TODO()

	err = r.defaultTx(ctx).begin(func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, QUERY_GET_TASK_STATUS, id)

		err := row.Scan(&task.Status, &task.CreatedAt, &task.Duration)

		return err
	})

	if err != nil {
		return entity.Task{}, errors.Join(ErrGetStatusByID, err)
	}

	return entity.Task{}, nil
}

func (r sqlRepository) UpdateTaskInfo(t entity.Task) (err error) {
	ctx := context.TODO()

	err = r.defaultTx(ctx).begin(func(tx *sql.Tx) error {
		res, err := tx.ExecContext(ctx, QUERY_UPDATE_TASK,
			t.ID,
			t.Status,
			t.Duration,
			t.StartedAt,
			t.Result,
		)
		if err != nil {
			return err
		}

		c, err := res.RowsAffected()
		if err != nil {
			return err // TODO join
		}
		if c != 1 {
			return err // TODO join
		}

		return err
	})

	if err != nil {
		return errors.Join(ErrUpdateTaskInfo, err)
	}

	return nil
}

// .
func (r sqlRepository) GetAllTasks() (tasks []entity.Task, err error) {
	ctx := context.TODO()

	err = r.defaultTx(ctx).begin(func(tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, QUERY_GET_ALL_TASKS)
		if err != nil {
			return err
		}
		defer rows.Close()

		tasks = make([]entity.Task, 0)

		for rows.Next() {
			var task entity.Task
			err := rows.Scan(&task.ID, &task.Result, &task.Status, &task.CreatedAt, &task.Duration, &task.StartedAt)
			if err != nil {
				continue
			}

			tasks = append(tasks, task)
		}

		if err := rows.Err(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

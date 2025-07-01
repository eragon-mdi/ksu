package repository

import "errors"

const prefic = "repository: "

var (
	ErrInsertTask     = errors.New(prefic + "failed insert task")
	ErrDeleteTask     = errors.New(prefic + "failed delete task")
	ErrGetResultByID  = errors.New(prefic + "failed get task result by id")
	ErrGetStatusByID  = errors.New(prefic + "failed get task status by id")
	ErrUpdateTaskInfo = errors.New(prefic + "failed update task info")
)

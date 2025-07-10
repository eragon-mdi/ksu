package fakerepo

import (
	"errors"

	"github.com/eragon-mdi/ksu/pkg/apperrors"
)

const prefic = "repository: "

var (
	NotFound          = apperrors.NotFound
	ErrInsertTask     = errors.New(prefic + "failed insert task")
	ErrDeleteTask     = errors.New(prefic + "failed delete task")
	ErrGetResultByID  = errors.New(prefic + "failed get task result by id")
	ErrGetStatusByID  = errors.New(prefic + "failed get task status by id")
	ErrUpdateTaskInfo = errors.New(prefic + "failed update task info")
)

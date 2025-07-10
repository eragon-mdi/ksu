package handlers

import (
	"context"
	"errors"
	"net/http"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
	applog "github.com/eragon-mdi/ksu/pkg/log"
	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateTask(context.Context) (entity.TaskCreateResponse, error)
	DropTask(context.Context, string) error
	GetTaskResult(string) (entity.TaskResultResponse, error)
	GetTaskStatus(string) (entity.TaskStatusResponse, error)
	GetTasksAll() ([]entity.TaskResponse, error)
}

// .
func (h handler) NewTask(c echo.Context) error {
	l := applog.GetRequestCtxLogger(c)
	ctx := applog.CtxWithLogger(l)

	taskResponse, err := h.service.CreateTask(ctx)
	if err != nil {
		l.Error("failed to create task", withErr(err))
		return echo.NewHTTPError(http.StatusInternalServerError, apperrors.ErrInternal)
	}

	l = l.With("task", taskResponse)
	l.Debug("task created OK")

	return c.JSON(http.StatusCreated, taskResponse)
}

// удаление мягкое
func (h handler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	l := applog.GetRequestCtxLogger(c).With("task_id", id)
	l.Debug("deleting task")

	if !validateId(id) {
		l.Error("invalid task id")
		return echo.NewHTTPError(http.StatusBadRequest, apperrors.ErrInvalidID)
	}

	ctx := applog.CtxWithLogger(l)
	if err := h.service.DropTask(ctx, id); err != nil {
		l.Error("failed to delete task", withErr(err))

		if errors.Is(err, apperrors.NotFound) {
			return echo.NewHTTPError(http.StatusNotFound, apperrors.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, apperrors.ErrInternal)
	}

	l.Debug("task deleted OK")

	return c.NoContent(http.StatusNoContent)
}

// .
func (h handler) GetTaskResult(c echo.Context) error {
	id := c.Param("id")

	l := applog.GetRequestCtxLogger(c).With("key", id)
	l.Debug("getting task result")

	if !validateId(id) {
		l.Error("invalid task id")
		return echo.NewHTTPError(http.StatusBadRequest, apperrors.ErrInvalidID)
	}

	taskResult, err := h.service.GetTaskResult(id)
	if err != nil {
		l.Error("failed to get task result", withErr(err))
		return echo.NewHTTPError(http.StatusInternalServerError, apperrors.ErrInternal)
	}

	l = l.With("task_status", taskResult)
	l.Debug("task result getted OK")

	return c.JSON(http.StatusOK, taskResult)
}

// .
func (h handler) GetTaskStatus(c echo.Context) error {
	id := c.Param("id")

	l := applog.GetRequestCtxLogger(c).With("key", id)
	l.Debug("getting task status")

	if !validateId(id) {
		l.Error("invalid task id")
		return echo.NewHTTPError(http.StatusBadRequest, apperrors.ErrInvalidID)
	}

	taskStatus, err := h.service.GetTaskStatus(id)
	if err != nil {
		l.Error("failed to get task status", withErr(err))
		return echo.NewHTTPError(http.StatusInternalServerError, apperrors.ErrInternal)
	}

	l = l.With("task_status", taskStatus)
	l.Debug("task result status OK")

	return c.JSON(http.StatusOK, taskStatus)
}

// .
func (h handler) GetAllTasks(c echo.Context) error {
	l := applog.GetRequestCtxLogger(c)
	l.Debug("getting tasks")

	tasks, err := h.service.GetTasksAll()
	if err != nil {
		l.Error("failed to get tasks", withErr(err))
		return echo.NewHTTPError(http.StatusInternalServerError, apperrors.ErrInternal)
	}

	l.Debug("tasks get status OK")

	return c.JSON(http.StatusOK, tasks)
}

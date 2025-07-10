package routes

import (
	"github.com/eragon-mdi/ksu/pkg/server/router"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	NewTask(c echo.Context) error
	DeleteTask(c echo.Context) error
	GetTaskResult(c echo.Context) error
	GetTaskStatus(c echo.Context) error

	GetAllTasks(c echo.Context) error
}

func WithTaskRoutes(r router.Router, h Handler) router.Router {
	group := r.Echo().Group("/task")

	group.POST("", h.NewTask)
	group.DELETE("/:id", h.DeleteTask)
	group.GET("/:id/result", h.GetTaskResult)
	group.GET("/:id/status", h.GetTaskStatus)

	// такой ручки по ТЗ не было, но она удобна для демонстрации корректной работы сервиса
	group.GET("", h.GetAllTasks)

	return r
}

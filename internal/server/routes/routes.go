package routes

import (
	"github.com/eragon-mdi/ksu/internal/handlers"
	"github.com/eragon-mdi/ksu/pkg/server/router"
)

func WithTaskRoutes(r router.Router, h handlers.Handler) router.Router {
	group := r.Echo().Group("/task")

	group.POST("", h.NewTask)
	group.DELETE("/:id", h.DeleteTask)
	group.GET("/:id/result", h.GetTaskResult)
	group.GET("/:id/status", h.GetTaskStatus)

	// такой ручки по ТЗ не было, но она удобна для демонстрации корректной работы сервиса
	group.GET("", h.GetAllTasks)

	return r
}

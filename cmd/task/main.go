// Михайлюк Дмитрий Игоревич
// тестовое в компанию work-mate
package main

import (
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/eragon-mdi/ksu/internal/handlers"
	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/internal/service"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
	applog "github.com/eragon-mdi/ksu/pkg/log"
	"github.com/eragon-mdi/ksu/pkg/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	// pprof
	//	go func() {
	//		http.ListenAndServe("0.0.0.0:6060", nil)
	//	}()

	// слои
	fakeStorage, err := storage.Get()
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.New(fakeStorage)
	service := service.New(repository)
	handlers := handlers.New(service)

	//
	e := echo.New()
	e.Use(applog.InitMiddlewareLogging())
	e.HTTPErrorHandler = apperrors.CustomHTTPErrorHandler

	//
	taskGroup := e.Group("/task")
	taskGroup.POST("", handlers.NewTask)
	taskGroup.DELETE("/:id", handlers.DeleteTask)
	taskGroup.GET("/:id/result", handlers.GetTaskResult)
	taskGroup.GET("/:id/status", handlers.GetTaskStatus)
	// такой ручки по ТЗ не было, но она удобна для демонстрации корректной работы сервиса
	taskGroup.GET("", handlers.GetAllTasks)

	if err := customHttpServer(e).ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func customHttpServer(e *echo.Echo) *http.Server {
	return &http.Server{
		Addr:              "0.0.0.0:8080",
		Handler:           e,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
}

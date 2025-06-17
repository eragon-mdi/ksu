// Михайлюк Дмитрий Игоревич
// тестовое в компанию work-mate
package main

import (
	"log"

	"github.com/eragon-mdi/ksu/internal/handlers"
	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/internal/service"
	"github.com/eragon-mdi/ksu/pkg/apperrors"
	"github.com/eragon-mdi/ksu/pkg/storage"
	"github.com/labstack/echo/v4"
)

func main() {

	// слои
	fakeStorage, err := storage.Get() //  8*8*8
	if err != nil {
		log.Fatal(err)
	}
	r := repository.New(fakeStorage)
	s := service.New(r)
	h := handlers.New(s)

	//
	e := echo.New()
	e.HTTPErrorHandler = apperrors.CustomHTTPErrorHandler

	//
	taskGroup := e.Group("/task")
	_ = taskGroup
	_ = h
	//taskGroup.POST("", h.NewTask)
	//taskGroup.DELETE("/:id", h.DeleteTask)
	//taskGroup.GET("/:id/result", h.GetTaskResult)
	//taskGroup.GET("/:id/status", h.GetTaskStatus)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}

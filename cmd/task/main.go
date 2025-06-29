// Михайлюк Дмитрий Игоревич
// тестовое в компанию work-mate
package main

import (
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/eragon-mdi/ksu/internal/handlers"
	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/internal/server/routes"
	"github.com/eragon-mdi/ksu/internal/service"
	"github.com/eragon-mdi/ksu/pkg/config"
	applog "github.com/eragon-mdi/ksu/pkg/log"
	"github.com/eragon-mdi/ksu/pkg/server"
	"github.com/eragon-mdi/ksu/pkg/server/router"
	"github.com/eragon-mdi/ksu/pkg/storage"
)

func main() {
	//pprof
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	applog.SetDefaultBaseLogger(cfg)

	fakeStorage, err := storage.Get()
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.New(fakeStorage)
	service := service.New(repository)
	handlers := handlers.New(service)

	//
	router := router.New()
	router = routes.WithTaskRoutes(router, handlers)
	serv := server.New(router.Handler(), cfg)

	go serv.Start()

	server.WaitingForShutdownSignal()

	serv.GracefulShutdown()
	// db.Close()
}

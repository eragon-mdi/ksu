// Михайлюк Дмитрий Игоревич
// тестовое в компанию work-mate
package main

import (
	"log"

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
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	applog.SetDefaultBaseLogger(cfg)

	fakeStorage, err := storage.Get()
	if err != nil {
		log.Fatal(err)
	}
	r := repository.New(fakeStorage)
	s := service.New(cfg, r)
	h := handlers.New(s)

	//
	router := router.New()
	router = routes.WithTaskRoutes(router, h)
	serv := server.New(router.Handler(), cfg)

	go serv.Start()

	server.WaitingForShutdownSignal()

	serv.GracefulShutdown()
}

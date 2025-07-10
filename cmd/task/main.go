package main

import (
	"log"

	"github.com/eragon-mdi/ksu/internal/handlers"
	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/internal/server/routes"
	"github.com/eragon-mdi/ksu/internal/service"
	"github.com/eragon-mdi/ksu/internal/service/executor"
	taskstate "github.com/eragon-mdi/ksu/internal/service/task_state"
	"github.com/eragon-mdi/ksu/pkg/config"
	applog "github.com/eragon-mdi/ksu/pkg/log"
	"github.com/eragon-mdi/ksu/pkg/log/clickhouse"
	"github.com/eragon-mdi/ksu/pkg/server"
	"github.com/eragon-mdi/ksu/pkg/server/router"
	"github.com/eragon-mdi/ksu/pkg/storage"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	clickH := clickhouse.New(cfg)
	applog.SetDefaultBaseLogger(cfg, clickH)

	stor, err := storage.Get(cfg)
	if err != nil {
		log.Fatal(err)
	}
	r := repository.New(stor)
	ts := taskstate.New(r)
	e := executor.New(cfg, ts)
	s := service.New(cfg, r, e, ts)
	h := handlers.New(s)

	//
	router := router.New()
	router = routes.WithTaskRoutes(router, h)
	serv := server.New(router.Handler(), cfg)

	go serv.Start()

	server.WaitingForShutdownSignal()

	serv.GracefulShutdown()
	clickH.GracefulShutdown()
	storage.GracefulShutdown()
}

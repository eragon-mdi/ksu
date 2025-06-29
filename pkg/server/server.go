package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/eragon-mdi/ksu/pkg/config"
)

type Server interface {
	Start()
	GracefulShutdown()
}

type server struct {
	*http.Server
}

func New(rHandler http.Handler, cfg config.Config) Server {
	srvCfg := cfg.Server()

	return server{
		Server: &http.Server{
			Addr:              fmt.Sprintf("%s:%s", srvCfg.Addr(), srvCfg.Port()),
			Handler:           rHandler,
			ReadTimeout:       srvCfg.ReadTimeout(),
			WriteTimeout:      srvCfg.WriteTimeout(),
			ReadHeaderTimeout: srvCfg.ReadTimeout(),
			IdleTimeout:       srvCfg.ReadTimeout(),
		},
	}
}

func (s server) Start() {
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	slog.Default().Debug("server started", slog.Any("addr:", s.Addr))
}

func (s server) GracefulShutdown() {
	if err := s.Shutdown(context.Background()); err != nil {
		slog.Default().Error("error ocured on server shutdown", slog.Any("cause", err))
	}
}

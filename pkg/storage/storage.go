package storage

import (
	"log/slog"
	"sync"

	"github.com/eragon-mdi/ksu/internal/repository"
	"github.com/eragon-mdi/ksu/pkg/config"
)

type storageImplement interface {
	repository.Storage
	Init
}

type Init interface {
	Connect(config.Config) error
	Migrate(config.Config) error
	Shutdown() error
}

var (
	once sync.Once
	stor storageImplement

	fabric = map[string]Init{}
)

func Get(cfg config.Config) (storageImplement, error) {
	var err error
	once.Do(func() {
		strCfg := cfg.Storage()

		stor, err = getStorageByType(strCfg.Type())
		if err != nil {
			return
		}

		err = stor.Connect(cfg)
		if err != nil {
			return
		}

		if cfg.Storage().NeedMigrate() {
			err = stor.Migrate(cfg)
		}
	})
	if err != nil {
		return nil, err
	}

	return stor, err
}

func GracefulShutdown() {
	if err := stor.Shutdown(); err != nil {
		slog.Default().Error("error ocured on storage drop connection", slog.Any("cause", err))
	}
}

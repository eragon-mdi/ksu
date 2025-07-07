package storage

import (
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
}

var (
	once sync.Once
	stor storageImplement

	// fabric = map[string]InitFuncs{}
	fabric = map[string]Init{}
)

func Get(cfg config.Config) (storageImplement, error) {
	var err error
	once.Do(func() {
		strCfg := cfg.Storage()

		//var f *InitFuncs
		// f, err = getStorageByType(strCfg.Type())
		stor, err = getStorageByType(strCfg.Type())
		if err != nil {
			return
		}

		//storage, err = f.Connect(cfg)
		err = stor.Connect(cfg)
		if err != nil {
			return
		}

		//if cfg.Storage().NeedMigrate() && f.Migrate != nil {
		//	err = f.Migrate(cfg)
		//}
		if cfg.Storage().NeedMigrate() {
			err = stor.Migrate(cfg)
		}
	})
	if err != nil {
		return nil, err
	}

	return stor, err
}

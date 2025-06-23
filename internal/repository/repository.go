package repository

import (
	"github.com/eragon-mdi/ksu/pkg/storage"
)

type Repositorier interface {
	Tasker
}

type repository struct {
	storage *storage.Type
}

func New(fake *storage.Type) Repositorier {
	return repository{
		storage: fake,
	}
}

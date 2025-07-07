package fake

import (
	"errors"
	"sync"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

// alias заглушка
type data = entity.Task

// замена БД,
// по тз храню в "памяти сервиса", напомнило принцип fake
type StorageType struct {
	data sync.Map
}

var storage *StorageType

type CRUDer interface {
	Creater
	Reader
	Updater
	Deleter
}

func New() (*StorageType, error) {
	if storage != nil {
		return nil, errors.New("storage Init() call second time")
	}

	return &StorageType{}, nil
}

var (
	ErrKeyIsReserved = errors.New("key is used")
	ErrBadKey        = errors.New("no content by key")
	ErrNoData        = errors.New("no data by key")
	ErrBadData       = errors.New("bad data type")
)

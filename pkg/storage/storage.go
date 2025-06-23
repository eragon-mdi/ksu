package storage

import (
	"sync"

	"github.com/eragon-mdi/ksu/pkg/storage/fake"
)

// Используется для изоляции репозитория от самой бд/fake
type Type = fake.StorageType // для примера потом поменять на sql.DB или на мок

// синглтон
var (
	once    sync.Once
	storage *Type
)

func Get() (*Type, error) {
	var err error
	once.Do(func() {
		storage, err = fake.New()
	})

	return storage, err
}

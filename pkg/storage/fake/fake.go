package fake

import "errors"

// замена БД,
// по тз храню в "памяти сервиса", напомнило принцип fake
type StorageType map[string]data

var storage StorageType

// изолировать данные и не обращаться к ним на прямую
type CRUDer interface {
	Creater
	Reader
	Updater
	Deleter
}

// мапа ссылочный тип, но возвращаю указатель для привычного объявления в storage
func Init() (*StorageType, error) {
	if storage != nil {
		return &StorageType{}, errors.New("storage Init() call second time")
	}

	storage = make(StorageType, 8*8*8)

	return &storage, nil
}

type data struct {
	//
	//
	//
}

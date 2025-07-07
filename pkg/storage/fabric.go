package storage

import (
	"errors"
	"log/slog"
)

//type InitFuncs struct {
//	Connect func(config.Config) (storageImplement, error)
//	Migrate func(config.Config) error
//}

// func getStorageByType(typ string) (*InitFuncs, error) {
func getStorageByType(typ string) (Init, error) {
	f, ok := fabric[typ]
	if !ok {
		return nil, errors.New("unknown storage type: " + typ)
	}

	return f, nil
	// return &f, nil
}

// any must by InitFuncs type
// func register(typ string, f *InitFuncs) {
func register(typ string, f Init) {
	l := slog.Default().With("type", typ)

	_, ok := fabric[typ]
	if ok {
		l.Warn("dublicate storage type")
		return
	}

	//v, ok := f.(InitFuncs)
	//if !ok {
	//	l.Warn("invalid New() storage function")
	//	return
	//}

	//if f.Connect == nil || f.Migrate == nil {
	//	l.Warn("Connect or Migrate funcs nil")
	//	return
	//}

	//fabric[typ] = *f
	fabric[typ] = f
}

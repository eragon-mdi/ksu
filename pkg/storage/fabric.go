package storage

import (
	"errors"
	"log/slog"
)

func getStorageByType(typ string) (Init, error) {
	f, ok := fabric[typ]
	if !ok {
		return nil, errors.New("unknown storage type: " + typ)
	}

	return f, nil
}

func register(typ string, f Init) {
	l := slog.Default().With("type", typ)

	_, ok := fabric[typ]
	if ok {
		l.Warn("dublicate storage type")
		return
	}

	fabric[typ] = f
}

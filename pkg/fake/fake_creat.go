package fake

import "log/slog"

type Creater interface {
	InsertTaskWithReturns(key string, data data) (data, error)
}

func (s *StorageType) InsertTaskWithReturns(key string, vals data) (data, error) {
	value, loaded := s.data.LoadOrStore(key, vals)
	if loaded {
		slog.Error(ErrKeyIsReserved.Error(), slog.Any("key", key))
		return data{}, ErrKeyIsReserved
	}

	stored, ok := value.(data)
	if !ok {
		slog.Error(ErrBadData.Error(), slog.Any("data", value))
		return data{}, ErrBadData
	}

	slog.Debug("create new elem", slog.Any("elem", stored))

	return stored, nil
}

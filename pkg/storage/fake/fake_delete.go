package fake

import "log/slog"

type Deleter interface {
	DeleteTaskById(string) error
}

func (s *StorageType) DeleteTaskById(key string) error {
	deleted, ok := s.data.LoadAndDelete(key)
	if !ok {
		slog.Error(ErrBadKey.Error(), slog.Any("key", key))
		return ErrBadKey
	}

	slog.Debug("deleted elem", slog.Any("elem", deleted))

	return nil
}

package fake

import "log/slog"

type Updater interface {
	UpdateTask(data) error
}

func (s *StorageType) UpdateTask(d data) error {
	key := d.ID

	updated, loaded := s.data.Swap(key, d)
	if !loaded {
		return ErrBadKey
	}

	slog.Debug("update elem", slog.Any("elem", updated))

	return nil
}

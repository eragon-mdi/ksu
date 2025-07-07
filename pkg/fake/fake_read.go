package fake

import "log/slog"

type Reader interface {
	SelectAllInfoById(string) (data, error)

	SelectAll() []data
}

func (s *StorageType) SelectAllInfoById(key string) (data, error) {
	value, ok := s.data.Load(key)
	if !ok {
		slog.Error(ErrNoData.Error(), slog.Any("key", key))
		return data{}, ErrNoData
	}

	selected, ok := value.(data)
	if !ok {
		slog.Error(ErrBadData.Error(), slog.Any("data", value))
		return data{}, ErrBadData
	}

	slog.Debug("get elem by id", slog.Any("elem", selected))

	return selected, nil
}

func (s *StorageType) SelectAll() []data {
	dArr := make([]data, 0, 10)

	s.data.Range(func(key, value any) bool {
		d, ok := value.(data)
		if !ok {
			slog.Error(ErrBadData.Error(), slog.Any("data", value))
			return false
		}

		dArr = append(dArr, d)
		return true
	})

	return dArr
}

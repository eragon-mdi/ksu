package mapwithmutex

import "sync"

type MaperWithMutex[T any] interface {
	Set(string, T)
	Get(string) (T, bool)
	Delete(string)
}

type mapWithMutex[T any] struct {
	maap map[string]T

	mutex sync.RWMutex
}

func New[T any](countMapElems int) MaperWithMutex[T] {
	return &mapWithMutex[T]{
		maap: make(map[string]T, countMapElems),
	}
}

func (m *mapWithMutex[T]) Set(key string, val T) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.maap[key] = val
}
func (m *mapWithMutex[T]) Get(key string) (T, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := m.maap[key]

	return v, ok
}
func (m *mapWithMutex[T]) Delete(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.maap, key)
}

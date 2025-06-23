package semaphor

import "context"

type Semaphorer interface {
	AcquireCtx(context.Context) error
	Release() // может запаниковать!!!
}

type semaphor struct {
	tickets chan struct{}
}

func New(capacity int) Semaphorer {
	return &semaphor{
		tickets: make(chan struct{}, capacity),
	}
}

func (s *semaphor) AcquireCtx(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.tickets <- struct{}{}:
		return nil
	}
}
func (s *semaphor) Release() {
	select {
	case <-s.tickets:
	// защита от неправильного использования и потенциальной блокировки
	default:
		panic("semaphore: release empty tickets query")
	}
}

package batch

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type BatchQueue[T any] interface {
	Push(context.Context, T)
	GoBatchByTimer(context.Context)

	clearAndDo(ctx context.Context)
}

type batchQueue[T any] struct {
	sizeQ          int
	queue          []T
	ticker         *time.Ticker
	tickerDuration time.Duration

	m      sync.Mutex
	doFunc func(context.Context, []T)
}

func New[T any](queueCap int, tInterval time.Duration, f func(context.Context, []T)) BatchQueue[T] {
	return &batchQueue[T]{
		sizeQ:          queueCap,
		queue:          make([]T, 0, queueCap),
		ticker:         time.NewTicker(tInterval),
		tickerDuration: tInterval,

		m:      sync.Mutex{},
		doFunc: f,
	}
}

// опционально вызвать в горутине, так как есть блокировка в мьютексе
func (q *batchQueue[T]) Push(ctx context.Context, l T) {
	q.m.Lock()
	defer q.m.Unlock()

	q.queue = append(q.queue, l)

	// батчинг по кол-ву
	if len(q.queue) >= q.sizeQ { // cap(q.queue) { /////////////////////////////////////
		q.clearAndDo(ctx)

		q.ticker.Reset(q.tickerDuration)
	}
}

func (q *batchQueue[T]) GoBatchByTimer(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-q.ticker.C:
				q.m.Lock()
				q.clearAndDo(ctx)
				q.m.Unlock()
			}
		}
	}()
}

// вызывает функцию над очередью и обнуляет очередь
// вызывать только в мьютексе
func (q *batchQueue[T]) clearAndDo(ctx context.Context) {
	if len(q.queue) == 0 {
		return
	}

	temp := make([]T, len(q.queue))
	copy(temp, q.queue)
	q.queue = q.queue[:0]

	go q.doFunc(ctx, temp)

	fmt.Println("do called", len(temp), time.Now())
}

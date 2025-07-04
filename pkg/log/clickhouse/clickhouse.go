package clickhouse

import (
	"context"
	"io"
	"log/slog"
	"sync"

	clickH "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/eragon-mdi/ksu/pkg/batch"
)

type Clickhouse interface {
	io.Writer
	GracefulShutdown()
}

type clickhouse struct {
	storCon clickH.Conn
	queue   batch.BatchQueue[[]byte]

	syncDbReady chan struct{}
}

var (
	once sync.Once
	ch   clickhouse
)

// реализация интерфейса write через батчинг
func (ch *clickhouse) Write(p []byte) (int, error) {

	// передвать только копию, иначе данные режутся
	// владелец p после завершения Write может перезаписать p
	copyP := make([]byte, len(p))
	copy(copyP, p)
	go ch.queue.Push(context.TODO(), copyP)

	return len(p), nil
}

func (ch *clickhouse) GracefulShutdown() {
	if err := ch.storCon.Close(); err != nil {
		slog.Default().Error("error ocured on drop connection with clickhouse", slog.Any("cause", err))
	}
}

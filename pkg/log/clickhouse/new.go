package clickhouse

import (
	"context"
	"log"
	"time"

	clickH "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/eragon-mdi/ksu/pkg/batch"
	"github.com/eragon-mdi/ksu/pkg/config"
)

func New(cfg config.Config) Clickhouse {
	ctx := context.Background()
	chCfg := cfg.ClickHouse()

	// var err error
	once.Do(func() {
		ch.queue = batch.New(chCfg.ButchSize(), chCfg.ButchInteval(), ch.insertManyLogs)
		ch.syncDbReady = make(chan struct{})

		go func() {
			defer close(ch.syncDbReady)

			var err error

			for range chCfg.ConnAttempts() {
				ch.storCon, err = oneTryConnectCH(cfg)

				if err != nil {
					continue
				}

				time.Sleep(chCfg.TryConnPeriod())
			}

			if err != nil {
				log.Fatal(err)
			}

			if err = ch.storCon.Ping(ctx); err != nil {
				log.Fatal(err)
			}

			if err = ch.ensureLogsTable(ctx); err != nil {
				log.Fatal(err)
			}

			ch.queue.GoBatchByTimer(ctx) // горутина
		}()
	})

	return &ch
}

func oneTryConnectCH(cfg config.Config) (clickH.Conn, error) {
	chCfg := cfg.ClickHouse()

	return clickH.Open(&clickH.Options{
		Addr: []string{chCfg.Addr()},
		Auth: clickH.Auth{
			Database: chCfg.Db(),
			Username: chCfg.User(),
			Password: chCfg.Pass(),
		},
	})
}

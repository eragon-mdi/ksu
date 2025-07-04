package clickhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

const query = `INSERT INTO logs_raw FORMAT JSONEachRow`

func (ch *clickhouse) insertManyLogs(ctx context.Context, batchLogs [][]byte) {
	<-ch.syncDbReady

	batch, err := ch.storCon.PrepareBatch(ctx, query)
	if err != nil {
		fmt.Println("batch init fail:", err)
		return
	}

	for _, raw := range batchLogs {
		if !json.Valid(raw) {
			fmt.Println("invalid JSON skipped:", string(raw))
			continue
		}

		// оборачиваем: {"raw": <raw>}
		wrapped, err := json.Marshal(map[string]json.RawMessage{
			"raw": raw,
		})
		if err != nil {
			fmt.Println("marshal failed:", err)
			continue
		}

		if err := batch.Append(wrapped); err != nil {
			fmt.Println("append failed:", err)
			return
		}
	}

	if err := batch.Send(); err != nil {
		fmt.Println("send failed:", err)
	}
}

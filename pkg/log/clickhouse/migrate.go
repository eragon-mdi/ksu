package clickhouse

import "context"

const makeLogsTable = `
	CREATE TABLE logs_raw (
		raw       JSON
		)
	ENGINE = MergeTree()
	ORDER BY tuple();
`

func (ch *clickhouse) ensureLogsTable(ctx context.Context) error {
	return ch.storCon.Exec(ctx, makeLogsTable)
}

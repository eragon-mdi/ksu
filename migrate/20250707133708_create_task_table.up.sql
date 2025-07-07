CREATE TABLE IF NOT EXISTS task (
    id TEXT PRIMARY KEY,
    result JSONB,
    status INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    duration BIGINT NOT NULL,         -- в наносекундах (time.Duration = int64)
    started_at TIMESTAMPTZ NOT NULL
);

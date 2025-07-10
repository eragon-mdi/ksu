CREATE TYPE task_status AS ENUM ('pending', 'running', 'completed', 'failed');

CREATE TABLE IF NOT EXISTS task (
    id TEXT PRIMARY KEY,
    result TEXT NOT NULL,
    status task_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    duration BIGINT NOT NULL,           -- в наносекундах (time.Duration = int64)
    started_at TIMESTAMPTZ NOT NULL
);
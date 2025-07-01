package taskstate

import (
	"context"
	"errors"
	"log/slog"
	"time"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
	applog "github.com/eragon-mdi/ksu/pkg/log"
)

type TaskUtils interface {
	Result(entity.Task) (entity.Task, error)
	Duration(entity.Task) entity.Task
}

var statusMapResult = map[int]any{
	entity.STATUS_PENDING: "no result, task pending",
	entity.STATUS_RUNNING: "no result, task runnibg",
	entity.STATUS_FAILED:  "no result, task failed",
}

func (taskState) Result(t entity.Task) (entity.Task, error) {
	if s, ok := statusMapResult[t.Status]; ok {
		t.Result = s
		return t, errors.New("task service: bad status for result")
	}
	return t, nil
}

func (taskState) Duration(t entity.Task) entity.Task {
	if t.Status == entity.STATUS_RUNNING {
		t.Duration = time.Duration(time.Since(t.StartedAt).Seconds())
	}
	return t
}

func (ts taskState) saveToRepositiry(ctx context.Context, t entity.Task) entity.Task {
	l := applog.GetCtxLogger(ctx)

	if err := ts.repository.UpdateTaskInfo(t); err != nil {
		l.Warn("taskState: err updated task status", slog.Any("cause", err))
	}
	return t
}

package taskstate

import (
	"context"

	entity "github.com/eragon-mdi/ksu/internal/entity/task"
)

func (ts taskState) Advanced(ctx context.Context, t entity.Task) entity.Task {
	next, ok := ts.transitions[t.Status]
	if ok {
		t.Status = next
	}

	return ts.saveToRepositiry(ctx, t)
}

func (ts taskState) Failed(ctx context.Context, t entity.Task) entity.Task {
	if t.Status != entity.STATUS_COMPLETED {
		t.Status = entity.STATUS_FAILED
	}

	return ts.saveToRepositiry(ctx, t)
}

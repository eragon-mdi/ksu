package entity

import "time"

func (t Task) SetResult(r ResultType) Task {
	t.Result = r
	return t
}

func (t *Task) SetStartedAtTime() {
	t.StartedAt = time.Now()
}

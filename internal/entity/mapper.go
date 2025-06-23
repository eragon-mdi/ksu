// mapping
package entity

import "time"

var statusMap = map[int]string{
	STATUS_COMPLETED: "completed",
	STATUS_FAILED:    "failed by io bound",
	STATUS_PENDING:   "created, but waiting to start",
	STATUS_RUNNING:   "in process",
}

func (t Task) ResponseCreate() TaskCreateResponse {
	return TaskCreateResponse{
		ID:         t.ID,
		TaskStatus: t.TaskStatus,
	}
}

func (t Task) Response() TaskResponse {
	status, ok := statusMap[t.Status]
	if !ok {
		status = "unknown"
	}

	return TaskResponse{
		ID:        t.ID,
		Result:    t.Result,
		Status:    status,
		CreatedAt: t.CreatedAt,
		Duration:  t.updateDuration(),
	}
}

func (t TaskStatus) Response() TaskStatusResponse {
	status, ok := statusMap[t.Status]
	if !ok {
		status = "unknown"
	}

	return TaskStatusResponse{
		StatusString: status,
		TaskStatus:   t,
	}
}

func (t Task) Update(status int, result ...ResultType) Task {
	res := t.Result
	if len(result) == 1 {
		res = result[0]
	}

	//t.Duration = updateDuration(status, t.StartedAt, t.Status)
	//t.Status = status

	return Task{
		ID:        t.ID,
		StartedAt: t.StartedAt,
		TaskResult: TaskResult{
			Result: res,
		},
		TaskStatus: TaskStatus{
			Status:    status,
			CreatedAt: t.CreatedAt,
			Duration:  t.updateDuration(), //updateDuration(status, t.StartedAt, t.Duration),
		},
	}
}

func (t Task) EntityTaskStatus() TaskStatus {
	return TaskStatus{
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
		Duration:  t.updateDuration(),
	}
}

// func updateDuration(status int, startedAt time.Time, oldStatus ...int) time.Duration { //status int, startedAt time.Time, tD time.Duration
func (t Task) updateDuration() time.Duration {
	//if len(oldStatus) == 1 {
	//	if oldStatus[0] == STATUS_COMPLETED || oldStatus[0] == STATUS_FAILED {
	//		return updateDuration(status, startedAt)
	//	}
	//}

	if t.Status == STATUS_RUNNING {
		return time.Duration(time.Since(t.StartedAt).Seconds())
	}

	return t.Duration
}

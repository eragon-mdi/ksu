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

	return Task{
		ID:        t.ID,
		StartedAt: t.StartedAt,
		TaskResult: TaskResult{
			Result: res,
		},
		TaskStatus: TaskStatus{
			Status:    status,
			CreatedAt: t.CreatedAt,
			Duration:  t.updateDuration(),
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

func (t Task) updateDuration() time.Duration {

	if t.Status == STATUS_RUNNING {
		return time.Duration(time.Since(t.StartedAt).Seconds())
	}

	return t.Duration
}

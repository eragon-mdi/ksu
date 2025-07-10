package entity

var statusMap = map[TaskStatusType]string{
	STATUS_COMPLETED: "completed",
	STATUS_FAILED:    "failed by io bound",
	STATUS_PENDING:   "created, but waiting to start",
	STATUS_RUNNING:   "in process",
}

func (t Task) strStatus() string {
	status, ok := statusMap[t.Status]
	if !ok {
		status = "unknown"
	}

	return status
}

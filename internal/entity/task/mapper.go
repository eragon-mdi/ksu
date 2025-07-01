package entity

func (t Task) CreateResponse() TaskCreateResponse {
	return TaskCreateResponse{
		ID:         t.ID,
		TaskStatus: t.TaskStatus,
	}
}

func (t Task) ResultResponse() TaskResultResponse {
	return TaskResultResponse{
		Result: t.Result,
	}
}

func (t Task) StatusResponse() TaskStatusResponse {
	return TaskStatusResponse{
		StatusString: t.strStatus(),
		TaskStatus:   t.TaskStatus,
	}
}

func (t Task) Response() TaskResponse {
	return TaskResponse{
		ID:        t.ID,
		Result:    t.Result,
		Status:    t.strStatus(),
		CreatedAt: t.CreatedAt,
		Duration:  t.Duration,
	}
}

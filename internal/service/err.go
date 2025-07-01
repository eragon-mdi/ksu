package service

import "fmt"

const prefix = "service: "

var (
	ErrBySave        = fmt.Errorf("%sfailed to save task", prefix)
	ErrByDelete      = fmt.Errorf("%sfailed to delete task", prefix)
	ErrGetTaskResult = fmt.Errorf("%sfailed to get task result", prefix)
	ErrGetTaskStatus = fmt.Errorf("%sfailed to get task status", prefix)
	ErrGetAllTask    = fmt.Errorf("%sfailed to get all task statuses", prefix)
	ErrGetTasks      = fmt.Errorf("%sservice: no any tasks", prefix)
)

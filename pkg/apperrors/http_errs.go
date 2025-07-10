package apperrors

var (
	NotFound           = newAppErr("not found")
	ErrInternal        = newAppErr("internal server err")
	ErrInvalidID       = newAppErr("invalid id")
	ErrInvalidData     = newAppErr("invalid data")
	ErrTaskNotFound    = newAppErr("task not found")
	ErrInvalidTaskData = newAppErr("invalid task data")
	ErrCantCreateTask  = newAppErr("err staring task")
	ErrCantDeleteTask  = newAppErr("err droping task")
	ErrTaskNoComplete  = newAppErr("task running or failed, no result, check status")
)

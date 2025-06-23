package apperrors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppErr struct {
	message string
}

func newAppErr(m string) *AppErr {
	return &AppErr{message: m}
}

func (a AppErr) Error() string {
	return a.message
}

// Кастоиная реализация обработчика ошибок
// для простого написания хендлеров
//
// основан на
// // HTTPErrorHandler is a centralized HTTP error handler.
// // type HTTPErrorHandler func(err error, c Context)
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code    int = http.StatusInternalServerError
		message any = ErrInternal.Error()
	)

	if echoError, ok := err.(*echo.HTTPError); ok {
		code = echoError.Code
		message = echoError.Message

		if appErr, ok := echoError.Message.(*AppErr); ok {
			message = appErr.Error()
		}
	}

	if !c.Response().Committed {
		c.JSON(code, echo.Map{"error": message})
	}
}

func HandlePanic() {
	if r := recover(); r != nil {
		// TODO log
	}
}

var (
	ErrInternal        = newAppErr("internal server err")
	ErrInvalidID       = newAppErr("invalid id")
	ErrInvalidData     = newAppErr("invalid data")
	ErrTaskNotFound    = newAppErr("task not found")
	ErrInvalidTaskData = newAppErr("invalid task data")
	ErrCantCreateTask  = newAppErr("err staring task")
	ErrCantDeleteTask  = newAppErr("err droping task")
	ErrTaskNoComplete  = newAppErr("task running or failed, no result, check status")
)

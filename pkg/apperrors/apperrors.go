package apperrors

import (
	"log/slog"
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

func HandlePanic(l *slog.Logger) {
	if r := recover(); r != nil {
		l.Warn("was panic", slog.Any("cause", r))
	}
}

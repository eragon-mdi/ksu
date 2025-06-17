package apperrors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Кастоиная реализация обработчика ошибок
// для простого написания хендлеров
//
// основан на 
// // HTTPErrorHandler is a centralized HTTP error handler.
// // type HTTPErrorHandler func(err error, c Context)
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code    int = http.StatusInternalServerError
		message any = "internal server error"
	)

	echoError, ok := err.(*echo.HTTPError)
	if ok {
		code = echoError.Code
		message = echoError.Message
	}

	if !c.Response().Committed {
		c.JSON(code, echo.Map{"error": message})
	}
}

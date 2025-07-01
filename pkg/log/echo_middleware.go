package applog

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// кастомный логгер
// MiddlewareFunc defines a function to process middleware.
// type MiddlewareFunc func(next HandlerFunc) HandlerFunc
func InitMiddlewareLogging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			loggerChild := slog.Default().With(
				slog.String("trace", uuid.NewString()),
				slog.String("method", c.Request().Method),
				slog.String("uri", c.Request().URL.Path),
			)

			// логгер в context
			ctx := context.WithValue(c.Request().Context(), ctxLoggerKey{}, loggerChild)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

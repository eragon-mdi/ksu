package applog

import (
	"log/slog"
	"os"

	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const LOGLEVEL = slog.LevelWarn // slog.LevelDebug

type ctxLoggerKey struct{}

func GetCtxLogger(c echo.Context) *slog.Logger {
	logger, ok := c.Request().Context().Value(ctxLoggerKey{}).(*slog.Logger)
	if !ok {
		slog.Error("can't get custom logger, set default")
		return slog.Default()
	}
	return logger
}

// кастомный логгер
// MiddlewareFunc defines a function to process middleware.
// type MiddlewareFunc func(next HandlerFunc) HandlerFunc
func InitMiddlewareLogging() echo.MiddlewareFunc {
	baseLogger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     LOGLEVEL,
			},
		),
	)
	slog.SetDefault(baseLogger)

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

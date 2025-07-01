package applog

import (
	"log/slog"

	"context"

	"github.com/labstack/echo/v4"
)

type ctxLoggerKey struct{}

func CtxWithLogger(l *slog.Logger) context.Context {
	return context.WithValue(context.Background(), ctxLoggerKey{}, l)
}

func GetCtxLogger(c context.Context) *slog.Logger {
	logger, ok := c.Value(ctxLoggerKey{}).(*slog.Logger)
	if !ok {
		slog.Error("can't get custom logger, set default")
		return slog.Default()
	}
	return logger
}

func GetRequestCtxLogger(c echo.Context) *slog.Logger {
	logger, ok := c.Request().Context().Value(ctxLoggerKey{}).(*slog.Logger)
	if !ok {
		slog.Error("can't get custom logger, set default")
		return slog.Default()
	}
	return logger
}

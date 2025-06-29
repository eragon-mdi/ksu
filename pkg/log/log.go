package applog

import (
	"log"
	"log/slog"
	"os"

	"context"

	"github.com/eragon-mdi/ksu/pkg/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	LoggerHandlerType_JSON = "json"
	LoggerHandlerType_TEXT = "text"
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

func SetDefaultBaseLogger(cfg config.Config) {
	logCfg := cfg.Logger()

	hOpt := slog.HandlerOptions{
		AddSource: true,
		Level:     logCfg.Level(),
	}

	var file *os.File
	if logCfg.Output() == "internal" {
		file = os.Stdout
	} // TODO add other cases

	switch logCfg.Handler() {
	case LoggerHandlerType_JSON:
		slog.SetDefault(slog.New(
			slog.NewJSONHandler(
				file,
				&hOpt,
			),
		))
	case LoggerHandlerType_TEXT:
		slog.SetDefault(slog.New(
			slog.NewTextHandler(
				file,
				&hOpt,
			),
		))
	default:
		log.Fatal("can't init logger")
	}
}

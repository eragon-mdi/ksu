package applog

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/eragon-mdi/ksu/pkg/config"
)

const (
	LoggerHandlerType_JSON = "json"
	LoggerHandlerType_TEXT = "text"
)

var loggerTypes = map[string]func(w io.Writer, ho slog.HandlerOptions) *slog.Logger{
	LoggerHandlerType_JSON: func(w io.Writer, ho slog.HandlerOptions) *slog.Logger {
		return slog.New(slog.NewJSONHandler(w, &ho))
	},
	LoggerHandlerType_TEXT: func(w io.Writer, ho slog.HandlerOptions) *slog.Logger {
		return slog.New(slog.NewTextHandler(w, &ho))
	},
}

func SetDefaultBaseLogger(cfg config.Config, w ...io.Writer) {
	logCfg := cfg.Logger()

	hOpt := slog.HandlerOptions{
		AddSource: true,
		Level:     logCfg.Level(),
	}

	var file io.Writer = os.Stdout
	if len(w) == 1 {
		file = w[0]
	}

	getLogger, ok := loggerTypes[logCfg.Handler()]
	if !ok {
		log.Fatal("can't init logger")
	}

	slog.SetDefault(getLogger(file, hOpt))
}

package applog

import (
	"log"
	"log/slog"
	"os"

	"github.com/eragon-mdi/ksu/pkg/config"
)

const (
	LoggerHandlerType_JSON = "json"
	LoggerHandlerType_TEXT = "text"
)

var loggerTypes = map[string]func(f *os.File, ho slog.HandlerOptions) *slog.Logger{
	LoggerHandlerType_JSON: func(f *os.File, ho slog.HandlerOptions) *slog.Logger {
		return slog.New(slog.NewJSONHandler(f, &ho))
	},
	LoggerHandlerType_TEXT: func(f *os.File, ho slog.HandlerOptions) *slog.Logger {
		return slog.New(slog.NewTextHandler(f, &ho))
	},
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
	} // TODO add other cases // with map

	getLogger, ok := loggerTypes[logCfg.Handler()]
	if !ok {
		log.Fatal("can't init logger")
	}

	slog.SetDefault(getLogger(file, hOpt))
}

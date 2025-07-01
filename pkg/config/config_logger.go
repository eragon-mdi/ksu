package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type loggerConfig interface {
	Handler() string
	Level() slog.Level
	Output() string
}

const loggerTag = "logger"

type logger struct {
	HandlerType string `mapstructure:"handler"`
	Loglevel    string `mapstructure:"level"`
	OutputDir   string `mapstructure:"output"`
}

func loggerSetDefaultRoutes(v *viper.Viper) {
	v.SetDefault(loggerTag+".handler", "text")
	v.SetDefault(loggerTag+".level", "error")
	v.SetDefault(loggerTag+".output", "internal")
}

func (l logger) Handler() string {
	return l.HandlerType
}

var logLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func (l *logger) Level() slog.Level {
	level, ok := logLevels[l.Loglevel]
	if !ok {
		l.Loglevel = "debug"
		return slog.LevelDebug
	}

	return level
}

func (l logger) Output() string {
	return l.OutputDir
}

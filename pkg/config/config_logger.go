package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

func init() {
	configSetters = append(configSetters, loggerSetDefault)
}

func (c config) Logger() loggerConfig {
	return &c.LoggerCfg
}

type loggerConfig interface {
	Handler() string
	Level() slog.Level
	WriteInternal() bool
}

const loggerTag = "logger"

type logger struct {
	HandlerType string `mapstructure:"handler"`
	Loglevel    string `mapstructure:"level"`
	OutputDir   bool   `mapstructure:"output_internal"`
}

func loggerSetDefault(v *viper.Viper) {
	v.SetDefault(loggerTag+".handler", "text")
	v.SetDefault(loggerTag+".level", "error")
	v.SetDefault(loggerTag+".output_internal", true)
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

func (l logger) WriteInternal() bool {
	return l.OutputDir
}

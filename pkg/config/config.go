package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type Config interface {
	// group cfgs by feature domains
	Server() serverConfig
	Logger() loggerConfig
}

type config struct {
	ServerCfg server `mapstructure:"server"`
	LoggerCfg logger `mapstructure:"logger"`
	// add other config entities
}

func Init() (Config, error) {
	log := slog.Default()

	v := viper.New()
	setDefault(v)

	v.SetConfigFile("/config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Warn("can't read config yaml", slog.Any("cause", err))
	}

	var cfg config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Debug("set configs:", slog.Any("conf", cfg))

	return &cfg, nil
}

func (c config) Server() serverConfig {
	return &c.ServerCfg
}

func (c config) Logger() loggerConfig {
	return &c.LoggerCfg
}

func setDefault(v *viper.Viper) {
	serverSetDefaultRoutes(v)
	loggerSetDefaultRoutes(v)
}

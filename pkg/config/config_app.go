package config

import "github.com/spf13/viper"

type app struct {
	MaxSemaphore int `mapstructure:"semaphore"`
}

type appConfig interface {
	Semaphore() int
}

const appTag = "app"

func appSetDefaultRoutes(v *viper.Viper) {
	v.SetDefault(appTag+".semaphore", 8)
}

func (a app) Semaphore() int {
	return a.MaxSemaphore
}

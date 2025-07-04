package config

import "github.com/spf13/viper"

func init() {
	configSetters = append(configSetters, appSetDefault)
}

func (c config) App() appConfig {
	return &c.AppCfg
}

type app struct {
	MaxSemaphore int `mapstructure:"semaphore"`
}

type appConfig interface {
	Semaphore() int
}

const appTag = "app"

func appSetDefault(v *viper.Viper) {
	v.SetDefault(appTag+".semaphore", 8)
}

func (a app) Semaphore() int {
	return a.MaxSemaphore
}

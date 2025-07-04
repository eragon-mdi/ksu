package config

import "github.com/spf13/viper"

var configSetters []func(v *viper.Viper)

func setDefault(v *viper.Viper) {
	for _, set := range configSetters {
		set(v)
	}
}

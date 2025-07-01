package config

import (
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const serverTag = "server"

type server struct {
	Address      string        `mapstructure:"addr"`
	PortStr      int           `mapstructure:"port"`
	ReadTt       time.Duration `mapstructure:"read_timeout"`
	WriteTt      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTt time.Duration `mapstructure:"read_header_timeout"`
	IdleTt       time.Duration `mapstructure:"idle_timeout"`
}

type serverConfig interface {
	Addr() string
	Port() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	ReadHeaderTimeout() time.Duration
	IdleTimeout() time.Duration
}

func serverSetDefaultRoutes(v *viper.Viper) {
	v.SetDefault(serverTag+".addr", "0.0.0.0")
	v.SetDefault(serverTag+".port", 8080)
	v.SetDefault(serverTag+".read_timeout", "5s")
	v.SetDefault(serverTag+".write_timeout", "10s")
	v.SetDefault(serverTag+".read_header_timeout", "2s")
	v.SetDefault(serverTag+".idle_timeout", "30s")
}

func (s server) Addr() string {
	return s.Address
}
func (s server) Port() string {
	return strconv.Itoa(s.PortStr)
}
func (s server) ReadTimeout() time.Duration {
	return s.ReadTt
}
func (s server) WriteTimeout() time.Duration {
	return s.WriteTt
}
func (s server) ReadHeaderTimeout() time.Duration {
	return s.ReadHeaderTt
}
func (s server) IdleTimeout() time.Duration {
	return s.IdleTt
}

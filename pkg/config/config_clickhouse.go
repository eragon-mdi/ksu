package config

import (
	"time"

	"github.com/spf13/viper"
)

func init() {
	configSetters = append(configSetters, clickHouseSetDefault)
}

func (c config) ClickHouse() chConfig {
	return &c.ClickHouseCfg
}

const clickHouseTag = "clickhouse"

type chConfig interface {
	Addr() string
	Db() string
	User() string
	Pass() string
	ButchSize() int
	ButchInteval() time.Duration

	TryConnPeriod() time.Duration
	ConnAttempts() int
}

type clickhouse struct {
	Address           string        `mapstructure:"address"`
	Database          string        `mapstructure:"database"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	ButchCapacity     int           `mapstructure:"batch_size"`
	ButchTimerInteval time.Duration `mapstructure:"batch_inteval"`

	TryConncetionPeriod time.Duration `mapstructure:"try_connection_perod"`
	ConnectionAttempts  int           `mapstructure:"connection_attempts"`
}

func clickHouseSetDefault(v *viper.Viper) {
	v.SetDefault(clickHouseTag+".address", "0.0.0.0:9000")
	v.SetDefault(clickHouseTag+".database", "default")
	v.SetDefault(clickHouseTag+".username", "user")
	v.SetDefault(clickHouseTag+".password", "password")
	v.SetDefault(clickHouseTag+".batch_size", 100)
	v.SetDefault(clickHouseTag+".batch_inteval", 2*time.Minute)

	v.SetDefault(clickHouseTag+".try_connection_perod", 3*time.Second)
	v.SetDefault(clickHouseTag+".connection_attempts", 5)
}

func (ch clickhouse) Addr() string {
	return ch.Address
}
func (ch clickhouse) Db() string {
	return ch.Database
}
func (ch clickhouse) User() string {
	return ch.Username
}
func (ch clickhouse) Pass() string {
	return ch.Password
}
func (ch clickhouse) ButchSize() int {
	return ch.ButchCapacity
}
func (ch clickhouse) ButchInteval() time.Duration {
	return ch.ButchTimerInteval
}

func (ch clickhouse) TryConnPeriod() time.Duration {
	return ch.TryConncetionPeriod
}
func (ch clickhouse) ConnAttempts() int {
	return ch.ConnectionAttempts
}

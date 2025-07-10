package config

import "github.com/spf13/viper"

func init() {
	configSetters = append(configSetters, storageSetDefault)
}

func (c config) Storage() stroageCfg {
	return &c.StorageCfg
}

const storageTag = "storage"

type stroageCfg interface {
	Type() string
	Host() string
	Port() string
	User() string
	Password() string
	DBname() string
	SSLmode() string
	NeedMigrate() bool
	MigrateSrc() string
}

type storage struct {
	TypeOfStr     string `mapstructure:"type"`
	HostDB        string `mapstructure:"host"`
	PortDB        string `mapstructure:"port"`
	UserDB        string `mapstructure:"user"`
	PasswordDB    string `mapstructure:"password"`
	Dbname        string `mapstructure:"name"`
	SSLmodeDB     string `mapstructure:"ssl_mode"`
	NeedMigrateDB bool   `mapstructure:"need_migrate"`
	MigrateSrcDB  string `mapstructure:"migaret_src"`
}

func storageSetDefault(v *viper.Viper) {
	v.SetDefault(storageTag+".type", "internal")
	v.SetDefault(storageTag+".host", "0.0.0.0")
	v.SetDefault(storageTag+".port", "5437")
	v.SetDefault(storageTag+".user", "default")
	v.SetDefault(storageTag+".password", "secret")
	v.SetDefault(storageTag+".name", "db")
	v.SetDefault(storageTag+".ssl_mode", "disable")
	v.SetDefault(storageTag+".need_migrate", false)
	v.SetDefault(storageTag+".migaret_src", "/migrate/")
}

func (s storage) Type() string {
	return s.TypeOfStr
}
func (s storage) Host() string {
	return s.HostDB
}
func (s storage) Port() string {
	return s.PortDB
}
func (s storage) User() string {
	return s.UserDB
}
func (s storage) Password() string {
	return s.PasswordDB
}
func (s storage) DBname() string {
	return s.Dbname
}
func (s storage) SSLmode() string {
	return s.SSLmodeDB
}
func (s storage) NeedMigrate() bool {
	return s.NeedMigrateDB
}
func (s storage) MigrateSrc() string {
	return s.MigrateSrcDB
}

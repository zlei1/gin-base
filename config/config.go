package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigFile struct {
	Name string
}

type Config struct {
	App   AppConfig
	Log   LogConfig
	Pgsql PgsqlConfig
	Redis RedisConfig
}

type AppConfig struct {
	Name    string
	RunMode string
	Port    string
	Host    string
	Secret  string
}

type LogConfig struct {
	Writers    string
	File       string
	WarnFile   string
	ErrorFile  string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

type PgsqlConfig struct {
	Host     string
	Database string
	Username string
	Password string
	Port     string
}

type RedisConfig struct {
	Addr         string
	Password     string
	Db           int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolSize     int
}

func Setup(cfgPath string) *Config {
	f := ConfigFile{
		Name: cfgPath,
	}

	config := f.read()
	f.watch()

	return config
}

func (f *ConfigFile) read() *Config {
	if f.Name != "" {
		viper.SetConfigFile(f.Name)
	} else {
		viper.AddConfigPath("config/")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := new(Config)
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}

func (f *ConfigFile) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

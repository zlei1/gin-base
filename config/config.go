package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigFile struct {
	Name string
}

type Config struct {
	Project ProjectConfig `yaml:"project"`
	Log     LogConfig     `yaml:"log"`
	Pgsql   PgsqlConfig   `yaml:"pgsql"`
	Redis   RedisConfig   `yaml:"redis"`
}

type ProjectConfig struct {
	Name    string `yaml:"name"`
	RunMode string `yaml:"run_mode"`
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	Secret  string `yaml:"secret"`
}

type LogConfig struct {
	Writers    string `yaml:"writers"`
	Level      string `yaml:"level"`
	File       string `yaml:"file"`
	WarnFile   string `yaml:"warn_file"`
	ErrorFile  string `yaml:"error_file"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

type PgsqlConfig struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

type RedisConfig struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	Db           int    `yaml:"db"`
	DialTimeout  int    `yaml:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	PoolSize     int    `yaml:"pool_size"`
}

func Perform(cfgPath string) *Config {
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
		log.Fatalf("%s: %v", "Config Read Failed", err)
	}

	config := new(Config)
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("%s: %v", "Config Unmarshal Failed", err)
	}

	return config
}

func (f *ConfigFile) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

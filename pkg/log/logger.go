package log

import "errors"

var log Logger

var (
	invalidLogInstance = errors.New("invalid log instance")
)

const (
	InstanceZapLogger int = iota
)

type Fields map[string]interface{}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(keyValues Fields) Logger
}

type Config struct {
	Writers    string `yaml:"writers"`
	Level      string `yaml:"level"`
	File       string `yaml:"file"`
	WarnFile   string `yaml:"warn_file"`
	ErrorFile  string `yaml:"error_file"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

func New(cfg *Config, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZap(cfg)
		if err != nil {
			return err
		}

		log = logger
		return nil
	default:
		return invalidLogInstance
	}
}

// Debug logger
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logger
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logger
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logger
func Error(args ...interface{}) {
	log.Error(args...)
}

// Fatal logger
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Debugf logger
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf logger
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf logger
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// WithFields
// output more field, eg:
//   contextLogger := log.WithFields(log.Fields{"key1": "value1"})
//   contextLogger.Info("log")
// or more sample to use:
// 	 log.WithFields(log.Fields{"key1": "value1"}).Info("log")
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}

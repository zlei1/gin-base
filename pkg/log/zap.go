package log

import (
	"io"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gin-base/pkg/global"
	"gin-base/pkg/utils"
)

const (
	WriterStdOut = "stdout"
	WriterFile   = "file"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func newZap(cfg *Config) (Logger, error) {
	encoder := getJSONEncoder()

	var cores []zapcore.Core
	var options []zap.Option

	option := zap.Fields(
		zap.String("ip", utils.GetLocalIP()),
		zap.String("app", global.App.Conf.Project.Name),
	)
	options = append(options, option)

	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel
	})

	writers := strings.Split(cfg.Writers, ",")
	for _, w := range writers {
		if w == WriterStdOut {
			core := zapcore.NewCore(
				encoder,
				zapcore.AddSync(os.Stdout),
				zapcore.DebugLevel,
			)
			cores = append(cores, core)
		}

		if w == WriterFile {
			infoWrite := getWrite(
				cfg.File,
				cfg.MaxSize,
				cfg.MaxBackups,
				cfg.MaxAge,
			)
			infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl <= zapcore.InfoLevel
			})
			core := zapcore.NewCore(
				encoder,
				zapcore.AddSync(infoWrite),
				infoLevel,
			)
			cores = append(cores, core)

			warnWrite := getWrite(
				cfg.WarnFile,
				cfg.MaxSize,
				cfg.MaxBackups,
				cfg.MaxAge,
			)
			warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				stacktrace := zap.AddStacktrace(zapcore.WarnLevel)
				options = append(options, stacktrace)
				return lvl == zapcore.WarnLevel
			})
			core = zapcore.NewCore(
				encoder,
				zapcore.AddSync(warnWrite),
				warnLevel,
			)
			cores = append(cores, core)

			errorWrite := getWrite(
				cfg.ErrorFile,
				cfg.MaxSize,
				cfg.MaxBackups,
				cfg.MaxAge,
			)
			errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				stacktrace := zap.AddStacktrace(zapcore.ErrorLevel)
				options = append(options, stacktrace)
				return lvl >= zapcore.ErrorLevel
			})
			core = zapcore.NewCore(
				encoder,
				zapcore.AddSync(errorWrite),
				errorLevel,
			)
			cores = append(cores, core)
		}

		if w != WriterFile && w != WriterStdOut {
			core := zapcore.NewCore(
				encoder,
				zapcore.AddSync(os.Stdout),
				zapcore.DebugLevel,
			)
			cores = append(cores, core)

			writer := getWrite(
				cfg.File,
				cfg.MaxSize,
				cfg.MaxBackups,
				cfg.MaxAge,
			)
			core = zapcore.NewCore(
				encoder,
				zapcore.AddSync(writer),
				allLevel,
			)
			cores = append(cores, core)
		}
	}

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	options = append(options, caller)
	// 开启文件及行号
	development := zap.Development()
	options = append(options, development)
	// 跳过文件调用层数
	addCallerSkip := zap.AddCallerSkip(2)
	options = append(options, addCallerSkip)

	logger := zap.New(combinedCore, options...).Sugar()

	return &zapLogger{sugaredLogger: logger}, nil
}

func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "app",
		CallerKey:      "line",
		StacktraceKey:  "trace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWrite(filename string, maxSize, maxBackups, maxAge int) io.Writer {
	ljl := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}

	return &ljl
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Panicf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

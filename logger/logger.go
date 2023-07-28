package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Debug(msg string)

	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Debugf(format string, args ...any)
}

type DefaultLogger struct {
	log *zap.Logger
}

func NewDefaultLogger() (*DefaultLogger, error) {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	cfg.Level.SetLevel(zapcore.DebugLevel)
	log, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &DefaultLogger{log: log}, nil
}

func (dl *DefaultLogger) Info(msg string) {
	dl.log.Info(msg)
}

func (dl *DefaultLogger) Warn(msg string) {
	dl.log.Warn(msg)
}

func (dl *DefaultLogger) Error(msg string) {
	dl.log.Warn(msg)
}

func (dl *DefaultLogger) Fatal(msg string) {
	dl.log.Error(msg)
}

func (dl *DefaultLogger) Debug(msg string) {
	dl.log.Debug(msg)
}

func (dl *DefaultLogger) Infof(format string, args ...any) {
	dl.log.Info(fmt.Sprintf(format, args...))
}

func (dl *DefaultLogger) Warnf(format string, args ...any) {
	dl.log.Warn(fmt.Sprintf(format, args...))
}

func (dl *DefaultLogger) Errorf(format string, args ...any) {
	dl.log.Error(fmt.Sprintf(format, args...))
}

func (dl *DefaultLogger) Fatalf(format string, args ...any) {
	dl.log.Fatal(fmt.Sprintf(format, args...))
}

func (dl *DefaultLogger) Debugf(format string, args ...any) {
	dl.log.Debug(fmt.Sprintf(format, args...))
}

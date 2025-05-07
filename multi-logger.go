package mlog

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type MultiLogger struct {
	loggers []*slog.Logger
}

func (m MultiLogger) log(level slog.Level, msg interface{}, keysAndValues ...interface{}) {
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg.(string), pcs[0])
	r.Add(keysAndValues...)
	for _, logger := range m.loggers {
		_ = logger.Handler().Handle(context.Background(), r)
	}
}

func (m MultiLogger) Debug(msg interface{}, keysAndValues ...interface{}) {
	m.log(slog.LevelDebug, msg, keysAndValues...)
}

func (m MultiLogger) Info(msg interface{}, keysAndValues ...interface{}) {
	m.log(slog.LevelInfo, msg, keysAndValues...)
}

func (m MultiLogger) Warn(msg interface{}, keysAndValues ...interface{}) {
	m.log(slog.LevelWarn, msg, keysAndValues...)
}

func (m MultiLogger) Error(msg interface{}, keysAndValues ...interface{}) {
	m.log(slog.LevelError, msg, keysAndValues...)
}

func (m MultiLogger) Fatal(msg interface{}, keysAndValues ...interface{}) {
	m.Error(msg, keysAndValues...)
	os.Exit(69420)
}

func (m MultiLogger) SetPrefix(prefix string) {
	for _, logger := range m.loggers {
		logger.WithGroup(prefix)
	}
}

func (m MultiLogger) WithPrefix(prefix string) *MultiLogger {
	var newLoggers []*slog.Logger
	for _, logger := range m.loggers {
		newLoggers = append(newLoggers, logger.WithGroup(prefix))
	}
	return &MultiLogger{loggers: newLoggers}
}

func NewMultiLogger(loggers ...slog.Handler) *MultiLogger {
	var newLoggers []*slog.Logger
	for _, logger := range loggers {
		newLoggers = append(newLoggers, slog.New(logger))
	}
	return &MultiLogger{loggers: newLoggers}
}

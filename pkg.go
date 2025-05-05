package mlog

import (
	"sync"
	"sync/atomic"
)

var (
	defaultLogger     atomic.Pointer[MultiLogger]
	defaultLoggerOnce sync.Once
)

func Default() *MultiLogger {
	dl := defaultLogger.Load()
	if dl == nil {
		defaultLoggerOnce.Do(func() {
			defaultLogger.CompareAndSwap(nil, NewMultiLogger())
		})
		dl = defaultLogger.Load()
	}
	return dl
}

func SetDefault(logger *MultiLogger) {
	defaultLogger.Store(logger)
}

func Debug(msg string, keysAndValues ...interface{}) {
	Default().Debug(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	Default().Info(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	Default().Warn(msg, keysAndValues...)
}

func Error(msg interface{}, keysAndValues ...interface{}) {
	Default().Error(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	Default().Fatal(msg, keysAndValues...)
}

package logger

import (
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_logger atomic.Value
	// defaultLogger logger for default usage
	_defaultLogger = newDefaultLogger()
	// RunningAtomicLevel supports changing level on the fly
	_runningAtomicLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
)

// GetLogger return logger with module name.
func GetLogger(module string) *Logger {
	return &Logger{
		module: module,
	}
}

// L returns under zap logger.
func L() *Logger {
	return new(Logger)
}

// newDefaultLogger creates a default logger for uninitialized usage.
func newDefaultLogger() *zap.Logger {
	logConf := zap.NewProductionConfig()
	logConf.DisableStacktrace = true
	logConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	l, err := logConf.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}
	return l
}

// InitLogger initializes a zap logger from user config.
func InitLogger(cfg Config) error {
	if err := initLogger(cfg); err != nil {
		return err
	}
	return nil
}

// initLogger initializes a zap logger for different module.
func initLogger(cfg Config) error {
	// parse logging level
	if err := _runningAtomicLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return err
	}
	conf := zap.NewProductionConfig()
	conf.Level = _runningAtomicLevel
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.DisableStacktrace = true
	l, err := conf.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	_logger.Store(l)
	return nil
}

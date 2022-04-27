package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field represents filed of zap logger.
type Field = zap.Field

// Logger is wrapper for zap logger with module, it is singleton.
type Logger struct {
	module string
	logger *zap.Logger
}

// getInitializedOrDefaultLogger try get initialized zap logger,
// if failure, it will use the default logger.
func (l *Logger) getInitializedOrDefaultLogger() *zap.Logger {
	if l.logger != nil {
		return l.logger
	}
	item := _logger.Load()
	if item == nil {
		return _defaultLogger
	}
	log, ok := item.(*zap.Logger)
	if ok {
		l.logger = log
	} else {
		l.logger = _defaultLogger
	}
	return l.logger
}

// getLogger returns under logger impl.
func (l *Logger) getLogger() *zap.Logger {
	if l.module == "" {
		return l.getInitializedOrDefaultLogger()
	}
	return l.getInitializedOrDefaultLogger().Named(l.module)
}

// With creates a child logger and adds structured context to it.
// Fields added to the child don't affect the parent, and vice versa.
func (l *Logger) With(fields ...Field) *Logger {
	return &Logger{
		module: l.module,
		logger: l.getLogger().With(fields...),
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Debug(msg string, fields ...Field) {
	l.getInitializedOrDefaultLogger().Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Info(msg string, fields ...Field) {
	l.getInitializedOrDefaultLogger().Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Warn(msg string, fields ...Field) {
	l.getInitializedOrDefaultLogger().Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Error(msg string, fields ...Field) {
	l.getInitializedOrDefaultLogger().Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.getInitializedOrDefaultLogger().Fatal(msg, fields...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.getInitializedOrDefaultLogger().Sugar().Infof(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.getInitializedOrDefaultLogger().Sugar().Errorf(template, args...)
}

// String constructs a field with the given key and value.
func String(key string, val string) Field {
	return Field{Key: key, Type: zapcore.StringType, String: val}
}

// Error is shorthand for the common idiom NamedError("error", err).
func Error(err error) Field {
	return zap.NamedError("error", err)
}

// Uint16 constructs a field with the given key and value.
func Uint16(key string, val uint16) Field {
	return Field{Key: key, Type: zapcore.Uint16Type, Integer: int64(val)}
}

// Uint32 constructs a field with the given key and value.
func Uint32(key string, val uint32) Field {
	return Field{Key: key, Type: zapcore.Uint32Type, Integer: int64(val)}
}

// Uint64 constructs a field with the given key and value.
func Uint64(key string, val uint64) Field {
	return Field{Key: key, Type: zapcore.Uint64Type, Integer: int64(val)}
}

// Stack constructs a field that stores a stacktrace of the current goroutine
// under provided key. Keep in mind that taking a stacktrace is eager and
// expensive (relatively speaking); this function both makes an allocation and
// takes about two microseconds.
func Stack() Field {
	return zap.Stack("stack")
}

// Reflect constructs a field with the given key and an arbitrary object. It uses
// an encoding-appropriate, reflection-based function to lazily serialize nearly
// any object into the logging context, but it's relatively slow and
// allocation-heavy. Outside tests, Any is always a better choice.
//
// If encoding fails (e.g., trying to serialize a map[int]string to JSON), Reflect
// includes the error message in the final log output.
func Reflect(key string, val interface{}) Field {
	return zap.Reflect(key, val)
}

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}

// Int32 constructs a field with the given key and value.
func Int32(key string, val int32) Field {
	return Field{Key: key, Type: zapcore.Int32Type, Integer: int64(val)}
}

// Int64 constructs a field with the given key and value.
func Int64(key string, val int64) Field {
	return Field{Key: key, Type: zapcore.Int64Type, Integer: val}
}

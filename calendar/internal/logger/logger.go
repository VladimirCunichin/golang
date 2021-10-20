package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger
var once sync.Once

//Init logger
func Init(logLevel, path string) {
	once.Do(func() {
		var level zapcore.Level
		err := level.UnmarshalText([]byte(logLevel))
		if err != nil {
			log.Fatalf("can't marshal level string: %v", logLevel)
		}
		cfg := zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(level),
			OutputPaths:      []string{path, "stderr"},
			ErrorOutputPaths: []string{path, "stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		logger, err := cfg.Build()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}
		defer logger.Sync()
		sugar = logger.Sugar()
	})
}

// Debug logs a debug message with the given fields
func Debug(message string, fields ...interface{}) {
	sugar.Debugw(message, fields...)
}

// Info logs a debug message with the given fields
func Info(message string, fields ...interface{}) {
	sugar.Infow(message, fields...)
}

// Warn logs a debug message with the given fields
func Warn(message string, fields ...interface{}) {
	sugar.Warnw(message, fields...)
}

// Error logs a debug message with the given fields
func Error(message string, fields ...interface{}) {
	sugar.Errorw(message, fields...)
}

// Fatal logs a message than calls os.Exit(1)
func Fatal(message string, fields ...interface{}) {
	sugar.Fatalw(message, fields...)
}

package logger

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

// NewLogger creates a new Logger instance with custom configurations
func NewLogger() *Logger {
	config := zap.Config{
		Encoding:         "console", // Use "json" for structured logs
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "logger",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, err := config.Build()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	return &Logger{logger: zapLogger}
}

// Info logs an informational message
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error message
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

// LogJSON converts JSON to string for readable logging
func (l *Logger) LogJSON(msg string, jsonData interface{}) {
	jsonString, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		l.Error("Failed to convert JSON to string", zap.Error(err))
		return
	}
	l.Info(msg, zap.String("data", string(jsonString)))
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() {
	_ = l.logger.Sync()
}

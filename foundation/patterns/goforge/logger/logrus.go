package logger

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yourusername/foundation/patterns/goforge/config"
)

// LogrusLogger implements the Logger interface using logrus package
type LogrusLogger struct {
	logger *logrus.Logger
	name   string
}

// NewLogrusLogger creates a new logrus-based logger
func NewLogrusLogger(cfg config.LoggerConfig) (Logger, error) {
	logger := logrus.New()

	// Determine output
	var output io.Writer
	switch cfg.Output {
	case "stderr":
		output = os.Stderr
	case "stdout", "":
		output = os.Stdout
	default:
		// Try to open file
		if file, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
			output = file
		} else {
			output = os.Stdout // fallback to stdout
		}
	}
	logger.SetOutput(output)

	// Determine format and level
	level := parseLogrusLevel(cfg.Level)
	logger.SetLevel(level)

	switch cfg.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text", "":
		logger.SetFormatter(&logrus.TextFormatter{})
	default:
		// fallback to text format
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	return &LogrusLogger{
		logger: logger,
		name:   "logrus-logger",
	}, nil
}

// parseLogrusLevel converts string level to logrus.Level
func parseLogrusLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info", "":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

// Service interface implementation
func (l *LogrusLogger) Start(ctx context.Context) error {
	l.logger.WithField("name", l.name).Info("Logger started")
	return nil
}

func (l *LogrusLogger) Stop(ctx context.Context) error {
	l.logger.WithField("name", l.name).Info("Logger stopped")
	return nil
}

func (l *LogrusLogger) Name() string {
	return l.name
}

// Logger interface implementation
func (l *LogrusLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg)
}

func (l *LogrusLogger) Info(msg string, args ...any) {
	l.logger.Info(msg)
}

func (l *LogrusLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg)
}

func (l *LogrusLogger) Error(msg string, args ...any) {
	l.logger.Error(msg)
}

func (l *LogrusLogger) With(args ...any) Logger {
	// For logrus, we'll create a new logger with fields
	// This is a simplified implementation
	return &LogrusLogger{
		logger: l.logger,
		name:   l.name,
	}
}

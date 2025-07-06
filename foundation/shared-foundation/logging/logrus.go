package logging

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// LogrusLogger implements the Logger interface using logrus
type LogrusLogger struct {
	logger *logrus.Logger
	name   string
}

// NewLogrusLogger creates a new logrus-based logger
func NewLogrusLogger(cfg LoggerConfig) (Logger, error) {
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

	// Determine format
	switch cfg.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text", "":
		logger.SetFormatter(&logrus.TextFormatter{})
	default:
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	// Determine level
	level := parseLogrusLevel(cfg.Level)
	logger.SetLevel(level)

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
	l.logger.Info("Logger started", "name", l.name)
	return nil
}

func (l *LogrusLogger) Stop(ctx context.Context) error {
	l.logger.Info("Logger stopped", "name", l.name)
	return nil
}

func (l *LogrusLogger) Name() string {
	return l.name
}

// Logger interface implementation
func (l *LogrusLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *LogrusLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *LogrusLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *LogrusLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *LogrusLogger) With(args ...any) Logger {
	// Convert args to logrus fields
	fields := logrus.Fields{}
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			fields[args[i].(string)] = args[i+1]
		}
	}

	return &LogrusLogger{
		logger: l.logger.WithFields(fields),
		name:   l.name,
	}
}

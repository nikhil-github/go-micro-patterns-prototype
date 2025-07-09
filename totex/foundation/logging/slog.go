package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// SlogLogger implements the Logger interface using Go's slog package
type SlogLogger struct {
	logger *slog.Logger
	name   string
}

// LoggerConfig configuration for the logger

// NewSlogLogger creates a new slog-based logger

func NewSlogLogger(cfg core.LoggerConfig) (Logger, error) {
	var handler slog.Handler

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

	// Determine format and level
	level := parseLogLevel(cfg.Level)
	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	case "text", "":
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	default:
		// fallback to text format
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	}

	return &SlogLogger{
		logger: slog.New(handler),
		name:   "slog-logger",
	}, nil
}

// parseLogLevel converts string level to slog.Level
func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info", "":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Service interface implementation
func (l *SlogLogger) Start(ctx context.Context) error {
	l.logger.Info("Logger started", "name", l.name)
	return nil
}

func (l *SlogLogger) Stop(ctx context.Context) error {
	l.logger.Info("Logger stopped", "name", l.name)
	return nil
}

func (l *SlogLogger) Name() string {
	return l.name
}

// Logger interface implementation
func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *SlogLogger) With(args ...any) Logger {
	return &SlogLogger{
		logger: l.logger.With(args...),
		name:   l.name,
	}
}

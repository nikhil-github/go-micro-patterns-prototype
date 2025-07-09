package logging

import (
	"log/slog"
	"os"
)

// SlogLogger implements Logger using slog
type SlogLogger struct {
	name string
	slog *slog.Logger
}

// NewSlogLogger creates a new slog-based logger
func NewSlogLogger(name string, level, format, output string) Logger {
	var h slog.Handler

	// Set log level
	levelOpt := &slog.LevelVar{}
	switch level {
	case "debug":
		levelOpt.Set(slog.LevelDebug)
	case "warn":
		levelOpt.Set(slog.LevelWarn)
	case "error":
		levelOpt.Set(slog.LevelError)
	default:
		levelOpt.Set(slog.LevelInfo)
	}

	// Set format
	if format == "json" {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: levelOpt,
		})
	} else {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: levelOpt,
		})
	}

	return &SlogLogger{
		name: name,
		slog: slog.New(h),
	}
}

// NewDefaultSlogLogger creates a default slog logger
func NewDefaultSlogLogger() Logger {
	return &SlogLogger{
		name: "default-logger",
		slog: slog.Default(),
	}
}

func (l *SlogLogger) Debug(msg string, args ...any) { l.slog.Debug(msg, args...) }
func (l *SlogLogger) Info(msg string, args ...any)  { l.slog.Info(msg, args...) }
func (l *SlogLogger) Warn(msg string, args ...any)  { l.slog.Warn(msg, args...) }
func (l *SlogLogger) Error(msg string, args ...any) { l.slog.Error(msg, args...) }
func (l *SlogLogger) With(args ...any) Logger       { return l }
func (l *SlogLogger) Name() string                  { return l.name }

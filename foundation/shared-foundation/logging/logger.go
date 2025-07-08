package logging

import (
	"github.com/yourusername/shared-foundation/core"
)

// Logger interface for structured logging
type Logger interface {
	core.Service
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
}

func WithSlogLogger(config core.LoggerConfig) any {
	panic("unimplemented")
}

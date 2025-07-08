package core

import "context"

// Service is the base interface that all services must implement.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

// LoggerConfig holds configuration for loggers, moved here for shared use.
type LoggerConfig struct {
	Type   string `mapstructure:"type"` // slog, logrus, zap, etc.
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

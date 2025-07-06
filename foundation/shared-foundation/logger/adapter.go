package logger

import (
	"context"
)

// LoggerAdapter adapts the logger package Logger to the foundation Logger interface
type LoggerAdapter struct {
	logger Logger
}

// NewLoggerAdapter creates a new adapter
func NewLoggerAdapter(logger Logger) *LoggerAdapter {
	return &LoggerAdapter{logger: logger}
}

// Start implements foundation.Service
func (a *LoggerAdapter) Start(ctx context.Context) error {
	return a.logger.Start(ctx)
}

// Stop implements foundation.Service
func (a *LoggerAdapter) Stop(ctx context.Context) error {
	return a.logger.Stop(ctx)
}

// Name implements foundation.Service
func (a *LoggerAdapter) Name() string {
	return a.logger.Name()
}

// Debug implements foundation.Logger
func (a *LoggerAdapter) Debug(msg string, args ...any) {
	a.logger.Debug(msg, args...)
}

// Info implements foundation.Logger
func (a *LoggerAdapter) Info(msg string, args ...any) {
	a.logger.Info(msg, args...)
}

// Warn implements foundation.Logger
func (a *LoggerAdapter) Warn(msg string, args ...any) {
	a.logger.Warn(msg, args...)
}

// Error implements foundation.Logger
func (a *LoggerAdapter) Error(msg string, args ...any) {
	a.logger.Error(msg, args...)
}

// With implements foundation.Logger
func (a *LoggerAdapter) With(args ...any) interface{} {
	return NewLoggerAdapter(a.logger.With(args...))
}

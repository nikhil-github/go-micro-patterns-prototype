package config

import (
	"log/slog"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := Load()

	if cfg.ConnectRPC.Address == "" {
		t.Error("expected non-empty ConnectRPC address")
	}

	if cfg.Logger.Level == "" {
		t.Error("expected non-empty logger level")
	}

	if cfg.Logger.Format == "" {
		t.Error("expected non-empty logger format")
	}

	if cfg.Logger.Output == "" {
		t.Error("expected non-empty logger output")
	}
}

func TestCreateLogger_TextFormat(t *testing.T) {
	cfg := &Config{
		Logger: Logger{
			Level:  "info",
			Format: "text",
			Output: "stdout",
		},
	}

	logger := cfg.CreateLogger()
	if logger == nil {
		t.Error("expected non-nil logger")
	}
}

func TestCreateLogger_JSONFormat(t *testing.T) {
	cfg := &Config{
		Logger: Logger{
			Level:  "debug",
			Format: "json",
			Output: "stdout",
		},
	}

	logger := cfg.CreateLogger()
	if logger == nil {
		t.Error("expected non-nil logger")
	}
}

func TestCreateLogger_StderrOutput(t *testing.T) {
	cfg := &Config{
		Logger: Logger{
			Level:  "warn",
			Format: "text",
			Output: "stderr",
		},
	}

	logger := cfg.CreateLogger()
	if logger == nil {
		t.Error("expected non-nil logger")
	}
}

func TestCreateLogger_FileOutput(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "test-log")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	cfg := &Config{
		Logger: Logger{
			Level:  "error",
			Format: "json",
			Output: tmpFile.Name(),
		},
	}

	logger := cfg.CreateLogger()
	if logger == nil {
		t.Error("expected non-nil logger")
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"error", slog.LevelError},
		{"", slog.LevelInfo},
		{"invalid", slog.LevelInfo},
	}

	for _, test := range tests {
		result := parseLogLevel(test.input)
		if result != test.expected {
			t.Errorf("parseLogLevel(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

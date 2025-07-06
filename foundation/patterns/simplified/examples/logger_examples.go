package main

import (
	"github.com/yourusername/foundation/patterns/simplified/config"
	"github.com/yourusername/foundation/patterns/simplified/servicefactory"
)

// Example 1: Default logger configuration
func exampleDefaultLogger() {
	cfg := config.Load()
	logger := cfg.CreateLogger() // Uses defaults: info level, text format, stdout

	factory := servicefactory.New(cfg, logger)
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	logger.Info("Starting with default logger configuration")
	// factory.Run()
}

// Example 2: JSON logger for production
func exampleJSONLogger() {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
		Logger: config.Logger{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		},
	}

	logger := cfg.CreateLogger()
	factory := servicefactory.New(cfg, logger)
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	logger.Info("Starting with JSON logger", "environment", "production")
	// factory.Run()
}

// Example 3: Debug logger for development
func exampleDebugLogger() {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
		Logger: config.Logger{
			Level:  "debug",
			Format: "text",
			Output: "stdout",
		},
	}

	logger := cfg.CreateLogger()
	factory := servicefactory.New(cfg, logger)
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	logger.Debug("Starting with debug logger", "environment", "development")
	logger.Info("Debug mode enabled")
	// factory.Run()
}

// Example 4: File-based logger
func exampleFileLogger() {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
		Logger: config.Logger{
			Level:  "warn",
			Format: "json",
			Output: "/var/log/myapp.log",
		},
	}

	logger := cfg.CreateLogger()
	factory := servicefactory.New(cfg, logger)
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	logger.Warn("Starting with file-based logger", "log_file", "/var/log/myapp.log")
	// factory.Run()
}

// Example 5: Error-only logger for minimal output
func exampleErrorOnlyLogger() {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
		Logger: config.Logger{
			Level:  "error",
			Format: "text",
			Output: "stderr",
		},
	}

	logger := cfg.CreateLogger()
	factory := servicefactory.New(cfg, logger)
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	// Only error messages will be logged
	logger.Error("Starting with error-only logger")
	// factory.Run()
}

// Example 6: Custom logger with environment variables
func exampleEnvBasedLogger() {
	// This would be set via environment variables:
	// export LOGGER_LEVEL="debug"
	// export LOGGER_FORMAT="json"
	// export LOGGER_OUTPUT="stdout"

	cfg := config.Load() // Loads from environment variables
	logger := cfg.CreateLogger()

	factory := servicefactory.New(cfg, logger)
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	logger.Info("Starting with environment-based logger configuration")
	// factory.Run()
}

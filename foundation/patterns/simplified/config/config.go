package config

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type ConnectRPC struct {
	Address string
}

type Logger struct {
	Level  string
	Format string
	Output string
}

type Config struct {
	ConnectRPC ConnectRPC
	Logger     Logger
}

func Load() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("CONNECTRPC_ADDRESS", ":8080")
	viper.SetDefault("LOGGER_LEVEL", "info")
	viper.SetDefault("LOGGER_FORMAT", "text")
	viper.SetDefault("LOGGER_OUTPUT", "stdout")

	return &Config{
		ConnectRPC: ConnectRPC{
			Address: viper.GetString("CONNECTRPC_ADDRESS"),
		},
		Logger: Logger{
			Level:  viper.GetString("LOGGER_LEVEL"),
			Format: viper.GetString("LOGGER_FORMAT"),
			Output: viper.GetString("LOGGER_OUTPUT"),
		},
	}
}

// CreateLogger creates a logger based on the configuration
func (c *Config) CreateLogger() *slog.Logger {
	var handler slog.Handler

	// Determine output
	var output *os.File
	switch c.Logger.Output {
	case "stderr":
		output = os.Stderr
	case "stdout", "":
		output = os.Stdout
	default:
		// Try to open file
		if file, err := os.OpenFile(c.Logger.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
			output = file
		} else {
			output = os.Stdout // fallback to stdout
		}
	}

	// Determine format
	switch c.Logger.Format {
	case "json":
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: parseLogLevel(c.Logger.Level),
		})
	case "text", "":
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: parseLogLevel(c.Logger.Level),
		})
	default:
		// fallback to text format
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: parseLogLevel(c.Logger.Level),
		})
	}

	return slog.New(handler)
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

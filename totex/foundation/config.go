package foundation

import (
	"os"
)

// AppConfig holds all configuration for the application
type AppConfig struct {
	Logger  LoggerConfig
	Tracer  TracerConfig
	Metrics MetricsConfig
	Servers []ServerConfig
}

// LoggerConfig configuration for the logger
type LoggerConfig struct {
	Type   string
	Level  string
	Format string
	Output string
}

// TracerConfig configuration for the tracer
type TracerConfig struct {
	Type     string
	Endpoint string
}

// MetricsConfig configuration for the metrics
type MetricsConfig struct {
	Type string
	Port string
}

// ServerConfig configuration for servers
type ServerConfig struct {
	Type string // "connectrpc", "http", etc.
	Name string
	Addr string
}

// LoadConfigFromEnv loads configuration from environment variables
func LoadConfigFromEnv() AppConfig {
	// Set defaults for missing environment variables
	setDefaultEnv("LOGGER_TYPE", "slog")
	setDefaultEnv("LOGGER_LEVEL", "info")
	setDefaultEnv("LOGGER_FORMAT", "text")
	setDefaultEnv("LOGGER_OUTPUT", "stdout")
	setDefaultEnv("TRACER_TYPE", "noop")
	setDefaultEnv("TRACER_ENDPOINT", "")
	setDefaultEnv("METRICS_TYPE", "noop")
	setDefaultEnv("METRICS_PORT", "9090")
	setDefaultEnv("SERVER_NAME", "server")
	setDefaultEnv("SERVER_ADDR", ":8080")

	// Parse server configuration
	servers := parseServerConfig()

	return AppConfig{
		Logger: LoggerConfig{
			Type:   os.Getenv("LOGGER_TYPE"),
			Level:  os.Getenv("LOGGER_LEVEL"),
			Format: os.Getenv("LOGGER_FORMAT"),
			Output: os.Getenv("LOGGER_OUTPUT"),
		},
		Tracer: TracerConfig{
			Type:     os.Getenv("TRACER_TYPE"),
			Endpoint: os.Getenv("TRACER_ENDPOINT"),
		},
		Metrics: MetricsConfig{
			Type: os.Getenv("METRICS_TYPE"),
			Port: os.Getenv("METRICS_PORT"),
		},
		Servers: servers,
	}
}

// setDefaultEnv sets an environment variable if it's not already set
func setDefaultEnv(key, value string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, value)
	}
}

// parseServerConfig parses server configuration from environment variables
func parseServerConfig() []ServerConfig {
	// For now, we support a single server configuration
	// This can be extended to support multiple servers via SERVER_1_TYPE, SERVER_1_ADDR, etc.
	serverType := os.Getenv("SERVER_TYPE")
	if serverType == "" {
		serverType = "connectrpc" // Default to ConnectRPC
	}

	return []ServerConfig{
		{
			Type: serverType,
			Name: os.Getenv("SERVER_NAME"),
			Addr: os.Getenv("SERVER_ADDR"),
		},
	}
}

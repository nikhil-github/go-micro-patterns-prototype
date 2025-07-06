package foundation

import (
	"time"

	"github.com/spf13/viper"
)

// AppConfig holds application metadata
type AppConfig struct {
	Name    string
	Version string
	Env     string
}

// LoggerConfig configuration
type LoggerConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

// TracerConfig configuration
type TracerConfig struct {
	Type        string  `mapstructure:"type"` // jaeger, zipkin, etc.
	Endpoint    string  `mapstructure:"endpoint"`
	ServiceName string  `mapstructure:"service_name"`
	SampleRate  float64 `mapstructure:"sample_rate"`
}

// MetricsConfig configuration
type MetricsConfig struct {
	Type string `mapstructure:"type"` // prometheus, statsd, etc.
	Port int    `mapstructure:"port"`
	Path string `mapstructure:"path"`
}

// ServiceDiscoveryConfig configuration
type ServiceDiscoveryConfig struct {
	Type     string `mapstructure:"type"` // consul, etcd, etc.
	Endpoint string `mapstructure:"endpoint"`
	Token    string `mapstructure:"token"`
}

// BrokerConfig configuration
type BrokerConfig struct {
	Type    string   `mapstructure:"type"` // kafka, rabbitmq, etc.
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`
}

// CacheConfig configuration
type CacheConfig struct {
	Type    string        `mapstructure:"type"` // redis, memcached, etc.
	Address string        `mapstructure:"address"`
	TTL     time.Duration `mapstructure:"ttl"`
}

// DatabaseConfig configuration
type DatabaseConfig struct {
	Type string `mapstructure:"type"` // postgres, mysql, etc.
	DSN  string `mapstructure:"dsn"`
}

// ConnectRPCConfig configuration
type ConnectRPCConfig struct {
	Address string `mapstructure:"address"`
}

// Config holds all configuration
type Config struct {
	App              AppConfig              `mapstructure:"app"`
	Logger           LoggerConfig           `mapstructure:"logger"`
	Tracer           TracerConfig           `mapstructure:"tracer"`
	Metrics          MetricsConfig          `mapstructure:"metrics"`
	ServiceDiscovery ServiceDiscoveryConfig `mapstructure:"service_discovery"`
	Broker           BrokerConfig           `mapstructure:"broker"`
	Cache            CacheConfig            `mapstructure:"cache"`
	Database         DatabaseConfig         `mapstructure:"database"`
	ConnectRPC       ConnectRPCConfig       `mapstructure:"connectrpc"`
}

// Load loads configuration from environment variables
func Load() *Config {
	viper.AutomaticEnv()

	// App defaults
	viper.SetDefault("APP_NAME", "microservice")
	viper.SetDefault("APP_VERSION", "1.0.0")
	viper.SetDefault("APP_ENV", "development")

	// Logger defaults
	viper.SetDefault("LOGGER_LEVEL", "info")
	viper.SetDefault("LOGGER_FORMAT", "text")
	viper.SetDefault("LOGGER_OUTPUT", "stdout")

	// Tracer defaults
	viper.SetDefault("TRACER_TYPE", "jaeger")
	viper.SetDefault("TRACER_ENDPOINT", "http://jaeger:14268")
	viper.SetDefault("TRACER_SAMPLE_RATE", 1.0)

	// Metrics defaults
	viper.SetDefault("METRICS_TYPE", "prometheus")
	viper.SetDefault("METRICS_PORT", 9090)
	viper.SetDefault("METRICS_PATH", "/metrics")

	// Service Discovery defaults
	viper.SetDefault("SERVICE_DISCOVERY_TYPE", "consul")
	viper.SetDefault("SERVICE_DISCOVERY_ENDPOINT", "http://consul:8500")

	// Broker defaults
	viper.SetDefault("BROKER_TYPE", "kafka")
	viper.SetDefault("BROKER_BROKERS", "kafka:9092")

	// Cache defaults
	viper.SetDefault("CACHE_TYPE", "redis")
	viper.SetDefault("CACHE_ADDRESS", "redis:6379")
	viper.SetDefault("CACHE_TTL", "1h")

	// Database defaults
	viper.SetDefault("DATABASE_TYPE", "postgres")
	viper.SetDefault("DATABASE_DSN", "postgres://user:pass@db:5432/mydb")

	// ConnectRPC defaults
	viper.SetDefault("CONNECTRPC_ADDRESS", ":8080")

	var cfg Config
	viper.Unmarshal(&cfg)

	return &cfg
}

package foundation

import (
	"github.com/yourusername/shared-foundation/logging"
)

// AppOption is a function that configures an App
type AppOption func(*App)

// WithLogger configures the logger
func WithLogger(logger Logger) AppOption {
	return func(a *App) {
		a.logger = logger
	}
}

// WithSlogLogger creates a slog logger with the given config
func WithSlogLogger(cfg LoggerConfig) AppOption {
	return func(a *App) {
		loggerCfg := logging.LoggerConfig{
			Level:  cfg.Level,
			Format: cfg.Format,
			Output: cfg.Output,
		}
		if logger, err := logging.NewSlogLogger(loggerCfg); err == nil {
			a.logger = logger
		}
	}
}

// WithLogrusLogger creates a logrus logger with the given config
func WithLogrusLogger(cfg LoggerConfig) AppOption {
	return func(a *App) {
		loggerCfg := logging.LoggerConfig{
			Level:  cfg.Level,
			Format: cfg.Format,
			Output: cfg.Output,
		}
		if logger, err := logging.NewLogrusLogger(loggerCfg); err == nil {
			a.logger = logger
		}
	}
}

// WithTracer configures the tracer
func WithTracer(tracer Tracer) AppOption {
	return func(a *App) {
		a.tracer = tracer
	}
}

// WithMockTracer creates a mock tracer
func WithMockTracer() AppOption {
	return func(a *App) {
		a.tracer = &MockTracer{name: "tracer"}
	}
}

// WithMetrics configures the metrics
func WithMetrics(metrics Metrics) AppOption {
	return func(a *App) {
		a.metrics = metrics
	}
}

// WithMockMetrics creates a mock metrics
func WithMockMetrics() AppOption {
	return func(a *App) {
		a.metrics = &MockMetrics{name: "metrics"}
	}
}

// WithBroker configures the broker
func WithBroker(broker Broker) AppOption {
	return func(a *App) {
		a.broker = broker
	}
}

// WithMockBroker creates a mock broker
func WithMockBroker() AppOption {
	return func(a *App) {
		a.broker = &MockBroker{name: "broker"}
	}
}

// WithCache configures the cache
func WithCache(cache Cache) AppOption {
	return func(a *App) {
		a.cache = cache
	}
}

// WithMockCache creates a mock cache
func WithMockCache() AppOption {
	return func(a *App) {
		a.cache = &MockCache{name: "cache"}
	}
}

// WithDatabase configures the database
func WithDatabase(database Database) AppOption {
	return func(a *App) {
		a.database = database
	}
}

// WithMockDatabase creates a mock database
func WithMockDatabase() AppOption {
	return func(a *App) {
		a.database = &MockDatabase{name: "database"}
	}
}

// WithConnectRPCServer configures the ConnectRPC server
func WithConnectRPCServer(server ConnectRPCServer) AppOption {
	return func(a *App) {
		a.connectRPCServer = server
	}
}

// WithMockConnectRPCServer creates a mock ConnectRPC server
func WithMockConnectRPCServer() AppOption {
	return func(a *App) {
		a.connectRPCServer = &MockConnectRPCServer{name: "connectrpc-server"}
	}
}

// WithServiceDiscovery configures the service discovery
func WithServiceDiscovery(discovery ServiceDiscovery) AppOption {
	return func(a *App) {
		a.serviceDiscovery = discovery
	}
}

// WithMockServiceDiscovery creates a mock service discovery
func WithMockServiceDiscovery() AppOption {
	return func(a *App) {
		a.serviceDiscovery = &MockServiceDiscovery{name: "service-discovery"}
	}
}

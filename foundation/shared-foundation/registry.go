package foundation

import (
	"fmt"
	"sync"

	"github.com/yourusername/shared-foundation/logging"
)

// CreatorFunc is a function that creates a service instance
type CreatorFunc func(config interface{}, deps ...interface{}) (interface{}, error)

// Registry holds service creators
type Registry struct {
	creators map[string]CreatorFunc
	mu       sync.RWMutex
}

// NewRegistry creates a new service registry
func NewRegistry() *Registry {
	return &Registry{
		creators: make(map[string]CreatorFunc),
	}
}

// Register adds a service creator to the registry
func (r *Registry) Register(serviceType string, creator CreatorFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.creators[serviceType] = creator
}

// Create creates a service instance using the registered creator
func (r *Registry) Create(serviceType string, config interface{}, deps ...interface{}) (interface{}, error) {
	r.mu.RLock()
	creator, exists := r.creators[serviceType]
	r.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("unknown service type: %s", serviceType)
	}

	return creator(config, deps...)
}

// Global registry instance
var globalRegistry = NewRegistry()

// RegisterLogger registers a logger creator
func RegisterLogger(loggerType string, creator CreatorFunc) {
	globalRegistry.Register("logger:"+loggerType, creator)
}

// RegisterTracer registers a tracer creator
func RegisterTracer(tracerType string, creator CreatorFunc) {
	globalRegistry.Register("tracer:"+tracerType, creator)
}

// RegisterMetrics registers a metrics creator
func RegisterMetrics(metricsType string, creator CreatorFunc) {
	globalRegistry.Register("metrics:"+metricsType, creator)
}

// RegisterBroker registers a broker creator
func RegisterBroker(brokerType string, creator CreatorFunc) {
	globalRegistry.Register("broker:"+brokerType, creator)
}

// RegisterCache registers a cache creator
func RegisterCache(cacheType string, creator CreatorFunc) {
	globalRegistry.Register("cache:"+cacheType, creator)
}

// RegisterDatabase registers a database creator
func RegisterDatabase(dbType string, creator CreatorFunc) {
	globalRegistry.Register("database:"+dbType, creator)
}

// RegisterServiceDiscovery registers a service discovery creator
func RegisterServiceDiscovery(discoveryType string, creator CreatorFunc) {
	globalRegistry.Register("service_discovery:"+discoveryType, creator)
}

// RegisterConnectRPCServer registers a ConnectRPC server creator
func RegisterConnectRPCServer(serverType string, creator CreatorFunc) {
	globalRegistry.Register("connectrpc:"+serverType, creator)
}

// CreateLogger creates a logger using the registry
func CreateLogger(loggerType string, cfg LoggerConfig) (Logger, error) {
	result, err := globalRegistry.Create("logger:"+loggerType, cfg)
	if err != nil {
		return nil, err
	}
	return result.(Logger), nil
}

// CreateTracer creates a tracer using the registry
func CreateTracer(tracerType string, cfg TracerConfig, logger Logger) (Tracer, error) {
	result, err := globalRegistry.Create("tracer:"+tracerType, cfg, logger)
	if err != nil {
		return nil, err
	}
	return result.(Tracer), nil
}

// CreateMetrics creates metrics using the registry
func CreateMetrics(metricsType string, cfg MetricsConfig, logger Logger) (Metrics, error) {
	result, err := globalRegistry.Create("metrics:"+metricsType, cfg, logger)
	if err != nil {
		return nil, err
	}
	return result.(Metrics), nil
}

// CreateBroker creates a broker using the registry
func CreateBroker(brokerType string, cfg BrokerConfig, logger Logger, tracer Tracer) (Broker, error) {
	result, err := globalRegistry.Create("broker:"+brokerType, cfg, logger, tracer)
	if err != nil {
		return nil, err
	}
	return result.(Broker), nil
}

// CreateCache creates a cache using the registry
func CreateCache(cacheType string, cfg CacheConfig, logger Logger, tracer Tracer) (Cache, error) {
	result, err := globalRegistry.Create("cache:"+cacheType, cfg, logger, tracer)
	if err != nil {
		return nil, err
	}
	return result.(Cache), nil
}

// CreateDatabase creates a database using the registry
func CreateDatabase(dbType string, cfg DatabaseConfig, logger Logger, tracer Tracer) (Database, error) {
	result, err := globalRegistry.Create("database:"+dbType, cfg, logger, tracer)
	if err != nil {
		return nil, err
	}
	return result.(Database), nil
}

// CreateServiceDiscovery creates a service discovery using the registry
func CreateServiceDiscovery(discoveryType string, cfg ServiceDiscoveryConfig, logger Logger) (ServiceDiscovery, error) {
	result, err := globalRegistry.Create("service_discovery:"+discoveryType, cfg, logger)
	if err != nil {
		return nil, err
	}
	return result.(ServiceDiscovery), nil
}

// CreateConnectRPCServer creates a ConnectRPC server using the registry
func CreateConnectRPCServer(serverType string, cfg ConnectRPCConfig, logger Logger, tracer Tracer, metrics Metrics) (ConnectRPCServer, error) {
	result, err := globalRegistry.Create("connectrpc:"+serverType, cfg, logger, tracer, metrics)
	if err != nil {
		return nil, err
	}
	return result.(ConnectRPCServer), nil
}

// init registers default implementations
func init() {
	// Register default logger implementations
	RegisterLogger("slog", func(config interface{}, deps ...interface{}) (interface{}, error) {
		cfg := config.(LoggerConfig)
		loggerCfg := logging.LoggerConfig{
			Level:  cfg.Level,
			Format: cfg.Format,
			Output: cfg.Output,
		}
		return logging.NewSlogLogger(loggerCfg)
	})

	RegisterLogger("logrus", func(config interface{}, deps ...interface{}) (interface{}, error) {
		cfg := config.(LoggerConfig)
		loggerCfg := logging.LoggerConfig{
			Level:  cfg.Level,
			Format: cfg.Format,
			Output: cfg.Output,
		}
		return logging.NewLogrusLogger(loggerCfg)
	})

	// Register default tracer implementations
	RegisterTracer("mock", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockTracer{name: "tracer"}, nil
	})

	// Register default metrics implementations
	RegisterMetrics("mock", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockMetrics{name: "metrics"}, nil
	})

	// Register default broker implementations
	RegisterBroker("mock", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockBroker{name: "broker"}, nil
	})

	// Register default cache implementations
	RegisterCache("mock", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockCache{name: "cache"}, nil
	})

	// Register default database implementations
	RegisterDatabase("mock", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockDatabase{name: "database"}, nil
	})

	// Register default service discovery implementations
	RegisterServiceDiscovery("mock", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockServiceDiscovery{name: "service-discovery"}, nil
	})

	// Register default ConnectRPC server implementations
	RegisterConnectRPCServer("default", func(config interface{}, deps ...interface{}) (interface{}, error) {
		return &MockConnectRPCServer{name: "connectrpc-server"}, nil
	})
}

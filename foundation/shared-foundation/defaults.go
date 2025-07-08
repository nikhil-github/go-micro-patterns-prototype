package foundation

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/yourusername/shared-foundation/core"
	"github.com/yourusername/shared-foundation/messaging"
	"github.com/yourusername/shared-foundation/tracing"
)

// DefaultLogger creates a slog logger with sensible defaults
func DefaultLogger() Logger {
	// Use environment variables or sensible defaults
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	format := os.Getenv("LOG_FORMAT")
	if format == "" {
		format = "text"
	}

	output := os.Getenv("LOG_OUTPUT")
	if output == "" {
		output = "stdout"
	}

	loggerCfg := core.LoggerConfig{
		Level:  level,
		Format: format,
		Output: output,
	}

	logger, err := NewSlogLogger(loggerCfg)
	if err != nil {
		// Logging is critical - panic if we can't create a logger
		panic(fmt.Sprintf("failed to create default logger: %v", err))
	}

	return logger
}

// NewSlogLogger creates a new slog-based logger
func NewSlogLogger(cfg core.LoggerConfig) (Logger, error) {
	// This would call the actual slog implementation
	// For now, return a mock to avoid import cycles
	return &MockLogger{name: "slog-logger"}, nil
}

// DefaultTracer creates a no-op tracer (lightweight, no external dependencies)
func DefaultTracer() Tracer {
	return &NoOpTracer{name: "noop-tracer"}
}

// DefaultMetrics creates a no-op metrics (lightweight, no external dependencies)
func DefaultMetrics() Metrics {
	return &NoOpMetrics{name: "noop-metrics"}
}

// DefaultBroker creates a no-op broker (lightweight, no external dependencies)
func DefaultBroker() Broker {
	return &NoOpBroker{name: "noop-broker"}
}

// DefaultCache creates a no-op cache (lightweight, no external dependencies)
func DefaultCache() Cache {
	return &NoOpCache{name: "noop-cache"}
}

// DefaultDatabase creates a no-op database (lightweight, no external dependencies)
func DefaultDatabase() Database {
	return &NoOpDatabase{name: "noop-database"}
}

// DefaultServiceDiscovery creates a no-op service discovery (lightweight, no external dependencies)
func DefaultServiceDiscovery() ServiceDiscovery {
	return &NoOpServiceDiscovery{name: "noop-service-discovery"}
}

// DefaultConnectRPCServer creates a no-op ConnectRPC server (lightweight, no external dependencies)
func DefaultConnectRPCServer() ConnectRPCServer {
	return &NoOpConnectRPCServer{name: "noop-connectrpc-server"}
}

// NoOp implementations that are lightweight and don't require external dependencies

type NoOpTracer struct {
	name string
}

func (n *NoOpTracer) Start(ctx context.Context) error { return nil }
func (n *NoOpTracer) Stop(ctx context.Context) error  { return nil }
func (n *NoOpTracer) Name() string                    { return n.name }
func (n *NoOpTracer) StartSpan(name string, opts ...tracing.SpanOption) tracing.Span {
	return &NoOpSpan{}
}
func (n *NoOpTracer) Inject(span tracing.Span, format interface{}, carrier interface{}) error {
	return nil
}
func (n *NoOpTracer) Extract(format interface{}, carrier interface{}) (tracing.Span, error) {
	return &NoOpSpan{}, nil
}

type NoOpSpan struct{}

func (n *NoOpSpan) SetTag(key, value string) {}
func (n *NoOpSpan) SetError(err error)       {}
func (n *NoOpSpan) Finish()                  {}
func (n *NoOpSpan) Context() context.Context { return context.Background() }

type NoOpMetrics struct {
	name string
}

func (n *NoOpMetrics) Start(ctx context.Context) error                        { return nil }
func (n *NoOpMetrics) Stop(ctx context.Context) error                         { return nil }
func (n *NoOpMetrics) Name() string                                           { return n.name }
func (n *NoOpMetrics) Counter(name string, value float64, labels ...string)   {}
func (n *NoOpMetrics) Gauge(name string, value float64, labels ...string)     {}
func (n *NoOpMetrics) Histogram(name string, value float64, labels ...string) {}
func (n *NoOpMetrics) Summary(name string, value float64, labels ...string)   {}

type NoOpBroker struct {
	name string
}

func (n *NoOpBroker) Start(ctx context.Context) error                                { return nil }
func (n *NoOpBroker) Stop(ctx context.Context) error                                 { return nil }
func (n *NoOpBroker) Name() string                                                   { return n.name }
func (n *NoOpBroker) Publish(topic string, message []byte) error                     { return nil }
func (n *NoOpBroker) Subscribe(topic string, handler messaging.MessageHandler) error { return nil }
func (n *NoOpBroker) Unsubscribe(topic string) error                                 { return nil }

type NoOpCache struct {
	name string
}

func (n *NoOpCache) Start(ctx context.Context) error                       { return nil }
func (n *NoOpCache) Stop(ctx context.Context) error                        { return nil }
func (n *NoOpCache) Name() string                                          { return n.name }
func (n *NoOpCache) Get(key string) ([]byte, error)                        { return nil, nil }
func (n *NoOpCache) Set(key string, value []byte, ttl time.Duration) error { return nil }
func (n *NoOpCache) Delete(key string) error                               { return nil }
func (n *NoOpCache) Exists(key string) (bool, error)                       { return false, nil }
func (n *NoOpCache) Incr(key string) (int64, error)                        { return 0, nil }

type NoOpDatabase struct {
	name string
}

func (n *NoOpDatabase) Start(ctx context.Context) error { return nil }
func (n *NoOpDatabase) Stop(ctx context.Context) error  { return nil }
func (n *NoOpDatabase) Name() string                    { return n.name }
func (n *NoOpDatabase) Connect() error                  { return nil }
func (n *NoOpDatabase) Disconnect() error               { return nil }
func (n *NoOpDatabase) Health() error                   { return nil }

type NoOpServiceDiscovery struct {
	name string
}

func (n *NoOpServiceDiscovery) Start(ctx context.Context) error    { return nil }
func (n *NoOpServiceDiscovery) Stop(ctx context.Context) error     { return nil }
func (n *NoOpServiceDiscovery) Name() string                       { return n.name }
func (n *NoOpServiceDiscovery) Register(service ServiceInfo) error { return nil }
func (n *NoOpServiceDiscovery) Deregister(serviceID string) error  { return nil }
func (n *NoOpServiceDiscovery) GetService(name string) ([]ServiceInfo, error) {
	return []ServiceInfo{}, nil
}
func (n *NoOpServiceDiscovery) Watch(name string) (<-chan []ServiceInfo, error) {
	return make(chan []ServiceInfo), nil
}

type NoOpConnectRPCServer struct {
	name string
}

func (n *NoOpConnectRPCServer) Start(ctx context.Context) error                        { return nil }
func (n *NoOpConnectRPCServer) Stop(ctx context.Context) error                         { return nil }
func (n *NoOpConnectRPCServer) Name() string                                           { return n.name }
func (n *NoOpConnectRPCServer) RegisterHandler(path string, handler interface{}) error { return nil }
func (n *NoOpConnectRPCServer) GetHandler() interface{}                                { return nil }

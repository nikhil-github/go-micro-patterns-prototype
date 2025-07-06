package foundation

import (
	"context"
	"time"
)

// Service is the base interface that all services must implement.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

// Logger interface for structured logging
type Logger interface {
	Service
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
}

// Tracer interface for distributed tracing
type Tracer interface {
	Service
	StartSpan(name string, opts ...SpanOption) Span
	Inject(span Span, format interface{}, carrier interface{}) error
	Extract(format interface{}, carrier interface{}) (Span, error)
}

// Span represents a tracing span
type Span interface {
	SetTag(key, value string)
	SetError(err error)
	Finish()
	Context() context.Context
}

// SpanOption configures span creation
type SpanOption func(Span)

// Metrics interface for observability
type Metrics interface {
	Service
	Counter(name string, value float64, labels ...string)
	Gauge(name string, value float64, labels ...string)
	Histogram(name string, value float64, labels ...string)
	Summary(name string, value float64, labels ...string)
}

// ServiceDiscovery interface for service registration and discovery
type ServiceDiscovery interface {
	Service
	Register(service ServiceInfo) error
	Deregister(serviceID string) error
	GetService(name string) ([]ServiceInfo, error)
	Watch(name string) (<-chan []ServiceInfo, error)
}

// ServiceInfo represents a service in the discovery system
type ServiceInfo struct {
	ID      string
	Name    string
	Address string
	Port    int
	Tags    []string
	Meta    map[string]string
}

// Broker interface for message queuing/pub-sub
type Broker interface {
	Service
	Publish(topic string, message []byte) error
	Subscribe(topic string, handler MessageHandler) error
	Unsubscribe(topic string) error
}

// MessageHandler processes messages from broker
type MessageHandler func(topic string, message []byte) error

// Cache interface for caching
type Cache interface {
	Service
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
	Exists(key string) (bool, error)
	Incr(key string) (int64, error)
}

// Database interface for data persistence
type Database interface {
	Service
	Connect() error
	Disconnect() error
	Health() error
	// Add more database-specific methods as needed
	// Query(query string, args ...interface{}) (Result, error)
	// Transaction() (Transaction, error)
}

// ConnectRPCServer interface for gRPC/Connect-RPC server
type ConnectRPCServer interface {
	Service
	RegisterHandler(path string, handler interface{}) error
	GetHandler() interface{} // Returns the underlying handler for registration
}

package foundation

import (
	"context"
	"time"
)

// Mock implementations for demonstration

type MockLogger struct {
	name string
}

func (m *MockLogger) Start(ctx context.Context) error { return nil }
func (m *MockLogger) Stop(ctx context.Context) error  { return nil }
func (m *MockLogger) Name() string                    { return m.name }
func (m *MockLogger) Debug(msg string, args ...any)   {}
func (m *MockLogger) Info(msg string, args ...any)    {}
func (m *MockLogger) Warn(msg string, args ...any)    {}
func (m *MockLogger) Error(msg string, args ...any)   {}
func (m *MockLogger) With(args ...any) Logger         { return m }

type MockTracer struct {
	name string
}

func (m *MockTracer) Start(ctx context.Context) error { return nil }
func (m *MockTracer) Stop(ctx context.Context) error  { return nil }
func (m *MockTracer) Name() string                    { return m.name }
func (m *MockTracer) StartSpan(name string, opts ...SpanOption) Span {
	return &MockSpan{}
}
func (m *MockTracer) Inject(span Span, format interface{}, carrier interface{}) error {
	return nil
}
func (m *MockTracer) Extract(format interface{}, carrier interface{}) (Span, error) {
	return &MockSpan{}, nil
}

type MockSpan struct{}

func (m *MockSpan) SetTag(key, value string) {}
func (m *MockSpan) SetError(err error)       {}
func (m *MockSpan) Finish()                  {}
func (m *MockSpan) Context() context.Context { return context.Background() }

type MockMetrics struct {
	name string
}

func (m *MockMetrics) Start(ctx context.Context) error                        { return nil }
func (m *MockMetrics) Stop(ctx context.Context) error                         { return nil }
func (m *MockMetrics) Name() string                                           { return m.name }
func (m *MockMetrics) Counter(name string, value float64, labels ...string)   {}
func (m *MockMetrics) Gauge(name string, value float64, labels ...string)     {}
func (m *MockMetrics) Histogram(name string, value float64, labels ...string) {}
func (m *MockMetrics) Summary(name string, value float64, labels ...string)   {}

type MockServiceDiscovery struct {
	name string
}

func (m *MockServiceDiscovery) Start(ctx context.Context) error    { return nil }
func (m *MockServiceDiscovery) Stop(ctx context.Context) error     { return nil }
func (m *MockServiceDiscovery) Name() string                       { return m.name }
func (m *MockServiceDiscovery) Register(service ServiceInfo) error { return nil }
func (m *MockServiceDiscovery) Deregister(serviceID string) error  { return nil }
func (m *MockServiceDiscovery) GetService(name string) ([]ServiceInfo, error) {
	return []ServiceInfo{}, nil
}
func (m *MockServiceDiscovery) Watch(name string) (<-chan []ServiceInfo, error) {
	return make(chan []ServiceInfo), nil
}

type MockBroker struct {
	name string
}

func (m *MockBroker) Start(ctx context.Context) error                      { return nil }
func (m *MockBroker) Stop(ctx context.Context) error                       { return nil }
func (m *MockBroker) Name() string                                         { return m.name }
func (m *MockBroker) Publish(topic string, message []byte) error           { return nil }
func (m *MockBroker) Subscribe(topic string, handler MessageHandler) error { return nil }
func (m *MockBroker) Unsubscribe(topic string) error                       { return nil }

type MockCache struct {
	name string
}

func (m *MockCache) Start(ctx context.Context) error                       { return nil }
func (m *MockCache) Stop(ctx context.Context) error                        { return nil }
func (m *MockCache) Name() string                                          { return m.name }
func (m *MockCache) Get(key string) ([]byte, error)                        { return nil, nil }
func (m *MockCache) Set(key string, value []byte, ttl time.Duration) error { return nil }
func (m *MockCache) Delete(key string) error                               { return nil }
func (m *MockCache) Exists(key string) (bool, error)                       { return false, nil }
func (m *MockCache) Incr(key string) (int64, error)                        { return 0, nil }

type MockDatabase struct {
	name string
}

func (m *MockDatabase) Start(ctx context.Context) error { return nil }
func (m *MockDatabase) Stop(ctx context.Context) error  { return nil }
func (m *MockDatabase) Name() string                    { return m.name }
func (m *MockDatabase) Connect() error                  { return nil }
func (m *MockDatabase) Disconnect() error               { return nil }
func (m *MockDatabase) Health() error                   { return nil }

type MockConnectRPCServer struct {
	name string
}

func (m *MockConnectRPCServer) Start(ctx context.Context) error                        { return nil }
func (m *MockConnectRPCServer) Stop(ctx context.Context) error                         { return nil }
func (m *MockConnectRPCServer) Name() string                                           { return m.name }
func (m *MockConnectRPCServer) RegisterHandler(path string, handler interface{}) error { return nil }
func (m *MockConnectRPCServer) GetHandler() interface{}                                { return nil }

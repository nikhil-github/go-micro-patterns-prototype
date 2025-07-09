package metrics

// Metrics interface for observability
type Metrics interface {
	Counter(name string, value float64, labels ...string)
	Gauge(name string, value float64, labels ...string)
	Histogram(name string, value float64, labels ...string)
	Summary(name string, value float64, labels ...string)
	Name() string
}

// NoopMetrics is a no-operation metrics implementation
type NoopMetrics struct {
	name string
}

// NewNoopMetrics creates a new noop metrics implementation
func NewNoopMetrics(name string) Metrics {
	return &NoopMetrics{name: name}
}

// NewDefaultMetrics creates a default noop metrics implementation
func NewDefaultMetrics() Metrics {
	return &NoopMetrics{name: "default-metrics"}
}

func (m *NoopMetrics) Counter(name string, value float64, labels ...string)   {}
func (m *NoopMetrics) Gauge(name string, value float64, labels ...string)     {}
func (m *NoopMetrics) Histogram(name string, value float64, labels ...string) {}
func (m *NoopMetrics) Summary(name string, value float64, labels ...string)   {}
func (m *NoopMetrics) Name() string                                           { return m.name }

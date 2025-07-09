package metrics

// Metrics interface for observability
type Metrics interface {
	Counter(name string, value float64, labels ...string)
	Gauge(name string, value float64, labels ...string)
	Histogram(name string, value float64, labels ...string)
	Summary(name string, value float64, labels ...string)
	Name() string
}

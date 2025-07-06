package metrics

import "github.com/yourusername/shared-foundation/core"

// Metrics interface for observability
type Metrics interface {
	core.Service
	Counter(name string, value float64, labels ...string)
	Gauge(name string, value float64, labels ...string)
	Histogram(name string, value float64, labels ...string)
	Summary(name string, value float64, labels ...string)
}

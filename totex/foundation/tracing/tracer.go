package tracing

import (
	"context"
)

// Tracer interface for distributed tracing
type Tracer interface {
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

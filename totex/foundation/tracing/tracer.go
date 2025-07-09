package tracing

import (
	"context"
)

// Tracer interface for distributed tracing
type Tracer interface {
	StartSpan(name string, opts ...SpanOption) Span
	Inject(span Span, format interface{}, carrier interface{}) error
	Extract(format interface{}, carrier interface{}) (Span, error)
	Name() string
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

// NoopTracer is a no-operation tracer implementation
type NoopTracer struct {
	name string
}

// NoopSpan is a no-operation span implementation
type NoopSpan struct{}

// NewNoopTracer creates a new noop tracer implementation
func NewNoopTracer(name string) Tracer {
	return &NoopTracer{name: name}
}

// NewDefaultTracer creates a default noop tracer implementation
func NewDefaultTracer() Tracer {
	return &NoopTracer{name: "default-tracer"}
}

func (t *NoopTracer) StartSpan(name string, opts ...SpanOption) Span {
	return &NoopSpan{}
}

func (t *NoopTracer) Inject(span Span, format interface{}, carrier interface{}) error {
	return nil
}

func (t *NoopTracer) Extract(format interface{}, carrier interface{}) (Span, error) {
	return &NoopSpan{}, nil
}

func (t *NoopTracer) Name() string {
	return t.name
}

func (s *NoopSpan) SetTag(key, value string) {}
func (s *NoopSpan) SetError(err error)       {}
func (s *NoopSpan) Finish()                  {}
func (s *NoopSpan) Context() context.Context {
	return context.Background()
}

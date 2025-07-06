package core

import "context"

// Service is the base interface that all services must implement.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

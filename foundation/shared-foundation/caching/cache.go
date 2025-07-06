package caching

import (
	"time"

	"github.com/yourusername/shared-foundation/core"
)

// Cache interface for caching
type Cache interface {
	core.Service
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
	Exists(key string) (bool, error)
	Incr(key string) (int64, error)
}

package foundation

import (
	"github.com/yourusername/shared-foundation/caching"
	"github.com/yourusername/shared-foundation/connectrpc"
	"github.com/yourusername/shared-foundation/core"
	"github.com/yourusername/shared-foundation/database"
	"github.com/yourusername/shared-foundation/logging"
	"github.com/yourusername/shared-foundation/messaging"
	"github.com/yourusername/shared-foundation/metrics"
	"github.com/yourusername/shared-foundation/tracing"
)

// Re-export interfaces for convenience
type (
	Service          = core.Service
	Logger           = logging.Logger
	Tracer           = tracing.Tracer
	Metrics          = metrics.Metrics
	Broker           = messaging.Broker
	Cache            = caching.Cache
	Database         = database.Database
	ConnectRPCServer = connectrpc.ConnectRPCServer
)

// ServiceDiscovery interface for service registration and discovery
type ServiceDiscovery interface {
	core.Service
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

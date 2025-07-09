package foundation

import (
	"context"
	"fmt"
	"sync"

	"github.com/yourusername/foundation/connectrpc"
	"github.com/yourusername/foundation/logging"
	"github.com/yourusername/foundation/metrics"
	"github.com/yourusername/foundation/tracing"
)

type Span = tracing.Span
type SpanOption = tracing.SpanOption
type Logger = logging.Logger
type Tracer = tracing.Tracer
type Metrics = metrics.Metrics

// Server represents a lifecycle-managed server component
// (e.g., HTTP servers, gRPC servers) that can be started and stopped.
type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

// App represents the main application with cross-cutting concerns
type App struct {
	name       string
	version    string
	logger     logging.Logger
	tracer     tracing.Tracer
	metrics    metrics.Metrics
	connectRPC *connectrpc.Server

	servers []Server
	ctx     context.Context
	cancel  context.CancelFunc
	mu      sync.Mutex
}

// NewWithConfig returns an App with logger, metrics, tracing, and servers using AppConfig
func NewWithConfig(name, version string, cfg AppConfig) *App {
	ctx, cancel := context.WithCancel(context.Background())
	logger := NewLoggerFromConfig(cfg.Logger)
	metrics := NewMetricsFromConfig(cfg.Metrics)
	tracer := NewTracerFromConfig(cfg.Tracer)

	app := &App{
		name:    name,
		version: version,
		logger:  logger,
		tracer:  tracer,
		metrics: metrics,
		ctx:     ctx,
		cancel:  cancel,
	}

	for _, serverCfg := range cfg.Servers {
		server := createServerFromConfig(serverCfg, logger)
		if server != nil {
			app.AddServer(server)
			if serverCfg.Type == "connectrpc" {
				if connectServer, ok := server.(*connectrpc.Server); ok {
					app.connectRPC = connectServer
				}
			}
		}
	}

	return app
}

// New returns an App with default configuration
func New(name, version string) *App {
	return NewWithConfig(name, version, LoadConfigFromEnv())
}

// AddServer adds a server to the app's lifecycle management
func (a *App) AddServer(server Server) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.servers = append(a.servers, server)
}

// Start starts all servers
func (a *App) Start(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.logger.Info("Starting app", "name", a.name, "version", a.version)
	for _, server := range a.servers {
		if err := server.Start(ctx); err != nil {
			a.logger.Error("Failed to start server", "server", server.Name(), "error", err)
			return fmt.Errorf("failed to start %s: %w", server.Name(), err)
		}
		a.logger.Info("Started server", "server", server.Name())
	}
	a.logger.Info("All servers started successfully")
	return nil
}

// Stop stops all servers gracefully
func (a *App) Stop(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.logger.Info("Stopping app", "name", a.name, "version", a.version)
	a.cancel()
	for i := len(a.servers) - 1; i >= 0; i-- {
		server := a.servers[i]
		if err := server.Stop(ctx); err != nil {
			a.logger.Error("Failed to stop server", "server", server.Name(), "error", err)
		} else {
			a.logger.Info("Stopped server", "server", server.Name())
		}
	}
	a.logger.Info("All servers stopped")
	return nil
}

// Logger returns the logger
func (a *App) Logger() logging.Logger { return a.logger }

// Metrics returns the metrics
func (a *App) Metrics() metrics.Metrics { return a.metrics }

// Tracer returns the tracer
func (a *App) Tracer() tracing.Tracer { return a.tracer }

// Name returns the app name
func (a *App) Name() string { return a.name }

// Version returns the app version
func (a *App) Version() string { return a.version }

// ConnectRPC returns the ConnectRPC server
func (a *App) ConnectRPC() *connectrpc.Server { return a.connectRPC }

// GetServers returns all registered servers
func (a *App) GetServers() []Server {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.servers
}

// GetServerByName returns a server by name
func (a *App) GetServerByName(name string) Server {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, server := range a.servers {
		if server.Name() == name {
			return server
		}
	}
	return nil
}

// createServerFromConfig creates a server based on the configuration
func createServerFromConfig(cfg ServerConfig, logger logging.Logger) Server {
	switch cfg.Type {
	case "connectrpc":
		return connectrpc.NewServer(cfg.Name, cfg.Addr, logger)
	// Add more server types here as needed
	default:
		logger.Error("Unknown server type", "type", cfg.Type)
		return nil
	}
}

// NewLoggerFromConfig creates a logger using LoggerConfig
func NewLoggerFromConfig(cfg LoggerConfig) logging.Logger {
	return logging.NewSlogLogger("configured-logger", cfg.Level, cfg.Format, cfg.Output)
}

// NewMetricsFromConfig creates metrics using MetricsConfig
func NewMetricsFromConfig(cfg MetricsConfig) metrics.Metrics {
	// For now, return default metrics. Extend for prometheus/statsd as needed.
	return metrics.NewDefaultMetrics()
}

// NewTracerFromConfig creates tracer using TracerConfig
func NewTracerFromConfig(cfg TracerConfig) tracing.Tracer {
	// For now, return default tracer. Extend for jaeger/zipkin as needed.
	return tracing.NewDefaultTracer()
}

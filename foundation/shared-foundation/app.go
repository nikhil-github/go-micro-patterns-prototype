package foundation

import (
	"context"
	"fmt"
	"sync"
)

// App represents the main application with all services
type App struct {
	config *Config

	// Core services
	logger           Logger
	tracer           Tracer
	metrics          Metrics
	serviceDiscovery ServiceDiscovery
	broker           Broker
	cache            Cache
	database         Database
	connectRPCServer ConnectRPCServer

	// Lifecycle management
	services []Service
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
}

// New creates a new App with the given name and version
func New(name, version string, options ...AppOption) *App {
	ctx, cancel := context.WithCancel(context.Background())

	// Load configuration
	cfg := Load()
	cfg.App.Name = name
	cfg.App.Version = version

	app := &App{
		config: cfg,
		ctx:    ctx,
		cancel: cancel,
	}

	// Apply options
	for _, option := range options {
		option(app)
	}

	// Set sensible defaults if not provided
	if app.logger == nil {
		app.logger = DefaultLogger()
	}
	if app.tracer == nil {
		app.tracer = DefaultTracer()
	}
	if app.metrics == nil {
		app.metrics = DefaultMetrics()
	}
	if app.broker == nil {
		app.broker = DefaultBroker()
	}
	if app.cache == nil {
		app.cache = DefaultCache()
	}
	if app.database == nil {
		app.database = DefaultDatabase()
	}
	if app.serviceDiscovery == nil {
		app.serviceDiscovery = DefaultServiceDiscovery()
	}
	if app.connectRPCServer == nil {
		app.connectRPCServer = DefaultConnectRPCServer()
	}

	return app
}

// Init initializes all services (now handled by functional options)
func (a *App) Init() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Add all services to the services slice for lifecycle management
	a.services = []Service{
		a.logger,
		a.tracer,
		a.metrics,
		a.serviceDiscovery,
		a.broker,
		a.cache,
		a.database,
		a.connectRPCServer,
	}

	a.logger.Info("App initialized successfully",
		"name", a.config.App.Name,
		"version", a.config.App.Version,
		"services", len(a.services))

	return nil
}

// Start starts all services
func (a *App) Start(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.logger.Info("Starting app", "name", a.config.App.Name, "version", a.config.App.Version)

	for _, service := range a.services {
		if err := service.Start(ctx); err != nil {
			a.logger.Error("Failed to start service", "service", service.Name(), "error", err)
			return fmt.Errorf("failed to start %s: %w", service.Name(), err)
		}
		a.logger.Info("Started service", "service", service.Name())
	}

	a.logger.Info("All services started successfully")
	return nil
}

// Stop stops all services gracefully
func (a *App) Stop(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.logger.Info("Stopping app", "name", a.config.App.Name)
	a.cancel()

	// Stop services in reverse order
	for i := len(a.services) - 1; i >= 0; i-- {
		service := a.services[i]
		if err := service.Stop(ctx); err != nil {
			a.logger.Error("Failed to stop service", "service", service.Name(), "error", err)
		} else {
			a.logger.Info("Stopped service", "service", service.Name())
		}
	}

	a.logger.Info("All services stopped")
	return nil
}

// Name returns the app name
func (a *App) Name() string {
	return a.config.App.Name
}

// Service getters
func (a *App) Logger() Logger                     { return a.logger }
func (a *App) Tracer() Tracer                     { return a.tracer }
func (a *App) Metrics() Metrics                   { return a.metrics }
func (a *App) ServiceDiscovery() ServiceDiscovery { return a.serviceDiscovery }
func (a *App) Broker() Broker                     { return a.broker }
func (a *App) Cache() Cache                       { return a.cache }
func (a *App) Database() Database                 { return a.database }
func (a *App) ConnectRPCServer() ConnectRPCServer { return a.connectRPCServer }

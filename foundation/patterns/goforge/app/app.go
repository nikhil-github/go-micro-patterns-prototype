package goforge

import (
	"context"
	"fmt"
	"sync"

	"github.com/yourusername/foundation/patterns/goforge/config"
	"github.com/yourusername/foundation/patterns/goforge/logger"
	"github.com/yourusername/foundation/patterns/goforge/services"
)

// Import types from main package
type (
	Service          interface{}
	Logger           interface{}
	Tracer           interface{}
	Metrics          interface{}
	ServiceDiscovery interface{}
	Broker           interface{}
	Cache            interface{}
	Database         interface{}
	ConnectRPCServer interface{}
)

// App represents the main application with all services
type App struct {
	config *config.Config

	// Core services
	logger           logger.Logger
	tracer           services.Tracer
	metrics          services.Metrics
	serviceDiscovery services.ServiceDiscovery
	broker           services.Broker
	cache            services.Cache
	database         services.Database
	connectRPCServer services.ConnectRPCServer

	// Lifecycle management
	services []services.Service
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
}

// New creates a new App with the given name and version
func New(name, version string) *App {
	ctx, cancel := context.WithCancel(context.Background())

	// Load configuration
	cfg := config.Load()
	cfg.App.Name = name
	cfg.App.Version = version

	return &App{
		config: cfg,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Init initializes all services based on configuration
func (a *App) Init() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Create service factories
	loggerFactory := services.NewLoggerFactory()
	tracerFactory := services.NewTracerFactory()
	metricsFactory := services.NewMetricsFactory()
	discoveryFactory := services.NewServiceDiscoveryFactory()
	brokerFactory := services.NewBrokerFactory()
	cacheFactory := services.NewCacheFactory()
	dbFactory := services.NewDatabaseFactory()
	connectFactory := services.NewConnectRPCServerFactory()

	// Initialize services in dependency order
	var err error

	// 1. Logger (no dependencies)
	a.logger, err = loggerFactory.Create(a.config.Logger)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	a.services = append(a.services, a.logger)

	// 2. Tracer (depends on logger)
	a.tracer, err = tracerFactory.Create(a.config.Tracer, a.logger)
	if err != nil {
		return fmt.Errorf("failed to create tracer: %w", err)
	}
	a.services = append(a.services, a.tracer)

	// 3. Metrics (depends on logger)
	a.metrics, err = metricsFactory.Create(a.config.Metrics, a.logger)
	if err != nil {
		return fmt.Errorf("failed to create metrics: %w", err)
	}
	a.services = append(a.services, a.metrics)

	// 4. Service Discovery (depends on logger)
	a.serviceDiscovery, err = discoveryFactory.Create(a.config.ServiceDiscovery, a.logger)
	if err != nil {
		return fmt.Errorf("failed to create service discovery: %w", err)
	}
	a.services = append(a.services, a.serviceDiscovery)

	// 5. Broker (depends on logger, tracer)
	a.broker, err = brokerFactory.Create(a.config.Broker, a.logger, a.tracer)
	if err != nil {
		return fmt.Errorf("failed to create broker: %w", err)
	}
	a.services = append(a.services, a.broker)

	// 6. Cache (depends on logger, tracer)
	a.cache, err = cacheFactory.Create(a.config.Cache, a.logger, a.tracer)
	if err != nil {
		return fmt.Errorf("failed to create cache: %w", err)
	}
	a.services = append(a.services, a.cache)

	// 7. Database (depends on logger, tracer)
	a.database, err = dbFactory.Create(a.config.Database, a.logger, a.tracer)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	a.services = append(a.services, a.database)

	// 8. ConnectRPC Server (depends on logger, tracer, metrics)
	a.connectRPCServer, err = connectFactory.Create(a.config.ConnectRPC, a.logger, a.tracer, a.metrics)
	if err != nil {
		return fmt.Errorf("failed to create connectRPC server: %w", err)
	}
	a.services = append(a.services, a.connectRPCServer)

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
func (a *App) Logger() logger.Logger                       { return a.logger }
func (a *App) Tracer() services.Tracer                     { return a.tracer }
func (a *App) Metrics() services.Metrics                   { return a.metrics }
func (a *App) ServiceDiscovery() services.ServiceDiscovery { return a.serviceDiscovery }
func (a *App) Broker() services.Broker                     { return a.broker }
func (a *App) Cache() services.Cache                       { return a.cache }
func (a *App) Database() services.Database                 { return a.database }
func (a *App) ConnectRPCServer() services.ConnectRPCServer { return a.connectRPCServer }

package servicefactory

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"connectrpc.com/connect"
	"github.com/spf13/viper"
	"github.com/yourusername/foundation/patterns/current/connectrpc"
)

// Service is the interface that all services must implement.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// ServiceFactory encapsulates logic for creating services.
type ServiceFactory struct {
	ctx      context.Context
	cancel   context.CancelFunc
	logger   *slog.Logger
	services []Service
	mu       sync.Mutex
}

// NewServiceFactory creates a new ServiceFactory with default logger
func NewServiceFactory() *ServiceFactory {
	ctx, cancel := context.WithCancel(context.Background())
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &ServiceFactory{
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

// NewServiceFactoryWithLogger creates a new ServiceFactory with custom logger
func NewServiceFactoryWithLogger(logger *slog.Logger) *ServiceFactory {
	ctx, cancel := context.WithCancel(context.Background())
	return &ServiceFactory{
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

// Init initializes the factory and loads configuration
func (sf *ServiceFactory) Init() {
	viper.AutomaticEnv()
	sf.logger.Info("ServiceFactory initialized")
}

// AddConnectRPCServer adds a connectRPC server with auto-loaded config
func (sf *ServiceFactory) AddConnectRPCServer(handler *connect.Handler) error {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	cfg := connectrpc.LoadConfig()
	sf.logger.Info("Adding connectRPC server", "address", cfg.Address)

	service := connectrpc.NewServer(cfg, handler, sf.logger)
	sf.services = append(sf.services, service)
	return nil
}

// AddConnectRPCServerWithConfig adds a connectRPC server with custom config
func (sf *ServiceFactory) AddConnectRPCServerWithConfig(cfg connectrpc.Config, handler *connect.Handler) error {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Adding connectRPC server with custom config", "address", cfg.Address)

	service := connectrpc.NewServer(cfg, handler, sf.logger)
	sf.services = append(sf.services, service)
	return nil
}

// GetLogger returns the logger instance
func (sf *ServiceFactory) GetLogger() *slog.Logger {
	return sf.logger
}

// GetContext returns the service context
func (sf *ServiceFactory) GetContext() context.Context {
	return sf.ctx
}

// Start starts all registered services
func (sf *ServiceFactory) Start() error {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Starting all services", "count", len(sf.services))
	for _, service := range sf.services {
		if err := service.Start(sf.ctx); err != nil {
			sf.logger.Error("Failed to start service", "error", err)
			return err
		}
	}
	return nil
}

// Stop stops all registered services gracefully
func (sf *ServiceFactory) Stop() error {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Stopping all services", "count", len(sf.services))
	sf.cancel()

	var wg sync.WaitGroup
	errors := make(chan error, len(sf.services))

	for _, service := range sf.services {
		wg.Add(1)
		go func(s Service) {
			defer wg.Done()
			if err := s.Stop(sf.ctx); err != nil {
				errors <- err
			}
		}(service)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			sf.logger.Error("Error during service shutdown", "error", err)
			return err
		}
	}

	return nil
}

// Run starts all services and waits for shutdown signal
func (sf *ServiceFactory) Run() error {
	if err := sf.Start(); err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sf.logger.Info("Waiting for shutdown signal...")
	<-sigChan

	sf.logger.Info("Shutdown signal received, stopping services...")
	return sf.Stop()
}

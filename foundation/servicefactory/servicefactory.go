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
	"github.com/yourusername/foundation/connectrpc"
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

// Build initializes the global config and loads all service configs
func (sf *ServiceFactory) Build() {
	viper.AutomaticEnv()
	sf.logger.Info("ServiceFactory bootstrapped")
}

// GetContext returns the service context for use by the microservice
func (sf *ServiceFactory) GetContext() context.Context {
	return sf.ctx
}

// CreateConnectRPCServer creates a connectRPC server with config loaded from env
func (sf *ServiceFactory) CreateConnectRPCServer(handler *connect.Handler) (Service, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	cfg := connectrpc.LoadConfig()
	sf.logger.Info("Creating connectRPC server", "address", cfg.Address)
	service := connectrpc.NewServer(cfg, handler, sf.logger)
	sf.services = append(sf.services, service)
	return service, nil
}

// CreateConnectRPCServerWithConfig creates a connectRPC server with custom config
func (sf *ServiceFactory) CreateConnectRPCServerWithConfig(cfg connectrpc.Config, handler *connect.Handler) (Service, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Creating connectRPC server with custom config", "address", cfg.Address)
	service := connectrpc.NewServer(cfg, handler, sf.logger)
	sf.services = append(sf.services, service)
	return service, nil
}

// GetConnectRPCConfig returns the current connectRPC config (useful for inspection/modification)
func (sf *ServiceFactory) GetConnectRPCConfig() connectrpc.Config {
	return connectrpc.LoadConfig()
}

// GetLogger returns the logger instance for external use
func (sf *ServiceFactory) GetLogger() *slog.Logger {
	return sf.logger
}

// StartAll starts all registered services
func (sf *ServiceFactory) StartAll() error {
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

// StopAll stops all registered services gracefully
func (sf *ServiceFactory) StopAll() error {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Stopping all services", "count", len(sf.services))
	sf.cancel() // Cancel the context to signal shutdown

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

	// Check for any errors during shutdown
	for err := range errors {
		if err != nil {
			sf.logger.Error("Error during service shutdown", "error", err)
			return err
		}
	}

	return nil
}

// WaitForShutdown waits for shutdown signals and gracefully stops all services
func (sf *ServiceFactory) WaitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sf.logger.Info("Waiting for shutdown signal...")
	<-sigChan

	sf.logger.Info("Shutdown signal received, stopping services...")
	if err := sf.StopAll(); err != nil {
		sf.logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}
	sf.logger.Info("All services stopped gracefully")
}

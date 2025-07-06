package servicefactory

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"connectrpc.com/connect"
	"github.com/yourusername/foundation/patterns/simplified/config"
)

// Service is the interface that all services must implement.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// ServiceFactory focuses only on service creation and lifecycle
type ServiceFactory struct {
	config   *config.Config
	logger   *slog.Logger
	services []Service
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
}

// New creates a new ServiceFactory with explicit dependencies
func New(cfg *config.Config, logger *slog.Logger) *ServiceFactory {
	ctx, cancel := context.WithCancel(context.Background())
	return &ServiceFactory{
		config: cfg,
		logger: logger,
		ctx:    ctx,
		cancel: cancel,
	}
}

// AddConnectRPCServer adds a connectRPC server with default config
func (sf *ServiceFactory) AddConnectRPCServer(handler *connect.Handler) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Adding connectRPC server", "address", sf.config.ConnectRPC.Address)
	service := NewConnectRPCServer(sf.config.ConnectRPC, handler, sf.logger)
	sf.services = append(sf.services, service)
}

// AddConnectRPCServerWithConfig adds a connectRPC server with custom config
func (sf *ServiceFactory) AddConnectRPCServerWithConfig(cfg config.ConnectRPC, handler *connect.Handler) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.logger.Info("Adding connectRPC server with custom config", "address", cfg.Address)
	service := NewConnectRPCServer(cfg, handler, sf.logger)
	sf.services = append(sf.services, service)
}

// Run starts all services and waits for shutdown signal
func (sf *ServiceFactory) Run() error {
	// Start all services
	if err := sf.start(); err != nil {
		return err
	}

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sf.logger.Info("Waiting for shutdown signal...")
	<-sigChan

	sf.logger.Info("Shutdown signal received, stopping services...")
	return sf.stop()
}

// start starts all registered services
func (sf *ServiceFactory) start() error {
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

// stop stops all registered services gracefully
func (sf *ServiceFactory) stop() error {
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

// ConnectRPCServer implementation
type ConnectRPCServer struct {
	config  config.ConnectRPC
	handler *connect.Handler
	logger  *slog.Logger
}

func NewConnectRPCServer(cfg config.ConnectRPC, handler *connect.Handler, logger *slog.Logger) Service {
	return &ConnectRPCServer{
		config:  cfg,
		handler: handler,
		logger:  logger,
	}
}

func (s *ConnectRPCServer) Start(ctx context.Context) error {
	s.logger.Info("Starting connectRPC server", "address", s.config.Address)
	// TODO: Implement actual server start logic
	return nil
}

func (s *ConnectRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping connectRPC server", "address", s.config.Address)
	// TODO: Implement actual server stop logic
	return nil
}

func NewDummyHandler() *connect.Handler {
	return connect.NewUnaryHandler("/test.TestService/TestMethod", func(ctx context.Context, req *connect.Request[any]) (*connect.Response[any], error) {
		return connect.NewResponse[any](nil), nil
	})
}

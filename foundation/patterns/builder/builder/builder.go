package builder

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"connectrpc.com/connect"
	"github.com/yourusername/foundation/patterns/builder/config"
	"github.com/yourusername/foundation/patterns/builder/services"
)

type Microservice struct {
	config   *config.Config
	logger   *slog.Logger
	services []services.Service
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
}

type Builder struct {
	config *config.Config
	logger *slog.Logger
}

func NewMicroservice() *Builder {
	return &Builder{
		config: config.New(),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (b *Builder) WithLogger(logger *slog.Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) WithConfig(cfg *config.Config) *Builder {
	b.config = cfg
	return b
}

func (b *Builder) AddConnectRPCServer(handler *connect.Handler) *Builder {
	service := services.NewConnectRPCServer(b.config.ConnectRPC, handler, b.logger)
	b.config.Services = append(b.config.Services, service)
	return b
}

func (b *Builder) Build() *Microservice {
	ctx, cancel := context.WithCancel(context.Background())
	return &Microservice{
		config:   b.config,
		logger:   b.logger,
		services: b.config.Services,
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (m *Microservice) Start() *Microservice {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.logger.Info("Starting all services", "count", len(m.services))
	for _, service := range m.services {
		if err := service.Start(m.ctx); err != nil {
			m.logger.Error("Failed to start service", "error", err)
			return m
		}
	}
	return m
}

func (m *Microservice) WaitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	m.logger.Info("Waiting for shutdown signal...")
	<-sigChan

	m.logger.Info("Shutdown signal received, stopping services...")
	m.StopAll()
}

func (m *Microservice) StopAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.logger.Info("Stopping all services", "count", len(m.services))
	m.cancel()

	var wg sync.WaitGroup
	for _, service := range m.services {
		wg.Add(1)
		go func(s services.Service) {
			defer wg.Done()
			if err := s.Stop(m.ctx); err != nil {
				m.logger.Error("Error during service shutdown", "error", err)
			}
		}(service)
	}

	wg.Wait()
	m.logger.Info("All services stopped gracefully")
}

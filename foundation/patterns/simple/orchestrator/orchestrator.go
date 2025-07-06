package orchestrator

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yourusername/foundation/patterns/simple/services"
)

type Orchestrator struct {
	services []services.Service
	logger   *slog.Logger
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
}

func New(logger *slog.Logger) *Orchestrator {
	ctx, cancel := context.WithCancel(context.Background())
	return &Orchestrator{
		logger: logger,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (o *Orchestrator) Add(service services.Service) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.services = append(o.services, service)
}

func (o *Orchestrator) Start() error {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.logger.Info("Starting all services", "count", len(o.services))
	for _, service := range o.services {
		if err := service.Start(o.ctx); err != nil {
			o.logger.Error("Failed to start service", "error", err)
			return err
		}
	}
	return nil
}

func (o *Orchestrator) WaitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	o.logger.Info("Waiting for shutdown signal...")
	<-sigChan

	o.logger.Info("Shutdown signal received, stopping services...")
	o.Stop()
}

func (o *Orchestrator) Stop() {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.logger.Info("Stopping all services", "count", len(o.services))
	o.cancel()

	var wg sync.WaitGroup
	for _, service := range o.services {
		wg.Add(1)
		go func(s services.Service) {
			defer wg.Done()
			if err := s.Stop(o.ctx); err != nil {
				o.logger.Error("Error during service shutdown", "error", err)
			}
		}(service)
	}

	wg.Wait()
	o.logger.Info("All services stopped gracefully")
}

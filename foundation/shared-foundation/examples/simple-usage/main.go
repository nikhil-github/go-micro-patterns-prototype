package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	foundation "github.com/yourusername/shared-foundation"
)

func main() {
	// Simple usage with functional options - much more readable!
	app := foundation.New("user-service", "1.0.0",
		foundation.WithSlogLogger(foundation.LoggerConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		}),
		foundation.WithMockTracer(),
		foundation.WithMockMetrics(),
		foundation.WithMockBroker(),
		foundation.WithMockCache(),
		foundation.WithMockDatabase(),
		foundation.WithMockServiceDiscovery(),
		foundation.WithMockConnectRPCServer(),
	)

	// Initialize (now just sets up lifecycle management)
	if err := app.Init(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start all services
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}

	// Get injected dependencies
	logger := app.Logger()
	logger.Info("User service started successfully")

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Info("Shutdown signal received, stopping app...")
	if err := app.Stop(ctx); err != nil {
		logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}
}

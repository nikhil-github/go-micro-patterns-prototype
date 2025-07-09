package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	foundation "github.com/yourusername/foundation"
)

func main() {
	// Create app with name and version
	application := foundation.New("demo-service", "1.0.0")

	// Initialize all services
	if err := application.Init(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start all services
	if err := application.Start(ctx); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}

	// Get logger for demonstration
	logger := application.Logger()
	logger.Info("Demo service started successfully")

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	logger.Info("Shutdown signal received, stopping app...")
	if err := application.Stop(ctx); err != nil {
		logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}
}

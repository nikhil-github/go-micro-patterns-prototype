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
	// Set environment variables for defaults (optional)
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("LOG_FORMAT", "json")
	os.Setenv("LOG_OUTPUT", "stdout")

	// Create app with no options - uses all defaults!
	app := foundation.New("user-service", "1.0.0", foundation.WithSlogLogger(foundation.LoggerConfig{
		Level:  "debug",
		Format: "json",
		Output: "stdout",
	}))

	// Initialize (sets up lifecycle management)
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

	// Get the default logger (slog with env config)
	logger := app.Logger()
	logger.Info("User service started with defaults",
		"logger", logger.Name(),
		"tracer", app.Tracer().Name(),
		"metrics", app.Metrics().Name(),
		"broker", app.Broker().Name(),
		"cache", app.Cache().Name(),
		"database", app.Database().Name(),
	)

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

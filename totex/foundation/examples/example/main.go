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
	// Set environment variables for server configuration
	os.Setenv("SERVER_NAME", "example-server")
	os.Setenv("SERVER_ADDR", ":8080")

	// Create app with name and version (automatically creates ConnectRPC server)
	app := foundation.New("example-service", "1.0.0")

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get injected dependencies for business logic
	logger := app.Logger()
	tracer := app.Tracer()
	metrics := app.Metrics()

	// Get the automatically created ConnectRPC server
	connectServer := app.ConnectRPC()
	if connectServer == nil {
		log.Fatal("ConnectRPC server not found")
	}

	// Example: Create business logic with injected dependencies
	exampleHandler := &ExampleHandler{
		Logger:  logger,
		Tracer:  tracer,
		Metrics: metrics,
	}

	// Example: Register ConnectRPC handlers
	if err := connectServer.RegisterHandler("/example.ExampleService/", exampleHandler); err != nil {
		logger.Error("Failed to register ConnectRPC handler", "error", err)
	}

	// Start all servers
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Example service started successfully")
	<-sigChan

	logger.Info("Shutdown signal received, stopping app...")
	if err := app.Stop(ctx); err != nil {
		logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}
}

// ExampleHandler demonstrates how to use the injected dependencies
type ExampleHandler struct {
	Logger  foundation.Logger
	Tracer  foundation.Tracer
	Metrics foundation.Metrics
}

// Example method showing how to use the dependencies
func (h *ExampleHandler) ProcessRequest(ctx context.Context, requestID string, data string) error {
	// Start a trace span
	span := h.Tracer.StartSpan("process_request")
	defer span.Finish()

	// Record metrics
	h.Metrics.Counter("request_processed", 1, "service", "example-service")

	// Log the operation
	h.Logger.Info("Processing request", "request_id", requestID, "data", data)

	h.Logger.Info("Request processed successfully", "request_id", requestID)
	return nil
}

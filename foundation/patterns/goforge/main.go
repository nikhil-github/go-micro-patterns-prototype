package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/foundation/patterns/goforge"
)

func main() {
	// Create app with name and version
	application := goforge.New("user-service", "1.0.0")

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

	// Example: Use the services
	logger := application.Logger()
	_ = application.Tracer() // Available for use
	metrics := application.Metrics()
	broker := application.Broker()
	cache := application.Cache()
	_ = application.Database()         // Available for use
	_ = application.ConnectRPCServer() // Available for use

	// Example: Register ConnectRPC handlers
	// handler := &UserServiceHandler{
	//     logger: logger,
	//     tracer: tracer,
	//     metrics: metrics,
	//     broker: broker,
	//     cache: cache,
	//     database: database,
	// }
	// connectServer.RegisterHandler("/user.UserService/", handler)

	// Example: Use other services
	logger.Info("User service started successfully")

	// Publish a message
	broker.Publish("user.created", []byte(`{"user_id": "123"}`))

	// Cache some data
	cache.Set("user:123", []byte(`{"name": "John"}`), 0)

	// Record metrics
	metrics.Counter("user_requests_total", 1, "service", "user-service")

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Waiting for shutdown signal...")
	<-sigChan

	logger.Info("Shutdown signal received, stopping app...")

	// Stop all services gracefully
	if err := application.Stop(ctx); err != nil {
		logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("App stopped gracefully")
}

// Example service handler (commented out for demonstration)
/*
type UserServiceHandler struct {
	logger   interfaces.Logger
	tracer   interfaces.Tracer
	metrics  interfaces.Metrics
	broker   interfaces.Broker
	cache    interfaces.Cache
	database interfaces.Database
}

func (h *UserServiceHandler) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	span := h.tracer.StartSpan("create_user")
	defer span.Finish()

	h.metrics.Counter("user_creation_requests", 1)

	// Business logic here...

	return &CreateUserResponse{UserId: "123"}, nil
}
*/

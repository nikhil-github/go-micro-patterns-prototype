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
	application := foundation.New("user-service", "1.0.0")

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

	// Get injected dependencies for business logic
	logger := application.Logger()
	tracer := application.Tracer()
	metrics := application.Metrics()
	broker := application.Broker()
	cache := application.Cache()
	database := application.Database()
	connectServer := application.ConnectRPCServer()

	// Example: Create business logic with injected dependencies
	userHandler := &UserHandler{
		Logger:   logger,
		Tracer:   tracer,
		Metrics:  metrics,
		Broker:   broker,
		Cache:    cache,
		Database: database,
	}

	// Example: Register ConnectRPC handlers
	if err := connectServer.RegisterHandler("/user.UserService/", userHandler); err != nil {
		logger.Error("Failed to register ConnectRPC handler", "error", err)
	}

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("User service started successfully")
	<-sigChan

	logger.Info("Shutdown signal received, stopping app...")
	if err := application.Stop(ctx); err != nil {
		logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}
}

// UserHandler demonstrates how to use the injected dependencies
type UserHandler struct {
	Logger   foundation.Logger
	Tracer   foundation.Tracer
	Metrics  foundation.Metrics
	Broker   foundation.Broker
	Cache    foundation.Cache
	Database foundation.Database
}

// Example method showing how to use the dependencies
func (h *UserHandler) CreateUser(ctx context.Context, userID string, name string) error {
	// Start a trace span
	span := h.Tracer.StartSpan("create_user")
	defer span.Finish()

	// Record metrics
	h.Metrics.Counter("user_creation_requests", 1, "service", "user-service")

	// Log the operation
	h.Logger.Info("Creating user", "user_id", userID, "name", name)

	// Example: Cache user data
	if err := h.Cache.Set("user:"+userID, []byte(name), 0); err != nil {
		h.Logger.Error("Failed to cache user data", "error", err)
	}

	// Example: Publish event
	eventData := []byte(`{"user_id": "` + userID + `", "name": "` + name + `"}`)
	if err := h.Broker.Publish("user.created", eventData); err != nil {
		h.Logger.Error("Failed to publish user created event", "error", err)
	}

	h.Logger.Info("User created successfully", "user_id", userID)
	return nil
}

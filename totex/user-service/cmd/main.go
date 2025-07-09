package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	foundation "github.com/yourusername/foundation"
	userv1connect "github.com/yourusername/schema/gen/user/v1/userv1connect"
	"github.com/yourusername/user-service/internal/user"
)

func main() {
	// Set environment variables for server configuration
	os.Setenv("SERVER_NAME", "user-service-server")
	os.Setenv("SERVER_ADDR", ":8080")

	// Create app with name and version (automatically creates ConnectRPC server)
	app := foundation.New("user-service", "1.0.0")
	logger := app.Logger()

	// User service implementation (no DB)
	userSvc := user.NewService()
	path, handler := userv1connect.NewUserServiceHandler(userSvc)

	// Get the automatically created ConnectRPC server
	connectServer := app.ConnectRPC()
	if connectServer == nil {
		logger.Error("ConnectRPC server not found")
		os.Exit(1)
	}

	// Register handler with the auto-created server
	if err := connectServer.RegisterHandler(path, handler); err != nil {
		logger.Error("Failed to register ConnectRPC handler", "error", err)
		os.Exit(1)
	}

	// Start all servers (including HTTP server)
	if err := app.Start(context.Background()); err != nil {
		logger.Error("Failed to start app", "error", err)
		os.Exit(1)
	}

	logger.Info("User service started successfully")

	// Graceful shutdown on SIGINT/SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	logger.Info("Shutdown signal received", "signal", sig)
	if err := app.Stop(context.Background()); err != nil {
		logger.Error("Error during shutdown", "error", err)
		os.Exit(1)
	}
	logger.Info("Shutdown complete")
}

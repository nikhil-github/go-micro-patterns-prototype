package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	foundation "github.com/yourusername/foundation"
	"github.com/yourusername/foundation/connectrpc"
	userv1connect "github.com/yourusername/schema/gen/user/v1/userv1connect"
	"github.com/yourusername/user-service/internal/user"
)

func main() {
	app := foundation.New()
	logger := app.Logger()

	// Create ConnectRPC server as a separate server
	server := connectrpc.NewServer("user-service-server", ":8080", logger)

	// User service implementation (no DB)
	userSvc := user.NewService()
	path, handler := userv1connect.NewUserServiceHandler(userSvc)

	// Register handler with the server
	if err := server.RegisterHandler(path, handler); err != nil {
		logger.Error("Failed to register ConnectRPC handler", "error", err)
		os.Exit(1)
	}

	// Add the server to the app's lifecycle management
	app.AddServer(server)

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

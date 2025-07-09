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
	app := foundation.New()
	logger := app.Logger()

	// User service implementation (no DB)
	userSvc := user.NewService()
	path, handler := userv1connect.NewUserServiceHandler(userSvc)

	// Register handler using the ConnectRPCServer interface directly
	server := app.ConnectRPCServer()
	if err := server.RegisterHandler(path, handler); err != nil {
		logger.Error("Failed to register ConnectRPC handler", "error", err)
		os.Exit(1)
	}

	// Start all services (including HTTP server)
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

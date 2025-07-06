package main

import (
	"log/slog"

	"github.com/yourusername/foundation/patterns/separated/config"
	"github.com/yourusername/foundation/patterns/separated/lifecycle"
	"github.com/yourusername/foundation/patterns/separated/services"
)

func main() {
	logger := slog.Default()

	// Config management
	cfg := config.New().Load()

	// Service creation
	factory := services.NewFactory(logger)
	service := factory.CreateConnectRPCServer(cfg.ConnectRPC, services.NewDummyHandler())

	// Lifecycle management
	orchestrator := lifecycle.NewOrchestrator(logger)
	orchestrator.Add(service)

	if err := orchestrator.Start(); err != nil {
		logger.Error("Failed to start services", "error", err)
		return
	}

	// Wait for shutdown signal
	orchestrator.WaitForShutdown()
}

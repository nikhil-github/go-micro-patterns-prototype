package main

import (
	"log/slog"

	"github.com/yourusername/foundation/patterns/simple/orchestrator"
	"github.com/yourusername/foundation/patterns/simple/services"
)

func main() {
	logger := slog.Default()

	// Create service directly (no factory needed)
	service := services.NewConnectRPCServer(":8080", services.NewDummyHandler(), logger)

	// Simple orchestrator for lifecycle
	orchestrator := orchestrator.New(logger)
	orchestrator.Add(service)

	if err := orchestrator.Start(); err != nil {
		logger.Error("Failed to start services", "error", err)
		return
	}

	// Wait for shutdown signal
	orchestrator.WaitForShutdown()
}

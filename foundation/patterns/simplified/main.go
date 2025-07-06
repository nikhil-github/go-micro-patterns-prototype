package main

import (
	"github.com/yourusername/foundation/patterns/simplified/config"
	"github.com/yourusername/foundation/patterns/simplified/servicefactory"
)

func main() {
	// Load config (separated concern)
	cfg := config.Load()

	// Create logger from config
	logger := cfg.CreateLogger()

	// Create factory with explicit dependencies
	factory := servicefactory.New(cfg, logger)

	// Add services explicitly
	handler := servicefactory.NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	// Add another service with custom config
	customCfg := config.ConnectRPC{Address: ":9090"}
	factory.AddConnectRPCServerWithConfig(customCfg, handler)

	// Run everything (start + wait for shutdown)
	if err := factory.Run(); err != nil {
		logger.Error("Failed to run services", "error", err)
		return
	}
}

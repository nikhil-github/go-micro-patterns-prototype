package main

import (
	"github.com/yourusername/foundation/connectrpc"
	"github.com/yourusername/foundation/servicefactory"
)

func main() {
	factory := servicefactory.NewServiceFactory()

	// Build the factory (loads all configs from env)
	factory.Build()

	// Get the service context for use by the microservice
	ctx := factory.GetContext()

	// Example 1: Create connectRPC server with auto-loaded config
	handler := connectrpc.NewDummyHandler()
	_, err := factory.CreateConnectRPCServer(handler)
	if err != nil {
		factory.GetLogger().Error("Failed to create connectRPC server", "err", err)
		return
	}

	// Example 2: Create connectRPC server with custom config override
	customCfg := factory.GetConnectRPCConfig()
	customCfg.Address = ":9090" // Override the address

	_, err = factory.CreateConnectRPCServerWithConfig(customCfg, handler)
	if err != nil {
		factory.GetLogger().Error("Failed to create connectRPC server with custom config", "err", err)
		return
	}

	// Start all services
	if err := factory.StartAll(); err != nil {
		factory.GetLogger().Error("Failed to start services", "err", err)
		return
	}

	// Example: Use the factory context for microservice operations
	go func() {
		<-ctx.Done()
		factory.GetLogger().Info("Context cancelled, shutting down...")
	}()

	// Wait for shutdown signal and gracefully stop all services
	factory.WaitForShutdown()
}

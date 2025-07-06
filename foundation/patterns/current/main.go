package main

import (
	"github.com/yourusername/foundation/patterns/current/connectrpc"
	"github.com/yourusername/foundation/patterns/current/servicefactory"
)

func main() {
	factory := servicefactory.NewServiceFactory()

	// Initialize the factory
	factory.Init()

	// Get the service context for use by the microservice
	ctx := factory.GetContext()

	// Add connectRPC server with auto-loaded config
	handler := connectrpc.NewDummyHandler()
	if err := factory.AddConnectRPCServer(handler); err != nil {
		factory.GetLogger().Error("Failed to add connectRPC server", "err", err)
		return
	}

	// Add connectRPC server with custom config override
	customCfg := connectrpc.Config{Address: ":9090"}
	if err := factory.AddConnectRPCServerWithConfig(customCfg, handler); err != nil {
		factory.GetLogger().Error("Failed to add connectRPC server with custom config", "err", err)
		return
	}

	// Example: Use the factory context for microservice operations
	go func() {
		<-ctx.Done()
		factory.GetLogger().Info("Context cancelled, shutting down...")
	}()

	// Run the factory (starts services and waits for shutdown)
	if err := factory.Run(); err != nil {
		factory.GetLogger().Error("Failed to run services", "err", err)
		return
	}
}

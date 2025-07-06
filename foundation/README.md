# Service Factory Library

This Go library provides a comprehensive, thread-safe ServiceFactory for bootstrapping infrastructure services in a microservices architecture. The ServiceFactory automatically manages logger initialization, configuration loading (using viper), context management, and service lifecycle with graceful shutdown.

## Features
- Comprehensive ServiceFactory that handles logger initialization, config loading, context management, and service lifecycle
- Automatic environment variable configuration via viper
- Interface-based service abstraction for extensibility
- Factory methods for connectRPC server with both auto-config and custom config options
- Thread-safe, dependency-injection friendly design
- Built-in context management for microservice lifecycle
- Graceful shutdown with signal handling
- Logging via log/slog
- Example usage in `main.go`
- Unit tests using `testing`

## Usage Example

```
go run main.go
```

Example `main.go`:
```go
package main

import (
	"context"

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
```

## Service Lifecycle

The ServiceFactory manages the complete lifecycle of all services:

1. **Creation**: Services are created and registered with the factory
2. **Starting**: `StartAll()` starts all registered services with the factory's context
3. **Running**: Services run until shutdown is requested
4. **Shutdown**: `StopAll()` gracefully stops all services when context is cancelled
5. **Signal Handling**: `WaitForShutdown()` listens for SIGINT/SIGTERM and initiates graceful shutdown

## Context Management

The factory provides a shared context for all services:
- `GetContext()` returns the factory's context for use by the microservice
- Context cancellation triggers graceful shutdown of all services
- All services receive the same context for coordinated lifecycle management

## Custom Logger

If you need a custom logger configuration:

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
factory := servicefactory.NewServiceFactoryWithLogger(logger)
```

## Environment Variables

The factory automatically loads configuration from environment variables:
- `CONNECTRPC_ADDRESS`: ConnectRPC server address (default: ":8080")

## Testing

Run unit tests:
```
go test ./servicefactory -v
```

## Dependencies
- [connectrpc.com/connect](https://connectrpc.com/)
- [github.com/spf13/viper](https://github.com/spf13/viper)
- [log/slog](https://pkg.go.dev/log/slog) 
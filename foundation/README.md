# Foundation - Go Microservices Shared Library

This repository contains a Go-idiomatic shared library for microservices with infrastructure components. The library provides a comprehensive foundation for building microservices with standardized logging, tracing, metrics, messaging, caching, and database access.

## Structure

The library is organized in a Go-idiomatic structure:

```
foundation/
├── shared-foundation/              # Main shared library
│   ├── go.mod
│   ├── README.md
│   ├── types.go                    # All interfaces and types
│   ├── config.go                   # Configuration management
│   ├── app.go                      # Main App orchestrator
│   ├── factory.go                  # Service factories
│   ├── mocks.go                    # Mock implementations
│   ├── logger/                     # Logger implementations
│   ├── tracer/                     # Tracer implementations
│   ├── metrics/                    # Metrics implementations
│   ├── broker/                     # Message broker implementations
│   ├── cache/                      # Cache implementations
│   ├── database/                   # Database implementations
│   ├── connectrpc/                 # ConnectRPC implementations
│   └── examples/                   # Usage examples
└── README.md                       # This file
```

## Features

- **Go-Idiomatic Structure**: Follows standard Go conventions and patterns
- **Interface-Based Design**: Clean abstractions for all infrastructure components
- **Factory Pattern**: Easy service creation and configuration
- **Environment Configuration**: Automatic config loading via environment variables
- **Graceful Shutdown**: Coordinated service lifecycle management
- **Mock Implementations**: Ready for testing and development
- **Extensible Architecture**: Easy to add new service implementations

## Quick Start

### 1. Use the Shared Library

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/yourusername/shared-foundation"
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

    // Use injected dependencies
    logger := application.Logger()
    tracer := application.Tracer()
    metrics := application.Metrics()
    broker := application.Broker()
    cache := application.Cache()
    database := application.Database()
    connectServer := application.ConnectRPCServer()

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
```

### 2. Configuration

The library uses environment variables for configuration:

```bash
# Application
export APP_NAME="user-service"
export APP_VERSION="1.0.0"
export APP_ENV="development"

# Logger
export LOGGER_LEVEL="info"
export LOGGER_FORMAT="json"
export LOGGER_OUTPUT="stdout"

# Tracer
export TRACER_TYPE="jaeger"
export TRACER_ENDPOINT="http://jaeger:14268"

# Metrics
export METRICS_TYPE="prometheus"
export METRICS_PORT="9090"

# Broker
export BROKER_TYPE="kafka"
export BROKER_BROKERS="kafka:9092"

# Cache
export CACHE_TYPE="redis"
export CACHE_ADDRESS="redis:6379"

# Database
export DATABASE_TYPE="postgres"
export DATABASE_DSN="postgres://user:pass@db:5432/mydb"

# ConnectRPC
export CONNECTRPC_ADDRESS=":8080"
```

## Status

This is a **skeleton implementation** with:
- ✅ **Architecture** - Complete interface definitions and structure
- ✅ **Configuration** - Environment-based configuration loading
- ✅ **Factory Pattern** - Service creation factories
- ✅ **Mock Implementations** - For testing and development
- 🚧 **Real Implementations** - Placeholder implementations (TODO)

## Development

### Adding New Service Implementations

1. Create implementation in the appropriate package (e.g., `logger/logrus.go`)
2. Add factory method in `factory.go`
3. Add configuration in `config.go`
4. Add tests

### Testing

```bash
cd shared-foundation
go test ./...
```

## Dependencies

- [connectrpc.com/connect](https://connectrpc.com/)
- [github.com/spf13/viper](https://github.com/spf13/viper)
- [log/slog](https://pkg.go.dev/log/slog)

## Next Steps

1. Implement actual service providers (Jaeger, Prometheus, Redis, etc.)
2. Add comprehensive tests
3. Create usage examples
4. Add documentation for each package 
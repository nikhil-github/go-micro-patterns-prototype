# Shared Foundation Library

A Go-idiomatic shared library for microservices with infrastructure components.

## Structure

```
shared-foundation/
├── go.mod
├── go.sum
├── README.md
├── types.go                        # All interfaces and types in one file
├── config.go                       # Configuration
├── app.go                          # Main App orchestrator
├── factory.go                      # Service factories
├── mocks.go                        # Mock implementations for testing
├── logger/                         # Logger package
│   ├── slog.go
│   ├── logrus.go
│   └── slog_test.go
├── tracer/                         # Tracer package
│   ├── jaeger.go
│   ├── zipkin.go
│   └── jaeger_test.go
├── metrics/                        # Metrics package
│   ├── prometheus.go
│   ├── statsd.go
│   └── prometheus_test.go
├── broker/                         # Broker package
│   ├── kafka.go
│   ├── rabbitmq.go
│   ├── nats.go
│   └── kafka_test.go
├── cache/                          # Cache package
│   ├── redis.go
│   ├── memcached.go
│   └── redis_test.go
├── database/                       # Database package
│   ├── postgres.go
│   ├── mysql.go
│   └── postgres_test.go
├── connectrpc/                     # ConnectRPC package
│   ├── server.go
│   └── server_test.go
└── examples/                       # Usage examples
    ├── user-service/
    ├── order-service/
    ├── payment-service/
    └── notification-service/
```

## Usage

### In a Microservice

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

    // Create business logic with injected dependencies
    userHandler := &handlers.UserHandler{
        Logger:   application.Logger(),
        Tracer:   application.Tracer(),
        Metrics:  application.Metrics(),
        Broker:   application.Broker(),
        Cache:    application.Cache(),
        Database: application.Database(),
    }

    // Register ConnectRPC handlers
    connectServer := application.ConnectRPCServer()
    connectServer.RegisterHandler("/user.UserService/", userHandler)

    // Wait for shutdown signal
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    application.Logger().Info("User service started successfully")
    <-sigChan

    application.Logger().Info("Shutdown signal received, stopping app...")
    if err := application.Stop(ctx); err != nil {
        application.Logger().Error("Error during shutdown", "error", err)
        os.Exit(1)
    }
}
```

## Configuration

The library uses environment variables for configuration:

- `APP_NAME`: Application name
- `APP_VERSION`: Application version
- `APP_ENV`: Environment (development, staging, production)
- `LOGGER_LEVEL`: Log level (debug, info, warn, error)
- `LOGGER_FORMAT`: Log format (text, json)
- `LOGGER_OUTPUT`: Log output (stdout, stderr, file path)
- `TRACER_TYPE`: Tracer type (jaeger, zipkin)
- `TRACER_ENDPOINT`: Tracer endpoint
- `METRICS_TYPE`: Metrics type (prometheus, statsd)
- `METRICS_PORT`: Metrics port
- `BROKER_TYPE`: Broker type (kafka, rabbitmq, nats)
- `CACHE_TYPE`: Cache type (redis, memcached)
- `DATABASE_TYPE`: Database type (postgres, mysql)
- `CONNECTRPC_ADDRESS`: ConnectRPC server address

## Status

This is a **skeleton implementation** with:
- ✅ **Architecture** - Complete interface definitions and structure
- ✅ **Configuration** - Environment-based configuration loading
- ✅ **Factory Pattern** - Service creation factories
- ✅ **Mock Implementations** - For testing and development
- 🚧 **Real Implementations** - Placeholder implementations (TODO)

## Next Steps

1. Implement actual service providers (Jaeger, Prometheus, Redis, etc.)
2. Add comprehensive tests
3. Create usage examples
4. Add documentation for each package 
# Shared Library Structure (Go Idiomatic)

## Overview
This document explains how to structure the shared library for your 4 microservices with 8+ infrastructure components, following Go conventions.

## Recommended Shared Library Structure (Go Idiomatic)

```
shared-foundation/
├── go.mod
├── go.sum
├── README.md
├── types/                          # All service interfaces and types
│   └── types.go
├── config/                         # Configuration management
│   └── config.go
├── app/                            # Main App orchestrator
│   └── app.go
├── factory/                        # Service factories
│   ├── factory.go
│   └── mocks.go
├── logger/                         # Logger implementations
│   ├── slog.go                     # slog implementation
│   ├── logrus.go                   # logrus implementation
│   └── slog_test.go
├── tracer/                         # Tracer implementations
│   ├── jaeger.go
│   ├── zipkin.go
│   └── jaeger_test.go
├── metrics/                        # Metrics implementations
│   ├── prometheus.go
│   ├── statsd.go
│   └── prometheus_test.go
├── broker/                         # Broker implementations
│   ├── kafka.go
│   ├── rabbitmq.go
│   ├── nats.go
│   └── kafka_test.go
├── cache/                          # Cache implementations
│   ├── redis.go
│   ├── memcached.go
│   └── redis_test.go
├── database/                       # Database implementations
│   ├── postgres.go
│   ├── mysql.go
│   └── postgres_test.go
├── connectrpc/                     # ConnectRPC implementations
│   ├── server.go
│   └── server_test.go
└── examples/                       # Usage examples
    ├── user-service/
    ├── order-service/
    ├── payment-service/
    └── notification-service/
```

## Alternative Structure (Even More Go-like)

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

## Go Module Structure

### Shared Library (shared-foundation)
```go
// go.mod
module github.com/yourusername/shared-foundation

go 1.21

require (
    github.com/spf13/viper v1.17.0
    github.com/sirupsen/logrus v1.9.3  // optional for logrus logger
    // Add other dependencies as needed
)
```

### Microservice (user-service)
```go
// go.mod
module github.com/yourusername/user-service

go 1.21

require (
    github.com/yourusername/shared-foundation v0.1.0
    // Add other dependencies as needed
)
```

## Usage in Microservices

### User Service Example
```go
// user-service/main.go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/yourusername/shared-foundation/app"
    "github.com/yourusername/user-service/handlers"
)

func main() {
    // Create app with name and version
    application := app.New("user-service", "1.0.0")

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

### User Handler Example
```go
// user-service/handlers/user_handler.go
package handlers

import (
    "context"

    "github.com/yourusername/shared-foundation/types"
    "github.com/yourusername/user-service/models"
)

type UserHandler struct {
    Logger   types.Logger
    Tracer   types.Tracer
    Metrics  types.Metrics
    Broker   types.Broker
    Cache    types.Cache
    Database types.Database
}

func (h *UserHandler) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
    span := h.Tracer.StartSpan("create_user")
    defer span.Finish()

    h.Metrics.Counter("user_creation_requests", 1, "service", "user-service")

    // Business logic here...
    user := &models.User{
        ID:   "123",
        Name: req.Name,
        Email: req.Email,
    }

    // Cache user data
    h.Cache.Set("user:123", []byte(user.Name), 0)

    // Publish event
    h.Broker.Publish("user.created", []byte(`{"user_id": "123"}`))

    h.Logger.Info("User created successfully", "user_id", user.ID)

    return &models.CreateUserResponse{UserID: user.ID}, nil
}
```

## Package Organization Benefits

### ✅ **Go Idiomatic**
- **types.go** - Common Go pattern for interfaces and types
- **Package-based organization** - Each component in its own package
- **Flat structure** - No deep nesting
- **Clear naming** - Descriptive package names

### ✅ **Import Clarity**
```go
import (
    "github.com/yourusername/shared-foundation/types"
    "github.com/yourusername/shared-foundation/logger"
    "github.com/yourusername/shared-foundation/tracer"
    "github.com/yourusername/shared-foundation/broker"
)
```

### ✅ **Easy Testing**
```go
import (
    "github.com/yourusername/shared-foundation/mocks"
)

func TestUserHandler(t *testing.T) {
    mockLogger := &mocks.MockLogger{}
    mockTracer := &mocks.MockTracer{}
    // ... test with mocks
}
```

## Migration from Current Structure

If you want to migrate from the current structure:

1. **Move interfaces** from `interfaces/interfaces.go` to `types.go`
2. **Move implementations** from `implementations/logger/` to `logger/`
3. **Update imports** in all files
4. **Update factory** to use new package structure

## Recommended Approach

I recommend the **second structure** (Alternative Structure) because:

- ✅ **More Go-like** - Follows standard Go conventions
- ✅ **Simpler imports** - `types.Logger` instead of `interfaces.Logger`
- ✅ **Better organization** - Each component in its own package
- ✅ **Easier to navigate** - Flat structure with clear package names
- ✅ **Standard Go patterns** - Similar to how Go standard library is organized

This structure is much more familiar to Go developers and follows the conventions they expect! 
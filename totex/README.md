# Totex - Go Microservices Foundation Library

A minimal, production-ready Go library for microservice lifecycle management with cross-cutting concerns.

## Overview

The foundation library provides a clean, simple way to bootstrap microservices with:
- **Lifecycle Management** - Start/stop coordination for servers
- **Cross-cutting Concerns** - Logging, metrics, tracing
- **Auto-configured Servers** - ConnectRPC servers created from environment variables
- **Graceful Shutdown** - Proper resource cleanup
- **Dependency Injection** - Common services for business logic

## Structure

```
totex/
├── foundation/              # Main shared library
│   ├── app.go              # Main App orchestrator
│   ├── config.go           # Configuration management
│   ├── connectrpc/         # ConnectRPC server implementation
│   │   └── server.go
│   ├── logging/            # Logger interfaces and implementations
│   │   ├── logger.go       # Interface
│   │   └── slog.go         # Default implementation
│   ├── metrics/            # Metrics interfaces
│   │   └── metrics.go      # Interface + default implementation
│   ├── tracing/            # Tracing interfaces
│   │   └── tracer.go       # Interface + default implementation
│   ├── examples/           # Usage examples
│   │   └── example/        # Comprehensive example
│   ├── go.mod
│   └── README.md
├── schema/                 # Protocol Buffers and ConnectRPC
│   ├── user/v1/           # User service definitions
│   ├── order/v1/          # Order service definitions
│   └── gen/               # Generated Go code
├── user-service/          # Example microservice
│   ├── cmd/main.go
│   └── internal/user/
├── go.mod                 # Workspace configuration
└── README.md              # This file
```

## Quick Start

### 1. Create a Microservice

```go
package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"

    foundation "github.com/yourusername/foundation"
)

func main() {
    // Configure server via environment variables
    os.Setenv("SERVER_NAME", "my-service-server")
    os.Setenv("SERVER_ADDR", ":8080")

    // Create the app (factory method)
    app := foundation.New("my-service", "1.0.0")

    // Get the auto-created ConnectRPC server and register handlers
    connectServer := app.ConnectRPC()
    connectServer.RegisterHandler("/my.MyService/", myHandler)

    // Start all servers
    if err := app.Start(context.Background()); err != nil {
        app.Logger().Error("Failed to start app", "error", err)
        os.Exit(1)
    }

    // Graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    if err := app.Stop(context.Background()); err != nil {
        app.Logger().Error("Error during shutdown", "error", err)
        os.Exit(1)
    }
}
```

### 2. Use Injected Dependencies

```go
// Get injected dependencies
logger := app.Logger()
tracer := app.Tracer()
metrics := app.Metrics()

// Use in business logic
func (h *MyHandler) ProcessRequest(ctx context.Context, data string) error {
    // Start tracing
    span := h.Tracer.StartSpan("process_request")
    defer span.Finish()

    // Record metrics
    h.Metrics.Counter("requests_processed", 1, "service", "my-service")

    // Log operation
    h.Logger.Info("Processing request", "data", data)
    
    return nil
}
```

## Configuration

The library uses environment variables for configuration:

```bash
# Server configuration
export SERVER_NAME="my-service-server"
export SERVER_ADDR=":8080"
export SERVER_TYPE="connectrpc"     # connectrpc (default)

# Logger configuration
export LOGGER_TYPE="slog"           # slog, logrus
export LOGGER_LEVEL="info"          # debug, info, warn, error
export LOGGER_FORMAT="text"         # text, json
export LOGGER_OUTPUT="stdout"       # stdout, stderr, file path

# Tracer configuration
export TRACER_TYPE="noop"           # noop, jaeger, zipkin
export TRACER_ENDPOINT=""           # tracer endpoint URL

# Metrics configuration
export METRICS_TYPE="noop"          # noop, prometheus, statsd
export METRICS_PORT="9090"          # metrics port
```

## Key Features

### 1. **Service Factory Method**
- `foundation.New(name, version string) *App` - Creates app with all dependencies
- `app.ConnectRPC()` - Direct access to ConnectRPC server
- `app.Logger()`, `app.Metrics()`, `app.Tracer()` - Cross-cutting dependencies

### 2. **Auto-configured Servers**
- ConnectRPC servers created automatically from environment variables
- No manual server creation for common use cases
- Easy to extend with additional server types

### 3. **Lifecycle Management**
- Coordinated start/stop of all servers
- Graceful shutdown with proper resource cleanup
- Error handling and logging throughout

### 4. **Clean Architecture**
- Default implementations in their respective packages
- Clear separation of concerns
- Easy to extend and maintain

## Examples

### User Service
See `user-service/` for a complete example microservice using the foundation library.

### Foundation Examples
See `foundation/examples/` for comprehensive examples showing:
- Foundation initialization
- ConnectRPC server integration
- Dependency injection
- Graceful shutdown
- Business logic implementation

```bash
cd foundation/examples/example
go run main.go
```

## Status

This is a **working implementation** with:
- ✅ **Core Architecture** - Complete foundation with lifecycle management
- ✅ **Auto-configured Servers** - ConnectRPC servers created from environment variables
- ✅ **Cross-cutting Concerns** - Logging, metrics, tracing
- ✅ **Examples** - Working examples and documentation
- ✅ **Graceful Shutdown** - Proper resource cleanup
- ✅ **Clean Architecture** - Default implementations in their respective packages
- 🚧 **Advanced Features** - Additional server types, real implementations (TODO)

## Development

### Building

```bash
# Build foundation library
cd foundation && go build .

# Build user service
cd user-service && go build cmd/main.go

# Build examples
cd foundation/examples/example && go build main.go
```

### Testing

```bash
cd foundation
go test ./...
```

## Dependencies

- [connectrpc.com/connect](https://connectrpc.com/) - ConnectRPC implementation
- [log/slog](https://pkg.go.dev/log/slog) - Structured logging
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Alternative logger
- [github.com/spf13/viper](https://github.com/spf13/viper) - Configuration management

## Architecture Benefits

### 1. **Simplified Setup**
- Servers created automatically from configuration
- Minimal boilerplate for common use cases
- Environment-based configuration
- Name and version as simple arguments

### 2. **Separation of Concerns**
- Foundation handles cross-cutting concerns only
- Transport layer is separate and optional
- Business logic is clean and focused

### 3. **Flexibility**
- Support for multiple transport protocols
- Configurable server addresses and settings
- Easy to extend with new server types

### 4. **Testability**
- Clear interfaces for mocking
- Separation enables unit testing
- Mock implementations provided

### 5. **Maintainability**
- Small, focused modules
- Clear responsibilities
- Easy to understand and extend

## Next Steps

1. **Real Implementations** - Replace noop implementations with real ones (Jaeger, Prometheus, etc.)
2. **Additional Server Types** - gRPC, WebSocket, HTTP servers
3. **Health Checks** - Auto-created health check servers
4. **Service Discovery** - Auto-registration with service discovery
5. **Configuration Validation** - Validate environment variables at startup

# Foundation Library

A Go-idiomatic shared library for microservices with cross-cutting concerns and lifecycle management.

## Overview

The foundation library provides:
- **Cross-cutting concerns**: Logging, metrics, tracing, configuration
- **Lifecycle management**: Start/stop coordination for servers
- **Dependency injection**: Common services for business logic
- **Graceful shutdown**: Proper resource cleanup

## Structure

```
foundation/
├── app.go                          # Main App orchestrator
├── config.go                       # Configuration management
├── connectrpc/                     # ConnectRPC server implementation
│   └── server.go
├── logging/                        # Logger interfaces and implementations
│   ├── logger.go
│   ├── slog.go
│   └── logrus.go
├── metrics/                        # Metrics interfaces
│   └── metrics.go
├── tracing/                        # Tracing interfaces
│   └── tracer.go
└── examples/                       # Usage examples
    └── example/                    # Comprehensive example
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"

    foundation "github.com/yourusername/foundation"
    "github.com/yourusername/foundation/connectrpc"
)

func main() {
    // Create app with default configuration
    app := foundation.New()

    // Create ConnectRPC server
    server := connectrpc.NewServer("my-service", ":8080", app.Logger())
    
    // Add server to lifecycle management
    app.AddServer(server)

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

### With Business Logic

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

## Key Features

### 1. Cross-Cutting Concerns
- **Logging**: Structured logging with multiple backends (slog, logrus)
- **Metrics**: Observability metrics (counters, gauges, histograms)
- **Tracing**: Distributed tracing support
- **Configuration**: Environment-based configuration

### 2. Lifecycle Management
- **Server Registry**: Add multiple servers with `AddServer()`
- **Ordered Startup**: Servers start in registration order
- **Graceful Shutdown**: Servers stop in reverse order
- **Error Handling**: Proper error propagation

### 3. Dependency Injection
- **Logger**: `app.Logger()` - Structured logging
- **Tracer**: `app.Tracer()` - Distributed tracing
- **Metrics**: `app.Metrics()` - Observability metrics

## Configuration

The library uses environment variables for configuration:

```bash
# Logger configuration
export LOGGER_TYPE="slog"           # slog, logrus
export LOGGER_LEVEL="info"          # debug, info, warn, error
export LOGGER_FORMAT="text"         # text, json
export LOGGER_OUTPUT="stdout"       # stdout, stderr, file path
```

## Examples

See the `examples/` directory for a comprehensive example showing:
- Foundation initialization
- ConnectRPC server integration
- Dependency injection
- Graceful shutdown
- Business logic implementation

```bash
cd examples/example
go run main.go
```

## Architecture Benefits

### 1. **Separation of Concerns**
- Foundation handles cross-cutting concerns only
- Transport layer is separate and optional
- Business logic is clean and focused

### 2. **Flexibility**
- Support for multiple transport protocols
- Configurable server addresses and settings
- Easy to extend with new server types

### 3. **Testability**
- Clear interfaces for mocking
- Separation enables unit testing
- Mock implementations provided

### 4. **Maintainability**
- Small, focused modules
- Clear responsibilities
- Easy to understand and extend

## Migration from Previous Versions

If you're migrating from an older version:

1. **Update imports**: Add `"github.com/yourusername/foundation/connectrpc"`
2. **Create server**: Use `connectrpc.NewServer()` instead of `app.ConnectRPCServer()`
3. **Register server**: Call `app.AddServer(server)` to add to lifecycle management
4. **No other changes needed**: All other foundation APIs remain the same

## Status

This is a **working implementation** with:
- ✅ **Core Architecture** - Complete foundation with lifecycle management
- ✅ **ConnectRPC Server** - HTTP server for ConnectRPC
- ✅ **Cross-cutting Concerns** - Logging, metrics, tracing
- ✅ **Examples** - Working examples and documentation
- ✅ **Graceful Shutdown** - Proper resource cleanup
- 🚧 **Advanced Features** - Additional server types, real implementations (TODO) 
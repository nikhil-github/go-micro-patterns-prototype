# Examples

This directory contains example microservices that demonstrate how to use the shared-foundation library.

## Available Examples

### demo-service
A simple service that demonstrates basic usage of the shared-foundation library.

```bash
cd demo-service
go run main.go
```

### user-service
A more comprehensive example showing how to use all the injected dependencies (logger, tracer, metrics, broker, cache, database, ConnectRPC).

```bash
cd user-service
go run main.go
```

## Running Examples

Each example is a separate Go module. To run an example:

1. Navigate to the example directory
2. Run the service:
   ```bash
   go run main.go
   ```

## Configuration

The examples use environment variables for configuration. You can set them before running:

```bash
export APP_NAME="demo-service"
export APP_VERSION="1.0.0"
export LOGGER_LEVEL="info"
export LOGGER_FORMAT="json"
go run main.go
```

## Expected Output

When running an example, you should see:
- Service initialization messages
- "Service started successfully" message
- Graceful shutdown when you press Ctrl+C

## Note

These examples use mock implementations of the infrastructure services. In a real environment, you would:
1. Implement actual service providers (Jaeger, Prometheus, Redis, etc.)
2. Configure real infrastructure endpoints
3. Handle actual business logic 
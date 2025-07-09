# Foundation Examples

This directory contains examples demonstrating how to use the foundation library.

## Example

The `example/` directory contains a comprehensive example showing:

- **Basic foundation setup** with default configuration
- **ConnectRPC server integration** with lifecycle management
- **Dependency injection** of logger, tracer, and metrics
- **Graceful shutdown** handling
- **Business logic** demonstrating how to use the injected dependencies

### Key Features Demonstrated

1. **Foundation Initialization**
   ```go
   app := foundation.New()
   ```

2. **Server Lifecycle Management**
   ```go
   server := connectrpc.NewServer("example-server", ":8080", logger)
   app.AddServer(server)
   ```

3. **Dependency Injection**
   ```go
   logger := app.Logger()
   tracer := app.Tracer()
   metrics := app.Metrics()
   ```

4. **Graceful Shutdown**
   ```go
   sigChan := make(chan os.Signal, 1)
   signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
   <-sigChan
   app.Stop(ctx)
   ```

### Running the Example

```bash
cd example
go run main.go
```

The example will start a ConnectRPC server on port 8080 and demonstrate proper lifecycle management with graceful shutdown.

## Usage Pattern

This example demonstrates the recommended pattern for using the foundation library:

1. **Create the app** with default or custom configuration
2. **Add servers** to the app's lifecycle management
3. **Start all servers** with proper error handling
4. **Handle graceful shutdown** on system signals
5. **Use injected dependencies** in business logic

This pattern ensures proper resource management, observability, and graceful shutdown for microservices.

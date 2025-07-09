# Foundation Examples

This directory contains examples demonstrating how to use the foundation library.

## Example

The `example/` directory contains a comprehensive example showing:

- **Auto-configured foundation setup** with environment variables
- **ConnectRPC server auto-creation** from configuration
- **Dependency injection** of logger, tracer, and metrics
- **Graceful shutdown** handling
- **Business logic** demonstrating how to use the injected dependencies

### Key Features Demonstrated

1. **Environment Configuration**
   ```bash
   export SERVER_NAME="example-server"
   export SERVER_ADDR=":8080"
   ```

2. **Foundation Initialization**
   ```go
   app := foundation.New("example-service", "1.0.0")  // Auto-creates servers from env vars
   ```

3. **Auto-created Server Access**
   ```go
   server := app.GetServerByName("example-server")
   connectServer := server.(*connectrpc.Server)
   ```

4. **Dependency Injection**
   ```go
   logger := app.Logger()
   tracer := app.Tracer()
   metrics := app.Metrics()
   ```

5. **Graceful Shutdown**
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

1. **Set environment variables** for server configuration
2. **Create the app** with `foundation.New(name, version)` (auto-creates servers)
3. **Access auto-created servers** with `app.GetServerByName()`
4. **Register handlers** with the servers
5. **Start all servers** with proper error handling
6. **Handle graceful shutdown** on system signals
7. **Use injected dependencies** in business logic

This pattern ensures:
- **Minimal boilerplate** for common use cases
- **Proper resource management** and observability
- **Graceful shutdown** for microservices
- **Environment-based configuration** for easy deployment
- **Clean API** with name and version as simple arguments

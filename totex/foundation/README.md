# Foundation

A minimal Go library for microservice lifecycle, logging, metrics, and tracing.

## Quick Start

```go
import (
    "os"
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
    app.Start(context.Background())
}
```

## Service Factory Method

- `foundation.New(name, version string) *App` — creates an app with logging, metrics, tracing, and servers from environment variables.
- `app.ConnectRPC()` — get the ConnectRPC server directly.
- `app.Logger()`, `app.Metrics()`, `app.Tracer()` — access cross-cutting dependencies.

## Usage

1. Set environment variables for server config:
   ```bash
   export SERVER_NAME="my-service-server"
   export SERVER_ADDR=":8080"
   ```
2. Call `foundation.New("my-service", "1.0.0")` in your main.
3. Register handlers with `app.ConnectRPC()`.
4. Start the app.

See the `examples/` directory for a full example. 
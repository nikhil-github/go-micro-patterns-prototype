package main

import (
	"log/slog"

	"github.com/yourusername/foundation/patterns/builder/builder"
	"github.com/yourusername/foundation/patterns/builder/services"
)

func main() {
	// Builder pattern with method chaining
	app := builder.NewMicroservice().
		WithLogger(slog.Default()).
		AddConnectRPCServer(services.NewDummyHandler()).
		Build().
		Start()

	// Wait for shutdown signal
	app.WaitForShutdown()
}

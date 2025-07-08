package main

import (
	"log"
	"net/http"
	"os"

	userv1connect "github.com/yourusername/microservices-protos/gen/user/v1/userv1connect"
	foundation "github.com/yourusername/shared-foundation"
	"github.com/yourusername/user-service/internal/user"
)

func main() {
	cfg := foundation.Load()

	// Use Zap logger from shared-foundation (assume WithZapLogger exists)
	app := foundation.New("user-service", "1.0.0",
		foundation.WithSlogLogger(cfg.Logger),
	)
	if err := app.Init(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	logger := app.Logger()

	// User service implementation (no DB)
	userSvc := user.NewService()

	// ConnectRPC handler
	path, handler := userv1connect.NewUserServiceHandler(userSvc)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:    cfg.ConnectRPC.Address,
		Handler: mux,
	}

	logger.Info("Starting user-service", "address", cfg.ConnectRPC.Address)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server error", "error", err)
		os.Exit(1)
	}
}

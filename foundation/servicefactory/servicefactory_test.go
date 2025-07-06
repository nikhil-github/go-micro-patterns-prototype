package servicefactory

import (
	"log/slog"
	"testing"
	"time"

	"github.com/yourusername/foundation/connectrpc"
)

func TestServiceFactory_Build(t *testing.T) {
	factory := NewServiceFactory()
	factory.Build()
	// Build should not panic and should initialize viper
}

func TestServiceFactory_GetContext(t *testing.T) {
	factory := NewServiceFactory()
	ctx := factory.GetContext()
	if ctx == nil {
		t.Fatal("expected non-nil context")
	}
}

func TestCreateConnectRPCServer_Success(t *testing.T) {
	factory := NewServiceFactory()
	factory.Build()

	handler := connectrpc.NewDummyHandler()
	service, err := factory.CreateConnectRPCServer(handler)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if service == nil {
		t.Fatal("expected service, got nil")
	}
}

func TestCreateConnectRPCServerWithConfig_Success(t *testing.T) {
	factory := NewServiceFactory()
	factory.Build()

	cfg := connectrpc.Config{Address: ":9090"}
	handler := connectrpc.NewDummyHandler()
	service, err := factory.CreateConnectRPCServerWithConfig(cfg, handler)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if service == nil {
		t.Fatal("expected service, got nil")
	}
}

func TestGetConnectRPCConfig(t *testing.T) {
	factory := NewServiceFactory()
	factory.Build()

	cfg := factory.GetConnectRPCConfig()
	if cfg.Address == "" {
		t.Fatal("expected non-empty address from config")
	}
}

func TestNewServiceFactoryWithLogger(t *testing.T) {
	logger := slog.Default()
	factory := NewServiceFactoryWithLogger(logger)
	if factory.GetLogger() != logger {
		t.Fatal("expected custom logger to be set")
	}
}

func TestStartAllAndStopAll(t *testing.T) {
	factory := NewServiceFactory()
	factory.Build()

	// Create a service
	handler := connectrpc.NewDummyHandler()
	_, err := factory.CreateConnectRPCServer(handler)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	// Start all services
	if err := factory.StartAll(); err != nil {
		t.Fatalf("failed to start services: %v", err)
	}

	// Stop all services
	if err := factory.StopAll(); err != nil {
		t.Fatalf("failed to stop services: %v", err)
	}
}

func TestContextCancellation(t *testing.T) {
	factory := NewServiceFactory()
	ctx := factory.GetContext()

	// Cancel the context
	factory.cancel()

	// Wait a bit for cancellation to propagate
	time.Sleep(10 * time.Millisecond)

	select {
	case <-ctx.Done():
		// Expected
	default:
		t.Fatal("expected context to be cancelled")
	}
}

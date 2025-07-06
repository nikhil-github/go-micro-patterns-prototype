package servicefactory

import (
	"log/slog"
	"testing"
	"time"

	"github.com/yourusername/foundation/patterns/current/connectrpc"
)

func TestServiceFactory_Init(t *testing.T) {
	factory := NewServiceFactory()
	factory.Init()
	// Init should not panic and should initialize viper
}

func TestServiceFactory_GetContext(t *testing.T) {
	factory := NewServiceFactory()
	ctx := factory.GetContext()
	if ctx == nil {
		t.Fatal("expected non-nil context")
	}
}

func TestAddConnectRPCServer_Success(t *testing.T) {
	factory := NewServiceFactory()
	factory.Init()

	handler := connectrpc.NewDummyHandler()
	err := factory.AddConnectRPCServer(handler)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestAddConnectRPCServerWithConfig_Success(t *testing.T) {
	factory := NewServiceFactory()
	factory.Init()

	cfg := connectrpc.Config{Address: ":9090"}
	handler := connectrpc.NewDummyHandler()
	err := factory.AddConnectRPCServerWithConfig(cfg, handler)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewServiceFactoryWithLogger(t *testing.T) {
	logger := slog.Default()
	factory := NewServiceFactoryWithLogger(logger)
	if factory.GetLogger() != logger {
		t.Fatal("expected custom logger to be set")
	}
}

func TestStartAndStop(t *testing.T) {
	factory := NewServiceFactory()
	factory.Init()

	// Add a service
	handler := connectrpc.NewDummyHandler()
	err := factory.AddConnectRPCServer(handler)
	if err != nil {
		t.Fatalf("failed to add service: %v", err)
	}

	// Start services
	if err := factory.Start(); err != nil {
		t.Fatalf("failed to start services: %v", err)
	}

	// Stop services
	if err := factory.Stop(); err != nil {
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

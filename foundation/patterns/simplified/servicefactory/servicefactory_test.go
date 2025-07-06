package servicefactory

import (
	"log/slog"
	"testing"

	"github.com/yourusername/foundation/patterns/simplified/config"
)

func TestNew(t *testing.T) {
	cfg := &config.Config{}
	logger := slog.Default()

	factory := New(cfg, logger)

	if factory.config != cfg {
		t.Error("expected config to be set")
	}
	if factory.logger != logger {
		t.Error("expected logger to be set")
	}
	if factory.ctx == nil {
		t.Error("expected context to be created")
	}
}

func TestAddConnectRPCServer(t *testing.T) {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
	}
	logger := slog.Default()
	factory := New(cfg, logger)

	handler := NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	if len(factory.services) != 1 {
		t.Errorf("expected 1 service, got %d", len(factory.services))
	}
}

func TestAddConnectRPCServerWithConfig(t *testing.T) {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
	}
	logger := slog.Default()
	factory := New(cfg, logger)

	handler := NewDummyHandler()
	customCfg := config.ConnectRPC{Address: ":9090"}
	factory.AddConnectRPCServerWithConfig(customCfg, handler)

	if len(factory.services) != 1 {
		t.Errorf("expected 1 service, got %d", len(factory.services))
	}
}

func TestConnectRPCServer_Start(t *testing.T) {
	cfg := config.ConnectRPC{Address: ":8080"}
	logger := slog.Default()
	handler := NewDummyHandler()

	server := NewConnectRPCServer(cfg, handler, logger)

	if err := server.Start(nil); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestConnectRPCServer_Stop(t *testing.T) {
	cfg := config.ConnectRPC{Address: ":8080"}
	logger := slog.Default()
	handler := NewDummyHandler()

	server := NewConnectRPCServer(cfg, handler, logger)

	if err := server.Stop(nil); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestNewDummyHandler(t *testing.T) {
	handler := NewDummyHandler()
	if handler == nil {
		t.Error("expected non-nil handler")
	}
}

func TestServiceFactory_StartStop(t *testing.T) {
	cfg := &config.Config{
		ConnectRPC: config.ConnectRPC{Address: ":8080"},
	}
	logger := slog.Default()
	factory := New(cfg, logger)

	// Add a service
	handler := NewDummyHandler()
	factory.AddConnectRPCServer(handler)

	// Test start
	if err := factory.start(); err != nil {
		t.Errorf("expected no error on start, got %v", err)
	}

	// Test stop
	if err := factory.stop(); err != nil {
		t.Errorf("expected no error on stop, got %v", err)
	}
}

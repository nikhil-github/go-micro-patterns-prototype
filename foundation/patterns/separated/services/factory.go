package services

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/yourusername/foundation/patterns/separated/config"
)

type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Factory struct {
	logger *slog.Logger
}

func NewFactory(logger *slog.Logger) *Factory {
	return &Factory{logger: logger}
}

func (f *Factory) CreateConnectRPCServer(cfg config.ConnectRPC, handler *connect.Handler) Service {
	return &ConnectRPCServer{
		config:  cfg,
		handler: handler,
		logger:  f.logger,
	}
}

type ConnectRPCServer struct {
	config  config.ConnectRPC
	handler *connect.Handler
	logger  *slog.Logger
}

func (s *ConnectRPCServer) Start(ctx context.Context) error {
	s.logger.Info("Starting connectRPC server", "address", s.config.Address)
	// TODO: Implement actual server start logic
	return nil
}

func (s *ConnectRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping connectRPC server", "address", s.config.Address)
	// TODO: Implement actual server stop logic
	return nil
}

func NewDummyHandler() *connect.Handler {
	return connect.NewUnaryHandler("/test.TestService/TestMethod", func(ctx context.Context, req *connect.Request[any]) (*connect.Response[any], error) {
		return connect.NewResponse[any](nil), nil
	})
}

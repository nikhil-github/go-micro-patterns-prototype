package services

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/yourusername/foundation/patterns/builder/config"
)

type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type ConnectRPCServer struct {
	config  config.ConnectRPCConfig
	handler *connect.Handler
	logger  *slog.Logger
}

func NewConnectRPCServer(cfg config.ConnectRPCConfig, handler *connect.Handler, logger *slog.Logger) Service {
	return &ConnectRPCServer{
		config:  cfg,
		handler: handler,
		logger:  logger,
	}
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

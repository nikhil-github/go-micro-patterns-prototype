package services

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type ConnectRPCServer struct {
	address string
	handler *connect.Handler
	logger  *slog.Logger
}

func NewConnectRPCServer(address string, handler *connect.Handler, logger *slog.Logger) Service {
	return &ConnectRPCServer{
		address: address,
		handler: handler,
		logger:  logger,
	}
}

func (s *ConnectRPCServer) Start(ctx context.Context) error {
	s.logger.Info("Starting connectRPC server", "address", s.address)
	// TODO: Implement actual server start logic
	return nil
}

func (s *ConnectRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping connectRPC server", "address", s.address)
	// TODO: Implement actual server stop logic
	return nil
}

func NewDummyHandler() *connect.Handler {
	return connect.NewUnaryHandler("/test.TestService/TestMethod", func(ctx context.Context, req *connect.Request[any]) (*connect.Response[any], error) {
		return connect.NewResponse[any](nil), nil
	})
}

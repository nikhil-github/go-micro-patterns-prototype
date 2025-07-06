package connectrpc

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/spf13/viper"
	"github.com/yourusername/shared-foundation/core"
)

type Config struct {
	Address string
}

func LoadConfig() Config {
	viper.SetDefault("CONNECTRPC_ADDRESS", ":8080")
	viper.BindEnv("CONNECTRPC_ADDRESS")
	return Config{
		Address: viper.GetString("CONNECTRPC_ADDRESS"),
	}
}

type Server struct {
	config  Config
	handler *connect.Handler
	logger  *slog.Logger
}

func NewServer(cfg Config, handler *connect.Handler, logger *slog.Logger) *Server {
	return &Server{
		config:  cfg,
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting connectRPC server", "address", s.config.Address)
	// TODO: Implement actual server start logic
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Stopping connectRPC server", "address", s.config.Address)
	// TODO: Implement actual server stop logic
	return nil
}

func (s *Server) Name() string {
	return "connectrpc-server"
}

func NewDummyHandler() *connect.Handler {
	return connect.NewUnaryHandler("/test.TestService/TestMethod", func(ctx context.Context, req *connect.Request[any]) (*connect.Response[any], error) {
		return connect.NewResponse[any](nil), nil
	})
}

// ConnectRPCServer interface for gRPC/Connect-RPC server
type ConnectRPCServer interface {
	core.Service
	RegisterHandler(path string, handler interface{}) error
	GetHandler() interface{} // Returns the underlying handler for registration
}

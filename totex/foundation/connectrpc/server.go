package connectrpc

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/spf13/viper"
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
	config   Config
	handlers map[string]interface{}
	logger   *slog.Logger
}

func NewServer(cfg Config, logger *slog.Logger) *Server {
	return &Server{
		config:   cfg,
		handlers: make(map[string]interface{}),
		logger:   logger,
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

func (s *Server) RegisterHandler(path string, handler interface{}) error {
	s.handlers[path] = handler
	s.logger.Info("Registered handler", "path", path)
	return nil
}

func (s *Server) GetHandler() interface{} {
	return s.handlers
}

func NewDummyHandler() *connect.Handler {
	return connect.NewUnaryHandler("/test.TestService/TestMethod", func(ctx context.Context, req *connect.Request[any]) (*connect.Response[any], error) {
		return connect.NewResponse[any](nil), nil
	})
}

// ConnectRPCServer interface for gRPC/Connect-RPC server
type ConnectRPCServer interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
	RegisterHandler(path string, handler interface{}) error
	GetHandler() interface{}
}

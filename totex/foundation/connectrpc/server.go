package connectrpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yourusername/foundation/logging"
)

// Server represents a ConnectRPC HTTP server
type Server struct {
	name   string
	logger logging.Logger
	mux    *http.ServeMux
	addr   string
	server *http.Server
}

// NewServer creates a new ConnectRPC server
func NewServer(name, addr string, logger logging.Logger) *Server {
	return &Server{
		name:   name,
		logger: logger,
		mux:    http.NewServeMux(),
		addr:   addr,
	}
}

// RegisterHandler registers a ConnectRPC handler
func (s *Server) RegisterHandler(path string, handler interface{}) error {
	h, ok := handler.(http.Handler)
	if !ok {
		s.logger.Error("Handler does not implement http.Handler", "path", path)
		return fmt.Errorf("handler for %s does not implement http.Handler", path)
	}
	s.mux.Handle(path, h)
	s.logger.Info("Registered handler", "path", path)
	return nil
}

// GetHandler returns the underlying http.Handler
func (s *Server) GetHandler() http.Handler {
	return s.mux
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting ConnectRPC HTTP server", "address", s.addr)
	s.server = &http.Server{
		Addr:    s.addr,
		Handler: s.mux,
	}
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("HTTP server error", "error", err)
		}
	}()
	return nil
}

// Stop stops the HTTP server gracefully
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Stopping ConnectRPC HTTP server", "address", s.addr)
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// Name returns the server name
func (s *Server) Name() string { return s.name }

package foundation

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/viper"
	"github.com/yourusername/foundation/logging"
	"github.com/yourusername/foundation/metrics"
	"github.com/yourusername/foundation/tracing"
)

type Span = tracing.Span
type SpanOption = tracing.SpanOption

// ConnectRPCServer interface for gRPC/Connect-RPC server
type ConnectRPCServer interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
	RegisterHandler(path string, handler interface{}) error
	GetHandler() interface{}
}

// Service represents a lifecycle-managed component
// (e.g., servers) that can be started and stopped.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

// App represents the main application with all services
type App struct {
	logger           logging.Logger
	tracer           tracing.Tracer
	metrics          metrics.Metrics
	connectRPCServer ConnectRPCServer

	services []Service
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
}

// Config holds configuration for logger and other services
type Config struct {
	Logger LoggerConfig
	// Add other configs as needed
}

// LoggerConfig configuration for the logger
type LoggerConfig struct {
	Type   string
	Level  string
	Format string
	Output string
}

// LoadConfig loads configuration using viper and returns a Config struct
func LoadConfig() Config {
	viper.AutomaticEnv()
	viper.SetDefault("LOGGER_TYPE", "slog")
	viper.SetDefault("LOGGER_LEVEL", "info")
	viper.SetDefault("LOGGER_FORMAT", "text")
	viper.SetDefault("LOGGER_OUTPUT", "stdout")

	return Config{
		Logger: LoggerConfig{
			Type:   viper.GetString("LOGGER_TYPE"),
			Level:  viper.GetString("LOGGER_LEVEL"),
			Format: viper.GetString("LOGGER_FORMAT"),
			Output: viper.GetString("LOGGER_OUTPUT"),
		},
	}
}

// NewWithConfig returns an App with logger, metrics, tracing, and connectRPC server using config
func NewWithConfig(cfg Config) *App {
	ctx, cancel := context.WithCancel(context.Background())
	logger := NewLoggerFromConfig(cfg.Logger)
	metrics := NewDefaultMetrics()
	tracer := NewDefaultTracer()
	server := NewDefaultConnectRPCServer(logger)

	app := &App{
		logger:           logger,
		tracer:           tracer,
		metrics:          metrics,
		connectRPCServer: server,
		ctx:              ctx,
		cancel:           cancel,
	}
	app.services = []Service{
		server,
	}
	return app
}

// New returns an App with default logger, metrics, tracing, and connectRPC server
func New() *App {
	return NewWithConfig(LoadConfig())
}

// Start starts all services
func (a *App) Start(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.logger.Info("Starting app")
	for _, service := range a.services {
		if err := service.Start(ctx); err != nil {
			a.logger.Error("Failed to start service", "service", service.Name(), "error", err)
			return fmt.Errorf("failed to start %s: %w", service.Name(), err)
		}
		a.logger.Info("Started service", "service", service.Name())
	}
	a.logger.Info("All services started successfully")
	return nil
}

// Stop stops all services gracefully
func (a *App) Stop(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.logger.Info("Stopping app")
	a.cancel()
	for i := len(a.services) - 1; i >= 0; i-- {
		service := a.services[i]
		if err := service.Stop(ctx); err != nil {
			a.logger.Error("Failed to stop service", "service", service.Name(), "error", err)
		} else {
			a.logger.Info("Stopped service", "service", service.Name())
		}
	}
	a.logger.Info("All services stopped")
	return nil
}

// Logger returns the logger
func (a *App) Logger() logging.Logger { return a.logger }

// Metrics returns the metrics
func (a *App) Metrics() metrics.Metrics { return a.metrics }

// Tracer returns the tracer
func (a *App) Tracer() tracing.Tracer { return a.tracer }

// ConnectRPCServer returns the connectRPC server
func (a *App) ConnectRPCServer() ConnectRPCServer { return a.connectRPCServer }

// --- Default Implementations ---

// NewDefaultLogger returns a stub logger (slog-based)
func NewDefaultLogger() logging.Logger {
	return &SlogLogger{name: "default-logger", slog: slog.Default()}
}

// NewLoggerFromConfig creates a logger using LoggerConfig (currently only slog, extend as needed)
func NewLoggerFromConfig(cfg LoggerConfig) logging.Logger {
	// For demonstration, only slog is implemented. Extend for logrus/zap as needed.
	var h slog.Handler
	if cfg.Format == "json" {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	} else {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	}
	l := slog.New(h)
	return &SlogLogger{name: "configured-logger", slog: l}
}

type SlogLogger struct {
	name string
	slog *slog.Logger
}

func (l *SlogLogger) Debug(msg string, args ...any)   { l.slog.Debug(msg, args...) }
func (l *SlogLogger) Info(msg string, args ...any)    { l.slog.Info(msg, args...) }
func (l *SlogLogger) Warn(msg string, args ...any)    { l.slog.Warn(msg, args...) }
func (l *SlogLogger) Error(msg string, args ...any)   { l.slog.Error(msg, args...) }
func (l *SlogLogger) With(args ...any) logging.Logger { return l }
func (l *SlogLogger) Name() string                    { return l.name }

// NewDefaultMetrics returns a stub metrics implementation
func NewDefaultMetrics() metrics.Metrics {
	return &NoopMetrics{name: "default-metrics"}
}

type NoopMetrics struct{ name string }

func (m *NoopMetrics) Counter(name string, value float64, labels ...string)   {}
func (m *NoopMetrics) Gauge(name string, value float64, labels ...string)     {}
func (m *NoopMetrics) Histogram(name string, value float64, labels ...string) {}
func (m *NoopMetrics) Summary(name string, value float64, labels ...string)   {}
func (m *NoopMetrics) Start(ctx context.Context) error                        { return nil }
func (m *NoopMetrics) Stop(ctx context.Context) error                         { return nil }
func (m *NoopMetrics) Name() string                                           { return m.name }

// NewDefaultTracer returns a stub tracer implementation
func NewDefaultTracer() tracing.Tracer {
	return &NoopTracer{name: "default-tracer"}
}

type NoopTracer struct{ name string }

func (t *NoopTracer) StartSpan(name string, opts ...SpanOption) Span { return &NoopSpan{} }
func (t *NoopTracer) Inject(span Span, format interface{}, carrier interface{}) error {
	return nil
}
func (t *NoopTracer) Extract(format interface{}, carrier interface{}) (Span, error) {
	return &NoopSpan{}, nil
}
func (t *NoopTracer) Start(ctx context.Context) error { return nil }
func (t *NoopTracer) Stop(ctx context.Context) error  { return nil }
func (t *NoopTracer) Name() string                    { return t.name }

type NoopSpan struct{}

func (s *NoopSpan) SetTag(key, value string) {}
func (s *NoopSpan) SetError(err error)       {}
func (s *NoopSpan) Finish()                  {}
func (s *NoopSpan) Context() context.Context { return context.Background() }

// NewDefaultConnectRPCServer returns a real connectrpc server implementation
func NewDefaultConnectRPCServer(logger logging.Logger) ConnectRPCServer {
	return &HTTPConnectRPCServer{
		name:   "connectrpc-server",
		logger: logger,
		mux:    http.NewServeMux(),
		addr:   ":8080",
	}
}

type HTTPConnectRPCServer struct {
	name   string
	logger logging.Logger
	mux    *http.ServeMux
	addr   string
	server *http.Server
}

func (s *HTTPConnectRPCServer) RegisterHandler(path string, handler interface{}) error {
	h, ok := handler.(http.Handler)
	if !ok {
		s.logger.Error("Handler does not implement http.Handler", "path", path)
		return fmt.Errorf("handler for %s does not implement http.Handler", path)
	}
	s.mux.Handle(path, h)
	s.logger.Info("Registered handler", "path", path)
	return nil
}

func (s *HTTPConnectRPCServer) GetHandler() interface{} {
	return s.mux
}

func (s *HTTPConnectRPCServer) Start(ctx context.Context) error {
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

func (s *HTTPConnectRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping ConnectRPC HTTP server", "address", s.addr)
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *HTTPConnectRPCServer) Name() string { return s.name }

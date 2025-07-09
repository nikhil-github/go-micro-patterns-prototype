package foundation

import (
	"context"
	"fmt"
	"log/slog"
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
}

// App represents the main application with all services
type App struct {
	logger           logging.Logger
	tracer           tracing.Tracer
	metrics          metrics.Metrics
	connectRPCServer ConnectRPCServer

	services []interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
		Name() string
	}
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
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
	app.services = []interface {
		Start(context.Context) error
		Stop(context.Context) error
		Name() string
	}{
		logger, metrics, tracer, server,
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
		h = slog.NewJSONHandler(nil, &slog.HandlerOptions{})
	} else {
		h = slog.NewTextHandler(nil, &slog.HandlerOptions{})
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
func (l *SlogLogger) Start(ctx context.Context) error {
	l.Info("Logger started", "name", l.name)
	return nil
}
func (l *SlogLogger) Stop(ctx context.Context) error {
	l.Info("Logger stopped", "name", l.name)
	return nil
}
func (l *SlogLogger) Name() string { return l.name }

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

// NewDefaultConnectRPCServer returns a stub connectrpc server implementation
func NewDefaultConnectRPCServer(logger Logger) ConnectRPCServer {
	return &DummyConnectRPCServer{name: "default-connectrpc-server", logger: logger}
}

type DummyConnectRPCServer struct {
	name   string
	logger Logger
}

func (s *DummyConnectRPCServer) Start(ctx context.Context) error {
	s.logger.Info("ConnectRPC server started", "name", s.name)
	return nil
}
func (s *DummyConnectRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("ConnectRPC server stopped", "name", s.name)
	return nil
}
func (s *DummyConnectRPCServer) Name() string { return s.name }

// --- gRPC Server Factory Example ---

// GrpcServer is a minimal gRPC server interface
type GrpcServer interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

// NewDefaultGrpcServer returns a stub gRPC server implementation
func NewDefaultGrpcServer(logger Logger) GrpcServer {
	return &DummyGrpcServer{name: "default-grpc-server", logger: logger}
}

type DummyGrpcServer struct {
	name   string
	logger Logger
}

func (s *DummyGrpcServer) Start(ctx context.Context) error {
	s.logger.Info("gRPC server started", "name", s.name)
	return nil
}
func (s *DummyGrpcServer) Stop(ctx context.Context) error {
	s.logger.Info("gRPC server stopped", "name", s.name)
	return nil
}
func (s *DummyGrpcServer) Name() string { return s.name }

# goforge

## Overview
A comprehensive, interface-first approach for microservices with 8+ infrastructure components. This pattern uses the `App` struct as the main orchestrator, making it perfect for Kubernetes deployments and framework swapping.

## Key Features

### ✅ **Interface-First Design**
- All components defined as interfaces
- Easy to swap underlying frameworks
- Mock implementations for testing

### ✅ **App-Based Orchestration**
- Single `App` struct manages all services
- Initialized with name and version
- Dependency-aware startup/shutdown

### ✅ **Kubernetes Ready**
- Environment-based configuration
- Health checks and graceful shutdown
- Service discovery integration

### ✅ **Framework Agnostic**
- Swap logger: slog ↔ logrus ↔ zap
- Swap tracer: jaeger ↔ zipkin ↔ otel
- Swap broker: kafka ↔ rabbitmq ↔ nats
- Swap cache: redis ↔ memcached ↔ in-memory
- Swap database: postgres ↔ mysql ↔ mongodb

## Structure
```
patterns/separated-interfaces/
├── interfaces/           # All service interfaces
├── config/              # Configuration management
├── services/            # Service factories + mocks
├── app/                 # Main App orchestrator
└── main.go             # Usage example
```

## Usage

### Basic Usage
```go
// Create app with name and version
app := app.New("user-service", "1.0.0")

// Initialize all services
if err := app.Init(); err != nil {
    log.Fatalf("Failed to initialize app: %v", err)
}

// Start all services
if err := app.Start(ctx); err != nil {
    log.Fatalf("Failed to start app: %v", err)
}

// Use services
logger := app.Logger()
broker := app.Broker()
cache := app.Cache()

// Graceful shutdown
app.Stop(ctx)
```

### Environment Configuration
```bash
# App metadata
export APP_NAME="user-service"
export APP_VERSION="1.0.0"
export APP_ENV="production"

# Logger
export LOGGER_LEVEL="info"
export LOGGER_FORMAT="json"
export LOGGER_OUTPUT="stdout"

# Tracer
export TRACER_TYPE="jaeger"
export TRACER_ENDPOINT="http://jaeger:14268"
export TRACER_SERVICE_NAME="user-service"

# Metrics
export METRICS_TYPE="prometheus"
export METRICS_PORT="9090"
export METRICS_PATH="/metrics"

# Service Discovery
export SERVICE_DISCOVERY_TYPE="consul"
export SERVICE_DISCOVERY_ENDPOINT="http://consul:8500"

# Broker
export BROKER_TYPE="kafka"
export BROKER_BROKERS="kafka1:9092,kafka2:9092"

# Cache
export CACHE_TYPE="redis"
export CACHE_ADDRESS="redis:6379"
export CACHE_TTL="1h"

# Database
export DATABASE_TYPE="postgres"
export DATABASE_DSN="postgres://user:pass@db:5432/mydb"

# ConnectRPC
export CONNECTRPC_ADDRESS=":8080"
```

## Interface Definitions

### Core Service Interface
```go
type Service interface {
    Start(ctx context.Context) error
    Stop(ctx context.Context) error
    Name() string
}
```

### Component Interfaces
- **Logger**: Structured logging with levels and formatting
- **Tracer**: Distributed tracing with spans and context propagation
- **Metrics**: Observability with counters, gauges, histograms
- **ServiceDiscovery**: Service registration and discovery
- **Broker**: Message queuing and pub/sub
- **Cache**: Key-value caching with TTL
- **Database**: Data persistence with health checks
- **ConnectRPCServer**: gRPC/Connect-RPC server

## Framework Swapping Examples

### Logger Swapping
```go
// In LoggerFactory.Create()
switch cfg.Type {
case "slog":
    return slog.New(slog.NewJSONHandler(os.Stdout, nil))
case "logrus":
    return logrus.New()
case "zap":
    return zap.NewProduction()
}
```

### Tracer Swapping
```go
// In TracerFactory.Create()
switch cfg.Type {
case "jaeger":
    return jaeger.NewTracer(cfg.Endpoint)
case "zipkin":
    return zipkin.NewTracer(cfg.Endpoint)
case "otel":
    return otel.NewTracer(cfg.Endpoint)
}
```

### Broker Swapping
```go
// In BrokerFactory.Create()
switch cfg.Type {
case "kafka":
    return kafka.NewBroker(cfg.Brokers)
case "rabbitmq":
    return rabbitmq.NewBroker(cfg.Endpoint)
case "nats":
    return nats.NewBroker(cfg.Endpoint)
}
```

## Kubernetes Deployment

### Dockerfile Example
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### Kubernetes ConfigMap
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: user-service-config
data:
  APP_NAME: "user-service"
  APP_VERSION: "1.0.0"
  LOGGER_LEVEL: "info"
  LOGGER_FORMAT: "json"
  TRACER_TYPE: "jaeger"
  METRICS_TYPE: "prometheus"
  BROKER_TYPE: "kafka"
  CACHE_TYPE: "redis"
  DATABASE_TYPE: "postgres"
```

### Kubernetes Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: user-service
  template:
    metadata:
      labels:
        app: user-service
    spec:
      containers:
      - name: user-service
        image: user-service:latest
        ports:
        - containerPort: 8080
        - containerPort: 9090
        envFrom:
        - configMapRef:
            name: user-service-config
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
```

## Testing Strategy

### Unit Testing
```go
func TestAppInitialization(t *testing.T) {
    app := app.New("test-service", "1.0.0")
    assert.NoError(t, app.Init())
    assert.NotNil(t, app.Logger())
    assert.NotNil(t, app.Tracer())
}

func TestServiceDependencies(t *testing.T) {
    app := app.New("test-service", "1.0.0")
    app.Init()
    
    // Test that services can be accessed
    logger := app.Logger()
    tracer := app.Tracer()
    
    assert.NotNil(t, logger)
    assert.NotNil(t, tracer)
}
```

### Integration Testing
```go
func TestAppLifecycle(t *testing.T) {
    app := app.New("test-service", "1.0.0")
    app.Init()
    
    ctx := context.Background()
    assert.NoError(t, app.Start(ctx))
    assert.NoError(t, app.Stop(ctx))
}
```

## Benefits for 4 Microservices

### ✅ **Consistency**
- Same App structure across all 4 services
- Consistent configuration patterns
- Uniform service lifecycle

### ✅ **Maintainability**
- Interface-first design makes changes easy
- Clear separation of concerns
- Easy to test individual components

### ✅ **Scalability**
- Easy to add new services
- Framework swapping without code changes
- Kubernetes-ready deployment

### ✅ **Team Collaboration**
- Different teams can work on different components
- Clear interfaces for integration
- Mock implementations for parallel development

## Comparison with Simplified Pattern

| Aspect | Simplified (2 components) | Separated Interfaces (8+ components) |
|--------|---------------------------|--------------------------------------|
| **Complexity** | Low | Medium |
| **Framework Swapping** | Hard | Easy |
| **Testing** | Good | Excellent |
| **Kubernetes Ready** | Basic | Full |
| **Team Collaboration** | Limited | High |
| **Maintenance** | Good | Excellent |

## Conclusion

For your 4 microservices with 8+ infrastructure components, the **Separated Interfaces Pattern** is the clear winner because:

1. **Interface-first design** makes framework swapping trivial
2. **App-based orchestration** provides clean, consistent API
3. **Kubernetes-ready** with environment-based configuration
4. **Excellent testability** with mock implementations
5. **Scalable architecture** that grows with your needs

The simplified pattern would become unwieldy with 8+ components, while this pattern is designed exactly for complex, infrastructure-heavy microservices. 

The folder structure the shared-foundation has

shared-foundation/
├── types.go                        # All interfaces and types in one file
├── config.go                       # Configuration
├── app.go                          # Main App orchestrator
├── factory.go                      # Service factories
├── mocks.go                        # Mock implementations for testing
├── logger/                         # Logger package
│   ├── slog.go
│   ├── logrus.go
│   └── slog_test.go
├── tracer/                         # Tracer package
│   ├── jaeger.go
│   ├── zipkin.go
│   └── jaeger_test.go
├── metrics/                        # Metrics package
│   ├── prometheus.go
│   ├── statsd.go
│   └── prometheus_test.go
├── broker/                         # Broker package
│   ├── kafka.go
│   ├── rabbitmq.go
│   ├── nats.go
│   └── kafka_test.go
├── cache/                          # Cache package
│   ├── redis.go
│   ├── memcached.go
│   └── redis_test.go
├── database/                       # Database package
│   ├── postgres.go
│   ├── mysql.go
│   └── postgres_test.go
├── connectrpc/                     # ConnectRPC package
│   ├── server.go
│   └── server_test.go
└── examples/                       # Usage examples
# Separated Concerns: Modular Components

## Overview
Separate, focused components for each responsibility: config management, service creation, and lifecycle orchestration. This pattern follows the single responsibility principle and is designed for maintainable, testable microservices.

## Structure
```
patterns/separated/
├── config/             # Configuration management
├── services/           # Service factory and implementations
├── lifecycle/          # Lifecycle orchestrator
└── main.go            # Usage example
```

## Pros
- **Clear separation of concerns** - each component has one job
- **Easy to test individual components** - test config, factory, and orchestrator separately
- **Highly modular and reusable** - use components across different microservices
- **Explicit dependencies** - clear what depends on what
- **Full configuration management** - viper-based config with environment variables
- **Easy to extend** - add new services without modifying existing components
- **Consistent across microservices** - same patterns everywhere

## Cons
- More boilerplate code
- Multiple objects to manage
- Need to understand multiple components
- Higher learning curve initially

## Usage
```go
// Config management
cfg := config.New().Load()

// Service creation
factory := services.NewFactory(logger)
service := factory.CreateConnectRPCServer(cfg.ConnectRPC, handler)

// Lifecycle management
orchestrator := lifecycle.NewOrchestrator(logger)
orchestrator.Add(service)
orchestrator.Start()
orchestrator.WaitForShutdown()
```

## Comparison with Current Pattern

| Aspect | Separated | Current |
|--------|-----------|---------|
| **API Complexity** | 3 focused APIs | 8 methods on one object |
| **Testability** | Easy - test each component | Hard - everything coupled |
| **Extensibility** | Easy - add new components | Hard - modify factory |
| **Dependencies** | Explicit - manual registration | Hidden - auto-registration |
| **Learning Curve** | High - need to learn 3 components | Medium - need to learn factory API |
| **Maintenance** | Easy - focused responsibilities | Hard - all logic in one place |
| **Reusability** | High - mix and match components | Low - tied to factory |
| **Multi-service consistency** | High - same patterns everywhere | Low - each service might differ |

## Benefits for Multiple Microservices

### Consistency
- Same config management across all services
- Same service creation patterns
- Same lifecycle management
- Easy to enforce standards

### Reusability
- Share config components between services
- Reuse service factories
- Common orchestrator patterns
- DRY principle across services

### Maintainability
- Update config logic in one place
- Add new services without touching existing code
- Test components independently
- Clear separation makes debugging easier

### Scalability
- Easy to add new service types
- Simple to extend configuration
- Modular approach scales with complexity
- Team can work on different components

## When to Use Separated Pattern
- Complex microservices with many dependencies
- Production systems requiring high testability
- When you need to reuse components across services
- When you prefer explicit control and dependencies
- Multiple microservices that need consistency
- Teams working on different parts of the system 
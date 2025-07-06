# Current Pattern: Comprehensive ServiceFactory

## Overview
A single ServiceFactory that handles everything: logger initialization, config management, service creation, context management, and lifecycle orchestration. This is a "batteries-included" approach that tries to do everything in one place.

## Structure
```
patterns/current/
├── servicefactory/     # Single comprehensive factory
├── connectrpc/         # Service implementation
└── main.go            # Usage example
```

## Pros
- Single entry point for everything
- Automatic service registration
- Built-in context management
- Graceful shutdown handling
- Less boilerplate in main.go

## Cons
- **Mixed responsibilities** - violates single responsibility principle
- **Complex API** - multiple steps to remember (Build, Create, StartAll, WaitForShutdown)
- **Hard to test** - can't test individual components in isolation
- **Hidden dependencies** - auto-registration makes dependencies unclear
- **Tight coupling** - everything depends on the factory
- **Difficult to extend** - adding new services requires modifying the factory

## Usage
```go
factory := servicefactory.NewServiceFactory()
factory.Build()
factory.CreateConnectRPCServer(handler)
factory.StartAll()
factory.WaitForShutdown()
```

## Comparison with Separated Pattern

| Aspect | Current | Separated |
|--------|---------|-----------|
| **API Complexity** | 8 methods on one object | 3 focused APIs |
| **Testability** | Hard - everything coupled | Easy - test each component |
| **Extensibility** | Hard - modify factory | Easy - add new components |
| **Dependencies** | Hidden - auto-registration | Explicit - manual registration |
| **Learning Curve** | Medium - need to learn factory API | High - need to learn 3 components |
| **Maintenance** | Hard - all logic in one place | Easy - focused responsibilities |
| **Reusability** | Low - tied to factory | High - mix and match components |

## When to Use Current Pattern
- Simple microservices with few dependencies
- Quick prototyping
- When you prefer less boilerplate over flexibility
- When you don't need fine-grained control

## When to Use Separated Pattern
- Complex microservices with many dependencies
- Production systems requiring high testability
- When you need to reuse components across services
- When you prefer explicit control and dependencies 
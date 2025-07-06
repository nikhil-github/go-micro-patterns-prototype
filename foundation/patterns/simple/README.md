# Simple: Direct Service Creation + Orchestrator

## Overview
The most minimal approach possible - create services directly and use a simple orchestrator for lifecycle management. No factories, no config management, just the essentials.

## Structure
```
patterns/simple/
├── orchestrator/       # Lifecycle management only
├── services/           # Direct service implementations
└── main.go            # Usage example
```

## Pros
- Extremely simple and minimal
- No abstractions or factories
- Direct service creation
- Easy to understand and use
- Minimal boilerplate
- Clear responsibilities

## Cons
- No configuration management
- Manual service creation
- Less reusable
- Hardcoded values

## Usage
```go
// Create service directly (no factory needed)
service := services.NewConnectRPCServer(":8080", handler, logger)

// Simple orchestrator for lifecycle
orchestrator := orchestrator.New(logger)
orchestrator.Add(service)
orchestrator.Start()
orchestrator.WaitForShutdown()
```

## Trade-offs
- **Simplicity**: Extremely simple to understand
- **Flexibility**: Limited - hardcoded values
- **Testability**: Good - simple components
- **Maintainability**: Good - minimal and focused 
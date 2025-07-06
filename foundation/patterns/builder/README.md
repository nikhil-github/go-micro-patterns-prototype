# Builder Pattern: Method Chaining

## Overview
A builder pattern with fluent interface that allows explicit configuration and method chaining for a more readable API.

## Structure
```
patterns/builder/
├── builder/            # Builder with method chaining
├── config/             # Configuration management
├── services/           # Service implementations
├── lifecycle/          # Lifecycle management
└── main.go            # Usage example
```

## Pros
- Explicit and readable API
- Method chaining for better flow
- Clear separation of configuration and execution
- Easy to understand what's happening

## Cons
- More complex implementation
- Multiple objects to manage
- Potential for long chains

## Usage
```go
app := builder.NewMicroservice().
    WithLogger(logger).
    WithConfig(config).
    AddConnectRPCServer(handler).
    AddDatabase(dbConfig).
    Build().
    Start()

app.WaitForShutdown()
```

## Trade-offs
- **Simplicity**: Clear, readable API
- **Flexibility**: Good - explicit configuration
- **Testability**: Better - can test each step
- **Maintainability**: Separated concerns 
# Simplified Pattern: Focused ServiceFactory

## Overview
A simplified version of the current pattern that addresses testability concerns while maintaining the single entry point benefit. This pattern separates configuration management from service creation while keeping the factory focused on its core responsibilities.

## Structure
```
patterns/simplified/
├── config/             # Configuration management (separated)
├── servicefactory/     # Focused service factory
└── main.go            # Usage example
```

## Key Improvements Over Current Pattern

### 1. **Separated Configuration**
- Config management is now a separate concern
- Factory receives config as explicit dependency
- Easy to test config loading independently
- **Logger configuration via environment variables**

### 2. **Explicit Dependencies**
- Factory constructor takes config and logger explicitly
- No hidden dependencies or auto-initialization
- Clear what the factory depends on

### 3. **Simplified API**
- Removed complex multi-step initialization
- Single `Run()` method for start + shutdown
- Clearer service registration

### 4. **Better Testability**
- Can test config loading separately
- Can test factory with mock config
- Can test individual service creation
- No hidden state or auto-registration

## Usage
```go
// Load config (separated concern)
cfg := config.Load()

// Create logger from config
logger := cfg.CreateLogger()

// Create factory with explicit dependencies
factory := servicefactory.New(cfg, logger)

// Add services explicitly
handler := servicefactory.NewDummyHandler()
factory.AddConnectRPCServer(handler)

// Run everything (start + wait for shutdown)
factory.Run()
```

## Environment Variables

### ConnectRPC Configuration
- `CONNECTRPC_ADDRESS`: Server address (default: ":8080")

### Logger Configuration
- `LOGGER_LEVEL`: Log level - debug, info, warn, error (default: "info")
- `LOGGER_FORMAT`: Log format - text, json (default: "text")
- `LOGGER_OUTPUT`: Log output - stdout, stderr, or file path (default: "stdout")

### Example Environment Setup
```bash
# ConnectRPC
export CONNECTRPC_ADDRESS=":9090"

# Logger
export LOGGER_LEVEL="debug"
export LOGGER_FORMAT="json"
export LOGGER_OUTPUT="/var/log/myapp.log"
```

## Comparison with Current Pattern

| Aspect | Current | Simplified |
|--------|---------|------------|
| **API Complexity** | 8 methods, multi-step | 4 methods, single-step |
| **Dependencies** | Hidden, auto-initialized | Explicit, injected |
| **Config Management** | Mixed in factory | Separated concern |
| **Logger Configuration** | None | Full env-based config |
| **Testability** | Hard - everything coupled | Easy - test each part |
| **Learning Curve** | Medium - complex API | Low - simple API |
| **Maintenance** | Hard - mixed concerns | Easy - focused concerns |
| **Flexibility** | Low - tied to factory | Medium - config separate |

## Benefits

### Testability
```go
// Test config separately
func TestConfig(t *testing.T) {
    cfg := config.Load()
    assert.Equal(t, ":8080", cfg.ConnectRPC.Address)
}

// Test logger configuration
func TestLoggerConfig(t *testing.T) {
    cfg := &config.Config{
        Logger: config.Logger{
            Level: "debug",
            Format: "json",
            Output: "stdout",
        },
    }
    logger := cfg.CreateLogger()
    assert.NotNil(t, logger)
}

// Test factory with mock config
func TestFactory(t *testing.T) {
    cfg := &config.Config{ConnectRPC: config.ConnectRPC{Address: ":8080"}}
    factory := servicefactory.New(cfg, logger)
    // Test factory methods...
}
```

### Maintainability
- Config changes don't affect factory
- Factory focuses only on service lifecycle
- Clear separation of concerns
- Easy to understand and modify
- Logger configuration is centralized

### Flexibility
- Can use different config sources
- Can inject mock config for testing
- Can extend config without touching factory
- Can reuse config across different factories
- Full logger customization via environment

## When to Use Simplified Pattern
- You want the benefits of current pattern but better testability
- You prefer explicit dependencies over hidden ones
- You want to separate config management from service creation
- You need to test individual components
- You're building production systems that need maintainability
- You need configurable logging via environment variables

## Trade-offs
- **Pros**: Better testability, explicit dependencies, separated concerns, configurable logging
- **Cons**: Slightly more boilerplate than current pattern, still single factory 
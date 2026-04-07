---
name: "error-handling"
description: "Return errors up, no log-and-return, %w wrapping"
domain: "quality"
confidence: "high"
source: "hyoka/internal/eval/engine.go, hyoka/internal/config/config.go patterns"
---

## Context

Hyoka follows Go's idiomatic error handling pattern: errors bubble up the call stack, and are logged at the top level (in CLI commands). This separation of concerns makes the codebase testable and keeps error flows explicit.

## Patterns

### Return Errors Up (No Log-and-Return)

**✓ Correct:**
```go
// In internal function
func loadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("load config: %w", err)  // Wrap and return
    }
    // ...
}

// In CLI command (top level)
func Execute() error {
    cfg, err := loadConfig(cfgPath)
    if err != nil {
        slog.Error("Failed to load config", "error", err)  // Log once at top
        return err
    }
    // ...
}
```

**✗ Incorrect:**
```go
// Internal function logs AND returns
func loadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        slog.Error("Failed to read config", "error", err)  // Wrong: logging here
        return nil, err
    }
}
```

### Error Wrapping with `%w`

Always wrap lower-level errors with `%w` to preserve the error chain. This enables `errors.Is()` and `errors.As()` in caller code.

**✓ Correct:**
```go
if err != nil {
    return fmt.Errorf("failed to build project: %w", err)
}

// Caller can inspect the chain
if errors.Is(err, os.ErrNotExist) {
    // Handle missing file
}
```

**✗ Incorrect:**
```go
if err != nil {
    return fmt.Errorf("failed to build project: %v", err)  // Loses error chain
}
```

### Context in Error Messages

Provide actionable context at the point where error is wrapped, not repeated at logging.

**✓ Correct:**
```go
// In eval engine
if err := runGrader(ctx, grader); err != nil {
    return fmt.Errorf("grade with %q: %w", grader.Name(), err)
}

// In CLI command
if err != nil {
    slog.Error("Evaluation failed", "error", err)  // Message says what failed, details in error chain
}
```

## Anti-Patterns

- Logging at multiple levels (log once at top-level CLI)
- Using `%v` instead of `%w` for error wrapping
- Ignoring errors silently with `_ = someFunc()`
- Returning `nil` error when error occurred
- Logging and then returning without wrapping

## Related Code Locations

- **Error wrapping examples:** `hyoka/internal/eval/engine.go`
- **CLI error handling:** `hyoka/cmd/*.go`
- **Logging configuration:** `hyoka/internal/logging/`

---
name: "logging-conventions"
description: "slog usage, no third-party logging"
domain: "quality"
confidence: "high"
source: "hyoka/internal/logging/logging.go, all internal packages"
---

## Context

Hyoka uses Go's standard `log/slog` package (available since Go 1.21). No third-party logging libraries are used. Logging follows a simple hierarchy: diagnostic logs go to slog, user-facing output goes to stdout/stderr.

## slog Basics

```go
import "log/slog"

// Debug: development/troubleshooting info
slog.Debug("Session started", "session_id", sessID, "model", model)

// Info: important events users might care about
slog.Info("Evaluation complete", "prompt_id", id, "duration_s", elapsed)

// Warn: recoverable issues (grader timeout, build skipped)
slog.Warn("Grader timeout", "grader_name", name, "duration_s", 30)

// Error: problems that don't stop evaluation (reviewer model unavailable)
slog.Error("Failed to load skill", "path", p, "error", err)
```

## Initialization

slog is typically initialized in `main()` or in the logging package:

```go
// main.go or logging.go
package main

import (
    "log/slog"
    "os"
)

func init() {
    // Default: JSON output to stderr
    handler := slog.NewJSONHandler(os.Stderr, nil)
    slog.SetDefault(slog.New(handler))
}
```

## Log Levels

Control verbosity with `--log-level` flag:

```bash
go run ./hyoka run --prompt-id ... --log-level debug    # Very verbose
go run ./hyoka run --prompt-id ... --log-level info     # Important events
go run ./hyoka run --prompt-id ... --log-level warn     # Warnings only
```

Mapping (from least to most verbose):
1. **ERROR** — Errors only
2. **WARN** — Warnings + errors
3. **INFO** — Important events + warnings + errors
4. **DEBUG** — Development diagnostics + everything else

## Structured Attributes

Always use key-value pairs, not formatted strings:

**✓ Correct:**
```go
slog.Info("Evaluation started",
    "prompt_id", promptID,
    "config", configName,
    "model", model,
)
```

**✗ Incorrect:**
```go
slog.Info(fmt.Sprintf("Evaluation %s with config %s using model %s",
    promptID, configName, model))
```

Structured attributes enable filtering and aggregation in log analysis tools.

## Logging Pattern: Errors

Always log at the point where you can provide context:

```go
// Load config
cfg, err := loadConfig(cfgPath)
if err != nil {
    // Log with context
    slog.Error("Failed to load config",
        "path", cfgPath,
        "error", err,
    )
    return err
}

// Process config (no logging here — error means we already logged above)
if err := processConfig(cfg); err != nil {
    slog.Error("Failed to process config",
        "config_name", cfg.Name,
        "error", err,
    )
    return err
}
```

## Logging Pattern: Lifecycle Events

Use **Info** for important milestones:

```go
func (e *Engine) Run(ctx context.Context, prompt *Prompt, cfg *Config) (*Result, error) {
    slog.Info("Evaluation starting",
        "prompt_id", prompt.ID,
        "config", cfg.Name,
    )
    defer func(start time.Time) {
        duration := time.Since(start)
        slog.Info("Evaluation complete",
            "prompt_id", prompt.ID,
            "duration_s", duration.Seconds(),
        )
    }(time.Now())
    
    // ... evaluation logic
}
```

## Logging Pattern: Debug Traces

Use **Debug** for developer troubleshooting:

```go
// generation phase
slog.Debug("Copilot session created",
    "session_id", sess.ID(),
    "model", cfg.Generator.Model,
)

slog.Debug("Action captured",
    "type", action.Type,
    "tool", action.Tool,
    "duration_ms", action.DurationMs,
)

// review phase
slog.Debug("Review panel score",
    "reviewer_model", model,
    "score", score,
)
```

## Log File Output

When `--log-file` is specified, slog output is both written to the file and echoed to stderr:

```bash
go run ./hyoka run ... --log-file hyoka-debug.log
```

The file captures full debug output; users see warnings and errors in console.

## slog Configuration (if needed)

For custom output formatting:

```go
import (
    "log/slog"
    "os"
)

// Text format (instead of JSON)
handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})
slog.SetDefault(slog.New(handler))

// Or with pretty-printing
handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    Level:     slog.LevelDebug,
    AddSource: true,  // Include file:line
})
```

## Anti-Patterns

- Logging errors AND returning them with same context (return once, log once)
- Using `log.Printf` or `fmt.Printf` for diagnostic output (use slog)
- Logging user output (progress, results) — write to stdout directly
- String formatting in log arguments (use structured attributes)
- Logging secrets or sensitive data
- `slog.Error()` for non-critical issues (use `slog.Warn()`)

## Related Code Locations

- **slog setup:** `hyoka/internal/logging/logging.go`
- **CLI log-level flag:** `hyoka/cmd/root.go`
- **Example logging patterns:** `hyoka/internal/eval/engine.go`

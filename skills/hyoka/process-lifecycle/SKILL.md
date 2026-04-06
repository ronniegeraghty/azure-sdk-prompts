---
name: "process-lifecycle"
description: "Session creation, workspace isolation, cleanup"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/eval/workspace.go, hyoka/internal/eval/copilot.go, hyoka/internal/clean/"
---

## Context

Each hyoka evaluation runs in an isolated workspace with its own session. Understanding process lifecycle is critical for debugging eval failures, preventing resource leaks, and writing tests.

## Workspace Lifecycle

### Creation

When an evaluation starts:

```go
// Create workspace directory
ws, err := NewWorkspace(ctx, "eval-run-123")
if err != nil {
    return fmt.Errorf("create workspace: %w", err)
}
defer ws.Close()  // Cleanup deferred

// Start Copilot session in workspace
sess, err := copilot.NewSession(ctx, &SessionOptions{
    WorkingDir: ws.Path(),
    Model:      "claude-opus-4.6",
    // ...
})
```

Workspace is created as a temporary directory:
- **Linux/Mac:** `/tmp/hyoka-{run_id}-XXXXXX`
- **Windows:** `%TEMP%\hyoka-{run_id}-XXXXXX`
- Contains all generated files, logs, and session metadata

### Session Initialization

Copilot SDK launches the agent process:

```
┌─────────────────────────────────────┐
│ hyoka (CLI process, PID 1234)       │
└──────────────┬──────────────────────┘
               │
               ├─→ Workspace: /tmp/hyoka-eval-1
               │
               └─→ Copilot process (PID 1235) [child]
                   ├─→ MCP server (PID 1236) [grandchild]
                   └─→ Tool subprocess (PID 1237) [grandchild]
```

The CLI tracks all child PIDs for cleanup on signal.

### File Operations

Agent writes generated code to workspace:

```
/tmp/hyoka-eval-1/
  main.py           # Generated code
  test_main.py      # Generated tests
  Copilot.log       # Session log
  action_timeline.json  # Captured actions
```

The CLI captures all file operations in the action timeline.

### Cleanup Phase

When evaluation completes (success or failure):

```go
// Workspace cleanup
if err := ws.Close(); err != nil {
    slog.Warn("Workspace cleanup failed", "path", ws.Path(), "error", err)
}

// Copilot session cleanup
if err := sess.Close(); err != nil {
    slog.Warn("Session cleanup failed", "error", err)
}

// Orphaned process cleanup (if signal received)
// kill all child PIDs
```

Cleanup removes:
- Workspace directory + all files
- Copilot process + child processes
- Temporary MCP servers

## Process Tracking

Hyoka tracks child processes to ensure cleanup on interruption:

```go
// hyoka/internal/eval/proctracker.go
tracker := NewProcessTracker()

// Register Copilot process
tracker.Add(sess.PID())

// On SIGINT, cleanup
<-sigChan
tracker.KillAll()  // Kills all tracked processes
```

### PID Files

For long-running evaluations, PIDs are written to files:

```
reports/{run_id}/
  pids.txt     # List of tracked PIDs
```

If eval crashes, `hyoka clean` reads these files and kills orphaned processes.

## Session State

### Before Generation
```
Created ✓
Tools loaded ✓
MCP servers started ✓
Skills injected ✓
Ready for agent ✓
```

### During Generation
```
Agent running...
Tool calls executed...
Files written...
Actions captured...
```

### After Generation
```
Agent stopped ✓
Build phase (optional)
Review phase (if enabled)
Workspace preserved (for debugging)
```

### Cleanup
```
Workspace deleted ✓
Session closed ✓
Processes killed ✓
```

## Workspace Isolation

Each eval is isolated to prevent:
- **State leakage:** Code from eval 1 not visible to eval 2
- **File conflicts:** Two evals writing to same file
- **Process conflicts:** Child processes from eval 1 interfering with eval 2

This is achieved by:
- Unique workspace directories (per run ID)
- Separate Copilot sessions (each with own process tree)
- Deferred cleanup (filesystem cleanup guaranteed even on panic)

## Cleanup Command

```bash
go run ./hyoka clean
```

This command:
1. Finds all abandoned sessions (workspace dirs left on disk)
2. Reads PIDs from `reports/{run_id}/pids.txt`
3. Kills orphaned processes
4. Removes workspace directories
5. Reports cleaned sessions

Useful after:
- Interrupted eval runs
- Process crashes
- Development/testing (always run before committing)

## Session Limits

Per-evaluation guardrails prevent runaway sessions:

```yaml
limits:
  max_turns: 25                    # Max agent turns
  max_files: 50                    # Max files created
  max_output_size: 1048576         # Max output (1 MB)
  max_session_actions: 50          # Max tool calls
```

When limit exceeded:
```
Generation phase stops → Session killed → Error reported
```

## Testing with Temporary Workspaces

In tests, use stub workspaces:

```go
// Stub workspace that doesn't create disk files
type stubWorkspace struct {
    files map[string]string  // In-memory files
}

func (sw *stubWorkspace) WriteFile(path, content string) error {
    sw.files[path] = content
    return nil
}
```

Never create real workspaces in tests (use stubs or temp directories).

## Code Locations

- **Workspace management:** `hyoka/internal/eval/workspace.go`
- **Process tracking:** `hyoka/internal/eval/proctracker.go`
- **Cleanup command:** `hyoka/cmd/clean.go`
- **Session lifecycle:** `hyoka/internal/eval/copilot.go`

## Anti-Patterns

- Not deferring workspace cleanup (may leak on early return)
- Hardcoding workspace paths (use WorkingDir option)
- Not tracking child processes (orphans may remain)
- Assuming workspace persists after eval (it's deleted)
- Ignoring cleanup errors in tests (cleanup failures indicate leaks)

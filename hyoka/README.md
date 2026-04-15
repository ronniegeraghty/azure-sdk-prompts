# hyoka — Developer Guide

> Internal developer reference for the `hyoka` Go source. For user-facing
> documentation (installation, CLI usage, configuration), see the
> [root README](../README.md) and [docs/](../docs/).

## Package Architecture

The CLI entry point is `main.go` → `cmd.Execute()`. All domain logic lives
under `internal/`.

### `cmd/`

Cobra command definitions. Each file exports a `func xxxCmd() *cobra.Command`
that is registered in `root.go` via `root.AddCommand(xxxCmd())`.

### `internal/` packages

| Package | What it does |
|---------|-------------|
| `checkenv` | Verifies required toolchains, Copilot CLI, MCP servers, and auth |
| `clean` | Cleans stale Copilot session state, orphaned processes, and logs |
| `comparison` | Config-vs-config, run-vs-run, and temporal diff analysis |
| `config` | Loads and parses evaluation config YAML (generator, reviewer, tools) |
| `criteria` | Tiered evaluation criteria with conditional YAML rule matching |
| `eval` | Core evaluation engine — Copilot SDK session orchestration and workspace lifecycle |
| `graders` | Pluggable grading system (6 kinds) with weighted scoring and gate semantics |
| `logging` | Structured logging helpers on top of `log/slog` |
| `pairwise` | Tool-ablation test variant expansion (disable one tool per variant) |
| `pidfile` | Tracks SDK-spawned Copilot processes via PID files for orphan detection |
| `plugin` | Composable plugin system bundling skills, MCP servers, and hooks |
| `progress` | Progress display with live (ANSI), log, and off rendering modes |
| `prompt` | Loads and parses `.prompt.md` files with YAML frontmatter and filters |
| `report` | Generates JSON, HTML, and Markdown reports with aggregated statistics |
| `rerender` | Re-renders HTML/Markdown reports from existing `report.json` files |
| `review` | Multi-model LLM-as-judge code review via Copilot sessions |
| `serve` | Local web server + API for browsing reports (serves the `site/` SPA) |
| `trends` | Cross-run trend analysis (stable / improving / regressing / flaky) |
| `utils` | Shared utilities — file I/O, string helpers, workspace handling |
| `validate` | Prompt schema validation (required fields, enums, naming conventions) |

### `site/`

React + TypeScript SPA (Vite, Tailwind CSS). Proxies `/api` to the Go
server during development. See `site/package.json` for scripts.

## Build & Test

```bash
# From the repo root (uses go.work):

# Build
go build ./hyoka/...

# Test (always use -race)
go test -race ./hyoka/...

# Run the CLI
go run ./hyoka <command> [flags]

# Build the dashboard SPA
cd site && npm ci && npm run build
```

## How to Add a New Command

1. Create `hyoka/cmd/<name>.go`:

```go
package cmd

import "github.com/spf13/cobra"

func fooCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "foo",
        Short: "One-line description",
        RunE: func(cmd *cobra.Command, args []string) error {
            // implementation
            return nil
        },
    }
    cmd.Flags().StringP("bar", "b", "", "Flag description")
    return cmd
}
```

2. Register in `hyoka/cmd/root.go`:

```go
root.AddCommand(fooCmd())
```

3. Add a test in `hyoka/cmd/<name>_test.go` — at minimum verify the command
   creates without error and flag defaults are correct.

## How to Add a New Grader

Graders implement the `Grader` interface (`internal/graders/grader.go`):

```go
type Grader interface {
    Kind() string                                                    // e.g. "file", "program"
    Name() string                                                    // human-readable instance name
    Grade(ctx context.Context, input GraderInput) (GraderResult, error)
}
```

Steps:

1. **Define the grader** — create `internal/graders/<kind>.go` with a struct
   that satisfies `Grader`. Return a `GraderResult` with `Score` (0.0–1.0),
   `Pass`, and optionally typed `*<Kind>Details`.

2. **Define its config** — add a typed config struct and register it in
   `DecodeConfig()` inside `internal/graders/config.go`.

3. **Register in the factory** — add a `case` to the `switch` in
   `NewGrader()` (`internal/graders/registry.go`).

4. **Add the kind constant** — add `Kind<Name> = "<name>"` alongside the
   existing constants.

5. **Test** — add table-driven tests in `internal/graders/<kind>_test.go`.

Existing kinds: `file`, `program`, `prompt`, `behavior`, `action_sequence`,
`tool_constraint`.

## How to Add a New Internal Package

1. Create `hyoka/internal/<name>/` with a descriptive package name
   (singular, lowercase).
2. Export only what other packages need. Keep implementation details
   unexported.
3. Add `<name>_test.go` in the same directory — use table-driven tests
   (see [Testing Patterns](#testing-patterns)).
4. Do **not** introduce third-party logging — use `log/slog`.
5. Return errors up the call stack with `fmt.Errorf("context: %w", err)`.
   Never log-and-return.

## Debugging Tips

```bash
# Verbose logging — writes to both stderr and the log file
go run ./hyoka run --prompt-id <id> --config <cfg> \
    --log-level debug --log-file hyoka-debug.log

# Check for orphaned Copilot processes after a failed run
go run ./hyoka clean

# Verify environment prerequisites
go run ./hyoka check-env

# Dry-run to see which prompts match without executing
go run ./hyoka run --service storage --language python --dry-run

# Grep debug logs for role-tagged entries
grep "role=" hyoka-debug.log | head -20
```

## Testing Patterns

### Table-driven tests

```go
func TestParseFoo(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    Foo
        wantErr bool
    }{
        {name: "valid", input: "bar", want: Foo{Val: "bar"}},
        {name: "empty", input: "", wantErr: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParseFoo(tt.input)
            if (err != nil) != tt.wantErr {
                t.Fatalf("ParseFoo(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
            }
            if !tt.wantErr && got != tt.want {
                t.Errorf("ParseFoo(%q) = %v, want %v", tt.input, got, tt.want)
            }
        })
    }
}
```

### Race detection

Always run tests with `-race`:

```bash
go test -race ./hyoka/...
```

### Stubs and dependency injection

Use package-level function variables for external dependencies so tests can
override them:

```go
// production code
var runCommand = exec.Command

// test code
func TestFoo(t *testing.T) {
    orig := runCommand
    defer func() { runCommand = orig }()
    runCommand = func(name string, args ...string) *exec.Cmd {
        return exec.Command("echo", "stub")
    }
    // ...
}
```

For Cobra commands, redirect stdout/stderr to `io.Discard` to suppress
output during tests:

```go
cmd := fooCmd()
cmd.SetOut(io.Discard)
cmd.SetErr(io.Discard)
```

### Test fixtures

Place fixture files under `hyoka/testdata/` — Go's test runner automatically
excludes this directory from builds.

### Logger suppression

Suppress log output in tests that don't need it:

```go
func TestMain(m *testing.M) {
    slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
    os.Exit(m.Run())
}
```

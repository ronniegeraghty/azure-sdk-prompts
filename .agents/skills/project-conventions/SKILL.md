---
name: "project-conventions"
description: "Core conventions and patterns for this codebase"
domain: "project-conventions"
confidence: "high"
source: "go.mod, hyoka/cmd/, hyoka/internal/, AGENTS.md"
---

## Context

hyoka is a Go CLI tool that evaluates AI agent-generated Azure SDK code. These are the stable, high-level conventions that apply across the entire codebase.

## Patterns

### Language & Dependencies

- Go standard library preferred — minimize third-party dependencies
- CLI framework: `github.com/spf13/cobra` with `gopkg.in/yaml.v3` for config
- No other frameworks for HTTP (`net/http`), JSON (`encoding/json`), etc.

### Logging

- Use `log/slog` exclusively — no third-party logging libraries
- User-facing output (progress bars, results) goes to stdout/stderr directly
- Diagnostic/debug output goes through slog

### Error Handling

- Return errors up the call stack with `fmt.Errorf("context: %w", err)`
- Log errors once at the top-level handler — never log-and-return
- Never call `log.Fatal` in library code; only in `main()` or CLI command runners

### Testing

- Table-driven tests: define test cases as slices of structs, iterate with `t.Run`
- Always run with race detector: `go test -race ./...`
- Use stubs and interfaces for external dependencies (Copilot SDK, HTTP clients)
- Test files live alongside source: `foo.go` → `foo_test.go`

### CLI Conventions

- Cobra commands live in `hyoka/cmd/`
- Flags use kebab-case: `--log-level`, `--criteria-dir`, `--prompt-id`
- Config values use snake_case in YAML files
- `--config` takes the `name:` field from inside the YAML, not the filename

### Code Style

- Package names are short, lowercase, single-word (e.g., `eval`, `prompt`, `config`)
- Exported types have doc comments; unexported helpers may omit them
- No blank imports except for side-effect registrations

### File Structure

```
hyoka/
  main.go            # Entry point (delegates to cmd.Execute())
  cmd/               # Cobra commands
  internal/          # All internal packages
    eval/            # Evaluation engine
    config/          # Config loading
    prompt/          # Prompt parsing
    review/          # Multi-model review panel
    ...
configs/             # Evaluation config YAML files
prompts/             # Prompt library
criteria/            # Attribute-matched grader criteria
reports/             # Generated output (git-ignored)
```

### Build & Test Commands

```bash
go build ./...           # Build all packages
go test -race ./...      # Run all tests with race detector
go vet ./...             # Static analysis
```

### Git Conventions

- Branch from main: `{username}/issue-{N}-{short-description}`
- Commit trailers: `Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>`
- Use `ronniegeraghty` GitHub identity for pushes

## Anti-Patterns

- **No log-and-return** — either log the error and stop, or return it; never both
- **No third-party loggers** — only `log/slog`
- **No `log.Fatal` in library code** — return errors to callers
- **No camelCase CLI flags** — use `--kebab-case`
- **No inline struct fields for prompt metadata** — use the `Properties` map

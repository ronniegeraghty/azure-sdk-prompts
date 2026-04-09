# Contributing Guide

## Prerequisites

- Go 1.26.1+
- Node.js 18+ (for Copilot CLI)
- GitHub CLI (`gh`)
- Copilot CLI (`copilot`)

## Building

```bash
# From repo root (uses go.work)
cd /path/to/hyoka
go build ./hyoka/...
```

## Running Tests

```bash
# All tests
go test ./hyoka/...

# Specific package
go test ./hyoka/internal/eval/ -v

# With race detection
go test -race ./hyoka/...
```

## Project Structure

```
hyoka/              # Go module root
  main.go           # CLI entry point (13 lines, delegates to cmd.Execute())
  cmd/              # Cobra commands
    root.go         # Root command setup
    run.go          # hyoka run — evaluation orchestration
    list.go         # hyoka list — list prompts
    compare.go      # hyoka compare — compare results
    init.go         # hyoka init — scaffold .hyoka project
    validate.go     # hyoka validate — validate prompts
    check_env.go    # hyoka check-env — check prerequisites
    clean.go        # hyoka clean — cleanup orphaned processes
    serve.go        # hyoka serve — local web UI
    trends.go       # hyoka trends — cross-run trend analysis
    new_prompt.go   # hyoka new-prompt — scaffold prompt
    ...
  internal/
    build/          # Language-specific build verification
    checkenv/       # Environment prerequisite checks
    comparison/     # Config/run/temporal comparison logic
    config/         # YAML config loading and parsing
    criteria/       # Tiered evaluation criteria system
    eval/           # Evaluation engine, process tracker, resource monitor
    graders/        # Grader registry (6 types: builder, complexity, etc.)
    history/        # Run history tracking
    logging/        # Structured logging (slog)
    manifest/       # Dependency manifest
    pairwise/       # Tool-ablation pairwise expansion
    progress/       # Live progress display
    prompt/         # Prompt loading, parsing, filtering, validation
    report/         # Report generation (JSON, HTML, Markdown)
    rerender/       # Re-rendering past reports
    review/         # Multi-model review panel, rubric
    serve/          # Local web server for report browsing
    skills/         # Skill fetching (local + remote)
    tools/          # Tool-related utilities and registry
    trends/         # Cross-run trend analysis
    utils/          # Shared utilities
    validate/       # Prompt and config validation
.hyoka/             # Project directory (created by hyoka init)
  configs/          # Evaluation configs
  prompts/          # Prompt library
  criteria/         # Grader criteria files
  skills/           # Copilot skills
  reports/          # Evaluation output (git-ignored)
configs/            # Evaluation configurations
criteria/           # Attribute-matched criteria (per-language, per-service)
prompts/            # Prompt library
skills/             # Copilot skills (generator + reviewer)
```

## Adding a New Command

1. Create the command function in `main.go`:
   ```go
   func myCmd() *cobra.Command {
       return &cobra.Command{
           Use:   "my-command",
           Short: "Description",
           RunE: func(cmd *cobra.Command, args []string) error {
               // implementation
           },
       }
   }
   ```
2. Register it in `rootCmd()`: `root.AddCommand(myCmd())`
3. Add tests in `main_test.go`

## Adding a New Report Format

1. Add a generation function in `hyoka/internal/report/`
2. Call it from the engine's report-writing section in `engine.go`
3. Add format-specific tests

## Conventions

- Use Go standard library where possible (`log/slog`, `net/http`, `html/template`)
- Return errors up the call stack; don't log-and-return
- User-facing output → stdout/stderr directly
- Diagnostic logging → `log/slog`
- CLI framework: `github.com/spf13/cobra`
- Config format: YAML with `gopkg.in/yaml.v3`

## Git Workflow

- Branch naming: `{user}/issue-{N}-{description}` or `{user}/dev`
- Always include co-author trailer: `Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>`

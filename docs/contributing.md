# Contributing to Hyoka

How to build, test, and extend hyoka.

## Prerequisites

- **Go 1.24.5+** — `go version`
- **Git** — `git --version`
- **GitHub Copilot CLI** — for running real evaluations (optional for development)

## Building

The repo uses a Go workspace (`go.work`), so all commands run from the repo root:

```bash
# Build the CLI binary
go build ./hyoka

# Or run directly without building
go run ./hyoka version
```

## Testing

```bash
# Run all tests
go test ./hyoka/...

# Run tests with verbose output
go test -v ./hyoka/...

# Run tests for a specific package
go test ./hyoka/internal/config/...
go test ./hyoka/internal/prompt/...
go test ./hyoka/internal/review/...

# Run tests with race detection
go test -race ./hyoka/...

# Run tests with coverage
go test -cover ./hyoka/...
```

### Test Patterns

Tests use Go's standard `testing` package:

- **Unit tests** test individual functions and structs
- **Integration tests** (in `hyoka/eval.test/`, `hyoka/report.test/`, `hyoka/review.test/`) test full evaluation flows with stub data
- **Test data** lives in `hyoka/testdata/` for fixtures

## Project Structure

```
hyoka/
├── main.go                    # CLI entry point (Cobra commands)
├── go.mod / go.sum            # Go module dependencies
├── internal/                  # Internal packages (not importable)
│   ├── eval/                  # Evaluation engine
│   ├── config/                # Configuration loading and parsing
│   ├── prompt/                # Prompt loading, parsing, filtering
│   ├── review/                # Multi-model code review
│   ├── report/                # Report generation (JSON/HTML/MD)
│   ├── build/                 # Build verification
│   ├── trends/                # Trend analysis
│   ├── validate/              # Prompt validation
│   ├── checkenv/              # Environment checks
│   ├── logging/               # Structured logging
│   ├── progress/              # Progress display
│   ├── skills/                # Skill resolution
│   ├── rerender/              # Report re-rendering
│   ├── manifest/              # Prompt manifest
│   ├── history/               # Run history queries
│   └── utils/                 # Utilities
├── eval.test/                 # Integration test fixtures
├── report.test/               # Report test fixtures
├── review.test/               # Review test fixtures
└── testdata/                  # Shared test data
```

## Coding Conventions

### General

- **Go stdlib preferred** — minimize external dependencies
- **Cobra** for CLI commands
- **slog** for structured logging (Go 1.21+ stdlib)
- **YAML v3** (`gopkg.in/yaml.v3`) for config/frontmatter parsing
- **html/template** for report generation

### Style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Keep functions focused — one responsibility each
- Use meaningful variable names over comments
- Error wrapping with context: `fmt.Errorf("loading config: %w", err)`

### Package Design

- `internal/` packages are not importable by external code
- Each package has a clear responsibility (config, prompt, eval, review, report, etc.)
- Interfaces used for testability (e.g., `Evaluator`, `Reviewer`)
- `//go:embed` for static assets (rubric, templates)

## Adding a New CLI Command

1. Create a command function in `main.go`:

```go
func myNewCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "my-command",
        Short: "Brief description",
        RunE: func(cmd *cobra.Command, args []string) error {
            // Implementation
            return nil
        },
    }
    cmd.Flags().StringVar(&someVar, "flag-name", "default", "Description")
    return cmd
}
```

2. Register it in `rootCmd()`:

```go
root.AddCommand(myNewCmd())
```

3. Add tests and documentation.

## Adding a New Internal Package

1. Create the directory: `hyoka/internal/mypackage/`
2. Add Go files with `package mypackage`
3. Add tests: `mypackage_test.go`
4. Import from other internal packages as needed

## Adding a New Report Format

Reports are generated in `internal/report/`:

1. Create a new generator (e.g., `csv.go`)
2. Implement report writing using the `EvalReport` and `RunSummary` structs
3. Call it from the report orchestration in `generator.go`

## Working with Configs

When adding config fields:

1. Add the field to the appropriate struct in `internal/config/config.go` (`ToolConfig`, `GeneratorConfig`, or `ReviewerConfig`)
2. If legacy support is needed, add migration logic to `Normalize()`
3. Add accessor method if the field can come from legacy or new format
4. Add validation in `Parse()`
5. Update tests

## Working with Prompts

When extending prompt metadata:

1. Add the field to `Prompt` struct in `internal/prompt/types.go`
2. Update `ParsePromptFile()` in `internal/prompt/parser.go` if the field needs special parsing
3. Update validation rules in `internal/validate/validate.go`
4. Update the prompt template in `templates/prompt-template.prompt.md`

## Git Workflow

### Branch Naming

```
treebeard/issue-{N}-short-description
```

### Commit Messages

```
docs: add comprehensive documentation site

Closes #48

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>
```

### Git Identity

```bash
git config user.email "ronniegeraghty@users.noreply.github.com"
git config user.name "Ronnie Geraghty"
```

## Release Process

1. Update version in `main.go`
2. Run full test suite: `go test ./hyoka/...`
3. Build and verify: `go build ./hyoka && ./hyoka version`
4. Tag: `git tag v0.X.0`
5. Push: `git push origin main --tags`
6. Build binaries for distribution (if applicable)

## Running the Full Pipeline

For development, use stub mode to test without Copilot:

```bash
# Quick test of the entire pipeline
go run ./hyoka run --stub --prompt-id storage-dp-python-crud -y

# Validate everything
go run ./hyoka validate

# Check environment
go run ./hyoka check-env

# Generate trends from past runs
go run ./hyoka trends --no-analyze
```

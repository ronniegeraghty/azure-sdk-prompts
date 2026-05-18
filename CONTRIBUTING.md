# Contributing to hyoka

Thank you for your interest in contributing to hyoka — an evaluation tool for AI-generated Azure SDK code.

## Prerequisites

- **Go 1.26.1+** — check with `go version`
- **GitHub CLI** (`gh`) — [install](https://cli.github.com/)
- **Copilot SDK** — installed and authenticated
- **git** — configured with GitHub SSH access

## Clone and build

```bash
git clone https://github.com/ronniegeraghty/hyoka.git
cd hyoka

# Download dependencies (uses go.work for workspace)
go mod download

# Install npm dependencies and git hooks
npm install

# Build the CLI
go build ./hyoka/...

# Run the CLI
go run . <command>
```

### Git hooks

A **pre-commit hook** is installed via npm when you run `npm install`. If you modify `site/src/**`, the hook automatically rebuilds `site/dist/` and stages it for commit. This ensures the go:embed'd bundle never goes stale.

**If the site build fails:** Fix the TypeScript/CSS errors, then retry your commit. The hook will run again.

## Running tests

```bash
# All tests with race detection (required before committing)
go test -race ./hyoka/...

# Specific package
go test ./hyoka/internal/eval/ -v
```

## Development workflow

### 1. Create a branch from `main`

```bash
git checkout main
git pull origin main
git checkout -b ronniegeraghty/issue-{N}-{short-description}
```

Branch naming convention: `{username}/issue-{N}-{short-description}`

### 2. Configure git identity

```bash
git config user.name "ronniegeraghty"
git config user.email "ronniegeraghty@users.noreply.github.com"
```

### 3. Make changes

Edit code in `hyoka/cmd/` (CLI commands) or `hyoka/internal/` (packages). Run the formatter:

```bash
go fmt ./hyoka/...
```

### 4. Test manually with an eval run

For the fastest feedback loop, run a single prompt against a single config:

```bash
go run . run --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6 \
  --log-level debug --log-file hyoka-debug.log
```

Python prompts finish quickest (5–10 minutes). After each run, clean up orphaned sessions:

```bash
go run . clean
```

Browse results locally:

```bash
go run . serve
# Open http://localhost:8080
```

### 5. Commit with co-author trailer

Every commit **must** include the Copilot co-author trailer:

```bash
git commit -m "Fix eval engine timeout handling

- Add timeout context to generation phase
- Capture timeout error in action timeline

Closes #162

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### 6. Push and create a PR

```bash
gh auth switch --user ronniegeraghty
git push origin ronniegeraghty/issue-{N}-{short-description}
gh pr create --base main \
  --title "Fix: eval engine timeout" \
  --body "Closes #162"
```

## Project structure

```
hyoka/
  main.go           # Entry point (delegates to cmd.Execute())
  cmd/              # Cobra CLI commands
  internal/
    build/          # Language-specific build verification
    checkenv/       # Environment prerequisite checks
    config/         # YAML config loading and parsing
    criteria/       # Tiered evaluation criteria
    eval/           # Evaluation engine (generation + review)
    graders/        # Grader registry (6 types)
    progress/       # Live progress display
    prompt/         # Prompt loading, filtering, validation
    report/         # Report generation (JSON, HTML, Markdown)
    review/         # Multi-model review panel
    serve/          # Local web server for report browsing
    ...
configs/            # Evaluation config YAML files
criteria/           # Attribute-matched grader criteria
prompts/            # Prompt library (by language/service)
skills/             # Copilot skills (generator + reviewer)
```

## Adding a new CLI command

1. Create `hyoka/cmd/mycommand.go`:

   ```go
   var myCmd = &cobra.Command{
       Use:   "mycommand",
       Short: "One-line description",
       RunE: func(cmd *cobra.Command, args []string) error {
           // implementation
       },
   }

   func init() { rootCmd.AddCommand(myCmd) }
   ```

2. Add tests in `hyoka/cmd/mycommand_test.go`.
3. Document the command in its `Long` help text.

## Adding a new grader type

1. Create `hyoka/internal/graders/my_grader.go` implementing the grader interface.
2. Register it in `hyoka/internal/graders/registry.go`.
3. Write tests in `hyoka/internal/graders/my_grader_test.go`.
4. Update config examples in `configs/*.yaml` if applicable.

## Coding conventions

- **Go standard library preferred** — `log/slog` for logging, `net/http` for HTTP, `html/template` for templates.
- **CLI framework:** `github.com/spf13/cobra`
- **Config format:** YAML with `gopkg.in/yaml.v3`
- **Error handling:** Return errors up the call stack. Don't log-and-return.
- **User-facing output** → stdout/stderr directly (progress bars, results).
- **Diagnostic logging** → `log/slog`.

## Code review process

1. Submit your PR with `Closes #{issue}` in the description.
2. Address review comments in follow-up commits.
3. Ensure CI tests pass.
4. Squash if requested for history cleanliness.

## Anti-patterns to avoid

- Working directly on `main` — always use feature branches.
- Committing without running `go test -race`.
- Leaving orphaned Copilot sessions — always run `hyoka clean` after test runs.
- Hardcoding paths or config names in code.

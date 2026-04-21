# Agent Instructions for hyoka

## Overview

hyoka is a Go CLI tool that evaluates AI agents generating code. It uses GitHub Copilot sessions to generate code from prompts, then runs a multi-model review panel to score the output using extensible graders.

## Repository Structure

```
hyoka/              # Go source (module: github.com/ronniegeraghty/hyoka)
  main.go
  cmd/              # CLI commands
  internal/         # All packages (19 modules)
.hyoka/             # Project directory (created by hyoka init)
  configs/          # Evaluation configs
  prompts/          # Prompt library
  criteria/         # Grader criteria files
  skills/           # Copilot skills
  reports/          # Evaluation output (git-ignored)
configs/            # Evaluation config YAML files
prompts/            # Prompt library (organized by language/service)
skills/             # Copilot skills (generator/ and reviewer/)
criteria/           # Attribute-matched criteria (language/ and service/ subdirs)
reports/            # Generated evaluation output (gitignored)
docs/               # Design docs and getting started guide
```

To see a complete package inventory with descriptions, run:

```bash
# List all internal packages
go list ./hyoka/internal/...

# Or inspect the directory
ls -la hyoka/internal/
```

## Build & Test

```bash
# Build (from repo root)
go build ./hyoka/...

# Run tests
go test ./hyoka/...

# Run the CLI
go run . <command>

# Common commands
go run . list
go run . run --all-configs
go run . validate
go run . check-env
go run . clean
```

Go version: 1.26.1+ required. Module path: `github.com/ronniegeraghty/hyoka`.

## Running Evaluations

### Config Naming Convention

Config YAML files live in `configs/`. The `--config` flag takes the `name:` field from **inside** the YAML file, **NOT** the filename.

To discover available configs:

```bash
# List all available configs
go run . configs

# Or inspect the configs directory directly
ls configs/ | grep -E '\.ya?ml$'
for f in configs/*.yaml; do echo "File: $f"; grep '^name:' "$f"; done
```

Example: `configs/azure-mcp-opus.yaml` contains `name: azure-mcp/claude-opus-4.6` → use `--config azure-mcp/claude-opus-4.6`

### Prompt ID Patterns

- `--prompt-id` accepts a **single** prompt ID (not multiple, not comma-separated)
- Prompt IDs follow the pattern: `{service}-{plane-abbrev}-{language}-{short-name}`
  - e.g., `identity-dp-python-default-credential`, `key-vault-dp-python-crud-secrets`
  - `dp` = data-plane, `mp` = management-plane
- To run multiple prompts, use filter flags: `--service`, `--language`, `--plane`, `--category`

### Command Examples

```bash
# Single prompt, single config:
go run . run --prompt-id identity-dp-python-default-credential \
  --config baseline/claude-opus-4.6

# Single prompt, multiple configs (MUST quote comma-separated values):
go run . run --prompt-id identity-dp-python-default-credential \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6"

# Filter by service + language (runs ALL matching prompts):
go run . run --service key-vault --language python \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6"

# Full debug logging with log file:
go run ./hyoka run --service identity --language python \
  --config azure-mcp/claude-opus-4.6 \
  --log-level debug --log-file hyoka-debug.log

# Dry run (list matching prompts without executing):
go run ./hyoka run --service storage --language dotnet --dry-run

# All configs (requires explicit --all-configs flag):
go run ./hyoka run --prompt-id identity-dp-python-default-credential --all-configs

# With resource monitoring:
go run ./hyoka run --service key-vault --language python \
  --config azure-mcp/claude-opus-4.6 --monitor-resources
```

### Important Flag Rules

- `--config` values with commas **MUST** be quoted: `--config "config1,config2"`
- `--prompt-id` is singular — pass **ONE** ID only
- `--tags` is also comma-separated and must be quoted: `--tags "auth,crud"`
- Without `--config` or `--all-configs`, the run will fail
- `--log-level debug` enables verbose logging; pair with `--log-file` to capture to file
- `--max-session-actions` (default: 50) limits actions per Copilot session
- Prompt-level overrides: Prompts can override `--max-session-actions` and `--max-turns` via frontmatter (see [configuration.md](docs/configuration.md#prompt-level-limits) for examples). Resolution order: prompt frontmatter > config YAML > CLI flag > default.

### Available Filter Flags

```
--service        Azure service (e.g., identity, key-vault, storage, cosmos-db)
--language       Programming language (e.g., python, dotnet, java, js-ts, go, rust, cpp)
--plane          data-plane or management-plane
--category       Use-case category (e.g., auth, crud, pagination)
--tags           Comma-separated tags (must quote)
--prompt-id      Single prompt ID
```

## Testing Changes with Live Runs

When working on hyoka itself, test your changes by running real evaluations:

```bash
# Run 1 prompt on 1 config (fastest feedback loop — Python prompts finish quickest):
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config "baseline/claude-opus-4.6" \
  --log-level debug --log-file hyoka-debug.log

# After each run, clean up orphaned Copilot sessions:
go run ./hyoka clean

# Check the log file for role-prefixed output:
grep "role=" hyoka-debug.log | head -20

# Check the serve command to browse results:
go run ./hyoka serve
```

**Guidelines:**
- Run **1 prompt × 1 config** at a time when iterating — multi-eval runs can take 10+ minutes
- Always run `hyoka clean` after test runs to ensure no sessions were orphaned
- Python prompts tend to complete fastest; use them for quick iteration
- Use `--progress off` when you need clean stderr output (no live display interference)
- The `--log-file` flag writes to BOTH the file AND stderr, so you see debug output in the console too
- Compare console output and log file to verify logging changes work end-to-end

## Git Workflow

- **Branch naming**: `{username}/issue-{N}-{short-description}`
- **Commit trailers**: Always include `Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>`
- **Git identity**: Configure your GitHub account:
  ```
  git config user.name "{your-github-username}"
  git config user.email "{your-github-email}"
  ```
- **Push auth**: Use `gh auth switch` to select your account before pushing

## Coding Conventions

**Quick Reference (critical conventions):**
- **Go standard library preferred** — use `log/slog` for logging, `net/http` for HTTP
- **CLI framework**: `github.com/spf13/cobra`
- **Config format**: YAML with `gopkg.in/yaml.v3`
- **Error handling**: Return errors up the call stack; don't log-and-return

For detailed patterns and conventions, refer to:
- **Logging**: See `logging-conventions` skill
- **Error handling**: See `error-handling` skill
- **Testing patterns**: See `testing-patterns` skill
- **Go best practices**: See `golang-patterns` skill

## Key Architectural Patterns

For comprehensive architectural documentation, see [`docs/architecture.md`](docs/architecture.md).

Quick overview:
- **Multi-model review panel**: Multiple LLMs review generated code independently, then a consolidator merges scores
- **Config-driven evaluation**: Each YAML config defines a generator model, reviewer models, skills, and MCP servers
- **Prompt frontmatter**: Prompts have YAML frontmatter with `id`, `service`, `language`, `plane`, `category`, `difficulty`
- **Guardrails**: Turn limits (25), file limits (50), output size limits (1 MB)

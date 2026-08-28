# Filtering and CLI Flags Guide

## Overview

hyoka provides powerful filtering and control flags to run evaluations at any scale — from a single prompt to the entire library. This guide covers all the ways to select, filter, and control evaluation runs.

## Prompt Selection

### Single Prompt by ID

```bash
# Run exactly one prompt
go run ./hyoka run --prompt-id identity-dp-python-default-credential \
  --config baseline/claude-opus-4.6
```

**Important:** `--prompt-id` accepts a **single** prompt ID only (not comma-separated).

### Filter by Service

```bash
# Run all Key Vault prompts
go run ./hyoka run --service key-vault \
  --config azure-mcp/claude-opus-4.6

# Run all Identity prompts
go run ./hyoka run --service identity \
  --config baseline/claude-opus-4.6
```

Available services: `identity`, `key-vault`, `storage`, `cosmos-db`, `event-hubs`, `service-bus`, `app-config`, etc.

### Filter by Language

```bash
# Run all Python prompts
go run ./hyoka run --language python \
  --config azure-mcp/claude-opus-4.6

# Run all .NET prompts
go run ./hyoka run --language dotnet \
  --config baseline/claude-opus-4.6
```

Available languages: `python`, `dotnet`, `java`, `js-ts`, `go`, `rust`, `cpp`

### Filter by Plane

```bash
# Data-plane prompts only
go run ./hyoka run --plane data-plane \
  --config azure-mcp/claude-opus-4.6

# Management-plane prompts only
go run ./hyoka run --plane management-plane \
  --config baseline/claude-opus-4.6
```

### Filter by Category

```bash
# Authentication prompts
go run ./hyoka run --category auth \
  --config azure-mcp/claude-opus-4.6

# CRUD operations
go run ./hyoka run --category crud \
  --config baseline/claude-opus-4.6

# Pagination prompts
go run ./hyoka run --category pagination \
  --config azure-mcp/claude-opus-4.6
```

### Filter by Tags

```bash
# Prompts tagged with specific keywords (comma-separated, must quote)
go run ./hyoka run --tags "getting-started,crud" \
  --config azure-mcp/claude-opus-4.6
```

### Filter by Difficulty

```bash
# Basic difficulty only
go run ./hyoka run --difficulty basic \
  --config baseline/claude-opus-4.6

# Advanced prompts
go run ./hyoka run --difficulty advanced \
  --config azure-mcp/claude-opus-4.6
```

Difficulty levels: `basic`, `intermediate`, `advanced`

### Combine Multiple Filters

```bash
# Python + Key Vault + data-plane + auth
go run ./hyoka run --service key-vault \
  --language python \
  --plane data-plane \
  --category auth \
  --config azure-mcp/claude-opus-4.6
```

Filters are applied with AND logic — all must match.

## Config Selection

### Single Config

```bash
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6
```

### Multiple Configs (Comma-Separated)

**Important:** Comma-separated config names **MUST** be quoted.

```bash
# Compare baseline vs azure-mcp configs
go run ./hyoka run --prompt-id identity-dp-python-default-credential \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6"

# Run across three configs
go run ./hyoka run --service key-vault --language python \
  --config "baseline/claude-opus-4.6,baseline/gpt-5.3-codex,azure-mcp/claude-opus-4.6"
```

### All Configs

**Important:** Requires explicit `--all-configs` flag.

```bash
# Run one prompt on every config in configs/
go run ./hyoka run --prompt-id identity-dp-python-default-credential \
  --all-configs
```

### Config Name Format

Config names come from the `name:` field **inside** the YAML file, not the filename.

Example: `configs/azure-mcp-opus.yaml` contains:
```yaml
configs:
  - name: azure-mcp/claude-opus-4.6
```

So use `--config azure-mcp/claude-opus-4.6`, not `--config azure-mcp-opus`.

## Run Control Flags

### Dry Run

Preview what prompts would run without executing:

```bash
go run ./hyoka run --service storage --language dotnet \
  --config azure-mcp/claude-opus-4.6 --dry-run
```

Outputs matching prompts with their IDs, descriptions, and filters.

### Progress Display Modes

Control how evaluation progress is displayed:

```bash
# Live progress bar (default)
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6 --progress live

# Log-only output (no progress bar, just timestamped events)
go run ./hyoka run --service identity --language python \
  --config azure-mcp/claude-opus-4.6 --progress log

# No progress output (silent except errors)
go run ./hyoka run --service key-vault --language python \
  --config baseline/claude-opus-4.6 --progress off
```

**Tip:** Use `--progress off` when you need clean stderr output, or `--progress log` for CI environments where live updates don't render well.

### Resource Monitoring

Track CPU, memory, and disk usage during evaluations:

```bash
go run ./hyoka run --service key-vault --language python \
  --config azure-mcp/claude-opus-4.6 --monitor-resources
```

Resource metrics appear in the JSON report and HTML dashboard.

### Session Action Limits

Control how many actions a Copilot session can take:

```bash
# Increase action limit for complex prompts
go run ./hyoka run --prompt-id cosmos-db-dp-python-pagination \
  --config azure-mcp/claude-opus-4.6 --max-session-actions 75
```

Default: 50 actions per session.

## Logging Flags

### Log Level

Control verbosity of diagnostic logs:

```bash
# Debug logging for troubleshooting
go run ./hyoka run --service identity --language python \
  --config azure-mcp/claude-opus-4.6 --log-level debug

# Info logging (default)
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6 --log-level info

# Warn or error only
go run ./hyoka run --service storage --language python \
  --config azure-mcp/claude-opus-4.6 --log-level warn
```

Levels: `debug`, `info`, `warn`, `error`

### Log to File

Write logs to a file (in addition to stderr):

```bash
go run ./hyoka run --service key-vault --language python \
  --config azure-mcp/claude-opus-4.6 \
  --log-level debug --log-file hyoka-debug.log
```

**Note:** Logs are written to **both** the file and stderr, so you see output in real-time.

## Tool Ablation

### Pairwise Testing

Generate tool-ablation variants (baseline + N "without-tool" configs):

```bash
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config azure-mcp/claude-opus-4.6 --pairwise
```

See `examples/pairwise-testing.md` for full details.

## Example Workflows

### Quick Single-Prompt Test

```bash
go run ./hyoka run --prompt-id identity-dp-python-default-credential \
  --config baseline/claude-opus-4.6
```

### Compare Two Configs on Key Vault + Python

```bash
go run ./hyoka run --service key-vault --language python \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6"
```

### Full Language Suite (All Python Prompts, All Configs)

```bash
go run ./hyoka run --language python --all-configs
```

### Debug a Failing Prompt

```bash
go run ./hyoka run --prompt-id cosmos-db-dp-python-pagination \
  --config azure-mcp/claude-opus-4.6 \
  --log-level debug --log-file debug.log --progress log
```

### Pairwise Tool Impact Analysis with Resource Monitoring

```bash
go run ./hyoka run --service identity --language python \
  --config azure-mcp/claude-opus-4.6 \
  --pairwise --monitor-resources
```

## Tips

- **Always quote comma-separated values:** `--config "cfg1,cfg2"`, `--tags "tag1,tag2"`
- **Use dry-run first:** Preview matches before running expensive evaluations
- **Filter early:** Run on a subset (one service, one language) before scaling to all prompts
- **Python prompts are fast:** Use Python prompts for quick iteration when testing hyoka changes
- **Monitor first runs:** Use `--monitor-resources` and `--log-level debug` on unfamiliar prompts to catch issues early

## See Also

- `examples/configs/` — Example config files demonstrating all features
- `examples/prompts/` — Example prompts with various frontmatter fields
- `examples/pairwise-testing.md` — Deep dive on tool ablation testing
- `docs/configuration.md` — Full configuration reference
- `docs/prompt-authoring.md` — Prompt frontmatter and structure guide

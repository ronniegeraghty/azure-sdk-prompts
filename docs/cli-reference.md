# CLI Reference

Complete command and flag reference for the hyoka evaluation tool.

## Global Flags

These flags apply to all commands:

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--log-level` | string | `warn` | Log level: `debug`, `info`, `warn`, `error` |
| `--log-file` | string | _(empty)_ | Redirect log output to a file (stderr stays clean) |

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| [`run`](#run) | | Run evaluations against prompts |
| [`list`](#list) | `ls` | List prompts matching filters |
| [`configs`](#configs) | | Show available tool configurations |
| [`validate`](#validate) | | Validate prompt frontmatter and config files |
| [`check-env`](#check-env) | `env` | Check for required language toolchains and tools |
| [`trends`](#trends) | | Generate historical trend reports with AI analysis |
| [`report`](#report) | | Re-render HTML/MD reports from existing JSON data |
| [`new-prompt`](#new-prompt) | | Scaffold a new prompt file interactively |
| [`version`](#version) | | Print version |

---

## `run`

Run evaluations with optional filters against the prompt library.

```bash
go run ./hyoka run [flags]
```

### Filter Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--prompts` | string | `./prompts` | Path to prompt library directory |
| `--service` | string | | Filter by Azure service (e.g., `storage`, `key-vault`) |
| `--language` | string | | Filter by programming language (e.g., `dotnet`, `python`) |
| `--plane` | string | | Filter by plane (`data-plane`, `management-plane`) |
| `--category` | string | | Filter by category (e.g., `crud`, `authentication`) |
| `--tags` | string | | Filter by tags (comma-separated) |
| `--prompt-id` | string | | Run a single prompt by ID |

### Config Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--config` | string | | Config name(s) to use (comma-separated) |
| `--config-file` | string | | Path to a specific config YAML file |
| `--config-dir` | string | `./configs` | Directory containing config YAML files |
| `--model` | string | | Override model for all configs |
| `--all-configs` | bool | `false` | Required when running all configs without a `--config` filter |

### Execution Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--workers` | int | CPU cores (max 8) | Parallel evaluation workers |
| `--max-sessions` | int | workers × 3 | Maximum concurrent Copilot sessions |
| `--timeout` | int | `600` | Per-prompt generation timeout in seconds _(deprecated: use `--generate-timeout`)_ |
| `--generate-timeout` | int | _(defaults to `--timeout`)_ | Generation phase timeout in seconds |
| `--build-timeout` | int | `300` | Build verification timeout in seconds |
| `--review-timeout` | int | `300` | Review phase timeout in seconds |
| `--output` | string | `./reports` | Report output directory |

### Feature Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--skip-tests` | bool | `false` | Skip test generation |
| `--skip-review` | bool | `false` | Skip code review |
| `--verify-build` | bool | `false` | Run build verification on generated code |
| `--skip-trends` | bool | `false` | Skip automatic trend analysis after run |
| `--dry-run` | bool | `false` | List matching prompts without running |
| `--stub` | bool | `false` | Use stub evaluator (no Copilot SDK) |
| `-y`, `--yes` | bool | `false` | Skip confirmation prompt for large runs (>10 evaluations) |

### Guardrail Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--max-turns` | int | `25` | Maximum conversation turns per generation |
| `--max-files` | int | `50` | Maximum generated files per evaluation |
| `--max-output-size` | string | `1MB` | Maximum total output size (supports `KB`, `MB` suffixes) |
| `--allow-cloud` | bool | `false` | Allow generated code to provision real Azure resources |

### Display Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--progress` | string | `auto` | Progress display mode: `auto`, `live`, `log`, `off` |

### Examples

```bash
# Run a single prompt with a specific config
go run ./hyoka run --prompt-id storage-dp-dotnet-auth --config baseline

# Run all storage prompts
go run ./hyoka run --service storage

# Run all prompts × all configs (requires --all-configs)
go run ./hyoka run --all-configs -y

# Dry run — see what would execute
go run ./hyoka run --service storage --dry-run

# Stub mode — test pipeline without Copilot
go run ./hyoka run --stub

# Tighten guardrails for faster iteration
go run ./hyoka run --max-turns 10 --max-files 20

# Limit concurrency on shared machines
go run ./hyoka run --max-sessions 4 --workers 2
```

---

## `list`

List prompts matching the given filters. Alias: `ls`.

```bash
go run ./hyoka list [flags]
```

### Flags

Inherits all [filter flags](#filter-flags) from `run`, plus:

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--json` | bool | `false` | Output prompts as JSON array |

### Examples

```bash
# List all prompts
go run ./hyoka list

# Filter by service
go run ./hyoka list --service storage

# JSON output for scripting
go run ./hyoka list --json

# Combine filters
go run ./hyoka list --service storage --language python --plane data-plane
```

---

## `configs`

List available evaluation configurations.

```bash
go run ./hyoka configs [flags]
```

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--config-file` | string | | Path to a specific configuration YAML file |
| `--config-dir` | string | `./configs` | Directory containing configuration YAML files |

### Examples

```bash
# List all configs
go run ./hyoka configs

# List configs from a specific file
go run ./hyoka configs --config-file configs/baseline-sonnet.yaml
```

---

## `validate`

Validate all prompt files against schema rules, naming conventions, and config file syntax.

```bash
go run ./hyoka validate [flags]
```

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--prompts` | string | `./prompts` | Path to prompt library directory |

### What It Checks

- **Required frontmatter fields** — `id`, `service`, `plane`, `language`, `category`, `difficulty`, `description`, `created`, `author`
- **Enum validation** — field values against allowed values
- **ID naming convention** — must match `{service}-{dp|mp}-{language}-{slug}`
- **Prompt content** — must have a non-empty `## Prompt` section
- **Config YAML syntax** — valid structure and field types

### Examples

```bash
# Validate all prompts
go run ./hyoka validate

# Validate prompts in a custom directory
go run ./hyoka validate --prompts ~/my-prompts
```

---

## `check-env`

Check for required language toolchains and tools. Alias: `env`.

```bash
go run ./hyoka check-env
```

### What It Checks

- **Language toolchains:** `dotnet`, `python`, `go`, `node`, `java`, `rust`, `cargo`, `cmake`
- **Copilot CLI** availability and authentication
- **MCP prerequisites** (Node.js for Azure MCP server)

No flags — runs a fixed set of checks and prints results.

---

## `trends`

Generate historical trend reports with time-series performance data and AI-powered analysis.

```bash
go run ./hyoka trends [flags]
```

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--prompt-id` | string | | Filter trends by prompt ID |
| `--service` | string | | Filter trends by Azure service |
| `--language` | string | | Filter trends by programming language |
| `--reports-dir` | string | `./reports` | Directory containing past evaluation reports |
| `--output` | string | `./reports/trends` | Output directory for trend reports |
| `--analyze` | bool | `true` | Run AI-powered analysis of trends |
| `--no-analyze` | bool | `false` | Skip AI-powered trend analysis |
| `--open` | bool | `false` | Auto-open the HTML trend report in browser |

### Examples

```bash
# Generate trends for all past runs
go run ./hyoka trends

# Filter by service
go run ./hyoka trends --service storage

# Skip AI analysis for faster output
go run ./hyoka trends --no-analyze

# Auto-open the HTML report
go run ./hyoka trends --open
```

---

## `report`

Re-render HTML and Markdown reports from existing `report.json` files without re-running evaluations. Useful after template improvements.

```bash
go run ./hyoka report [run-id] [flags]
```

### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `run-id` | No | Specific run ID to re-render. Omit if using `--all`. |

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--reports-dir` | string | `./reports` | Directory containing evaluation reports |
| `--all` | bool | `false` | Re-render all runs |

### Examples

```bash
# Re-render a specific run
go run ./hyoka report 20250728-143022

# Re-render all runs
go run ./hyoka report --all
```

---

## `new-prompt`

Scaffold a new prompt file interactively. Asks for metadata and generates a file with populated frontmatter at the correct directory path.

```bash
go run ./hyoka new-prompt
```

No flags — fully interactive. Prompts for:

1. **Service** — which Azure service
2. **Plane** — data-plane or management-plane
3. **Language** — target programming language
4. **Category** — use-case category
5. **Difficulty** — basic, intermediate, or advanced
6. **Description** — what the prompt tests
7. **Slug** — short filename identifier

Creates a file at `prompts/{service}/{plane}/{language}/{slug}.prompt.md` with populated frontmatter ready for you to add the prompt text and evaluation criteria.

---

## `version`

Print the hyoka version.

```bash
go run ./hyoka version
```

Output:
```
hyoka version 0.2.0
```

---

## Smart Path Detection

hyoka automatically resolves common paths:

- **Prompts:** Checks `./prompts` then `../prompts`
- **Configs:** Checks `./configs` then `../configs`
- **Reports:** Defaults to `./reports`

This means running from either the repo root or the `hyoka/` directory works without extra flags.

## Filtering

All filter flags use AND logic — only prompts matching **all** specified criteria are selected:

```bash
# Matches: service=storage AND language=dotnet AND plane=data-plane
go run ./hyoka run --service storage --language dotnet --plane data-plane
```

Filters work with `run`, `list`, and other prompt-aware commands.

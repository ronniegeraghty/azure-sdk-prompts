# sdk-eval — CLI Reference

The `sdk-eval` tool evaluates AI agent code generation quality by running prompts from the `azure-sdk-prompts` library through configurable Copilot sessions, verifying builds, scoring code via LLM-as-judge review, and generating JSON + HTML reports.

## Installation

### Run from source (recommended for development)

```bash
cd azure-sdk-prompts
go run ./tool/cmd/sdk-eval <command> [flags]
```

### Install globally

```bash
go install github.com/ronniegeraghty/azure-sdk-prompts/tool/cmd/sdk-eval@latest
sdk-eval <command> [flags]
```

> **Pinned version:** `go install github.com/ronniegeraghty/azure-sdk-prompts/tool/cmd/sdk-eval@tool/v0.2.0`

## Features

### Phase 1 (v0.1.0) ✅
- Prompt library loading, filtering, and validation
- Build verification for 9 languages (dotnet, Python, Go, Java, JS, TS, Rust, C++)
- JSON report generation with directory hierarchy
- Manifest generation and prompt validation

### Phase 2 (v0.2.0) ✅
- **Copilot SDK integration** — Real code generation via `github.com/github/copilot-sdk/go`
- **LLM-as-judge code review** — Separate Copilot session scores generated code on 7 dimensions
- **Reference answer comparison** — Optional reference code included in review prompt
- **HTML reports** — Per-evaluation reports with score visualization and collapsible build output
- **Summary dashboard** — Cross-config comparison matrix (prompt × config) with scores and build status
- **Graceful fallback** — Falls back to stub evaluator if Copilot CLI is unavailable

## Authentication

The Copilot SDK evaluator requires a running Copilot CLI with valid authentication. The SDK will:
1. Try `GITHUB_TOKEN` environment variable
2. Try the logged-in user's GitHub CLI (`gh`) auth token
3. If neither is available, fall back to the stub evaluator with a warning

Use `--stub` to explicitly skip SDK initialization and use the stub evaluator.

## Commands

### `sdk-eval run`

Run evaluations against the prompt library.

```bash
sdk-eval run [flags]
```

| Flag | Default | Description |
|------|---------|-------------|
| `--prompts` | `./prompts` (auto-detected) | Path to prompt library directory |
| `--service` | | Filter by Azure service |
| `--language` | | Filter by programming language |
| `--plane` | | Filter by data-plane / management-plane |
| `--category` | | Filter by use-case category |
| `--tags` | | Filter by tags (comma-separated) |
| `--prompt-id` | | Run a single prompt by ID |
| `--config` | all | Config name(s) (comma-separated) |
| `--config-file` | `./configs/all.yaml` (auto-detected) | Path to configuration YAML |
| `--workers` | `4` | Parallel evaluation workers |
| `--timeout` | `300` | Per-prompt timeout in seconds |
| `--model` | | Override model for all configs |
| `--output` | `./reports` | Report output directory |
| `--skip-tests` | `false` | Skip test generation |
| `--skip-review` | `false` | Skip LLM-as-judge code review |
| `--stub` | `false` | Force stub evaluator (no Copilot SDK) |
| `--debug` | `false` | Verbose output |
| `--dry-run` | `false` | List matches without executing |

**Examples:**

```bash
# Run all prompts with all configs (real Copilot SDK)
sdk-eval run

# Run with stub evaluator (no SDK needed)
sdk-eval run --stub

# Run storage prompts with the baseline config, skip review
sdk-eval run --service storage --config baseline --skip-review

# Run a single prompt
sdk-eval run --prompt-id storage-dp-dotnet-auth

# Compare configs
sdk-eval run --service storage --config baseline,azure-mcp
```

### `sdk-eval list`

List prompts matching the given filters (no evaluation).

```bash
sdk-eval list [flags]
```

Takes the same filter flags as `run`. Output shows prompt ID, service/plane/language, category, and description.

### `sdk-eval manifest`

Regenerate `manifest.yaml` from prompt files.

```bash
sdk-eval manifest [flags]
```

| Flag | Default | Description |
|------|---------|-------------|
| `--prompts` | `./prompts` (auto-detected) | Path to prompt library directory |
| `--output` | `./manifest.yaml` (auto-detected) | Output path for manifest |

### `sdk-eval validate`

Validate prompt frontmatter against the schema.

```bash
sdk-eval validate [flags]
```

Checks required fields, enum values, ID naming conventions, and `## Prompt` section presence. Exits with code 1 on validation failure.

### `sdk-eval configs`

List available tool configurations.

```bash
sdk-eval configs [--config-file PATH]
```

### `sdk-eval version`

Print the tool version.

## Code Review (LLM-as-Judge)

After code generation, `sdk-eval` creates a **separate** Copilot session to review the generated code. This avoids self-bias — the reviewer didn't generate the code.

### Scoring Dimensions (1-10)

| Dimension | What it measures |
|-----------|-----------------|
| Correctness | Does the code correctly implement the prompt? |
| Completeness | Are all requirements addressed? |
| Best Practices | Azure SDK patterns (DefaultAzureCredential, disposal, async) |
| Error Handling | Proper error handling, retries, timeouts |
| Package Usage | Correct and up-to-date SDK packages |
| Code Quality | Clean, readable, well-structured code |
| Reference Similarity | Match to reference answer (if provided) |

### Reference Answers

If a prompt has a `reference_answer` field pointing to a directory of reference code, that code is included in the review prompt for comparison.

## Report Formats

### JSON (machine-readable)

```
reports/runs/<timestamp>/
├── summary.json          # Aggregate run statistics
└── results/
    └── <service>/<plane>/<language>/<category>/<config>/
        └── report.json   # Individual evaluation result (with review scores)
```

### HTML (human-readable)

```
reports/runs/<timestamp>/
├── summary.html          # Cross-config comparison matrix dashboard
└── results/
    └── <service>/<plane>/<language>/<category>/<config>/
        └── report.html   # Individual report with score visualization
```

The **summary.html** shows a matrix of prompt × config with overall scores and build pass/fail indicators:

| Prompt | baseline | azure-mcp | azure-mcp-plus-skills |
|---|---|---|---|
| storage-dp-dotnet-auth | 6/10 ✅ | 8/10 ✅ | 9/10 ✅ |
| storage-dp-python-crud | 5/10 ❌ | 7/10 ✅ | 8/10 ✅ |

## Configuration Matrix

Configurations live in the `configs/` directory at the repo root:

| File | Description |
|------|-------------|
| `configs/all.yaml` | Both configs — used for matrix runs (default) |
| `configs/baseline.yaml` | No MCP servers, no skills — raw Copilot |
| `configs/azure-mcp.yaml` | Azure MCP server attached |

**Sample config file:**

```yaml
configs:
  - name: baseline
    description: "No MCP servers, no skills — just base Copilot"
    model: "claude-sonnet-4.5"
    mcp_servers: {}
    skill_directories: []
    available_tools: []
    excluded_tools: []
```

### Config Fields

| Field | Type | SDK Mapping | Description |
|-------|------|-------------|-------------|
| `name` | string | — | Unique config identifier |
| `description` | string | — | Human-readable description |
| `model` | string | `SessionConfig.Model` | AI model to use |
| `mcp_servers` | map | `SessionConfig.MCPServers` | MCP server definitions |
| `skill_directories` | list | `SessionConfig.SkillDirectories` | Paths to skill directories |
| `available_tools` | list | `SessionConfig.AvailableTools` | Allowed tool names |
| `excluded_tools` | list | `SessionConfig.ExcludedTools` | Blocked tool names |

## Smart Path Detection

`sdk-eval` automatically resolves paths when flags aren't explicitly set:

| Flag | Candidates checked |
|------|--------------------|
| `--prompts` | `./prompts` → `../prompts` |
| `--config-file` | `./configs/all.yaml` → `../configs/all.yaml` → `./configs.yaml` → `../configs.yaml` |
| `--output` (manifest) | `./manifest.yaml` → `../manifest.yaml` |

## Project Structure

```
tool/
├── cmd/sdk-eval/main.go        # CLI entry point (cobra)
├── go.mod / go.sum
├── internal/
│   ├── config/                  # Config file parsing
│   ├── prompt/                  # Prompt loading, parsing, filtering
│   ├── eval/                    # Engine, workspace, CopilotSDKEvaluator
│   ├── build/                   # Build verification per language
│   ├── report/                  # JSON + HTML report generation
│   ├── review/                  # LLM-as-judge code review
│   ├── manifest/                # Manifest generation from prompts
│   └── validate/                # Prompt frontmatter validation
└── testdata/                    # Test fixtures
```

## Roadmap

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1 | ✅ Done | Prompt library, build verification, JSON reports (stub evaluator) |
| Phase 2 | ✅ Done | Copilot SDK integration, LLM-as-judge review, HTML reports |
| Phase 3 | Planned | Auto-generated tests, historical trend tracking, event trace reports |

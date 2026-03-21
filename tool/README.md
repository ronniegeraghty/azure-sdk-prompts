# azsdk-prompt-eval — CLI Reference

The `azsdk-prompt-eval` tool evaluates AI agent code generation quality by running prompts from the `azure-sdk-prompts` library through configurable Copilot sessions, verifying code with Copilot-based verification, scoring code via LLM-as-judge review, and generating JSON, HTML, and Markdown reports.

## Prerequisites

- **Go 1.24.5+** — to build and run the tool
- **GitHub Copilot CLI** — the SDK communicates with Copilot via the CLI in server mode. Must be installed and authenticated:
  - Install: follow [GitHub Copilot CLI setup](https://docs.github.com/en/copilot/how-tos/set-up/install-copilot-cli)
  - Authenticate: run `copilot` once to complete OAuth device flow, or set `COPILOT_GITHUB_TOKEN` / `GH_TOKEN` env var
  - Without this, the tool falls back to stub mode (no real evaluations)
- **GitHub CLI (`gh`)** — optional but recommended for auth token management
- **For `azure-mcp` config:** `npx` (Node.js) must be available since the Azure MCP server is launched via `npx -y @azure/mcp@latest`

## Installation

### Run from source (recommended for development)

```bash
cd azure-sdk-prompts
go run ./tool/cmd/azsdk-prompt-eval <command> [flags]
```

### Install globally

```bash
go install github.com/ronniegeraghty/azure-sdk-prompts/tool/cmd/azsdk-prompt-eval@latest
azsdk-prompt-eval <command> [flags]
```

> **Pinned version:** `go install github.com/ronniegeraghty/azure-sdk-prompts/tool/cmd/azsdk-prompt-eval@tool/v0.3.0`

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

### Phase 2.1 (v0.3.0) ✅
- **Copilot-based verification** — Separate Copilot session verifies code meets requirements (replaces build-only verification as default)
- **Build verification optional** — Use `--verify-build` to also run language-specific build checks
- **Session transcripts** — Full event capture (tool calls, assistant messages, errors) in JSON + HTML reports
- **Failure diagnostics** — Failed evals show detailed error info, session events, and stub mode indicator
- **Debug mode** — `--debug` streams real-time session events to stderr (tool calls, messages, verification/review status)
- **Flat report structure** — Reports write to `reports/{timestamp}/` instead of `reports/runs/{timestamp}/`
- **Expected Coverage** — Parser extracts `## Expected Coverage` sections from prompt files for verification

## Authentication

The Copilot SDK evaluator requires a running Copilot CLI with valid authentication. The SDK will:
1. Try `GITHUB_TOKEN` environment variable
2. Try the logged-in user's GitHub CLI (`gh`) auth token
3. If neither is available, fall back to the stub evaluator with a warning

Use `--stub` to explicitly skip SDK initialization and use the stub evaluator.

## Commands

### `azsdk-prompt-eval run`

Run evaluations against the prompt library.

```bash
azsdk-prompt-eval run [flags]
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
| `--verify-build` | `false` | Also run build verification (in addition to Copilot verification) |
| `--stub` | `false` | Force stub evaluator (no Copilot SDK) |
| `--debug` | `false` | Verbose output |
| `--dry-run` | `false` | List matches without executing |

**Examples:**

```bash
# Run all prompts with all configs (real Copilot SDK)
azsdk-prompt-eval run

# Run with stub evaluator (no SDK needed)
azsdk-prompt-eval run --stub

# Run storage prompts with the baseline config, skip review
azsdk-prompt-eval run --service storage --config baseline --skip-review

# Run a single prompt
azsdk-prompt-eval run --prompt-id storage-dp-dotnet-auth

# Compare configs
azsdk-prompt-eval run --service storage --config baseline,azure-mcp
```

### `azsdk-prompt-eval list`

List prompts matching the given filters (no evaluation).

```bash
azsdk-prompt-eval list [flags]
```

Takes the same filter flags as `run`. Output shows prompt ID, service/plane/language, category, and description.

### `azsdk-prompt-eval manifest`

Regenerate `manifest.yaml` from prompt files.

```bash
azsdk-prompt-eval manifest [flags]
```

| Flag | Default | Description |
|------|---------|-------------|
| `--prompts` | `./prompts` (auto-detected) | Path to prompt library directory |
| `--output` | `./manifest.yaml` (auto-detected) | Output path for manifest |

### `azsdk-prompt-eval validate`

Validate prompt frontmatter against the schema.

```bash
azsdk-prompt-eval validate [flags]
```

Checks required fields, enum values, ID naming conventions, and `## Prompt` section presence. Exits with code 1 on validation failure.

### `azsdk-prompt-eval configs`

List available tool configurations.

```bash
azsdk-prompt-eval configs [--config-file PATH]
```

### `azsdk-prompt-eval version`

Print the tool version.

### `azsdk-prompt-eval check-env`

Check for required language toolchains and tools.

```bash
azsdk-prompt-eval check-env
```

Reports availability of Python, .NET, Go, Node.js, Java, Rust, C/C++, Copilot CLI, gh authentication, and npx (for Azure MCP server). Uses ✅/❌ indicators with version strings.

## Code Review (LLM-as-Judge)

After code generation, `azsdk-prompt-eval` creates a **separate** Copilot session to review the generated code. This avoids self-bias — the reviewer didn't generate the code.

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
reports/<timestamp>/
├── summary.json          # Aggregate run statistics
└── results/
    └── <service>/<plane>/<language>/<category>/<config>/
        └── report.json   # Individual evaluation result (with review scores)
```

### HTML (human-readable)

```
reports/<timestamp>/
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

### Markdown (portable, git-friendly)

```
reports/<timestamp>/
├── summary.md            # Cross-config comparison matrix (Markdown)
└── results/
    └── <service>/<plane>/<language>/<category>/<config>/
        └── report.md     # Individual evaluation report (Markdown)
```

Markdown reports contain the same information as HTML reports (scores, tool calls, verification, review) in a clean, readable format suitable for viewing in GitHub, VS Code, or any Markdown renderer.

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

`azsdk-prompt-eval` automatically resolves paths when flags aren't explicitly set:

| Flag | Candidates checked |
|------|--------------------|
| `--prompts` | `./prompts` → `../prompts` |
| `--config-file` | `./configs/all.yaml` → `../configs/all.yaml` → `./configs.yaml` → `../configs.yaml` |
| `--output` (manifest) | `./manifest.yaml` → `../manifest.yaml` |

## Project Structure

```
tool/
├── cmd/azsdk-prompt-eval/main.go        # CLI entry point (cobra)
├── go.mod / go.sum
├── internal/
│   ├── checkenv/                # Environment check (check-env command)
│   ├── config/                  # Config file parsing
│   ├── prompt/                  # Prompt loading, parsing, filtering
│   ├── eval/                    # Engine, workspace, CopilotSDKEvaluator
│   ├── build/                   # Build verification per language
│   ├── report/                  # JSON + HTML report generation
│   ├── review/                  # LLM-as-judge code review
│   ├── verify/                  # Copilot-based code verification
│   ├── manifest/                # Manifest generation from prompts
│   └── validate/                # Prompt frontmatter validation
└── testdata/                    # Test fixtures
```

## Roadmap

| Phase | Status | Description |
|-------|--------|-------------|
| Phase 1 | ✅ Done | Prompt library, build verification, JSON reports (stub evaluator) |
| Phase 2 | ✅ Done | Copilot SDK integration, LLM-as-judge review, HTML reports |
| Phase 2.1 | ✅ Done | Copilot verification, session transcripts, debug mode, failure diagnostics |
| Phase 3 | ✅ Done | Tool matrix, MCP server attachment, skill loading, cross-config comparison |
| Phase 4 | In Progress | Evaluation quality — check-env, expected_tools, reviewer skills |
| Phase 5 | Planned | Polish — report re-rendering, embedded CLI, progress bars |

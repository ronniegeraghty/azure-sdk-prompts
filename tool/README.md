# SDK Eval Tool

A Go-based evaluation tool for testing AI agent code generation quality. It reads prompts from a YAML-frontmatter prompt library, runs them through the Copilot SDK, verifies builds, and generates JSON reports.

## Quick Start

```bash
# Build
go build ./cmd/sdk-eval

# List available prompts
./sdk-eval list --prompts ./path/to/prompts

# Run evaluations
./sdk-eval run --prompts ./path/to/prompts --language dotnet --config baseline

# List configurations
./sdk-eval configs
```

## CLI Commands

| Command | Description |
|---------|-------------|
| `sdk-eval run` | Run evaluations with filter flags |
| `sdk-eval list` | List matching prompts (dry-run) |
| `sdk-eval configs` | List available configurations |
| `sdk-eval version` | Print version |

## Filter Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--prompts` | Path to prompt library directory | `./prompts` |
| `--service` | Filter by Azure service | |
| `--language` | Filter by programming language | |
| `--plane` | Filter by data-plane/management-plane | |
| `--category` | Filter by use-case category | |
| `--tags` | Filter by tags (comma-separated) | |
| `--prompt-id` | Run a single prompt by ID | |
| `--config` | Config name(s) (comma-separated) | all |
| `--config-file` | Path to configuration YAML | `./configs.yaml` |
| `--workers` | Parallel workers | `4` |
| `--timeout` | Per-prompt timeout (seconds) | `300` |
| `--model` | Override model for all configs | |
| `--output` | Report output directory | `./reports` |
| `--skip-tests` | Skip test generation | `false` |
| `--skip-review` | Skip code review | `false` |
| `--debug` | Verbose output | `false` |
| `--dry-run` | List matching prompts without running | `false` |

## Configuration

Configurations are defined in `configs.yaml`:

```yaml
configs:
  - name: baseline
    description: "No MCP servers, no skills"
    model: "claude-sonnet-4.5"
    mcp_servers: {}
    skill_directories: []
    available_tools: []
    excluded_tools: []
```

## Project Structure

```
azure-sdk-prompts/tool/
├── cmd/sdk-eval/main.go          # CLI entry point
├── configs.yaml                   # Default tool configs
├── internal/
│   ├── config/                    # Config file parsing
│   ├── prompt/                    # Prompt loading and filtering
│   ├── eval/                      # Evaluation engine
│   ├── build/                     # Build verification
│   └── report/                    # Report generation
└── testdata/                      # Test fixtures
```

## Reports

Reports are written to timestamped directories:

```
reports/runs/<timestamp>/
├── summary.json
└── results/<service>/<plane>/<language>/<category>/<config>/
    └── report.json
```

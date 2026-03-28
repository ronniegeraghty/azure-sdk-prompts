# hyoka

A curated library of prompts for evaluating how well AI agents generate Azure SDK code, paired with a Go evaluation tool (`hyoka`) that runs prompts through the Copilot SDK, reviews code via a multi-model panel, and produces criteria-based pass/fail reports.

## Quick Start

### Prerequisites

- **Go 1.24.5+** — to build and run the tool
- **GitHub Copilot CLI** — the SDK communicates with Copilot via the CLI in server mode. Must be installed and authenticated:
  - Install: follow [GitHub Copilot CLI setup](https://docs.github.com/en/copilot/how-tos/set-up/install-copilot-cli)
  - Authenticate: run `copilot` once to complete OAuth device flow, or set `COPILOT_GITHUB_TOKEN` / `GH_TOKEN` env var
  - Without this, the tool falls back to stub mode (no real evaluations)
- **GitHub CLI (`gh`)** — optional but recommended for auth token management
- **For `azure-mcp` config:** `npx` (Node.js) must be available since the Azure MCP server is launched via `npx -y @azure/mcp@latest`

### Run from the repo (recommended)

The repo root has a `go.work` file, so all commands run from the repo root:

```bash
git clone https://github.com/ronniegeraghty/hyoka.git
cd hyoka

# List prompts
go run ./hyoka list

# Run all evaluations (auto-generates trend analysis after)
# Note: requires --all-configs if multiple configs exist
go run ./hyoka run --all-configs

# Filter by service and language
go run ./hyoka run --service storage --language dotnet
```

### Install as a CLI

```bash
go install github.com/ronniegeraghty/hyoka@latest

# When run from the repo root, prompts are auto-detected
cd hyoka
hyoka run --service storage

# Or specify the prompts path explicitly
hyoka run --prompts ~/projects/hyoka/prompts
```

> **Smart path detection:** `hyoka` checks `./prompts` then `../prompts` automatically. Running from the repo root or the `hyoka/` directory both work without extra flags.

## Safety & Guardrails

hyoka includes built-in protections that keep evaluation runs safe, bounded, and predictable by default. No extra flags are needed — everything below is on unless you opt out.

### Generator Guardrails

Every code-generation session is automatically aborted if it exceeds any of these limits:

| Limit | Default | Flag | Purpose |
|-------|---------|------|---------|
| Turn count | 25 | `--max-turns` | Prevents runaway conversations |
| File count | 50 | `--max-files` | Prevents excessive file creation |
| Output size | 1 MB | `--max-output-size` | Prevents oversized outputs (supports KB, MB suffixes) |

When a guardrail trips, the evaluation is marked as failed with a clear reason (e.g., `guardrail: turn count 26 exceeded limit of 25`).

### Safety Boundaries (No Real Azure Resources)

By default, generators receive a system instruction that **prevents real Azure resource provisioning**. The agent will:
- Use mock data, environment variables, and local emulators (Azurite, CosmosDB emulator)
- Generate Bicep/ARM/Terraform templates instead of running live `az` CLI commands
- Use placeholder values like `os.Getenv("AZURE_STORAGE_CONNECTION_STRING")`

Use `--allow-cloud` to opt out and permit real resource provisioning.

### Fan-Out Confirmation

When a run would execute **more than 10 evaluations**, hyoka shows a summary and asks for confirmation before proceeding. Use `-y` / `--yes` to skip the prompt (useful in CI). If multiple configs exist and no `--config` filter is specified, `--all-configs` is required to prevent accidental full-matrix runs.

### Process Lifecycle

hyoka tracks all spawned Copilot processes and terminates them on completion or interrupt (Ctrl+C). The cleanup sequence sends SIGTERM, waits up to 5 seconds, then SIGKILL — no more orphaned processes consuming resources after a run.

### Smart Concurrency

Workers default to **CPU core count** (capped at 8) instead of a hardcoded 4. The `--max-sessions` flag limits total concurrent Copilot instances (default: workers × 3) to prevent resource exhaustion on shared machines.

### Prompt Discovery

`validate` and `run` now fail with a clear error when zero prompts are found. Near-miss detection suggests corrections for misnamed files:

```
no prompts found in ./prompts

Did you mean one of these?
  prompts/storage/data-plane/dotnet/auth-prompt.md → auth.prompt.md
  prompts/key-vault/crud.prompt.txt → crud.prompt.md
```

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `hyoka run` | | Run evaluations against prompts |
| `hyoka list` | `ls` | List prompts matching filters |
| `hyoka configs` | | Show available tool configurations |
| `hyoka validate` | | Validate prompt frontmatter against schema |
| `hyoka check-env` | `env` | Check for required language toolchains and tools |
| `hyoka trends` | | Generate historical trend reports with AI analysis |
| `hyoka report` | | Re-render HTML/MD reports from existing JSON data |
| `hyoka new-prompt` | | Scaffold a new prompt file interactively |
| `hyoka version` | | Print version |

### Filtering

All filter flags work with `run`, `list`, and other prompt-aware commands:

```bash
# By service
hyoka run --service storage

# By language
hyoka run --language dotnet

# Combine filters (AND logic)
hyoka run --service storage --language dotnet --plane data-plane

# By category
hyoka run --category authentication

# By tags
hyoka run --tags identity

# Single prompt by ID
hyoka run --prompt-id storage-dp-dotnet-auth

# Dry run — list matches without executing
hyoka run --service storage --dry-run

# JSON output for scripting
hyoka list --json
```

### Run Command Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--analyze` | `true` | AI-powered trend analysis after run |
| `--skip-trends` | `false` | Skip automatic trend analysis after run |
| `--progress` | `auto` | Progress display mode: `auto`, `live`, `log`, `off` |
| `--skip-tests` | `false` | Skip test generation |
| `--skip-review` | `false` | Skip code review |
| `--verify-build` | `false` | Run build verification on generated code |
| `--stub` | `false` | Use stub evaluator (no Copilot SDK) |
| `--dry-run` | `false` | List matching prompts without running |
| `--workers` | CPU cores (max 8) | Parallel evaluation workers |
| `--max-sessions` | workers × 3 | Maximum concurrent Copilot sessions |
| `--timeout` | `300` | Per-prompt timeout in seconds |
| `-y` / `--yes` | `false` | Skip confirmation prompt for large runs (>10 evaluations) |
| `--all-configs` | `false` | Required when running all configs without a `--config` filter |
| `--max-turns` | `25` | Maximum conversation turns per generation before aborting |
| `--max-files` | `50` | Maximum generated files per evaluation before aborting |
| `--max-output-size` | `1MB` | Maximum total output size per evaluation (supports KB, MB suffixes) |
| `--allow-cloud` | `false` | Allow generated code to provision real Azure resources |
| `--sandbox` | `true` | Alias confirming safe/local-only mode (default behavior) |

### Run Command Examples

```bash
# Skip confirmation for large runs (CI-friendly)
go run ./hyoka run --prompt-id my-prompt --config baseline -y

# Run all prompts × all configs (requires --all-configs + -y for non-interactive)
go run ./hyoka run --all-configs -y

# Tighten guardrails for faster iteration
go run ./hyoka run --max-turns 10 --max-files 20

# Allow real Azure resource provisioning (use with caution)
go run ./hyoka run --allow-cloud

# Limit concurrent Copilot sessions on a shared machine
go run ./hyoka run --max-sessions 4 --workers 2
```

### Validating Prompts

```bash
# Validate all prompts
hyoka validate
```

### Tool Configurations

Each config file defines **one generator model** and a **multi-model review panel**. The `configs/` directory contains configs auto-discovered via `LoadDir()`:

```bash
# List configs
hyoka configs

# Run with a specific config file
hyoka run --config-file configs/baseline-sonnet.yaml --prompt-id storage-dp-dotnet-auth

# Run all configs (default — auto-discovers configs/ directory)
hyoka run --prompt-id storage-dp-dotnet-auth

# Run with a specific config name
hyoka run --config baseline/claude-sonnet-4.5

# Run multiple configs (produces comparison data)
hyoka run --config baseline/claude-sonnet-4.5,azure-mcp/claude-sonnet-4.5
```

#### Custom Configs

Create your own config YAML in the `configs/` directory. The config has two clear sections — `generator` for the code generation agent and `reviewer` for the review/grading plane:

```yaml
configs:
  - name: my-custom-config
    description: "My custom evaluation config"
    generator:
      model: "claude-sonnet-4.5"
      skills:
        - type: remote
          name: azure-keyvault-py
          repo: microsoft/skills
        - type: local
          path: "./skills/generator"
      mcp_servers:
        azure:
          type: local
          command: npx
          args: ["-y", "@azure/mcp@latest"]
          tools: ["*"]
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gemini-3-pro-preview"
        - "gpt-4.1"
      skills:
        - type: local
          path: "./skills/reviewer"
```

Then run with: `hyoka run --config-file configs/my-custom-config.yaml`

> **Backward compatibility:** Legacy top-level fields (`model`, `reviewer_models`, `skill_directories`, `generator_skill_directories`, etc.) still work. They are automatically migrated to the `generator`/`reviewer` sub-structs at parse time.

#### Unified Skills

Skills give agents domain-specific knowledge (SDK patterns, API examples, acceptance criteria) that improve code generation and review quality. The unified `skills:` list replaces the old `skill_directories`, `generator_skill_directories`, and `reviewer_skill_directories` fields.

Each skill has a `type`:

| Type | Fields | Description |
|------|--------|-------------|
| `local` | `path` | Local directory containing a `SKILL.md` file. Supports glob patterns (e.g., `"./skills/generator/*"`) |
| `remote` | `name`, `repo` | Skill fetched from a GitHub repository via `npx skills add` |

**Example — generator with local + remote skills:**

```yaml
generator:
  model: "claude-sonnet-4.5"
  skills:
    - type: remote
      name: azure-keyvault-py
      repo: microsoft/skills
    - type: local
      path: "./skills/generator"
```

**Example — reviewer with local skills:**

```yaml
reviewer:
  models:
    - "claude-opus-4.6"
    - "gpt-4.1"
  skills:
    - type: local
      path: "./skills/reviewer"
```

> **Tip:** The [microsoft/skills](https://github.com/microsoft/skills) repo contains 132+ skills across Azure SDK scenarios. Browse the repo or run `npx skills add microsoft/skills` to see what's available.

See `configs/example-full.yaml` for a complete example with all options.

## Adding a New Prompt

Add a `.prompt.md` file to `prompts/` and run the tool — it discovers prompts automatically.

```bash
# 1. Copy the template
cp templates/prompt-template.prompt.md \
   prompts/<service>/<plane>/<language>/<use-case>.prompt.md

# 2. Edit the file — fill in frontmatter and write your prompt

# 3. Validate
go run ./hyoka validate

# 4. Commit
git add prompts/
git commit -m "prompt: add <service> <plane> <language> <category>"
```

## Repo Structure

```
hyoka/
├── README.md
├── go.work                            # Go workspace (run commands from repo root)
├── configs/                           # Evaluation configs (one generator per file)
│   ├── baseline-sonnet.yaml           # Baseline + Claude Sonnet 4.5
│   ├── baseline-opus.yaml             # Baseline + Claude Opus 4.6
│   ├── baseline-opus-skills.yaml      # Baseline + Claude Opus 4.6 + generator skills
│   ├── baseline-codex.yaml            # Baseline + GPT Codex
│   ├── azure-mcp-sonnet.yaml          # Azure MCP + Claude Sonnet 4.5
│   ├── azure-mcp-opus.yaml            # Azure MCP + Claude Opus 4.6
│   └── azure-mcp-codex.yaml           # Azure MCP + GPT Codex
├── prompts/                           # Prompt library
│   ├── storage/
│   │   ├── data-plane/
│   │   │   ├── dotnet/
│   │   │   │   └── authentication.prompt.md
│   │   │   └── python/
│   │   │       └── pagination-list-blobs.prompt.md
│   │   └── management-plane/
│   │       └── ...
│   └── key-vault/
│       └── ...
├── skills/                            # Copilot skills for eval sessions
│   ├── generator/                     # Skills for the generator agent (install via npx skills add or type: remote)
│   └── reviewer/                      # Skills for the review panel agents
│       ├── code-review-comments/
│       ├── reviewer-build/
│       └── sdk-version-check/
├── hyoka/                              # Go eval tool (hyoka)
│   ├── cmd/hyoka/main.go
│   ├── go.mod / go.sum
│   └── internal/                      # config, prompt, eval, build, report,
│       │                              #   validate, trends, review
│       ├── config/
│       ├── prompt/
│       ├── eval/
│       ├── build/
│       ├── report/
│       ├── trends/
│       ├── review/
│       │   └── rubric.md              # Criteria-based scoring rubric (embedded)
│       └── validate/
├── reports/                           # Evaluation output
│   └── <run-id>/
│       ├── summary.{json,html,md}
│       └── results/<service>/<plane>/<language>/<category>/<config>/
│           └── report.{json,html,md}
├── docs/                              # Documentation
│   ├── getting-started.md
│   └── cleanup-plan.md
└── templates/
    └── prompt-template.prompt.md
```

## Tagging System

Every prompt uses YAML frontmatter:

| Field | Required | Values |
|---|---|---|
| `id` | ✅ | `{service}-{dp\|mp}-{lang}-{category-slug}` |
| `service` | ✅ | `storage`, `key-vault`, `cosmos-db`, `event-hubs`, `app-configuration`, `purview`, `digital-twins`, `identity`, `resource-manager`, `service-bus` |
| `plane` | ✅ | `data-plane`, `management-plane` |
| `language` | ✅ | `dotnet`, `java`, `js-ts`, `python`, `go`, `rust`, `cpp` |
| `category` | ✅ | `authentication`, `pagination`, `polling`, `retries`, `error-handling`, `crud`, `batch`, `streaming`, `auth`, `provisioning` |
| `difficulty` | ✅ | `basic`, `intermediate`, `advanced` |
| `description` | ✅ | What this prompt tests (1-3 sentences) |
| `created` | ✅ | Date (YYYY-MM-DD) |
| `author` | ✅ | GitHub username |
| `sdk_package` | ❌ | SDK package name |
| `doc_url` | ❌ | Library reference docs (API overview, pkg.go.dev, docs.rs) |
| `tags` | ❌ | Free-form tags for additional filtering |

## Roadmap

- **Phase 1:** ✅ Prompt library, build verification, report generation with stub evaluator
- **Phase 2:** ✅ Copilot SDK integration — live agent evaluation with code generation and criteria-based review panel
- **Phase 3:** ✅ Tool matrix, MCP server attachment, skill loading, cross-config comparison
- **Phase 4:** ✅ Guardrails, safety boundaries, smart concurrency, process lifecycle, prompt discovery
- **Phase 5:** In progress — Evaluation quality (check-env, expected_tools, reviewer skills)
- **Phase 6:** Planned — Polish (report re-rendering, embedded CLI, progress bars)

See [`hyoka/README.md`](hyoka/README.md) for full CLI reference and configuration docs.

## License

[MIT](LICENSE)

## Documentation

For detailed documentation beyond this README, see the [`docs/`](docs/README.md) directory:

- **[Getting Started](docs/getting-started.md)** — Prerequisites, installation, first run
- **[CLI Reference](docs/cli-reference.md)** — Complete command and flag reference
- **[Configuration Guide](docs/configuration.md)** — Config YAML schema, skills, MCP servers
- **[Prompt Authoring](docs/prompt-authoring.md)** — Frontmatter schema, evaluation criteria
- **[Architecture Overview](docs/architecture-overview.md)** — End-to-end pipeline
- **[Evaluation Criteria](docs/evaluation-criteria.md)** — Scoring methodology
- **[Reports & Trends](docs/reports-and-trends.md)** — Report format, trend analysis
- **[Guardrails & Safety](docs/guardrails-and-safety.md)** — Limits, safety boundaries
- **[Troubleshooting](docs/troubleshooting.md)** — Common issues, diagnostics
- **[Contributing](docs/contributing.md)** — Build, test, extend hyoka

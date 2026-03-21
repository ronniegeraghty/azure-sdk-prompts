# azure-sdk-prompts

A curated library of prompts for evaluating how well AI agents generate Azure SDK code, paired with a Go evaluation tool (`sdk-eval`) that runs prompts through the Copilot SDK, verifies builds, and produces scored reports.

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

```bash
git clone https://github.com/ronniegeraghty/azure-sdk-prompts.git
cd azure-sdk-prompts

# List prompts
go run ./tool/cmd/sdk-eval list

# Run all evaluations
go run ./tool/cmd/sdk-eval run

# Filter by service and language
go run ./tool/cmd/sdk-eval run --service storage --language dotnet
```

### Install as a CLI

```bash
go install github.com/ronniegeraghty/azure-sdk-prompts/tool/cmd/sdk-eval@latest

# When run from the repo root, prompts are auto-detected
cd azure-sdk-prompts
sdk-eval run --service storage

# Or specify the prompts path explicitly
sdk-eval run --prompts ~/projects/azure-sdk-prompts/prompts
```

> **Smart path detection:** `sdk-eval` checks `./prompts` then `../prompts` automatically. Running from the repo root or the `tool/` directory both work without extra flags.

## Commands

| Command | Description |
|---------|-------------|
| `sdk-eval run` | Run evaluations against prompts |
| `sdk-eval list` | List prompts matching filters |
| `sdk-eval configs` | Show available tool configurations |
| `sdk-eval manifest` | Regenerate manifest.yaml from prompt files |
| `sdk-eval validate` | Validate prompt frontmatter against schema |
| `sdk-eval version` | Print version |

### Filtering

All filter flags work with `run`, `list`, and other prompt-aware commands:

```bash
# By service
sdk-eval run --service storage

# By language
sdk-eval run --language dotnet

# Combine filters (AND logic)
sdk-eval run --service storage --language dotnet --plane data-plane

# By category
sdk-eval run --category authentication

# By tags
sdk-eval run --tags identity

# Single prompt by ID
sdk-eval run --prompt-id storage-dp-dotnet-auth

# Dry run — list matches without executing
sdk-eval run --service storage --dry-run
```

### Validating Prompts

```bash
# Validate all prompts
sdk-eval validate

# Regenerate the manifest
sdk-eval manifest
```

### Tool Configurations

Evaluations can run prompts against different Copilot configurations (models, MCP servers, skills) defined in the `configs/` directory:

```bash
# List configs
sdk-eval configs

# Run with baseline only (no MCP, no skills)
sdk-eval run --config-file configs/baseline.yaml --prompt-id storage-dp-dotnet-auth

# Run with azure-mcp only
sdk-eval run --config-file configs/azure-mcp.yaml --prompt-id storage-dp-dotnet-auth

# Run matrix (both configs — default)
sdk-eval run --prompt-id storage-dp-dotnet-auth

# Run with a specific config name from the default config file
sdk-eval run --config baseline

# Run multiple configs (produces comparison data)
sdk-eval run --config baseline,azure-mcp
```

#### Custom Configs

Create your own config YAML in the `configs/` directory:

```yaml
configs:
  - name: my-custom-config
    description: "My custom evaluation config"
    model: "claude-sonnet-4.5"
    mcp_servers: {}
    skill_directories: []
    available_tools: []
    excluded_tools: []
```

Then run with: `sdk-eval run --config-file configs/my-custom-config.yaml`

## Adding a New Prompt

```bash
# 1. Copy the template
cp templates/prompt-template.prompt.md \
   prompts/<service>/<plane>/<language>/<use-case>.prompt.md

# 2. Edit the file — fill in frontmatter and write your prompt

# 3. Validate
go run ./tool/cmd/sdk-eval validate

# 4. Regenerate the manifest
go run ./tool/cmd/sdk-eval manifest

# 5. Commit
git add prompts/ manifest.yaml
git commit -m "prompt: add <service> <plane> <language> <category>"
```

## Repo Structure

```
azure-sdk-prompts/
├── README.md
├── manifest.yaml                      # Auto-generated prompt index
├── configs/                           # Evaluation config matrix
│   ├── all.yaml                       # Both configs (default for matrix runs)
│   ├── baseline.yaml                  # No MCP, no skills — raw Copilot
│   └── azure-mcp.yaml                # Azure MCP server attached
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
├── tool/                              # Go eval tool (sdk-eval)
│   ├── cmd/sdk-eval/main.go
│   ├── go.mod / go.sum
│   ├── internal/                      # config, prompt, eval, build, report,
│   │   │                              #   manifest, validate
│   │   ├── config/
│   │   ├── prompt/
│   │   ├── eval/
│   │   ├── build/
│   │   ├── report/
│   │   ├── manifest/
│   │   └── validate/
│   └── testdata/
├── reports/                           # Evaluation output
│   └── runs/<timestamp>/
│       ├── summary.json
│       └── results/<service>/<plane>/<language>/<category>/<config>/
│           └── report.json
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
| `doc_url` | ❌ | Link to the docs page being evaluated |
| `tags` | ❌ | Free-form tags for additional filtering |

## Roadmap

- **Phase 1 (current):** Prompt library, build verification, report generation with stub evaluator
- **Phase 2:** Copilot SDK integration — live agent evaluation with code generation and LLM-as-judge scoring

See [`tool/README.md`](tool/README.md) for full CLI reference and configuration docs.

## License

[MIT](LICENSE)

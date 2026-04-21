# hyoka

**A comprehensive evaluation tool for AI agents.** hyoka runs prompts through AI agents (GitHub Copilot SDK), scores generated outputs against extensible grader criteria, and produces detailed pass/fail reports with workspace delta tracking.

## Installation

**Prerequisites:**
- **Go 1.26.1+**
- **GitHub Copilot CLI** — must be installed and authenticated ([setup guide](https://docs.github.com/en/copilot/how-tos/set-up/install-copilot-cli))

**From source (recommended for development):**

```bash
git clone https://github.com/ronniegeraghty/hyoka.git
cd hyoka
go install ./...
hyoka version
```

**From GitHub (latest release):**

```bash
go install github.com/ronniegeraghty/hyoka@latest
hyoka version
```

## Quick Start

Run your first evaluation in under 5 minutes:

```bash
# 1. List available prompts
hyoka list --service key-vault --language python

# 2. Run a single evaluation
hyoka run \
  --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6

# 3. View results in your browser
hyoka serve
# Open http://localhost:8080
```

**What just happened?** hyoka:
1. Loaded the `key-vault-dp-python-crud` prompt
2. Spawned a Copilot session with Claude Opus 4.6
3. Captured the agent's output
4. Evaluated it against 5 grader types (builder, complexity, prompt adherence, behavior, AI review)
5. Produced a pass/fail report with detailed grading breakdown

## Examples

### Sample Prompt

Prompts are markdown files with YAML frontmatter. Here's a minimal example:

```markdown
---
id: key-vault-dp-python-crud
properties:
  service: key-vault
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: Create, read, update, and delete secrets using Azure Key Vault SDK for Python
  sdk_package: azure-keyvault-secrets
  doc_url: https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme
  created: '2025-01-15'
  author: ronniegeraghty
tags:
- secrets
- crud
---

Write a Python script that demonstrates CRUD operations for Azure Key Vault secrets. The script should:

1. Create a secret with name and value
2. Retrieve the secret value
3. Update the secret value
4. Delete the secret
5. Use DefaultAzureCredential for authentication
```

**See more:** [Prompt Authoring Guide](docs/prompt-authoring.md)

### Sample Config

Configs define the generator model, reviewer models, and available tools:

```yaml
configs:
  - name: baseline/claude-opus-4.6
    description: "Claude Opus 4.6 with no MCP or skills"
    generator:
      model: "claude-opus-4.6"
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gpt-5.3-codex"
```

**See more:** [Configuration Guide](docs/configuration.md)

### Sample Criteria

Criteria files define grading rules matched by prompt attributes:

```yaml
attributes:
  service: key-vault
  language: python
criteria:
  - name: SDK Package Import
    type: code_pattern
    pattern: 'from azure\.keyvault\.secrets import'
    weight: 0.15
  - name: DefaultAzureCredential Usage
    type: code_pattern
    pattern: 'DefaultAzureCredential\(\)'
    weight: 0.20
```

**See more:** [Grader Configuration Schema](docs/grader-config-schema.md)

## Commands

hyoka provides commands for running evaluations, browsing results, and managing prompts. For complete flag documentation, see [CLI Reference](docs/cli-reference.md).

| Command | Description |
|---------|-------------|
| `hyoka run` | Run evaluations against prompts with specified configs |
| `hyoka list` | List prompts matching filter criteria |
| `hyoka serve` | Launch local web UI for browsing reports |
| `hyoka compare` | Compare evaluation results (configs, runs, or time periods) |
| `hyoka init` | Scaffold a `.hyoka` project directory |
| `hyoka validate` | Validate prompt frontmatter against schema |
| `hyoka check-env` | Check for required language toolchains and tools |
| `hyoka clean` | Remove stale session state and orphaned processes |

**Filtering prompts:**

```bash
# By service
hyoka run --service key-vault --config baseline/claude-opus-4.6

# By language
hyoka run --language python --config baseline/claude-opus-4.6

# Combine filters (AND logic)
hyoka run --service storage --language dotnet --plane data-plane \
  --config baseline/claude-opus-4.6

# Single prompt by ID
hyoka run --prompt-id storage-dp-dotnet-auth \
  --config baseline/claude-opus-4.6

# Dry run — list matches without executing
hyoka run --service storage --config baseline/claude-opus-4.6 --dry-run
```

**See all commands and flags:** [CLI Reference](docs/cli-reference.md)

## Safety & Guardrails

hyoka includes built-in protections that keep evaluation runs safe, bounded, and predictable by default.

### Generator Guardrails

Every evaluation session is automatically aborted if it exceeds any of these limits:

| Limit | Default | Flag | Purpose |
|-------|---------|------|---------|
| Session actions | 50 | `--max-session-actions` | Limits reasoning, response, and tool call actions per session |
| File count | 50 | `--max-files` | Prevents excessive file creation (counts new files + deleted starters) |
| Output size | — | — | Removed in #566 — review no longer inlines file contents and review-side caps prevent runaway memory |

Prompts can override defaults via frontmatter. Resolution order: prompt frontmatter > config YAML > CLI flag > engine default.

### Safety Boundaries

By default, generators **prevent real Azure resource provisioning**. Agents use:
- Mock data, environment variables, and local emulators
- Bicep/ARM/Terraform templates instead of live `az` CLI commands
- Placeholder values like `os.Getenv("AZURE_STORAGE_CONNECTION_STRING")`

Use `--allow-cloud` to permit real resource provisioning.

### Other Protections

- **Fan-out confirmation:** Prompts for confirmation when >10 evaluations would run (skip with `-y`)
- **Process lifecycle:** Auto-terminates spawned Copilot processes on completion or interrupt (SIGTERM → SIGKILL)
- **Smart concurrency:** Defaults to CPU core count (capped at 8); `--max-sessions` limits concurrent instances
- **Prompt discovery:** Clear error messages with near-miss detection for misnamed files

**See more:** [Guardrails Documentation](docs/guardrails.md)

## Contributing

We welcome contributions! To get started:

1. **Read the contributing guide:** [CONTRIBUTING.md](CONTRIBUTING.md)
2. **Check the architecture docs:** [docs/architecture.md](docs/architecture.md)
3. **Browse open issues:** [GitHub Issues](https://github.com/ronniegeraghty/hyoka/issues)

**Quick development loop:**

```bash
# Clone and build
git clone https://github.com/ronniegeraghty/hyoka.git
cd hyoka
go build ./...

# Run tests
go test -race ./...

# Test site (frontend)
cd site && npm test

# Rebuild the site AND refresh the embedded bundle that `hyoka serve` ships.
# Required whenever you change anything under site/src/** — a CI check
# (site-embed-freshness) will fail the PR otherwise.
make site-embed

# Test with a live eval (fastest feedback)
hyoka run --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6
```

**See also:**
- [Roadmap](docs/roadmap.md) — completed phases and what's planned
- [CLI Reference](docs/cli-reference.md) — all commands and flags
- [Configuration Guide](docs/configuration.md) — config file format and options

## License

[MIT](LICENSE)

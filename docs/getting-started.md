# Getting Started with hyoka

This guide walks you through cloning the repo, running your first evaluation, and viewing results.

## Prerequisites

| Tool | Version | Check |
|------|---------|-------|
| Go | 1.26.1+ | `go version` |
| GitHub Copilot CLI | Latest | `copilot --version` |
| Git | Any | `git --version` |
| Node.js (for Azure MCP config) | 18+ | `node --version` |

### Copilot Authentication

The tool uses the Copilot SDK which requires an authenticated Copilot CLI:

```bash
# Option A: OAuth device flow (interactive)
copilot

# Option B: Environment variable
export COPILOT_GITHUB_TOKEN="your-token"
```

Without Copilot auth, the tool falls back to **stub mode** (no real agent evaluations).

## 1. Clone and Install

```bash
git clone https://github.com/ronniegeraghty/hyoka.git
cd hyoka
go install ./...
```

Verify the installation:

```bash
hyoka version
```

Expected output:
```
hyoka version v0.x.x
```

Check your environment:

```bash
hyoka check-env
```

This reports which language toolchains and tools are available.

## 2. Explore Available Prompts

```bash
# List all prompts
hyoka list

# Filter by service
hyoka list --service storage

# JSON output (for scripting)
hyoka list --json
```

## 3. Run Your First Evaluation

Start with a single prompt to keep it quick:

```bash
hyoka run \
  --prompt-id storage-dp-dotnet-auth \
  --config baseline
```

Or use **stub mode** to test the pipeline without Copilot:

```bash
hyoka run \
  --prompt-id storage-dp-dotnet-auth \
  --stub
```

> **Confirmation prompt:** If a run would execute more than 10 evaluations, hyoka asks for confirmation before proceeding. Use `-y` to skip in CI or scripted runs. If you run without a `--config` filter and multiple configs exist, add `--all-configs` to confirm you intend to run all of them.

Expected output:
```
Found 1 prompt(s), 1 config(s) → 1 evaluation(s)
Using Copilot SDK evaluator

Run Summary:
  Run ID:      20250728-143022
  Evaluations: 1
  Passed:      1
  Failed:      0
  Errors:      0
  Duration:    45.20s

────────────────────────────────────────────────────────
📊 Generating trend analysis...
...
```

## 4. View Results

Reports are generated in `reports/<run-id>/`:

```bash
# Browse reports in the web dashboard
hyoka serve

# Or view the JSON summary directly
cat reports/<run-id>/summary.json | python3 -m json.tool
```

The summary includes:
- **Prompt × Config Matrix** — pass/fail status with scores
- **Duration Analysis** — min/avg/max per config and prompt
- **Config Comparison** — side-by-side pass rates
- **Tool Usage** — aggregate tool call statistics
- **Detailed Results** — individual eval data

Individual reports at `reports/<run-id>/results/.../report.json` contain the full agent session: prompt, reasoning, tool calls, generated code, and review scores. Markdown reports are also generated alongside each JSON report.

## 5. Run Trend Analysis

After multiple runs, generate trend reports:

```bash
hyoka trends
```

This scans all past runs and produces:
- Pass rate timelines
- Duration trends
- Config comparisons
- AI-powered insights (enabled by default)

Open the trend report:

```bash
hyoka trends --open
```

## 6. Create a New Prompt

Use the interactive scaffolder:

```bash
hyoka new-prompt
```

Or copy the template manually:

```bash
cp templates/prompt-template.prompt.md \
   prompts/<service>/<plane>/<language>/<slug>.prompt.md
```

Validate after editing:

```bash
hyoka validate
```

## Common Workflows

### Run a full evaluation matrix

```bash
# All prompts × all configs (requires --all-configs since multiple configs exist)
hyoka run --all-configs

# Skip confirmation prompt for CI
hyoka run --all-configs -y
```

### Run with specific configs

```bash
# Just baseline
hyoka run --config baseline

# Both configs for one service
hyoka run --service storage

# Run with multiple configs (compare baseline vs azure-mcp):
hyoka run --service identity --language python \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6" \
  --log-level debug --log-file hyoka-debug.log
```

> **Tip:** The `--config` flag takes config *names* (the `name:` field inside the YAML file), not filenames. For example, `configs/azure-mcp-opus.yaml` defines a config named `azure-mcp/claude-opus-4.6`.

### Adjust guardrails

```bash
# Tighter limits for faster iteration
hyoka run --max-session-actions 10 --max-files 20

# Allow real Azure resource provisioning
hyoka run --allow-cloud

# Limit concurrent sessions on a shared machine
hyoka run --max-sessions 4 --workers 2
```

### Useful Flag Combos

```bash
# Debug + log file (keeps terminal clean, logs to file)
hyoka run --config baseline --log-level debug --log-file hyoka-debug.log

# Monitor resource usage during evaluation
hyoka run --config baseline --monitor-resources

# Skip review for quick generation iteration
hyoka run --config baseline --skip-review

# Dry run to preview what would be evaluated
hyoka run --service storage --dry-run
```

### Re-render reports after template changes

```bash
hyoka report --all
```

### Skip AI analysis for faster iteration

```bash
hyoka run --skip-trends
hyoka trends --no-analyze
```

## Guardrails & Safety

hyoka applies sensible defaults to keep evaluation runs safe and bounded. All limits are configurable via CLI flags.

### Generator Limits

Every code-generation session is automatically aborted if it exceeds:

| Limit | Default | Flag |
|-------|---------|------|
| Session actions | 50 | `--max-session-actions` |

When a limit is hit, the evaluation stops and the report shows the specific guardrail that triggered (e.g., `guardrail: file count 51 exceeded limit of 50`).

### Safety Boundaries

By default, generators are instructed **not to provision real Azure resources**. They'll use:
- Local emulators (Azurite, CosmosDB emulator)
- Environment variable placeholders (`os.Getenv("AZURE_STORAGE_CONNECTION_STRING")`)
- Bicep/ARM/Terraform templates instead of live `az` CLI commands

To opt out: `--allow-cloud`.

### Process Cleanup

All spawned Copilot processes are tracked and terminated on run completion or Ctrl+C. The cleanup sends SIGTERM, waits up to 5 seconds, then escalates to SIGKILL.

### Prompt Discovery

If `validate` or `run` finds zero prompts, it scans for near-miss filenames and suggests fixes:
- `auth-prompt.md` → `auth.prompt.md` (hyphen instead of dot)
- `crud.prompt.txt` → `crud.prompt.md` (wrong extension)

## Browse Results with Serve

Start the built-in report viewer:

```bash
hyoka serve
```

This launches a local web server at `http://localhost:8080` with a dashboard for browsing all evaluation runs and their reports.

## Next Steps

- [CLI Reference](cli-reference.md) — Full command and flag documentation
- [Configuration Guide](configuration.md) — Config YAML format and options
- [Prompt Authoring Guide](prompt-authoring.md) — How to write evaluation prompts
- [Guardrails and Safety](guardrails.md) — Limits, process cleanup, and safety boundaries
- [Architecture Overview](architecture.md) — How hyoka works end-to-end
- [Contributing Guide](../CONTRIBUTING.md) — Building, testing, and adding features

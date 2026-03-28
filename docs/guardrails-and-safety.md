# Guardrails & Safety

hyoka includes built-in protections that keep evaluation runs safe, bounded, and predictable. All guardrails are on by default — no extra configuration needed.

## Generator Guardrails

Every code-generation session is automatically aborted if it exceeds any of these limits:

| Limit | Default | Flag | Description |
|-------|---------|------|-------------|
| Turn count | 25 | `--max-turns` | Maximum conversation turns before aborting |
| File count | 50 | `--max-files` | Maximum generated files per evaluation |
| Output size | 1 MB | `--max-output-size` | Maximum total output size (supports `KB`, `MB` suffixes) |

### How Guardrails Trigger

After code generation completes, hyoka checks each limit:

1. **Turn count:** Counts `assistant.turn_end` events in the session transcript
2. **File count:** Counts files in the generator workspace
3. **Output size:** Sums the byte size of all generated files

When a guardrail trips, the evaluation is marked as failed with a clear reason:

```
guardrail: turn count 26 exceeded limit of 25
guardrail: file count 51 exceeded limit of 50
guardrail: size 1048577 bytes exceeded limit of 1048576
```

### Adjusting Limits

```bash
# Tighter limits for faster iteration
go run ./hyoka run --max-turns 10 --max-files 20 --max-output-size 512KB

# Looser limits for complex prompts
go run ./hyoka run --max-turns 50 --max-files 100 --max-output-size 5MB
```

## Safety Boundaries

By default, generators receive a system instruction that **prevents real Azure resource provisioning**. The agent will:

- Use mock data, environment variables, and local emulators (Azurite, CosmosDB emulator)
- Generate Bicep/ARM/Terraform templates instead of running live `az` CLI commands
- Use placeholder values like `os.Getenv("AZURE_STORAGE_CONNECTION_STRING")`

This ensures evaluations are safe to run on any machine without risking:

- Unintended Azure resource creation or deletion
- Unexpected billing charges
- Data exposure or modification in production environments

### Opting Out

To allow real Azure resource provisioning:

```bash
go run ./hyoka run --allow-cloud
```

> **⚠️ Use with caution.** This permits the agent to run `az` CLI commands, create resources, and interact with live Azure services. Only use in sandboxed environments with appropriate permissions.

## Fan-Out Confirmation

When a run would execute **more than 10 evaluations**, hyoka shows a summary and asks for confirmation:

```
Found 87 prompt(s), 7 config(s) → 609 evaluation(s)

This will run 609 evaluations. Continue? [y/N]
```

### Skipping Confirmation

```bash
# Skip prompt for CI/scripted runs
go run ./hyoka run --all-configs -y
```

### Multi-Config Safety

If multiple configs exist and no `--config` filter is specified, hyoka requires `--all-configs` to prevent accidental full-matrix runs:

```bash
# ERROR: multiple configs found, specify --config or --all-configs
go run ./hyoka run

# OK: explicitly running all configs
go run ./hyoka run --all-configs
```

## Process Lifecycle

hyoka tracks all spawned Copilot processes and ensures cleanup on completion or interrupt.

### Process Tracking

The `ProcessTracker` registers every child process (Copilot SDK sessions, MCP servers):

```
[Register] PID 12345 — copilot session for storage-dp-python-crud
[Register] PID 12346 — MCP server (azure)
```

### Cleanup Sequence

On run completion or Ctrl+C (SIGINT/SIGTERM):

1. **SIGTERM** sent to all tracked processes
2. **Wait up to 5 seconds** for graceful shutdown
3. **SIGKILL** sent to any remaining processes
4. **Orphan detection** warns about processes that outlived their session

### Signal Handling

- **First Ctrl+C:** Graceful shutdown — terminate all tracked processes, write partial reports
- **Second Ctrl+C:** Force exit immediately

## Smart Concurrency

### Workers

Workers control how many evaluations run in parallel:

```bash
# Default: CPU core count (capped at 8)
go run ./hyoka run

# Manual setting
go run ./hyoka run --workers 2
```

### Session Limiting

The `--max-sessions` flag limits total concurrent Copilot instances to prevent resource exhaustion:

```bash
# Default: workers × 3
go run ./hyoka run

# Limit on shared machines
go run ./hyoka run --max-sessions 4 --workers 2
```

Each evaluation may spawn multiple sessions (generation + review), so `max-sessions` should be higher than `workers`.

## Timeout Control

Independent timeouts for each evaluation phase:

| Phase | Flag | Default | Description |
|-------|------|---------|-------------|
| Generation | `--generate-timeout` | 600s | Time for the agent to generate code |
| Build | `--build-timeout` | 300s | Time for build verification |
| Review | `--review-timeout` | 300s | Time for the review panel |

```bash
# Quick iterations with shorter timeouts
go run ./hyoka run --generate-timeout 120 --review-timeout 60

# Complex prompts that need more time
go run ./hyoka run --generate-timeout 900
```

When a timeout triggers, the evaluation is marked as failed with `ErrorCategory: "timeout"`.

## Prompt Discovery Safety

If `validate` or `run` finds zero prompts in the specified directory, it fails with a clear error instead of silently proceeding:

```
no prompts found in ./prompts
```

### Near-Miss Detection

hyoka scans for common naming mistakes and suggests fixes:

```
no prompts found in ./prompts

Did you mean one of these?
  prompts/storage/data-plane/dotnet/auth-prompt.md → auth.prompt.md
  prompts/key-vault/crud.prompt.txt → crud.prompt.md
```

Detected patterns:
- `*-prompt.md` instead of `*.prompt.md` (hyphen vs dot)
- `*.prompt.txt` instead of `*.prompt.md` (wrong extension)
- `.md` files with YAML frontmatter (might be prompts with wrong naming)

## Workspace Containment

After generation, hyoka validates that all files stayed within the workspace:

1. **Snapshot** home directory and CWD before evaluation
2. **Detect** files created outside the workspace during generation
3. **Recover** code files moved back to the workspace
4. **Clean up** junk directories (`__pycache__`, `node_modules`, `.venv`, `target`, `dist`)
5. **Validate** containment — warn if files escaped

This prevents agents from modifying the repository, leaving files in unexpected locations, or polluting the host environment.

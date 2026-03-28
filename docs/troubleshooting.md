# Troubleshooting

Common issues when running hyoka and how to diagnose them.

## Diagnostic Tools

### Debug Logging

Enable detailed logging to see what hyoka is doing:

```bash
# Debug output to stderr
go run ./hyoka run --log-level debug --prompt-id my-prompt

# Redirect logs to a file (keeps terminal clean)
go run ./hyoka run --log-level debug --log-file hyoka-debug.log --prompt-id my-prompt
```

Log levels: `debug`, `info`, `warn`, `error`

### Environment Check

Verify that all required tools are installed:

```bash
go run ./hyoka check-env
```

This checks language toolchains (dotnet, python, go, node, java, rust), Copilot CLI, and MCP prerequisites.

### Prompt Validation

Verify prompt files are correctly formatted:

```bash
go run ./hyoka validate
```

This checks frontmatter fields, naming conventions, and section presence.

## Common Issues

### "no prompts found"

**Symptom:** `no prompts found in ./prompts`

**Causes:**
1. Running from the wrong directory — hyoka expects `./prompts` relative to CWD
2. Prompt files don't end in `.prompt.md`
3. Custom `--prompts` path doesn't exist

**Fixes:**
- Run from the repo root: `cd /path/to/hyoka && go run ./hyoka run`
- Check file naming: files must end in `.prompt.md` (not `.md`, not `-prompt.md`)
- Check near-miss suggestions in the error output

### "config not found"

**Symptom:** Error about missing or unrecognized config

**Causes:**
1. Config name doesn't match any file in `configs/`
2. Typo in `--config` flag value
3. Config directory not found

**Fixes:**
- List available configs: `go run ./hyoka configs`
- Check config names: `go run ./hyoka configs --config-dir ./configs`
- Use `--config-file` to point to a specific file

### Copilot Authentication Failures

**Symptom:** "failed to create Copilot session" or stub mode fallback

**Causes:**
1. Copilot CLI not installed
2. Not authenticated
3. Token expired

**Fixes:**
```bash
# Check Copilot CLI is installed
copilot --version

# Authenticate via OAuth device flow
copilot

# Or set token directly
export COPILOT_GITHUB_TOKEN="your-token"
# or
export GH_TOKEN="your-token"
```

### Orphaned Copilot Processes

**Symptom:** `copilot` processes running after hyoka exits, high CPU/memory usage

**Causes:**
1. hyoka was force-killed (kill -9) without cleanup
2. Bug in process tracking (rare)

**Fixes:**
```bash
# Find orphaned processes
ps aux | grep copilot

# Kill specific processes
kill <PID>
```

hyoka normally handles cleanup automatically via the process tracker. Use Ctrl+C for graceful shutdown instead of `kill -9`.

### Timeout Errors

**Symptom:** `ErrorCategory: "timeout"` in report, evaluation marked as failed

**Causes:**
1. Complex prompt requires more generation time
2. Slow network (for MCP/skill fetching)
3. Model is generating excessively verbose output

**Fixes:**
```bash
# Increase generation timeout
go run ./hyoka run --generate-timeout 900

# Increase review timeout
go run ./hyoka run --review-timeout 600

# Increase all phase timeouts
go run ./hyoka run --generate-timeout 900 --build-timeout 600 --review-timeout 600
```

### Guardrail Triggers

**Symptom:** Report shows `guardrail: turn count X exceeded limit of Y`

**Causes:**
1. Agent is in a loop (asking clarifying questions, retrying)
2. Complex prompt legitimately needs more turns/files
3. Default limits too restrictive for the scenario

**Fixes:**
```bash
# Increase limits
go run ./hyoka run --max-turns 50 --max-files 100 --max-output-size 5MB
```

Check the session events in the HTML report to see what the agent was doing — if it's looping, the prompt may need refinement rather than higher limits.

### Build Verification Failures

**Symptom:** Build result shows failure when using `--verify-build`

**Causes:**
1. Missing language toolchain (check with `check-env`)
2. Generated code has syntax errors (expected — this is what build verification tests)
3. Missing dependencies (agent didn't include package management files)

**Fixes:**
- Install required toolchains (e.g., `dotnet`, `python3`, `go`)
- Build failures are informational — they don't prevent the review from running
- Check the build result in the HTML report for specific error messages

### Zero Generated Files

**Symptom:** `ErrorCategory: "no_files"`, evaluation failed

**Causes:**
1. Agent didn't produce any files (conversation went off-track)
2. Files were created outside the workspace and not recovered
3. Workspace containment issue

**Fixes:**
- Check session events in the HTML report to see what happened
- Enable debug logging to see workspace recovery: `--log-level debug`
- Try a different model or add skills for better domain knowledge

### MCP Server Launch Failures

**Symptom:** Azure MCP config fails to start, error about `npx` or `@azure/mcp`

**Causes:**
1. Node.js not installed (MCP server requires `npx`)
2. Network issues (npm can't fetch `@azure/mcp@latest`)
3. Node.js version too old

**Fixes:**
```bash
# Check Node.js installation
node --version  # need 18+
npx --version

# Test MCP server launch manually
npx -y @azure/mcp@latest
```

## Reading Reports

### Finding Your Results

```bash
# List recent runs
ls -lt reports/ | head -5

# Open the latest summary
open reports/$(ls -t reports/ | head -1)/summary.html
```

### Understanding Failures

In the HTML report, check:

1. **Error category** — timeout, sdk_error, no_files, generation_failure, review_failure
2. **Session events** — expand the action history to see what the agent did step-by-step
3. **Guardrail status** — check if any limits were hit
4. **Review panel** — see individual reviewer scores to understand disagreements
5. **Generated code** — review the actual output to assess quality

### Comparing Configs

Run the same prompt with multiple configs and compare:

```bash
go run ./hyoka run --prompt-id my-prompt --all-configs
```

The summary HTML includes a config comparison matrix showing pass rates side-by-side.

## Getting Help

- **Check environment:** `go run ./hyoka check-env`
- **Validate prompts:** `go run ./hyoka validate`
- **Debug logging:** `--log-level debug --log-file debug.log`
- **Dry run:** `go run ./hyoka run --dry-run` (see what would run without executing)
- **Stub mode:** `go run ./hyoka run --stub` (test pipeline without Copilot)

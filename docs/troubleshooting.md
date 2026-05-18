# Troubleshooting Guide

This guide helps you diagnose and resolve common issues when running hyoka evaluations.

## Tool Load Failures

If your evaluation aborts with an error like:

```
Error: required skill 'generator-skills' failed to load: SDK did not report skill as loaded
```

or

```
Error: required mcp 'azure' failed to load: SDK did not report MCP server as loaded
```

This means hyoka started a Copilot session but the SDK failed to load one or more of your configured tools within 10 seconds. The evaluation is immediately aborted to prevent silent failures.

### Diagnosis

To see detailed diagnostic information about tool load failures, run with debug logging:

```bash
hyoka run --prompt-id <prompt-id> --config <config> \
  --log-level debug --log-file hyoka-debug.log

# Search the log for tool-related errors
grep -E "tool|skill|mcp|verifier|failed" hyoka-debug.log
```

The debug log will show:
- Which tools were expected to load
- Which tools actually loaded according to the SDK
- Timestamps and SDK event details

### Common Causes and Fixes

#### Skill not found or `SKILL.md` missing

**Symptom:** `"required skill 'X' failed to load: SDK did not report skill as loaded"`

**Causes:**
- The configured skill directory doesn't exist
- The skill directory exists but doesn't contain a `SKILL.md` file
- The path is relative to the wrong directory (e.g., config file's directory vs. current working directory)

**How to fix:**

1. **Check the skill directory exists:**
   ```bash
   ls -la skills/generator/
   ```

2. **Verify each skill directory contains `SKILL.md`:**
   ```bash
   find skills/generator -name "SKILL.md"
   ```

3. **Check your config's `path` field:**
   - Paths in `generator.tools` are relative to the **config file's directory**
   - If your config is in `configs/my-eval.yaml` and you specify `path: ./skills/generator`, hyoka looks for `configs/./skills/generator`
   - Use `../skills/generator` instead (one level up from `configs/`)

   Example:
   ```yaml
   generator:
     tools:
       - name: generator-skills
         type: skill
         source: local
         path: ../skills/generator  # Correct: relative to config file
   ```

4. **Test with an absolute path** to isolate path resolution issues:
   ```bash
   # First, find the absolute path
   cd skills/generator && pwd
   # /home/user/hyoka/skills/generator
   
   # Update your config temporarily to test
   path: /home/user/hyoka/skills/generator
   ```

#### Glob pattern produces no matches

**Symptom:** Skill directory configured with a glob pattern (e.g., `./skills/generator/*`) loads zero skills

**Cause:** The glob pattern doesn't match any directories, or all directories are filtered out

**How to fix:**

1. **Check what the glob matches:**
   ```bash
   ls skills/generator/
   ls skills/generator/*  # Should list subdirectories
   ```

2. **Ensure only directories are matched** — if using a glob, only directories containing `SKILL.md` are included:
   ```bash
   find skills/generator -type d -name "SKILL.md"
   ```

3. **If glob matches but skills still don't load**, verify the `SKILL.md` in each directory is readable:
   ```bash
   cat skills/generator/*/SKILL.md
   ```

#### Remote skill download fails

**Symptom:** `"required skill 'X' failed to load"` for a remote skill configured with `source: remote`

**Causes:**
- Network connectivity issue
- GitHub repository doesn't exist or is inaccessible
- No credentials to access a private repository
- The skill name doesn't exist in the specified repository

**How to fix:**

1. **Test GitHub access:**
   ```bash
   gh repo view microsoft/skills
   ```

2. **Check credentials:**
   ```bash
   gh auth status
   ```
   If you're not authenticated, run:
   ```bash
   gh auth login
   ```

3. **Manually test the skill fetch** using the same command hyoka uses:
   ```bash
   npx skills add microsoft/skills --directory .skills-cache/microsoft/skills
   ```
   If this fails, hyoka will also fail.

4. **Check your config's `repo` field:**
   ```yaml
   tools:
     - name: azure-keyvault-py
       type: skill
       source: remote
       repo: microsoft/skills  # Must be owner/repo format
   ```

#### MCP server fails to start

**Symptom:** `"required mcp 'azure' failed to load: SDK did not report MCP server as loaded"`

**Causes:**
- The `command` specified doesn't exist or isn't in PATH
- The command exists but the arguments are invalid
- The MCP server starts but crashes immediately
- Node packages (for `npx` commands) aren't installed

**How to fix:**

1. **Verify the command exists:**
   ```bash
   which npx
   ```
   If it's not found, install Node.js:
   ```bash
   # macOS
   brew install node
   
   # Linux (Ubuntu/Debian)
   sudo apt-get install nodejs npm
   
   # Windows
   choco install nodejs
   ```

2. **Test the command manually:**
   ```bash
   # Example: test the Azure MCP server
   npx -y @azure/mcp@latest server start
   ```
   If this fails or hangs, fix it before retrying hyoka.

3. **Check your config's MCP entry:**
   ```yaml
   tools:
     - name: azure
       type: mcp
       command: npx
       args: ["-y", "@azure/mcp@latest", "server", "start"]
       mcp_tools: ["*"]
   ```

4. **Ensure `mcp_tools` is set:**
   Without this field, the server starts but its tools won't be available to the agent. Always include:
   ```yaml
   mcp_tools: ["*"]  # or ["specific-tool-1", "specific-tool-2"]
   ```

5. **Increase diagnostics** — check if the server is hanging or crashing:
   ```bash
   # Run with debug logging
   hyoka run --prompt-id <id> --config <cfg> --log-level debug --log-file hyoka-debug.log
   
   # Check for timeout messages in the log
   grep -i "timeout\|hang\|crash" hyoka-debug.log
   ```

#### SDK timeout (10 seconds)

**Symptom:** `"required skill/mcp X failed to load"` and debug logs show tool events didn't arrive

**Cause:** The Copilot SDK took longer than 10 seconds to load the tools

**How to fix:**

This is rare and usually indicates one of the above issues that's just slower than expected. Steps:

1. **Verify network connectivity:**
   ```bash
   ping github.com
   ```

2. **Check for slow remote skills:**
   - Pre-download remote skills to `.skills-cache/` to speed up subsequent runs
   - Manually run `npx skills add <repo>` first

3. **Verify system resources:**
   - Check CPU, memory, and disk usage while hyoka runs
   - Close other resource-heavy applications

4. **Report the issue** if you've confirmed the tools are valid and the timeout persists

### Quick Checklist

Before re-running an evaluation after a tool load failure:

- [ ] All skill directories exist and contain `SKILL.md`
- [ ] All `path` fields in the config are correct (relative to the config file)
- [ ] All `command` fields for MCP servers are valid and in PATH
- [ ] All remote skill repositories are accessible (GitHub auth working)
- [ ] All MCP tools have `mcp_tools: ["*"]` or explicit tool names specified
- [ ] Network connectivity is working (for remote skills/MCP)
- [ ] Debug log shows no error messages related to tools

---

## Other Issues

### Evaluation times out

If an evaluation runs longer than expected or times out:

1. **Check guardrail limits:**
   ```bash
   # Review your config's limits section
   grep -A5 "limits:" <config-file>
   
   # Or check the CLI flags you used
   hyoka run --help | grep -E "max-turns|max-files|max-session"
   ```

2. **Increase limits if needed** (in config or CLI):
   ```bash
   hyoka run --prompt-id <id> --config <cfg> \
     --max-turns 50 --max-files 100 --max-session-actions 100
   ```

### Model not found or not authorized

If you see errors about a model not being available:

1. **Check the model is installed:**
   ```bash
   copilot models list
   ```

2. **Verify authentication:**
   ```bash
   copilot auth status
   ```

3. **Check your config's `generator.model` and `reviewer.models`:**
   ```yaml
   generator:
     model: claude-opus-4.6
   reviewer:
     models:
       - claude-opus-4.6
       - gpt-5.3-codex
   ```

### Session initialization fails

If the Copilot session fails to start:

1. **Check Copilot CLI is installed and working:**
   ```bash
   copilot --version
   ```

2. **Verify you have an active Copilot license or trial:**
   ```bash
   copilot auth status
   ```

3. **Run with debug logging:**
   ```bash
   hyoka run --prompt-id <id> --config <cfg> --log-level debug --log-file hyoka-debug.log
   ```

---

## Getting Help

If you're still stuck:

1. **Check the logs:**
   ```bash
   # Run with debug logging
   hyoka run --prompt-id <id> --config <cfg> \
     --log-level debug --log-file hyoka-debug.log
   
   # Search for ERROR, WARN, or specific keywords
   grep -E "ERROR|WARN|tool|skill|mcp" hyoka-debug.log
   ```

2. **Check the report:**
   ```bash
   # Browse the generated report
   hyoka serve
   ```

3. **Open an issue** on GitHub with:
   - The error message (full text from `hyoka run` output)
   - The relevant section from the debug log
   - Your config file (sanitized of any secrets)
   - Output of `hyoka check-env`

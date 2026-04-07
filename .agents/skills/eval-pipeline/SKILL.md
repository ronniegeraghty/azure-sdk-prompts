---
name: "eval-pipeline"
description: "How the eval engine works: generate → grade → report"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/eval/engine.go, architecture.md"
---

## Context

The evaluation pipeline is the core workflow of hyoka. It orchestrates AI agent code generation, multi-model grading, and report generation. Understanding the pipeline is essential for debugging eval failures and extending the engine.

## Pipeline Phases

### 1. Generation Phase
- **Input:** Prompt (with service, language, plane metadata)
- **Process:** 
  - Creates a workspace (temporary directory)
  - Launches Copilot SDK session with generator model
  - Injects skills and MCP servers from config
  - Runs max 25 turns (configurable via `--max-session-actions`)
  - Captures all tool calls, file reads/writes, and stdout
- **Output:** Generated code + action timeline

### 2. Build Verification Phase (optional)
- **Input:** Generated code
- **Process:**
  - Runs language-specific build command (e.g., `cargo build` for Rust)
  - Collects stderr/stdout for report
  - Marks success/failure
- **Output:** Build status + logs

### 3. Grading Phase
- **Input:** Generated code + action timeline + config
- **Process:**
  - Runs 6 pluggable grader types independently
  - Each grader inspects code/timeline and scores against criteria
  - Collects all grader results
- **Output:** Grader scores (pass/fail + numeric score per grader)

### 4. Review Panel Phase
- **Input:** Generated code + build status
- **Process:**
  - Launches N reviewer sessions (one per reviewer model in config)
  - Sends code + rubric to each reviewer independently
  - Each reviewer scores against semantic criteria (e.g., "Handles errors correctly")
  - Consolidator model merges individual reviews → consensus score
- **Output:** Panel scores + individual reviewer insights

### 5. Report Generation Phase
- **Input:** All phase outputs + metadata
- **Process:**
  - Renders JSON report with full timeline
  - Renders HTML/Markdown from JSON
  - Saves to `reports/{run_id}/`
- **Output:** Multi-format reports

## Key Patterns

**Workspace Isolation:**
- Each eval run gets its own temp directory (`/tmp/hyoka-{run_id}`)
- Ensures no state leakage between runs

**Action Timeline Capture:**
- Every tool call, file operation, bash command is recorded with duration
- Used by graders to verify required/forbidden tool usage
- Enables post-hoc analysis of agent behavior

**Async/Parallel Processing:**
- Multiple prompts can run in parallel (via `--workers` flag)
- Generator and review phases are sequential per prompt (generation must complete before review)

**Error Handling:**
- Generation timeout → captured as error in report, review skipped
- Build failure → reported but review proceeds
- Grader timeout → grader marked as timeout, engine continues
- Single reviewer timeout → other reviewers still score

## Code Locations

- **Engine orchestration:** `hyoka/internal/eval/engine.go`
- **Copilot session management:** `hyoka/internal/eval/copilot.go`
- **Action timeline:** `hyoka/internal/eval/action.go`
- **Workspace creation:** `hyoka/internal/eval/workspace.go`
- **Process tracking:** `hyoka/internal/eval/proctracker.go`

## Anti-Patterns

- Don't call `log.Fatal` in pipeline steps — return errors
- Don't skip phases silently; report skipped phases in output
- Don't assume workspace directory persists after eval (it's cleaned up)
- Don't hardcode turn/file/output limits in engine (use SessionLimits)

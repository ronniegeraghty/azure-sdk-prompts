# Guardrail Enforcement Bug — maxTurns/maxFiles Stale Runner State

**Date:** 2026-04-23  
**Investigator:** Morpheus  
**Status:** ✅ CONFIRMED  

---

## Bug Summary

**Confirmed:** YES

The `CopilotPromptRunner` is constructed once at CLI startup with CLI flag defaults, **before** any per-config or per-prompt limits are loaded. Its `maxTurns` and `maxFiles` fields remain stale throughout execution, causing real-time enforcement to use wrong values.

**Impact:** A config with `max_turns: 100` will still kill sessions at 25 turns because the runner's `e.maxTurns` was set to `0` (CLI default) at construction, which falls back to `25` in the enforcement code.

---

## Root Cause Analysis

### Construction Timeline

1. **CLI startup** (`hyoka/cmd/run.go:356-361`):
   ```go
   sdkEval := eval.NewCopilotPromptRunner(eval.PromptRunnerOptions{
       AllowCloud:        f.allowCloud,
       MaxSessionActions: f.maxSessionActions,  // CLI flag: default 50
       MaxTurns:          f.maxTurns,           // CLI flag: default 0
       MaxFiles:          f.maxFiles,           // CLI flag: default 50
   })
   ```
   - The runner is created **once** with CLI defaults
   - `f.maxTurns` is `0` (default from flag line 101)
   - The runner stores these as `e.maxTurns = 0`, `e.maxFiles = 50`, `e.maxSessionActions = 50`

2. **Per-eval config resolution** (`hyoka/internal/eval/engine_eval.go:73-76`):
   ```go
   lim := e.resolveLimits(task.Config, task.Prompt)
   evalReport.GuardrailMaxTurns = lim.maxTurns        // ✅ CORRECT: resolved to 100
   evalReport.GuardrailMaxFiles = lim.maxFiles        // ✅ CORRECT
   evalReport.GuardrailMaxSessionActions = lim.maxSessionActions  // ✅ CORRECT
   ```
   - The Engine correctly merges CLI flags → config YAML → prompt frontmatter
   - Resolution order: `prompt > config > CLI flag/engine default`
   - The **report** gets the correct resolved value

3. **Real-time enforcement** (`hyoka/internal/eval/copilot.go:223-231`):
   ```go
   maxTurnsLimit := e.maxTurns   // ❌ STALE: still 0 from construction
   if maxTurnsLimit <= 0 {
       maxTurnsLimit = 25        // Falls back to hardcoded default
   }
   maxFilesLimit := e.maxFiles   // ✅ OK if CLI flag matches needs
   if maxFilesLimit <= 0 {
       maxFilesLimit = 50
   }
   ```
   - The OnEvent callback (lines 303-308) uses `maxTurnsLimit` for real-time cancellation
   - This value is **never updated** with the resolved config/prompt limits
   - The session is killed at turn 25, not turn 100

### Post-Hoc Guardrail Check

After generation completes, `engine_eval.go` does a post-hoc check using the **correct** resolved limits. But by then, the session may have already been killed by the stale real-time enforcement.

---

## Bug Class Extension

### Does this affect `maxSessionActions`?

**NO** — `maxSessionActions` is safe because:
1. It has a CLI flag (`--max-session-actions`, default `50`)
2. Per-config overrides via `config.Limits.MaxSessionActions` **do** exist in the schema
3. Real-time enforcement at `copilot.go:316-320` uses `e.maxSessionActions` directly
4. **BUT:** If the config sets a *higher* limit than the CLI default (e.g., config says `250`), the CLI default `50` will win during real-time enforcement

**Conditions for bug:**
- Config or prompt sets `max_session_actions > CLI flag value`
- Real-time enforcement will use the CLI value, killing the session early
- The report will show the resolved higher value, but the session already died

### Does this affect `maxFiles`?

**YES** — Same bug pattern:
1. CLI flag `--max-files` defaults to `50`
2. Per-config overrides via `config.Limits.MaxFiles` exist
3. Real-time enforcement at `copilot.go:339-344` uses `maxFilesLimit` (derived from `e.maxFiles`)
4. If config sets `max_files: 200`, real-time enforcement still uses CLI default `50`

**Conditions for bug:**
- Config sets `max_files != CLI flag value` (either higher or lower)
- Real-time enforcement ignores the config value

---

## Test Evidence

The test suite **validates the post-hoc guardrail check**, not the real-time enforcement:

- `engine_test.go:693` (`TestConfigLimitsRespectedByGuardrail`) confirms the **report** has correct resolved limits
- It uses a `manyTurnsRunner` stub that emits events, but doesn't test whether the runner's `e.maxTurns` was updated
- The stub runner doesn't have the real-time cancellation logic that depends on stale `e.maxTurns`

**Gap:** No test verifies that real-time `genCancel()` fires at the **resolved** limit, not the CLI default.

---

## Recommended Fix

**Option A: Pass resolved limits to runner per-eval** ✅ RECOMMENDED

**Approach:**
- Add a method to `CopilotPromptRunner`: `SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)`
- Call it from `engine_eval.go` right before `e.evaluator.Run()` (after `resolveLimits`)
- The runner stores these in per-eval fields, overriding the CLI-level defaults
- Real-time enforcement uses the per-eval fields if set, otherwise falls back to CLI defaults

**Why this is best:**
- Minimal structural change — runner already exists, just needs state update
- Per-eval granularity is the right abstraction (one runner, many evals)
- Backward compatible — if limits aren't set, CLI defaults still work
- Testable — stub runner can verify the method was called

**Pseudo-code:**
```go
// In copilot.go
func (e *CopilotPromptRunner) SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int) {
    e.evalMaxTurns = maxTurns
    e.evalMaxFiles = maxFiles
    e.evalMaxSessionActions = maxSessionActions
}

// In Run() method, use evalMaxTurns instead of e.maxTurns
maxTurnsLimit := e.evalMaxTurns
if maxTurnsLimit <= 0 {
    maxTurnsLimit = e.maxTurns  // Fall back to CLI default
}
if maxTurnsLimit <= 0 {
    maxTurnsLimit = 25  // Hard default
}
```

```go
// In engine_eval.go, before calling evaluator.Run()
lim := e.resolveLimits(task.Config, task.Prompt)
if copilotRunner, ok := e.evaluator.(*eval.CopilotPromptRunner); ok {
    copilotRunner.SetLimitsForEval(lim.maxTurns, lim.maxFiles, lim.maxSessionActions)
}
result, err := e.evaluator.Run(genCtx, task.Prompt, &task.Config, genDir)
```

---

### Option B: Move runner construction inside per-eval loop

**Approach:**
- Construct a new `CopilotPromptRunner` for each eval task
- Pass resolved limits directly to constructor

**Why not:**
- Breaks resource reuse (each runner creates a new SDK client)
- Higher overhead per-eval
- Doesn't match current architecture (runner is long-lived, engine fans out)

---

### Option C: Thread limits through a per-eval context

**Approach:**
- Add `EvalContext` struct with resolved limits
- Pass it through `evaluator.Run()`
- OnEvent callback reads from context

**Why not:**
- Requires changing the `PromptRunner` interface signature
- Breaks all stub runners in tests
- More invasive than Option A

---

## Smoke Test Plan

**Cheapest live eval to verify fix:**

```bash
# Create a test config with max_turns: 100
cat > configs/test-high-turns.yaml <<EOF
name: test-high-turns
generator:
  model: claude-opus-4.6
  tools: []
limits:
  max_turns: 100
EOF

# Run with a Python prompt (fastest)
hyoka run --prompt-id identity-dp-python-default-credential \
  --config test-high-turns \
  --log-level debug --log-file verify-turns.log

# Verify in log that turn 25 does NOT trigger cancellation:
grep "Turn limit reached" verify-turns.log  # Should be empty if prompt finishes <25 turns

# Or artificially trigger it with a prompt that loops:
# Check that it cancels at turn 100, not turn 25
```

**Expected behavior after fix:**
- Logs show "Turn limit reached" at turn 100 (if triggered)
- Report `guardrail_max_turns: 100` matches real-time enforcement
- No premature cancellation at turn 25

**Regression check:**
- CLI flag `--max-turns 10` should still override config defaults
- Prompt-level `max_turns` in frontmatter should override config

---

## Related Files

- `hyoka/cmd/run.go:356-361` — Runner construction with CLI defaults
- `hyoka/internal/eval/copilot.go:27-35` — Runner struct with stale fields
- `hyoka/internal/eval/copilot.go:223-231` — Real-time limit initialization
- `hyoka/internal/eval/copilot.go:303-308` — Turn limit enforcement
- `hyoka/internal/eval/copilot.go:339-344` — File limit enforcement
- `hyoka/internal/eval/copilot.go:316-320` — Action limit enforcement
- `hyoka/internal/eval/engine.go:394-421` — `resolveLimits()` function
- `hyoka/internal/eval/engine_eval.go:73-76` — Resolved limits written to report
- `hyoka/internal/eval/engine_test.go:693-705` — Test for post-hoc guardrail (not real-time)

---

## Impact Assessment

**Severity:** HIGH

- Real-time enforcement is critical for cost control and runaway protection
- Users setting config-level limits expect them to be honored
- The bug creates a trust gap: "I set `max_turns: 100` but it killed at 25"

**Affected scenarios:**
1. Any config with `max_turns` > CLI default (or with CLI flag = 0)
2. Any config with `max_files` != CLI default
3. Any config with `max_session_actions` > CLI default

**Current workaround:**
- Set CLI flags to match config maximums: `--max-turns 100 --max-files 200`
- But this breaks per-config flexibility — all configs in the run get the same limits

---

## Next Steps

1. **Implement Option A** — Neo or Tank (backend guardrail logic)
2. **Add real-time enforcement test** — Switch (verify genCancel fires at resolved limit)
3. **Update documentation** — Oracle (note the fix in configuration.md)
4. **Run smoke test** — Morpheus or Neo (live verification after merge)

---

**Filed by:** Morpheus  
**Date:** 2026-04-23  
**Priority:** High  
**Component:** Guardrails / Copilot Runner  

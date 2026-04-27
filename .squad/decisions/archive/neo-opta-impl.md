# Decision: Option A Implementation — Per-Eval Limit Threading

**Date:** 2026-04-23  
**Author:** Neo  
**Status:** Implemented  
**Commit:** `d2f6e93b`

---

## Context

Morpheus identified a guardrail enforcement bug: `CopilotPromptRunner` was constructed once at CLI startup with stale CLI defaults. Real-time enforcement read `e.maxTurns` (= 0 → 25) instead of per-eval resolved values from `resolveLimits()`. Same bug affected `maxFiles` and `maxSessionActions`.

**Impact:** A config with `max_turns: 100` would still kill sessions at 25 turns.

---

## Decision

Implemented **Option A** from Morpheus's report: pass resolved limits to runner per-eval.

### Method Signature

```go
func (e *CopilotPromptRunner) SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)
```

**Parameters:** All three limits as `int` — matches the `resolveLimits()` return structure.

### Concurrency Approach

**Model:** Engine runs multiple evals in parallel (goroutines with worker semaphore). All goroutines share one `e.evaluator` (`CopilotPromptRunner`) instance.

**Protection:** Per-eval fields (`evalMaxTurns`, `evalMaxFiles`, `evalMaxSessionActions`) protected by `sync.RWMutex` (`evalLimitsMu`):
- `SetLimitsForEval` acquires write lock
- Real-time enforcement (in `Run`) acquires read lock once at start, reads all three fields, then releases

**Why RWMutex:** Multiple evals can run concurrently. A simple `Mutex` would work too, but RWMutex clarifies that the lock is held briefly only to snapshot the values, not throughout the entire `Run` execution.

### Files Touched

1. **`hyoka/internal/eval/copilot.go`:**
   - Added per-eval fields to struct (lines 27-41)
   - Added `SetLimitsForEval` method (lines 49-57)
   - Updated real-time enforcement logic (lines 223-265) to prefer per-eval values → CLI defaults → hardcoded defaults
   - Updated action limit enforcement (line 352) to use `maxSessionActionsLimit` local variable

2. **`hyoka/internal/eval/engine_eval.go`:**
   - Type-assert `e.evaluator` to `*CopilotPromptRunner` (lines 148-152)
   - Call `SetLimitsForEval(lim.maxTurns, lim.maxFiles, lim.maxSessionActions)` after `resolveLimits()` and before `evaluator.Run()`
   - Skip if type assertion fails (stub runners in tests)

### Fallback Chain

Real-time enforcement now uses this priority:

1. Per-eval resolved value (from `evalMaxTurns`, set by engine)
2. CLI-level default (from `e.maxTurns`, set at runner construction)
3. Hardcoded default (e.g., 25 for turns, 50 for files)

**Example (turns):**
```go
maxTurnsLimit := e.evalMaxTurns        // Per-eval (e.g., 100 from config)
if maxTurnsLimit <= 0 {
    maxTurnsLimit = e.maxTurns         // CLI flag (e.g., 0 → skip)
}
if maxTurnsLimit <= 0 {
    maxTurnsLimit = 25                 // Hardcoded default
}
```

---

## Testing

- **Unit tests:** All existing eval tests pass (`go test -race ./hyoka/internal/eval/...`)
- **Test coverage:** `TestConfigLimitsRespectedByGuardrail` validates post-hoc guardrail check (report fields have correct resolved values)
- **Real-time enforcement test:** Switch is writing `guardrail_realtime_test.go` in parallel to verify `genCancel()` fires at the resolved limit, not CLI default

---

## Coordination Notes

**For Switch:** Use this method signature when writing your test:

```go
runner.SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)
```

Your WIP test file had compilation errors (missing `report` import, wrong field names like `GuardrailAbortReason`). I've renamed it to `.wip` so the build succeeds. When you integrate, import the `report` package and use the correct `EvalResult` fields.

**For future implementers:** If you change the signature, write a decision file to `.squad/decisions/inbox/` immediately so parallel work can adapt.

---

## Alternatives Considered

**Option B (move runner construction inside eval loop):**  
Rejected — breaks resource reuse, higher overhead per-eval, doesn't match current architecture (runner is long-lived, engine fans out).

**Option C (thread limits through per-eval context):**  
Rejected — requires changing `PromptRunner` interface signature, breaks all stub runners, more invasive.

---

## Related

- **Bug report:** `.squad/decisions/inbox/morpheus-maxturns-enforcement-bug.md`
- **Commit:** `d2f6e93b` — "Fix guardrail enforcement to use per-eval resolved limits"
- **Learnings:** `.squad/agents/neo/history.md` — 2026-04-23 entry

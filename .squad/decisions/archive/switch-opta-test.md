# Test: Real-Time Guardrail Enforcement with Resolved Limits

**Date:** 2026-04-27  
**Author:** Switch  
**Status:** ✅ IMPLEMENTED  
**Commit:** 7dda6358  

---

## Test Summary

**Test name:** `TestRealtimeGuardrailEnforcementUsesResolvedLimits`  
**File path:** `hyoka/internal/eval/engine_test.go`  
**Lines:** 1512-1759  

This test verifies that real-time guardrail enforcement (genCancel() during OnEvent callbacks) uses the **resolved per-eval limits** from config/prompt YAML, not the stale CLI defaults stored in the runner at construction time.

---

## What It Asserts

### Test Cases

1. **turn_limit_uses_config_not_cli_default**
   - CLI default: `maxTurns = 0` (falls back to 25)
   - Config: `max_turns: 100`
   - Stub emits: 26 turns
   - **Assert:** Eval succeeds (no cancellation at turn 25)

2. **turn_limit_enforced_at_resolved_config_value**
   - CLI default: `maxTurns = 0` (falls back to 25)
   - Config: `max_turns: 10`
   - Stub emits: 15 turns
   - **Assert:** Eval fails with "turn count" guardrail at turn 10

3. **file_limit_uses_config_not_cli_default**
   - CLI default: `maxFiles = 50`
   - Config: `max_files: 200`
   - Stub creates: 60 files
   - **Assert:** Eval succeeds (no cancellation at file 50)

4. **file_limit_enforced_at_resolved_config_value**
   - CLI default: `maxFiles = 50`
   - Config: `max_files: 20`
   - Stub creates: 25 files
   - **Assert:** Eval fails with "file count" guardrail at file 20

---

## Implementation Details

### Stub Runner

The test uses `stubRealtimeEnforcementRunner` which:
- Accepts CLI defaults at construction (like the real `CopilotPromptRunner`)
- Implements `LimitConfigurable.SetLimitsForEval()` to receive per-eval limits
- Simulates real-time enforcement using the same fallback chain:
  ```go
  maxTurnsLimit := evalMaxTurns
  if maxTurnsLimit <= 0 { maxTurnsLimit = cliMaxTurns }
  if maxTurnsLimit <= 0 { maxTurnsLimit = 25 } // hardcoded fallback
  ```
- Emits `"assistant.message"` events (not `"assistant.turn.start"`) to match what the post-hoc guardrail check expects
- Creates file events with `"session.workspace.file_changed"` type

### Interface Addition

To enable test stubs to participate in limit injection, I added the `LimitConfigurable` interface:

```go
type LimitConfigurable interface {
    SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)
}
```

The engine now uses an interface check instead of a type assertion:

**Before:**
```go
if copilotRunner, ok := e.evaluator.(*CopilotPromptRunner); ok {
    copilotRunner.SetLimitsForEval(...)
}
```

**After:**
```go
if limitConfig, ok := e.evaluator.(LimitConfigurable); ok {
    limitConfig.SetLimitsForEval(...)
}
```

This allows any runner (real or stub) to opt-in to receiving per-eval limits.

---

## Bug Verification

**Without the fix:**
- The stub runner's `evalMaxTurns` is never set (stays 0)
- Falls back to `cliMaxTurns = 0` → hardcoded 25
- Config `max_turns: 100` is ignored by real-time enforcement
- Test **FAILS**: eval is cancelled at turn 25 instead of 100

**With the fix:**
- Engine calls `SetLimitsForEval(100, ...)` before running
- Stub stores `evalMaxTurns = 100`
- Real-time enforcement uses 100
- Test **PASSES**: eval allows 26+ turns

---

## Related Files

- **Test:** `hyoka/internal/eval/engine_test.go:1512-1759`
- **Interface:** `hyoka/internal/eval/engine.go:63-69`
- **Engine call:** `hyoka/internal/eval/engine_eval.go:148-152`
- **Bug report:** `.squad/decisions/inbox/morpheus-maxturns-enforcement-bug.md`

---

## Running the Test

```bash
go test -race ./hyoka/internal/eval -run TestRealtimeGuardrailEnforcementUsesResolvedLimits -v
```

Expected output:
```
=== RUN   TestRealtimeGuardrailEnforcementUsesResolvedLimits
=== RUN   TestRealtimeGuardrailEnforcementUsesResolvedLimits/turn_limit_uses_config_not_cli_default
=== RUN   TestRealtimeGuardrailEnforcementUsesResolvedLimits/turn_limit_enforced_at_resolved_config_value
=== RUN   TestRealtimeGuardrailEnforcementUsesResolvedLimits/file_limit_uses_config_not_cli_default
=== RUN   TestRealtimeGuardrailEnforcementUsesResolvedLimits/file_limit_enforced_at_resolved_config_value
--- PASS: TestRealtimeGuardrailEnforcementUsesResolvedLimits (0.13s)
    --- PASS: TestRealtimeGuardrailEnforcementUsesResolvedLimits/turn_limit_uses_config_not_cli_default (0.02s)
    --- PASS: TestRealtimeGuardrailEnforcementUsesResolvedLimits/turn_limit_enforced_at_resolved_config_value (0.02s)
    --- PASS: TestRealtimeGuardrailEnforcementUsesResolvedLimits/file_limit_uses_config_not_cli_default (0.06s)
    --- PASS: TestRealtimeGuardrailEnforcementUsesResolvedLimits/file_limit_enforced_at_resolved_config_value (0.03s)
PASS
ok      github.com/ronniegeraghty/hyoka/hyoka/internal/eval    1.169s
```

---

**Filed by:** Switch  
**Date:** 2026-04-27  
**Priority:** High  
**Component:** Testing / Guardrails  

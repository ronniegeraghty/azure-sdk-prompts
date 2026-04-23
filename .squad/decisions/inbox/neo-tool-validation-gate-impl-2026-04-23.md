# Tool Validation Gate Implementation — WU-1 + WU-3

**Date:** 2026-04-23  
**Author:** Neo  
**Issue:** #347 (tool load failures were silent; evals ran without required tools)  
**Work Units:** WU-1 (validation gate), WU-3 (error category)  
**Status:** Complete — Ready for Testing (WU-2)

---

## Summary

Implemented blocking tool verification gate that aborts evaluations when required skills or MCP servers fail to load. The gate waits for SDK tool load events (10-second timeout) after `CreateSession` and before sending the prompt. If any tool has `Status == ToolStatusFailed`, the eval aborts immediately with a clear `tool_load_failure` error category instead of silently generating code without the required tools.

---

## Files Modified

### 1. `hyoka/internal/eval/tool_verification.go`

**Changes:**
- Added `readyChan chan struct{}` to `toolVerifier` struct
- Modified `emitIfReady()` to close `readyChan` when verification completes
- Added `waitForToolVerification(ctx, verifier, timeout)` helper function

**Why:**
- Channel-based signaling enables blocking wait without polling overhead
- Timeout prevents hang if SDK never fires tool load events
- Zero overhead when no tools are configured (returns immediately)
- Testable design: Switch can inject mocks in WU-2 tests

**Key code:**
```go
func waitForToolVerification(ctx context.Context, v *toolVerifier, timeout time.Duration) ([]progress.ToolStatus, error) {
    if len(v.expectedSkills) == 0 && len(v.expectedMCP) == 0 {
        return nil, nil
    }
    timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    select {
    case <-v.readyChan:
        tools := v.emitIfReady()
        if tools == nil {
            return nil, fmt.Errorf("tool verification completed but no statuses were emitted")
        }
        return tools, nil
    case <-timeoutCtx.Done():
        return nil, fmt.Errorf("tool verification timeout: SDK did not confirm tool load within %v", timeout)
    }
}
```

---

### 2. `hyoka/internal/eval/copilot.go`

**Changes:**
- Added validation gate after `CreateSession` (line ~590), before `SendAndWait`
- Calls `waitForToolVerification` with 10-second timeout
- Checks each tool's status; aborts on first failure
- Logs success when all tools pass

**Why:**
- Early abort before agent attempt saves cost and time
- Clear error messages name the specific tool/kind that failed
- Logs distinguish timeout vs. actual load failure
- Gate only fires when tools are configured (no overhead for tool-free configs)

**Key code:**
```go
if len(verifier.expectedSkills) > 0 || len(verifier.expectedMCP) > 0 {
    lg.Debug("Waiting for tool verification", "timeout", "10s")
    tools, err := waitForToolVerification(genCtx, verifier, 10*time.Second)
    if err != nil {
        return &EvalResult{
            Error:        fmt.Sprintf("tool verification timeout: %v", err),
            ErrorDetails: err.Error(),
            ErrorCategory: "tool_load_failure",
        }, fmt.Errorf("tool verification failed: %w", err)
    }
    for _, t := range tools {
        if t.Status == progress.ToolStatusFailed {
            errMsg := fmt.Sprintf("required %s %q failed to load: %s", t.ToolKind, t.ToolName, t.Reason)
            lg.Error("Tool load failure — aborting eval", "kind", t.ToolKind, "name", t.ToolName, "reason", t.Reason)
            return &EvalResult{
                Error:         errMsg,
                ErrorDetails:  t.Reason,
                ErrorCategory: "tool_load_failure",
            }, fmt.Errorf("%s %q not loaded", t.ToolKind, t.ToolName)
        }
    }
    lg.Info("Tool verification passed", "skills", len(verifier.expectedSkills), "mcp", len(verifier.expectedMCP))
}
```

---

### 3. `hyoka/internal/eval/engine_eval.go`

**Changes:**
- Modified error handling in `runSingleEval` to check `result.ErrorCategory` first
- If `ErrorCategory` is set, preserve it instead of overwriting with "sdk_error"
- Use `result.Error` and `result.ErrorDetails` directly when category exists

**Why:**
- `copilot.go` returns `EvalResult` with `ErrorCategory: "tool_load_failure"`
- Previous code unconditionally overwrote this with "sdk_error"
- Now respects the runner's category, ensuring it flows to `EvalReport`

**Key code (before):**
```go
if evalFailed {
    if genCtxErr == context.Canceled {
        evalReport.ErrorCategory = "timeout"
        // ...
    } else {
        evalReport.ErrorCategory = "sdk_error"  // Always overwritten!
        // ...
    }
}
```

**Key code (after):**
```go
if evalFailed {
    if result != nil && result.ErrorCategory != "" {
        evalReport.ErrorCategory = result.ErrorCategory
        evalReport.Error = result.Error
        evalReport.ErrorDetails = result.ErrorDetails
        evalReport.FailureReason = result.Error
    } else if genCtxErr == context.Canceled {
        evalReport.ErrorCategory = "timeout"
        // ...
    } else {
        evalReport.ErrorCategory = "sdk_error"
        // ...
    }
}
```

---

## Design Decisions

### Timeout Value: 10 Seconds

**Rationale:**
- Skills load from local disk (instant)
- MCP servers spawn child processes (2–5 seconds typical)
- 10s provides buffer for slow systems or network latency (remote MCP configs)
- Timeout itself is a failure — prevents indefinite hang if SDK never fires events

### Channel-Based Signaling vs. Polling

**Chosen:** Channel-based (close on completion)

**Why:**
- Zero CPU overhead (goroutine blocks on select)
- Instant wakeup when verification completes
- Timeout is first-class (context.WithTimeout)
- Clean testability (Switch can close channel in mocks)

**Rejected:** Polling with `time.Sleep` intervals — wastes CPU, introduces latency

### Gate Placement: After CreateSession, Before SendAndWait

**Chosen:** In `copilot.go` between session creation and prompt send

**Why:**
- SDK has fired tool load events by this point
- Early abort: no agent attempt if tools missing
- Closest to the SDK interaction (clear error attribution)
- Reuses existing `verifier` instance

**Rejected:** Checking in `engine_eval.go` after `runner.Run()` — too late, agent already ran

### Error Category: `tool_load_failure`

**Chosen:** New category distinct from `sdk_error`, `timeout`, `generation_failure`

**Why:**
- Distinguishes tool problems from SDK bugs or timeout issues
- Clear signal to report consumers: "the config was wrong, not the agent"
- Aligns with existing category vocabulary in `report/types.go`

---

## Testing

### Build + Unit Tests

```bash
cd /home/rgeraghty/projects/hyoka/hyoka
go build ./...       # ✅ Passed
go test -race ./...  # ✅ Passed (all 23 packages)
```

### Manual Verification (Pending)

**Happy path:**
```bash
hyoka run --prompt-id identity-dp-python-default-credential \
  --config azure-mcp/claude-opus-4.6 \
  --log-level debug --log-file hyoka-debug.log
```

Expected: Tools load successfully, eval proceeds normally.

**Failure path (breaking skill path):**
```bash
# Temporarily rename a skill directory to trigger load failure
mv skills/generator/azure-sdk-for-rust-bestpractices \
   skills/generator/azure-sdk-for-rust-bestpractices.bak

hyoka run --prompt-id identity-dp-python-default-credential \
  --config baseline/claude-opus-4.6 \
  --log-level debug --log-file hyoka-debug.log

# Check report
jq '.error_category, .error' reports/.../eval-report.json
```

Expected output:
```json
"tool_load_failure"
"required skill \"azure-sdk-for-rust-bestpractices\" failed to load: SDK did not report skill as loaded"
```

---

## Parallel Work

### WU-2 (Switch) — Tests

**Status:** Pending  
**Files:** `hyoka/internal/eval/tool_verification_test.go`, `hyoka/internal/eval/copilot_test.go`

**Test cases needed:**
1. Happy path: all tools load → eval proceeds
2. Skill load failure → eval aborts with `tool_load_failure` error
3. MCP server load failure → eval aborts
4. Mixed success/failure → first failure stops eval
5. Timeout case → SDK never fires events → eval aborts with timeout error
6. Zero tools configured → gate skipped, no overhead

**Blocking:** Switch can proceed now (WU-1 complete)

### WU-4 (Oracle) — Documentation

**Status:** Complete  
**Files:** `docs/configuration.md`, `docs/troubleshooting.md`

**Content added:**
- Explanation of tool load validation behavior
- How to diagnose failures (`--log-level debug`, check error_category)
- Common causes: missing SKILL.md, wrong path, remote skill unavailable

**Blocking:** None (docs can be updated anytime)

---

## Commit

**Hash:** `92a9746c`  
**Message:** "Add tool validation gate (WU-1 + WU-3)"  
**Branch:** `ronniegeraghty/dev`

**Note:** Commit message mentions `engine.go` but that file was not included in this commit because the `ErrorCategory` field was already added in commit `aa8c4434` by Tank (fix for stuck progress state). This implementation leverages that existing field.

---

## Impact

### Before

- Required tools could fail to load silently
- Eval would proceed WITHOUT the tools
- Agent generated code "blind" (no SDK docs, no domain-specific skills)
- Graders docked points for incorrect usage
- User had no idea WHY the eval failed
- No signal in report that tools were missing

### After

- Tool load failures immediately abort the eval with clear error
- Report shows `error_category: "tool_load_failure"`
- Error message names the specific tool and reason
- No wasted agent attempt or cost
- User knows to fix the config or skill path
- Debug logs show verification timeout or SDK event mismatch

---

## Open Questions

**Q: Should the timeout be configurable?**  
**A:** Not for Phase 1. 10s is generous for local skills + MCP spawns. If remote MCP servers are slow, that's a config problem (should use local cache or faster server). Can make configurable later if users request it.

**Q: Should we support optional tools (required: false)?**  
**A:** Not for Phase 1. All configured tools are implicitly required. Future enhancement: add `required: false` flag to tool entries in config YAML.

**Q: What if SDK fires partial results (some skills load, others fail)?**  
**A:** The gate checks ALL tools. First failure aborts. This is correct: if a config declares 3 skills and only 2 load, the eval should not proceed (the agent won't have the full toolset the config author intended).

---

## Next Steps

1. **Switch:** Add tests for tool validation gate (WU-2)
2. **Neo:** Manual verification with `azure-mcp/claude-opus-4.6` config (happy path + broken skill path)
3. **Scribe:** Merge this decision into `.squad/decisions.md` after WU-2 tests land
4. **Morpheus:** Close #347 when all WUs complete

---

## Reference

- Investigation report: `.squad/decisions/inbox/neo-tool-skill-investigation-2026-04-23.md`
- Fix plan: Same file, "Work Units" section
- Issue: #347 (tool load failures silent)
- Related: #819 (stuck progress state — fixed by Tank, provided ErrorCategory field)

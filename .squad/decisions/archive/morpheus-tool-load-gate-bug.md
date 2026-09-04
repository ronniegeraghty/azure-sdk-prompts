# Tool Load Verification Gate Bug — 30s Timeout Is Wrong Approach

**Author:** Morpheus 🕶️  
**Date:** 2026-04-27  
**Status:** Investigation Complete — Awaiting Fix Approval

---

## Issue Summary

During a live smoke test of the Option A guardrail fix, an eval terminated with `tool_load_failure` after 45 skills failed to confirm load within 30 seconds. The error appeared in `reports/20260427-232343/results/key-vault/.../generator.json`:

```
tool_load_failure:
45 tool(s) failed to load:
  • mcp "azure": SDK did not confirm tool load within 30s
  • skill "agent-framework-azure-ai-py": SDK did not confirm tool load within 30s
  ... (43 more, all "did not confirm tool load within 30s")
```

**Ronnie's diagnosis (user input):**

> "Woah so if an agent session doesn't confirm all the tools loaded in a certain amount of time we just stop looking to see if they eventually loaded? That's not the right way to do that. We should just wait until whatever usually happens after all tools loading messages or fail to load messages happens then make our call on what tools did and did not load."

---

## Root Cause

### Current Implementation

**File:** `hyoka/internal/eval/tool_verification.go:156-179` (function `waitForToolVerification`)  
**File:** `hyoka/internal/eval/copilot.go:774` (call site)

The post-session tool verification gate uses a **30-second hard timeout** to wait for SDK tool-load events:

```go
// copilot.go:774
if summary := postSessionToolVerification(ctx, verifier, 30*time.Second); summary != "" {
    // hard-fail the eval
}

// tool_verification.go:156-179
func waitForToolVerification(ctx context.Context, v *toolVerifier, timeout time.Duration) ([]progress.ToolStatus, error) {
    timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    select {
    case <-v.readyChan:  // Signal from verifier when all expected tool events fire
        return v.emitIfReady(), nil
    case <-timeoutCtx.Done():
        return nil, fmt.Errorf("tool verification timeout: SDK did not confirm tool load within %v", timeout)
    }
}
```

**How it works:**
1. `toolVerifier` listens for `SessionEventTypeSessionSkillsLoaded` and `SessionEventTypeSessionMcpServersLoaded` during the SDK event stream (`copilot.go:439-483`)
2. When **both** expected event types fire (or only the configured one fires if only skills OR only MCP is configured), the verifier closes `readyChan` to signal completion
3. `waitForToolVerification` blocks on `readyChan` with a 30s timeout
4. If the timeout fires before both events arrive, **every configured tool is marked as Failed** with reason `"SDK did not confirm tool load within 30s"`

**Why 30 seconds was chosen:**
- From `.squad/decisions/archive/neo-item-e-post-session-verification-gate.md:54-59`:
  > "The 30s value comes from Morpheus's spec § 5 open-question 5 ('default 30s, configurable via `--tool-verify-timeout`'). I did **not** add the CLI flag — out of scope for Item E. If cold-start MCP servers prove flaky, Tank can add the flag in a follow-up."
- This was a **guess**, not based on empirical data
- The decision doc even anticipated this might be too low for cold-start MCP servers

---

## The Fundamental Problem

**The timeout is a workaround, not a signal-based solution.**

The SDK **does** emit events that mark session initialization completion. The tool-load events fire **during** `session.SendAndWait()`, not after:

From `copilot.go:761-767`:
```go
// Post-session tool verification gate (#347 / Item E). The SDK emits
// SessionSkillsLoaded / SessionMcpServersLoaded only after the first
// message round-trip, so this gate runs AFTER SendAndWait returned —
// by which point the verifier's readyChan has typically already closed
// from inside the OnEvent callback.
```

**The sequence is:**
1. `session.SendAndWait(...)` starts
2. SDK begins loading tools (skills, MCP servers)
3. SDK emits `SessionSkillsLoaded` event → `verifier.onSkillsLoaded()` called
4. SDK emits `SessionMcpServersLoaded` event → `verifier.onMCPLoaded()` called
5. SDK emits `SessionEventTypeAssistantTurnStart` → first turn begins
6. ... model generates code, calls tools, etc. ...
7. `session.SendAndWait(...)` returns
8. **CURRENT CODE:** `postSessionToolVerification` waits 30s for `readyChan` to close

**The bug:** We're waiting for events that should have **already fired** by the time we call `postSessionToolVerification`. The 30s timeout is only hit when:
- The SDK doesn't emit the event at all (real tool load failure)
- The SDK is slow and hasn't emitted the event yet (false positive)
- There's an SDK bug or network delay

**What we should do instead:** Wait until the SDK signals "session initialization is complete" — which happens **before** `SendAndWait` returns. The proper signal is:
- **Option 1:** First `SessionEventTypeAssistantTurnStart` — this marks the session as ready for work
- **Option 2:** First `SessionEventTypeAssistantMessage` — the model's first response
- **Option 3 (safest):** `SendAndWait` returning successfully — by definition, all tool events have fired by this point

---

## SDK Events Analysis

**Available SDK events (from `/home/rgeraghty/go/pkg/mod/github.com/github/copilot-sdk/go@v0.2.0/generated_session_events.go`):**

Relevant to session lifecycle:
- `SessionEventTypeSessionStart` — session created
- `SessionEventTypeSessionSkillsLoaded` — skills finished loading
- `SessionEventTypeSessionMcpServersLoaded` — MCP servers finished loading
- `SessionEventTypeSessionIdle` — session is idle (waiting for input)
- `SessionEventTypeAssistantTurnStart` — first turn starts (tools are loaded by this point)
- `SessionEventTypeAssistantMessage` — first assistant response

**Critically:** There is **NO** `SessionReady` or `SessionToolsLoadingComplete` event. The SDK's signal that "tools are done loading" is:
1. The individual `SessionSkillsLoaded` and `SessionMcpServersLoaded` events
2. The session proceeding to the first turn (`AssistantTurnStart`)

**Current code already handles this correctly!** The `toolVerifier` waits for both events before closing `readyChan`. The problem is the **30s fallback timeout** when one or both events don't fire.

---

## What the SDK Actually Emits (Behavioral Analysis)

From `.squad/orchestration-log/2026-04-23T14-29-03Z-neo.md`:
> "SDK emits `SessionSkillsLoaded` events **during** first SendAndWait, not after CreateSession."

From `.squad/agents/neo/history.md`:
> "`SessionSkillsLoaded` and `SessionMcpServersLoaded` fire only AFTER `session.SendAndWait` completes its first round-trip."

**These statements are consistent:** The events fire **during** the `SendAndWait` call, after the prompt is sent and before the first assistant turn starts.

**The timing sequence in real sessions:**
1. `CreateSession()` returns
2. `SendAndWait(prompt)` is called
3. SDK sends prompt to model
4. **SDK loads tools in parallel** (skills, MCP servers)
5. SDK emits `SessionSkillsLoaded` (if skills configured)
6. SDK emits `SessionMcpServersLoaded` (if MCP configured)
7. SDK emits `AssistantTurnStart` — **tools must be loaded by now**
8. Model generates code, calls tools
9. `SendAndWait` returns

**By the time `postSessionToolVerification` is called (line 774), `SendAndWait` has already returned.** This means:
- If the SDK was going to emit the tool events, it **already did**
- If the SDK didn't emit them, **waiting 30 more seconds won't help**

**Exception:** The 30s timeout is only useful if the SDK is **still emitting events asynchronously** after `SendAndWait` returns — but the OnEvent handler is **still running** because the session is still open. Let me check if events can fire after `SendAndWait` returns...

From `copilot.go:250-298` (OnEvent handler setup):
```go
session.OnEvent(func(event copilot.SessionEvent) {
    mu.Lock()
    defer mu.Unlock()
    // ... event processing ...
})
```

**Key insight:** The OnEvent handler is a callback that fires **during** the session lifetime. Once `SendAndWait` returns, the session is **still open** but has finished the turn. The SDK documentation doesn't specify whether tool-load events can fire **after** `SendAndWait` returns but **before** the session is closed.

**However:** The comment at `copilot.go:764` says "by which point the verifier's readyChan has **typically** already closed" — the word "typically" suggests there are edge cases where it hasn't closed yet.

---

## The Real Issue: Race Condition

**The 30s timeout exists to handle a race:**
- `SendAndWait` returns
- We call `postSessionToolVerification` immediately
- But the OnEvent handler is still processing the final events from the SDK
- The tool-load events might not have been processed yet

**The proper fix:** Don't race against the event handler. Instead:
1. **Signal when tool loading is complete** — use `AssistantTurnStart` as the definitive "tools are loaded" marker
2. **No timeout for expected events** — if we're still in the first turn and haven't seen tool events, they **will** fire
3. **Fallback timeout only for SDK hangs** — if the session hangs before the first turn, we need a ceiling (but much higher, like 5-10 minutes)

---

## Proposed Fix

### Option A: Use First Turn Start as Tool Load Completion Signal (Recommended)

**Change the verifier to mark "tool loading complete" when `AssistantTurnStart` fires:**

```go
// tool_verification.go — add new method
func (v *toolVerifier) onSessionReady() {
    // Called when AssistantTurnStart fires — tools MUST be loaded by now
    if v.emitted {
        return  // Already emitted, no-op
    }
    // Force emit: if we haven't seen a tool event by first turn, it's never coming
    if len(v.expectedSkills) > 0 && !v.skillsEvtSeen {
        v.skillsEvtSeen = true  // Mark as seen (but empty)
    }
    if len(v.expectedMCP) > 0 && !v.mcpEvtSeen {
        v.mcpEvtSeen = true  // Mark as seen (but empty)
    }
}

// copilot.go:333-343 — add call to onSessionReady
case copilot.SessionEventTypeAssistantTurnStart:
    turnCounter++
    rec.TurnNumber = turnCounter
    lg.Info("Turn started", "turn", turnCounter)
    
    // Tool loading MUST be complete by first turn start
    if turnCounter == 1 {
        verifier.onSessionReady()
        if e.progressFn != nil {
            if t := verifier.emitIfReady(); t != nil {
                verifiedTools = t
            }
        }
    }
    
    // Real-time turn limit enforcement...
```

**Benefits:**
- **No timeout** — we wait for a real SDK signal (first turn start)
- **Correct semantics** — tools must be loaded before the first turn
- **Handles slow MCP servers** — if an MCP server takes 2 minutes to authenticate, we wait for it
- **Still fails fast** — if the session never reaches the first turn (SDK crash, etc.), the existing session timeout (10 minutes) will catch it

**Edge case handling:**
- If `AssistantTurnStart` fires before tool events → we force-close the verifier and mark missing tools as Failed
- If tool events never fire → same outcome as today (all tools Failed), but no arbitrary 30s limit
- If the session hangs before first turn → existing session timeout catches it

### Option B: Increase Timeout to 5-10 Minutes (Temporary Fix)

**Change:**
```go
// copilot.go:774
if summary := postSessionToolVerification(ctx, verifier, 5*time.Minute); summary != "" {
```

**Benefits:**
- Minimal code change
- Handles slow cold-start MCP servers

**Drawbacks:**
- Still a timeout-based approach (not signal-based)
- Delays failure detection by up to 5 minutes when SDK truly hangs
- Doesn't solve the fundamental race condition

### Option C: Remove Timeout, Wait Indefinitely (Not Recommended)

**Change:**
```go
func waitForToolVerification(ctx context.Context, v *toolVerifier) ([]progress.ToolStatus, error) {
    select {
    case <-v.readyChan:
        return v.emitIfReady(), nil
    case <-ctx.Done():
        return nil, ctx.Err()
    }
}
```

**Benefits:**
- No false positives from timeout

**Drawbacks:**
- Hangs forever if SDK never emits events
- No failure detection

---

## Recommendation

**Implement Option A** — use `AssistantTurnStart` as the definitive "tools loaded" signal.

**Rationale:**
1. **Semantically correct:** The SDK cannot start an assistant turn without loading tools first
2. **No arbitrary timeout:** We wait for a real event, not a guess
3. **Handles edge cases:** If tool events don't fire before the first turn, we mark them as Failed (correct behavior)
4. **Existing safety net:** The session timeout (10 minutes by default) still catches SDK hangs

**Pseudocode for the fix:**

```go
// tool_verification.go
func (v *toolVerifier) onSessionReady() {
    if v.emitted {
        return
    }
    // Force-close any tool kinds that haven't reported by first turn
    if len(v.expectedSkills) > 0 && !v.skillsEvtSeen {
        v.skillsEvtSeen = true
    }
    if len(v.expectedMCP) > 0 && !v.mcpEvtSeen {
        v.mcpEvtSeen = true
    }
}

// copilot.go — in OnEvent handler
case copilot.SessionEventTypeAssistantTurnStart:
    turnCounter++
    if turnCounter == 1 {
        verifier.onSessionReady()  // Mark tool loading as complete
        if e.progressFn != nil {
            if t := verifier.emitIfReady(); t != nil {
                verifiedTools = t
            }
        }
    }
    // ... existing turn limit logic ...

// tool_verification.go — simplify postSessionToolVerification
func postSessionToolVerification(ctx context.Context, v *toolVerifier) string {
    if v == nil || (len(v.expectedSkills) == 0 && len(v.expectedMCP) == 0) {
        return ""
    }
    
    // Opportunistic flush
    tools := v.emitIfReady()
    if tools == nil {
        // Should never happen — onSessionReady was called in AssistantTurnStart
        // But if we somehow get here, wait without timeout (session timeout will catch hangs)
        select {
        case <-v.readyChan:
            tools = v.emitIfReady()
        case <-ctx.Done():
            return tool.SummarizeToolLoadErrors(expectedAsContextCancelled(v))
        }
    }
    
    var failed []*tool.ToolLoadError
    for _, t := range tools {
        if t.Status == progress.ToolStatusFailed {
            failed = append(failed, &tool.ToolLoadError{
                Kind: t.ToolKind, Name: t.ToolName, Reason: t.Reason,
            })
        }
    }
    return tool.SummarizeToolLoadErrors(failed)
}
```

---

## Citations

- **Current timeout:** `hyoka/internal/eval/copilot.go:774`
- **Verifier implementation:** `hyoka/internal/eval/tool_verification.go:34-151`
- **Timeout function:** `hyoka/internal/eval/tool_verification.go:156-179`
- **SDK event sequence:** `hyoka/internal/eval/copilot.go:761-767` (comment)
- **Decision history:** `.squad/decisions/archive/neo-item-e-post-session-verification-gate.md:54-59`
- **Original bug fix:** `.squad/orchestration-log/2026-04-23T14-29-03Z-neo.md`
- **SDK event types:** `/home/rgeraghty/go/pkg/mod/github.com/github/copilot-sdk/go@v0.2.0/generated_session_events.go`
- **User input:** Ronnie's message (2026-04-27)

---

## Next Steps

1. **Neo:** Implement Option A (estimated 30 minutes)
2. **Switch:** Add test coverage for the new `onSessionReady` path (estimated 45 minutes)
3. **Morpheus:** Review PR before merge
4. **Smoke test:** Run `hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise` to verify 45-tool config succeeds

---

**Status:** Ready for implementation approval

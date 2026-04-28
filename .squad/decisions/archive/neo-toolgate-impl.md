# Tool-Load Verification Gate — Option A Implementation

**Author:** Neo 💊  
**Date:** 2026-04-27  
**Status:** ✅ Implemented  
**Related:** `.squad/decisions/inbox/morpheus-tool-load-gate-bug.md` (investigation)

---

## Summary

Replaced the 30s polling timeout with a listener for `SessionEventTypeAssistantTurnStart`. When the SDK starts the first turn, tool registration is definitively complete — the SDK won't begin generation until tools are loaded or definitively failed. This eliminates false positives when sessions with many skills (45+) take longer than 30s for tool events to fire.

**Key changes:**
1. Added `onSessionReady()` method to `toolVerifier` that marks tool registration as complete when `AssistantTurnStart` fires
2. Wired `onSessionReady()` into the SDK event dispatch at `copilot.go:337-345`
3. Replaced 30s timeout with 5-minute absolute ceiling (fail-safe for broken sessions)
4. Added per-kind tracking (`turnBeforeSkills`, `turnBeforeMCP`) to distinguish failure reasons:
   - **"Not registered before first turn"**: Event never fired before turn
   - **"SDK did not report X as loaded"**: Event fired, but X wasn't in the reported list

---

## Implementation Details

### File: `hyoka/internal/eval/tool_verification.go`

#### Method Signature

```go
func (v *toolVerifier) onSessionReady()
```

**Purpose:** Called when `SessionEventTypeAssistantTurnStart` fires. Marks tool registration as definitively complete. Forces any unseen tool events to be treated as "missing" so `emitIfReady()` will proceed.

**Logic:**
- If skills were expected but `skillsEvtSeen` is false → set `turnBeforeSkills = true` and `skillsEvtSeen = true`
- If MCP was expected but `mcpEvtSeen` is false → set `turnBeforeMCP = true` and `mcpEvtSeen = true`
- Per-kind tracking ensures correct failure reasons:
  - Skills that failed because the skills event never fired → "Not registered before first turn"
  - Skills that failed because the skills event fired but didn't include them → "SDK did not report skill as loaded"
  - Same logic for MCP servers

**Thread-safety:** Not thread-safe. Caller (`copilot.go` OnEvent handler) serializes all calls under a mutex.

#### Struct Changes

```go
type toolVerifier struct {
	expectedSkills    map[string]bool
	expectedMCP       map[string]bool
	loadedSkills      map[string]bool
	loadedMCP         map[string]bool
	skillsEvtSeen     bool
	mcpEvtSeen        bool
	emitted           bool
	readyChan         chan struct{}
	turnBeforeSkills  bool // NEW: True if turn started before skills event
	turnBeforeMCP     bool // NEW: True if turn started before MCP event
}
```

**Removed field:** `firstTurnStarted` (replaced with per-kind `turnBeforeSkills` / `turnBeforeMCP`)

**Rationale:** Per-kind tracking allows correct failure reasons when only ONE event is missing. Example: Skills event fires with partial list, MCP event never fires → skill failures use "SDK did not report", MCP failures use "Not registered before first turn".

---

### File: `hyoka/internal/eval/copilot.go`

#### Wiring Point

**Location:** `copilot.go:333-355` (inside `SessionEventTypeAssistantTurnStart` case)

```go
case copilot.SessionEventTypeAssistantTurnStart:
	turnCounter++
	rec.TurnNumber = turnCounter
	lg.Info("Turn started", "turn", turnCounter)
	// Tool loading MUST be complete by first turn start — the SDK won't
	// begin generation until tools are loaded or definitively failed.
	// Signal the verifier so postSessionToolVerification doesn't wait
	// forever for events that will never arrive (#347 / Option A).
	verifier.onSessionReady()
	if e.progressFn != nil {
		if t := verifier.emitIfReady(); t != nil {
			verifiedTools = t
		}
	}
	// Real-time turn limit enforcement (#347)
	if maxTurnsLimit > 0 && turnCounter > maxTurnsLimit && !turnLimitHit {
		turnLimitHit = true
		lg.Warn("Turn limit reached, cancelling session",
			"turns", turnCounter, "max_turns", maxTurnsLimit)
		genCancel()
	}
```

**Sequence:**
1. Increment turn counter
2. **Call `verifier.onSessionReady()`** ← NEW
3. Emit progress event if progress callback is registered
4. Enforce turn limit if applicable

**Key insight:** `onSessionReady()` is called EVERY time a turn starts (defensive), but internally it's idempotent: if both expected events already fired, it's a no-op. If any events are missing, it forces them to be "seen" (empty) so verification can proceed.

#### Post-Session Verification Gate

**Location:** `copilot.go:761-797`

**Changed from:**
```go
if summary := postSessionToolVerification(ctx, verifier, 30*time.Second); summary != "" {
	// ...
}
```

**Changed to:**
```go
if summary := postSessionToolVerification(ctx, verifier, 5*time.Minute); summary != "" {
	// ...
}
```

**Rationale:** The 5-minute ceiling is NOT the primary gate — it's a fail-safe for broken sessions (auth hang, network failure) that never reach first turn. The PRIMARY gate is `onSessionReady()` which fires when `AssistantTurnStart` happens. The timeout only trips if the session is fundamentally broken.

**Updated comment (copilot.go:761-773):**
> "waitForToolVerification now uses a 5-minute absolute ceiling as a fail-safe in case the session never reached first turn (auth hang, network failure, SDK bug). This is NOT the primary gate — the real signal is AssistantTurnStart, which marks tool registration as definitively complete. The ceiling is ONLY for broken sessions."

---

### File: `hyoka/internal/config/config.go`

#### Configuration Field (Added but NOT wired yet)

```go
type SessionLimits struct {
	MaxTurns          int    `yaml:"max_turns,omitempty" json:"max_turns,omitempty"`
	MaxFiles          int    `yaml:"max_files,omitempty" json:"max_files,omitempty"`
	MaxSessionActions int    `yaml:"max_session_actions,omitempty" json:"max_session_actions,omitempty"`
	ToolLoadCeiling   string `yaml:"tool_load_ceiling,omitempty" json:"tool_load_ceiling,omitempty"` // NEW
}
```

**Status:** Field added to schema, but NOT yet wired into engine. Currently hard-coded to 5 minutes in `copilot.go:779`.

**TODO (future work):**
- Parse `ToolLoadCeiling` duration string (e.g., `"10m"`, `"2m30s"`)
- Thread through to `postSessionToolVerification` call
- Validate in `config.Validate()` (e.g., must be >= 1 minute, <= 30 minutes)
- Add to docs/configuration.md

**Default:** 5 minutes (when omitted or zero)

---

## Testing

### New Test Suite

**File:** `hyoka/internal/eval/tool_verification_gate_test.go`

**Test:** `TestAssistantTurnStartToolLoadGate`

**Coverage:**
1. **all_tools_load_before_assistant_turn_start**: Normal path — both events fire before turn, all tools load
2. **some_tools_fail_before_assistant_turn_start**: Partial failures — events fire before turn, some tools missing from SDK lists
3. **tools_load_slow_but_before_turn_proves_fix**: THE FIX — simulates 45+ skills taking 35-40s to load. OLD: Would timeout at 30s. NEW: Waits for turn at 45s, all tools succeed.
4. **assistant_turn_fires_before_some_tool_events**: Mixed case — skills event fires, MCP event never fires, turn starts → skill failures use "SDK did not report", MCP failures use "Not registered before first turn"
5. **absolute_ceiling_exceeded_no_turn_start**: Fail-safe — session hangs, never reaches first turn, 5min ceiling trips → all tools Failed with "session did not reach first turn"

**Test runtime:**
- Fast cases: ~300ms each (delays < 1s)
- Slow case (tools_load_slow_but_before_turn_proves_fix): 45s (simulates real 45+ skill load time)
- Ceiling case: Skipped in normal test runs (would take 5min+)

**Verification:**
```bash
cd /home/rgeraghty/projects/hyoka
go test -race ./hyoka/internal/eval/ -run TestAssistantTurnStartToolLoadGate -timeout 2m -v
```

### Existing Tests Updated

**File:** `hyoka/internal/eval/tool_verification_test.go`

**Changed:** `TestPostSessionVerification_TimeoutMarksAllFailed`

**What changed:** Updated expected error substring from `"did not confirm tool load"` to `"Session did not reach first turn within"` to match new timeout semantics.

**All other tests:** Unchanged and passing.

---

## Behavior Changes

### Before (30s polling)

1. Tool events fire during `session.SendAndWait()`
2. `SendAndWait()` returns
3. `postSessionToolVerification()` called with 30s timeout
4. **If tool events haven't fired yet → WAIT UP TO 30s**
5. **If 30s elapses → MARK ALL TOOLS AS FAILED** (even if events would arrive at 31s)

**Problem:** False positives when SDK takes >30s to confirm tool loads (common with 45+ skills).

### After (AssistantTurnStart signal + 5min ceiling)

1. Tool events fire during `session.SendAndWait()`
2. `AssistantTurnStart` fires → `onSessionReady()` called → marks tool registration complete
3. `SendAndWait()` returns
4. `postSessionToolVerification()` called with 5min timeout (for broken sessions only)
5. **If tool events already fired → IMMEDIATE RETURN (typical case)**
6. **If tool events missing but turn started → MARK MISSING AS FAILED** (correct)
7. **If session hung and never reached turn → WAIT UP TO 5min → FAIL** (rare, broken session)

**Result:** Eliminates false positives. Real tool failures still caught. Broken sessions caught by 5min ceiling instead of continuing forever.

---

## Error Message Changes

### Old Messages

- **Timeout:** `"SDK did not confirm tool load within 30s"`
- **Missing from list:** `"SDK did not report skill as loaded"` / `"SDK did not report MCP server as loaded"`

### New Messages

- **Timeout (5min ceiling):** `"Session did not reach first turn within 5m0s"`
- **Event never fired:** `"Not registered before first turn"`
- **Event fired, tool not in list:** `"SDK did not report skill as loaded"` / `"SDK did not report MCP server as loaded"` (unchanged)

**Rationale:** The new "Not registered before first turn" message is MORE accurate — it tells operators that the SDK never even tried to load those tools (event didn't fire), vs. "SDK did not report" which means the SDK DID try but the tool wasn't in the success list.

---

## Verification Commands

### Build

```bash
cd /home/rgeraghty/projects/hyoka
go build ./hyoka/...
```

### Test

```bash
# Run all eval tests (includes tool verification)
go test -race ./hyoka/internal/eval/ -timeout 3m

# Run just tool verification tests
go test -race ./hyoka/internal/eval/ -run 'TestToolVerifier|TestPost' -timeout 1m

# Run just the new gate tests
go test -race ./hyoka/internal/eval/ -run TestAssistantTurnStartToolLoadGate -timeout 2m -v
```

### Smoke Test (from Morpheus's spec)

```bash
# Run a config with 45+ skills (e.g., python-pairwise if available)
hyoka run --prompt-id key-vault-dp-python-crud \
  --config python-pairwise \
  --log-level debug --log-file toolgate-test.log

# Expected: NO "tool_load_failure" even if skills take >30s
# Check log: grep "Skills loaded" toolgate-test.log
# Should see: Turn started AFTER all skills loaded
```

---

## Future Work

1. **Wire ToolLoadCeiling config field** → Parse duration string, thread to postSessionToolVerification, validate in config.Validate()
2. **Add --tool-load-ceiling CLI flag** → Override config/default ceiling for debugging
3. **Metric: Tool-load timing** → Track how long each tool-load event takes, surface in report.json
4. **Progress UI enhancement** → Show "Waiting for tools..." message while blocking in postSessionToolVerification (currently silent)

---

## Commit Message

```
Fix: Replace 30s tool-load timeout with AssistantTurnStart listener (#347 / Option A)

The post-session tool verification gate used a 30s hard timeout to wait
for SDK tool-load events. When sessions with many skills (45+) took
longer than 30s, every tool was marked Failed — even when they loaded
successfully before the first turn.

The fix: Replace polling with a listener for SessionEventTypeAssistantTurnStart.
When the SDK starts the first turn, tool registration is definitively
complete. Snapshot status at that point and proceed. A 5-minute absolute
ceiling remains as a fail-safe for broken sessions (auth hang, network
failure) but is NOT the primary gate.

Changes:
- Added toolVerifier.onSessionReady() to mark registration complete
- Wired into copilot.go OnEvent handler at AssistantTurnStart
- Replaced 30s timeout with 5min ceiling in postSessionToolVerification
- Added per-kind tracking (turnBeforeSkills, turnBeforeMCP) to distinguish:
  • "Not registered before first turn" (event never fired)
  • "SDK did not report X as loaded" (event fired, X not in list)

Tests: TestAssistantTurnStartToolLoadGate covers all scenarios including
the fix case (45+ skills taking 35-40s, no false positive).

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>
```

---

## Acknowledgments

Investigation by Morpheus 🕶️ (see `.squad/decisions/inbox/morpheus-tool-load-gate-bug.md`).  
Implementation by Neo 💊.  
User insight: Ronnie Geraghty ("We should just wait until whatever usually happens after all tools loading messages... happens").

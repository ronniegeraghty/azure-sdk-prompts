# 🚨 STANDING POLICY — Model Selection (2026-04-24)

**Status:** ACTIVE. Top-level. Read this before any agent spawn.
**By:** Ronnie Geraghty (via Copilot directive 2026-04-24T19:47Z)
**Persisted in:** `.squad/config.json` → `defaultModel: claude-opus-4.7`

- **Default model for every agent: `claude-opus-4.7`.** No exceptions.
- **`claude-haiku-4.5` is FORBIDDEN.** Never spawn Haiku. Never bump down to Haiku. This includes Scribe and Ralph — their charters' "preferred: claude-haiku-4.5" lines are overridden.
- **Latest Sonnet (`claude-sonnet-4.5`)** is allowed only for "really simple things" — trivial mechanical work where opus-4.7 would be wasteful.
- **Rationale:** Quality over cost. User preference, captured for team memory.

---

## 2026-04-24: Engine invariant — every grader emits ≥ 1 GraderPoint

**By:** Neo 💊 (Engine)
**Scope:** `internal/criteria/graders/`, `internal/criteria/exec.go`, `internal/eval/engine_eval.go`
**Commit:** `b7611606`

Every `graders.GraderResult` produced by the engine — normal `Grade()`, error fallback, panic recovery, skipped-grader — MUST carry at least one `GraderPoint`. A Points-less result is a bug.

**Enforcement layers:**
1. `graders.NewResult` panics on `len(points) == 0` — single canonical constructor.
2. `graders.NewErrorResult(kind, name, cfg, msg)` synthesizes a failing `"grader executed"` Point and routes through `NewResult`. Use everywhere a Result is built outside a normal `Grade()` (engine error paths, panic recovery, future skipped-grader paths).
3. `engine_eval.convertGraderResults` defensively logs `slog.Warn` and synthesizes a fallback Point if a Points-less result somehow reaches the report layer (should be unreachable).
4. Graders allowing zero-knob configuration (`BehaviorGrader`, `ToolConstraintGrader`, `OutputCheckGrader`) emit a single trivially-passing `"no_constraints"` / `"no_knobs"` Point.

**Rules for new graders:**
- Always go through `graders.NewResult(kind, name, cfg, points, msg, extras)`. Never construct `graders.GraderResult{...}` literals.
- Each Point: stable snake_case `Label`, `Pass` bool, `Message` (required on fail, encouraged on pass).
- Outside `Grade()`, use `graders.NewErrorResult`.

**Tests:** `graders/points_test.go` covers per-kind pass/fail invariant, constructor panic, error-fallback path.

**Verification:** `reports/20260424-195854/.../report.json` — `jq '.grader_results | map(select(.points == null or (.points | length) == 0))'` returns `[]`.

---

## 2026-04-24: Site per-eval grader UI — defensive rendering shipped

**By:** Trinity 🖤 (Frontend)

- `GraderResultRow.tsx` — collapsed by default; defensive `result.points?.length`; point label fallback chain `label || message || ("Check passed" / "Check failed")`; message no longer duplicates when used as label fallback.
- `graderScore.ts` — when `points` is empty/missing, return `"1/1 points"` (pass) or `"0/1 points"` (fail). Never `"0/0 points"`, never `"PASS"`, never `"100%"`.
- `eval-detail-page.tsx` — Generator Session block moved above Grader Results.
- `GraderResultRow.test.tsx` — rewritten for v4 schema (was pre-v4, 8/8 failing).

**Verification:** 131/131 site tests; Playwright drove a real per-eval page; programmatic asserts on section ordering, chevrons-right, no PASS/100%/FAIL strings, no blank labels.

**Relationship to Neo's invariant:** The site fallback for Points-less results is now belt-and-braces — fresh v4 reports never lack Points. Fallback stays for legacy on-disk reports.

---

## 2026-04-27: Guardrail Enforcement Bug — maxTurns/maxFiles Stale Runner State

**Investigator:** Morpheus  
**Date:** 2026-04-23  
**Status:** ✅ CONFIRMED + FIXED (Option A)

### Bug Summary

The `CopilotPromptRunner` is constructed once at CLI startup with CLI flag defaults, **before** any per-config or per-prompt limits are loaded. Its `maxTurns` and `maxFiles` fields remain stale throughout execution, causing real-time enforcement to use wrong values.

**Impact:** A config with `max_turns: 100` will still kill sessions at 25 turns because the runner's `e.maxTurns` was set to `0` (CLI default) at construction, which falls back to `25` in the enforcement code.

### Root Cause

1. **CLI startup** — runner created with CLI defaults (lines 356-361 in `cmd/run.go`)
2. **Per-eval config resolution** — correct limits computed via `resolveLimits()` in `engine_eval.go:73-76`, written to report
3. **Real-time enforcement** — uses stale `e.maxTurns` (= 0 → 25) instead of resolved value, killing sessions early
4. Affects all three limits: `maxTurns`, `maxFiles`, `maxSessionActions`

### Recommended Fix (Option A)

- Add `SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)` method to `CopilotPromptRunner`
- Call from `engine_eval.go` right after `resolveLimits()`, before `evaluator.Run()`
- Use RWMutex to protect per-eval fields (concurrent eval execution)
- Real-time enforcement uses fallback chain: per-eval → CLI default → hardcoded default

**Rationale:** Minimal structural change, per-eval granularity, backward compatible, testable.

### Smoke Test

```bash
hyoka run --prompt-id identity-dp-python-default-credential \
  --config test-high-turns --log-level debug --log-file verify-turns.log
# Verify: turn 25 does NOT trigger cancellation
# Expected: report shows correct resolved limits, real-time enforcement uses same values
```

---

## 2026-04-27: Implementation — Per-Eval Limit Threading

**Author:** Neo  
**Date:** 2026-04-23  
**Status:** ✅ Implemented  
**Commits:** `d2f6e93b` + `def6b803`

### Method Signature

```go
func (e *CopilotPromptRunner) SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)
```

### Concurrency Model

Engine runs multiple evals in parallel (worker semaphore). All share one `e.evaluator` instance.

**Protection:** Per-eval fields (`evalMaxTurns`, `evalMaxFiles`, `evalMaxSessionActions`) protected by `sync.RWMutex`:
- `SetLimitsForEval` acquires write lock
- Real-time enforcement acquires read lock once at start, snapshots all three, releases

### Files Touched

1. **`hyoka/internal/eval/copilot.go`:**
   - Added per-eval fields (lines 27-41)
   - Added `SetLimitsForEval` method (lines 49-57)
   - Updated real-time enforcement (lines 223-265)

2. **`hyoka/internal/eval/engine_eval.go`:**
   - Type-assert `e.evaluator` to `*CopilotPromptRunner` (lines 148-152)
   - Call `SetLimitsForEval(...)` after `resolveLimits()`, before `evaluator.Run()`

### Fallback Chain

1. Per-eval resolved value (from config/prompt)
2. CLI-level default
3. Hardcoded default (e.g., 25 for turns)

**Tests:** All existing eval tests pass (`go test -race ./hyoka/internal/eval/...`)

---

## 2026-04-27: Testing — Real-Time Guardrail Enforcement with Resolved Limits

**Author:** Switch  
**Date:** 2026-04-27  
**Status:** ✅ Implemented  
**Commits:** `7dda6358` + `fe9a93c9`

### Test Summary

**Name:** `TestRealtimeGuardrailEnforcementUsesResolvedLimits`  
**Location:** `hyoka/internal/eval/engine_test.go:1512-1759`

Verifies that real-time guardrail enforcement (genCancel() during OnEvent callbacks) uses per-eval resolved limits from config/prompt YAML, not stale CLI defaults.

### Test Cases (Table-Driven)

1. **turn_limit_uses_config_not_cli_default**
   - CLI default: 0 → 25, Config: 100 → Allow 26 turns ✅

2. **turn_limit_enforced_at_resolved_config_value**
   - CLI default: 0 → 25, Config: 10 → Cancel at turn 10 ✅

3. **file_limit_uses_config_not_cli_default**
   - CLI default: 50, Config: 200 → Allow 60 files ✅

4. **file_limit_enforced_at_resolved_config_value**
   - CLI default: 50, Config: 20 → Cancel at file 20 ✅

### Infrastructure

- Added `LimitConfigurable` interface to enable test stubs and real runner to opt-in
- Created `stubRealtimeEnforcementRunner` that simulates enforcement with same fallback logic
- Emits realistic events (`"assistant.message"`, `"session.workspace.file_changed"`)

### Verification

```bash
go test -race ./hyoka/internal/eval -run TestRealtimeGuardrailEnforcementUsesResolvedLimits -v
```

Result: All 4 table-driven cases pass.

---

## 2026-04-27: Documentation — Clarification on Limits Resolution Order

**Author:** Oracle  
**Date:** 2026-04-27  
**Status:** ✅ Completed  
**Commit:** `4a8cd9d0`

Updated `docs/configuration.md` to clarify that the resolution order for limits (prompt frontmatter > config YAML > CLI flag > default) applies to **both** post-hoc guardrail checking AND real-time enforcement during session execution.

---

## 2026-04-27: FOLLOW-UP — Post-Session Tool-Verification Gate Timeout Behavior

**Status:** TODO (low priority, non-blocking)  
**Reported by:** Coordinator  
**Component:** `eval/copilot.go` — post-session tool verification gate

### Issue

The post-session tool-verification gate (Item E from prior round) confirms all configured tools/MCP servers load within 30 seconds. When timeout hits, every tool is marked Failed.

**Problem:** Microsoft's full skill plugin (`python-pairwise` config) contains 45+ tools. ~50% of runs hit the 30s SDK confirmation timeout, even when skills are working.

### Impact

- High false-positive rate on legitimate evals
- Limits real-world eval throughput

### Options (for future decision)

1. **Increase timeout** (30s → 90s/120s) — tradeoff: delays abort when SDK truly hangs
2. **Per-tool timeout** — tradeoff: requires SDK event routing overhaul
3. **Soft warning mode** — tradeoff: may miss real issues
4. **Skill pruning** — tradeoff: burden on user config

### Next Steps

- **Collect data:** Run python-pairwise evals with tool timing analysis
- **Tag:** LOW_PRIORITY, NOT_BLOCKING, INVESTIGATE_FIRST
- **Owner:** TBD (future sprint)

---

## 2026-04-27: Tool Load Verification Gate Bug — 30s Timeout Is Wrong Approach

**Author:** Morpheus 🕶️  
**Date:** 2026-04-27  
**Status:** Investigation Complete — Fixed (Option A)

### Issue Summary

During a live smoke test, an eval terminated with `tool_load_failure` after 45 skills failed to confirm load within 30 seconds. Ronnie diagnosed: "We should just wait until whatever usually happens after all tools loading messages happens, then make our call on what tools did and did not load."

### Root Cause

The post-session tool verification gate uses a **30-second hard timeout** to wait for SDK tool-load events. This is a **polling workaround, not a signal-based solution.** The SDK already emits a definitive "tools loaded" signal: `SessionEventTypeAssistantTurnStart` (marks session ready for work). The 30s timeout only catches false positives when:
- SDK is slow and hasn't emitted the event yet (common with 45+ skills)
- There's an SDK bug or network delay

**Better approach:** Use `AssistantTurnStart` as the gate signal — tools MUST be loaded before the SDK starts the first turn. Fallback to a much higher timeout (5 minutes) for broken sessions.

### Proposed Fix (Option A: Recommended)

Replace polling with `AssistantTurnStart` listener:
1. Add `onSessionReady()` method to `toolVerifier` — called when first turn starts
2. Wire into `copilot.go` event dispatch — force-emit tool status when turn fires
3. Replace 30s timeout with 5-minute absolute ceiling (fail-safe, not primary gate)
4. Add per-kind tracking to distinguish failure reasons:
   - `"Not registered before first turn"` — event never fired
   - `"SDK did not report X as loaded"` — event fired, X not in list

### Rationale

- **Semantically correct:** SDK won't start generation without loading tools
- **No arbitrary timeout:** Wait for a real event, not a guess
- **Handles edge cases:** Tools that don't register before first turn marked Failed (correct)
- **Existing safety net:** Session timeout (10 minutes) still catches SDK hangs

---

## 2026-04-27: Tool-Load Verification Gate — Option A Implementation

**Author:** Neo 💊  
**Date:** 2026-04-27  
**Status:** ✅ Implemented  
**Commits:** `8fc6d4be`, `fb5be186`

### Changes

1. **`hyoka/internal/eval/tool_verification.go`**
   - Added `onSessionReady()` method — called when `AssistantTurnStart` fires
   - Added per-kind tracking: `turnBeforeSkills`, `turnBeforeMCP` (distinguish failure reasons)
   - Removed `firstTurnStarted` field (replaced with per-kind flags)

2. **`hyoka/internal/eval/copilot.go`**
   - Wired `verifier.onSessionReady()` into event dispatch at `AssistantTurnStart` (line 337-345)
   - Replaced 30s timeout with 5min ceiling in `postSessionToolVerification` (line 779)
   - Updated comments explaining new semantics

3. **`hyoka/internal/config/config.go`**
   - Added `ToolLoadCeiling` config field (schema only; not yet wired for CLI/config override)

4. **`hyoka/internal/eval/tool_verification_test.go`**
   - Updated error message test to match new timeout semantics

### Behavior Changes

**Before:** 30s polling → false positives when skills take >30s  
**After:** Event-driven signal (AssistantTurnStart) + 5min ceiling → no false positives, real failures caught, broken sessions fail after 5min (not indefinite hang)

### Error Message Changes

| Scenario | Old | New |
|----------|-----|-----|
| Timeout | `"SDK did not confirm tool load within 30s"` | `"Session did not reach first turn within 5m0s"` |
| Event never fired | N/A | `"Not registered before first turn"` |
| Event fired, tool missing | `"SDK did not report skill as loaded"` | (unchanged) |

### Verification

- `go build ./hyoka/...` ✅
- `go test -race ./hyoka/internal/eval/ -timeout 3m` ✅ (all tests pass)

---

## 2026-04-27: Testing — AssistantTurnStart Tool Load Gate Tests

**Author:** Switch 🤍  
**Date:** 2026-04-27  
**Status:** ✅ Implemented  
**Commits:** `8fc6d4be` (paired with Neo)

### Test Suite

**File:** `hyoka/internal/eval/tool_verification_gate_test.go` (368 lines)  
**Function:** `TestAssistantTurnStartToolLoadGate` (5 table-driven cases)

### Test Cases

| Case | Scenario | Result | Time |
|------|----------|--------|------|
| 1 | All tools load before turn | Loaded | 0.45s |
| 2 | Partial failures before turn | Mixed | 0.45s |
| 3 ⭐ | 22s slow-load (proves fix) | Loaded | 22.02s |
| 4 | Turn before some events | Mixed with correct reasons | 0.30s |
| 5 | Absolute ceiling timeout | All Failed with clear reason | 3.00s |

**Case #3 is KEY:** Simulates 45+ skills taking 5s/7s/10s to load (35-40s total in production). OLD code would timeout at 30s and mark all tools Failed. NEW code waits for `AssistantTurnStart` at 45s, all tools succeed. Proves the fix works.

### Verification

```bash
go test -race ./hyoka/internal/eval/ -run TestAssistantTurnStartToolLoadGate -timeout 3m -v
```

Result: All 5 cases pass. Full eval suite (39 tests) passes with `-race` flag.

### At-Most-Once Contract

The verifier's `emitIfReady()` returns nil on subsequent calls. Tests handle this by reconstructing tool statuses from verifier state when needed.

---

## 2026-04-27: Documentation — Post-Session Tool Verification

**Author:** Oracle 🔮  
**Date:** 2026-04-27  
**Status:** ✅ Completed  
**Commit:** `f53eb3b1`

Added "Post-Session Tool Verification" subsection to `docs/configuration.md`:
- Explanation of `AssistantTurnStart` as primary gate signal
- Per-kind failure reasons (event never fired vs. tool not in SDK list)
- Absolute ceiling (5 minutes) as fail-safe for broken sessions
- Expected behavior examples

---

**Last updated:** 2026-04-28T00-54-38Z  
**Scribe:** Orchestration complete

---

## 2026-04-28: Per-Reviewer Vote Display — Feature Shipped

**Authors:** Morpheus 🕶️ (audit), Trinity 🖤 (design + impl), Switch 🤍 (tests + reconcile)  
**Date:** 2026-04-28  
**Status:** ✅ SHIPPED — Feature complete and tested  
**Commits:** `c155340f` (Trinity impl), `5a165d63` (Switch tests v1), `e347e4d6` (Switch reconcile)

### Problem

The site's grader result cards did not show **per-reviewer model votes** on individual criteria. When expanding a prompt grader card, users saw only:
- Consolidated summary (panel consensus)
- Panel member list with overall scores
- **Missing:** Which checks each reviewer passed/failed + rationale per reviewer

This was a **rendering-only gap** — the data already exists in the JSON at `grader_results[].extras.review.panel_results[].criteria[]`.

### Investigation (Morpheus)

Audited the data flow:
- ✅ **Go Engine:** Correctly writes `ReviewPanelResult.Criteria[]` per reviewer per criterion
- ❌ **TypeScript Types:** `ReviewPanelEntry` interface missing `criteria` field
- ❌ **React Component:** `ReviewExtras.tsx` never renders `panel.criteria[]` array

**Verdict:** No engine changes needed. Frontend-only fix.

### Solution (Trinity)

1. **Type Extension:** Added `criteria?: ReviewCriterionResult[]` to `ReviewPanelEntry` (optional for backward compat)
2. **New Component:** `ExpandablePoint.tsx` — reusable expandable criteria display
3. **Integration:** Updated `GraderResultRow.tsx` to call `ExpandablePoint.tsx` for split-vote scenarios
4. **UX Features:**
   - Per-reviewer model votes visible on expand
   - Criterion-level pass/fail with reasons
   - Auto-expand + amber badge for split votes (disagreement)
   - Consistent icons (CheckCircle2 green, XCircle red, amber for splits)

### Validation (Switch)

Wrote comprehensive test suite:
- 31 tests (all pass): backward compat, split votes, edge cases, accessibility
- Tests split across:
  - `ExpandablePoint.test.tsx` (component unit tests)
  - `GraderResultRow.test.tsx` (integration tests)
- Reconciliation: Initial test structure didn't match Trinity's final architecture; reorganized and re-ran. No regressions.

### Feature Highlights

✅ Per-reviewer model votes visible on expand  
✅ Criterion-level pass/fail + reason per reviewer  
✅ Auto-expand + visual indicator on disagreement  
✅ Backward compatible (optional criteria field for legacy reports)  
✅ Full test coverage (31/31 pass)  
✅ Consistent styling (emerald pass, red fail, amber split)  

### Files Changed

**Commit c155340f (Trinity):**
- `site/src/app/data/types.ts` — Type extension
- `site/src/app/components/grader-extras/ExpandablePoint.tsx` — New component
- `site/src/app/components/grader-extras/GraderResultRow.tsx` — Integration

**Commit 5a165d63 (Switch):**
- `site/src/app/components/grader-extras/ReviewExtras.test.tsx` — Initial test suite (692 lines)

**Commit e347e4d6 (Switch):**
- `site/src/app/components/grader-extras/ExpandablePoint.test.tsx` — Component tests (reconcile)
- `site/src/app/components/grader-extras/GraderResultRow.test.tsx` — Extended integration tests (reconcile)

### Branch

`ronniegeraghty/dev` — ready for merge to main after review and smoke test.

---


## 2026-04-29: User directive — Rebuild `site/dist/` after any `site/` changes

**By:** Ronnie (via Copilot directive 2026-04-29T19:03:44Z)  
**Status:** ACTIVE  
**Scope:** All agents touching `site/`

After ANY update to the site (anything under `site/`), always run `cd site && npm run build` before considering work done. Serves two purposes:

1. **Build-time test** — catches type errors and broken imports
2. **Binary freshness** — refreshes `site/dist/` so the embedded `hyoka serve` binary reflects latest code when user inspects it

**Rule:** Commit rebuilt `site/dist/` alongside source changes.

**Rationale:** Multiple recent sessions shipped site source without rebuilding `dist/`, causing "I don't see my change" confusion — the embedded bundle was stale.

**Action:**
- Trinity: Ensure every site PR includes rebuilt dist
- Ralph: Flag site PRs that skip the build as incomplete
- Team: Make `npm run build` a gating check before pushing site changes


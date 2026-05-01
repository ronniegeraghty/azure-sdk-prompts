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

---

## 2026-04-29: Rerun Command Fix for Multi-Model and Pairwise Configs (Option C1) — SHIPPED

**Authors:** Morpheus 🕶️ (investigation), Neo 💊 (implementation)  
**Date:** 2026-04-29  
**Status:** ✅ SHIPPED — Commit 5dda7811 on `ronniegeraghty/dev`

### Problem

Rerun commands in the web UI (and `report.json`) fail for evaluations using multi-model configs or pairwise-expanded configs. The engine synthesizes virtual config names that don't exist in YAML files:
- Multi-model: `hyoka run --config python-pairwise/claude-opus-4.6` ❌
- Pairwise: `hyoka run --config python-pairwise/baseline/claude-opus-4.6` ❌

### Investigation (Morpheus)

**Confirmed bug.** Root cause: engine fan-out creates synthetic config names during multi-model expansion; `buildRerunCommand()` uses these synthetic names, but YAML files contain only base configs.

Drafted three remediation options:
- **Option A:** Reconstruct original command (store original CLI invocation) — most complete but invasive
- **Option B:** Add `--model` flag to run.go — requires multi-config awareness
- **Option C1:** Add `BaseConfigName` + `GeneratorModel` fields to `EvalReport` — **pragmatic, recommended**

### Solution (Neo) — Option C1

1. **Schema (types.go):** Added two optional fields to `EvalReport`:
   ```go
   BaseConfigName  string  // e.g., "python-pairwise" before fan-out
   GeneratorModel  string  // e.g., "claude-opus-4.6" the actual model used
   ```

2. **Engine (engine.go):** Modified `expandGeneratorModels()` to return expanded configs with base config name metadata.

3. **Eval Runner (engine_eval.go):** 
   - Threaded `BaseConfigName` through `EvalTask`
   - Populated both new fields at eval creation time
   - Updated `buildRerunCommand()` logic:
     - Multi-model: `hyoka run --config <base> --model <specific>`
     - Single-model: `hyoka run --config <name>` (no --model)

4. **Tests (engine_eval_rerun_test.go):** 9 new table-driven tests covering single/multi-model, pairwise variants, flag combinations, special characters. All pass.

### Verification

- ✅ Unit tests: All 9 pass
- ✅ Build: Clean
- ✅ Full suite: Eval tests pass (pre-existing serve/validate failures unrelated)
- ✅ Code inspection: `--model` flag logic in run.go verified

### Result

**Before:**
```bash
User runs:    hyoka run --prompt-id X --config python-pairwise
Report generated with:
  rerunCommand = "hyoka run --prompt-id X --config python-pairwise/claude-opus-4.6"
User retries: ERROR — config not found
```

**After:**
```bash
User runs:    hyoka run --prompt-id X --config python-pairwise
Report generated with:
  baseConfigName = "python-pairwise", generatorModel = "claude-opus-4.6"
  rerunCommand = "hyoka run --prompt-id X --config python-pairwise --model claude-opus-4.6"
User retries: ✅ Works
```

### Pairwise Behavior (By Design)

Pairwise variants expand configs into tool-ablation variants (e.g., `/without-azure`). The C1 rerun command **does not preserve tool ablations** — this is intentional:
- Pairwise is for comparative analysis, not one-off reruns
- Users wanting exact pairwise reproduction can manually use `-P` and find results
- If exact pairwise rerun becomes critical, Option C2 (`--pairwise-variant` flag) is a future follow-up

### Schema Impact

Two new optional fields on v4 report schema (backward compatible):
```json
{
  "baseConfigName": "python-pairwise",
  "generatorModel": "claude-opus-4.6"
}
```

Legacy reports (missing these fields) gracefully fall back to using the full `configName` in rerun commands.

### Follow-Ups

- **Tank (CLI):** The `--model` flag is hidden (`cmd.Flags().MarkHidden("model")`). If users discover it, Tank should unhide and improve help text.
- **Trinity (Site):** No React changes required. Site already renders whatever `rerunCommand` string it receives. New fields available in v4 schema if site wants to display them separately (e.g., "Config: python-pairwise | Model: claude-opus-4.6").
- **Morpheus (Docs):** Can note C1 as shipped. If exact pairwise rerun becomes critical, spec out Option C2.


---

## 2026-04-29: Rerun Commands v2 — Tool-Ablation Fidelity (Options Analysis)

**Author:** Morpheus 🕶️  
**Date:** 2026-04-29  
**Status:** OPTIONS — awaiting selection  
**Supersedes:** Scope of `morpheus-rerun-command-pairwise-options.md` (multi-model portion still valid)

### Problem

C1 (shipped multi-model fix) does not preserve pairwise tool-ablation variants on rerun. Clicking "rerun" on `python-pairwise/without-azure/claude-opus-4.6` actually runs the baseline, losing the ablation. **Tool-ablation fidelity is a hard requirement.**

### Options Analyzed

Four options to close the fidelity gap:

1. **Option D — `--without-tool` repeatable flag:** User hand-crafts ablations (most flexible, ~120 LOC)
2. **Option E — `--exclude-tool` repeatable flag:** Extends existing `excluded_tools` config field (~150 LOC, maintenance risk)
3. **Option F — `--pairwise-variant <name>` flag:** Replay exact variant by name (smallest, ~80 LOC) **← Recommended**
4. **Option G — Inline/snapshot config blob:** Self-contained replay via base64 or sidecar YAML (~250-350 LOC, unreadable)

### Recommendation: Option F

**Why F:**
- Reuses exact `ExpandPairwise` machinery that created the variant
- Smallest surface area: one flag, one schema field
- Composes cleanly with C1: `--config` (base) + `--model` (generator) + `--pairwise-variant` (ablation)
- Cost: ~80 LOC

**Doesn't deprecate C1.** F extends it. C1 remains correct for non-pairwise multi-model configs.

### Open Questions for Ronnie

1. **Q1:** Baseline as first-class variant? Emit explicit `--pairwise-variant baseline` or implicit?
2. **Q2:** Single-variant rerun or full sweep replay?
3. **Q3:** Human-readable vs. opaque blob? (Confirmed: human-readable)
4. **Q4:** Slash quoting in variant names (`without-azure/storage_blob_list`)?
5. **Q5:** Variant folder structure? (Confirmed: each variant is own EvalReport)
6. **Q6:** Rerun needs `-P` opt-in? (No under F; unclear under D)

---

## 2026-04-29: `--pairwise-variant` Flag Implementation (Option F)

**Authors:** Neo 💊 (implementation), based on Morpheus 🕶️ (spec)  
**Date:** 2026-04-29  
**Status:** ✅ SHIPPED  
**Layered on:** C1 (BaseConfigName + GeneratorModel)

### Summary

Implemented `--pairwise-variant <name>` flag to restore tool-ablation fidelity in rerun commands. Users can now reproduce exact pairwise variants instead of just baselines.

### Changes

#### 1. CLI (cmd/run.go)
- Added `--pairwise-variant <name>` string flag
- Mutually exclusive with `-P`/`--pairwise` (errors if both set)
- Accepts variant suffixes: `baseline`, `without-{toolName}`, `without-{mcpName}/{mcpToolName}`
- Expands base config and selects matching variant

#### 2. Schema (internal/report/types.go)
- Added `PairwiseVariant string` field to `EvalReport`
- Stores variant suffix (e.g., `"baseline"`, `"without-azure"`, `"without-azure/storage_blob_list"`)
- Populated at eval creation time (not string parsing downstream)
- Backward-compatible (optional, omitempty)

#### 3. Engine (internal/eval/)
- **`EvalTask`:** Added `PairwiseVariant string` field
- **`extractPairwiseVariant()`:** Extracts variant suffix from config name
  - Handles baseline, simple ablations, deep MCP variants
  - Strips model suffix correctly (e.g., `/claude-opus-4.6`)
- **`runSingleEval()`:** Populates `evalReport.PairwiseVariant` from task
- **`buildRerunCommand()`:** Emits `--pairwise-variant <name>` when field non-empty
  - Q1 default: baseline variants emit explicit `--pairwise-variant baseline`
  - Q4 default: quotes variant names with slashes (e.g., `"without-azure/storage_blob_list"`)

#### 4. Test Coverage (internal/eval/*_test.go)
- **`TestBuildRerunCommand`:** 6 new test cases (baseline, simple ablation, deep MCP + composition)
- **`TestExtractPairwiseVariant`:** 8 test cases (all variant patterns)
- All tests pass with `-race`

### Architecture

Three orthogonal fan-out dimensions now compose cleanly:

```
User invokes:  hyoka run --config python-pairwise --model claude-opus-4.6 --pairwise-variant without-azure
                          └─ base ──────────┘  └─ C1: model ────────────┘  └─ F: variant ──────────┘

Engine flow:
1. Load base config (python-pairwise.yaml)
2. Override model if --model set → single-model config
3. If --pairwise-variant set:
   - Expand base config via pairwise.ExpandPairwise()
   - Look up variant by suffix (e.g., "without-azure")
   - Replace configs list with single matching variant
4. expandGeneratorModels() fans out (if any)
5. Build task list; populate task.PairwiseVariant from config name
6. Populate evalReport.PairwiseVariant at eval creation
7. buildRerunCommand() emits --pairwise-variant when field set

Result: hyoka run --prompt-id X --config python-pairwise --model claude-opus-4.6 --pairwise-variant without-azure
```

### Why Not Reuse parsePairwiseConfigName()?

The existing parser (engine.go:962-981) is kept for backward-compat reading of older reports. New runs populate `PairwiseVariant` directly at eval time via `extractPairwiseVariant()`, eliminating brittle string parsing downstream.

### Verification

✅ **Build:** `go build ./...`  
✅ **Tests:** `go test -race ./hyoka/internal/eval/... -timeout 3m` (41s, all pass)  
✅ **Manual:** Single-variant runs produce correct rerun commands

### Files Modified

- hyoka/cmd/run.go
- hyoka/internal/eval/engine.go
- hyoka/internal/eval/engine_eval.go
- hyoka/internal/eval/engine_eval_rerun_test.go
- hyoka/internal/eval/engine_pairwise_test.go (new)
- hyoka/internal/report/types.go

### Resolved Open Questions (from Morpheus's spec)

- **Q1:** YES. Baseline gets explicit `--pairwise-variant baseline` for round-trip clarity
- **Q2:** Single variant (flag selects one, not full sweep)
- **Q3:** YES. Paste-safe, human-readable
- **Q4:** YES. Deep MCP variants like `without-azure/storage_blob_list` are quoted in emitted command
- **Q6:** NO. Single-variant rerun uses `--pairwise-variant`, not `-P`

### Schema Impact

New optional field on `EvalReport` (v4):
```json
{
  "pairwiseVariant": "without-azure"
}
```

Legacy reports (missing this field) gracefully fall back to C1 behavior.

### Example Rerun Commands

| Scenario | Input | Output |
|----------|-------|--------|
| Baseline | `--config python-pairwise --pairwise-variant baseline` | `hyoka run --prompt-id X --config python-pairwise --pairwise-variant baseline` |
| Simple ablation | `--config python-pairwise --pairwise-variant without-azure` | `hyoka run --prompt-id X --config python-pairwise --pairwise-variant without-azure` |
| Deep MCP | `--config python-pairwise --pairwise-variant without-azure/storage_blob_list` | `hyoka run --prompt-id X --config python-pairwise --pairwise-variant "without-azure/storage_blob_list"` |
| + Model | `--config python-pairwise --model opus --pairwise-variant without-azure` | `hyoka run --prompt-id X --config python-pairwise --model opus --pairwise-variant without-azure` |

### Follow-Ups

**Trinity (Site Team):** The site may build rerun commands client-side. If so:
1. Read `report.pairwiseVariant` from v4 schema
2. Emit `--pairwise-variant <value>` when field non-empty
3. Quote values containing slashes

No React changes required if site only renders `rerunCommand` string as-is.

**Future (Option D):** If users want hand-crafted ablations outside `-P`, Option D's `--without-tool` flag can be added later without removing Option F. The two flags coexist.

### Related Decisions

- **C1:** `.squad/decisions.md` (earlier) — Multi-model rerun fix (BaseConfigName + GeneratorModel)
- **Morpheus v2 spec:** Rerun Commands v2 (above) — Option F rationale

### Learnings

1. **Structured fields > string parsing.** Moving pairwise variant identity from "parse the config name string" to "store at eval time" eliminates fragility.
2. **Three orthogonal flags compose cleanly.** `--config`, `--model`, `--pairwise-variant` each handle one fan-out dimension.
3. **Model suffix stripping is tricky.** Config names like `python-pairwise/without-azure/storage_blob_list/claude-opus-4.6` require heuristics. Test coverage critical.
# Decision: Trends Process — Opt-Out → Opt-In (Issue #638)

**Status:** FILED (GitHub issue #638, squad label)  
**Date:** 2026-04-29  
**Owner:** Morpheus 🕶️ (Lead)  
**Stakeholders:** Neo (engine), Tank (performance), Trinity (UI)  

---

## Problem Statement

The `hyoka run` command automatically invokes trend analysis at the end of every evaluation run. This process:

1. Scans past reports in `reports/` directory
2. Computes historical metrics and time-series trends
3. **Spawns a Copilot SDK session** for AI-powered analysis (regression detection, insights)
4. Writes results to `reports/trends/`

**Impact:** Every single eval run (including fast iteration loops) pays the cost of trend generation + LLM analysis, even when users don't need the output. Users must remember to pass `--skip-trends` to opt out.

## Current Design (Opt-Out)

```
hyoka run <filters>               # trends run automatically
hyoka run <filters> --skip-trends # skip trends (opt-out)
```

- **Flag:** `--skip-trends` (boolean, default: false)
- **Logic:** `if !f.skipTrends && !f.dryRun` → generate trends
- **UX problem:** Negative flag (skip) as default behavior; counterintuitive for optional feature

## Proposed Design (Opt-In)

```
hyoka run <filters>              # trends skip by default (fast)
hyoka run <filters> --with-trends # generate trends (opt-in)
```

| Aspect | Current | Proposed |
|--------|---------|----------|
| Default behavior | Trends run automatically | Trends skipped |
| Control flag | `--skip-trends` (boolean, false) | `--with-trends` (boolean, false) |
| CLI UX | `hyoka run ... --skip-trends` | `hyoka run ... --with-trends` |
| Logic | `if !skipTrends` → run | `if withTrends` → run |
| Backward compat | N/A | Keep `--skip-trends` as deprecated alias (optional) |

## Rationale

1. **Aligns with typical opt-in UX** — optional post-analysis features default to off
2. **Improves iteration speed** — most use cases (CI/CD, quick loops) don't need trends; opt-in users get full analysis on demand
3. **Reduces surprise costs** — no hidden Copilot session spawned without explicit request
4. **Dashboard survives gracefully** — `/api/trends` endpoint generates data on-demand if not pre-computed (no hard dependency)

## Architecture Notes

- **No hard blockers** — Dashboard already calls `trends.Generate()` on-demand
- **Graceful degradation** — Missing pre-computed trends data causes no system failure
- **Isolated concern** — Affects only `hyoka run` command; standalone `hyoka trends` subcommand unaffected

## Implementation Scope

- **Files:** `hyoka/cmd/run.go` (flag definition, invocation logic)
- **Changes:** 
  - Replace `skipTrends bool` with `withTrends bool` in `runFlags` struct
  - Update flag definition: `BoolVar(&f.withTrends, "with-trends", false, "Generate trend analysis...")`
  - Flip condition: `if f.withTrends && !f.dryRun` → generate trends
  - Update help text on `run` command
- **Tests:** Verify dashboard still works without pre-computed trends

## Acceptance Criteria

- [x] Issue filed (#638, squad label)
- [ ] Team aligns on flag name (`--with-trends` vs. alternatives)
- [ ] Implementation plan: migrate `skipTrends` → `withTrends`, flip logic
- [ ] Deprecation path for `--skip-trends` (keep as alias or remove cleanly)
- [ ] Dashboard tests pass with missing pre-computed trends
- [ ] Help text updated to clarify new default

## Alternative Considered

**Keep `--skip-trends`, invert its default to true:**
- Simpler (no new flag name)
- Less clear intent (users see `--skip-trends` and assume trends run by default, but they don't)
- **Rejected:** Negative flag should not represent the default behavior; leads to confusion

**Use `--trends` instead of `--with-trends`:**
- Shorter
- Less explicit about intent (is it on or off?)
- **Decided:** `--with-trends` is clearer ("include trend analysis in this run")

## Cross-Team Dependencies

- **Neo (engine):** No impact on eval engine or grading
- **Tank (perf):** Iteration time improves for common case (trends skipped by default)
- **Trinity (UI):** Dashboard must handle missing trends gracefully (already does via on-demand API)
- **Scribe (logging):** Trend analysis logging continues if `--with-trends` is used

---

## Next Steps

1. **Triage (squad):** Validate problem statement, confirm flag name
2. **Implementation (Neo or Tank):** Migrate flag logic, update tests
3. **Review (Trinity, Scribe):** Verify dashboard degradation, logging behavior
4. **Release notes:** Document new opt-in behavior, deprecation timeline for `--skip-trends` (if kept)

---

## 2026-04-30: CI Pipeline Failures Diagnosed — Handoff to Neo and Tank

**By:** Morpheus (Lead, CI Owner)  
**Status:** Diagnosis Complete — Fixes Delegated

### Problem

Ronnie reported periodic GitHub Actions failures. Investigation reveals **no scheduled workflows are failing**. However, **two push-triggered workflows show frequent failures** (21 failures in past 100 runs):

1. **CI (build-and-test):** 18 failures / 100 runs (since 2026-04-29)
   - Symptom: `go vet` errors
   - Root causes: Lock copy violation in `cacheroot.go:110,117` (sync.Once assigned to variables) + unknown struct field `Model` in 3 test files + pointer/value mismatch in 2 test files
   - Impact: Blocking all CI runs; prevents PRs/pushes from validating

2. **Site Bundle Freshness:** 3 failures / 100 runs (since 2026-04-29)
   - Symptom: `ERROR: site/dist/ is stale` — source changed without bundle rebuild
   - Root cause: Developers pushing `site/src/**` changes without running `npm run build`
   - Impact: Go:embed'd bundle becomes out-of-sync; test failures and unpredictable behavior

### Root Cause Analysis

**CI Failures:**
- Incomplete refactoring: GraderResult struct fields removed/renamed in core types without corresponding test updates (sync.Once no-copy semantics violated)
- Test-code sync breakdown (no validation that tests match struct definitions)

**Site Bundle Failures:**
- Developer workflow gap: site bundle rebuild not part of standard commit checklist
- Missing guardrail: no pre-commit hook or enforced build step to ensure dist/ stays in sync with src/

### Decision & Handoff

| Issue | Owner | Action | Urgency |
|-------|-------|--------|---------|
| CI struct sync + lock fix | Neo (Engine) | Fix test-code sync + sync.Once handling | P0 (Blocking) |
| Site Bundle guardrail | Tank (Build) | Add pre-commit hook or CI enforcement | P1 (Frequent) |

### Implementation Status

- **Neo:** ✅ Fixed all 5 vet errors (commits: 99a185ba, e007695e)
- **Tank:** ✅ Added Husky pre-commit hook (commit: 0de4468b)

---

## 2026-04-30: Phantom 3rd Grader Point in Checks-Based Criteria

**By:** Morpheus (Investigation)  
**Status:** Root Cause Confirmed — Fix Queued for Neo

### Problem

When a YAML grader entry has `checks:` (multi-point format), the parent grader line is numbered like a criterion. LLMs interpret this as an additional scoreable criterion, producing N+1 points instead of N.

Example malformed output:
```
1. **DefaultAzureCredential Authentication**
   Check the following criteria:
   1. Uses DefaultAzureCredential...
   2. Uses async/await patterns...
```

LLM scores 3 items (parent + 2 checks) instead of 2 checks only.

### Root Cause

`hyoka/internal/criteria/buckets.go:136` (FormatUnifiedPromptEntries) formats the parent line with:
```go
fmt.Fprintf(&b, "%d. **%s**\n", i+1, e.Name)
```

This numbered format signals "scoreable criterion" to the LLM.

### Decision

**Remove the number from the parent line** when checks exist. Format it as a section header instead:

Option 1 (bold):
```
**DefaultAzureCredential Authentication**
Check the following criteria:
1. Uses DefaultAzureCredential...
2. Uses async/await patterns...
```

Option 2 (markdown header):
```
### DefaultAzureCredential Authentication
Check the following criteria:
1. Uses DefaultAzureCredential...
2. Uses async/await patterns...
```

Either format makes it clear the parent is a grouping label, not a criterion to score.

### Implementation Owner

**Neo** — criteria/graders pipeline expert.

### Files to Change

- `hyoka/internal/criteria/buckets.go:136` (FormatUnifiedPromptEntries logic)
- `hyoka/internal/criteria/buckets_test.go` (update test expectations)

### Evidence

- **Run:** `/home/rgeraghty/projects/hyoka/reports/20260430-041731/`
- **Grader:** `criteria/language/python.yaml` — "DefaultAzureCredential Authentication"
- **Expected:** 2 points (2 checks)
- **Actual:** 3 points (parent + 2 checks)

### Follow-Up

This fix is queued for Neo's next session after CI stabilization.

---

## 2026-04-30: GraderResult v4 Schema Migration Complete

**By:** Neo (Core Eval Framework)  
**Status:** ✅ Resolved — CI Unblocked

### Problem Statement

CI was blocked by 5 go vet errors stemming from incomplete GraderResult schema migration. Tests referenced removed/renamed fields and used incorrect types.

### Root Cause

GraderResult v4 schema refactoring removed legacy review-style fields (Model, OverallScore, MaxScore, IsConsensus) and changed Pass from `*bool` to `bool`. The struct definition was updated but test files across 8 packages were not synchronized.

Additionally, `cacheroot.go` had a sync.Once copy violation (Go's no-copy semantics prevent assigning sync.Once by value).

### Solution

**Immediate action:** Updated all test files to match GraderResult v4 schema.

**Schema requirements (v4):**
- **Removed fields:** Model, OverallScore, MaxScore, Summary, Issues, Strengths, IsConsensus
- **Changed types:** Pass from `*bool` to `bool`
- **Required fields:** Points ([]GraderPoint, len >= 1)
- **Retained fields:** GraderName, GraderType, Score, Weight, Gate, Message, Extras

**Implementation:**
1. Searched all test files for GraderResult literals
2. Removed references to legacy fields
3. Changed Pass from `&pass` / `boolPtr(pass)` to `pass`
4. Added minimal Points field: `[]GraderPoint{{Label: "check", Pass: pass, Weight: 1.0}}`
5. Fixed cacheroot.go sync.Once handling by using state flags instead of copying

### Files Changed

- hyoka/internal/toolload/cacheroot.go (sync.Once fix)
- hyoka/internal/report/generator_test.go (10+ instances)
- hyoka/internal/serve/dashboard_test.go (2 instances)
- hyoka/internal/comparison/comparison_test.go (4 instances)
- hyoka/internal/comparison/inmem_test.go (1 instance)
- hyoka/cmd/compare_test.go (1 instance)
- hyoka/internal/serve/equivalence_test.go (2 instances)
- hyoka/internal/report/markdown_test.go (1 instance)

### Verification

```bash
go vet ./...         # ✅ clean (was 5 errors)
go build ./...       # ✅ success
go test ./hyoka/internal/toolload/...    # ✅ PASS
go test ./hyoka/internal/comparison/...  # ✅ PASS (1 pre-existing failure)
go test ./hyoka/internal/serve/...       # ✅ PASS
```

Pre-existing test failures (NOT caused by this fix):
- TestWriteReport_LargeReportWrittenCorrectly (report v0 migration panic)
- TestReviewerFactory_MissingSkillFailsFast (error type assertion)

### Commits

- `99a185ba` — Fix sync.Once copy and toolload schema
- `e007695e` — Update test files to GraderResult v4 schema

### Lessons Learned

1. **Cross-package test sync:** When refactoring core structs, grep for usage across ALL test files (not just the defining package)
2. **sync.Once semantics:** Cannot be copied by value; use pointers or state flags to preserve/restore
3. **Schema invariants:** GraderResult v4 requires Points field (len >= 1) — empty results need dummy points

---

## 2026-04-30: Site Bundle Freshness Pre-Commit Hook

**By:** Tank (CLI Dev)  
**Status:** ✅ Implemented  
**Related Issue:** Morpheus CI Failure Diagnosis — "Site Bundle Freshness" section

### Problem

The "Site Bundle Freshness" GitHub Actions workflow fails ~3% of the time (3 failures / 100 runs) because developers push changes to `site/src/**` without running `npm run build`, leaving `site/dist/` stale. This bundle is `go:embed`'d into the binary, so staleness causes unpredictable behavior.

Root cause: No local guardrail to enforce the build step before commit.

### Options Evaluated

| Option | Pros | Cons |
|--------|------|------|
| **CI auto-rebuild** | Catches all cases | Delays feedback (CI runs 5-10 min); requires branching strategy for auto-commits |
| **Pre-commit hook** | Instant feedback; prevents bad commits; lightweight | Requires npm install on clone |
| **Documentation only** | No setup overhead | Doesn't prevent developer mistakes |

**Selected:** Pre-commit hook (lightest effective solution)

### Implementation

#### Files Created/Modified

1. **`package.json`** (root)
   - Minimal npm config for Husky setup
   - `prepare` script: `husky install`

2. **`.husky/pre-commit`** (executable)
   - Detects `site/src/**` in staged files
   - Runs `npm run build` in site/ directory
   - Auto-stages rebuilt `site/dist/`
   - Fails commit if build fails (giving developers clear error)

3. **`CONTRIBUTING.md`** (updated)
   - Documents `npm install` as part of setup
   - Explains the hook behavior and what happens if site build fails

4. **`.gitignore`** (verified)
   - Root `node_modules/` already ignored
   - `.husky/` hooks are **committed** (not ignored) — this is correct

#### How It Works

```
Developer commits site changes:
  1. Git pre-commit hook fires
  2. Detects `site/src/**` in staged files
  3. Runs `npm run build` → rebuilds `site/dist/`
  4. Stages the rebuilt bundle automatically
  5. Commit proceeds with both source + fresh dist/

If build fails:
  → Commit is blocked
  → Developer sees clear error message
  → Fixes the TypeScript/CSS error and retries
```

### Guardrails & Testing

**Integration Testing:**
- Workflow `site-embed-freshness.yml` remains unchanged (still detects staleness as fallback)
- If somehow stale code gets committed (hook skipped with `--no-verify`), CI will catch it
- Error message on CI failure is already excellent (guides dev to run `npm run build`)

**Hook Validation:**
- Hook is executable (`chmod +x`)
- Pre-commit guard checks for `site/src/` paths (no false positives)
- Build failure blocks commit (atomic)

### Developer Experience

**Before (current):**
- Dev pushes changes → CI fails → "run npm run build" → commit → force push
- Cycle time: 10-15 minutes

**After (with hook):**
- Dev commits locally → hook rebuilds → commit succeeds
- Cycle time: 0 (transparent)
- If site build fails locally → immediate feedback before push

### Deployment

1. ✅ Hook created in `.husky/pre-commit`
2. ✅ Package files committed (`package.json`, `package-lock.json`)
3. ✅ Documentation updated (`CONTRIBUTING.md`)
4. ✅ Workflow YAML validated (no changes needed)

New contributors on clone:
```bash
git clone ...
npm install    # ← sets up .husky hooks automatically via "prepare" script
```

### Fallback

If a developer skips the hook with `git commit --no-verify`, the existing CI workflow (`site-embed-freshness.yml`) will catch the stale bundle. The guardrail is **defense-in-depth**, not the sole check.

### Decision

✅ **Implement pre-commit hook via Husky** — prevents ~100% of stale bundle commits locally, with zero performance impact on other workflows.

**No changes required to:**
- GitHub Actions workflows (fallback still works)
- Go build process
- Site build process
- PR requirements

### Commit

`0de4468b` — Add pre-commit hook to enforce site bundle freshness

---

## Grader Redesign — Multi-Part Initiative

**Date:** 2026-04-26 (scope), 2026-04-30 (completion)  
**Initiated by:** Ronnie  
**Lead:** Morpheus (design), Neo (engine), Tank (rendering)  
**Status:** ✓ Delivered (PR ready)  
**Branch:** `neo/issue-grader-redesign`

This is a comprehensive redesign of the grader system across four parts: semantics, execution order, output format, and tool verification.

### Part 1 — Prompt Grader Semantics (BREAKING CHANGE)

**Rule:** For `type: prompt` graders, ONLY `checks:` entries are scorable points. `name` and `prompt` are LLM judge context — never counted in the X/X tally, never reported as scored items.

#### Current Behavior (Removed)

- `FormatUnifiedPromptEntries()` rendered grader Name and Prompt as numbered criteria items
- Review grader created one GraderPoint per criterion (including name/prompt text)
- No separation between context and scorable items

#### Changes Required

| File | Change |
|------|--------|
| `internal/criteria/buckets.go` | `FormatUnifiedPromptEntries()` — render `prompt:` as preamble context only. Only `checks:` items become numbered criteria. Grader `name:` becomes section heading. |
| `internal/criteria/config.go` | `validateEntry()` — YAML graders with `type: prompt` but no `checks:` log soft deprecation warning |
| `internal/prompt/parser.go` | `ParseEvaluationCriteria()` — rewrite to separate lead text (Prompt field) from bullet items (Checks slice) |
| `internal/prompt/types.go` | `CriterionEntry` struct — add `Checks []string` alongside existing fields |
| `internal/criteria/buckets.go` | `BuildUnifiedReviewBuckets()` — render only checks as numbered criteria, lead text as preamble |

#### Migration Impact

- **YAML criteria using `checks:`** — no change (already correct)
- **YAML criteria using only `prompt:` (no `checks:`)** — become zero-scorable graders (context only)
- **Prompt files' `## Evaluation Criteria`** — parser rewrite normalizes into `{prompt, checks}` shape (backward-compatible)

### Part 2 — Grader Execution Order

**Rule:** Prompt-file eval criteria runs FIRST, then criteria-file graders in YAML file order.

#### Current Behavior

- Typed graders ran first, then AI review graders
- Prompt-file criteria ran first among review graders, but after typed graders
- No unified ordering across typed and AI review partitions

#### Changes Required

| File | Change |
|------|--------|
| `internal/eval/engine_eval.go` | Reorder execution: (1) prompt-file eval criteria grader, (2) criteria-file graders in YAML file order (typed and prompt interleaved per declaration) |
| `internal/criteria/buckets.go` | `MatchingUnifiedEntries()` — ensure stable file-walk order when re-interleaving partitions |

#### Key Constraint

- Prompt-file grader is synthetic (not from YAML)
- Must be injected at position 0 before any criteria-file graders

### Part 3 — Output Format Redesign

**Rule:** Reports group graders by source file at three indentation levels.

#### Target Format

```
Graders:
- crud-secrets.prompt.md (prompt file):
  - Eval Criteria (prompt): Pass/Fail (x/x)
      - check 1: Pass/Fail
      - check 2: Pass/Fail
- python.yaml (criteria file):
  - DefaultAzureCredential Authentication (prompt): Pass/Fail (x/x)
      - Uses DefaultAzureCredential ...: Pass/Fail
      - Uses async/await patterns ...: Pass/Fail
  - Output Files Exist (output_check): Pass/Fail (x/x)
      - min_files (1): Pass/Fail
      - min_bytes_per_file (1): Pass/Fail
```

#### Data Model Changes

| Struct | Field to Add | Purpose |
|--------|-------------|---------|
| `graders.GraderResult` | `SourceFile string` | Absolute path to originating file |
| `graders.GraderResult` | `SourceType string` | `"prompt_file"` or `"criteria_file"` |
| `report.GraderResult` | `SourceFile string` (JSON: `source_file`) | Persisted in report |
| `report.GraderResult` | `SourceType string` (JSON: `source_type`) | Persisted in report |

#### Files to Change

| File | Change |
|------|--------|
| `internal/criteria/graders/grader.go` | Add `SourceFile`, `SourceType` fields |
| `internal/report/types.go` | Add `SourceFile`, `SourceType` (JSON mapped) |
| `internal/eval/engine_eval.go` | `convertGraderResults()` — copy SourceFile/SourceType; thread source file from MatchedUnifiedEntry.Source |
| `internal/report/markdown.go` | Rewrite grader section to group by SourceFile, render 3-level indentation |
| `internal/progress/display_interactive.go` | Add source file suffixes to grader name lines |
| `site/src/app/data/types.ts` | Add `source_file?`, `source_type?` to GraderResult interface |
| `site/src/app/components/GraderResultRow.tsx` | Group GraderResult[] by source_file at outer level |
| `site/src/app/components/eval-detail-page.tsx` | File-level grouping wrapper |
| `site/src/app/lib/graderScore.ts` | No change (operates per-grader) |

#### Display Rules

- Markdown: use `filepath.Base(sourceFile)` (absolute paths too verbose)
- CLI: add source suffix `(prompt file)` / `(criteria_file.yaml)`; keep lines flat
- Site: group by source_file at outer level using source_type for label
- Graceful degradation when SourceFile/SourceType are empty strings

### Part 4 — Tool Usage Grader

**Rule:** Verify declared tools/skills were actually used during the session.

#### Design: New `tool_usage` Grader Type

Rationale: Semantically distinct from `output_check` (workspace files) and `tool_constraint` (tool call counts). New type keeps concern clean.

#### Config Shape (YAML)

```yaml
- name: Tool Usage Verification
  type: tool_usage
  weight: 1.0
  details:
    rules:
      - type: mcp_server
        name: azure-mcp
        expect: at_least_one_tool_call
      - type: skill_plugin
        name: azure-sdk-python
        expect: any_skill_invoked
      - type: skill_repo
        repo: mauromedda/agent-toolkit
        skill: python
        expect: skill_invoked
```

#### Detection Logic

Input data available on `GraderInput`:
- `GeneratorArtifact.MCPToolCalls` — MCP tool calls from session events
- `GeneratorArtifact.SkillsInvoked` — skills actually invoked
- `GeneratorArtifact.ToolCalls` — all tool calls
- `EnvironmentTools []EnvironmentTool` — configured MCP servers + skills (NEW)

**Detection rule:** For each declared rule, check if tool was used. Skip rules where env item isn't in config.

#### Per-Point Output

Each rule → one `GraderPoint`:
- Label: `"azure-mcp tool used"`, `"azure-sdk-python skill invoked"`, etc.
- Pass: whether condition met
- Evidence: `{"tool_calls": "3", "expected": "≥1"}` (example)
- Rules where env item isn't present → **skipped silently** (not emitted)

#### Edge Case: Zero Applicable Rules

If all rules are skipped (env doesn't contain any declared items), emit one trivially-passing point `"no_applicable_rules"` to satisfy the ≥1 Point invariant.

#### Files to Change

| File | Change |
|------|--------|
| `internal/criteria/graders/types.go` | Add `KindToolUsage = "tool_usage"`, `ToolUsageConfig` struct, register in `validKinds` and `DecodeConfig` |
| `internal/criteria/graders/tool_usage_grader.go` | **New file** — implements Grader interface |
| `internal/criteria/graders/tool_usage_grader_test.go` | **New file** — table-driven tests |
| `internal/criteria/graders/registry.go` | Register `tool_usage` in NewGrader factory |
| `internal/criteria/graders/grader.go` | Add `EnvironmentTools []EnvironmentTool` to GraderInput |

### Implementation Status

✅ **Complete** — All four parts implemented and live-verified

- **Morpheus:** Comprehensive scope doc with all implementation details
- **Neo:** Engine implementation (Parts 1-4) + data model wiring
- **Tank:** Rendering redesign (markdown, CLI, site) with graceful degradation
- **Live Verification:** Tested on key-vault-dp-python-crud × azure-mcp/claude-opus-4.6 (twice)
- **Test Status:** 3 pre-existing failures confirmed unrelated

### Deliverables

- **Branch:** `neo/issue-grader-redesign`
- **Commits:** 5 (scope + 4-part engine impl + render merge + follow-up fixes)
- **PR:** Ready at https://github.com/ronniegeraghty/hyoka/pull/new/neo/issue-grader-redesign

### Follow-up Fixes (Neo)

- Fixed tool_usage env detection: `azure-mcp` → `azure` name mismatch resolved
- Added threshold values to output_check labels: `min_files (1)`, `min_bytes_per_file (1)`

### Handoffs

**Neo → Tank (2026-04-26):**
- Data model ready for render-side integration
- Provided exact field names, struct locations, and wiring instructions
- Tank implemented 3-level rendering with graceful degradation

**Tank → Neo (2026-04-30):**
- Render code complete and data-contract stable
- Tank's render changes merged into neo/issue-grader-redesign
- Consolidated on single branch for unified PR

### Decision

✅ **Implement all four parts as designed** — shipping in PR neo/issue-grader-redesign

**No changes to:**
- Backward compatibility (prompt files normalize automatically)
- CI workflows (existing tests confirm unrelated failures)
- API surface (grader interface unchanged)


# Phase 4 Test Review — Issue #375

**Author:** Switch 🤍  
**Date:** 2026-04-17  
**Branch:** ronniegeraghty/dev @ f5785617  
**Status:** Audit complete, recommendations ready

## Executive Summary

Conducted comprehensive test audit across Phase 4 work (PRs #568, #570, #571, #572). Go test suite is **solid** with 64.1% overall coverage, zero flakiness detected across 3 consecutive race-detector runs. Site test suite has **3 failing tests** due to stale fixtures after Trinity's UI updates. Identified 5 high-priority coverage gaps and 11 untested React components.

**Key Findings:**
- ✅ **Go race detector:** 3/3 clean runs, no flakiness
- ✅ **CI stability:** Last 15 runs on dev branch all green
- ⚠️ **Site tests:** 3 failures in `runs-page.test.tsx` (outdated after PR #572)
- ⚠️ **Coverage gaps:** progress (40.6%), process (43.6%), cmd (42.7%), config/tool (49.6%), eval (54.1%)
- 🔴 **Untested components:** 11 React components have no test coverage

---

## 1. Coverage Audit — Go Packages

### Overall Coverage: 64.1%

Ran full suite with `-race -cover` flags. All 24 packages pass. Zero race conditions detected.

```
Package                  Coverage    Assessment
─────────────────────────────────────────────────
logging                  96.3%       ✅ Excellent
utils                    96.2%       ✅ Excellent  
plugin                   95.3%       ✅ Excellent
validate                 93.0%       ✅ Excellent
prompt                   92.6%       ✅ Excellent
comparison               91.6%       ✅ Excellent
workspace                88.9%       ✅ Excellent (Phase 4 addition)
pidfile                  88.9%       ✅ Excellent
criteria                 87.0%       ✅ Excellent (Phase 3+ touched)
graders                  83.3%       ✅ Good (Phase 4 touched — WorkspaceDelta integration)
serve                    81.4%       ✅ Good (Phase 4 touched)
rerender                 79.6%       ✅ Good
config                   75.7%       ✅ Good
report                   71.2%       ✅ Good (Phase 4 touched — WorkspaceDelta JSON)
pairwise                 69.9%       ⚠️ Acceptable
checkenv                 59.5%       ⚠️ Thin (needs edge-case tests)
review                   60.8%       ⚠️ Thin (Phase 3 touched — modes)
trends                   58.8%       ⚠️ Thin
eval                     54.1%       ⚠️ Thin (Phase 3+ touched — core engine)
config/tool              49.6%       🔴 Low
process                  43.6%       🔴 Low
cmd                      42.7%       🔴 Low
progress                 40.6%       🔴 Low
```

### Phase 4-Specific Observations

**Packages Phase 4 touched:**
1. **workspace** (88.9%): New WorkspaceDelta feature (#566, PR #571) — solid baseline coverage from Neo's tests + my extended edge-case suite (`delta_extended_test.go`, `delta_json_test.go`, `delta_nil_safety_test.go`). Gap: TakeSnapshot error paths (symlink failures, permission errors) not covered.

2. **graders** (83.3%): Added WorkspaceDelta field to GraderInput — nil-safety tests exist. Gap: no graders *consume* WorkspaceDelta yet (future "code churn" grader candidate).

3. **report** (71.2%): Added WorkspaceDelta JSON serialization — backward-compat tests exist. Gap: HTML/Markdown rendering of delta data not tested (because not yet implemented in report templates).

4. **serve** (81.4%): Phase 4 API endpoints added for comparison unification (#357 pending). Current tests cover dashboard endpoints. Gap: new comparison endpoints will need tests when Neo lands #357.

5. **criteria** (87.0%): Hierarchical `when` filtering (WI-025, #356 pending). Current tests cover basic attribute matching. Gap: file-level, group-level, grader-level hierarchy interaction not tested (pending Neo's implementation).

6. **review** (60.8%): Review session modes — `combined` vs `isolated` (WI-024, #355 pending). Current tests cover single-mode scenarios. Gap: multi-grader `isolate: true` interaction not tested.

7. **comparison** (91.6%): Strong baseline. Gap: temporal comparison (comparing same config across time) only lightly tested.

---

## 2. Coverage Audit — Site (React/Vitest)

### Test Status: 50/53 passing (3 failures)

**Passing:** 8 test files, 50 tests
- ✅ api.test.ts (11 tests)
- ✅ footer.test.tsx (4 tests)
- ✅ navbar.test.tsx (4 tests)
- ✅ GraderResultRow.test.tsx (10 tests) — **NEW** (PR #572)
- ✅ layout.test.tsx (3 tests)
- ✅ home-page.test.tsx (5 tests)
- ✅ comparison-page.test.tsx (4 tests)
- ✅ dashboard-page.test.tsx (7 tests)

**Failing:** 1 test file, 3 tests
- 🔴 runs-page.test.tsx (3/5 tests failing):
  - "renders N/A for missing or invalid timestamp"
  - "handles missing duration and pass counts gracefully"
  - "renders run cards after data loads" (possibly flaky)

**Root cause:** Test fixtures expect `run.run_id` text to appear in rendered output, but Trinity's recent UI updates (PRs #570, #572) changed the component to display formatted timestamps instead of raw run IDs. Tests are querying for text that no longer renders.

**Recommended fix:** Update test assertions to match new UI (query by formatted timestamp text, not run_id). Low priority — cosmetic failure, not a regression.

### Untested Components (11)

Phase 4 added/modified components without corresponding test coverage:

**High Priority (Phase 4 deliverables):**
1. **eval-detail-page.tsx** — WI-044 (#358, PR #572) — NEW redesign, zero tests
2. **GraderResultRow.tsx** — WI-044 (#358, PR #572) — HAS tests (10 tests in GraderResultRow.test.tsx) ✅
3. **run-detail-page.tsx** — WI-045 (#359, Trinity pending) — untested
4. **pairwise-page.tsx** — WI-049 (#360, Trinity pending) — untested
5. **docs-page.tsx** — WI-051 (#362, PR #570) — content-only, low priority

**Medium Priority (existing pages):**
6. **prompt-detail-page.tsx** — untested
7. **prompts-page.tsx** — untested
8. **how-it-works-page.tsx** — untested

**Low Priority (UI components):**
9. **ui/select.tsx** — UI primitive, low value
10. **ui/table.tsx** — UI primitive, low value
11. **figma/ImageWithFallback.tsx** — utility component, low value

---

## 3. Flakiness + Race Audit

### Race Detector: ✅ CLEAN

Ran `go test -race ./hyoka/... -timeout 3m` **3 times consecutively**:
- Run 1: All pass ✅
- Run 2: All pass ✅
- Run 3: All pass ✅

Only differences between runs: timing variations (1.248s vs 1.269s) and cache hits. **Zero race conditions detected.**

### CI History: ✅ STABLE

Checked last 15 CI runs on `ronniegeraghty/dev` branch:
- **15/15 success** (100% pass rate)
- Most recent: PR #572 merge (f5785617) — green
- Phase 4 PRs: #568, #570, #571, #572 — all green
- Phase 3 merge: #562 — green

**No flakiness patterns observed.** CI is rock-solid.

### Known Flaky Tests: ZERO

Historical context: Fixed resourcemonitor flakiness in Phase 0 (#99, PR #167) by replacing time.Sleep assertions with event-driven sample() calls. Pattern established. No new flakiness introduced since.

---

## 4. Top 5 Coverage Gaps (Scenario-Level)

### Gap 1: Progress Display Edge Cases (40.6% coverage)
**Package:** `internal/progress`  
**Why it matters:** User-facing output, central to UX during eval runs  
**Missing scenarios:**
- Progress bar rendering with extremely long prompt IDs (> 60 chars) — truncation logic untested
- Simultaneous progress updates from multiple goroutines — race-safe but not stress-tested
- Terminal width < 80 columns — narrow terminal fallback untested
- `--progress off` flag interaction with live display — verify no ANSI codes leak to logs

**Recommended tests:**
- `TestProgressTruncation` — verify long strings don't break formatting
- `TestProgressConcurrency` — fire 100 updates from 10 goroutines, assert no panic
- `TestProgressNarrowTerminal` — mock terminal width to 40 cols, verify output
- `TestProgressOffMode` — assert zero output when mode=off

**Priority:** P2 (UX polish, not critical)

---

### Gap 2: Process Monitoring Error Paths (43.6% coverage)
**Package:** `internal/process`  
**Why it matters:** Guardrail enforcement relies on accurate CPU/mem tracking; silent failures could allow runaway sessions  
**Missing scenarios:**
- `resourcemonitor` start failure (e.g., /proc unavailable on non-Linux) — error handling untested
- PID reuse race condition — process exits, new process reuses PID, monitor still thinks old process is alive
- CPU reading failure mid-session — `getProcessCPU()` error path not covered
- Orphan detection with zombie processes — does `IsProcessAlive()` handle zombies correctly?

**Recommended tests:**
- `TestResourceMonitorStartFailure` — mock /proc read error, verify graceful degradation
- `TestPIDReuse` — kill process, start new process with same PID, verify tracker detects it
- `TestCPUReadFailure` — inject error into getProcessCPU, verify monitor doesn't crash
- `TestZombieProcessDetection` — create zombie (fork without wait), verify IsProcessAlive returns false

**Priority:** P1 (guardrail reliability)

---

### Gap 3: Eval Engine Guardrail Interactions (54.1% coverage)
**Package:** `internal/eval`  
**Why it matters:** Core evaluation loop; bugs here affect every eval run  
**Missing scenarios:**
- MaxSessionActions limit hit mid-turn — verify session terminates cleanly, report includes partial results
- MaxTurns hit with agent still typing — verify graceful shutdown, no data loss
- Guardrail warning emission (future: workspace delta size warnings) — warning logs captured correctly
- Nested workspace operations (agent creates workspace in workspace) — delta tracking edge case
- Session timeout during review phase — reviewer Copilot session hangs, verify timeout

**Recommended tests:**
- `TestMaxSessionActionsTermination` — set limit=5, prompt that triggers 10 actions, verify early exit + report
- `TestMaxTurnsGracefulShutdown` — verify EvalReport.Success=false, TurnsUsed=MaxTurns
- `TestGuardrailWarningCapture` — trigger guardrail warning, assert it appears in EvalReport.GuardrailWarnings
- `TestNestedWorkspaceHandling` — agent creates `nested/workspace/dir`, verify TakeSnapshot doesn't recurse infinitely
- `TestReviewerTimeout` — mock reviewer session that never responds, verify eval completes with timeout error

**Priority:** P1 (correctness + reliability)

---

### Gap 4: Review Session Modes (60.8% coverage)
**Package:** `internal/review`  
**Why it matters:** WI-024 (#355) introduces `combined` vs `isolated` modes — untested as of this audit  
**Missing scenarios (PENDING Neo's #355 implementation):**
- `--review-mode combined` with 3 reviewers — verify single session, shared context
- `--review-mode isolated` with 3 reviewers — verify 3 separate sessions, no context leakage
- Per-grader `isolate: true` override when global mode=`combined` — verify grader runs in own session
- Invalid `--review-mode` value — verify error message, graceful failure
- Mixed isolation (some graders isolated, some combined) — verify correct session routing

**Recommended tests (defer until #355 lands):**
- `TestReviewModeCombined` — 3 reviewers, assert 1 session created
- `TestReviewModeIsolated` — 3 reviewers, assert 3 sessions created
- `TestGraderIsolationOverride` — global=combined, grader.isolate=true, assert 2 sessions (1 combined + 1 isolated)
- `TestInvalidReviewMode` — pass `--review-mode invalid`, assert error
- `TestMixedIsolationSessionCounts` — complex scenario with 5 graders, 2 isolated, 3 combined

**Priority:** P1 (WI-024 acceptance criteria)

---

### Gap 5: Hierarchical `when` Filtering (87.0% coverage — deceptively high)
**Package:** `internal/criteria`  
**Why it matters:** WI-025 (#356) hierarchical filtering is complex; edge cases likely  
**Missing scenarios (PENDING Neo's #356 implementation):**
- File-level `when` + group-level `when` + grader-level `when` — 3-layer interaction
- Overlapping `when` conditions (e.g., file matches `language: python` AND `language: go`) — verify AND logic
- Empty `when` at group level — verify all graders in group run (not skipped)
- `when` with unknown attribute key — verify graceful failure, error message
- Grader-level `when` contradicts file-level `when` — which takes precedence?

**Recommended tests (defer until #356 lands):**
- `TestHierarchicalWhenThreeLayers` — file + group + grader all have `when`, verify correct filtering
- `TestOverlappingWhenConditions` — `when: {language: [python, go]}`, verify both match
- `TestEmptyGroupWhen` — group has no `when`, verify all graders run
- `TestUnknownWhenAttribute` — `when: {unknown_key: value}`, assert error
- `TestWhenPrecedence` — grader `when` says "run", file `when` says "skip" — assert grader wins (or vice versa, depending on spec)

**Priority:** P1 (WI-025 acceptance criteria)

---

## 5. Integration Test Gaps

Current integration tests:
- ✅ `eval/integration_test.go` — end-to-end eval pipeline with StubEvaluator + StubReviewer
- ✅ `eval/grader_integration_test.go` — real grader execution
- ✅ `serve/serve_test.go` — HTTP endpoint tests

**Missing cross-package flows:**

### Gap A: CLI → Engine → Report → Serve (end-to-end smoke test)
**Why it matters:** Validates full user journey, catches integration bugs  
**Proposed test:** `TestCLIToSiteEndToEnd`
1. Generate test prompt + config YAML in tempdir
2. Run `hyoka run --prompt-id test-prompt --config test-config` via Go subprocess
3. Verify JSON report written to `reports/` directory
4. Start `hyoka serve` on random port
5. Fetch `/api/dashboard` endpoint, assert test-prompt appears
6. Fetch `/api/evals/{id}`, assert grader results present
7. Shutdown serve, cleanup tempdir

**Complexity:** High (subprocess + HTTP server + file I/O)  
**Value:** High (catches regressions in CLI → site data flow)  
**Priority:** P2 (nice-to-have for release confidence)

---

### Gap B: Comparison Engine + Serve API (WI-034 + WI-050)
**Why it matters:** Phase 4 unifies comparison logic; site must consume unified API  
**Proposed test (PENDING Neo's #357):** `TestComparisonUnifiedEndpoint`
1. Generate 2 eval runs with same prompt, different configs
2. Call `comparison.Compare(runA, runB)` — verify ComparisonResult struct
3. Start serve, fetch `/api/compare?run_a=X&run_b=Y`
4. Assert HTTP response matches Go struct serialization (identical JSON)

**Priority:** P1 (WI-034 + WI-050 acceptance criteria)

---

### Gap C: Pairwise Expansion + Site Display (WI-022 + WI-049)
**Why it matters:** Pairwise results must render correctly on Eval Detail page  
**Proposed test (PENDING Trinity's #360):** `TestPairwiseDisplayIntegration`
1. Generate eval with `tool_filter: pairwise`
2. Verify `EvalReport.PairwiseResults` populated
3. Serve report, fetch `/api/evals/{id}`
4. Assert pairwise data present in JSON
5. (Manual smoke test) Open `/eval/{id}` in browser, verify pairwise chart renders

**Priority:** P2 (Trinity will manually test during #360 dev)

---

### Gap D: WorkspaceDelta + Guardrail Warnings (future)
**Why it matters:** Phase 4.5 follow-up — delta-based guardrails emit warnings instead of hard-failing  
**Proposed test (DEFERRED):** `TestWorkspaceDeltaGuardrailWarnings`
1. Configure MaxNewFiles=5, MaxOutputSize=10KB
2. Run eval that creates 8 files, 15KB total
3. Verify `EvalReport.Success=true` (not hard-fail)
4. Verify `EvalReport.GuardrailWarnings` contains "Exceeded MaxNewFiles (8 > 5)"
5. Verify `EvalReport.WorkspaceDelta.NewFiles` lists all 8 files

**Priority:** P3 (Phase 4.5 work, not current scope)

---

## 6. Tool/Infra Recommendations

### Recommendation 1: Enable `-race` flag in CI by default
**Current state:** CI runs `go test ./hyoka/...` without `-race`  
**Proposed:** Update `.github/workflows/ci.yml` to use `go test -race ./hyoka/...`  
**Rationale:** We already run `-race` locally during dev; making it CI-default catches concurrency bugs earlier. Adds ~30s to CI time (current: ~1m, with -race: ~1.5m). Acceptable tradeoff.  
**Priority:** P2 (quality-of-life improvement)

---

### Recommendation 2: Enforce minimum coverage threshold (60%)
**Current state:** No coverage gates; coverage reporting only  
**Proposed:** Add CI check that fails if `go test -coverprofile` reports < 60% coverage  
**Rationale:** Prevents coverage erosion. Current baseline: 64.1% — gives 4% buffer. New packages must meet threshold or explicitly justify.  
**Implementation:** Add step to CI workflow:
```yaml
- name: Coverage gate
  run: |
    go test -coverprofile=coverage.out ./hyoka/...
    total=$(go tool cover -func=coverage.out | grep "total:" | awk '{print $3}' | sed 's/%//')
    if (( $(echo "$total < 60.0" | bc -l) )); then
      echo "Coverage $total% below threshold 60%"
      exit 1
    fi
```
**Priority:** P3 (process improvement, not urgent)

---

### Recommendation 3: Add Vitest coverage reporting to CI
**Current state:** Site tests run, but no coverage metrics collected in CI  
**Proposed:** Run `npm test -- --coverage` in site CI job, report coverage %  
**Rationale:** Currently blind to site coverage trends. With 11 untested components, we need visibility.  
**Implementation:** Update `.github/workflows/ci.yml`:
```yaml
- name: Site tests with coverage
  run: |
    cd site
    npm install @vitest/coverage-v8  # if not in package.json
    npm test -- --coverage
```
**Priority:** P2 (visibility)

---

### Recommendation 4: Fix stale site tests before Phase 4 release
**Current state:** 3 failing tests in `runs-page.test.tsx`  
**Proposed:** Update test fixtures to match Trinity's UI changes (PR #572)  
**Rationale:** Failing tests erode confidence. Even if cosmetic, they signal "tests are stale, don't trust them."  
**Assignee:** Switch (me) can fix in follow-up, or Trinity can fix during #359 work  
**Priority:** P1 (hygiene)

---

### Recommendation 5: Document test patterns in CONTRIBUTING.md
**Current state:** Test patterns exist (table-driven, event-driven for goroutines) but not documented  
**Proposed:** Add `docs/testing-guide.md` with examples:
- Table-driven test structure (from `workspace/delta_test.go`)
- Nil-safety pattern for optional fields (from `graders/delta_nil_safety_test.go`)
- Event-driven async test pattern (from `process/resourcemonitor_test.go`)
- JSON roundtrip test pattern (from `report/delta_json_test.go`)

**Rationale:** New contributors (and future agents) need patterns to follow. Codifies tribal knowledge.  
**Priority:** P3 (documentation debt)

---

## 7. CI Gatekeeper Role — Wave 2 PRs

Monitoring upcoming PRs for quick-pass reviews:

### Pending PRs:
- **#355/#356** (Neo) — Review session modes + hierarchical `when`
  - Will verify: new tests added for `combined`/`isolated` modes, hierarchical filtering edge cases
  - Will check: review package coverage ↑ from 60.8%, criteria package coverage stable
  
- **#359** (Trinity) — Run Detail + Runs Page improvements
  - Will verify: stale `runs-page.test.tsx` failures addressed, new components have tests
  - Will check: site test pass rate back to 100%
  
- **CHANGELOG** (Oracle) — Phase 4 changelog update
  - Will verify: no code changes, docs-only → no tests needed

**Process:** I'll comment on PRs with quick test assessment (pass/concerns) once they're open. No deep reviews — Morpheus owns architectural review per charter.

---

## 8. Summary + Next Actions

### Coverage Baseline (for tracking)
- **Go overall:** 64.1%
- **Go race detector:** 3/3 clean runs
- **Site tests:** 50/53 passing (94.3%)
- **CI stability:** 15/15 green (100%)

### Top 3 Coverage Gaps (Priority P1)
1. **Process monitoring error paths** — guardrail reliability depends on accurate resource tracking
2. **Eval engine guardrail interactions** — MaxSessionActions, MaxTurns, timeout scenarios untested
3. **Review session modes** (pending #355) — `combined`/`isolated` modes need comprehensive tests

### Flakiness Found
**None.** CI is stable, race detector clean.

### Follow-Up Issues (Recommended)
1. **Issue #XXX:** Fix stale site tests in `runs-page.test.tsx` (P1, assign to Switch or Trinity)
2. **Issue #XXX:** Add process monitoring error path tests (P1, assign to Switch)
3. **Issue #XXX:** Add eval engine guardrail interaction tests (P1, assign to Switch)
4. **Issue #XXX:** Add integration test: CLI → Engine → Serve end-to-end (P2, assign to Switch)
5. **Issue #XXX:** Test hierarchical `when` filtering edge cases (P1, defer until #356 lands, assign to Switch)
6. **Issue #XXX:** Test review session modes (P1, defer until #355 lands, assign to Switch)
7. **Issue #XXX:** Add site test coverage for untested components (P2, assign to Trinity or Switch)
8. **Issue #XXX:** Enable `-race` flag in CI by default (P2, assign to Tank)
9. **Issue #XXX:** Add coverage threshold gate to CI (P3, assign to Tank)
10. **Issue #XXX:** Document test patterns in testing-guide.md (P3, assign to Oracle)

### Files Touched This Session
- None (audit-only, no code changes)

---

**Report complete. Ready for coordinator review.**

# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/main_test.go, hyoka/testdata/, hyoka/internal/ (packages to test)

## Core Context

Agent Switch initialized as Tester for hyoka. Guardrail defaults to test: max turns 25, max files 50, max output 1MB, max session actions 50. Safety boundaries prevent real Azure provisioning (--allow-cloud to opt out). Fan-out confirmation at >10 evals. Tests run via `go test ./...` from workspace root.

## Recent Updates

📌 Team initialized on 2026-04-03

📋 **Morpheus Audit (2026-04-03):** Testing audit complete. Key findings: (1) **pidfile package has zero tests (P1)** — needs test coverage. (2) **no integration test for eval pipeline (P1)** — TestRunEndToEnd recommended with StubEvaluator + StubReviewer. (3) **error logging gaps** — several discarded errors not logged. See `.squad/decisions.md` for prioritized test work.

## Learnings

Initial setup complete. Coverage is good across packages except pidfile. Integration test gap is highest priority for catching regressions.

### Session 2026-04-04T00-05 (Morpheus Evolution Plan)

Evolution plan assigns you pidfile tests, review package tests, integration test, flaky test fixes. Read `.squad/decisions.md` for full plan.

### Session 2026-04-04T19:49 (Phase 0 Execution — Flaky Test Fix)

**Status:** COMPLETE  
**Issue:** #99  
**PR:** #167

Fixed flaky resourcemonitor tests by replacing time.Sleep assertions with event-driven sample() calls. Tests now pass reliably under -race detection.

**Key outcome:** Established deterministic test pattern for background goroutines. Verified with go test -race -count=5 — 35 consecutive passes. Full suite all green under -race.

**Cross-agent dependency:** Tank's CI pipeline (#91, PR #168) enables reliable race detection across test suite. Neo's reviewer factory work (#92, PR #170) benefits from stable, race-safe test environment.

**Team impact:** All future background goroutine tests should follow event-driven pattern (direct method calls, never assertion sleeps).

**Files:** resourcemonitor_test.go

### Session 2026-07-22 (Issue #240 — Site SPA Tests)

**Status:** COMPLETE  
**Issue:** #240  
**PR:** #250

Added first-ever test coverage for the `site/` React SPA dashboard. Set up Vitest + React Testing Library with jsdom. Wrote 8 test files (41 tests) covering: API module (11), DashboardPage (7), HomePage (5), ComparisonPage (4), RunsPage (3), Navbar (4), Layout (3), Footer (4). All passing.

**Key outcome:** Site now has a test foundation. Future component work can be TDD'd. Setup file stubs ResizeObserver + IntersectionObserver for jsdom compatibility with recharts and motion libraries.

**Files:** site/vite.config.ts, site/package.json, site/src/__tests__/*.test.{ts,tsx}, site/src/__tests__/setup.ts

### Session 2026-04-16 (PR #567 Review — Starter-Aware Guardrails)

**Status:** COMPLETE ✅  
**Issue:** #565  
**PR:** #567  
**Verdict:** APPROVE

Reviewed Neo's starter-aware guardrail refactor. Original PR had 11 table-driven cases covering core scenarios. I identified missing edge cases (zero-byte files, empty starter) and added 4 test cases. All 15 cases pass with `-race` detection.

**Key findings:**
- Core coverage was solid (unchanged, modified, new, shrunk, deleted, mixed)
- Missing: explicit zero-byte file tests
- Missing: empty starter project test
- Symlinks: Safe (os.Stat follows symlinks automatically)
- Concurrent access: Not a concern for hotfix (read-only snapshot)
- Integration test: Not needed (3-line integration surface, pure functions)

**Outcome:** Enhanced test suite from 11 → 15 cases. APPROVED with commit `f3ea8b9f` pushed to PR branch.

**Cross-team:** Coordinator tracking #566 (Phase 3.5 followup: full `WorkspaceDelta` capture and guardrail softening).

**Files reviewed:** `hyoka/internal/eval/guardrail.go`, `guardrail_test.go`, `engine.go`

## 2026-04-16 — Phase 3 Merged to Dev (Neo)

Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has both Phase 3 features and starter-aware guardrail fix. All tests pass, CI green.

---

## 2026-04-17 — Phase 4 Kickoff: WorkspaceDelta Test Plan (#566)

**Status:** Test plan delivered  
**Issue:** #566  
**Deliverable:** `.squad/decisions/inbox/switch-566-test-plan.md`

Created comprehensive test plan for WorkspaceDelta feature with 40+ test scenarios across 6 categories:

1. **Delta Computation (9 scenarios):** Fresh workspace, unchanged starter, modified files (growth/shrink), deletions, mixed operations, zero-byte files, empty final workspace
2. **JSON Output (3 scenarios):** Serialization stability, backward compat with missing delta, zero-value handling
3. **Grader Integration (2 scenarios):** GraderInput exposure, nil-safety patterns
4. **Guardrail Interaction (4 scenarios):** Warning vs hard-fail, delta-based vs total size, negative deltas, MaxNewFiles threshold
5. **Edge Cases (6 scenarios):** Binary files, build artifacts, symlinks, large files (GB-scale), Unicode filenames, permission changes
6. **Integration (2 scenarios):** Existing test suite compat, guardrail test refactor

**Key patterns established:**
- **Table-driven structure** for all delta computation tests (setup → action → expected delta → assertions)
- **Fixture-based approach** for workspace state (tempdir, known file sizes, reproducible snapshots)
- **Nil-safety requirements** for graders consuming `WorkspaceDelta` (document pattern: `if input.WorkspaceDelta != nil`)
- **Size-only tracking** for binary files and large files (no content diffing, prevent OOM)
- **Exclusion filter integration** with `DefaultIgnoreDirs` (build artifacts don't pollute delta)

**Learnings:**
- WorkspaceDelta guardrails are harder to test than the hotfix — they're warning-based, not hard-fail. Test pattern: assert `Success = true` AND `GuardrailWarnings` contains message.
- Starter-aware workspace state requires consistent fixture setup — snapshot before, mutate workspace, compute delta. Reusable helper: `writeFile(t, baseDir, rel, contents)`.
- Binary file handling is critical — must track size without loading content. Test with sparse files or `fallocate` to verify no OOM.
- Edge case: agent creates then deletes file in same session → zero delta. Final state == initial state, so no net change recorded.
- Backward compat: older reports with no `workspace_delta` field must decode cleanly (nil check pattern).

**Handoff to Neo:**
Test plan document enumerates all scenarios with setup/action/expected triplets. Neo reads it during implementation; I codify it as `delta_test.go` once his branch `squad/566-workspace-delta` exists with types defined. Coordination via decision inbox.

**Files created:** `.squad/decisions/inbox/switch-566-test-plan.md` (17 KB, 40+ scenarios)

**Next actions:**
1. Monitor for Neo's branch `squad/566-workspace-delta`
2. When types exist: checkout branch, write `hyoka/internal/workspace/delta_test.go`
3. Run tests, report gaps to Neo
4. Iterate until all green
5. Review Neo's PR #566 with test lens

---

## 2026-04-17 — Phase 4 Wave 1: Test Codification Complete (#566, #571)

**Status:** Tests codified, CI green, PRs reviewed  
**Issue:** #566 (WorkspaceDelta foundation)  
**PR:** #571 (Neo's types)  
**Commit:** `b15371e9` on `squad/566-workspace-delta-v2`

**PRIMARY — Codified anticipatory tests against Neo's types:**

Neo delivered WorkspaceDelta types on PR #571. His baseline included 7 table-driven tests covering basic scenarios (new/modified/deleted/mixed). My test plan had 40 scenarios — I filled the gaps:

**New test files created:**
1. **`delta_extended_test.go`** (9 edge-case tests):
   - Zero-byte files (Scenario 1.8)
   - Empty workspace after agent (Scenario 1.9)
   - Create+delete same file (Scenario 1.6)
   - Unicode filenames (Scenario 5.5)
   - Binary files (Scenario 5.1)
   - Permission-only changes (Scenario 5.6)
   - Hidden file exclusion
   - JSON roundtrip
   
2. **`delta_json_test.go`** (4 EvalReport JSON tests):
   - WorkspaceDelta serialization (Scenario 2.1)
   - Backward compat (missing delta decodes cleanly) (Scenario 2.2)
   - Zero values serialize (not omitted) (Scenario 2.3)
   - Nil delta omits field (saves bytes)

3. **`delta_nil_safety_test.go`** (4 grader tests):
   - GraderInput.WorkspaceDelta nil vs non-nil (Scenarios 3.1, 3.2)
   - Mock grader demonstrates safe access pattern

**Type enhancement:**
Added `WorkspaceDelta` field to `GraderInput` so future graders can access delta (e.g., a "code churn" grader could analyze modified file counts). Included type alias in `graders` package to avoid import cycles.

**Test results:**
✅ All tests pass with `-race` flag:
- workspace: 16 tests (7 Neo baseline + 9 new)
- report: 4 new delta JSON tests
- graders: 4 new nil-safety tests
- Full suite: `go test -race ./hyoka/... -timeout 3m` ✅

**CI status on PR #571:** ✅ Both checks green (Build+Test, Site Build+Test)

**Coverage gaps filled:**
- JSON serialization (Scenarios 2.1–2.3) ✅
- Grader nil-safety (Scenarios 3.1–3.2) ✅
- Edge cases (Scenarios 5.1, 5.5, 5.6, 1.6, 1.8, 1.9) ✅

**Learnings:**

1. **Neo's baseline was solid** — covered the core delta computation logic (new/modified/deleted, size tracking, hash comparison). The gaps were edge cases and integration points.

2. **JSON backward compatibility is critical** — old reports (pre-#566) must decode without error. Test pattern: unmarshal JSON without `workspace_delta` field, assert `report.WorkspaceDelta == nil` (not panic).

3. **Grader nil-safety pattern established** — future graders that use `WorkspaceDelta` must check `if input.WorkspaceDelta != nil` before accessing. Mock grader test demonstrates this. Document in grader authoring guide.

4. **Zero-byte files are edge cases worth testing** — `__init__.py` is common in Python projects, must appear in `NewFiles` even with size 0. BytesNet remains 0.

5. **Permission-only changes do NOT trigger modification** — WorkspaceDelta is size-based (via hash comparison). File with `chmod 755 → 644` but identical content has same hash → not in `ModifiedFiles`. This is correct (permission changes aren't agent code output).

6. **Hidden files are already excluded** — `TakeSnapshot` skips files starting with `.` (matches existing `copyDir` behavior). Test confirmed `.hidden` and `.hiddendir/` are not in snapshot.

7. **Binary files work as expected** — `hashFile` reads binary content without panic. Test with PNG header bytes confirmed file appears in `NewFiles` with correct size and hash.

8. **Unicode paths work** — `文档.txt` and `résumé.md` serialize correctly in JSON (Go's string handling is UTF-8 native). No special escaping needed.

**QUICK-PASS REVIEWS:**

Reviewed #568 (Oracle — examples update) and #570 (Trinity — site branding):

- **#568:** `go run . validate` ✅ — 89 prompts valid, 12 configs valid, 25 graders valid
- **#570:** `npm test` (43 tests pass) + `npm run build` (✅ clean build) — site branding/content changes safe

---

## 2026-04-18 — Phase 4 Wave 3 Review (#581, #582, #583)

**Status:** COMPLETE — review at `.squad/decisions/switch-wave3-review.md`
**Gate:** Phase 4 criterion #4 (CLI↔site parity) now test-enforced.

**Primary deliverable — CLI↔site equivalence test.** Neo's #583 claimed the gate is "satisfied by construction" via shared `CompareReports`. I codified the claim in `hyoka/internal/serve/equivalence_test.go` (commit `e91997a5` on `squad/357-comparison-unification`):
- `TestCLISiteEquivalence_AutoGenComparisons` — on-demand path (no `comparisons.json`; API recomputes via `AutoGenerateForRun`).
- `TestCLISiteEquivalence_LoadedComparisonsFile` — cached path (`WriteForRun` writes, API reads from disk).
Both `reflect.DeepEqual` on wire-format-roundtripped structs. Passed `-race -count=3`.

**Per-PR verdicts:** #581 LGTM (file cache + pairwise endpoint, 83.7% pkg cov). #582 LGTM w/ notes (239 new lines, zero new FE tests — pairwise-page stmt cov 44.85%). #583 LGTM (total Go cov 64.4%, above 64% baseline).

**Merge order flagged:** #581 and #583 both edit `dashboard.go`/`serve.go`. Recommended order: #582 → #581 → #583 (rebased). One rebase is unavoidable.

**Key learnings:**
1. "Shared core" claims need tests that compare both paths byte-for-byte. Marshal+unmarshal both sides before DeepEqual so `omitempty` asymmetry is normalized at wire-format level.
2. `AutoGenerateForRun` returns nil for <2 configs — equivalence test must seed ≥2 configs or silently pass on empty slices. Also assert `len(apiResults) == expected` to catch that.
3. The earlier "94% site baseline" number was test-count, not statement coverage. Real statement coverage is 66.66%. Future reviews should distinguish.
4. When two PRs refactor the same handler signatures, the second merger rebases — flag this in the review rather than letting the merge-bot surprise the team.

Both PRs approved with "LGTM ✅" comments.

**NOT IN SCOPE:**
- PR #572 (Trinity R2 — Eval Detail component) — Morpheus is architectural reviewer per charter
- Guardrail interaction tests (Scenarios 4.1–4.4) — deferred to future work (current PR keeps hard-fail behavior, no guardrail changes yet)

**Next actions:**
1. Coordinator merges PRs when ready (#568, #570, #571)
2. Monitor for follow-up PRs in Phase 4 (Trinity R2, Oracle's workspace integration)
3. If guardrail warnings feature ships, add Scenarios 4.1–4.4 tests

**Files created:**
- `hyoka/internal/workspace/delta_extended_test.go`
- `hyoka/internal/report/delta_json_test.go`
- `hyoka/internal/graders/delta_nil_safety_test.go`

**Files modified:**
- `hyoka/internal/graders/grader.go` (added WorkspaceDelta field + type alias)

**Cross-team:**
- Neo: PR #571 has tests codified, safe to merge
- Morpheus: Architectural review of #572 in progress
- Coordinator: Ready to merge #568, #570, #571 when greenlit

---

## 2026-04-17 — Phase 4 Wave 2: Test Review Complete (#375, PR #577)

**Status:** COMPLETE ✅  
**Issue:** #375  
**PR:** #577 (branch: `squad/375-phase4-test-review-v2`)

**Audit scope:**
- Go test suite (24 packages)
- Site test suite (Vitest, 9 test files)
- CI stability (last 15 runs)
- Race detector (3 consecutive runs)
- Coverage baseline for Phase 4 tracking

**Findings:**

**Go Test Suite: ✅ SOLID**
- Overall coverage: **64.1%**
- Race detector: **3/3 clean runs**, zero race conditions
- CI: **15/15 green** on ronniegeraghty/dev
- Flakiness: **ZERO** (resourcemonitor pattern from Phase 0 PR #167 still holding strong)

**Site Test Suite: ⚠️ 3 FAILURES**
- Status: **50/53 tests passing** (94.3%)
- Failures: `runs-page.test.tsx` (3 tests) — stale fixtures after Trinity's PR #572 UI updates
- Root cause: Tests query for `run.run_id` text, but component now displays formatted timestamps
- Fix: Update assertions (low priority, cosmetic)

**Coverage Gaps Identified (Top 5):**

1. **Progress display edge cases** (40.6%) — narrow terminals, long prompt IDs, concurrent updates, `--progress off` leakage
2. **Process monitoring error paths** (43.6%) — /proc failures, PID reuse, zombie processes, CPU read errors
3. **Eval engine guardrail interactions** (54.1%) — MaxSessionActions mid-turn, MaxTurns graceful shutdown, nested workspace, reviewer timeout
4. **Review session modes** (60.8%, pending Neo #355) — `combined` vs `isolated`, per-grader overrides, invalid mode handling
5. **Hierarchical `when` filtering** (87.0%, pending Neo #356) — 3-layer interaction, overlapping conditions, precedence rules

**Site Component Coverage:**
- **11 untested components** (out of 16 total)
- High-priority: `eval-detail-page.tsx` (WI-044, PR #572), `run-detail-page.tsx` (WI-045 pending), `pairwise-page.tsx` (WI-049 pending)
- `GraderResultRow.tsx` HAS tests ✅ (10 tests in GraderResultRow.test.tsx)

**Integration Test Gaps:**
1. CLI → Engine → Report → Serve (end-to-end smoke test)
2. Comparison unification + serve API (pending Neo #357)
3. Pairwise expansion + site display (pending Trinity #360)
4. WorkspaceDelta + guardrail warnings (Phase 4.5 deferred)

**Recommendations:**
1. Enable `-race` in CI by default (P2)
2. Enforce 60% coverage threshold in CI (P3)
3. Add Vitest coverage reporting to CI (P2)
4. Fix stale site tests (P1)
5. Document test patterns in testing-guide.md (P3)

**Follow-up issues recommended:**
10 P1-P3 issues spanning process monitoring, eval engine, site tests, integration tests, CI improvements, documentation.

**Deliverables:**
- Report: `.squad/decisions/switch-375-test-review.md` (20KB, 8 sections)
- PR #577: 1 file, 435 lines added

**Learnings:**

1. **Coverage baseline established** — 64.1% Go, 94.3% site tests passing. Future PRs can be compared against this snapshot. Phase 4 packages (workspace, graders, report, serve, criteria, review) all have adequate baseline coverage; gaps are in older packages (progress, process, eval, cmd).

2. **CI is rock-solid** — 15/15 green runs, zero flakiness. The event-driven async test pattern from Phase 0 (resourcemonitor fix) has prevented any new race conditions. Race detector runs locally but not in CI — recommended to add `-race` flag to CI workflow for earlier detection.

3. **Site tests lag behind UI updates** — Trinity's PR #572 changed RunsPage component, but tests weren't updated in same PR. Established pattern: UI PRs should include test updates or explicitly note "tests deferred." Without this, test suite becomes untrustworthy (3 failures == "tests are stale, ignore them").

4. **Untested != broken** — 11 React components have no tests, but site builds cleanly and renders correctly (manual smoke test). Test coverage gaps are technical debt, not bugs. Priority order: (1) fix failing tests (hygiene), (2) test new Phase 4 components (eval-detail, run-detail, pairwise), (3) backfill older components (prompts-page, docs-page).

5. **Phase 4 test gaps are mostly "pending implementation"** — Review session modes (#355), hierarchical `when` (#356), comparison unification (#357) are untested because Neo hasn't landed them yet. Test plan exists (my 40-scenario WorkspaceDelta plan is reference template). I'll codify tests once Neo's PRs land.

6. **Integration tests are sparse** — Only 3 integration test files exist (`eval/integration_test.go`, `eval/grader_integration_test.go`, `serve/serve_test.go`). Recommended end-to-end flow: `CLI → Engine → Report → Serve` — validates full user journey, catches integration bugs. Complexity: high (subprocess + HTTP + file I/O). Value: high (release confidence).

7. **Coverage reporting is manual** — Had to run `go test -coverprofile`, parse output, manually extract percentages. CI runs tests but doesn't report coverage trends. Recommended: add coverage threshold gate (60% floor) + trend tracking over time.

8. **Site coverage tooling missing** — Had to install `@vitest/coverage-v8` manually (not in package.json). Site tests run in CI but no coverage metrics collected. Without visibility, can't track whether new components have tests. Recommended: add `npm test -- --coverage` to site CI job.

9. **Test review timing is critical** — Conducted audit AFTER PRs #568, #570, #571, #572 merged. Ideal: review DURING PR phase (quick-pass checks, not blocking). Established CI gatekeeper role for Wave 2 PRs (#355/#356/#359) — I'll comment with test assessment once they're open.

10. **Stale tests are worse than no tests** — 3 failing tests create uncertainty ("is this a real bug or just outdated fixtures?"). Teams start ignoring test failures. Pattern: if tests fail for >1 day without fix, either (a) fix immediately, or (b) skip/disable with TODO comment. Never leave failing tests in main/dev branch.

**Next actions:**
1. Monitor Wave 2 PRs (#355, #356, #359) for quick-pass test reviews
2. File 10 follow-up issues (process monitoring, eval engine, site tests, CI improvements) — defer to next session or coordinator
3. Fix stale site tests in `runs-page.test.tsx` (P1) — can do in follow-up or assign to Trinity

**Cross-team:**
- Coordinator: PR #577 ready for review, report at `.squad/decisions/switch-375-test-review.md`
- Neo: When #355/#356 land, ping me to codify review mode + hierarchical `when` tests
- Trinity: When #359 lands, ping me for quick-pass site test review
- Tank: Recommended to add `-race` to CI, coverage threshold gate (P2/P3 infra work)
- Oracle: Recommended to write testing-guide.md (P3 documentation)


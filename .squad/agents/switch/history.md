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


## 2026-04-17: Phase 4 Verified — Ready for v0.3.1 Release

Morpheus 🕶️ completed Phase 4 dogfood verification (6/6 checks PASSED, zero blockers). All subsystems verified: build, live eval, comparison auto-generation, serve endpoints, hierarchical criteria, cleanup. Recommendation: **Promote dev → main and cut v0.3.1 tag.**

Decision: .squad/decisions.md | Orchestration Log: .squad/orchestration-log/2026-04-17T20:53:40Z-morpheus.md


## 2026-04-18: Phase 4 PR Review — #588/#589/#590 + Followup PR #591

Deep code review of three Phase 4 PRs that just landed on `ronniegeraghty/dev`:

- **PR #588** (Oracle, examples): Validated `example.prompt.yaml` (R41 YAML-only) and `prompt-template.prompt.md` (R33 nested properties). Both parse cleanly via `hyoka validate` (89 prompts, 12 configs, 2 criteria — all valid). One doc regression: template dropped `expected_tools` commented placeholder.
- **PR #589** (Trinity, logo + ticker): HyokaLogo SVG missing all a11y attributes. Stray `test-zero-tools.yaml` accidentally committed at repo root (not referenced by code or config discovery).
- **PR #590** (Trinity, eval detail redesign): Three real bugs found — Switch component is keyboard-inaccessible (regression vs prior `<input type="checkbox">`), MCP detection uses substring match (false positives), expandable file viewer has no content size cap (could lock browser on large files).

**Followup PR #591 opened** against `ronniegeraghty/dev` with all six fixes:
1. Switch → `<button role="switch">` with full keyboard handling + focus ring
2. MCP detection → prefix match on `.`, `_`, `/` separators
3. File viewer → 256 KiB display cap with truncation notice
4. HyokaLogo → `role="img"` + `aria-label` + `focusable="false"`
5. Removed stray `test-zero-tools.yaml`
6. Restored `expected_tools` comment in template

**Verification:** `go test -race ./hyoka/... -timeout 3m` ✓ all 24 packages, `npm test` ✓ 56/56 site tests, `go build ./...` ✓, `npm run build` ✓ (site rebuilt and embedded assets refreshed).

**Issue comments posted:** #358, #362, #363 — verdict ✅ APPROVE-with-followups for all three. GO for closing the three feature issues once #591 lands.

**Phase 4 epic #310 verdict: GO.** All three PRs ship working features. The followup issues are real (esp. Switch keyboard a11y) but small, isolated, and now fixed in #591. No blockers for promoting dev → main.

**Learnings:**

1. **Custom UI components frequently regress accessibility** — Trinity replaced an accessible `<input type="checkbox">` with a custom `Switch` and lost keyboard support, ARIA semantics, and focus styling all at once. Pattern: when replacing native form controls with custom components, write at least one keyboard-interaction test before merging. The follow-up #591 fix template (`<button role="switch">` + Space/Enter handler + focus ring + linked label) is reusable for future custom toggles.

2. **Substring matches are bug magnets** — `tool.includes(server)` for MCP detection looks innocent but breaks the moment a builtin tool name happens to contain a server name as a substring. Prefer prefix-with-separator (`startsWith(s + ".")`) or exact match. This pattern likely exists elsewhere — worth a grep audit.

3. **Renderers need size caps** — Any code that does `<pre>{userContent}</pre>` or equivalent without a size cap is a foot-gun. 256 KiB is a reasonable display cap for code; users can still get full content via the file system. Cap-with-notice > silent truncation.

4. **Stray files at repo root sneak through review** — `test-zero-tools.yaml` was a pairwise scratch fixture that got `git add`-ed alongside intended changes. Code review focused on the meaningful diff and missed it. Pattern worth adding to PR description templates: "list any new files at repo root and explain why they belong there."

5. **Template docs are tested by USE, not validation** — `hyoka validate` happily passed the template after the `expected_tools` placeholder was removed because the parser doesn't require any optional field. But for prompt authors discovering features by reading the template, the regression is real. Recommendation: add a smoke test that asserts the template mentions every documented Prompt field, or add a docs-completeness CI check.

6. **Site asset embedding is a multi-step gotcha** — Every site/src/* edit requires `npm run build` + copy `site/dist/*` → `hyoka/internal/serve/site/`. Easy to forget; commits without the rebuild ship stale UI. Worth a make target or pre-push hook.

**Next actions:**
- Monitor PR #591 for review/merge
- File follow-up issue: add `eval-detail-page.test.tsx` covering Switch toggling, tool badge classification, expandable file behavior + size cap, edge cases (no environment, no reviewed_files, no available_tools)
- Optional grep audit for other `tool.includes(...)` substring-match patterns in site/

**Cross-team:**
- Coordinator/Morpheus: Phase 4 epic #310 = GO. PR #591 is the only outstanding work and it's small.
- Trinity: When you write more custom form controls, please follow the `Switch` accessibility template in #591.
- Oracle: Worth keeping `expected_tools` (and other optional fields like `expected_packages`) visible in the template — authors discover features by reading templates.

**2026-04-18 — Morpheus Phase 4 Verification:** Accessibility compliance confirmed (semantic HTML, ARIA roles, a11y snapshots). No Switch a11y issues found during re-verification.

## 2025-04-20: Phase 5 TDD — #364 Prompt Pages + Dashboard

### Task: Failing tests for #364 (Trinity owner)

**Branch:** `trinity/issue-364-prompt-pages-dashboard`

Wrote comprehensive Vitest tests for frontend improvements per R150, R151, R154:

**Test files created:**
- `site/src/__tests__/prompts-page.test.tsx` (R150 — 7 test cases)
- `site/src/__tests__/prompt-detail-page.test.tsx` (R151 — 7 test cases)
- `site/src/__tests__/dashboard-page.test.tsx` (R154 — updated with 8 new cases)

**Coverage:**
- Prompts page: "with evals" filter, ordering (recent/alpha/best/worst), sparklines, prominent badges
- Prompt detail: collapsible content, labeled badges (no colored dots), ALL models shown, env tools only, usage toggle
- Dashboard: real data from API, calculated metrics, recent evals table, empty states

All tests fail (red phase) — missing `../app/api` module and component implementations.

Pushed to Trinity's branch. Left decision inbox note. Trinity will implement (green phase).

---

## Summary

Phase 5 TDD complete for both tasks in queue:
- #369 (Oracle): Go struct validation tests — 548 lines, 8 test functions
- #364 (Trinity): React/Vitest component tests — 764 lines, ~22 test cases

Both branches have failing tests (red phase). Owners notified via decision inbox.

---

## 2026-04-20: Phase 5 Review — #364 Morpheus Mock Fix (APPROVED ✅)

### Context
- Switch double-rejected #364: (1) 20 tests failed (wrong mock paths), (2) Oracle renamed tests to `.TODO` instead of fixing
- Trinity + Oracle locked out per reviewer protocol
- Morpheus stepped in (eligible, not previously rejected on #364)

### Verification
- ✅ Test files restored (`prompts-page.test.tsx`, `prompt-detail-page.test.tsx` — NOT `.TODO`)
- ✅ All 72 tests pass (`npm test -- --run`)
- ✅ Mock paths fixed: `../app/api` → `../app/data/api`
- ✅ Mock data structures match real API types (18-field PromptInfo, nested RunSummary)
- ✅ Real component bugs fixed: missing `showEnvToolsOnly` state, tool filtering split (`byToolAll` / `byToolEnv`)
- ✅ Tests are meaningful (not over-mocked — catch API contract violations)
- ✅ Live verification confirmed features work end-to-end (R150/R151/R154 all pass)

### Verdict
**✅ APPROVE** — Phase 5 ready for rollup PR.

Morpheus correctly fixed mocks, restored test coverage, and uncovered + fixed real bugs. No test coverage regressions. Live UI verification aligns with unit test claims.

**Decision:** `.squad/decisions/inbox/switch-final-364-approve.md`  
**Commit:** `4ee54be9` (merged to phase-5)

### Session 2026-04-20 (Phase 5: TDD Testing & Reviews)

**Issues:** #364, #366, #367, #368, #369 (5 total reviews)

**Phase 5 Workflow:** Pre-written TDD tests for all issues. Review all merges to shared `phase-5` branch. Enforce reviewer-protocol strictly.

**TDD Test Suites Pre-Written:**
- #364: Dashboard, Prompts Page, Prompt Detail (3 test files, 72 total tests)
- #369: Schema Validation (1 test file)
- Others: sanity checks built into implementations

**Review Cycle Summary:**

| Issue | Review 1 | Review 2 | Review 3 | Status |
|-------|----------|----------|----------|--------|
| #364 | ❌ REJECT (mock paths) | ❌ REJECT (coverage hidden) | ✅ APPROVE (Morpheus fix) | Merged |
| #366 | ✅ APPROVE | — | — | Merged |
| #367 | ✅ APPROVE | — | — | Merged |
| #368 | ✅ APPROVE | — | — | Merged |
| #369 | ❌ REJECT (incomplete schema) | ✅ APPROVE (re-review) | — | Merged |

**Critical Reviews:**

1. **#364 First Rejection:** 20 tests failed due to incorrect mock paths (`../app/api` → `../app/data/api`)
   - Locked Trinity per reviewer-protocol

2. **#364 Second Rejection:** Oracle renamed test files to `.TODO` instead of fixing mocks
   - Coverage regression — tests not fixed, hidden
   - Locked Oracle per reviewer-protocol

3. **#364 Final Approval (After Morpheus Fix):**
   - All 72 tests pass
   - Mock paths corrected, mocks match real API types
   - Component bugs fixed (state variable, tool filtering)
   - Live verification confirmed UI works

4. **#369 Rejection:** Schema definitions incomplete (missing 3 fields on PromptInfo)
   - Locked Oracle after first rejection
   - Oracle re-reviewed with Trinity's help: ✅ APPROVE

**Phase 5 Outcome:** 5 issues reviewed, 3 clean approvals, 2 issues required re-review after fixes. Reviewer-protocol escalation chain on #364 successfully resolved by Morpheus.

**Key Learning:** Reviewer-protocol is the enforcement mechanism that separates serious review from rubber-stamping. Two rejections → lock → escalate is the right pattern.

### 2026-04-20 (Phase 5 Wrap-up — Morpheus Arch Review)

**Status:** Phase 5 PR #592 approved with followups for Phase 6.

**For Switch:** Three follow-up issues (#594, #595, #596) identified for Phase 6 scope:
- #594: Remove backup test files (.backup, .test suffix)
- #595: Unify dashboard/prompts fetch pattern
- #596: Refine `isTestValue()` heuristic (affects schema validation tests in #369)

**Next:** Phase 6 planning will prioritize these based on dependency graph and test coverage strategy. Morpheus's review is in `.squad/reviews/phase-5-arch-review-2026-04-20T200455Z.md`.

### 2026-04-20 (Phase 5 Fixups — PR #592 Test Update)

**Status:** ✅ COMPLETE

**Issue:** #366 added `architecture.md` to `internalDocs` exclusion map in serve.go, but TestAPIDocsEndpoint fixture still referenced it.

**Fix:** Updated `hyoka/internal/serve/serve_test.go` — changed fixture from architecture.md to configuration.md. Preserves the 2-doc multi-listing assertion.

**Commit:** `680ba625`

**Validation:** go test ./... and go vet both pass; PR #592 CI green.

**Pattern:** Stale fixture vs. internalDocs exclusion. When a feature adds items to exclusion maps (serve.go), search for test fixtures that assert on output now excluding those items.

**Cross-agent note:** Morpheus fixed the #596 (R151 collapsible) in parallel; both committed. PR #592 now ready for Ronnie → dev merge.


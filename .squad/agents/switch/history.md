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

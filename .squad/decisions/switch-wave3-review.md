# Wave 3 PR Review — #581 / #582 / #583

**Author:** Switch 🤍
**Date:** 2026-04-18
**Base:** `squad/phase-4-remainder`
**Gate ref:** `.squad/decisions/archive/morpheus-phase4-kickoff.md` §6

## Summary

| PR  | Title                           | Verdict           | Go cov (pkg) | Site tests |
|-----|---------------------------------|-------------------|--------------|------------|
| #581 | #361 serve cache + pairwise    | ✅ LGTM          | serve 83.7%  | 56/56 pass |
| #582 | #360 pairwise methodology/chart| 🟡 LGTM w/ notes | n/a (FE)     | 53/53 pass |
| #583 | #357 comparison unification    | ✅ LGTM          | comparison 91.6% · serve 80.0% · total 64.4% | — |

Go suite runs **race-clean 3×** on all three branches.

## Phase 4 Gate Criterion #4 — CLI↔site parity

**Status: ENFORCED IN TESTS.** Neo's "satisfied by construction" claim is now load-bearing rather than aspirational.

Added `hyoka/internal/serve/equivalence_test.go` to `squad/357-comparison-unification` — **commit `e91997a5`**. Two tests:

1. `TestCLISiteEquivalence_AutoGenComparisons` — forces on-demand path (no `comparisons.json` on disk) → the API recomputes via `AutoGenerateForRun`. Compares the engine's direct output to the JSON returned by `GET /api/runs/{runID}/comparisons` via `reflect.DeepEqual` on wire-format-roundtripped structs.
2. `TestCLISiteEquivalence_LoadedComparisonsFile` — uses `comparison.WriteForRun` (end-of-run path) → API reads from disk. Same equivalence assertion.

Both passed under `-race -count=3`. If the shared-core invariant ever regresses, these tests fail.

## Per-PR notes

### PR #581 — LGTM ✅
- File cache is correct: mtime+size fingerprint busts on modification (`TestFileCache_MtimeChangeBustsCache`) and survives file removal (`TestFileCache_ReadThenCacheHit`). Race-clean.
- Pairwise endpoint (`/api/runs/{runID}/pairwise`) has 404 coverage for missing runs.
- **Minor:** `registerDashboardRoutes` receives `cache` but currently only stashes it as `_ = cache`. Dead-flag acceptable as reserved capacity — comment says so explicitly.

### PR #582 — 🟡 LGTM with documented test gaps
- 239 lines added to `pairwise-page.tsx` (MethodologyInfo, ToolUsageFrequencyChart, `computeToolFrequency`) and 21 lines to `run-detail-page.tsx` (cross-link), **no new frontend tests**. Existing 53 tests still pass and the build is clean, so this isn't a blocker, but I'm flagging it.
- Site statement coverage on `pairwise-page.tsx` is **44.85%** (uncovered: 287–292, 501–565 — exactly the new code). The "94% site baseline" in my earlier review was test-count, not statement coverage; the actual all-files statement coverage is **66.66%**. The new code pushes this down slightly, not up.
- `computeToolFrequency` is pure and cheap to test — recommend a follow-up issue to add at least one table-driven test over it.
- Integration: #582 reads `pairwise_results` embedded in `RunSummary` (existing plumbing). It does **not** consume #581's new `/api/runs/{runID}/pairwise` endpoint. They are independent features — no wiring gap, but also no cross-PR synergy realized yet.

### PR #583 — LGTM ✅
- `ComparisonResult` is a clean unification; `CompareReports` is the correct seam. Existing tests (`inmem_test.go`, `autogen_test.go`) plus the new equivalence tests give end-to-end coverage.
- CLI refactor in `cmd/compare.go` is mostly deletions (−208) — good sign that shared-core claim is real.
- Total Go coverage **64.4%**, just above the 64% baseline. No regression.

## Cross-PR integration / merge-order

**Both #581 and #583 modify `hyoka/internal/serve/dashboard.go` and `serve.go` — non-trivial conflicts expected.**

- #581 renames `handleAPIGraders`, `handleAPITimeline`, `handleAPIScoreBreakdown` signatures to accept `cache *fileCache` and routes them through `registerDashboardRoutes(mux, opts, cache)`.
- #583 adds a new handler `handleAPIRunComparisons` and a `/api/runs/{runID}/comparisons` mux route.

**Recommended merge order:**

1. **#582 first** — frontend-only (site/ + trinity history). Zero overlap with #581/#583.
2. **#581 second** — cache refactor lands cleanly on `squad/phase-4-remainder`.
3. **#583 last, rebased onto #581** — Neo must rebase `squad/357-comparison-unification` after #581 merges, reapply `handleAPIRunComparisons` on top of the cache-aware registration. At that point the new handler should also take `*fileCache` for consistency (even if unused initially, like the current `_ = cache`).

Alternate: #583 before #581 also works (smaller rebase surface for Neo), but then #581 must adopt the cache pattern in `handleAPIRunComparisons` when it rebases. Either way, **one of the two will rebase.**

## Coverage check vs baseline

| Surface | Baseline | Wave 3 | Status |
|---------|----------|--------|--------|
| Go total | 64% | 64.4% (@#583) | ✅ no regression |
| Go comparison | (new) | 91.6% | ✅ healthy |
| Go serve | ~75% | 80.0% (#583) / 83.7% (#581) | ✅ improved |
| Site statement | 66.66% actual | 66.66% (#582) | ⚠ flat — new FE code uncovered |

## Flakiness

`go test -race -count=3` on each branch: **zero flakes** across serve, comparison, and full `./hyoka/...` suites.

## Action items

1. **Neo:** rebase onto merged #581; re-run equivalence tests post-merge on the combined base before dismissing the gate.
2. **Trinity (follow-up issue, non-blocking):** add a `pairwise-page.test.tsx` covering `computeToolFrequency` and rendering of the methodology explainer / tool usage chart. Target 60%+ stmt coverage on `pairwise-page.tsx`.
3. **Morpheus:** Phase 4 gate criterion #4 can be checked — test-enforced, not hand-waved.

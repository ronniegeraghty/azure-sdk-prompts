# Decision: PR #603 wiring-layer tests added (reviewer-lockout reassignment)

**Author:** Tank 📡
**Date:** 2026-04-21
**Status:** Complete
**PR:** #603 (`ronniegeraghty/issue-580-review-session-splitting`)
**Issue:** #580
**Commit:** 04579b47

## Context

Switch (Test Reviewer) issued REQUEST CHANGES on PR #603 because the implementation of `--review-mode isolated` (Neo's #580 work) had excellent coverage at the bucket-construction layer (14 tests in `criteria/buckets_test.go`) but **zero coverage** at the wiring layers above it. This is the exact failure mode that caused PR #587 to revert PR #355's first attempt at the same feature: tests pass, runtime behavior absent.

Per strict reviewer-lockout protocol, Neo (the PR author) was locked out of revisions. Tank was the named alternate.

## Decision

Pushed 16 new tests (22 subtests counting table cases) directly to Neo's PR branch, closing all four gaps Switch identified, plus an end-to-end runtime integration test as bonus regression armor against the #587 failure mode.

## What Was Added

| Surface | File | Coverage Closed |
|---|---|---|
| `Engine.reviewBuckets()` (eval/engine.go:275–290) | `internal/eval/engine_reviewbuckets_test.go` | 0% → 100% line coverage; combined / default-empty / isolated+isolation / isolated-degraded with `slog.Warn` capture / no-graders edge case |
| Engine→grader→Reviewer chain end-to-end | `internal/eval/engine_reviewmode_runtime_test.go` | Recording-reviewer integration tests through `engine.Run` proving `ReviewBuckets()` actually fires in isolated mode and `Review()` in combined — locks in "flag has runtime effect" |
| `(*PromptReviewGrader).gradeSingle` branch selection | `internal/graders/prompt_review_grader_buckets_test.go` | `len(buckets) ∈ {0, 1, 3}` table; non-MultiBucketReviewer fallback to `joinCriteria` |
| `mergeBucketResults` (review/buckets.go:156) | `internal/review/buckets_test.go` | Name-prefixing rules ("" / "combined" / named); cross-bucket aggregation; nil-part skip |
| CLI flag validation | `cmd/run_validate.go` (extracted helper) + `cmd/run_validate_test.go` | All six branches of `validateReviewMode` + cobra flag-registration & default check |

## Coverage Delta

- `Engine.reviewBuckets`: **0% → 100%** line coverage (per-line coverprofile)
- `internal/review`: 48.6% → **53.5%** (+4.9 pp)
- `internal/graders`: 79.9% → **82.9%** (+3.0 pp)
- `cmd`: 42.4% → **42.6%** (+0.2 pp)
- `internal/eval`: 54.5% → 54.5% at package level (eval/engine.go is 918 lines; the small `reviewBuckets` function fully exercised but doesn't move the package %)

## Gap Closed

Before this commit, a refactor that dropped `e.opts.ReviewMode` from `reviewBuckets()` (or removed the entire function) would compile cleanly and pass every unit test in the repo. The integration tests at `engine_reviewmode_runtime_test.go` lock in the runtime contract: with `ReviewMode: "isolated"` and an isolate-marked grader, the engine MUST invoke `Reviewer.ReviewBuckets()` with ≥2 buckets. The slog-capturing test at `engine_reviewbuckets_test.go::TestEngineReviewBuckets_IsolatedDegradesWithoutIsolation` locks in the "observably no-op rather than silently dead" promise from #580's design.

## Verification

- `go test -race ./hyoka/... -timeout 3m` — green across all 24 packages
- `go run . run --review-mode bogus --dry-run --prompt-id ... --config ...` — exits 1 with `Error: invalid --review-mode "bogus": must be "combined" or "isolated"`
- End-to-end isolated-mode runtime path verified via integration test (real Copilot session not invoked — would consume real quota; the integration test exercises the same engine pipeline against stubs, which is the same evidence chain other engine integration tests use, e.g. `TestIntegrationStubEvalReviewPipeline`).

## Process Notes

- Neo was NOT consulted (lockout enforced).
- Tests were pushed directly to `ronniegeraghty/issue-580-review-session-splitting` so they land inside PR #603 — fewer rebases, single PR for Switch to re-review.
- Co-authored-by trailer: Tank + Copilot only. Neo is not a co-author of the test commit per lockout.
- Switch tagged for re-review on PR #603.

## Implications for Future Work

- Establishes a pattern for "wiring-layer regression tests" on any future feature-flag work: in addition to unit-testing the leaf function, add an integration test through `engine.Run` with a recording stub asserting the runtime path. This is the cheapest defense against the "tests pass, behavior gone" failure mode that has now bitten this team twice (#355→#587, almost #580→#603).
- The `validateReviewMode` extraction is a reusable seam pattern: when a flag validator lives deep inside a Cobra `RunE`, extracting it as a package-level pure function gives 100% test coverage at near-zero cost. Worth applying to other validators in `cmd/run.go` if/when they grow.

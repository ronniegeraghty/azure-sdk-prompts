# Switch — Phase 6 Round-1 Test Review

**Date:** 2026-04-21  
**Author:** Switch 🤍  
**Branch context:** All three PRs target `phase-6` integration branch.

## Verdicts

| PR | Issue | Author | Verdict | Reassign on reject |
|----|-------|--------|---------|--------------------|
| #601 | #365 (Compare page redesign) | Trinity (per task spec) | ✅ APPROVE | — |
| #602 | #598 (Configurable prompt directory) | Tank | ✅ APPROVE | — |
| #603 | #580 (Review session splitting) | Neo | ❌ REQUEST CHANGES | **Tank** (Neo locked out) |

## Summary

### #601 — APPROVE
- 31 new tests across `comparison-groups.test.ts` (21 pure-function) and `comparison-page.test.tsx` (10 UI flow). All 99/99 vitest cases pass.
- Edge cases well-covered: top-bin overflow (1.0 → bin 9), missing prompt_metadata, malformed localStorage JSON, malformed group entries dropped, AND/OR filter semantics, persistence round-trip, no-evals-match warning.
- Gap (non-blocking): `group-builder.tsx` has no isolated test file — only exercised through page tests. Filed as follow-up suggestion.

### #602 — APPROVE
- 11 new tests in `prompt_dir_test.go` covering all new code paths: relative/absolute resolution, default-empty, LoadDir propagation, conflict detection (with error-msg substring assertion), `ResolvePromptDirCandidates` (override/no-override/nonexistent), `PeekPromptDirectory` (find/absent/missing-dir).
- `go test -race ./hyoka/... -timeout 3m` green across all 24 packages.
- Backwards-compat ("default behavior unchanged") locked in by `TestLoad_NoPromptDirectoryDefaultsEmpty` + `TestResolvePromptDirCandidates_NoOverrideDelegates`.
- Minor gaps (non-blocking): no test that `--prompts` flag wins over config (priority #1 vs #2); no malformed-YAML peek test.

### #603 — REQUEST CHANGES
**This is the #587 dead-flag risk repeating.** Unit-level `BuildReviewBuckets` coverage (14 tests) is excellent, but the wiring layers above it have **zero** test coverage:

1. **`Engine.reviewBuckets()` untested.** The cmd-flag → engine-mode → builder bridge is exactly what #587 caught as dead. A future refactor that drops `e.opts.ReviewMode` from the chain compiles + passes every test.
2. **`PromptReviewGrader.gradeSingle/gradeWithPanel` branch selection untested.** No stub-reviewer test asserts which path fires for `len(EvalCriteriaBuckets) ∈ {0, 1, 3}`.
3. **`mergeBucketResults` untested** (review/buckets.go:156). PR description calls this load-bearing for vote dedup; criterion-name prefixing is not asserted.
4. **CLI flag validation untested.** `cmd/run.go:289-293` rejects bogus `--review-mode`; no test locks this.
5. No "default = combined produces byte-identical structure" regression test.

Per reviewer-protocol, Neo is locked out. Reassigning to **Tank**: just shipped #602 (cmd-layer adjacent territory), the asks are bounded (~4 small test additions in `eval/engine_test.go`, `graders/prompt_review_grader_test.go`, new `review/buckets_test.go`, `cmd/run_test.go`).

## Pattern observed (for history.md)

When a PR re-implements a previously dead-flagged feature (#580 ↔ #587), unit tests of the new pure functions are necessary but not sufficient. The test of record needs to exercise the **wiring layer** — typically Engine/cmd plumbing — because that's where the previous attempt failed. Reviewer should require integration-style stub tests as gating, not optional follow-ups.

---

**Cross-team:**
- Tank: please pick up the four test additions outlined in the PR #603 review comment. Should be ~150 lines total. Once green, retag for re-review.
- Neo: locked out of #603 fix per reviewer-protocol; you'll see the full diff once Tank lands the tests.
- Trinity / Tank (#601, #602): clean approvals — proceed to merge whenever Coordinator wraps Round 1.

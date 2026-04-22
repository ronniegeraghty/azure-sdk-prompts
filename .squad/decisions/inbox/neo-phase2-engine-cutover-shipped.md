# Phase 2 engine cutover shipped (#625)

**Date:** 2026-04-22
**Branch:** ronniegeraghty/dev
**Commit:** a8a6d2d4

## What shipped

The engine now executes the unified `graders.Bundle` directly. Dual-storage of
`criteria.GraderConfig` + `graders.GraderConfig` is gone from the execution
path; `EngineOptions.GradersDir` is gone; the pre-#625 `loadCriteria` /
`loadGraders` split collapsed into a single `loadBundle` that calls
`graders.LoadUnifiedDir(CriteriaDir)`.

## Execution pipeline (runSingleEval)

1. `Bundle.MatchingErrors(props)` — if any file-level `when:` matches the
   eval's props AND the file failed to load, fail **this** eval with a clear
   message and skip grading. Other evals in the run are unaffected. (Q4
   deferred-error semantics.)
2. `MatchingUnifiedEntries(bundle, props)` → matched entries with group
   metadata preserved.
3. `PartitionMatched` → typed entries (file, program, output_check, ...) vs
   prompt entries (LLM review panel input).
4. Typed entries flow through `UnifiedGraderEntry.ToRuntimeConfig()` →
   `InstantiateGraders` → `RunGraders`.
5. Prompt entries + the prompt's own evaluation criteria feed
   `BuildUnifiedReviewBuckets` and `MergeUnifiedCriteria`, which populate
   `GraderInput.EvalCriteriaBuckets` and `GraderInput.EvalCriteria`.
6. Every result (typed + review) goes through a single `AggregateResults`
   call.

## Gate removal

`AggregateResults` no longer short-circuits on `Gate=true && Pass=false`.
Every grader contributes to the weighted score. `AggregateResult.Pass` is
now the AND of per-result `Pass`. `GateFailed` stays on the struct for
report-schema back-compat but is never set true.

## What's NOT in this cutover (deferred to Phase 3)

- `internal/criteria/` package — still on disk, still called by `cmd/list`,
  `cmd/validate`, and `internal/validate/schema.go`. The engine no longer
  touches it.
- The legacy `graders.GraderConfig` YAML shape (`kind` / `config`) remains
  readable via `graders.LoadDir`, but the engine no longer calls `LoadDir`.
  Phase 3 can decide whether to delete `LoadDir` and the legacy shape, or
  keep it for external tooling.
- No on-disk criteria YAML was migrated — the existing files already work
  under the unified schema's back-compat translator (legacy `prompt`-only
  entries get an implicit `type: prompt`).

## Test suite

`go test -race ./... -timeout 3m` is green. Integration-tagged tests
(`-tags=integration`) also pass. Rewrote 6 tests (engine_reviewbuckets_test,
engine_reviewmode_runtime_test, engine_test, grader_integration_test,
grader_test) to target the Bundle API and no-gate aggregation.

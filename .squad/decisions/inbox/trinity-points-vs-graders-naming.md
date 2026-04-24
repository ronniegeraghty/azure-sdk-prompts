# Decision Inbox — Schema field naming: `graders_total` actually counts Points

**Filed by:** Trinity 🖤  
**Date:** 2026-04-24  
**Status:** For team consideration (not blocking)

## Observation

`hyoka/internal/eval/engine_eval.go:690-691` writes:

```go
evalReport.GradersTotal  = countTotalPoints(agg.Results)
evalReport.GradersPassed = countPassedPoints(agg.Results)
```

The JSON field names are `graders_total` / `graders_passed`, but the values
are POINTS counts (sum of grader sub-checks across all graders), not grader
counts. For the test fixture
`reports/20260424-173723/.../test-dp-test-hello-markdown/test/sonnet/`, the
report has 6 graders and 14 total points. `graders_total = 14`,
`graders_passed = 13`.

## Why it bit us

The site's `evalGraderTotals(r)` returns `{passed: graders_passed, total: graders_total}`
verbatim when the engine totals are present, so the `graders_total` semantic
leaks straight into the UI. The per-eval headline subtitle was rendering
"across 14 graders" — wrong noun for the wrong number. Tank's run-detail
table also names the locals `gradersPassed` / `gradersTotal` (`run-detail-page.tsx:261-262`)
even though they hold points counts.

## Options

1. **Rename schema fields** (v4 bump): `graders_total` → `points_total`, etc.
   Update `report.EvalReport` JSON tags, all callers (engine, site, lib helpers,
   trends, comparison), and bump the schema version. Cleanest but cross-cutting.
2. **Add separate fields**: keep `graders_total` for backward compat (deprecated)
   and add `points_total` / `points_passed` as authoritative. Site reads new
   fields when present.
3. **Status quo + helpers**: leave the wire format alone, document the
   semantic, and standardize on `pointTotals.passed` / `pointTotals.graders`
   from `evalPointTotals(r)` in the React layer (which derives both from
   `grader_results` directly). Trinity used this approach for the immediate
   fix.

## Trinity's recommendation

Option 2 in the medium term — start emitting `points_total` / `points_passed`
alongside the legacy fields, mark the legacy fields deprecated in
`report/types.go`, and migrate the site/trends/comparison readers in one
follow-up PR. Avoids a hard schema break while paying down the naming debt.

## Scope of immediate fix (already shipped on dev)

Only the per-eval detail page subtitle. Run-detail page table left as-is —
its labels say "Score" so the misnamed locals aren't user-visible. Worth a
follow-up rename pass when somebody touches run-detail next.

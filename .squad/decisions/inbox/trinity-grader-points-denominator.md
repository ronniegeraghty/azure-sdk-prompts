# Decision: Grader Score Denominator Counts All Grader Points

**Date:** 2026-04-24  
**Agent:** Trinity 🌐  
**Branch:** ronniegeraghty/dev  
**Commit:** c06ca9e2

## Context

User reported: "On the site we don't expand out the grader points for the total score. For instance the last run I did had two graders, a file output one and a combined AI one, but they each have multiple grader points yet on the site it only shows scores of x/2 instead of x/<all-grader-points>."

## Root Cause

In `hyoka/internal/eval/engine_eval.go:636-637`, the `GradersTotal` and `GradersPassed` fields were populated using `len(agg.Results)` and `countPassed(agg.Results)`, which counted the number of graders, not the number of grader points.

Example:
- `file_check` grader with 3 points (file1, file2, file3)
- `output_check` grader with 2 points (min_files, require_files)
- Old denominator: 2 (number of graders)
- **Correct denominator: 5 (number of grader points)**

This caused the site to display `3/2` when it should display `3/5`.

## Decision

**The total score denominator is the sum of all grader points across all graders.**

1. **Added `countTotalPoints()` helper** that sums `len(g.Points)` for each grader in `agg.Results`.
   - Graders with `len(Points) == 0` are treated as having 1 point for backward compatibility with legacy graders that don't populate the Points slice.

2. **Added `countPassedPoints()` helper** that counts passed points across all graders.
   - For graders with Points, count the number of points where `Point.Pass == true`.
   - For graders with no Points (legacy), use the grader's overall `Pass` field (1 if true, 0 if false).

3. **Updated `engine_eval.go:636-637`** to use these helpers instead of `len(graders)` and `countPassed(graders)`.

## Implementation

**Files changed:**
- `hyoka/internal/eval/engine_eval.go` — Added `countTotalPoints()` and `countPassedPoints()` helpers, updated roll-up computation.
- `hyoka/internal/eval/grader_scoring_test.go` — Table-driven tests covering multiple points per grader, zero-point graders (legacy), mixed scenarios, empty results.

## Verification

- All tests pass: `go test ./hyoka/...`
- New unit tests cover edge cases: graders with multiple points, zero points, mixed legacy + modern graders.

## Impact

- **Fixes bug:** Site now displays accurate `X/Y` scores where Y is the total number of grader points, not the number of graders.
- **Backward compatible:** Graders with no Points (legacy) are treated as 1 point, so existing reports round-trip correctly.
- **Consistent with grader architecture:** The grader-points design (Phase 2, #GraderPoints) is now correctly reflected in the score roll-ups consumed by the site.

## Reusable Rule

**When aggregating grader scores, always count grader points, not graders.** The denominator is `Σ len(g.Points) for g in graders`, with a fallback of 1 for graders that don't populate Points.

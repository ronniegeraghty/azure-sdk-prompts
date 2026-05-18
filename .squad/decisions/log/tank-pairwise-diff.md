## Status: SHIPPED (2026-05-01)
# Decision: Pairwise Check Diff API Contract

**Date:** 2026-05-01  
**Agent:** Tank 📡  
**Commit:** 8a584f17  
**Status:** Implemented  

## Context

User reported: "I'm also seeing that our pages on the site for viewing the data from a pairwise run aren't really showing the differences between the evals and pointing out what tools made a difference, which improved the run and which made it worse and which made no difference."

The pairwise pipeline already computed tool-level impact (aggregate score differences), but the site needed per-check granularity to show which specific grader points improved or regressed when a tool was removed.

## Decision

Extended `PairwiseReport` with a new `check_diffs` field:

```go
type PairwiseReport struct {
    PromptID   string
    Baseline   VariantResult
    Variants   []VariantResult
    Impacts    []ToolImpact
    CheckDiffs map[string][]PairwiseCheckDiff `json:"check_diffs,omitempty"`
}
```

Each `PairwiseCheckDiff` captures:
- `GraderName`, `GraderType`: which grader the check belongs to
- `CheckID`: stable identifier (e.g., "check_0", "check_1")
- `CheckLabel`: human-readable label from the grader point
- `BaselinePassed`, `VariantPassed`: pass/fail for baseline and variant
- `Movement`: "improved" | "regressed" | "unchanged"
- `Reasoning`: optional message from variant's grader point if it failed

## API Contract for Trinity

The `/api/pairwise/<id>` endpoint now includes:

```json
{
  "check_diffs": {
    "without-tool-a": [
      {
        "grader_name": "file_check",
        "grader_type": "file",
        "check_id": "check_0",
        "check_label": "main.py exists",
        "baseline_passed": false,
        "variant_passed": true,
        "movement": "improved",
        "reasoning": "file found"
      },
      {
        "grader_name": "build_test",
        "grader_type": "program",
        "check_id": "check_0",
        "check_label": "build succeeds",
        "baseline_passed": true,
        "variant_passed": false,
        "movement": "regressed",
        "reasoning": "build failed: syntax error"
      }
    ],
    "without-tool-b": [ ... ]
  }
}
```

**Movement Classification:**
- **improved**: baseline failed → variant passed
- **regressed**: baseline passed → variant failed
- **unchanged**: both passed or both failed

**Edge Cases:**
- Missing checks in variant: treated as unchanged with `VariantPassed=false`
- Extra checks in variant (rare): included with `BaselinePassed=false`, `Movement="unchanged"`

## Implementation Notes

Avoided import cycle (pairwise → report → pairwise) by defining lightweight mirror types:
- `EvalReportData` with `ConfigName` and `Graders`
- `GraderData` with `Name`, `Type`, `Points`
- `PointData` with `Label`, `Pass`, `Message`

Conversion happens in `engine.go:evalReportToData()` before calling `pairwise.ComputeCheckDiffs()`.

## Testing

Added `hyoka/internal/pairwise/checkdiff_test.go`:
- `TestComputeCheckDiffs`: improved, regressed, unchanged scenarios
- `TestComputeCheckDiffsNilInputs`: nil safety
- `TestComputeCheckDiffsMissingChecks`: graceful handling of missing checks
- `TestComputeCheckDiffsExtraChecksInVariant`: rare case of variant-only checks
- `TestIndexPoints`: indexing logic validation

All tests pass with race detector.

## Next Steps for Trinity

Trinity (frontend) can now:
1. Fetch pairwise.json via `/api/pairwise/<runId>`
2. Access `check_diffs[variantConfigName]` for each variant
3. Group by `Movement` to render "Improved", "Regressed", "Unchanged" sections
4. Link to `GraderName` for context
5. Show `Reasoning` on hover/expand for failed checks

## Status: SHIPPED (2026-05-01)
# Decision: JSON Dual-Emit Migration for checks/points (C10)

**Status:** Implemented  
**Commit:** d955565e  
**Date:** 2026-05-01  
**Author:** Tank 📡

## Context

Neo's C5 renamed `GraderPoint` → `GraderCheck` (with one-release alias) and field `Points` → `Checks`, but kept the JSON field tag as `json:"points,omitempty"` to preserve on-disk compatibility with existing reports. This decision documents Tank's C10 implementation of the JSON migration horizon so the site can transition without a flag day.

## Problem

**Without dual-emit:**
- Old reports have `"points": [...]` keys
- New code writes `"checks": [...]` keys
- Site must handle EITHER shape depending on report age
- No clean migration path — requires version sniffing or try-both fallback logic scattered across site

**With dual-emit:**
- New code writes BOTH `"checks"` AND `"points"` keys (same data)
- Site reads from either key, preferring `"checks"`
- One release later, drop `"points"` key — all readers already switched to `"checks"`
- Clean migration: no version sniffing, no scattered fallbacks

## Decision

**Implement dual-emit for one release:**

### Go Side

In `hyoka/internal/report/types.go`:
- Changed `Checks []GraderCheck` JSON tag from `json:"points,omitempty"` to `json:"-"` (omit from default marshal)
- Implemented custom `MarshalJSON()` on `GraderResult`:
  - Uses type alias to avoid recursion
  - Emits BOTH `"checks"` and `"points"` keys with identical content
- Implemented custom `UnmarshalJSON()` on `GraderResult`:
  - Reads from either key
  - Prefers `"checks"` if both are present (future-proof)
- Added TODO comment: `// TODO(next-release): drop "points" JSON key; readers should consume "checks" only.`
- Added test coverage (`dual_emit_test.go`):
  - Unmarshal legacy (points-only) reports
  - Unmarshal new (checks-only) reports
  - Unmarshal both (prefers checks)
  - Marshal emits both keys with same data

### Site Side

In `site/src/app/data/types.ts`:
- `GraderResult` interface now has BOTH:
  - `checks?: GraderPoint[]` (new canonical key)
  - `points?: GraderPoint[]` (legacy back-compat key)
- Added `getChecks(result)` helper:
  - Returns `result.checks ?? result.points ?? []`
  - Prefers `checks` over `points` if both present
  - Centralizes fallback logic in one place

Updated all call sites:
- `lib/evalPass.ts`
- `lib/graderScore.ts`
- `components/GraderResultRow.tsx`
- `components/eval-detail-page.tsx`

All now call `getChecks(result)` instead of direct field access.

## Migration Horizon

**One release:**
1. **Current release (this commit):** Go dual-emits both keys; site reads either.
2. **Next release:** Drop `"points"` key from Go marshal. Site already reads `"checks"` first, so no breakage.

## Alternatives Considered

**1. Version field in report JSON:**
- PRO: Explicit version tracking
- CON: More complex; requires version logic in every reader
- CON: Doesn't help with mixed-version deployments (old reports + new site)

**2. Site-side try-both fallback everywhere:**
- PRO: No Go changes needed
- CON: Scattered `?? []` logic across 10+ call sites
- CON: Hard to audit when migration is complete
- CON: Doesn't help old reports if we drop the key

**3. Immediate breaking change (drop "points" now):**
- PRO: Simplest code
- CON: Breaks all existing reports on disk
- CON: Requires users to re-run all evals or manually migrate JSON

**Chosen approach (dual-emit) is the cleanest:** one-line helper + one release cycle + clean removal.

## Implementation Notes

**ProgressEvent fields (GraderChecksPassed/Total):**
- Did NOT need dual-emit
- These fields are only used in-process via callbacks, never JSON-serialized
- Confirmed by grepping for `json.Marshal.*ProgressEvent` (zero results)

**Test coverage:**
- Go: `TestGraderResultDualEmit` and `TestGraderResultMarshalDualEmit` verify both marshal and unmarshal paths
- Site: TypeScript build validates type safety; no runtime changes to test (just field access)

**Pre-existing failures:**
- `hyoka/internal/criteria/graders/tool_grader.go` (Neo's WIP)
- `hyoka/internal/report/generator_test.go:TestWriteReport` (baseline failure)
- Confirmed NOT introduced by this change

## Risks

**Low:** 
- Dual-emit increases report JSON size by ~5% (one duplicate array per grader)
- Mitigated: one-release horizon means short-lived overhead

**None:**
- Breaking existing reports: dual-emit preserves back-compat
- Site regressions: `getChecks()` helper centralizes all access

## Rollback Plan

If issues arise:
1. Revert commit d955565e
2. GraderResult reverts to `json:"points,omitempty"` tag
3. Site reverts to `points: GraderPoint[]` (non-optional)
4. No data loss — reports on disk unchanged

## Next Steps

**Next release (post-merge):**
1. Drop `"points"` key from Go MarshalJSON (keep UnmarshalJSON for one more release)
2. Remove `points?:` field from site TypeScript (keep `getChecks()` helper for consistency)
3. Update TODO comment → DONE

**Monitoring:**
- No specific metrics needed — this is a format migration, not a feature
- Site build + Go tests will catch any regressions

## Related

- Neo's C5: `refactor(graders): rename GraderPoint to GraderCheck with one-release alias` (commit 3c04d9a4)
- Trinity's task: CLI display strings (separate from this JSON migration)

// Single canonical pass-rollup for any eval row.
//
// History: each page used to compute "did this eval pass?" on its own,
// and the answers drifted (Phase 4 had three different shadows). This
// helper is the only place that decision lives. Every page-level pass/fail
// rollup MUST call evalPassFromPoints.
//
// v4 schema: GraderResult.pass is always boolean, derived from Points.
// The source-of-truth precedence is simplified:
//   1. Schema-v4 roll-up totals on the report (graders_passed / graders_total)
//      — engine already AND-rolled the unified grader aggregate. Trust them.
//   2. grader_results[*].pass — v4: always boolean, AND of all points.
//   3. r.success — final fallback for older reports or error rows.
//
// The shape is intentionally loose so it accepts both EvalReport (full
// detail) and EvalResult (summary). Pages that only have summaries
// transparently degrade to the engine's success bit.
import type { GraderResult } from "../data/types";

export interface EvalPassInput {
  success?: boolean;
  grader_results?: GraderResult[];
  graders_passed?: number;
  graders_total?: number;
}

/**
 * Returns the canonical pass verdict for an eval row.
 *
 * Use everywhere a page needs to render pass/fail on a single eval. Do
 * NOT recompute pass/fail inline — divergence between pages is the bug
 * this helper exists to prevent.
 */
export function evalPassFromPoints(r: EvalPassInput): boolean {
  // Schema v4: trust the engine's roll-up totals when present.
  if (typeof r.graders_total === "number" && r.graders_total > 0) {
    return (r.graders_passed ?? 0) === r.graders_total;
  }

  // Fall back to per-grader AND rollup over pass field (v4: always boolean).
  if (r.grader_results && r.grader_results.length > 0) {
    return r.grader_results.every((g) => g.pass);
  }

  // No grader data at all (error rows, etc.). Trust the engine's success bit.
  return r.success === true;
}

/**
 * True iff a single grader is considered "passed".
 * v4: GraderResult.pass is always boolean (derived from Points in the engine).
 */
export function graderPasses(g: GraderResult): boolean {
  return g.pass;
}

/**
 * Roll-up totals for the eval. Prefers v4 engine-side totals; falls back
 * to recomputing from grader_results. Returned object is `{passed, total}`.
 *
 * Used by the eval-detail score card and any other surface that wants to
 * render `N/M` instead of just a pass/fail boolean.
 */
export function evalGraderTotals(r: EvalPassInput): { passed: number; total: number } {
  if (typeof r.graders_total === "number" && r.graders_total > 0) {
    return { passed: r.graders_passed ?? 0, total: r.graders_total };
  }
  const grs = r.grader_results ?? [];
  return {
    passed: grs.filter((g) => g.pass).length,
    total: grs.length,
  };
}

/**
 * Fractional pass rate (0–100) over the eval's Points. Returns 0 when
 * no points are available (e.g. error rows). Use this everywhere the UI
 * wants partial-credit % rather than binary pass/fail.
 */
export function pointsPassRate(r: EvalPassInput): number {
  const t = evalPointTotals(r);
  return t.total > 0 ? (t.passed / t.total) * 100 : 0;
}

/**
 * Total Points seen and total Points passed across every grader on this
 * eval. For graders that did not emit Points we synthesize a single
 * virtual point so the count remains meaningful (`✓ N / N points across M
 * graders` reads correctly even when one grader is points-less).
 */
export function evalPointTotals(
  r: EvalPassInput
): { passed: number; total: number; graders: number } {
  const grs = r.grader_results ?? [];
  let passed = 0;
  let total = 0;
  for (const g of grs) {
    if (g.points && g.points.length > 0) {
      total += g.points.length;
      passed += g.points.filter((p) => p.pass).length;
    } else {
      // Synthesize a single point from this grader's verdict.
      total += 1;
      if (g.pass) passed += 1;
    }
  }
  return { passed, total, graders: grs.length };
}


// Single canonical pass-rollup for any eval row.
//
// History: each page used to compute "did this eval pass?" on its own,
// and the answers drifted (Phase 4 had three different shadows). This
// helper is the only place that decision lives. Every page-level pass/fail
// rollup MUST call evalPassFromPoints.
//
// Source-of-truth precedence:
//   1. Schema-v3 roll-up totals on the report (graders_passed / graders_total)
//      — engine already AND-rolled the unified grader aggregate. Trust them.
//   2. grader_results[*].points — AND every Point across every grader. This
//      is the v3 fallback when the roll-up totals are absent (e.g. test
//      fixtures, partial reports).
//   3. grader_results[*].pass — older v3 fallback when a grader emitted
//      no Points but did set pass.
//   4. r.success — final fallback for v2 reports / EvalResult summaries
//      that don't carry grader_results.
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
  // Schema v3: trust the engine's roll-up totals when present.
  if (typeof r.graders_total === "number" && r.graders_total > 0) {
    return (r.graders_passed ?? 0) === r.graders_total;
  }

  // Fall back to per-grader AND rollup over Points / pass / derived score.
  if (r.grader_results && r.grader_results.length > 0) {
    return r.grader_results.every(graderPasses);
  }

  // No grader data at all (v2 EvalResult summary, error rows). Trust the
  // engine's success bit.
  return r.success === true;
}

/**
 * True iff a single grader is considered "passed". Mirrors the engine's
 * tri-state aware rollup: Points beat pass beats derived-from-score.
 */
export function graderPasses(g: GraderResult): boolean {
  // Points are authoritative when present.
  if (g.points && g.points.length > 0) {
    return g.points.every(p => p.pass);
  }

  // Explicit pass field (v2 graders that don't emit points).
  if (g.pass === true) return true;
  if (g.pass === false) return false;

  // pass is null/undefined — fall back to derived truth.
  const criteria = g.scores?.criteria;
  if (criteria && criteria.length > 0) {
    return criteria.every(c => c.passed);
  }
  if (
    g.overall_score != null &&
    g.max_score != null &&
    g.max_score > 0 &&
    g.overall_score === g.max_score
  ) {
    return true;
  }
  return false;
}

/**
 * Roll-up totals for the eval. Prefers v3 engine-side totals; falls back
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
    passed: grs.filter(graderPasses).length,
    total: grs.length,
  };
}

/**
 * Total Points seen and total Points passed across every grader on this
 * eval. For graders that did not emit Points we synthesize a single
 * virtual point so the count remains meaningful (`✓ N / N points across M
 * graders` reads correctly even when one grader is points-less).
 */
export function evalPointTotals(r: EvalPassInput): { passed: number; total: number; graders: number } {
  const grs = r.grader_results ?? [];
  let passed = 0;
  let total = 0;
  for (const g of grs) {
    if (g.points && g.points.length > 0) {
      total += g.points.length;
      passed += g.points.filter(p => p.pass).length;
    } else {
      // Synthesize a single point from this grader's verdict.
      total += 1;
      if (graderPasses(g)) passed += 1;
    }
  }
  return { passed, total, graders: grs.length };
}

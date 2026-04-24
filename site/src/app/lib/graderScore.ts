// Canonical grader score formatting for v4 reports.
// Per Morpheus's plan: ALWAYS "N/M points", even for single-point graders.
// No "Passed", no "100%", no special cases.

import type { GraderResult } from "../data/types";

/**
 * Returns the canonical score string for a grader: "N/M points".
 * This is the ONLY score string format shown in the row header.
 */
export function formatGraderScore(result: GraderResult): string {
  const passed = result.points.filter((p) => p.pass).length;
  const total = result.points.length;
  return `${passed}/${total} points`;
}

/**
 * Returns true if the grader passed (all points passed).
 * For v4 reports, this is just `result.pass`, but keeping the helper
 * for consistency with older code patterns.
 */
export function graderPasses(result: GraderResult): boolean {
  return result.pass;
}

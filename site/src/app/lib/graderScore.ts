// Canonical grader score formatting for v4 reports.
// Per Morpheus's plan: ALWAYS "N/M points", even for single-point graders.
// No "Passed", no "100%", no special cases.

import type { GraderResult } from "../data/types";

/**
 * Returns the canonical score string for a grader: "N/M points".
 * This is the ONLY score string format shown in the row header.
 */
export function formatGraderScore(result: GraderResult): string {
  const points = result.points ?? [];
  const total = points.length;
  // Defensive: if no detail points were emitted, treat the grader itself as a
  // single implicit point so we never render bare "PASS" / "100%" / "0/0".
  if (total === 0) {
    return result.pass ? "1/1 points" : "0/1 points";
  }
  const passed = points.filter((p) => p.pass).length;
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

### 2026-04-22T21:08:00Z: User directive
**By:** ronniegeraghty (via Copilot)
**What:** Graders never gate (hard-fail) an eval. Every grader just runs and reports its result into the results object. The eval continues regardless of any single grader's score.
**Why:** "Graders shouldn't error out an eval, a grader should just run and report their result in the results. Why would we want to stop checking the rest of the graders." Resolves Q3 (typed-grader Gate) and Q9 (prompt-grader Gate) in the unification proposal — both default to no-gate. Implication: drop `Gate` from the unified schema entirely, OR keep the field but make it a no-op (reporting-only flag for downstream consumers). Lead's call.

### 2026-04-22T21:08:00Z: User decision — Q1 keep --criteria-dir
**By:** ronniegeraghty (via Copilot)
**What:** Keep the `--criteria-dir` CLI flag as-is. Do not rename to `--graders-dir`.
**Why:** Established name, zero breaking change, "criteria" still semantically reasonable for typed graders.

# Grader Coverage Investigation (2026-04-23)

## Problem Statement

User reported: "Graders aren't running on all evals."

## Investigation Summary

Traced grader execution end-to-end from CLI flags → engine options → bundle loading → grader execution.

## Key Findings

### 1. Grader Pipeline Architecture (Post-#625)

The unified grader system has two types:

- **Typed graders** (`output_check`, `file`, `program`, `behavior`, etc.) - Run via `criteria.RunGradersWithHooks()`
- **Prompt-type graders** (`type: prompt`) - Fed into AI review panel as evaluation criteria

Both types live in the same `criteria/*.yaml` files and are loaded via `criteria.LoadUnifiedDir()`.

### 2. Auto-Discovery Works

`--criteria-dir` has automatic fallback resolution (cmd/helpers.go:79-88):
```go
candidates := config.ResolveCandidates(proj, "criteria", "./criteria", "../criteria")
```

So users do NOT need to explicitly pass `--criteria-dir` if `./criteria` or `../criteria` exists.

### 3. When Graders Don't Run

**Scenario A: Empty/Missing Criteria Directory**
- `CriteriaDir == ""` or dir doesn't exist → `loadBundle()` returns early
- `e.graderBundle` stays nil → zero graders run
- Result: `grader_results: null` in report

**Scenario B: Review Skipped (`--skip-review`)**
- Typed graders run normally
- Prompt-type graders are SKIPPED (they're embedded in the review criteria)
- Result: User only sees typed graders (e.g., just "Output Files Exist")

**Scenario C: When Filter Doesn't Match**
- Grader file has `when: language: python`
- Prompt has `language: go`
- No graders match → zero graders run for that eval
- This is CORRECT behavior (graders are attribute-matched)

### 4. Evidence from Real Reports

| Run | Criteria Dir | Skip Review | Graders Run | Notes |
|-----|-------------|-------------|-------------|-------|
| 20260423-025247 | auto (./criteria) | YES | 1 (typed only) | DefaultAzureCredential skipped |
| 20260423-025359 | explicit (../criteria) | YES | 1 (typed only) | DefaultAzureCredential skipped |
| 20260423-025531 | explicit (../criteria) | NO | 1 typed + review | Both grader types ran |
| 20260407-011814 | (old system) | YES | 0 | Pre-#625 unification |

### 5. Current Behavior is CORRECT

After reviewing the code and test runs:
- Graders ARE running when criteria files exist
- Auto-discovery works
- Typed vs. prompt partition is working as designed
- The user's observation may be from:
  - Old reports (pre-#625)
  - Using `--skip-review` and not realizing prompt-type graders are review-bound
  - Looking at evals that failed during generation (no files → no grading phase)

## Recommendations

### A. UX/Documentation Improvements (Small)

1. **Emit a warning when no graders match**
   - Currently silent when `len(matched) == 0`
   - User should know "no graders matched this prompt's properties"

2. **Document the typed vs. prompt split**
   - Users may not realize `type: prompt` graders require review to run
   - Add to docs: "To run all graders, omit `--skip-review`"

3. **Improve `hyoka list` output**
   - Show which graders would apply to each prompt
   - `hyoka list --prompt-id X --show-graders`

### B. No Code Changes Needed (Yet)

The engine is working as designed. The "graders not running" observation is likely:
- User confusion about when prompt-type graders run
- Looking at old reports from before unification
- Not understanding attribute matching (when filters)

## Proposed Next Steps

1. **Update docs** (Oracle task):
   - Explain typed vs. prompt grader execution
   - Show example criteria files
   - Document `--skip-review` behavior

2. **Add debug logging** (Neo task, trivial):
   - Log "matched N graders for prompt X"
   - Log "skipped M prompt-type graders (review disabled)"

3. **Defer typed grader CLI surface (#622)**
   - This is a known gap (graders from `internal/graders/` not CLI-accessible yet)
   - Not urgent — users CAN define graders in `criteria/*.yaml`

## Conclusion

**No bug found.** Graders are running as designed. The issue is UX/observability:
- Users don't see which graders matched
- Prompt-type graders are "hidden" inside review results
- No clear signal when zero graders matched a prompt

Recommend: Add logging + documentation, defer structural changes.

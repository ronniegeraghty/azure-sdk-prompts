# Grader Source-of-Truth Identity Principle

## Core Rule

**A grader's identity is determined by its source-of-truth, not its review mode.**

When multiple criteria sources contribute to an evaluation (prompt frontmatter, criteria files, future: remote criteria), each source must produce a separate grader instance if we want distinguishable pass/fail reporting.

## Why This Matters

**Problem:** If prompt-frontmatter criteria and criteria-file entries share the same AI review grader, failures are aggregated and users cannot distinguish which source contributed which failing criteria.

**Solution:** Each criteria source → separate bucket → separate Copilot review session → separate grader display entry.

## Current Sources

1. **Prompt frontmatter** — `evaluation_criteria:` block in `.prompt.md` files
   - Bucket name: `"Criteria from prompt file"`
   - Always isolated from other sources

2. **Criteria files** — `criteria/*.yaml` files matched against prompt properties
   - Bucket name(s): `"combined"` (combined mode) or per-entry/group names (isolated mode)
   - Mode controls how criteria-FILE entries are bucketed among themselves

## Implementation Location

`hyoka/internal/criteria/buckets.go` → `BuildUnifiedReviewBuckets`

The function enforces source-separation as a FIRST-CLASS partitioning rule:
1. If `promptCriteria` is non-empty, create a dedicated bucket named "Criteria from prompt file"
2. Then bucket matched criteria-file entries based on mode (combined or isolated)
3. Each bucket produces its own `PromptReviewGrader` instance

## Review Mode is Secondary

`--review-mode combined` vs `isolated` controls how criteria-FILE entries are bucketed **within their source**. It does NOT control whether prompt-frontmatter criteria are separated from criteria-file entries. Source-separation is a harder rule than mode-separation.

## Edge Cases

- `promptCriteria == ""` + criteria files exist → only criteria-file bucket(s)
- `promptCriteria != ""` + no criteria files → only prompt-frontmatter bucket
- Both empty → zero buckets (no review graders run)

## Display Flow

Bucket Name → `graders.ReviewBucket.Name` → `PromptReviewGrader.Name()` → `progress.ProgressEvent.GraderID` → interactive display grader name

Users see TWO AI review grader entries when both sources exist:
- "Criteria from prompt file" (grader 1)
- "combined" or other bucket name (grader 2)

## Historical Context

This decision was implemented in commit `27c04c71` (2026-04-24) to fix a user-reported bug: "I'm only seeing one group of ai review graders running but I thought we decided that if we wanted grader points to be graded in the same review session they would have to be grader points on the same grader."

The fix applies the existing principle ("different sources = different graders = different sessions") to the prompt-frontmatter vs criteria-files boundary, which had been incorrectly merged into a single `combined` bucket.

## Future Extensions

When new criteria sources are added (e.g., remote criteria fetched from an API), apply the same source-separation rule:
1. Identify the source (local prompt, local files, remote API, etc.)
2. Create a dedicated bucket for that source with a meaningful name
3. Ensure the bucket flows through the review pipeline as a separate grader

## Testing

- `hyoka/internal/criteria/buckets_test.go`: Tests source-separation logic
- `hyoka/internal/eval/engine_reviewbuckets_test.go`: Tests engine-level bucket construction
- `hyoka/internal/eval/engine_reviewmode_runtime_test.go`: Tests runtime behavior (ReviewBuckets() vs Review())

All tests enforce that prompt-frontmatter criteria produce a separate bucket from criteria-file entries.

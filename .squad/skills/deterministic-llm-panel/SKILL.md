---
name: Deterministic LLM Panel
description: Stable id-based voting for multi-model review panels
created: 2026-04-27
updated: 2026-05-01
status: shipped
---

# Deterministic LLM Panel Pattern

## Problem

Multi-model review panels produce non-deterministic results when each reviewer LLM paraphrases the same check text differently. Vote aggregation keyed by exact string match on `c.Name` causes paraphrases to split into separate criteria → inconsistent point counts across runs.

## Solution

**Stable check IDs** assigned once at criteria-bundling time flow through the entire pipeline:
1. Reviewer prompt: `check_1: File hello.md exists...`
2. Reviewer JSON contract: `{"id": "check_1", "passed": true, "reasoning": "..."}`
3. Vote aggregation: key by `bucket::check_1` (not by paraphrased name)
4. Final label: Canonical YAML text (never from LLM echo)

## Implementation References

- **ID assignment:** `hyoka/internal/criteria/buckets.go:326` (`FormatUnifiedPromptEntries` emits `ReviewCheck` with `check_<n>` IDs)
- **Prompt rendering:** `hyoka/internal/review/prompt.go:87` (`BuildReviewPrompt` renders `check_1: text` format)
- **Validator:** `hyoka/internal/review/reviewer.go:278` (`parseReviewResponseV2` validates id-set match)
- **Vote keying:** `hyoka/internal/review/reviewer.go:694` (`averageReview` keys by `bucket::id`, not `Name`)

## Key Patterns

### 1. Stable ID Format

`check_<n>` where `n` is 1-based index within the rendered bucket. Deterministic, short, easy for LLMs to echo back.

### 2. Reviewer JSON Contract

```json
{
  "criteria": [
    {"id": "check_1", "passed": true, "reasoning": "..."}
  ],
  "summary": "...",
  "issues": ["..."],
  "strengths": ["..."]
}
```

**Dropped from contract:** `criterion` text field (LLM no longer trusted to echo label).

### 3. Validator Retry-Then-Drop

On validation error, re-prompt with precise feedback:
```
Your response is missing ids: [check_2]. Extra ids: [check_99].
Please return exactly: [check_1, check_2, check_3].
```

After max retries (2), drop reviewer with `slog.Warn`:
```go
slog.Warn("reviewer dropped: invalid criteria ids after retries",
    "model", model, "validation_errors", validationErrors)
```

### 4. Bucket::ID Vote Keying

- Vote key: `<bucket-name>::<check_id>` for non-`combined` buckets, plain `<check_id>` for `combined`
- Display label: Canonical YAML text (bucket-prefixed if applicable)

Example:
```go
// Bucket "security" has check_1: "No hardcoded secrets"
voteKey := "security::check_1"
displayLabel := "[security] No hardcoded secrets"
```

### 5. Canonical Labels from YAML

`parseReviewResponseV2` sets `CriterionResult.Name` from `expected[i].Text` (YAML source of truth), **never** from LLM echo. This ensures identical labels across runs.

## Smoke Test Proof

Ran `test-dp-test-hello-markdown` twice with `test/baseline` config (2 models, panel review). Grader breakdowns **identical**:
- Same point counts per grader
- Same pass/fail verdicts per check
- No label drift or duplicate points

## When to Use This Pattern

- Multi-model review panels where LLMs vote on the same criteria
- Any system where LLM outputs are aggregated/compared (e.g., grading, consensus)
- Scenarios requiring reproducible, auditable results across LLM runs

## Contract for Future Changes

Any code touching the reviewer-pipeline MUST preserve:
1. **Stable IDs assigned once** (at criteria-bundling time, 1-based index)
2. **Canonical labels from YAML** (never from LLM echo)
3. **Bucket::id vote keys** (no paraphrased name keys)

Violating these rules will re-introduce non-determinism.


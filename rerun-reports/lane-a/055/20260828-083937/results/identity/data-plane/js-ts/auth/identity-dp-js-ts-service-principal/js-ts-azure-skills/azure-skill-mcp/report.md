# Evaluation Report: identity-dp-js-ts-service-principal

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 2206.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-js-ts-service-principal` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 0/11 |
| Duration | 2206.0s |
| Timestamp | 2026-08-28T00:39:37Z |
| Files Generated | 0 |
| Event Count | 14 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 2186.3s |
| Review | 12.7s |
| **Total** | **2206.0s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** js-ts-azure-skills/azure-skill-mcp

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |

## Error

```
evaluation failed: sending prompt: waiting for session.idle: context deadline exceeded
```

**Details:**

```
sending prompt: waiting for session.idle: context deadline exceeded
```

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (0/1)
      - grader executed: Fail
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Fail (0/1)
      - grader executed: Fail
  - @azure/identity for Authentication (prompt): Fail (0/1)
      - grader executed: Fail
  - Client Constructor with Endpoint and Credential (prompt): Fail (0/1)
      - grader executed: Fail
  - Async/Await Pattern (prompt): Fail (0/1)
      - grader executed: Fail
  - Pagination with for-await-of (prompt): Fail (0/1)
      - grader executed: Fail
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Fail (0/1)
      - grader executed: Fail
  - RestError Exception Handling (prompt): Fail (0/1)
      - grader executed: Fail
  - No Deprecated Packages (prompt): Fail (0/1)
      - grader executed: Fail
  - Logging via @azure/logger (prompt): Fail (0/1)
      - grader executed: Fail
  - package.json with Correct Dependencies (prompt): Fail (0/1)
      - grader executed: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 11.00** | **Σ 0.0000** | **0.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-js-ts-service-principal --config js-ts-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

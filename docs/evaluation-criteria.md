# Evaluation Criteria System

This document explains how hyoka scores generated code — the criteria system, the multi-model review panel, and how results are consolidated.

## Overview

hyoka uses an **LLM-as-judge** approach: after generating code, a panel of reviewer models evaluates it against a set of criteria. Each criterion is scored as **pass** or **fail**. The overall score is the number of passed criteria out of the total.

## Criteria Tiers

hyoka uses a two-tier criteria system (with a third tier planned):

### Tier 1: General Criteria (Always Applied)

Five criteria from the embedded rubric (`rubric.md`) that apply to **every** evaluation:

| # | Criterion | What It Tests |
|---|-----------|---------------|
| 1 | **Code Builds** | Does the generated code compile/build without errors? |
| 2 | **Latest Package Versions** | Are the Azure SDK packages the latest stable versions? |
| 3 | **Best Practices** | Does it follow Azure SDK best practices? (DefaultAzureCredential, proper disposal, async patterns) |
| 4 | **Error Handling** | Are errors handled properly? Retries? Timeouts? |
| 5 | **Code Quality** | Clean, readable, well-structured code? |

These general criteria ensure a baseline quality bar across all prompts regardless of what's being tested.

### Tier 2: Prompt-Specific Criteria (Author-Defined)

Custom criteria defined in the `## Evaluation Criteria` section of each prompt file. These test scenario-specific requirements:

```markdown
## Evaluation Criteria

- Uses `azure-storage-blob` and `azure-identity` packages
- `BlobServiceClient` created with `DefaultAzureCredential`
- `BlobClient.upload_blob()` called with `overwrite` parameter
- `ContainerClient.list_blobs()` iteration to enumerate blobs
- `HttpResponseError` and `ResourceExistsError` handling
```

Each bullet becomes an individual criterion scored as pass/fail.

### Tier 3: Attribute-Matched Criteria (Planned)

A future tier where criteria YAML files in a `criteria/` directory are automatically matched to prompts by metadata attributes (language, service, plane, category). For example, a `java.yaml` criteria file would apply to all Java prompts.

## The Review Panel

### How It Works

1. **Three reviewer models** run in parallel, each in an independent Copilot session
2. Each reviewer receives:
   - The original prompt text
   - All generated code files
   - The embedded rubric (general criteria)
   - Prompt-specific evaluation criteria (if provided)
   - Optional reference answer for comparison
3. Each reviewer returns a JSON response with pass/fail per criterion
4. The first reviewer model **consolidates** all results

### Majority Voting

The consolidation uses majority voting per criterion:

```
Criterion: "Uses DefaultAzureCredential"

  Reviewer 1 (Claude Opus):  PASS ✓
  Reviewer 2 (Gemini Pro):   PASS ✓
  Reviewer 3 (GPT-4.1):      FAIL ✗

  Consolidated: PASS (2/3 majority)
```

A criterion passes if **more than half** of the reviewers marked it as passed.

### Consolidation Process

1. **Primary path:** First reviewer model synthesizes all panel results into a consolidated review
2. **Fallback path:** If LLM consolidation fails, `averageReview()` applies mechanical majority voting
3. Issues and strengths from all reviewers are merged (union)
4. The consolidated result becomes the final `ReviewResult`

### Why Multiple Models?

- **Reduces bias** — a single model may have blind spots or systematic biases
- **Increases reliability** — majority voting smooths out individual errors
- **Prevents self-bias** — each reviewer gets its own session (can't see generation reasoning)
- **Cross-model validation** — if all three models agree, confidence is high

## Scoring

### Per-Criterion Score

Each criterion is either:
- **Passed** (`true`) — criterion is fully met
- **Failed** (`false`) — criterion is not met

With a `reason` explaining why:

```json
{
  "name": "Code Builds",
  "passed": true,
  "reason": "The Python script uses valid syntax and all imports are correct."
}
```

### Overall Score

```
overall_score = count of passed criteria
max_score     = total number of criteria
```

For example, if 3 general + 5 prompt-specific criteria are evaluated and 6 pass:
- `overall_score = 6`
- `max_score = 8`
- Score percentage: 75%

### Success Definition

An evaluation is considered **successful** when:
- All criteria pass (`passed_count == total_count`)
- No errors occurred during generation, build, or review
- No guardrails were triggered

## The Rubric

The rubric is embedded in the binary via `//go:embed rubric.md` and included in every review prompt. It instructs reviewers to:

1. Actively verify where possible (attempt to build code, check package versions)
2. Evaluate both general and prompt-specific criteria
3. Respond with **only** a JSON object (no markdown, no explanation)
4. Include a `reason` for each criterion

### Rubric Output Format

Reviewers must respond with this exact JSON structure:

```json
{
  "scores": {
    "criteria": [
      {"name": "Code Builds", "passed": true, "reason": "Compiles without errors"},
      {"name": "Latest Package Versions", "passed": false, "reason": "Uses azure-storage-blob 12.14.0, latest is 12.19.0"},
      {"name": "Uses DefaultAzureCredential", "passed": true, "reason": "Correctly imports and uses DefaultAzureCredential"}
    ]
  },
  "overall_score": 2,
  "max_score": 3,
  "summary": "Code is functional but uses outdated package versions.",
  "issues": ["Outdated azure-storage-blob package"],
  "strengths": ["Clean code structure", "Proper credential management"]
}
```

## Review Prompt Construction

The `BuildReviewPrompt()` function assembles the full review prompt:

1. **System message** — "You are a senior Azure SDK code reviewer..."
2. **Original prompt** — the exact prompt text from the `.prompt.md` file
3. **Prompt-specific criteria** — from the `## Evaluation Criteria` section
4. **Generated code** — all files produced by the generator
5. **Reference answer** — if provided via `reference_answer` frontmatter field
6. **Embedded rubric** — general criteria + output format instructions
7. **JSON-only instruction** — "Respond with ONLY a valid JSON object"

## Data Structures

### CriterionResult

```go
type CriterionResult struct {
    Name   string `json:"name"`
    Passed bool   `json:"passed"`
    Reason string `json:"reason,omitempty"`
}
```

### ReviewScores

```go
type ReviewScores struct {
    Criteria []CriterionResult `json:"criteria"`
}
```

### ReviewResult

```go
type ReviewResult struct {
    Model        string         // reviewer model name
    Scores       ReviewScores   // per-criterion results
    OverallScore int            // count of passed criteria
    MaxScore     int            // total criteria count
    Summary      string         // human-readable summary
    Issues       []string       // problems found
    Strengths    []string       // things done well
    Events       []ReviewEvent  // session event timeline
}
```

## Example: Full Scoring Flow

1. **Prompt** `storage-dp-python-crud` has 6 evaluation criteria
2. **General rubric** adds 5 criteria → **11 total criteria**
3. **Three reviewers** each score all 11 criteria independently
4. **Majority voting** consolidates:

| Criterion | Reviewer 1 | Reviewer 2 | Reviewer 3 | Consolidated |
|-----------|-----------|-----------|-----------|-------------|
| Code Builds | ✓ | ✓ | ✓ | ✓ |
| Latest Packages | ✗ | ✗ | ✓ | ✗ |
| Best Practices | ✓ | ✓ | ✓ | ✓ |
| Error Handling | ✓ | ✓ | ✗ | ✓ |
| Code Quality | ✓ | ✓ | ✓ | ✓ |
| Uses azure-storage-blob | ✓ | ✓ | ✓ | ✓ |
| DefaultAzureCredential | ✓ | ✓ | ✓ | ✓ |
| upload_blob() | ✓ | ✓ | ✓ | ✓ |
| list_blobs() | ✓ | ✓ | ✓ | ✓ |
| download_blob() | ✓ | ✓ | ✗ | ✓ |
| Error handling | ✓ | ✗ | ✓ | ✓ |

**Result:** 10/11 passed (90.9%) — evaluation marked as **failed** (not all criteria passed).

> **Note:** Success requires **all** criteria to pass. A 90% score is still a failure. This strict standard ensures generated code is production-ready.

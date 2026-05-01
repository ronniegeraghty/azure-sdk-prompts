---
name: "deterministic-llm-panel"
description: "How to build deterministic multi-model LLM voting panels via server-controlled stable IDs"
domain: "llm-orchestration, evaluation, multi-model-consensus"
confidence: "high"
source: "earned (hyoka prompt-grader determinism fix, 2026-04-27)"
---

## Context

Use this pattern any time you have:
- Multiple LLMs reviewing the same set of items (criteria, claims, options).
- A vote / merge step that must produce identical output for identical input.
- Items defined by free-text labels that the LLMs would otherwise echo back paraphrased.

The failure mode this prevents: paraphrase drift across reviewers causes the vote function to split logically-identical items into separate buckets, producing non-deterministic counts run-to-run.

## The Pattern

### 1. Assign stable IDs at item-definition time

The application — never the LLM — owns the IDs. Format: short, stable, easy to echo (`check_1`, `option_a`, `claim_3`). Slug-from-text is tempting but fragile (collisions, length, instability under text edits).

```go
type Item struct {
    ID    string // "check_1" — owned by us
    Text  string // human-readable label, source of truth
    Group string // optional namespace (bucket/section)
}
```

### 2. Send IDs in the prompt; rule out paraphrase

```
You MUST return one judgment per id below. No extras, no missing entries.

- check_1: <text>
- check_2: <text>

Schema:
{"results":[{"id":"check_1","passed":true,"reasoning":"..."}, ...]}

Rules:
- The "id" field MUST be one of: check_1, check_2.
- Do NOT invent ids. Do NOT omit ids. Do NOT echo the item text.
```

The "Do NOT echo" rule matters — it removes a degree of freedom the LLM would otherwise use to drift.

### 3. Validate the response against the expected ID set

```go
func validate(expected []string, got []Result) []string {
    var errs []string
    expectedSet := setOf(expected)
    seen := map[string]bool{}
    for _, r := range got {
        if !expectedSet[r.ID]   { errs = append(errs, "extra id: " + r.ID) }
        if seen[r.ID]           { errs = append(errs, "duplicate id: " + r.ID) }
        seen[r.ID] = true
    }
    for _, id := range expected {
        if !seen[id] { errs = append(errs, "missing id: " + id) }
    }
    return errs
}
```

### 4. Re-prompt once on shape failure, then drop the reviewer

Three-tier resilience: parse → validate → drop. Don't fail the whole panel because one reviewer was flaky; don't accept a malformed reviewer because then the whole point of the panel is undermined.

```go
for attempt := 0; attempt <= maxRetries; attempt++ {
    resp := callLLM(prompt)
    if errs := validate(expected, resp); len(errs) > 0 && attempt < maxRetries {
        prompt = repromptWithSpecificErrors(errs)
        continue
    } else if len(errs) > 0 {
        slog.Warn("reviewer dropped: invalid ids after retries", "model", m)
        return nil // panel handles missing reviewer
    }
    return resp
}
```

### 5. Vote keys by ID; display label comes from the application

```go
type agg struct {
    id        string
    label     string // from the application's Item.Text — NEVER from any LLM
    failCount int
    total     int
}
m := map[string]*agg{} // key = id (or "group::id" for namespaced)
```

Display layer reads `agg.label`. Vote logic reads `agg.id`. Truncation/styling is a display concern, not a vote concern.

### 6. Treat missing-from-reviewer as fail (strict consensus)

If a reviewer returns a partial response that somehow passed validation, count any expected-but-not-returned id as "failed by that reviewer". Matches the strict any-fail philosophy: silence is not consent.

## Anti-patterns

- **❌ Trusting the LLM to echo the item text verbatim.** "Use the original text as the id" is a request, not a contract.
- **❌ Slugs derived from text as IDs.** Edit the text → ID changes → historical comparisons break. Collisions on similar items.
- **❌ String concatenation for namespaced keys without sanitization.** `bucket::id` collides if `bucket` contains `::`.
- **❌ Failing the whole eval on one bad reviewer.** That's why the panel exists. Drop with a warning.
- **❌ Validating only "criteria array non-empty".** Misses extras, missing, duplicates, malformed shapes.

## When NOT to apply

- Single-LLM evaluation (no vote → no key collisions).
- LLM is producing freeform output that's not being merged with other LLMs' output.
- Item set is itself LLM-generated (no canonical IDs to assign — different problem).

## Verification

A determinism regression test belongs in the test suite: run the same input twice with a deterministic stub that returns paraphrased labels. Assert identical structure. If this test isn't in CI, the pattern isn't truly enforced.

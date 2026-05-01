# Grader Reference

## Overview

hyoka's grading system evaluates generated code and outputs using composable, single-concern graders. Each grader checks exactly one thing: file existence, build success, LLM-based code review, tool constraint compliance, and so on. Results are aggregated by weighted scoring to produce a final evaluation report.

## Key Concepts

**Graders never gate evaluations.** Each grader runs independently and reports its result. Failed graders don't stop other graders or halt the evaluation—all graders run and contribute to the final score.

**Load-time validation.** When hyoka loads a criteria file containing malformed grader configuration, it validates at load time. If the file is referenced in an active evaluation, that evaluation fails. If the file is not used, the validation error is logged but doesn't affect other evaluations.

**Name uniqueness.** Grader names must be unique within a criteria file. You can use the same grader type multiple times with different `name` values and different configurations.

## Unified Schema

All graders are configured using a flat YAML structure with a `type:` discriminator:

```yaml
graders:
  - name: <string>              # Required: human-readable name, unique in this file
    type: <string>              # Required: grader type (prompt, output_check, file, etc.)
    weight: <float>             # Optional: scoring weight 0.0–1.0 (default: 1.0)
    when: <map[string]string>   # Optional: property-based applicability conditions

    # ── For type=prompt ONLY ─────────────────────────────────────────
    prompt: <string>            # Optional preamble shown to the LLM judge
    checks:                     # List of individual pass/fail items the judge must
      - <string>                # evaluate. Each non-empty string becomes one line in
      - <string>                # the rendered review criteria AND one Point in the
      # ...                     # resulting GraderResult. At least one of prompt/
                                # checks must be non-empty.

    # ── For every OTHER type ─────────────────────────────────────────
    details: <object>           # Required: type-specific configuration (shape varies)
```

> **v4 invariant — "every grader emits Points":** Each `prompt` grader produces one
> `Point` per entry in `checks:` (or a single Point when only `prompt:` is set). The
> pre-v4 magic that split a single prompt blob into multiple checks by parsing
> bullets is gone for YAML graders — list each sub-check explicitly under `checks:`.

The `type:` field is the discriminator:
- `prompt` — LLM-based review grader
- `output_check` — Workspace file and size checks
- `program` — Run a custom program/script and evaluate exit code
- `tool` — Unified tool-perspective checks (canonical)
- `action_sequence` — Verify expected action sequence

**Deprecated types** (still functional, emit warnings at load time):
- `file` — use `output_check` with `require_files` instead
- `behavior` — use `tool` instead
- `tool_constraint` — use `tool` instead
- `tool_usage` — use `tool` instead

**Engine-internal types** (not valid in user YAML):
- `prompt_review` — AI review panel orchestration (used by the engine, not user-configurable)

### Example: Mixed Prompt and Typed Graders

```yaml
graders:
  # LLM prompt grader — `prompt:` and `checks:` are TOP-LEVEL fields,
  # not inside `details:` (validation rejects details for type=prompt).
  - name: Code Quality Review
    type: prompt
    weight: 0.7
    prompt: "Review the generated Python code against each of the following:"
    checks:
      - Readable variable names and small focused functions
      - Errors are caught and handled appropriately
      - Adheres to PEP 8 style guide

  # File existence check
  - name: Main File Exists
    type: file
    weight: 0.2
    details:
      path: main.py
      must_exist: true

  # Output check (file count and size)
  - name: Generated Some Output
    type: output_check
    weight: 0.1
    details:
      min_files: 1
      min_bytes_per_file: 10

  # Conditional: only run for Python service + data-plane
  - name: Python Async Patterns
    type: prompt
    weight: 0.5
    when:
      language: python
      plane: data-plane
    prompt: Check that async functions use await correctly.
```

### Field Reference

| Field     | Type                | Required | Default | Description                                                                     |
|-----------|---------------------|----------|---------|---------------------------------------------------------------------------------|
| `name`    | string              | yes      | —       | Human-readable name. Must be unique within the grader list.                     |
| `type`    | string              | yes      | —       | Grader type identifier. See Grader Types below.                                 |
| `weight`  | float64             | no       | 1.0     | Scoring weight (0.0–1.0). Normalized across all graders in aggregation.         |
| `prompt`  | string              | type=prompt only | — | Preamble / instructions for the LLM judge. Must be empty for typed graders. |
| `checks`  | list&lt;string&gt;  | type=prompt only | — | One pass/fail item per entry. Each becomes one Point. Must be empty for typed graders. |
| `details` | object              | typed graders only | — | Type-specific configuration. Must be empty for `type: prompt`.       |
| `when`    | map\[string\]string | no       | —       | Property conditions for applicability. All entries must match (AND logic).       |

### Applicability (`when`)

The `when` field applies a grader conditionally based on prompt properties using exact, case-insensitive string matching. All specified keys must match for the grader to run. If `when` is omitted, the grader applies to all prompts.

```yaml
when:
  language: python        # Only apply to Python prompts
  service: key-vault      # AND only Key Vault service
```

Property keys are not restricted to a fixed set—any prompt property can be used (language, service, plane, category, tags).

## Grader Types

Detailed documentation for each grader type:

| Type               | Purpose                                           | Page                                      |
|--------------------|---------------------------------------------------|-------------------------------------------|
| `prompt`           | LLM-based review and evaluation                   | [prompt.md](./prompt.md)                  |
| `output_check`     | File count and size verification                  | [output_check.md](./output_check.md)       |
| `file`             | File existence and content pattern matching       | [file.md](./file.md)                      |
| `program`          | Run external program and evaluate exit code       | [program.md](./program.md)                |
| `behavior`         | Tool and action constraint checking               | [behavior.md](./behavior.md)              |
| `action_sequence`  | Expected action sequence verification             | [action_sequence.md](./action_sequence.md)|
| `tool_constraint`  | Tool usage and call count constraints             | [tool_constraint.md](./tool_constraint.md)|
| `prompt_review`    | AI review panel orchestration (engine-internal; not user-configurable) | [prompt_review.md](./prompt_review.md)    |

## Score Aggregation

After all graders run, scores are aggregated using a weighted average:

```
score = Σ(grader_score × grader_weight) / Σ(grader_weight)
```

- Each grader produces a score between 0 and 1.
- Weights default to 1.0 if not specified.
- The final aggregated score is normalized to 0–1.
- Failed graders score 0; passed graders score 1 (typical behavior, varies by grader type).

Graders do not gate. A failed grader contributes 0 to the weighted average but does not stop evaluation or fail other graders.

## Validation

Criteria files are validated at load time:
- Each grader entry must have a valid `type` from the supported kinds.
- The `details` object must match the schema for that grader type.
- Grader names must be unique within the file.
- Unknown fields in a grader entry are rejected (strict schema validation).

If a criteria file with validation errors is **not** used in any active evaluation, the error is logged but doesn't affect other evaluations. If the file **is** used, the evaluation fails with a validation error.

## Next Steps

- **Getting started**: See [../getting-started.md](../getting-started.md) for end-to-end examples.
- **Configuration guide**: See [../configuration.md](../configuration.md) for prompt and config file structure.
- **Criteria authoring**: See [../criteria-authoring.md](../criteria-authoring.md) for advanced patterns (hierarchical `when`, groups, isolation).

# Grader Reference

## Overview

hyoka's grading system evaluates generated code and agent behavior using composable, single-concern graders. Each grader checks exactly one aspect: file creation, build success, LLM-based code review, tool usage patterns, or session activity. Results are aggregated by weighted scoring to produce a final evaluation report.

## Key Concepts

**Graders never gate evaluations.** Each grader runs independently and reports its result. Failed graders don't stop other graders or halt the evaluation—all graders run and contribute to the final weighted score.

**Load-time validation.** When hyoka loads a criteria file containing malformed grader configuration, it validates at load time. If the file is referenced in an active evaluation, that evaluation fails. If the file is not used, the validation error is logged but doesn't affect other evaluations.

**Name uniqueness.** Grader names must be unique within a criteria file. You can use the same grader type multiple times with different `name` values and different configurations.

## Unified Schema

All graders are configured using a flat YAML structure with a `type:` discriminator. Canonical graders (prompt, program, workspace, tool, activity) flatten their checks to the **top level** of each grader entry, not inside a nested `details:` object:

```yaml
graders:
  - name: <string>              # Required: human-readable name, unique in this file
    type: <string>              # Required: grader type (prompt, program, workspace, tool, activity)
    weight: <float>             # Optional: scoring weight 0.0–1.0 (default: 1.0)
    when: <map[string]string>   # Optional: property-based applicability conditions

    # ── For type=prompt ──────────────────────────────────────────────
    prompt: <string>            # Optional: preamble shown to the LLM before checks
    checks:                     # Each string becomes one independent pass/fail check
      - <string>                # judged by the LLM (one Point per check in results)

    # ── For type=program ─────────────────────────────────────────────
    checks:                     # Each command is one check
      - kind: command           # Only kind supported: "command"
        command: <string>       # Command to execute
        args: [<string>, ...]   # Command arguments (optional)
        timeout: <int>          # Timeout in seconds (optional)

    # ── For type=workspace ───────────────────────────────────────────
    checks:                     # Each check validates file state or delta
      - kind: <string>          # One of: require_to_create, forbidden_to_create,
                                # required_to_update, required_to_delete,
                                # forbidden_to_delete, file

    # ── For type=tool ────────────────────────────────────────────────
    checks:                     # Each check validates tool usage patterns
      - kind: <string>          # One of: tool_used, tool_not_used,
                                # any_from_group, none_from_group

    # ── For type=activity ────────────────────────────────────────────
    checks:                     # Each check validates session activity
      - kind: <string>          # One of: turn_limit, action_count,
                                # tool_call_count, contains_subsequence,
                                # contains_action, excludes_action, terminated_by
```

> **Current schema (flat checks):** All canonical graders now have `checks:` at the top level of the grader entry, not nested under `details:`. This flattening makes the schema more uniform and easier to read. Engine-internal graders like `prompt_review` continue to use their own internal structure.

### Example: All Five Canonical Grader Types

```yaml
graders:
  # 1. LLM-based prompt grader
  - name: Code Quality
    type: prompt
    weight: 0.7
    prompt: "Review the code against:"
    checks:
      - Readable variable names and focused functions
      - Proper error handling

  # 2. Workspace (file state) checks
  - name: Files Created
    type: workspace
    weight: 0.5
    checks:
      - kind: require_to_create
        files: [main.py, README.md]
      - kind: file
        name: main.py
        state: present
        min_bytes: 100

  # 3. Tool usage checks
  - name: Tool Usage
    type: tool
    weight: 0.5
    checks:
      - kind: tool_used
        tool: bash
      - kind: tool_not_used
        tool: dangerous_tool

  # 4. Session activity checks
  - name: Session Behavior
    type: activity
    weight: 0.3
    checks:
      - kind: turn_limit
        max: 25
      - kind: contains_action
        type: message
      - kind: terminated_by
        equals: completed

  # 5. Program execution checks
  - name: Build Success
    type: program
    weight: 0.4
    checks:
      - kind: command
        command: go
        args: [build, ./...]
        timeout: 30
```

### Field Reference

| Field     | Type                | Required | Default | Description                                                             |
|-----------|---------------------|----------|---------|-------------------------------------------------------------------------|
| `name`    | string              | yes      | —       | Human-readable name, unique within this file.                           |
| `type`    | string              | yes      | —       | Grader type: `prompt`, `program`, `workspace`, `tool`, `activity`.     |
| `weight`  | float64             | no       | 1.0     | Scoring weight for aggregation (0.0–1.0).                              |
| `when`    | map\[string\]string | no       | —       | Conditional applicability by prompt properties (AND logic).             |
| `prompt`  | string              | type=prompt only | — | LLM preamble/instructions (optional if `checks:` is set).              |
| `checks`  | list               | yes (for canonical) | — | Type-specific checks (each canonical grader type defines its own schema). |

### Applicability (`when`)

The `when` field applies a grader conditionally based on prompt properties using exact, case-insensitive string matching. All specified keys must match for the grader to run. If `when` is omitted, the grader applies to all prompts.

```yaml
when:
  language: python        # Only apply to Python prompts
  service: key-vault      # AND only Key Vault service
  plane: data-plane       # AND only data-plane prompts
```

Property keys are not restricted to a fixed set—any prompt property can be used (language, service, plane, category, tags).

#### Config-aware properties

In addition to prompt-derived properties, the engine injects properties derived from the **eval config under test**, so a grader can gate itself to configs that actually load the tools it depends on. These are the keys available in `when:`:

| Key                       | Value    | Source                                            |
|---------------------------|----------|---------------------------------------------------|
| `generator`               | model name | `generator.model` from the eval config          |
| `config`                  | config name | `name:` field of the eval config               |
| `skill:<name>`            | `"true"` | one entry per `generator.tools` with `type: skill` |
| `mcp_server:<name>`       | `"true"` | one entry per `generator.tools` with `type: mcp`   |
| `plugin:<name>`           | `"true"` | one entry per `generator.tools` with `type: plugin` |

**Example — gate a `tool_used` grader to configs that load the `azure` MCP server:**

```yaml
graders:
  - name: uses-azure-list-resources
    kind: tool
    when:
      "mcp_server:azure": "true"   # quote keys containing ':' (YAML requirement)
    config:
      checks:
        - kind: tool_used
          tool: list-resources
          source: mcp
          mcp_server: azure
```

On a config that doesn't include the `azure` MCP server, this grader is silently skipped instead of failing every eval.

> **YAML quoting.** Map keys containing `:` must be quoted (e.g. `"mcp_server:azure": "true"`). Unquoted, YAML parses `mcp_server:azure` as a nested mapping. The `:`-prefixed namespace is reserved for engine-injected config props; do not use it for prompt frontmatter.

## Canonical Grader Types

hyoka defines five canonical grader types, each with a flat `checks:` list at the top level:

| Type      | Purpose                                    | Check Kinds                                                   | Details               |
|-----------|--------------------------------------------|---------|-----------------------------------------|
| `prompt`  | LLM-judged subjective review              | Each string in `checks:` is one independent check          | [prompt.md](./prompt.md)   |
| `program` | Execute command and grade exit code        | `command` (only kind: runs command, exit 0 = pass)         | [program.md](./program.md) |
| `workspace` | Validate file creation/modification state  | `require_to_create`, `forbidden_to_create`, `required_to_update`, `required_to_delete`, `forbidden_to_delete`, `file` | [workspace.md](./workspace.md) |
| `tool`    | Validate tool usage patterns               | `tool_used`, `tool_not_used`, `any_from_group`, `none_from_group` | [tool.md](./tool.md) |
| `activity` | Validate session activity and action log   | `turn_limit`, `action_count`, `tool_call_count`, `contains_subsequence`, `contains_action`, `excludes_action`, `terminated_by` | [activity.md](./activity.md) |

**Engine-internal type** (not user-configurable):
- `prompt_review` — AI review panel orchestration (used internally; see [prompt_review.md](./prompt_review.md) for reference)

## Score Aggregation

After all graders run, scores are aggregated using a weighted average:

```
score = Σ(grader_score × grader_weight) / Σ(grader_weight)
```

- Each grader produces a score between 0 and 1.
- Weights default to 1.0 if not specified.
- The final aggregated score is normalized to 0–1.
- Failed graders score 0; passed graders score 1 (typical behavior; varies by grader type).

Graders do not gate. A failed grader contributes 0 to the weighted average but does not stop evaluation or fail other graders.

## Validation

Criteria files are validated at load time:
- Each grader entry must have a valid `type` from the supported kinds.
- The `checks` array (for canonical types) must match the expected schema for that type.
- Grader names must be unique within the file.
- Unknown fields in a grader entry are rejected (strict schema validation).

If a criteria file with validation errors is **not** used in any active evaluation, the error is logged but doesn't affect other evaluations. If the file **is** used, the evaluation fails with a validation error.

## Deprecated Grader Types

The following legacy grader types are no longer supported. If you encounter them in existing configs, migrate to the canonical types listed:

| Legacy Type      | Replacement Guidance                                                 |
|------------------|----------------------------------------------------------------------|
| `file`           | Use `workspace` grader with `kind: file` checks instead              |
| `output_check`   | Use `workspace` grader with `kind: require_to_create`, `forbidden_to_create` checks |
| `behavior`       | Use `activity` grader for session behavior; use `tool` for tool constraints |
| `action_sequence` | Use `activity` grader with `kind: contains_subsequence` and `contains_action` checks |
| `tool_constraint` | Use `tool` grader with `kind: tool_used`, `tool_not_used` checks    |
| `tool_usage`     | Use `tool` grader instead                                            |

## Next Steps

- **Getting started**: See [../getting-started.md](../getting-started.md) for end-to-end examples.
- **Configuration guide**: See [../configuration.md](../configuration.md) for prompt and config file structure.
- **Criteria authoring**: See [../criteria-authoring.md](../criteria-authoring.md) for advanced patterns (hierarchical `when`, groups, isolation).

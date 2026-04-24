# Grader Config YAML Schema (LEGACY — superseded)

> ⚠️ **This document described the pre-v4 grader schema** (`kind:` discriminator,
> `config:` payload, `gate: true` semantics). That schema is no longer accepted
> by the loader. The current schema uses a flat `type:` discriminator, a
> per-type `details:` payload, and `prompt:` + `checks:` top-level fields for
> the `prompt` grader. **No grader gates evaluation.**

The authoritative grader schema reference now lives at:

- **[docs/graders/index.md](./graders/index.md)** — Unified schema, field
  reference, applicability (`when`), `checks:` / `prompt:` rules, score
  aggregation, validation behavior.
- **Per-type pages** under [docs/graders/](./graders/):
  - [`prompt`](./graders/prompt.md) — LLM-judged checks (top-level `prompt:` +
    `checks:`).
  - [`output_check`](./graders/output_check.md) — workspace file/size knobs.
  - [`file`](./graders/file.md) — file existence + content pattern.
  - [`program`](./graders/program.md) — exit-code-based external checks.
  - [`behavior`](./graders/behavior.md), [`action_sequence`](./graders/action_sequence.md),
    [`tool_constraint`](./graders/tool_constraint.md) — tool/turn constraints.
  - [`prompt_review`](./graders/prompt_review.md) — engine-internal review panel
    (not user-configurable in YAML).

## Migration Cheat Sheet (legacy → current)

| Pre-v4 (REMOVED)                                | Current schema                                                |
|-------------------------------------------------|---------------------------------------------------------------|
| `kind: <type>`                                  | `type: <type>`                                                |
| `config: { ... }`                               | `details: { ... }` (typed graders) — empty for `type: prompt` |
| `kind: prompt` + `config.rubric: <text>`        | `type: prompt` + top-level `prompt:` and/or `checks:`         |
| Splitting a single prompt into bullets implicitly | Each item must be listed explicitly under `checks:`         |
| `gate: true`                                    | (removed — no grader gates evaluation; failed graders score 0 and are reported) |

# Grader Config YAML Schema (LEGACY — superseded)

> ⚠️ **This document is OBSOLETE.** The schema described here (v3, with `kind:`, `config:`, `gate:`, and `details:`) was completely replaced in hyoka v0.4.  
> This page is kept for historical reference only.

**Use these instead:**
- **[docs/graders/index.md](./graders/index.md)** — Current unified grader schema with five canonical types
- **Per-type pages** under [docs/graders/](./graders/):
  - [`prompt`](./graders/prompt.md) — LLM-judged review criteria
  - [`program`](./graders/program.md) — External command execution
  - [`workspace`](./graders/workspace.md) — File creation/modification/deletion validation
  - [`tool`](./graders/tool.md) — Tool usage patterns
  - [`activity`](./graders/activity.md) — Session activity and action sequences
  - [`prompt_review`](./graders/prompt_review.md) — Engine-internal multi-model review panel

## Removed in v0.4

The following grader kinds no longer exist. Use their replacements:

| Removed Kind | Reason | Replacement |
|--|--|--|
| `output_check` | Too specific (files + bytes only) | `workspace` grader with full delta support |
| `file` | Superseded by delta awareness | `workspace` grader |
| `behavior` | Split into tool/activity graders | `tool` or `activity` grader |
| `action_sequence` | Renamed to match modern semantics | `activity` grader with `contains_subsequence` check |
| `tool_constraint` | Consolidated into tool grader | `tool` grader |
| `tool_usage` | Consolidated into tool grader | `tool` grader |

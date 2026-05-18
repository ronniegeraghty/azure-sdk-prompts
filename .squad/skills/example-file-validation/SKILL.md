---
name: example-file-validation
description: Validating in-repo prompt/config/criteria example files against the live schemas, including paths the `hyoka validate` CLI does not cover.
---

# Example-file validation

`hyoka validate` only loads `./prompts`, `./configs`, and `./criteria` (the
runtime paths). It does **not** validate `examples/`, `.hyoka/` seed
templates, or doc snippets — drift can sit in those for releases.

## Where drift hides

| Path                                    | Validated by `hyoka validate`? |
|-----------------------------------------|--------------------------------|
| `prompts/`, `configs/`, `criteria/`     | ✅ Yes                         |
| `.hyoka/{prompts,configs,criteria}/`    | ✅ When discovered as project root |
| `examples/`                             | ❌ No                          |
| `docs/**` snippets                      | ❌ No                          |
| `hyoka/cmd/init.go` const string seeds  | ❌ No (compiled in)            |
| Per-prompt `criteria.yaml` (if any)     | ❌ Not via the validate cmd    |

## Schema sources of truth

- **Prompt frontmatter:** `hyoka/internal/prompt/parser.go` (`frontmatter` struct)
  + `types.go`. **All** typed string metadata (service, language, plane, etc.)
  lives under `properties:`. Top-level fields are limited to: `id`, `tags`,
  `project_context`, `starter_project`, `reference_answer`, `timeout`,
  `max_session_actions`, `max_turns`, `expected_packages`, `expected_tools`,
  `group`, `properties`, `prompt_text`, `evaluation_criteria`. Nothing else
  is consumed — extra keys are silently ignored.
- **Config YAML:** `hyoka/internal/config/config.go` (`ConfigFile`/`ToolConfig`).
  Top-level wrapper is `configs: [ ... ]` — **a single ToolConfig at root
  will not parse**. `model:` (singular) and `models:` (plural) both work;
  `models:` wins when both set.
- **Criteria YAML (v4 unified):** `hyoka/internal/criteria/config.go`
  (`UnifiedGraderEntry`). Flat `type:` discriminator. For `type: prompt`,
  use top-level `prompt:` and/or `checks:` — `details:` is **rejected**.
  For typed graders (file/program/behavior/action_sequence/tool_constraint/
  output_check), use `details:`. `prompt_review` is **engine-internal**
  and not a valid YAML `type:` value.

## Checking files outside `hyoka validate`

Drop a tiny audit binary inside the module (so it can import internal
packages) and exercise the loaders directly:

```go
// hyoka/cmd_audit/main.go (delete after run)
package main
import (
  "github.com/.../hyoka/internal/config"
  "github.com/.../hyoka/internal/criteria"
  "github.com/.../hyoka/internal/prompt"
)
// ... call config.Load, criteria.LoadUnifiedFile, prompt.ParsePromptFile ...
```

Run with `go run ./hyoka/cmd_audit/`. Files outside `hyoka/` cannot import
the `internal/` packages — putting the binary under `hyoka/` works around
the internal-import rule.

## Translation gotchas

The unified loader silently coerces "no `type:` + has top-level `prompt:` +
no `details:`" → `type: prompt` (`criteria/bundle.go:84-97`). This makes
pre-v4 criteria files validate **as if** they were already migrated. When
auditing, always grep for graders missing an explicit `type:` — they
"work" but are stale by today's standards and break the moment the
translation is removed.

## v4 invariants (do not regress)

- Every grader must emit Points (`graders/grader.go`). A grader with zero
  Points is a configuration error.
- `prompt:` + `checks:` for `type: prompt` are **top-level**, never under
  `details:`. The pre-v4 magic that split a single `prompt:` blob into
  per-bullet checks is gone for YAML graders.
- No grader gates evaluation. The `gate:` field is removed from the
  unified schema; failed graders score 0 and are reported.

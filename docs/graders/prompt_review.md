# Prompt Review Grader

The `prompt_review` grader orchestrates hyoka's multi-model AI review panel. It is
**created internally by the engine** and is not a valid `type:` value in user-
authored criteria YAML files. Schema validation rejects entries with
`type: prompt_review` at load time.

## When You'll See It

`prompt_review` results appear in evaluation reports whenever a config defines a
multi-model `reviewer:` panel. The engine constructs one `PromptReviewGrader`
instance per evaluation, runs all reviewer models against the generated code, and
emits a single `GraderResult` whose Points correspond to per-criterion verdicts.

For user-defined LLM checks in criteria files, use the [`prompt`](./prompt.md)
grader instead — it is the user-facing equivalent and supports `prompt:` +
`checks:` for explicit per-Point control.

## Result Structure

The `prompt_review` grader produces:
- **Per-criterion Points**: One Point per evaluation criterion, weighted by the
  criterion's max score.
- **`ReviewExtras`**: Summary, issues, strengths, and per-model panel results
  (rendered by the site under `grader-extras/ReviewExtras.tsx`).
- **Aggregated score**: Weighted average across all panel models.

## Notes

- **Not user-configurable**: There is no valid YAML schema for `type: prompt_review`
  in criteria files. Validation fails at load time if you try.
- **Configured via reviewer panel**: Add reviewer models under `reviewer.models:`
  in the config YAML; the engine wires up the `PromptReviewGrader` automatically.
- **Prefer `prompt` for custom checks**: For ad-hoc LLM-judged criteria, use the
  [`prompt`](./prompt.md) grader.

See [`prompt.md`](./prompt.md) for user-facing LLM review configuration and
[`../configuration.md`](../configuration.md) for `reviewer.models:` syntax.

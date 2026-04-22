# Prompt Review Grader

The `prompt_review` grader is an internal, advanced grader that orchestrates the AI review panel.

## When to Use

This grader is primarily used internally by hyoka's evaluation engine to coordinate multi-model review panels. Most users should use the [`prompt`](./prompt.md) grader instead, which provides a simpler interface for LLM-based code review.

Use `prompt_review` only if you need:
- Custom review panel orchestration
- Multiple reviewers with specific configurations
- Internal evaluation workflows
- Advanced model selection and aggregation

## Configuration

```yaml
graders:
  - name: Review Panel
    type: prompt_review
    details:
      # Internal configuration — see internal/graders/prompt_review_grader.go
      {}
```

## Result Structure

The `prompt_review` grader produces:
- **Aggregated score**: Consensus score from all review models
- **Per-model scores**: Individual scores from each reviewer
- **Consolidated feedback**: Merged insights across all reviewers
- **Review details**: Full reasoning and notes from each model

## Notes

- **Not recommended for user criteria**: This grader is maintained for engine compatibility.
- **Prefer `prompt` grader**: For custom LLM reviews, use the simpler [`prompt`](./prompt.md) grader instead.
- **Engine integration**: Automatically invoked during evaluation if configured at the config level.

See [`prompt.md`](./prompt.md) for user-facing LLM review configuration.

TODO: Document the internal API and integration points for advanced use cases.

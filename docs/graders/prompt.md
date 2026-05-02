# Prompt Grader

The `prompt` grader sends generated code and a custom evaluation prompt to an LLM for subjective assessment. Use this grader for code quality, architecture, style, and other concerns that require human-like judgment.

## When to Use

- **Code quality reviews**: Readability, maintainability, adherence to conventions
- **Architectural patterns**: Use of design patterns, separation of concerns
- **Security considerations**: Input validation, error handling, safe practices
- **Documentation and comments**: Clarity of inline and external docs
- **Language idioms**: Proper use of language-specific patterns and features

For objective, deterministic checks (file existence, build success, tool constraints), use canonical graders like [`workspace`](./workspace.md), [`program`](./program.md), or [`tool`](./tool.md) instead.

## Configuration

The `prompt` grader uses TWO top-level fields (NOT inside `details:` — validation
rejects `details:` for `type: prompt`):

- `prompt:` — optional preamble shown to the LLM judge before the numbered checks.
- `checks:` — list of individual pass/fail items. Each non-empty string becomes one
  line in the rendered review criteria AND one Point in the resulting GraderResult.

At least one of `prompt:` or `checks:` must be non-empty.

```yaml
graders:
  - name: Code Quality Review
    type: prompt
    weight: 0.7
    prompt: |
      Review the generated Python code against each of the following checks.
      For each check, return PASS or FAIL with one-sentence reasoning.
    checks:
      - Variable naming is clear and functions are small and focused
      - Errors are caught and handled appropriately (no bare except)
      - Code adheres to PEP 8 style conventions
      - Type annotations are present and correct on public APIs
```

### Field Reference

| Field    | Type            | Required             | Description                                                              |
|----------|-----------------|----------------------|--------------------------------------------------------------------------|
| `prompt` | string          | one of prompt/checks | Preamble text shown to the judge before the numbered checks.             |
| `checks` | list&lt;string&gt; | one of prompt/checks | Individual pass/fail items. Each becomes one Point in `GraderResult`.    |

> **v4 invariant:** Each item in `checks:` becomes one Point. The pre-v4 magic that
> split a single `prompt:` blob into multiple checks by parsing bullets is **gone**
> for YAML graders — list each sub-check explicitly under `checks:`.

## Example

```yaml
graders:
  - name: TypeScript Best Practices
    type: prompt
    weight: 0.6
    when:
      language: js-ts
    prompt: "Review the TypeScript code against each of the following:"
    checks:
      - Strict mode and type checking are enabled
      - Optional types are used for nullable values (no any-typed null handling)
      - Async/await is used (not raw Promises with .then chains)
      - Imports are organized (external, internal, relative)
      - Public functions have JSDoc-style documentation

  - name: Go Idioms Check
    type: prompt
    weight: 0.5
    when:
      language: go
    prompt: "Evaluate this Go code for idiomatic usage:"
    checks:
      - Errors are returned and explicitly checked (no panic in library code)
      - defer is used for cleanup paths
      - Interfaces are defined at the consumer, not the producer
      - Concurrency uses goroutines + channels (not shared mutable state)
      - Package layout follows standard Go conventions
```

## Result Structure

Each `prompt` grader result includes:
- **Score**: LLM-assigned numeric score (typically 0–10, normalized to 0–1 for aggregation)
- **Reasoning**: The LLM's textual explanation of the score
- **Details**: Full review panel response and model used

Results are visible in the evaluation report under `grader_results`.

## Notes

- **One model per grader**: Each `prompt` grader instance runs exactly one LLM model. If you want multiple models reviewing the same code, create multiple `prompt` graders with different LLM configurations.
- **Reproducibility**: LLM reviews are not deterministic. Identical code may receive slightly different scores across multiple runs.
- **Scoring scale**: Different LLMs may use different scoring conventions (0–5, 0–10, etc.). Normalize to 0–1 in aggregation.
- **Conditional evaluation**: Use `when:` to apply prompt graders only to relevant prompts (e.g., Python-only code quality checks).

See [index.md](./index.md#applicability-when) for `when:` syntax and conditional grader application.

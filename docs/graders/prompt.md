# Prompt Grader

The `prompt` grader sends generated code and a custom evaluation prompt to an LLM for subjective assessment. Use this grader for code quality, architecture, style, and other concerns that require human-like judgment.

## When to Use

- **Code quality reviews**: Readability, maintainability, adherence to conventions
- **Architectural patterns**: Use of design patterns, separation of concerns
- **Security considerations**: Input validation, error handling, safe practices
- **Documentation and comments**: Clarity of inline and external docs
- **Language idioms**: Proper use of language-specific patterns and features

For objective, deterministic checks (file existence, build success, tool constraints), use typed graders like [`output_check`](./output_check.md), [`file`](./file.md), or [`tool_constraint`](./tool_constraint.md) instead.

## Configuration

```yaml
graders:
  - name: Code Quality Review
    type: prompt
    weight: 0.7
    details:
      prompt: |
        Review the generated Python code for these aspects:
        
        1. **Readability**: Is variable naming clear? Are functions small and focused?
        2. **Error handling**: Are exceptions caught and handled appropriately?
        3. **PEP 8 compliance**: Does the code follow PEP 8 style guide?
        4. **Type hints**: Are type annotations present and correct?
        
        Provide a score from 0 (poor) to 10 (excellent) and brief reasoning.
```

### `details` Schema

| Field    | Type   | Required | Description                                    |
|----------|--------|----------|------------------------------------------------|
| `prompt` | string | yes      | Evaluation rubric/instructions for the LLM.    |

The `prompt` field contains the full evaluation criteria. It is sent to an LLM reviewer along with the generated code to produce a scored review.

## Example

```yaml
graders:
  - name: TypeScript Best Practices
    type: prompt
    weight: 0.6
    when:
      language: js-ts
    details:
      prompt: |
        Review this TypeScript code for adherence to best practices:
        
        - Proper use of strict mode and type checking
        - Null safety (proper use of optional types)
        - Async/await patterns
        - Import organization
        - Function documentation

  - name: Go Idioms Check
    type: prompt
    weight: 0.5
    when:
      language: go
    details:
      prompt: |
        Evaluate this Go code for idiomatic usage:
        
        - Error handling (defer, explicit error checks)
        - Interface design
        - Concurrency patterns (goroutines, channels)
        - Package organization
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

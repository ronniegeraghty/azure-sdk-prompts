# Tool Constraint Grader

The `tool_constraint` grader enforces constraints on tool availability, usage, and call counts.

## When to Use

- Verify specific tools were used (or not used)
- Enforce minimum/maximum call counts for tools
- Test behavior under tool constraints
- Validate tool-specific workflows

## Configuration

```yaml
graders:
  - name: Tool Usage Constraints
    type: tool_constraint
    weight: 0.3
    details:
      required:
        - bash
        - file_search
      forbidden:
        - dangerous_tool
      min_calls:
        bash: 2
      max_calls:
        bash: 10
```

### `details` Schema

| Field      | Type              | Required | Description                              |
|------------|-------------------|----------|------------------------------------------|
| `required` | []string          | no       | Tools that must be used at least once   |
| `forbidden` | []string         | no       | Tools that must not be used             |
| `min_calls` | map[string]int    | no       | Minimum call counts per tool            |
| `max_calls` | map[string]int    | no       | Maximum call counts per tool            |

## Example

```yaml
graders:
  - name: Efficient Tool Usage
    type: tool_constraint
    weight: 0.4
    details:
      required:
        - bash
        - file_search
      max_calls:
        bash: 5
        file_search: 10

  - name: No Exec in Production Paths
    type: tool_constraint
    weight: 0.2
    details:
      forbidden:
        - exec_shell_escape
```

## Result Structure

- **Pass/Fail**: Binary result based on constraint violations
- **Tool usage summary**: Call counts per tool
- **Constraint violations**: Which constraints were violated
- **Required tools missing**: Tools that should have been used
- **Forbidden tools used**: Tools that should not have been used

## Notes

- **Tool names**: Must match tool IDs in the SDK
- **Call counts**: Tallied across entire session
- **All-or-nothing constraints**: Violations fail the grader; no partial credit
- **No soft limits (v1)**: min/max are strict cutoffs

TODO: Add SDK-specific tool naming conventions and cross-platform compatibility examples.

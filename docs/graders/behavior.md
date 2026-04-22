# Behavior Grader

The `behavior` grader checks constraints on tool usage and session actions.

## When to Use

- Verify required tools were used
- Forbid use of specific tools
- Limit total number of session turns/actions
- Check tool invocation patterns

## Configuration

```yaml
graders:
  - name: Required Tools Used
    type: behavior
    weight: 0.3
    details:
      required_tools:
        - bash
        - file_search
      forbidden_tools:
        - dangerous_tool
      max_turns: 20
```

### `details` Schema

| Field             | Type     | Required | Default | Description                            |
|-------------------|----------|----------|---------|----------------------------------------|
| `required_tools`  | []string | no       | —       | Tools that must be used at least once |
| `forbidden_tools` | []string | no       | —       | Tools that must not be used           |
| `max_turns`       | int      | no       | —       | Maximum session turns allowed         |

## Example

```yaml
graders:
  - name: No Forbidden API Usage
    type: behavior
    weight: 0.4
    details:
      forbidden_tools:
        - exec
        - shell_escape

  - name: Uses Required Debugging Tools
    type: behavior
    weight: 0.2
    details:
      required_tools:
        - debug_console
        - logging
```

## Result Structure

- **Pass/Fail**: Binary result based on constraint violations
- **Used tools**: List of tools actually invoked
- **Missing tools**: Required tools not invoked
- **Forbidden tools used**: Tools that should not have been used
- **Turn count**: Total session turns recorded

## Notes

- **Tool names**: Must match tool IDs in the Copilot SDK
- **No partial credit (v1)**: Constraints are all-or-nothing
- **Turn limits**: Applied at session level, not per-grader

TODO: Add detail on tool matching, case sensitivity, and SDK-specific tool names.

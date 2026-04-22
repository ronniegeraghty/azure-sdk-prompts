# Action Sequence Grader

The `action_sequence` grader verifies that the agent followed an expected sequence of actions.

## When to Use

- Validate step-by-step problem-solving approach
- Ensure diagnostic phase before implementation
- Check for specific action ordering
- Verify methodology compliance

## Configuration

```yaml
graders:
  - name: Follows Diagnostic First Approach
    type: action_sequence
    weight: 0.3
    details:
      expected_actions:
        - "analyze"
        - "design"
        - "implement"
```

### `details` Schema

| Field              | Type     | Required | Description                            |
|--------------------|----------|----------|----------------------------------------|
| `expected_actions` | []string | yes      | Sequence of action names to verify    |

## Example

```yaml
graders:
  - name: Problem Analysis Before Coding
    type: action_sequence
    weight: 0.4
    details:
      expected_actions:
        - "read_problem"
        - "analyze_requirements"
        - "design_solution"
        - "write_code"
        - "test_solution"
```

## Result Structure

- **Pass/Fail**: Binary result based on sequence match
- **Actions observed**: Actual action sequence from session
- **Match details**: Which steps were followed in order

## Notes

- **Exact sequence**: All actions must appear in the specified order
- **Action names**: Must match action identifiers in the Copilot SDK
- **No branching (v1)**: Sequence is linear; conditional branches not yet supported
- **Partial matches**: Only exact sequence matches are accepted

TODO: Add examples of different action types and integration with Copilot SDK action identifiers.

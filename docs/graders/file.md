# File Grader

The `file` grader checks for the existence of specific files and optionally matches file content against a regex pattern.

## When to Use

- Verify a specific file was created (e.g., `main.py`, `README.md`)
- Check file content against a pattern (e.g., "contains `def main():`")
- Verify file absence (inverted check)
- Gate on required output artifacts

## Configuration

```yaml
graders:
  - name: Main File Exists
    type: file
    weight: 0.3
    details:
      path: main.py
      pattern: "def main"
      must_exist: true
```

### `details` Schema

| Field        | Type   | Required | Default | Description                                    |
|--------------|--------|----------|---------|------------------------------------------------|
| `path`       | string | yes      | —       | File path relative to workspace root.          |
| `pattern`    | string | no       | —       | Regex pattern to match in file content.        |
| `must_exist` | bool   | no       | true    | If true, file must exist; if false, must not. |

## Example

```yaml
graders:
  - name: README Documentation
    type: file
    weight: 0.2
    details:
      path: README.md
      must_exist: true

  - name: No Debug Files
    type: file
    weight: 0.1
    details:
      path: debug.log
      must_exist: false
```

## Result Structure

- **Pass/Fail**: Binary result based on file existence and pattern match
- **File details**: Path checked, content match result

## Notes

- **Pattern matching**: Uses Go regex syntax. See [golang.org/pkg/regexp](https://golang.org/pkg/regexp/) for pattern syntax.
- **Relative paths**: Paths are resolved relative to the workspace root.
- **No line-by-line checks (v1)**: Pattern matching operates on entire file content.

TODO: Add more detailed examples and edge cases.

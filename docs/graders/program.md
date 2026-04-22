# Program Grader

The `program` grader runs an external program or script in the workspace and evaluates its exit code.

## When to Use

- Run build commands (e.g., `go build`, `npm run build`)
- Execute linters or format checkers
- Run custom validation scripts
- Check exit code for pass/fail

## Configuration

```yaml
graders:
  - name: Code Builds Successfully
    type: program
    weight: 0.4
    details:
      command: "go"
      args: ["build", "./..."]
      timeout: 30
```

### `details` Schema

| Field   | Type     | Required | Default | Description                              |
|---------|----------|----------|---------|------------------------------------------|
| `command` | string | yes      | —       | Program name or path to execute.        |
| `args`  | []string | no       | —       | Command arguments.                      |
| `timeout` | int    | no       | —       | Execution timeout in seconds.           |

## Example

```yaml
graders:
  - name: Python Tests Pass
    type: program
    weight: 0.5
    when:
      language: python
    details:
      command: "pytest"
      args: ["-q", "tests/"]
      timeout: 60

  - name: Linter Clean
    type: program
    weight: 0.2
    details:
      command: "pylint"
      args: ["*.py"]
      timeout: 30
```

## Result Structure

- **Pass/Fail**: Binary result based on exit code (0 = pass, non-zero = fail)
- **Output**: Captured stdout/stderr from program execution
- **Exit code**: Actual exit code returned

## Notes

- **Exit code semantics**: Exit code 0 = pass, non-zero = fail
- **Working directory**: Program runs in the workspace root
- **Timeout**: If specified, program is terminated if it exceeds timeout seconds
- **Environment**: Program inherits the evaluation environment

TODO: Add details on environment variables, path handling, and cross-platform considerations.

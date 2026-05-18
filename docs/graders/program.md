# Program Grader

The `program` grader runs one or more commands in the evaluation workspace and grades based on exit codes. Each command check produces a pass/fail result (exit 0 = pass, non-zero = fail). The overall grader score is calculated as passed_checks / total_checks.

## When to Use

- Run build commands (e.g., `go build`, `npm run build`, `cargo build`)
- Execute test suites
- Run linters or format checkers
- Validate generated code with custom scripts
- Verify pre/post-generation state with shell commands

## Configuration

The `program` grader uses a top-level `checks:` array. Each check runs one command:

```yaml
graders:
  - name: Build Success
    type: program
    weight: 0.4
    checks:
      - kind: command
        command: go
        args: [build, ./...]
        timeout: 30
```

### `checks` Schema

Each check in `checks:` must have `kind: command`. Only the `command` kind is currently supported.

| Field     | Type     | Required | Default | Description                                              |
|-----------|----------|----------|---------|----------------------------------------------------------|
| `kind`    | string   | yes      | —       | Must be `"command"` (only supported kind).               |
| `command` | string   | yes      | —       | Program name or path to execute (e.g., `go`, `python`, `npm`). |
| `args`    | []string | no       | —       | Command arguments passed to the program.                |
| `timeout` | int      | no       | 30     | Execution timeout in seconds. If exceeded, check fails. |

## Examples

### Basic Build Check

```yaml
graders:
  - name: Python Tests Pass
    type: program
    weight: 0.5
    when:
      language: python
    checks:
      - kind: command
        command: pytest
        args: [-q, tests/]
        timeout: 60

  - name: Linter Clean
    type: program
    weight: 0.2
    checks:
      - kind: command
        command: pylint
        args: [src/]
        timeout: 30
```

### Multiple Checks

```yaml
graders:
  - name: Go Build & Test
    type: program
    weight: 0.6
    when:
      language: go
    checks:
      - kind: command
        command: go
        args: [build, ./...]
        timeout: 30
      - kind: command
        command: go
        args: [test, -race, -v, ./...]
        timeout: 120
      - kind: command
        command: go
        args: [vet, ./...]
        timeout: 30
```

## Result Structure

Each program grader produces:
- **Pass/Fail**: Binary result (one per check) based on exit code
- **Exit code**: Actual exit code returned (0 = pass, non-zero = fail)
- **Output**: Captured stdout and stderr from command execution
- **Duration**: Time taken to execute (in milliseconds)
- **Overall score**: passed_checks / total_checks

Results are visible in evaluation reports under `grader_results`.

## Data Visible to Grader

Program graders can access:
- **Workspace files**: All files present in the workspace (workspace_path + generated files)
- **Exit codes**: Returned by each command
- **Command output**: Stdout and stderr (captured; limited to 500 chars of stderr in result summary)
- **Execution time**: Duration of each command execution

## Notes

- **Exit code semantics**: Exit code 0 = check passes, non-zero = check fails (standard Unix convention)
- **Working directory**: Commands execute in `workspace_path` (the root of the generated agent output)
- **Timeout behavior**: If a command exceeds the timeout, it is terminated (SIGKILL) and the check fails
- **Default timeout**: 30 seconds if not specified
- **Environment inheritance**: Commands inherit the evaluation environment (PATH, shell, etc.)
- **Sequential execution**: Multiple checks run sequentially; failure of one check does not prevent others from running
- **Cross-platform**: Commands must be portable to the deployment environment (Windows/Linux/macOS considerations apply)

## Troubleshooting

- **Command not found**: Ensure the command is available in the workspace or PATH during evaluation
- **Timeout exceeded**: Increase the `timeout` value if legitimate builds/tests are being cut off
- **Permission denied**: Ensure scripts have execute permissions if using relative script paths
- **Working directory issues**: Use relative paths from workspace root; avoid absolute paths

See [index.md](./index.md) for general grader concepts and [../configuration.md](../configuration.md) for config file structure.

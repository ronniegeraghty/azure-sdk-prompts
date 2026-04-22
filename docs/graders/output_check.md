# Output Check Grader

The `output_check` grader verifies that the agent produced files in the workspace and that those files meet minimum size thresholds. Use this grader to ensure the agent actually generated output (as opposed to failing silently or producing empty files).

## When to Use

- **Verify output generation**: Did the agent produce any files at all?
- **Minimum content checks**: Ensure generated files are not empty or trivially small
- **Workspace state validation**: Confirm the workspace was modified with meaningful content
- **Gating minimal output**: Gate evaluations on "at least 1 file with content"

For more specific checks (exact filename, content patterns, build success), use [`file`](./file.md) or [`program`](./program.md) graders instead.

## Configuration

```yaml
graders:
  - name: Generated Output Exists
    type: output_check
    weight: 0.2
    details:
      min_files: 1
      min_bytes_per_file: 10
```

### `details` Schema

| Field                | Type   | Required | Default | Description                                                     |
|----------------------|--------|----------|---------|--------------------------------------------------------------|
| `min_files`          | int    | no       | 1       | Minimum number of files (with content) required to pass.     |
| `min_bytes_per_file` | int64  | no       | 1       | Minimum file size in bytes for a file to count as valid.     |
| `min_total_bytes`    | int64  | no       | 0       | Optional: minimum total bytes across all qualifying files. 0 = disabled. |

### Defaults

- If `min_files` is not set or ≤ 0, defaults to 1 (grader passes if at least one file exists with content).
- If `min_bytes_per_file` is not set or ≤ 0, defaults to 1 (files with 1+ bytes are counted; empty files are ignored).
- If `min_total_bytes` is 0 (default), this check is disabled.

## Examples

### Basic: Any Generated Output

```yaml
graders:
  - name: Generate Some Output
    type: output_check
    weight: 0.1
    details:
      min_files: 1
```

This passes if at least one file with 1+ bytes was created.

### Stricter: Multiple Files with Substantial Content

```yaml
graders:
  - name: Complete Implementation
    type: output_check
    weight: 0.3
    details:
      min_files: 3
      min_bytes_per_file: 100
      min_total_bytes: 1000
```

This passes if:
- At least 3 files exist
- Each file has ≥ 100 bytes
- Total content across all files is ≥ 1000 bytes

### Conditional: Language-Specific Output Checks

```yaml
graders:
  - name: Python Files Generated
    type: output_check
    weight: 0.15
    when:
      language: python
    details:
      min_files: 2
      min_bytes_per_file: 50

  - name: Minimal Go Output
    type: output_check
    weight: 0.1
    when:
      language: go
    details:
      min_files: 1
      min_bytes_per_file: 1
```

## Result Structure

Each `output_check` grader result includes:
- **Pass/Fail**: Binary result based on whether minimums are met
- **Files checked**: List of all files in the workspace with their sizes
- **Qualifying count**: Number of files meeting the size threshold
- **Total bytes**: Sum of bytes in qualifying files

Results are visible in the evaluation report under `grader_results`.

## Coming in v1 (Future Enhancements)

The following features are planned for future releases:

- **Filename presence checking**: Verify specific file names exist (e.g., "README.md must be present")
- **Updated file detection**: Use WorkspaceDelta to verify specific files were modified/created by the agent (not pre-existing)
- **File pattern matching**: Include/exclude files by glob patterns before applying size checks
- **Extension filtering**: Count only files with specific extensions (e.g., "only .py files count")

Until these features ship, use the [`file`](./file.md) grader for specific filename checks and combine multiple `output_check` graders with different thresholds and `when:` conditions for language-specific validation.

## Notes

- **Empty file handling**: Files with 0 bytes are never counted toward `min_files`, regardless of configuration.
- **Workspace state**: The grader inspects the final workspace state after generation is complete.
- **No pattern matching (v1)**: This version checks only file count and total size. Specific filename or content pattern checks use the [`file`](./file.md) grader.
- **Deterministic**: Unlike LLM-based graders, results are fully deterministic given the same workspace state.

See [index.md](./index.md#applicability-when) for `when:` syntax and conditional grader application.

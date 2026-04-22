# Output Check Grader

The `output_check` grader evaluates the files an agent produced or modified during a run against a set of configured checks. It uses `WorkspaceDelta` to track created and edited files (excluding starter files the agent did not touch), then runs independent sub-checks on file counts, specific paths, and file sizes. All sub-checks are reported, and the overall result is the AND of all configured checks.

## When to Use

- **Verify output generation**: Did the agent create or modify files?
- **Ensure specific files exist**: Require particular files in the output (e.g., `README.md`)
- **Validate file modifications**: Confirm the agent edited existing files, not just created new ones
- **Enforce size constraints**: Check that generated files meet minimum and/or maximum size thresholds
- **Forbid sensitive files**: Reject runs that produce secrets or config files accidentally

Use [`file`](./file.md) or [`program`](./program.md) graders for content pattern matching or build validation.

## Configuration

```yaml
graders:
  - name: Output File Checks
    type: output_check
    weight: 0.3
    details:
      min_files: 1
      require_files: [README.md]
      forbid_files: [.env, secrets.json]
      min_bytes_per_file: 10
      max_bytes_per_file: 1048576
```

### `details` Schema

All fields in `details` are optional. Unconfigured knobs are skipped; no implicit defaults apply.

| Field                  | Type      | Semantics |
|------------------------|-----------|-----------|
| `min_files`            | int       | Require ≥ N produced files (created or modified). 0 = unset (check skipped). |
| `max_files`            | int       | Require ≤ N produced files. 0 = unset (check skipped). |
| `require_files`        | []string  | Every listed path must appear in produced files (created or modified). Optional; no paths = check skipped. |
| `forbid_files`         | []string  | None of the listed paths may appear in produced files. Optional; no paths = check skipped. |
| `require_updated`      | []string  | Every listed path must appear in the *modified set* (agent edited existing content). Optional; no paths = check skipped. |
| `min_bytes_per_file`   | int64     | Every produced file must be ≥ N bytes. 0 = unset (check skipped). Vacuously true if zero files produced. |
| `max_bytes_per_file`   | int64     | Every produced file must be ≤ N bytes. 0 = unset (check skipped). |

**"Produced files"** = files the agent created (NewFiles) ∪ files the agent modified (ModifiedFiles). Starter files the agent did not touch are not counted.

## Result Structure

Each `output_check` grader result includes:

- **Pass/Fail**: Boolean; true iff all configured sub-checks pass.
- **Score**: 1.0 if Pass, else 0.0 (boolean grader, no partial credit).
- **OutputCheckDetails** (in JSON result):
  - `ProducedFiles`: Sorted list of all created/modified file paths
  - `SubChecks`: Array of check results, one per configured knob
    - `Check`: Name of the check (e.g., `"min_files"`, `"require_files"`)
    - `Pass`: Boolean result for this check
    - `Message`: Human-readable explanation

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

Passes if the agent created or modified at least one file.

### Require Specific Files

```yaml
graders:
  - name: Generated Implementation
    type: output_check
    weight: 0.3
    details:
      require_files: [main.py, README.md]
      min_bytes_per_file: 50
```

Passes if:
- Both `main.py` and `README.md` are in the produced files
- Every produced file is ≥ 50 bytes

### Enforce Size Constraints

```yaml
graders:
  - name: Reasonable Output Size
    type: output_check
    weight: 0.2
    details:
      min_files: 1
      max_bytes_per_file: 1048576  # 1 MB max per file
      forbid_files: [.env, secrets.json, config.json]
```

Passes if:
- At least one file was produced
- No file exceeds 1 MB
- The forbidden sensitive paths are not present

### Validate File Modifications

```yaml
graders:
  - name: Updated Existing Files
    type: output_check
    weight: 0.15
    details:
      require_updated: [src/main.py, src/config.py]
```

Passes if the agent modified (not just created) both `src/main.py` and `src/config.py`.

### Complete Example (All Knobs)

```yaml
graders:
  - name: Output Check - All Knobs
    type: output_check
    weight: 0.5
    details:
      min_files: 1                    # At least 1 file
      max_files: 50                   # At most 50 files
      require_files: [README.md]      # Must have README.md
      require_updated: [src/main.py]  # Must have modified src/main.py
      min_bytes_per_file: 10          # Every file ≥ 10 bytes
      max_bytes_per_file: 1048576     # Every file ≤ 1 MB
      forbid_files: [.env]            # Must NOT contain .env
```

### Conditional: Language-Specific Checks

```yaml
graders:
  - name: Python Package Structure
    type: output_check
    weight: 0.2
    when:
      language: python
    details:
      min_files: 2
      require_files: [__init__.py, setup.py]

  - name: Go Module Minimal
    type: output_check
    weight: 0.15
    when:
      language: go
    details:
      min_files: 1
      require_files: [go.mod]
```

## Behavior

### Sub-Check Execution

- Each configured knob runs as an independent sub-check.
- **No early exit**: All sub-checks run and are reported, even if earlier ones fail.
- **Deterministic**: Results depend only on the produced files and their sizes; fully reproducible.

### Nil or Empty WorkspaceDelta

If the engine could not compute a `WorkspaceDelta` (e.g., due to workspace access issues), the grader treats it as an empty delta (zero files produced). This ensures meaningful failures: `min_files=1` will correctly fail with a clear message rather than silently skip.

### Construction-Time Validation

`NewOutputCheckGrader` rejects structurally invalid configs:

- `min_files < 0`, `max_files < 0`
- `min_files > max_files` (when both > 0)
- `min_bytes_per_file < 0`, `max_bytes_per_file < 0`
- `min_bytes_per_file > max_bytes_per_file` (when both > 0)

### No Implicit Defaults

- Unconfigured knobs are skipped entirely.
- A config with no knobs specified trivially passes (all zero sub-checks pass).
- There are no defaults like "min_files defaults to 1" — only explicit values matter.

## Notes

- **Empty file handling**: Files with 0 bytes are included in the produced set and subject to all checks (e.g., `min_bytes_per_file: 1` will fail for a 0-byte file).
- **File ordering**: Produced files are reported in sorted order for stable result presentation.
- **No pattern matching (v1)**: This version checks exact paths. Glob or regex matching is deferred to v2.
- **Workspace state**: The grader runs after agent generation is complete, inspecting the final workspace state.

See [index.md](./index.md#applicability-when) for `when:` syntax and conditional grader application.

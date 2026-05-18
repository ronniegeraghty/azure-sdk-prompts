# Output Check Grader (Deprecated)

> **DEPRECATED:** The `output_check` grader type is no longer the canonical approach. Use the [`workspace`](./workspace.md) grader instead.

The `output_check` grader validates file counts and sizes in the agent's workspace output.

## Migration Guide

### Old `output_check` Configuration

```yaml
graders:
  - name: Generated Some Output
    type: output_check
    weight: 0.1
    details:
      min_files: 1
      min_bytes_per_file: 10
      max_bytes_per_file: 50000
```

### New `workspace` Equivalent

Use the [`workspace`](./workspace.md) grader with `kind: file` checks:

```yaml
graders:
  - name: Generated Some Output
    type: workspace
    weight: 0.1
    checks:
      - kind: require_to_create
        files: ["*"]  # Requires at least one file created
      - kind: file
        name: "<your-output-file>"
        state: present
        min_bytes: 10
        max_bytes: 50000
```

## Why the Change?

The `workspace` grader:
- Operates on **actual workspace delta** (file creation, modification, deletion history) instead of just size constraints
- Provides **more granular control** with check kinds for create/update/delete operations
- Unifies file validation with **state and content checks** in one canonical grader type
- Offers **better error messages** and transparent scoring logic

## Legacy Reference

If you encounter `output_check` in existing configs, the field mapping is:

| `output_check` field | `workspace` equivalent |
|---------------------|------------------------|
| `min_files` | Use `require_to_create` with appropriate files |
| `max_files` | Not directly supported; use `workspace` checks to limit deletions |
| `require_files` | Use `require_to_create` with specific file paths |
| `forbid_files` | Use `forbidden_to_create` with specific file paths |
| `require_updated` | Use `required_to_update` with specific file paths |
| `min_bytes_per_file` | Use `file` check kind with `min_bytes` |
| `max_bytes_per_file` | Use `file` check kind with `max_bytes` |

See [workspace.md](./workspace.md) for full documentation and [index.md](./index.md) for grader type overview.

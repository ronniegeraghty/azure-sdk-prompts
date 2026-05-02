# Workspace Grader

The `workspace` grader validates the agent's file modifications during a session. It checks file creation, modification, and deletion patterns against configured expectations. The grader passes only if **all** checks pass (boolean semantics).

This grader replaces the legacy `output_check` and `file` graders with a unified, delta-aware approach.

## Important: WorkspaceDelta Availability

The workspace grader operates on `WorkspaceDelta` — a summary of file changes (created, modified, deleted). WorkspaceDelta is:
- **Available** in evaluations generated with hyoka v0.4+ (it's created by the engine and included in GraderInput)
- **Possibly nil** in evaluation reports created with older hyoka versions (pre-#566). Graders handle nil gracefully by treating it as "no file changes available."

If all checks pass when `WorkspaceDelta` is nil, the grader reports success. If any check **requires** a file to be in NewFiles/ModifiedFiles/DeletedFiles and WorkspaceDelta is nil, the check fails.

## When to Use

- Verify required files were created (e.g., `main.py`, `README.md`)
- Forbid creation of sensitive files (e.g., `.env`, `secrets.txt`)
- Require modification of existing files
- Verify file deletion patterns
- Validate file size and content (for created/modified files)
- Prevent destructive operations (forbid deletion of key files)

## Configuration

The `workspace` grader uses a top-level `checks:` array. Checks validate workspace state using WorkspaceDelta (file creation, modification, and deletion history):

```yaml
graders:
  - name: Workspace Check
    type: workspace
    weight: 0.5
    checks:
      - kind: require_to_create
        files: [main.py, README.md]
      - kind: file
        name: main.py
        state: present
        min_bytes: 100
```

### `checks` Schema

Each check defines one file validation rule. Six check kinds are supported:

#### 1. `require_to_create`

File(s) must appear in the NewFiles list (created by the agent).

| Field | Type     | Required | Description               |
|-------|----------|----------|---------------------------|
| `kind` | string  | yes      | Must be `"require_to_create"` |
| `files` | []string | yes     | List of file paths that must be created |

#### 2. `forbidden_to_create`

File(s) must NOT appear in the NewFiles list.

| Field | Type     | Required | Description               |
|-------|----------|----------|---------------------------|
| `kind` | string  | yes      | Must be `"forbidden_to_create"` |
| `files` | []string | yes     | List of file paths that must not be created |

#### 3. `required_to_update`

File(s) must appear in the ModifiedFiles list (modified by the agent).

| Field | Type     | Required | Description               |
|-------|----------|----------|---------------------------|
| `kind` | string  | yes      | Must be `"required_to_update"` |
| `files` | []string | yes     | List of file paths that must be modified |

#### 4. `required_to_delete`

File(s) must appear in the DeletedFiles list (deleted by the agent).

| Field | Type     | Required | Description               |
|-------|----------|----------|---------------------------|
| `kind` | string  | yes      | Must be `"required_to_delete"` |
| `files` | []string | yes     | List of file paths that must be deleted |

#### 5. `forbidden_to_delete`

File(s) must NOT appear in the DeletedFiles list. Use `files: ["*"]` to forbid ALL deletions.

| Field | Type     | Required | Description               |
|-------|----------|----------|---------------------------|
| `kind` | string  | yes      | Must be `"forbidden_to_delete"` |
| `files` | []string | yes     | List of file paths that must not be deleted, or `["*"]` to forbid all deletions |

#### 6. `file`

Validate file state on disk: existence, size, and content matching. Checks both file presence and optional content constraints.

| Field | Type     | Required | Description               |
|-------|----------|----------|---------------------------|
| `kind` | string  | yes      | Must be `"file"` |
| `name` | string  | yes      | File path relative to workspace root |
| `state` | string  | yes      | Either `"present"` (file must exist) or `"absent"` (file must not exist) |
| `min_bytes` | int64 | no (present only) | Minimum file size in bytes |
| `max_bytes` | int64 | no (present only) | Maximum file size in bytes |
| `contains` | string | no (present only) | Substring that must appear in file content |
| `excludes` | string | no (present only) | Substring that must NOT appear in file content |

> Note: `min_bytes`, `max_bytes`, `contains`, and `excludes` are only valid when `state: present`. Using them with `state: absent` is an error.

## Examples

### Basic File Creation

```yaml
graders:
  - name: Main Files Created
    type: workspace
    weight: 0.3
    checks:
      - kind: require_to_create
        files: [main.py, tests/test_main.py, README.md]
```

### Forbid Sensitive Files

```yaml
graders:
  - name: No Secrets Committed
    type: workspace
    weight: 0.5
    checks:
      - kind: forbidden_to_create
        files: [.env, secrets.txt, api_key.txt, config.secret]
```

### Verify File Content

```yaml
graders:
  - name: Generated Code Quality
    type: workspace
    weight: 0.4
    checks:
      - kind: file
        name: src/main.py
        state: present
        min_bytes: 100
        max_bytes: 50000
        contains: "def main"
        excludes: "TODO"
```

### Prevent Destructive Changes

```yaml
graders:
  - name: No File Deletion
    type: workspace
    weight: 0.2
    checks:
      - kind: forbidden_to_delete
        files: ["*"]  # Forbid all deletions
```

### Update Requirement

```yaml
graders:
  - name: Must Update Config
    type: workspace
    weight: 0.2
    checks:
      - kind: required_to_update
        files: [config.yaml, setup.py]
```

### Comprehensive Workspace Validation

```yaml
graders:
  - name: Comprehensive Workspace Check
    type: workspace
    weight: 0.6
    checks:
      # Created files
      - kind: require_to_create
        files: [main.py]

      # Forbidden files
      - kind: forbidden_to_create
        files: [.env, debug.log, secrets.txt]

      # No deletions allowed
      - kind: forbidden_to_delete
        files: ["*"]

      # File state and content
      - kind: file
        name: main.py
        state: present
        min_bytes: 50
        contains: "def main"
```

## Result Structure

Each workspace grader produces:
- **Pass/Fail**: Binary result (true only if ALL checks pass)
- **Check results**: Individual pass/fail for each configured check with explanation
- **Score**: 1.0 if all checks pass, 0.0 if any check fails (boolean grader)

Results visible in evaluation reports under `grader_results`.

## Data Visible to Grader

Workspace graders can access:
- **WorkspaceDelta**: File creation, modification, and deletion history
  - `NewFiles`: List of files created during generation
  - `ModifiedFiles`: List of files modified during generation
  - `DeletedFiles`: List of files deleted during generation
- **Workspace files**: Current state of files on disk (for `kind: file` checks)
- **File sizes**: Byte counts for size validation
- **File content**: Full content for substring matching (contains/excludes)

## Notes

- **Delta-driven**: Checks for create/update/delete use WorkspaceDelta, not disk state. This captures the agent's actions independent of the current file state.
- **Absolute paths in `files:`**: Paths in `require_to_create`, `forbidden_to_create`, etc. are matched against file paths in the delta (typically relative to workspace root).
- **Wildcard support**: Use `files: ["*"]` in `forbidden_to_delete` to forbid **any** file deletion.
- **Substring matching**: `contains` and `excludes` are case-sensitive substring matches, not regex patterns.
- **All-or-nothing**: Grader passes only if **every** check passes. A single failed check fails the grader.
- **State combinations**: A file can be created AND then modified; it will appear in both `NewFiles` and `ModifiedFiles`.

## Troubleshooting

- **Check always fails**: Verify file paths match exactly (case-sensitive) and use paths relative to workspace root.
- **Size constraints not working**: Ensure `min_bytes` and `max_bytes` use integers (not strings) and are in the correct order (min ≤ max).
- **Content checks failing**: Remember that `contains` and `excludes` are substring matches, not regex; use escaped strings for special characters.
- **Wildcard not working**: Use `files: ["*"]` only with `forbidden_to_delete`; other check kinds require specific file paths.

See [index.md](./index.md) for general grader concepts and [../configuration.md](../configuration.md) for config file structure.

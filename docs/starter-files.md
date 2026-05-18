# Starter File Reference Format

**Status:** DRAFT
**Date:** 2025-07-25
**Owner:** Morpheus (Architect)
**Issue:** [#117](https://github.com/ronniegeraghty/hyoka/issues/117)

## 1. Overview

Starter files enable "fix this broken code" and "extend this existing project" style
prompts. Before a Copilot session begins, hyoka copies referenced files into the
agent's workspace so the agent works against a pre-existing codebase rather than
starting from scratch.

This design follows **Waza's `ResourceFile` pattern** — prompt authors declare
resources alongside the prompt, and the evaluation harness materializes them in the
workspace before generation begins.

## 2. Frontmatter Format

### 2.1 Option A: Directory Reference (Recommended)

A single `starter_project` field points to a directory whose entire contents are
copied into the workspace root.

```yaml
---
id: storage-dp-dotnet-fix-retry-logic
properties:
  service: storage
  plane: data-plane
  language: dotnet
  category: debugging
  difficulty: intermediate
  description: Fix broken retry logic in an existing Azure Storage client
starter_project: ./starters/
---
```

The directory tree is copied as-is:

```
prompts/storage/data-plane/dotnet/
├── fix-retry-logic.prompt.md
└── starters/
    ├── Program.cs
    ├── StorageApp.csproj
    └── appsettings.json
```

After copy, the agent workspace contains:

```
<workspace>/
├── Program.cs
├── StorageApp.csproj
└── appsettings.json
```

> **Why recommend directory reference?** Most starter projects involve multiple
> interdependent files (source, project file, config). A single directory keeps them
> co-located, avoids long file lists in frontmatter, and mirrors how developers think
> about projects.

### 2.2 Precedence Rules

| Condition | Behavior |
|-----------|----------|
| `starter_project` set | Copy entire directory contents to workspace root |
| Neither set | Empty workspace (current default behavior) |

## 3. Path Resolution

All paths are resolved **relative to the prompt file's directory**, matching the
existing implementation in `eval/copilot.go`:

```go
starterDir := p.StarterProject
if !filepath.IsAbs(starterDir) && p.FilePath != "" {
    starterDir = filepath.Join(filepath.Dir(p.FilePath), starterDir)
}
```

Absolute paths are supported but discouraged — they break portability across machines.

> **Security:** `copyDir` already skips symlinks and logs a warning. Validation
> should additionally reject paths that escape the `prompts/` tree via `../`
> traversal.

## 4. File Placement

Files are copied into the workspace **before** the Copilot session is created. The
sequence is:

```
1. Create temp workspace directory
2. Copy starter files/project into workspace root
3. List copied files (for metadata)
4. Create Copilot client with workspace as cwd
5. Start session — agent sees files immediately
```

This matches the current implementation. The `starterFiles` list returned from step 3
can be included in `EvalResult` metadata so reviewers know which files were
pre-existing vs. agent-generated.

## 5. Recommended Directory Layout

### Convention

Place starter files in a `starters/` subdirectory next to the prompt file. This keeps
the prompt directory tidy and makes the relationship explicit.

```
prompts/
└── key-vault/
    └── data-plane/
        └── python/
            ├── crud-secrets.prompt.md           # blank project prompt
            ├── fix-error-handling.prompt.md      # starter project prompt
            └── fix-error-handling.starters/
                ├── main.py
                ├── requirements.txt
                └── README.md
```

Naming the directory `<prompt-short-name>.starters/` prevents collisions when
multiple prompts in the same directory each have their own starter files.

### Interaction with `project_context`

| `project_context` | `starter_project` | Meaning |
|--------------------|--------------------|---------|
| `blank` (default)  | not set            | Agent starts from scratch |
| `existing`         | set                | Agent works on pre-existing code |
| `existing`         | not set            | Validation warning — existing context declared but no files provided |
| `blank`            | set                | Validation warning — starter files provided but context says blank |

## 6. Validation

The `hyoka validate` command should check starter file references. Add these checks
to `internal/validate/validate.go`:

### 6.1 Checks for `starter_project`

| Check | Severity | Message |
|-------|----------|---------|
| Path exists and is a directory | Error | `starter_project path does not exist or is not a directory: %s` |
| Path does not escape `prompts/` tree | Error | `starter_project must not reference paths outside prompts/ directory` |
| Directory is not empty | Warning | `starter_project directory is empty: %s` |

### 6.2 Cross-field consistency

| Check | Severity | Message |
|-------|----------|---------|
| `project_context: existing` without starter files | Warning | `project_context is 'existing' but no starter_project provided` |
| `project_context: blank` with starter files | Warning | `project_context is 'blank' but starter_project is configured` |

## 7. Complete Example Prompt

```markdown
---
id: key-vault-dp-python-fix-error-handling
tags: [error-handling, secrets, existing-project]
properties:
  service: key-vault
  plane: data-plane
  language: python
  category: debugging
  difficulty: intermediate
  description: Fix missing error handling in an Azure Key Vault secrets client
  sdk_package: azure-keyvault-secrets
  doc_url: https://learn.microsoft.com/python/api/azure-keyvault-secrets/
  created: "2025-07-25"
  author: morpheus
project_context:
  type: existing
starter_project: ./fix-error-handling.starters/
timeout: 300
expected_packages:
  - azure-keyvault-secrets
  - azure-identity
---

## Prompt

The project in your workspace contains an Azure Key Vault client that retrieves
secrets but has no error handling. The `main.py` file will crash on network errors
or missing secrets.

Add proper error handling:
1. Wrap Key Vault operations in try/except blocks
2. Handle `ResourceNotFoundError` for missing secrets
3. Handle `HttpResponseError` for service errors
4. Add retry logic using the SDK's built-in retry policy
5. Ensure the client is properly closed in a finally block

## Evaluation Criteria

- Error handling covers all Key Vault operations
- Uses SDK-specific exception types, not bare `except`
- Retry policy is configured on the client
- Client cleanup is guaranteed (context manager or finally)
- Existing functionality is preserved
```



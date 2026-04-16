---
id: key-vault-dp-python-fix-error-handling

# ──────────────────────────────────────────────────────────────────────
# starter_project  (string, optional)
#   Path to a directory whose contents are copied into the agent's
#   workspace BEFORE the Copilot session begins.  The path is resolved
#   relative to THIS prompt file.
#
#   At runtime hyoka will:
#     1. Create an empty temp workspace directory
#     2. Copy every file from the starter_project directory into that
#        workspace root (preserving subdirectory structure)
#     3. Start the Copilot session — the agent sees those files
#        immediately and can modify them in place
#
#   Convention: name the directory <prompt-short-name>.starters/ to
#   avoid collisions when multiple prompts live in the same folder.
# ──────────────────────────────────────────────────────────────────────
starter_project: ./existing-files-example.starters/

# ──────────────────────────────────────────────────────────────────────
# project_context  (map, optional)
#   Tells hyoka whether the agent is working on pre-existing code or
#   starting from a blank workspace.
#
#   type: existing  — agent works on files from starter_project
#   type: blank     — agent starts from scratch (this is the default
#                     when project_context is omitted entirely)
#
#   If you set type: existing, you should also set starter_project so
#   hyoka has files to copy.  hyoka will warn if these two fields are
#   inconsistent (e.g., type: existing but no starter_project).
# ──────────────────────────────────────────────────────────────────────
project_context:
  type: existing

properties:
  service: key-vault
  plane: data-plane
  language: python
  category: error-handling
  difficulty: intermediate
  description: >
    Fix missing error handling in an existing Azure Key Vault secrets
    client.  The starter project contains a working but fragile script
    that crashes on network errors or missing secrets.
  sdk_package: azure-keyvault-secrets
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/keyvault-secrets-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
  - error-handling
  - secrets
  - existing-project
  - starter-files
---

# Fix Error Handling: Azure Key Vault Secrets (Python)

## Prompt

The project in your workspace contains an Azure Key Vault client that
retrieves and stores secrets but has **no error handling**. The `main.py`
file will crash on network errors, missing secrets, or authentication
failures.

Your job is to **fix the existing code** — do not rewrite it from
scratch. Add proper error handling while preserving the current
functionality:

1. Wrap Key Vault operations in try/except blocks using SDK-specific
   exception types (`ResourceNotFoundError`, `HttpResponseError`,
   `ClientAuthenticationError`) — never use bare `except`.
2. Handle missing secrets gracefully by logging a warning and
   returning `None` instead of crashing.
3. Add a retry policy to the `SecretClient` using the SDK's built-in
   `RetryPolicy` or the `retry_total` keyword argument.
4. Ensure the `SecretClient` is properly closed using a context
   manager (`with` statement) or a `finally` block.
5. Keep the existing `get_secret` and `set_secret` helper functions —
   fix them in place.

## Evaluation Criteria

The generated code should demonstrate:

- Uses `azure.core.exceptions` types, not bare `except Exception`
- `ResourceNotFoundError` handled for missing secrets
- `HttpResponseError` handled for service-level errors
- `ClientAuthenticationError` handled for credential failures
- Retry policy configured on the client
- Client cleanup is guaranteed (context manager or finally)
- Existing function signatures and behavior are preserved
- `requirements.txt` is unchanged (dependencies are already correct)

## Context

This example demonstrates a "fix broken code" prompt — the most common
use case for `starter_project`.  Instead of generating code from
scratch, the agent must understand existing code and surgically add
error handling.  This tests whether the agent respects existing
structure rather than rewriting everything.

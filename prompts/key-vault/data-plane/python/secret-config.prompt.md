---
id: key-vault-dp-python-secret-config
properties:
  service: key-vault
  plane: data-plane
  language: python
  category: crud
  difficulty: intermediate
  description: 'Can an agent generate a Key Vault-backed configuration provider with secret versioning, expiry inspection,
    in-memory caching with bulk-load, and safe secret rotation using long-running delete operations?

    '
  sdk_package: azure-keyvault-secrets
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/keyvault-secrets-readme
  created: '2026-04-10'
  author: copilot
tags:
- key-vault
- secrets
- caching
- secret-rotation
- lro
- versioning
- expiry
- async
---

# Secret Config Provider: Azure Key Vault (Python)

## Prompt

Create a Python project that implements an application configuration provider backed by Azure Key Vault.

**Write the code to files (use file-write tools, do not reply with code blocks).**

The project needs:

- A **secret provider module** (both sync and async versions) that retrieves secrets from Key Vault by name, with graceful handling when a secret doesn't exist (return a default value instead of crashing). It should also be able to retrieve a specific version of a secret (not just the latest), and inspect a secret's expiry date so the caller can tell if a secret is about to expire.

- A **caching layer** on top of the provider that stores secret values in memory after first retrieval. It should support bulk-loading a predefined set of required config keys at startup, on-demand refresh of individual keys, and automatic re-fetch of any secret whose expiry date is within a configurable warning window (e.g., 7 days out).

- A **configuration/factory module** that connects securely to the Key Vault using the vault URL from an environment variable. The application runs in Azure and should authenticate using `DefaultAzureCredential` — no client secrets or certificates in code.

- A **secret rotation helper** that safely rotates a secret: delete the old secret, ensure the deletion is fully complete, then create the new secret with an updated value and expiry date. The rotation must be safe — don't assume deletion is instantaneous, since Key Vault's soft-delete feature means the secret may not be immediately gone. Use the long-running operation poller returned by `begin_delete_secret()`.

- A **main script** that demos both implementations: loading several config keys at startup, reading them from cache, refreshing one, printing a warning if any secret is near expiry, and performing a secret rotation (delete old, wait for completion, create new). Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Scenario-Specific Patterns
- Secret versioning: retrieves specific version via `get_secret(name, version=version)`
- Secret expiry: accesses `secret.properties.expires_on`
- Configurable warning window for near-expiry detection
- In-memory caching (e.g., `dict`) with bulk-load and single-key refresh
- Secret rotation uses `begin_delete_secret()` as a long-running operation
- Sync uses `LROPoller` — calls `.wait()` or `.result()` to wait for delete completion
- Async uses the async poller — `await poller.wait()` for delete completion
- Creates new secret only after delete completes (not concurrently)
- Async version uses `azure.keyvault.secrets.aio.SecretClient`

### Scenario-Specific Error Handling
- Returns a default value when secret is not found (`ResourceNotFoundError`)

### Anti-Patterns (scenario-specific)
- NOT using fire-and-forget `delete_secret()` without the long-running operation
- NOT ignoring the poller and creating the new secret immediately

## Context

This goes beyond basic secret CRUD (covered by `crud-secrets.prompt.md`) to test production
patterns: secret versioning, expiry-aware caching, and safe secret rotation using long-running
operations. The rotation pattern is critical — Key Vault uses soft-delete, so `begin_delete_secret()`
returns an `LROPoller` that must be waited on before the new secret can be created. LLMs
frequently generate a simple `delete_secret()` call without waiting, which fails in production.
The caching layer tests whether the agent can build a practical config provider on top of the
raw Key Vault client.

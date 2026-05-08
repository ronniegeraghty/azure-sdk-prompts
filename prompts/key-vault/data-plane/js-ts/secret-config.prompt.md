---
id: key-vault-dp-js-ts-secret-config
properties:
  service: key-vault
  plane: data-plane
  language: js-ts
  category: crud
  difficulty: intermediate
  description: >
    Can an agent generate a Key Vault-backed configuration provider in
    TypeScript with secret versioning, expiry inspection, in-memory caching
    with bulk-load, and safe secret rotation using version-based updates and
    long-running delete operations?
  sdk_package: '@azure/keyvault-secrets'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/keyvault-secrets-readme
  created: '2026-04-30'
  author: kaghiya
tags:
  - key-vault
  - secrets
  - caching
  - secret-rotation
  - lro
  - poller
  - versioning
  - expiry
---

# Secret Config Provider: Azure Key Vault (TypeScript)

## Prompt

Create a TypeScript Node.js project that implements an application configuration provider backed by Azure Key Vault.

The project needs:

- A **secret provider class** that retrieves secrets from Key Vault by name, with graceful handling when a secret doesn't exist (return a default value instead of crashing) — use `RestError` from `@azure/core-rest-pipeline` with `statusCode` checks (e.g., 404) to detect not-found vs other failures. It should also be able to retrieve a specific version of a secret (not just the latest), and inspect a secret's expiry date so the caller can tell if a secret is about to expire.

- A **caching layer** on top of the provider that stores secret values in memory after first retrieval. It should support bulk-loading a predefined set of required config keys at startup, on-demand refresh of individual keys, and automatic re-fetch of any secret whose expiry date is within a configurable warning window (e.g., 7 days out).

- A **configuration module** that connects securely to the Key Vault using the vault URL from an environment variable. The application runs in Azure and should authenticate using managed identity — no client secrets or certificates in code.

- A **secret rotation helper** that safely rotates a secret: create a new version of the secret with an updated value and expiry date (since Key Vault supports multiple versions per secret name), then optionally clean up old versions by deleting and purging the previous secret if full name reuse is needed. The cleanup must be safe — use the long-running delete operation and wait for completion before purging, since Key Vault's soft-delete feature means the secret is not immediately gone.

- A **main script** that demos the full flow: loading several config keys at startup, reading them from cache, refreshing one, printing a warning if any secret is near expiry, and performing a secret rotation (creating a new version, then demonstrating the delete-and-purge cleanup flow). Print results at each step.

Enable SDK diagnostic logging using `@azure/logger` with a configurable log level for debugging.

Include a complete `package.json` with the necessary Azure SDK dependencies and a `tsconfig.json`.

## Evaluation Criteria

### Scenario-Specific Patterns
- Secret versioning: retrieves specific version via `client.getSecret(name, { version })`
- Secret expiry: accesses `properties.expiresOn` on the secret response
- Configurable warning window for near-expiry detection (compares expiresOn to current date)
- In-memory caching (e.g., `Map<string, KeyVaultSecret>`) with bulk-load and single-key refresh
- Version-based rotation: calls `client.setSecret(name, newValue, { expiresOn })` to create a new version
- Cleanup uses `client.beginDeleteSecret(name)` as a long-running operation
- Awaits `poller.pollUntilDone()` before calling `client.purgeDeletedSecret(name)`
- Creates new secret only after delete+purge completes (not concurrently)

### Scenario-Specific Error Handling
- Returns a default value when secret is not found (404), does not crash
- Handles RestError with statusCode check for not-found scenarios

### Anti-Patterns (scenario-specific)
- NOT using fire-and-forget `deleteSecret()` without waiting for completion
- NOT assuming deletion is instantaneous (must use poller)
- NOT hardcoding vault URL or credentials

## Context

This goes beyond basic secret CRUD (covered by `crud-secrets.prompt.md`) to test production
patterns: secret versioning, expiry-aware caching, and safe secret rotation. The rotation
pattern is important — Key Vault supports multiple versions per secret name, so the simplest
rotation is just `setSecret()` with the same name to create a new version. For full cleanup,
`beginDeleteSecret()` returns a Poller that must be awaited via `pollUntilDone()` before
purging. LLMs frequently generate a simple `deleteSecret()` call without the poller, which
silently fails in soft-delete-enabled vaults.

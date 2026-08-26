# Azure Key Vault configuration provider

A Java 17 example with synchronous and asynchronous Key Vault secret providers,
in-memory caches, expiry-aware refresh, and safe secret rotation.

## Configuration

Set `KEY_VAULT_URL` to the vault URL. For a user-assigned managed identity, also
set `AZURE_CLIENT_ID`; otherwise the Azure host's system-assigned identity is
used. No client secret or certificate is required.

The managed identity needs secret read/write/delete/purge permissions for the
full demo. Rotation purges the soft-deleted secret and waits until its name can
be reused. It therefore cannot run against a vault with purge protection
enabled; in that case Key Vault intentionally prevents immediate recreation
under the same name.

## Build and run

```powershell
mvn package
mvn exec:java
```

`Main` loads the required keys, reads them from cache without printing secret
values, refreshes one key, reports secrets in the expiry warning window, and
rotates `demo-rotating-secret` with the sync API followed by the async API.

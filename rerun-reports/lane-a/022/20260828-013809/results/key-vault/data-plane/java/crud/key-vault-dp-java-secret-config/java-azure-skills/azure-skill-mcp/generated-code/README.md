# Azure Key Vault configuration provider

A small Java 17 example with synchronous and Reactor-based asynchronous providers, in-memory caches, expiry-aware refresh, and safe secret rotation.

## Configuration

Set `KEY_VAULT_URL` to an HTTPS vault URL. The application authenticates only with Azure managed identity. For a user-assigned identity, also set `AZURE_CLIENT_ID`; otherwise the system-assigned identity is used.

The identity needs secret read/list permissions for the provider. The rotation demo additionally needs delete and purge permissions.

```powershell
$env:KEY_VAULT_URL = "https://your-vault.vault.azure.net"
mvn compile exec:java
```

`DEMO_SECRET_VERSION` is optional and demonstrates fetching an exact secret version. Secret values are deliberately not printed.

## Rotation behavior

Key Vault soft-delete retains a deleted secret name, so the helper waits for deletion, purges the deleted record, waits until that record is no longer visible, and only then creates the replacement. Rotation therefore requires purge permission and a vault without purge protection. When purge protection is enabled, delete-and-recreate rotation is intentionally rejected by Key Vault; use normal secret versioning instead.

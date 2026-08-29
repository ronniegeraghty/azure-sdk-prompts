# Azure Key Vault configuration provider

Java 17 sample with synchronous and asynchronous Key Vault secret providers,
in-memory caching, expiry-aware refresh, and soft-delete-safe rotation.

## Configuration

Set these environment variables on the Azure-hosted application:

- `KEY_VAULT_URL`: required, for example `https://my-vault.vault.azure.net`
- `AZURE_CLIENT_ID`: optional client ID for a user-assigned managed identity;
  omit it to use the system-assigned managed identity
- `ROTATION_SECRET_NAME`: optional demo secret name; defaults to
  `rotating-demo-secret`

Grant the managed identity only the Key Vault data-plane permissions it needs.
The rotation demo needs read, set, delete, and purge permissions. It cannot
recreate the same secret name when purge protection is enabled; in that case,
rotate by creating a new version with `setSecret` instead of delete-and-purge.

## Build and run

```powershell
mvn clean package
mvn exec:java
```

The demo intentionally performs real secret deletion, purge, and creation
against the configured vault. Use a non-production secret and vault.

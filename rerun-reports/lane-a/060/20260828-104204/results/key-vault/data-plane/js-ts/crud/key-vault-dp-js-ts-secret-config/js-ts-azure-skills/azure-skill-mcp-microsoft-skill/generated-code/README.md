# Azure Key Vault configuration provider

TypeScript configuration loading backed by Azure Key Vault Secrets. It uses
`ManagedIdentityCredential`, caches values in memory, refreshes secrets near expiry,
supports version-specific reads, and provides guarded rotation and cleanup helpers.

## Run

```powershell
npm install
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
npm start
```

For a user-assigned managed identity, also set `AZURE_CLIENT_ID`. The identity needs
Key Vault secret `get` permissions for reads, plus `set`, `delete`, and `purge` only
when those demo operations are enabled.

Optional demo settings:

| Variable | Purpose |
|---|---|
| `SECRET_EXPIRY_WARNING_DAYS` | Near-expiry refresh window; defaults to `7` |
| `REFRESH_SECRET_NAME` | Secret refreshed on demand |
| `DEMO_SECRET_VERSION_NAME`, `DEMO_SECRET_VERSION` | Read one exact version |
| `ROTATION_SECRET_NAME`, `ROTATION_SECRET_VALUE` | Create a new secret version |
| `ROTATION_EXPIRY_DAYS` | New version lifetime; defaults to `90` |
| `PURGE_ROTATED_SECRET=true` | Enable permanent name-level cleanup |
| `PURGE_CONFIRM_SECRET_NAME` | Must exactly match the name being purged |

Key Vault cannot delete only one secret version. Normal rotation therefore creates a
new version and keeps history. `deleteAndPurgeSecret` is a separate, explicitly
confirmed operation that deletes **every version** under the name, waits for the
soft-delete operation to complete, and then purges it so the name can be reused.
Vaults with purge protection intentionally reject immediate purge.

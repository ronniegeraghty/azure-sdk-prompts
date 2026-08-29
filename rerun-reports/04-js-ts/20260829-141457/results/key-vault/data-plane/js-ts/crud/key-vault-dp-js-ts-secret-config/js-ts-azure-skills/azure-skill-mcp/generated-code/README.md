# Azure Key Vault configuration provider

TypeScript/Node.js configuration provider backed by Azure Key Vault. It uses managed identity only, caches values in memory, refreshes near-expiry secrets, supports version-specific reads, and provides explicit rotation and delete/purge operations.

## Run

```powershell
npm install
npm test
npm run build
$env:KEY_VAULT_URL = "https://your-vault.vault.azure.net"
npm start
```

The Azure host's managed identity needs data-plane permissions to read secrets. Rotation additionally needs `set`; cleanup needs `delete` and `purge`. For a user-assigned managed identity, set `AZURE_CLIENT_ID` to its client ID.

Copy `.env.example` values into the host environment. The program does not load `.env` files or accept client secrets. Secret values are redacted in logs.

## Safety

`setSecret` creates a new version when a secret name already exists. Azure Key Vault cannot delete just one secret version: deleting a secret name soft-deletes every version under that name. For that reason, delete/purge is separate from rotation, waits for `beginDeleteSecret(...).pollUntilDone()`, and requires both `ENABLE_PURGE_DEMO=true` and a `PURGE_SECRET_NAME`. Purge is irreversible and can be blocked by vault purge protection.

## References

- [Azure Key Vault JavaScript quickstart](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-node)
- [Delete, recover, and purge a secret](https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-delete-secret)
- [Managed identities for Azure resources](https://learn.microsoft.com/entra/identity/managed-identities-azure-resources/overview)

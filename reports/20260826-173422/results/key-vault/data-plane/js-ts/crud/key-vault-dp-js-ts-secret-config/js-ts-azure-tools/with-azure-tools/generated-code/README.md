# Azure Key Vault configuration provider

A TypeScript Node.js configuration provider with:

- managed-identity authentication to Azure Key Vault;
- latest and version-specific secret reads with defaults for missing secrets;
- expiry metadata and warning-window checks;
- in-memory startup loading, individual refresh, and near-expiry refresh;
- new-version rotation and guarded delete/purge cleanup.

## Run the offline demo

```powershell
npm install
npm run demo
npm test
```

The demo uses `InMemorySecretClient`; it does not contact Azure or print secret
values. Application code can construct the production SDK client with:

```typescript
import { createKeyVaultSecretClient } from "./configuration.js";

const client = createKeyVaultSecretClient();
```

Set `KEY_VAULT_URL` to the HTTPS vault origin. For a user-assigned managed
identity, optionally set `AZURE_MANAGED_IDENTITY_CLIENT_ID`; otherwise the
system-assigned identity is used. Grant only the Key Vault data-plane
permissions needed by the application.

## Cleanup semantics

`setSecret` creates a new version under the same name. Azure Key Vault cannot
delete one secret version: `beginDeleteSecret` deletes the name and every
version. `deleteAndPurgeForNameReuse` therefore requires an exact-name
confirmation, waits for the delete long-running operation, and only then
purges. Purge will fail when purge protection or RBAC policy disallows it.

References:

- https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-get-started
- https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-get-set-secrets
- https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-delete-secret
- https://learn.microsoft.com/javascript/api/@azure/keyvault-secrets/secretclient
- https://learn.microsoft.com/azure/developer/javascript/sdk/authentication/azure-hosted-apps

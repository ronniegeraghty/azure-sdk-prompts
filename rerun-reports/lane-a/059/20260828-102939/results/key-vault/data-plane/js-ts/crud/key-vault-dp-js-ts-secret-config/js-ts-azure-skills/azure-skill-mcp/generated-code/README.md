# Azure Key Vault configuration provider

A TypeScript Node.js example that reads versioned secrets from Azure Key Vault, caches configuration in memory, refreshes near-expiry values, and rotates secrets by creating new versions.

## Run locally without Azure

The default demo uses an in-memory Key Vault-compatible client, so it is safe to run offline:

```powershell
npm install
npm run demo
```

## Run against Azure Key Vault

Set `KEY_VAULT_DEMO_MODE=azure` and `KEY_VAULT_URL` to the vault URL. The application uses `ManagedIdentityCredential`; it contains no client secret or certificate. `AZURE_CLIENT_ID` is optional and selects a user-assigned managed identity. Otherwise, the system-assigned managed identity is used.

The identity needs **Key Vault Secrets User** to read configuration and **Key Vault Secrets Officer** (or a narrower custom role with equivalent secret permissions) for rotation and deletion.

```powershell
$env:KEY_VAULT_DEMO_MODE = "azure"
$env:KEY_VAULT_URL = "https://your-vault-name.vault.azure.net/"
npm run demo
```

The Azure demo does not perform permanent cleanup unless `RUN_DESTRUCTIVE_CLEANUP=true`. Purging requires additional permission and fails when purge protection is enabled.

## Important deletion behavior

Calling `setSecret` for an existing name creates a new version; this is the normal rotation flow. Azure Key Vault cannot delete one historical version. `deleteAndPurgeSecret` therefore deletes and permanently purges the secret name **and every version**, waits for the long-running delete operation with `pollUntilDone()`, and requires the secret name as an explicit confirmation value.

Secret values are masked in demo output to avoid leaking credentials into logs.

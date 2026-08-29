# Azure Key Vault configuration provider

This TypeScript project provides application configuration backed by Azure Key
Vault Secrets, including defaults for missing secrets, version-specific reads,
expiry inspection, in-memory caching, startup bulk loading, refresh, rotation,
and explicit delete-and-purge cleanup.

## Run locally

The default demo uses an in-memory Key Vault substitute and does not contact
Azure:

```bash
npm install
npm test
npm run demo
```

Secret values are not printed; the demo reports their source, version, and
length.

## Run in Azure

Enable a system-assigned or user-assigned managed identity on the host and grant
only the required Key Vault data-plane permissions. Then configure:

```text
DEMO_MODE=azure
KEY_VAULT_URL=https://your-vault-name.vault.azure.net
AZURE_CLIENT_ID=<user-assigned-managed-identity-client-id> # optional
```

`ManagedIdentityCredential` is used directly. No client secret or certificate is
accepted by the configuration module.

The demo's permanent cleanup is disabled in Azure unless both
`RUN_DESTRUCTIVE_CLEANUP=true` and
`PURGE_CONFIRM_SECRET_NAME=demo-rotating-secret` are set. Purging is irreversible
and fails when purge protection is enabled.

## Rotation semantics

`setSecret` creates a new version under the same secret name, so normal rotation
does not require deletion. Azure Key Vault cannot delete only one secret
version: `beginDeleteSecret(name)` soft-deletes the name and all its versions.
The cleanup helper therefore requires an exact-name confirmation, waits for the
long-running delete operation with `pollUntilDone()`, and only then requests a
purge.

## References

- [Azure Key Vault Secrets client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/keyvault-secrets-readme)
- [Authenticate Azure-hosted JavaScript apps with managed identity](https://learn.microsoft.com/azure/developer/javascript/sdk/authentication/system-assigned-managed-identity)

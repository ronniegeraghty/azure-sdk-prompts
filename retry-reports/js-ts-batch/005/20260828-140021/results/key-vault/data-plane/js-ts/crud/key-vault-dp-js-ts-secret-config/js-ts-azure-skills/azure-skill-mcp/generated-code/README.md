# Azure Key Vault configuration provider

This TypeScript project loads application configuration from Azure Key Vault,
caches values in memory, refreshes secrets that are near expiry, retrieves
specific secret versions, and demonstrates safe secret rotation.

## Prerequisites

- Node.js 20 or later
- An Azure-hosted workload with a system-assigned or user-assigned managed
  identity
- Key Vault data-plane permissions for the operations used by the application

No client secret or certificate is used. `ManagedIdentityCredential` uses the
Azure workload's managed identity. Set `AZURE_CLIENT_ID` only when selecting a
user-assigned managed identity.

## Configure and run

```text
KEY_VAULT_URL=https://your-vault-name.vault.azure.net
SECRET_EXPIRY_WARNING_DAYS=7
ROTATION_SECRET_NAME=rotation-demo
ROTATION_SECRET_VALUE=set-this-securely-in-the-host-environment
ROTATION_EXPIRY_DAYS=90
ENABLE_DELETE_AND_PURGE_DEMO=false
```

Install and build:

```text
npm install
npm run check
npm run demo
```

The rotation helper calls `setSecret`, which creates a new version under the
same secret name. Key Vault does not support deleting one old secret version
through `SecretClient`: deleting by name soft-deletes every version. For that
reason, delete-and-purge is a separate, explicitly enabled demo. It waits for
the long-running delete operation to finish before purging. Purge will fail
when purge protection is enabled or when the identity lacks purge permission.

Do not log secret values. The demo redacts cached values and only prints secret
names, versions, and expiry metadata.

# Azure Key Vault secret CRUD (TypeScript)

## Required packages

```powershell
npm install @azure/identity @azure/keyvault-secrets
npm install --save-dev typescript tsx @types/node
```

## Run

Set the vault URL and authenticate with any identity supported by
`DefaultAzureCredential`, such as Azure CLI credentials for local development:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
npm install
npm start
```

The identity needs secret `get`, `set`, `delete`, and `purge` permissions. With
Azure RBAC, the **Key Vault Secrets Officer** role includes these operations.
Purging is irreversible and fails when purge protection is enabled; in that
case, the soft-deleted secret remains until the retention period expires.

# Azure Key Vault secret CRUD (TypeScript)

This sample creates `my-secret`, reads it, updates it to `updated-value`, then
soft-deletes and permanently purges it.

## Install

```powershell
npm install
```

The Azure SDK packages required at runtime are:

```powershell
npm install @azure/identity @azure/keyvault-secrets
```

## Run

Set the vault URL, authenticate with a credential supported by
`DefaultAzureCredential`, and run the script:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
az login
npm start
```

In Azure-hosted environments, `DefaultAzureCredential` can use managed identity
instead of Azure CLI credentials.

The identity needs secret `get`, `set`, `delete`, and `purge` data-plane
permissions. Purging is irreversible and fails when purge protection is enabled.

Reference: [Azure Key Vault Secrets client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/keyvault-secrets-readme)

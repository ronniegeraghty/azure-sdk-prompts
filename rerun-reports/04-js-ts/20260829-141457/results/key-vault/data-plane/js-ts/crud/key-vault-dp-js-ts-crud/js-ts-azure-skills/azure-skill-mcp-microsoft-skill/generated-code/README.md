# Azure Key Vault secret CRUD (TypeScript)

This sample creates, reads, updates, soft-deletes, and permanently purges
`my-secret` by using `DefaultAzureCredential`.

## Required packages

Runtime:

```powershell
npm install @azure/identity @azure/keyvault-secrets
```

TypeScript tooling used by this project:

```powershell
npm install --save-dev typescript tsx @types/node
```

## Run

Authenticate locally with any credential supported by `DefaultAzureCredential`,
then set the vault URL:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
npm install
npm start
```

The identity needs permission to get, set, delete, and purge secrets. Purging is
irreversible and fails when purge protection is enabled.

SDK reference:
https://learn.microsoft.com/javascript/api/overview/azure/keyvault-secrets-readme

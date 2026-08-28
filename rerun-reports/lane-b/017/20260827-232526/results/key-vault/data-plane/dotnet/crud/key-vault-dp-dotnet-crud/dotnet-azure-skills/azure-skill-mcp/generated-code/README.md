# Azure Key Vault secret CRUD sample

This console application creates `my-secret`, reads and prints its value, updates
it, then deletes and permanently purges it.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

The project currently references:

- `Azure.Identity` 1.21.0
- `Azure.Security.KeyVault.Secrets` 4.11.0

## Authentication and permissions

`DefaultAzureCredential` can use a developer login from Azure CLI or Visual
Studio locally, and managed identity when hosted in Azure. The authenticated
identity needs permission to get, set, delete, and purge secrets. With Key Vault
access policies, grant the `Get`, `Set`, `Delete`, and `Purge` secret
permissions. With Azure RBAC, assign an appropriate Key Vault data-plane role
that includes these operations.

## Run

Set the vault URI, not its resource ID:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

The vault must have soft delete enabled, and purge protection must be disabled
for immediate purge to succeed. The sample intentionally prints secret values
as requested; avoid logging secret values in production applications.

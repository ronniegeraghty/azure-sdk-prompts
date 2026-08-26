# Azure Key Vault Secrets CRUD

This .NET console application uses `DefaultAzureCredential` to create, read,
update, delete, and purge the `my-secret` secret.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

The project currently resolves these package versions:

- `Azure.Identity` 1.21.0
- `Azure.Security.KeyVault.Secrets` 4.11.0

## Run

Set the vault URL without putting credentials in source code:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

`DefaultAzureCredential` can use local developer credentials and automatically
use managed identity when hosted in Azure. The authenticated identity needs Key
Vault data-plane permissions to get, set, delete, and purge secrets. The vault
must have soft delete enabled and purge protection disabled for immediate purge.

Updating the value with `SetSecretAsync` creates a new version of the existing
secret. The application waits for the delete operation before purging it.

Reference: [Azure Key Vault secret client library quickstart for
.NET](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-net)

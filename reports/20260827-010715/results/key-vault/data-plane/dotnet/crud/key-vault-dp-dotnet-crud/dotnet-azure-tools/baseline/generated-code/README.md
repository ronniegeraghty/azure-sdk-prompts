# Azure Key Vault secret CRUD sample

This console application creates, reads, updates, deletes, and permanently
purges `my-secret`. It authenticates with `DefaultAzureCredential`.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity --version 1.13.2
dotnet add package Azure.Security.KeyVault.Secrets --version 4.7.0
```

The packages are already declared in `KeyVaultCrud.csproj`.

## Run

Set the vault URL, authenticate using any credential supported by
`DefaultAzureCredential`, and run the project:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

The authenticated identity needs permissions to get, set, delete, and purge
secrets. For an RBAC-enabled vault, assign an appropriate data-plane role such
as **Key Vault Administrator**. Purge protection must be disabled; a
purge-protected secret cannot be permanently purged until its retention period
expires.

# Azure Key Vault secrets CRUD

This .NET 8 console application creates, reads, updates, deletes, and purges
the `my-secret` secret. It authenticates with `DefaultAzureCredential`.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity --version 1.13.2
dotnet add package Azure.Security.KeyVault.Secrets --version 4.7.0
```

The project file already includes these package references.

## Run

Set the vault URI through an environment variable, then run the application:

```powershell
$env:KEY_VAULT_URI = "https://your-vault-name.vault.azure.net/"
dotnet run
```

The authenticated identity needs permissions to set, get, delete, and purge
secrets. For a vault using Azure RBAC, the **Key Vault Secrets Officer** role
includes these secret-management permissions. Purge also requires that purge
protection is disabled; Azure Key Vault does not allow purging when purge
protection is enabled.

`DefaultAzureCredential` tries supported credential sources in order, such as
environment-based service principal credentials, workload identity, managed
identity, and developer credentials.

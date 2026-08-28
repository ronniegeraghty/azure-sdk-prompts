# Azure Key Vault secret CRUD console app

This .NET 8 console application uses `DefaultAzureCredential` and performs
these operations on the `my-secret` secret:

1. Creates it with the value `my-secret-value`.
2. Reads it and prints the value.
3. Updates it to `updated-value` by creating a new secret version.
4. Deletes it, waits for soft deletion to complete, and purges it.

## Required NuGet packages

```powershell
dotnet add package Azure.Security.KeyVault.Secrets --version 4.11.0
```

The package is already referenced by `KeyVaultSecretCrud.csproj`. Version
4.11.0 depends on `Azure.Core` 1.53 or later, which includes
`DefaultAzureCredential`; a separate `Azure.Identity` reference is not needed
with this SDK generation.

## Configuration and run

Set the vault URL without putting credentials in source code:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

For local development, `DefaultAzureCredential` can use a supported developer
login, such as Azure CLI or Visual Studio authentication. In Azure, assign a
managed identity to the host. The authenticated identity needs the **Key Vault
Secrets Officer** RBAC role (or equivalent `get`, `set`, `delete`, and `purge`
secret permissions) scoped to the vault.

The vault must have soft delete enabled and purge protection disabled for the
immediate purge step to succeed. Running the application changes and then
permanently purges the named secret.

Reference: [Quickstart: Azure Key Vault secret client library for .NET](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-net)

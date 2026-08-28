# Azure Key Vault secret CRUD console app

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

The project file already references both packages.

## Authentication and permissions

`DefaultAzureCredential` can use local developer credentials or a managed
identity. The identity needs secret permissions to get, set, delete, and purge.
For an RBAC-enabled vault, assign an appropriate data-plane role such as
**Key Vault Administrator**. Purging is unavailable while purge protection is
enabled and may also be blocked by organizational policy.

Set the vault URL before running:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

For service-principal authentication, set `AZURE_TENANT_ID`,
`AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`. In Azure-hosted environments,
`DefaultAzureCredential` can use managed identity instead.

The application creates `my-secret`, reads and prints its value, creates a new
version with `updated-value`, waits for soft deletion to finish, and purges the
deleted secret.

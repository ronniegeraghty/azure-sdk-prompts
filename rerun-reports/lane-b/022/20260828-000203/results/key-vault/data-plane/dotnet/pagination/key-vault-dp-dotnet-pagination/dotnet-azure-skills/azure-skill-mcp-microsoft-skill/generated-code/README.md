# Azure Key Vault secret paging sample

This console program lists secret metadata page-by-page. It does not retrieve or
print secret values. Disabled secrets are included because listing properties
does not require reading their values.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

The project file pins known compatible package versions.

## Run

The authenticated identity needs permission to list secrets, such as the
**Key Vault Secrets User** role when the vault uses Azure RBAC.

```powershell
$env:AZURE_KEYVAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

`DefaultAzureCredential` can authenticate through supported local development
credentials or a managed identity when the program runs in Azure.

## Paging concepts

- `Pageable<SecretProperties>` is the synchronous sequence returned by
  `GetPropertiesOfSecrets()`.
- `AsyncPageable<SecretProperties>` is the asynchronous sequence returned by
  `GetPropertiesOfSecretsAsync()`.
- Calling `AsPages()` exposes each service response as a
  `Page<SecretProperties>`.
- `Page<T>.Values` contains the items in that response.
- `Page<T>.ContinuationToken` identifies the next page. The Azure SDK follows
  it automatically as the `foreach` or `await foreach` advances.
- `pageSizeHint` is a requested page size, not a guarantee; the service controls
  the actual number of items returned.
- Enumeration is lazy, so the program keeps only one page of a large result set
  in memory at a time.

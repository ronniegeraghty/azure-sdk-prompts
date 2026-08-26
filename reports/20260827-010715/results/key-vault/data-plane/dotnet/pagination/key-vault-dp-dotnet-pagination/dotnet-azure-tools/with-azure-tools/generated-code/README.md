# Azure Key Vault secret pagination

This console sample lists secret metadata with `SecretClient` and
`DefaultAzureCredential`. It does not retrieve or print secret values.
`GetPropertiesOfSecrets` returns both enabled and disabled secrets.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

`Azure.Core`, which defines `Page<T>`, `Pageable<T>`, and
`AsyncPageable<T>`, is included transitively.

## Run

Set the vault URI and authenticate using any credential supported by
`DefaultAzureCredential`. The identity needs permission to list secrets, such
as the **Key Vault Secrets User** role when the vault uses Azure RBAC.

```powershell
$env:AZURE_KEY_VAULT_URI = "https://my-vault.vault.azure.net/"

# Preferred for network I/O
dotnet run -- --async

# Synchronous equivalent
dotnet run -- --sync

# Demonstrate both (lists the same secrets twice)
dotnet run -- --both
```

## How pagination works

- `GetPropertiesOfSecretsAsync()` returns
  `AsyncPageable<SecretProperties>`. It is lazy: no page is requested until
  `await foreach` advances the sequence.
- `AsPages(pageSizeHint: 25)` exposes each response as
  `Page<SecretProperties>`. The service can choose a different page size, so
  the hint is not a guarantee.
- `Page<T>.Values` is the current page's `IReadOnlyList<T>`.
- `Page<T>.ContinuationToken` identifies the next page. The SDK follows it
  automatically as iteration continues; pass a saved token to `AsPages` to
  resume later.
- `Pageable<T>` and a normal `foreach` provide the synchronous equivalent,
  but asynchronous iteration is preferred for HTTP operations.

The listing operation returns only metadata for the current version of each
secret. It does not return secret values or individual secret versions.

Reference: [Pagination with the Azure SDK for .NET](https://learn.microsoft.com/dotnet/azure/sdk/pagination)

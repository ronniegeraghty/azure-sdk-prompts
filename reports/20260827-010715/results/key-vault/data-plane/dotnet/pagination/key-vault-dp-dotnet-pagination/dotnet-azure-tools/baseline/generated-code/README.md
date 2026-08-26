# Azure Key Vault pagination with C#

This sample lists secret **properties** without downloading secret values. Listing
properties includes disabled secrets, while calling `GetSecret` for a disabled
secret would fail.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity --version 1.14.2
dotnet add package Azure.Security.KeyVault.Secrets --version 4.8.0
```

`Azure.Security.KeyVault.Secrets` brings in the `Azure.Core` dependency that
defines `Page<T>`, `Pageable<T>`, and `AsyncPageable<T>`.

## Run

Authenticate locally with any credential supported by `DefaultAzureCredential`,
then pass the vault URL and the iteration mode:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"

dotnet run -- "https://<vault-name>.vault.azure.net/" async
dotnet run -- "https://<vault-name>.vault.azure.net/" sync
```

The identity needs permission to list secrets. For Azure RBAC, the
**Key Vault Secrets User** role includes that permission.

## How pagination works

- `GetPropertiesOfSecretsAsync()` returns `AsyncPageable<SecretProperties>`.
  It represents a lazily fetched result set; it does not load hundreds of
  secrets immediately.
- `AsPages(pageSizeHint: 100)` changes iteration from individual items to
  `Page<SecretProperties>` objects. The service may return fewer or more items
  because the page size is only a hint.
- Each `Page<T>` contains `Values`, the items returned in that response, and
  `ContinuationToken`, the opaque position used to request the next page.
- `await foreach` requests each page asynchronously as it is needed. The sync
  equivalent is `Pageable<T>` plus ordinary `foreach`.
- To resume from a saved token, pass it to
  `AsPages(continuationToken: savedToken, pageSizeHint: 100)`. Treat the token
  as opaque and persist it only if the application needs resumable scans.

The sample prints whether a continuation token exists rather than printing the
token itself.

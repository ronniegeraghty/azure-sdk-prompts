# Azure Key Vault pagination sample

This console program lists secret **properties** from an Azure Key Vault. It
does not download secret values. `GetPropertiesOfSecrets` includes both enabled
and disabled secrets, so disabled entries are printed with a `disabled` status
instead of being skipped or read.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

`Azure.Security.KeyVault.Secrets` brings in `Azure.Core`, which defines
`Pageable<T>`, `AsyncPageable<T>`, and `Page<T>`.

## Run

Authenticate with any identity supported by `DefaultAzureCredential`, then set
the vault URL:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://my-vault.vault.azure.net/"
dotnet run -- --mode both
```

Valid modes are `sync`, `async`, and `both`; the default is `both`. The identity
must have permission to list secrets, such as the **Key Vault Secrets User**
RBAC role at an appropriate scope.

## How paging works

- `Pageable<SecretProperties>` supports synchronous `foreach`.
- `AsyncPageable<SecretProperties>` implements `IAsyncEnumerable<T>` and
  supports `await foreach`.
- Calling `AsPages()` on either sequence changes item-at-a-time iteration into
  page-at-a-time iteration.
- Each `Page<SecretProperties>` contains `Values` for that response and a
  `ContinuationToken` for resuming with the next page.
- Enumeration is lazy. The SDK requests pages as the loops advance rather than
  loading hundreds of entries into memory first.
- `pageSizeHint` is only a request to the service; a service may ignore it or
  return fewer items. The loops must not assume a fixed page size.

To resume async enumeration from a saved token, pass it back to `AsPages`:

```csharp
await foreach (Page<SecretProperties> page in
    client.GetPropertiesOfSecretsAsync().AsPages(savedToken, pageSizeHint: 25))
{
    // Process page.Values and persist page.ContinuationToken after the page.
}
```

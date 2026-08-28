# Azure Key Vault pagination sample

This console app lists secret **properties**, including enabled and disabled
secrets, without downloading secret values. It uses `SecretClient` with
`DefaultAzureCredential` and supports explicit synchronous and asynchronous
page iteration.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
```

The project pins `Azure.Identity` 1.21.0 and
`Azure.Security.KeyVault.Secrets` 4.11.0 for reproducible builds. The core
`Azure`, `Page<T>`, `Pageable<T>`, and `AsyncPageable<T>` types arrive as
transitive dependencies of the Key Vault package.

## Run

Set the vault URL; `DefaultAzureCredential` then uses the available local
developer credential or the workload's managed identity:

```powershell
$env:KEY_VAULT_URL = "https://my-vault.vault.azure.net/"

dotnet run -- async
dotnet run -- sync
dotnet run -- both
```

The identity needs the `secrets/list` data-plane permission. In production,
prefer a managed identity and grant only the permissions it needs.

## How pagination works

`GetPropertiesOfSecretsAsync()` returns an
`AsyncPageable<SecretProperties>`. It is lazy: creating it does not retrieve
every secret or build one large in-memory list. Iteration performs service
requests as pages are needed.

Calling `AsPages(pageSizeHint: 50)` exposes each response as a
`Page<SecretProperties>`:

- `Page<T>.Values` is the read-only collection of items in that response.
- `Page<T>.ContinuationToken` identifies where another request can continue;
  `null` means that enumeration has reached the end.
- The page size is a hint. The service can return a different number of items.
- `await foreach` waits asynchronously between page requests. Normal `foreach`
  blocks while the synchronous request completes.

The SDK automatically follows continuation tokens during enumeration. To
resume manually, retain a token and pass it as the `continuationToken`
argument to `AsPages`. Avoid converting hundreds of results to a list unless
all items really must be held in memory.

`GetPropertiesOfSecrets[Async]()` returns only metadata for the current
versions and includes disabled secrets. The sample checks the nullable
`Enabled` property and prints `disabled` without attempting to fetch the
secret value.

## References

- [Pagination with the Azure SDK for .NET](https://learn.microsoft.com/dotnet/azure/sdk/pagination)
- [Azure Key Vault secret client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme)
- [`SecretClient.GetPropertiesOfSecretsAsync`](https://learn.microsoft.com/dotnet/api/azure.security.keyvault.secrets.secretclient.getpropertiesofsecretsasync)

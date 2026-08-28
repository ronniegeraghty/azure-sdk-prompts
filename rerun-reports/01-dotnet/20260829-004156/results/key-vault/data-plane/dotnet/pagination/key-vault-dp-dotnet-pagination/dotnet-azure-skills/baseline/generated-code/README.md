# Azure Key Vault secret paging

This sample lists secret metadata with `SecretClient` and
`DefaultAzureCredential`. It does not retrieve or print secret values.

## Required packages

```powershell
dotnet add package Azure.Identity --version 1.14.2
dotnet add package Azure.Security.KeyVault.Secrets --version 4.8.0
```

`Azure.Security.KeyVault.Secrets` brings in the `Azure.Core` types used for
paging, including `Pageable<T>`, `AsyncPageable<T>`, and `Page<T>`.

## Run

Sign in with a credential supported by `DefaultAzureCredential`, such as Azure
CLI, Visual Studio, environment variables, or a managed identity. The identity
needs permission to list secrets in the vault.

```powershell
dotnet run -- https://my-vault.vault.azure.net/ --async
dotnet run -- https://my-vault.vault.azure.net/ --sync
dotnet run -- https://my-vault.vault.azure.net/ --both
```

The default mode is `--both`. It lists the same metadata twice so the two
iteration styles can be compared.

## How paging works

`GetPropertiesOfSecretsAsync()` returns `AsyncPageable<SecretProperties>`.
No complete in-memory list is created. An `await foreach` requests results
lazily as iteration advances.

Calling `AsPages()` changes the iteration unit from one
`SecretProperties` object to one `Page<SecretProperties>`:

- `page.Values` contains the items returned by that service request.
- `page.ContinuationToken` identifies the next page and is `null` on the last
  page.
- `page.GetRawResponse()` provides the underlying Azure HTTP response when
  headers or status information are needed.
- `pageSizeHint` is a request hint; the service can return a different number
  of items.

The synchronous equivalent is `Pageable<SecretProperties>`, returned by
`GetPropertiesOfSecrets()`. Iterating either pageable directly hides page
boundaries:

```csharp
await foreach (SecretProperties secret in client.GetPropertiesOfSecretsAsync())
{
    Console.WriteLine(secret.Name);
}
```

Using `AsPages()` is preferable when logging page progress, checkpointing a
continuation token, or processing large results in page-sized batches.

Disabled secrets are still returned by the metadata-list operation. The sample
checks the nullable `SecretProperties.Enabled` property and prints `Disabled`,
`Enabled`, or `Not specified`; it does not call `GetSecret`, so disabled
secrets do not interrupt the listing.

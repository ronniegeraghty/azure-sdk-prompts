# Azure Storage management-plane LRO sample

This console app starts storage-account creation with `WaitUntil.Started`, then
either lets the SDK poll with `WaitForCompletionAsync` or polls explicitly with
`UpdateStatusAsync`.

## Packages

```powershell
dotnet add package Azure.ResourceManager.Storage --version 1.7.0
dotnet add package Azure.Identity --version 1.21.0
```

`Azure.ResourceManager.Storage` brings in the core `Azure.ResourceManager`
dependency that defines `ArmClient` and `ArmOperation<T>`.

## Configuration

The example contains no credentials or fixed Azure resource identifiers.
`DefaultAzureCredential` is convenient for local development and uses the
developer identity configured in the environment. For an Azure-hosted
production app, use a specific `ManagedIdentityCredential` instead.

Set these environment variables before running:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-lowercase-name>"
$env:AZURE_LOCATION = "eastus"       # optional; default: eastus
$env:LRO_TIMEOUT_SECONDS = "600"     # optional; default: 600
$env:LRO_POLL_SECONDS = "10"         # optional; default: 10
```

The identity needs permission to create storage accounts in the target resource
group. Apply least privilege at the narrowest practical scope.

> Running either command below creates or updates a real Azure resource and may
> incur charges. This repository only builds the sample; it does not execute it.

## Automatic SDK polling

```powershell
dotnet run -- --mode automatic
```

`CreateOrUpdateAsync(WaitUntil.Started, ...)` returns an
`ArmOperation<StorageAccountResource>` as soon as the service accepts the
request. `WaitForCompletionAsync` then follows the service's LRO protocol,
honors retry guidance from response headers, and returns
`Response<StorageAccountResource>` when complete. The cancellation token bounds
the client-side wait. A timeout stops polling but does **not** cancel the
server-side create operation.

## Manual polling

```powershell
dotnet run -- --mode manual
```

Manual polling calls `UpdateStatusAsync`, examines `HasCompleted`, `HasValue`,
the latest HTTP response, and waits between calls. It is useful when an
application needs per-poll logging or custom orchestration. The application is
then responsible for polling cadence and cancellation. Prefer
`WaitForCompletionAsync` when no custom behavior is required.

`WaitForCompletion` is the synchronous equivalent of
`WaitForCompletionAsync`; both hide the polling loop. In asynchronous
applications, prefer the async method so no thread is blocked.

## `ArmOperation<T>` pattern

- `Id` identifies the LRO when the service supplies an operation identifier.
- `HasCompleted` indicates terminal success or failure.
- `HasValue` indicates successful completion with a final resource.
- `Value` is valid only after successful completion.
- `GetRawResponse()` exposes the most recent HTTP response.
- `UpdateStatusAsync()` performs one explicit status refresh.
- `WaitForCompletionAsync()` repeatedly refreshes until completion.

References:

- [StorageAccountCollection.CreateOrUpdateAsync](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.storage.storageaccountcollection.createorupdateasync?view=azure-dotnet)
- [ArmOperation&lt;T&gt;](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.armoperation-1?view=azure-dotnet)
- [Azure Identity authentication guidance](https://learn.microsoft.com/dotnet/azure/sdk/authentication/)

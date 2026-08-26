# Azure Storage management-plane LRO sample

This console app creates an Azure Storage account through
`Azure.ResourceManager.Storage`. It starts the request with
`CreateOrUpdateAsync(WaitUntil.Started, ...)`, then demonstrates either SDK-managed
waiting or explicit polling. The sample is provided for review and local
compilation; running it performs a real Azure create/update operation.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager.Storage
```

`Azure.ResourceManager.Storage` brings in the core Azure Resource Manager
dependencies transitively. `Azure.Identity` supplies `DefaultAzureCredential`.

## Configuration and execution

`DefaultAzureCredential` checks its supported credential sources in order, such
as environment-based service-principal credentials, workload identity, Azure
CLI login, and developer-tool credentials.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-name>"
$env:AZURE_LOCATION = "eastus"

# SDK-managed polling:
dotnet run

# Manual polling with status output:
dotnet run -- --manual
```

The resource group must already exist. The storage account name must be 3-24
lowercase letters or digits and globally unique.

## `ArmOperation<T>` pattern

Azure Resource Manager create, update, and delete requests are often
long-running operations (LROs). Passing `WaitUntil.Started` makes
`CreateOrUpdateAsync` return an `ArmOperation<StorageAccountResource>` as soon as
the service accepts the request. It does not mean the storage account is ready.

`ArmOperation<T>` tracks the latest service response:

- `HasCompleted` reports whether the LRO reached a terminal state.
- `HasValue` reports whether a successful result is available.
- `Value` is the final `StorageAccountResource`; access it only after successful
  completion.
- `GetRawResponse()` exposes the most recent HTTP response.
- `UpdateStatusAsync()` performs one explicit status refresh.
- `WaitForCompletionAsync()` repeatedly refreshes status and returns the final
  `Response<T>`.

`WaitForCompletionAsync` is concise and should be the default when no per-poll
behavior is needed. Manual polling is useful for logging progress, publishing
metrics, or performing work between polls, but the caller must implement the
delay, status refresh, timeout, and cancellation behavior correctly.

Both sample paths enforce a ten-minute timeout with a linked
`CancellationTokenSource`. Timing out stops the local wait only. It does not
cancel the server-side Azure operation, which may continue and eventually
succeed or fail.

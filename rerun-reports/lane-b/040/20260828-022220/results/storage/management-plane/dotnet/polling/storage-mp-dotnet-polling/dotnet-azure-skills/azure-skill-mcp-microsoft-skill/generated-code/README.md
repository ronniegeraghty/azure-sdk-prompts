# Azure Storage account LRO sample

This console program starts a storage account create/update request with
`WaitUntil.Started`, then completes it using either SDK-managed or manual
polling.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager.Storage
```

`Azure.ResourceManager.Storage` brings in the common
`Azure.ResourceManager` and `Azure.Core` dependencies transitively. This
project currently resolves `Azure.Identity` 1.21.0 and
`Azure.ResourceManager.Storage` 1.7.0.

## Configuration

The signed-in identity needs permission to create storage accounts in the
target resource group, such as a narrowly scoped custom role or Storage
Account Contributor assignment.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-name>"
$env:AZURE_LOCATION = "eastus"
```

The program defaults to a local dry run. Enabling the following switch makes
it send a real management-plane create request:

```powershell
$env:AZURE_ENABLE_LIVE_CREATION = "true"
```

Use the SDK-managed wait:

```powershell
dotnet run -- wait
```

Use manual polling:

```powershell
dotnet run -- manual
```

## `ArmOperation<T>` and LRO handling

`CreateOrUpdateAsync(WaitUntil.Started, ...)` returns an
`ArmOperation<StorageAccountResource>` after Azure accepts the request. The
operation is a local handle to the server-side long-running operation, not the
finished storage account.

- `Id` identifies the operation.
- `UpdateStatusAsync` makes one status request and refreshes the handle.
- `HasCompleted` says whether the LRO reached a terminal state.
- `HasValue` says whether a successful result is available.
- `Value` returns the `StorageAccountResource` only after successful
  completion.
- `GetRawResponse()` exposes the latest HTTP response.

In `wait` mode, `WaitForCompletionAsync` owns the polling loop, honors the
requested polling interval and cancellation token, and surfaces a failed Azure
operation as `RequestFailedException`. This is the preferred approach for most
applications.

In `manual` mode, the application calls `UpdateStatusAsync`, inspects
`HasCompleted`, logs each response, and delays between requests. Manual polling
is useful when each poll must drive custom progress reporting or orchestration,
but the application must implement delay, cancellation, and terminal-state
handling correctly.

The ten-minute `CancellationTokenSource` bounds both approaches. A timeout
cancels the client's wait; it does not cancel the server-side create operation.
The resource may therefore finish provisioning after this program exits.

`DefaultAzureCredential` is convenient for local development. For an
Azure-hosted production application, prefer a specific managed identity
credential so authentication behavior is deterministic.

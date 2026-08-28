# Azure Storage account LRO sample

This console app starts a storage-account create/update operation with
`WaitUntil.Started`, then either delegates polling to `WaitForCompletionAsync`
or polls explicitly with `UpdateStatusAsync`.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager.Storage
```

`Azure.ResourceManager.Storage` brings in the common Azure Resource Manager and
Azure Core dependencies transitively.

## Run

Authenticate locally with a developer credential supported by
`DefaultAzureCredential`, then run one of the polling modes:

```powershell
dotnet run -- <resource-group> <globally-unique-account-name> eastus wait 300
dotnet run -- <resource-group> <globally-unique-account-name> eastus manual 300
```

The identity needs permission to read the resource group and create storage
accounts in it. `DefaultAzureCredential` is convenient for local development;
Azure-hosted production applications should normally select a specific managed
identity credential.

## `ArmOperation<T>` and polling

`CreateOrUpdateAsync(WaitUntil.Started, ...)` returns as soon as Azure accepts
the request. Its `ArmOperation<StorageAccountResource>` represents the
server-side long-running operation:

- `HasCompleted` reports whether the latest observed state is terminal.
- `GetRawResponse().Status` exposes the HTTP status from the latest request.
- `UpdateStatusAsync` performs one explicit status refresh.
- `Value` contains the created resource after successful completion.
- `WaitForCompletionAsync` performs the status refreshes for you and returns
  the final `Response<StorageAccountResource>`.

The `wait` mode is the normal choice: the SDK honors service polling guidance
and owns the polling loop. The `manual` mode is useful when the application
needs per-poll logging, a custom cadence, or work between polls, but the
application must implement delay, cancellation, and status refresh correctly.

Both modes use a timeout-backed cancellation token. A timeout stops this
process from waiting; it does **not** cancel or roll back the Azure operation.
The operation may finish after the program exits, so production code should
record enough resource context to reconcile its state later.

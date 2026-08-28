# Azure Storage management-plane LRO sample

This .NET 8 sample starts a storage-account create operation with
`CreateOrUpdateAsync(WaitUntil.Started, ...)`, then completes it with either
SDK-managed polling or a manual polling loop. It uses `DefaultAzureCredential`
and performs no Azure operation unless `--execute` is supplied.

## Required packages

```powershell
dotnet add package Azure.Identity --version 1.17.0
dotnet add package Azure.ResourceManager.Storage --version 1.4.0
```

`Azure.ResourceManager.Storage` brings in the ARM core and resource-management
dependencies used by the sample.

## Run

Authenticate using any credential supported by `DefaultAzureCredential`, such
as Azure CLI login, Visual Studio, managed identity, or service-principal
environment variables. The principal needs permission to create storage
accounts in the target resource group.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-name>"
$env:AZURE_LOCATION = "eastus"
$env:LRO_TIMEOUT_MINUTES = "10"

dotnet run -- --execute wait
dotnet run -- --execute manual
```

Running `dotnet run` without `--execute` is a local dry run and does not
authenticate or send an Azure request.

## The `ArmOperation<T>` pattern

ARM create/update/delete calls often return before the server-side work is
finished. Passing `WaitUntil.Started` makes `CreateOrUpdateAsync` return an
`ArmOperation<StorageAccountResource>` as soon as the request has been accepted.
The operation exposes:

- `Id`: the server-side operation identifier.
- `HasCompleted`: whether the terminal state has been reached.
- `HasValue` and `Value`: whether and what resource was returned on success.
- `GetRawResponse()`: the latest HTTP response known to the operation.
- `UpdateStatusAsync()`: one explicit status refresh.
- `WaitForCompletionAsync()`: the SDK-managed polling loop.

`WaitForCompletionAsync(interval, token)` is concise and lets the SDK follow
Azure's LRO protocol until completion. Manual polling repeatedly calls
`UpdateStatusAsync`, inspects the response and `HasCompleted`, and controls the
delay itself. Manual polling is useful for progress logging or application-level
scheduling, but the application must avoid aggressive polling and correctly
handle cancellation and terminal failures.

The timeout is implemented with a cancellation token. Cancellation only stops
the client from waiting; it does not necessarily cancel the operation running
in Azure. Preserve the operation ID so its eventual status can be investigated.

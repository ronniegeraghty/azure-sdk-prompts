using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;
using System.Diagnostics;

const string Usage =
    "Usage: dotnet run -- <resource-group> <storage-account> <location> <wait|manual> [timeout-seconds]";

if (args.Length is < 4 or > 5 ||
    !Enum.TryParse<PollingMode>(args[3], ignoreCase: true, out PollingMode mode) ||
    (args.Length == 5 &&
     (!int.TryParse(args[4], out int parsedTimeoutSeconds) || parsedTimeoutSeconds <= 0)))
{
    Console.Error.WriteLine(Usage);
    return 2;
}

string resourceGroupName = args[0];
string storageAccountName = args[1];
AzureLocation location = new(args[2]);
TimeSpan timeout = TimeSpan.FromSeconds(
    args.Length == 5 ? int.Parse(args[4]) : 300);

var armClient = new ArmClient(new DefaultAzureCredential());
SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();
ResourceGroupResource resourceGroup =
    await subscription.GetResourceGroupAsync(resourceGroupName);
StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

var content = new StorageAccountCreateOrUpdateContent(
    new StorageSku(StorageSkuName.StandardLrs),
    StorageKind.StorageV2,
    location)
{
    AllowBlobPublicAccess = false,
    EnableHttpsTrafficOnly = true,
    MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2
};

Console.WriteLine(
    $"Starting create/update for '{storageAccountName}' with a {timeout.TotalSeconds:F0}-second timeout.");

ArmOperation<StorageAccountResource> operation =
    await storageAccounts.CreateOrUpdateAsync(
        WaitUntil.Started,
        storageAccountName,
        content);

Console.WriteLine(
    $"Started. HTTP status: {operation.GetRawResponse().Status}; completed: {operation.HasCompleted}");

try
{
    StorageAccountResource account = mode switch
    {
        PollingMode.Wait => await WaitForCompletionAsync(operation, timeout),
        PollingMode.Manual => await PollManuallyAsync(operation, timeout),
        _ => throw new UnreachableException()
    };

    Console.WriteLine($"Succeeded: {account.Id}");
    return 0;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine(
        $"Timed out after {timeout}. The local wait stopped, but the Azure operation may still be running.");
    Console.Error.WriteLine(
        $"Last HTTP status: {operation.GetRawResponse().Status}; completed: {operation.HasCompleted}");
    return 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed ({ex.Status}, {ex.ErrorCode ?? "no error code"}): {ex.Message}");
    return 1;
}

static async Task<StorageAccountResource> WaitForCompletionAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan timeout)
{
    using var timeoutSource = new CancellationTokenSource(timeout);

    Console.WriteLine(
        "Using WaitForCompletionAsync: the SDK chooses the polling cadence and updates the operation.");

    Response<StorageAccountResource> response =
        await operation.WaitForCompletionAsync(timeoutSource.Token);

    Console.WriteLine(
        $"Final HTTP status: {operation.GetRawResponse().Status}; completed: {operation.HasCompleted}");
    return response.Value;
}

static async Task<StorageAccountResource> PollManuallyAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan timeout)
{
    using var timeoutSource = new CancellationTokenSource(timeout);
    TimeSpan pollingInterval = TimeSpan.FromSeconds(5);

    Console.WriteLine(
        "Using manual polling: this code controls the cadence and calls UpdateStatusAsync.");

    while (!operation.HasCompleted)
    {
        Console.WriteLine(
            $"In progress. HTTP status: {operation.GetRawResponse().Status}; " +
            $"next poll in {pollingInterval.TotalSeconds:F0}s.");

        await Task.Delay(pollingInterval, timeoutSource.Token);
        await operation.UpdateStatusAsync(timeoutSource.Token);
    }

    Console.WriteLine(
        $"Final HTTP status: {operation.GetRawResponse().Status}; completed: {operation.HasCompleted}");
    return operation.Value;
}

enum PollingMode
{
    Wait,
    Manual
}

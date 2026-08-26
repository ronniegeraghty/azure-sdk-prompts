using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

const string ManualOption = "--manual";

string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string storageAccountName = GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
string location = Environment.GetEnvironmentVariable("AZURE_LOCATION") ?? "eastus";

TimeSpan timeout = TimeSpan.FromMinutes(10);
TimeSpan pollingInterval = TimeSpan.FromSeconds(10);
bool useManualPolling = args.Contains(ManualOption, StringComparer.OrdinalIgnoreCase);

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential, subscriptionId);
SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();
ResourceGroupResource resourceGroup =
    await subscription.GetResourceGroups().GetAsync(resourceGroupName);
StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

var content = new StorageAccountCreateOrUpdateContent(
    new StorageSku(StorageSkuName.StandardLrs),
    StorageKind.StorageV2,
    new AzureLocation(location));

using var shutdown = new CancellationTokenSource();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    shutdown.Cancel();
};

try
{
    Console.WriteLine($"Starting creation of '{storageAccountName}'...");

    // WaitUntil.Started returns after Azure accepts the request, not after creation finishes.
    ArmOperation<StorageAccountResource> operation =
        await storageAccounts.CreateOrUpdateAsync(
            WaitUntil.Started,
            storageAccountName,
            content,
            shutdown.Token);

    Console.WriteLine(
        $"Accepted: HTTP {operation.GetRawResponse().Status}; " +
        $"completed={operation.HasCompleted}");

    StorageAccountResource account = useManualPolling
        ? await WaitWithManualPollingAsync(
            operation,
            pollingInterval,
            timeout,
            shutdown.Token)
        : await WaitWithSdkAsync(
            operation,
            pollingInterval,
            timeout,
            shutdown.Token);

    Console.WriteLine($"Created storage account: {account.Id}");
}
catch (OperationCanceledException) when (shutdown.IsCancellationRequested)
{
    Console.Error.WriteLine("Canceled by the user.");
    Environment.ExitCode = 2;
}
catch (TimeoutException ex)
{
    Console.Error.WriteLine($"Timed out: {ex.Message}");
    Environment.ExitCode = 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed: HTTP {ex.Status}, code={ex.ErrorCode}, {ex.Message}");
    Environment.ExitCode = 1;
}

static async Task<StorageAccountResource> WaitWithSdkAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan pollingInterval,
    TimeSpan timeout,
    CancellationToken cancellationToken)
{
    using var timeoutCts = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
    timeoutCts.CancelAfter(timeout);

    try
    {
        Console.WriteLine(
            "Waiting with ArmOperation.WaitForCompletionAsync. " +
            "Use --manual to display every poll.");

        Response<StorageAccountResource> response =
            await operation.WaitForCompletionAsync(pollingInterval, timeoutCts.Token);

        return response.Value;
    }
    catch (OperationCanceledException)
        when (!cancellationToken.IsCancellationRequested && timeoutCts.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"The operation did not finish within {timeout}. " +
            "Cancellation stops this client from waiting; it does not cancel the Azure operation.");
    }
}

static async Task<StorageAccountResource> WaitWithManualPollingAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan pollingInterval,
    TimeSpan timeout,
    CancellationToken cancellationToken)
{
    using var timeoutCts = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
    timeoutCts.CancelAfter(timeout);

    try
    {
        while (!operation.HasCompleted)
        {
            Response latestResponse = operation.GetRawResponse();
            Console.WriteLine(
                $"Polling: HTTP {latestResponse.Status}; " +
                $"completed={operation.HasCompleted}; hasValue={operation.HasValue}");

            await Task.Delay(pollingInterval, timeoutCts.Token);

            // Refreshes HasCompleted, HasValue, Value, and the latest raw response.
            await operation.UpdateStatusAsync(timeoutCts.Token);
        }
    }
    catch (OperationCanceledException)
        when (!cancellationToken.IsCancellationRequested && timeoutCts.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"The operation did not finish within {timeout}. " +
            "Cancellation stops polling; it does not cancel the Azure operation.");
    }

    Console.WriteLine(
        $"Final status: HTTP {operation.GetRawResponse().Status}; " +
        $"completed={operation.HasCompleted}; hasValue={operation.HasValue}");

    return operation.Value;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Set the required environment variable '{name}'.");
    }

    return value;
}

using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

const string ManualPollingArgument = "--manual";

string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string storageAccountName = GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
string location = Environment.GetEnvironmentVariable("AZURE_LOCATION") ?? "eastus";
TimeSpan timeout = TimeSpan.FromSeconds(GetPositiveInteger("AZURE_LRO_TIMEOUT_SECONDS", 600));
TimeSpan pollingInterval = TimeSpan.FromSeconds(
    GetPositiveInteger("AZURE_LRO_POLL_INTERVAL_SECONDS", 10));
bool useManualPolling = args.Contains(ManualPollingArgument, StringComparer.OrdinalIgnoreCase);

using var applicationStopping = new CancellationTokenSource();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    applicationStopping.Cancel();
};

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential, subscriptionId);

ResourceIdentifier resourceGroupId =
    ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);
StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

var content = new StorageAccountCreateOrUpdateContent(
    new StorageSku(StorageSkuName.StandardLrs),
    StorageKind.StorageV2,
    new AzureLocation(location))
{
    AllowBlobPublicAccess = false,
    EnableHttpsTrafficOnly = true,
    MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2
};

try
{
    Console.WriteLine($"Starting create/update for '{storageAccountName}'...");

    // WaitUntil.Started returns as soon as Azure accepts the request.
    ArmOperation<StorageAccountResource> operation =
        await storageAccounts.CreateOrUpdateAsync(
            WaitUntil.Started,
            storageAccountName,
            content,
            applicationStopping.Token);

    PrintStatus("Started", operation);

    using var timeoutTokenSource =
        CancellationTokenSource.CreateLinkedTokenSource(applicationStopping.Token);
    timeoutTokenSource.CancelAfter(timeout);

    StorageAccountResource account = useManualPolling
        ? await PollManuallyAsync(operation, pollingInterval, timeoutTokenSource.Token)
        : await WaitWithSdkPollingAsync(operation, pollingInterval, timeoutTokenSource.Token);

    Console.WriteLine($"Completed: {account.Id}");
    Console.WriteLine($"Provisioning state: {account.Data.ProvisioningState}");
}
catch (OperationCanceledException) when (!applicationStopping.IsCancellationRequested)
{
    Console.Error.WriteLine(
        $"Timed out after {timeout}. Stopping local polling does not cancel the Azure operation.");
    Environment.ExitCode = 2;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine("Canceled.");
    Environment.ExitCode = 3;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Authentication failed: {ex.Message}");
    Environment.ExitCode = 4;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
    Environment.ExitCode = 5;
}

static async Task<StorageAccountResource> WaitWithSdkPollingAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan pollingInterval,
    CancellationToken cancellationToken)
{
    // A status refresh is optional; it demonstrates inspection before handing polling to the SDK.
    if (!operation.HasCompleted)
    {
        Response statusResponse = await operation.UpdateStatusAsync(cancellationToken);
        PrintStatus($"Status check (HTTP {statusResponse.Status})", operation);
    }

    Console.WriteLine("Using WaitForCompletionAsync; the SDK now owns the polling loop.");
    Response<StorageAccountResource> completedResponse =
        await operation.WaitForCompletionAsync(pollingInterval, cancellationToken);

    return completedResponse.Value;
}

static async Task<StorageAccountResource> PollManuallyAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan pollingInterval,
    CancellationToken cancellationToken)
{
    Console.WriteLine("Using manual polling; the application owns delay and status refresh.");

    while (!operation.HasCompleted)
    {
        await Task.Delay(pollingInterval, cancellationToken);
        Response statusResponse = await operation.UpdateStatusAsync(cancellationToken);
        PrintStatus($"Polled (HTTP {statusResponse.Status})", operation);
    }

    if (!operation.HasValue)
    {
        throw new InvalidOperationException(
            "The operation completed without returning a storage account.");
    }

    return operation.Value;
}

static void PrintStatus(string label, ArmOperation<StorageAccountResource> operation)
{
    Response response = operation.GetRawResponse();
    Console.WriteLine(
        $"{label}: HTTP={response.Status}, HasCompleted={operation.HasCompleted}, " +
        $"HasValue={operation.HasValue}, OperationId={operation.Id}");
}

static string GetRequiredEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name) is { Length: > 0 } value
        ? value
        : throw new InvalidOperationException(
            $"Set the required environment variable {name}.");

static int GetPositiveInteger(string name, int defaultValue)
{
    string? value = Environment.GetEnvironmentVariable(name);
    if (value is null)
    {
        return defaultValue;
    }

    return int.TryParse(value, out int parsed) && parsed > 0
        ? parsed
        : throw new InvalidOperationException($"{name} must be a positive integer.");
}

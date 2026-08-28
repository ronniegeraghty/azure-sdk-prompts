using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

const string ManualMode = "manual";
const string ManagedMode = "managed";

string mode = args.FirstOrDefault()?.ToLowerInvariant() ?? ManagedMode;
if (mode is not (ManagedMode or ManualMode))
{
    Console.Error.WriteLine("Usage: dotnet run -- [managed|manual]");
    return 2;
}

try
{
    string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
    string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
    string accountName = GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
    AzureLocation location = new(
        Environment.GetEnvironmentVariable("AZURE_LOCATION") ?? "eastus");

    TimeSpan timeout = TimeSpan.FromMinutes(10);
    TimeSpan pollingInterval = TimeSpan.FromSeconds(10);

    TokenCredential credential = new DefaultAzureCredential();
    ArmClient armClient = new(credential, subscriptionId);

    ResourceIdentifier resourceGroupId =
        ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
    ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);
    StorageAccountCollection accounts = resourceGroup.GetStorageAccounts();

    StorageAccountCreateOrUpdateContent content = new(
        new StorageSku(StorageSkuName.StandardLrs),
        StorageKind.StorageV2,
        location)
    {
        AllowSharedKeyAccess = false
    };

    StorageAccountResource account = mode == ManagedMode
        ? await CreateWithSdkPollingAsync(
            accounts, accountName, content, pollingInterval, timeout)
        : await CreateWithManualPollingAsync(
            accounts, accountName, content, pollingInterval, timeout);

    Console.WriteLine($"Storage account ready: {account.Id}");
    return 0;
}
catch (TimeoutException ex)
{
    Console.Error.WriteLine($"Timeout: {ex.Message}");
    return 3;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Authentication failed: {ex.Message}");
    return 4;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed. Status={ex.Status}, ErrorCode={ex.ErrorCode}, Message={ex.Message}");
    return 5;
}
catch (ArgumentException ex)
{
    Console.Error.WriteLine(ex.Message);
    return 2;
}

static async Task<StorageAccountResource> CreateWithSdkPollingAsync(
    StorageAccountCollection accounts,
    string accountName,
    StorageAccountCreateOrUpdateContent content,
    TimeSpan pollingInterval,
    TimeSpan timeout)
{
    using CancellationTokenSource timeoutSource = new(timeout);

    try
    {
        ArmOperation<StorageAccountResource> operation =
            await accounts.CreateOrUpdateAsync(
                WaitUntil.Started,
                accountName,
                content,
                timeoutSource.Token);

        PrintStatus("Started", operation);

        // The SDK performs UpdateStatusAsync calls internally until the LRO finishes.
        Response<StorageAccountResource> completed =
            await operation.WaitForCompletionAsync(pollingInterval, timeoutSource.Token);

        PrintStatus("Completed", operation);
        return completed.Value;
    }
    catch (OperationCanceledException ex) when (timeoutSource.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"The SDK-managed wait exceeded {timeout}. The Azure operation may still be running.",
            ex);
    }
}

static async Task<StorageAccountResource> CreateWithManualPollingAsync(
    StorageAccountCollection accounts,
    string accountName,
    StorageAccountCreateOrUpdateContent content,
    TimeSpan pollingInterval,
    TimeSpan timeout)
{
    using CancellationTokenSource timeoutSource = new(timeout);

    try
    {
        ArmOperation<StorageAccountResource> operation =
            await accounts.CreateOrUpdateAsync(
                WaitUntil.Started,
                accountName,
                content,
                timeoutSource.Token);

        PrintStatus("Started", operation);

        while (!operation.HasCompleted)
        {
            await Task.Delay(pollingInterval, timeoutSource.Token);

            Response response = await operation.UpdateStatusAsync(timeoutSource.Token);
            Console.WriteLine(
                $"Polled at {DateTimeOffset.UtcNow:O}: " +
                $"HTTP {response.Status}, HasCompleted={operation.HasCompleted}");
        }

        PrintStatus("Completed", operation);

        if (!operation.HasValue)
        {
            throw new InvalidOperationException(
                "The operation completed without producing a storage account.");
        }

        return operation.Value;
    }
    catch (OperationCanceledException ex) when (timeoutSource.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"Manual polling exceeded {timeout}. The Azure operation may still be running.",
            ex);
    }
}

static void PrintStatus(string stage, ArmOperation<StorageAccountResource> operation)
{
    Response response = operation.GetRawResponse();
    Console.WriteLine(
        $"{stage}: OperationId={operation.Id}, HTTP={response.Status}, " +
        $"HasCompleted={operation.HasCompleted}, HasValue={operation.HasValue}");
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new ArgumentException($"Set the required environment variable {name}.");
}

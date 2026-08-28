using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

namespace StorageAccountLro;

internal static class Program
{
    private static readonly TimeSpan PollingInterval = TimeSpan.FromSeconds(10);
    private static readonly TimeSpan OperationTimeout = TimeSpan.FromMinutes(10);

    public static async Task<int> Main(string[] args)
    {
        string mode = args.FirstOrDefault()?.ToLowerInvariant() ?? "wait";

        if (mode is not ("wait" or "manual"))
        {
            Console.Error.WriteLine("Usage: dotnet run -- [wait|manual]");
            return 2;
        }

        string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
        string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
        string storageAccountName = GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
        string location = Environment.GetEnvironmentVariable("AZURE_LOCATION") ?? "eastus";

        // The sample is safe to build and inspect without changing an Azure subscription.
        if (!string.Equals(
                Environment.GetEnvironmentVariable("AZURE_ENABLE_LIVE_CREATION"),
                "true",
                StringComparison.OrdinalIgnoreCase))
        {
            Console.WriteLine(
                "Dry run only. Set AZURE_ENABLE_LIVE_CREATION=true to allow the create request.");
            Console.WriteLine(
                $"Mode={mode}, account={storageAccountName}, resourceGroup={resourceGroupName}, location={location}");
            return 0;
        }

        ArmClient armClient = new(new DefaultAzureCredential());
        ResourceIdentifier resourceGroupId =
            ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
        ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);
        StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

        StorageAccountCreateOrUpdateContent content = new(
            new StorageSku(StorageSkuName.StandardLrs),
            StorageKind.StorageV2,
            new AzureLocation(location))
        {
            AllowBlobPublicAccess = false,
            MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2,
            EnableHttpsTrafficOnly = true
        };

        using CancellationTokenSource timeout = new(OperationTimeout);

        try
        {
            // WaitUntil.Started returns as soon as Azure accepts the request. It does not
            // wait for the storage account to finish provisioning.
            ArmOperation<StorageAccountResource> operation =
                await storageAccounts.CreateOrUpdateAsync(
                    WaitUntil.Started,
                    storageAccountName,
                    content,
                    timeout.Token);

            Console.WriteLine($"Started operation {operation.Id}");

            StorageAccountResource account = mode == "manual"
                ? await WaitWithManualPollingAsync(operation, PollingInterval, timeout.Token)
                : await WaitWithSdkPollingAsync(operation, PollingInterval, timeout.Token);

            Console.WriteLine($"Created storage account: {account.Id}");
            return 0;
        }
        catch (OperationCanceledException) when (timeout.IsCancellationRequested)
        {
            Console.Error.WriteLine(
                $"Timed out after {OperationTimeout}. Only local polling was canceled; " +
                "the Azure operation may still be running.");
            return 3;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure request failed: HTTP {ex.Status}, code={ex.ErrorCode}, message={ex.Message}");
            return 1;
        }
    }

    private static async Task<StorageAccountResource> WaitWithSdkPollingAsync(
        ArmOperation<StorageAccountResource> operation,
        TimeSpan pollingInterval,
        CancellationToken cancellationToken)
    {
        // A single explicit refresh shows the status API before handing polling to the SDK.
        Response status = await operation.UpdateStatusAsync(cancellationToken);
        WriteStatus("SDK wait", operation, status);

        if (!operation.HasCompleted)
        {
            await operation.WaitForCompletionAsync(pollingInterval, cancellationToken);
        }

        return GetCompletedValue(operation);
    }

    private static async Task<StorageAccountResource> WaitWithManualPollingAsync(
        ArmOperation<StorageAccountResource> operation,
        TimeSpan pollingInterval,
        CancellationToken cancellationToken)
    {
        while (!operation.HasCompleted)
        {
            Response status = await operation.UpdateStatusAsync(cancellationToken);
            WriteStatus("Manual poll", operation, status);

            if (!operation.HasCompleted)
            {
                await Task.Delay(pollingInterval, cancellationToken);
            }
        }

        return GetCompletedValue(operation);
    }

    private static void WriteStatus(
        string strategy,
        ArmOperation<StorageAccountResource> operation,
        Response response)
    {
        Console.WriteLine(
            $"{DateTimeOffset.UtcNow:O} [{strategy}] " +
            $"HTTP={response.Status}, completed={operation.HasCompleted}, hasValue={operation.HasValue}");
    }

    private static StorageAccountResource GetCompletedValue(
        ArmOperation<StorageAccountResource> operation)
    {
        if (!operation.HasCompleted)
        {
            throw new InvalidOperationException("The operation has not completed.");
        }

        if (!operation.HasValue)
        {
            Response response = operation.GetRawResponse();
            throw new InvalidOperationException(
                $"The operation completed without a value. Last HTTP status: {response.Status}.");
        }

        return operation.Value;
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        return Environment.GetEnvironmentVariable(name)
            ?? throw new InvalidOperationException(
                $"Set the required environment variable {name}.");
    }
}

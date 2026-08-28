using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

namespace StorageAccountManagement;

internal static class Program
{
    private const string ExecuteFlag = "--execute";

    public static async Task<int> Main(string[] args)
    {
        if (!TryReadConfiguration(args, out Configuration configuration))
        {
            PrintUsage();
            return 2;
        }

        if (!configuration.Execute)
        {
            Console.WriteLine("Dry run only; no Azure requests were sent.");
            Console.WriteLine(
                $"Would create, list, inspect, enable blob versioning, and delete " +
                $"'{configuration.StorageAccountName}' in resource group " +
                $"'{configuration.ResourceGroupName}' ({AzureLocation.EastUS}).");
            Console.WriteLine($"Add {ExecuteFlag} to perform these operations.");
            return 0;
        }

        using CancellationTokenSource cancellationSource = new();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        return await ManageStorageAccountAsync(configuration, cancellationSource.Token);
    }

    private static async Task<int> ManageStorageAccountAsync(
        Configuration configuration,
        CancellationToken cancellationToken)
    {
        StorageAccountResource? createdAccount = null;
        Exception? operationError = null;

        try
        {
            ArmClientOptions clientOptions = new()
            {
                Retry =
                {
                    Mode = RetryMode.Exponential,
                    Delay = TimeSpan.FromSeconds(1),
                    MaxDelay = TimeSpan.FromSeconds(10),
                    MaxRetries = 5,
                    NetworkTimeout = TimeSpan.FromSeconds(100)
                }
            };

            DefaultAzureCredential credential = new();
            ArmClient armClient = new(credential, configuration.SubscriptionId, clientOptions);

            ResourceIdentifier resourceGroupId = ResourceGroupResource.CreateResourceIdentifier(
                configuration.SubscriptionId,
                configuration.ResourceGroupName);
            ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);

            Console.WriteLine($"Creating storage account '{configuration.StorageAccountName}'...");
            StorageAccountCreateOrUpdateContent createContent = new(
                new StorageSku(StorageSkuName.StandardLrs),
                StorageKind.StorageV2,
                AzureLocation.EastUS)
            {
                AllowBlobPublicAccess = false,
                AllowSharedKeyAccess = false,
                EnableHttpsTrafficOnly = true,
                MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2
            };

            StorageAccountCollection accounts = resourceGroup.GetStorageAccounts();
            ArmOperation<StorageAccountResource> createOperation =
                await accounts.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    configuration.StorageAccountName,
                    createContent,
                    cancellationToken);
            createdAccount = createOperation.Value;

            Console.WriteLine("Storage accounts in the resource group:");
            await foreach (StorageAccountResource account in
                accounts.GetAllAsync(cancellationToken: cancellationToken))
            {
                Console.WriteLine($"  {account.Data.Name}");
            }

            Response<StorageAccountResource> getResponse =
                await createdAccount.GetAsync(cancellationToken: cancellationToken);
            StorageAccountData accountData = getResponse.Value.Data;
            Console.WriteLine("Created storage account properties:");
            Console.WriteLine($"  Resource ID: {accountData.Id}");
            Console.WriteLine($"  Location:    {accountData.Location}");
            Console.WriteLine($"  Kind:        {accountData.Kind}");
            Console.WriteLine($"  SKU:         {accountData.Sku.Name}");
            Console.WriteLine($"  State:       {accountData.ProvisioningState}");

            BlobServiceResource blobService = createdAccount.GetBlobService();
            BlobServiceData blobServiceData = new()
            {
                IsVersioningEnabled = true
            };
            await blobService.CreateOrUpdateAsync(
                WaitUntil.Completed,
                blobServiceData,
                cancellationToken);
            Console.WriteLine("Blob versioning enabled.");
        }
        catch (AuthenticationFailedException exception)
        {
            operationError = exception;
            Console.Error.WriteLine($"Authentication failed: {exception.Message}");
        }
        catch (RequestFailedException exception)
        {
            operationError = exception;
            Console.Error.WriteLine(
                $"Azure request failed ({exception.Status}, {exception.ErrorCode}): " +
                exception.Message);
        }
        catch (OperationCanceledException exception)
        {
            operationError = exception;
            Console.Error.WriteLine("Operation canceled.");
        }

        if (createdAccount is not null)
        {
            try
            {
                Console.WriteLine($"Deleting storage account '{createdAccount.Data.Name}'...");
                await createdAccount.DeleteAsync(WaitUntil.Completed, CancellationToken.None);
                Console.WriteLine("Storage account deleted.");
            }
            catch (RequestFailedException exception)
            {
                operationError ??= exception;
                Console.Error.WriteLine(
                    $"Cleanup failed ({exception.Status}, {exception.ErrorCode}): " +
                    exception.Message);
            }
        }

        return operationError is null ? 0 : 1;
    }

    private static bool TryReadConfiguration(
        string[] args,
        out Configuration configuration)
    {
        string? subscriptionId = Environment.GetEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
        string? resourceGroupName = Environment.GetEnvironmentVariable("AZURE_RESOURCE_GROUP");
        string? storageAccountName = Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
        bool execute = args.Contains(ExecuteFlag, StringComparer.OrdinalIgnoreCase);

        if (string.IsNullOrWhiteSpace(subscriptionId) ||
            string.IsNullOrWhiteSpace(resourceGroupName) ||
            string.IsNullOrWhiteSpace(storageAccountName))
        {
            configuration = null!;
            return false;
        }

        configuration = new Configuration(
            subscriptionId,
            resourceGroupName,
            storageAccountName,
            execute);
        return true;
    }

    private static void PrintUsage()
    {
        Console.Error.WriteLine(
            "Set AZURE_SUBSCRIPTION_ID, AZURE_RESOURCE_GROUP, and " +
            "AZURE_STORAGE_ACCOUNT_NAME.");
        Console.Error.WriteLine(
            $"Run without {ExecuteFlag} for a dry run, or add {ExecuteFlag} to perform changes.");
        Console.Error.WriteLine(
            "Storage account names must be globally unique, 3-24 characters, " +
            "and contain only lowercase letters and digits.");
    }

    private sealed record Configuration(
        string SubscriptionId,
        string ResourceGroupName,
        string StorageAccountName,
        bool Execute);
}

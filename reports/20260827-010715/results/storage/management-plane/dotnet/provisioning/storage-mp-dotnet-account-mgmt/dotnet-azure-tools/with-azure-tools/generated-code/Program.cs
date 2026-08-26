using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

namespace StorageAccountManager;

internal static class Program
{
    private static async Task<int> Main()
    {
        using CancellationTokenSource cancellationSource = new();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        try
        {
            Settings settings = Settings.FromEnvironment();
            await ManageStorageAccountAsync(settings, cancellationSource.Token);
            return 0;
        }
        catch (CredentialUnavailableException exception)
        {
            Console.Error.WriteLine($"No credential was available: {exception.Message}");
            return 2;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine($"Azure authentication failed: {exception.Message}");
            return 3;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed ({exception.Status}, {exception.ErrorCode}): {exception.Message}");
            return 4;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The operation was canceled.");
            return 5;
        }
        catch (ArgumentException exception)
        {
            Console.Error.WriteLine($"Configuration error: {exception.Message}");
            return 6;
        }
        catch (InvalidOperationException exception)
        {
            Console.Error.WriteLine($"Operation rejected: {exception.Message}");
            return 7;
        }
    }

    private static async Task ManageStorageAccountAsync(
        Settings settings,
        CancellationToken cancellationToken)
    {
        DefaultAzureCredential credential = new();
        ArmClientOptions clientOptions = new()
        {
            Retry =
            {
                Mode = RetryMode.Exponential,
                MaxRetries = 5,
                Delay = TimeSpan.FromSeconds(1),
                MaxDelay = TimeSpan.FromSeconds(16),
                NetworkTimeout = TimeSpan.FromSeconds(100),
            },
        };

        ArmClient armClient = new(credential, settings.SubscriptionId, clientOptions);
        SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync(cancellationToken);
        ResourceGroupResource resourceGroup =
            await subscription.GetResourceGroupAsync(settings.ResourceGroupName, cancellationToken);
        StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

        if (await storageAccounts.ExistsAsync(
            settings.StorageAccountName,
            cancellationToken: cancellationToken))
        {
            throw new InvalidOperationException(
                $"Storage account '{settings.StorageAccountName}' already exists. " +
                "Choose a new name so this sample cannot modify or delete an existing account.");
        }

        StorageAccountResource? createdAccount = null;

        try
        {
            Console.WriteLine($"Creating storage account '{settings.StorageAccountName}'...");
            StorageAccountCreateOrUpdateContent createContent = new(
                new StorageSku(StorageSkuName.StandardLrs),
                StorageKind.StorageV2,
                AzureLocation.EastUS)
            {
                AllowBlobPublicAccess = false,
                AllowSharedKeyAccess = false,
                EnableHttpsTrafficOnly = true,
                MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2,
            };

            ArmOperation<StorageAccountResource> createOperation =
                await storageAccounts.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    settings.StorageAccountName,
                    createContent,
                    cancellationToken);
            createdAccount = createOperation.Value;

            Console.WriteLine(
                $"Created {createdAccount.Data.Name} with SKU {createdAccount.Data.Sku.Name} " +
                $"in {createdAccount.Data.Location}.");

            Console.WriteLine($"\nStorage accounts in resource group '{settings.ResourceGroupName}':");
            await foreach (StorageAccountResource account in
                storageAccounts.GetAllAsync(cancellationToken: cancellationToken))
            {
                Console.WriteLine($"- {account.Data.Name} ({account.Data.Location})");
            }

            StorageAccountResource accountWithProperties =
                await storageAccounts.GetAsync(
                    settings.StorageAccountName,
                    cancellationToken: cancellationToken);
            StorageAccountData properties = accountWithProperties.Data;

            Console.WriteLine("\nCreated account properties:");
            Console.WriteLine($"  Resource ID: {properties.Id}");
            Console.WriteLine($"  Name:        {properties.Name}");
            Console.WriteLine($"  Location:    {properties.Location}");
            Console.WriteLine($"  Kind:        {properties.Kind}");
            Console.WriteLine($"  SKU:         {properties.Sku.Name}");
            Console.WriteLine($"  State:       {properties.ProvisioningState}");
            Console.WriteLine($"  Blob URI:    {properties.PrimaryEndpoints?.BlobUri}");

            Console.WriteLine("\nEnabling blob versioning...");
            BlobServiceResource blobService = accountWithProperties.GetBlobService();
            BlobServiceData blobServiceProperties = new()
            {
                IsVersioningEnabled = true,
            };

            ArmOperation<BlobServiceResource> updateOperation =
                await blobService.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    blobServiceProperties,
                    cancellationToken);

            Console.WriteLine(
                $"Blob versioning enabled: {updateOperation.Value.Data.IsVersioningEnabled}");
        }
        finally
        {
            if (createdAccount is not null)
            {
                Console.WriteLine($"\nDeleting storage account '{createdAccount.Data.Name}'...");
                using CancellationTokenSource cleanupSource = new(TimeSpan.FromMinutes(10));
                await createdAccount.DeleteAsync(WaitUntil.Completed, cleanupSource.Token);
                Console.WriteLine("Storage account deleted.");
            }
        }
    }

    private sealed record Settings(
        string SubscriptionId,
        string ResourceGroupName,
        string StorageAccountName)
    {
        public static Settings FromEnvironment()
        {
            string subscriptionId = GetRequiredVariable("AZURE_SUBSCRIPTION_ID");
            string resourceGroupName = GetRequiredVariable("AZURE_RESOURCE_GROUP");
            string storageAccountName = GetRequiredVariable("AZURE_STORAGE_ACCOUNT_NAME");

            if (!Guid.TryParse(subscriptionId, out _))
            {
                throw new ArgumentException("AZURE_SUBSCRIPTION_ID must be a valid GUID.");
            }

            if (storageAccountName.Length is < 3 or > 24 ||
                storageAccountName.Any(character =>
                    character is not (>= 'a' and <= 'z') and not (>= '0' and <= '9')))
            {
                throw new ArgumentException(
                    "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters or digits.");
            }

            return new Settings(subscriptionId, resourceGroupName, storageAccountName);
        }

        private static string GetRequiredVariable(string name)
        {
            string? value = Environment.GetEnvironmentVariable(name);
            return string.IsNullOrWhiteSpace(value)
                ? throw new ArgumentException($"Set the {name} environment variable.")
                : value.Trim();
        }
    }
}

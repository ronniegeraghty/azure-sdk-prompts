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
        using var cancellationSource = new CancellationTokenSource();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        try
        {
            string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
            string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
            string storageAccountName = GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

            ValidateStorageAccountName(storageAccountName);

            await ManageStorageAccountAsync(
                subscriptionId,
                resourceGroupName,
                storageAccountName,
                cancellationSource.Token);

            return 0;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The operation was canceled.");
            return 2;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            return 3;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure request failed. Status: {ex.Status}; ErrorCode: {ex.ErrorCode}; Message: {ex.Message}");
            return 4;
        }
        catch (ArgumentException ex)
        {
            Console.Error.WriteLine($"Configuration error: {ex.Message}");
            return 5;
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Unexpected error: {ex}");
            return 1;
        }
    }

    private static async Task ManageStorageAccountAsync(
        string subscriptionId,
        string resourceGroupName,
        string storageAccountName,
        CancellationToken cancellationToken)
    {
        TokenCredential credential = new DefaultAzureCredential();
        ArmClient armClient = new(credential);

        SubscriptionResource subscription = armClient.GetSubscriptionResource(
            SubscriptionResource.CreateResourceIdentifier(subscriptionId));

        ResourceGroupResource resourceGroup =
            (await subscription.GetResourceGroupAsync(resourceGroupName, cancellationToken)).Value;

        StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();
        StorageAccountResource? createdAccount = null;
        Exception? operationError = null;

        try
        {
            var createContent = new StorageAccountCreateOrUpdateContent(
                new StorageSku(StorageSkuName.StandardLrs),
                StorageKind.StorageV2,
                AzureLocation.EastUS);

            Console.WriteLine($"Creating storage account '{storageAccountName}'...");
            ArmOperation<StorageAccountResource> createOperation =
                await storageAccounts.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    storageAccountName,
                    createContent,
                    cancellationToken);

            createdAccount = createOperation.Value;
            Console.WriteLine($"Created: {createdAccount.Id}");

            Console.WriteLine($"\nStorage accounts in resource group '{resourceGroupName}':");
            await foreach (StorageAccountResource account in
                storageAccounts.GetAllAsync(cancellationToken: cancellationToken))
            {
                Console.WriteLine($"- {account.Data.Name} ({account.Data.Location})");
            }

            StorageAccountResource accountWithProperties =
                (await createdAccount.GetAsync(cancellationToken: cancellationToken)).Value;

            Console.WriteLine("\nCreated storage account properties:");
            Console.WriteLine($"Name:               {accountWithProperties.Data.Name}");
            Console.WriteLine($"Location:           {accountWithProperties.Data.Location}");
            Console.WriteLine($"SKU:                {accountWithProperties.Data.Sku.Name}");
            Console.WriteLine($"Kind:               {accountWithProperties.Data.Kind}");
            Console.WriteLine($"Provisioning state: {accountWithProperties.Data.ProvisioningState}");
            Console.WriteLine($"Blob endpoint:      {accountWithProperties.Data.PrimaryEndpoints?.BlobUri}");

            BlobServiceResource blobService = accountWithProperties.GetBlobService();
            BlobServiceData blobServiceData = (await blobService.GetAsync(cancellationToken)).Value.Data;
            blobServiceData.IsVersioningEnabled = true;

            Console.WriteLine("\nEnabling blob versioning...");
            await blobService.CreateOrUpdateAsync(
                WaitUntil.Completed,
                blobServiceData,
                cancellationToken);
            Console.WriteLine("Blob versioning enabled.");
        }
        catch (Exception ex)
        {
            operationError = ex;
            throw;
        }
        finally
        {
            if (createdAccount is not null)
            {
                try
                {
                    Console.WriteLine($"\nDeleting storage account '{storageAccountName}'...");
                    await createdAccount.DeleteAsync(WaitUntil.Completed, cancellationToken);
                    Console.WriteLine("Storage account deleted.");
                }
                catch (Exception cleanupError) when (operationError is not null)
                {
                    Console.Error.WriteLine(
                        $"Cleanup also failed while handling another error: {cleanupError.Message}");
                }
            }
        }
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        return string.IsNullOrWhiteSpace(value)
            ? throw new ArgumentException($"Environment variable '{name}' is required.")
            : value;
    }

    private static void ValidateStorageAccountName(string name)
    {
        bool isValid = name.Length is >= 3 and <= 24 &&
                       name.All(character => character is >= 'a' and <= 'z' or >= '0' and <= '9');

        if (!isValid)
        {
            throw new ArgumentException(
                "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters or digits.");
        }
    }
}

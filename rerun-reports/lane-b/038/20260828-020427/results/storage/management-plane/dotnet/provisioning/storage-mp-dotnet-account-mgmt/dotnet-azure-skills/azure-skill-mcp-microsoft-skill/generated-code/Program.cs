using System.Text.RegularExpressions;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

internal static partial class Program
{
    private const string Location = "eastus";

    public static async Task<int> Main()
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
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            return 1;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure request failed (status {ex.Status}, code {ex.ErrorCode ?? "unknown"}): {ex.Message}");
            return 1;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Operation canceled.");
            return 2;
        }
        catch (ArgumentException ex)
        {
            Console.Error.WriteLine($"Configuration error: {ex.Message}");
            return 2;
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
        var credential = new DefaultAzureCredential();
        var armClient = new ArmClient(credential);
        SubscriptionResource subscription = armClient.GetSubscriptionResource(
            SubscriptionResource.CreateResourceIdentifier(subscriptionId));
        ResourceGroupResource resourceGroup =
            (await subscription.GetResourceGroupAsync(resourceGroupName, cancellationToken)).Value;
        StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

        if (await storageAccounts.ExistsAsync(
            storageAccountName,
            cancellationToken: cancellationToken))
        {
            throw new InvalidOperationException(
                $"Storage account '{storageAccountName}' already exists in resource group " +
                $"'{resourceGroupName}'. Choose a new name so this sample cannot delete an existing account.");
        }

        StorageAccountResource? createdAccount = null;
        bool deleted = false;

        try
        {
            Console.WriteLine($"Creating storage account '{storageAccountName}' in {Location}...");
            var createContent = new StorageAccountCreateOrUpdateContent(
                new StorageSku(StorageSkuName.StandardLrs),
                StorageKind.StorageV2,
                new AzureLocation(Location));

            ArmOperation<StorageAccountResource> createOperation =
                await storageAccounts.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    storageAccountName,
                    createContent,
                    cancellationToken);
            createdAccount = createOperation.Value;

            Console.WriteLine($"\nStorage accounts in resource group '{resourceGroupName}':");
            await foreach (StorageAccountResource account in
                storageAccounts.GetAllAsync(cancellationToken: cancellationToken))
            {
                Console.WriteLine($"- {account.Data.Name} ({account.Data.Location})");
            }

            StorageAccountResource accountWithProperties =
                (await createdAccount.GetAsync(cancellationToken: cancellationToken)).Value;
            Console.WriteLine("\nCreated account properties:");
            Console.WriteLine($"  Resource ID: {accountWithProperties.Id}");
            Console.WriteLine($"  Name:        {accountWithProperties.Data.Name}");
            Console.WriteLine($"  Location:    {accountWithProperties.Data.Location}");
            Console.WriteLine($"  SKU:         {accountWithProperties.Data.Sku.Name}");
            Console.WriteLine($"  Kind:        {accountWithProperties.Data.Kind}");

            Console.WriteLine("\nEnabling blob versioning...");
            BlobServiceResource blobService = createdAccount.GetBlobService();
            BlobServiceResource currentBlobService =
                (await blobService.GetAsync(cancellationToken)).Value;
            BlobServiceData blobServiceProperties = currentBlobService.Data;
            blobServiceProperties.IsVersioningEnabled = true;

            BlobServiceResource updatedBlobService =
                (await currentBlobService.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    blobServiceProperties,
                    cancellationToken)).Value;
            Console.WriteLine(
                $"Blob versioning enabled: {updatedBlobService.Data.IsVersioningEnabled}");

            Console.WriteLine($"\nDeleting storage account '{storageAccountName}'...");
            await createdAccount.DeleteAsync(WaitUntil.Completed, cancellationToken);
            deleted = true;
            Console.WriteLine("Storage account deleted.");
        }
        finally
        {
            if (createdAccount is not null && !deleted)
            {
                try
                {
                    Console.Error.WriteLine(
                        $"Cleaning up storage account '{storageAccountName}' after an earlier failure...");
                    await createdAccount.DeleteAsync(WaitUntil.Completed, CancellationToken.None);
                }
                catch (RequestFailedException cleanupException)
                {
                    Console.Error.WriteLine(
                        $"Cleanup failed (status {cleanupException.Status}, " +
                        $"code {cleanupException.ErrorCode ?? "unknown"}): {cleanupException.Message}");
                }
            }
        }
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        return string.IsNullOrWhiteSpace(value)
            ? throw new ArgumentException($"Environment variable {name} is required.")
            : value;
    }

    private static void ValidateStorageAccountName(string name)
    {
        if (!StorageAccountNamePattern().IsMatch(name))
        {
            throw new ArgumentException(
                "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters and digits.");
        }
    }

    [GeneratedRegex("^[a-z0-9]{3,24}$", RegexOptions.CultureInvariant)]
    private static partial Regex StorageAccountNamePattern();
}

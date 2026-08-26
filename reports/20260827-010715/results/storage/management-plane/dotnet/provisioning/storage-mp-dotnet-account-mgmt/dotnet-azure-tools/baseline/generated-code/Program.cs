using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

const string location = "eastus";

string? subscriptionId = Environment.GetEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string? resourceGroupName = Environment.GetEnvironmentVariable("AZURE_RESOURCE_GROUP");
string? storageAccountName = Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

if (string.IsNullOrWhiteSpace(subscriptionId) ||
    string.IsNullOrWhiteSpace(resourceGroupName) ||
    string.IsNullOrWhiteSpace(storageAccountName))
{
    Console.Error.WriteLine(
        "Set AZURE_SUBSCRIPTION_ID, AZURE_RESOURCE_GROUP, and " +
        "AZURE_STORAGE_ACCOUNT_NAME before running this program.");
    return 2;
}

StorageAccountResource? createdAccount = null;
int exitCode;

try
{
    TokenCredential credential = new DefaultAzureCredential();
    ArmClient armClient = new(credential, subscriptionId);

    ResourceIdentifier resourceGroupId =
        ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
    ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);

    Console.WriteLine($"Creating storage account '{storageAccountName}'...");
    StorageAccountCollection accounts = resourceGroup.GetStorageAccounts();
    StorageAccountCreateOrUpdateContent createContent = new(
        new StorageSku(StorageSkuName.StandardLrs),
        StorageKind.StorageV2,
        location);

    ArmOperation<StorageAccountResource> createOperation =
        await accounts.CreateOrUpdateAsync(
            WaitUntil.Completed,
            storageAccountName,
            createContent);
    createdAccount = createOperation.Value;
    Console.WriteLine($"Created: {createdAccount.Id}");

    Console.WriteLine($"\nStorage accounts in resource group '{resourceGroupName}':");
    await foreach (StorageAccountResource account in accounts.GetAllAsync())
    {
        Console.WriteLine($"- {account.Data.Name} ({account.Data.Location})");
    }

    Response<StorageAccountResource> getResponse =
        await accounts.GetAsync(storageAccountName);
    StorageAccountData properties = getResponse.Value.Data;
    Console.WriteLine(
        $"\nProperties for '{properties.Name}': " +
        $"location={properties.Location}, kind={properties.Kind}, sku={properties.Sku.Name}");

    Console.WriteLine("\nEnabling blob versioning...");
    BlobServiceResource blobService =
        (await createdAccount.GetBlobService().GetAsync()).Value;
    blobService.Data.IsVersioningEnabled = true;

    ArmOperation<BlobServiceResource> updateOperation =
        await blobService.CreateOrUpdateAsync(WaitUntil.Completed, blobService.Data);
    Console.WriteLine(
        $"Blob versioning enabled: {updateOperation.Value.Data.IsVersioningEnabled}");

    exitCode = 0;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Azure authentication failed: {ex.Message}");
    exitCode = 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed ({ex.Status}, {ex.ErrorCode ?? "no error code"}): {ex.Message}");
    exitCode = 4;
}
catch (Exception ex)
{
    Console.Error.WriteLine($"Unexpected error: {ex.Message}");
    exitCode = 5;
}
finally
{
    if (createdAccount is not null)
    {
        try
        {
            Console.WriteLine($"\nDeleting storage account '{storageAccountName}'...");
            await createdAccount.DeleteAsync(WaitUntil.Completed);
            Console.WriteLine("Storage account deleted.");
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Cleanup failed ({ex.Status}, {ex.ErrorCode ?? "no error code"}): {ex.Message}");
            exitCode = 6;
        }
    }
}

return exitCode;

using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

const string executeArgument = "--execute";

if (!args.Contains(executeArgument, StringComparer.OrdinalIgnoreCase))
{
    Console.WriteLine("Dry run: no Azure operations were performed.");
    Console.WriteLine(
        "Set AZURE_SUBSCRIPTION_ID, AZURE_RESOURCE_GROUP, and optionally " +
        "AZURE_STORAGE_ACCOUNT_NAME, then pass --execute to run the sample.");
    return;
}

string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string storageAccountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME")
    ?? $"st{Guid.NewGuid():N}"[..24];

ArmClient? armClient = null;
StorageAccountResource? createdAccount = null;
bool deleted = false;

try
{
    // DefaultAzureCredential supports local developer credentials and managed identity.
    armClient = new ArmClient(
        new DefaultAzureCredential(),
        subscriptionId);

    ResourceIdentifier resourceGroupId = ResourceGroupResource.CreateResourceIdentifier(
        subscriptionId,
        resourceGroupName);
    ResourceGroupResource resourceGroup =
        armClient.GetResourceGroupResource(resourceGroupId);
    StorageAccountCollection accounts = resourceGroup.GetStorageAccounts();

    var createContent = new StorageAccountCreateOrUpdateContent(
        new StorageSku(StorageSkuName.StandardLrs),
        StorageKind.StorageV2,
        AzureLocation.EastUS);

    Console.WriteLine($"Creating storage account '{storageAccountName}'...");
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
        Console.WriteLine(
            $"- {account.Data.Name} | {account.Data.Location} | {account.Data.Sku.Name}");
    }

    Response<StorageAccountResource> getResponse = await createdAccount.GetAsync();
    StorageAccountData accountData = getResponse.Value.Data;
    Console.WriteLine("\nCreated storage account properties:");
    Console.WriteLine($"Name: {accountData.Name}");
    Console.WriteLine($"Location: {accountData.Location}");
    Console.WriteLine($"Kind: {accountData.Kind}");
    Console.WriteLine($"SKU: {accountData.Sku.Name}");
    Console.WriteLine($"Provisioning state: {accountData.ProvisioningState}");
    Console.WriteLine($"Primary blob endpoint: {accountData.PrimaryEndpoints?.BlobUri}");

    BlobServiceResource blobService = createdAccount.GetBlobService();
    Response<BlobServiceResource> blobServiceResponse = await blobService.GetAsync();
    BlobServiceData blobServiceData = blobServiceResponse.Value.Data;
    blobServiceData.IsVersioningEnabled = true;

    Console.WriteLine("\nEnabling blob versioning...");
    ArmOperation<BlobServiceResource> updateOperation =
        await blobService.CreateOrUpdateAsync(WaitUntil.Completed, blobServiceData);
    Console.WriteLine(
        $"Blob versioning enabled: {updateOperation.Value.Data.IsVersioningEnabled}");

    Console.WriteLine($"\nDeleting storage account '{storageAccountName}'...");
    await createdAccount.DeleteAsync(WaitUntil.Completed);
    deleted = true;
    Console.WriteLine("Storage account deleted.");
}
catch (CredentialUnavailableException ex)
{
    Console.Error.WriteLine($"No Azure credential is available: {ex.Message}");
    Environment.ExitCode = 1;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Azure authentication failed: {ex.Message}");
    Environment.ExitCode = 1;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed. Status={ex.Status}, ErrorCode={ex.ErrorCode}, " +
        $"Message={ex.Message}");
    Environment.ExitCode = 1;
}
catch (ArgumentException ex)
{
    Console.Error.WriteLine($"Invalid configuration: {ex.Message}");
    Environment.ExitCode = 1;
}
finally
{
    if (createdAccount is not null && !deleted)
    {
        try
        {
            Console.Error.WriteLine(
                $"Cleaning up storage account '{storageAccountName}' after failure...");
            await createdAccount.DeleteAsync(WaitUntil.Completed);
            Console.Error.WriteLine("Cleanup completed.");
        }
        catch (RequestFailedException cleanupException)
        {
            Console.Error.WriteLine(
                $"Cleanup failed. Delete '{storageAccountName}' manually. " +
                $"Status={cleanupException.Status}, ErrorCode={cleanupException.ErrorCode}, " +
                $"Message={cleanupException.Message}");
            Environment.ExitCode = 1;
        }
    }
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new ArgumentException($"Environment variable '{name}' is required.");
}

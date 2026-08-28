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

        StorageAccountResource? createdAccount = null;

        try
        {
            string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
            string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
            string accountName = GetStorageAccountName();

            var credential = new DefaultAzureCredential();
            var armClient = new ArmClient(credential, subscriptionId);

            ResourceIdentifier resourceGroupId =
                ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
            ResourceGroupResource resourceGroup =
                (await armClient.GetResourceGroupResource(resourceGroupId)
                    .GetAsync(cancellationSource.Token)).Value;
            StorageAccountCollection accounts = resourceGroup.GetStorageAccounts();

            if (await accounts.ExistsAsync(
                accountName,
                cancellationToken: cancellationSource.Token))
            {
                throw new InvalidOperationException(
                    $"Storage account '{accountName}' already exists. Choose a new name so this sample cannot delete an existing account.");
            }

            var createContent = new StorageAccountCreateOrUpdateContent(
                new StorageSku(StorageSkuName.StandardLrs),
                StorageKind.StorageV2,
                AzureLocation.EastUS)
            {
                AllowBlobPublicAccess = false,
                MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2
            };

            Console.WriteLine($"Creating storage account '{accountName}' in eastus...");
            ArmOperation<StorageAccountResource> createOperation =
                await accounts.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    accountName,
                    createContent,
                    cancellationSource.Token);
            createdAccount = createOperation.Value;

            Console.WriteLine($"\nStorage accounts in resource group '{resourceGroupName}':");
            await foreach (StorageAccountResource account in
                accounts.GetAllAsync(cancellationToken: cancellationSource.Token))
            {
                Console.WriteLine($"- {account.Data.Name} ({account.Data.Location})");
            }

            StorageAccountResource currentAccount =
                (await createdAccount.GetAsync(
                    cancellationToken: cancellationSource.Token)).Value;
            StorageAccountData properties = currentAccount.Data;

            Console.WriteLine("\nCreated storage account properties:");
            Console.WriteLine($"  Resource ID:        {properties.Id}");
            Console.WriteLine($"  Location:           {properties.Location}");
            Console.WriteLine($"  SKU:                {properties.Sku.Name}");
            Console.WriteLine($"  Kind:               {properties.Kind}");
            Console.WriteLine($"  Provisioning state: {properties.ProvisioningState}");
            Console.WriteLine($"  Blob endpoint:      {properties.PrimaryEndpoints.BlobUri}");

            BlobServiceResource blobService = currentAccount.GetBlobService();
            BlobServiceData blobServiceData;

            try
            {
                blobServiceData =
                    (await blobService.GetAsync(cancellationSource.Token)).Value.Data;
            }
            catch (RequestFailedException ex) when (ex.Status == 404)
            {
                blobServiceData = new BlobServiceData();
            }

            blobServiceData.IsVersioningEnabled = true;
            await blobService.CreateOrUpdateAsync(
                WaitUntil.Completed,
                blobServiceData,
                cancellationSource.Token);
            Console.WriteLine("\nBlob versioning enabled.");

            await createdAccount.DeleteAsync(
                WaitUntil.Completed,
                cancellationSource.Token);
            createdAccount = null;
            Console.WriteLine("Storage account deleted.");

            return 0;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            return 2;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure request failed (HTTP {ex.Status}, code {ex.ErrorCode ?? "unknown"}): {ex.Message}");
            return 3;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Operation canceled.");
            return 4;
        }
        catch (InvalidOperationException ex)
        {
            Console.Error.WriteLine($"Configuration error: {ex.Message}");
            return 5;
        }
        finally
        {
            if (createdAccount is not null)
            {
                try
                {
                    Console.Error.WriteLine(
                        $"Cleaning up storage account '{createdAccount.Data.Name}'...");
                    await createdAccount.DeleteAsync(
                        WaitUntil.Completed,
                        CancellationToken.None);
                }
                catch (RequestFailedException ex)
                {
                    Console.Error.WriteLine(
                        $"Cleanup failed (HTTP {ex.Status}, code {ex.ErrorCode ?? "unknown"}): {ex.Message}");
                }
            }
        }
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        return string.IsNullOrWhiteSpace(value)
            ? throw new InvalidOperationException(
                $"Set the {name} environment variable before running the program.")
            : value;
    }

    private static string GetStorageAccountName()
    {
        string? configuredName =
            Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

        if (string.IsNullOrWhiteSpace(configuredName))
        {
            return $"stmgmt{Guid.NewGuid():N}"[..24];
        }

        if (configuredName.Length is < 3 or > 24 ||
            configuredName.Any(character =>
                character is not (>= 'a' and <= 'z') &&
                character is not (>= '0' and <= '9')))
        {
            throw new InvalidOperationException(
                "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters or digits.");
        }

        return configuredName;
    }
}

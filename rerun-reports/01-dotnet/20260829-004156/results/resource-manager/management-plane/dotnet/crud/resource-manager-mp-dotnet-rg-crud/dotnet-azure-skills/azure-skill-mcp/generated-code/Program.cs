using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

namespace AzureResourceGroupManager;

internal static class Program
{
    private static async Task<int> Main()
    {
        ResourceGroupResource? createdResourceGroup = null;

        try
        {
            DefaultAzureCredential credential = new();
            ArmClient armClient = new(credential);
            SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();
            ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();

            string resourceGroupName =
                $"sdk-rg-{DateTimeOffset.UtcNow:yyyyMMddHHmmss}-{Guid.NewGuid():N}"[..39];

            Console.WriteLine(
                $"Creating resource group '{resourceGroupName}' in '{AzureLocation.EastUS}'...");

            ResourceGroupData resourceGroupData = new(AzureLocation.EastUS);
            ArmOperation<ResourceGroupResource> createOperation =
                await resourceGroups.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    resourceGroupName,
                    resourceGroupData);

            createdResourceGroup = createOperation.Value;
            Console.WriteLine($"Created: {createdResourceGroup.Id}");

            Console.WriteLine("\nResource groups in the subscription:");
            await foreach (ResourceGroupResource resourceGroup in resourceGroups.GetAllAsync())
            {
                Console.WriteLine(
                    $"- {resourceGroup.Data.Name} ({resourceGroup.Data.Location})");
            }

            ResourceGroupResource resourceGroupDetails =
                await resourceGroups.GetAsync(resourceGroupName);

            Console.WriteLine("\nCreated resource group details:");
            Console.WriteLine($"Name: {resourceGroupDetails.Data.Name}");
            Console.WriteLine($"Location: {resourceGroupDetails.Data.Location}");
            Console.WriteLine($"Resource ID: {resourceGroupDetails.Id}");
            Console.WriteLine($"Tag count: {resourceGroupDetails.Data.Tags.Count}");

            const string tagName = "ManagedBy";
            const string tagValue = "Azure.ResourceManager";
            createdResourceGroup =
                await resourceGroupDetails.AddTagAsync(tagName, tagValue);

            Console.WriteLine($"\nAdded tag: {tagName}={createdResourceGroup.Data.Tags[tagName]}");

            Console.WriteLine($"\nDeleting resource group '{resourceGroupName}'...");
            await createdResourceGroup.DeleteAsync(WaitUntil.Completed);
            createdResourceGroup = null;
            Console.WriteLine("Resource group deleted.");

            return 0;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine(
                $"Authentication failed. Configure a credential supported by " +
                $"DefaultAzureCredential. Details: {exception.Message}");
            return 1;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed (HTTP {exception.Status}, " +
                $"error code '{exception.ErrorCode ?? "unknown"}'): {exception.Message}");
            return 1;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The operation was canceled.");
            return 1;
        }
        finally
        {
            if (createdResourceGroup is not null)
            {
                try
                {
                    Console.Error.WriteLine(
                        $"\nCleaning up resource group '{createdResourceGroup.Data.Name}'...");
                    await createdResourceGroup.DeleteAsync(WaitUntil.Completed);
                    Console.Error.WriteLine("Cleanup completed.");
                }
                catch (RequestFailedException exception)
                {
                    Console.Error.WriteLine(
                        $"Cleanup failed (HTTP {exception.Status}, " +
                        $"error code '{exception.ErrorCode ?? "unknown"}'): {exception.Message}");
                }
            }
        }
    }
}

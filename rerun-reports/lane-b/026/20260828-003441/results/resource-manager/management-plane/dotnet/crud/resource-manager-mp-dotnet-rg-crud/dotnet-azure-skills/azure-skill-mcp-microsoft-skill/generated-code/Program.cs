using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

internal static class Program
{
    private static async Task<int> Main(string[] args)
    {
        string resourceGroupName = args.Length > 0
            ? args[0]
            : $"rg-sdk-demo-{DateTime.UtcNow:yyyyMMddHHmmss}";

        ResourceGroupResource? createdResourceGroup = null;
        bool deleted = false;

        try
        {
            var credential = new DefaultAzureCredential();
            var armClient = new ArmClient(credential);

            SubscriptionResource subscription =
                await armClient.GetDefaultSubscriptionAsync();
            ResourceGroupCollection resourceGroups =
                subscription.GetResourceGroups();

            Console.WriteLine(
                $"Creating resource group '{resourceGroupName}' in eastus...");

            var resourceGroupData =
                new ResourceGroupData(AzureLocation.EastUS);
            ArmOperation<ResourceGroupResource> createOperation =
                await resourceGroups.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    resourceGroupName,
                    resourceGroupData);

            createdResourceGroup = createOperation.Value;
            Console.WriteLine(
                $"Created: {createdResourceGroup.Data.Id}");

            Console.WriteLine("\nResource groups in the subscription:");
            await foreach (ResourceGroupResource resourceGroup
                in resourceGroups.GetAllAsync())
            {
                Console.WriteLine(
                    $"- {resourceGroup.Data.Name} " +
                    $"({resourceGroup.Data.Location})");
            }

            Response<ResourceGroupResource> getResponse =
                await resourceGroups.GetAsync(resourceGroupName);
            ResourceGroupResource resourceGroupDetails = getResponse.Value;

            Console.WriteLine("\nCreated resource group details:");
            Console.WriteLine($"  Name:     {resourceGroupDetails.Data.Name}");
            Console.WriteLine($"  Location: {resourceGroupDetails.Data.Location}");
            Console.WriteLine($"  ID:       {resourceGroupDetails.Data.Id}");

            Response<ResourceGroupResource> tagResponse =
                await resourceGroupDetails.AddTagAsync(
                    "managed-by",
                    "Azure.ResourceManager");
            createdResourceGroup = tagResponse.Value;

            Console.WriteLine(
                "\nAdded tag: managed-by=Azure.ResourceManager");

            Console.WriteLine(
                $"\nDeleting resource group '{resourceGroupName}'...");
            await createdResourceGroup.DeleteAsync(WaitUntil.Completed);
            deleted = true;

            Console.WriteLine("Resource group deleted.");
            return 0;
        }
        catch (CredentialUnavailableException exception)
        {
            Console.Error.WriteLine(
                $"No credential is available: {exception.Message}");
            return 1;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine(
                $"Authentication failed: {exception.Message}");
            return 1;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed. Status={exception.Status}, " +
                $"ErrorCode={exception.ErrorCode}, Message={exception.Message}");
            return 1;
        }
        finally
        {
            if (createdResourceGroup is not null && !deleted)
            {
                try
                {
                    Console.Error.WriteLine(
                        $"Cleaning up resource group '{resourceGroupName}'...");
                    await createdResourceGroup.DeleteAsync(WaitUntil.Completed);
                }
                catch (RequestFailedException cleanupException)
                {
                    Console.Error.WriteLine(
                        $"Cleanup failed. Status={cleanupException.Status}, " +
                        $"ErrorCode={cleanupException.ErrorCode}, " +
                        $"Message={cleanupException.Message}");
                }
            }
        }
    }
}

using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

namespace AzureResourceGroupManager;

internal static class Program
{
    public static async Task<int> Main(string[] args)
    {
        string resourceGroupName = args.Length > 0
            ? args[0]
            : $"rg-sdk-demo-{DateTime.UtcNow:yyyyMMddHHmmss}";

        using var cancellationSource = new CancellationTokenSource();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        try
        {
            await ManageResourceGroupAsync(resourceGroupName, cancellationSource.Token);
            return 0;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine($"Azure authentication failed: {exception.Message}");
            Console.Error.WriteLine(
                "Configure a credential supported by DefaultAzureCredential and try again.");
            return 1;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed (HTTP {exception.Status}, code {exception.ErrorCode ?? "unknown"}): " +
                exception.Message);
            return 2;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The operation was canceled.");
            return 3;
        }
    }

    private static async Task ManageResourceGroupAsync(
        string resourceGroupName,
        CancellationToken cancellationToken)
    {
        var credential = new DefaultAzureCredential();
        var armClient = new ArmClient(credential);

        SubscriptionResource subscription =
            await armClient.GetDefaultSubscriptionAsync(cancellationToken);
        ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();

        ResourceGroupResource? createdResourceGroup = null;

        try
        {
            Console.WriteLine(
                $"Creating resource group '{resourceGroupName}' in '{AzureLocation.EastUS}'...");

            var resourceGroupData = new ResourceGroupData(AzureLocation.EastUS);
            ArmOperation<ResourceGroupResource> createOperation =
                await resourceGroups.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    resourceGroupName,
                    resourceGroupData,
                    cancellationToken);

            createdResourceGroup = createOperation.Value;
            Console.WriteLine($"Created: {createdResourceGroup.Id}");

            Console.WriteLine("\nResource groups in the subscription:");
            await foreach (ResourceGroupResource resourceGroup in
                resourceGroups.GetAllAsync(cancellationToken: cancellationToken))
            {
                Console.WriteLine(
                    $"- {resourceGroup.Data.Name} ({resourceGroup.Data.Location})");
            }

            Response<ResourceGroupResource> getResponse =
                await resourceGroups.GetAsync(resourceGroupName, cancellationToken);
            ResourceGroupResource resourceGroupDetails = getResponse.Value;

            Console.WriteLine("\nCreated resource group details:");
            Console.WriteLine($"  Name:     {resourceGroupDetails.Data.Name}");
            Console.WriteLine($"  ID:       {resourceGroupDetails.Id}");
            Console.WriteLine($"  Location: {resourceGroupDetails.Data.Location}");
            Console.WriteLine($"  Type:     {resourceGroupDetails.Data.ResourceType}");

            var tags = new Dictionary<string, string>(resourceGroupDetails.Data.Tags)
            {
                ["managed-by"] = "Azure.ResourceManager"
            };

            Response<ResourceGroupResource> tagResponse =
                await resourceGroupDetails.SetTagsAsync(tags, cancellationToken);

            createdResourceGroup = tagResponse.Value;
            Console.WriteLine("\nAdded tag: managed-by=Azure.ResourceManager");
        }
        finally
        {
            if (createdResourceGroup is not null)
            {
                Console.WriteLine($"\nDeleting resource group '{resourceGroupName}'...");
                await createdResourceGroup.DeleteAsync(
                    WaitUntil.Completed,
                    forceDeletionTypes: null,
                    cancellationToken: CancellationToken.None);
                Console.WriteLine("Resource group deleted.");
            }
        }
    }
}

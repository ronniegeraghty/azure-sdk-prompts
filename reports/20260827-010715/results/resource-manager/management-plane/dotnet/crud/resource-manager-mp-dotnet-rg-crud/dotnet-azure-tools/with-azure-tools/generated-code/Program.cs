using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

namespace hyoka_resource_manager_mp_dotnet_rg_crud_dotnet_azure_tools_with_azure_tools_121524720;

internal static class Program
{
    private static async Task<int> Main(string[] args)
    {
        string resourceGroupName = args.Length > 0
            ? args[0]
            : $"resource-manager-demo-{Guid.NewGuid():N}"[..37];

        using CancellationTokenSource cancellationTokenSource = new();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationTokenSource.Cancel();
        };

        ResourceGroupResource? resourceGroup = null;
        bool deleted = false;
        int exitCode = 0;

        try
        {
            TokenCredential credential = new DefaultAzureCredential();
            ArmClient armClient = new(credential);
            SubscriptionResource subscription =
                await armClient.GetDefaultSubscriptionAsync(cancellationTokenSource.Token);
            ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();

            Console.WriteLine(
                $"Creating resource group '{resourceGroupName}' in '{AzureLocation.EastUS}'...");
            ResourceGroupData resourceGroupData = new(AzureLocation.EastUS);
            ArmOperation<ResourceGroupResource> createOperation =
                await resourceGroups.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    resourceGroupName,
                    resourceGroupData,
                    cancellationTokenSource.Token);
            resourceGroup = createOperation.Value;
            Console.WriteLine($"Created: {resourceGroup.Id}");

            Console.WriteLine("\nResource groups in the subscription:");
            await foreach (ResourceGroupResource item in
                resourceGroups.GetAllAsync().WithCancellation(cancellationTokenSource.Token))
            {
                Console.WriteLine($"- {item.Data.Name} ({item.Data.Location})");
            }

            Console.WriteLine($"\nGetting details for '{resourceGroupName}'...");
            resourceGroup = (await resourceGroups.GetAsync(
                resourceGroupName,
                cancellationTokenSource.Token)).Value;
            PrintDetails(resourceGroup.Data);

            const string tagName = "managed-by";
            const string tagValue = "azure-resource-manager-sdk";
            Console.WriteLine($"\nAdding tag '{tagName}={tagValue}'...");
            resourceGroup = (await resourceGroup.AddTagAsync(
                tagName,
                tagValue,
                cancellationTokenSource.Token)).Value;
            PrintDetails(resourceGroup.Data);

            Console.WriteLine($"\nDeleting resource group '{resourceGroupName}'...");
            await resourceGroup.DeleteAsync(
                WaitUntil.Completed,
                cancellationToken: cancellationTokenSource.Token);
            deleted = true;
            Console.WriteLine("Resource group deleted.");
        }
        catch (CredentialUnavailableException exception)
        {
            Console.Error.WriteLine(
                $"No credential was available for DefaultAzureCredential: {exception.Message}");
            exitCode = 1;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine($"Azure authentication failed: {exception.Message}");
            exitCode = 1;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed (HTTP {exception.Status}, code '{exception.ErrorCode}'): " +
                exception.Message);
            exitCode = 1;
        }
        catch (OperationCanceledException) when (cancellationTokenSource.IsCancellationRequested)
        {
            Console.Error.WriteLine("Operation canceled.");
            exitCode = 2;
        }
        finally
        {
            if (resourceGroup is not null && !deleted)
            {
                try
                {
                    Console.WriteLine(
                        $"\nCleaning up resource group '{resourceGroupName}' after the failure...");
                    await resourceGroup.DeleteAsync(
                        WaitUntil.Completed,
                        cancellationToken: CancellationToken.None);
                    Console.WriteLine("Cleanup completed.");
                }
                catch (RequestFailedException cleanupException)
                {
                    Console.Error.WriteLine(
                        $"Cleanup failed (HTTP {cleanupException.Status}, " +
                        $"code '{cleanupException.ErrorCode}'): {cleanupException.Message}");
                    exitCode = 1;
                }
            }
        }

        return exitCode;
    }

    private static void PrintDetails(ResourceGroupData data)
    {
        Console.WriteLine($"Name: {data.Name}");
        Console.WriteLine($"Location: {data.Location}");
        Console.WriteLine($"Resource ID: {data.Id}");
        Console.WriteLine(
            $"Tags: {(data.Tags.Count == 0 ? "(none)" : string.Join(", ", data.Tags))}");
    }
}

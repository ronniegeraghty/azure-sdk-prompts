using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

namespace ResourceGroupManager;

internal static class Program
{
    private const string Location = "eastus";
    private const string TagName = "managed-by";
    private const string TagValue = "azure-resource-manager-sdk";

    public static async Task<int> Main(string[] args)
    {
        using CancellationTokenSource cancellationSource = new();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        string resourceGroupName = args.Length > 0
            ? args[0]
            : $"rg-sdk-sample-{DateTimeOffset.UtcNow:yyyyMMddHHmmss}";

        ResourceGroupResource? createdResourceGroup = null;
        bool deleted = false;

        try
        {
            DefaultAzureCredential credential = new();
            ArmClient armClient = new(credential);
            SubscriptionResource subscription =
                await armClient.GetDefaultSubscriptionAsync(cancellationSource.Token);
            ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();

            Console.WriteLine(
                $"Creating resource group '{resourceGroupName}' in '{Location}'...");

            ArmOperation<ResourceGroupResource> createOperation =
                await resourceGroups.CreateOrUpdateAsync(
                    WaitUntil.Completed,
                    resourceGroupName,
                    new ResourceGroupData(new AzureLocation(Location)),
                    cancellationSource.Token);

            createdResourceGroup = createOperation.Value;
            Console.WriteLine($"Created: {createdResourceGroup.Data.Id}");

            Console.WriteLine("\nResource groups in the subscription:");
            await foreach (ResourceGroupResource resourceGroup in
                resourceGroups.GetAllAsync(cancellationToken: cancellationSource.Token))
            {
                Console.WriteLine(
                    $"- {resourceGroup.Data.Name} ({resourceGroup.Data.Location})");
            }

            ResourceGroupResource details =
                (await resourceGroups.GetAsync(
                    resourceGroupName,
                    cancellationSource.Token)).Value;

            Console.WriteLine("\nCreated resource group details:");
            Console.WriteLine($"  ID:       {details.Data.Id}");
            Console.WriteLine($"  Name:     {details.Data.Name}");
            Console.WriteLine($"  Location: {details.Data.Location}");

            ResourceGroupResource taggedResourceGroup =
                (await details.AddTagAsync(
                    TagName,
                    TagValue,
                    cancellationSource.Token)).Value;

            Console.WriteLine(
                $"\nAdded tag '{TagName}={taggedResourceGroup.Data.Tags[TagName]}'.");

            Console.WriteLine($"\nDeleting resource group '{resourceGroupName}'...");
            await taggedResourceGroup.DeleteAsync(
                WaitUntil.Completed,
                cancellationToken: cancellationSource.Token);
            deleted = true;
            Console.WriteLine("Resource group deleted.");

            return 0;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine(
                $"Authentication failed. Configure a DefaultAzureCredential source: {exception.Message}");
            return 1;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed ({exception.Status}, {exception.ErrorCode}): {exception.Message}");
            return 2;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Operation canceled.");
            return 3;
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
                catch (RequestFailedException exception)
                {
                    Console.Error.WriteLine(
                        $"Cleanup failed ({exception.Status}, {exception.ErrorCode}). " +
                        $"Delete '{resourceGroupName}' manually. {exception.Message}");
                }
            }
        }
    }
}

using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

internal static class Program
{
    private const string SubscriptionIdVariable = "AZURE_SUBSCRIPTION_ID";

    public static async Task<int> Main(string[] args)
    {
        using var cancellationTokenSource = new CancellationTokenSource();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationTokenSource.Cancel();
        };

        try
        {
            string subscriptionId = GetRequiredEnvironmentVariable(SubscriptionIdVariable);
            string resourceGroupName =
                Environment.GetEnvironmentVariable("AZURE_RESOURCE_GROUP_NAME")
                ?? $"rg-sdk-demo-{DateTime.UtcNow:yyyyMMddHHmmss}-{Guid.NewGuid():N}"[..34];

            await ManageResourceGroupAsync(
                subscriptionId,
                resourceGroupName,
                cancellationTokenSource.Token);

            return 0;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine($"Authentication failed: {exception.Message}");
            Console.Error.WriteLine(
                "Sign in with a supported DefaultAzureCredential method, such as Azure CLI, " +
                "Visual Studio, workload identity, or managed identity.");
            return 1;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure request failed ({exception.Status}, {exception.ErrorCode}): {exception.Message}");
            return 1;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Operation canceled.");
            return 2;
        }
        catch (InvalidOperationException exception)
        {
            Console.Error.WriteLine(exception.Message);
            return 2;
        }
        catch (Exception exception)
        {
            Console.Error.WriteLine($"Unexpected error: {exception}");
            return 1;
        }
    }

    private static async Task ManageResourceGroupAsync(
        string subscriptionId,
        string resourceGroupName,
        CancellationToken cancellationToken)
    {
        var credential = new DefaultAzureCredential();
        var armClient = new ArmClient(credential, subscriptionId);
        SubscriptionResource subscription = armClient.GetSubscriptionResource(
            SubscriptionResource.CreateResourceIdentifier(subscriptionId));
        ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();

        Console.WriteLine($"Creating resource group '{resourceGroupName}' in eastus...");
        var resourceGroupData = new ResourceGroupData(AzureLocation.EastUS);
        ArmOperation<ResourceGroupResource> createOperation =
            await resourceGroups.CreateOrUpdateAsync(
                WaitUntil.Completed,
                resourceGroupName,
                resourceGroupData,
                cancellationToken);
        ResourceGroupResource createdResourceGroup = createOperation.Value;

        Console.WriteLine("Resource groups in the subscription:");
        await foreach (ResourceGroupResource resourceGroup in
                       resourceGroups.GetAllAsync(cancellationToken: cancellationToken))
        {
            Console.WriteLine($"- {resourceGroup.Data.Name} ({resourceGroup.Data.Location})");
        }

        Response<ResourceGroupResource> getResponse =
            await resourceGroups.GetAsync(resourceGroupName, cancellationToken);
        ResourceGroupResource fetchedResourceGroup = getResponse.Value;
        Console.WriteLine(
            $"Created resource group: Name={fetchedResourceGroup.Data.Name}, " +
            $"Location={fetchedResourceGroup.Data.Location}, Id={fetchedResourceGroup.Id}");

        const string tagName = "ManagedBy";
        const string tagValue = "Azure.ResourceManager";
        var patch = new ResourceGroupPatch();
        patch.Tags[tagName] = tagValue;

        Response<ResourceGroupResource> updateResponse =
            await createdResourceGroup.UpdateAsync(
                patch,
                cancellationToken);
        Console.WriteLine(
            $"Added tag '{tagName}={updateResponse.Value.Data.Tags[tagName]}'.");

        Console.WriteLine($"Deleting resource group '{resourceGroupName}'...");
        await updateResponse.Value.DeleteAsync(
            WaitUntil.Completed,
            cancellationToken: cancellationToken);
        Console.WriteLine("Resource group deleted.");
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        if (string.IsNullOrWhiteSpace(value))
        {
            throw new InvalidOperationException(
                $"Set the {name} environment variable to the target Azure subscription ID.");
        }

        return value;
    }
}

using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

const string location = "eastus";
const string tagName = "managed-by";
const string tagValue = "azure-resource-manager-sdk";

string? subscriptionId = Environment.GetEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
if (string.IsNullOrWhiteSpace(subscriptionId))
{
    Console.Error.WriteLine(
        "Set AZURE_SUBSCRIPTION_ID to the subscription in which the sample should run.");
    return 1;
}

string resourceGroupName =
    Environment.GetEnvironmentVariable("AZURE_RESOURCE_GROUP_NAME")
    ?? $"rg-sdk-sample-{Guid.NewGuid():N}"[..32];

ResourceGroupResource? createdResourceGroup = null;

try
{
    TokenCredential credential = new DefaultAzureCredential();
    ArmClient armClient = new(credential);

    ResourceIdentifier subscriptionResourceId =
        SubscriptionResource.CreateResourceIdentifier(subscriptionId);
    SubscriptionResource subscription =
        armClient.GetSubscriptionResource(subscriptionResourceId);
    ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();

    NullableResponse<ResourceGroupResource> existingResourceGroup =
        await resourceGroups.GetIfExistsAsync(resourceGroupName);
    if (existingResourceGroup.HasValue)
    {
        throw new InvalidOperationException(
            $"Resource group '{resourceGroupName}' already exists. " +
            "Choose a different AZURE_RESOURCE_GROUP_NAME.");
    }

    Console.WriteLine(
        $"Creating resource group '{resourceGroupName}' in '{location}'...");
    ResourceGroupData resourceGroupData = new(location);
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

    Response<ResourceGroupResource> getResponse =
        await resourceGroups.GetAsync(resourceGroupName);
    ResourceGroupResource fetchedResourceGroup = getResponse.Value;
    Console.WriteLine(
        $"\nDetails: name={fetchedResourceGroup.Data.Name}, " +
        $"location={fetchedResourceGroup.Data.Location}, " +
        $"id={fetchedResourceGroup.Data.Id}");

    ResourceGroupPatch patch = new();
    patch.Tags.Add(tagName, tagValue);
    Response<ResourceGroupResource> updateResponse =
        await fetchedResourceGroup.UpdateAsync(patch);
    createdResourceGroup = updateResponse.Value;
    Console.WriteLine($"Added tag: {tagName}={tagValue}");

    Console.WriteLine($"\nDeleting resource group '{resourceGroupName}'...");
    await createdResourceGroup.DeleteAsync(WaitUntil.Completed);
    createdResourceGroup = null;
    Console.WriteLine("Resource group deleted.");

    return 0;
}
catch (CredentialUnavailableException ex)
{
    Console.Error.WriteLine($"No DefaultAzureCredential source was available: {ex.Message}");
    return 2;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Azure authentication failed: {ex.Message}");
    return 2;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed (status {ex.Status}, code {ex.ErrorCode}): {ex.Message}");
    return 3;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine("The operation was canceled.");
    return 4;
}
catch (InvalidOperationException ex)
{
    Console.Error.WriteLine(ex.Message);
    return 5;
}
finally
{
    if (createdResourceGroup is not null)
    {
        try
        {
            Console.WriteLine(
                $"\nCleaning up resource group '{resourceGroupName}' after an error...");
            await createdResourceGroup.DeleteAsync(WaitUntil.Completed);
            Console.WriteLine("Cleanup completed.");
        }
        catch (RequestFailedException cleanupException)
        {
            Console.Error.WriteLine(
                $"Cleanup failed (status {cleanupException.Status}, " +
                $"code {cleanupException.ErrorCode}): {cleanupException.Message}");
        }
    }
}

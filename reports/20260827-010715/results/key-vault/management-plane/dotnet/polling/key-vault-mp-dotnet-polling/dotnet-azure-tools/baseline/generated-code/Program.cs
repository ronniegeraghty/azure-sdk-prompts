using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.KeyVault;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.Resources;
using Azure.Security.KeyVault.Secrets;

const string secretsOfficerRoleId = "b86a8fe4-44ce-4948-aee5-eccb2c155cd7";

if (!args.Contains("--execute", StringComparer.OrdinalIgnoreCase))
{
    Console.WriteLine(
        "Dry run only. Pass --execute after setting AZURE_TENANT_ID, " +
        "AZURE_RESOURCE_GROUP, and AZURE_KEY_VAULT_NAME.");
    return;
}

Guid tenantId = GetRequiredGuid("AZURE_TENANT_ID");
string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string vaultName = GetRequiredEnvironmentVariable("AZURE_KEY_VAULT_NAME");
Guid? principalObjectId = GetOptionalGuid("AZURE_PRINCIPAL_OBJECT_ID");

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential);

SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();
ResourceGroupResource resourceGroup =
    await subscription.GetResourceGroups().GetAsync(resourceGroupName);

var properties = new KeyVaultProperties(
    tenantId,
    new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))
{
    EnableRbacAuthorization = true,
    EnableSoftDelete = true,
    SoftDeleteRetentionInDays = 90,
    EnablePurgeProtection = true
};

var createContent = new KeyVaultCreateOrUpdateContent(
    AzureLocation.EastUS,
    properties);

KeyVaultCollection vaults = resourceGroup.GetKeyVaults();

// Start the LRO without blocking, then explicitly poll it to completion.
ArmOperation<KeyVaultResource> createOperation =
    await vaults.CreateOrUpdateAsync(WaitUntil.Started, vaultName, createContent);

Console.WriteLine($"Vault creation started. Operation ID: {createOperation.Id}");

Response<KeyVaultResource> completedResponse =
    await createOperation.WaitForCompletionAsync();
KeyVaultResource vault = completedResponse.Value;

Console.WriteLine($"Vault creation completed: {vault.Id}");

if (principalObjectId is Guid principalId)
{
    string subscriptionId = subscription.Data.SubscriptionId;
    var roleDefinitionId = new ResourceIdentifier(
        $"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/" +
        $"roleDefinitions/{secretsOfficerRoleId}");

    var assignmentContent =
        new RoleAssignmentCreateOrUpdateContent(roleDefinitionId, principalId);

    await vault.GetRoleAssignments().CreateOrUpdateAsync(
        WaitUntil.Completed,
        Guid.NewGuid().ToString(),
        assignmentContent);

    Console.WriteLine(
        $"Assigned Key Vault Secrets Officer to principal {principalId}.");
}
else
{
    Console.WriteLine(
        "No data-plane role was assigned. Set AZURE_PRINCIPAL_OBJECT_ID " +
        "to assign Key Vault Secrets Officer.");
}

Uri vaultUri = vault.Data.Properties.VaultUri
    ?? throw new InvalidOperationException("The completed vault has no vault URI.");

var secretClient = new SecretClient(vaultUri, credential);
Console.WriteLine($"SecretClient created for {secretClient.VaultUri}");

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException(
            $"Set the required environment variable {name}.")
        : value;
}

static Guid GetRequiredGuid(string name)
{
    string value = GetRequiredEnvironmentVariable(name);
    return Guid.TryParse(value, out Guid result)
        ? result
        : throw new InvalidOperationException($"{name} must be a GUID.");
}

static Guid? GetOptionalGuid(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    if (string.IsNullOrWhiteSpace(value))
    {
        return null;
    }

    return Guid.TryParse(value, out Guid result)
        ? result
        : throw new InvalidOperationException($"{name} must be a GUID.");
}

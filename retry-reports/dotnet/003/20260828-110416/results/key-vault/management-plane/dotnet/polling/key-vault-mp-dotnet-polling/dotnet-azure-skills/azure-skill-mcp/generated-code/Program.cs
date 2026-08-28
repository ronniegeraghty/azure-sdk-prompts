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

const string secretsUserRoleDefinitionId = "4633458b-17de-408a-b874-0445c86b69e6";

string subscriptionId = RequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string tenantId = RequiredEnvironmentVariable("AZURE_TENANT_ID");
string resourceGroupName = RequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string vaultName = RequiredEnvironmentVariable("AZURE_KEY_VAULT_NAME");
string? rbacPrincipalObjectId = Environment.GetEnvironmentVariable(
    "AZURE_KEY_VAULT_RBAC_PRINCIPAL_OBJECT_ID");

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential, subscriptionId);

SubscriptionResource subscription = armClient.GetSubscriptionResource(
    SubscriptionResource.CreateResourceIdentifier(subscriptionId));
ResourceGroupResource resourceGroup =
    (await subscription.GetResourceGroups().GetAsync(resourceGroupName)).Value;

var properties = new KeyVaultProperties(
    Guid.Parse(tenantId),
    new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))
{
    EnableRbacAuthorization = true,
    EnableSoftDelete = true,
    SoftDeleteRetentionInDays = 90,
    EnablePurgeProtection = true
};

var content = new KeyVaultCreateOrUpdateContent(AzureLocation.EastUS, properties);
KeyVaultCollection vaults = resourceGroup.GetKeyVaults();

// WaitUntil.Started exposes the ArmOperation so callers control when to poll.
ArmOperation<KeyVaultResource> createOperation =
    await vaults.CreateOrUpdateAsync(WaitUntil.Started, vaultName, content);

Console.WriteLine($"Vault creation started. Operation ID: {createOperation.Id}");
await createOperation.WaitForCompletionAsync();
KeyVaultResource vault = createOperation.Value;
Console.WriteLine($"Vault created: {vault.Id}");

if (Guid.TryParse(rbacPrincipalObjectId, out Guid principalId))
{
    ResourceIdentifier roleDefinitionId = new(
        $"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/" +
        $"roleDefinitions/{secretsUserRoleDefinitionId}");

    var roleContent = new RoleAssignmentCreateOrUpdateContent(
        roleDefinitionId,
        principalId);

    RoleAssignmentCollection roleAssignments = vault.GetRoleAssignments();
    ArmOperation<RoleAssignmentResource> roleOperation =
        await roleAssignments.CreateOrUpdateAsync(
            WaitUntil.Completed,
            Guid.NewGuid().ToString(),
            roleContent);

    Console.WriteLine(
        $"Assigned Key Vault Secrets User to {principalId}: {roleOperation.Value.Id}");
}
else
{
    Console.WriteLine(
        "No data-plane role was assigned. Set " +
        "AZURE_KEY_VAULT_RBAC_PRINCIPAL_OBJECT_ID to a principal object ID.");
}

Uri vaultUri = vault.Data.Properties.VaultUri
    ?? throw new InvalidOperationException("The completed vault has no vault URI.");
var secretClient = new SecretClient(vaultUri, credential);

// Constructing SecretClient does not make a network request; listing one page does.
await foreach (Page<Azure.Security.KeyVault.Secrets.SecretProperties> _ in secretClient
    .GetPropertiesOfSecretsAsync()
    .AsPages(pageSizeHint: 1))
{
    break;
}

Console.WriteLine($"SecretClient successfully accessed {vaultUri}");

static string RequiredEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException(
        $"Set the required environment variable {name}.");

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

string subscriptionId = RequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string resourceGroupName = RequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string vaultName = RequiredEnvironmentVariable("KEY_VAULT_NAME");
Guid tenantId = Guid.Parse(RequiredEnvironmentVariable("AZURE_TENANT_ID"));
Guid principalObjectId = Guid.Parse(RequiredEnvironmentVariable("KEY_VAULT_PRINCIPAL_OBJECT_ID"));

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential, subscriptionId);

SubscriptionResource subscription = armClient.GetSubscriptionResource(
    SubscriptionResource.CreateResourceIdentifier(subscriptionId));
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

var createContent = new KeyVaultCreateOrUpdateContent(AzureLocation.EastUS, properties);
KeyVaultCollection vaults = resourceGroup.GetKeyVaults();

// Start the long-running operation without blocking, then explicitly poll it.
ArmOperation<KeyVaultResource> createOperation =
    await vaults.CreateOrUpdateAsync(WaitUntil.Started, vaultName, createContent);

Console.WriteLine($"Vault creation started. Operation ID: {createOperation.Id}");
KeyVaultResource vault = await createOperation.WaitForCompletionAsync();
Console.WriteLine($"Vault created: {vault.Data.Id}");

// RBAC assignments cannot be embedded in the vault create payload. Create one at
// the vault scope after the resource exists.
ResourceIdentifier roleDefinitionId = new(
    $"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/" +
    $"roleDefinitions/{secretsOfficerRoleId}");
var roleAssignmentContent = new RoleAssignmentCreateOrUpdateContent(
    roleDefinitionId,
    principalObjectId);

RoleAssignmentCollection roleAssignments = vault.GetRoleAssignments();
ArmOperation<RoleAssignmentResource> roleOperation =
    await roleAssignments.CreateOrUpdateAsync(
        WaitUntil.Completed,
        Guid.NewGuid().ToString(),
        roleAssignmentContent);

Console.WriteLine($"RBAC role assignment created: {roleOperation.Value.Data.Id}");

Uri vaultUri = vault.Data.Properties.VaultUri
    ?? throw new InvalidOperationException("Azure did not return a vault URI.");
var secretClient = new SecretClient(vaultUri, credential);

await VerifyDataPlaneAccessAsync(secretClient);
Console.WriteLine($"SecretClient successfully accessed {secretClient.VaultUri}");

static async Task VerifyDataPlaneAccessAsync(SecretClient client)
{
    const int maximumAttempts = 12;

    for (int attempt = 1; attempt <= maximumAttempts; attempt++)
    {
        try
        {
            await using IAsyncEnumerator<Page<Azure.Security.KeyVault.Secrets.SecretProperties>> pages = client
                .GetPropertiesOfSecretsAsync()
                .AsPages(pageSizeHint: 1)
                .GetAsyncEnumerator();
            await pages.MoveNextAsync();
            return;
        }
        catch (RequestFailedException exception)
            when (exception.Status == 403 && attempt < maximumAttempts)
        {
            Console.WriteLine(
                $"Waiting for RBAC propagation ({attempt}/{maximumAttempts})...");
            await Task.Delay(TimeSpan.FromSeconds(10));
        }
    }
}

static string RequiredEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException(
        $"Set the required environment variable '{name}'.");

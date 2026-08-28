using System.Security.Cryptography;
using System.Text;
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

const string keyVaultSecretsUserRoleId = "4633458b-17de-408a-b874-0445c86b69e6";

string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
Guid tenantId = Guid.Parse(GetRequiredEnvironmentVariable("AZURE_TENANT_ID"));
string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string vaultName = GetRequiredEnvironmentVariable("KEY_VAULT_NAME");
Guid principalObjectId = Guid.Parse(
    GetRequiredEnvironmentVariable("KEY_VAULT_PRINCIPAL_OBJECT_ID"));

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential, subscriptionId);

SubscriptionResource subscription = armClient.GetSubscriptionResource(
    SubscriptionResource.CreateResourceIdentifier(subscriptionId));
ResourceGroupResource resourceGroup =
    (await subscription.GetResourceGroupAsync(resourceGroupName)).Value;

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

// WaitUntil.Started returns after Azure accepts the request, leaving polling to ArmOperation.
ArmOperation<KeyVaultResource> createOperation = await vaults.CreateOrUpdateAsync(
    WaitUntil.Started,
    vaultName,
    createContent);

Console.WriteLine($"Vault creation started. Operation ID: {createOperation.Id}");

await createOperation.WaitForCompletionAsync();
KeyVaultResource vault = createOperation.Value;

Console.WriteLine($"Vault creation completed: {vault.Id}");

// Data-plane RBAC assignments are separate ARM resources, so the vault must exist first.
ResourceIdentifier roleDefinitionId = new(
    $"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/" +
    $"roleDefinitions/{keyVaultSecretsUserRoleId}");

var roleAssignmentContent = new RoleAssignmentCreateOrUpdateContent(
    roleDefinitionId,
    principalObjectId)
{
    Description = "Allows this sample's identity to read Key Vault secrets."
};

Guid roleAssignmentName = CreateDeterministicGuid(
    $"{vault.Id}|{principalObjectId}|{keyVaultSecretsUserRoleId}");

ArmOperation<RoleAssignmentResource> roleAssignmentOperation =
    await armClient.GetRoleAssignments(vault.Id).CreateOrUpdateAsync(
        WaitUntil.Completed,
        roleAssignmentName.ToString(),
        roleAssignmentContent);

Console.WriteLine($"Role assignment ready: {roleAssignmentOperation.Value.Id}");

Uri vaultUri = vault.Data.Properties.VaultUri
    ?? throw new InvalidOperationException("Azure did not return the vault URI.");
var secretClient = new SecretClient(vaultUri, credential);

// A client can be constructed before RBAC propagation finishes. Make a harmless list request
// and retry only the expected temporary 403 response.
await VerifySecretClientAccessAsync(secretClient);
Console.WriteLine($"SecretClient can access {secretClient.VaultUri}");

static async Task VerifySecretClientAccessAsync(
    SecretClient client,
    CancellationToken cancellationToken = default)
{
    const int maximumAttempts = 10;

    for (int attempt = 1; ; attempt++)
    {
        try
        {
            await using IAsyncEnumerator<SecretProperties> secrets =
                client.GetPropertiesOfSecretsAsync(cancellationToken)
                    .GetAsyncEnumerator(cancellationToken);

            _ = await secrets.MoveNextAsync();
            return;
        }
        catch (RequestFailedException exception)
            when (exception.Status == 403 && attempt < maximumAttempts)
        {
            TimeSpan delay = TimeSpan.FromSeconds(Math.Min(60, attempt * 10));
            Console.WriteLine(
                $"Waiting for RBAC propagation (attempt {attempt}/{maximumAttempts})...");
            await Task.Delay(delay, cancellationToken);
        }
    }
}

static Guid CreateDeterministicGuid(string value)
{
    byte[] hash = SHA256.HashData(Encoding.UTF8.GetBytes(value));
    return new Guid(hash.AsSpan(0, 16));
}

static string GetRequiredEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name) is { Length: > 0 } value
        ? value
        : throw new InvalidOperationException(
            $"Set the required environment variable {name}.");

// Access-policy alternative (do not combine this with EnableRbacAuthorization = true):
//
// var permissions = new IdentityAccessPermissions();
// permissions.Secrets.Add(IdentityAccessSecretPermission.Get);
// permissions.Secrets.Add(IdentityAccessSecretPermission.List);
//
// var accessPolicyProperties = new KeyVaultProperties(
//     tenantId,
//     new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))
// {
//     EnableRbacAuthorization = false,
//     EnableSoftDelete = true,
//     SoftDeleteRetentionInDays = 90,
//     EnablePurgeProtection = true
// };
// accessPolicyProperties.AccessPolicies.Add(
//     new KeyVaultAccessPolicy(
//         tenantId,
//         principalObjectId.ToString(),
//         permissions));
//
// Pass accessPolicyProperties to KeyVaultCreateOrUpdateContent. Unlike an RBAC role
// assignment, this access policy is included directly in the vault creation request.

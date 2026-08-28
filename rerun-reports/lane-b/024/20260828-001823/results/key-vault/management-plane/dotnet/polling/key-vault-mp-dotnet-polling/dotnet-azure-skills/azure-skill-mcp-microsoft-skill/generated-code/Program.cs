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

string subscriptionId = RequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string resourceGroupName = RequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string vaultName = RequiredEnvironmentVariable("KEY_VAULT_NAME");
Guid tenantId = Guid.Parse(RequiredEnvironmentVariable("AZURE_TENANT_ID"));

var credential = new DefaultAzureCredential();
var armClient = new ArmClient(credential, subscriptionId);

ResourceIdentifier resourceGroupId =
    ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);
KeyVaultCollection vaults = resourceGroup.GetKeyVaults();

var properties = new KeyVaultProperties(
    tenantId,
    new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))
{
    EnableRbacAuthorization = true,
    EnableSoftDelete = true,
    SoftDeleteRetentionInDays = 90,
    EnablePurgeProtection = true,
    PublicNetworkAccess = "Enabled"
};

var content = new KeyVaultCreateOrUpdateContent(
    new AzureLocation("eastus"),
    properties);

Console.WriteLine($"Starting creation of Key Vault '{vaultName}'...");

ArmOperation<KeyVaultResource> createOperation =
    await vaults.CreateOrUpdateAsync(WaitUntil.Started, vaultName, content);

Console.WriteLine(
    $"Operation started. HTTP status: {createOperation.GetRawResponse().Status}");

await createOperation.WaitForCompletionAsync();
KeyVaultResource vault = createOperation.Value;

Console.WriteLine($"Vault provisioning completed: {vault.Data.Id}");

string? principalObjectIdValue =
    Environment.GetEnvironmentVariable("PRINCIPAL_OBJECT_ID");

if (!string.IsNullOrWhiteSpace(principalObjectIdValue))
{
    Guid principalObjectId = Guid.Parse(principalObjectIdValue);
    ResourceIdentifier roleDefinitionId = new(
        $"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/" +
        $"roleDefinitions/{keyVaultSecretsUserRoleId}");

    var roleAssignmentContent =
        new RoleAssignmentCreateOrUpdateContent(roleDefinitionId, principalObjectId);

    string roleAssignmentName = CreateDeterministicGuid(
        $"{vault.Data.Id}|{principalObjectId}|{keyVaultSecretsUserRoleId}")
        .ToString();

    ArmOperation<RoleAssignmentResource> roleOperation =
        await vault.GetRoleAssignments().CreateOrUpdateAsync(
            WaitUntil.Started,
            roleAssignmentName,
            roleAssignmentContent);

    await roleOperation.WaitForCompletionAsync();
    Console.WriteLine("Assigned the Key Vault Secrets User role.");
}
else
{
    Console.WriteLine(
        "PRINCIPAL_OBJECT_ID was not set; assuming this identity already has " +
        "a Key Vault data-plane role.");
}

Uri vaultUri = vault.Data.Properties.VaultUri
    ?? new Uri($"https://{vaultName}.vault.azure.net");
var secretClient = new SecretClient(vaultUri, credential);

await VerifySecretClientAccessAsync(secretClient);
Console.WriteLine($"SecretClient can access {secretClient.VaultUri}.");

static async Task VerifySecretClientAccessAsync(SecretClient client)
{
    const int attempts = 12;

    for (int attempt = 1; attempt <= attempts; attempt++)
    {
        try
        {
            await foreach (Page<Azure.Security.KeyVault.Secrets.SecretProperties> _ in
                client.GetPropertiesOfSecretsAsync().AsPages(pageSizeHint: 1))
            {
                break;
            }

            return;
        }
        catch (RequestFailedException ex) when (
            ex.Status == 403 && attempt < attempts)
        {
            Console.WriteLine(
                $"Waiting for RBAC propagation ({attempt}/{attempts - 1})...");
            await Task.Delay(TimeSpan.FromSeconds(10));
        }
    }
}

static Guid CreateDeterministicGuid(string value)
{
    byte[] hash = SHA256.HashData(Encoding.UTF8.GetBytes(value));
    return new Guid(hash.AsSpan(0, 16));
}

static string RequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new InvalidOperationException(
            $"Set the {name} environment variable before running the sample.");
}

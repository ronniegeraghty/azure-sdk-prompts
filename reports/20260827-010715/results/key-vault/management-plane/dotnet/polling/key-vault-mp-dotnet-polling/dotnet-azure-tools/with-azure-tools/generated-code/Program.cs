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

namespace KeyVaultProvisioning;

internal static class Program
{
    private const string KeyVaultSecretsOfficerRoleId =
        "4633458b-17de-408a-b874-0445c86b69e6";

    private static async Task<int> Main()
    {
        try
        {
            string subscriptionId = GetRequiredSetting("AZURE_SUBSCRIPTION_ID");
            string resourceGroupName = GetRequiredSetting("AZURE_RESOURCE_GROUP");
            string vaultName = GetRequiredSetting("AZURE_KEY_VAULT_NAME");
            Guid tenantId = Guid.Parse(GetRequiredSetting("AZURE_TENANT_ID"));
            Guid principalObjectId =
                Guid.Parse(GetRequiredSetting("AZURE_PRINCIPAL_OBJECT_ID"));

            var credential = new DefaultAzureCredential();
            var armClient = new ArmClient(credential);

            var subscriptionResourceId =
                SubscriptionResource.CreateResourceIdentifier(subscriptionId);
            SubscriptionResource subscription =
                armClient.GetSubscriptionResource(subscriptionResourceId);
            ResourceGroupResource resourceGroup =
                await subscription.GetResourceGroupAsync(resourceGroupName);

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

            Console.WriteLine($"Starting creation of Key Vault '{vaultName}'...");

            ArmOperation<KeyVaultResource> createOperation =
                await resourceGroup.GetKeyVaults().CreateOrUpdateAsync(
                    WaitUntil.Started,
                    vaultName,
                    createContent);

            Console.WriteLine($"Operation ID: {createOperation.Id}");
            await createOperation.WaitForCompletionResponseAsync(
                TimeSpan.FromSeconds(10));

            KeyVaultResource vault = createOperation.Value;
            Console.WriteLine($"Vault created: {vault.Id}");

            await AssignSecretsOfficerRoleAsync(
                armClient,
                vault,
                subscriptionId,
                principalObjectId);

            Uri vaultUri = vault.Data.Properties.VaultUri
                ?? new Uri($"https://{vaultName}.vault.azure.net");
            var secretClient = new SecretClient(vaultUri, credential);

            Console.WriteLine(
                $"SecretClient created for {secretClient.VaultUri}. " +
                "The client is ready for data-plane secret operations.");

            return 0;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            return 1;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
            return 1;
        }
        catch (FormatException ex)
        {
            Console.Error.WriteLine($"A GUID setting is invalid: {ex.Message}");
            return 1;
        }
        catch (InvalidOperationException ex)
        {
            Console.Error.WriteLine(ex.Message);
            return 1;
        }
    }

    private static async Task AssignSecretsOfficerRoleAsync(
        ArmClient armClient,
        KeyVaultResource vault,
        string subscriptionId,
        Guid principalObjectId)
    {
        var roleDefinitionId = new ResourceIdentifier(
            $"/subscriptions/{subscriptionId}" +
            $"/providers/Microsoft.Authorization/roleDefinitions/" +
            KeyVaultSecretsOfficerRoleId);
        var assignmentContent = new RoleAssignmentCreateOrUpdateContent(
            roleDefinitionId,
            principalObjectId);

        RoleAssignmentCollection roleAssignments =
            armClient.GetRoleAssignments(vault.Id);
        ArmOperation<RoleAssignmentResource> assignmentOperation =
            await roleAssignments.CreateOrUpdateAsync(
                WaitUntil.Completed,
                Guid.NewGuid().ToString(),
                assignmentContent);

        Console.WriteLine($"RBAC role assigned: {assignmentOperation.Value.Id}");
    }

    // Access policies are ignored when EnableRbacAuthorization is true.
    // Use this payload instead of the RBAC payload only for a legacy policy-based vault.
    private static KeyVaultCreateOrUpdateContent CreateAccessPolicyContent(
        Guid tenantId,
        Guid principalObjectId)
    {
        var permissions = new IdentityAccessPermissions();
        permissions.Secrets.Add(IdentityAccessSecretPermission.Get);
        permissions.Secrets.Add(IdentityAccessSecretPermission.List);
        permissions.Secrets.Add(IdentityAccessSecretPermission.Set);

        var properties = new KeyVaultProperties(
            tenantId,
            new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))
        {
            EnableRbacAuthorization = false,
            EnableSoftDelete = true,
            SoftDeleteRetentionInDays = 90,
            EnablePurgeProtection = true
        };
        properties.AccessPolicies.Add(
            new KeyVaultAccessPolicy(
                tenantId,
                principalObjectId.ToString(),
                permissions));

        return new KeyVaultCreateOrUpdateContent(
            AzureLocation.EastUS,
            properties);
    }

    private static string GetRequiredSetting(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        return string.IsNullOrWhiteSpace(value)
            ? throw new InvalidOperationException(
                $"Required environment variable '{name}' is not set.")
            : value;
    }
}

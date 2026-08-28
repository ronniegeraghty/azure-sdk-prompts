using Azure.Core;
using Azure.ResourceManager.KeyVault.Models;

internal static class LegacyAccessPolicyExample
{
    // Access policies are the legacy alternative to RBAC. Do not combine the
    // policy model with EnableRbacAuthorization = true.
    public static KeyVaultCreateOrUpdateContent BuildContent(
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
            new AzureLocation("eastus"),
            properties);
    }
}

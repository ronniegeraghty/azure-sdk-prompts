# Azure Key Vault management-plane LRO sample

This sample uses `DefaultAzureCredential` and
`Azure.ResourceManager.KeyVault` to create an RBAC-enabled vault in `eastus`.
It starts the create/update long-running operation with `WaitUntil.Started`,
then explicitly polls it with `WaitForCompletionAsync`.

## Required packages

```powershell
dotnet add package Azure.Identity --version 1.21.0
dotnet add package Azure.ResourceManager --version 1.14.0
dotnet add package Azure.ResourceManager.KeyVault --version 1.4.0
dotnet add package Azure.Security.KeyVault.Secrets --version 4.11.0

# Required only when creating the optional RBAC role assignment:
dotnet add package Azure.ResourceManager.Authorization --version 1.1.7
```

The resource group must already exist. Vault creation is an Azure control-plane
operation and requires an identity with permission such as Key Vault
Contributor at the resource-group scope. Creating an RBAC role assignment also
requires `Microsoft.Authorization/roleAssignments/write`.

Set the inputs, authenticate through any credential supported by
`DefaultAzureCredential`, and explicitly opt in to resource creation:

```powershell
$env:AZURE_TENANT_ID = "<tenant-guid>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_KEY_VAULT_NAME = "<globally-unique-vault-name>"

# Optional: assign Key Vault Secrets Officer after vault creation.
$env:AZURE_PRINCIPAL_OBJECT_ID = "<user-service-principal-or-managed-identity-object-id>"

dotnet run -- --execute
```

Without `--execute`, the program performs a local dry run and exits without
contacting Azure.

## RBAC versus access policies

For the recommended RBAC model, set this property in the creation payload:

```csharp
properties.EnableRbacAuthorization = true;
```

Azure RBAC role assignments cannot be included inside the Key Vault creation
payload. They are separate `Microsoft.Authorization/roleAssignments`
resources, so the sample creates the vault first and then assigns the built-in
**Key Vault Secrets Officer** role at vault scope:

```csharp
var assignment = new RoleAssignmentCreateOrUpdateContent(
    roleDefinitionId,
    principalObjectId);

await vault.GetRoleAssignments().CreateOrUpdateAsync(
    WaitUntil.Completed,
    Guid.NewGuid().ToString(),
    assignment);
```

To use the legacy access-policy model instead, disable RBAC and add policies
directly to `KeyVaultProperties` before creating the vault:

```csharp
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
```

Use either RBAC or access policies for data-plane authorization, not both.
Constructing `SecretClient` verifies that the completed management-plane result
contains a usable vault endpoint. An actual data-plane request may briefly fail
after a new RBAC assignment because role propagation is eventually consistent.

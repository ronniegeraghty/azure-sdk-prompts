# Azure Key Vault management-plane LRO sample

This sample uses `DefaultAzureCredential` and `Azure.ResourceManager.KeyVault`
to create an RBAC-enabled vault in `eastus`. Soft delete has a 90-day retention
period, purge protection is enabled, and the `ArmOperation<T>` returned by
`WaitUntil.Started` is explicitly polled to completion. It then assigns the
built-in **Key Vault Secrets Officer** role and verifies data-plane access with
`SecretClient`.

The project references these packages:

```powershell
dotnet add package Azure.Identity --version 1.21.0
dotnet add package Azure.ResourceManager.Authorization --version 1.1.7
dotnet add package Azure.ResourceManager.KeyVault --version 1.4.0
dotnet add package Azure.Security.KeyVault.Secrets --version 4.11.0
```

Set the inputs before running:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:KEY_VAULT_NAME = "<globally-unique-vault-name>"
$env:KEY_VAULT_PRINCIPAL_OBJECT_ID = "<user-group-or-service-principal-object-id>"
dotnet run
```

The signed-in identity needs permission to create vaults and role assignments.
`KEY_VAULT_PRINCIPAL_OBJECT_ID` is an Entra object ID, not an application/client
ID. No command in this repository runs against Azure unless you explicitly run
the program.

## Access policies instead of RBAC

Access policies and RBAC are alternative authorization models. To put a legacy
access policy in the vault creation request, set RBAC to `false` and add a
policy before constructing `KeyVaultCreateOrUpdateContent`:

```csharp
var permissions = new IdentityAccessPermissions
{
    Secrets =
    {
        IdentityAccessSecretPermission.Get,
        IdentityAccessSecretPermission.List,
        IdentityAccessSecretPermission.Set
    }
};

properties.EnableRbacAuthorization = false;
properties.AccessPolicies.Add(
    new KeyVaultAccessPolicy(
        tenantId,
        principalObjectId.ToString(),
        permissions));
```

When using this variant, remove the `RoleAssignmentCollection` block from
`Program.cs`. Access policies are part of the vault create payload; Azure RBAC
role assignments are separate `Microsoft.Authorization/roleAssignments`
resources and therefore can only be created at the vault scope once the vault
resource exists.

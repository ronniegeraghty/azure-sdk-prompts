# Azure Key Vault management-plane sample

This local console sample uses `DefaultAzureCredential` to create an
RBAC-enabled Key Vault in `eastus`, waits for the management-plane
`ArmOperation<KeyVaultResource>` to finish, optionally creates a vault-scoped
role assignment, and verifies data-plane access with `SecretClient`.

## Required packages

```powershell
dotnet add package Azure.Identity --version 1.13.2
dotnet add package Azure.ResourceManager.Authorization --version 1.1.1
dotnet add package Azure.ResourceManager.KeyVault --version 1.3.2
dotnet add package Azure.Security.KeyVault.Secrets --version 4.7.0
```

## Configuration

The resource group must already exist. Set these variables before running:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_KEY_VAULT_NAME = "<globally-unique-vault-name>"

# Optional: assign Key Vault Secrets User at the new vault's scope.
$env:AZURE_KEY_VAULT_RBAC_PRINCIPAL_OBJECT_ID = "<principal-object-id>"

dotnet run
```

The authenticated identity needs permission to create vaults and role
assignments. The optional principal ID is the Microsoft Entra **object ID**,
not an application/client ID. To make the final access check succeed, use the
object ID of the user, service principal, or managed identity selected by
`DefaultAzureCredential`.

Azure RBAC can take several minutes to propagate. A `403 Forbidden` from the
final data-plane request immediately after role assignment can therefore be
transient; retry the program after propagation completes.

## Access policies versus Azure RBAC

This sample sets `EnableRbacAuthorization = true`. With that setting, legacy
Key Vault access policies do not grant data-plane access; use Azure role
assignments such as the vault-scoped assignment shown in `Program.cs`.
Role assignments are separate ARM resources, so the management SDK creates
one after the vault creation operation completes.

For a legacy access-policy vault instead, set
`EnableRbacAuthorization = false` and add policies before creating the vault:

```csharp
properties.AccessPolicies.Add(new KeyVaultAccessPolicy(
    tenantId: Guid.Parse(tenantId),
    objectId: Guid.Parse("<principal-object-id>"),
    permissions: new KeyVaultAccessPolicyPermissions
    {
        Secrets = { KeyVaultSecretPermission.Get, KeyVaultSecretPermission.List }
    }));
```

Choose one authorization model. Do not configure access policies as a
fallback on an RBAC-enabled vault.

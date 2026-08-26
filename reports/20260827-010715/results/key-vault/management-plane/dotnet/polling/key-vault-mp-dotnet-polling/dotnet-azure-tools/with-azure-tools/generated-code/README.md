# Azure Key Vault management-plane LRO sample

This console program:

- authenticates with `DefaultAzureCredential`;
- creates a standard Key Vault in `eastus`;
- enables Azure RBAC authorization, soft delete with 90-day retention, and
  purge protection;
- starts creation with `WaitUntil.Started`, then explicitly polls the
  `ArmOperation<KeyVaultResource>` to completion;
- assigns the built-in **Key Vault Secrets Officer** role at vault scope; and
- constructs a data-plane `SecretClient` using the created vault URI.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager.KeyVault
dotnet add package Azure.ResourceManager.Authorization
dotnet add package Azure.Security.KeyVault.Secrets
```

The project currently resolves these versions:

| Package | Version |
|---|---:|
| `Azure.Identity` | 1.21.0 |
| `Azure.ResourceManager.KeyVault` | 1.4.0 |
| `Azure.ResourceManager.Authorization` | 1.1.7 |
| `Azure.Security.KeyVault.Secrets` | 4.11.0 |

## Configuration

Set these environment variables before running:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-guid>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_KEY_VAULT_NAME = "<globally-unique-vault-name>"
$env:AZURE_TENANT_ID = "<tenant-guid>"
$env:AZURE_PRINCIPAL_OBJECT_ID = "<user-or-service-principal-object-guid>"

dotnet run
```

`AZURE_PRINCIPAL_OBJECT_ID` is the Microsoft Entra **object ID**, not an
application/client ID. The caller needs permission to create the vault and to
write role assignments at the vault scope, such as appropriately scoped
management-plane roles.

`DefaultAzureCredential` can use local developer credentials or environment
variables. In Azure-hosted production code, prefer a managed identity.

## RBAC versus access policies

For the recommended RBAC model, the creation payload sets
`EnableRbacAuthorization = true`. A role assignment cannot be embedded in the
Key Vault creation payload, so the sample creates it as a separate ARM
operation immediately after the vault LRO completes.

For the legacy access-policy model, set `EnableRbacAuthorization = false` and
add `KeyVaultAccessPolicy` entries to `KeyVaultProperties.AccessPolicies`.
`CreateAccessPolicyContent` in `Program.cs` shows that alternative. Do not
combine the models: access policies are ignored when RBAC authorization is
enabled.

Constructing `SecretClient` validates the resulting vault endpoint and client
configuration without writing a secret. Azure RBAC assignments can take a
short time to propagate before the first data-plane request succeeds.

## References

- [Azure.ResourceManager.KeyVault package overview](https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.keyvault-readme)
- [Azure Key Vault RBAC guide](https://learn.microsoft.com/azure/key-vault/general/rbac-guide)
- [DefaultAzureCredential overview](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains)
- [Azure Key Vault Secrets client library](https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme)

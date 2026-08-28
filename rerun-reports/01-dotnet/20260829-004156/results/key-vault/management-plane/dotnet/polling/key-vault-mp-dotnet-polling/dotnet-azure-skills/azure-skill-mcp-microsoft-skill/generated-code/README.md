# Azure Key Vault management SDK sample

This sample creates an RBAC-enabled Key Vault in `eastus`, explicitly enables
soft delete with 90-day retention and purge protection, waits for the
management-plane long-running operation, creates a vault-scoped **Key Vault
Secrets User** role assignment, and verifies data-plane access with
`SecretClient`.

## Packages

```powershell
dotnet add package Azure.Identity --version 1.17.0
dotnet add package Azure.ResourceManager.Authorization --version 1.1.7
dotnet add package Azure.ResourceManager.KeyVault --version 1.4.0
dotnet add package Azure.Security.KeyVault.Secrets --version 4.8.0
```

## Configuration

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:KEY_VAULT_NAME = "<globally-unique-vault-name>"
$env:KEY_VAULT_PRINCIPAL_OBJECT_ID = "<object-id-used-by-DefaultAzureCredential>"
dotnet run
```

`DefaultAzureCredential` uses your developer credential locally and managed
identity when Azure-hosted. The principal must be able to create vaults and
role assignments at the chosen scope. `KEY_VAULT_PRINCIPAL_OBJECT_ID` is the
Microsoft Entra object ID, not an application/client ID.

An RBAC role assignment cannot be embedded in Key Vault's create payload
because its scope is the completed vault resource. The program therefore
waits for vault creation and then creates the role assignment. The commented
alternative in `Program.cs` shows legacy access-policy authorization, which
can be embedded during creation but is ignored when
`EnableRbacAuthorization` is `true`.

## References

- [Azure Key Vault management library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.keyvault-readme?view=azure-dotnet)
- [Azure Key Vault RBAC guide](https://learn.microsoft.com/azure/key-vault/general/rbac-guide)
- [DefaultAzureCredential documentation](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains)

# Evaluation Report: key-vault-mp-dotnet-polling

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 191.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-mp-dotnet-polling` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 191.7s |
| Timestamp | 2026-08-26T18:03:00Z |
| Files Generated | 3 |
| Event Count | 5216 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 149.2s |
| Review | 41.7s |
| **Total** | **191.7s** |

## Configuration

- **name:** dotnet-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=384117 out=5082 |
| Turn Count | 12 |

## Prompt Sent

```
Write a C# program that creates
an Azure Key Vault using the management plane SDK and handles the long-running
operation:
1. Authenticate using DefaultAzureCredential
2. Create a Key Vault with RBAC authorization enabled in "eastus"
3. Configure soft-delete and purge protection
4. Wait for the vault creation to complete using the ArmOperation pattern
5. Verify the vault is accessible by creating a SecretClient pointing to it

Use Azure.ResourceManager.KeyVault. Show required NuGet packages and how
to set access policies or RBAC roles during creation.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-tools-baseline-4047949117","pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-tools-baseline-4047949117","pattern":"*.sln"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-tools-baseline-4047949117","pattern":"*.csproj"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":2000,"url":"https://api.nuget.org/v3-flatcontainer/azure.identity/index.json"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":2000,"url":"https://api.nuget.org/v3-flatcontainer/azure.resourcemanager/index.json"}
```

### 🔧 6. web_fetch

**Input:**

```json
{"max_length":2000,"url":"https://api.nuget.org/v3-flatcontainer/azure.resourcemanager.keyvault/index.json"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":2000,"url":"https://api.nuget.org/v3-flatcontainer/azure.security.keyvault.secrets/index.json"}
```

### 🔧 8. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":5,"query":"\"KeyVaultCreateOrUpdateContent\" language:C#"}
```

### 🔧 9. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":5,"query":"\"VaultCreateOrUpdateContent\" language:C# Azure.ResourceManager.KeyVault"}
```

### 🔧 10. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":10,"query":"\"EnableRbacAuthorization\" \"KeyVaultProperties\" language:C#"}
```

### 🔧 11. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"GetRoleAssignments()\" \"RoleAssignmentCreateOrUpdateContent\" language:C#"}
```

### 🔧 12. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"new KeyVaultProperties(\" language:C# path:sdk/keyvault"}
```

### 🔧 13. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"new KeyVaultAccessPolicy(\" language:C# path:sdk/keyvault"}
```

### 🔧 14. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/Azure.ResourceManager.KeyVault/api/Azure.ResourceManager.KeyVault.net8.0.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

### 🔧 15. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Azure.ResourceManager.KeyVault.csproj","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

### 🔧 16. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/identity/Azure.Identity/src/Azure.Identity.csproj","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

### 🔧 17. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/Azure.Security.KeyVault.Secrets/src/Azure.Security.KeyVault.Secrets.csproj","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

### 🔧 18. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Azure.ResourceManager.Authorization.csproj","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

### 🔧 19. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/resourcemanager/Azure.ResourceManager/src/Azure.ResourceManager.csproj","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

### 🔧 20. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: KeyVaultProvisioningSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Identity\" Version=\"1.21.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.ResourceManager\" Version=\"1.14.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.ResourceManager.Authorization\" Version=\"1.1.7\" /\u003e\n+    \u003cPackageReference Include=\"Azure.ResourceManager.KeyVault\" Version=\"1.4.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.Security.KeyVault.Secrets\" Version=\"4.11.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure;\n+using Azure.Core;\n+using Azure.Identity;\n+using Azure.ResourceManager;\n+using Azure.ResourceManager.Authorization;\n+using Azure.ResourceManager.Authorization.Models;\n+using Azure.ResourceManager.KeyVault;\n+using Azure.ResourceManager.KeyVault.Models;\n+using Azure.ResourceManager.Resources;\n+using Azure.Security.KeyVault.Secrets;\n+\n+const string secretsOfficerRoleId = \"b86a8fe4-44ce-4948-aee5-eccb2c155cd7\";\n+\n+if (!args.Contains(\"--execute\", StringComparer.OrdinalIgnoreCase))\n+{\n+    Console.WriteLine(\n+        \"Dry run only. Pass --execute after setting AZURE_TENANT_ID, \" +\n+        \"AZURE_RESOURCE_GROUP, and AZURE_KEY_VAULT_NAME.\");\n+    return;\n+}\n+\n+Guid tenantId = GetRequiredGuid(\"AZURE_TENANT_ID\");\n+string resourceGroupName = GetRequiredEnvironmentVariable(\"AZURE_RESOURCE_GROUP\");\n+string vaultName = GetRequiredEnvironmentVariable(\"AZURE_KEY_VAULT_NAME\");\n+Guid? principalObjectId = GetOptionalGuid(\"AZURE_PRINCIPAL_OBJECT_ID\");\n+\n+var credential = new DefaultAzureCredential();\n+var armClient = new ArmClient(credential);\n+\n+SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();\n+ResourceGroupResource resourceGroup =\n+    await subscription.GetResourceGroups().GetAsync(resourceGroupName);\n+\n+var properties = new KeyVaultProperties(\n+    tenantId,\n+    new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))\n+{\n+    EnableRbacAuthorization = true,\n+    EnableSoftDelete = true,\n+    SoftDeleteRetentionInDays = 90,\n+    EnablePurgeProtection = true\n+};\n+\n+var createContent = new KeyVaultCreateOrUpdateContent(\n+    AzureLocation.EastUS,\n+    properties);\n+\n+KeyVaultCollection vaults = resourceGroup.GetKeyVaults();\n+\n+// Start the LRO without blocking, then explicitly poll it to completion.\n+ArmOperation\u003cKeyVaultResource\u003e createOperation =\n+    await vaults.CreateOrUpdateAsync(WaitUntil.Started, vaultName, createContent);\n+\n+Console.WriteLine($\"Vault creation started. Operation ID: {createOperation.Id}\");\n+\n+ArmOperation\u003cKeyVaultResource\u003e completedOperation =\n+    await createOperation.WaitForCompletionAsync();\n+KeyVaultResource vault = completedOperation.Value;\n+\n+Console.WriteLine($\"Vault creation completed: {vault.Id}\");\n+\n+if (principalObjectId is Guid principalId)\n+{\n+    string subscriptionId = subscription.Data.SubscriptionId;\n+    var roleDefinitionId = new ResourceIdentifier(\n+        $\"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/\" +\n+        $\"roleDefinitions/{secretsOfficerRoleId}\");\n+\n+    var assignmentContent =\n+        new RoleAssignmentCreateOrUpdateContent(roleDefinitionId, principalId);\n+\n+    await vault.GetRoleAssignments().CreateOrUpdateAsync(\n+        WaitUntil.Completed,\n+        Guid.NewGuid().ToString(),\n+        assignmentContent);\n+\n+    Console.WriteLine(\n+        $\"Assigned Key Vault Secrets Officer to principal {principalId}.\");\n+}\n+else\n+{\n+    Console.WriteLine(\n+        \"No data-plane role was assigned. Set AZURE_PRINCIPAL_OBJECT_ID \" +\n+        \"to assign Key Vault Secrets Officer.\");\n+}\n+\n+Uri vaultUri = vault.Data.Properties.VaultUri\n+    ?? throw new InvalidOperationException(\"The completed vault has no vault URI.\");\n+\n+var secretClient = new SecretClient(vaultUri, credential);\n+Console.WriteLine($\"SecretClient created for {secretClient.VaultUri}\");\n+\n+static string GetRequiredEnvironmentVariable(string name)\n+{\n+    string? value = Environment.GetEnvironmentVariable(name);\n+    return string.IsNullOrWhiteSpace(value)\n+        ? throw new InvalidOperationException(\n+            $\"Set the required environment variable {name}.\")\n+        : value;\n+}\n+\n+static Guid GetRequiredGuid(string name)\n+{\n+    string value = GetRequiredEnvironmentVariable(name);\n+    return Guid.TryParse(value, out Guid result)\n+        ? result\n+        : throw new InvalidOperationException($\"{name} must be a GUID.\");\n+}\n+\n+static Guid? GetOptionalGuid(string name)\n+{\n+    string? value = Environment.GetEnvironmentVariable(name);\n+    if (string.IsNullOrWhiteSpace(value))\n+    {\n+        return null;\n+    }\n+\n+    return Guid.TryParse(value, out Guid result)\n+        ? result\n+        : throw new InvalidOperationException($\"{name} must be a GUID.\");\n+}\n*** Add File: README.md\n+# Azure Key Vault management-plane LRO sample\n+\n+This sample uses `DefaultAzureCredential` and\n+`Azure.ResourceManager.KeyVault` to create an RBAC-enabled vault in `eastus`.\n+It starts the create/update long-running operation with `WaitUntil.Started`,\n+then explicitly polls it with `WaitForCompletionAsync`.\n+\n+## Required packages\n+\n+```powershell\n+dotnet add package Azure.Identity --version 1.21.0\n+dotnet add package Azure.ResourceManager --version 1.14.0\n+dotnet add package Azure.ResourceManager.KeyVault --version 1.4.0\n+dotnet add package Azure.Security.KeyVault.Secrets --version 4.11.0\n+\n+# Required only when creating the optional RBAC role assignment:\n+dotnet add package Azure.ResourceManager.Authorization --version 1.1.7\n+```\n+\n+The resource group must already exist. Vault creation is an Azure control-plane\n+operation and requires an identity with permission such as Key Vault\n+Contributor at the resource-group scope. Creating an RBAC role assignment also\n+requires `Microsoft.Authorization/roleAssignments/write`.\n+\n+Set the inputs, authenticate through any credential supported by\n+`DefaultAzureCredential`, and explicitly opt in to resource creation:\n+\n+```powershell\n+$env:AZURE_TENANT_ID = \"\u003ctenant-guid\u003e\"\n+$env:AZURE_RESOURCE_GROUP = \"\u003cexisting-resource-group\u003e\"\n+$env:AZURE_KEY_VAULT_NAME = \"\u003cglobally-unique-vault-name\u003e\"\n+\n+# Optional: assign Key Vault Secrets Officer after vault creation.\n+$env:AZURE_PRINCIPAL_OBJECT_ID = \"\u003cuser-service-principal-or-managed-identity-object-id\u003e\"\n+\n+dotnet run -- --execute\n+```\n+\n+Without `--execute`, the program performs a local dry run and exits without\n+contacting Azure.\n+\n+## RBAC versus access policies\n+\n+For the recommended RBAC model, set this property in the creation payload:\n+\n+```csharp\n+properties.EnableRbacAuthorization = true;\n+```\n+\n+Azure RBAC role assignments cannot be included inside the Key Vault creation\n+payload. They are separate `Microsoft.Authorization/roleAssignments`\n+resources, so the sample creates the vault first and then assigns the built-in\n+**Key Vault Secrets Officer** role at vault scope:\n+\n+```csharp\n+var assignment = new RoleAssignmentCreateOrUpdateContent(\n+    roleDefinitionId,\n+    principalObjectId);\n+\n+await vault.GetRoleAssignments().CreateOrUpdateAsync(\n+    WaitUntil.Completed,\n+    Guid.NewGuid().ToString(),\n+    assignment);\n+```\n+\n+To use the legacy access-policy model instead, disable RBAC and add policies\n+directly to `KeyVaultProperties` before creating the vault:\n+\n+```csharp\n+var permissions = new IdentityAccessPermissions();\n+permissions.Secrets.Add(IdentityAccessSecretPermission.Get);\n+permissions.Secrets.Add(IdentityAccessSecretPermission.List);\n+permissions.Secrets.Add(IdentityAccessSecretPermission.Set);\n+\n+var properties = new KeyVaultProperties(\n+    tenantId,\n+    new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))\n+{\n+    EnableRbacAuthorization = false,\n+    EnableSoftDelete = true,\n+    SoftDeleteRetentionInDays = 90,\n+    EnablePurgeProtection = true\n+};\n+\n+properties.AccessPolicies.Add(\n+    new KeyVaultAccessPolicy(\n+        tenantId,\n+        principalObjectId.ToString(),\n+        permissions));\n+```\n+\n+Use either RBAC or access policies for data-plane authorization, not both.\n+Constructing `SecretClient` verifies that the completed management-plane result\n+contains a usable vault endpoint. An actual data-plane request may briefly fail\n+after a new RBAC assignment because role propagation is eventually consistent.\n*** End Patch\n"
```

### 🔧 21. powershell

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Compile the Key Vault sample","initial_wait":180,"mode":"sync"}
```

### 🔧 22. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: Program.cs\n@@\n-ArmOperation\u003cKeyVaultResource\u003e completedOperation =\n+Response\u003cKeyVaultResource\u003e completedResponse =\n     await createOperation.WaitForCompletionAsync();\n-KeyVaultResource vault = completedOperation.Value;\n+KeyVaultResource vault = completedResponse.Value;\n*** End Patch\n"
```

### 🔧 23. powershell

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Rebuild the Key Vault sample","initial_wait":180,"mode":"sync"}
```

### 🔧 24. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-tools-baseline-4047949117\\Program.cs"}
```

### 🔧 25. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-tools-baseline-4047949117\\README.md","view_range":[1,160]}
```

### 🔧 26. powershell

**Input:**

```json
{"command":"dotnet run --no-build --nologo","description":"Run the local dry-run path","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `KeyVaultProvisioningSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll create a compile-ready .NET sample with RBAC and access-policy alternatives, while keeping execution gated so no Azure resource is created here.

The implementation will use `WaitUntil.Started` followed by `WaitForCompletionAsync`, then optionally assign the built-in **Key Vault Secrets Officer** role because RBAC assignments are separate ARM resources and cannot be embedded in the vault creation payload.

The sample and authorization alternatives are in place; I’m compiling it now to catch SDK API or package-version mismatches without executing any Azure operation.

Implemented a compile-ready .NET 8 sample:

- `Program.cs` — authenticates with `DefaultAzureCredential`, creates an RBAC-enabled Key Vault in `eastus`, configures 90-day soft delete and purge protection, handles the `ArmOperation` LRO, optionally assigns **Key Vault Secrets Officer**, and constructs a `SecretClient`.
- `KeyVaultProvisioningSample.csproj` — includes all required Azure SDK packages.
- `README.md` — documents setup, execution, permissions, and the alternative legacy access-policy configuration.

Execution is protected by an explicit `--execute` flag; without it, the program exits locally without contacting Azure.

## Reviewer Notes (LLM-as-Judge)

### Summary

7/8 criteria passed. The sample correctly demonstrates management-plane creation, RBAC configuration, and explicit ArmOperation completion, but omits existing-vault and soft-deleted-vault recovery handling.

### Strengths

- Correctly uses CreateOrUpdateAsync with ArmOperation<KeyVaultResource> and explicit WaitForCompletionAsync.
- Configures RBAC authorization, soft delete, purge protection, location, tenant ID, and optional principal role assignment.
- Documents both RBAC and legacy access-policy authorization models and lists the required NuGet packages.

### Issues

- No RequestFailedException handling distinguishes an existing vault from a soft-deleted vault name conflict, and no recovery or purge workflow is shown.

## Grader Results

- polling-create-vault.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `KeyVaultCollection.CreateOrUpdateAsync()` returning `ArmOperation<KeyVaultResource>`: Pass
      - `KeyVaultCreateOrUpdateContent` with `KeyVaultProperties`: Pass
      - Configuring `EnableRbacAuthorization`, `EnableSoftDelete`, `EnablePurgeProtection`: Pass
      - `VaultAccessPolicy` vs RBAC authorization model: Pass
      - `ArmOperation<T>.WaitForCompletionAsync()` for completion: Pass
      - `WaitUntil.Completed` vs `WaitUntil.Started`: Pass
      - Tenant ID and object ID configuration: Pass
      - Error handling for existing vaults and soft-deleted vaults: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-mp-dotnet-polling --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)

# Evaluation Report: key-vault-mp-dotnet-polling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 169.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-mp-dotnet-polling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 169.0s |
| Timestamp | 2026-08-28T18:22:19Z |
| Files Generated | 3 |
| Event Count | 4982 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 150.9s |
| Review | 16.9s |
| **Total** | **169.0s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-security-keyvault-keys-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=539158 out=7704 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 686ms

**Input:**

```json
{"skill":"azure-security-keyvault-keys-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure.Security.KeyVault.Keys (.NET)

Client library for managing cryptographic keys in Azure Key Vault and Managed HSM.

## Installation

```bash
dotnet add package Azure.Security.KeyVault.Keys
dotnet add package Azure.Identity
```

**Current Version**: 4.7.0 (stable)

## Environment Variables

```bash
KEY_VAULT_NAME=<your-key-vault-name>  # Required: Key Vault name
AZURE_KEYVAULT_URL=https://<vault-name>.vault.azure.net  # Optional: full Key Vault URL
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
```

## Client Hierarchy

```
KeyClient (key management)
├── CreateKey / CreateRsaKey / CreateEcKey
├── GetKey / GetKeys
├── UpdateKeyProperties
├── DeleteKey / PurgeDeletedKey
├── BackupKey / RestoreKey
└── GetCryptographyClient() → CryptographyClient

CryptographyClient (cryptographic operations)
├── Encrypt / Decrypt
├── WrapKey / UnwrapKey
├── Sign / Verify
└── SignData / VerifyData

KeyResolver (key resolution)
└── Resolve(keyId) → CryptographyClient
```

## Authentication

### Microsoft Entra Token Credential

```csharp
using Azure.Identity;
using Azure.Security.KeyVault.Keys;

var keyVaultName = Environment.GetEnvironmentVariable("KEY_VAULT_NAME");
var kvUri = $"https://{keyVaultName}.vault.azure.net";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
var credential = new DefaultAzureCredential(
    DefaultAzureCredential.DefaultEnvironmentVariableName
);
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet#credential-classes
// var credential = new ManagedIdentityCredential();
var client = new KeyClient(new Uri(kvUri), credential);
```

### Service Principal

```csharp
var credential = new ClientSecretCredential(
    tenantId: "<tenant-id>
... (truncated)
```

### ❌ 2. view — 3299ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 3297ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 3294ms

**Input:**

```json
{"intent":"Generate a local-only C# example using Azure.ResourceManager.KeyVault, DefaultAzureCredential, ArmOperation polling, RBAC-enabled vault creation, soft delete, purge protection, role assignment or access-policy alternatives, and SecretClient construction.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 3614ms

**Input:**

```json
{"intent":"Find current official .NET Azure.ResourceManager.KeyVault APIs for creating a vault, KeyVaultCreateOrUpdateContent/KeyVaultProperties, ArmOperation WaitUntil.Started and WaitForCompletionAsync, access policies, and RBAC authorization.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. glob — 4197ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. glob — 3864ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-dotnet\skills\azure-security-keyvault-keys-dotnet\SKILL.md
```

### ✅ 8. powershell — 4922ms

**Input:**

```json
{"command":"dotnet --version","description":"Check installed .NET SDK version"}
```

**Output:**

```
10.0.111
<shellId: 0 completed with exit code 0>
```

### ✅ 9. azure-get_azure_bestpractices (via azure) — 4566ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code generation best practices for a C# Key Vault management SDK sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 5650ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official management SDK documentation for C# Key Vault creation and ArmOperation polling.","parameters":{"query":"Azure.ResourceManager.KeyVault C# KeyVaultCollection CreateOrUpdateAsync WaitUntil.Started ArmOperation KeyVaultCreateOrUpdateContent EnableRbacAuthorization EnablePurgeProtection SoftDeleteRetentionInDays access policies"}}
```

**Output:**

```
{"results":[{"title":"KeyVaultSecretCollection.CreateOrUpdateAsync Method-Definition","content":"# KeyVaultSecretCollection.CreateOrUpdateAsync Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure.ResourceManager.KeyVault](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.keyvault?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.ResourceManager.KeyVault.dll\r\n\r\n- Package:\r\n    - Azure.ResourceManager.KeyVault v1.4.0\r\n\r\n- Package:\r\n    - Azure.ResourceManager.KeyVault v1.2.0-beta.2\r\n\r\n- Source:\r\n    - [KeyVaultSecretCollection.cs](https://github.com/Azure/azure-sdk-for-net/blob/75a7ef8ab6eae983be2334991e4584a4f5f91d71/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/KeyVaultSecretCollection.cs)\r\n\r\n- Source:\r\n    - [KeyVaultSecretCollection.cs](https://github.com/Azure/azure-sdk-for-net/blob/6d2569b8458c5a311035d400204825724ff299ec/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/KeyVaultSecretCollection.cs)\r\n\r\nCreate or update a secret in a key vault in the specified subscription. NOTE: This API is intended for internal use in ARM deployments. Users should use the data-plane REST service for interaction with vault secrets.\r\n\r\n- ** Request Path. **: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}.\r\n- ** Operation Id. **: Secrets\\_CreateOrUpdate.\r\n- ** Default Api Version. **: 2026-02-01.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic virtual System.Threading.Tasks.Task\u003CAzure.ResourceManager.ArmOperation\u003CAzure.ResourceManager.KeyVault.KeyVaultSecretResource\u003E\u003E CreateOrUpdateAsync(Azure.WaitUntil waitUntil, string secretName, Azure.ResourceManager.KeyVault.Models.KeyVaultSecretCreateOrUpdateContent content, System.Threading.CancellationToken cancellationToken = default);\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- waitUntil\r\n    - [WaitUntil](https://learn.microsoft.com/dotnet/ap
... (truncated)
```

### ✅ 11. azure-documentation (via azure) — 5285ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for creating a Key Vault with Azure.ResourceManager.KeyVault.","parameters":{"language":"csharp","query":"Azure.ResourceManager.KeyVault create vault KeyVaultCreateOrUpdateContent ArmOperation WaitForCompletionAsync access policy RBAC role assignment"}}
```

**Output:**

```
{"results":[{"description":"description: Installs the Microsoft Azure Key Vault management library for .NET using the dotnet CLI and NuGet package manager.\npackage: Azure.ResourceManager.KeyVault\nlanguage: csharp\n","codeSnippet":"dotnet add package Azure.ResourceManager.KeyVault","link":"https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.keyvault-readme?view=azure-dotnet#getting-started","language":"csharp"},{"description":"description: Creates a role assignment scoped to a specific Key Vault using the Azure CLI, assigning the \u0027Key Vault Secrets Officer\u0027 role to a user principal.\npackage: azure-cli\nlanguage: azurecli\n","codeSnippet":"az role assignment create --role \u0022Key Vault Secrets Officer\u0022 --assignee \u003Cuser-principal-name\u003E --scope /subscriptions/\u003Csubscription-id\u003E/resourcegroups/\u003Cresource-group\u003E/providers/Microsoft.KeyVault/vaults/\u003Cvault-name\u003E","link":"https://learn.microsoft.com/azure/key-vault/general/rbac-guide#using-azure-rbac-secret-key-and-certificate-permissions-with-key-vault","language":"azurecli"},{"description":"description: Creates an Azure Key Vault with purge protection and RBAC authorization enabled, then retrieves the key vault ID for role assignment.\nlanguage: azurecli\n","codeSnippet":"az keyvault create --name $KVName --resource-group $resource_group_name --location $Location --enable-purge-protection --enable-rbac-authorization\r\n    \r\n    # retrieve the key vault ID for role assignment\r\n    $KVId=$(az keyvault show --resource-group $resource_group_name --name $KVName --query id --output tsv)","link":"https://learn.microsoft.com/azure/aks/aksarc/workload-identity#example-grant-permissions-to-access-azure-key-vault","language":"azurecli"},{"description":"description: Defines RBAC role IDs for Key Vault access, retrieves resource and principal IDs, and assigns Key Vault RBAC roles to a user-assigned managed identity and the current user to enable certificat
... (truncated)
```

### ✅ 12. rg — 3631ms

**Input:**

```json
{"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787941375536-copilot-tool-output-99b5861b8e9447ccb6f51cac9745b9ad.txt","pattern":"KeyVaultCreateOrUpdateContent|KeyVaultProperties|EnableRbacAuthorization|EnablePurgeProtection|SoftDeleteRetentionInDays|AccessPolicies|WaitForCompletionAsync|KeyVaultAccessPolicy"}
```

**Output:**

```
Output too large to read at once (68.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787941390091-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787941375536-copilot-tool-output-99b5861b8e9447ccb6f51cac9745b9ad.txt:{"results":[{"title":"KeyVaultSecretCollection.CreateOrUpdateAsync Method-Definition","content":"# KeyVaultSecretCollection.CreateOrUpdateAsync Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure.ResourceManager.KeyVault](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.keyvault?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.ResourceManager.KeyVault.dll\r\n\r\n-
```

### ❌ 13. web_fetch — 3637ms

**Input:**

```json
{"max_length":5000,"url":"https://api.nuget.org/v3-flatcontainer/azure.resourcemanager.keyvault/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ❌ 14. web_fetch — 3634ms

**Input:**

```json
{"max_length":5000,"url":"https://api.nuget.org/v3-flatcontainer/azure.resourcemanager.authorization/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ❌ 15. web_fetch — 3633ms

**Input:**

```json
{"max_length":5000,"url":"https://api.nuget.org/v3-flatcontainer/azure.identity/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ❌ 16. web_fetch — 3632ms

**Input:**

```json
{"max_length":5000,"url":"https://api.nuget.org/v3-flatcontainer/azure.security.keyvault.secrets/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ✅ 17. azure-documentation (via azure) — 4272ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the official Azure.ResourceManager.KeyVault .NET package overview and samples.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.keyvault-readme?view=azure-dotnet"}}
```

**Output:**

```
# KeyVaultProperties Class

## Definition

- Namespace:
    - [Azure.ResourceManager.KeyVault.Models](azure.resourcemanager.keyvault.models?view=azure-dotnet)

- Assembly:
    - Azure.ResourceManager.KeyVault.dll

- Package:
    - Azure.ResourceManager.KeyVault v1.4.0

- Package:
    - Azure.ResourceManager.KeyVault v1.2.0-beta.2

- Source:
    - [KeyVaultProperties.cs](https://github.com/Azure/azure-sdk-for-net/blob/75a7ef8ab6eae983be2334991e4584a4f5f91d71/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Customization/Models/KeyVaultProperties.cs)

- Source:
    - [KeyVaultProperties.cs](https://github.com/Azure/azure-sdk-for-net/blob/6d2569b8458c5a311035d400204825724ff299ec/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Customization/Models/KeyVaultProperties.cs)

Properties of the vault.

```csharp
public class KeyVaultProperties : System.ClientModel.Primitives.IJsonModel<Azure.ResourceManager.KeyVault.Models.KeyVaultProperties>, System.ClientModel.Primitives.IPersistableModel<Azure.ResourceManager.KeyVault.Models.KeyVaultProperties>
```

```csharp
public class KeyVaultProperties
```

```fsharp
type KeyVaultProperties = class
<output too long - dropped 39 lines from the middle>
| [EnablePurgeProtection](azure.resourcemanager.keyvault.models.keyvaultproperties.enablepurgeprotection?view=azure-dotnet#azure-resourcemanager-keyvault-models-keyvaultproperties-enablepurgeprotection) | Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value. |
| [EnableRbacAuthorization](azure.resourcemanager.keyvault.models.keyvaultproperties.enablerbacauthorization?view=azure-dotnet#azure-resourcem
... (truncated)
```

### ✅ 18. azure-documentation (via azure) — 3934ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the official KeyVaultCreateOrUpdateContent API documentation.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/azure.resourcemanager.keyvault.models.keyvaultcreateorupdatecontent?view=azure-dotnet"}}
```

**Output:**

```
# Azure Authorization management client library for .NET - version 1.1.7

The Microsoft Azure Authorization service provides authorization functionality as Azure core service which can serve for public and as well as for internal tenants.

This library supports managing Microsoft Azure Attestation resources.

This library follows the [new Azure SDK guidelines](https://azure.github.io/azure-sdk/general_introduction.html), and provides many core capabilities:

```
- Support MSAL.NET, Azure.Identity is out of box for supporting MSAL.NET.
- Support [OpenTelemetry](https://opentelemetry.io/) for distributed tracing.
- HTTP pipeline with custom policies.
- Better error-handling.
- Support uniform telemetry across all languages.
```

## Getting started

### Install the package

Install the Microsoft Azure Authorization management library for .NET with [NuGet](https://www.nuget.org/):

```dotnetcli
dotnet add package Azure.ResourceManager.Authorization
```

### Prerequisites

- You must have an [Microsoft Azure subscription](https://azure.microsoft.com/free/dotnet/).

### Authenticate the Client

To create an authenticated client and start interacting with Microsoft Azure resources, see the [quickstart guide here](https://github.com/Azure/azure-sdk-for-net/blob/Azure.ResourceManager.Authorization_1.1.7/doc/dev/mgmt_quickstart.md).

## Key concepts

Key concepts of the Microsoft Azure SDK for .NET can be found [here](https://azure.github.io/azure-sdk/dotnet_introduction.html).

## Documentation

Documentation is available to help you learn how to use this package:

- [Quickstart](https://github.com/Azure/azure-sdk-for-net/blob/Azure.ResourceManager.Authorization_1.1.7/doc/dev/mgmt_quickstart.md).
- [API References](/en-us/dotnet/api/?view=azure-dotnet).
- [Authentication](https://github.com/Azure/azure-sdk-for-net/blob/Azure.ResourceManager.Authorization_1.1.7/sdk/identity/Azure.Identity/README.md).

## Examples

Code samples f
... (truncated)
```

### ✅ 19. azure-documentation (via azure) — 3588ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the official KeyVaultProperties API documentation.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/azure.resourcemanager.keyvault.models.keyvaultproperties?view=azure-dotnet"}}
```

**Output:**

```
# KeyVaultCreateOrUpdateContent Class

## Definition

- Namespace:
    - [Azure.ResourceManager.KeyVault.Models](azure.resourcemanager.keyvault.models?view=azure-dotnet)

- Assembly:
    - Azure.ResourceManager.KeyVault.dll

- Package:
    - Azure.ResourceManager.KeyVault v1.4.0

- Package:
    - Azure.ResourceManager.KeyVault v1.2.0-beta.2

- Source:
    - [KeyVaultCreateOrUpdateContent.cs](https://github.com/Azure/azure-sdk-for-net/blob/75a7ef8ab6eae983be2334991e4584a4f5f91d71/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultCreateOrUpdateContent.cs)

- Source:
    - [KeyVaultCreateOrUpdateContent.cs](https://github.com/Azure/azure-sdk-for-net/blob/6d2569b8458c5a311035d400204825724ff299ec/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultCreateOrUpdateContent.cs)

Parameters for creating or updating a vault.

```csharp
public class KeyVaultCreateOrUpdateContent : System.ClientModel.Primitives.IJsonModel<Azure.ResourceManager.KeyVault.Models.KeyVaultCreateOrUpdateContent>, System.ClientModel.Primitives.IPersistableModel<Azure.ResourceManager.KeyVault.Models.KeyVaultCreateOrUpdateContent>
```

```csharp
public class KeyVaultCreateOrUpdateContent
```

```fsharp
type KeyVaultCreateOrUpdateContent = class
    interface IJsonModel<KeyVaultCreateOrUpdateContent>
    interface IPersistableModel<KeyVaultCreateOrUpdateContent>
```

```fsharp
type KeyVaultCreateOrUpdateContent = class
```

```vb
Public Class KeyVaultCreateOrUpdateContent
Implements IJsonModel(Of KeyVaultCreateOrUpdateContent), IPersistableModel(Of KeyVaultCreateOrUpdateContent)
```

```vb
Public Class KeyVaultCreateOrUpdateContent
```

- Inheritance
    - [Object](/en-us/dotnet/api/system.object)
KeyVaultCreateOrUpdateContent

- Implements
    - [IJsonModel](system.clientmodel.primitives.ijsonmodel-1?view=azure-dotnet)&lt;[KeyVaultCreateOrUpdateContent](azure.resourcemanager.keyvault.models.keyvaultcreateor
... (truncated)
```

### ✅ 20. azure-documentation (via azure) — 2885ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official .NET role assignment API documentation for Azure Resource Manager Authorization.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.authorization-readme?view=azure-dotnet"}}
```

**Output:**

```
# Microsoft Azure Key Vault management client library for .NET

Microsoft Azure Key Vault is one of several key management solutions in Azure, and helps solve the following problems:

- **Secrets Management** - Azure Key Vault can be used to Securely store and tightly control access to tokens, passwords, certificates, API keys, and other secrets
- **Key Management** - Azure Key Vault can be used as a Key Management solution. Azure Key Vault makes it easy to create and control the encryption keys used to encrypt your data.
- **Certificate Management** - Azure Key Vault lets you easily provision, manage, and deploy public and private Transport Layer Security/Secure Sockets Layer (TLS/SSL) certificates for use with Azure and your internal connected resources.

This library supports managing Microsoft Azure Key Vault resources.

This library follows the [new Azure SDK guidelines](https://azure.github.io/azure-sdk/general_introduction.html), and provides many core capabilities:

```
- Support MSAL.NET, Azure.Identity is out of box for supporting MSAL.NET.
- Support [OpenTelemetry](https://opentelemetry.io/) for distributed tracing.
- HTTP pipeline with custom policies.
- Better error-handling.
- Support uniform telemetry across all languages.
```

## Getting started

### Install the package

Install the Microsoft Azure Key Vault management library for .NET with [NuGet](https://www.nuget.org/):

```dotnetcli
dotnet add package Azure.ResourceManager.KeyVault
```

### Prerequisites

- You must have an [Azure subscription](https://azure.microsoft.com/free/dotnet/).

### Authenticate the Client

To create an authenticated client and start interacting with Microsoft Azure resources, see the [quickstart guide here](https://github.com/Azure/azure-sdk-for-net/blob/Azure.ResourceManager.KeyVault_1.4.0/doc/dev/mgmt_quickstart.md).

## Key concepts

Key concepts of the Microsoft Azure SDK for .NET can be found [here](https://azure.github.io/az
... (truncated)
```

### ✅ 21. powershell — 3229ms

**Input:**

```json
{"command":"$json = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787941375536-copilot-tool-output-99b5861b8e9447ccb6f51cac9745b9ad.txt' -Raw | ConvertFrom-Json; $json.results | Where-Object { $_.content -match 'KeyVaultCreateOrUpdateContent|KeyVaultProperties|EnableRbacAuthorization|EnablePurgeProtection|SoftDeleteRetentionInDays|KeyVaultAccessPolicy' } | Select-Object -First 10 title,content,contentUrl | ConvertTo-Json -Depth 4","description":"Extract relevant Key Vault documentation snippets"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … 99b5861b8e9447ccb6f51cac9745b9ad.txt' -Raw | ConvertFrom-Json; $json. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 22. github-mcp-server-search_code (via github-mcp-server) — 6976ms

**Input:**

```json
{"fields":["path","repository","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"class RoleAssignmentCreateOrUpdateContent\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"namespace Azure.ResourceManager.KeyVault.Models\n{\n    /// <summary> An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. </summary>\n    public partial class KeyVaultAccessPolicy\n    {\n        /// <summary> Keeps track of any properties unknown to the library. </summary>\n        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;","matches":[{"indices":[237,257],"text":"KeyVaultAccessPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/keyvault/Azure.Provisioning.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"namespace Azure.Provisioning.KeyVault\n{\n    /// <summary> An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. </summary>\n    public partial class KeyVaultAccessPolicy : ProvisionableConstruct\n    {\n        private BicepValue<Guid> _tenantId;\n        private BicepValue<string> _objectId;","matches":[{"indices":[227,247],"text":"KeyVaultAccessPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/keyvault/Azure.Provisioning.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.Serialization.cs","repository":"Azure/azure-sdk-for-net","text_matches
... (truncated)
```

### ✅ 23. github-mcp-server-search_code (via github-mcp-server) — 6652ms

**Input:**

```json
{"fields":["path","repository","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"GetRoleAssignments(ResourceIdentifier\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Custom/Mocking/MockableAuthorizationArmResource.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"            return await GetDenyAssignments().GetAsync(denyAssignmentId, cancellationToken).ConfigureAwait(false);\n        }\n\n        /// <inheritdoc cref=\"MockableAuthorizationArmClient.GetRoleAssignments(ResourceIdentifier)\"/>\n        public virtual RoleAssignmentCollection GetRoleAssignments()\n        {\n            return GetCachedClient(client => new RoleAssignmentCollection(client, Id));","matches":[{"indices":[187,224],"text":"GetRoleAssignments(ResourceIdentifier"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/authorization/Azure.ResourceManager.Authorization/src/Custom/Mocking/MockableAuthorizationArmResource.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/Extensions/AuthorizationExtensions.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"        /// Gets a collection of <see cref=\"RoleAssignmentCollection\"/> objects within the specified scope.\n        /// <item>\n        /// <term> Mocking. </term>\n        /// <description> To mock this method, please mock <see cref=\"MockableAuthorizationArmClient.GetRoleAssignments(ResourceIdentifier)\"/> instead. </description>\n        /// </item>\n        /// </summary>\n        /// <param name=\"client\"> The <see cref=\"ArmClient\"/> the method will execute against. </param>","matches":[{"indices":[264,301],"text":"GetRoleAssignments(ResourceIdentifier"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/Extensions/AuthorizationExtensions.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"
... (truncated)
```

### ✅ 24. github-mcp-server-search_code (via github-mcp-server) — 6327ms

**Input:**

```json
{"fields":["path","repository","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"class KeyVaultAccessPolicy\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/keyvault/Azure.Provisioning.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"namespace Azure.Provisioning.KeyVault\n{\n    /// <summary> An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. </summary>\n    public partial class KeyVaultAccessPolicy : ProvisionableConstruct\n    {\n        private BicepValue<Guid> _tenantId;\n        private BicepValue<string> _objectId;","matches":[{"indices":[221,247],"text":"class KeyVaultAccessPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/keyvault/Azure.Provisioning.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"namespace Azure.ResourceManager.KeyVault.Models\n{\n    /// <summary> An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. </summary>\n    public partial class KeyVaultAccessPolicy\n    {\n        /// <summary> Keeps track of any properties unknown to the library. </summary>\n        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;","matches":[{"indices":[231,257],"text":"class KeyVaultAccessPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/keyvault/Azure.Provisioning.KeyVault/api/Azure.Provisioning.KeyVault.net8.0.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fra
... (truncated)
```

### ✅ 25. github-mcp-server-search_code (via github-mcp-server) — 6008ms

**Input:**

```json
{"fields":["path","repository","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"IdentityAccessPermissions\" KeyVaultAccessPolicy"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Custom/Models/RoleAssignmentCreateOrUpdateContent.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"\nnamespace Azure.ResourceManager.Authorization.Models\n{\n    public partial class RoleAssignmentCreateOrUpdateContent\n    {\n        // The generator emits getter-only flattened properties even though the generated inner properties are settable.\n        // These forwarding setters restore the shipped mutable surface without changing inner storage or serialization: https://github.com/Azure/azure-sdk-for-net/issues/61114.","matches":[{"indices":[75,116],"text":"class RoleAssignmentCreateOrUpdateContent"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/authorization/Azure.ResourceManager.Authorization/src/Custom/Models/RoleAssignmentCreateOrUpdateContent.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/Models/RoleAssignmentCreateOrUpdateContent.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"namespace Azure.ResourceManager.Authorization.Models\n{\n    /// <summary> Role assignment create parameters. </summary>\n    public partial class RoleAssignmentCreateOrUpdateContent\n    {\n        /// <summary> Keeps track of any properties unknown to the library. </summary>\n        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;","matches":[{"indices":[138,179],"text":"class RoleAssignmentCreateOrUpdateContent"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/Models/RoleAssignmentCreateOrUpdateContent.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"path":"sdk/authorization/Azure.ResourceManager.
... (truncated)
```

### ✅ 26. azure-documentation (via azure) — 5681ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official .NET API signatures for vault-scoped role assignment creation.","parameters":{"query":"Azure.ResourceManager.Authorization RoleAssignmentCollection ArmClient GetRoleAssignments ResourceIdentifier RoleAssignmentCreateOrUpdateContent constructor C#"}}
```

**Output:**

```
{"results":[{"title":"RoleAssignmentCollection.CreateOrUpdate Method-Definition","content":"# RoleAssignmentCollection.CreateOrUpdate Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure.ResourceManager.Authorization](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.authorization?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.ResourceManager.Authorization.dll\r\n\r\n- Package:\r\n    - Azure.ResourceManager.Authorization v1.1.7\r\n\r\n- Package:\r\n    - Azure.ResourceManager.Authorization v1.2.0-beta.1\r\n\r\n- Source:\r\n    - [RoleAssignmentCollection.cs](https://github.com/Azure/azure-sdk-for-net/blob/f1ce8e8a389e35c76d1c108ec29b06e58ad7960c/sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/RoleAssignmentCollection.cs#L123C13-L143C10)\r\n\r\n- Source:\r\n    - [RoleAssignmentCollection.cs](https://github.com/Azure/azure-sdk-for-net/blob/c97bde6212bedc476d9a8977711310c2c3cc1296/sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/RoleAssignmentCollection.cs#L126C13-L154C10)\r\n\r\nCreate or update a role assignment by scope and name.\r\n\r\n- ** Request Path. **: /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.\r\n- ** Operation Id. **: RoleAssignments\\_Create.\r\n- ** Default Api Version. **: 2022-04-01.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic virtual Azure.ResourceManager.ArmOperation\u003CAzure.ResourceManager.Authorization.RoleAssignmentResource\u003E CreateOrUpdate(Azure.WaitUntil waitUntil, string roleAssignmentName, Azure.ResourceManager.Authorization.Models.RoleAssignmentCreateOrUpdateContent content, System.Threading.CancellationToken cancellationToken = default);\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- waitUntil\r\n    - [WaitUntil](https://learn.microsoft.com/dotnet/api/azure.waituntil?view=azure-dotnet)\r\n\r\n[Completed](https://learn.microsoft.com/dotnet/api/azure.waituntil?view=azure-dotnet#azure-waituntil-completed) if the method should w
... (truncated)
```

### ✅ 27. github-mcp-server-get_file_contents (via github-mcp-server) — 9220ms

**Input:**

```json
{"owner":"Azure","path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/Models/RoleAssignmentCreateOrUpdateContent.cs","ref":"refs/tags/Azure.ResourceManager.Authorization_1.1.7","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: c97b87a825bb64ecdb8403e1b04ad7d20c3372ca)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using Azure.ResourceManager.KeyVault;

namespace Azure.ResourceManager.KeyVault.Models
{
    /// <summary> SKU details. </summary>
    public partial class KeyVaultSku
    {
        /// <summary> Keeps track of any properties unknown to the library. </summary>
        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;

        /// <summary> Initializes a new instance of <see cref="KeyVaultSku"/>. </summary>
        /// <param name="family"> SKU family name. </param>
        /// <param name="name"> SKU name to specify whether the key vault is a standard vault or a premium vault. </param>
        public KeyVaultSku(KeyVaultSkuFamily family, KeyVaultSkuName name)
        {
            Family = family;
            Name = name;
        }

        /// <summary> Initializes a new instance of <see cref="KeyVaultSku"/>. </summary>
        /// <param name="family"> SKU family name. </param>
        /// <param name="name"> SKU name to specify whether the key vault is a standard vault or a premium vault. </param>
        /// <param name="additionalBinaryDataProperties"> Keeps track of any properties unknown to the library. </param>
        internal KeyVaultSku(KeyVaultSkuFamily family, KeyVaultSkuName name, IDictionary<string, BinaryData> additionalBinaryDataProperties)
        {
            Family = family;
            Name = name;
            _additionalBinaryDataProperties = additionalBinaryDataProperties;
        }

        /// <summary> SKU family name. </summary>
        [WirePath("family")]
        public KeyVaultSkuFamily Family { get; set; }

        /// <summary> SKU name to specify whether the key vault is a standard vault or a premium vault. </summary>
        [Wi
... (truncated)
```

### ✅ 28. github-mcp-server-get_file_contents (via github-mcp-server) — 8913ms

**Input:**

```json
{"owner":"Azure","path":"sdk/authorization/Azure.ResourceManager.Authorization/src/Generated/Extensions/AuthorizationExtensions.cs","ref":"refs/tags/Azure.ResourceManager.Authorization_1.1.7","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 4892fce514532150124ee8208640d1465e727d90)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using Azure.ResourceManager.KeyVault;

namespace Azure.ResourceManager.KeyVault.Models
{
    /// <summary> Permissions the identity has for keys, secrets, certificates and storage. </summary>
    public partial class IdentityAccessPermissions
    {
        /// <summary> Keeps track of any properties unknown to the library. </summary>
        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;

        /// <summary> Initializes a new instance of <see cref="IdentityAccessPermissions"/>. </summary>
        public IdentityAccessPermissions()
        {
            Keys = new ChangeTrackingList<IdentityAccessKeyPermission>();
            Secrets = new ChangeTrackingList<IdentityAccessSecretPermission>();
            Certificates = new ChangeTrackingList<IdentityAccessCertificatePermission>();
            Storage = new ChangeTrackingList<IdentityAccessStoragePermission>();
        }

        /// <summary> Initializes a new instance of <see cref="IdentityAccessPermissions"/>. </summary>
        /// <param name="keys"> Permissions to keys. </param>
        /// <param name="secrets"> Permissions to secrets. </param>
        /// <param name="certificates"> Permissions to certificates. </param>
        /// <param name="storage"> Permissions to storage accounts. </param>
        /// <param name="additionalBinaryDataProperties"> Keeps track of any properties unknown to the library. </param>
        internal IdentityAccessPermissions(IList<IdentityAccessKeyPermission> keys, IList<IdentityAccessSecretPermission> secrets, IList<IdentityAccessCertificatePermission> certificates, IList<IdentityAccessStoragePermission> storage, IDictionary<string, BinaryData> additionalBinaryDataPrope
... (truncated)
```

### ✅ 29. github-mcp-server-get_file_contents (via github-mcp-server) — 8601ms

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultAccessPolicy.cs","ref":"refs/tags/Azure.ResourceManager.KeyVault_1.4.0","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: e85a2f63b2ecbbf4d875fb797ee79c23c0fe1d2c)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core;
using Azure.ResourceManager.Authorization.Mocking;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Resources;

namespace Azure.ResourceManager.Authorization
{
    /// <summary> A class to add extension methods to Azure.ResourceManager.Authorization. </summary>
    public static partial class AuthorizationExtensions
    {
        private static MockableAuthorizationArmClient GetMockableAuthorizationArmClient(ArmClient client)
        {
            return client.GetCachedClient(client0 => new MockableAuthorizationArmClient(client0));
        }

        private static MockableAuthorizationArmResource GetMockableAuthorizationArmResource(ArmResource resource)
        {
            return resource.GetCachedClient(client => new MockableAuthorizationArmResource(client, resource.Id));
        }

        private static MockableAuthorizationResourceGroupResource GetMockableAuthorizationResourceGroupResource(ArmResource resource)
        {
            return resource.GetCachedClient(client => new MockableAuthorizationResourceGroupResource(client, resource.Id));
        }

        private static MockableAuthorizationSubscriptionResource GetMockableAuthorizationSubscriptionResource(ArmResource resource)
        {
            return resource.GetCachedClient(client => new MockableAuthorizationSubscriptionResource(client, resource.Id));
        }

        private static MockableAuthorizationTenantResource GetMockableAuthorizationTenantResource(ArmResource resource)
        {
            return resource.GetCachedClient(client => new MockableAuthorizationTenantResource(client, resource.Id));
        }

        /// <summary>
        /// Gets a collection of
... (truncated)
```

### ✅ 30. github-mcp-server-get_file_contents (via github-mcp-server) — 8273ms

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/IdentityAccessPermissions.cs","ref":"refs/tags/Azure.ResourceManager.KeyVault_1.4.0","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 780f7abf0582349e302c3e1a209f19cc1d448fe9)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using Azure.ResourceManager.KeyVault;

namespace Azure.ResourceManager.KeyVault.Models
{
    /// <summary> An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. </summary>
    public partial class KeyVaultAccessPolicy
    {
        /// <summary> Keeps track of any properties unknown to the library. </summary>
        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;

        /// <summary> Initializes a new instance of <see cref="KeyVaultAccessPolicy"/>. </summary>
        /// <param name="tenantId"> The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault. </param>
        /// <param name="objectId"> The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies. </param>
        /// <param name="permissions"> Permissions the identity has for keys, secrets and certificates. </param>
        /// <exception cref="ArgumentNullException"> <paramref name="objectId"/> or <paramref name="permissions"/> is null. </exception>
        public KeyVaultAccessPolicy(Guid tenantId, string objectId, IdentityAccessPermissions permissions)
        {
            Argument.AssertNotNull(objectId, nameof(objectId));
            Argument.AssertNotNull(permissions, nameof(permissions));

            TenantId = tenantId;
            ObjectId = objectId;
            Permissions = permissions;
        }

        /// <summary> Initializes a new instance of <see cref="KeyVaultAccessPolicy"/>. </summary>
        /// <param name="tenantId"> The Azure Active Direc
... (truncated)
```

### ✅ 31. github-mcp-server-get_file_contents (via github-mcp-server) — 7958ms

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/Azure.ResourceManager.KeyVault/src/Generated/Models/KeyVaultSku.cs","ref":"refs/tags/Azure.ResourceManager.KeyVault_1.4.0","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 64e3594bd5a020d2255e0e1f46f91a9f41de66fc)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Authorization.Models
{
    /// <summary> Role assignment create parameters. </summary>
    public partial class RoleAssignmentCreateOrUpdateContent
    {
        /// <summary>
        /// Keeps track of any properties unknown to the library.
        /// <para>
        /// To assign an object to the value of this property use <see cref="BinaryData.FromObjectAsJson{T}(T, System.Text.Json.JsonSerializerOptions?)"/>.
        /// </para>
        /// <para>
        /// To assign an already formatted json string to this property use <see cref="BinaryData.FromString(string)"/>.
        /// </para>
        /// <para>
        /// Examples:
        /// <list type="bullet">
        /// <item>
        /// <term>BinaryData.FromObjectAsJson("foo")</term>
        /// <description>Creates a payload of "foo".</description>
        /// </item>
        /// <item>
        /// <term>BinaryData.FromString("\"foo\"")</term>
        /// <description>Creates a payload of "foo".</description>
        /// </item>
        /// <item>
        /// <term>BinaryData.FromObjectAsJson(new { key = "value" })</term>
        /// <description>Creates a payload of { "key": "value" }.</description>
        /// </item>
        /// <item>
        /// <term>BinaryData.FromString("{\"key\": \"value\"}")</term>
        /// <description>Creates a payload of { "key": "value" }.</description>
        /// </item>
        /// </list>
        /// </para>
        /// </summary>
        private IDictionary<string, BinaryData> _serializedAdditionalRawData;

        /// <summary> Initializes a new instance of <see cref="RoleAssignmentCreateOrUpdateContent"/>. </summary>
        /// <param name="roleDefinitionI
... (truncated)
```

### ✅ 32. apply_patch — 670ms

**Input:**

```json
"*** Begin Patch\n*** Add File: KeyVaultManagementSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Identity\" Version=\"1.17.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.ResourceManager.Authorization\" Version=\"1.1.7\" /\u003e\n+    \u003cPackageReference Include=\"Azure.ResourceManager.KeyVault\" Version=\"1.4.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.Security.KeyVault.Secrets\" Version=\"4.8.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using System.Security.Cryptography;\n+using System.Text;\n+using Azure;\n+using Azure.Core;\n+using Azure.Identity;\n+using Azure.ResourceManager;\n+using Azure.ResourceManager.Authorization;\n+using Azure.ResourceManager.Authorization.Models;\n+using Azure.ResourceManager.KeyVault;\n+using Azure.ResourceManager.KeyVault.Models;\n+using Azure.ResourceManager.Resources;\n+using Azure.Security.KeyVault.Secrets;\n+\n+const string keyVaultSecretsUserRoleId = \"4633458b-17de-408a-b874-0445c86b69e6\";\n+\n+string subscriptionId = GetRequiredEnvironmentVariable(\"AZURE_SUBSCRIPTION_ID\");\n+Guid tenantId = Guid.Parse(GetRequiredEnvironmentVariable(\"AZURE_TENANT_ID\"));\n+string resourceGroupName = GetRequiredEnvironmentVariable(\"AZURE_RESOURCE_GROUP\");\n+string vaultName = GetRequiredEnvironmentVariable(\"KEY_VAULT_NAME\");\n+Guid principalObjectId = Guid.Parse(\n+    GetRequiredEnvironmentVariable(\"KEY_VAULT_PRINCIPAL_OBJECT_ID\"));\n+\n+var credential = new DefaultAzureCredential();\n+var armClient = new ArmClient(credential, subscriptionId);\n+\n+SubscriptionResource subscription = armClient.GetSubscriptionResource(\n+    SubscriptionResource.CreateResourceIdentifier(subscriptionId));\n+ResourceGroupResource resourceGroup =\n+    (await subscription.GetResourceGroupAsync(resourceGroupName)).Value;\n+\n+var properties = new KeyVaultProperties(\n+    tenantId,\n+    new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))\n+{\n+    EnableRbacAuthorization = true,\n+    EnableSoftDelete = true,\n+    SoftDeleteRetentionInDays = 90,\n+    EnablePurgeProtection = true\n+};\n+\n+var createContent = new KeyVaultCreateOrUpdateContent(\n+    AzureLocation.EastUS,\n+    properties);\n+\n+KeyVaultCollection vaults = resourceGroup.GetKeyVaults();\n+\n+// WaitUntil.Started returns after Azure accepts the request, leaving polling to ArmOperation.\n+ArmOperation\u003cKeyVaultResource\u003e createOperation = await vaults.CreateOrUpdateAsync(\n+    WaitUntil.Started,\n+    vaultName,\n+    createContent);\n+\n+Console.WriteLine($\"Vault creation started. Operation ID: {createOperation.Id}\");\n+\n+await createOperation.WaitForCompletionAsync();\n+KeyVaultResource vault = createOperation.Value;\n+\n+Console.WriteLine($\"Vault creation completed: {vault.Id}\");\n+\n+// Data-plane RBAC assignments are separate ARM resources, so the vault must exist first.\n+ResourceIdentifier roleDefinitionId = new(\n+    $\"/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/\" +\n+    $\"roleDefinitions/{keyVaultSecretsUserRoleId}\");\n+\n+var roleAssignmentContent = new RoleAssignmentCreateOrUpdateContent(\n+    roleDefinitionId,\n+    principalObjectId)\n+{\n+    Description = \"Allows this sample's identity to read Key Vault secrets.\"\n+};\n+\n+Guid roleAssignmentName = CreateDeterministicGuid(\n+    $\"{vault.Id}|{principalObjectId}|{keyVaultSecretsUserRoleId}\");\n+\n+ArmOperation\u003cRoleAssignmentResource\u003e roleAssignmentOperation =\n+    await armClient.GetRoleAssignments(vault.Id).CreateOrUpdateAsync(\n+        WaitUntil.Completed,\n+        roleAssignmentName.ToString(),\n+        roleAssignmentContent);\n+\n+Console.WriteLine($\"Role assignment ready: {roleAssignmentOperation.Value.Id}\");\n+\n+Uri vaultUri = vault.Data.Properties.VaultUri\n+    ?? throw new InvalidOperationException(\"Azure did not return the vault URI.\");\n+var secretClient = new SecretClient(vaultUri, credential);\n+\n+// A client can be constructed before RBAC propagation finishes. Make a harmless list request\n+// and retry only the expected temporary 403 response.\n+await VerifySecretClientAccessAsync(secretClient);\n+Console.WriteLine($\"SecretClient can access {secretClient.VaultUri}\");\n+\n+static async Task VerifySecretClientAccessAsync(\n+    SecretClient client,\n+    CancellationToken cancellationToken = default)\n+{\n+    const int maximumAttempts = 10;\n+\n+    for (int attempt = 1; ; attempt++)\n+    {\n+        try\n+        {\n+            await using IAsyncEnumerator\u003cSecretProperties\u003e secrets =\n+                client.GetPropertiesOfSecretsAsync(cancellationToken)\n+                    .GetAsyncEnumerator(cancellationToken);\n+\n+            _ = await secrets.MoveNextAsync();\n+            return;\n+        }\n+        catch (RequestFailedException exception)\n+            when (exception.Status == 403 \u0026\u0026 attempt \u003c maximumAttempts)\n+        {\n+            TimeSpan delay = TimeSpan.FromSeconds(Math.Min(60, attempt * 10));\n+            Console.WriteLine(\n+                $\"Waiting for RBAC propagation (attempt {attempt}/{maximumAttempts})...\");\n+            await Task.Delay(delay, cancellationToken);\n+        }\n+    }\n+}\n+\n+static Guid CreateDeterministicGuid(string value)\n+{\n+    byte[] hash = SHA256.HashData(Encoding.UTF8.GetBytes(value));\n+    return new Guid(hash.AsSpan(0, 16));\n+}\n+\n+static string GetRequiredEnvironmentVariable(string name) =\u003e\n+    Environment.GetEnvironmentVariable(name) is { Length: \u003e 0 } value\n+        ? value\n+        : throw new InvalidOperationException(\n+            $\"Set the required environment variable {name}.\");\n+\n+// Access-policy alternative (do not combine this with EnableRbacAuthorization = true):\n+//\n+// var permissions = new IdentityAccessPermissions();\n+// permissions.Secrets.Add(IdentityAccessSecretPermission.Get);\n+// permissions.Secrets.Add(IdentityAccessSecretPermission.List);\n+//\n+// var accessPolicyProperties = new KeyVaultProperties(\n+//     tenantId,\n+//     new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))\n+// {\n+//     EnableRbacAuthorization = false,\n+//     EnableSoftDelete = true,\n+//     SoftDeleteRetentionInDays = 90,\n+//     EnablePurgeProtection = true\n+// };\n+// accessPolicyProperties.AccessPolicies.Add(\n+//     new KeyVaultAccessPolicy(\n+//         tenantId,\n+//         principalObjectId.ToString(),\n+//         permissions));\n+//\n+// Pass accessPolicyProperties to KeyVaultCreateOrUpdateContent. Unlike an RBAC role\n+// assignment, this access policy is included directly in the vault creation request.\n*** Add File: README.md\n+# Azure Key Vault management SDK sample\n+\n+This sample creates an RBAC-enabled Key Vault in `eastus`, explicitly enables\n+soft delete with 90-day retention and purge protection, waits for the\n+management-plane long-running operation, creates a vault-scoped **Key Vault\n+Secrets User** role assignment, and verifies data-plane access with\n+`SecretClient`.\n+\n+## Packages\n+\n+```powershell\n+dotnet add package Azure.Identity --version 1.17.0\n+dotnet add package Azure.ResourceManager.Authorization --version 1.1.7\n+dotnet add package Azure.ResourceManager.KeyVault --version 1.4.0\n+dotnet add package Azure.Security.KeyVault.Secrets --version 4.8.0\n+```\n+\n+## Configuration\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:AZURE_TENANT_ID = \"\u003ctenant-id\u003e\"\n+$env:AZURE_RESOURCE_GROUP = \"\u003cexisting-resource-group\u003e\"\n+$env:KEY_VAULT_NAME = \"\u003cglobally-unique-vault-name\u003e\"\n+$env:KEY_VAULT_PRINCIPAL_OBJECT_ID = \"\u003cobject-id-used-by-DefaultAzureCredential\u003e\"\n+dotnet run\n+```\n+\n+`DefaultAzureCredential` uses your developer credential locally and managed\n+identity when Azure-hosted. The principal must be able to create vaults and\n+role assignments at the chosen scope. `KEY_VAULT_PRINCIPAL_OBJECT_ID` is the\n+Microsoft Entra object ID, not an application/client ID.\n+\n+An RBAC role assignment cannot be embedded in Key Vault's create payload\n+because its scope is the completed vault resource. The program therefore\n+waits for vault creation and then creates the role assignment. The commented\n+alternative in `Program.cs` shows legacy access-policy authorization, which\n+can be embedded during creation but is ignored when\n+`EnableRbacAuthorization` is `true`.\n+\n+## References\n+\n+- [Azure Key Vault management library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.keyvault-readme?view=azure-dotnet)\n+- [Azure Key Vault RBAC guide](https://learn.microsoft.com/azure/key-vault/general/rbac-guide)\n+- [DefaultAzureCredential documentation](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788/KeyVaultManagementSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788/KeyVaultManagementSample.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788/KeyVaultManagementSample.csproj
@@ -1,0 +1,18 @@
+<Project Sdk="Microsoft.NET.Sdk">
+
+  <PropertyGroup>
+    <OutputType>Exe</OutputType>
+    <TargetFramework>net8.0</TargetFramework>
+    <ImplicitUsings>enable</ImplicitUsings>
+    <Nullable>enable</Nullable>
+  </PropertyGroup>
+
+  <ItemGroup>
+    <PackageReference Include="Azure.Identity" Version="1.17.0" />
+    <PackageReference Include="Azure.ResourceManager.Authorization" Version="1.1.7" />
+    <PackageReference Include="Azure.ResourceManager.KeyVault" Version="1.4.0" />
+    <PackageReference Include="Azure.Security.KeyVault.Secrets" Version="4.8.0" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788/Program.cs
@@ -1,0 +1,155 @@
+using System.Security.Cryptography;
+using System.Text;
+using Azure;
+using Azure.Core;
+using Azure.Identity;
+using Azure.ResourceManager;
+using Azure.ResourceManager.Authorization;
+using Azure.ResourceManager.Authorization.Models;
+using Azure.ResourceManager.
... (truncated)
```

### ✅ 33. powershell — 2449ms

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Restore packages and compile Key Vault sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\Program.cs(23,22): error CS0433: The type 'DefaultAzureCredential' exists in both 'Azure.Core, Version=1.57.0.0, Culture=neutral, PublicKeyToken=92742159e12e44c8' and 'Azure.Identity, Version=1.17.0.0, Culture=neutral, PublicKeyToken=92742159e12e44c8' [C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\KeyVaultManagementSample.csproj]
C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\Program.cs(102,42): error CS0104: 'SecretProperties' is an ambiguous reference between 'Azure.ResourceManager.KeyVault.Models.SecretProperties' and 'Azure.Security.KeyVault.Secrets.SecretProperties' [C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\KeyVaultManagementSample.csproj]

Build FAILED.

C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\Program.cs(23,22): error CS0433: The type 'DefaultAzureCredential' exists in both 'Azure.Core, Version=1.57.0.0, Culture=neutral, PublicKeyToken=92742159e12e44c8' and 'Azure.Identity, Version=1.17.0.0, Culture=neutral, PublicKeyToken=92742159e12e44c8' [C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\KeyVaultManagementSample.csproj]
C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-dotnet-polling-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1606642788\Program.cs(102,42): error CS0104: 'SecretProperties' is an ambiguous reference between 'Azure.ResourceManager.KeyVault.Models.SecretProperties' and 'Azure.Security.KeyVault.Secrets.SecretProperties' [C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-mp-d
... (truncated)
```

## Generated Files

- `KeyVaultManagementSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll ground the sample in the current management SDK’s creation and polling APIs, then compile it locally without contacting Azure.

The key authorization detail is that an RBAC-enabled vault ignores access policies; the sample will create the vault first, wait for the LRO, then create a vault-scoped role assignment. It will also include the mutually exclusive access-policy creation variant.

The first compile exposed a package-generation compatibility issue: the selected Identity package duplicates `DefaultAzureCredential` now present in the resolved Azure.Core. I’m aligning the current package set and removing the model-name ambiguity.

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The sample thoroughly implements creation, LRO completion, RBAC assignment, and data-plane verification, but omits the requested existing-vault and soft-deleted-vault error handling.

### Strengths

- Correctly uses the management-plane ArmOperation<KeyVaultResource> pattern.
- Clearly distinguishes RBAC role assignments from legacy access policies.
- Configures all requested vault security properties and verifies real SecretClient access with RBAC propagation retries.

### Issues

- No targeted handling for existing vault conflicts or vault names retained in the soft-deleted state.
- The final response reports unresolved compilation work rather than clearly presenting the completed sample outcome.

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
hyoka run --prompt-id key-vault-mp-dotnet-polling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

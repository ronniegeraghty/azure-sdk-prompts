# Evaluation Report: identity-dp-dotnet-managed-identity

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 153.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-managed-identity` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 153.0s |
| Timestamp | 2026-08-26T17:39:44Z |
| Files Generated | 0 |
| Event Count | 4978 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 122.7s |
| Review | 29.2s |
| **Total** | **153.0s** |

## Configuration

- **name:** dotnet-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=423142 out=4173 |
| Turn Count | 10 |

## Prompt Sent

```
Show me how to use
Managed Identity to authenticate Azure SDK clients in C#. Cover:
1. System-assigned vs user-assigned managed identity differences
2. How to create a ManagedIdentityCredential for each type
3. Using it with an Azure SDK client (e.g., KeyVaultClient or BlobServiceClient)
4. How to test locally when managed identity isn't available
5. Common pitfalls and error handling

Provide examples for both system-assigned and user-assigned identity.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide modern C# examples for authenticating Azure SDK clients with system-assigned and user-assigned managed identities, local development fallback, and robust error handling.","learn":true}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation guidance for C# Azure SDK authentication with managed identity.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet"}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Microsoft documentation for Azure.Identity ManagedIdentityCredential in .NET, including system-assigned and user-assigned identity constructors, local development authentication guidance, error handling, and use with BlobServiceClient or SecretClient.","learn":true}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Ground managed identity guidance in current official docs.","parameters":{"query":"Azure Identity .NET ManagedIdentityCredential system assigned user assigned client ID resource ID local development DefaultAzureCredential best practices"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for managed identity credentials and Azure SDK clients.","parameters":{"language":"csharp","query":"C# ManagedIdentityCredential ManagedIdentityId.SystemAssigned FromUserAssignedClientId BlobServiceClient SecretClient DefaultAzureCredential"}}
```

### 🔧 9. rg

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766051663-copilot-tool-output-e2ccdbb077e84f27b21d2d91447cfd67.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

### 🔧 10. rg

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766052120-copilot-tool-output-dfa7a11afaa04dd9ba8448ea1dd75d25.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"$files = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766052120-copilot-tool-output-dfa7a11afaa04dd9ba8448ea1dd75d25.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766051663-copilot-tool-output-e2ccdbb077e84f27b21d2d91447cfd67.txt'); foreach ($file in $files) { $json = Get-Content -Raw $file | ConvertFrom-Json; $json.results | Select-Object -First 10 | ForEach-Object { [PSCustomObject]@{ Title = if ($_.title) {$_.title} else {$_.description.Split(\"`n\")[0]}; Url = if ($_.url) {$_.url} elseif ($_.sourceUrl) {$_.sourceUrl} elseif ($_.link) {$_.link} else {''} } } | Format-Table -AutoSize }","description":"Extract documentation titles and URLs"}
```

### 🔧 12. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766052120-copilot-tool-output-dfa7a11afaa04dd9ba8448ea1dd75d25.txt","view_range":[1,6]}
```

### 🔧 13. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766051663-copilot-tool-output-e2ccdbb077e84f27b21d2d91447cfd67.txt","view_range":[1,6]}
```

### 🔧 14. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official guidance for system-assigned managed identity in .NET.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/authentication/system-assigned-managed-identity"}}
```

### 🔧 15. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official guidance for user-assigned managed identity in .NET.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/authentication/user-assigned-managed-identity"}}
```

### 🔧 16. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official Azure Identity best practices for production and local development.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/authentication/best-practices"}}
```

## Copilot Response

Managed identities let Azure SDK clients obtain Microsoft Entra tokens without storing credentials. Use `ManagedIdentityCredential` explicitly in Azure-hosted production code and `DefaultAzureCredential` for local development.

Install the current stable packages:

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
dotnet add package Azure.Security.KeyVault.Secrets
```

> `KeyVaultClient` belongs to older SDKs. New applications should use service-specific clients such as `SecretClient`, `KeyClient`, or `CertificateClient`.

## 1. System-assigned vs. user-assigned

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created on and deleted with the hosting resource | Independent Azure resource |
| Sharing | Belongs to one hosting resource | Can be attached to multiple resources |
| Selection in code | No identifier required | Specify client ID, resource ID, or object ID |
| Typical use | One application with its own permissions | Shared permissions, deployment slots, multiple applications |
| Operational impact | Replacing the host creates a new principal and requires new role assignments | Identity and role assignments survive host replacement |

In both cases:

1. Enable or attach the identity on the Azure compute resource.
2. Grant it the required **data-plane** Azure RBAC role at the narrowest practical scope.
3. Pass a `ManagedIdentityCredential` to the SDK client.

For example, reading blobs commonly requires **Storage Blob Data Reader**. Reading Key Vault secrets commonly requires **Key Vault Secrets User** when the vault uses Azure RBAC.

## 2. Creating `ManagedIdentityCredential`

### System-assigned identity

```csharp
using Azure.Core;
using Azure.Identity;

TokenCredential credential =
    new ManagedIdentityCredential(ManagedIdentityId.SystemAssigned);
```

The older parameterless form also selects the system-assigned identity when it is the only available identity:

```csharp
var credential = new ManagedIdentityCredential();
```

Using `ManagedIdentityId.SystemAssigned` makes the intent clearer.

### User-assigned identity using its client ID

```csharp
using Azure.Core;
using Azure.Identity;

string clientId =
    Environment.GetEnvironmentVariable("MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "MANAGED_IDENTITY_CLIENT_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(clientId));
```

The **client ID** is normally the most portable identifier. A resource ID can also be used:

```csharp
using Azure.Core;
using Azure.Identity;

string resourceId =
    Environment.GetEnvironmentVariable("MANAGED_IDENTITY_RESOURCE_ID")
    ?? throw new InvalidOperationException(
        "MANAGED_IDENTITY_RESOURCE_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedResourceId(
        new ResourceIdentifier(resourceId)));
```

Do not accidentally pass the principal/object ID to `FromUserAssignedClientId`. If you intentionally have an object ID, use `FromUserAssignedObjectId`.

## 3. Using the credential with Azure SDK clients

### System-assigned identity with Blob Storage and Key Vault

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;
using Azure.Storage.Blobs;

var credential =
    new ManagedIdentityCredential(ManagedIdentityId.SystemAssigned);

// Reuse the credential and clients; they are thread-safe.
var blobServiceClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);

var secretClient = new SecretClient(
    new Uri("https://myvault.vault.azure.net"),
    credential);

BlobContainerClient container =
    blobServiceClient.GetBlobContainerClient("documents");

await foreach (BlobItem blob in container.GetBlobsAsync())
{
    Console.WriteLine(blob.Name);
}

KeyVaultSecret secret =
    await secretClient.GetSecretAsync("database-password");

Console.WriteLine($"Loaded secret version {secret.Properties.Version}");
```

### User-assigned identity with Blob Storage and Key Vault

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;
using Azure.Storage.Blobs;

string managedIdentityClientId =
    Environment.GetEnvironmentVariable("MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "MANAGED_IDENTITY_CLIENT_ID is required.");

var credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(managedIdentityClientId));

var blobServiceClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);

var secretClient = new SecretClient(
    new Uri("https://myvault.vault.azure.net"),
    credential);

await foreach (BlobContainerItem container
    in blobServiceClient.GetBlobContainersAsync())
{
    Console.WriteLine(container.Name);
}

KeyVaultSecret secret =
    await secretClient.GetSecretAsync("database-password");
```

The user-assigned identity must be attached to the Azure host. Merely creating the identity or granting it RBAC permissions is insufficient.

## 4. Testing locally

Managed identity endpoints exist only on supported Azure hosts. A local process therefore cannot directly authenticate as a managed identity.

Use a developer credential locally and grant your developer identity the same permissions as the managed identity:

```csharp
using Azure.Core;
using Azure.Identity;

static TokenCredential CreateCredential(
    bool isAzureHosted,
    string? userAssignedClientId = null)
{
    if (!isAzureHosted)
    {
        return new DefaultAzureCredential(
            new DefaultAzureCredentialOptions
            {
                // Avoid probing for an unavailable managed identity locally.
                ExcludeManagedIdentityCredential = true
            });
    }

    return string.IsNullOrWhiteSpace(userAssignedClientId)
        ? new ManagedIdentityCredential(ManagedIdentityId.SystemAssigned)
        : new ManagedIdentityCredential(
            ManagedIdentityId.FromUserAssignedClientId(
                userAssignedClientId));
}
```

Usage:

```csharp
bool isAzureHosted =
    string.Equals(
        Environment.GetEnvironmentVariable("AZURE_HOSTED"),
        "true",
        StringComparison.OrdinalIgnoreCase);

string? managedIdentityClientId =
    Environment.GetEnvironmentVariable("MANAGED_IDENTITY_CLIENT_ID");

TokenCredential credential =
    CreateCredential(isAzureHosted, managedIdentityClientId);

var blobServiceClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);
```

`DefaultAzureCredential` can use credentials from tools such as:

- Visual Studio
- Azure CLI
- Azure PowerShell
- Azure Developer CLI
- VS Code with the appropriate broker package and sign-in configuration

For more deterministic local behavior, choose one explicitly:

```csharp
TokenCredential localCredential = new AzureCliCredential();
```

Local authentication validates SDK behavior and RBAC, but it does **not** prove that the managed identity is attached correctly. Perform an integration test on the actual Azure hosting service before release.

## 5. Error handling and common pitfalls

```csharp
using Azure;
using Azure.Identity;

try
{
    await foreach (BlobContainerItem container
        in blobServiceClient.GetBlobContainersAsync())
    {
        Console.WriteLine(container.Name);
    }
}
catch (CredentialUnavailableException ex)
{
    // No usable credential endpoint or developer credential was found.
    Console.Error.WriteLine(
        $"Credential unavailable: {ex.Message}");
}
catch (AuthenticationFailedException ex)
{
    // A credential was found, but token acquisition failed.
    Console.Error.WriteLine(
        $"Microsoft Entra authentication failed: {ex.Message}");
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // Authentication succeeded, but the identity lacks permission.
    Console.Error.WriteLine(
        $"Access denied. Verify data-plane RBAC assignments. " +
        $"Azure error: {ex.ErrorCode}");
}
catch (RequestFailedException ex) when (ex.Status == 429)
{
    Console.Error.WriteLine(
        $"The service throttled the request. Azure error: {ex.ErrorCode}");
    throw;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed with HTTP {ex.Status}: " +
        $"{ex.ErrorCode} - {ex.Message}");
    throw;
}
```

Key pitfalls:

- **Wrong RBAC role:** Management roles such as Contributor do not necessarily grant access to blob data or Key Vault secrets.
- **Wrong identifier:** Client ID, object/principal ID, and resource ID are different values.
- **Identity not attached:** A user-assigned identity must be associated with the VM, App Service, Container App, or other host.
- **Using `DefaultAzureCredential` in production:** It can select an unexpected credential. Prefer an explicit `ManagedIdentityCredential`.
- **Testing local managed identity directly:** Local authentication uses your developer account, not the managed identity.
- **Role propagation:** New role assignments can take several minutes to become effective.
- **Network restrictions:** A valid token does not bypass Key Vault firewalls, storage firewalls, private endpoints, or DNS configuration.
- **Creating clients repeatedly:** Credentials and SDK clients are thread-safe and should generally be reused to benefit from token caching and connection pooling.
- **Logging secrets or tokens:** Log status codes, request IDs, and Azure error codes—not access tokens, secret values, or credential material.
- **Over-customizing retries:** Azure SDK clients already retry many transient service errors. `ManagedIdentityCredential` also has a resilient retry strategy when used directly.

**References:**

- [System-assigned managed identity for .NET](https://learn.microsoft.com/dotnet/azure/sdk/authentication/system-assigned-managed-identity)
- [User-assigned managed identity for .NET](https://learn.microsoft.com/dotnet/azure/sdk/authentication/user-assigned-managed-identity)
- [Azure Identity authentication best practices](https://learn.microsoft.com/dotnet/azure/sdk/authentication/best-practices)
- [Azure Identity client library](https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme)

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response thoroughly covers explicit managed identity usage, Azure SDK clients, local development, and error handling, but misses the two required credential-chain examples.

### Strengths

- Provides correct system-assigned and user-assigned ManagedIdentityCredential examples.
- Shows practical Blob Storage and Key Vault client integration.
- Includes CredentialUnavailableException and other useful Azure authentication and authorization error handling.

### Issues

- DefaultAzureCredential is configured to exclude managed identity rather than demonstrating managed identity in its chain.
- ChainedTokenCredential is not used for local fallback.

## Grader Results

- managed-identity-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/6)
      - `ManagedIdentityCredential` class and constructors: Pass
      - System-assigned: no parameters needed: Pass
      - User-assigned: passing the client ID: Pass
      - Integration with `DefaultAzureCredential` (managed identity in the chain): Fail
      - `CredentialUnavailableException` when not running in Azure: Pass
      - Combining with `ChainedTokenCredential` for local fallback: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 67% | 1.00 | 0.6667 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6667** | **66.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-managed-identity --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)

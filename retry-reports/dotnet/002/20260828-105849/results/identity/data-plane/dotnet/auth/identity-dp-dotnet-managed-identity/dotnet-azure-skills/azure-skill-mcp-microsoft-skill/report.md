# Evaluation Report: identity-dp-dotnet-managed-identity

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 325.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-managed-identity` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 3/6 |
| Duration | 325.1s |
| Timestamp | 2026-08-28T02:58:49Z |
| Files Generated | 0 |
| Event Count | 3963 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 280.5s |
| Review | 43.5s |
| **Total** | **325.1s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=147210 out=3252 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 848ms

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for .NET

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```bash
dotnet add package Azure.Identity

# For ASP.NET Core integration
dotnet add package Microsoft.Extensions.Azure

# For brokered authentication and Visual Studio Code credential support
dotnet add package Azure.Identity.Broker
```

## Environment Variables

### Service Principal with Secret

```bash
AZURE_CLIENT_ID=<application-client-id>
AZURE_TENANT_ID=<directory-tenant-id>
AZURE_CLIENT_SECRET=<client-secret-value>
```

### Service Principal with Certificate

```bash
AZURE_CLIENT_ID=<application-client-id>
AZURE_TENANT_ID=<directory-tenant-id>
AZURE_CLIENT_CERTIFICATE_PATH=<path-to-pfx-or-pem>
AZURE_CLIENT_CERTIFICATE_PASSWORD=<certificate-password>  # Optional
```

### Managed Identity

```bash
AZURE_CLIENT_ID=<user-assigned-managed-identity-client-id>  # Only for user-assigned
```

## DefaultAzureCredential

The recommended credential for most scenarios. Tries multiple authentication methods in order. See [DefaultAzureCredential overview](https://aka.ms/azsdk/net/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

### Basic Usage

```csharp
using Azure.Identity;
using Azure.Storage.Blobs;

var credential = new DefaultAzureCredential();
var blobClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);
```

### ASP.NET Core with Dependency Injection

```csharp
using Azure.Identity;
using Microsoft.Extensions.Azure;

builder.Services.AddAzureClients(clientBuilder =>
{
    clientBuilder.AddBlobServiceClient(
        new Uri("https://myaccount.blob.core.windows.net"));
    clientBuilder.AddSecretClient(
        new Uri("https://myvault.vault.azure.net"));
    
    // Uses DefaultAzureCredential by default
    clientBuilder.UseCredential(new DefaultAzureCredential());
});
```

### Customizing DefaultAz
... (truncated)
```

### ❌ 2. view — 183291ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 183288ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 183293ms

**Input:**

```json
{"intent":"Get current best practices for C# Azure SDK authentication with system-assigned and user-assigned managed identities, local development credentials, and error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 183295ms

**Input:**

```json
{"intent":"Find official current Microsoft guidance for using ManagedIdentityCredential in .NET for system-assigned and user-assigned managed identities, local development with DefaultAzureCredential, and troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. view — 1434ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-identity-dotnet b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-identity-dotnet
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-identity-dotnet
@@ -1,0 +1,1 @@
 SKILL.md

```

## Copilot Response

I’ll align the examples with the current Azure Identity SDK guidance, including the managed-identity-specific credential constructors and local development chain.

## 1. System-assigned vs. user-assigned identity

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created and deleted with the Azure resource | Independent Azure resource |
| Assignment | One identity per hosting resource | Can be assigned to multiple resources |
| Selection | Azure SDK uses the host resource’s identity | Specify its client ID, resource ID, or object ID |
| Typical use | Simple, resource-specific workloads | Shared permissions, stable identity across redeployments |
| Local availability | Not available | Not available |

Both identities require explicit authorization. Enabling an identity does **not** automatically grant access to Blob Storage, Key Vault, or other services.

## 2. Install the SDK packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
dotnet add package Azure.Security.KeyVault.Secrets
```

`SecretClient` is the current Azure SDK client for Key Vault secrets; the older `KeyVaultClient` belongs to the legacy SDK.

## 3. System-assigned managed identity

Enable system-assigned identity on the App Service, Function App, VM, Container App, or other supported Azure host, then grant that identity the necessary RBAC role.

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;
using Azure.Storage.Blobs;

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.SystemAssigned);

var blobServiceClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);

var secretClient = new SecretClient(
    new Uri("https://myvault.vault.azure.net"),
    credential);

// Network calls happen here, not in the client constructors.
BlobContainerClient container = blobServiceClient.GetBlobContainerClient("documents");
await foreach (BlobItem blob in container.GetBlobsAsync())
{
    Console.WriteLine(blob.Name);
}

KeyVaultSecret secret = await secretClient.GetSecretAsync("DatabasePassword");
Console.WriteLine($"Secret retrieved: {secret.Name}");
```

Example roles include:

- **Storage Blob Data Reader** or **Storage Blob Data Contributor**
- **Key Vault Secrets User**

Prefer narrowly scoped role assignments—for example, at the container or vault level rather than the entire subscription.

## 4. User-assigned managed identity

When several user-assigned identities are attached to the host, the application must identify which one to use. The **client ID** is normally the simplest selector.

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;
using Azure.Storage.Blobs;

string managedIdentityClientId =
    Environment.GetEnvironmentVariable("MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "MANAGED_IDENTITY_CLIENT_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(managedIdentityClientId));

var blobServiceClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);

var secretClient = new SecretClient(
    new Uri("https://myvault.vault.azure.net"),
    credential);

BlobContainerClient container = blobServiceClient.GetBlobContainerClient("documents");
await container.GetPropertiesAsync();

KeyVaultSecret secret = await secretClient.GetSecretAsync("DatabasePassword");
```

You can also select a user-assigned identity by Azure resource ID:

```csharp
string identityResourceId =
    "/subscriptions/<subscription-id>" +
    "/resourceGroups/<resource-group>" +
    "/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity-name>";

var credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedResourceId(
        new ResourceIdentifier(identityResourceId)));
```

Using `ManagedIdentityId` makes it explicit whether a value is a client ID, object ID, or resource ID and avoids confusing those identifiers.

## 5. Local development

A managed identity exists only in its supported Azure hosting environment. Calling `ManagedIdentityCredential` directly on a developer machine normally produces `CredentialUnavailableException`.

Use a developer credential locally and a deterministic managed identity credential in Azure:

```csharp
using Azure.Core;
using Azure.Identity;

static TokenCredential CreateCredential(
    bool isLocalDevelopment,
    string? userAssignedClientId = null)
{
    if (isLocalDevelopment)
    {
        return new DefaultAzureCredential(
            new DefaultAzureCredentialOptions
            {
                // Avoid probing the managed identity endpoint locally.
                ExcludeManagedIdentityCredential = true
            });
    }

    return string.IsNullOrWhiteSpace(userAssignedClientId)
        ? new ManagedIdentityCredential(ManagedIdentityId.SystemAssigned)
        : new ManagedIdentityCredential(
            ManagedIdentityId.FromUserAssignedClientId(userAssignedClientId));
}
```

Use it with either identity:

```csharp
bool isLocal = string.Equals(
    Environment.GetEnvironmentVariable("DOTNET_ENVIRONMENT"),
    "Development",
    StringComparison.OrdinalIgnoreCase);

// null in Azure selects the system-assigned identity.
// Supply the client ID to select a user-assigned identity.
string? userAssignedClientId =
    Environment.GetEnvironmentVariable("MANAGED_IDENTITY_CLIENT_ID");

TokenCredential credential = CreateCredential(isLocal, userAssignedClientId);

var secretClient = new SecretClient(
    new Uri("https://myvault.vault.azure.net"),
    credential);
```

Locally, `DefaultAzureCredential` can use supported developer authentication such as Visual Studio, Azure CLI, Azure PowerShell, or Azure Developer CLI. The developer’s account needs the same data-plane permissions required by the application.

For CI, prefer workload identity federation or another explicitly configured non-secret mechanism instead of pretending managed identity is locally available.

## 6. Error handling

Handle credential availability, authentication, and service authorization separately:

```csharp
using Azure;
using Azure.Identity;

try
{
    KeyVaultSecret secret = await secretClient.GetSecretAsync("DatabasePassword");
}
catch (CredentialUnavailableException ex)
{
    // No supported identity endpoint or local developer credential was available.
    Console.Error.WriteLine($"Credential unavailable: {ex.Message}");
    throw;
}
catch (AuthenticationFailedException ex)
{
    // An identity was found, but token acquisition failed.
    Console.Error.WriteLine($"Authentication failed: {ex.Message}");
    throw;
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // Authentication succeeded, but the identity lacks permission.
    Console.Error.WriteLine(
        $"Access denied. Check RBAC or Key Vault access configuration. " +
        $"Code={ex.ErrorCode}, RequestId={ex.ClientRequestId}");
    throw;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure request failed. Status={ex.Status}, " +
        $"Code={ex.ErrorCode}, RequestId={ex.ClientRequestId}");
    throw;
}
```

Catch `CredentialUnavailableException` before `AuthenticationFailedException` because it is a more specific authentication exception.

## Common pitfalls

- **Missing RBAC:** Successfully obtaining a token does not grant access. A `403` generally indicates an authorization problem rather than a credential problem.
- **Wrong role type:** Blob operations require data-plane roles such as **Storage Blob Data Reader**; management-plane roles such as **Reader** are insufficient.
- **Wrong user-assigned identifier:** Client ID, object/principal ID, and Azure resource ID are different values. Use the corresponding `ManagedIdentityId.From...` method.
- **Identity not attached to the host:** Creating a user-assigned identity is not enough; it must also be assigned to the App Service, VM, Function App, or other host.
- **Multiple user-assigned identities:** Explicitly select the intended identity. Otherwise token acquisition may fail because Azure cannot infer which identity to use.
- **RBAC propagation delays:** New identity assignments and role assignments may take several minutes to become effective.
- **Using managed identity locally:** The Azure managed identity metadata endpoint does not exist on ordinary developer machines. Use `DefaultAzureCredential` or another developer credential.
- **Using a broad credential chain in production:** Prefer an explicit `ManagedIdentityCredential` in production so an unexpected environment variable or developer credential cannot silently change the application identity.
- **Creating clients repeatedly:** Azure credentials and service clients are thread-safe. Create and reuse them as singletons where practical.
- **Logging secrets or tokens:** Log error codes and request IDs, but never access tokens, client secrets, or retrieved Key Vault secret values.

References: [Managed identities with Azure SDKs](https://learn.microsoft.com/dotnet/azure/sdk/authentication/managed-identity), [Azure Identity authentication best practices](https://learn.microsoft.com/dotnet/azure/sdk/authentication/best-practices), and [`ManagedIdentityCredential` API](https://learn.microsoft.com/dotnet/api/azure.identity.managedidentitycredential).

## Reviewer Notes (LLM-as-Judge)

### Summary

3 of 6 criteria passed. The response thoroughly explains managed identities and Azure client usage, but misses three specifically required credential-construction and chaining patterns.

### Strengths

- Clearly distinguishes system-assigned and user-assigned managed identities.
- Correctly selects a user-assigned identity using its client ID.
- Provides realistic Blob Storage and Key Vault integrations plus useful exception handling and pitfall guidance.

### Issues

- The system-assigned example does not use the parameterless ManagedIdentityCredential constructor.
- DefaultAzureCredential is configured to exclude managed identity rather than demonstrating managed identity in its chain.
- ChainedTokenCredential is not used for local fallback.

## Grader Results

- managed-identity-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (3/6)
      - `ManagedIdentityCredential` class and constructors: Pass
      - System-assigned: no parameters needed: Fail
      - User-assigned: passing the client ID: Pass
      - Integration with `DefaultAzureCredential` (managed identity in the chain): Fail
      - `CredentialUnavailableException` when not running in Azure: Pass
      - Combining with `ChainedTokenCredential` for local fallback: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 50% | 1.00 | 0.5000 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.5000** | **50.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-managed-identity --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

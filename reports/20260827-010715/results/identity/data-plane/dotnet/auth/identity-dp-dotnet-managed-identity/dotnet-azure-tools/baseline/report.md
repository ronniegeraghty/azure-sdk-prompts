# Evaluation Report: identity-dp-dotnet-managed-identity

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 201.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-managed-identity` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 5/6 |
| Duration | 201.2s |
| Timestamp | 2026-08-26T17:36:23Z |
| Files Generated | 0 |
| Event Count | 4394 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 177.2s |
| Review | 23.2s |
| **Total** | **201.2s** |

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
| Token Usage | in=45491 out=2869 |
| Turn Count | 2 |

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

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com .NET Azure.Identity ManagedIdentityCredential system-assigned user-assigned client ID resource ID 2026"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com .NET Azure SDK authenticate locally DefaultAzureCredential ManagedIdentityCredential Key Vault BlobServiceClient managed identities"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Identity .NET credential chains ManagedIdentityCredential troubleshooting CredentialUnavailableException AuthenticationFailedException"}
```

## Copilot Response

I’ll ground the examples in the current Azure Identity and SDK APIs, then show deploy-time and local-development patterns.

Managed identity lets Azure-hosted code obtain Microsoft Entra tokens without storing client secrets. Most modern Azure SDK clients accept an `Azure.Core.TokenCredential`, including `SecretClient` for Key Vault and `BlobServiceClient` for Storage.

Install the relevant packages:

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Security.KeyVault.Secrets
dotnet add package Azure.Storage.Blobs
```

> `KeyVaultClient` is from an older SDK. For new applications, use `SecretClient`, `KeyClient`, or `CertificateClient`.

## 1. System-assigned vs. user-assigned

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created and deleted with the Azure host resource | Independent Azure resource |
| Sharing | Used only by its host resource | Can be attached to multiple resources |
| Selection | No identity ID is normally required | Select using its **client ID** or Azure **resource ID** |
| Typical use | One application with one identity | Shared identity, stable permissions, deployment-slot or multi-resource scenarios |
| Multiple identities | One system-assigned identity per host | Several may be attached, so explicit selection is important |

In either case, two separate configurations are required:

1. Attach or enable the identity on the Azure compute resource.
2. Grant that identity authorization on the destination resource—for example, **Key Vault Secrets User** or **Storage Blob Data Reader**.

Authentication can succeed while the subsequent request receives `403 Forbidden` if step 2 is missing.

## 2. Creating `ManagedIdentityCredential`

### System-assigned identity

With no identity selector, `ManagedIdentityCredential` uses the system-assigned identity:

```csharp
using Azure.Core;
using Azure.Identity;

TokenCredential credential = new ManagedIdentityCredential();
```

Current Azure.Identity versions also allow the intent to be expressed explicitly:

```csharp
TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.SystemAssigned);
```

### User-assigned identity

The preferred selector is usually the identity's **client ID**:

```csharp
using Azure.Core;
using Azure.Identity;

string clientId = Environment.GetEnvironmentVariable(
    "AZURE_MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "AZURE_MANAGED_IDENTITY_CLIENT_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(clientId));
```

You can alternatively select it by its Azure resource ID:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;

string resourceId = Environment.GetEnvironmentVariable(
    "AZURE_MANAGED_IDENTITY_RESOURCE_ID")
    ?? throw new InvalidOperationException(
        "AZURE_MANAGED_IDENTITY_RESOURCE_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedResourceId(
        new ResourceIdentifier(resourceId)));
```

Example resource ID:

```text
/subscriptions/<subscription-id>/resourceGroups/<resource-group>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity-name>
```

Do not confuse these identifiers:

- **Client ID/application ID:** commonly used to select a user-assigned identity.
- **Object ID/principal ID:** identifies its service principal and is commonly used in role assignments.
- **Resource ID:** Azure Resource Manager path for the managed-identity resource.

## 3. Using the credential with Azure SDK clients

Azure SDK clients are thread-safe and should generally be created once and reused.

### System-assigned identity with Key Vault

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

Uri vaultUri = new(
    Environment.GetEnvironmentVariable("KEY_VAULT_URI")
    ?? throw new InvalidOperationException("KEY_VAULT_URI is not configured."));

TokenCredential credential = new ManagedIdentityCredential();
SecretClient secretClient = new(vaultUri, credential);

try
{
    KeyVaultSecret secret =
        await secretClient.GetSecretAsync("database-connection-string");

    Console.WriteLine($"Retrieved secret version {secret.Properties.Version}.");
}
catch (CredentialUnavailableException ex)
{
    Console.Error.WriteLine(
        $"Managed identity is unavailable on this host: {ex.Message}");
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine(
        $"Managed identity token acquisition failed: {ex.Message}");
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    Console.Error.WriteLine(
        "Authentication succeeded, but the identity cannot read this secret.");
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Key Vault request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
}
```

### User-assigned identity with Blob Storage

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

string clientId = Environment.GetEnvironmentVariable(
    "AZURE_MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "AZURE_MANAGED_IDENTITY_CLIENT_ID is not configured.");

string accountName = Environment.GetEnvironmentVariable("STORAGE_ACCOUNT_NAME")
    ?? throw new InvalidOperationException(
        "STORAGE_ACCOUNT_NAME is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(clientId));

Uri serviceUri = new($"https://{accountName}.blob.core.windows.net");
BlobServiceClient blobServiceClient = new(serviceUri, credential);

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
    Console.Error.WriteLine(
        $"The selected managed identity is unavailable: {ex.Message}");
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine(
        $"Token acquisition failed: {ex.Message}");
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    Console.Error.WriteLine(
        "The identity needs an appropriate Storage Blob Data role.");
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Storage request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
}
```

## 4. Local development and testing

`ManagedIdentityCredential` normally cannot authenticate on a developer workstation because the Azure managed-identity endpoint is absent. Use a developer credential locally instead of emulating or storing managed-identity secrets.

A clean pattern is to select credentials by environment:

```csharp
using Azure.Core;
using Azure.Identity;

static TokenCredential CreateCredential(bool isDevelopment)
{
    if (isDevelopment)
    {
        return new DefaultAzureCredential(
            new DefaultAzureCredentialOptions
            {
                // Fail faster locally instead of probing for managed identity.
                ExcludeManagedIdentityCredential = true
            });
    }

    return new ManagedIdentityCredential();
}
```

`DefaultAzureCredential` can use credentials from supported developer tools such as Azure CLI, Visual Studio, or Azure Developer CLI. The signed-in developer must be granted the necessary Key Vault or Storage roles.

For a user-assigned identity in production:

```csharp
static TokenCredential CreateCredential(
    bool isDevelopment,
    string managedIdentityClientId)
{
    if (isDevelopment)
    {
        return new DefaultAzureCredential(
            new DefaultAzureCredentialOptions
            {
                ExcludeManagedIdentityCredential = true
            });
    }

    return new ManagedIdentityCredential(
        ManagedIdentityId.FromUserAssignedClientId(
            managedIdentityClientId));
}
```

Another option is to use `DefaultAzureCredential` in both environments:

```csharp
TokenCredential credential = new DefaultAzureCredential(
    new DefaultAzureCredentialOptions
    {
        ManagedIdentityClientId =
            Environment.GetEnvironmentVariable(
                "AZURE_MANAGED_IDENTITY_CLIENT_ID")
    });
```

Locally, the chain uses a developer login; in Azure, it selects the configured user-assigned identity. For production systems, an explicit `ManagedIdentityCredential` is more deterministic because it cannot unexpectedly authenticate using another credential from the chain.

For offline unit tests, inject `TokenCredential` or the Azure client into your application and substitute a fake/mock. Do not make unit tests depend on a managed-identity endpoint. Integration tests can use `DefaultAzureCredential` with a dedicated, least-privileged developer or test identity.

## 5. Common pitfalls

- **The identity is enabled but not authorized:** Token acquisition succeeds, but Key Vault or Storage returns `403`.
- **Wrong user-assigned identifier:** The client ID is not the object/principal ID. Validate which value your configuration contains.
- **Identity is not attached to the host:** Creating a user-assigned identity alone is insufficient; it must also be attached to the App Service, VM, Function, container environment, or other supported host.
- **Ambiguous identity selection:** If several user-assigned identities are attached, always provide a client ID or resource ID.
- **Expecting constructor-time failure:** Azure SDK clients usually acquire tokens lazily. Authentication errors normally occur during the first service operation, not during client construction.
- **Immediate use after configuration:** Identity attachment and RBAC changes can take time to propagate.
- **Using management-plane roles for data access:** Generic roles such as Contributor do not necessarily grant secret or blob data access. Use the appropriate data-plane role.
- **Testing locally with `ManagedIdentityCredential`:** It will normally produce `CredentialUnavailableException`; use a developer credential locally.
- **Catching only authentication exceptions:** Authorization and service failures are generally reported as `RequestFailedException`.
- **Excessive custom retries:** Azure SDK clients already retry supported transient failures. Avoid retrying permanent `400`, `401`, or `403` responses indefinitely.
- **Creating clients per request:** Reuse `SecretClient`, `BlobServiceClient`, and credential instances.

Official references: [ManagedIdentityCredential](https://learn.microsoft.com/dotnet/api/azure.identity.managedidentitycredential), [ManagedIdentityId](https://learn.microsoft.com/dotnet/api/azure.identity.managedidentityid), [local development authentication](https://learn.microsoft.com/dotnet/azure/sdk/authentication/local-development-dev-accounts), and [Azure Identity authentication best practices](https://learn.microsoft.com/dotnet/azure/sdk/authentication/best-practices).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 6 criteria passed. The response thoroughly covers managed identity types, credential construction, Azure client integration, DefaultAzureCredential, local development, and error handling, but omits the specifically required ChainedTokenCredential fallback example.

### Strengths

- Provides correct system-assigned and user-assigned ManagedIdentityCredential examples.
- Integrates credentials with modern Key Vault and Blob Storage clients.
- Clearly distinguishes authentication, authorization, and service errors.
- Explains DefaultAzureCredential behavior and practical local-development options.

### Issues

- No ChainedTokenCredential example is provided for explicitly combining managed identity with a local developer credential.

## Grader Results

- managed-identity-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/6)
      - `ManagedIdentityCredential` class and constructors: Pass
      - System-assigned: no parameters needed: Pass
      - User-assigned: passing the client ID: Pass
      - Integration with `DefaultAzureCredential` (managed identity in the chain): Pass
      - `CredentialUnavailableException` when not running in Azure: Pass
      - Combining with `ChainedTokenCredential` for local fallback: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8333** | **83.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-managed-identity --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)

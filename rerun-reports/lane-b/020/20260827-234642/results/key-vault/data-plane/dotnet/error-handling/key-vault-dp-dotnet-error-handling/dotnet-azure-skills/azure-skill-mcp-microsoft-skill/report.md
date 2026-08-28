# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 359.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 4/7 |
| Duration | 359.1s |
| Timestamp | 2026-08-27T15:46:42Z |
| Files Generated | 0 |
| Event Count | 3261 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 333.4s |
| Review | 24.6s |
| **Total** | **359.1s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-security-keyvault-keys-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=155517 out=2731 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 5065ms

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

### ❌ 2. view — 187476ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 187472ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 187478ms

**Input:**

```json
{"intent":"Provide modern .NET Azure SDK best practices for handling Azure Key Vault Secrets errors, RequestFailedException status/error codes, conflicts, and throttling retries.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 187480ms

**Input:**

```json
{"intent":"Find official Microsoft guidance for Azure.Security.KeyVault.Secrets .NET RequestFailedException handling, SecretClient retry configuration, HTTP 403, 404, 409 or 412 conflicts, and 429 throttling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_search (via github-mcp-server) — 75320ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure.Security.KeyVault.Secrets SecretClientOptions Retry RequestFailedException 429 .NET ErrorCode Status"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The `RequestFailedException` in the Azure SDK for .NET is an exception type thrown when a request to an Azure service fails. This exception provides two useful properties for troubleshooting:\n\n- **ErrorCode**: Gets the service-specific error code, allowing you to understand the precise reason for the failure from the Azure service side【3:2†source】.\n- **Status**: Gets the HTTP status code of the response (e.g., 404, 403, 400). If the response was not received, this will be `0`【3:3†source】.\n\n### Usage\n\n- Check the **Status** property to know the HTTP status code of the failed request.\n- Check the **ErrorCode** property to see the specific Azure service error code (these are usually documented per Azure service).\n- Use the `GetRawResponse()` method to retrieve the raw response that led to the exception, which can help further debug the error【3:6†source】.\n\n**Example:**\n```csharp\ntry\n{\n    // Azure SDK code that may fail\n}\ncatch (RequestFailedException ex)\n{\n    Console.WriteLine($\"Status: {ex.Status}\");      // HTTP status code\n    Console.WriteLine($\"ErrorCode: {ex.ErrorCode}\"); // Service-specific error code\n    Console.WriteLine($\"Message: {ex.Message}\");    // Exception message\n}\n```\n**Common Causes:**\n- Authentication/authorization failures (e.g., \"AuthorizationFailure\" with Status 403)\n- Resource not found (e.g., \"ResourceNotFound\" with Status 404)\n- Missing/invalid headers or request formatting issues【3:5†source】【3:7†source】.\n\nFor official documentation and more implementation details on the `RequestFailedException`, see Microsoft Learn:  \n- [RequestFailedException Class](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception)【3:0†source】\n- [ErrorCode Property](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.errorcode)【3:2†source】\n- [Status Property](https://learn.microsoft.com/dotnet/api/azure.requestfailedex
... (truncated)
```

### ✅ 7. web_search (via github-mcp-server) — 73268ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Key Vault secrets 409 Conflict SecretClient SetSecretAsync deleted secret recoverable .NET"}
```

**Output:**

```
{"type":"output_text","text":{"value":"When you receive a `409 Conflict` error while calling `SecretClient.SetSecretAsync` for an Azure Key Vault secret, it's often because the secret with that name was recently deleted and is still in a \"soft-deleted\" (recoverable) state. If \"soft delete\" is enabled on your Key Vault (which is now mandatory in all new vaults), deleted secrets are not immediately purged and remain recoverable for a retention period (default 90 days).\n\n**Behavior:**\n- If you try to set (add) a secret with the same name as a deleted-but-not-purged secret, you'll see a `409 Conflict`:\n  > \"A secret with (name) is currently deleted and recoverable but not purged. You must permanently delete it (purge) before re-creating with this name.\"\n- To resolve this, you must purge the deleted secret or recover it.\n\n**What you can do:**\n\n1. **List deleted secrets**  \n   (Optional: to confirm it's in a deleted state)\n\n   ```csharp\n   await foreach(var deletedSecret in client.GetDeletedSecretsAsync())\n   {\n       Console.WriteLine(deletedSecret.Name);\n   }\n   ```\n\n2. **Recover the deleted secret**  \n   (if you want to restore its value and versions)\n\n   ```csharp\n   await client.StartRecoverDeletedSecretAsync(\"secret-name\");\n   ```\n\n3. **Purge the deleted secret**  \n   (if you want to permanently delete it and re-create)\n\n   ```csharp\n   await client.PurgeDeletedSecretAsync(\"secret-name\");\n   ```\n\n4. **Re-attempt SetSecretAsync**  \n   After purge (may take time to complete), you can now create the new secret as expected.\n\n**References:**\n- [SecretClient.SetSecretAsync documentation](https://learn.microsoft.com/dotnet/api/azure.security.keyvault.secrets.secretclient.setsecretasync) mentions that RequestFailedException is thrown, and “the server returned an error” such as when a secret is in a deleted and recoverable state【3:0†source】.\n- [Azure Key Vault Recovery Overview](https://learn.microsoft.com/azure/key-v
... (truncated)
```

### ✅ 8. web_search (via github-mcp-server) — 71495ms

**Input:**

```json
{"query":"site:learn.microsoft.com RequestFailedException ErrorCode Status .NET Azure SDK"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To handle `RequestFailedException` with status code 429 (Too Many Requests) using the `Azure.Security.KeyVault.Secrets` library in .NET, you should configure retry behavior via the `SecretClientOptions` when initializing your `SecretClient`. \n\n**Key Points:**\n- **429** indicates throttling from the Azure Key Vault service.\n- `SecretClientOptions` has a `Retry` property (type: `RetryOptions`) that controls the retry policy for handling such throttling errors.\n- You can customize max retries, delay between retries, and other settings.\n\n**Example:**\n```csharp\nvar options = new SecretClientOptions()\n{\n    Retry =\n    {\n        MaxRetries = 5,                 // Number of retries\n        Delay = TimeSpan.FromSeconds(2), // Initial delay\n        MaxDelay = TimeSpan.FromSeconds(16), // Max delay\n        Mode = RetryMode.Exponential     // Retry mode (Exponential or Fixed)\n    }\n};\n\nvar client = new SecretClient(new Uri(\"https://<your-keyvault-name>.vault.azure.net/\"), new DefaultAzureCredential(), options);\n```\n\nWhen a 429 error is encountered, the Azure SDK will automatically apply the retry policy defined in `SecretClientOptions.Retry`.\n\n- You can catch and inspect a `RequestFailedException` in your code:\n```csharp\ntry\n{\n    // Key Vault operation here\n}\ncatch (RequestFailedException ex) when (ex.Status == 429)\n{\n    // Handle throttling (maybe log or implement custom backoff)\n}\n```\n\n**References:**\n- [SecretClientOptions Class Documentation](https://learn.microsoft.com/dotnet/api/azure.security.keyvault.secrets.secretclientoptions)【3:0†source】\n- [Retry Policy in Azure SDK](https://learn.microsoft.com/dotnet/api/azure.core.retryoptions)【3:0†source】\n- [Handling 429 and Throttling](https://learn.microsoft.com/azure/key-vault/general/overview-throttling) (general docs)\n\nLet me know if you need more detailed error handling code or advanced retry customization!","annotations":[{"tex
... (truncated)
```

### ✅ 9. web_fetch — 69730ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclientoptions?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclientoptions?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: SecretClientOptions Class (Azure.Security.KeyVault.Secrets) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclientoptions?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Security.KeyVault.Secrets.SecretClientOptions
namespace: Azure.Security.KeyVault.Secrets
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-net/issues
author: azure-sdk
ms.author: azsdkdocs
ms.topic: generated-reference
ms.devlang: csharp
ms.date: 2021-04-13T00:00:00.0000000Z
uhfHeaderId: Azure
breadcrumb_path: ../azure_breadcrumb/toc.json
apiPlatform: dotnet
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/f488294d-f483-456e-94e3-755f933b811b
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Security.KeyVault.Secrets.SecretClientOptions
api_location:
- Azure.Security.KeyVault.Secrets.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 1e6565e2-dd1d-2d17-8576-515d366464da
document_version_independent_id: b21dd2d1-5543-59e8-9b84-cb112a3a0c5c
updated_at: 2026-05-06T11:36:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Security.KeyVault.Secrets/SecretClientOptions.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/f3535d7b864f1a7a9a07185098c8ba5bc53a31d7/xml/Azure.Security.KeyVault.Secrets/SecretClientOptions.xml
git_commit_id: f3535d7b864f1a7a9a07185098c8ba5bc53a31d7
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class

... (truncated)
```

## Copilot Response

I’m grounding the patterns in the SDK’s retry and exception guidance, including the subtle distinction between creating a new secret version and conditional-update conflicts.

Use `Azure.RequestFailedException` for service failures. Its `Status` property contains the HTTP status code, while `ErrorCode` contains Key Vault’s service-specific code.

```bash
dotnet add package Azure.Security.KeyVault.Secrets
dotnet add package Azure.Identity
```

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

Uri vaultUri = new(
    Environment.GetEnvironmentVariable("AZURE_KEYVAULT_URL")
    ?? throw new InvalidOperationException("AZURE_KEYVAULT_URL is not set."));

var options = new SecretClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(16),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

var client = new SecretClient(
    vaultUri,
    new DefaultAzureCredential(),
    options);
```

The Azure SDK retry policy automatically retries transient failures such as 429 responses and respects service-provided retry delays such as `Retry-After`.

## Reading a secret

```csharp
static async Task<string?> GetSecretValueAsync(
    SecretClient client,
    string secretName,
    string? version = null,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.GetSecretAsync(secretName, version, cancellationToken);

        return response.Value.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        Console.Error.WriteLine(
            "Key Vault denied access. " +
            $"Status={ex.Status}, ErrorCode={ex.ErrorCode}, Message={ex.Message}");

        // Check the caller's Key Vault RBAC role or legacy access policy.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 404)
    {
        Console.Error.WriteLine(
            $"Secret '{secretName}' was not found. " +
            $"Status={ex.Status}, ErrorCode={ex.ErrorCode}");

        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        // The SDK has already exhausted the configured retries.
        Console.Error.WriteLine(
            "Key Vault throttling persisted after retries. " +
            $"Status={ex.Status}, ErrorCode={ex.ErrorCode}");

        throw;
    }
    catch (RequestFailedException ex)
    {
        Console.Error.WriteLine(
            "Key Vault request failed. " +
            $"Status={ex.Status}, ErrorCode={ex.ErrorCode}, Message={ex.Message}");

        throw;
    }
}
```

For required configuration, throwing on 404 is usually better than returning `null`; use the return pattern only when a missing secret is an expected state.

## Creating a new secret version

`SetSecretAsync` creates a new immutable version when the secret already exists:

```csharp
static async Task<KeyVaultSecret> SetSecretAsync(
    SecretClient client,
    string secretName,
    string value,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.SetSecretAsync(secretName, value, cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        Console.Error.WriteLine(
            $"Not authorized to set secret '{secretName}'. " +
            $"ErrorCode={ex.ErrorCode}");

        throw;
    }
    catch (RequestFailedException ex) when (ex.Status is 409 or 412)
    {
        Console.Error.WriteLine(
            $"Secret state or version conflict for '{secretName}'. " +
            $"Status={ex.Status}, ErrorCode={ex.ErrorCode}");

        // Re-read the secret/deleted-secret state before deciding whether
        // to recover it, wait, or retry the operation.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        Console.Error.WriteLine(
            $"Setting '{secretName}' remained throttled after SDK retries. " +
            $"ErrorCode={ex.ErrorCode}");

        throw;
    }
}
```

A common `409 Conflict` from `SetSecretAsync` does **not** mean two callers created the same version—each successful set creates a distinct version. It commonly means the name belongs to a soft-deleted, recoverable secret. Resolve that deliberately by recovering the secret or, when permanently deleting it is explicitly intended and permitted, purging it. Do not blindly retry a 409.

Existing secret values and versions are immutable. To change a value, create a new version with `SetSecretAsync`; use `UpdateSecretPropertiesAsync` only for metadata such as enabled state, expiration, content type, and tags.

## Centralized inspection

When operation-specific handling is unnecessary, catch once and inspect the exception:

```csharp
try
{
    KeyVaultSecret secret =
        await client.GetSecretAsync("database-password");

    Console.WriteLine($"Version: {secret.Properties.Version}");
}
catch (RequestFailedException ex)
{
    int httpStatus = ex.Status;        // 0 if no HTTP response was received
    string? serviceCode = ex.ErrorCode;
    string message = ex.Message;

    Response? rawResponse = ex.GetRawResponse();
    string? requestId = null;

    if (rawResponse is not null &&
        rawResponse.Headers.TryGetValue("x-ms-request-id", out string? value))
    {
        requestId = value;
    }

    Console.Error.WriteLine(
        $"Key Vault failure: Status={httpStatus}, " +
        $"ErrorCode={serviceCode ?? "<none>"}, " +
        $"RequestId={requestId ?? "<none>"}, " +
        $"Message={message}");

    throw;
}
```

Use `Status` as the primary classification because `ErrorCode` can vary by operation and service version. Log `ErrorCode` and `x-ms-request-id` for diagnosis, but never log the secret value.

| Status | Meaning | Appropriate handling |
|---|---|---|
| `403` | Authenticated but unauthorized | Fix RBAC/access policy; do not retry |
| `404` | Secret or requested version absent | Return an expected “missing” result or fail configuration |
| `409` / `412` | Resource-state or concurrency conflict | Re-read state and reconcile; do not blindly retry |
| `429` | Key Vault throttling | Let the SDK honor `Retry-After`; reduce request rate/cache values |

For RBAC, reads generally require **Key Vault Secrets User**, while secret management requires **Key Vault Secrets Officer** or another appropriately scoped custom role. A `401` is different from `403`: it generally indicates credential or token acquisition/authentication trouble rather than missing data-plane authorization.

Official references: [RequestFailedException](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception), [SecretClientOptions](https://learn.microsoft.com/dotnet/api/azure.security.keyvault.secrets.secretclientoptions), and [Key Vault throttling guidance](https://learn.microsoft.com/azure/key-vault/general/overview-throttling).

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 7 criteria passed. The response provides strong RequestFailedException, retry, conflict, and throttling patterns, but lacks complete diagnostic handling for authorization modes, deleted-secret detection, and purge protection.

### Strengths

- Correctly uses RequestFailedException and exposes Status and ErrorCode.
- Provides practical SecretClientOptions exponential retry configuration.
- Accurately explains 429 behavior after SDK retries and warns against blindly retrying conflicts.
- Correctly notes that SetSecretAsync creates new immutable versions rather than overwriting a version.

### Issues

- RBAC and access-policy authorization models are named but not meaningfully distinguished for diagnosis.
- A 404 is not checked against the deleted-secrets collection or GetDeletedSecretAsync.
- Purge-protection restrictions and their associated failure handling are not covered.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/7)
      - `RequestFailedException` for all Key Vault errors: Pass
      - Extracting `Status` and `ErrorCode` properties: Pass
      - 403 handling: diagnosing RBAC vs. access policy misconfiguration: Fail
      - 404 handling: secret not found vs. deleted secret: Fail
      - 429 throttling: Key Vault rate limits and retry behavior: Pass
      - Soft-delete and purge protection error scenarios: Fail
      - `SecretClientOptions` retry configuration: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 57% | 1.00 | 0.5714 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.5714** | **57.1%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

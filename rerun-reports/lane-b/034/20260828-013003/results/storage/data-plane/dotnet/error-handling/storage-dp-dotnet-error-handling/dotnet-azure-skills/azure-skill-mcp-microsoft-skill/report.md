# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 466.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 466.9s |
| Timestamp | 2026-08-27T17:30:03Z |
| Files Generated | 0 |
| Event Count | 2370 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 444.3s |
| Review | 21.5s |
| **Total** | **466.9s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=154484 out=2501 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 4471ms

**Input:**

```json
{"skill":"azure-storage"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Services

## Services

| Service | Use When | MCP Tools | CLI |
|---------|----------|-----------|-----|
| Blob Storage | Objects, files, backups, static content | `azure__storage` | `az storage blob` |
| File Shares | SMB file shares, lift-and-shift | - | `az storage file` |
| Queue Storage | Async messaging, task queues | - | `az storage queue` |
| Table Storage | NoSQL key-value (consider Cosmos DB) | - | `az storage table` |
| Data Lake | Big data analytics, hierarchical namespace | - | `az storage fs` |

## MCP Server (Preferred)

When Azure MCP is enabled:

- `azure__storage` with command `storage_account_list` - List storage accounts
- `azure__storage` with command `storage_container_list` - List containers in account
- `azure__storage` with command `storage_blob_list` - List blobs in container
- `azure__storage` with command `storage_blob_get` - Download blob content
- `azure__storage` with command `storage_blob_put` - Upload blob content

**If Azure MCP is not enabled:** Run `/azure:setup` or enable via `/mcp`.

## CLI Fallback

```bash
# List storage accounts
az storage account list --output table

# List containers
az storage container list --account-name ACCOUNT --output table

# List blobs
az storage blob list --account-name ACCOUNT --container-name CONTAINER --output table

# Download blob
az storage blob download --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH

# Upload blob
az storage blob upload --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH
```

## Storage Account Tiers

| Tier | Use Case | Performance |
|------|----------|-------------|
| Standard | General purpose, backup | Milliseconds |
| Premium | Databases, high IOPS | Sub-millisecond |

## Blob Access Tiers

| Tier | Access Frequency | Cost |
|------|-----------------|------|
| Hot | Frequent | Higher storage, lower access |
| Cool | Infrequent (30+ days) | Lower storage, higher ac
... (truncated)
```

### ✅ 2. view — 188028ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
@@ -1,115 +1,115 @@
 # Azure Storage SDK Usage
 
 SDK packages and quick start examples for Azure Storage services.
 
 ## Storage SDKs by Language
 
 | Language | Blob | Queue | File Share | Data Lake |
 |----------|------|-------|------------|----------|
 | .NET | `Azure.Storage.Blobs` | `Azure.Storage.Queues` | `Azure.Storage.Files.Shares` | `Azure.Storage.Files.DataLake` |
 | Java | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | JavaScript | `@azure/storage-blob` | `@azure/storage-queue` | `@azure/storage-file-share` | `@azure/storage-file-datalake` |
 | Python | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | Go | `azblob` | `azqueue` | `azfile` | `azdatalake` |
 | Rust | `azure_storage_blob` | `azure_storage_queue` | - | - |
 
 ## Installation Commands
 
 | Language | Install Blob SDK + Identity |
 |----------|-----------------------------|
 | .NET | `dotnet add package Azure.Storage.Blobs` `dotnet add package Azure.Identity` |
 | Java | Maven: `com.azure:azure-storage-blob` `com.azure:azure-identity` |
 | JavaScript | `npm install @azure/storage-blob @azure/identity` |
 | Python | `pip install azure-storage-blob azure-identity` |
 | Go | `go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob github.com/Azure/azure-sdk-for-g
... (truncated)
```

### ✅ 3. view — 186232ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks environment variables, CLI tokens, and oth
... (truncated)
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 188033ms

**Input:**

```json
{"intent":"Provide current best practices for idiomatic .NET error handling with Azure.Storage.Blobs, including RequestFailedException status/error code handling and retry behavior for 404, 403, and 429.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 190296ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Storage.Blobs .NET RequestFailedException handling, including Status, ErrorCode, catch filters, and retry configuration for HTTP 404, 403, and 429.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_fetch — 190292ms

**Input:**

```json
{"max_length":6000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.requestfailedexception?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy:
---
layout: Conceptual
title: Implement a retry policy using the Azure Storage client library for .NET - Azure Storage | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy
breadcrumb_path: /azure/bread/toc.json
feedback_help_link_url: https://learn.microsoft.com/answers/tags/125/azure-blob-storage/
feedback_help_link_type: get-help-at-qna
feedback_product_url: https://feedback.azure.com/d365community/forum/a8bb4a47-3525-ec11-b6e6-000d3a4f0f84
feedback_system: Standard
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
uhfHeaderId: azure
ms.suite: office
adobe-target: true
learn_banner_products:
- azure
manager: akashdubey
ms.reviewer: akashdubey-ms
description: Learn about retry policies and how to implement them for Blob Storage. This article helps you set up a retry policy for Blob Storage requests using the Azure Storage client library for .NET.
author: stevenmatthew
ms.author: shaas
ms.service: azure-blob-storage
ms.topic: how-to
ms.date: 2024-11-19T00:00:00.0000000Z
ms.custom: devx-track-dotnet, devguide-csharp
locale: en-us
document_id: 3b78918f-3432-e8b0-aa74-df6998979c4d
document_version_independent_id: 4434fc07-9da4-1db4-6bf5-61401ecf97bc
updated_at: 2025-08-15T05:11:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-pr/blob/live/articles/storage/blobs/storage-retry-policy.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-pr/blob/9656d6c3a103d19d44c6fcc49b1391adc8590214/articles/storage/blobs/storage-retry-policy.md
git_commit_id: 9656d6c3a103d19d44c6fcc49b1391adc8590214
site_name: Docs
depot_name: Azure.azure-documents
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/Azure.azure-documents/{branchName}{pdfName}
word_count: 854
asset_id: storage/blobs/storage-retry-policy
moniker_rang
... (truncated)
```

### ✅ 7. web_fetch — 188931ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.requestfailedexception?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: RequestFailedException Class (Azure) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.requestfailedexception?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.RequestFailedException
namespace: Azure
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.RequestFailedException
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 72ae24bf-d6bd-72f9-1431-8c972a27e080
document_version_independent_id: 3db31ccd-e355-a729-5860-6f430b066c08
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure/RequestFailedException.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure/RequestFailedException.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.RequestFailedException
description: 'An exception thrown when service request fails. '
toc_rel: azure_ref_toc/_splitted/azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang:
... (truncated)
```

## Copilot Response

I’m grounding the patterns in the current Azure SDK guidance, especially retry behavior and authentication-related 403 handling.

The key distinction is that all three service responses surface as `RequestFailedException`; status identifies the HTTP class, while `ErrorCode` provides the storage-specific reason.

`Azure.Storage.Blobs` reports Blob service failures as `Azure.RequestFailedException`. Use `Status` for the HTTP status and `ErrorCode` for the storage-specific code.

| Scenario | `Status` | Typical `ErrorCode` |
|---|---:|---|
| Container missing | `404` | `ContainerNotFound` |
| Access denied | `403` | `AuthorizationPermissionMismatch`, `AuthorizationFailure`, or `AuthenticationFailed` |
| Throttled | `429` | Often `ServerBusy`; use the status as the authoritative signal |

```csharp
using Azure;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;

public async Task DownloadBlobAsync(
    BlobClient blobClient,
    Stream destination,
    CancellationToken cancellationToken)
{
    try
    {
        await blobClient.DownloadToAsync(destination, cancellationToken);
    }
    catch (RequestFailedException ex)
        when (ex.Status == 404 &&
              ex.ErrorCode == BlobErrorCode.ContainerNotFound.ToString())
    {
        // Translate an expected condition into an application-specific result.
        throw new InvalidOperationException(
            "The blob container does not exist.", ex);
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        // ErrorCode explains whether authentication, RBAC, SAS permissions,
        // or another authorization condition caused the rejection.
        Console.Error.WriteLine(
            $"Blob access denied. Status={ex.Status}, " +
            $"ErrorCode={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        throw; // Preserve the original exception and stack trace.
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        // The SDK has normally exhausted its configured automatic retries
        // before this exception reaches application code.
        Console.Error.WriteLine(
            $"Blob request throttled. Status={ex.Status}, " +
            $"ErrorCode={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        throw;
    }
    catch (RequestFailedException ex)
    {
        // Other Blob service failures, such as 409, 412, or 5xx.
        Console.Error.WriteLine(
            $"Blob request failed. Status={ex.Status}, " +
            $"ErrorCode={ex.ErrorCode}, Message={ex.Message}, " +
            $"ClientRequestId={ex.ClientRequestId}");

        throw;
    }
}
```

For code that only needs to determine whether a container exists, use `ExistsAsync` rather than using exceptions as normal control flow:

```csharp
BlobContainerClient container = serviceClient.GetBlobContainerClient("documents");

Response<bool> result = await container.ExistsAsync(cancellationToken);

if (!result.Value)
{
    // The container does not exist.
}
```

If you need to distinguish an absent container using an operation such as `GetPropertiesAsync`, catch `ContainerNotFound` explicitly:

```csharp
try
{
    await container.GetPropertiesAsync(cancellationToken: cancellationToken);
}
catch (RequestFailedException ex)
    when (ex.Status == 404 &&
          ex.ErrorCode == BlobErrorCode.ContainerNotFound.ToString())
{
    Console.WriteLine("Container not found.");
}
```

Configure retries on the client rather than writing a manual retry loop around every call. The SDK honors server-provided `Retry-After` information:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(10),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(100)
    }
};

var serviceClient = new BlobServiceClient(
    new Uri("https://<account-name>.blob.core.windows.net"),
    new DefaultAzureCredential(), // Appropriate for local development.
    options);
```

Important details:

- **Do not manually retry 403 or 404.** They are normally permanent until credentials, permissions, or resource state changes.
- **429 is transient.** Let the SDK retry it first; a caught 429 means those retries were exhausted. Avoid immediately starting another aggressive retry loop.
- Blob Storage may also indicate overload with **503**, so applications concerned with throttling should treat retryable 5xx responses according to the configured SDK policy.
- A 403 does not have one universal error code. Inspect `ErrorCode` to distinguish invalid authentication from insufficient RBAC/SAS permissions.
- Use `throw;`, not `throw ex;`, when rethrowing.
- Let `OperationCanceledException` propagate separately; cancellation is not a Blob service failure.
- Avoid logging SAS tokens, account keys, authorization headers, or full request URLs containing credentials.

References: [RequestFailedException API](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception) and [Blob Storage retry policy for .NET](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy).

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response thoroughly covers RequestFailedException, exception metadata, retries, and practical logging, but lacks explicit 409 handling and conditional failure examples for ETags and leases.

### Strengths

- Idiomatic filtered RequestFailedException catch blocks.
- Correct extraction of Status and ErrorCode.
- Appropriate BlobClientOptions retry configuration.
- Useful diagnostic context and guidance for protecting credentials in logs.

### Issues

- No dedicated handling or explanation for HTTP 409 conflicts.
- No examples for ETag precondition failures or lease-related failures.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/6)
      - `RequestFailedException` as the primary exception type: Pass
      - Extracting `Status` and `ErrorCode` from the exception: Pass
      - Handling specific HTTP status codes (404, 403, 409, 429): Fail
      - Retry policy configuration via `BlobClientOptions`: Pass
      - Conditional request failures (ETags, leases): Fail
      - Logging and diagnostics for troubleshooting: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 67% | 1.00 | 0.6667 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6667** | **66.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 303.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 3/6 |
| Duration | 303.4s |
| Timestamp | 2026-08-27T17:24:58Z |
| Files Generated | 0 |
| Event Count | 2080 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 279.2s |
| Review | 23.1s |
| **Total** | **303.4s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=152276 out=2092 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 3462ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 188307ms

**Input:**

```json
{"intent":"Provide idiomatic .NET Azure.Storage.Blobs error-handling examples for 404 container-not-found, 403 access-denied, and 429 throttling, including RequestFailedException status and error-code extraction.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 188308ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Storage.Blobs .NET RequestFailedException, BlobErrorCode values, retry behavior, and handling HTTP 404, 403, and 429 responses.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 188302ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md,sdk-usage.md,auth-best-practices.md}"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-storage\references\sdk-usage.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-storage\references\auth-best-practices.md
```

### ✅ 5. view — 6702ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
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

### ✅ 6. view — 6381ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
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

### ✅ 7. web_fetch — 7707ms

**Input:**

```json
{"max_length":6000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.requestfailedexception?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.models.bloberrorcode?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobErrorCode Struct (Azure.Storage.Blobs.Models) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.models.bloberrorcode?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Models.BlobErrorCode
namespace: Azure.Storage.Blobs.Models
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Storage.Blobs.Models.BlobErrorCode
api_location:
- Azure.Storage.Blobs.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 2e616006-277a-b7ca-704a-9dac55550fd8
document_version_independent_id: 1d1187f4-7724-c221-8390-22a9baf110b3
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Models/BlobErrorCode.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Models/BlobErrorCode.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: struct
ms.assetid: Azure.Storage.Blobs.Models.BlobErrorCode
description: 'Error codes returned by t
... (truncated)
```

### ✅ 8. web_fetch — 7377ms

**Input:**

```json
{"max_length":6000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.models.bloberrorcode?view=azure-dotnet"}
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

### ✅ 9. web_fetch — 7034ms

**Input:**

```json
{"max_length":6000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}
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

## Copilot Response

`Azure.Storage.Blobs` reports service-side failures as `Azure.RequestFailedException`. The two key properties are:

- `Status`: HTTP status code as an `int`
- `ErrorCode`: Azure Storage’s service-specific error code, such as `ContainerNotFound`

```csharp
using Azure;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;

BlobContainerClient containerClient =
    blobServiceClient.GetBlobContainerClient("documents");

try
{
    await containerClient.GetPropertiesAsync();
}
catch (RequestFailedException ex)
    when (ex.Status == 404 &&
          ex.ErrorCode == BlobErrorCode.ContainerNotFound.ToString())
{
    Console.Error.WriteLine(
        $"Container does not exist. Status={ex.Status}, Code={ex.ErrorCode}");
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // Possible codes include AuthorizationPermissionMismatch,
    // AuthorizationFailure, or AuthenticationFailed.
    Console.Error.WriteLine(
        $"Access denied. Status={ex.Status}, Code={ex.ErrorCode}");

    throw; // Usually not transient; surface or correct credentials/RBAC.
}
catch (RequestFailedException ex) when (ex.Status == 429)
{
    Console.Error.WriteLine(
        $"Request throttled after SDK retries. Status={ex.Status}, Code={ex.ErrorCode}");

    throw; // Preserve stack trace; let an outer resilience policy decide what to do.
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Blob Storage request failed. Status={ex.Status}, " +
        $"Code={ex.ErrorCode ?? "<none>"}, Message={ex.Message}");

    throw;
}
```

For a broader operation where every `404` means “not found,” filtering only by status is reasonable:

```csharp
try
{
    await containerClient.DeleteAsync();
}
catch (RequestFailedException ex) when (ex.Status == 404)
{
    // Idempotent delete: the desired state already exists.
}
```

When distinguishing resources matters, check both status and error code:

```csharp
catch (RequestFailedException ex)
    when (ex.Status == 404 &&
          ex.ErrorCode == BlobErrorCode.BlobNotFound.ToString())
{
    // The blob is missing.
}
catch (RequestFailedException ex)
    when (ex.Status == 404 &&
          ex.ErrorCode == BlobErrorCode.ContainerNotFound.ToString())
{
    // The containing container is missing.
}
```

For existence checks, prefer the SDK’s non-exceptional API. `ExistsAsync` returns `false` for a normal not-found response but still throws for failures such as authorization errors:

```csharp
try
{
    Response<bool> result = await containerClient.ExistsAsync();

    if (!result.Value)
    {
        Console.WriteLine("Container not found.");
    }
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    Console.Error.WriteLine($"Access denied: {ex.ErrorCode}");
    throw;
}
```

Configure retries on the client rather than writing a manual retry loop around every operation:

```csharp
using Azure.Core;
using Azure.Storage.Blobs;

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(30),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(100)
    }
};

var blobServiceClient = new BlobServiceClient(
    new Uri("https://<account>.blob.core.windows.net"),
    credential,
    options);
```

The SDK automatically retries transient responses, including `429`, and honors the service’s `Retry-After` header. Therefore, catching `429` normally means all configured SDK retries have already been exhausted.

Also account for Blob Storage commonly signaling overload with `503 ServerBusy`, not only `429`:

```csharp
catch (RequestFailedException ex)
    when (ex.Status == 429 ||
          (ex.Status == 503 &&
           ex.ErrorCode == BlobErrorCode.ServerBusy.ToString()))
{
    Console.Error.WriteLine(
        $"Storage remains throttled or busy after retries: " +
        $"Status={ex.Status}, Code={ex.ErrorCode}");
    throw;
}
```

Avoid catching plain `Exception` for expected storage failures, comparing localized `Message` text, or manually constructing a replacement exception. Catch `RequestFailedException`, use `Status` plus `ErrorCode`, and use `throw;` when propagating it so the original stack trace and diagnostics are preserved.

## Reviewer Notes (LLM-as-Judge)

### Summary

3 of 6 criteria passed. The response handles the primary exception, extracts service error details, and configures retries well, but omits 409 handling, conditional failures, and substantive diagnostics guidance.

### Strengths

- Correctly centers RequestFailedException and preserves stack traces with throw;.
- Uses both HTTP status and storage-specific error codes to distinguish failures.
- Shows idiomatic BlobClientOptions retry configuration and explains exhausted retries.
- Adds useful coverage for 503 ServerBusy and recommends ExistsAsync for existence checks.

### Issues

- No handling for HTTP 409 Conflict.
- No coverage of ETag or lease-related conditional request failures.
- Troubleshooting uses basic console output rather than structured logging and Azure SDK diagnostic facilities.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (3/6)
      - `RequestFailedException` as the primary exception type: Pass
      - Extracting `Status` and `ErrorCode` from the exception: Pass
      - Handling specific HTTP status codes (404, 403, 409, 429): Fail
      - Retry policy configuration via `BlobClientOptions`: Pass
      - Conditional request failures (ETags, leases): Fail
      - Logging and diagnostics for troubleshooting: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 50% | 1.00 | 0.5000 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.5000** | **50.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

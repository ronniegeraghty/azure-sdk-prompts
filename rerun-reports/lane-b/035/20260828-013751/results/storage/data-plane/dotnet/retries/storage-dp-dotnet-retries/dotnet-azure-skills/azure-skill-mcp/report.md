# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ✅ PASSED | **Duration:** 412.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ✅ PASSED |
| Score | 8/8 |
| Duration | 412.9s |
| Timestamp | 2026-08-27T17:37:51Z |
| Files Generated | 0 |
| Event Count | 4083 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 382.4s |
| Review | 29.3s |
| **Total** | **412.9s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=269598 out=5878 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3477ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 193138ms

**Input:**

```json
{"intent":"Get current .NET Azure SDK best practices for Azure.Storage.Blobs retry configuration, cancellation-based operation timeouts, transient error handling, custom HttpPipelinePolicy usage, and circuit breakers.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 193141ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Storage.Blobs BlobClientOptions RetryOptions in .NET, default retryable HTTP status codes, NetworkTimeout semantics, custom per-call pipeline policies, RequestFailedException classification, and circuit breaker guidance.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. view — 193131ms

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

### ✅ 5. view — 191575ms

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

### ✅ 6. glob — 190140ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. web_search (via github-mcp-server) — 46054ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure.Storage.Blobs .NET configure retry policy RetryOptions default retry status codes 408 429 500 502 503 504 NetworkTimeout"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 8. web_search (via github-mcp-server) — 45700ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure.Core.ClientOptions RetryPolicy custom RetryPolicy HttpPipelinePolicy BlobClientOptions .NET"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The Azure.Storage.Blobs .NET client library allows you to configure retry policies for blob operations via the `RetryOptions` class, set in `BlobClientOptions.Retry`. The retry policy is designed to handle transient faults like network issues or brief service outages to improve resiliency.\n\n### Default Retry Status Codes\nBy default, the retry policy for Azure.Storage.Blobs automatically retries operations that fail with transient error status codes, which include:\n- 408 (Request Timeout)\n- 429 (Too Many Requests)\n- 500 (Internal Server Error)\n- 502 (Bad Gateway)\n- 503 (Service Unavailable)\n- 504 (Gateway Timeout)\n\n**NetworkTimeout** is not an HTTP status code but refers to the policy’s timeout for each network operation; the default is 100 seconds.\n\n### Default RetryOptions in .NET\n- **Delay**: 0.8 seconds (delay between retries for exponential backoff)\n- **MaxDelay**: 1 minute (maximum allowed delay between retries)\n- **MaxRetries**: 5 (before giving up)\n- **Mode**: Exponential (backoff strategy)\n- **NetworkTimeout**: 100 seconds per operation\n\nThe retryable transient failures, as described above, are retried automatically via these policies when you use the Azure.Storage.Blobs .NET client, unless you override them with customized settings【3:0†source】【3:2†source】.\n\n### References\n- [Configure retry policy using Azure.Storage.Blobs .NET](https://learn.microsoft.com/en-us/azure/storage/blobs/retry-policy-dotnet?tabs=azure-cli)【3:0†source】\n\nLet me know if you need a code example for configuration!","annotations":[{"text":"","start_index":1260,"end_index":1272,"url_citation":{"title":"Implement a retry policy using the Azure Storage client library for ...","url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}},{"text":"","start_index":1272,"end_index":1284,"url_citation":{"title":"BlobClientOptions Class (Azure.Storage.Blobs) - Azure for .NET ...","url":"ht
... (truncated)
```

### ✅ 9. web_fetch — 45388ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: RetryOptions Class (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.RetryOptions
namespace: Azure.Core
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
- Azure.Core.RetryOptions
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 15e6555f-f376-e6f7-03f5-7bb9884708bc
document_version_independent_id: 44778f21-538e-09b6-e9f4-0a3bacbdddef
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/RetryOptions.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core/RetryOptions.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Core.RetryOptions
description: 'The set of options that can be specified to influence how retry attempts are made, and a failure is eligible to be retried. '
toc_rel: azure_ref_toc/_splitted/azure.core/toc.json
feedback_help_link_type: ''
feed
... (truncated)
```

### ✅ 10. web_fetch — 45076ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions?view=azure-dotnet"}
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

### ✅ 11. web_fetch — 44740ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: ClientOptions.RetryPolicy Property (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.ClientOptions.RetryPolicy*
namespace: Azure.Core
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
- Azure.Core.ClientOptions.RetryPolicy
- Azure.Core.ClientOptions.get_RetryPolicy
- Azure.Core.ClientOptions.set_RetryPolicy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 4600e65f-3027-abe5-f2ea-aa6125fe7d79
document_version_independent_id: b0e545bb-9918-8c2f-f6a2-85ac3dd21f2f
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/ClientOptions.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core/ClientOptions.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: property
ms.assetid: Azure.Core.ClientOptions.RetryPolicy*
description: "Gets or sets the policy to use for re
... (truncated)
```

### ✅ 12. github-mcp-server-get_file_contents (via github-mcp-server) — 35211ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: c8102635d8888d976a9ec9269e44a09f8595bc44)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using Azure.Core.Pipeline;
using Microsoft.Extensions.Configuration;

namespace Azure.Core
{
    /// <summary>
    /// Base type for all client option types, exposes various common client options like <see cref="Diagnostics"/>, <see cref="Retry"/>, <see cref="Transport"/>.
    /// </summary>
    public abstract class ClientOptions
    {
        private HttpPipelineTransport _transport;
        internal bool IsCustomTransportSet { get; private set; }

        /// <summary>
        /// Gets the default set of <see cref="ClientOptions"/>. Changes to the <see cref="Default"/> options would be reflected
        /// in new instances of <see cref="ClientOptions"/> type created after changes to <see cref="Default"/> were made.
        /// </summary>
        public static ClientOptions Default { get; private set; } = new DefaultClientOptions();

        // For testing
        internal static void ResetDefaultOptions()
        {
            Default = new DefaultClientOptions();
        }

        /// <summary>
        /// Creates a new instance of <see cref="ClientOptions"/>.
        /// </summary>
        protected ClientOptions() : this(Default, null)
        {
        }

        /// <summary>
        /// Creates a new instance of <see cref="ClientOptions"/> with the specified <see cref="DiagnosticsOptions"/>.
        /// </summary>
        /// <param name="diagnostics"><see cref="DiagnosticsOptions"/> to be used for <see cref="Diagnostics"/>.</param>
        protected ClientOptions(DiagnosticsOptions? diagnostics)
            : this(Default, diagnostics)
        {
        }

        internal ClientOptions(ClientOptions? clientOptions, DiagnosticsOptions? diagnostics)
        {
       
... (truncated)
```

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server) — 34786ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/ClientOptions.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: a729cb2a06b751b0f53499f0ca00d31ac7c919b1)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Runtime.ExceptionServices;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core.Diagnostics;

namespace Azure.Core.Pipeline
{
    /// <summary>
    /// Represents a policy that can be overridden to customize whether or not a request will be retried and how long to wait before retrying.
    /// </summary>
    public class RetryPolicy : HttpPipelinePolicy
    {
        private readonly int _maxRetries;

        /// <summary>
        /// Gets the delay to use for computing the interval between retry attempts.
        /// </summary>
        private readonly DelayStrategy _delayStrategy;

        /// <summary>
        /// Initializes a new instance of the <see cref="RetryPolicy"/> class.
        /// </summary>
        /// <param name="maxRetries">The maximum number of retries to attempt.</param>
        /// <param name="delayStrategy">The delay to use for computing the interval between retry attempts.</param>
        public RetryPolicy(int maxRetries = RetryOptions.DefaultMaxRetries, DelayStrategy? delayStrategy = default)
        {
            _maxRetries = maxRetries;
            _delayStrategy = delayStrategy ?? DelayStrategy.CreateExponentialDelayStrategy();
        }

        /// <summary>
        /// This method can be overridden to take full control over the retry policy. If this is overridden and the base method isn't called,
        /// it is the implementer's responsibility to populate the <see cref="HttpMessage.ProcessingContext"/> property.
        /// This method will only be called for async methods.
        /// </summary>
        /// <param name="message">The <see cref="HttpMessage"/> this policy would be applied to.</param>
        /// <param name="pipeline">The set of <see cre
... (truncated)
```

### ✅ 14. web_search (via github-mcp-server) — 35576ms

**Input:**

```json
{"query":"Polly v8 AddCircuitBreaker ResiliencePipelineBuilder ShouldHandle PredicateBuilder RequestFailedException C# example"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s a **C# example using Polly v8** showing how to build a circuit breaker with `ResiliencePipelineBuilder`, configure `ShouldHandle` using `PredicateBuilder` to handle `RequestFailedException`, and explain each piece:\n\n### Step-by-Step Example\n\nFirst, make sure you have the necessary Polly v8 package:\n```sh\ndotnet add package Polly.Core\n```\n\nThen implement your pipeline:\n\n```csharp\nusing Polly;\nusing Polly.CircuitBreaker;\nusing Polly.Extensions;\nusing Polly.Predicates;\nusing System;\nusing Azure;\nusing System.Net.Http;\n\n// 1. Define your circuit breaker options\nvar breakerOptions = new CircuitBreakerStrategyOptions\n{\n    FailureRatio = 0.5, // Break when >= 50% handled calls fail\n    SamplingDuration = TimeSpan.FromSeconds(10), // Check failure ratio within a 10s window\n    MinimumThroughput = 4, // Sample at least 4 executions\n    BreakDuration = TimeSpan.FromSeconds(30), // Break for 30 seconds\n    ShouldHandle = new PredicateBuilder()\n        .Handle<RequestFailedException>() // Break on RequestFailedException (Azure SDK)\n    // Add additional .Handle<T>() or .HandleResult() if needed\n};\n\n// 2. Build your resilience pipeline with the circuit breaker\nvar pipeline = new ResiliencePipelineBuilder()\n    .AddCircuitBreaker(breakerOptions)\n    .Build();\n\n// 3. Use the pipeline for your calls\ntry\n{\n    await pipeline.ExecuteAsync(async token =>\n    {\n        // Simulate call to Azure resource that could throw RequestFailedException\n        await SomeAzureCallAsync();\n    });\n}\ncatch (BrokenCircuitException)\n{\n    // Handle open circuit, e.g., log or return fallback response\n}\ncatch (RequestFailedException)\n{\n    // Handle actual Azure failure\n}\n\nasync Task SomeAzureCallAsync()\n{\n    // Imagine this calls Azure and can throw RequestFailedException\n    throw new RequestFailedException(\"Simulated Azure failure\");\n}\n```\n\n**Key notes:**\n- `PredicateBuilder().Handle<Re
... (truncated)
```

### ❌ 15. web_fetch — 34407ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/configure-retry-policy"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/configure-retry-policy - status code 404
```

### ✅ 16. web_fetch — 34402ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: RetryPolicy Class (Azure.Core.Pipeline) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.Pipeline.RetryPolicy
namespace: Azure.Core.Pipeline
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
- Azure.Core.Pipeline.RetryPolicy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: e36854d9-73a0-4858-e4fc-604b07c03719
document_version_independent_id: 82fa277a-e6fd-df80-ffa2-193c3a4b5102
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core.Pipeline/RetryPolicy.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core.Pipeline/RetryPolicy.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Core.Pipeline.RetryPolicy
description: 'Represents a policy that can be overriden to customize whether or not a request will be retried and how long to wait before retrying. '
to
... (truncated)
```

### ✅ 17. github-mcp-server-get_file_contents (via github-mcp-server) — 4658ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/DelayStrategy.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 89f71bd86b7924997b5f722b70748a88d3c1b501)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using Azure.Core.Pipeline;

namespace Azure.Core
{
    /// <summary>
    /// An abstraction to control delay behavior.
    /// </summary>
    public abstract class DelayStrategy
    {
        internal const double DefaultJitterFactor = 0.2;

        private readonly Random _random = new ThreadSafeRandom();
        private readonly double _minJitterFactor;
        private readonly double _maxJitterFactor;
        private readonly TimeSpan _maxDelay;

        /// <summary>
        /// Constructs a new instance of <see cref="DelayStrategy"/>. This constructor can be used by derived classes to customize the jitter factor and max delay.
        /// </summary>
        /// <param name="maxDelay">The max delay value to apply on an individual delay.</param>
        /// <param name="jitterFactor">The jitter factor to apply to each delay. For example, if the delay is 1 second with a jitterFactor of 0.2, the actual
        /// delay used will be a random double between 0.8 and 1.2. If set to 0, no jitter will be applied.</param>
        protected DelayStrategy(TimeSpan? maxDelay = default, double jitterFactor = DefaultJitterFactor)
        {
            _minJitterFactor = 1.0 - jitterFactor;
            _maxJitterFactor = 1.0 + jitterFactor;

            // use same defaults as RetryOptions
            _maxDelay = maxDelay ?? RetryOptions.DefaultMaxDelay;
        }

        /// <summary>
        /// Constructs an exponential delay with jitter.
        /// </summary>
        /// <param name="initialDelay">The initial delay to use.</param>
        /// <param name="maxDelay">The maximum delay to use.</param>
        /// <returns>The <see cref="DelayStrategy"/> instance.</returns>
        public static DelayStrategy CreateExponentialDelayStrategy(
            TimeSpan? initialDelay = default,

... (truncated)
```

### ✅ 18. web_fetch — 4314ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.delaystrategy.createexponentialdelaystrategy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy.shouldretryasync?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: RetryPolicy.ShouldRetryAsync(HttpMessage, Exception) Method (Azure.Core.Pipeline) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy.shouldretryasync?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.Pipeline.RetryPolicy.ShouldRetryAsync*
namespace: Azure.Core.Pipeline
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
- Azure.Core.Pipeline.RetryPolicy.ShouldRetryAsync
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: f4e0fb9f-9e5f-0981-2c06-27a9df563198
document_version_independent_id: f7025d53-aa32-39df-91c1-4635a58881f4
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core.Pipeline/RetryPolicy.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core.Pipeline/RetryPolicy.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: method
ms.assetid: Azure.Core.Pipeline.RetryPolicy.ShouldRetryAsync*
description: 'This met
... (truncated)
```

### ✅ 19. web_fetch — 3964ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy.shouldretryasync?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.delaystrategy.createexponentialdelaystrategy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: DelayStrategy.CreateExponentialDelayStrategy Method (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.delaystrategy.createexponentialdelaystrategy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.DelayStrategy.CreateExponentialDelayStrategy*
namespace: Azure.Core
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
- Azure.Core.DelayStrategy.CreateExponentialDelayStrategy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 1c5e276b-a6d6-2a0f-dd98-5845b30cf730
document_version_independent_id: f3536192-f4e3-dbbd-8b1f-106b1d6d0d66
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/DelayStrategy.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core/DelayStrategy.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: method
ms.assetid: Azure.Core.DelayStrategy.CreateExponentialDelayStrategy*
description: 'Constructs an
... (truncated)
```

### ✅ 20. web_fetch — 3629ms

**Input:**

```json
{"max_length":12000,"url":"https://www.pollydocs.org/strategies/circuit-breaker.html"}
```

**Output:**

```
Contents of https://www.pollydocs.org/strategies/circuit-breaker.html:
# Circuit breaker resilience strategy

## About

* * *

The circuit breaker **reactive** resilience strategy shortcuts the execution if the underlying resource is detected as unhealthy. The detection process is done via sampling. If the sampled executions' failure-success ratio exceeds a predefined threshold then a circuit breaker will prevent any new executions by throwing a `BrokenCircuitException`. After a preset duration the circuit breaker performs a probe, because the assumption is that this period was enough for the resource to self-heal. Depending on the outcome of the probe, the circuit will either allow new executions or continue to block them. If an execution is blocked by the circuit breaker, the thrown exception may indicate the amount of time executions will continue to be blocked through its `RetryAfter` property.

##### Note

Be aware that the Circuit Breaker strategy [rethrows all exceptions](https://github.com/App-vNext/Polly/wiki/Circuit-Breaker#exception-handling), including those that are handled. A Circuit Breaker's role is to monitor faults and break the circuit when a certain threshold is reached; it does not manage retries. Combine the Circuit Breaker with a Retry strategy if needed.

## Usage

```
// Circuit breaker with default options.
// See https://www.pollydocs.org/strategies/circuit-breaker#defaults for defaults.
var optionsDefaults = new CircuitBreakerStrategyOptions();

// Circuit breaker with customized options:
// The circuit will break if more than 50% of actions result in handled exceptions,
// within any 10-second sampling duration, and at least 8 actions are processed.
var optionsComplex = new CircuitBreakerStrategyOptions
{
    FailureRatio = 0.5,
    SamplingDuration = TimeSpan.FromSeconds(10),
    MinimumThroughput = 8,
    BreakDuration = TimeSpan.FromSeconds(30),
    ShouldHandle = new PredicateBuilder().Handle<SomeExceptionType>()
};

// Circuit break
... (truncated)
```

## Copilot Response

I’m grounding the examples in the current Azure SDK retry model, including the distinction between transport failures, HTTP failures, cancellation, and per-operation policies.

The built-in retry configuration is straightforward; the key design choice is to keep ordinary and high-value operations on separate client pipelines so their retry budgets do not multiply. I’m checking the exact current API surface and default status set before giving the code.

Configure retries on `BlobClientOptions`, then enforce an **end-to-end deadline** with a cancellation token. `NetworkTimeout` limits each individual network attempt; it does not bound the total upload, including retries and multipart transfers.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

var accountUri = new Uri(
    Environment.GetEnvironmentVariable("AZURE_STORAGE_BLOB_URI")
    ?? throw new InvalidOperationException("AZURE_STORAGE_BLOB_URI is required."));

TokenCredential credential = new ManagedIdentityCredential();

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,

        // Five retries after the initial request: up to six attempts total.
        MaxRetries = 5,

        // Base delay used by exponential backoff.
        Delay = TimeSpan.FromSeconds(1),

        // Caps SDK-calculated delays unless the service supplies Retry-After.
        MaxDelay = TimeSpan.FromSeconds(30),

        // Timeout for each individual network operation/attempt.
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

var serviceClient = new BlobServiceClient(accountUri, credential, options);
var blobClient = serviceClient
    .GetBlobContainerClient("documents")
    .GetBlobClient("report.pdf");
```

The exponential strategy includes jitter. When Azure Storage returns a `Retry-After` header, the SDK honors that server-provided delay.

## End-to-end operation timeout

Use `CancellationTokenSource.CancelAfter` to limit the complete logical upload:

```csharp
static async Task UploadWithDeadlineAsync(
    BlobClient blob,
    Stream content,
    TimeSpan timeout,
    CancellationToken callerToken)
{
    using var deadline = CancellationTokenSource.CreateLinkedTokenSource(callerToken);
    deadline.CancelAfter(timeout);

    try
    {
        await blob.UploadAsync(
            content,
            overwrite: true,
            cancellationToken: deadline.Token);
    }
    catch (OperationCanceledException) when (!callerToken.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"Blob upload exceeded its {timeout} operation deadline.");
    }
}

// The entire upload, including SDK retries, is limited to two minutes.
await UploadWithDeadlineAsync(
    blobClient,
    contentStream,
    TimeSpan.FromMinutes(2),
    applicationStoppingToken);
```

Without the outer cancellation deadline, five retries can take considerably longer than `NetworkTimeout`, because that setting applies independently to network operations and retry delays are added between attempts.

## Transient versus authentication errors

The default policy already avoids retrying `401` and `403`. You should preserve that distinction in logging, circuit-breaker predicates, and any application-level resilience logic:

```csharp
using System.Net.Http;

static bool IsTransientStatus(int status) =>
    status is 408 or 429 or 500 or 502 or 503 or 504;

try
{
    await UploadWithDeadlineAsync(
        blobClient,
        contentStream,
        TimeSpan.FromMinutes(2),
        applicationStoppingToken);
}
catch (AuthenticationFailedException ex)
{
    // Credential acquisition failed. Fix identity configuration; don't retry here.
    Console.Error.WriteLine($"Credential failure: {ex.Message}");
    throw;
}
catch (RequestFailedException ex) when (ex.Status is 401 or 403)
{
    // Token was rejected or the identity lacks RBAC/data-plane permission.
    Console.Error.WriteLine(
        $"Blob authorization failed: HTTP {ex.Status}, code {ex.ErrorCode}");
    throw;
}
catch (RequestFailedException ex) when (IsTransientStatus(ex.Status))
{
    // The SDK has already exhausted its configured retries.
    Console.Error.WriteLine(
        $"Transient storage failure after retries: HTTP {ex.Status}, code {ex.ErrorCode}");
    throw;
}
catch (Exception ex) when (ex is HttpRequestException or IOException)
{
    // A transport failure remained after the SDK retry pipeline completed.
    Console.Error.WriteLine($"Network failure after retries: {ex.Message}");
    throw;
}
```

Do not add an unconditional second retry loop around the SDK. Otherwise, five SDK retries combined with five outer retries can produce up to 36 HTTP attempts.

## A separate policy for high-value operations

Azure clients are thread-safe and intended to be reused. Create and reuse a second client with its own pipeline rather than changing options on each request:

```csharp
var highValueOptions = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 8,
        Delay = TimeSpan.FromSeconds(2),
        MaxDelay = TimeSpan.FromMinutes(1),
        NetworkTimeout = TimeSpan.FromSeconds(45)
    }
};

var highValueServiceClient =
    new BlobServiceClient(accountUri, credential, highValueOptions);

var highValueBlob = highValueServiceClient
    .GetBlobContainerClient("critical-documents")
    .GetBlobClient("settlement-report.pdf");

await UploadWithDeadlineAsync(
    highValueBlob,
    highValueStream,
    TimeSpan.FromMinutes(5),
    applicationStoppingToken);
```

This keeps the normal retry budget at five while granting selected operations a larger budget. Only increase retries for operations that are safe to repeat. Blob uploads are generally designed for retryability, but custom multi-request workflows should use conditions, leases, deterministic block IDs, or another idempotency mechanism.

For rules beyond `RetryOptions`, `BlobClientOptions.RetryPolicy` accepts a custom `HttpPipelinePolicy` and replaces the built-in retry policy. Prefer `RetryOptions` unless you genuinely need to redefine failure classification; implementing a pipeline retry policy also makes you responsible for preserving details such as cancellation, response disposal, retry accounting, and `Retry-After`.

## Circuit breaker with Polly v8

A circuit breaker should count only transient failures—not authentication, authorization, caller cancellation, or semantic errors such as `404` and `409`.

```powershell
dotnet add package Polly.Core
```

```csharp
using Polly;
using Polly.CircuitBreaker;

var circuitState = new CircuitBreakerStateProvider();

ResiliencePipeline storageCircuit = new ResiliencePipelineBuilder()
    .AddCircuitBreaker(new CircuitBreakerStrategyOptions
    {
        FailureRatio = 0.5,
        SamplingDuration = TimeSpan.FromSeconds(30),
        MinimumThroughput = 8,
        BreakDuration = TimeSpan.FromSeconds(30),
        StateProvider = circuitState,

        ShouldHandle = new PredicateBuilder()
            .Handle<RequestFailedException>(
                ex => IsTransientStatus(ex.Status))
            .Handle<HttpRequestException>()
            .Handle<IOException>()
    })
    .Build();

try
{
    await storageCircuit.ExecuteAsync(
        async cancellationToken =>
        {
            await UploadWithDeadlineAsync(
                highValueBlob,
                highValueStream,
                TimeSpan.FromMinutes(5),
                cancellationToken);
        },
        applicationStoppingToken);
}
catch (BrokenCircuitException ex)
{
    // Fail fast, enqueue for later processing, or return a temporary-unavailable result.
    Console.Error.WriteLine(
        $"Blob Storage circuit is open. Retry after: {ex.RetryAfter}");
    throw;
}
```

The breaker observes one **logical SDK operation**. It opens only after the SDK has exhausted that operation’s internal retries. Share the pipeline across calls; creating one breaker per request prevents it from accumulating failure history.

## Default retry classification

By default, Azure.Core considers transport exceptions and these HTTP responses retryable:

| Status | Meaning |
|---:|---|
| `408` | Request Timeout |
| `429` | Too Many Requests |
| `500` | Internal Server Error |
| `502` | Bad Gateway |
| `503` | Service Unavailable |
| `504` | Gateway Timeout |

Notable non-transient responses include:

- `400`: invalid request
- `401`: invalid or expired authentication
- `403`: insufficient authorization, firewall, or storage access
- `404`: resource not found
- `409`: state/conflict error
- `412`: failed ETag or other access condition

Those should normally be corrected rather than retried. One special case exists for geo-redundant reads: when `GeoRedundantSecondaryUri` is configured, a `404` from the secondary causes later retries to avoid that secondary because replication may not have reached it yet.

References: [Blob Storage retry policy](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), [`RetryOptions`](https://learn.microsoft.com/dotnet/api/azure.core.retryoptions), and [`ClientOptions.RetryPolicy`](https://learn.microsoft.com/dotnet/api/azure.core.clientoptions.retrypolicy).

## Reviewer Notes (LLM-as-Judge)

### Summary

8/8 criteria passed. The response comprehensively covers SDK retry configuration, timeout boundaries, error classification, operation-specific policies, geo-redundant behavior, and circuit breaking.

### Strengths

- Correctly distinguishes per-attempt NetworkTimeout from an end-to-end CancellationToken deadline.
- Clearly separates transient transport/service failures from authentication and authorization failures.
- Avoids multiplicative retry loops and scopes the Polly circuit breaker around logical SDK operations.
- Provides correct default retryable status codes and appropriate non-retryable examples.

## Grader Results

- retry-configuration.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - `BlobClientOptions.Retry` configuration with `RetryOptions`: Pass
      - `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed): Pass
      - `NetworkTimeout` for per-request timeouts: Pass
      - Default retryable status codes (408, 429, 500, 502, 503, 504): Pass
      - Non-retryable errors (400, 401, 403, 404, 409): Pass
      - Per-operation `CancellationToken` for timeout control: Pass
      - Geo-redundant retry with `GeoRedundantSecondaryUri`: Pass
      - Interaction with Polly or other resilience libraries: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

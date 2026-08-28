# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 407.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 4/7 |
| Duration | 407.2s |
| Timestamp | 2026-08-27T15:39:54Z |
| Files Generated | 0 |
| Event Count | 3235 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 383.8s |
| Review | 22.0s |
| **Total** | **407.2s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-diagnostics |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=322127 out=3784 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 3710ms

**Input:**

```json
{"skill":"azure-diagnostics"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Diagnostics

> **AUTHORITATIVE GUIDANCE — MANDATORY COMPLIANCE**
>
> This document is the **official source** for debugging and troubleshooting Azure production issues. Follow these instructions to diagnose and resolve common Azure service problems systematically.

## Triggers

Activate this skill when user wants to:
- Debug or troubleshoot production issues
- Diagnose errors in Azure services
- Analyze application logs or metrics
- Fix image pull, cold start, or health probe issues
- Investigate why Azure resources are failing
- Find root cause of application errors
- Troubleshoot App Service issues (high CPU, deployment failures, crashes, slow responses, TLS/custom domains)
- Respond to prompts like "troubleshoot app service", "app service high CPU", or "app service deployment failure"
- Troubleshoot Azure Function Apps (invocation failures, timeouts, binding errors)
- Find the App Insights or Log Analytics workspace linked to a Function App
- Troubleshoot AKS clusters, nodes, pods, ingress, or Kubernetes networking issues
- Troubleshoot Azure VM connectivity issues (RDP/SSH failures, port 3389/22 timeouts, NSG or firewall blocking, credential resets)
- Troubleshoot Azure Messaging SDK issues (Event Hubs, Service Bus connection failures, AMQP errors, message lock issues)

## Rules

1. Start with systematic diagnosis flow
2. Use AppLens (MCP) for AI-powered diagnostics when available
3. Check resource health before deep-diving into logs
4. Select appropriate troubleshooting guide based on service type
5. Document findings and attempted remediation steps
6. Route AKS incidents to the dedicated AKS troubleshooting document

---

## Quick Diagnosis Flow

1. **Identify symptoms** - What's failing?
2. **Check resource health** - Is Azure healthy?
3. **Review logs** - What do logs show?
4. **Analyze metrics** - Performance patterns?
5. **Investigate recent changes** - What changed?

---

## Troubleshooting Guides by Service

| Ser
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 184961ms

**Input:**

```json
{"intent":"Provide current .NET Azure SDK best practices for handling Azure Key Vault Secrets failures, RequestFailedException, optimistic concurrency, and throttling retries.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 184965ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Security.KeyVault.Secrets .NET RequestFailedException handling, SecretClientOptions retry settings, 403 Forbidden, 404 SecretNotFound, 409 version or update conflicts, and 429 throttling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. web_search (via github-mcp-server) — 67870ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure RequestFailedException Status ErrorCode GetRawResponse Headers Retry-After"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s a summary and reference for the main Azure Key Vault REST API secret error codes you mentioned (403, 404, 409, 429), their meanings, and related error strings, sourced directly from Microsoft Learn:\n\n---\n\n### Azure Key Vault REST Secret Error Codes\n\n#### 403 Forbidden (`Forbidden`, `AccessDenied`, `ForbiddenByFirewall`)\n- **Meaning:** Caller is authenticated but doesn’t have enough permissions to perform the operation.\n    - You might lack the appropriate access policy or RBAC assignment for the requested operation.\n    - Firewall rules may also block the client.\n- **Error strings:** `Forbidden`, `AccessDenied`, `ForbiddenByFirewall`\n- **Resolution:** Assign the correct permissions via Access Policies or Azure RBAC, and check that the firewall allows the call source. Make sure your token is valid and has the correct permissions【3:1†source】【3:5†source】【3:6†source】【3:9†source】.\n\n#### 404 Not Found (`SecretNotFound`, `ResourceNotFound`)\n- **Meaning:** The requested secret does not exist. This could be due to a typo in the secret name or the secret being deleted.\n- **Error strings:** `SecretNotFound`, `ResourceNotFound`\n- **Resolution:** Verify the secret name and ensure the secret has been created or is not deleted (including soft delete)【3:1†source】.\n\n#### 409 Conflict (`Conflict`, `ConflictError`)\n- **Meaning:** The request could not be completed due to a conflict with the current state of the resource. This often happens when multiple requests are made to create/update/delete the same secret simultaneously.\n- **Error string:** `Conflict`, `ConflictError`\n- **Resolution:** Retry the operation after some time or sequence operations to avoid concurrent changes【3:1†source】.\n\n#### 429 Too Many Requests (`Throttling`)\n- **Meaning:** Too many requests have been sent in a short period. Azure Key Vault throttles requests to protect the service.\n- **Error string:** `
... (truncated)
```

### ✅ 5. web_search (via github-mcp-server) — 66204ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure.Security.KeyVault.Secrets SecretClientOptions Retry RetryOptions .NET"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 6. web_search (via github-mcp-server) — 64945ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Key Vault REST secret errors 403 404 409 429 error codes SecretNotFound Forbidden Conflict Throttling"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 7. web_fetch — 7826ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.requestfailedexception?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/key-vault/general/rest-error-codes:
---
layout: Conceptual
title: REST API error codes - Azure Key Vault | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/key-vault/general/rest-error-codes
breadcrumb_path: /azure/breadcrumb/azure-security/toc.json
feedback_help_link_url: https://learn.microsoft.com/answers/tags/5/azure-key-vault/
feedback_help_link_type: get-help-at-qna
feedback_product_url: https://feedback.azure.com/d365community/forum/285c5ae0-f524-ec11-b6e6-000d3a4f0da0
feedback_system: Standard
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
uhfHeaderId: azure
ms.suite: office
author: msmbaldwin
ms.author: mbaldwin
learn_banner_products:
- azure
description: An operation on an Azure Key Vault web service may return the following error codes.
services: key-vault
ms.service: azure-key-vault
ms.subservice: general
ms.topic: reference
ms.date: 2026-04-10T00:00:00.0000000Z
locale: en-us
document_id: 8a71ce24-5593-c767-e4fb-83dfbcd1b6ce
document_version_independent_id: 7284734d-db6b-93c3-b9d1-9de6533bf680
updated_at: 2026-06-12T22:35:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-security-docs-pr/blob/live/articles/key-vault/general/rest-error-codes.md
gitcommit: https://github.com/MicrosoftDocs/azure-security-docs-pr/blob/cc9ac0f545a688e4f2016a1541f565f136e23129/articles/key-vault/general/rest-error-codes.md
git_commit_id: cc9ac0f545a688e4f2016a1541f565f136e23129
site_name: Docs
depot_name: Learn.azure-security
page_type: conceptual
toc_rel: toc.json
word_count: 1110
asset_id: key-vault/general/rest-error-codes
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/key-vault/general/rest-error-codes.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/f488294d-f483-456e-94e3-755f933b811b
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/02662057-0b9b-40f4-a3c7
... (truncated)
```

### ✅ 8. web_fetch — 7475ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclientoptions?view=azure-dotnet"}
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

### ✅ 9. web_fetch — 7124ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions?view=azure-dotnet"}
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

### ✅ 10. web_fetch — 6772ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclient.updatesecretproperties?view=azure-dotnet"}
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

### ✅ 11. web_fetch — 6440ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/azure/key-vault/general/rest-error-codes"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclient.updatesecretproperties?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: SecretClient.UpdateSecretProperties Method (Azure.Security.KeyVault.Secrets) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.security.keyvault.secrets.secretclient.updatesecretproperties?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Security.KeyVault.Secrets.SecretClient.UpdateSecretProperties*
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
- Azure.Security.KeyVault.Secrets.SecretClient.UpdateSecretProperties
api_location:
- Azure.Security.KeyVault.Secrets.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 6dba525b-4f3c-0ba4-6625-2f8ff2d5ffb3
document_version_independent_id: 45bfd510-43f4-eceb-927f-5cdb8a7e071a
updated_at: 2026-05-06T11:36:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Security.KeyVault.Secrets/SecretClient.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/f3535d7b864f1a7a9a07185098c8ba5bc53a31d7/xml/Azure.Security.KeyVault.Secrets/SecretClient.xml
git_commit_id: f3535d7b864f1a7a9a07185098c8ba5bc53a31d7
default_moniker: azure-dotnet
site_name: Do
... (truncated)
```

### ✅ 12. web_fetch — 3246ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/keyvault/Azure.Security.KeyVault.Secrets/src/SecretClient.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/keyvault/Azure.Security.KeyVault.Secrets/src/SecretClient.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.ClientModel.Primitives;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Pipeline;
using Azure.Security.KeyVault.Secrets.Models;

namespace Azure.Security.KeyVault.Secrets
{
    /// <summary>
    /// The SecretClient provides synchronous and asynchronous methods to manage <see cref="KeyVaultSecret"/> in the Azure Key Vault. The client
    /// supports creating, retrieving, updating, deleting, purging, backing up, restoring, and listing <see cref="KeyVaultSecret"/>.
    /// The client also supports listing <see cref="DeletedSecret"/> for a soft-delete enabled Azure Key Vault.
    /// </summary>
    /// <remarks>
    /// Internally, all transport — request building, response parsing, paging, model
    /// (de)serialization, and LRO polling — is delegated to the TypeSpec-generated
    /// <c>KeyVaultSecretsClient</c> (internal). Public method signatures, return
    /// types, exception contracts and recorded HTTP traffic match every previously
    /// shipped version of this package, so adopting this build is a no-op for
    /// existing consumers. The legacy hand-written transport (KeyVaultPipeline,
    /// SecretBackup, JSON read/write methods on the model classes) is no longer
    /// invoked.
    /// </remarks>
    [CallerShouldAudit("https://aka.ms/azsdk/callershouldaudit/security-keyvault-secrets")]
    public class SecretClient
    {
        private const string OTelSecretNameKey    = "az.keyvault.secret.name";
        private const string OTelSecretVersionKey = "az.keyvault.secret.version";

   
... (truncated)
```

### ✅ 13. web_fetch — 2916ms

**Input:**

```json
{"max_length":15000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/keyvault/Azure.Security.KeyVault.Secrets/src/SecretProperties.cs"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/key-vault/general/overview-throttling:
---
layout: Conceptual
title: Azure Key Vault throttling guidance | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/key-vault/general/overview-throttling
breadcrumb_path: /azure/breadcrumb/azure-security/toc.json
feedback_help_link_url: https://learn.microsoft.com/answers/tags/5/azure-key-vault/
feedback_help_link_type: get-help-at-qna
feedback_product_url: https://feedback.azure.com/d365community/forum/285c5ae0-f524-ec11-b6e6-000d3a4f0da0
feedback_system: Standard
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
uhfHeaderId: azure
ms.suite: office
author: msmbaldwin
ms.author: mbaldwin
learn_banner_products:
- azure
description: Key Vault throttling limits the number of concurrent calls to prevent overuse of resources.
services: key-vault
ms.service: azure-key-vault
ms.subservice: general
ms.topic: best-practice
ms.date: 2026-07-15T00:00:00.0000000Z
locale: en-us
document_id: fbc6bfba-ef99-bfbd-23e6-987d570b81d4
document_version_independent_id: c5f7c3c6-ee98-2c29-3a5f-9a82b5bd57c8
updated_at: 2026-07-17T17:33:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-security-docs-pr/blob/live/articles/key-vault/general/overview-throttling.md
gitcommit: https://github.com/MicrosoftDocs/azure-security-docs-pr/blob/bdb26e62f581710e3c6a59994e8b564e60c1a861/articles/key-vault/general/overview-throttling.md
git_commit_id: bdb26e62f581710e3c6a59994e8b564e60c1a861
site_name: Docs
depot_name: Learn.azure-security
page_type: conceptual
toc_rel: toc.json
word_count: 949
asset_id: key-vault/general/overview-throttling
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/key-vault/general/overview-throttling.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/f488294d-f483-456e-94e3-755f933b811b
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devre
... (truncated)
```

### ✅ 14. web_fetch — 2570ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/azure/key-vault/general/overview-throttling"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/keyvault/Azure.Security.KeyVault.Secrets/src/SecretProperties.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Threading;
using Azure.Core;

namespace Azure.Security.KeyVault.Secrets
{
    /// <summary>
    /// <see cref="SecretProperties"/> is the resource containing all the properties of the secret except its value.
    /// </summary>
    public class SecretProperties : IJsonDeserializable, IJsonSerializable
    {
        private const string IdPropertyName = "id";
        private const string ContentTypePropertyName = "contentType";
        private const string KidPropertyName = "kid";
        private const string ManagedPropertyName = "managed";
        private const string AttributesPropertyName = "attributes";
        private const string TagsPropertyName = "tags";
        private const string PreviousVersionPropertyName = "previousVersion";

        private static readonly JsonEncodedText s_contentTypePropertyNameBytes = JsonEncodedText.Encode(ContentTypePropertyName);
        private static readonly JsonEncodedText s_attributesPropertyNameBytes = JsonEncodedText.Encode(AttributesPropertyName);
        private static readonly JsonEncodedText s_tagsPropertyNameBytes = JsonEncodedText.Encode(TagsPropertyName);

        private SecretAttributes _attributes;
        private Dictionary<string, string> _tags;
        private string _keyId;

        internal SecretProperties()
        {
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SecretProperties"/> class.
        /// </summary>
        /// <param name="name">The name of the secret.</param>
        /// <exception cref="ArgumentException"><paramref name="name"/> i
... (truncated)
```

### ✅ 15. web_fetch — 1475ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.responseheaders?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.responseheaders?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: ResponseHeaders Struct (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.responseheaders?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.ResponseHeaders
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
- Azure.Core.ResponseHeaders
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: d8d2d294-93ec-4031-2ae6-41a712d76f36
document_version_independent_id: 14f32499-fdd8-19fe-e8ca-26072ad88057
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/ResponseHeaders.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core/ResponseHeaders.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: struct
ms.assetid: Azure.Core.ResponseHeaders
description: 'Headers received as part of the Response. '
toc_rel: azure_ref_toc/_splitted/azure.core/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp vb
... (truncated)
```

## Copilot Response

I’m grounding the patterns in the current Azure SDK guidance, especially retry behavior and Key Vault error-code semantics.

The Azure documentation endpoint timed out, so I’m using the official API references directly rather than relying on memory for exception properties and retry defaults.

`Azure.Security.KeyVault.Secrets` reports service failures through `Azure.RequestFailedException`. Use `Status` for reliable control flow and `ErrorCode` for diagnostics—the latter can be `null` or vary by operation.

```csharp
using Azure;
using Azure.Core;
using Azure.Core.Pipeline;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;
using Microsoft.Extensions.Logging;

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
    new Uri("https://<vault-name>.vault.azure.net"),
    new DefaultAzureCredential(),
    options);
```

The SDK retry pipeline handles transient failures, including HTTP 429, before throwing. It also honors the service's `Retry-After` header when present.

### Catching common failures

```csharp
public static async Task<KeyVaultSecret?> GetSecretOrNullAsync(
    SecretClient client,
    string secretName,
    ILogger logger,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.GetSecretAsync(
                secretName,
                version: null,
                cancellationToken: cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        LogFailure(logger, ex, secretName);

        // Usually missing Key Vault Secrets User/Officer RBAC permissions,
        // an incorrect access policy, or a firewall/private endpoint restriction.
        throw;
    }
    catch (RequestFailedException ex) when (
        ex.Status == 404 ||
        string.Equals(
            ex.ErrorCode,
            "SecretNotFound",
            StringComparison.OrdinalIgnoreCase))
    {
        // Treat absence as expected only when that matches application semantics.
        logger.LogInformation(
            "Key Vault secret {SecretName} was not found. Code={ErrorCode}",
            secretName,
            ex.ErrorCode);

        return null;
    }
    catch (RequestFailedException ex) when (
        ex.Status is 409 or 412)
    {
        LogFailure(logger, ex, secretName);

        // 409: the vault or secret is in a conflicting state, such as an
        // overlapping delete/recover/create operation.
        //
        // 412: a conditional request no longer matches the current resource.
        //
        // Re-read the current state before deciding whether retrying is safe.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        LogFailure(logger, ex, secretName);

        // The configured SDK retries have already been exhausted.
        // Surface the failure, shed load, or retry the larger operation later.
        throw;
    }
    catch (RequestFailedException ex)
    {
        LogFailure(logger, ex, secretName);
        throw;
    }
}
```

### Extracting status, error code, request ID, and retry headers

```csharp
private static void LogFailure(
    ILogger logger,
    RequestFailedException exception,
    string secretName)
{
    Response? response = exception.GetRawResponse();

    string? retryAfter = GetHeader(response, "Retry-After");
    string? retryAfterMilliseconds = GetHeader(response, "retry-after-ms");
    string? requestId = response?.Headers.RequestId;

    logger.LogError(
        exception,
        """
        Key Vault request failed for secret {SecretName}.
        HttpStatus={HttpStatus}
        ErrorCode={ErrorCode}
        RequestId={RequestId}
        RetryAfter={RetryAfter}
        RetryAfterMilliseconds={RetryAfterMilliseconds}
        """,
        secretName,
        exception.Status,
        exception.ErrorCode,
        requestId,
        retryAfter,
        retryAfterMilliseconds);
}

private static string? GetHeader(Response? response, string name)
{
    return response is not null &&
           response.Headers.TryGetValue(name, out string value)
        ? value
        : null;
}
```

The key exception members are:

```csharp
catch (RequestFailedException ex)
{
    int httpStatus = ex.Status;       // 0 if no HTTP response was received
    string? errorCode = ex.ErrorCode; // Service-specific; may be null
    string message = ex.Message;
    Response? rawResponse = ex.GetRawResponse();
}
```

### Handling write/version conflicts

Secret values are immutable per version. `SetSecretAsync` creates a new version; it does not overwrite the value of an existing version. `UpdateSecretPropertiesAsync` updates metadata for one specific version.

```csharp
public static async Task<KeyVaultSecret> CreateSecretVersionAsync(
    SecretClient client,
    string name,
    string value,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.SetSecretAsync(
                new KeyVaultSecret(name, value),
                cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status is 409 or 412)
    {
        // Do not blindly retry a version-creating write: if the original response
        // was lost, retrying may create another version. Re-read the latest
        // version and reconcile it with the intended operation.
        throw;
    }
}
```

For property updates, fetch the exact version first:

```csharp
Response<KeyVaultSecret> current =
    await client.GetSecretAsync(name, version, cancellationToken: cancellationToken);

current.Value.Properties.ExpiresOn = DateTimeOffset.UtcNow.AddDays(90);

try
{
    await client.UpdateSecretPropertiesAsync(
        current.Value.Properties,
        cancellationToken);
}
catch (RequestFailedException ex) when (ex.Status is 409 or 412)
{
    // Re-fetch the version and decide whether the update should be reapplied.
    throw;
}
```

Use status codes as the primary classification:

| Status | Typical handling |
|---|---|
| `403` | Do not retry; correct RBAC, access policy, firewall, or private networking configuration |
| `404` | Return “not found” only if absence is expected; otherwise propagate |
| `409` / `412` | Re-read and reconcile state; retry only when the operation is safe |
| `429` | Let the SDK honor `Retry-After` with exponential backoff; cache reads and reduce request volume if retries are exhausted |

Avoid logging secret values, credentials, or authorization headers. Logging the secret name may also be inappropriate if names contain sensitive business information.

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 7 criteria passed. The response provides strong RequestFailedException, diagnostics, conflict, retry, and throttling examples, but omits deleted-secret and purge-protection handling and does not fully explain RBAC-versus-access-policy diagnosis.

### Strengths

- Uses RequestFailedException consistently and preserves unexpected failures by rethrowing.
- Clearly demonstrates Status, ErrorCode, raw response, request ID, and retry-header extraction.
- Provides practical SecretClientOptions exponential retry settings and sound 429 guidance.
- Explains secret version creation semantics and cautions against blindly retrying version-creating writes.

### Issues

- No deleted-secret lookup or recovery pattern to distinguish soft-deleted secrets from ordinary 404 responses.
- No coverage of soft-delete name conflicts, purge operations, or purge-protection failures.
- The 403 guidance lists possible authorization causes without showing how to identify whether the vault uses Azure RBAC or legacy access policies.

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
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

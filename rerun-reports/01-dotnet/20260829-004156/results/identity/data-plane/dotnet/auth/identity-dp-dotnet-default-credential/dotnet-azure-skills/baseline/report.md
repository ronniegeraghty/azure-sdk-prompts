# Evaluation Report: identity-dp-dotnet-default-credential

**Config:** dotnet-azure-skills/baseline | **Result:** ✅ PASSED | **Duration:** 132.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-default-credential` |
| Config | dotnet-azure-skills/baseline |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 132.1s |
| Timestamp | 2026-08-28T17:28:09Z |
| Files Generated | 0 |
| Event Count | 3963 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 107.2s |
| Review | 24.3s |
| **Total** | **132.1s** |

## Configuration

- **name:** dotnet-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=109591 out=2810 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. web_search (via github-mcp-server) — 51808ms

**Input:**

```json
{"query":"Official Microsoft Azure Identity .NET DefaultAzureCredential credential chain order 2026 Azure.Identity BlobServiceClient C# troubleshooting"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To create a runnable C# example using `BlobServiceClient` with `DefaultAzureCredential`, you'll need the following required NuGet packages:\n\n- `Azure.Storage.Blobs`\n- `Azure.Identity`\n\n## Step 1: Install Required NuGet Packages\n\nRun these commands in your project directory:\n\n```shell\ndotnet add package Azure.Storage.Blobs\ndotnet add package Azure.Identity\n```\n\n## Step 2: Add Using Directives\n\n```csharp\nusing Azure.Identity;\nusing Azure.Storage.Blobs;\n```\n\n## Step 3: Example – Authenticate and List Containers\n\nHere is a runnable C# example that connects to Azure Blob Storage using `DefaultAzureCredential`:\n\n```csharp\nusing System;\nusing System.Threading.Tasks;\nusing Azure.Identity;\nusing Azure.Storage.Blobs;\n\nclass Program\n{\n    static async Task Main(string[] args)\n    {\n        // Replace with your Blob service URL\n        string blobServiceUri = \"https://<your-storage-account>.blob.core.windows.net/\";\n\n        // Authenticate with DefaultAzureCredential\n        var client = new BlobServiceClient(new Uri(blobServiceUri), new DefaultAzureCredential());\n\n        await foreach (var container in client.GetBlobContainersAsync())\n        {\n            Console.WriteLine($\"Container name: {container.Name}\");\n        }\n    }\n}\n```\n\n### Notes\n\n- `<your-storage-account>` should be replaced with your actual Azure Storage Account name.\n- For `DefaultAzureCredential` to work locally, ensure you are signed in using `az login` or have the appropriate environment variables set. In Azure, it will use the managed identity of the resource by default.\n\n### Official Docs\n\n- [Get started with Azure Blob Storage and .NET](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-dotnet)\n- [DefaultAzureCredential Class – Azure.Identity](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential)【3:1†source】【3:9†source】【3:8†sour
... (truncated)
```

### ✅ 2. web_search (via github-mcp-server) — 50785ms

**Input:**

```json
{"query":"Official Microsoft docs C# BlobServiceClient DefaultAzureCredential required NuGet packages Azure.Identity Azure.Storage.Blobs runnable example"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The official Microsoft Azure Identity .NET `DefaultAzureCredential` credential chain (as of 2026) attempts several authentication mechanisms in a specific order, stopping at the first one that succeeds. This makes it the recommended option for most Azure SDK client scenarios, such as constructing a `BlobServiceClient` with C#.\n\n**Current `DefaultAzureCredential` chain order (2026):**\n\n1. **EnvironmentCredential** – Uses environment variables like `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_SECRET` for service principal or user authentication.\n2. **WorkloadIdentityCredential** – For Kubernetes or other workload identity federation.\n3. **ManagedIdentityCredential** – Uses Azure Managed Identities (system- or user-assigned, e.g., when running in Azure App Service, VM, etc.).\n4. **VisualStudioCredential** – Uses an account signed in to Visual Studio (for local dev).\n5. **VisualStudioCodeCredential** – Uses an account signed in to VS Code (if SSO and broker are enabled).\n6. **AzureCliCredential** – Uses credentials from `az login` (Azure CLI).\n7. **AzurePowerShellCredential** – Uses credentials from `Az` PowerShell module.\n8. **AzureDeveloperCliCredential** – Uses credentials from the Azure Developer CLI (`azd login`).\n9. **InteractiveBrowserCredential** – Pops up a browser window for sign-in (not enabled by default).\n10. **BrokerCredential** – SSO using the Azure Identity Broker, only if specifically configured.\n\nYou can control which credentials are enabled/disabled via `DefaultAzureCredentialOptions`.\n\n### Example usage (C#)\n\n```csharp\nusing Azure.Identity;\nusing Azure.Storage.Blobs;\n\nvar credential = new DefaultAzureCredential();\nvar client = new BlobServiceClient(new Uri(\"https://<account>.blob.core.windows.net/\"), credential);\n```\n\n### Troubleshooting Tips\n\n- **CredentialUnavailableException**: Indicates no credential in the chain could provide a token. Check that 
... (truncated)
```

### ❌ 3. web_fetch — 7533ms

**Input:**

```json
{"max_length":16000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting - status code 404
```

### ✅ 4. web_fetch — 7533ms

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/on-premises-apps"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains:
---
layout: Conceptual
title: Credential chains in the Azure Identity library for .NET - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: concept-article
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: This article describes the DefaultAzureCredential and ChainedTokenCredential classes in the Azure Identity library.
ms.date: 2025-08-13T00:00:00.0000000Z
locale: en-us
document_id: a5e37cff-dad3-2449-9853-9c4a12abd30d
document_version_independent_id: 42104a50-1199-6f33-12f2-e4edefa13934
updated_at: 2026-03-30T21:10:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/authentication/credential-chains.md
gitcommit: https://github.com/dotnet/docs/blob/156931bb4ec1e81b028c76ea983553f2e9778bdd/docs/azure/sdk/authentication/credential-chains.md
git_commit_id: 156931bb4ec1e81b028c76ea983553f2e9778bdd
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1586
asset_id: azure/sdk/authentication/credential-chains
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/authentication/credential-chains.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/90370425-aca4-4a39-9533-d52e5e002a5d
platformId: a77e8a0e-e3e5-f851-27fc-13be2980c0c2
---

... (truncated)
```

### ✅ 5. web_fetch — 7216ms

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/on-premises-apps:
---
layout: Conceptual
title: Authenticate to Azure resources from .NET apps hosted on-premises - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/on-premises-apps
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: how-to
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: This article describes how to authenticate your application to Azure services when using the Azure SDK for .NET in on-premises hosted apps.
ms.date: 2026-03-20T00:00:00.0000000Z
ms.custom:
- devx-track-dotnet
- engagement-fy23
- sfi-image-nochange
locale: en-us
document_id: 5fff20ea-2ad6-faab-47b2-a80474f9d1d9
document_version_independent_id: 6a6aa86b-4f67-51df-4b16-4fdf439832d3
updated_at: 2026-03-23T20:42:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/authentication/on-premises-apps.md
gitcommit: https://github.com/dotnet/docs/blob/870c28a07336caa46b0e4c36071881831a9b077e/docs/azure/sdk/authentication/on-premises-apps.md
git_commit_id: 870c28a07336caa46b0e4c36071881831a9b077e
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1544
asset_id: azure/sdk/authentication/on-premises-apps
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/authentication/on-premises-apps.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
- https://authoring-docs-microsoft.poolparty.biz/devrel/7696cda6-0510-
... (truncated)
```

### ✅ 6. web_fetch — 6893ms

**Input:**

```json
{"max_length":13000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access:
---
layout: Conceptual
title: Assign an Azure Role for Blob Data Access - Azure Storage | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access
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
ms.reviewer: dineshm
description: Assign Azure roles for blob data access with Azure RBAC and Microsoft Entra ID. Follow this guide to grant least-privilege permissions securely.
author: stevenmatthew
ms.author: shaas
ms.service: azure-blob-storage
ms.topic: how-to
ms.date: 2026-05-08T00:00:00.0000000Z
ms.devlang: powershell
ms.custom: devx-track-azurepowershell, devx-track-azurecli
locale: en-us
document_id: 1ca94e03-5f05-6703-1672-8062f31c0780
document_version_independent_id: a6a8d9e9-9aac-c0ed-c474-34cbd0609584
updated_at: 2026-05-09T06:18:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-pr/blob/live/articles/storage/blobs/assign-azure-role-data-access.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-pr/blob/99ea6b4bf72db19a72caba74818eead4a2043db7/articles/storage/blobs/assign-azure-role-data-access.md
git_commit_id: 99ea6b4bf72db19a72caba74818eead4a2043db7
site_name: Docs
depot_name: Azure.azure-documents
page_type: conceptual
interactive_type: azurecli
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/Azure.azure-documents/{branchName}{pdfName}
word_count: 1262
asset_id: storage/blobs/assign-azure-rol
... (truncated)
```

### ✅ 7. web_fetch — 6562ms

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity:
---
layout: Conceptual
title: Authenticate Azure-hosted .NET apps to Azure resources using a system-assigned managed identity - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: how-to
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn how to authenticate Azure-hosted .NET apps to other Azure services using a system-assigned managed identity.
ms.custom: devx-track-dotnet, engagement-fy23, devx-track-azurecli
ms.date: 2025-11-10T00:00:00.0000000Z
locale: en-us
document_id: aafadbb3-d54c-03c6-864e-bdeb330ce49f
document_version_independent_id: 11530b71-136a-5f66-c1d7-0cd8f196e569
updated_at: 2026-03-20T17:26:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/authentication/system-assigned-managed-identity.md
gitcommit: https://github.com/dotnet/docs/blob/863ac42cec71c43cff89653ed8951574358361b1/docs/azure/sdk/authentication/system-assigned-managed-identity.md
git_commit_id: 863ac42cec71c43cff89653ed8951574358361b1
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1449
asset_id: azure/sdk/authentication/system-assigned-managed-identity
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/authentication/system-assigned-managed-identity.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-
... (truncated)
```

### ✅ 8. web_fetch — 3180ms

**Input:**

```json
{"max_length":7000,"start_index":15500,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/logging:
---
layout: Conceptual
title: Logging with the Azure SDK for .NET - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/logging
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: how-to
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn how to enable logging with the Azure SDK for .NET client libraries
ms.date: 2025-04-05T00:00:00.0000000Z
ms.custom:
- devx-track-dotnet
- engagement-fy23
- sfi-ropc-nochange
locale: en-us
document_id: 337e8ed0-9a24-93ce-c7cf-478c0e24dc06
document_version_independent_id: 2151de27-a308-a249-277d-7addb659775e
updated_at: 2026-03-30T21:10:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/logging.md
gitcommit: https://github.com/dotnet/docs/blob/156931bb4ec1e81b028c76ea983553f2e9778bdd/docs/azure/sdk/logging.md
git_commit_id: 156931bb4ec1e81b028c76ea983553f2e9778bdd
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1528
asset_id: azure/sdk/logging
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/logging.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/7696cda6-0510-47f6-8302-71bb5d2e28cf
- https://authoring-docs-microsoft.poolparty.biz/devrel/fd7d5d12-dbbc-4585-98a0-c6a0a5324f97
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/69c76c32-967e-4c65-b89a-74cc527db725
- https://authoring-docs
... (truncated)
```

### ❌ 9. web_fetch — 2845ms

**Input:**

```json
{"max_length":14000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshoot"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshoot - status code 404
```

### ✅ 10. web_fetch — 2842ms

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/logging"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: DefaultAzureCredential Class (Azure.Identity) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Identity.DefaultAzureCredential
namespace: Azure.Identity
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/b1cfdec6-b0c3-4209-818c-736879856e0e
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Identity.DefaultAzureCredential
api_location:
- Azure.Identity.dll
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 414d8832-9224-3bb5-2bf4-1197c41fdedf
document_version_independent_id: 02d95dd6-bf1b-5cef-bc94-ac187b133c45
updated_at: 2026-08-28T11:57:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Identity/DefaultAzureCredential.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/c5d2f1e44889c00f914e0d66fa7d0148914ebc8a/xml/Azure.Identity/DefaultAzureCredential.xml
git_commit_id: c5d2f1e44889c00f914e0d66fa7d0148914ebc8a
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Identity.DefaultAzureCredential
description: "Simplifies authentication while developing apps tha
... (truncated)
```

### ✅ 11. web_fetch — 2505ms

**Input:**

```json
{"max_length":6000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains:
dential` implementation, such as `ManagedIdentityCredential`. See the [**Derived** list](/en-us/dotnet/api/azure.core.tokencredential?view=azure-dotnet&amp;preserve-view=true#definition) for options.

Here's why:

- **Debugging challenges**: When authentication fails, it can be challenging to debug and identify the offending credential. You must enable logging to see the progression from one credential to the next and the success/failure status of each. For more information, see Debug a chained credential.
- **Performance overhead**: The process of sequentially trying multiple credentials can introduce performance overhead. For example, when running on a local development machine, managed identity is unavailable. Consequently, `ManagedIdentityCredential` always fails in the local development environment, unless explicitly disabled via its corresponding `Exclude`-prefixed property.
- **Unpredictable behavior**: `DefaultAzureCredential` checks for the presence of certain [environment variables](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/identity/Azure.Identity/README.md#environment-variables). It's possible that someone could add or modify these environment variables at the system level on the host machine. Those changes apply globally and therefore alter the behavior of `DefaultAzureCredential` at runtime in any app running on that machine. For more information on unpredictability, see [Use deterministic credentials in production environments](best-practices#use-deterministic-credentials-in-production-environments).

## Debug a chained credential

To diagnose an unexpected issue or to understand what a chained credential is doing, [enable logging](../logging) in your app. Optionally, filter the logs to only those events emitted from the Azure Identity library. For example:

```csharp
using AzureEventSourceListener listener = new((args, message) 
... (truncated)
```

## Copilot Response

`DefaultAzureCredential` lets the same client code use developer credentials locally and workload/managed identity credentials when deployed. It obtains Microsoft Entra tokens; Azure RBAC still determines what the authenticated identity can do.

## 1. Required NuGet packages

```powershell
dotnet new console --name BlobCredentialExample
Set-Location BlobCredentialExample

dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

- [`Azure.Identity`](https://www.nuget.org/packages/Azure.Identity): provides `DefaultAzureCredential`.
- [`Azure.Storage.Blobs`](https://www.nuget.org/packages/Azure.Storage.Blobs): provides `BlobServiceClient`.
- Optional: [`Azure.Identity.Broker`](https://www.nuget.org/packages/Azure.Identity.Broker) for VS Code authentication and OS account broker support.

## 2. Complete runnable example

Replace `Program.cs` with:

```csharp
using Azure;
using Azure.Identity;
using Azure.Storage.Blobs;

string? blobServiceUri = Environment.GetEnvironmentVariable("BLOB_SERVICE_URI");

if (string.IsNullOrWhiteSpace(blobServiceUri))
{
    Console.Error.WriteLine(
        "Set BLOB_SERVICE_URI to a value such as " +
        "https://mystorageaccount.blob.core.windows.net.");
    return 1;
}

if (!Uri.TryCreate(blobServiceUri, UriKind.Absolute, out Uri? serviceUri))
{
    Console.Error.WriteLine("BLOB_SERVICE_URI is not a valid absolute URI.");
    return 1;
}

// Reuse this instance. Azure Identity caches and refreshes tokens internally.
var credential = new DefaultAzureCredential();

// The Blob SDK requests the appropriate Microsoft Entra token when an
// operation is performed; constructing the client does not make a request.
var blobServiceClient = new BlobServiceClient(serviceUri, credential);

try
{
    Console.WriteLine($"Containers in {blobServiceClient.AccountName}:");

    await foreach (var container in blobServiceClient.GetBlobContainersAsync())
    {
        Console.WriteLine($"- {container.Name}");
    }

    return 0;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine("Microsoft Entra authentication failed.");
    Console.Error.WriteLine(ex.Message);
    return 2;
}
catch (RequestFailedException ex) when (ex.Status is 401 or 403)
{
    Console.Error.WriteLine(
        $"Authentication or authorization failed ({ex.Status}, {ex.ErrorCode}).");
    Console.Error.WriteLine(
        "Confirm that the selected identity has a Blob data-plane RBAC role.");
    Console.Error.WriteLine(ex.Message);
    return 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure Storage request failed ({ex.Status}, {ex.ErrorCode}).");
    Console.Error.WriteLine(ex.Message);
    return 4;
}
```

Configure and run it:

```powershell
$env:BLOB_SERVICE_URI = "https://<storage-account>.blob.core.windows.net"
dotnet run
```

Locally, sign in through a supported development tool—for example, Visual Studio, Azure CLI (`az login`), Azure PowerShell (`Connect-AzAccount`), or Azure Developer CLI (`azd auth login`). The selected identity needs at least **Storage Blob Data Reader** to list containers, or **Storage Blob Data Contributor** if the application writes data. Management-plane roles such as **Reader** or **Contributor** do not automatically grant access to blob data.

The essential client construction is:

```csharp
var credential = new DefaultAzureCredential();

var client = new BlobServiceClient(
    new Uri("https://<storage-account>.blob.core.windows.net"),
    credential);
```

## 3. Credential chain order

With current `Azure.Identity` releases, credentials are attempted in this order until one returns a token:

| Order | Credential | Typical source |
|---:|---|---|
| 1 | `EnvironmentCredential` | `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and a secret, certificate, or federated configuration |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly Kubernetes |
| 3 | `ManagedIdentityCredential` | System-assigned or user-assigned Azure managed identity |
| 4 | `VisualStudioCredential` | Account signed into Visual Studio |
| 5 | `VisualStudioCodeCredential` | VS Code Azure account; requires broker support |
| 6 | `AzureCliCredential` | Account authenticated through `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated through `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated through `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser login; **disabled by default** |
| 10 | Broker credential | OS account broker; requires `Azure.Identity.Broker` |

The exact chain can vary with the installed `Azure.Identity` version and options. See Microsoft’s current [credential-chain documentation](https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains).

Environment credentials take precedence over developer logins. Therefore, incomplete or stale `AZURE_*` variables are a common source of surprising behavior.

## 4. Local development versus Azure deployment

**Local development**

The environment, workload identity, and managed identity entries are normally unavailable, so the chain proceeds to developer tools such as Visual Studio or Azure CLI. The application operates as the signed-in developer, and that developer must have the necessary Blob data role.

**Deployed to Azure**

The normal choice is:

- `WorkloadIdentityCredential` for a federated workload identity; or
- `ManagedIdentityCredential` for App Service, Functions, Container Apps, VMs, and other managed-identity-capable services.

Enable an identity on the hosting resource and grant that identity a Blob data role on the narrowest appropriate scope. No client secret should be stored in the application.

For a user-assigned managed identity, identify which one should be used:

```csharp
var credential = new DefaultAzureCredential(
    new DefaultAzureCredentialOptions
    {
        ManagedIdentityClientId =
            Environment.GetEnvironmentVariable("AZURE_CLIENT_ID")
    });
```

Although `DefaultAzureCredential` is convenient during development and deployment, Microsoft recommends using a deterministic credential such as `ManagedIdentityCredential` once production authentication requirements are fixed. Alternatively, `AZURE_TOKEN_CREDENTIALS=prod` limits recent `Azure.Identity` versions to deployed-service credentials.

## 5. Troubleshooting failures

1. **Enable Azure Identity diagnostics** to see each attempted credential and the one selected:

   ```csharp
   using System.Diagnostics.Tracing;
   using Azure.Core.Diagnostics;

   using var listener = new AzureEventSourceListener(
       (eventArgs, message) =>
       {
           if (eventArgs.EventSource.Name == "Azure-Identity")
           {
               Console.Error.WriteLine(message);
           }
       },
       EventLevel.Verbose);
   ```

   Add this before creating or using the client. Avoid leaving verbose identity logging enabled unnecessarily in production.

2. **Interpret the failure correctly:**
   - `CredentialUnavailableException`: a particular credential isn't configured or available. This is often normal while the chain is searching.
   - `AuthenticationFailedException`: Microsoft Entra authentication was attempted but rejected—for example, an expired secret, tenant mismatch, or invalid federation configuration.
   - HTTP `401`: the token is absent, invalid, or intended for the wrong tenant/resource.
   - HTTP `403`: authentication usually succeeded, but the selected identity lacks authorization. Check Blob **data-plane** RBAC roles.
   - DNS, timeout, or connection errors: check the service URI, storage firewall, private endpoint DNS, proxy, and outbound networking.

3. **Check which identity was selected.** A developer may be signed into multiple tenants or tools. Environment variables are tried before Visual Studio and CLI credentials, and Visual Studio is tried before Azure CLI.

4. **Check configuration:**
   - Service URI is `https://<account>.blob.core.windows.net`.
   - Local developer or deployed identity belongs to the expected tenant.
   - `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` are complete and current if environment authentication is intentional.
   - A user-assigned identity's client ID is configured when several identities are attached.
   - The machine clock is accurate.

5. **Check authorization separately.** Assign **Storage Blob Data Reader**, **Storage Blob Data Contributor**, or another suitable data role at the container or storage-account scope. New role assignments can take several minutes to propagate.

Official references: [DefaultAzureCredential API](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential), [Azure Identity credential chains](https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains), [Azure SDK logging](https://learn.microsoft.com/en-us/dotnet/azure/sdk/logging), and [Blob data-role assignment](https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access).

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response comprehensively covers package setup, credential construction and configuration, chain order, client integration, environment-specific behavior, and authentication troubleshooting.

### Strengths

- Complete runnable BlobServiceClient example
- Accurate ordered credential-chain explanation with version caveat
- Strong distinction between authentication and Blob data-plane authorization failures
- Actionable Azure Identity diagnostic logging example

## Grader Results

- default-azure-credential.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - `Azure.Identity` NuGet package installation: Pass
      - `DefaultAzureCredential` constructor and options: Pass
      - Credential chain: Environment → Workload Identity → Managed Identity → Azure CLI → etc.: Pass
      - Passing credential to any Azure SDK client constructor: Pass
      - `AuthenticationFailedException` handling and diagnostics: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-default-credential --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

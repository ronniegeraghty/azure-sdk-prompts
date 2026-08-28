# Evaluation Report: identity-dp-dotnet-default-credential

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 386.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-default-credential` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 4/5 |
| Duration | 386.2s |
| Timestamp | 2026-08-27T14:40:00Z |
| Files Generated | 0 |
| Event Count | 3540 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 353.7s |
| Review | 31.3s |
| **Total** | **386.2s** |

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
| Token Usage | in=251023 out=3417 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3281ms

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

### ✅ 2. view — 189270ms

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

### ✅ 3. view — 187893ms

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

### ❌ 4. azure-get_azure_bestpractices (via azure) — 189276ms

**Input:**

```json
{"intent":"Provide modern C# code showing Azure BlobServiceClient authentication with DefaultAzureCredential, including local and hosted behavior and troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 189276ms

**Input:**

```json
{"intent":"Find current official documentation for DefaultAzureCredential credential chain order and continuation policy in Azure.Identity for .NET, plus BlobServiceClient authentication with token credentials.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_search (via github-mcp-server) — 69251ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains DefaultAzureCredential .NET credential order 2026 WorkloadIdentityCredential BrokerCredential"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 7. web_search (via github-mcp-server) — 67989ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api/azure.identity.defaultazurecredential Azure.Identity latest credential order continuation policy"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 8. web_search (via github-mcp-server) — 66673ms

**Input:**

```json
{"query":"site:learn.microsoft.com/azure/storage/blobs/storage-blob-dotnet-get-started DefaultAzureCredential BlobServiceClient RBAC Storage Blob Data Contributor troubleshooting"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 9. web_fetch — 7640ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains"}
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

### ✅ 10. web_fetch — 7166ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-dotnet-get-started:
---
layout: Conceptual
title: Get Started with Azure Blob Storage and .NET - Azure Storage | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-dotnet-get-started
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
description: Learn how to connect a .NET app to Azure Blob Storage, set up secure authorization, and start building now with this step-by-step guide.
services: storage
author: stevenmatthew
ms.author: shaas
ms.service: azure-blob-storage
ms.topic: how-to
ms.date: 2026-06-02T00:00:00.0000000Z
ms.devlang: csharp
ms.custom:
- template-how-to
- devguide-csharp
- devx-track-dotnet
- sfi-ropc-nochange
locale: en-us
document_id: 914843bb-eff9-df8d-8e6b-cb812785888e
document_version_independent_id: 4157c7c4-1f14-26b5-c169-c4b3e0e3ee00
updated_at: 2026-06-03T05:14:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-pr/blob/live/articles/storage/blobs/storage-blob-dotnet-get-started.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-pr/blob/5d4ec1910b8f3ea24610482cf67c649da6ec133b/articles/storage/blobs/storage-blob-dotnet-get-started.md
git_commit_id: 5d4ec1910b8f3ea24610482cf67c649da6ec133b
site_name: Docs
depot_name: Azure.azure-documents
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/Azure.azure-documents/{branchName}{pdfName}
word_count: 1297
asset_id: storag
... (truncated)
```

### ❌ 11. web_fetch — 6647ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-dotnet-get-started"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting - status code 404
```

### ✅ 12. web_fetch — 6643ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting"}
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
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Identity/DefaultAzureCredential.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Identity/DefaultAzureCredential.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Identity.DefaultAzureCredential
description: "Simplifies authentication while developing apps tha
... (truncated)
```

### ✅ 13. web_fetch — 4485ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting-default-azure-credential"}
```

**Output:**

```
Contents of https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/identity/Azure.Identity/TROUBLESHOOTING.md:
# Troubleshoot Azure Identity authentication issues

[](about:blank#troubleshoot-azure-identity-authentication-issues)

This troubleshooting guide covers failure investigation techniques, common errors for the credential types in the Azure Identity library for .NET, and mitigation steps to resolve these errors.

## Handle Azure Identity exceptions

[](about:blank#handle-azure-identity-exceptions)

### AuthenticationFailedException

[](about:blank#authenticationfailedexception)

Exceptions arising from authentication errors can be raised on any service client method that makes a request to the service. This is because the token is requested from the credential on:

*   The first call to the service.
*   Any subsequent requests to the service that need to refresh the token.

To distinguish these failures from failures in the service client, Azure Identity classes raise the `AuthenticationFailedException` with details describing the source of the error in the exception message and possibly the error message. Depending on the application, these errors may or may not be recoverable.

usingAzure.Identity;usingAzure.Security.KeyVault.Secrets;// Create a secret client using the DefaultAzureCredentialvarclient=newSecretClient(newUri("https://myvault.vault.azure.net/"),newDefaultAzureCredential());try{KeyVaultSecretsecret=awaitclient.GetSecretAsync("secret1");}catch(AuthenticationFailedExceptione){Console.WriteLine($"Authentication Failed. {e.Message}");}

### CredentialUnavailableException

[](about:blank#credentialunavailableexception)

The `CredentialUnavailableException` is a special exception type derived from `AuthenticationFailedException`. This exception type is used to indicate that the credential can't authenticate in the current environment, due to lack of required configuration or setup. This exception is also used as a signal to chained credential type
... (truncated)
```

### ✅ 14. web_fetch — 3855ms

**Input:**

```json
{"max_length":15000,"url":"https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/identity/Azure.Identity/TROUBLESHOOTING.md"}
```

**Output:**

```
Contents of https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/identity/Azure.Identity/README.md:
# Azure Identity client library for .NET

[](about:blank#azure-identity-client-library-for-net)

The Azure Identity library provides [Microsoft Entra ID](https://learn.microsoft.com/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [`TokenCredential`](https://learn.microsoft.com/dotnet/api/azure.core.tokencredential?view=azure-dotnet) implementations that can be used to construct Azure SDK clients that support Microsoft Entra token authentication.

[Source code](https://github.com/Azure/azure-sdk-for-net/tree/main/sdk/identity/Azure.Identity/src) | [Package (NuGet)](https://www.nuget.org/packages/Azure.Identity) | [API reference documentation](https://learn.microsoft.com/dotnet/api/azure.identity?view=azure-dotnet) | [Microsoft Entra ID documentation](https://learn.microsoft.com/entra/identity/)

Install the Azure Identity client library for .NET with NuGet:

```
dotnet add package Azure.Identity
```

Note

Starting with `Azure.Core` 1.53.0, all credential types (including `DefaultAzureCredential`) are bundled directly in `Azure.Core`. If your project already targets `Azure.Core` 1.53.0+ or does so through a transitive dependency, you should omit the `Azure.Identity` package reference entirely. See the [Migration Guide](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/identity/Azure.Identity/MigrationGuide.md) for upgrade and compatibility details.

*   An [Azure subscription](https://azure.microsoft.com/free/dotnet/).
*   The [Azure CLI](https://learn.microsoft.com/cli/azure) can also be useful for authenticating in a development environment, creating accounts, and managing account roles.

When debugging and executing code locally, it's typical for a developer to use their own account for authenticating calls to Azure services. There are several developer tools that can be used to perform this authent
... (truncated)
```

### ❌ 15. web_fetch — 3222ms

**Input:**

```json
{"max_length":15000,"url":"https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/identity/Azure.Identity/README.md"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting-default-azure-credential - status code 404
```

## Copilot Response

I’m grounding the example in the current .NET identity chain and Azure Storage SDK guidance, since the continuation order can vary by `Azure.Identity` version.

The SDK references confirm the two package names and passwordless client construction. I’m checking the version-sensitive chain details against Microsoft’s current documentation before listing the exact sequence.

Use `Azure.Identity` for `DefaultAzureCredential` and `Azure.Storage.Blobs` for `BlobServiceClient`.

## 1. NuGet packages

```powershell
dotnet new console --name DefaultAzureCredentialBlobDemo
Set-Location DefaultAzureCredentialBlobDemo

dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

Optional:

```powershell
dotnet add package Azure.Identity.Broker
```

`Azure.Identity.Broker` enables broker-backed authentication and Visual Studio Code SSO on supported platforms. It isn't required for Azure CLI, Visual Studio, workload identity, managed identity, or environment-based authentication.

## 2. Complete runnable example

Replace `Program.cs` with:

```csharp
using Azure.Identity;
using Azure.Storage.Blobs;

string serviceUri =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_BLOB_SERVICE_URI")
    ?? "https://exampleaccount.blob.core.windows.net";

var credential = new DefaultAzureCredential();

var blobServiceClient = new BlobServiceClient(
    new Uri(serviceUri),
    credential);

Console.WriteLine("BlobServiceClient created successfully.");
Console.WriteLine($"Service endpoint: {blobServiceClient.Uri}");
```

Run it:

```powershell
$env:AZURE_STORAGE_BLOB_SERVICE_URI = "https://<storage-account>.blob.core.windows.net"
dotnet run
```

Client construction does not make a network request, so this program can run without contacting Azure. Authentication happens lazily when the first service operation is performed. For example:

```csharp
await foreach (var container in blobServiceClient.GetBlobContainersAsync())
{
    Console.WriteLine(container.Name);
}
```

That operation requires connectivity, a valid Azure identity, and an appropriate data-plane role such as **Storage Blob Data Reader**. Creating or modifying blobs normally requires **Storage Blob Data Contributor**.

Reuse credential and client instances rather than creating them for every request; both are thread-safe and credentials cache acquired tokens.

## 3. Credential chain order

With the current `Azure.Identity` behavior, `DefaultAzureCredential` tries:

| Order | Credential | Typical source |
|---:|---|---|
| 1 | `EnvironmentCredential` | `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and a secret or certificate |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS |
| 3 | `ManagedIdentityCredential` | System-assigned or user-assigned Azure managed identity |
| 4 | `VisualStudioCredential` | Account signed into Visual Studio |
| 5 | `VisualStudioCodeCredential` | VS Code Azure sign-in; requires `Azure.Identity.Broker` |
| 6 | `AzureCliCredential` | Account selected by `az login` |
| 7 | `AzurePowerShellCredential` | Account selected by `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account selected by `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser login; **disabled by default** |
| 10 | Broker credential | OS account broker; requires `Azure.Identity.Broker` |

The chain stops when a credential obtains a token.

Since `Azure.Identity` 1.10.1, failures from developer-tool credentials don't prevent later developer credentials from being attempted. Deployed-service credentials behave more strictly: if one is configured and can attempt authentication but token acquisition fails, the exception stops the chain. This prevents a misconfigured production identity from silently falling back to an unrelated developer identity.

Interactive browser authentication can be explicitly enabled:

```csharp
var credential = new DefaultAzureCredential(
    new DefaultAzureCredentialOptions
    {
        ExcludeInteractiveBrowserCredential = false
    });
```

Avoid enabling interactive authentication in unattended applications.

## 4. Local development versus Azure deployment

**Local development:** environment, workload identity, and managed identity are usually unavailable, so the chain reaches developer credentials. Sign in using one supported tool, such as Visual Studio, Azure CLI, Azure PowerShell, or Azure Developer CLI. The signed-in user must also have a Blob Storage data-plane role.

**Deployed to Azure:** workload identity or managed identity should normally provide the token before developer credentials are considered. Enable an identity on the hosting resource and assign it an appropriate role on the storage account or container.

For a user-assigned managed identity:

```csharp
var credential = new DefaultAzureCredential(
    new DefaultAzureCredentialOptions
    {
        ManagedIdentityClientId = "<managed-identity-client-id>"
    });
```

You can alternatively set `AZURE_CLIENT_ID`.

Although `DefaultAzureCredential` is convenient when the same build runs locally and in Azure, production applications often benefit from using the expected credential directly:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

TokenCredential credential = new ManagedIdentityCredential();

var client = new BlobServiceClient(
    new Uri("https://<storage-account>.blob.core.windows.net"),
    credential);
```

This is more deterministic and avoids unnecessary fallback attempts.

## 5. Troubleshooting authentication failures

### Enable Azure Identity diagnostics

```csharp
using Azure.Core.Diagnostics;
using Azure.Identity;
using Azure.Storage.Blobs;

using AzureEventSourceListener listener =
    AzureEventSourceListener.CreateConsoleLogger();

var credential = new DefaultAzureCredential(
    new DefaultAzureCredentialOptions
    {
        Diagnostics =
        {
            IsAccountIdentifierLoggingEnabled = true
        }
    });

var client = new BlobServiceClient(
    new Uri("https://<storage-account>.blob.core.windows.net"),
    credential);

await foreach (var container in client.GetBlobContainersAsync())
{
    Console.WriteLine(container.Name);
}
```

Account identifiers can be sensitive, so enable that logging only during controlled troubleshooting.

### Interpret the failure

- **`CredentialUnavailableException`**: a credential isn't configured for the current environment. In a chain this normally means “try the next credential.”
- **`AuthenticationFailedException`**: a configured credential attempted authentication but failed. Inspect the inner errors, Microsoft Entra error code, correlation ID, and timestamp.
- **HTTP 401**: the token is missing, invalid, expired, for the wrong tenant, or has the wrong audience.
- **HTTP 403**: authentication probably succeeded, but the identity lacks permission. Assign a Blob **data-plane** role—not merely `Owner` or `Contributor`, which are management-plane roles.

Also check:

1. Remove incomplete `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, or `AZURE_CLIENT_SECRET` variables that unintentionally activate `EnvironmentCredential`.
2. Confirm the developer tool is signed into the intended tenant and account.
3. Use diagnostic logs to identify which credential actually supplied the token.
4. Ensure the endpoint is exactly `https://<account>.blob.core.windows.net`.
5. For managed identity, verify that the identity is enabled, the correct user-assigned client ID is configured, and role assignment propagation has completed.
6. Narrow the chain while diagnosing. With `Azure.Identity` 1.15.0 or later, `AZURE_TOKEN_CREDENTIALS` can be set to `dev`, `prod`, or a credential name such as `AzureCliCredential`.

Official references: [DefaultAzureCredential API](https://learn.microsoft.com/dotnet/api/azure.identity.defaultazurecredential), [credential chains](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains), and [Blob Storage with .NET](https://learn.microsoft.com/azure/storage/blobs/storage-blob-dotnet-get-started).

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 5 criteria passed. The response comprehensively covers packages, credential construction and options, chain order, client integration, environment differences, and diagnostics, but omits executable AuthenticationFailedException handling.

### Strengths

- Provides correct package installation commands and a runnable BlobServiceClient example.
- Clearly documents the credential chain and local-versus-deployed behavior.
- Includes useful Azure Identity diagnostic logging and distinguishes authentication failures from authorization failures.

### Issues

- No try/catch or equivalent handling for AuthenticationFailedException is included in the runnable or troubleshooting examples.

## Grader Results

- default-azure-credential.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/5)
      - `Azure.Identity` NuGet package installation: Pass
      - `DefaultAzureCredential` constructor and options: Pass
      - Credential chain: Environment → Workload Identity → Managed Identity → Azure CLI → etc.: Pass
      - Passing credential to any Azure SDK client constructor: Pass
      - `AuthenticationFailedException` handling and diagnostics: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 80% | 1.00 | 0.8000 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8000** | **80.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-default-credential --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

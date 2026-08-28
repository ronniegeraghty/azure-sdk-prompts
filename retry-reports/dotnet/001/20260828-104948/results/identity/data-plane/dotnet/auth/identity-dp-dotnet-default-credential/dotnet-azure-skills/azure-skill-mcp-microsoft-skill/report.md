# Evaluation Report: identity-dp-dotnet-default-credential

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 537.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-default-credential` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 537.4s |
| Timestamp | 2026-08-28T02:49:48Z |
| Files Generated | 0 |
| Event Count | 3951 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 505.8s |
| Review | 30.4s |
| **Total** | **537.4s** |

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
| Token Usage | in=237342 out=4096 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 1184ms

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

### ❌ 2. view — 183898ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 183896ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 183900ms

**Input:**

```json
{"intent":"Get current best practices for C# code using Azure.Identity DefaultAzureCredential with Azure.Storage.Blobs BlobServiceClient, including local development, managed identity, and authentication troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 183900ms

**Input:**

```json
{"intent":"Find current official Microsoft documentation for DefaultAzureCredential in .NET, including exact credential chain order, continuation policy, local versus Azure-hosted behavior, and troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. glob — 181666ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-dotnet\skills\azure-identity-dotnet\SKILL.md
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 181670ms

**Input:**

```json
{"intent":"C# Azure SDK authentication code generation with DefaultAzureCredential and BlobServiceClient.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 8. azure-documentation (via azure) — 181673ms

**Input:**

```json
{"intent":"DefaultAzureCredential .NET credential chain order and troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 9. web_search (via github-mcp-server) — 34184ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains DefaultAzureCredential order .NET WorkloadIdentityCredential VisualStudioCredential AzureCliCredential AzurePowerShellCredential AzureDeveloperCliCredential broker 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To troubleshoot `DefaultAzureCredential` issues and enable detailed logging in Azure Identity for .NET, you can use `AzureEventSourceListener`. This provides visibility into what each credential in the chain is doing and why an `AuthenticationFailedException` might be thrown.\n\n### Enabling Logging with AzureEventSourceListener\n\nYou can wire up detailed logs (to the console, for example) by registering an event source listener. The Azure SDK recommends using `Azure.Core.Diagnostics.AzureEventSourceListener`:\n\n```csharp\nusing Azure.Core.Diagnostics;\n\n// At the start of your application's Program.cs or Main method:\nusing var listener = AzureEventSourceListener.CreateConsoleLogger();\n```\n\nThis will log all events from the Azure SDK, including information about which credentials are being attempted in `DefaultAzureCredential`, what is being excluded, and why authentication is failing.\n\n### Understanding AuthenticationFailedException\n\n`DefaultAzureCredential` tries multiple credential sources in a specific order (including environment variables, managed identity, Azure CLI, etc.). If every possible credential fails, you'll get an `AuthenticationFailedException`, and the logs will show which source failed and why.\n\nCommon reasons might include:\n\n- Missing required environment variables for `EnvironmentCredential`\n- No managed identity assigned, or unavailable endpoint for `ManagedIdentityCredential`\n- Not logged in with Azure CLI or Visual Studio\n\n### Next Steps for Troubleshooting\n\n- **Check the logs** generated by `AzureEventSourceListener` for messages about excluded credentials, missing configuration, or other issues.\n- **Explicitly include or exclude credentials** by configuring `DefaultAzureCredentialOptions` if you have a specific authentication flow.\n- **Review Azure Identity troubleshooting documentation** for detailed guidance on resolving failures: [Azure Identity troubleshooting for .NET](https
... (truncated)
```

### ✅ 10. web_search (via github-mcp-server) — 33836ms

**Input:**

```json
{"query":"site:learn.microsoft.com azure identity dotnet troubleshooting DefaultAzureCredential logging AzureEventSourceListener AuthenticationFailedException"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 11. web_fetch — 33426ms

**Input:**

```json
{"max_length":12000,"url":"https://aka.ms/azsdk/net/identity/credential-chains#defaultazurecredential-overview"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains#defaultazurecredential-overview (redirected from https://aka.ms/azsdk/net/identity/credential-chains#defaultazurecredential-overview):
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
- https://au
... (truncated)
```

### ✅ 12. web_fetch — 4295ms

**Input:**

```json
{"max_length":10000,"start_index":12000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity (redirected from https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/azure-hosted-apps):
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
source_path: docs/azure/sdk/authentication/system-assi
... (truncated)
```

### ❌ 13. web_fetch — 3919ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/on-premises-apps-dev-overview"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting - status code 404
```

### ❌ 14. web_fetch — 3917ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/azure-hosted-apps"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/on-premises-apps-dev-overview - status code 404
```

### ✅ 15. web_fetch — 3910ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshooting"}
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

### ✅ 16. web_fetch — 3521ms

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains:
ain looks as follows:

![DefaultAzureCredential with AZURE_TOKEN_CREDENTIALS set to 'prod'](../media/mermaidjs/default-azure-credential-environment-variable-production.svg)

When a value of `dev` is used, the chain looks as follows:

![DefaultAzureCredential with AZURE_TOKEN_CREDENTIALS set to 'dev'](../media/mermaidjs/default-azure-credential-environment-variable-development.svg)

To ensure the environment variable is defined and set to a supported string, use constructor overload [DefaultAzureCredential(String, DefaultAzureCredentialOptions)](/en-us/dotnet/api/azure.identity.defaultazurecredential.-ctor#azure-identity-defaultazurecredential-ctor%28system-string-azure-identity-defaultazurecredentialoptions%29).

#### Use a specific credential

To exclude all credentials except for one, set environment variable `AZURE_TOKEN_CREDENTIALS` to the credential name. For example, you can reduce the `DefaultAzureCredential` chain to `VisualStudioCredential` by setting `AZURE_TOKEN_CREDENTIALS` to `VisualStudioCredential`. The string comparison is performed in a case-insensitive manner. Valid string values for the environment variable include:

- `AzureCliCredential`
- `AzureDeveloperCliCredential`
- `AzurePowerShellCredential`
- `BrokerCredential`
- `EnvironmentCredential`
- `InteractiveBrowserCredential`
- `ManagedIdentityCredential`
- `VisualStudioCredential`
- `VisualStudioCodeCredential`
- `WorkloadIdentityCredential`

Important

The `AZURE_TOKEN_CREDENTIALS` environment variable supports individual credential names in `Azure.Identity` package versions 1.15.0 and later.

To ensure the environment variable is defined and set to a supported string, use constructor overload [DefaultAzureCredential(String, DefaultAzureCredentialOptions)](/en-us/dotnet/api/azure.identity.defaultazurecredential.-ctor#azure-identity-defaultazurecredential-ctor%28system-
... (truncated)
```

## Copilot Response

I’ll ground the example in the current Azure.Identity chain and SDK guidance so the credential order and troubleshooting steps match the latest .NET behavior.

## 1. NuGet packages

Create a console app and install the authentication and Blob Storage libraries:

```powershell
dotnet new console --framework net8.0 -n DefaultAzureCredentialDemo
Set-Location DefaultAzureCredentialDemo

dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

| Package | Purpose |
|---|---|
| `Azure.Identity` | Provides `DefaultAzureCredential` and other Microsoft Entra credentials |
| `Azure.Storage.Blobs` | Provides `BlobServiceClient` and Blob Storage APIs |
| `Azure.Identity.Broker` | Optional: enables brokered OS authentication and Visual Studio Code credentials |
| `Microsoft.Extensions.Azure` | Optional: Azure SDK dependency injection for ASP.NET Core |

`Azure.Core` is required internally but is installed transitively.

## 2. Complete runnable example

Replace `Program.cs` with:

```csharp
using System.Diagnostics.Tracing;
using Azure;
using Azure.Core.Diagnostics;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;

string accountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME")
    ?? throw new InvalidOperationException(
        "Set the AZURE_STORAGE_ACCOUNT_NAME environment variable.");

Uri serviceUri = new($"https://{accountName}.blob.core.windows.net");

// Enable Azure Identity diagnostics when requested.
using AzureEventSourceListener? listener =
    Environment.GetEnvironmentVariable("AZURE_IDENTITY_LOGGING") == "1"
        ? new AzureEventSourceListener(
            (args, message) =>
            {
                if (args.EventSource.Name == "Azure-Identity")
                {
                    Console.Error.WriteLine(message);
                }
            },
            EventLevel.LogAlways)
        : null;

// Credentials are thread-safe and should be reused between SDK clients.
var credential = new DefaultAzureCredential();

var blobServiceClient = new BlobServiceClient(serviceUri, credential);

Console.WriteLine($"Created BlobServiceClient for {blobServiceClient.Uri}");

// Client construction is local and doesn't request a token. Pass --list to
// make an authenticated request and list containers.
if (!args.Contains("--list", StringComparer.OrdinalIgnoreCase))
{
    Console.WriteLine("Run with --list to make an authenticated Azure request.");
    return;
}

try
{
    await foreach (
        BlobContainerItem container in blobServiceClient.GetBlobContainersAsync())
    {
        Console.WriteLine(container.Name);
    }
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Authentication failed: {ex.Message}");
    Environment.ExitCode = 1;
}
catch (RequestFailedException ex) when (ex.Status is 401 or 403)
{
    Console.Error.WriteLine(
        $"Azure rejected the request ({ex.Status}, {ex.ErrorCode}). " +
        "Authentication may have succeeded, but the identity might not have " +
        $"permission to access this storage account. {ex.Message}");

    Environment.ExitCode = 1;
}
```

Run it locally:

```powershell
$env:AZURE_STORAGE_ACCOUNT_NAME = "your-storage-account-name"

# Constructs the client without making a network request:
dotnet run

# Authenticates and lists containers:
dotnet run -- --list

# Enable detailed credential-chain diagnostics:
$env:AZURE_IDENTITY_LOGGING = "1"
dotnet run -- --list
```

`DefaultAzureCredential` doesn't authenticate during construction. It requests a token only when an SDK operation such as `GetBlobContainersAsync` sends a request.

The selected identity needs a Blob **data-plane** role, such as **Storage Blob Data Reader** to list/read or **Storage Blob Data Contributor** to write. Subscription `Owner` or `Contributor` alone doesn't necessarily grant access to Blob data.

## 3. Credential chain order

In the current `Azure.Identity` documentation, `DefaultAzureCredential` tries credentials in this order:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal variables such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly on AKS |
| 3 | `ManagedIdentityCredential` | System-assigned or user-assigned Azure managed identity |
| 4 | `VisualStudioCredential` | Account signed into Visual Studio |
| 5 | `VisualStudioCodeCredential` | VS Code Azure account; requires `Azure.Identity.Broker` |
| 6 | `AzureCliCredential` | Account authenticated through `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated through `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated through `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser login; **disabled by default** |
| 10 | `BrokerCredential` | OS account broker; requires `Azure.Identity.Broker` |

The chain stops after one credential obtains a token. The exact chain can vary with the installed `Azure.Identity` version and options.

Developer-tool credentials continue through the developer portion of the chain when one fails. Deployed-service credentials are stricter: if one appears configured but fails authentication, the error is normally surfaced rather than silently falling through to a developer identity. This helps prevent a deployed application from unexpectedly using another account.

## 4. Local development versus deployed Azure

**Local development:** Workload and managed identity usually report themselves unavailable. The chain then reaches Visual Studio, VS Code, Azure CLI, Azure PowerShell, or Azure Developer CLI and uses your signed-in developer account. That account must have the required Blob data role.

**Deployed to Azure:** The same code typically authenticates using workload identity or managed identity. No password or client secret needs to be stored in application configuration. The Azure hosting resource must have an identity enabled, and that identity must receive the required RBAC role on the storage account or container.

For a user-assigned managed identity, explicitly select it:

```csharp
var credential = new DefaultAzureCredential(
    new DefaultAzureCredentialOptions
    {
        ManagedIdentityClientId =
            Environment.GetEnvironmentVariable("AZURE_MANAGED_IDENTITY_CLIENT_ID")
    });
```

Although `DefaultAzureCredential` is convenient across environments, production applications often benefit from a deterministic credential such as `ManagedIdentityCredential`, which avoids unexpected fallback behavior:

```csharp
var credential = new ManagedIdentityCredential(
    ManagedIdentityId.SystemAssigned);
```

## 5. Troubleshooting failures

1. **Enable Azure Identity logging.** In the example, set `AZURE_IDENTITY_LOGGING=1`. Look for `DefaultAzureCredential credential selected:` and the reason each preceding credential was unavailable or failed.

2. **Read the complete exception, including inner exceptions.** `DefaultAzureCredential` reports the result of every credential it attempted.

3. **Distinguish authentication from authorization.**
   - `AuthenticationFailedException` generally means no credential obtained a usable token.
   - HTTP `401` usually indicates an invalid token, tenant, authority, or audience.
   - HTTP `403` usually means authentication succeeded but the identity lacks a Blob data role.

4. **Check local login state.** Confirm that Visual Studio or the relevant CLI is signed in to the intended tenant and subscription. Multiple cached accounts are a frequent source of confusion.

5. **Check environment variables.** Remove stale or partially configured `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, or `AZURE_CLIENT_SECRET` values. Because `EnvironmentCredential` is first, valid-looking but incorrect values can take precedence over your local developer login.

6. **Check managed identity configuration.** Confirm that the identity is enabled on the Azure host. If multiple user-assigned identities exist, specify `ManagedIdentityClientId`.

7. **Check Blob data-plane RBAC.** Assign the identity **Storage Blob Data Reader** or **Storage Blob Data Contributor** at the narrowest appropriate scope. New assignments can take several minutes to propagate.

8. **Reduce the chain while diagnosing.** With `Azure.Identity` 1.15 or later, `AZURE_TOKEN_CREDENTIALS` can select one credential:

```powershell
$env:AZURE_TOKEN_CREDENTIALS = "AzureCliCredential"
dotnet run -- --list
```

Official references: [DefaultAzureCredential chain](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains), [Azure Identity logging](https://learn.microsoft.com/dotnet/azure/sdk/logging), and [Blob data-access roles](https://learn.microsoft.com/azure/storage/blobs/assign-azure-role-data-access).

## Reviewer Notes (LLM-as-Judge)

### Summary

All 5 criteria passed. The response provides a complete runnable BlobServiceClient example and thoroughly explains authentication behavior, credential ordering, environment differences, and diagnostics.

### Strengths

- Includes all required NuGet installation commands and a runnable C# example.
- Accurately explains lazy authentication and local-versus-Azure credential selection.
- Provides explicit authentication exception handling, diagnostic logging, RBAC guidance, and practical troubleshooting.

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
hyoka run --prompt-id identity-dp-dotnet-default-credential --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

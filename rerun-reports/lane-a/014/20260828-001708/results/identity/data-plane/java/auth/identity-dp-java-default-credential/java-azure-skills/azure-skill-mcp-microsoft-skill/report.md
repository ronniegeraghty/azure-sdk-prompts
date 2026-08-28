# Evaluation Report: identity-dp-java-default-credential

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 670.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-default-credential` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 670.5s |
| Timestamp | 2026-08-27T16:17:08Z |
| Files Generated | 0 |
| Event Count | 5305 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 406.0s |
| Review | 263.3s |
| **Total** | **670.5s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=321419 out=4431 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 4786ms

**Input:**

```json
{"skill":"azure-identity-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for Java

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.15.0</version>
</dependency>
```

## Key Concepts

| Credential | Use Case |
|------------|----------|
| `DefaultAzureCredential` | **Recommended** - Works in dev and production |
| `ManagedIdentityCredential` | Azure-hosted apps (App Service, Functions, VMs) |
| `EnvironmentCredential` | CI/CD pipelines with env vars |
| `ClientSecretCredential` | Service principals with secret |
| `ClientCertificateCredential` | Service principals with certificate |
| `AzureCliCredential` | Local dev using `az login` |
| `InteractiveBrowserCredential` | Interactive login flow |
| `DeviceCodeCredential` | Headless device authentication |

## DefaultAzureCredential (Recommended)

The `DefaultAzureCredential` tries multiple authentication methods in order. See [DefaultAzureCredential overview](https://aka.ms/azsdk/java/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

```java
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

// Simple usage
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();

// Use with any Azure client
BlobServiceClient blobClient = new BlobServiceClientBuilder()
    .endpoint("https://<storage-account>.blob.core.windows.net")
    .credential(credential)
    .buildClient();

KeyClient keyClient = new KeyClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();
```

### Configure DefaultAzureCredential

```java
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId("<user-assigned-identity-client-id>")  // For user-assigned MI
    .tenantId("<tenant-id
... (truncated)
```

### ✅ 2. view — 190959ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
@@ -1,495 +1,495 @@
 # Azure Identity library for Java - Examples
 
 Comprehensive code examples for the Azure Identity library for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [DefaultAzureCredential](#defaultazurecredential)
 - [ChainedTokenCredential](#chainedtokencredential)
 - [ClientSecretCredential](#clientsecretcredential)
 - [ClientCertificateCredential](#clientcertificatecredential)
 - [ManagedIdentityCredential](#managedidentitycredential)
 - [EnvironmentCredential](#environmentcredential)
 - [InteractiveBrowserCredential](#interactivebrowsercredential)
 - [DeviceCodeCredential](#devicecodecredential)
 - [AzureCliCredential](#azureclicredential)
 - [Using Credentials with Azure SDK Clients](#using-credentials-with-azure-sdk-clients)
 
 ## Maven Dependency
 
 ```xml
 <!-- Using Azure SDK BOM (recommended) -->
 <dependencyManagement>
     <dependencies>
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-sdk-bom</artifactId>
             <version>1.2.29</version>
             <type>pom</type>
             <scope>import</scope>
         </dependency>
     </dependencies>
 </dependencyManagement>
 
 <dependencies>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-identity</artifactId>
     </dependency>
 </dependencies>
 
 <!-- Or direct depend
... (truncated)
```

### ✅ 3. glob — 189476ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 190961ms

**Input:**

```json
{"intent":"Provide a modern local-only Java example using DefaultAzureCredential to create an Azure Key Vault SecretClient, including dependencies, credential chain, environment behavior, and troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 190962ms

**Input:**

```json
{"intent":"Find current official Microsoft documentation for Java DefaultAzureCredential credential chain order, continuation policy, local development versus Azure-hosted authentication, troubleshooting, and creating a Key Vault SecretClient.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_search (via github-mcp-server) — 87792ms

**Input:**

```json
{"query":"site:learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains DefaultAzureCredential Java order Environment Workload Identity Managed Identity 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To use the latest Azure SDK for Java Bill of Materials (BOM) with Maven for dependency management, you add the `azure-sdk-bom` to your `<dependencyManagement>` section in your `pom.xml`. This allows you to align all Azure SDK dependencies to compatible versions without specifying individual versions for each library.\n\n**Example snippet for your `pom.xml`:**\n\n```xml\n<dependencyManagement>\n    <dependencies>\n        <dependency>\n            <groupId>com.azure</groupId>\n            <artifactId>azure-sdk-bom</artifactId>\n            <version>{bom_version_to_target}</version>\n            <type>pom</type>\n            <scope>import</scope>\n        </dependency>\n    </dependencies>\n</dependencyManagement>\n```\nReplace `{bom_version_to_target}` with the latest released version (check [azure-sdk-bom on Maven Central](https://search.maven.org/search?q=g:com.azure%20AND%20a:azure-sdk-bom) for the most current version).\n\nThen, in your `<dependencies>` section, add the Azure SDK libraries you need **without the version tag**:\n\n```xml\n<dependencies>\n    <dependency>\n        <groupId>com.azure</groupId>\n        <artifactId>azure-identity</artifactId>\n    </dependency>\n    <dependency>\n        <groupId>com.azure</groupId>\n        <artifactId>azure-storage-blob</artifactId>\n    </dependency>\n    <!-- Add other Azure SDK libraries as needed -->\n</dependencies>\n```\nThis approach ensures all Azure client libraries use versions tested for compatibility together【3:5†source】【3:8†source】【3:9†source】.\n\n**References:**\n- [Azure SDK for Java BOM README](https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable) (see \"Include the BOM file\")\n- [Azure SDK for Java official docs](https://learn.microsoft.com/en-us/java/azure/javasdk-dependency-version-management?view=azure-java-stable)","annotations":[{"text":"","start_index":1462,"end_index":1474,"url_citation":{"t
... (truncated)
```

### ✅ 7. web_search (via github-mcp-server) — 86465ms

**Input:**

```json
{"query":"site:learn.microsoft.com Java Key Vault SecretClient DefaultAzureCredential Maven azure-security-keyvault-secrets azure-identity"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 8. web_search (via github-mcp-server) — 84655ms

**Input:**

```json
{"query":"site:learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-default-azure-credential Java DefaultAzureCredential troubleshoot logging"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here is how to use the Azure Key Vault SecretClient with DefaultAzureCredential in a Java Maven project using the `azure-security-keyvault-secrets` and `azure-identity` SDKs:\n\n### Maven dependencies\n\nIn your `pom.xml` file, include the following dependencies. You can use the SDK BOM (Bill of Materials) to manage versions:\n\n```xml\n<dependencyManagement>\n  <dependencies>\n    <dependency>\n      <groupId>com.azure</groupId>\n      <artifactId>azure-sdk-bom</artifactId>\n      <version>1.2.16</version><!-- replace with latest BOM version -->\n      <type>pom</type>\n      <scope>import</scope>\n    </dependency>\n  </dependencies>\n</dependencyManagement>\n\n<dependencies>\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-security-keyvault-secrets</artifactId>\n  </dependency>\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-identity</artifactId>\n  </dependency>\n</dependencies>\n```\nAlternatively, specify explicit versions if not using the BOM.\n\n### Setting up SecretClient with DefaultAzureCredential\n\n```java\nimport com.azure.identity.DefaultAzureCredentialBuilder;\nimport com.azure.security.keyvault.secrets.SecretClient;\nimport com.azure.security.keyvault.secrets.SecretClientBuilder;\n\npublic class KeyVaultExample {\n    public static void main(String[] args) {\n        String keyVaultUrl = \"https://<your-key-vault-name>.vault.azure.net/\";\n\n        SecretClient secretClient = new SecretClientBuilder()\n            .vaultUrl(keyVaultUrl)\n            .credential(new DefaultAzureCredentialBuilder().build())\n            .buildClient();\n\n        // Example: Set and get a secret\n        secretClient.setSecret(\"mySecretName\", \"mySecretValue\");\n        String secretValue = secretClient.getSecret(\"mySecretName\").getValue();\n        System.out.println(\"Secret value: \" + secretValue);\n    }\n}\n```\n- `DefaultAzureCredential` works seamlessly both for local
... (truncated)
```

### ✅ 9. web_search (via github-mcp-server) — 83290ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure SDK for Java BOM Maven latest azure-sdk-bom dependency management"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 10. web_fetch — 4564ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Identity client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-08-24T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-identity, entra-id
ms.service: entra-id
locale: en-us
document_id: a0c04cc9-4127-ee61-109d-2d672636a46e
document_version_independent_id: 0e5b5560-efe9-0c0a-9a63-38dc58e6af8b
updated_at: 2026-08-25T11:08:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/identity-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/75b1775676c4b42b2562bb1c23a9a8a867297328/docs-ref-services/latest/identity-readme.md
git_commit_id: 75b1775676c4b42b2562bb1c23a9a8a867297328
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 2370
asset_id: api/overview/azure/identity-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/identity-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
spProducts:
- https://authoring-docs-microsoft
... (truncated)
```

### ✅ 11. web_fetch — 4177ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-azure-hosted"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains:
---
layout: Conceptual
title: Credential chains in the Azure Identity library for Java - Java on Azure | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains
breadcrumb_path: /azure/developer/breadcrumb/toc.json
uhfHeaderId: azure
feedback_system: Standard
recommendations: true
feedback_help_link_type: get-help-at-qna
feedback_help_link_url: https://learn.microsoft.com/answers/questions/
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
author: bmitchell287
manager: nitinme
ms.author: brendm
ms.devlang: java
ms.service: azure-java
ms.subservice: java-sdk
description: This article describes the DefaultAzureCredential and ChainedTokenCredential classes in the Azure Identity library.
ms.date: 2026-02-24T00:00:00.0000000Z
ms.topic: concept-article
ms.reviewer: scaddie
ms.custom: devx-track-java
ai-usage: ai-assisted
locale: en-us
document_id: 54718866-4bea-1757-d2ff-144acdb1bb70
document_version_independent_id: 54718866-4bea-1757-d2ff-144acdb1bb70
updated_at: 2026-03-04T00:04:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/live/articles/java/sdk/authentication/credential-chains.md
gitcommit: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/ac67d79b6d52421f5635629e7be7d54c779d6c95/articles/java/sdk/authentication/credential-chains.md
git_commit_id: ac67d79b6d52421f5635629e7be7d54c779d6c95
site_name: Docs
depot_name: MSDN.azure-dev-docs
page_type: conceptual
toc_rel: ../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/MSDN.azure-dev-docs/{branchName}{pdfName}
word_count: 1453
asset_id: java/sdk/authentication/credential-chains
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/java/sdk/authentication/credential-chains.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/d
... (truncated)
```

### ❌ 12. web_fetch — 3796ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-dev-credential"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-dev-credential - status code 404
```

### ✅ 13. web_fetch — 3794ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-azure-hosted:
---
layout: Conceptual
title: Troubleshoot Azure-hosted application authentication - Azure SDK for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-azure-hosted
breadcrumb_path: /azure/developer/breadcrumb/toc.json
uhfHeaderId: azure
feedback_system: Standard
recommendations: true
feedback_help_link_type: get-help-at-qna
feedback_help_link_url: https://learn.microsoft.com/answers/questions/
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
author: bmitchell287
manager: nitinme
ms.author: brendm
ms.devlang: java
ms.service: azure-java
ms.subservice: java-sdk
description: Provides an overview of how to troubleshoot Azure-hosted authentication issues.
ms.date: 2026-07-27T00:00:00.0000000Z
ms.topic: troubleshooting-general
ms.reviewer: jogiles
ms.custom:
- devx-track-java
- devx-track-extended-java
- sfi-ropc-nochange
locale: en-us
document_id: 2d7b985e-8a97-686b-97c8-fe0f67c354ca
document_version_independent_id: 2d7b985e-8a97-686b-97c8-fe0f67c354ca
updated_at: 2026-07-27T17:05:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/live/articles/java/sdk/troubleshooting-authentication-azure-hosted.md
gitcommit: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/242d11c3bf5e37d84c7db9948dc37e357f9881f5/articles/java/sdk/troubleshooting-authentication-azure-hosted.md
git_commit_id: 242d11c3bf5e37d84c7db9948dc37e357f9881f5
site_name: Docs
depot_name: MSDN.azure-dev-docs
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/MSDN.azure-dev-docs/{branchName}{pdfName}
word_count: 1368
asset_id: java/sdk/troubleshooting-authentication-azure-hosted
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/java/sdk/troubleshooting-authentication-azure-hosted.md
... (truncated)
```

### ❌ 14. web_fetch — 31774ms

**Input:**

```json
{"max_length":4000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-sdk-bom%22\u0026rows=1\u0026wt=json"}
```

**Error:**

```
TimeoutError: The operation was aborted due to timeout
```

### ✅ 15. web_fetch — 31770ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/key-vault/general/rbac-guide:
---
layout: Conceptual
title: Grant permission to applications to access an Azure key vault using Azure RBAC | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/key-vault/general/rbac-guide
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
description: Learn how to provide access to keys, secrets, and certificates using Azure role-based access control.
services: key-vault
ms.service: azure-key-vault
ms.subservice: general
ms.custom: devx-track-azurecli, devx-track-azurepowershell, sfi-image-nochange, copilot-scenario-highlight
ms.topic: how-to
ms.date: 2026-06-16T00:00:00.0000000Z
locale: en-us
document_id: 44b9a84f-6765-67dc-bb7b-04303e070c03
document_version_independent_id: 32fc20a0-4321-c74a-34aa-909f43a72db0
updated_at: 2026-08-21T17:34:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-security-docs-pr/blob/live/articles/key-vault/general/rbac-guide.md
gitcommit: https://github.com/MicrosoftDocs/azure-security-docs-pr/blob/6502b5a33690f91073651cbd073ebd42edae2634/articles/key-vault/general/rbac-guide.md
git_commit_id: 6502b5a33690f91073651cbd073ebd42edae2634
site_name: Docs
depot_name: Learn.azure-security
page_type: conceptual
toc_rel: toc.json
word_count: 2657
asset_id: key-vault/general/rbac-guide
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/key-vault/general/rbac-guide.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/f48
... (truncated)
```

### ✅ 16. web_fetch — 31425ms

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/azure/key-vault/general/rbac-guide"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview:
---
layout: Conceptual
title: Troubleshoot Azure Identity authentication - Azure SDK for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview
breadcrumb_path: /azure/developer/breadcrumb/toc.json
uhfHeaderId: azure
feedback_system: Standard
recommendations: true
feedback_help_link_type: get-help-at-qna
feedback_help_link_url: https://learn.microsoft.com/answers/questions/
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
author: bmitchell287
manager: nitinme
ms.author: brendm
ms.devlang: java
ms.service: azure-java
ms.subservice: java-sdk
description: Learn how to troubleshoot Azure Identity authentication issues in the Azure SDK for Java. Resolve common errors and get mitigation steps.
ms.date: 2026-07-27T00:00:00.0000000Z
ms.topic: troubleshooting-general
ms.custom: devx-track-java, devx-track-extended-java
ms.reviewer: jogiles
locale: en-us
document_id: b76e7ea0-b79d-a913-1fdb-777cff4947af
document_version_independent_id: b76e7ea0-b79d-a913-1fdb-777cff4947af
updated_at: 2026-07-27T17:05:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/live/articles/java/sdk/troubleshooting-authentication-overview.md
gitcommit: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/242d11c3bf5e37d84c7db9948dc37e357f9881f5/articles/java/sdk/troubleshooting-authentication-overview.md
git_commit_id: 242d11c3bf5e37d84c7db9948dc37e357f9881f5
site_name: Docs
depot_name: MSDN.azure-dev-docs
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/MSDN.azure-dev-docs/{branchName}{pdfName}
word_count: 811
asset_id: java/sdk/troubleshooting-authentication-overview
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/java/sdk/troubleshooting-authentication-overview.md
... (truncated)
```

### ✅ 17. web_fetch — 1571ms

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Key Vault Secret client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-05-27T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-security-keyvault-secrets, keyvault
ms.service: keyvault
locale: en-us
document_id: 009f13e3-83f0-8f58-4be1-ebb785e6a344
document_version_independent_id: 6c845f98-c6a3-d3cd-c0b7-23f6775ef5ad
updated_at: 2026-05-28T11:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/security-keyvault-secrets-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/4f5bf09b69d6f35fcbd5a7a92ad75413fdbedaf8/docs-ref-services/latest/security-keyvault-secrets-readme.md
git_commit_id: 4f5bf09b69d6f35fcbd5a7a92ad75413fdbedaf8
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 1484
asset_id: api/overview/azure/security-keyvault-secrets-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/security-keyvault-secrets-readme.md
cmProducts:
- https://au
... (truncated)
```

## Copilot Response

I’ll ground the example in the current Java SDK credential-chain behavior and Key Vault client guidance.

The local skill references don’t include acceptance criteria, so I’m checking Microsoft’s current chain definition and package guidance directly.

`DefaultAzureCredential` lets the same Java code use developer credentials locally and workload/managed identity when deployed to Azure. Credential construction is lazy: authentication normally happens on the first service operation, such as `getSecret()`.

## 1. Maven dependencies

For a self-contained example, these versions match the current stable Microsoft documentation as of August 2026:

```xml
<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-identity</artifactId>
        <version>1.18.5</version>
    </dependency>

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
        <version>4.11.0</version>
    </dependency>
</dependencies>
```

For larger applications, Microsoft recommends importing the [`azure-sdk-bom`](https://learn.microsoft.com/java/azure/javasdk-dependency-version-management) and omitting versions from individual Azure dependencies. This keeps Azure libraries on mutually compatible versions.

If you want Visual Studio Code or OS-broker authentication, also add the optional broker package:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity-broker</artifactId>
    <version>1.1.13</version>
</dependency>
```

Use the current broker version from Maven Central if it has advanced.

## 2. Creating and using `DefaultAzureCredential`

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder().build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://my-vault.vault.azure.net")
    .credential(credential)
    .buildClient();
```

The credential can be reused by multiple Azure SDK clients. It handles token acquisition, caching, and refresh; do not create a new credential for each request.

For a user-assigned managed identity:

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder()
        .managedIdentityClientId(
            System.getenv("AZURE_CLIENT_ID"))
        .build();
```

## 3. Credential-chain order

The current Java `DefaultAzureCredential` chain tries:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service principal configuration in environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS |
| 3 | `ManagedIdentityCredential` | Azure managed identity |
| 4 | `IntelliJCredential` | Azure Toolkit for IntelliJ login |
| 5 | `VisualStudioCodeCredential` | VS Code Azure Resources login; broker package required |
| 6 | `AzureCliCredential` | Account from `az login` |
| 7 | `AzurePowerShellCredential` | Account from `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account from `azd auth login` |
| 9 | Broker credential | OS account broker; broker package required |

The first credential that successfully returns a token wins.

The continuation policy is important:

- **Developer credentials** are tried until one succeeds, even if an earlier developer credential reports an authentication error.
- **Deployed-service credentials**—environment, workload identity, and managed identity—stop the chain when they are configured and attempt authentication but fail. This makes production behavior more predictable.

With `azure-identity` 1.16.1 or later, you can restrict the chain:

```powershell
# Only deployed-service credentials
$env:AZURE_TOKEN_CREDENTIALS = "prod"

# Only developer credentials
$env:AZURE_TOKEN_CREDENTIALS = "dev"

# Exactly one credential; supported in 1.17.0+
$env:AZURE_TOKEN_CREDENTIALS = "AzureCliCredential"
```

## 4. Local development versus Azure

**Local development:** `DefaultAzureCredential` usually skips unavailable environment, workload, and managed identities, then authenticates through IntelliJ, VS Code, Azure CLI, Azure PowerShell, or `azd`.

A common setup is:

```powershell
az login
az account set --subscription "<subscription-id-or-name>"
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
```

The signed-in user must have Key Vault **data-plane** permission. Subscription `Owner` or `Contributor` alone does not necessarily grant permission to read secrets. For Azure RBAC, `Key Vault Secrets User` can read secret values; broader roles such as `Key Vault Secrets Officer` can manage secrets.

**Azure environments:** keep the application code unchanged and enable an identity on the host:

- App Service, Functions, or VM: managed identity.
- AKS: preferably workload identity.
- User-assigned managed identity: specify its client ID through `AZURE_CLIENT_ID` or `managedIdentityClientId()`.

Grant that identity access to the vault. Avoid client secrets in production when managed or workload identity is available.

Be careful with environment variables on Azure. A complete `AZURE_TENANT_ID`/`AZURE_CLIENT_ID`/`AZURE_CLIENT_SECRET` configuration activates `EnvironmentCredential` before managed identity.

## 5. Troubleshooting failures

1. **Enable Azure SDK logging.**

   ```powershell
   $env:AZURE_LOG_LEVEL = "verbose"
   ```

   Ensure the application has an SLF4J-compatible logging implementation. For temporary diagnosis, you can also use:

   ```java
   new DefaultAzureCredentialBuilder()
       .enableAccountIdentifierLogging()
       .build();
   ```

   Account identifier logging can expose identity-related information, so use it cautiously.

2. **Read the complete exception.** It lists each attempted credential and usually includes Microsoft Entra error codes, correlation IDs, and timestamps.

3. **Check local login state.**

   ```powershell
   az account show
   ```

   Confirm the expected tenant and subscription are active. If the wrong cached account succeeds first, sign out of that tool or restrict `AZURE_TOKEN_CREDENTIALS` to the desired credential.

4. **Check environment configuration.** For a service principal using a secret, all three are required:

   ```text
   AZURE_TENANT_ID
   AZURE_CLIENT_ID
   AZURE_CLIENT_SECRET
   ```

   A partially configured environment credential can produce confusing diagnostics.

5. **Distinguish authentication from authorization.**
   - `CredentialUnavailableException`: a credential cannot run in this environment.
   - `ClientAuthenticationException`: token acquisition failed, such as an invalid tenant, secret, or federation configuration.
   - HTTP `401`: token, tenant, audience, or authentication problem.
   - HTTP `403`: authentication probably succeeded, but the identity lacks Key Vault data-plane access or is blocked by vault networking.

6. **Check the resource configuration.**
   - Vault URL must be `https://<vault-name>.vault.azure.net`.
   - The authenticated identity needs an appropriate Key Vault role/access policy.
   - Key Vault firewall, private endpoint, DNS, and proxy configuration must allow the client.
   - For managed identity, verify that the requested identity is assigned to the Azure host.

## Complete example

`pom.xml`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0
                             https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>example</groupId>
    <artifactId>default-azure-credential-example</artifactId>
    <version>1.0.0</version>

    <properties>
        <maven.compiler.release>17</maven.compiler.release>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    </properties>

    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-identity</artifactId>
            <version>1.18.5</version>
        </dependency>

        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-security-keyvault-secrets</artifactId>
            <version>4.11.0</version>
        </dependency>
    </dependencies>
</project>
```

`src/main/java/example/KeyVaultExample.java`:

```java
package example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultExample {
    private KeyVaultExample() {
    }

    public static void main(String[] args) {
        String vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");
        String secretName = args.length > 0 ? args[0] : "example-secret";

        DefaultAzureCredential credential =
            new DefaultAzureCredentialBuilder().build();

        SecretClient secretClient =
            new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential)
                .buildClient();

        try {
            // The first service request triggers token acquisition.
            KeyVaultSecret secret = secretClient.getSecret(secretName);

            // Avoid printing the secret value to logs.
            System.out.printf(
                "Retrieved secret '%s', version '%s'.%n",
                secret.getName(),
                secret.getProperties().getVersion());
        } catch (CredentialUnavailableException exception) {
            System.err.println(
                "No credential in DefaultAzureCredential was available: "
                    + exception.getMessage());
        } catch (ClientAuthenticationException exception) {
            System.err.println(
                "Microsoft Entra authentication failed: "
                    + exception.getMessage());
        } catch (HttpResponseException exception) {
            System.err.printf(
                "Key Vault request failed with HTTP %d: %s%n",
                exception.getResponse().getStatusCode(),
                exception.getMessage());
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }
}
```

Run it after authenticating and setting the vault URL:

```powershell
az login
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
mvn compile exec:java `
  "-Dexec.mainClass=example.KeyVaultExample" `
  "-Dexec.args=<secret-name>"
```

References: [Java credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), [Azure Identity Java library](https://learn.microsoft.com/java/api/overview/azure/identity-readme), [Key Vault Secrets Java library](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme), and [authentication troubleshooting](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview).

## Reviewer Notes (LLM-as-Judge)

### Summary

The response satisfies the sole criterion because it does not instantiate any AutoCloseable Azure SDK clients.

### Strengths

- Uses SecretClient appropriately without unnecessary resource-management code.

## Grader Results

- default-azure-credential.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Maven dependency for `azure-identity`: Pass
      - `DefaultAzureCredentialBuilder` pattern: Pass
      - Credential chain order in Java SDK: Pass
      - Passing credential to client builders (e.g., `SecretClientBuilder`): Pass
      - Logging configuration for authentication troubleshooting: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Pass
  - Client Builder Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Pass
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Pass
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Pass
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.0000** | **92.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-default-credential --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

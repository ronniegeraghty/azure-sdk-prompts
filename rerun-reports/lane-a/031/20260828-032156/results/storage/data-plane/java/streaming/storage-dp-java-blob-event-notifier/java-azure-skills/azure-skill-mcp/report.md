# Evaluation Report: storage-dp-java-blob-event-notifier

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 893.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-blob-event-notifier` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 16/22 |
| Duration | 893.0s |
| Timestamp | 2026-08-27T19:21:56Z |
| Files Generated | 14 |
| Event Count | 17709 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 589.9s |
| Review | 302.3s |
| **Total** | **893.0s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=692924 out=12123 |
| Turn Count | 13 |

## Tool Calls

### ✅ 1. skill — 2482ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 193975ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Blob Storage and Event Grid SDKs with DefaultAzureCredential managed identity, synchronous and asynchronous clients, without deploying Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. view — 193972ms

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

### ✅ 4. view — 192549ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk\\azure-storage-blob-java.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md
@@ -1,37 +1,37 @@
 # Blob Storage — Java SDK Quick Reference
 
 > Condensed from **azure-storage-blob-java**. Full patterns (SAS tokens,
 > streaming, lease management, parallel uploads, proxy config)
 > in the **azure-storage-blob-java** plugin skill if installed.
 
 ## Install
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-storage-blob</artifactId>
     <version>12.33.0</version>
 </dependency>
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
 </dependency>
 ```
 
 ## Quick Start
 ```java
 import com.azure.storage.blob.BlobServiceClientBuilder;
 import com.azure.identity.DefaultAzureCredentialBuilder;
 var serviceClient = new BlobServiceClientBuilder()
     .endpoint("<storage-account-url>")
     .credential(new DefaultAzureCredentialBuilder().build())
     .buildClient();
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** — in production, use ManagedIdentityCredential. See [auth-best-practices.md](../auth-best-practices.md)
 - Use `BinaryData.fromString()` for string uploads
 - Use `createIfNotExists()` for idempotent container creation
 - Use `BlobParallelUploadOptions` for large file uploads with headers/metadata
 - Use `BlobInputStream`/`BlobOutputStream` for streaming 
... (truncated)
```

### ✅ 5. glob — 191216ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. glob — 190002ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 185676ms

**Input:**

```json
{"command":"get_best_practices","intent":"Need Java Azure SDK code-generation best practices for managed-identity Blob Storage access and Event Grid publishing in a local-only Maven demo. No resources or deployments.","learn":false,"parameters":{"action":"code-generation","language":"java","resource":"general"}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. web_fetch — 185670ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-messaging-eventgrid/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/messaging-eventgrid-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Event Grid client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/messaging-eventgrid-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-01-29T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-messaging-eventgrid, azure-event-grid
ms.service: azure-event-grid
locale: en-us
document_id: 688a6327-61e5-55ed-bb09-3301c8941cd5
document_version_independent_id: 0ec271f3-07c3-8cb9-480e-e46700721546
updated_at: 2026-01-30T12:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/messaging-eventgrid-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/d986a215fb382f879424164aba89c1642e21c742/docs-ref-services/latest/messaging-eventgrid-readme.md
git_commit_id: d986a215fb382f879424164aba89c1642e21c742
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 2655
asset_id: api/overview/azure/messaging-eventgrid-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/messaging-eventgrid-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty
... (truncated)
```

### ✅ 9. web_fetch — 185326ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-messaging-eventgrid/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-messaging-eventgrid</artifactId>
  <versioning>
    <latest>4.31.8</latest>
    <release>4.31.8</release>
    <versions>
      <version>2.0.0-beta.1</version>
      <version>2.0.0-beta.2</version>
      <version>2.0.0-beta.3</version>
      <version>2.0.0-beta.4</version>
      <version>4.0.0</version>
      <version>4.1.0</version>
      <version>4.2.0</version>
      <version>4.3.0</version>
      <version>4.4.0</version>
      <version>4.5.0</version>
      <version>4.6.0</version>
      <version>4.6.1</version>
      <version>4.7.0</version>
      <version>4.7.1</version>
      <version>4.8.0</version>
      <version>4.9.0</version>
      <version>4.10.0</version>
      <version>4.11.0</version>
      <version>4.11.1</version>
      <version>4.11.2</version>
      <version>4.12.0-beta.1</version>
      <version>4.12.0-beta.2</version>
      <version>4.12.0</version>
      <version>4.12.1</version>
      <version>4.12.2</version>
      <version>4.12.3</version>
      <version>4.13.0</version>
      <version>4.13.1</version>
      <version>4.14.0</version>
      <version>4.14.1</version>
      <version>4.14.2</version>
      <version>4.15.0</version>
      <version>4.15.1</version>
      <version>4.16.0</version>
      <version>4.17.0-beta.1</version>
      <version>4.17.0</version>
      <version>4.17.1</version>
      <version>4.17.2</version>
      <version>4.18.0</version>
      <version>4.19.0</version>
      <version>4.20.0</version>
      <version>4.20.1</version>
      <version>4.21.0-beta.1</version>
      <version>4.21.0</version>
      <version>4.22.0</version>
      <version>4.22.1</version>
      <version>4.22.2</version>
      <version>4.22.3</version>
      <version>4.22.4<
... (truncated)
```

### ✅ 10. web_fetch — 185007ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-storage-blob</artifactId>
  <versioning>
    <latest>12.36.0-beta.1</latest>
    <release>12.36.0-beta.1</release>
    <versions>
      <version>12.7.0</version>
      <version>12.8.0-beta.1</version>
      <version>12.8.0</version>
      <version>12.9.0-beta.1</version>
      <version>12.9.0-beta.2</version>
      <version>12.9.0</version>
      <version>12.10.0-beta.1</version>
      <version>12.10.0</version>
      <version>12.10.1</version>
      <version>12.10.2</version>
      <version>12.11.0-beta.1</version>
      <version>12.11.0-beta.2</version>
      <version>12.11.0-beta.3</version>
      <version>12.11.0</version>
      <version>12.11.1</version>
      <version>12.12.0-beta.1</version>
      <version>12.12.0</version>
      <version>12.13.0</version>
      <version>12.14.0-beta.1</version>
      <version>12.14.0</version>
      <version>12.14.1</version>
      <version>12.14.2</version>
      <version>12.14.3</version>
      <version>12.14.4</version>
      <version>12.15.0-beta.1</version>
      <version>12.15.0-beta.2</version>
      <version>12.15.0-beta.3</version>
      <version>12.15.0</version>
      <version>12.16.0</version>
      <version>12.16.1</version>
      <version>12.17.0-beta.1</version>
      <version>12.17.0</version>
      <version>12.17.1</version>
      <version>12.18.0-beta.1</version>
      <version>12.18.0</version>
      <version>12.19.0</version>
      <version>12.19.1</version>
      <version>12.20.0-beta.1</version>
      <version>12.20.0</version>
      <version>12.20.1</version>
      <version>12.20.2</version>
      <version>12.20.3</version>
      <version>12.21.0-beta.1</version>
      <version>12.21.0</version>
      <version>12.21.1</version>
     
... (truncated)
```

### ✅ 11. web_fetch — 184683ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/messaging-eventgrid-readme?view=azure-java-stable"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <versioning>
    <latest>1.19.0-beta.2</latest>
    <release>1.19.0-beta.2</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.0-preview.1</version>
      <version>1.0.0-preview.2</version>
      <version>1.0.0-preview.3</version>
      <version>1.0.0-preview.4</version>
      <version>1.0.0-preview.5</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.0.7</version>
      <version>1.0.8</version>
      <version>1.0.9</version>
      <version>1.1.0-beta.1</version>
      <version>1.1.0-beta.2</version>
      <version>1.1.0-beta.3</version>
      <version>1.1.0-beta.4</version>
      <version>1.1.0-beta.5</version>
      <version>1.1.0-beta.6</version>
      <version>1.1.0-beta.7</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.1.2</version>
      <version>1.1.3</version>
      <version>1.2.0-beta.1</version>
      <version>1.2.0-beta.2</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.3.0-beta.1</version>
      <version>1.3.0-beta.2</version>
      <version>1.3.0</version>
      <version>1.3.1</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version>1.3.4</version>
      <version>1.3.5</version>
      <version>1.3.6</version>
      <version>1.3.7</version>
      <version>1.4.0-beta.1</version>
      <version>1.4.0</version>
      <version>1.4.1</version>
      <version>
... (truncated)
```

### ✅ 12. apply_patch — 707ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eblob-event-notifier\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0-SNAPSHOT\u003c/version\u003e\n+    \u003cname\u003eAzure Blob Event Notifier\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.eventgrid.version\u003e4.31.8\u003c/azure.eventgrid.version\u003e\n+        \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.35.1\u003c/azure.storage.blob.version\u003e\n+        \u003cjackson.version\u003e2.20.0\u003c/jackson.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventgrid\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.eventgrid.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.blobevents.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/blobevents/BlobLifecycleEvent.java\n+package com.example.blobevents;\n+\n+import com.fasterxml.jackson.databind.JsonNode;\n+import java.time.OffsetDateTime;\n+\n+public record BlobLifecycleEvent(\n+    String id,\n+    String eventType,\n+    String subject,\n+    OffsetDateTime eventTime,\n+    JsonNode data,\n+    EventSchema schema\n+) {\n+    public enum EventSchema {\n+        EVENT_GRID,\n+        CLOUD_EVENTS_1_0\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPayloadParser.java\n+package com.example.blobevents;\n+\n+import com.example.blobevents.BlobLifecycleEvent.EventSchema;\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.JsonNode;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import com.fasterxml.jackson.databind.node.ArrayNode;\n+import java.time.OffsetDateTime;\n+import java.util.ArrayList;\n+import java.util.List;\n+\n+public final class EventPayloadParser {\n+    private final ObjectMapper objectMapper;\n+\n+    public EventPayloadParser() {\n+        this(new ObjectMapper());\n+    }\n+\n+    EventPayloadParser(ObjectMapper objectMapper) {\n+        this.objectMapper = objectMapper;\n+    }\n+\n+    public List\u003cBlobLifecycleEvent\u003e parse(String payload) {\n+        if (payload == null || payload.isBlank()) {\n+            throw new IllegalArgumentException(\"Event payload must not be blank\");\n+        }\n+\n+        try {\n+            JsonNode root = objectMapper.readTree(payload);\n+            List\u003cBlobLifecycleEvent\u003e events = new ArrayList\u003c\u003e();\n+            if (root.isArray()) {\n+                for (JsonNode node : (ArrayNode) root) {\n+                    events.add(parseEvent(node));\n+                }\n+            } else if (root.isObject()) {\n+                events.add(parseEvent(root));\n+            } else {\n+                throw new IllegalArgumentException(\"Event payload must be a JSON object or array\");\n+            }\n+            return List.copyOf(events);\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Event payload is not valid JSON\", exception);\n+        }\n+    }\n+\n+    private BlobLifecycleEvent parseEvent(JsonNode node) {\n+        if (node.hasNonNull(\"specversion\")) {\n+            String specVersion = requiredText(node, \"specversion\");\n+            if (!\"1.0\".equals(specVersion)) {\n+                throw new IllegalArgumentException(\"Unsupported CloudEvents specversion: \" + specVersion);\n+            }\n+            return new BlobLifecycleEvent(\n+                requiredText(node, \"id\"),\n+                requiredText(node, \"type\"),\n+                requiredText(node, \"subject\"),\n+                optionalTime(node, \"time\"),\n+                node.path(\"data\"),\n+                EventSchema.CLOUD_EVENTS_1_0\n+            );\n+        }\n+\n+        if (node.hasNonNull(\"eventType\")) {\n+            return new BlobLifecycleEvent(\n+                requiredText(node, \"id\"),\n+                requiredText(node, \"eventType\"),\n+                requiredText(node, \"subject\"),\n+                optionalTime(node, \"eventTime\"),\n+                node.path(\"data\"),\n+                EventSchema.EVENT_GRID\n+            );\n+        }\n+\n+        throw new IllegalArgumentException(\"Event is neither Event Grid schema nor CloudEvents 1.0 schema\");\n+    }\n+\n+    private static String requiredText(JsonNode node, String field) {\n+        JsonNode value = node.get(field);\n+        if (value == null || !value.isTextual() || value.textValue().isBlank()) {\n+            throw new IllegalArgumentException(\"Event field '\" + field + \"' must be a non-empty string\");\n+        }\n+        return value.textValue();\n+    }\n+\n+    private static OffsetDateTime optionalTime(JsonNode node, String field) {\n+        JsonNode value = node.get(field);\n+        return value == null || value.isNull() ? null : OffsetDateTime.parse(requiredText(node, field));\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobEventHandler.java\n+package com.example.blobevents;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.models.BlobProperties;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import java.net.URLDecoder;\n+import java.nio.charset.StandardCharsets;\n+import java.util.Map;\n+import java.util.Objects;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+public final class BlobEventHandler {\n+    public static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    public static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);\n+    private static final String CONTAINER_MARKER = \"/containers/\";\n+    private static final String BLOB_MARKER = \"/blobs/\";\n+\n+    private final BlobServiceClient blobServiceClient;\n+    private final BlobServiceAsyncClient blobServiceAsyncClient;\n+    private final Map\u003cBlobAddress, BlobSummary\u003e demoBlobs;\n+\n+    public BlobEventHandler(\n+        BlobServiceClient blobServiceClient,\n+        BlobServiceAsyncClient blobServiceAsyncClient\n+    ) {\n+        this.blobServiceClient = Objects.requireNonNull(blobServiceClient, \"blobServiceClient\");\n+        this.blobServiceAsyncClient = Objects.requireNonNull(blobServiceAsyncClient, \"blobServiceAsyncClient\");\n+        this.demoBlobs = null;\n+    }\n+\n+    public static BlobEventHandler forDemo(Map\u003cBlobAddress, BlobSummary\u003e demoBlobs) {\n+        return new BlobEventHandler(Map.copyOf(demoBlobs));\n+    }\n+\n+    private BlobEventHandler(Map\u003cBlobAddress, BlobSummary\u003e demoBlobs) {\n+        this.blobServiceClient = null;\n+        this.blobServiceAsyncClient = null;\n+        this.demoBlobs = demoBlobs;\n+    }\n+\n+    public void handle(BlobLifecycleEvent event) {\n+        switch (event.eventType()) {\n+            case BLOB_CREATED -\u003e handleCreated(event);\n+            case BLOB_DELETED -\u003e logDeletion(event);\n+            default -\u003e LOGGER.warn(\"Ignoring unrecognized event type '{}' for event {}\", event.eventType(), event.id());\n+        }\n+    }\n+\n+    public Mono\u003cVoid\u003e handleAsync(BlobLifecycleEvent event) {\n+        return switch (event.eventType()) {\n+            case BLOB_CREATED -\u003e handleCreatedAsync(event);\n+            case BLOB_DELETED -\u003e Mono.fromRunnable(() -\u003e logDeletion(event));\n+            default -\u003e Mono.fromRunnable(\n+                () -\u003e LOGGER.warn(\"Ignoring unrecognized event type '{}' for event {}\", event.eventType(), event.id())\n+            );\n+        };\n+    }\n+\n+    private void handleCreated(BlobLifecycleEvent event) {\n+        BlobAddress address = parseAddress(event.subject());\n+        if (demoBlobs != null) {\n+            printSummary(address, requireDemoBlob(address));\n+            return;\n+        }\n+\n+        BlobClient blobClient = blobServiceClient\n+            .getBlobContainerClient(address.container())\n+            .getBlobClient(address.blobName());\n+        try {\n+            BlobProperties properties = blobClient.getProperties();\n+            BinaryData ignoredContent = blobClient.downloadContent();\n+            printSummary(address, fromProperties(properties));\n+        } catch (BlobStorageException exception) {\n+            handleStorageRace(address, exception);\n+        }\n+    }\n+\n+    private Mono\u003cVoid\u003e handleCreatedAsync(BlobLifecycleEvent event) {\n+        BlobAddress address = parseAddress(event.subject());\n+        if (demoBlobs != null) {\n+            return Mono.fromRunnable(() -\u003e printSummary(address, requireDemoBlob(address)));\n+        }\n+\n+        return blobServiceAsyncClient\n+            .getBlobContainerAsyncClient(address.container())\n+            .getBlobAsyncClient(address.blobName())\n+            .getProperties()\n+            .flatMap(properties -\u003e blobServiceAsyncClient\n+                .getBlobContainerAsyncClient(address.container())\n+                .getBlobAsyncClient(address.blobName())\n+                .downloadContent()\n+                .doOnNext(ignored -\u003e printSummary(address, fromProperties(properties))))\n+            .onErrorResume(BlobStorageException.class, exception -\u003e {\n+                handleStorageRace(address, exception);\n+                return Mono.empty();\n+            })\n+            .then();\n+    }\n+\n+    private static void logDeletion(BlobLifecycleEvent event) {\n+        BlobAddress address = parseAddress(event.subject());\n+        LOGGER.info(\"Blob deleted: container='{}', name='{}'\", address.container(), address.blobName());\n+    }\n+\n+    private static BlobSummary fromProperties(BlobProperties properties) {\n+        String contentType = properties.getContentType() == null ? \"application/octet-stream\" : properties.getContentType();\n+        String accessTier = properties.getAccessTier() == null ? \"unknown\" : properties.getAccessTier().toString();\n+        return new BlobSummary(properties.getBlobSize(), contentType, accessTier);\n+    }\n+\n+    private static void printSummary(BlobAddress address, BlobSummary summary) {\n+        LOGGER.info(\n+            \"Blob created: name='{}', container='{}', size={} bytes, contentType='{}', accessTier='{}'\",\n+            address.blobName(),\n+            address.container(),\n+            summary.size(),\n+            summary.contentType(),\n+            summary.accessTier()\n+        );\n+    }\n+\n+    private BlobSummary requireDemoBlob(BlobAddress address) {\n+        BlobSummary summary = demoBlobs.get(address);\n+        if (summary == null) {\n+            LOGGER.warn(\"Blob disappeared before it could be downloaded: {}/{}\", address.container(), address.blobName());\n+            return new BlobSummary(0, \"unknown\", \"unknown\");\n+        }\n+        return summary;\n+    }\n+\n+    private static void handleStorageRace(BlobAddress address, BlobStorageException exception) {\n+        if (exception.getStatusCode() == 404) {\n+            LOGGER.warn(\"Blob disappeared before it could be downloaded: {}/{}\", address.container(), address.blobName());\n+            return;\n+        }\n+        if (exception.getStatusCode() == 409) {\n+            LOGGER.warn(\n+                \"Blob is currently unavailable, possibly because its access tier changed: {}/{} ({})\",\n+                address.container(),\n+                address.blobName(),\n+                exception.getErrorCode()\n+            );\n+            return;\n+        }\n+        throw exception;\n+    }\n+\n+    public static BlobAddress parseAddress(String subject) {\n+        int containerStart = subject.indexOf(CONTAINER_MARKER);\n+        int blobStart = subject.indexOf(BLOB_MARKER, containerStart + CONTAINER_MARKER.length());\n+        if (containerStart \u003c 0 || blobStart \u003c 0) {\n+            throw new IllegalArgumentException(\"Blob event subject has an unexpected format: \" + subject);\n+        }\n+        String container = subject.substring(containerStart + CONTAINER_MARKER.length(), blobStart);\n+        String blobName = URLDecoder.decode(subject.substring(blobStart + BLOB_MARKER.length()), StandardCharsets.UTF_8);\n+        if (container.isBlank() || blobName.isBlank()) {\n+            throw new IllegalArgumentException(\"Blob event subject must identify a container and blob: \" + subject);\n+        }\n+        return new BlobAddress(container, blobName);\n+    }\n+\n+    public record BlobAddress(String container, String blobName) {\n+    }\n+\n+    public record BlobSummary(long size, String contentType, String accessTier) {\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventReceiver.java\n+package com.example.blobevents;\n+\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class EventReceiver {\n+    private final EventPayloadParser parser;\n+    private final BlobEventHandler handler;\n+\n+    public EventReceiver(EventPayloadParser parser, BlobEventHandler handler) {\n+        this.parser = Objects.requireNonNull(parser, \"parser\");\n+        this.handler = Objects.requireNonNull(handler, \"handler\");\n+    }\n+\n+    public List\u003cBlobLifecycleEvent\u003e receive(String jsonPayload) {\n+        List\u003cBlobLifecycleEvent\u003e events = parser.parse(jsonPayload);\n+        events.forEach(handler::handle);\n+        return events;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventReceiverAsync.java\n+package com.example.blobevents;\n+\n+import java.util.List;\n+import java.util.Objects;\n+import reactor.core.publisher.Flux;\n+\n+public final class EventReceiverAsync {\n+    private final EventPayloadParser parser;\n+    private final BlobEventHandler handler;\n+\n+    public EventReceiverAsync(EventPayloadParser parser, BlobEventHandler handler) {\n+        this.parser = Objects.requireNonNull(parser, \"parser\");\n+        this.handler = Objects.requireNonNull(handler, \"handler\");\n+    }\n+\n+    public Flux\u003cBlobLifecycleEvent\u003e receive(String jsonPayload) {\n+        List\u003cBlobLifecycleEvent\u003e events = parser.parse(jsonPayload);\n+        return Flux.fromIterable(events)\n+            .concatMap(event -\u003e handler.handleAsync(event).thenReturn(event));\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/CustomEvent.java\n+package com.example.blobevents;\n+\n+import java.util.Objects;\n+\n+public record CustomEvent(String subject, String eventType, String dataVersion, Object data) {\n+    public CustomEvent {\n+        Objects.requireNonNull(subject, \"subject\");\n+        Objects.requireNonNull(eventType, \"eventType\");\n+        Objects.requireNonNull(dataVersion, \"dataVersion\");\n+        Objects.requireNonNull(data, \"data\");\n+        if (!subject.startsWith(\"/\")) {\n+            throw new IllegalArgumentException(\"subject must be an absolute hierarchy beginning with '/'\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import java.util.List;\n+import java.util.Objects;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+public final class EventPublisher {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(EventPublisher.class);\n+\n+    private final EventGridPublisherClient\u003cEventGridEvent\u003e client;\n+\n+    public EventPublisher(EventGridPublisherClient\u003cEventGridEvent\u003e client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public static EventPublisher dryRun() {\n+        return new EventPublisher();\n+    }\n+\n+    private EventPublisher() {\n+        this.client = null;\n+    }\n+\n+    public void publish(List\u003cCustomEvent\u003e events) {\n+        List\u003cEventGridEvent\u003e sdkEvents = toSdkEvents(events);\n+        if (client == null) {\n+            sdkEvents.forEach(event -\u003e LOGGER.info(\n+                \"Dry-run publish: type='{}', subject='{}'\",\n+                event.getEventType(),\n+                event.getSubject()\n+            ));\n+            return;\n+        }\n+        client.sendEvents(sdkEvents);\n+    }\n+\n+    static List\u003cEventGridEvent\u003e toSdkEvents(List\u003cCustomEvent\u003e events) {\n+        return events.stream()\n+            .map(event -\u003e new EventGridEvent(\n+                event.subject(),\n+                event.eventType(),\n+                event.data(),\n+                event.dataVersion()\n+            ))\n+            .toList();\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPublisherAsync.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import java.util.List;\n+import java.util.Objects;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+public final class EventPublisherAsync {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(EventPublisherAsync.class);\n+\n+    private final EventGridPublisherAsyncClient\u003cEventGridEvent\u003e client;\n+\n+    public EventPublisherAsync(EventGridPublisherAsyncClient\u003cEventGridEvent\u003e client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public static EventPublisherAsync dryRun() {\n+        return new EventPublisherAsync();\n+    }\n+\n+    private EventPublisherAsync() {\n+        this.client = null;\n+    }\n+\n+    public Mono\u003cVoid\u003e publish(List\u003cCustomEvent\u003e events) {\n+        List\u003cEventGridEvent\u003e sdkEvents = EventPublisher.toSdkEvents(events);\n+        if (client == null) {\n+            return Mono.fromRunnable(() -\u003e sdkEvents.forEach(event -\u003e LOGGER.info(\n+                \"Dry-run async publish: type='{}', subject='{}'\",\n+                event.getEventType(),\n+                event.getSubject()\n+            )));\n+        }\n+        return client.sendEvents(sdkEvents);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureConfiguration.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+import java.util.Objects;\n+\n+public final class AzureConfiguration {\n+    private final String storageEndpoint;\n+    private final String eventGridTopicEndpoint;\n+    private final TokenCredential credential;\n+\n+    public AzureConfiguration(\n+        String storageEndpoint,\n+        String eventGridTopicEndpoint,\n+        String managedIdentityClientId\n+    ) {\n+        this.storageEndpoint = requireHttpsEndpoint(storageEndpoint, \"storageEndpoint\");\n+        this.eventGridTopicEndpoint = requireHttpsEndpoint(eventGridTopicEndpoint, \"eventGridTopicEndpoint\");\n+        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            builder.clientId(managedIdentityClientId);\n+        }\n+        this.credential = builder.build();\n+    }\n+\n+    public BlobServiceClient blobServiceClient() {\n+        return new BlobServiceClientBuilder()\n+            .endpoint(storageEndpoint)\n+            .credential(credential)\n+            .buildClient();\n+    }\n+\n+    public BlobServiceAsyncClient blobServiceAsyncClient() {\n+        return new BlobServiceClientBuilder()\n+            .endpoint(storageEndpoint)\n+            .credential(credential)\n+            .buildAsyncClient();\n+    }\n+\n+    public EventGridPublisherClient\u003cEventGridEvent\u003e eventGridPublisherClient() {\n+        return new EventGridPublisherClientBuilder()\n+            .endpoint(eventGridTopicEndpoint)\n+            .credential(credential)\n+            .buildEventGridEventPublisherClient();\n+    }\n+\n+    public EventGridPublisherAsyncClient\u003cEventGridEvent\u003e eventGridPublisherAsyncClient() {\n+        return new EventGridPublisherClientBuilder()\n+            .endpoint(eventGridTopicEndpoint)\n+            .credential(credential)\n+            .buildEventGridEventPublisherAsyncClient();\n+    }\n+\n+    private static String requireHttpsEndpoint(String endpoint, String name) {\n+        Objects.requireNonNull(endpoint, name);\n+        if (!endpoint.startsWith(\"https://\")) {\n+            throw new IllegalArgumentException(name + \" must use HTTPS\");\n+        }\n+        return endpoint;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/Main.java\n+package com.example.blobevents;\n+\n+import com.example.blobevents.BlobEventHandler.BlobAddress;\n+import com.example.blobevents.BlobEventHandler.BlobSummary;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class Main {\n+    private static final String EVENT_GRID_PAYLOAD = \"\"\"\n+        [\n+          {\n+            \"id\": \"eg-created-001\",\n+            \"topic\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/invoices%2Finvoice-1001.pdf\",\n+            \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+            \"eventTime\": \"2026-08-27T19:00:00Z\",\n+            \"data\": {\n+              \"api\": \"PutBlob\",\n+              \"contentType\": \"application/pdf\",\n+              \"contentLength\": 24576,\n+              \"url\": \"https://demostore.blob.core.windows.net/documents/invoices/invoice-1001.pdf\"\n+            },\n+            \"dataVersion\": \"3\",\n+            \"metadataVersion\": \"1\"\n+          },\n+          {\n+            \"id\": \"eg-deleted-001\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/archive%2Finvoice-0999.pdf\",\n+            \"eventType\": \"Microsoft.Storage.BlobDeleted\",\n+            \"eventTime\": \"2026-08-27T19:01:00Z\",\n+            \"data\": {\n+              \"api\": \"DeleteBlob\",\n+              \"url\": \"https://demostore.blob.core.windows.net/documents/archive/invoice-0999.pdf\"\n+            },\n+            \"dataVersion\": \"3\",\n+            \"metadataVersion\": \"1\"\n+          }\n+        ]\n+        \"\"\";\n+\n+    private static final String CLOUD_EVENTS_PAYLOAD = \"\"\"\n+        [\n+          {\n+            \"specversion\": \"1.0\",\n+            \"id\": \"ce-created-001\",\n+            \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore\",\n+            \"type\": \"Microsoft.Storage.BlobCreated\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/reports%2Fquarterly.txt\",\n+            \"time\": \"2026-08-27T19:02:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+              \"api\": \"PutBlob\",\n+              \"contentType\": \"text/plain\",\n+              \"contentLength\": 1024,\n+              \"url\": \"https://demostore.blob.core.windows.net/documents/reports/quarterly.txt\"\n+            }\n+          },\n+          {\n+            \"specversion\": \"1.0\",\n+            \"id\": \"ce-deleted-001\",\n+            \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore\",\n+            \"type\": \"Microsoft.Storage.BlobDeleted\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/reports%2Fold-quarterly.txt\",\n+            \"time\": \"2026-08-27T19:03:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+              \"api\": \"DeleteBlob\",\n+              \"url\": \"https://demostore.blob.core.windows.net/documents/reports/old-quarterly.txt\"\n+            }\n+          }\n+        ]\n+        \"\"\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        BlobEventHandler handler = BlobEventHandler.forDemo(Map.of(\n+            new BlobAddress(\"documents\", \"invoices/invoice-1001.pdf\"),\n+            new BlobSummary(24_576, \"application/pdf\", \"HOT\"),\n+            new BlobAddress(\"documents\", \"reports/quarterly.txt\"),\n+            new BlobSummary(1_024, \"text/plain\", \"COOL\")\n+        ));\n+        EventPayloadParser parser = new EventPayloadParser();\n+        List\u003cCustomEvent\u003e downstreamEvents = List.of(new CustomEvent(\n+            \"/documents/invoices/processed\",\n+            \"Contoso.Documents.Processed\",\n+            \"1.0\",\n+            Map.of(\"documentId\", \"invoice-1001\", \"status\", \"processed\")\n+        ));\n+\n+        System.out.println(\"=== Synchronous implementation ===\");\n+        EventReceiver receiver = new EventReceiver(parser, handler);\n+        receiver.receive(EVENT_GRID_PAYLOAD);\n+        receiver.receive(CLOUD_EVENTS_PAYLOAD);\n+        EventPublisher.dryRun().publish(downstreamEvents);\n+\n+        System.out.println(\"=== Asynchronous implementation ===\");\n+        EventReceiverAsync asyncReceiver = new EventReceiverAsync(parser, handler);\n+        asyncReceiver.receive(EVENT_GRID_PAYLOAD).then().block();\n+        asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD).then().block();\n+        EventPublisherAsync.dryRun().publish(downstreamEvents).block();\n+    }\n+}\n*** Add File: README.md\n+# Azure Blob Event Notifier\n+\n+A Java 17 Maven sample that receives Azure Blob Storage lifecycle events in\n+Event Grid or CloudEvents 1.0 schema, downloads created blobs, logs deleted\n+blobs, and publishes downstream Event Grid events.\n+\n+The demo is local-only: it uses in-memory blob metadata and dry-run publishers.\n+\n+```powershell\n+mvn compile exec:java\n+```\n+\n+For Azure-hosted use, construct `AzureConfiguration` from environment-provided\n+HTTPS endpoints, then inject its sync or async clients into `BlobEventHandler`\n+and the matching publisher:\n+\n+```java\n+var config = new AzureConfiguration(\n+    System.getenv(\"AZURE_STORAGE_BLOB_ENDPOINT\"),\n+    System.getenv(\"AZURE_EVENT_GRID_TOPIC_ENDPOINT\"),\n+    System.getenv(\"AZURE_CLIENT_ID\")\n+);\n+```\n+\n+`AZURE_CLIENT_ID` is optional for a system-assigned identity and identifies a\n+user-assigned managed identity when set. Grant the identity **Storage Blob Data\n+Reader** on the storage account/container and **EventGrid Data Sender** on the\n+custom topic. No account keys, topic keys, connection strings, or SAS tokens\n+are used.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml
@@ -1,0 +1,69 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>blob-event-notifier</artifactId>
+    <version>1.0.0-SNAPSHOT</version>
+    <name>Azure Blob Event Notifier</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.eventgrid.version>4.31.8</azure.eventgrid.version>
+        <azure.identity.version>1.18.5</azure.identity.version>
+        <azure.storage.blob.version>12.35.1</azure.storage.blob.version>
+        <jackson.version>2.20.0</jackson.version>
+        <slf4j.version>2.0.17</slf4j.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-eventgrid</artifactId>
+            <version>${azure.eventgrid.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-storage-blob</artifactId>
+  
... (truncated)
```

### ✅ 13. powershell — 6729ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile exec:java","description":"Compile and run Java demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------< com.example:blob-event-notifier >-------------------
[INFO] Building Azure Blob Event Notifier 1.0.0-SNAPSHOT
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ blob-event-notifier ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ blob-event-notifier ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] Some messages have been simplified; recompile with -Xdiags:verbose to get full output
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/EventPublisher.java:[45,27] incompatible types: java.lang.Object cannot be converted to com.azure.core.util.BinaryData
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/EventPublisher.java:[48,20] incompatible types: java.util.List<j
... (truncated)
```

### ✅ 14. apply_patch — 742ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/blobevents/EventPublisher.java\n@@\n package com.example.blobevents;\n \n+import com.azure.core.util.BinaryData;\n import com.azure.messaging.eventgrid.EventGridEvent;\n@@\n                 event.subject(),\n                 event.eventType(),\n-                event.data(),\n+                BinaryData.fromObject(event.data()),\n                 event.dataVersion()\n             ))\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/EventPublisher.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/EventPublisher.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/EventPublisher.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/EventPublisher.java
@@ -1,5 +1,6 @@
 package com.example.blobevents;
 
+import com.azure.core.util.BinaryData;
 import com.azure.messaging.eventgrid.EventGridEvent;
 import com.azure.messaging.eventgrid.EventGridPublisherClient;
 import java.util.List;
@@ -42,7 +43,7 @@
             .map(event -> new EventGridEvent(
                 event.subject(),
                 event.eventType(),
-                event.data(),
+                BinaryData.fromObject(event.data()),
                 event.dataVersion()
             ))
             .toList();

```

### ✅ 15. powershell — 7185ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile exec:java","description":"Recompile and run Java demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------< com.example:blob-event-notifier >-------------------
[INFO] Building Azure Blob Event Notifier 1.0.0-SNAPSHOT
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ blob-event-notifier ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ blob-event-notifier ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- exec:3.5.1:java (default-cli) @ blob-event-notifier ---
=== Synchronous implementation ===
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob created: name='invoices/invoice-1001.pdf', container='documents', size=24576 bytes, contentType='application/pdf', accessTier='HOT'
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob deleted: container='documents', name='archive/invoice-0999.pdf'
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob created: name='reports/quarterly.txt', container='documents', size=1024 bytes, contentType='text/plain', accessTier='COOL'
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandle
... (truncated)
```

### ✅ 16. github-mcp-server-search_code (via github-mcp-server) — 3746ms

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"BLOB_ARCHIVED repo:Azure/azure-sdk-for-java path:sdk/storage/azure-storage-blob"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobErrorCode.java","path":"sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java","text_matches":[{"fragment":"     * Static value BlobArchived for BlobErrorCode.\n     */\n    @Generated\n    public static final BlobErrorCode BLOB_ARCHIVED = fromString(\"BlobArchived\");\n\n    /**\n     * Static value BlobNotArchived for BlobErrorCode.","matches":[{"indices":[113,126],"text":"BLOB_ARCHIVED"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"}]}],"total_count":1}
```

### ✅ 17. web_fetch — 3406ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/org/junit/junit-bom/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/org/junit/junit-bom/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>org.junit</groupId>
  <artifactId>junit-bom</artifactId>
  <versioning>
    <latest>6.1.3</latest>
    <release>6.1.3</release>
    <versions>
      <version>5.2.0-M1</version>
      <version>5.2.0-RC1</version>
      <version>5.2.0</version>
      <version>5.3.0-M1</version>
      <version>5.3.0-RC1</version>
      <version>5.3.0</version>
      <version>5.3.1</version>
      <version>5.3.2</version>
      <version>5.4.0-M1</version>
      <version>5.4.0-RC1</version>
      <version>5.4.0-RC2</version>
      <version>5.4.0</version>
      <version>5.4.1</version>
      <version>5.4.2</version>
      <version>5.5.0-M1</version>
      <version>5.5.0-RC1</version>
      <version>5.5.0-RC2</version>
      <version>5.5.0</version>
      <version>5.5.1</version>
      <version>5.5.2</version>
      <version>5.6.0-M1</version>
      <version>5.6.0-RC1</version>
      <version>5.6.0</version>
      <version>5.6.1</version>
      <version>5.6.2</version>
      <version>5.6.3</version>
      <version>5.7.0-M1</version>
      <version>5.7.0-RC1</version>
      <version>5.7.0</version>
      <version>5.7.1</version>
      <version>5.7.2</version>
      <version>5.8.0-M1</version>
      <version>5.8.0-RC1</version>
      <version>5.8.0</version>
      <version>5.8.1</version>
      <version>5.8.2</version>
      <version>5.9.0-M1</version>
      <version>5.9.0-RC1</version>
      <version>5.9.0</version>
      <version>5.9.1</version>
      <version>5.9.2</version>
      <version>5.9.3</version>
      <version>5.10.0-M1</version>
      <version>5.10.0-RC1</version>
      <version>5.10.0-RC2</version>
      <version>5.10.0</version>
      <version>5.10.1</version>
      <version>5.10.2</version>
      <version>5.10.3</version>
      <version>5.10.4</version>
... (truncated)
```

### ✅ 18. rg — 3059ms

**Input:**

```json
{"-n":true,"glob":"*.java","output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\\src","pattern":"URLDecoder|handleStorageRace|CLOUD_EVENTS_PAYLOAD|EVENT_GRID_PAYLOAD"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\BlobEventHandler.java:9:import java.net.URLDecoder;
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\BlobEventHandler.java:81:            handleStorageRace(address, exception);
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\BlobEventHandler.java:101:                handleStorageRace(address, exception);
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\BlobEventHandler.java:138:    private static void handleStorageRace(BlobAddress address, BlobStorageException exception) {
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\BlobEventHandler.java:162:        String blobName = URLDecoder.decode(subject.substring(blobStart + BLOB_MARKER.length()), StandardCharsets.UTF_8);
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\Main.java:9:    private static final String EVENT_GRID_PAYLOAD = """
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\Main.java:41:    private static final String CLOUD_EVENTS_PAYLOAD = """
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\java\com\example\blobevents\Main.java:94:        receiver.receive(EVENT_GRID_PAYLOAD);
C:\Users\weidxu\AppData\Local\Temp\hyoka-storage
... (truncated)
```

### ✅ 19. github-mcp-server-search_code (via github-mcp-server) — 3813ms

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"\"BlobBeingRehydrated\" repo:Azure/azure-sdk-for-java path:sdk/storage/azure-storage-blob"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobErrorCode.java","path":"sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java","text_matches":[{"fragment":"    /**\n     * Static value BlobNotArchived for BlobErrorCode.\n     */","matches":[{"indices":[28,43],"text":"BlobNotArchived"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"},{"fragment":"    @Generated\n    public static final BlobErrorCode BLOB_NOT_ARCHIVED = fromString(\"BlobNotArchived\");\n","matches":[{"indices":[85,100],"text":"BlobNotArchived"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"}]}],"total_count":1}
```

### ✅ 20. github-mcp-server-search_code (via github-mcp-server) — 3441ms

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"\"BlobNotArchived\" repo:Azure/azure-sdk-for-java path:sdk/storage/azure-storage-blob"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobErrorCode.java","path":"sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java","text_matches":[{"fragment":"    /**\n     * Static value BlobBeingRehydrated for BlobErrorCode.\n     */","matches":[{"indices":[28,47],"text":"BlobBeingRehydrated"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"},{"fragment":"    @Generated\n    public static final BlobErrorCode BLOB_BEING_REHYDRATED = fromString(\"BlobBeingRehydrated\");\n","matches":[{"indices":[89,108],"text":"BlobBeingRehydrated"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/models/BlobErrorCode.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"}]}],"total_count":1}
```

### ✅ 21. apply_patch — 647ms

**Input:**

```json
"*** Begin Patch\n*** Update File: pom.xml\n@@\n         \u003cjackson.version\u003e2.20.0\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.14.4\u003c/junit.version\u003e\n         \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n@@\n         \u003cdependency\u003e\n             \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n             \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n             \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n             \u003cscope\u003eruntime\u003c/scope\u003e\n         \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n@@\n             \u003cplugin\u003e\n                 \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n                 \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n@@\n                     \u003cmainClass\u003ecom.example.blobevents.Main\u003c/mainClass\u003e\n                 \u003c/configuration\u003e\n             \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.4\u003c/version\u003e\n+            \u003c/plugin\u003e\n         \u003c/plugins\u003e\n*** Update File: src/main/java/com/example/blobevents/BlobEventHandler.java\n@@\n import com.azure.storage.blob.BlobServiceAsyncClient;\n import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.models.BlobErrorCode;\n import com.azure.storage.blob.models.BlobProperties;\n import com.azure.storage.blob.models.BlobStorageException;\n-import java.net.URLDecoder;\n-import java.nio.charset.StandardCharsets;\n+import java.net.URI;\n@@\n-        if (exception.getStatusCode() == 409) {\n+        if (BlobErrorCode.BLOB_ARCHIVED.equals(exception.getErrorCode())\n+            || BlobErrorCode.BLOB_BEING_REHYDRATED.equals(exception.getErrorCode())) {\n             LOGGER.warn(\n                 \"Blob is currently unavailable, possibly because its access tier changed: {}/{} ({})\",\n@@\n         }\n         String container = subject.substring(containerStart + CONTAINER_MARKER.length(), blobStart);\n-        String blobName = URLDecoder.decode(subject.substring(blobStart + BLOB_MARKER.length()), StandardCharsets.UTF_8);\n+        String encodedBlobName = subject.substring(blobStart + BLOB_MARKER.length());\n+        String blobName = URI.create(encodedBlobName).getPath();\n*** Update File: src/main/java/com/example/blobevents/Main.java\n@@\n             \"dataVersion\": \"3\",\n             \"metadataVersion\": \"1\"\n+          },\n+          {\n+            \"id\": \"eg-unknown-001\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/invoices%2Finvoice-1001.pdf\",\n+            \"eventType\": \"Contoso.Storage.BlobReviewed\",\n+            \"eventTime\": \"2026-08-27T19:01:30Z\",\n+            \"data\": {\n+              \"review\": \"approved\"\n+            },\n+            \"dataVersion\": \"1\",\n+            \"metadataVersion\": \"1\"\n           }\n         ]\n*** Add File: src/test/java/com/example/blobevents/EventPayloadParserTest.java\n+package com.example.blobevents;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+import com.example.blobevents.BlobLifecycleEvent.EventSchema;\n+import java.util.List;\n+import org.junit.jupiter.api.Test;\n+\n+class EventPayloadParserTest {\n+    private final EventPayloadParser parser = new EventPayloadParser();\n+\n+    @Test\n+    void parsesEventGridAndCloudEventsSchemas() {\n+        List\u003cBlobLifecycleEvent\u003e eventGridEvents = parser.parse(\"\"\"\n+            [{\n+              \"id\": \"eg-1\",\n+              \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+              \"subject\": \"/blobServices/default/containers/docs/blobs/a.txt\",\n+              \"eventTime\": \"2026-08-27T19:00:00Z\",\n+              \"data\": {\"api\": \"PutBlob\"}\n+            }]\n+            \"\"\");\n+        List\u003cBlobLifecycleEvent\u003e cloudEvents = parser.parse(\"\"\"\n+            {\n+              \"specversion\": \"1.0\",\n+              \"id\": \"ce-1\",\n+              \"type\": \"Microsoft.Storage.BlobDeleted\",\n+              \"source\": \"/storageAccounts/demo\",\n+              \"subject\": \"/blobServices/default/containers/docs/blobs/a.txt\",\n+              \"time\": \"2026-08-27T19:01:00Z\",\n+              \"data\": {\"api\": \"DeleteBlob\"}\n+            }\n+            \"\"\");\n+\n+        assertEquals(EventSchema.EVENT_GRID, eventGridEvents.getFirst().schema());\n+        assertEquals(EventSchema.CLOUD_EVENTS_1_0, cloudEvents.getFirst().schema());\n+        assertEquals(BlobEventHandler.BLOB_CREATED, eventGridEvents.getFirst().eventType());\n+        assertEquals(BlobEventHandler.BLOB_DELETED, cloudEvents.getFirst().eventType());\n+    }\n+\n+    @Test\n+    void rejectsUnsupportedCloudEventsVersion() {\n+        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () -\u003e parser.parse(\"\"\"\n+            {\n+              \"specversion\": \"0.3\",\n+              \"id\": \"ce-1\",\n+              \"type\": \"Microsoft.Storage.BlobCreated\",\n+              \"subject\": \"/blobServices/default/containers/docs/blobs/a.txt\"\n+            }\n+            \"\"\"));\n+\n+        assertEquals(\"Unsupported CloudEvents specversion: 0.3\", exception.getMessage());\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/BlobEventHandlerTest.java\n+package com.example.blobevents;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+import com.example.blobevents.BlobEventHandler.BlobAddress;\n+import org.junit.jupiter.api.Test;\n+\n+class BlobEventHandlerTest {\n+    @Test\n+    void parsesEncodedNestedBlobNameWithoutTreatingPlusAsSpace() {\n+        BlobAddress address = BlobEventHandler.parseAddress(\n+            \"/blobServices/default/containers/documents/blobs/invoices%2F2026%2Ftotal+tax.pdf\"\n+        );\n+\n+        assertEquals(\"documents\", address.container());\n+        assertEquals(\"invoices/2026/total+tax.pdf\", address.blobName());\n+    }\n+\n+    @Test\n+    void rejectsMalformedSubject() {\n+        assertThrows(\n+            IllegalArgumentException.class,\n+            () -\u003e BlobEventHandler.parseAddress(\"/containers/documents\")\n+        );\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/pom.xml
@@ -16,6 +16,7 @@
         <azure.identity.version>1.18.5</azure.identity.version>
         <azure.storage.blob.version>12.35.1</azure.storage.blob.version>
         <jackson.version>2.20.0</jackson.version>
+        <junit.version>5.14.4</junit.version>
         <slf4j.version>2.0.17</slf4j.version>
     </properties>
 
@@ -46,6 +47,12 @@
             <version>${slf4j.version}</version>
             <scope>runtime</scope>
         </dependency>
+        <dependency>
+            <groupId>org.junit.jupiter</groupId>
+            <artifactId>junit-jupiter</artifactId>
+            <version>${junit.version}</version>
+            <scope>test</scope>
+        </dependency>
     </dependencies>
 
     <build>
@@ -63,6 +70,11 @@
                     <mainClass>com.example.blobevents.Main</mainClass>
                 </configuration>
             </plugin>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-surefire-plugin</artifactId>
+                <version>3.5.4</version>
+            </plugin>
         </plugins>
     </build>
 </project>


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/main/java/com/example/blobevents/BlobEventHandler.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-m
... (truncated)
```

### ✅ 22. apply_patch — 682ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/test/java/com/example/blobevents/EventPayloadParserTest.java\n@@\n-        assertEquals(EventSchema.EVENT_GRID, eventGridEvents.getFirst().schema());\n-        assertEquals(EventSchema.CLOUD_EVENTS_1_0, cloudEvents.getFirst().schema());\n-        assertEquals(BlobEventHandler.BLOB_CREATED, eventGridEvents.getFirst().eventType());\n-        assertEquals(BlobEventHandler.BLOB_DELETED, cloudEvents.getFirst().eventType());\n+        assertEquals(EventSchema.EVENT_GRID, eventGridEvents.get(0).schema());\n+        assertEquals(EventSchema.CLOUD_EVENTS_1_0, cloudEvents.get(0).schema());\n+        assertEquals(BlobEventHandler.BLOB_CREATED, eventGridEvents.get(0).eventType());\n+        assertEquals(BlobEventHandler.BLOB_DELETED, cloudEvents.get(0).eventType());\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/test/java/com/example/blobevents/EventPayloadParserTest.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/test/java/com/example/blobevents/EventPayloadParserTest.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/test/java/com/example/blobevents/EventPayloadParserTest.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324/src/test/java/com/example/blobevents/EventPayloadParserTest.java
@@ -33,10 +33,10 @@
             }
             """);
 
-        assertEquals(EventSchema.EVENT_GRID, eventGridEvents.getFirst().schema());
-        assertEquals(EventSchema.CLOUD_EVENTS_1_0, cloudEvents.getFirst().schema());
-        assertEquals(BlobEventHandler.BLOB_CREATED, eventGridEvents.getFirst().eventType());
-        assertEquals(BlobEventHandler.BLOB_DELETED, cloudEvents.getFirst().eventType());
+        assertEquals(EventSchema.EVENT_GRID, eventGridEvents.get(0).schema());
+        assertEquals(EventSchema.CLOUD_EVENTS_1_0, cloudEvents.get(0).schema());
+        assertEquals(BlobEventHandler.BLOB_CREATED, eventGridEvents.get(0).eventType());
+        assertEquals(BlobEventHandler.BLOB_DELETED, cloudEvents.get(0).eventType());
     }
 
     @Test

```

### ✅ 23. powershell — 14678ms

**Input:**

```json
{"command":"mvn --no-transfer-progress verify exec:java","description":"Test and run complete Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------< com.example:blob-event-notifier >-------------------
[INFO] Building Azure Blob Event Notifier 1.0.0-SNAPSHOT
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] Artifact org.junit:junit-bom:pom:5.14.4 is present in the local repository, but cached from a remote repository ID that is unavailable in current build context, verifying that is downloadable from [central (https://repo.maven.apache.org/maven2, default, releases)]
[INFO] Artifact org.junit:junit-bom:pom:5.14.4 is present in the local repository, but cached from a remote repository ID that is unavailable in current build context, verifying that is downloadable from [central (https://repo.maven.apache.org/maven2, default, releases)]
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ blob-event-notifier ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-326186324\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ blob-event-notifier ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ blob-event-notifier ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-blob-event-notifier-java-azure-sk
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\blobevents\AzureConfiguration.java`
- `src\main\java\com\example\blobevents\BlobEventHandler.java`
- `src\main\java\com\example\blobevents\BlobLifecycleEvent.java`
- `src\main\java\com\example\blobevents\CustomEvent.java`
- `src\main\java\com\example\blobevents\EventPayloadParser.java`
- `src\main\java\com\example\blobevents\EventPublisher.java`
- `src\main\java\com\example\blobevents\EventPublisherAsync.java`
- `src\main\java\com\example\blobevents\EventReceiver.java`
- `src\main\java\com\example\blobevents\EventReceiverAsync.java`
- `src\main\java\com\example\blobevents\Main.java`
- `src\test\java\com\example\blobevents\BlobEventHandlerTest.java`
- `src\test\java\com\example\blobevents\EventPayloadParserTest.java`

## Copilot Response

I’m grounding the project in the current Java Blob SDK and managed-identity guidance, then I’ll scaffold and compile the demo locally.

The storage guidance confirms deterministic `ManagedIdentityCredential` for Azure-hosted use. I’m resolving the Event Grid API surface and dependency versions before writing the source.

The project will use production managed-identity clients but keep `Main` fully offline with in-memory blob metadata and dry-run publishers, so the complete sync/async demo runs without touching an Azure subscription.

The source is in place; I’m compiling and running the offline demo now to catch SDK signature or reactive-type mismatches.

The demo now compiles and runs. I’m tightening the race-condition branch to recognize only the Azure “blob archived” conflict and adding focused parser/subject tests rather than swallowing unrelated storage conflicts.

Created the Java 17 Maven project with:

- Sync and async Event Grid/CloudEvents 1.0 receivers
- Blob-created/download and blob-deleted handling
- Graceful missing, archived, and rehydrating blob handling
- Sync and async Event Grid publishers
- Managed identity configuration with no keys or SAS tokens
- Offline `Main` demo covering both schemas and implementations
- Focused parser and subject-decoding tests
- Complete `pom.xml` and usage documentation

Run it with:

```powershell
mvn compile exec:java
```

The Azure client wiring is production-ready, while the demo uses in-memory blobs and dry-run publishers to avoid accessing live Azure resources.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- BlobServiceClient and BlobServiceAsyncClient do not require explicit closing.

### Issues

- The sync and async Event Grid publisher clients have no lifecycle-management path and can remain unclosed in production usage.

## Grader Results

- blob-event-notifier.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/10)
      - Handles Event Grid native schema via `EventGridEvent.fromString()` deserialization: Fail
      - Handles CloudEvents 1.0 schema via `CloudEvent.fromString()` deserialization: Fail
      - Does NOT manually parse JSON without the SDK's deserialization helpers: Fail
      - Routes events based on event type string (`Microsoft.Storage.BlobCreated`, `Microsoft.Storage.BlobDeleted`): Pass
      - Logs a warning for unrecognized event types: Pass
      - Parses container name and blob name from event subject (`/blobServices/default/containers/{container}/blobs/{blob}`): Pass
      - Publishes custom events with subject hierarchy for filtering: Pass
      - Retrieves and prints blob access tier from blob properties: Pass
      - Handles race condition: blob may no longer exist (catches `BlobStorageException` with 404 status): Pass
      - Catches Event Grid-specific exceptions for publishing errors: Fail
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
  - Try-With-Resources for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 60% | 1.00 | 0.6000 | 5.7% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 13.00** | **Σ 10.6000** | **81.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-blob-event-notifier --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

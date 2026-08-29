# Evaluation Report: key-vault-dp-java-crud

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 341.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-crud` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 341.3s |
| Timestamp | 2026-08-29T02:12:51Z |
| Files Generated | 3 |
| Event Count | 4386 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 113.6s |
| Review | 226.5s |
| **Total** | **341.3s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-security-keyvault-secrets-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=403145 out=3975 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 2401ms

**Input:**

```json
{"skill":"azure-security-keyvault-secrets-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Key Vault Secrets (Java)

Securely store and manage secrets like passwords, API keys, and connection strings.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-secrets</artifactId>
    <version>4.9.0</version>
</dependency>
```

## Client Creation

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

// Sync client
SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();

// Async client
SecretAsyncClient secretAsyncClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildAsyncClient();
```

## Create/Set Secret

```java
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

// Simple secret
KeyVaultSecret secret = secretClient.setSecret("database-password", "P@ssw0rd123!");
System.out.println("Secret name: " + secret.getName());
System.out.println("Secret ID: " + secret.getId());

// Secret with options
KeyVaultSecret secretWithOptions = secretClient.setSecret(
    new KeyVaultSecre
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 11596ms

**Input:**

```json
{"intent":"Generate a local-only Java Maven application using DefaultAzureCredential and Azure Key Vault Secrets SDK to create, read, update, soft-delete, and purge one secret with robust exception handling. Do not deploy or modify Azure infrastructure.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 13085ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md
@@ -1,396 +1,396 @@
 # Azure Key Vault Secrets SDK for Java - Examples
 
 Comprehensive code examples for the Azure Key Vault Secrets SDK for Java.
 
 ## Table of Contents
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Setting Secrets](#setting-secrets)
 - [Getting Secrets](#getting-secrets)
 - [Listing Secrets](#listing-secrets)
 - [Updating Secret Properties](#updating-secret-properties)
 - [Deleting and Recovering Secrets](#deleting-and-recovering-secrets)
 - [Purging Deleted Secrets](#purging-deleted-secrets)
 - [Backup and Restore](#backup-and-restore)
 - [Async Client Patterns](#async-client-patterns)
 - [Error Handling](#error-handling)
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-security-keyvault-secrets</artifactId>
     <version>4.11.0-beta.1</version>
 </dependency>
 
 <!-- Required for authentication -->
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.18.2</version>
 </dependency>
 ```
 
 ## Client Creation
 
 ### Sync SecretClient
 
 ```java
 import com.azure.identity.DefaultAzureCredentialBuilder;
 import com.azure.security.keyvault.secrets.SecretClient;
 import com.azure.security.keyvault.secrets.SecretClientBuilder;
 

... (truncated)
```

### ✅ 4. glob — 10312ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 9092ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 1689ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure Java SDK code-generation best practices to a local Maven sample for Azure Key Vault secret CRUD using DefaultAzureCredential; no infrastructure deployment.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. web_fetch — 6276ms

**Input:**

```json
{"max_length":5000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-security-keyvault-secrets%22\u0026rows=1\u0026wt=json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-identity%22&rows=1&wt=json:
{"responseHeader":{"status":0,"QTime":1,"params":{"q":"g:\"com.azure\" AND a:\"azure-identity\"","core":"","indent":"off","spellcheck":"true","fl":"id,g,a,latestVersion,p,ec,repositoryId,text,timestamp,versionCount","start":"","spellcheck.count":"5","sort":"score desc,timestamp desc,g asc,a asc","rows":"1","wt":"json","version":"2.2"}},"response":{"numFound":1,"start":0,"docs":[{"id":"com.azure:azure-identity","g":"com.azure","a":"azure-identity","latestVersion":"1.16.2","repositoryId":"central","p":"jar","timestamp":1749530875060,"versionCount":117,"text":["com.azure","azure-identity","-sources.jar.sha256","-sources.jar","-readme.md",".pom","-javadoc.jar","-changelog.md","-javadoc.jar.sha256",".jar","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"],"ec":["-sources.jar.sha256","-sources.jar","-readme.md",".pom","-javadoc.jar","-changelog.md","-javadoc.jar.sha256",".jar","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"]}]},"spellcheck":{"suggestions":[]}}
```

### ✅ 8. web_fetch — 5956ms

**Input:**

```json
{"max_length":5000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-identity%22\u0026rows=1\u0026wt=json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-security-keyvault-secrets%22&rows=1&wt=json:
{"responseHeader":{"status":0,"QTime":1,"params":{"q":"g:\"com.azure\" AND a:\"azure-security-keyvault-secrets\"","core":"","indent":"off","spellcheck":"true","fl":"id,g,a,latestVersion,p,ec,repositoryId,text,timestamp,versionCount","start":"","spellcheck.count":"5","sort":"score desc,timestamp desc,g asc,a asc","rows":"1","wt":"json","version":"2.2"}},"response":{"numFound":1,"start":0,"docs":[{"id":"com.azure:azure-security-keyvault-secrets","g":"com.azure","a":"azure-security-keyvault-secrets","latestVersion":"4.10.0","repositoryId":"central","p":"jar","timestamp":1749870559369,"versionCount":79,"text":["com.azure","azure-security-keyvault-secrets","-sources.jar","-sources.jar.sha256","-readme.md",".pom","-javadoc.jar","-changelog.md","-javadoc.jar.sha256",".jar","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"],"ec":["-sources.jar","-sources.jar.sha256","-readme.md",".pom","-javadoc.jar","-changelog.md","-javadoc.jar.sha256",".jar","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"]}]},"spellcheck":{"suggestions":[]}}
```

### ✅ 9. web_fetch — 5611ms

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

### ✅ 10. apply_patch — 716ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-secret-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+            \u003cversion\u003e4.11.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.18.2\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.KeyVaultSecretCrud\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/KeyVaultSecretCrud.java\n+package com.example;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.core.util.polling.SyncPoller;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+import com.azure.security.keyvault.secrets.models.DeletedSecret;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+public final class KeyVaultSecretCrud {\n+    private static final String VAULT_URL_ENV = \"AZURE_KEY_VAULT_URL\";\n+    private static final String SECRET_NAME = \"my-secret\";\n+\n+    private KeyVaultSecretCrud() {\n+    }\n+\n+    public static void main(String[] args) {\n+        try {\n+            SecretClient secretClient = createSecretClient();\n+            runCrudOperations(secretClient);\n+        } catch (ResourceNotFoundException exception) {\n+            System.err.printf(\"The secret or deleted secret was not found: %s%n\", exception.getMessage());\n+            System.exit(1);\n+        } catch (ClientAuthenticationException exception) {\n+            System.err.printf(\"Azure authentication failed: %s%n\", exception.getMessage());\n+            System.err.println(\"Sign in with Azure CLI or configure a supported DefaultAzureCredential identity.\");\n+            System.exit(1);\n+        } catch (HttpResponseException exception) {\n+            int statusCode = exception.getResponse() == null\n+                ? -1\n+                : exception.getResponse().getStatusCode();\n+            System.err.printf(\"Key Vault request failed (HTTP %d): %s%n\", statusCode, exception.getMessage());\n+            if (statusCode == 403) {\n+                System.err.println(\n+                    \"Verify secret data-plane permissions and that purge protection does not block purging.\");\n+            }\n+            System.exit(1);\n+        } catch (IllegalStateException exception) {\n+            System.err.println(exception.getMessage());\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static SecretClient createSecretClient() {\n+        String vaultUrl = System.getenv(VAULT_URL_ENV);\n+        if (vaultUrl == null || vaultUrl.isBlank()) {\n+            throw new IllegalStateException(\n+                \"Set \" + VAULT_URL_ENV + \" to a vault URL such as https://my-vault.vault.azure.net.\");\n+        }\n+\n+        return new SecretClientBuilder()\n+            .vaultUrl(vaultUrl)\n+            .credential(new DefaultAzureCredentialBuilder().build())\n+            .buildClient();\n+    }\n+\n+    private static void runCrudOperations(SecretClient secretClient) {\n+        KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, \"my-secret-value\");\n+        System.out.printf(\"Created secret \\\"%s\\\".%n\", createdSecret.getName());\n+\n+        KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);\n+        System.out.printf(\"Read secret value: %s%n\", readSecret.getValue());\n+\n+        // A new value creates a new secret version; properties are updated through a separate API.\n+        KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, \"updated-value\");\n+        System.out.printf(\n+            \"Updated secret \\\"%s\\\" to value \\\"%s\\\".%n\",\n+            updatedSecret.getName(),\n+            updatedSecret.getValue());\n+\n+        SyncPoller\u003cDeletedSecret, Void\u003e deletePoller = secretClient.beginDeleteSecret(SECRET_NAME);\n+        deletePoller.waitForCompletion();\n+        System.out.printf(\"Deleted secret \\\"%s\\\".%n\", SECRET_NAME);\n+\n+        secretClient.purgeDeletedSecret(SECRET_NAME);\n+        System.out.printf(\"Purged secret \\\"%s\\\" permanently.%n\", SECRET_NAME);\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault secret CRUD with Java\n+\n+This Maven application creates `my-secret`, reads and prints its value, creates\n+a new version with the value `updated-value`, then soft-deletes and purges the\n+secret.\n+\n+## Prerequisites\n+\n+- JDK 17 and Maven\n+- An existing Azure Key Vault with soft delete enabled\n+- An identity available to `DefaultAzureCredential`\n+- Secret data-plane permissions to set, get, delete, and purge secrets\n+\n+Set the vault URL without storing credentials in source code:\n+\n+```powershell\n+$env:AZURE_KEY_VAULT_URL = \"https://\u003cvault-name\u003e.vault.azure.net\"\n+```\n+\n+For local development, `DefaultAzureCredential` can use an existing Azure CLI\n+login. In Azure, prefer a managed identity with least-privilege Key Vault RBAC.\n+\n+## Build and run\n+\n+```powershell\n+mvn compile\n+mvn exec:java\n+```\n+\n+Purging is irreversible. If purge protection is enabled on the vault, Azure\n+Key Vault rejects manual purge until the retention period expires.\n+\n+Reference: [Azure Key Vault Secret client library for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/pom.xml
@@ -1,0 +1,42 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>key-vault-secret-crud</artifactId>
+    <version>1.0.0</version>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-security-keyvault-secrets</artifactId>
+            <version>4.11.0</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>1.18.2</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.codehaus.mojo</groupId>
+                <artifactId>exec-maven-plugin</artifactId>
+                <version>3.5.1</version>
+                <configuration>
+                    <mainClass>com.example.KeyVaultSecretCrud</mainClass>
+                </configuration>
+            </plugin>
+        </plugins>
+    </build>
+</project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyo
... (truncated)
```

### ✅ 11. powershell — 4417ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile the Maven application","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] -----------------< com.example:key-vault-secret-crud >------------------
[INFO] Building key-vault-secret-crud 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-secret-crud ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162\src\main\resources
[INFO] 
[INFO] --- compiler:3.13.0:compile (default-compile) @ key-vault-secret-crud ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  2.055 s
[INFO] Finished at: 2026-08-29T10:14:22+08:00
[INFO] ------------------------------------------------------------------------
<shellId: 0 completed with exit code 0>
```

### ✅ 12. apply_patch — 726ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/KeyVaultSecretCrud.java\n@@\n         KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, \"updated-value\");\n         System.out.printf(\n-            \"Updated secret \\\"%s\\\" to value \\\"%s\\\".%n\",\n+            \"Updated secret \\\"%s\\\" by creating version \\\"%s\\\".%n\",\n             updatedSecret.getName(),\n-            updatedSecret.getValue());\n+            updatedSecret.getProperties().getVersion());\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/src/main/java/com/example/KeyVaultSecretCrud.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/src/main/java/com/example/KeyVaultSecretCrud.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/src/main/java/com/example/KeyVaultSecretCrud.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162/src/main/java/com/example/KeyVaultSecretCrud.java
@@ -67,9 +67,9 @@
         // A new value creates a new secret version; properties are updated through a separate API.
         KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, "updated-value");
         System.out.printf(
-            "Updated secret \"%s\" to value \"%s\".%n",
+            "Updated secret \"%s\" by creating version \"%s\".%n",
             updatedSecret.getName(),
-            updatedSecret.getValue());
+            updatedSecret.getProperties().getVersion());
 
         SyncPoller<DeletedSecret, Void> deletePoller = secretClient.beginDeleteSecret(SECRET_NAME);
         deletePoller.waitForCompletion();

```

### ✅ 13. powershell — 4313ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Recompile after final code change","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] -----------------< com.example:key-vault-secret-crud >------------------
[INFO] Building key-vault-secret-crud 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-secret-crud ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1239748162\src\main\resources
[INFO] 
[INFO] --- compiler:3.13.0:compile (default-compile) @ key-vault-secret-crud ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  2.025 s
[INFO] Finished at: 2026-08-29T10:14:39+08:00
[INFO] ------------------------------------------------------------------------
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\KeyVaultSecretCrud.java`

## Copilot Response

I’ll scaffold a local Maven app, using the SDK’s delete poller before purge so the lifecycle is deterministic.

The workspace is empty, so I’m adding a complete Maven project with configuration and run instructions rather than only a standalone snippet.

Implemented a complete Maven application:

- `pom.xml` — Azure Key Vault Secrets `4.11.0` and Azure Identity `1.18.2`
- `src\main\java\com\example\KeyVaultSecretCrud.java` — create, read, update, delete, and purge using `DefaultAzureCredential`
- `README.md` — configuration, permissions, and execution instructions

Set the vault URL and run:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
mvn compile
mvn exec:java
```

The application handles authentication failures, missing secrets, HTTP errors, missing configuration, and purge-protection failures. Purging requires the relevant Key Vault data-plane permission and is unavailable when purge protection prevents manual purge.

Reference: [Azure Key Vault Secret client library for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable)

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The Azure SecretClient lifecycle is handled correctly for a client that does not require explicit closure.

## Grader Results

- crud-secrets.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Maven dependency for `azure-security-keyvault-secrets` and `azure-identity`: Pass
      - Creating a `SecretClient` with `SecretClientBuilder`: Pass
      - `setSecret()`, `getSecret()`, `beginDeleteSecret()`, `purgeDeletedSecret()`: Pass
      - Using `SyncPoller` to wait for delete completion before purge: Pass
      - Exception handling for `HttpResponseException`: Pass
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
hyoka run --prompt-id key-vault-dp-java-crud --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

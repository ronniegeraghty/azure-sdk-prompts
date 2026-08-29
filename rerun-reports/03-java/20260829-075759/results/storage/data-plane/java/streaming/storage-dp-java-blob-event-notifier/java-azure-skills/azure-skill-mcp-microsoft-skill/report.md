# Evaluation Report: storage-dp-java-blob-event-notifier

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 404.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-blob-event-notifier` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 20/22 |
| Duration | 404.3s |
| Timestamp | 2026-08-29T03:54:24Z |
| Files Generated | 17 |
| Event Count | 15128 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 162.0s |
| Review | 241.0s |
| **Total** | **404.3s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-eventgrid-java, azure-storage-blob-java, azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=468509 out=10614 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 2463ms

**Input:**

```json
{"skill":"azure-eventgrid-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Event Grid SDK for Java

Build event-driven applications using the Azure Event Grid SDK for Java.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-eventgrid</artifactId>
    <version>4.27.0</version>
</dependency>
```

## Client Creation

### EventGridPublisherClient

```java
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.azure.core.credential.AzureKeyCredential;

// With API Key
EventGridPublisherClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
    .endpoint("<topic-endpoint>")
    .credential(new AzureKeyCredential("<access-key>"))
    .buildEventGridEventPublisherClient();

// For CloudEvents
EventGridPublisherClient<CloudEvent> cloudClient = new EventGridPublisherClientBuilder()
    .endpoint("<topic-endpoint>")
    .credential(new AzureKeyCredential("<access-key>"))
    .buildCloudEventPublisherClient();
```

### With DefaultAzureCredential

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

EventGridPublisherClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
    .endpoint("<topic-endpoint>")
    .credential(credential)
    .buildEventGridEventP
... (truncated)
```

### ✅ 2. skill — 2881ms

**Input:**

```json
{"skill":"azure-storage-blob-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Blob SDK for Java

Build blob storage applications using the Azure Storage Blob SDK for Java.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-storage-blob</artifactId>
    <version>12.33.0</version>
</dependency>
```

## Client Creation

### BlobServiceClient

```java
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

// With SAS token
BlobServiceClient serviceClient = new BlobServiceClientBuilder()
    .endpoint("<storage-account-url>")
    .sasToken("<sas-token>")
    .buildClient();

// With connection string
BlobServiceClient serviceClient = new BlobServiceClientBuilder()
    .connectionString("<connection-string>")
    .buildClient();
```

### With DefaultAzureCredential

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

BlobServiceClient serviceClient = new BlobServiceClientBuilder()
    .endpoint("<storage-account-url>")
    .credential(credential)
    .buildClient();
```

### BlobContainerClient

```java
import com.azure.storage.blob.BlobContainerClient;

// From service client
BlobContainerClient containerClient = serviceClient.getBlobContainerClient("mycontainer");

// Direct construction
BlobContainerClient containerCli
... (truncated)
```

### ✅ 3. skill — 2971ms

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

### ✅ 4. azure-get_azure_bestpractices (via azure) — 9581ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Blob Storage, Event Grid EventGridEvent and CloudEvent schemas, managed identity authentication, synchronous and Reactor asynchronous clients. Do not deploy or contact Azure.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. view — 10548ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-eventgrid-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-eventgrid-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-eventgrid-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-eventgrid-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-eventgrid-java/references/examples.md
@@ -1,392 +1,392 @@
 # Azure Event Grid SDK for Java - Examples
 
 Comprehensive code examples for the Azure Event Grid SDK for Java.
 
 ## Table of Contents
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Publishing CloudEvents](#publishing-cloudevents)
 - [Publishing EventGridEvents](#publishing-eventgridevents)
 - [Publishing Custom Events](#publishing-custom-events)
 - [Async Client Patterns](#async-client-patterns)
 - [Batch Publishing](#batch-publishing)
 - [Error Handling](#error-handling)
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-messaging-eventgrid</artifactId>
     <version>4.32.0-beta.1</version>
 </dependency>
 
 <!-- For DefaultAzureCredential authentication -->
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.18.2</version>
 </dependency>
 ```
 
 ## Client Creation
 
 ### Sync Client with DefaultAzureCredential
 
 ```java
 import com.azure.core.models.CloudEvent;
 import com.azure.identity.DefaultAzureCredentialBuilder;
 import com.azure.messaging.eventgrid.EventGridPublisherClient;
 import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
 
 EventGridPublisherClient<CloudEvent> cloudEventClient = new EventGridPublisherClientBuilder()
     .endpoint("<endpoint of your ev
... (truncated)
```

### ✅ 6. view — 10221ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-storage-blob-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md
@@ -1,411 +1,411 @@
 # Azure Storage Blob Java SDK - Examples
 
 Comprehensive code examples for the Azure Storage Blob SDK for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Container Operations](#container-operations)
 - [Upload Blobs](#upload-blobs)
 - [Download Blobs](#download-blobs)
 - [List Blobs](#list-blobs)
 - [SAS Token Generation](#sas-token-generation)
 - [Error Handling](#error-handling)
 
 ---
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-storage-blob</artifactId>
     <version>12.33.0</version>
 </dependency>
 ```
 
 Or use the BOM:
 
 ```xml
 <dependencyManagement>
     <dependencies>
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-sdk-bom</artifactId>
             <version>{bom_version}</version>
             <type>pom</type>
             <scope>import</scope>
         </dependency>
     </dependencies>
 </dependencyManagement>
 
 <dependencies>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-storage-blob</artifactId>
     </dependency>
 </dependencies>
 ```
 
 ---
 
 ## Client Creation
 
 ### Using Shared Key Credential
 
 ```java
 import com.azure.storage.blob.BlobServiceClient;
 import com.azure.storage.blob.BlobSe
... (truncated)
```

### ✅ 7. view — 9894ms

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

### ✅ 8. glob — 9245ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. glob — 8932ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 10. azure-get_azure_bestpractices (via azure) — 709ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only Java 17 Maven sample using Azure Blob Storage, Event Grid EventGridEvent and CloudEvent schemas, managed identity authentication, synchronous and Reactor asynchronous clients. Do not deploy or contact Azure.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 11. apply_patch — 799ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eblob-event-notifier\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eBlob Event Notifier\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+        \u003cazure.eventgrid.version\u003e4.27.0\u003c/azure.eventgrid.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.33.0\u003c/azure.storage.blob.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+        \u003cjunit.version\u003e5.12.2\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventgrid\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.eventgrid.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.blobevents.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/blobevents/model/IncomingEvent.java\n+package com.example.blobevents.model;\n+\n+import com.azure.core.util.BinaryData;\n+import java.time.OffsetDateTime;\n+\n+public record IncomingEvent(\n+    Schema schema,\n+    String id,\n+    String type,\n+    String subject,\n+    OffsetDateTime time,\n+    BinaryData data\n+) {\n+    public enum Schema {\n+        EVENT_GRID,\n+        CLOUD_EVENT\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/model/BlobLocation.java\n+package com.example.blobevents.model;\n+\n+import java.net.URLDecoder;\n+import java.nio.charset.StandardCharsets;\n+\n+public record BlobLocation(String container, String blobName) {\n+    private static final String CONTAINERS = \"/containers/\";\n+    private static final String BLOBS = \"/blobs/\";\n+\n+    public static BlobLocation fromSubject(String subject) {\n+        if (subject == null) {\n+            throw new IllegalArgumentException(\"Event subject is required\");\n+        }\n+\n+        int containerStart = subject.indexOf(CONTAINERS);\n+        int blobMarker = subject.indexOf(BLOBS, containerStart + CONTAINERS.length());\n+        if (containerStart \u003c 0 || blobMarker \u003c 0) {\n+            throw new IllegalArgumentException(\"Unsupported blob event subject: \" + subject);\n+        }\n+\n+        String container = subject.substring(containerStart + CONTAINERS.length(), blobMarker);\n+        String blobName = subject.substring(blobMarker + BLOBS.length());\n+        if (container.isBlank() || blobName.isBlank()) {\n+            throw new IllegalArgumentException(\"Blob event subject has an empty container or blob name: \" + subject);\n+        }\n+        return new BlobLocation(decodePathValue(container), decodePathValue(blobName));\n+    }\n+\n+    private static String decodePathValue(String value) {\n+        return URLDecoder.decode(value.replace(\"+\", \"%2B\"), StandardCharsets.UTF_8);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/model/BlobSummary.java\n+package com.example.blobevents.model;\n+\n+public record BlobSummary(String name, long size, String contentType, String accessTier) {\n+}\n*** Add File: src/main/java/com/example/blobevents/model/CustomEvent.java\n+package com.example.blobevents.model;\n+\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public record CustomEvent(String subject, String type, Map\u003cString, Object\u003e data, String dataVersion) {\n+    public CustomEvent {\n+        Objects.requireNonNull(subject, \"subject\");\n+        Objects.requireNonNull(type, \"type\");\n+        data = Map.copyOf(Objects.requireNonNull(data, \"data\"));\n+        Objects.requireNonNull(dataVersion, \"dataVersion\");\n+        if (!subject.startsWith(\"/\")) {\n+            throw new IllegalArgumentException(\"Subject must be an absolute hierarchy beginning with '/'\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/storage/BlobStore.java\n+package com.example.blobevents.storage;\n+\n+import com.example.blobevents.model.BlobSummary;\n+import reactor.core.publisher.Mono;\n+\n+public interface BlobStore {\n+    BlobSummary download(String container, String blobName);\n+\n+    Mono\u003cBlobSummary\u003e downloadAsync(String container, String blobName);\n+}\n*** Add File: src/main/java/com/example/blobevents/storage/AzureBlobStore.java\n+package com.example.blobevents.storage;\n+\n+import com.azure.storage.blob.BlobAsyncClient;\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.models.BlobProperties;\n+import com.example.blobevents.model.BlobSummary;\n+import reactor.core.publisher.Mono;\n+\n+public final class AzureBlobStore implements BlobStore {\n+    private final BlobServiceClient syncClient;\n+    private final BlobServiceAsyncClient asyncClient;\n+\n+    public AzureBlobStore(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {\n+        this.syncClient = syncClient;\n+        this.asyncClient = asyncClient;\n+    }\n+\n+    @Override\n+    public BlobSummary download(String container, String blobName) {\n+        BlobClient blob = syncClient.getBlobContainerClient(container).getBlobClient(blobName);\n+        BlobProperties properties = blob.getProperties();\n+        blob.downloadContent();\n+        return toSummary(blobName, properties);\n+    }\n+\n+    @Override\n+    public Mono\u003cBlobSummary\u003e downloadAsync(String container, String blobName) {\n+        BlobAsyncClient blob = asyncClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(blobName);\n+        return blob.getProperties()\n+            .flatMap(properties -\u003e blob.downloadContent().thenReturn(toSummary(blobName, properties)));\n+    }\n+\n+    private static BlobSummary toSummary(String name, BlobProperties properties) {\n+        String tier = properties.getAccessTier() == null ? \"unknown\" : properties.getAccessTier().toString();\n+        return new BlobSummary(name, properties.getBlobSize(), properties.getContentType(), tier);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/storage/InMemoryBlobStore.java\n+package com.example.blobevents.storage;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.example.blobevents.model.BlobSummary;\n+import java.util.Map;\n+import reactor.core.publisher.Mono;\n+\n+public final class InMemoryBlobStore implements BlobStore {\n+    private final Map\u003cString, BlobSummary\u003e blobs;\n+\n+    public InMemoryBlobStore(Map\u003cString, BlobSummary\u003e blobs) {\n+        this.blobs = Map.copyOf(blobs);\n+    }\n+\n+    @Override\n+    public BlobSummary download(String container, String blobName) {\n+        BlobSummary summary = blobs.get(container + \"/\" + blobName);\n+        if (summary == null) {\n+            throw new BlobStorageException(\"Mock blob was deleted before it could be read\", null, null);\n+        }\n+        return summary;\n+    }\n+\n+    @Override\n+    public Mono\u003cBlobSummary\u003e downloadAsync(String container, String blobName) {\n+        return Mono.fromCallable(() -\u003e download(container, blobName));\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobEventHandler.java\n+package com.example.blobevents;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.example.blobevents.model.BlobLocation;\n+import com.example.blobevents.model.BlobSummary;\n+import com.example.blobevents.model.IncomingEvent;\n+import com.example.blobevents.storage.BlobStore;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+public final class BlobEventHandler {\n+    public static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    public static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);\n+    private final BlobStore blobStore;\n+\n+    public BlobEventHandler(BlobStore blobStore) {\n+        this.blobStore = blobStore;\n+    }\n+\n+    public void handleCreated(IncomingEvent event) {\n+        BlobLocation location = BlobLocation.fromSubject(event.subject());\n+        try {\n+            printSummary(location, blobStore.download(location.container(), location.blobName()));\n+        } catch (BlobStorageException exception) {\n+            if (isLifecycleRace(exception)) {\n+                LOGGER.warn(\"Blob {}/{} is no longer readable; lifecycle processing likely changed it\",\n+                    location.container(), location.blobName());\n+                return;\n+            }\n+            throw exception;\n+        }\n+    }\n+\n+    public Mono\u003cVoid\u003e handleCreatedAsync(IncomingEvent event) {\n+        BlobLocation location = BlobLocation.fromSubject(event.subject());\n+        return blobStore.downloadAsync(location.container(), location.blobName())\n+            .doOnNext(summary -\u003e printSummary(location, summary))\n+            .then()\n+            .onErrorResume(BlobStorageException.class, exception -\u003e {\n+                if (isLifecycleRace(exception)) {\n+                    LOGGER.warn(\"Blob {}/{} is no longer readable; lifecycle processing likely changed it\",\n+                        location.container(), location.blobName());\n+                    return Mono.empty();\n+                }\n+                return Mono.error(exception);\n+            });\n+    }\n+\n+    public void handleDeleted(IncomingEvent event) {\n+        BlobLocation location = BlobLocation.fromSubject(event.subject());\n+        LOGGER.info(\"Blob deleted: container={}, name={}\", location.container(), location.blobName());\n+    }\n+\n+    public Mono\u003cVoid\u003e handleDeletedAsync(IncomingEvent event) {\n+        return Mono.fromRunnable(() -\u003e handleDeleted(event));\n+    }\n+\n+    private static boolean isLifecycleRace(BlobStorageException exception) {\n+        return exception.getStatusCode() == 404 || exception.getStatusCode() == 409\n+            || exception.getResponse() == null;\n+    }\n+\n+    private static void printSummary(BlobLocation location, BlobSummary summary) {\n+        LOGGER.info(\"Blob created: container={}, name={}, size={} bytes, contentType={}, accessTier={}\",\n+            location.container(), summary.name(), summary.size(), summary.contentType(), summary.accessTier());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventReceiver.java\n+package com.example.blobevents;\n+\n+import com.azure.core.models.CloudEvent;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.example.blobevents.model.IncomingEvent;\n+import com.example.blobevents.model.IncomingEvent.Schema;\n+import java.util.List;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+public final class EventReceiver {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);\n+    private final BlobEventHandler handler;\n+\n+    public EventReceiver(BlobEventHandler handler) {\n+        this.handler = handler;\n+    }\n+\n+    public List\u003cIncomingEvent\u003e receive(String jsonPayload) {\n+        List\u003cIncomingEvent\u003e events = deserialize(jsonPayload);\n+        events.forEach(this::route);\n+        return events;\n+    }\n+\n+    public static List\u003cIncomingEvent\u003e deserialize(String jsonPayload) {\n+        if (jsonPayload == null || jsonPayload.isBlank()) {\n+            throw new IllegalArgumentException(\"Event Grid payload must not be empty\");\n+        }\n+\n+        if (looksLikeCloudEvents(jsonPayload)) {\n+            return CloudEvent.fromString(jsonPayload).stream()\n+                .map(event -\u003e new IncomingEvent(\n+                    Schema.CLOUD_EVENT,\n+                    event.getId(),\n+                    event.getType(),\n+                    event.getSubject(),\n+                    event.getTime(),\n+                    event.getData()))\n+                .toList();\n+        }\n+\n+        return EventGridEvent.fromString(jsonPayload).stream()\n+            .map(event -\u003e new IncomingEvent(\n+                Schema.EVENT_GRID,\n+                event.getId(),\n+                event.getEventType(),\n+                event.getSubject(),\n+                event.getEventTime(),\n+                event.getData()))\n+            .toList();\n+    }\n+\n+    private void route(IncomingEvent event) {\n+        switch (event.type()) {\n+            case BlobEventHandler.BLOB_CREATED -\u003e handler.handleCreated(event);\n+            case BlobEventHandler.BLOB_DELETED -\u003e handler.handleDeleted(event);\n+            default -\u003e LOGGER.warn(\"Ignoring unrecognized Event Grid event type: {}\", event.type());\n+        }\n+    }\n+\n+    private static boolean looksLikeCloudEvents(String payload) {\n+        return payload.contains(\"\\\"specversion\\\"\") || payload.contains(\"\\\"specVersion\\\"\");\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventReceiver.java\n+package com.example.blobevents;\n+\n+import com.example.blobevents.model.IncomingEvent;\n+import java.util.List;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Flux;\n+\n+public final class AsyncEventReceiver {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncEventReceiver.class);\n+    private final BlobEventHandler handler;\n+\n+    public AsyncEventReceiver(BlobEventHandler handler) {\n+        this.handler = handler;\n+    }\n+\n+    public Flux\u003cIncomingEvent\u003e receive(String jsonPayload) {\n+        return Flux.defer(() -\u003e {\n+            List\u003cIncomingEvent\u003e events = EventReceiver.deserialize(jsonPayload);\n+            return Flux.fromIterable(events);\n+        }).concatMap(event -\u003e route(event).thenReturn(event));\n+    }\n+\n+    private reactor.core.publisher.Mono\u003cVoid\u003e route(IncomingEvent event) {\n+        return switch (event.type()) {\n+            case BlobEventHandler.BLOB_CREATED -\u003e handler.handleCreatedAsync(event);\n+            case BlobEventHandler.BLOB_DELETED -\u003e handler.handleDeletedAsync(event);\n+            default -\u003e reactor.core.publisher.Mono.fromRunnable(\n+                () -\u003e LOGGER.warn(\"Ignoring unrecognized Event Grid event type: {}\", event.type()));\n+        };\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import com.example.blobevents.model.CustomEvent;\n+import java.util.List;\n+import java.util.function.Consumer;\n+\n+public final class EventPublisher {\n+    private final Consumer\u003cList\u003cEventGridEvent\u003e\u003e sender;\n+\n+    public EventPublisher(EventGridPublisherClient\u003cEventGridEvent\u003e client) {\n+        this(client::sendEvents);\n+    }\n+\n+    public EventPublisher(Consumer\u003cList\u003cEventGridEvent\u003e\u003e sender) {\n+        this.sender = sender;\n+    }\n+\n+    public void publish(List\u003cCustomEvent\u003e customEvents) {\n+        if (customEvents == null || customEvents.isEmpty()) {\n+            throw new IllegalArgumentException(\"At least one custom event is required\");\n+        }\n+        sender.accept(customEvents.stream().map(EventPublisher::toEventGridEvent).toList());\n+    }\n+\n+    static EventGridEvent toEventGridEvent(CustomEvent event) {\n+        return new EventGridEvent(\n+            event.subject(),\n+            event.type(),\n+            BinaryData.fromObject(event.data()),\n+            event.dataVersion());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import com.example.blobevents.model.CustomEvent;\n+import java.util.List;\n+import java.util.function.Function;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncEventPublisher {\n+    private final Function\u003cList\u003cEventGridEvent\u003e, Mono\u003cVoid\u003e\u003e sender;\n+\n+    public AsyncEventPublisher(EventGridPublisherAsyncClient\u003cEventGridEvent\u003e client) {\n+        this(client::sendEvents);\n+    }\n+\n+    public AsyncEventPublisher(Function\u003cList\u003cEventGridEvent\u003e, Mono\u003cVoid\u003e\u003e sender) {\n+        this.sender = sender;\n+    }\n+\n+    public Mono\u003cVoid\u003e publish(List\u003cCustomEvent\u003e customEvents) {\n+        if (customEvents == null || customEvents.isEmpty()) {\n+            return Mono.error(new IllegalArgumentException(\"At least one custom event is required\"));\n+        }\n+        return sender.apply(customEvents.stream().map(EventPublisher::toEventGridEvent).toList());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureConfiguration.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+import com.example.blobevents.storage.AzureBlobStore;\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class AzureConfiguration {\n+    private final String storageAccountUrl;\n+    private final String eventGridTopicEndpoint;\n+    private final TokenCredential credential;\n+\n+    public AzureConfiguration(String storageAccountUrl, String eventGridTopicEndpoint, String managedIdentityClientId) {\n+        this.storageAccountUrl = requireHttpsUrl(storageAccountUrl, \"storageAccountUrl\");\n+        this.eventGridTopicEndpoint = requireHttpsUrl(eventGridTopicEndpoint, \"eventGridTopicEndpoint\");\n+\n+        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            builder.clientId(managedIdentityClientId);\n+        }\n+        this.credential = builder.build();\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        Map\u003cString, String\u003e environment = System.getenv();\n+        return new AzureConfiguration(\n+            environment.get(\"AZURE_STORAGE_ACCOUNT_URL\"),\n+            environment.get(\"EVENT_GRID_TOPIC_ENDPOINT\"),\n+            environment.get(\"AZURE_CLIENT_ID\"));\n+    }\n+\n+    public AzureBlobStore blobStore() {\n+        BlobServiceClientBuilder builder = new BlobServiceClientBuilder()\n+            .endpoint(storageAccountUrl)\n+            .credential(credential);\n+        return new AzureBlobStore(builder.buildClient(), builder.buildAsyncClient());\n+    }\n+\n+    public EventPublisher eventPublisher() {\n+        return new EventPublisher(new EventGridPublisherClientBuilder()\n+            .endpoint(eventGridTopicEndpoint)\n+            .credential(credential)\n+            .buildEventGridEventPublisherClient());\n+    }\n+\n+    public AsyncEventPublisher asyncEventPublisher() {\n+        return new AsyncEventPublisher(new EventGridPublisherClientBuilder()\n+            .endpoint(eventGridTopicEndpoint)\n+            .credential(credential)\n+            .buildEventGridEventPublisherAsyncClient());\n+    }\n+\n+    private static String requireHttpsUrl(String value, String name) {\n+        Objects.requireNonNull(value, name + \" is required\");\n+        if (!value.startsWith(\"https://\")) {\n+            throw new IllegalArgumentException(name + \" must use HTTPS\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/Main.java\n+package com.example.blobevents;\n+\n+import com.example.blobevents.model.BlobSummary;\n+import com.example.blobevents.model.CustomEvent;\n+import com.example.blobevents.storage.InMemoryBlobStore;\n+import java.util.List;\n+import java.util.Map;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+public final class Main {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        InMemoryBlobStore demoStore = new InMemoryBlobStore(Map.of(\n+            \"documents/invoices/august-2026.pdf\",\n+            new BlobSummary(\"invoices/august-2026.pdf\", 184_320, \"application/pdf\", \"COOL\")));\n+        BlobEventHandler handler = new BlobEventHandler(demoStore);\n+        List\u003cCustomEvent\u003e downstreamEvents = List.of(new CustomEvent(\n+            \"/documents/invoices/processed\",\n+            \"Contoso.Documents.Processed\",\n+            Map.of(\"documentId\", \"august-2026\", \"status\", \"processed\"),\n+            \"1.0\"));\n+\n+        LOGGER.info(\"----- synchronous demo -----\");\n+        EventReceiver receiver = new EventReceiver(handler);\n+        receiver.receive(eventGridPayload());\n+        receiver.receive(cloudEventsPayload());\n+        new EventPublisher(events -\u003e events.forEach(event -\u003e\n+            LOGGER.info(\"Mock publish: type={}, subject={}\", event.getEventType(), event.getSubject())))\n+            .publish(downstreamEvents);\n+\n+        LOGGER.info(\"----- asynchronous demo -----\");\n+        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver(handler);\n+        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -\u003e Mono.fromRunnable(() -\u003e\n+            events.forEach(event -\u003e\n+                LOGGER.info(\"Mock async publish: type={}, subject={}\", event.getEventType(), event.getSubject()))));\n+\n+        asyncReceiver.receive(eventGridPayload())\n+            .thenMany(asyncReceiver.receive(cloudEventsPayload()))\n+            .then(asyncPublisher.publish(downstreamEvents))\n+            .block();\n+    }\n+\n+    private static String eventGridPayload() {\n+        return \"\"\"\n+            [\n+              {\n+                \"id\": \"2f01f1f8-4f52-4b6d-a9e1-0cf1a833f000\",\n+                \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/invoices/august-2026.pdf\",\n+                \"eventTime\": \"2026-08-29T03:50:00Z\",\n+                \"data\": {\n+                  \"api\": \"PutBlob\",\n+                  \"clientRequestId\": \"9f621d84-76c8-4a9d-81f3-fdf58d7b1077\",\n+                  \"requestId\": \"e4b171fe-501e-0013-137c-f0cacc000000\",\n+                  \"eTag\": \"0x8DEDEADBEEF0000\",\n+                  \"contentType\": \"application/pdf\",\n+                  \"contentLength\": 184320,\n+                  \"blobType\": \"BlockBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/invoices/august-2026.pdf\",\n+                  \"sequencer\": \"0000000000000000000000000000001\"\n+                },\n+                \"dataVersion\": \"\",\n+                \"metadataVersion\": \"1\"\n+              },\n+              {\n+                \"id\": \"576499b2-8f4d-48ab-a98d-a6cd53eaf000\",\n+                \"eventType\": \"Microsoft.Storage.BlobDeleted\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf\",\n+                \"eventTime\": \"2026-08-29T03:51:00Z\",\n+                \"data\": {\n+                  \"api\": \"DeleteBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf\",\n+                  \"sequencer\": \"0000000000000000000000000000002\"\n+                },\n+                \"dataVersion\": \"\",\n+                \"metadataVersion\": \"1\"\n+              }\n+            ]\n+            \"\"\";\n+    }\n+\n+    private static String cloudEventsPayload() {\n+        return \"\"\"\n+            [\n+              {\n+                \"specversion\": \"1.0\",\n+                \"type\": \"Microsoft.Storage.BlobCreated\",\n+                \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage\",\n+                \"id\": \"825af45f-1d7a-4f2f-a93b-13568c24f000\",\n+                \"time\": \"2026-08-29T03:52:00Z\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/invoices/august-2026.pdf\",\n+                \"datacontenttype\": \"application/json\",\n+                \"data\": {\n+                  \"api\": \"PutBlob\",\n+                  \"contentType\": \"application/pdf\",\n+                  \"contentLength\": 184320,\n+                  \"blobType\": \"BlockBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/invoices/august-2026.pdf\",\n+                  \"sequencer\": \"0000000000000000000000000000003\"\n+                }\n+              },\n+              {\n+                \"specversion\": \"1.0\",\n+                \"type\": \"Microsoft.Storage.BlobDeleted\",\n+                \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage\",\n+                \"id\": \"dfab8ab6-abdb-4ae5-acf5-bd605c56f000\",\n+                \"time\": \"2026-08-29T03:53:00Z\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf\",\n+                \"datacontenttype\": \"application/json\",\n+                \"data\": {\n+                  \"api\": \"DeleteBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf\",\n+                  \"sequencer\": \"0000000000000000000000000000004\"\n+                }\n+              }\n+            ]\n+            \"\"\";\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/EventReceiverTest.java\n+package com.example.blobevents;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+import com.example.blobevents.model.BlobLocation;\n+import com.example.blobevents.model.IncomingEvent.Schema;\n+import org.junit.jupiter.api.Test;\n+\n+class EventReceiverTest {\n+    @Test\n+    void deserializesEventGridSchema() {\n+        String payload = \"\"\"\n+            [{\"id\":\"1\",\"eventType\":\"Microsoft.Storage.BlobDeleted\",\n+              \"subject\":\"/blobServices/default/containers/docs/blobs/a.txt\",\n+              \"eventTime\":\"2026-08-29T03:00:00Z\",\"data\":{},\"dataVersion\":\"\",\"metadataVersion\":\"1\"}]\n+            \"\"\";\n+\n+        var events = EventReceiver.deserialize(payload);\n+\n+        assertEquals(1, events.size());\n+        assertEquals(Schema.EVENT_GRID, events.get(0).schema());\n+        assertEquals(BlobEventHandler.BLOB_DELETED, events.get(0).type());\n+    }\n+\n+    @Test\n+    void deserializesCloudEventsSchema() {\n+        String payload = \"\"\"\n+            [{\"specversion\":\"1.0\",\"type\":\"Microsoft.Storage.BlobCreated\",\"source\":\"/storage\",\n+              \"id\":\"2\",\"time\":\"2026-08-29T03:00:00Z\",\n+              \"subject\":\"/blobServices/default/containers/docs/blobs/a.txt\",\n+              \"datacontenttype\":\"application/json\",\"data\":{}}]\n+            \"\"\";\n+\n+        var events = EventReceiver.deserialize(payload);\n+\n+        assertEquals(1, events.size());\n+        assertEquals(Schema.CLOUD_EVENT, events.get(0).schema());\n+        assertEquals(BlobEventHandler.BLOB_CREATED, events.get(0).type());\n+    }\n+\n+    @Test\n+    void parsesEncodedBlobSubject() {\n+        BlobLocation location = BlobLocation.fromSubject(\n+            \"/blobServices/default/containers/my-docs/blobs/2026/paid%20invoice+copy.pdf\");\n+\n+        assertEquals(\"my-docs\", location.container());\n+        assertEquals(\"2026/paid invoice+copy.pdf\", location.blobName());\n+    }\n+}\n*** Add File: README.md\n+# Blob Event Notifier\n+\n+Small Java 17 sample for receiving Azure Blob Storage lifecycle events through Event Grid and publishing downstream events.\n+\n+The demo is intentionally offline: it parses realistic Event Grid and CloudEvents payloads, uses an in-memory blob store, and logs mock publishes. Production adapters use only `ManagedIdentityCredential`; no access keys or SAS tokens are accepted.\n+\n+## Run locally\n+\n+```powershell\n+mvn test\n+mvn exec:java\n+```\n+\n+## Use Azure-backed adapters\n+\n+Set:\n+\n+- `AZURE_STORAGE_ACCOUNT_URL=https://\u003caccount\u003e.blob.core.windows.net`\n+- `EVENT_GRID_TOPIC_ENDPOINT=https://\u003ctopic\u003e.\u003cregion\u003e-1.eventgrid.azure.net/api/events`\n+- `AZURE_CLIENT_ID=\u003cuser-assigned-managed-identity-client-id\u003e` only for a user-assigned identity\n+\n+Construct `AzureConfiguration.fromEnvironment()`, then obtain `blobStore()`, `eventPublisher()`, and `asyncEventPublisher()`. Assign the managed identity **Storage Blob Data Reader** on the required storage scope and **EventGrid Data Sender** on the custom topic scope.\n+\n+Event Grid webhook authentication and subscription-validation handling belong in the hosting HTTP framework. Pass the validated request body to `EventReceiver.receive` or `AsyncEventReceiver.receive`.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-472879499/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-472879499/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-472879499/pom.xml
@@ -1,0 +1,75 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>blob-event-notifier</artifactId>
+    <version>1.0.0</version>
+    <name>Blob Event Notifier</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.identity.version>1.18.2</azure.identity.version>
+        <azure.eventgrid.version>4.27.0</azure.eventgrid.version>
+        <azure.storage.blob.version>12.33.0</azure.storage.blob.version>
+        <slf4j.version>2.0.17</slf4j.version>
+        <junit.version>5.12.2</junit.version>
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
+            <artifactId>azure-
... (truncated)
```

### ✅ 12. powershell — 11624ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress --quiet \u0026\u0026 mvn exec:java --no-transfer-progress --quiet","description":"Compile, test, and run the offline Maven demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 0 completed with exit code 0>
```

### ✅ 13. powershell — 4282ms

**Input:**

```json
{"command":"mvn exec:java --no-transfer-progress","description":"Inspect the offline demo output","initial_wait":120,"mode":"sync"}
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
[INFO] Building Blob Event Notifier 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- exec:3.5.0:java (default-cli) @ blob-event-notifier ---
[com.example.blobevents.Main.main()] INFO com.example.blobevents.Main - ----- synchronous demo -----
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob created: container=documents, name=invoices/august-2026.pdf, size=184320 bytes, contentType=application/pdf, accessTier=COOL
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob deleted: container=documents, name=archive/old-invoice.pdf
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob created: container=documents, name=invoices/august-2026.pdf, size=184320 bytes, contentType=application/pdf, accessTier=COOL
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob deleted: container=documents, name=archive/old-invoice.pdf
[com.example.blobevents.Main.main()] INFO com.example.blobevents.Main - Mock publish: type=Contoso.Documents.Processed, subject=/documents/invoices/processed
[com.example.blobevents.Main.main()] INFO com.example.blobevents.Main - ----- asynchronous demo -----
[com.example.blobevents.Main.main()] INFO com.example.blobevents.BlobEventHandler - Blob created: container=documents, name=invoices/au
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\blobevents\AsyncEventPublisher.java`
- `src\main\java\com\example\blobevents\AsyncEventReceiver.java`
- `src\main\java\com\example\blobevents\AzureConfiguration.java`
- `src\main\java\com\example\blobevents\BlobEventHandler.java`
- `src\main\java\com\example\blobevents\EventPublisher.java`
- `src\main\java\com\example\blobevents\EventReceiver.java`
- `src\main\java\com\example\blobevents\Main.java`
- `src\main\java\com\example\blobevents\model\BlobLocation.java`
- `src\main\java\com\example\blobevents\model\BlobSummary.java`
- `src\main\java\com\example\blobevents\model\CustomEvent.java`
- `src\main\java\com\example\blobevents\model\IncomingEvent.java`
- `src\main\java\com\example\blobevents\storage\AzureBlobStore.java`
- `src\main\java\com\example\blobevents\storage\BlobStore.java`
- `src\main\java\com\example\blobevents\storage\InMemoryBlobStore.java`
- `src\test\java\com\example\blobevents\EventReceiverTest.java`

## Copilot Response

I’ll align the project with the current Azure Java SDK patterns, then scaffold and locally validate the Maven build without contacting Azure.

The workspace is empty, so I’m creating a self-contained sample with Azure-backed adapters plus in-memory demo adapters; `Main` will exercise every path offline while production construction remains managed-identity-only.

The implementation is in place. I’m compiling the Azure SDK integrations and running the offline demo now; any SDK signature mismatches will be corrected against the actual resolved dependencies.

Created a complete Java 17 Maven project with:

- Sync and async Event Grid receivers supporting native Event Grid and CloudEvents 1.0 schemas
- Blob-created/deleted routing with unknown-event warnings
- Blob subject parsing, download summaries, and lifecycle race handling
- Sync and async custom event publishers with hierarchical subjects
- Managed-identity-only Azure configuration
- Offline `Main` demo using realistic payloads and mock storage/publishing
- Unit tests and usage/RBAC documentation in `README.md`

Run locally with:

```powershell
mvn test
mvn exec:java
```

Azure SDK references: [Event Grid for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-eventgrid-readme), [Blob Storage for Java](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java), and [Azure Identity for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Client lifecycle handling is appropriate for the Azure SDK client types used.

## Grader Results

- blob-event-notifier.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (9/10)
      - Handles Event Grid native schema via `EventGridEvent.fromString()` deserialization: Pass
      - Handles CloudEvents 1.0 schema via `CloudEvent.fromString()` deserialization: Pass
      - Does NOT manually parse JSON without the SDK's deserialization helpers: Pass
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
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 90% | 1.00 | 0.9000 | 7.6% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.4% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.9000** | **91.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-blob-event-notifier --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

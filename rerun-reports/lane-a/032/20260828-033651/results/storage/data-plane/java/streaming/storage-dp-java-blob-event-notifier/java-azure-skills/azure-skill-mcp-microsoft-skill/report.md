# Evaluation Report: storage-dp-java-blob-event-notifier

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 601.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-blob-event-notifier` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 19/22 |
| Duration | 601.5s |
| Timestamp | 2026-08-27T19:36:51Z |
| Files Generated | 16 |
| Event Count | 15092 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 348.7s |
| Review | 251.5s |
| **Total** | **601.5s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-eventgrid-java, azure-storage-blob-java, azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=370258 out=9030 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 3970ms

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

### ✅ 2. skill — 3466ms

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

### ✅ 3. skill — 1524ms

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

### ❌ 4. azure-get_azure_bestpractices (via azure) — 181929ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Blob Storage, Event Grid, and managed identity, with synchronous and asynchronous receiving, blob handling, and publishing.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. view — 6782ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-eventgrid-java\\references\\examples.md"}
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

### ✅ 6. view — 6411ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-storage-blob-java\\references\\examples.md"}
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

### ✅ 7. view — 6035ms

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

### ✅ 8. glob — 5609ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-2414409379","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. apply_patch — 899ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eblob-event-notifier\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0-SNAPSHOT\u003c/version\u003e\n+    \u003cname\u003eAzure Blob Event Notifier\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.eventgrid.version\u003e4.27.0\u003c/azure.eventgrid.version\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.33.0\u003c/azure.storage.blob.version\u003e\n+        \u003cjackson.version\u003e2.17.2\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.11.0\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventgrid\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.eventgrid.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.16\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.13.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.4.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.blobevents.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/blobevents/model/BlobLifecycleEvent.java\n+package com.example.blobevents.model;\n+\n+import com.azure.core.util.BinaryData;\n+\n+import java.time.OffsetDateTime;\n+\n+public record BlobLifecycleEvent(\n+    String id,\n+    String type,\n+    String subject,\n+    OffsetDateTime time,\n+    BinaryData data,\n+    EventSchema schema\n+) {\n+    public enum EventSchema {\n+        EVENT_GRID,\n+        CLOUD_EVENTS\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/model/CustomEvent.java\n+package com.example.blobevents.model;\n+\n+import java.util.Objects;\n+\n+public record CustomEvent(String type, String subject, String dataVersion, Object data) {\n+    public CustomEvent {\n+        Objects.requireNonNull(type, \"type\");\n+        Objects.requireNonNull(subject, \"subject\");\n+        Objects.requireNonNull(dataVersion, \"dataVersion\");\n+        Objects.requireNonNull(data, \"data\");\n+        if (!subject.startsWith(\"/\")) {\n+            throw new IllegalArgumentException(\"subject must be a hierarchy beginning with '/'\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/blob/BlobSummary.java\n+package com.example.blobevents.blob;\n+\n+public record BlobSummary(String name, long size, String contentType, String accessTier) {\n+}\n*** Add File: src/main/java/com/example/blobevents/blob/BlobOperations.java\n+package com.example.blobevents.blob;\n+\n+import com.azure.core.util.BinaryData;\n+import reactor.core.publisher.Mono;\n+\n+public interface BlobOperations {\n+    DownloadedBlob download(String container, String name);\n+\n+    Mono\u003cDownloadedBlob\u003e downloadAsync(String container, String name);\n+\n+    record DownloadedBlob(BinaryData content, BlobSummary summary) {\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/blob/AzureBlobOperations.java\n+package com.example.blobevents.blob;\n+\n+import com.azure.storage.blob.BlobAsyncClient;\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.models.BlobProperties;\n+import reactor.core.publisher.Mono;\n+\n+public final class AzureBlobOperations implements BlobOperations {\n+    private final BlobServiceClient syncClient;\n+    private final BlobServiceAsyncClient asyncClient;\n+\n+    public AzureBlobOperations(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {\n+        this.syncClient = syncClient;\n+        this.asyncClient = asyncClient;\n+    }\n+\n+    @Override\n+    public DownloadedBlob download(String container, String name) {\n+        BlobClient blob = syncClient.getBlobContainerClient(container).getBlobClient(name);\n+        BlobProperties properties = blob.getProperties();\n+        return new DownloadedBlob(blob.downloadContent(), summary(name, properties));\n+    }\n+\n+    @Override\n+    public Mono\u003cDownloadedBlob\u003e downloadAsync(String container, String name) {\n+        BlobAsyncClient blob = asyncClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(name);\n+        return blob.getProperties()\n+            .flatMap(properties -\u003e blob.downloadContent()\n+                .map(content -\u003e new DownloadedBlob(content, summary(name, properties))));\n+    }\n+\n+    private static BlobSummary summary(String name, BlobProperties properties) {\n+        String tier = properties.getAccessTier() == null ? \"unknown\" : properties.getAccessTier().toString();\n+        return new BlobSummary(name, properties.getBlobSize(), properties.getContentType(), tier);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/blob/BlobEventHandler.java\n+package com.example.blobevents.blob;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.example.blobevents.model.BlobLifecycleEvent;\n+\n+import java.net.URLDecoder;\n+import java.nio.charset.StandardCharsets;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+import java.util.regex.Matcher;\n+import java.util.regex.Pattern;\n+\n+public final class BlobEventHandler {\n+    private static final Logger LOGGER = Logger.getLogger(BlobEventHandler.class.getName());\n+    private static final Pattern SUBJECT_PATTERN =\n+        Pattern.compile(\"^/blobServices/default/containers/([^/]+)/blobs/(.+)$\");\n+\n+    private final BlobOperations blobs;\n+\n+    public BlobEventHandler(BlobOperations blobs) {\n+        this.blobs = blobs;\n+    }\n+\n+    public void handleCreated(BlobLifecycleEvent event) {\n+        BlobLocation location = parseSubject(event.subject());\n+        try {\n+            BlobSummary summary = blobs.download(location.container(), location.name()).summary();\n+            LOGGER.info(() -\u003e \"Blob created: name=%s, size=%d, contentType=%s, accessTier=%s\"\n+                .formatted(summary.name(), summary.size(), summary.contentType(), summary.accessTier()));\n+        } catch (BlobStorageException exception) {\n+            if (isLifecycleRace(exception)) {\n+                LOGGER.warning(() -\u003e \"Blob is no longer readable after creation event: \"\n+                    + location.container() + \"/\" + location.name() + \" (\" + exception.getStatusCode() + \")\");\n+                return;\n+            }\n+            throw exception;\n+        }\n+    }\n+\n+    public void handleDeleted(BlobLifecycleEvent event) {\n+        BlobLocation location = parseSubject(event.subject());\n+        LOGGER.info(() -\u003e \"Blob deleted: \" + location.container() + \"/\" + location.name());\n+    }\n+\n+    static BlobLocation parseSubject(String subject) {\n+        Matcher matcher = SUBJECT_PATTERN.matcher(subject);\n+        if (!matcher.matches()) {\n+            throw new IllegalArgumentException(\"Unexpected blob event subject: \" + subject);\n+        }\n+        return new BlobLocation(decode(matcher.group(1)), decode(matcher.group(2)));\n+    }\n+\n+    private static String decode(String value) {\n+        return URLDecoder.decode(value.replace(\"+\", \"%2B\"), StandardCharsets.UTF_8);\n+    }\n+\n+    static boolean isLifecycleRace(BlobStorageException exception) {\n+        return exception.getStatusCode() == 404 || exception.getStatusCode() == 409;\n+    }\n+\n+    record BlobLocation(String container, String name) {\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/blob/AsyncBlobEventHandler.java\n+package com.example.blobevents.blob;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.example.blobevents.model.BlobLifecycleEvent;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.logging.Logger;\n+\n+public final class AsyncBlobEventHandler {\n+    private static final Logger LOGGER = Logger.getLogger(AsyncBlobEventHandler.class.getName());\n+\n+    private final BlobOperations blobs;\n+\n+    public AsyncBlobEventHandler(BlobOperations blobs) {\n+        this.blobs = blobs;\n+    }\n+\n+    public Mono\u003cVoid\u003e handleCreated(BlobLifecycleEvent event) {\n+        BlobEventHandler.BlobLocation location = BlobEventHandler.parseSubject(event.subject());\n+        return blobs.downloadAsync(location.container(), location.name())\n+            .doOnNext(download -\u003e {\n+                BlobSummary summary = download.summary();\n+                LOGGER.info(() -\u003e \"Blob created (async): name=%s, size=%d, contentType=%s, accessTier=%s\"\n+                    .formatted(summary.name(), summary.size(), summary.contentType(), summary.accessTier()));\n+            })\n+            .then()\n+            .onErrorResume(BlobStorageException.class, exception -\u003e {\n+                if (!BlobEventHandler.isLifecycleRace(exception)) {\n+                    return Mono.error(exception);\n+                }\n+                LOGGER.warning(() -\u003e \"Blob is no longer readable after creation event: \"\n+                    + location.container() + \"/\" + location.name() + \" (\" + exception.getStatusCode() + \")\");\n+                return Mono.empty();\n+            });\n+    }\n+\n+    public Mono\u003cVoid\u003e handleDeleted(BlobLifecycleEvent event) {\n+        BlobEventHandler.BlobLocation location = BlobEventHandler.parseSubject(event.subject());\n+        return Mono.fromRunnable(() -\u003e LOGGER.info(\n+            () -\u003e \"Blob deleted (async): \" + location.container() + \"/\" + location.name()));\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/receiver/EventPayloadParser.java\n+package com.example.blobevents.receiver;\n+\n+import com.azure.core.models.CloudEvent;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.example.blobevents.model.BlobLifecycleEvent;\n+import com.example.blobevents.model.BlobLifecycleEvent.EventSchema;\n+import com.fasterxml.jackson.databind.JsonNode;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+import java.io.IOException;\n+import java.io.UncheckedIOException;\n+import java.util.List;\n+\n+final class EventPayloadParser {\n+    private static final ObjectMapper MAPPER = new ObjectMapper();\n+\n+    private EventPayloadParser() {\n+    }\n+\n+    static List\u003cBlobLifecycleEvent\u003e parse(String payload) {\n+        try {\n+            JsonNode root = MAPPER.readTree(payload);\n+            if (!root.isArray() || root.isEmpty()) {\n+                throw new IllegalArgumentException(\"Event Grid webhook payload must be a non-empty JSON array\");\n+            }\n+            if (root.get(0).has(\"specversion\")) {\n+                return CloudEvent.fromString(payload).stream()\n+                    .map(EventPayloadParser::fromCloudEvent)\n+                    .toList();\n+            }\n+            return EventGridEvent.fromString(payload).stream()\n+                .map(EventPayloadParser::fromEventGridEvent)\n+                .toList();\n+        } catch (IOException exception) {\n+            throw new UncheckedIOException(\"Invalid event JSON payload\", exception);\n+        }\n+    }\n+\n+    private static BlobLifecycleEvent fromCloudEvent(CloudEvent event) {\n+        return new BlobLifecycleEvent(\n+            event.getId(), event.getType(), event.getSubject(), event.getTime(),\n+            event.getData(), EventSchema.CLOUD_EVENTS);\n+    }\n+\n+    private static BlobLifecycleEvent fromEventGridEvent(EventGridEvent event) {\n+        return new BlobLifecycleEvent(\n+            event.getId(), event.getEventType(), event.getSubject(), event.getEventTime(),\n+            event.getData(), EventSchema.EVENT_GRID);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/receiver/EventReceiver.java\n+package com.example.blobevents.receiver;\n+\n+import com.example.blobevents.blob.BlobEventHandler;\n+import com.example.blobevents.model.BlobLifecycleEvent;\n+\n+import java.util.List;\n+import java.util.logging.Logger;\n+\n+public final class EventReceiver {\n+    private static final Logger LOGGER = Logger.getLogger(EventReceiver.class.getName());\n+    private static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    private static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private final BlobEventHandler handler;\n+\n+    public EventReceiver(BlobEventHandler handler) {\n+        this.handler = handler;\n+    }\n+\n+    public List\u003cBlobLifecycleEvent\u003e receive(String jsonPayload) {\n+        List\u003cBlobLifecycleEvent\u003e events = EventPayloadParser.parse(jsonPayload);\n+        events.forEach(this::route);\n+        return events;\n+    }\n+\n+    private void route(BlobLifecycleEvent event) {\n+        switch (event.type()) {\n+            case BLOB_CREATED -\u003e handler.handleCreated(event);\n+            case BLOB_DELETED -\u003e handler.handleDeleted(event);\n+            default -\u003e LOGGER.warning(() -\u003e \"Ignoring unrecognized event type: \" + event.type());\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/receiver/AsyncEventReceiver.java\n+package com.example.blobevents.receiver;\n+\n+import com.example.blobevents.blob.AsyncBlobEventHandler;\n+import com.example.blobevents.model.BlobLifecycleEvent;\n+import reactor.core.publisher.Flux;\n+\n+import java.util.logging.Logger;\n+\n+public final class AsyncEventReceiver {\n+    private static final Logger LOGGER = Logger.getLogger(AsyncEventReceiver.class.getName());\n+    private static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    private static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private final AsyncBlobEventHandler handler;\n+\n+    public AsyncEventReceiver(AsyncBlobEventHandler handler) {\n+        this.handler = handler;\n+    }\n+\n+    public Flux\u003cBlobLifecycleEvent\u003e receive(String jsonPayload) {\n+        return Flux.fromIterable(EventPayloadParser.parse(jsonPayload))\n+            .concatMap(event -\u003e route(event).thenReturn(event));\n+    }\n+\n+    private reactor.core.publisher.Mono\u003cVoid\u003e route(BlobLifecycleEvent event) {\n+        return switch (event.type()) {\n+            case BLOB_CREATED -\u003e handler.handleCreated(event);\n+            case BLOB_DELETED -\u003e handler.handleDeleted(event);\n+            default -\u003e reactor.core.publisher.Mono.fromRunnable(\n+                () -\u003e LOGGER.warning(\"Ignoring unrecognized event type: \" + event.type()));\n+        };\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/publisher/EventPublisher.java\n+package com.example.blobevents.publisher;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import com.example.blobevents.model.CustomEvent;\n+import com.azure.core.credential.TokenCredential;\n+\n+import java.util.List;\n+import java.util.Objects;\n+import java.util.function.Consumer;\n+\n+public final class EventPublisher {\n+    private final Consumer\u003cList\u003cEventGridEvent\u003e\u003e sender;\n+\n+    public EventPublisher(String topicEndpoint, TokenCredential credential) {\n+        EventGridPublisherClient\u003cEventGridEvent\u003e client = new EventGridPublisherClientBuilder()\n+            .endpoint(Objects.requireNonNull(topicEndpoint, \"topicEndpoint\"))\n+            .credential(Objects.requireNonNull(credential, \"credential\"))\n+            .buildEventGridEventPublisherClient();\n+        this.sender = client::sendEvents;\n+    }\n+\n+    public EventPublisher(Consumer\u003cList\u003cEventGridEvent\u003e\u003e sender) {\n+        this.sender = Objects.requireNonNull(sender, \"sender\");\n+    }\n+\n+    public void publish(List\u003cCustomEvent\u003e events) {\n+        sender.accept(events.stream().map(EventPublisher::toEventGridEvent).toList());\n+    }\n+\n+    static EventGridEvent toEventGridEvent(CustomEvent event) {\n+        return new EventGridEvent(\n+            event.subject(), event.type(), BinaryData.fromObject(event.data()), event.dataVersion());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/publisher/AsyncEventPublisher.java\n+package com.example.blobevents.publisher;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import com.example.blobevents.model.CustomEvent;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+import java.util.Objects;\n+import java.util.function.Function;\n+\n+public final class AsyncEventPublisher {\n+    private final Function\u003cList\u003cEventGridEvent\u003e, Mono\u003cVoid\u003e\u003e sender;\n+\n+    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {\n+        EventGridPublisherAsyncClient\u003cEventGridEvent\u003e client = new EventGridPublisherClientBuilder()\n+            .endpoint(Objects.requireNonNull(topicEndpoint, \"topicEndpoint\"))\n+            .credential(Objects.requireNonNull(credential, \"credential\"))\n+            .buildEventGridEventPublisherAsyncClient();\n+        this.sender = client::sendEvents;\n+    }\n+\n+    public AsyncEventPublisher(Function\u003cList\u003cEventGridEvent\u003e, Mono\u003cVoid\u003e\u003e sender) {\n+        this.sender = Objects.requireNonNull(sender, \"sender\");\n+    }\n+\n+    public Mono\u003cVoid\u003e publish(List\u003cCustomEvent\u003e events) {\n+        return sender.apply(events.stream().map(EventPublisher::toEventGridEvent).toList());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureConfiguration.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+\n+import java.util.Map;\n+\n+public final class AzureConfiguration {\n+    private final String storageEndpoint;\n+    private final String topicEndpoint;\n+    private final TokenCredential credential;\n+\n+    public AzureConfiguration(String storageEndpoint, String topicEndpoint, String managedIdentityClientId) {\n+        this.storageEndpoint = requireHttps(storageEndpoint, \"storageEndpoint\");\n+        this.topicEndpoint = requireHttps(topicEndpoint, \"topicEndpoint\");\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            credentialBuilder.clientId(managedIdentityClientId);\n+        }\n+        this.credential = credentialBuilder.build();\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        Map\u003cString, String\u003e environment = System.getenv();\n+        return new AzureConfiguration(\n+            required(environment, \"AZURE_STORAGE_ACCOUNT_URL\"),\n+            required(environment, \"EVENT_GRID_TOPIC_ENDPOINT\"),\n+            environment.get(\"AZURE_CLIENT_ID\"));\n+    }\n+\n+    public BlobServiceClient blobServiceClient() {\n+        return new BlobServiceClientBuilder().endpoint(storageEndpoint).credential(credential).buildClient();\n+    }\n+\n+    public BlobServiceAsyncClient blobServiceAsyncClient() {\n+        return new BlobServiceClientBuilder().endpoint(storageEndpoint).credential(credential).buildAsyncClient();\n+    }\n+\n+    public EventGridPublisherClient\u003cEventGridEvent\u003e eventPublisherClient() {\n+        return new EventGridPublisherClientBuilder().endpoint(topicEndpoint).credential(credential)\n+            .buildEventGridEventPublisherClient();\n+    }\n+\n+    public EventGridPublisherAsyncClient\u003cEventGridEvent\u003e asyncEventPublisherClient() {\n+        return new EventGridPublisherClientBuilder().endpoint(topicEndpoint).credential(credential)\n+            .buildEventGridEventPublisherAsyncClient();\n+    }\n+\n+    public TokenCredential credential() {\n+        return credential;\n+    }\n+\n+    public String topicEndpoint() {\n+        return topicEndpoint;\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e values, String name) {\n+        String value = values.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is missing: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static String requireHttps(String value, String name) {\n+        if (value == null || !value.startsWith(\"https://\")) {\n+            throw new IllegalArgumentException(name + \" must use HTTPS\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/Main.java\n+package com.example.blobevents;\n+\n+import com.azure.core.util.BinaryData;\n+import com.example.blobevents.blob.AsyncBlobEventHandler;\n+import com.example.blobevents.blob.BlobEventHandler;\n+import com.example.blobevents.blob.BlobOperations;\n+import com.example.blobevents.blob.BlobSummary;\n+import com.example.blobevents.model.CustomEvent;\n+import com.example.blobevents.publisher.AsyncEventPublisher;\n+import com.example.blobevents.publisher.EventPublisher;\n+import com.example.blobevents.receiver.AsyncEventReceiver;\n+import com.example.blobevents.receiver.EventReceiver;\n+import reactor.core.publisher.Mono;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.time.OffsetDateTime;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.logging.Logger;\n+\n+public final class Main {\n+    private static final Logger LOGGER = Logger.getLogger(Main.class.getName());\n+\n+    private static final String EVENT_GRID_PAYLOAD = \"\"\"\n+        [\n+          {\n+            \"topic\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/invoices%2Finvoice-1001.pdf\",\n+            \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+            \"id\": \"11111111-1111-1111-1111-111111111111\",\n+            \"data\": {\n+              \"api\": \"PutBlob\",\n+              \"clientRequestId\": \"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa\",\n+              \"requestId\": \"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb\",\n+              \"eTag\": \"0x8DB123456789ABC\",\n+              \"contentType\": \"application/pdf\",\n+              \"contentLength\": 2048,\n+              \"blobType\": \"BlockBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/invoices/invoice-1001.pdf\",\n+              \"sequencer\": \"000000000000000000000000000000010000000000000001\"\n+            },\n+            \"dataVersion\": \"\",\n+            \"metadataVersion\": \"1\",\n+            \"eventTime\": \"2026-08-28T01:30:00Z\"\n+          },\n+          {\n+            \"topic\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/archive%2Fold-invoice.pdf\",\n+            \"eventType\": \"Microsoft.Storage.BlobDeleted\",\n+            \"id\": \"22222222-2222-2222-2222-222222222222\",\n+            \"data\": {\n+              \"api\": \"DeleteBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/archive/old-invoice.pdf\",\n+              \"sequencer\": \"000000000000000000000000000000020000000000000001\"\n+            },\n+            \"dataVersion\": \"\",\n+            \"metadataVersion\": \"1\",\n+            \"eventTime\": \"2026-08-28T01:31:00Z\"\n+          }\n+        ]\n+        \"\"\";\n+\n+    private static final String CLOUD_EVENTS_PAYLOAD = \"\"\"\n+        [\n+          {\n+            \"specversion\": \"1.0\",\n+            \"type\": \"Microsoft.Storage.BlobCreated\",\n+            \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/reports%2Fquarterly.csv\",\n+            \"id\": \"33333333-3333-3333-3333-333333333333\",\n+            \"time\": \"2026-08-28T01:32:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+              \"api\": \"PutBlob\",\n+              \"contentType\": \"text/csv\",\n+              \"contentLength\": 512,\n+              \"blobType\": \"BlockBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/reports/quarterly.csv\"\n+            }\n+          },\n+          {\n+            \"specversion\": \"1.0\",\n+            \"type\": \"Microsoft.Storage.BlobDeleted\",\n+            \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/reports%2Fold.csv\",\n+            \"id\": \"44444444-4444-4444-4444-444444444444\",\n+            \"time\": \"2026-08-28T01:33:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+              \"api\": \"DeleteBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/reports/old.csv\"\n+            }\n+          }\n+        ]\n+        \"\"\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        BlobOperations demoBlobs = new DemoBlobOperations();\n+        CustomEvent notification = new CustomEvent(\n+            \"Contoso.Documents.Processed\",\n+            \"/documents/invoices/processed\",\n+            \"1.0\",\n+            Map.of(\"documentId\", \"invoice-1001\", \"processedAt\", OffsetDateTime.now().toString()));\n+\n+        LOGGER.info(\"=== Synchronous demo ===\");\n+        EventReceiver receiver = new EventReceiver(new BlobEventHandler(demoBlobs));\n+        receiver.receive(EVENT_GRID_PAYLOAD);\n+        receiver.receive(CLOUD_EVENTS_PAYLOAD);\n+        new EventPublisher(events -\u003e events.forEach(event -\u003e\n+            LOGGER.info(() -\u003e \"Would publish sync event: type=\" + event.getEventType()\n+                + \", subject=\" + event.getSubject())))\n+            .publish(List.of(notification));\n+\n+        LOGGER.info(\"=== Asynchronous demo ===\");\n+        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver(new AsyncBlobEventHandler(demoBlobs));\n+        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -\u003e Mono.fromRunnable(() -\u003e\n+            events.forEach(event -\u003e LOGGER.info(() -\u003e \"Would publish async event: type=\"\n+                + event.getEventType() + \", subject=\" + event.getSubject()))));\n+\n+        asyncReceiver.receive(EVENT_GRID_PAYLOAD)\n+            .thenMany(asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD))\n+            .then(asyncPublisher.publish(List.of(notification)))\n+            .block();\n+    }\n+\n+    private static final class DemoBlobOperations implements BlobOperations {\n+        @Override\n+        public DownloadedBlob download(String container, String name) {\n+            byte[] content = (\"mock content for \" + container + \"/\" + name)\n+                .getBytes(StandardCharsets.UTF_8);\n+            String contentType = name.endsWith(\".pdf\") ? \"application/pdf\" : \"text/csv\";\n+            return new DownloadedBlob(\n+                BinaryData.fromBytes(content),\n+                new BlobSummary(name, content.length, contentType, \"HOT\"));\n+        }\n+\n+        @Override\n+        public Mono\u003cDownloadedBlob\u003e downloadAsync(String container, String name) {\n+            return Mono.fromSupplier(() -\u003e download(container, name));\n+        }\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/receiver/EventReceiverTest.java\n+package com.example.blobevents.receiver;\n+\n+import com.azure.core.util.BinaryData;\n+import com.example.blobevents.blob.BlobEventHandler;\n+import com.example.blobevents.blob.BlobOperations;\n+import com.example.blobevents.blob.BlobSummary;\n+import org.junit.jupiter.api.Test;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.concurrent.atomic.AtomicReference;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class EventReceiverTest {\n+    @Test\n+    void parsesCloudEventAndDecodesBlobPath() {\n+        AtomicReference\u003cString\u003e requestedBlob = new AtomicReference\u003c\u003e();\n+        BlobOperations operations = new BlobOperations() {\n+            @Override\n+            public DownloadedBlob download(String container, String name) {\n+                requestedBlob.set(container + \"/\" + name);\n+                return new DownloadedBlob(BinaryData.fromString(\"x\"),\n+                    new BlobSummary(name, 1, \"text/plain\", \"HOT\"));\n+            }\n+\n+            @Override\n+            public Mono\u003cDownloadedBlob\u003e downloadAsync(String container, String name) {\n+                return Mono.fromSupplier(() -\u003e download(container, name));\n+            }\n+        };\n+        String payload = \"\"\"\n+            [{\n+              \"specversion\":\"1.0\",\n+              \"type\":\"Microsoft.Storage.BlobCreated\",\n+              \"source\":\"/storage\",\n+              \"subject\":\"/blobServices/default/containers/docs/blobs/folder%2Ffile+one.txt\",\n+              \"id\":\"1\",\n+              \"time\":\"2026-08-28T00:00:00Z\",\n+              \"datacontenttype\":\"application/json\",\n+              \"data\":{\"url\":\"https://example.blob.core.windows.net/docs/folder/file+one.txt\"}\n+            }]\n+            \"\"\";\n+\n+        new EventReceiver(new BlobEventHandler(operations)).receive(payload);\n+\n+        assertEquals(\"docs/folder/file+one.txt\", requestedBlob.get());\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-2414409379/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-2414409379/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-microsoft-skill-2414409379/pom.xml
@@ -1,0 +1,80 @@
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
+        <azure.eventgrid.version>4.27.0</azure.eventgrid.version>
+        <azure.identity.version>1.18.2</azure.identity.version>
+        <azure.storage.blob.version>12.33.0</azure.storage.blob.version>
+        <jackson.version>2.17.2</jackson.version>
+        <junit.version>5.11.0</junit.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-eventgrid</artifactId>
+            <version>${azure.eventgrid.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-storage-blob</artifactId>
+            <version>${azure.storage.blob.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+
... (truncated)
```

### ✅ 10. powershell — 23600ms

**Input:**

```json
{"command":"mvn --quiet test \u0026\u0026 mvn --quiet exec:java","description":"Compile, test, and run the Java demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
Aug 28, 2026 3:42:24 AM com.example.blobevents.blob.BlobEventHandler handleCreated
INFO: Blob created: name=folder/file+one.txt, size=1, contentType=text/plain, accessTier=HOT
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
Aug 28, 2026 3:42:27 AM com.example.blobevents.Main main
INFO: === Synchronous demo ===
Aug 28, 2026 3:42:28 AM com.example.blobevents.blob.BlobEventHandler handleCreated
INFO: Blob created: name=invoices/invoice-1001.pdf, size=52, contentType=application/pdf, accessTier=HOT
Aug 28, 2026 3:42:28 AM com.example.blobevents.blob.BlobEventHandler handleDeleted
INFO: Blob deleted: documents/archive/old-invoice.pdf
Aug 28, 2026 3:42:28 AM com.example.blobevents.blob.BlobEventHandler handleCreated
INFO: Blob created: name=reports/quarterly.csv, size=48, contentType=text/csv, accessTier=HOT
Aug 28, 2026 3:42:28 AM com.example.blobevents.blob.BlobEventHandler handleDeleted
INFO: Blob deleted: documents/reports/old.csv
Aug 28, 2026 3:42:28 AM com.example.blobevents.Main lambda$main$1
INFO: Would publish sync event: type=Contoso.Documents.Processed, subject=/documents/invoices/processed
Aug 28, 2026 3:42:28 AM com.example.blobevents.Main m
... (truncated)
```

## Generated Files

- `pom.xml`
- `src\main\java\com\example\blobevents\AzureConfiguration.java`
- `src\main\java\com\example\blobevents\Main.java`
- `src\main\java\com\example\blobevents\blob\AsyncBlobEventHandler.java`
- `src\main\java\com\example\blobevents\blob\AzureBlobOperations.java`
- `src\main\java\com\example\blobevents\blob\BlobEventHandler.java`
- `src\main\java\com\example\blobevents\blob\BlobOperations.java`
- `src\main\java\com\example\blobevents\blob\BlobSummary.java`
- `src\main\java\com\example\blobevents\model\BlobLifecycleEvent.java`
- `src\main\java\com\example\blobevents\model\CustomEvent.java`
- `src\main\java\com\example\blobevents\publisher\AsyncEventPublisher.java`
- `src\main\java\com\example\blobevents\publisher\EventPublisher.java`
- `src\main\java\com\example\blobevents\receiver\AsyncEventReceiver.java`
- `src\main\java\com\example\blobevents\receiver\EventPayloadParser.java`
- `src\main\java\com\example\blobevents\receiver\EventReceiver.java`
- `src\test\java\com\example\blobevents\receiver\EventReceiverTest.java`

## Copilot Response

I’m aligning the project with the current Azure SDK patterns, then I’ll scaffold the implementation and validate it locally without contacting Azure.

The demo will use in-memory blob and publisher adapters so it is fully runnable offline, while the production constructors use managed identity and real Azure SDK clients.

The project is scaffolded with real managed-identity adapters and offline demo doubles; I’m now resolving any SDK API mismatches exposed by compilation.

Created the Java 17 Maven project with:

- Sync/async Event Grid and CloudEvents receivers
- Blob-created/deleted routing and lifecycle race handling
- Sync/async Blob Storage handlers
- Sync/async custom Event Grid publishers with hierarchical subjects
- Managed identity configuration using `ManagedIdentityCredential`
- Offline demo using realistic mock events and in-memory adapters
- Unit coverage for CloudEvents parsing and encoded blob paths

Run the demo with:

```powershell
mvn compile exec:java
```

Production configuration reads `AZURE_STORAGE_ACCOUNT_URL`, `EVENT_GRID_TOPIC_ENDPOINT`, and optional `AZURE_CLIENT_ID`. No access keys, connection strings, or SAS tokens are used.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK client is left unclosed.

## Grader Results

- blob-event-notifier.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (8/10)
      - Handles Event Grid native schema via `EventGridEvent.fromString()` deserialization: Pass
      - Handles CloudEvents 1.0 schema via `CloudEvent.fromString()` deserialization: Pass
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
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 80% | 1.00 | 0.8000 | 6.8% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.8000** | **90.8%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-blob-event-notifier --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

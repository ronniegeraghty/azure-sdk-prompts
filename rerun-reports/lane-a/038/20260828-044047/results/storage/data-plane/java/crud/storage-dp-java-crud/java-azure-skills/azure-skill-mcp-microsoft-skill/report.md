# Evaluation Report: storage-dp-java-crud

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 505.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-crud` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 17/19 |
| Duration | 505.7s |
| Timestamp | 2026-08-27T20:40:47Z |
| Files Generated | 3 |
| Event Count | 4028 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 293.7s |
| Review | 210.4s |
| **Total** | **505.7s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage-blob-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=198390 out=3041 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 3400ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 188098ms

**Input:**

```json
{"intent":"Generate a local Java Maven application using Azure Blob Storage SDK and DefaultAzureCredential, with CRUD operations and BlobStorageException handling. Do not deploy or modify Azure resources during development.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. view — 188093ms

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

### ✅ 4. glob — 186708ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-855555952","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. apply_patch — 2479ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-blob-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e12.33.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.17.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.BlobStorageCrudApp\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/BlobStorageCrudApp.java\n+package com.example;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+import com.azure.storage.blob.models.BlobItem;\n+import com.azure.storage.blob.models.BlobStorageException;\n+\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+\n+public final class BlobStorageCrudApp {\n+    private static final String CONTAINER_NAME = \"my-container\";\n+    private static final String BLOB_NAME = \"uploads/data.txt\";\n+    private static final Path SOURCE_FILE = Path.of(\"data.txt\");\n+    private static final Path DOWNLOADED_FILE = Path.of(\"data-downloaded.txt\");\n+\n+    private BlobStorageCrudApp() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String accountUrl = System.getenv(\"AZURE_STORAGE_ACCOUNT_URL\");\n+        if (accountUrl == null || accountUrl.isBlank()) {\n+            System.err.println(\n+                \"AZURE_STORAGE_ACCOUNT_URL must be set, for example \"\n+                    + \"https://\u003caccount-name\u003e.blob.core.windows.net.\");\n+            System.exit(1);\n+        }\n+\n+        if (!Files.isRegularFile(SOURCE_FILE)) {\n+            System.err.println(\"Source file does not exist: \" + SOURCE_FILE.toAbsolutePath());\n+            System.exit(1);\n+        }\n+\n+        try {\n+            runCrudOperations(accountUrl);\n+        } catch (BlobStorageException exception) {\n+            System.err.printf(\n+                \"Azure Blob Storage request failed (status=%d, errorCode=%s): %s%n\",\n+                exception.getStatusCode(),\n+                exception.getErrorCode(),\n+                exception.getServiceMessage());\n+            System.exit(1);\n+        } catch (ClientAuthenticationException exception) {\n+            System.err.println(\"Azure authentication failed: \" + exception.getMessage());\n+            System.exit(1);\n+        } catch (IOException exception) {\n+            System.err.println(\"Local file operation failed: \" + exception.getMessage());\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static void runCrudOperations(String accountUrl) throws IOException {\n+        BlobServiceClient serviceClient = new BlobServiceClientBuilder()\n+            .endpoint(accountUrl)\n+            .credential(new DefaultAzureCredentialBuilder().build())\n+            .buildClient();\n+\n+        BlobContainerClient containerClient =\n+            serviceClient.getBlobContainerClient(CONTAINER_NAME);\n+        boolean containerCreated = containerClient.createIfNotExists();\n+        System.out.printf(\n+            \"Container '%s' %s.%n\",\n+            CONTAINER_NAME,\n+            containerCreated ? \"created\" : \"already exists\");\n+\n+        BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);\n+        blobClient.uploadFromFile(SOURCE_FILE.toString(), true);\n+        System.out.printf(\"Uploaded '%s' as '%s'.%n\", SOURCE_FILE, BLOB_NAME);\n+\n+        System.out.println(\"Blobs in container:\");\n+        for (BlobItem item : containerClient.listBlobs()) {\n+            Long size = item.getProperties() == null\n+                ? null\n+                : item.getProperties().getContentLength();\n+            System.out.printf(\n+                \"  %s (%s bytes)%n\",\n+                item.getName(),\n+                size == null ? \"unknown\" : size);\n+        }\n+\n+        Files.deleteIfExists(DOWNLOADED_FILE);\n+        blobClient.downloadToFile(DOWNLOADED_FILE.toString());\n+        System.out.printf(\"Downloaded '%s' to '%s'.%n\", BLOB_NAME, DOWNLOADED_FILE);\n+\n+        blobClient.delete();\n+        System.out.printf(\"Deleted blob '%s'.%n\", BLOB_NAME);\n+\n+        containerClient.delete();\n+        System.out.printf(\"Deleted container '%s'.%n\", CONTAINER_NAME);\n+    }\n+}\n*** Add File: README.md\n+# Azure Blob Storage CRUD with Java\n+\n+This sample creates a container, uploads and lists a blob, downloads it, and\n+then deletes the blob and container. It uses `DefaultAzureCredential`; no\n+credentials are stored in source code.\n+\n+## Prerequisites\n+\n+- Java 17 or later\n+- Maven 3.9 or later\n+- An existing Azure Storage account\n+- A local `data.txt` file in the project directory\n+- An authenticated identity with permission to manage blobs and containers,\n+  such as the **Storage Blob Data Contributor** role\n+\n+Set the storage account Blob service URL:\n+\n+```powershell\n+$env:AZURE_STORAGE_ACCOUNT_URL = \"https://\u003caccount-name\u003e.blob.core.windows.net\"\n+```\n+\n+For local development, `DefaultAzureCredential` can use environment\n+credentials or an existing developer-tool login. For service principal\n+authentication, set `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and\n+`AZURE_CLIENT_SECRET`.\n+\n+## Build and run\n+\n+```powershell\n+mvn compile\n+mvn exec:java\n+```\n+\n+The application overwrites `uploads/data.txt` and removes any existing local\n+`data-downloaded.txt`. It then deletes `my-container`, so use a disposable\n+container and do not store unrelated blobs in it.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-855555952/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-855555952/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-855555952/pom.xml
@@ -1,0 +1,47 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>azure-blob-crud</artifactId>
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
+            <artifactId>azure-storage-blob</artifactId>
+            <version>12.33.0</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>1.17.0</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-compiler-plugin</artifactId>
+                <version>3.14.1</version>
+            </plugin>
+            <plugin>
+                <groupId>org.codehaus.mojo</groupId>
+                <artifactId>exec-maven-plugin</artifactId>
+                <version>3.5.1</version>
+                <configuration>
+                    <mainClass>com.example.Blob
... (truncated)
```

### ✅ 6. powershell — 5758ms

**Input:**

```json
{"command":"mvn --quiet compile","description":"Compile the Maven application","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\BlobStorageCrudApp.java`

## Copilot Response

I’ll inspect the starter workspace and align the implementation with the current Azure Java SDK patterns before adding the CRUD application.

The workspace is empty, so I’m creating a minimal Maven project with environment-based configuration, explicit Azure Identity and Blob dependencies, and operation-specific failure reporting.

Implemented a complete Java 17 Maven application for Azure Blob Storage CRUD.

- `pom.xml` — Azure Blob Storage and Identity dependencies
- `src\main\java\com\example\BlobStorageCrudApp.java` — creates the client/container, uploads, lists names and sizes, downloads, and deletes resources
- `README.md` — authentication, permissions, configuration, and run instructions

Set `AZURE_STORAGE_ACCOUNT_URL`, create `data.txt`, then run:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
mvn compile
mvn exec:java
```

The application handles `BlobStorageException`, authentication failures, missing configuration, missing input files, and local I/O failures. It deletes `my-container` at the end, so that container should not contain unrelated data.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Correctly avoids unnecessary resource-management constructs for Azure Blob Storage synchronous clients that are not AutoCloseable.

## Grader Results

- crud-blobs.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - Maven dependency for `azure-storage-blob` and `azure-identity`: Pass
      - `BlobServiceClientBuilder` with `DefaultAzureCredential`: Pass
      - `BlobContainerClient.create()` and `exists()` check: Fail
      - `BlobClient.uploadFromFile()` and `downloadToFile()`: Pass
      - `BlobContainerClient.listBlobs()` iteration: Pass
      - `BlobClient.delete()` and `BlobContainerClient.delete()`: Pass
      - `BlobStorageException` handling with status codes: Pass
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 7.2% | ❌ |
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
| **Final** | | | **Σ 13.00** | **Σ 11.8571** | **91.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-crud --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

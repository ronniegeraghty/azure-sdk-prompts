# Evaluation Report: key-vault-dp-java-crud

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 519.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-crud` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 519.6s |
| Timestamp | 2026-08-27T17:29:28Z |
| Files Generated | 3 |
| Event Count | 3430 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 272.0s |
| Review | 246.3s |
| **Total** | **519.6s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-security-keyvault-secrets-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=198953 out=2382 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 2333ms

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

### ✅ 2. view — 188295ms

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

### ✅ 3. glob — 187261ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1022778379","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 188297ms

**Input:**

```json
{"intent":"Generate a local Java Maven application using Azure Key Vault Secrets and DefaultAzureCredential to create, read, update, delete, and purge a secret, with proper exception handling. Do not deploy or modify Azure resources during development.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. apply_patch — 3346ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-secret-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+            \u003cversion\u003e4.9.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.18.2\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.KeyVaultSecretCrud\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/KeyVaultSecretCrud.java\n+package com.example;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.core.util.polling.SyncPoller;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+import com.azure.security.keyvault.secrets.models.DeletedSecret;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+public final class KeyVaultSecretCrud {\n+    private static final String SECRET_NAME = \"my-secret\";\n+    private static final String INITIAL_VALUE = \"my-secret-value\";\n+    private static final String UPDATED_VALUE = \"updated-value\";\n+\n+    private KeyVaultSecretCrud() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String vaultUrl = System.getenv(\"AZURE_KEYVAULT_URL\");\n+        if (vaultUrl == null || vaultUrl.isBlank()) {\n+            System.err.println(\n+                \"AZURE_KEYVAULT_URL must be set, for example: https://\u003cvault-name\u003e.vault.azure.net\");\n+            System.exit(2);\n+        }\n+\n+        try {\n+            SecretClient secretClient = new SecretClientBuilder()\n+                .vaultUrl(vaultUrl)\n+                .credential(new DefaultAzureCredentialBuilder().build())\n+                .buildClient();\n+\n+            KeyVaultSecret created = secretClient.setSecret(SECRET_NAME, INITIAL_VALUE);\n+            System.out.printf(\"Created secret \\\"%s\\\" (version %s).%n\",\n+                created.getName(), created.getProperties().getVersion());\n+\n+            KeyVaultSecret read = secretClient.getSecret(SECRET_NAME);\n+            System.out.printf(\"Read secret value: %s%n\", read.getValue());\n+\n+            // Secret values are immutable; setting the same name creates a new version.\n+            KeyVaultSecret updated = secretClient.setSecret(SECRET_NAME, UPDATED_VALUE);\n+            System.out.printf(\"Updated secret \\\"%s\\\" to version %s.%n\",\n+                updated.getName(), updated.getProperties().getVersion());\n+\n+            SyncPoller\u003cDeletedSecret, Void\u003e deletePoller =\n+                secretClient.beginDeleteSecret(SECRET_NAME);\n+            deletePoller.waitForCompletion();\n+            System.out.printf(\"Deleted secret \\\"%s\\\".%n\", SECRET_NAME);\n+\n+            secretClient.purgeDeletedSecret(SECRET_NAME);\n+            System.out.printf(\"Purged secret \\\"%s\\\" permanently.%n\", SECRET_NAME);\n+        } catch (ClientAuthenticationException e) {\n+            System.err.println(\"Authentication failed. Check the credentials available to \"\n+                + \"DefaultAzureCredential: \" + e.getMessage());\n+            System.exit(1);\n+        } catch (ResourceNotFoundException e) {\n+            System.err.println(\"The requested vault or secret was not found: \" + e.getMessage());\n+            System.exit(1);\n+        } catch (HttpResponseException e) {\n+            int statusCode = e.getResponse() == null ? -1 : e.getResponse().getStatusCode();\n+            System.err.printf(\"Azure Key Vault request failed (HTTP %d): %s%n\",\n+                statusCode, e.getMessage());\n+            if (statusCode == 403) {\n+                System.err.println(\"Ensure the identity has secret get, set, delete, and purge permissions. \"\n+                    + \"Purge also fails when purge protection is enabled.\");\n+            } else if (statusCode == 429) {\n+                System.err.println(\"The vault is throttling requests; retry after the server's delay.\");\n+            }\n+            System.exit(1);\n+        } catch (IllegalArgumentException e) {\n+            System.err.println(\"Invalid Key Vault configuration: \" + e.getMessage());\n+            System.exit(2);\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault secret CRUD\n+\n+This Maven application uses `DefaultAzureCredential` to create, read, update,\n+soft-delete, and purge the `my-secret` secret. Updating a secret value creates a\n+new version because existing secret values are immutable.\n+\n+Set the vault URL and provide any credential supported by\n+`DefaultAzureCredential`:\n+\n+```powershell\n+$env:AZURE_KEYVAULT_URL = \"https://\u003cvault-name\u003e.vault.azure.net\"\n+mvn compile exec:java\n+```\n+\n+The authenticated identity needs permissions to get, set, delete, and purge\n+secrets. Purging is not possible when purge protection is enabled.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1022778379/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1022778379/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1022778379/pom.xml
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
+            <version>4.9.0</version>
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
+                <version>3.5.0</version>
+                <configuration>
+                    <mainClass>com.example.KeyVaultSecretCrud</mainClass>
+                </configuration>
+            </plugin>
+        </plugins>
+    </build>
+</project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyok
... (truncated)
```

### ✅ 6. powershell — 11229ms

**Input:**

```json
{"command":"mvn --quiet compile","description":"Compile the Maven application","initial_wait":120,"mode":"sync"}
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
- `src\main\java\com\example\KeyVaultSecretCrud.java`

## Copilot Response

I’ll inspect the existing project structure and align the implementation with the current Azure Java SDK guidance before adding the CRUD application.

The workspace is empty, so I’m creating a minimal Maven project. The update will correctly use `setSecret` to create a new secret version, then wait for soft deletion to finish before purging.

Implemented a complete Java Maven application for Azure Key Vault secret CRUD.

- `pom.xml` includes `azure-security-keyvault-secrets` and `azure-identity`.
- `KeyVaultSecretCrud.java` creates, reads, updates, soft-deletes, and purges `my-secret`.
- Uses `DefaultAzureCredential` and `AZURE_KEYVAULT_URL`.
- Handles authentication, authorization, throttling, missing resources, and invalid configuration.
- Waits for deletion to complete before purging.
- `README.md` includes configuration and run instructions.

## Reviewer Notes (LLM-as-Judge)

### Summary

The sole criterion passes because no AutoCloseable Azure SDK client instance is used.

### Strengths

- Uses the synchronous SecretClient appropriately without unnecessary or invalid resource-closing logic.

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
hyoka run --prompt-id key-vault-dp-java-crud --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

# Evaluation Report: key-vault-dp-java-crud

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 437.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-crud` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 437.6s |
| Timestamp | 2026-08-26T12:48:05Z |
| Files Generated | 3 |
| Event Count | 4523 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 161.0s |
| Review | 275.3s |
| **Total** | **437.6s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-tools/with-azure-tools

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-security-keyvault-secrets-java, azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=418480 out=4158 |
| Turn Count | 9 |

## Prompt Sent

```
Write a Java application that performs
all four CRUD operations on Azure Key Vault secrets:
1. Create a new secret called "my-secret" with value "my-secret-value"
2. Read the secret back and print its value
3. Update the secret to a new value "updated-value"
4. Delete the secret and purge it (soft-delete enabled vault)

Use DefaultAzureCredential for authentication. Show the Maven dependency
for azure-security-keyvault-secrets and azure-identity. Include proper exception handling.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-security-keyvault-secrets-java"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local-only Java Maven console application demonstrating CRUD operations for Azure Key Vault secrets with DefaultAzureCredential, robust exception handling, and no Azure resource provisioning.","learn":true}
```

### 🔧 3. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java\\references\\examples.md"}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java","pattern":"**\\acceptance-criteria.md"}
```

### 🔧 5. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-java-crud-java-azure-tools-with-azure-tools-2835151058","pattern":"**\\*"}
```

### 🔧 6. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 7. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java SDK code generation best practices for a local Maven console app using DefaultAzureCredential and Azure Key Vault Secrets CRUD.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 8. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 9. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java","pattern":"**\\acceptance-criteria.md"}
```

### 🔧 10. web_search (via github-mcp-server)

**Input:**

```json
{"query":"As of August 2026, what are the latest stable Maven Central versions of com.azure:azure-security-keyvault-secrets and com.azure:azure-identity? Cite Maven Central or official Microsoft sources."}
```

### 🔧 11. web_fetch

**Input:**

```json
{"max_length":10000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml"}
```

### 🔧 12. web_fetch

**Input:**

```json
{"max_length":10000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

### 🔧 13. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-secret-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+            \u003cversion\u003e4.11.2\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.18.5\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.6.2\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.KeyVaultSecretCrudApp\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/KeyVaultSecretCrudApp.java\n+package com.example;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.core.util.polling.SyncPoller;\n+import com.azure.identity.CredentialUnavailableException;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+import com.azure.security.keyvault.secrets.models.DeletedSecret;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+public final class KeyVaultSecretCrudApp {\n+    private static final String VAULT_URL_ENVIRONMENT_VARIABLE = \"AZURE_KEY_VAULT_URL\";\n+    private static final String SECRET_NAME = \"my-secret\";\n+\n+    private KeyVaultSecretCrudApp() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String vaultUrl = System.getenv(VAULT_URL_ENVIRONMENT_VARIABLE);\n+        if (vaultUrl == null || vaultUrl.isBlank()) {\n+            System.err.printf(\n+                \"Set %s to a vault URL such as https://\u003cvault-name\u003e.vault.azure.net/.%n\",\n+                VAULT_URL_ENVIRONMENT_VARIABLE);\n+            System.exit(2);\n+        }\n+\n+        SecretClient secretClient = new SecretClientBuilder()\n+            .vaultUrl(vaultUrl)\n+            .credential(new DefaultAzureCredentialBuilder().build())\n+            .buildClient();\n+\n+        try {\n+            runCrudOperations(secretClient);\n+        } catch (CredentialUnavailableException exception) {\n+            System.err.println(\"No credential was available to DefaultAzureCredential: \"\n+                + exception.getMessage());\n+            System.exit(1);\n+        } catch (ClientAuthenticationException exception) {\n+            System.err.println(\"Azure authentication failed: \" + exception.getMessage());\n+            System.exit(1);\n+        } catch (ResourceNotFoundException exception) {\n+            System.err.printf(\"Secret \\\"%s\\\" was not found: %s%n\",\n+                SECRET_NAME, exception.getMessage());\n+            System.exit(1);\n+        } catch (HttpResponseException exception) {\n+            handleHttpError(exception);\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static void runCrudOperations(SecretClient secretClient) {\n+        KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, \"my-secret-value\");\n+        System.out.printf(\"Created secret \\\"%s\\\" (version %s).%n\",\n+            createdSecret.getName(), createdSecret.getProperties().getVersion());\n+\n+        KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);\n+        System.out.printf(\"Read secret value: %s%n\", readSecret.getValue());\n+\n+        KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, \"updated-value\");\n+        System.out.printf(\"Updated secret \\\"%s\\\" to a new version (%s).%n\",\n+            updatedSecret.getName(), updatedSecret.getProperties().getVersion());\n+\n+        SyncPoller\u003cDeletedSecret, Void\u003e deletePoller =\n+            secretClient.beginDeleteSecret(SECRET_NAME);\n+        deletePoller.waitForCompletion();\n+        System.out.printf(\"Deleted secret \\\"%s\\\".%n\", SECRET_NAME);\n+\n+        secretClient.purgeDeletedSecret(SECRET_NAME);\n+        System.out.printf(\"Purged secret \\\"%s\\\" permanently.%n\", SECRET_NAME);\n+    }\n+\n+    private static void handleHttpError(HttpResponseException exception) {\n+        int statusCode = exception.getResponse() == null\n+            ? -1\n+            : exception.getResponse().getStatusCode();\n+\n+        switch (statusCode) {\n+            case 403 -\u003e\n+                System.err.println(\"Access denied. Grant the identity permission to set, get, \"\n+                    + \"delete, and purge secrets.\");\n+            case 409 -\u003e\n+                System.err.println(\"The operation conflicted with the vault state. A previously \"\n+                    + \"deleted secret may still exist, or purge protection may be enabled.\");\n+            case 429 -\u003e\n+                System.err.println(\"Key Vault throttled the request after SDK retries. Try again later.\");\n+            default -\u003e\n+                System.err.printf(\"Key Vault request failed with HTTP status %d: %s%n\",\n+                    statusCode, exception.getMessage());\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault secret CRUD with Java\n+\n+This Maven console application creates `my-secret`, reads and prints its value,\n+creates a new version with the value `updated-value`, then soft-deletes and\n+purges the secret.\n+\n+## Prerequisites\n+\n+- Java 17 and Maven\n+- An existing Azure Key Vault with soft delete enabled and purge protection\n+  disabled\n+- An identity authorized to set, get, delete, and purge secrets\n+- Local authentication available to `DefaultAzureCredential`, or managed\n+  identity when hosted in Azure\n+\n+Set the vault URL and run the application from PowerShell:\n+\n+```powershell\n+$env:AZURE_KEY_VAULT_URL = \"https://\u003cvault-name\u003e.vault.azure.net/\"\n+mvn compile exec:java\n+```\n+\n+`DefaultAzureCredential` supports development credentials and managed identity\n+without storing credentials in source code. Purging is permanent and requires\n+the Key Vault purge permission.\n+\n+References:\n+\n+- [Azure Key Vault Secrets client library for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme)\n+- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)\n*** End Patch\n"
```

### 🔧 14. powershell

**Input:**

```json
{"command":"$mvn = Get-Command mvn -ErrorAction SilentlyContinue; $java = Get-Command java -ErrorAction SilentlyContinue; if (-not $mvn) { Write-Error 'Maven is not installed or not on PATH.'; exit 127 }; if (-not $java) { Write-Error 'Java is not installed or not on PATH.'; exit 127 }; mvn --quiet --batch-mode package","description":"Compile the Maven application","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\KeyVaultSecretCrudApp.java`

## Copilot Response

I’ll scaffold a local Maven console app using the SDK’s synchronous delete poller so purge only starts after soft deletion completes.

The app will take the vault URL from `AZURE_KEY_VAULT_URL`; credentials and secret values remain out of source configuration except for the explicitly requested demo values.

I’m using the latest stable Maven Central releases and a Java 17 build. Updating a Key Vault secret value correctly uses `setSecret` again, which creates a new version because metadata updates cannot change secret values.

Created a complete Maven Java 17 application:

- `pom.xml` — includes `azure-security-keyvault-secrets:4.11.2` and `azure-identity:1.18.5`
- `src\main\java\com\example\KeyVaultSecretCrudApp.java` — create, read, update via a new secret version, soft-delete, and purge
- `README.md` — configuration, permissions, and execution instructions

Run with:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
mvn compile exec:java
```

The application uses `DefaultAzureCredential`, waits for deletion to complete before purging, and handles authentication, authorization, missing-secret, conflict, and throttling errors. Purge protection must be disabled and the identity needs set, get, delete, and purge permissions.

References: [Key Vault Secrets SDK](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme) · [Azure Identity SDK](https://learn.microsoft.com/java/api/overview/azure/identity-readme)

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The Azure Key Vault SecretClient is managed appropriately for its lifecycle contract.

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
hyoka run --prompt-id key-vault-dp-java-crud --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)

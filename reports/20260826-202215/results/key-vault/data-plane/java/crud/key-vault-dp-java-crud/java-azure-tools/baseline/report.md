# Evaluation Report: key-vault-dp-java-crud

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 332.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-crud` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 332.1s |
| Timestamp | 2026-08-26T12:42:33Z |
| Files Generated | 2 |
| Event Count | 3227 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 86.1s |
| Review | 244.7s |
| **Total** | **332.1s** |

## Configuration

- **name:** java-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=165776 out=2122 |
| Turn Count | 7 |

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

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Latest stable Maven versions azure-security-keyvault-secrets azure-identity August 2026 Maven Central"}
```

### 🔧 3. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml"}
```

### 🔧 5. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-secret-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+            \u003cversion\u003e4.11.2\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.18.5\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.KeyVaultSecretCrud\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/KeyVaultSecretCrud.java\n+package com.example;\n+\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.util.polling.SyncPoller;\n+import com.azure.identity.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+import com.azure.identity.DefaultAzureCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+import com.azure.security.keyvault.secrets.models.DeletedSecret;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+public final class KeyVaultSecretCrud {\n+    private static final String SECRET_NAME = \"my-secret\";\n+\n+    private KeyVaultSecretCrud() {\n+    }\n+\n+    public static void main(String[] args) {\n+        try {\n+            String vaultUrl = requireEnvironmentVariable(\"AZURE_KEY_VAULT_URL\");\n+            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();\n+            SecretClient secretClient = new SecretClientBuilder()\n+                    .vaultUrl(vaultUrl)\n+                    .credential(credential)\n+                    .buildClient();\n+\n+            runCrudOperations(secretClient);\n+        } catch (CredentialUnavailableException exception) {\n+            System.err.println(\"No supported Azure credential was available: \" + exception.getMessage());\n+            System.exit(1);\n+        } catch (ClientAuthenticationException exception) {\n+            System.err.println(\"Azure authentication failed: \" + exception.getMessage());\n+            System.exit(1);\n+        } catch (HttpResponseException exception) {\n+            System.err.printf(\n+                    \"Key Vault request failed with status %d: %s%n\",\n+                    exception.getResponse().getStatusCode(),\n+                    exception.getMessage());\n+            System.exit(1);\n+        } catch (IllegalArgumentException exception) {\n+            System.err.println(\"Configuration error: \" + exception.getMessage());\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static void runCrudOperations(SecretClient secretClient) {\n+        KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, \"my-secret-value\");\n+        System.out.println(\"Created secret version: \" + createdSecret.getProperties().getVersion());\n+\n+        KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);\n+        System.out.println(\"Secret value: \" + readSecret.getValue());\n+\n+        KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, \"updated-value\");\n+        System.out.println(\"Updated secret value: \" + updatedSecret.getValue());\n+\n+        SyncPoller\u003cDeletedSecret, Void\u003e deletionPoller = secretClient.beginDeleteSecret(SECRET_NAME);\n+        deletionPoller.waitForCompletion();\n+        System.out.println(\"Deleted secret: \" + SECRET_NAME);\n+\n+        secretClient.purgeDeletedSecret(SECRET_NAME);\n+        System.out.println(\"Purged secret: \" + SECRET_NAME);\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(\n+                    name + \" must be set to a vault URL such as https://your-vault.vault.azure.net\");\n+        }\n+        return value;\n+    }\n+}\n*** End Patch\n"
```

### 🔧 6. powershell

**Input:**

```json
{"command":"mvn --quiet compile","description":"Compile the Maven application","initial_wait":180,"mode":"sync"}
```

### 🔧 7. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/KeyVaultSecretCrud.java\n@@\n import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.exception.ClientAuthenticationException;\n import com.azure.core.util.polling.SyncPoller;\n-import com.azure.identity.ClientAuthenticationException;\n import com.azure.identity.CredentialUnavailableException;\n*** End Patch\n"
```

### 🔧 8. powershell

**Input:**

```json
{"command":"mvn --quiet compile","description":"Recompile the Maven application","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `pom.xml`
- `src\main\java\com\example\KeyVaultSecretCrud.java`

## Copilot Response

I’ll create a minimal Maven project with environment-based vault configuration, the full secret lifecycle, and explicit Azure error handling.

The project is empty, so I’m adding the complete Maven setup and a single executable class. Deletion will wait for Key Vault’s long-running soft-delete operation before purging.

Compilation exposed an SDK package mismatch: `ClientAuthenticationException` belongs to Azure Core, not Azure Identity. I’m correcting the import and recompiling.

Implemented the Java Azure Key Vault CRUD application.

- `pom.xml` includes:
  - `azure-security-keyvault-secrets:4.11.2`
  - `azure-identity:1.18.5`
- `KeyVaultSecretCrud.java` creates, reads, updates, soft-deletes, waits for deletion, and purges `my-secret`.
- Uses `DefaultAzureCredential`.
- Handles configuration, credential, authentication, and Key Vault HTTP errors.
- Reads the vault endpoint from `AZURE_KEY_VAULT_URL`.

Run with:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://your-vault.vault.azure.net"
mvn exec:java
```

The authenticated identity needs secret `get`, `set`, `delete`, and `purge` permissions.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The application does not create any Azure SDK client instances that require explicit resource closure.

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
hyoka run --prompt-id key-vault-dp-java-crud --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)

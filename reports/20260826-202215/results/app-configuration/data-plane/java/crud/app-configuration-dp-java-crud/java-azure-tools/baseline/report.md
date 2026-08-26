# Evaluation Report: app-configuration-dp-java-crud

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 323.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-java-crud` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 15/19 |
| Duration | 323.2s |
| Timestamp | 2026-08-26T14:03:23Z |
| Files Generated | 2 |
| Event Count | 2249 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 85.5s |
| Review | 236.6s |
| **Total** | **323.2s** |

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
| Token Usage | in=103424 out=1441 |
| Turn Count | 5 |

## Prompt Sent

```
Write a Java program that manages
configuration settings in Azure App Configuration:
1. Create a ConfigurationClient using ConfigurationClientBuilder with a connection string
2. Set a configuration setting with key "app:Settings:FontSize" and value "24"
3. Set a setting with label "Production"
4. Get the setting by key and print its value
5. List settings with a key filter "app:Settings:*" using listConfigurationSettings
6. Create a feature flag configuration setting for "BetaFeature"
7. Delete the setting by key

Show required Maven dependency (com.azure:azure-data-appconfiguration) and
proper error handling with HttpResponseException.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. powershell

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace root"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Maven Central latest stable com.azure azure-data-appconfiguration version ConfigurationClient FeatureFlagConfigurationSetting Java"}
```

### 🔧 4. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eapp-configuration-example\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0-SNAPSHOT\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e11\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-data-appconfiguration\u003c/artifactId\u003e\n+            \u003cversion\u003e1.10.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/AppConfigurationExample.java\n+package com.example;\n+\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.data.appconfiguration.ConfigurationClient;\n+import com.azure.data.appconfiguration.ConfigurationClientBuilder;\n+import com.azure.data.appconfiguration.models.ConfigurationSetting;\n+import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;\n+import com.azure.data.appconfiguration.models.SettingSelector;\n+\n+public final class AppConfigurationExample {\n+    private static final String CONNECTION_STRING_ENV = \"AZURE_APP_CONFIGURATION_CONNECTION_STRING\";\n+    private static final String FONT_SIZE_KEY = \"app:Settings:FontSize\";\n+\n+    private AppConfigurationExample() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String connectionString = System.getenv(CONNECTION_STRING_ENV);\n+        if (connectionString == null || connectionString.isBlank()) {\n+            System.err.printf(\"Set the %s environment variable before running.%n\", CONNECTION_STRING_ENV);\n+            return;\n+        }\n+\n+        ConfigurationClient client = new ConfigurationClientBuilder()\n+            .connectionString(connectionString)\n+            .buildClient();\n+\n+        try {\n+            // Create an unlabeled setting.\n+            client.setConfigurationSetting(FONT_SIZE_KEY, null, \"24\");\n+\n+            // The same key can have a separate value for the Production label.\n+            client.setConfigurationSetting(FONT_SIZE_KEY, \"Production\", \"24\");\n+\n+            ConfigurationSetting setting = client.getConfigurationSetting(FONT_SIZE_KEY, null);\n+            System.out.printf(\"%s = %s%n\", setting.getKey(), setting.getValue());\n+\n+            SettingSelector selector = new SettingSelector().setKeyFilter(\"app:Settings:*\");\n+            client.listConfigurationSettings(selector).forEach(item -\u003e\n+                System.out.printf(\"Found: key=%s, label=%s, value=%s%n\",\n+                    item.getKey(), item.getLabel(), item.getValue()));\n+\n+            FeatureFlagConfigurationSetting featureFlag =\n+                new FeatureFlagConfigurationSetting(\"BetaFeature\", false);\n+            client.setConfigurationSetting(featureFlag);\n+\n+            client.deleteConfigurationSetting(FONT_SIZE_KEY, null);\n+            System.out.printf(\"Deleted unlabeled setting: %s%n\", FONT_SIZE_KEY);\n+        } catch (HttpResponseException exception) {\n+            int statusCode = exception.getResponse() == null\n+                ? -1\n+                : exception.getResponse().getStatusCode();\n+            System.err.printf(\"Azure App Configuration request failed (HTTP %d): %s%n\",\n+                statusCode, exception.getMessage());\n+        }\n+    }\n+}\n*** End Patch\n"
```

### 🔧 5. powershell

**Input:**

```json
{"command":"mvn -q -DskipTests compile","description":"Compile Maven example","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `pom.xml`
- `src\main\java\com\example\AppConfigurationExample.java`

## Copilot Response

I’ll inspect the workspace, then add a locally compilable Maven example without connecting to Azure.

The workspace is empty. I’ll create a minimal Maven project with all requested operations and explicit Azure HTTP error handling.

The example is in place; I’m compiling it to confirm the current SDK API signatures.

Created a complete Maven example:

- `pom.xml` includes `com.azure:azure-data-appconfiguration:1.10.0`
- `src\main\java\com\example\AppConfigurationExample.java` implements:
  - `ConfigurationClientBuilder` with a connection string
  - Unlabeled and `Production`-labeled settings
  - Get, print, filter/list, feature flag creation, and deletion
  - `HttpResponseException` handling with HTTP status reporting
- Uses the `AZURE_APP_CONFIGURATION_CONNECTION_STRING` environment variable rather than embedding credentials.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The client lifecycle is appropriate for the Azure App Configuration ConfigurationClient API.

## Grader Results

- config-values.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `azure-data-appconfiguration` Maven dependency: Pass
      - `ConfigurationClientBuilder` and `ConfigurationClient`: Pass
      - `setConfigurationSetting()` with key, value, label: Pass
      - `getConfigurationSetting()` by key and label: Pass
      - `listConfigurationSettings()` with `SettingSelector`: Pass
      - `FeatureFlagConfigurationSetting` for feature flags: Pass
      - `deleteConfigurationSetting()` and exception handling: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Fail
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Fail
  - Client Builder Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Fail
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 9.0000** | **69.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id app-configuration-dp-java-crud --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)

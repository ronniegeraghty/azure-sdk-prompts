# Evaluation Report: resource-manager-mp-java-rg-crud

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 509.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-java-rg-crud` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 15/19 |
| Duration | 509.8s |
| Timestamp | 2026-08-26T13:21:43Z |
| Files Generated | 3 |
| Event Count | 5642 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 216.6s |
| Review | 291.7s |
| **Total** | **509.8s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=545683 out=5371 |
| Turn Count | 11 |

## Prompt Sent

```
Write a Java application that manages
Azure Resource Groups using the management plane SDK:
1. Authenticate using DefaultAzureCredential
2. Create a new resource group in "eastus" region
3. List all resource groups in the subscription
4. Get details of the created resource group
5. Add a tag to the resource group
6. Delete the resource group

Show required Maven dependencies and include proper exception handling.
Use the modern azure-resourcemanager SDK.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local Java sample using the modern Azure Resource Manager management-plane SDK and DefaultAzureCredential to create, list, read, tag, and delete a resource group. Do not execute against Azure.","learn":true}
```

### 🔧 3. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-java-rg-crud-java-azure-tools-with-azure-tools-4078361727","pattern":"**/*"}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code generation guidance for a Java management-plane SDK sample that performs resource group CRUD locally as source code only.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official current Java Azure Resource Manager SDK documentation and examples for azure-resourcemanager, DefaultAzureCredential, AzureProfile, and resource group create/list/get/update tags/delete APIs.","learn":true}
```

### 🔧 7. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Java azure-resourcemanager resource groups create list update tags delete AzureResourceManager authenticate AzureProfile Maven dependency current"}
```

### 🔧 8. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Maven Central com.azure.resourcemanager azure-resourcemanager latest stable version August 2026"}
```

### 🔧 9. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Java Azure Resource Manager SDK docs for resourceGroups define withRegion create, list, getByName, update withTag apply, deleteByName and AzureResourceManager authenticate DefaultAzureCredential AzureProfile.","parameters":{"query":"Java AzureResourceManager resourceGroups define withRegion create list getByName update withTag apply deleteByName DefaultAzureCredential AzureProfile"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find latest official Java code samples for managing Azure resource groups using com.azure.resourcemanager AzureResourceManager and DefaultAzureCredential.","parameters":{"language":"java","query":"AzureResourceManager resourceGroups create list get update tags delete DefaultAzureCredential AzureProfile"}}
```

### 🔧 11. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Maven Central com.azure azure-identity latest stable version August 2026"}
```

### 🔧 12. web_fetch

**Input:**

```json
{"max_length":10000,"url":"https://central.sonatype.com/artifact/com.azure.resourcemanager/azure-resourcemanager/2.63.0"}
```

### 🔧 13. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787750613233-copilot-tool-output-08b40248e1654cc78461eb33ffb85ad5.txt","pattern":"withTag|withTags|getByName|deleteByName|define\\(|AzureProfile|DefaultAzureCredential|url"}
```

### 🔧 14. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787750613721-copilot-tool-output-6da038dbc87d42bda4d6ac84aa561cad.txt","pattern":"resourceGroups|withTag|withTags|getByName|deleteByName|AzureProfile|DefaultAzureCredential|sourceUrl|url"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$docs = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787750613233-copilot-tool-output-08b40248e1654cc78461eb33ffb85ad5.txt' | ConvertFrom-Json; $docs.results | Select-Object title, contentUrl, @{n='excerpt';e={($_.content -replace \"`r?`n\",' ') -replace '\\s+',' ' | Select-Object -First 1}} | ForEach-Object { \"TITLE: $($_.title)`nURL: $($_.contentUrl)`nEXCERPT: $($_.excerpt.Substring(0, [Math]::Min(700, $_.excerpt.Length)))`n\" }","description":"Summarize official SDK documentation results"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787750613233-copilot-tool-output-08b40248e1654cc78461eb33ffb85ad5.txt","view_range":[1,12]}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787750613721-copilot-tool-output-6da038dbc87d42bda4d6ac84aa561cad.txt","view_range":[1,8]}
```

### 🔧 18. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Azure Resource Manager Java SDK overview identified by documentation search.","parameters":{"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable"}}
```

### 🔧 19. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official DefaultAzureCredential Azure SDK for Java authentication overview identified by code sample search.","parameters":{"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/overview"}}
```

### 🔧 20. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.resources.models.resourcegroup?view=azure-java-stable"}
```

### 🔧 21. web_fetch

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.resources.resourcegroups?view=azure-java-stable"}
```

### 🔧 22. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-resource-group-manager\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.resourcemanager.version\u003e2.63.0\u003c/azure.resourcemanager.version\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure.resourcemanager\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-resourcemanager\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.resourcemanager.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.ResourceGroupManagerApp\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/ResourceGroupManagerApp.java\n+package com.example;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.core.management.AzureEnvironment;\n+import com.azure.core.management.Region;\n+import com.azure.core.management.exception.ManagementException;\n+import com.azure.core.management.profile.AzureProfile;\n+import com.azure.identity.CredentialUnavailableException;\n+import com.azure.identity.DefaultAzureCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.resourcemanager.AzureResourceManager;\n+import com.azure.resourcemanager.resources.models.ResourceGroup;\n+\n+import java.util.UUID;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+\n+public final class ResourceGroupManagerApp {\n+    private static final Logger LOGGER = Logger.getLogger(ResourceGroupManagerApp.class.getName());\n+    private static final Region REGION = Region.US_EAST;\n+    private static final String TAG_KEY = \"environment\";\n+    private static final String TAG_VALUE = \"sdk-demo\";\n+\n+    private ResourceGroupManagerApp() {\n+    }\n+\n+    public static void main(String[] args) {\n+        try {\n+            manageResourceGroup();\n+        } catch (CredentialUnavailableException exception) {\n+            LOGGER.log(Level.SEVERE,\n+                \"DefaultAzureCredential could not find a usable credential. Configure a supported \"\n+                    + \"developer credential or managed identity.\", exception);\n+            System.exit(1);\n+        } catch (ClientAuthenticationException exception) {\n+            LOGGER.log(Level.SEVERE,\n+                \"Microsoft Entra authentication failed. Verify the credential and tenant configuration.\",\n+                exception);\n+            System.exit(1);\n+        } catch (ManagementException exception) {\n+            LOGGER.log(Level.SEVERE,\n+                \"Azure Resource Manager rejected an operation. Verify the subscription, RBAC permissions, \"\n+                    + \"resource name, and regional policy.\", exception);\n+            System.exit(1);\n+        } catch (RuntimeException exception) {\n+            LOGGER.log(Level.SEVERE, \"The resource group workflow failed unexpectedly.\", exception);\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static void manageResourceGroup() {\n+        AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);\n+        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()\n+            .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())\n+            .build();\n+\n+        AzureResourceManager azure = AzureResourceManager\n+            .authenticate(credential, profile)\n+            .withDefaultSubscription();\n+\n+        String resourceGroupName = \"rg-java-sdk-demo-\" + UUID.randomUUID()\n+            .toString()\n+            .replace(\"-\", \"\")\n+            .substring(0, 12);\n+\n+        boolean created = false;\n+        RuntimeException operationFailure = null;\n+\n+        try {\n+            LOGGER.info(() -\u003e \"Creating resource group \" + resourceGroupName + \" in \" + REGION.name());\n+            ResourceGroup createdGroup = azure.resourceGroups()\n+                .define(resourceGroupName)\n+                .withRegion(REGION)\n+                .create();\n+            created = true;\n+            logResourceGroup(\"Created\", createdGroup);\n+\n+            LOGGER.info(\"Listing resource groups in subscription \" + azure.subscriptionId());\n+            for (ResourceGroup group : azure.resourceGroups().list()) {\n+                LOGGER.info(() -\u003e String.format(\"Resource group: name=%s, region=%s\",\n+                    group.name(), group.regionName()));\n+            }\n+\n+            LOGGER.info(() -\u003e \"Getting details for \" + resourceGroupName);\n+            ResourceGroup retrievedGroup = azure.resourceGroups().getByName(resourceGroupName);\n+            if (retrievedGroup == null) {\n+                throw new IllegalStateException(\n+                    \"The newly created resource group could not be retrieved: \" + resourceGroupName);\n+            }\n+            logResourceGroup(\"Retrieved\", retrievedGroup);\n+\n+            LOGGER.info(() -\u003e String.format(\"Adding tag %s=%s to %s\",\n+                TAG_KEY, TAG_VALUE, resourceGroupName));\n+            ResourceGroup taggedGroup = retrievedGroup.update()\n+                .withTag(TAG_KEY, TAG_VALUE)\n+                .apply();\n+            logResourceGroup(\"Tagged\", taggedGroup);\n+        } catch (RuntimeException exception) {\n+            operationFailure = exception;\n+            throw exception;\n+        } finally {\n+            if (created) {\n+                try {\n+                    LOGGER.info(() -\u003e \"Deleting resource group \" + resourceGroupName);\n+                    azure.resourceGroups().deleteByName(resourceGroupName);\n+                    LOGGER.info(() -\u003e \"Deleted resource group \" + resourceGroupName);\n+                } catch (RuntimeException deletionFailure) {\n+                    if (operationFailure != null) {\n+                        operationFailure.addSuppressed(deletionFailure);\n+                    } else {\n+                        throw deletionFailure;\n+                    }\n+                }\n+            }\n+        }\n+    }\n+\n+    private static void logResourceGroup(String action, ResourceGroup group) {\n+        LOGGER.info(() -\u003e String.format(\n+            \"%s resource group: name=%s, id=%s, region=%s, provisioningState=%s, tags=%s\",\n+            action,\n+            group.name(),\n+            group.id(),\n+            group.regionName(),\n+            group.provisioningState(),\n+            group.tags()));\n+    }\n+}\n*** Add File: README.md\n+# Azure Resource Group Manager (Java)\n+\n+This sample uses the modern Azure management-plane SDK to create, list, read, tag, and\n+delete an Azure resource group in the `eastus` region. The generated name is unique, and\n+the application attempts deletion in a `finally` block after a successful create.\n+\n+## Prerequisites\n+\n+- JDK 17 or later\n+- Maven 3.9 or later\n+- An Azure subscription\n+- An identity with permission to manage resource groups, scoped as narrowly as practical\n+\n+`DefaultAzureCredential` supports managed identity when the application runs in Azure.\n+For local development, sign in with a supported developer tool or configure these\n+environment variables for a service principal:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:AZURE_TENANT_ID = \"\u003ctenant-id\u003e\"\n+$env:AZURE_CLIENT_ID = \"\u003cclient-id\u003e\"\n+$env:AZURE_CLIENT_SECRET = \"\u003cclient-secret\u003e\"\n+```\n+\n+Do not store credentials in source code. Use managed identity for Azure-hosted workloads.\n+\n+## Maven dependencies\n+\n+The complete dependencies are in `pom.xml`:\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure.resourcemanager\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-resourcemanager\u003c/artifactId\u003e\n+    \u003cversion\u003e2.63.0\u003c/version\u003e\n+\u003c/dependency\u003e\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+    \u003cversion\u003e1.18.2\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+## Build and run\n+\n+```powershell\n+mvn clean compile\n+mvn exec:java\n+```\n+\n+Running the application performs real management-plane operations and deletes the resource\n+group at the end. The source can be built without connecting to Azure.\n+\n+## References\n+\n+- [Azure Resource Manager client library for Java](https://learn.microsoft.com/en-us/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable)\n+- [Authenticate Java apps with Azure Identity](https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/overview)\n+- [AzureResourceManager Java API](https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.azureresourcemanager?view=azure-java-stable)\n*** End Patch\n"
```

### 🔧 23. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile the Maven Java application","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\ResourceGroupManagerApp.java`

## Copilot Response

I’ll inspect the workspace and ground the implementation in the current Azure Java SDK guidance before creating the local-only sample.

The workspace is empty. I’m confirming the exact current APIs and compatible dependency versions, then I’ll create a complete Maven project with cleanup-safe deletion and operation-specific error handling.

The API shape is confirmed against `azure-resourcemanager` 2.63.0. I’m creating the Maven project now, with a unique resource-group name, explicit authentication/management error reporting, and guaranteed cleanup that preserves the original failure if deletion also fails.

The project is implemented. I’m compiling it locally only—no Azure sign-in or resource operation will be executed—to catch SDK signature or dependency issues.

Implemented a complete Maven Java application using the modern Azure management-plane SDK:

- `pom.xml` — `azure-resourcemanager:2.63.0` and `azure-identity:1.18.2`
- `src\main\java\com\example\ResourceGroupManagerApp.java` — authenticates with `DefaultAzureCredential`, creates in `eastus`, lists, retrieves, tags, and deletes the resource group
- `README.md` — prerequisites, environment configuration, build/run instructions, and official references

The application uses a unique resource-group name, handles authentication and ARM failures explicitly, and guarantees deletion in a `finally` block while preserving any original exception. No Azure operations were executed.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No Azure SDK client requiring resource closure is left unclosed.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - Maven dependency for `azure-resourcemanager` and `azure-identity`: Pass
      - `AzureResourceManager.authenticate()` with credential and profile: Pass
      - `resourceGroups().define().withRegion().create()`: Pass
      - `resourceGroups().list()` iteration: Pass
      - `resourceGroups().getByName()` for details: Pass
      - Tag management via `update().withTag().apply()`: Pass
      - `resourceGroups().deleteByName()` for cleanup: Pass
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
  - Client Builder Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Fail
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Fail
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Fail
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
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 9.0000** | **69.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-java-rg-crud --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)

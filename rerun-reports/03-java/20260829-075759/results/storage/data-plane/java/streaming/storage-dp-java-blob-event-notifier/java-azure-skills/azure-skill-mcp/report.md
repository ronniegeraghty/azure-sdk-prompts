# Evaluation Report: storage-dp-java-blob-event-notifier

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 447.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-blob-event-notifier` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 21/22 |
| Duration | 447.8s |
| Timestamp | 2026-08-29T03:46:56Z |
| Files Generated | 16 |
| Event Count | 17981 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 211.5s |
| Review | 235.0s |
| **Total** | **447.8s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=601639 out=13073 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 3023ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 8856ms

**Input:**

```json
{"command":"get_ai_application_best_practices","intent":"Create a local-only Java 17 Maven demo using Azure Blob Storage and Event Grid SDKs with managed identity, synchronous and asynchronous receivers, handlers, and publishers."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 9863ms

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

### ✅ 4. view — 9495ms

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

### ✅ 5. view — 9182ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
@@ -1,115 +1,115 @@
 # Azure Storage SDK Usage
 
 SDK packages and quick start examples for Azure Storage services.
 
 ## Storage SDKs by Language
 
 | Language | Blob | Queue | File Share | Data Lake |
 |----------|------|-------|------------|----------|
 | .NET | `Azure.Storage.Blobs` | `Azure.Storage.Queues` | `Azure.Storage.Files.Shares` | `Azure.Storage.Files.DataLake` |
 | Java | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | JavaScript | `@azure/storage-blob` | `@azure/storage-queue` | `@azure/storage-file-share` | `@azure/storage-file-datalake` |
 | Python | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | Go | `azblob` | `azqueue` | `azfile` | `azdatalake` |
 | Rust | `azure_storage_blob` | `azure_storage_queue` | - | - |
 
 ## Installation Commands
 
 | Language | Install Blob SDK + Identity |
 |----------|-----------------------------|
 | .NET | `dotnet add package Azure.Storage.Blobs` `dotnet add package Azure.Identity` |
 | Java | Maven: `com.azure:azure-storage-blob` `com.azure:azure-identity` |
 | JavaScript | `npm install @azure/storage-blob @azure/identity` |
 | Python | `pip install azure-storage-blob azure-identity` |
 | Go | `go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob github.com/Azure/azure-sdk-for-g
... (truncated)
```

### ✅ 6. glob — 8485ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 4289ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java code-generation practices for a local-only Blob Storage and Event Grid sample using managed identity.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 5642ms

**Input:**

```json
{"command":"search","intent":"Find official Java SDK documentation for deserializing Event Grid schema and CloudEvents JSON, publishing events with TokenCredential, and Blob sync/async download and properties APIs.","parameters":{"query":"Java Azure Event Grid SDK EventGridEvent CloudEvent fromString deserialize publish TokenCredential BlobClient download properties access tier"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 9. github-mcp-server-search_code (via github-mcp-server) — 5292ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-java \"EventGridEvent.fromString\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"CHANGELOG.md","path":"sdk/eventgrid/azure-messaging-eventgrid/CHANGELOG.md","sha":"582959a625bdda7381af2311ab5f3437e25c8e64","text_matches":[{"fragment":"  serialize custom events and the data of `CloudEvent` and `EventGridEvent`.\n- `EventGridPublisherClient` is changed to `EventGridPublisherClient<T>` that can be statically instantiated to send `CloudEvent`, `EventGridEvent` or custom events (use `BinaryData`)\n  with methods `sendEvents` and `sendEvent`. \n  `EventGridPublisherClientBuilder` now has `buildCloudEventPublisherClient`, `buildEventGridEventPublisherClient` and `buildCustomEventPublisherClient` to build the generic-instantiated clients respectively.\n  The async client has the same change.\n- `EventGridPublisherClientBuilder.endpoint()` now requires the EventGrid topic or domain full url endpoint because\n  different EventGrid service deployments may require different url patterns.","matches":[{"indices":[386,420],"text":"buildEventGridEventPublisherClient"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/eventgrid/azure-messaging-eventgrid/CHANGELOG.md?ref=59cbd241703b0f890f8635d59231aace4d481ce9","property":"content"}]}],"total_count":1}
```

### ✅ 10. github-mcp-server-search_code (via github-mcp-server) — 4960ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-java \"buildEventGridEventPublisherClient\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"EventGridEvent.java","path":"sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridEvent.java","sha":"8f59bf7f25bf25c2dd8997393f47fcdd2e587cce","text_matches":[{"fragment":" * of one or more of EventGridEvents, use {@link #fromString(String)} to deserialize them from the Json string.</p>\n *\n * <p><strong>Deserialize EventGridEvent Samples</strong></p>\n * <!-- src_embed com.azure.messaging.eventgrid.EventGridEvent.fromString -->\n * <pre>\n * List&lt;EventGridEvent&gt; eventGridEventList = EventGridEvent.fromString&#40;eventGridEventJsonString&#41;;\n * EventGridEvent eventGridEvent = eventGridEventList.get&#40;0&#41;;","matches":[{"indices":[229,254],"text":"EventGridEvent.fromString"},{"indices":[319,344],"text":"EventGridEvent.fromString"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridEvent.java?ref=59cbd241703b0f890f8635d59231aace4d481ce9","property":"content"}]}],"total_count":1}
```

### ✅ 11. glob — 4618ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 12. azure-documentation (via azure) — 8591ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official current Java SDK documentation for Azure Event Grid publishing/deserialization and Blob Storage downloads with managed identity.","parameters":{"query":"Azure Java SDK Event Grid EventGridEvent CloudEvent deserialize publish managed identity BlobClient BlobAsyncClient download properties access tier"}}
```

**Output:**

```
{"results":[{"title":"Azure Event Grid client library for Java - version 4.31.5","content":"# Azure Event Grid client library for Java - version 4.31.5\n## Examples\nIf the domain accepts \u0060CloudEvent\u0060 schema, the CloudEvent\u0027s attribute that is configured to map the \u0060topic\u0060 when the domain is created must be set. The default mapping attribute is \u0060source\u0060.\n### Receiving and Consuming Events\nThe Event Grid service doesn\u0027t store events. So this Event Grid SDK doesn\u0027t have an event receiver. Instead, events are stored in the Event Handlers, including ServiceBus, EventHubs, Storage Queue, WebHook endpoint, or many other supported Azure Services. However, currently all events will be sent and stored as encoded JSON data. Here is some basic code that details the deserialization of events after they\u0027re received by the event handlers. Again, the handling is different based on the event schema being received from the topic/subscription.\n#### Deserialize \u0060EventGridEvent\u0060 or \u0060CloudEvent\u0060 from a Json String\nThe Json String can have a single event or an array of events. The returned result is a list of events.\n\u0060\u0060\u0060java\n// Deserialize an EventGridEvent\nString eventGridEventJsonData = \u0022\u003Cyour EventGridEvent json String\u003E\u0022;\nList\u003CEventGridEvent\u003E eventGridEvents = EventGridEvent.fromString(eventGridEventJsonData);\n\n// Deserialize a CloudEvent\nString cloudEventJsonData = \u0022\u003Cyour CloudEvent json String\u003E\u0022;\nList\u003CCloudEvent\u003E cloudEvents = CloudEvent.fromString(cloudEventJsonData);\n\u0060\u0060\u0060\n#### Deserialize data from a \u0060CloudEvent\u0060 or \u0060EventGridEvent\u0060\nOnce you deserialize the \u0060EventGridEvent\u0060 or \u0060CloudEvent\u0060 from a Json String, you can use \u0060getData()\u0060 of \u0060CloudEvent\u0060 or \u0060EventGridEvent\u0060 to get the payload of the event. It returns a \u0060BinaryData\u0060 objec
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 8270ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find latest official Java examples for Event Grid publisher clients with TokenCredential and EventGridEvent/CloudEvent model deserialization.","parameters":{"language":"java","query":"EventGridPublisherClientBuilder TokenCredential EventGridEvent.fromString CloudEvent.fromString"}}
```

**Output:**

```
{"results":[{"description":"description: Deserializes a JSON string into a list of EventGridEvent objects and a list of CloudEvent objects.\npackage: com.azure.messaging.eventgrid\nlanguage: java\n","codeSnippet":"// Deserialize an EventGridEvent\r\nString eventGridEventJsonData = \u0022\u003Cyour EventGridEvent json String\u003E\u0022;\r\nList\u003CEventGridEvent\u003E eventGridEvents = EventGridEvent.fromString(eventGridEventJsonData);\r\n\r\n// Deserialize a CloudEvent\r\nString cloudEventJsonData = \u0022\u003Cyour CloudEvent json String\u003E\u0022;\r\nList\u003CCloudEvent\u003E cloudEvents = CloudEvent.fromString(cloudEventJsonData);","link":"https://learn.microsoft.com/java/api/overview/azure/messaging-eventgrid-readme?view=azure-java-stable#examples","language":"java"},{"description":"description: Demonstrates how to create an Event Grid client using topic credentials, construct a list of custom EventGridEvent objects, and publish these events to an Azure Event Grid topic endpoint.\npackage: azure-eventgrid\nlanguage: java\n","codeSnippet":"// Create an event grid client.\r\nTopicCredentials topicCredentials = new TopicCredentials(System.getenv(\u0022EVENTGRID_TOPIC_KEY\u0022));\r\nEventGridClient client = new EventGridClientImpl(topicCredentials);\r\n\r\n// Publish custom events to the EventGrid.\r\nSystem.out.println(\u0022Publish custom events to the EventGrid\u0022);\r\nList\u003CEventGridEvent\u003E eventsList = new ArrayList\u003C\u003E();\r\nfor (int i = 0; i \u003C 5; i\u002B\u002B) {\r\n    eventsList.add(new EventGridEvent(\r\n        UUID.randomUUID().toString(),\r\n        String.format(\u0022Door%d\u0022, i),\r\n        new ContosoItemReceivedEventData(\u0022Contoso Item SKU #1\u0022),\r\n        \u0022Contoso.Items.ItemReceived\u0022,\r\n        DateTime.now(),\r\n        \u00222.0\u0022\r\n    ));\r\n}\r\n\r\nString eventGridEndpoint = String.format(\u0022https://%s/\u0022, new URI(System.getenv(\u0022EVENTGRID_TOPIC_ENDPOINT\u0022)).getHost()
... (truncated)
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server) — 9562ms

**Input:**

```json
{"owner":"Azure","path":"sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridEvent.java","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 4480d939e51e7a09d95310f4034197ac5e4fbf1c)<!--
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for
 * license information.
 *
 * Code generated by Microsoft (R) AutoRest Code Generator.
-->
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.azure</groupId>
  <artifactId>azure-messaging-eventgrid</artifactId>
  <version>4.32.0-beta.1</version> <!-- {x-version-update;com.azure:azure-messaging-eventgrid;current} -->

  <packaging>jar</packaging>
  <name>Microsoft Azure SDK for eventgrid</name>
  <description>This package contains Microsoft Azure EventGrid SDK.</description>
  <parent>
    <groupId>com.azure</groupId>
    <artifactId>azure-client-sdk-parent</artifactId>
    <version>1.7.0</version> <!-- {x-version-update;com.azure:azure-client-sdk-parent;current} -->
    <relativePath>../../parents/azure-client-sdk-parent</relativePath>
  </parent>
  <url>https://github.com/Azure/azure-sdk-for-java</url>

  <licenses>
    <license>
      <name>The MIT License (MIT)</name>
      <url>http://opensource.org/licenses/MIT</url>
      <distribution>repo</distribution>
    </license>
  </licenses>

  <distributionManagement>
    <site>
      <id>azure-java-build-docs</id>
      <url>${site.url}/site/${project.artifactId}</url>
    </site>
  </distributionManagement>

  <scm>
    <url>scm:git:https://github.com/Azure/azure-sdk-for-java</url>
    <connection>scm:git:git@github.com:Azure/azure-sdk-for-java.git</connection>
    <tag>HEAD</tag>
  </scm>

  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    <legal><![CDATA[[INFO] Any downloads listed may be third party software.  Microsoft grants you no r
... (truncated)
```

### ✅ 15. github-mcp-server-get_file_contents (via github-mcp-server) — 9230ms

**Input:**

```json
{"owner":"Azure","path":"sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridPublisherClientBuilder.java","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 65ce0036329f691a7f0ec1369eb6cf5803e22f53)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.messaging.eventgrid;

import com.azure.core.annotation.ServiceClientBuilder;
import com.azure.core.client.traits.AzureKeyCredentialTrait;
import com.azure.core.client.traits.AzureSasCredentialTrait;
import com.azure.core.client.traits.ConfigurationTrait;
import com.azure.core.client.traits.EndpointTrait;
import com.azure.core.client.traits.HttpTrait;
import com.azure.core.client.traits.TokenCredentialTrait;
import com.azure.core.credential.AzureKeyCredential;
import com.azure.core.credential.AzureSasCredential;
import com.azure.core.credential.TokenCredential;
import com.azure.core.http.HttpClient;
import com.azure.core.http.HttpHeader;
import com.azure.core.http.HttpHeaders;
import com.azure.core.http.HttpPipeline;
import com.azure.core.http.HttpPipelineBuilder;
import com.azure.core.http.policy.AddDatePolicy;
import com.azure.core.http.policy.AddHeadersFromContextPolicy;
import com.azure.core.http.policy.AddHeadersPolicy;
import com.azure.core.http.policy.AzureKeyCredentialPolicy;
import com.azure.core.http.policy.BearerTokenAuthenticationPolicy;
import com.azure.core.http.policy.HttpLogOptions;
import com.azure.core.http.policy.HttpLoggingPolicy;
import com.azure.core.http.policy.HttpPipelinePolicy;
import com.azure.core.http.policy.HttpPolicyProviders;
import com.azure.core.http.policy.RequestIdPolicy;
import com.azure.core.http.policy.RetryOptions;
import com.azure.core.http.policy.RetryPolicy;
import com.azure.core.http.policy.UserAgentPolicy;
import com.azure.core.models.CloudEvent;
import com.azure.core.util.BinaryData;
import com.azure.core.util.ClientOptions;
import com.azure.core.util.Configuration;
import com.azure.core.util.CoreUtils;
import com.azure.core.util.TracingOptions;
import com.azure.core.util.builder.ClientBuilderUtil;
import com.azure.core.util.logg
... (truncated)
```

### ✅ 16. github-mcp-server-get_file_contents (via github-mcp-server) — 8899ms

**Input:**

```json
{"owner":"Azure","path":"sdk/eventgrid/azure-messaging-eventgrid/pom.xml","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 8f59bf7f25bf25c2dd8997393f47fcdd2e587cce)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.messaging.eventgrid;

import com.azure.core.annotation.Fluent;
import com.azure.core.util.BinaryData;
import com.azure.core.util.CoreUtils;
import com.azure.core.util.logging.ClientLogger;
import com.azure.core.util.serializer.TypeReference;
import com.azure.json.JsonProviders;
import com.azure.json.JsonReader;
import com.azure.json.JsonSerializable;
import com.azure.json.JsonToken;
import com.azure.json.JsonWriter;

import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Objects;

/**
 * Represents the EventGrid event conforming to the <a href="https://docs.microsoft.com/azure/event-grid/event-schema">
 *     EventGrid event schema</a>.
 *
 * <p>Depending on your scenario, you can either use the constructor
 * {@link #EventGridEvent(String, String, BinaryData, String)} to create an EventGridEvent,
 * or use the factory method {@link #fromString(String)} to deserialize EventGridEvent instances
 * from a Json String representation of EventGrid events.</p>
 *
 * <p>If you have the data payload of an EventGridEvent and want to send it out, use the constructor
 * {@link #EventGridEvent(String, String, BinaryData, String)} to create it. Then use
 * {@link EventGridPublisherAsyncClient} or {@link EventGridPublisherClient} to send it the EventGrid service.</p>
 *
 * <p><strong>Create EventGridEvent Samples</strong></p>
 * <!-- src_embed com.azure.messaging.eventgrid.EventGridEvent#constructor -->
 * <pre>
 * &#47;&#47; Use BinaryData.fromObject&#40;&#41; to create EventGridEvent data
 * &#47;&#47; From a model class
 * User user = new User&#40;&quot;Stephen&quot;, &quot;James&quot;&#41;;
 * EventGridEvent eventGridEventDataObject = new EventGridEvent&#40;&quot;&#47;EventGridEvent
... (truncated)
```

### ✅ 17. web_fetch — 7927ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-sdk-bom</artifactId>
  <versioning>
    <latest>1.3.8</latest>
    <release>1.3.8</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.2.6</version>
      <version>1.2.7</version>
      <version>1.2.8</version>
      <version>1.2.9</version>
      <version>1.2.10</version>
      <version>1.2.11</version>
      <version>1.2.12</version>
      <version>1.2.13</version>
      <version>1.2.14</version>
      <version>1.2.15</version>
      <version>1.2.16</version>
      <version>1.2.17</version>
      <version>1.2.18</version>
      <version>1.2.19</version>
      <version>1.2.20</version>
      <version>1.2.21</version>
      <version>1.2.22</version>
      <version>1.2.23</version>
      <version>1.2.24</version>
      <version>1.2.25</version>
      <version>1.2.26</version>
      <version>1.2.27</version>
      <version>1.2.28</version>
      <version>1.2.29</version>
      <version>1.2.30</version>
      <version>1.2.31</version>
      <version>1.2.32</version>
      <version>1.2.33</version>
      <version>1.2.34</version>
      <version>1.2.35</version>
      <version>1.2.36</version>
      <version>1.2.37</version>
      <version>1.2.38</version>
      <version>1.3.0</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version
... (truncated)
```

### ✅ 18. web_fetch — 3920ms

**Input:**

```json
{"max_length":20000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/1.3.8/azure-sdk-bom-1.3.8.pom"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/1.3.8/azure-sdk-bom-1.3.8.pom:
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <groupId>com.azure</groupId>
  <artifactId>azure-sdk-bom</artifactId>
  <version>1.3.8</version>
  <packaging>pom</packaging>
  <name>Azure Java SDK BOM (Bill of Materials)</name>
  <description>Azure Java SDK BOM (Bill of Materials)</description>
  <url>https://github.com/azure/azure-sdk-for-java</url>
  <licenses>
    <license>
      <name>The MIT License (MIT)</name>
      <url>http://opensource.org/licenses/MIT</url>
      <distribution>repo</distribution>
    </license>
  </licenses>
  <developers>
    <developer>
      <id>microsoft</id>
      <name>Microsoft Corporation</name>
    </developer>
  </developers>
  <scm>
    <connection>scm:git:git://github.com/azure/azure-sdk-for-java</connection>
    <developerConnection>scm:git:git://github.com/azure/azure-sdk-for-java</developerConnection>
    <url>https://github.com/azure/azure-sdk-for-java</url>
  </scm>
  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    <packageOutputDirectory>${project.build.directory}</packageOutputDirectory>
  </properties>
  <issueManagement>
    <system>GitHub</system>
    <url>https://github.com/azure/azure-sdk-for-java/issues</url>
  </issueManagement>
  <dependencyManagement>
    <dependencies>
      <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-ai-agents</artifactId>
        <version>2.2.0</version>
      </dependency>
      <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-ai-contentsafety</artifactId>
        <version>1.0.19</version>
      </dependency>
      <dependenc
... (truncated)
```

### ✅ 19. rg — 3594ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787975284332-copilot-tool-output-2264662374cf4ea9b6976c8d4cbd135f.txt","pattern":"Event Grid client library|version 4\\.|TokenCredential|ManagedIdentity|sendEvents|BlobAsyncClient|download"}
```

**Output:**

```
Output too large to read at once (40.8 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787975298353-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787975284332-copilot-tool-output-2264662374cf4ea9b6976c8d4cbd135f.txt:{"results":[{"title":"Azure Event Grid client library for Java - version 4.31.5","content":"# Azure Event Grid client library for Java - version 4.31.5\n## Examples\nIf the domain accepts \u0060CloudEvent\u0060 schema, the CloudEvent\u0027s attribute that is configured to map the \u0060topic\u0060 when the domain is created must be set. The default mapping attribute is \u0060source\u0060.\n#
```

### ✅ 20. github-mcp-server-search_code (via github-mcp-server) — 4556ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-java path:sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/specialized BlobAsyncClient downloadContent"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"EventGridPublisherAsyncClient.java","path":"sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridPublisherAsyncClient.java","sha":"aa4ba286dbf3043fcd452f420d8eef1a8eece10d","text_matches":[{"fragment":"    @ServiceMethod(returns = ReturnType.SINGLE)\n    public Mono<Void> sendEvents(Iterable<T> events) {\n        return withContext(context -> sendEvents(events, context));","matches":[{"indices":[59,80],"text":"Mono<Void> sendEvents"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridPublisherAsyncClient.java?ref=59cbd241703b0f890f8635d59231aace4d481ce9","property":"content"},{"fragment":"    @SuppressWarnings(\"unchecked\")\n    Mono<Void> sendEvents(Iterable<T> events, Context context) {\n        if (this.eventClass == CloudEvent.class) {","matches":[{"indices":[39,60],"text":"Mono<Void> sendEvents"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/eventgrid/azure-messaging-eventgrid/src/main/java/com/azure/messaging/eventgrid/EventGridPublisherAsyncClient.java?ref=59cbd241703b0f890f8635d59231aace4d481ce9","property":"content"}]}],"total_count":1}
```

### ✅ 21. github-mcp-server-search_code (via github-mcp-server) — 4224ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-java path:sdk/eventgrid/azure-messaging-eventgrid/src/main/java \"Mono\u003cVoid\u003e sendEvents\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobClientBase.java","path":"sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/specialized/BlobClientBase.java","sha":"b35c654a1132e1545f12a907820de7d2333bcb68","text_matches":[{"fragment":"     * @throws UncheckedIOException If an I/O error occurs.\n     */\n    @ServiceMethod(returns = ReturnType.SINGLE)\n    public BinaryData downloadContent() {\n        return blockWithOptionalTimeout(client.downloadContent(), null);\n    }\n","matches":[{"indices":[138,153],"text":"downloadContent"},{"indices":[205,220],"text":"downloadContent"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/specialized/BlobClientBase.java?ref=59cbd241703b0f890f8635d59231aace4d481ce9","property":"content"}]},{"name":"BlobAsyncClientBase.java","path":"sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/specialized/BlobAsyncClientBase.java","sha":"c3392bfef4f9848342698cfe5485e8e6c8303ab4","text_matches":[{"fragment":"     *\n     * <p><strong>Code Samples</strong></p>\n     *\n     * <!-- src_embed com.azure.storage.blob.BlobAsyncClient.downloadContent -->\n     * <pre>\n     * client.downloadContent&#40;&#41;.subscribe&#40;data -&gt; &#123;\n     *     System.out.printf&#40;&quot;Downloaded %s&quot;, data.toString&#40;&#41;&#41;;","matches":[{"indices":[103,118],"text":"BlobAsyncClient"},{"indices":[119,134],"text":"downloadContent"},{"indices":[166,181],"text":"downloadContent"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/storage/azure-storage-blob/src/main/java/com/azure/storage/blob/specialized/BlobAsyncClientBase.java?ref=59cbd241703b0f890f8635d59231aace4d481ce9","property":"content"}]}],"total_count":2}
```

### ✅ 22. apply_patch — 702ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eblob-event-notifier\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0-SNAPSHOT\u003c/version\u003e\n+    \u003cname\u003eAzure Blob Event Notifier\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n+        \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventgrid\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.4\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cuseModulePath\u003efalse\u003c/useModulePath\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.blobevents.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/blobevents/EventSchema.java\n+package com.example.blobevents;\n+\n+public enum EventSchema {\n+    EVENT_GRID,\n+    CLOUD_EVENTS\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobLifecycleEvent.java\n+package com.example.blobevents;\n+\n+import com.azure.core.util.BinaryData;\n+\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public record BlobLifecycleEvent(\n+    String id,\n+    String type,\n+    String subject,\n+    OffsetDateTime time,\n+    BinaryData data,\n+    EventSchema schema\n+) {\n+    public BlobLifecycleEvent {\n+        Objects.requireNonNull(id, \"id\");\n+        Objects.requireNonNull(type, \"type\");\n+        Objects.requireNonNull(subject, \"subject\");\n+        Objects.requireNonNull(data, \"data\");\n+        Objects.requireNonNull(schema, \"schema\");\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/CustomEvent.java\n+package com.example.blobevents;\n+\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public record CustomEvent(String eventType, String subject, Map\u003cString, Object\u003e data, String dataVersion) {\n+    public CustomEvent {\n+        Objects.requireNonNull(eventType, \"eventType\");\n+        Objects.requireNonNull(subject, \"subject\");\n+        Objects.requireNonNull(data, \"data\");\n+        Objects.requireNonNull(dataVersion, \"dataVersion\");\n+        if (!subject.startsWith(\"/\")) {\n+            throw new IllegalArgumentException(\"subject must be an absolute hierarchy starting with '/'\");\n+        }\n+        data = Map.copyOf(data);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPayloadDeserializer.java\n+package com.example.blobevents;\n+\n+import com.azure.core.models.CloudEvent;\n+import com.azure.json.JsonProviders;\n+import com.azure.json.JsonReader;\n+import com.azure.json.JsonToken;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+\n+import java.io.IOException;\n+import java.util.List;\n+\n+final class EventPayloadDeserializer {\n+    List\u003cBlobLifecycleEvent\u003e deserialize(String payload) {\n+        EventSchema schema = detectSchema(payload);\n+        if (schema == EventSchema.CLOUD_EVENTS) {\n+            return CloudEvent.fromString(payload).stream()\n+                .map(event -\u003e new BlobLifecycleEvent(\n+                    event.getId(),\n+                    event.getType(),\n+                    requireSubject(event.getSubject(), event.getId()),\n+                    event.getTime(),\n+                    event.getData(),\n+                    schema))\n+                .toList();\n+        }\n+\n+        return EventGridEvent.fromString(payload).stream()\n+            .map(event -\u003e new BlobLifecycleEvent(\n+                event.getId(),\n+                event.getEventType(),\n+                event.getSubject(),\n+                event.getEventTime(),\n+                event.getData(),\n+                schema))\n+            .toList();\n+    }\n+\n+    private EventSchema detectSchema(String payload) {\n+        if (payload == null || payload.isBlank()) {\n+            throw new IllegalArgumentException(\"Event payload must not be blank\");\n+        }\n+\n+        try (JsonReader reader = JsonProviders.createReader(payload)) {\n+            JsonToken token = reader.nextToken();\n+            if (token == JsonToken.START_ARRAY) {\n+                token = reader.nextToken();\n+            }\n+            if (token != JsonToken.START_OBJECT) {\n+                throw new IllegalArgumentException(\"Event payload must contain a JSON object or array\");\n+            }\n+\n+            while (reader.nextToken() != JsonToken.END_OBJECT) {\n+                String fieldName = reader.getFieldName();\n+                reader.nextToken();\n+                if (\"specversion\".equals(fieldName)) {\n+                    return EventSchema.CLOUD_EVENTS;\n+                }\n+                if (\"eventType\".equals(fieldName)) {\n+                    return EventSchema.EVENT_GRID;\n+                }\n+                reader.skipChildren();\n+            }\n+        } catch (IOException exception) {\n+            throw new IllegalArgumentException(\"Event payload is not valid JSON\", exception);\n+        }\n+\n+        throw new IllegalArgumentException(\"Payload is neither Event Grid schema nor CloudEvents 1.0 schema\");\n+    }\n+\n+    private String requireSubject(String subject, String eventId) {\n+        if (subject == null || subject.isBlank()) {\n+            throw new IllegalArgumentException(\"CloudEvent \" + eventId + \" is missing a subject\");\n+        }\n+        return subject;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventReceiver.java\n+package com.example.blobevents;\n+\n+import java.util.List;\n+import java.util.logging.Logger;\n+\n+public final class EventReceiver {\n+    private static final Logger LOGGER = Logger.getLogger(EventReceiver.class.getName());\n+    private static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    private static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private final EventPayloadDeserializer deserializer = new EventPayloadDeserializer();\n+\n+    public List\u003cBlobLifecycleEvent\u003e deserialize(String payload) {\n+        return deserializer.deserialize(payload);\n+    }\n+\n+    public void receive(String payload, BlobEventHandler handler) {\n+        for (BlobLifecycleEvent event : deserialize(payload)) {\n+            switch (event.type()) {\n+                case BLOB_CREATED -\u003e handler.handleCreated(event);\n+                case BLOB_DELETED -\u003e handler.handleDeleted(event);\n+                default -\u003e LOGGER.warning(() -\u003e \"Ignoring unrecognized event type \" + event.type()\n+                    + \" for event \" + event.id());\n+            }\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventReceiver.java\n+package com.example.blobevents;\n+\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.logging.Logger;\n+\n+public final class AsyncEventReceiver {\n+    private static final Logger LOGGER = Logger.getLogger(AsyncEventReceiver.class.getName());\n+    private static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    private static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private final EventPayloadDeserializer deserializer = new EventPayloadDeserializer();\n+\n+    public Flux\u003cBlobLifecycleEvent\u003e deserialize(String payload) {\n+        return Flux.fromIterable(deserializer.deserialize(payload));\n+    }\n+\n+    public Mono\u003cVoid\u003e receive(String payload, BlobEventHandler handler) {\n+        return deserialize(payload)\n+            .concatMap(event -\u003e switch (event.type()) {\n+                case BLOB_CREATED -\u003e handler.handleCreatedAsync(event);\n+                case BLOB_DELETED -\u003e handler.handleDeletedAsync(event);\n+                default -\u003e {\n+                    LOGGER.warning(() -\u003e \"Ignoring unrecognized event type \" + event.type()\n+                        + \" for event \" + event.id());\n+                    yield Mono.empty();\n+                }\n+            })\n+            .then();\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobEventHandler.java\n+package com.example.blobevents;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobAsyncClient;\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.models.BlobProperties;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import reactor.core.publisher.Mono;\n+\n+import java.net.URLDecoder;\n+import java.nio.charset.StandardCharsets;\n+import java.util.Objects;\n+import java.util.logging.Logger;\n+\n+public final class BlobEventHandler {\n+    private static final Logger LOGGER = Logger.getLogger(BlobEventHandler.class.getName());\n+    private static final String CONTAINERS_SEGMENT = \"/containers/\";\n+    private static final String BLOBS_SEGMENT = \"/blobs/\";\n+\n+    private final SyncBlobReader syncReader;\n+    private final AsyncBlobReader asyncReader;\n+\n+    public BlobEventHandler(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {\n+        this(\n+            location -\u003e readBlob(syncClient, location),\n+            location -\u003e readBlobAsync(asyncClient, location));\n+    }\n+\n+    BlobEventHandler(SyncBlobReader syncReader, AsyncBlobReader asyncReader) {\n+        this.syncReader = Objects.requireNonNull(syncReader, \"syncReader\");\n+        this.asyncReader = Objects.requireNonNull(asyncReader, \"asyncReader\");\n+    }\n+\n+    public void handleCreated(BlobLifecycleEvent event) {\n+        BlobLocation location = parseSubject(event.subject());\n+        try {\n+            logSummary(syncReader.read(location));\n+        } catch (BlobStorageException exception) {\n+            if (!handleExpectedRace(location, exception)) {\n+                throw exception;\n+            }\n+        }\n+    }\n+\n+    public Mono\u003cVoid\u003e handleCreatedAsync(BlobLifecycleEvent event) {\n+        BlobLocation location = parseSubject(event.subject());\n+        return asyncReader.read(location)\n+            .doOnNext(BlobEventHandler::logSummary)\n+            .then()\n+            .onErrorResume(BlobStorageException.class, exception -\u003e {\n+                if (handleExpectedRace(location, exception)) {\n+                    return Mono.empty();\n+                }\n+                return Mono.error(exception);\n+            });\n+    }\n+\n+    public void handleDeleted(BlobLifecycleEvent event) {\n+        BlobLocation location = parseSubject(event.subject());\n+        LOGGER.info(() -\u003e \"Blob deleted: container=\" + location.container() + \", name=\" + location.name());\n+    }\n+\n+    public Mono\u003cVoid\u003e handleDeletedAsync(BlobLifecycleEvent event) {\n+        handleDeleted(event);\n+        return Mono.empty();\n+    }\n+\n+    static BlobLocation parseSubject(String subject) {\n+        int containerStart = subject.indexOf(CONTAINERS_SEGMENT);\n+        int blobSeparator = subject.indexOf(BLOBS_SEGMENT, containerStart + CONTAINERS_SEGMENT.length());\n+        if (containerStart \u003c 0 || blobSeparator \u003c 0) {\n+            throw new IllegalArgumentException(\"Unexpected blob event subject: \" + subject);\n+        }\n+\n+        String container = subject.substring(containerStart + CONTAINERS_SEGMENT.length(), blobSeparator);\n+        String name = subject.substring(blobSeparator + BLOBS_SEGMENT.length());\n+        if (container.isBlank() || name.isBlank()) {\n+            throw new IllegalArgumentException(\"Blob event subject must contain a container and blob name: \" + subject);\n+        }\n+        return new BlobLocation(decodePathPart(container), decodePathPart(name));\n+    }\n+\n+    private static BlobSummary readBlob(BlobServiceClient serviceClient, BlobLocation location) {\n+        BlobClient blob = serviceClient.getBlobContainerClient(location.container()).getBlobClient(location.name());\n+        BinaryData ignoredContent = blob.downloadContent();\n+        BlobProperties properties = blob.getProperties();\n+        return toSummary(location, properties);\n+    }\n+\n+    private static Mono\u003cBlobSummary\u003e readBlobAsync(\n+        BlobServiceAsyncClient serviceClient,\n+        BlobLocation location\n+    ) {\n+        BlobAsyncClient blob = serviceClient.getBlobContainerAsyncClient(location.container())\n+            .getBlobAsyncClient(location.name());\n+        return blob.downloadContent()\n+            .then(blob.getProperties())\n+            .map(properties -\u003e toSummary(location, properties));\n+    }\n+\n+    private static BlobSummary toSummary(BlobLocation location, BlobProperties properties) {\n+        String accessTier = properties.getAccessTier() == null ? \"unknown\" : properties.getAccessTier().toString();\n+        return new BlobSummary(\n+            location.name(),\n+            properties.getBlobSize(),\n+            properties.getContentType(),\n+            accessTier);\n+    }\n+\n+    private static void logSummary(BlobSummary summary) {\n+        LOGGER.info(() -\u003e \"Blob created: name=\" + summary.name()\n+            + \", size=\" + summary.size()\n+            + \", contentType=\" + summary.contentType()\n+            + \", accessTier=\" + summary.accessTier());\n+    }\n+\n+    private static boolean handleExpectedRace(BlobLocation location, BlobStorageException exception) {\n+        String errorCode = exception.getErrorCode() == null ? \"\" : exception.getErrorCode().toString();\n+        if (exception.getStatusCode() == 404) {\n+            LOGGER.warning(() -\u003e \"Blob disappeared before it could be read: \" + location);\n+            return true;\n+        }\n+        if (exception.getStatusCode() == 409\n+            \u0026\u0026 (errorCode.contains(\"Archive\") || errorCode.contains(\"Rehydr\"))) {\n+            LOGGER.warning(() -\u003e \"Blob is temporarily unreadable in its current access tier: \" + location\n+                + \" (\" + errorCode + \")\");\n+            return true;\n+        }\n+        return false;\n+    }\n+\n+    private static String decodePathPart(String value) {\n+        return URLDecoder.decode(value.replace(\"+\", \"%2B\"), StandardCharsets.UTF_8);\n+    }\n+\n+    record BlobLocation(String container, String name) {\n+    }\n+\n+    record BlobSummary(String name, long size, String contentType, String accessTier) {\n+    }\n+\n+    @FunctionalInterface\n+    interface SyncBlobReader {\n+        BlobSummary read(BlobLocation location);\n+    }\n+\n+    @FunctionalInterface\n+    interface AsyncBlobReader {\n+        Mono\u003cBlobSummary\u003e read(BlobLocation location);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class EventPublisher {\n+    private final SyncEventSender sender;\n+\n+    public EventPublisher(String topicEndpoint, TokenCredential credential) {\n+        EventGridPublisherClient\u003cEventGridEvent\u003e client = new EventGridPublisherClientBuilder()\n+            .endpoint(topicEndpoint)\n+            .credential(credential)\n+            .buildEventGridEventPublisherClient();\n+        this.sender = client::sendEvents;\n+    }\n+\n+    EventPublisher(SyncEventSender sender) {\n+        this.sender = Objects.requireNonNull(sender, \"sender\");\n+    }\n+\n+    public void publish(List\u003cCustomEvent\u003e events) {\n+        sender.send(toEventGridEvents(events));\n+    }\n+\n+    static List\u003cEventGridEvent\u003e toEventGridEvents(List\u003cCustomEvent\u003e events) {\n+        Objects.requireNonNull(events, \"events\");\n+        if (events.isEmpty()) {\n+            throw new IllegalArgumentException(\"At least one custom event is required\");\n+        }\n+        return events.stream()\n+            .map(event -\u003e new EventGridEvent(\n+                event.subject(),\n+                event.eventType(),\n+                BinaryData.fromObject(event.data()),\n+                event.dataVersion()))\n+            .toList();\n+    }\n+\n+    @FunctionalInterface\n+    interface SyncEventSender {\n+        void send(Iterable\u003cEventGridEvent\u003e events);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class AsyncEventPublisher {\n+    private final AsyncEventSender sender;\n+\n+    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {\n+        EventGridPublisherAsyncClient\u003cEventGridEvent\u003e client = new EventGridPublisherClientBuilder()\n+            .endpoint(topicEndpoint)\n+            .credential(credential)\n+            .buildEventGridEventPublisherAsyncClient();\n+        this.sender = client::sendEvents;\n+    }\n+\n+    AsyncEventPublisher(AsyncEventSender sender) {\n+        this.sender = Objects.requireNonNull(sender, \"sender\");\n+    }\n+\n+    public Mono\u003cVoid\u003e publish(List\u003cCustomEvent\u003e events) {\n+        return sender.send(EventPublisher.toEventGridEvents(events));\n+    }\n+\n+    @FunctionalInterface\n+    interface AsyncEventSender {\n+        Mono\u003cVoid\u003e send(Iterable\u003cEventGridEvent\u003e events);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureConfiguration.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+\n+import java.net.URI;\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class AzureConfiguration {\n+    private final String storageEndpoint;\n+    private final String eventGridTopicEndpoint;\n+    private final TokenCredential credential;\n+\n+    private AzureConfiguration(\n+        String storageEndpoint,\n+        String eventGridTopicEndpoint,\n+        TokenCredential credential\n+    ) {\n+        this.storageEndpoint = requireHttpsEndpoint(storageEndpoint, \"storageEndpoint\");\n+        this.eventGridTopicEndpoint = requireHttpsEndpoint(eventGridTopicEndpoint, \"eventGridTopicEndpoint\");\n+        this.credential = Objects.requireNonNull(credential, \"credential\");\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        return fromEnvironment(System.getenv());\n+    }\n+\n+    static AzureConfiguration fromEnvironment(Map\u003cString, String\u003e environment) {\n+        String clientId = environment.get(\"AZURE_CLIENT_ID\");\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        if (clientId != null \u0026\u0026 !clientId.isBlank()) {\n+            credentialBuilder.clientId(clientId);\n+        }\n+        return new AzureConfiguration(\n+            required(environment, \"AZURE_STORAGE_BLOB_ENDPOINT\"),\n+            required(environment, \"AZURE_EVENTGRID_TOPIC_ENDPOINT\"),\n+            credentialBuilder.build());\n+    }\n+\n+    public BlobServiceClient blobServiceClient() {\n+        return blobClientBuilder().buildClient();\n+    }\n+\n+    public BlobServiceAsyncClient blobServiceAsyncClient() {\n+        return blobClientBuilder().buildAsyncClient();\n+    }\n+\n+    public EventPublisher eventPublisher() {\n+        return new EventPublisher(eventGridTopicEndpoint, credential);\n+    }\n+\n+    public AsyncEventPublisher asyncEventPublisher() {\n+        return new AsyncEventPublisher(eventGridTopicEndpoint, credential);\n+    }\n+\n+    private BlobServiceClientBuilder blobClientBuilder() {\n+        return new BlobServiceClientBuilder()\n+            .endpoint(storageEndpoint)\n+            .credential(credential);\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e environment, String name) {\n+        String value = environment.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Missing required environment variable \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static String requireHttpsEndpoint(String value, String name) {\n+        URI uri;\n+        try {\n+            uri = URI.create(Objects.requireNonNull(value, name));\n+        } catch (IllegalArgumentException exception) {\n+            throw new IllegalArgumentException(name + \" must be a valid URI\", exception);\n+        }\n+        if (!\"https\".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {\n+            throw new IllegalArgumentException(name + \" must be an absolute HTTPS endpoint\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/Main.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+import java.util.Map;\n+import java.util.logging.Logger;\n+import java.util.stream.StreamSupport;\n+\n+public final class Main {\n+    private static final Logger LOGGER = Logger.getLogger(Main.class.getName());\n+\n+    private static final String EVENT_GRID_PAYLOAD = \"\"\"\n+        [\n+          {\n+            \"id\": \"f5f0a761-8a0f-4f9d-bf19-8014f5997a4f\",\n+            \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/invoices/2026-08/invoice-1001.pdf\",\n+            \"eventTime\": \"2026-08-29T03:40:00Z\",\n+            \"data\": {\n+              \"api\": \"PutBlob\",\n+              \"clientRequestId\": \"4d9329b5-efec-4f3a-b68d-9f3eb31f7857\",\n+              \"requestId\": \"8caa1a6e-501e-0030-1d80-9d3a23000000\",\n+              \"eTag\": \"0x8DE000000000001\",\n+              \"contentType\": \"application/pdf\",\n+              \"contentLength\": 24576,\n+              \"blobType\": \"BlockBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/invoices/2026-08/invoice-1001.pdf\",\n+              \"sequencer\": \"000000000000000000000000000000A10000000000000123\"\n+            },\n+            \"dataVersion\": \"\",\n+            \"metadataVersion\": \"1\"\n+          },\n+          {\n+            \"id\": \"6f39f7e3-6368-45cd-8936-67bc4ea0376c\",\n+            \"eventType\": \"Microsoft.Storage.BlobDeleted\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf\",\n+            \"eventTime\": \"2026-08-29T03:41:00Z\",\n+            \"data\": {\n+              \"api\": \"DeleteBlob\",\n+              \"requestId\": \"c48e8542-301e-0059-2f80-9dd848000000\",\n+              \"contentType\": \"application/pdf\",\n+              \"blobType\": \"BlockBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/archive/old-invoice.pdf\",\n+              \"sequencer\": \"000000000000000000000000000000A20000000000000124\"\n+            },\n+            \"dataVersion\": \"\",\n+            \"metadataVersion\": \"1\"\n+          }\n+        ]\n+        \"\"\";\n+\n+    private static final String CLOUD_EVENTS_PAYLOAD = \"\"\"\n+        [\n+          {\n+            \"specversion\": \"1.0\",\n+            \"type\": \"Microsoft.Storage.BlobCreated\",\n+            \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/reports/quarterly-summary.csv\",\n+            \"id\": \"cc32dad8-5ed4-4e3f-b829-bd55d72fdf0a\",\n+            \"time\": \"2026-08-29T03:42:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+              \"api\": \"PutBlockList\",\n+              \"contentType\": \"text/csv\",\n+              \"contentLength\": 4096,\n+              \"blobType\": \"BlockBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/reports/quarterly-summary.csv\",\n+              \"sequencer\": \"000000000000000000000000000000A30000000000000125\"\n+            }\n+          },\n+          {\n+            \"specversion\": \"1.0\",\n+            \"type\": \"Microsoft.Storage.BlobDeleted\",\n+            \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example\",\n+            \"subject\": \"/blobServices/default/containers/documents/blobs/temp/upload.tmp\",\n+            \"id\": \"c18e061b-b927-48cb-9bda-4836d3424e0c\",\n+            \"time\": \"2026-08-29T03:43:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+              \"api\": \"DeleteBlob\",\n+              \"contentType\": \"application/octet-stream\",\n+              \"blobType\": \"BlockBlob\",\n+              \"url\": \"https://example.blob.core.windows.net/documents/temp/upload.tmp\",\n+              \"sequencer\": \"000000000000000000000000000000A40000000000000126\"\n+            }\n+          }\n+        ]\n+        \"\"\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        BlobEventHandler handler = offlineBlobHandler();\n+        List\u003cCustomEvent\u003e downstreamEvents = List.of(new CustomEvent(\n+            \"Contoso.Documents.DocumentProcessed\",\n+            \"/documents/invoices/processed\",\n+            Map.of(\"documentId\", \"invoice-1001\", \"status\", \"processed\"),\n+            \"1.0\"));\n+\n+        LOGGER.info(\"Running synchronous demo\");\n+        EventReceiver receiver = new EventReceiver();\n+        receiver.receive(EVENT_GRID_PAYLOAD, handler);\n+        receiver.receive(CLOUD_EVENTS_PAYLOAD, handler);\n+        new EventPublisher(events -\u003e logPublished(\"sync\", events)).publish(downstreamEvents);\n+\n+        LOGGER.info(\"Running asynchronous demo\");\n+        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver();\n+        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -\u003e {\n+            logPublished(\"async\", events);\n+            return Mono.empty();\n+        });\n+        asyncReceiver.receive(EVENT_GRID_PAYLOAD, handler)\n+            .then(asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD, handler))\n+            .then(asyncPublisher.publish(downstreamEvents))\n+            .block();\n+    }\n+\n+    private static BlobEventHandler offlineBlobHandler() {\n+        BlobEventHandler.SyncBlobReader syncReader = location -\u003e sampleSummary(location.name());\n+        BlobEventHandler.AsyncBlobReader asyncReader = location -\u003e Mono.just(sampleSummary(location.name()));\n+        return new BlobEventHandler(syncReader, asyncReader);\n+    }\n+\n+    private static BlobEventHandler.BlobSummary sampleSummary(String name) {\n+        String contentType = name.endsWith(\".pdf\") ? \"application/pdf\" : \"text/csv\";\n+        long size = name.endsWith(\".pdf\") ? 24_576 : 4_096;\n+        return new BlobEventHandler.BlobSummary(name, size, contentType, \"HOT\");\n+    }\n+\n+    private static void logPublished(String mode, Iterable\u003cEventGridEvent\u003e events) {\n+        StreamSupport.stream(events.spliterator(), false)\n+            .forEach(event -\u003e LOGGER.info(() -\u003e mode + \" published custom event: type=\"\n+                + event.getEventType() + \", subject=\" + event.getSubject()));\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/EventReceiverTest.java\n+package com.example.blobevents;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class EventReceiverTest {\n+    private final EventReceiver receiver = new EventReceiver();\n+\n+    @Test\n+    void deserializesEventGridSchema() {\n+        String payload = \"\"\"\n+            [{\n+              \"id\":\"1\",\n+              \"eventType\":\"Microsoft.Storage.BlobCreated\",\n+              \"subject\":\"/blobServices/default/containers/docs/blobs/a.pdf\",\n+              \"eventTime\":\"2026-08-29T03:40:00Z\",\n+              \"data\":{\"url\":\"https://example.blob.core.windows.net/docs/a.pdf\"},\n+              \"dataVersion\":\"\",\n+              \"metadataVersion\":\"1\"\n+            }]\n+            \"\"\";\n+\n+        BlobLifecycleEvent event = receiver.deserialize(payload).get(0);\n+\n+        assertEquals(EventSchema.EVENT_GRID, event.schema());\n+        assertEquals(\"Microsoft.Storage.BlobCreated\", event.type());\n+        assertEquals(\"1\", event.id());\n+    }\n+\n+    @Test\n+    void deserializesCloudEventsSchema() {\n+        String payload = \"\"\"\n+            [{\n+              \"specversion\":\"1.0\",\n+              \"type\":\"Microsoft.Storage.BlobDeleted\",\n+              \"source\":\"/storageAccounts/example\",\n+              \"subject\":\"/blobServices/default/containers/docs/blobs/a.pdf\",\n+              \"id\":\"2\",\n+              \"time\":\"2026-08-29T03:40:00Z\",\n+              \"datacontenttype\":\"application/json\",\n+              \"data\":{\"url\":\"https://example.blob.core.windows.net/docs/a.pdf\"}\n+            }]\n+            \"\"\";\n+\n+        BlobLifecycleEvent event = receiver.deserialize(payload).get(0);\n+\n+        assertEquals(EventSchema.CLOUD_EVENTS, event.schema());\n+        assertEquals(\"Microsoft.Storage.BlobDeleted\", event.type());\n+        assertEquals(\"2\", event.id());\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/BlobEventHandlerTest.java\n+package com.example.blobevents;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class BlobEventHandlerTest {\n+    @Test\n+    void parsesAndDecodesBlobSubject() {\n+        BlobEventHandler.BlobLocation location = BlobEventHandler.parseSubject(\n+            \"/blobServices/default/containers/my-container/blobs/folder%20one/report%2Bfinal.pdf\");\n+\n+        assertEquals(\"my-container\", location.container());\n+        assertEquals(\"folder one/report+final.pdf\", location.name());\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/EventPublisherTest.java\n+package com.example.blobevents;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.List;\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+class EventPublisherTest {\n+    @Test\n+    void preservesSubjectHierarchy() {\n+        CustomEvent customEvent = new CustomEvent(\n+            \"Contoso.Documents.Processed\",\n+            \"/documents/invoices/processed\",\n+            Map.of(\"id\", \"42\"),\n+            \"1.0\");\n+\n+        var event = EventPublisher.toEventGridEvents(List.of(customEvent)).get(0);\n+\n+        assertEquals(\"/documents/invoices/processed\", event.getSubject());\n+        assertEquals(\"Contoso.Documents.Processed\", event.getEventType());\n+    }\n+\n+    @Test\n+    void rejectsRelativeSubject() {\n+        assertThrows(IllegalArgumentException.class, () -\u003e new CustomEvent(\n+            \"Contoso.Documents.Processed\",\n+            \"documents/processed\",\n+            Map.of(),\n+            \"1.0\"));\n+    }\n+}\n*** Add File: README.md\n+# Azure Blob Event Notifier\n+\n+Small Java 17 Maven sample that receives Azure Blob Storage lifecycle events in either Event Grid schema or\n+CloudEvents 1.0 schema, handles blob-created and blob-deleted events, and publishes downstream custom events.\n+\n+The `Main` demo is fully local and uses injected in-memory adapters. Production constructors use\n+`ManagedIdentityCredential`; no keys, connection strings, or SAS tokens are used.\n+\n+## Run the offline demo\n+\n+```powershell\n+mvn test\n+mvn exec:java\n+```\n+\n+## Production configuration\n+\n+Set these environment variables before constructing `AzureConfiguration.fromEnvironment()`:\n+\n+| Variable | Purpose |\n+|---|---|\n+| `AZURE_STORAGE_BLOB_ENDPOINT` | Storage endpoint, such as `https://account.blob.core.windows.net` |\n+| `AZURE_EVENTGRID_TOPIC_ENDPOINT` | Full custom topic endpoint |\n+| `AZURE_CLIENT_ID` | Optional client ID for a user-assigned managed identity |\n+\n+The managed identity needs **Storage Blob Data Reader** on the required storage scope and an Event Grid data-plane\n+sender role, such as **EventGrid Data Sender**, on the custom topic scope.\n+\n+```java\n+AzureConfiguration configuration = AzureConfiguration.fromEnvironment();\n+BlobEventHandler handler = new BlobEventHandler(\n+    configuration.blobServiceClient(),\n+    configuration.blobServiceAsyncClient());\n+\n+new EventReceiver().receive(webhookJson, handler);\n+configuration.eventPublisher().publish(customEvents);\n+```\n+\n+The receiver expects one delivery payload to use one schema, matching Event Grid subscription configuration.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml
@@ -1,0 +1,78 @@
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
+        <azure.sdk.bom.version>1.3.8</azure.sdk.bom.version>
+        <junit.version>5.13.4</junit.version>
+    </properties>
+
+    <dependencyManagement>
+        <dependencies>
+            <dependency>
+                <groupId>com.azure</groupId>
+                <artifactId>azure-sdk-bom</artifactId>
+                <version>${azure.sdk.bom.version}</version>
+                <type>pom</type>
+                <scope>import</scope>
+            </dependency>
+        </dependencies>
+    </dependencyManagement>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-eventgrid</artifactId>
+        </dependency>
+        <dependency>
+            <
... (truncated)
```

### ✅ 23. powershell — 8569ms

**Input:**

```json
{"command":"mvn -q test; if ($LASTEXITCODE -eq 0) { mvn -q exec:java }","description":"Compile, test, and run the offline Maven demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
Aug 29, 2026 11:49:54 AM com.example.blobevents.Main main
INFO: Running synchronous demo
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
Aug 29, 2026 11:49:55 AM com.example.blobevents.BlobEventHandler logSummary
INFO: Blob created: name=invoices/2026-08/invoice-1001.pdf, size=24576, contentType=application/pdf, accessTier=HOT
Aug 29, 2026 11:49:55 AM com.example.blobevents.BlobEventHandler handleDeleted
INFO: Blob deleted: container=documents, name=archive/old-invoice.pdf
Aug 29, 2026 11:49:55 AM com.example.blobevents.BlobEventHandler logSummary
INFO: Blob created: name=reports/quarterly-summary.csv, size=4096, contentType=text/csv, accessTier=HOT
Aug 29, 2026 11:49:55 AM com.example.blobevents.BlobEventHandler handleDeleted
INFO: Blo
... (truncated)
```

### ✅ 24. apply_patch — 644ms

**Input:**

```json
"*** Begin Patch\n*** Update File: pom.xml\n@@\n         \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n         \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n+        \u003cslf4j.version\u003e1.7.36\u003c/slf4j.version\u003e\n@@\n         \u003cdependency\u003e\n             \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n             \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n         \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n*** Update File: src/main/java/com/example/blobevents/BlobEventHandler.java\n@@\n         BlobClient blob = serviceClient.getBlobContainerClient(location.container()).getBlobClient(location.name());\n-        BinaryData ignoredContent = blob.downloadContent();\n+        blob.downloadContent();\n         BlobProperties properties = blob.getProperties();\n*** Update File: src/main/java/com/example/blobevents/AsyncEventPublisher.java\n@@\n     public Mono\u003cVoid\u003e publish(List\u003cCustomEvent\u003e events) {\n-        return sender.send(EventPublisher.toEventGridEvents(events));\n+        return Mono.defer(() -\u003e sender.send(EventPublisher.toEventGridEvents(events)));\n     }\n*** Update File: src/test/java/com/example/blobevents/EventPublisherTest.java\n@@\n import java.util.List;\n import java.util.Map;\n+import java.util.concurrent.atomic.AtomicBoolean;\n@@\n     void rejectsRelativeSubject() {\n         assertThrows(IllegalArgumentException.class, () -\u003e new CustomEvent(\n             \"Contoso.Documents.Processed\",\n@@\n             \"1.0\"));\n     }\n+\n+    @Test\n+    void asyncPublishIsLazy() {\n+        AtomicBoolean called = new AtomicBoolean();\n+        AsyncEventPublisher publisher = new AsyncEventPublisher(events -\u003e {\n+            called.set(true);\n+            return reactor.core.publisher.Mono.empty();\n+        });\n+        CustomEvent event = new CustomEvent(\n+            \"Contoso.Documents.Processed\",\n+            \"/documents/processed\",\n+            Map.of(),\n+            \"1.0\");\n+\n+        var publication = publisher.publish(List.of(event));\n+        assertEquals(false, called.get());\n+\n+        publication.block();\n+        assertEquals(true, called.get());\n+    }\n }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/pom.xml
@@ -14,6 +14,7 @@
         <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
         <azure.sdk.bom.version>1.3.8</azure.sdk.bom.version>
         <junit.version>5.13.4</junit.version>
+        <slf4j.version>1.7.36</slf4j.version>
     </properties>
 
     <dependencyManagement>
@@ -42,6 +43,12 @@
             <artifactId>azure-storage-blob</artifactId>
         </dependency>
         <dependency>
+            <groupId>org.slf4j</groupId>
+            <artifactId>slf4j-simple</artifactId>
+            <version>${slf4j.version}</version>
+            <scope>runtime</scope>
+        </dependency>
+        <dependency>
             <groupId>org.junit.jupiter</groupId>
             <artifactId>junit-jupiter</artifactId>
             <version>${junit.version}</version>


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/src/main/java/com/example/blobevents/BlobEventHandler.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/src/main/java/com/example/blobevents/BlobEventHandler.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-blob-event-notifier-java-azure-skills-azure-skill-mcp-3863101358/src/main/java/com/example/blobevents/BlobEventHandler.java
+++ b/C:/Users/weidxu/AppData/Lo
... (truncated)
```

### ✅ 25. powershell — 8250ms

**Input:**

```json
{"command":"mvn -q test; if ($LASTEXITCODE -eq 0) { mvn -q exec:java }","description":"Re-run tests and the corrected offline demo","initial_wait":180,"mode":"sync"}
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
Aug 29, 2026 11:50:19 AM com.example.blobevents.Main main
INFO: Running synchronous demo
Aug 29, 2026 11:50:19 AM com.example.blobevents.BlobEventHandler logSummary
INFO: Blob created: name=invoices/2026-08/invoice-1001.pdf, size=24576, contentType=application/pdf, accessTier=HOT
Aug 29, 2026 11:50:19 AM com.example.blobevents.BlobEventHandler handleDeleted
INFO: Blob deleted: container=documents, name=archive/old-invoice.pdf
Aug 29, 2026 11:50:19 AM com.example.blobevents.BlobEventHandler logSummary
INFO: Blob created: name=reports/quarterly-summary.csv, size=4096, contentType=text/csv, accessTier=HOT
Aug 29, 2026 11:50:19 AM com.example.blobevents.BlobEventHandler handleDeleted
INFO: Blob deleted: container=documents, name=temp/upload.tmp
Aug 29, 2026 11:50:19 AM com.example.blobevents.Main lambda$logPublished$0
INFO: sync published custom event: type=Contoso.Documents.DocumentProcessed, subject=/documents/invoices/processed
Aug 29, 2026 11:50:19 AM com.example.blobevents.Main main
INFO: Running asynchronous demo
Aug 29, 2026 11:50:19 AM com.example.blobevents.BlobEventHandler logSummary
INFO: Blob 
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\blobevents\AsyncEventPublisher.java`
- `src\main\java\com\example\blobevents\AsyncEventReceiver.java`
- `src\main\java\com\example\blobevents\AzureConfiguration.java`
- `src\main\java\com\example\blobevents\BlobEventHandler.java`
- `src\main\java\com\example\blobevents\BlobLifecycleEvent.java`
- `src\main\java\com\example\blobevents\CustomEvent.java`
- `src\main\java\com\example\blobevents\EventPayloadDeserializer.java`
- `src\main\java\com\example\blobevents\EventPublisher.java`
- `src\main\java\com\example\blobevents\EventReceiver.java`
- `src\main\java\com\example\blobevents\EventSchema.java`
- `src\main\java\com\example\blobevents\Main.java`
- `src\test\java\com\example\blobevents\BlobEventHandlerTest.java`
- `src\test\java\com\example\blobevents\EventPublisherTest.java`
- `src\test\java\com\example\blobevents\EventReceiverTest.java`

## Copilot Response

I’m grounding the project in the current Azure Java SDK patterns, then I’ll scaffold and run the local Maven demo without contacting Azure.

The workspace is empty. I’m confirming the Event Grid model and publisher APIs before generating the project, including an injectable local demo path so the sample runs offline while production constructors use managed identity.

The API shape is confirmed: both schemas deserialize through Azure SDK model factories, and the Event Grid publisher builder accepts `TokenCredential`. I’m now adding the complete project with real managed-identity adapters plus deterministic offline demo adapters.

The project runs successfully. One reactive ordering issue surfaced in the demo: the async publisher’s injected sender was invoked during pipeline assembly rather than subscription, so I’m making publication properly lazy and removing Azure SDK logging warnings.

Created the complete Java 17 Maven project in the current directory.

It includes sync/async Event Grid receivers and publishers, Event Grid and CloudEvents 1.0 deserialization, blob lifecycle routing, managed-identity configuration, race-condition handling, subject hierarchy support, offline demo adapters, tests, and documentation. Run it with:

```powershell
mvn test
mvn exec:java
```

The demo runs synchronously first, then asynchronously, without contacting Azure. Production constructors use `ManagedIdentityCredential`.

References: [Azure Event Grid Java SDK](https://learn.microsoft.com/java/api/overview/azure/messaging-eventgrid-readme?view=azure-java-stable) and [Azure Blob Storage Java SDK](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- AutoCloseable resources are handled correctly, and Azure SDK clients without close semantics are not unnecessarily wrapped or closed.

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
  - Azure SDK BOM for Version Management (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Pass
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
| `Criteria from prompt file` | prompt_review | 90% | 1.00 | 0.9000 | 7.0% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 7.8% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.9000** | **99.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-blob-event-notifier --config java-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

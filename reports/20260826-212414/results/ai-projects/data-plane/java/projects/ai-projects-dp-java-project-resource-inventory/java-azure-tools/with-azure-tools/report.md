# Evaluation Report: ai-projects-dp-java-project-resource-inventory

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 434.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-projects-dp-java-project-resource-inventory` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 18/21 |
| Duration | 434.8s |
| Timestamp | 2026-08-26T13:24:14Z |
| Files Generated | 3 |
| Event Count | 3675 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 138.6s |
| Review | 294.3s |
| **Total** | **434.8s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-ai-projects-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=364551 out=3658 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill

**Input:**

```json
{"skill":"azure-ai-projects-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure AI Projects SDK for Java

High-level SDK for Azure AI Foundry project management with access to connections, datasets, indexes, and evaluations.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-ai-projects</artifactId>
    <version>1.0.0-beta.1</version>
</dependency>
```

## Environment Variables

```bash
PROJECT_ENDPOINT=https://<resource>.services.ai.azure.com/api/projects/<project> # Required for project configuration
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
```

## Authentication

```java
import com.azure.ai.projects.AIProjectClientBuilder;
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

AIProjectClientBuilder builder = new AIProjectClientBuilder()
    .endpoint(System.getenv("PROJECT_ENDPOINT"))
    .credential(credential);
```

## Client Hierarchy

The SDK provides multiple sub-clients for different operations:

| Client | Purpose |
|--------|---------|
| `ConnectionsClient` | Enumerate connected Azure resources |
| `DatasetsClient` | Upload documents and manage datasets |
| `DeploymentsClient` | Enumerate AI model deployments |
| `IndexesClient` | Create and manage search indexes |
| `EvaluationsClient` | Run AI model evaluations |
| `EvaluatorsClient` | Manage evaluator configurations |
| `SchedulesClient` | Manage scheduled operations |

```java
// Build sub-clients 
... (truncated)
```

### ✅ 2. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-ai-projects-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-ai-projects-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-ai-projects-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-ai-projects-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-ai-projects-java/references/examples.md
@@ -1,345 +1,345 @@
 # Azure AI Projects Java SDK - Examples
 
 Comprehensive code examples for the Azure AI Projects SDK for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Working with Connections](#working-with-connections)
 - [Working with Deployments](#working-with-deployments)
 - [Working with Datasets](#working-with-datasets)
 - [Async Clients](#async-clients)
 
 ---
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-ai-projects</artifactId>
     <version>1.0.0-beta.1</version>
 </dependency>
 
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.15.3</version>
 </dependency>
 ```
 
 ---
 
 ## Client Creation
 
 ### Creating AIProjectClient and Sub-Clients
 
 ```java
 import com.azure.ai.projects.*;
 import com.azure.core.util.Configuration;
 import com.azure.identity.DefaultAzureCredentialBuilder;
 
 public class ClientInitializationSample {
 
     public static void main(String[] args) {
         // Create the builder with endpoint and credentials
         AIProjectClientBuilder builder = new AIProjectClientBuilder()
             .endpoint(Configuration.getGlobalConfiguration().get("AZURE_AI_PROJECTS_ENDPOINT"))
             .credential(new DefaultAzureCredentia
... (truncated)
```

### ✅ 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-java-project-resource-inventory-java-azure-tools-with-azure-tools-2919423923","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Get Java code-generation best practices for a local console application using Azure AI Projects and DefaultAzureCredential; no Azure resources will be modified.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Java SDK documentation for com.azure:azure-ai-projects synchronous ConnectionsClient and DeploymentsClient list/get APIs and model types.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_ai_app","intent":"Apply Microsoft Foundry Java application code-generation best practices.","parameters":{}}
```

**Output:**

```
{"status":200,"message":"","results":["# Best Practices for Building AI Apps with Azure AI Services\r\n\r\nThis guide covers best practices for:\r\n- Generating code for AI agent applications using Microsoft Foundry and any other Azure AI models\r\n- Building agents with Microsoft Foundry\r\n- Implementing chatbots, workflows, assistants, and multi-agent systems with Microsoft Agent Framework\r\n\r\n> Note: Understanding User Intent of Microsoft Foundry\r\n> \r\n> **Two Types of Operations:**\r\n> \r\n> **1. Resource Management** - Use `foundry` tool\r\n> - Keywords: \"create/deploy/provision agent/embedding IN Foundry\", \"set up agent/embedding resource\"\r\n> - Creates Microsoft Foundry agent resource via commands under `foundry`\r\n> \r\n> **2. Application Development** - Use this best practices tool\r\n> - Keywords: \"build/write/implement agent\", \"develop chatbot/assistant\", \"agent code\"\r\n> - Generates code using Microsoft Agent Framework SDK\r\n> - Returns code samples and implementation patterns\r\n> \r\n> **When unsure, ask:** \"Do you want to create an agent resource in Microsoft Foundry, or generate agent application code?\"\r\n\r\n## Key Principles\r\n**Defaults and Recommendations:**\r\n- **Default language**: If the user does not specify a programming language, use **.NET** as the default programming language.\r\n- **Recommended SDK**: Use **Microsoft Agent Framework** for building AI agents, chatbots, assistants, and multi-agent systems - it provides flexible orchestration, multi-agent patterns, and cross-platform support (.NET and Python)\r\n- **Select model first**: Model selection is REQUIRED for ANY AI implementation - always select the appropriate AI model before coding if the user does not specify a model\r\n\r\n**Process:**\r\n1. **Gather information first**: Always search / fetch / query for sufficient knowledge (quickstarts, tutorials, feature guides, code samples) BEFORE thinking or writing any code - this ensures high-quality code ge
... (truncated)
```

### ✅ 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Locate current official Java SDK API details.","parameters":{"query":"azure-ai-projects Java ConnectionsClient DeploymentsClient ModelDeployment getConnection get deployment"}}
```

**Output:**

```
{"results":[{"title":"Azure Projects client library for Java - version 2.4.0","content":"# Azure Projects client library for Java - version 2.4.0\n## Examples\n\u0060\u0060\u0060java\n\nString connectionName = Configuration.getGlobalConfiguration().get(\u0022TEST_CONNECTION_NAME\u0022, \u0022\u0022);\nreturn connectionsAsyncClient.getConnectionWithCredentials(connectionName)\n    .doOnNext(connection -\u003E {\n        System.out.printf(\u0022Connection name: %s%n\u0022, connection.getName());\n        System.out.printf(\u0022Connection credentials: %s%n\u0022, connection.getCredential().getType());\n    });\n\u0060\u0060\u0060\n### Deployments operations\nDeployment operations allow you to enumerate and inspect the models deployed to your AI Foundry Project.\n#### List deployments\n\u0060\u0060\u0060java\n\nPagedIterable\u003CDeployment\u003E deployments = deploymentsClient.listDeployments();\nfor (Deployment deployment : deployments) {\n    System.out.printf(\u0022Deployment name: %s%n\u0022, deployment.getName());\n}\n\u0060\u0060\u0060\n#### Get a deployment\n\u0060\u0060\u0060java\n\nString deploymentName = Configuration.getGlobalConfiguration().get(\u0022FOUNDRY_MODEL_NAME\u0022, \u0022\u0022);\nDeployment deployment = deploymentsClient.getDeployment(deploymentName);\n\nSystem.out.printf(\u0022Deployment name: %s%n\u0022, deployment.getName());\nSystem.out.printf(\u0022Deployment type: %s%n\u0022, deployment.getType().getValue());\n\u0060\u0060\u0060\n#### Asynchronous deployment operations\n\u0060\u0060\u0060java\n\nreturn deploymentsAsyncClient.listDeployments()\n    .doOnNext(deployment -\u003E System.out.printf(\u0022Deployment name: %s%n\u0022, deployment.getName()));\n\u0060\u0060\u0060\n\u0060\u0060\u0060java\n\nString deploymentName = Configuration.getGlobalConfiguration().get(\u0022FOUNDRY_MODEL_NAME\u0022, \u0022\u0022);\nreturn deploymentsAsyncClient.getDeployment(deploymentName)\n    .doOnNext(deployment -\u003E {\n        System.out.printf(\u0022D
... (truncated)
```

### ✅ 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Locate current official Java SDK samples.","parameters":{"language":"java","query":"com.azure.ai.projects ConnectionsClient DeploymentsClient ModelDeployment synchronous list get connection without credentials"}}
```

**Output:**

```
{"results":[{"description":"description: Creates and authenticates a PersistentAgentsClient and related clients for interacting with various resources of an AI project, using environment variables for configuration and DefaultAzureCredential for authentication.\npackage: com.azure.ai.agents.persistent, com.azure.core.util, com.azure.identity\nlanguage: java\n","codeSnippet":"import com.azure.ai.agents.persistent.FilesClient;\r\nimport com.azure.ai.agents.persistent.MessagesClient;\r\nimport com.azure.ai.agents.persistent.PersistentAgentsAdministrationClient;\r\nimport com.azure.ai.agents.persistent.PersistentAgentsClient;\r\nimport com.azure.ai.agents.persistent.PersistentAgentsClientBuilder;\r\nimport com.azure.ai.agents.persistent.RunsClient;\r\nimport com.azure.ai.agents.persistent.ThreadsClient;\r\nimport com.azure.ai.agents.persistent.VectorStoresClient;\r\nimport com.azure.core.util.Configuration;\r\nimport com.azure.identity.DefaultAzureCredentialBuilder;\r\n\r\n// Get Connection information from environment variables\r\nString projectEndpoint = Configuration.getGlobalConfiguration().get(\u0022PROJECT_ENDPOINT\u0022, \u0022endpoint\u0022);\r\nString modelDeploymentName = Configuration.getGlobalConfiguration().get(\u0022MODEL_DEPLOYMENT_NAME\u0022, \u0022model\u0022);\r\n\r\n// Create the Agent Client\r\nPersistentAgentsClientBuilder clientBuilder = new PersistentAgentsClientBuilder()\r\n    .endpoint(projectEndpoint)\r\n    .credential(new DefaultAzureCredentialBuilder().build());\r\nPersistentAgentsClient agentsClient = clientBuilder.buildClient();\r\nPersistentAgentsAdministrationClient administrationClient = agentsClient.getPersistentAgentsAdministrationClient();\r\nThreadsClient threadsClient = agentsClient.getThreadsClient();\r\nMessagesClient messagesClient = agentsClient.getMessagesClient();\r\nRunsClient runsClient = agentsClient.getRunsClient();\r\nFilesClient filesClient = agentsClient.getFilesClient();\r\nVectorStoresClient vectorStoresClient = age
... (truncated)
```

### ✅ 9. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","type"],"owner":"Azure","path":"sdk/ai/azure-ai-projects","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
[{"name":"CHANGELOG.md","path":"sdk/ai/azure-ai-projects/CHANGELOG.md","type":"file"},{"name":"README.md","path":"sdk/ai/azure-ai-projects/README.md","type":"file"},{"name":"assets.json","path":"sdk/ai/azure-ai-projects/assets.json","type":"file"},{"name":"checkstyle-suppressions.xml","path":"sdk/ai/azure-ai-projects/checkstyle-suppressions.xml","type":"file"},{"name":"customizations","path":"sdk/ai/azure-ai-projects/customizations","type":"dir"},{"name":"pom.xml","path":"sdk/ai/azure-ai-projects/pom.xml","type":"file"},{"name":"revapi.json","path":"sdk/ai/azure-ai-projects/revapi.json","type":"file"},{"name":"src","path":"sdk/ai/azure-ai-projects/src","type":"dir"},{"name":"tsp-location.yaml.hide","path":"sdk/ai/azure-ai-projects/tsp-location.yaml.hide","type":"file"}]
```

### ✅ 10. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"class ModelDeployment repo:Azure/azure-sdk-for-java path:sdk/ai/azure-ai-projects"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 11. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":20,"query":"getConnectionWithCredentials repo:Azure/azure-sdk-for-java path:sdk/ai/azure-ai-projects/src/samples"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ConnectionsSample.java","path":"sdk/ai/azure-ai-projects/src/samples/java/com/azure/ai/projects/ConnectionsSample.java","sha":"d8ae3d21c1492cee0595a962db27937c4df2844c","text_matches":[{"fragment":"        // END:com.azure.ai.projects.ConnectionsSample.getConnectionWithoutCredentials\n    }\n\n    public static void getConnectionWithCredentials() {\n        // BEGIN:com.azure.ai.projects.ConnectionsSample.getConnectionWithCredentials\n\n        String connectionName = Configuration.getGlobalConfiguration().get(\"TEST_CONNECTION_NAME\", \"\");","matches":[{"indices":[117,145],"text":"getConnectionWithCredentials"},{"indices":[207,235],"text":"getConnectionWithCredentials"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/ai/azure-ai-projects/src/samples/java/com/azure/ai/projects/ConnectionsSample.java?ref=16d590f580d71799f31ad5f392cdeefadc3d52a0","property":"content"}]},{"name":"ConnectionsAsyncSample.java","path":"sdk/ai/azure-ai-projects/src/samples/java/com/azure/ai/projects/ConnectionsAsyncSample.java","sha":"30012b484eefbecb13d754bb1f519373149eeb53","text_matches":[{"fragment":"        // END:com.azure.ai.projects.ConnectionsAsyncSample.getConnectionWithoutCredentials\n    }\n\n    public static Mono<Connection> getConnectionWithCredentials() {\n        // BEGIN:com.azure.ai.projects.ConnectionsAsyncSample.getConnectionWithCredentials\n\n        String connectionName = Configuration.getGlobalConfiguration().get(\"TEST_CONNECTION_NAME\", \"\");","matches":[{"indices":[134,162],"text":"getConnectionWithCredentials"},{"indices":[229,257],"text":"getConnectionWithCredentials"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/ai/azure-ai-projects/src/samples/java/com/azure/ai/projects/ConnectionsAsyncSample.java?ref=16d590f580d71799f31ad5f392cdeefadc3d52a0","property":"content"}]}],"total_count":2}
```

### ✅ 12. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":20,"query":"instanceof ModelDeployment repo:Azure/azure-sdk-for-java path:sdk/ai/azure-ai-projects"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ModelDeployment.java","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/ModelDeployment.java","sha":"0c949b507dd61a1e2681ce2fec9a4f1a7eac58fd","text_matches":[{"fragment":" * Model Deployment Definition.\n */\n@Immutable\npublic final class ModelDeployment extends Deployment {\n\n    /*\n     * The type of the deployment","matches":[{"indices":[60,65],"text":"class"},{"indices":[66,81],"text":"ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/ModelDeployment.java?ref=16d590f580d71799f31ad5f392cdeefadc3d52a0","property":"content"}]},{"name":"ModelDeploymentSku.java","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/ModelDeploymentSku.java","sha":"ba16f14f50122f88a47c03ad74d53c7ed5490932","text_matches":[{"fragment":" * Sku information.\n */\n@Immutable\npublic final class ModelDeploymentSku implements JsonSerializable<ModelDeploymentSku> {\n\n    /*\n     * Sku capacity","matches":[{"indices":[48,53],"text":"class"},{"indices":[54,69],"text":"ModelDeployment"},{"indices":[101,116],"text":"ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/ModelDeploymentSku.java?ref=16d590f580d71799f31ad5f392cdeefadc3d52a0","property":"content"}]},{"name":"DeploymentType.java","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/DeploymentType.java","sha":"b3c3d57f693123db821ccb3d45647cd3c1d4bce8","text_matches":[{"fragment":" */\npublic final class DeploymentType extends ExpandableStringEnum<DeploymentType> {\n","matches":[{"indices":[17,22],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/pro
... (truncated)
```

### ✅ 13. powershell

**Input:**

```json
{"command":"$ErrorActionPreference='Continue'; java -version; mvn -version","description":"Check installed Java and Maven versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
openjdk version "25.0.1" 2025-10-21 LTS
OpenJDK Runtime Environment Microsoft-12574223 (build 25.0.1+8-LTS)
OpenJDK 64-Bit Server VM Microsoft-12574223 (build 25.0.1+8-LTS, mixed mode, sharing)
Apache Maven 3.9.11 (3e54c93a704957b63ee3494413a2b544fd3d825b)
Maven home: C:\apache-maven-3.9.11
Java version: 25.0.1, vendor: Microsoft, runtime: C:\Program Files\Microsoft\jdk-25.0.1.8-hotspot
Default locale: en_US, platform encoding: UTF-8
OS name: "windows 11", version: "10.0", arch: "amd64", family: "windows"
<shellId: 0 completed with exit code 0>
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/Connection.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 2a97028bc8d4feb05fbc2c52c980b66f0ec02a68)<!--
 ~ Copyright (c) Microsoft Corporation. All rights reserved.
Licensed under the MIT License.
Code generated by Microsoft (R) TypeSpec Code Generator.
 -->
<project xmlns="http://maven.apache.org/POM/4.0.0" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <modelVersion>4.0.0</modelVersion>
  <parent>
    <groupId>com.azure</groupId>
    <artifactId>azure-client-sdk-parent</artifactId>
    <version>1.7.0</version> <!-- {x-version-update;com.azure:azure-client-sdk-parent;current} -->
    <relativePath>../../parents/azure-client-sdk-parent</relativePath>
  </parent>

  <groupId>com.azure</groupId>
  <artifactId>azure-ai-projects</artifactId>
  <version>2.5.0-beta.1</version> <!-- {x-version-update;com.azure:azure-ai-projects;current} -->
  <packaging>jar</packaging>

  <name>Microsoft Azure SDK for Projects</name>
  <description>This package contains Microsoft Azure Projects client library.</description>
  <url>https://github.com/Azure/azure-sdk-for-java</url>

  <licenses>
    <license>
      <name>The MIT License (MIT)</name>
      <url>http://opensource.org/licenses/MIT</url>
      <distribution>repo</distribution>
    </license>
  </licenses>

  <scm>
    <url>https://github.com/Azure/azure-sdk-for-java</url>
    <connection>scm:git:git@github.com:Azure/azure-sdk-for-java.git</connection>
    <developerConnection>scm:git:git@github.com:Azure/azure-sdk-for-java.git</developerConnection>
    <tag>HEAD</tag>
  </scm>
  <developers>
    <developer>
      <id>microsoft</id>
      <name>Microsoft</name>
    </developer>
  </developers>


  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>

    <jacoco.min.linecoverage>0.0</jacoco.min.linecoverage>
    <jacoco.min.branchcoverage>0.0</jacoco.min.branchcoverage>
    <javaModulesSurefireArgLine>
      -
... (truncated)
```

### ✅ 15. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/models/ModelDeployment.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: d8ae3d21c1492cee0595a962db27937c4df2844c)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package com.azure.ai.projects;

import com.azure.ai.projects.models.Connection;
import com.azure.core.http.rest.PagedIterable;
import com.azure.core.util.Configuration;
import com.azure.identity.DefaultAzureCredentialBuilder;

public class ConnectionsSample {

    private static ConnectionsClient connectionsClient
        = new AIProjectClientBuilder().endpoint(Configuration.getGlobalConfiguration().get("FOUNDRY_PROJECT_ENDPOINT", "endpoint"))
        .credential(new DefaultAzureCredentialBuilder().build())
        .buildConnectionsClient();

    public static void main(String[] args) {

//        listConnections();
//        getConnectionWithoutCredentials();
//        getConnectionWithCredentials();
    }

    public static void listConnections() {
        // BEGIN:com.azure.ai.projects.ConnectionsSample.listConnections
        PagedIterable<Connection> connections = connectionsClient.listConnections();
        for (Connection connection : connections) {
            System.out.println("Connection name: " + connection.getName());
            System.out.println("Connection type: " + connection.getType());
            System.out.println("Connection credential type: " + connection.getCredential().getType());
            System.out.println("-------------------------------------------------");
        }
        // END:com.azure.ai.projects.ConnectionsSample.listConnections
    }

    public static void getConnectionWithoutCredentials() {
        // BEGIN:com.azure.ai.projects.ConnectionsSample.getConnectionWithoutCredentials

        String connectionName = Configuration.getGlobalConfiguration().get("TEST_CONNECTION_NAME", "");
        Connection connection = connectionsClient.getConnection(connectionName);

        System.out.printf("Connection name: %s%n", connection.getName());

        // END
... (truncated)
```

### ✅ 16. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/ConnectionsClient.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 89352bb0cbc2d054bd18d9c25cab0f59c38d74e4)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// Code generated by Microsoft (R) TypeSpec Code Generator.
package com.azure.ai.projects;

import com.azure.ai.projects.implementation.ConnectionsImpl;
import com.azure.ai.projects.models.Connection;
import com.azure.ai.projects.models.ConnectionType;
import com.azure.core.annotation.Generated;
import com.azure.core.annotation.ReturnType;
import com.azure.core.annotation.ServiceClient;
import com.azure.core.annotation.ServiceMethod;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceModifiedException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.http.rest.PagedIterable;
import com.azure.core.http.rest.RequestOptions;
import com.azure.core.http.rest.Response;
import com.azure.core.util.BinaryData;
import com.azure.core.util.logging.ClientLogger;

/**
 * Initializes a new instance of the synchronous AIProjectClient type.
 */
@ServiceClient(builder = AIProjectClientBuilder.class)
public final class ConnectionsClient {

    private static final ClientLogger LOGGER = new ClientLogger(ConnectionsClient.class);

    @Generated
    private final ConnectionsImpl serviceClient;

    /**
     * Initializes an instance of ConnectionsClient class.
     *
     * @param serviceClient the service client implementation.
     */
    @Generated
    ConnectionsClient(ConnectionsImpl serviceClient) {
        this.serviceClient = serviceClient;
    }

    /**
     * Get a connection
     *
     * Retrieves the specified connection and its configuration details without including credential values.
     * <p><strong>Response Body Schema</strong></p>
     * 
     * <pre>
     * {@code
     * {
     *     name: String (Required)
     *     id: String (Required)
     *     type
... (truncated)
```

### ✅ 17. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/src/main/java/com/azure/ai/projects/DeploymentsClient.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 25d4d04e434f937fa9a7994801a86052813ef23a)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// Code generated by Microsoft (R) TypeSpec Code Generator.
package com.azure.ai.projects;

import com.azure.ai.projects.implementation.DeploymentsImpl;
import com.azure.ai.projects.models.Deployment;
import com.azure.ai.projects.models.DeploymentType;
import com.azure.core.annotation.Generated;
import com.azure.core.annotation.ReturnType;
import com.azure.core.annotation.ServiceClient;
import com.azure.core.annotation.ServiceMethod;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceModifiedException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.http.rest.PagedIterable;
import com.azure.core.http.rest.RequestOptions;
import com.azure.core.http.rest.Response;
import com.azure.core.util.BinaryData;

/**
 * Initializes a new instance of the synchronous AIProjectClient type.
 */
@ServiceClient(builder = AIProjectClientBuilder.class)
public final class DeploymentsClient {

    @Generated
    private final DeploymentsImpl serviceClient;

    /**
     * Initializes an instance of DeploymentsClient class.
     *
     * @param serviceClient the service client implementation.
     */
    @Generated
    DeploymentsClient(DeploymentsImpl serviceClient) {
        this.serviceClient = serviceClient;
    }

    /**
     * Get a deployment
     *
     * Retrieves a deployed model.
     * <p><strong>Response Body Schema</strong></p>
     * 
     * <pre>
     * {@code
     * {
     *     type: String(ModelDeployment) (Required)
     *     name: String (Required)
     * }
     * }
     * </pre>
     * 
     * <p><strong>Response Headers</strong></p>
     * <table border="1">
     * <caption>Response Headers</caption>
     * <tr><th>Name</th><th>Type</th><th>Description<
... (truncated)
```

### ✅ 18. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/src/samples/java/com/azure/ai/projects/ConnectionsSample.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 0c949b507dd61a1e2681ce2fec9a4f1a7eac58fd)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// Code generated by Microsoft (R) TypeSpec Code Generator.
package com.azure.ai.projects.models;

import com.azure.core.annotation.Generated;
import com.azure.core.annotation.Immutable;
import com.azure.json.JsonReader;
import com.azure.json.JsonToken;
import com.azure.json.JsonWriter;
import java.io.IOException;
import java.util.Map;

/**
 * Model Deployment Definition.
 */
@Immutable
public final class ModelDeployment extends Deployment {

    /*
     * The type of the deployment
     */
    @Generated
    private DeploymentType type = DeploymentType.MODEL_DEPLOYMENT;

    /*
     * Publisher-specific name of the deployed model
     */
    @Generated
    private String modelName;

    /*
     * Publisher-specific version of the deployed model
     */
    @Generated
    private String modelVersion;

    /*
     * Name of the deployed model's publisher
     */
    @Generated
    private String modelPublisher;

    /*
     * Capabilities of deployed model
     */
    @Generated
    private Map<String, String> capabilities;

    /*
     * Sku of the model deployment
     */
    @Generated
    private ModelDeploymentSku sku;

    /*
     * Name of the connection the deployment comes from
     */
    @Generated
    private String connectionName;

    /**
     * Creates an instance of ModelDeployment class.
     */
    @Generated
    private ModelDeployment() {
    }

    /**
     * Get the type property: The type of the deployment.
     *
     * @return the type value.
     */
    @Generated
    @Override
    public DeploymentType getType() {
        return this.type;
    }

    /**
     * Get the modelName property: Publisher-specific name of the deployed model.
     *
     * @return the modelName value.
     */
    @Generated
    public String getModelName() {
        return this.modelName;
  
... (truncated)
```

### ✅ 19. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/pom.xml","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 0eb6a2b1bec50b49ff0a2ad6e8c847e1ef4606f2)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// Code generated by Microsoft (R) TypeSpec Code Generator.
package com.azure.ai.projects.models;

import com.azure.core.annotation.Generated;
import com.azure.core.annotation.Immutable;
import com.azure.json.JsonReader;
import com.azure.json.JsonSerializable;
import com.azure.json.JsonToken;
import com.azure.json.JsonWriter;
import java.io.IOException;
import java.util.Map;

/**
 * Response from the list and get connections operations.
 */
@Immutable
public final class Connection implements JsonSerializable<Connection> {

    /*
     * The friendly name of the connection, provided by the user.
     */
    @Generated
    private String name;

    /*
     * A unique identifier for the connection, generated by the service
     */
    @Generated
    private String id;

    /*
     * Category of the connection
     */
    @Generated
    private ConnectionType type;

    /*
     * The connection URL to be used for this service
     */
    @Generated
    private String target;

    /*
     * Whether the connection is tagged as the default connection of its type
     */
    @Generated
    private boolean isDefault;

    /*
     * Metadata of the connection
     */
    @Generated
    private Map<String, String> metadata;

    /**
     * Creates an instance of Connection class.
     */
    @Generated
    private Connection() {
    }

    /**
     * Get the name property: The friendly name of the connection, provided by the user.
     *
     * @return the name value.
     */
    @Generated
    public String getName() {
        return this.name;
    }

    /**
     * Get the id property: A unique identifier for the connection, generated by the service.
     *
     * @return the id value.
     */
    @Generated
    public String getId() {
        return this.id;
    }

    /**
     * Get the type propert
... (truncated)
```

### ✅ 20. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003efoundry-project-inventory\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eFoundry Project Inventory\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure-ai-projects.version\u003e2.4.0\u003c/azure-ai-projects.version\u003e\n+        \u003cazure-identity.version\u003e1.18.4\u003c/azure-identity.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-ai-projects\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure-ai-projects.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure-identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.foundry.ProjectInventory\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/foundry/ProjectInventory.java\n+package com.example.foundry;\n+\n+import com.azure.ai.projects.AIProjectClientBuilder;\n+import com.azure.ai.projects.ConnectionsClient;\n+import com.azure.ai.projects.DeploymentsClient;\n+import com.azure.ai.projects.models.Connection;\n+import com.azure.ai.projects.models.Deployment;\n+import com.azure.ai.projects.models.ModelDeployment;\n+import com.azure.core.http.rest.PagedIterable;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+\n+public final class ProjectInventory {\n+    private ProjectInventory() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String endpoint = requireEnvironmentVariable(\"FOUNDRY_PROJECT_ENDPOINT\");\n+        String connectionName = requireEnvironmentVariable(\"CONNECTION_NAME\");\n+        String deploymentName = requireEnvironmentVariable(\"DEPLOYMENT_NAME\");\n+\n+        AIProjectClientBuilder clientBuilder = new AIProjectClientBuilder()\n+            .endpoint(endpoint)\n+            .credential(new DefaultAzureCredentialBuilder().build());\n+\n+        ConnectionsClient connectionsClient = clientBuilder.buildConnectionsClient();\n+        DeploymentsClient deploymentsClient = clientBuilder.buildDeploymentsClient();\n+\n+        System.out.println(\"Project connections\");\n+        PagedIterable\u003cConnection\u003e connections = connectionsClient.listConnections();\n+        for (Connection connection : connections) {\n+            printConnection(connection);\n+        }\n+\n+        System.out.printf(\"%nConnection '%s' (credentials excluded)%n\", connectionName);\n+        Connection selectedConnection = connectionsClient.getConnection(connectionName, false);\n+        printConnection(selectedConnection);\n+\n+        System.out.println(\"\\nModel deployments\");\n+        PagedIterable\u003cDeployment\u003e deployments = deploymentsClient.listDeployments();\n+        for (Deployment deployment : deployments) {\n+            if (deployment instanceof ModelDeployment modelDeployment) {\n+                printModelDeployment(modelDeployment);\n+            }\n+        }\n+\n+        System.out.printf(\"%nModel deployment '%s'%n\", deploymentName);\n+        Deployment selectedDeployment = deploymentsClient.getDeployment(deploymentName);\n+        if (!(selectedDeployment instanceof ModelDeployment modelDeployment)) {\n+            throw new IllegalStateException(\n+                \"Deployment '%s' is not a model deployment (type: %s).\"\n+                    .formatted(deploymentName, selectedDeployment.getType()));\n+        }\n+        printModelDeployment(modelDeployment);\n+    }\n+\n+    private static void printConnection(Connection connection) {\n+        System.out.printf(\n+            \"name=%s, type=%s, target=%s, default=%s%n\",\n+            connection.getName(),\n+            connection.getType(),\n+            connection.getTarget(),\n+            connection.isDefault());\n+    }\n+\n+    private static void printModelDeployment(ModelDeployment deployment) {\n+        System.out.printf(\n+            \"name=%s, publisher=%s, model=%s, version=%s%n\",\n+            deployment.getName(),\n+            deployment.getModelPublisher(),\n+            deployment.getModelName(),\n+            deployment.getModelVersion());\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Foundry Project Inventory\n+\n+Java 17 console application that uses the synchronous `azure-ai-projects` clients to inspect a Microsoft Foundry project's connections and model deployments. Authentication uses `DefaultAzureCredential`; sign in locally with a supported developer credential before running.\n+\n+## Configuration\n+\n+Set these environment variables in PowerShell:\n+\n+```powershell\n+$env:FOUNDRY_PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:CONNECTION_NAME = \"\u003cconnection-name\u003e\"\n+$env:DEPLOYMENT_NAME = \"\u003cdeployment-name\u003e\"\n+```\n+\n+## Restore, build, and run\n+\n+```powershell\n+mvn dependency:resolve\n+mvn package\n+mvn exec:java\n+```\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-java-project-resource-inventory-java-azure-tools-with-azure-tools-2919423923/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-java-project-resource-inventory-java-azure-tools-with-azure-tools-2919423923/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-java-project-resource-inventory-java-azure-tools-with-azure-tools-2919423923/pom.xml
@@ -1,0 +1,50 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>foundry-project-inventory</artifactId>
+    <version>1.0.0</version>
+    <name>Foundry Project Inventory</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure-ai-projects.version>2.4.0</azure-ai-projects.version>
+        <azure-identity.version>1.18.4</azure-identity.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-ai-projects</artifactId>
+            <version>${azure-ai-projects.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure-identity.version}</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-compiler-plugin</artifactId>
+                <version>3.14.1</version>
+            </plugin
... (truncated)
```

### ✅ 21. powershell

**Input:**

```json
{"command":"mvn --quiet package","description":"Build the Maven console application","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\foundry\ProjectInventory.java`

## Copilot Response

I’ll align the project with the current synchronous SDK surface, then create and locally verify the Maven application.

The current stable SDK documentation is newer than the skill’s beta example, so I’m resolving the exact 2.4.0 model hierarchy and method signatures before writing the project.

The verified API supports credential-free retrieval via `getConnection(name, false)` and polymorphic `ModelDeployment` instances. I’m creating a Maven project pinned to the stable 2.4.0 SDK with a runnable entry point and usage manifest.

Created a complete Maven Java 17 application in the workspace:

- `pom.xml` — uses stable `com.azure:azure-ai-projects:2.4.0`
- `src\main\java\com\example\foundry\ProjectInventory.java`
- `README.md` — environment setup plus restore, build, and run commands

The application uses synchronous pageable clients, credential-free connection retrieval, typed `ModelDeployment` handling, and `DefaultAzureCredential`. Build succeeds with `mvn package`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- The application consistently uses the synchronous Azure SDK clients requested by the prompt.

### Issues

- ConnectionsClient and DeploymentsClient are not closed, potentially leaking underlying HTTP resources.

## Grader Results

- project-resource-inventory.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Uses `AIProjectClientBuilder` to create `ConnectionsClient` and: Pass
      - Iterates every `Connection` from `ConnectionsClient.listConnections` and reads its: Pass
      - Calls `ConnectionsClient.getConnection` for `CONNECTION_NAME` with credentials: Pass
      - Iterates every `Deployment` from `DeploymentsClient.listDeployments`.: Pass
      - Narrows deployments to `ModelDeployment` before reading publisher, model name, and: Pass
      - Calls `DeploymentsClient.getDeployment` for `DEPLOYMENT_NAME` and rejects a result: Pass
      - Does not replace `PagedIterable` traversal with a single assumed page.: Pass
      - Does not select resources only from the locally enumerated results instead of: Pass
      - Does not request or display connection credentials.: Pass
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
  - Service-Specific Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Fail
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 13.00** | **Σ 10.0000** | **76.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id ai-projects-dp-java-project-resource-inventory --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)

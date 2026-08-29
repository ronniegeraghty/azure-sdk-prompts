# Evaluation Report: service-bus-dp-java-order-processor

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 435.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-order-processor` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 18/24 |
| Duration | 435.3s |
| Timestamp | 2026-08-29T03:14:33Z |
| Files Generated | 11 |
| Event Count | 17440 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 202.9s |
| Review | 231.2s |
| **Total** | **435.3s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=605791 out=12694 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. skill — 655ms

**Input:**

```json
{"skill":"azure-messaging"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Messaging SDK Troubleshooting

## Quick Reference

| Property | Value |
|----------|-------|
| **Services** | Azure Event Hubs, Azure Service Bus |
| **MCP Tools** | `mcp_azure_mcp_eventhubs`, `mcp_azure_mcp_servicebus` |
| **Best For** | Diagnosing SDK connection, auth, and message processing issues |

## When to Use This Skill

- SDK connection failures, auth errors, or AMQP link errors
- Idle timeout, connection inactivity, or slow reconnection after disconnect
- AMQP link detach or detach-forced errors
- Message lock lost, message lock expired, lock renewal failures, or batch lock timeouts
- Session lock lost, session lock expired, or session receiver errors
- Event processor or message handler stops processing
- Duplicate events or checkpoint offset resets
- SDK configuration questions (retry, prefetch, batch size, receive batch behavior)

## MCP Tools

| Tool | Command | Use |
|------|---------|-----|
| `mcp_azure_mcp_eventhubs` | Namespace/hub ops | List namespaces, hubs, consumer groups |
| `mcp_azure_mcp_servicebus` | Queue/topic ops | List namespaces, queues, topics, subscriptions |
| `mcp_azure_mcp_monitor` | `logs_query` | Query diagnostic logs with KQL |
| `mcp_azure_mcp_resourcehealth` | `get` | Check service health status |
| `mcp_azure_mcp_documentation` | Doc search | Search Microsoft Learn for troubleshooting docs |

## Diagnosis Workflow

1. **Identify the SDK and version** — Check the prompt for SDK and version clues; if not stated, proceed with diagnosis and ask later if needed
2. **Check resource health** — Use `mcp_azure_mcp_resourcehealth` to verify the namespace is healthy
3. **Review the error message** — Match against language-specific troubleshooting guide
4. **Look up documentation** — Use `mcp_azure_mcp_documentation` to search Microsoft Learn for the error or topic
5. **Check configuration** — Verify connection string, entity name, consumer group
6. **Recommend fix** — Apply remediatio
... (truncated)
```

### ❌ 2. view — 4073ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 4072ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. glob — 3720ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 4067ms

**Input:**

```json
{"intent":"Create a local-only Java 17 Maven example using Azure Service Bus queues, sessions, managed identity, scheduled messages, batching, dead-lettering, and sync/async clients. No Azure resources may be provisioned or modified.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 4399ms

**Input:**

```json
{"intent":"Find current official Java Azure Service Bus SDK guidance for sessions, sync and async send/receive, batch creation and tryAddMessage, scheduled delivery, dead-lettering, dead-letter subqueue receivers, and DefaultAzureCredential.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 3220ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java SDK code-generation best practices for a local Service Bus sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 4309ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Service Bus Java SDK docs.","parameters":{"query":"Azure Service Bus Java SDK ServiceBusClientBuilder sessions ServiceBusMessageBatch tryAddMessage schedule deadLetter SubQueue DEAD_LETTER_QUEUE DefaultAzureCredential sync async processor client"}}
```

**Output:**

```
{"results":[{"title":"ServiceBusClientBuilder Class-Clients and sub-builders","content":"# ServiceBusClientBuilder Class\r\n\r\n- Package:\r\n    - com.azure.messaging.servicebus\r\n\r\n- Maven Artifact:\r\n    - [com.azure:azure-messaging-servicebus:7.17.17](https://search.maven.org/artifact/com.azure/azure-messaging-servicebus/7.17.17/jar)\r\n\r\n- java.lang.Object\r\n- - com.azure.messaging.servicebus.ServiceBusClientBuilder\r\n\r\n**Implements**\r\n\r\n[AmqpTrait](https://learn.microsoft.com/java/api/com.azure.core.amqp.client.traits.amqptrait?view=azure-java-stable)\u0026lt;[ServiceBusClientBuilder](https://learn.microsoft.com/java/api/com.azure.messaging.servicebus.servicebusclientbuilder?view=azure-java-stable)\u0026gt;,[AzureNamedKeyCredentialTrait](https://learn.microsoft.com/java/api/com.azure.core.client.traits.azurenamedkeycredentialtrait?view=azure-java-stable)\u0026lt;[ServiceBusClientBuilder](https://learn.microsoft.com/java/api/com.azure.messaging.servicebus.servicebusclientbuilder?view=azure-java-stable)\u0026gt;,[AzureSasCredentialTrait](https://learn.microsoft.com/java/api/com.azure.core.client.traits.azuresascredentialtrait?view=azure-java-stable)\u0026lt;[ServiceBusClientBuilder](https://learn.microsoft.com/java/api/com.azure.messaging.servicebus.servicebusclientbuilder?view=azure-java-stable)\u0026gt;,[ConfigurationTrait](https://learn.microsoft.com/java/api/com.azure.core.client.traits.configurationtrait?view=azure-java-stable)\u0026lt;[ServiceBusClientBuilder](https://learn.microsoft.com/java/api/com.azure.messaging.servicebus.servicebusclientbuilder?view=azure-java-stable)\u0026gt;,[ConnectionStringTrait](https://learn.microsoft.com/java/api/com.azure.core.client.traits.connectionstringtrait?view=azure-java-stable)\u0026lt;[ServiceBusClientBuilder](https://learn.microsoft.com/java/api/com.azure.messaging.servicebus.servicebusclientbuilder?view=azure-java-stable)\u0026gt;,[TokenCredentialTrait](https://learn.microsoft.com/java/api/com.azure.c
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3939ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official Java Service Bus SDK code samples for sender, processor, sessions, batching, scheduling and dead-lettering.","parameters":{"language":"java","query":"Azure Service Bus Java ServiceBusClientBuilder sender processor session processor batch tryAddMessage scheduled enqueue deadLetter dead letter queue"}}
```

**Output:**

```
{"results":[{"description":"description: Creates a synchronous ServiceBusSenderClient to send messages in batches to an Azure Service Bus queue, handling batch size limits and closing the client after use.\npackage: com.azure:azure-messaging-servicebus, com.azure:azure-identity\nlanguage: java\n","codeSnippet":"TokenCredential credential = new DefaultAzureCredentialBuilder().build();\r\n\r\n// \u0027fullyQualifiedNamespace\u0027 will look similar to \u0022{your-namespace}.servicebus.windows.net\u0022\r\nServiceBusSenderClient sender = new ServiceBusClientBuilder()\r\n    .credential(fullyQualifiedNamespace, credential)\r\n    .sender()\r\n    .queueName(queueName)\r\n    .buildClient();\r\n\r\nList\u003CServiceBusMessage\u003E messages = Arrays.asList(\r\n    new ServiceBusMessage(\u0022test-1\u0022),\r\n    new ServiceBusMessage(\u0022test-2\u0022));\r\n\r\n// Creating a batch without options set.\r\nServiceBusMessageBatch batch = sender.createMessageBatch();\r\nfor (ServiceBusMessage message : messages) {\r\n    if (batch.tryAddMessage(message)) {\r\n        continue;\r\n    }\r\n\r\n    // The batch is full. Send the current batch and create a new one.\r\n    sender.sendMessages(batch);\r\n\r\n    batch = sender.createMessageBatch();\r\n\r\n    // Add the message we couldn\u0027t before.\r\n    if (!batch.tryAddMessage(message)) {\r\n        throw new IllegalArgumentException(\u0022Message is too large for an empty batch.\u0022);\r\n    }\r\n}\r\n\r\n// Send the final batch if there are any messages in it.\r\nif (batch.getCount() \u003E 0) {\r\n    sender.sendMessages(batch);\r\n}\r\n\r\n// Continue using the sender and finally, dispose of the sender.\r\n// Clients should be long-lived objects as they require resources\r\n// and time to establish a connection to the service.\r\nsender.close();","link":"https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable#examples","language":"java"},{"description":"description: De
... (truncated)
```

### ✅ 10. powershell — 3581ms

**Input:**

```json
{"command":"Get-ChildItem -Force 'C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging' | Select-Object Name,FullName,Mode","description":"Inspect Azure messaging skill references"}
```

**Output:**

```
Name     FullName
----     --------                                                                                                      
SKILL.md C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-m…
<shellId: 0 completed with exit code 0>
```

### ✅ 11. rg — 2555ms

**Input:**

```json
{"head_limit":120,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973343752-copilot-tool-output-753b612931084cf1bcfcb1089144c5d1.txt"],"pattern":"ServiceBusSessionProcessor|sessionProcessor|sessionId|setSessionId|SubQueue|deadLetter|scheduleMessage|tryAddMessage"}
```

**Output:**

```
Output too large to read at once (441.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787973351920-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt:{"results":[{"description":"description: Creates a synchronous ServiceBusSenderClient to send messages in batches to an Azure Service Bus queue, handling batch size limits and closing the client after use.\npackage: com.azure:azure-messaging-servicebus, com.azure:azure-identity\nlanguage: java\n","codeSnippet":"TokenCredential credential = new DefaultAzureCredentialBuilder().build();\r\n\r\n
```

### ✅ 12. web_fetch — 3249ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-messaging-servicebus/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-messaging-servicebus/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-messaging-servicebus</artifactId>
  <versioning>
    <latest>7.18.0-beta.3</latest>
    <release>7.18.0-beta.3</release>
    <versions>
      <version>7.0.0-beta.1</version>
      <version>7.0.0-beta.2</version>
      <version>7.0.0-beta.3</version>
      <version>7.0.0-beta.4</version>
      <version>7.0.0-beta.5</version>
      <version>7.0.0-beta.6</version>
      <version>7.0.0-beta.7</version>
      <version>7.0.0</version>
      <version>7.0.1</version>
      <version>7.0.2</version>
      <version>7.1.0</version>
      <version>7.2.0-beta.1</version>
      <version>7.2.0</version>
      <version>7.2.1</version>
      <version>7.2.2</version>
      <version>7.2.3</version>
      <version>7.3.0-beta.1</version>
      <version>7.3.0-beta.2</version>
      <version>7.3.0</version>
      <version>7.4.0</version>
      <version>7.4.1</version>
      <version>7.4.2</version>
      <version>7.5.0</version>
      <version>7.5.1</version>
      <version>7.5.2</version>
      <version>7.6.0</version>
      <version>7.7.0</version>
      <version>7.8.0</version>
      <version>7.9.0</version>
      <version>7.9.1</version>
      <version>7.10.0</version>
      <version>7.10.1</version>
      <version>7.11.0</version>
      <version>7.12.0</version>
      <version>7.12.1</version>
      <version>7.13.0</version>
      <version>7.13.1</version>
      <version>7.13.2</version>
      <version>7.13.3</version>
      <version>7.13.4</version>
      <version>7.14.0-beta.1</version>
      <version>7.14.0</version>
      <version>7.14.1</version>
      <version>7.14.2</version>
      <version>7.14.3</version>
      <version>7.14.4</version>
      <version>7.14.5</version>
      <version>7.14.6</version>
... (truncated)
```

### ✅ 13. web_fetch — 2894ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <versioning>
    <latest>1.19.0-beta.2</latest>
    <release>1.19.0-beta.2</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.0-preview.1</version>
      <version>1.0.0-preview.2</version>
      <version>1.0.0-preview.3</version>
      <version>1.0.0-preview.4</version>
      <version>1.0.0-preview.5</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.0.7</version>
      <version>1.0.8</version>
      <version>1.0.9</version>
      <version>1.1.0-beta.1</version>
      <version>1.1.0-beta.2</version>
      <version>1.1.0-beta.3</version>
      <version>1.1.0-beta.4</version>
      <version>1.1.0-beta.5</version>
      <version>1.1.0-beta.6</version>
      <version>1.1.0-beta.7</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.1.2</version>
      <version>1.1.3</version>
      <version>1.2.0-beta.1</version>
      <version>1.2.0-beta.2</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.3.0-beta.1</version>
      <version>1.3.0-beta.2</version>
      <version>1.3.0</version>
      <version>1.3.1</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version>1.3.4</version>
      <version>1.3.5</version>
      <version>1.3.6</version>
      <version>1.3.7</version>
      <version>1.4.0-beta.1</version>
      <version>1.4.0</version>
      <version>1.4.1</version>
      <version>
... (truncated)
```

### ✅ 14. view — 2212ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\SKILL.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging/SKILL.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging/SKILL.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging/SKILL.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging/SKILL.md
@@ -1,57 +1,57 @@
 ---
 name: azure-messaging
 description: "Troubleshoot and resolve issues with Azure Messaging SDKs for Event Hubs and Service Bus. Covers connection failures, authentication errors, message processing issues, and SDK configuration problems. WHEN: event hub SDK error, service bus SDK issue, messaging connection failure, AMQP error, event processor host issue, message lock lost, message lock expired, lock renewal, lock renewal batch, send timeout, receiver disconnected, SDK troubleshooting, azure messaging SDK, event hub consumer, service bus queue issue, topic subscription error, enable logging event hub, service bus logging, eventhub python, servicebus java, eventhub javascript, servicebus dotnet, event hub checkpoint, event hub not receiving messages, service bus dead letter, batch processing lock, session lock expired, idle timeout, connection inactive, link detach, slow reconnect, session error, duplicate events, offset reset, receive batch."
 license: MIT
 metadata:
   author: Microsoft
   version: "1.2.1"
 ---
 
 # Azure Messaging SDK Troubleshooting
 
 ## Quick Reference
 
 | Property | Value |
 |----------|-------|
 | **Services** | Azure Event Hubs, Azure Service Bus |
 | **MCP Tools** | `mcp_azure_mcp_eventhubs`, `mcp_azure_mcp_servicebus` |
 | **Best For** | Diagnosing SDK connection, auth, and message processing issues |
 
 ## When to Use This Skill
 
 - SDK connection failures, a
... (truncated)
```

### ✅ 15. azure-documentation (via azure) — 4076ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Verify Java session processor and session receiver APIs.","parameters":{"query":"site:learn.microsoft.com/java/api/com.azure.messaging.servicebus ServiceBusSessionProcessorClient ServiceBusSessionReceiverAsyncClient acceptNextSession maxConcurrentSessions maxConcurrentCalls buildProcessorClient"}}
```

**Output:**

```
{"results":[{"title":"ServiceBusReceiver interface-Method Details","content":"# ServiceBusReceiver interface\r\n\r\n- Package:\r\n    - [@azure/service-bus](https://learn.microsoft.com/en-us/javascript/api/@azure/service-bus/)\r\n\r\nA receiver that does not handle sessions.\r\n\r\n## Method Details\r\n\r\n### deadLetterMessage(ServiceBusReceivedMessage, DeadLetterOptions \u0026 {[key: string]: number | boolean | string | Date | null})\r\n\r\nMoves the message to the deadletter sub-queue. To receive a deadletted message, create a new QueueClient/SubscriptionClient using the path for the deadletter sub-queue.\r\n\r\n\u0060\u0060\u0060typescript\r\nfunction deadLetterMessage(message: ServiceBusReceivedMessage, options?: DeadLetterOptions \u0026 {[key: string]: number | boolean | string | Date | null}): Promise\u003Cvoid\u003E\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- message\r\n    - [ServiceBusReceivedMessage](https://learn.microsoft.com/en-us/javascript/api/@azure/service-bus/servicebusreceivedmessage)\r\n\r\n- options\r\n    - [DeadLetterOptions](https://learn.microsoft.com/en-us/javascript/api/@azure/service-bus/deadletteroptions) \u0026 {[key: string]: number | boolean | string | Date | null}\r\n\r\nThe DeadLetter options that can be provided while rejecting the message.\r\n\r\n#### Returns\r\n\r\nPromise\u0026lt;void\u0026gt;","contentUrl":"https://learn.microsoft.com/javascript/api/@azure/service-bus/servicebusreceiver?view=azure-node-latest"},{"title":"Azure Service Bus JMS 2.0 developer guide","content":"# Azure Service Bus JMS 2.0 developer guide\n## JMS - Building blocks\nWithout an exception listener, connection drops are completely silent. Add a \u0060jakarta.jms.ExceptionListener\u0060 to both sender and listener factories for observability:\n\u0060\u0060\u0060java\nconnection.setExceptionListener(exception -\u003E {\n    log.error(\u0022JMS connection error: {}\u0022, exception.getMessage(), exception);\n});\n\u0060\u0060\u0060\nIn Spring Bo
... (truncated)
```

### ✅ 16. azure-documentation (via azure) — 3720ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Verify Java dead-letter and settlement APIs.","parameters":{"query":"site:learn.microsoft.com/java/api/com.azure.messaging.servicebus ServiceBusReceivedMessageContext deadLetter DeadLetterOptions ServiceBusReceiverClient complete SubQueue DEAD_LETTER_QUEUE sessionReceiver"}}
```

**Output:**

```
{"results":[{"title":"Troubleshoot Azure Service Bus","content":"# Troubleshoot Azure Service Bus\n## Concurrency in ServiceBusProcessorClient\n\u0060ServiceBusProcessorClient\u0060 enables you to configure how many calls to the message handler happen concurrently. This configuration makes it possible to process multiple messages in parallel. For a \u0060ServiceBusProcessorClient\u0060 consuming messages from a non-session entity, you can configure the desired concurrency by using the \u0060maxConcurrentCalls\u0060 API. For a session-enabled entity, the desired concurrency is \u0060maxConcurrentSessions\u0060 times \u0060maxConcurrentCalls\u0060.\nIf you observe fewer concurrent calls to the message handler than the configured concurrency, it might be because the thread pool isn\u0027t sized appropriately.\n\u0060ServiceBusProcessorClient\u0060 uses daemon threads from the Reactor global [boundedElastic](https://projectreactor.io/docs/core/release/api/reactor/core/scheduler/Schedulers.html#boundedElastic--) thread pool to invoke the message handler. The maximum number of concurrent threads in this pool is limited by a cap. By default, this cap is ten times the number of available CPU cores. For the \u0060ServiceBusProcessorClient\u0060 to effectively support the application\u0027s desired concurrency (\u0060maxConcurrentCalls\u0060 or \u0060maxConcurrentSessions\u0060 times \u0060maxConcurrentCalls\u0060), you must have a \u0060boundedElastic\u0060 pool cap value that\u0027s higher than the desired concurrency. You can override the default cap by setting the system property \u0060reactor.schedulers.defaultBoundedElasticSize\u0060.\nTune the thread pool and CPU allocation on a case-by-case basis. However, when you override the pool cap, as a starting point, limit the concurrent threads to approximately 20-30 per CPU core. Cap the desired concurrency per \u0060ServiceBusProcessorClient\u0060 instance to approximately 20-30. Profile and measure your specific use case a
... (truncated)
```

### ✅ 17. web_fetch — 3378ms

**Input:**

```json
{"max_length":3000,"start_index":7000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-databind/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.fasterxml.jackson.core</groupId>
  <artifactId>jackson-databind</artifactId>
  <versioning>
    <latest>2.22.2</latest>
    <release>2.22.2</release>
    <versions>
      <version>2.0.0-RC1</version>
      <version>2.0.0-RC2</version>
      <version>2.0.0-RC3</version>
      <version>2.0.0</version>
      <version>2.0.1</version>
      <version>2.0.2</version>
      <version>2.0.4</version>
      <version>2.0.5</version>
      <version>2.0.6</version>
      <version>2.1.0</version>
      <version>2.1.1</version>
      <version>2.1.2</version>
      <version>2.1.3</version>
      <version>2.1.4</version>
      <version>2.1.5</version>
      <version>2.2.0-rc1</version>
      <version>2.2.0</version>
      <version>2.2.1</version>
      <version>2.2.2</version>
      <version>2.2.3</version>
      <version>2.2.4</version>
      <version>2.3.0-rc1</version>
      <version>2.3.0</version>
      <version>2.3.1</version>
      <version>2.3.2</version>
      <version>2.3.3</version>
      <version>2.3.4</version>
      <version>2.3.5</version>
      <version>2.4.0-rc1</version>
      <version>2.4.0-rc2</version>
      <version>2.4.0-rc3</version>
      <version>2.4.0</version>
      <version>2.4.1.1</version>
      <version>2.4.1.2</version>
      <version>2.4.1.3</version>
      <version>2.4.1</version>
      <version>2.4.2</version>
      <version>2.4.3</version>
      <version>2.4.4</version>
      <version>2.4.5.1<

<note>Content truncated. Call the fetch tool with a start_index of 1500 to get more content.</note>
```

### ✅ 18. web_fetch — 3032ms

**Input:**

```json
{"max_length":1500,"url":"https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-databind/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<error>No more content available.</error>
```

### ✅ 19. powershell — 1413ms

**Input:**

```json
{"command":"$files = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973371105-copilot-tool-output-6bfabeabdc714b1da7bbed8e714ed6e3.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973371161-copilot-tool-output-c55bcff04fc14f2d9055c1b577efc862.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt'); foreach ($file in $files) { $json = Get-Content -Raw $file | ConvertFrom-Json; $json.results | Where-Object { $_.title -match 'Session|Processor|Receiver|Dead|Sender|Batch' -or $_.description -match 'session|dead|batch' } | Select-Object -First 12 | ForEach-Object { \"TITLE: $($_.title)\"; if ($_.codeSnippet) { $_.codeSnippet.Substring(0,[Math]::Min(2500,$_.codeSnippet.Length)) } else { $_.content.Substring(0,[Math]::Min(2500,$_.content.Length)) }; '---' } }","description":"Extract relevant Service Bus API samples"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … n $files) { $json = Get-Content -Raw $file | ConvertFrom-Json; $json. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 20. view — 1543ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973371105-copilot-tool-output-6bfabeabdc714b1da7bbed8e714ed6e3.txt","view_range":[1,5]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt b/C:/Users/weidxu/AppData/Local/Temp/1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt
@@ -1,5 +1,5 @@
 {"results":[{"description":"description: Creates a synchronous ServiceBusSenderClient to send messages in batches to an Azure Service Bus queue, handling batch size limits and closing the client after use.\npackage: com.azure:azure-messaging-servicebus, com.azure:azure-identity\nlanguage: java\n","codeSnippet":"TokenCredential credential = new DefaultAzureCredentialBuilder().build();\r\n\r\n// \u0027fullyQualifiedNamespace\u0027 will look similar to \u0022{your-namespace}.servicebus.windows.net\u0022\r\nServiceBusSenderClient sender = new ServiceBusClientBuilder()\r\n    .credential(fullyQualifiedNamespace, credential)\r\n    .sender()\r\n    .queueName(queueName)\r\n    .buildClient();\r\n\r\nList\u003CServiceBusMessage\u003E messages = Arrays.asList(\r\n    new ServiceBusMessage(\u0022test-1\u0022),\r\n    new ServiceBusMessage(\u0022test-2\u0022));\r\n\r\n// Creating a batch without options set.\r\nServiceBusMessageBatch batch = sender.createMessageBatch();\r\nfor (ServiceBusMessage message : messages) {\r\n    if (batch.tryAddMessage(message)) {\r\n        continue;\r\n    }\r\n\r\n    // The batch is full. Send the current batch and create a new one.\r\n    sender.sendMessages(batch);\r\n\r\n    batch = sender.createMessageBatch();\r\n\r\n    // Add the message we couldn\u0027t before.\r\n    if (!batch.tryAddMessage(message)) {\r\n        throw new IllegalArgumentException(\u0022Message is too large for an empty batch.\u0022);\r\n    }\r\n}\r\n\r\n// Send the final batch
... (truncated)
```

### ✅ 21. view — 1160ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787973343458-copilot-tool-output-0724d678bc4f4c0c82805ce28fa958fa.txt","view_range":[1,5]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787973371105-copilot-tool-output-6bfabeabdc714b1da7bbed8e714ed6e3.txt b/C:/Users/weidxu/AppData/Local/Temp/1787973371105-copilot-tool-output-6bfabeabdc714b1da7bbed8e714ed6e3.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787973371105-copilot-tool-output-6bfabeabdc714b1da7bbed8e714ed6e3.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787973371105-copilot-tool-output-6bfabeabdc714b1da7bbed8e714ed6e3.txt
@@ -1,5 +1,5 @@
 {"results":[{"title":"Troubleshoot Azure Service Bus","content":"# Troubleshoot Azure Service Bus\n## Concurrency in ServiceBusProcessorClient\n\u0060ServiceBusProcessorClient\u0060 enables you to configure how many calls to the message handler happen concurrently. This configuration makes it possible to process multiple messages in parallel. For a \u0060ServiceBusProcessorClient\u0060 consuming messages from a non-session entity, you can configure the desired concurrency by using the \u0060maxConcurrentCalls\u0060 API. For a session-enabled entity, the desired concurrency is \u0060maxConcurrentSessions\u0060 times \u0060maxConcurrentCalls\u0060.\nIf you observe fewer concurrent calls to the message handler than the configured concurrency, it might be because the thread pool isn\u0027t sized appropriately.\n\u0060ServiceBusProcessorClient\u0060 uses daemon threads from the Reactor global [boundedElastic](https://projectreactor.io/docs/core/release/api/reactor/core/scheduler/Schedulers.html#boundedElastic--) thread pool to invoke the message handler. The maximum number of concurrent threads in this pool is limited by a cap. By default, this cap is ten times the number of available CPU cores. For the \u0060ServiceBusProcessorClient\u0060 to effectively support the application\u0027s desired concurrency (\u0060maxConcurrentCalls\u0060 or \u0060maxConcurrentSessions\u0060 times \u0060maxConcurrentCalls\u0060), you must have a \u0060boundedElastic\u0060 pool cap value that\u0027s h
... (truncated)
```

### ✅ 22. apply_patch — 755ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eservice-bus-order-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eService Bus Order Processor\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.servicebus.version\u003e7.17.20\u003c/azure.servicebus.version\u003e\n+        \u003cazure.identity.version\u003e1.18.1\u003c/azure.identity.version\u003e\n+        \u003cjackson.version\u003e2.22.2\u003c/jackson.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.servicebus.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/orders/Order.java\n+package com.example.orders;\n+\n+import com.fasterxml.jackson.annotation.JsonCreator;\n+import com.fasterxml.jackson.annotation.JsonProperty;\n+\n+import java.math.BigDecimal;\n+import java.util.Objects;\n+\n+public final class Order {\n+    public enum Status {\n+        PENDING,\n+        PROCESSING,\n+        COMPLETED,\n+        FAILED\n+    }\n+\n+    private final String orderId;\n+    private final String customerName;\n+    private final String product;\n+    private final int quantity;\n+    private final BigDecimal totalPrice;\n+    private Status status;\n+\n+    @JsonCreator\n+    public Order(\n+            @JsonProperty(\"orderId\") String orderId,\n+            @JsonProperty(\"customerName\") String customerName,\n+            @JsonProperty(\"product\") String product,\n+            @JsonProperty(\"quantity\") int quantity,\n+            @JsonProperty(\"totalPrice\") BigDecimal totalPrice,\n+            @JsonProperty(\"status\") Status status) {\n+        this.orderId = Objects.requireNonNull(orderId, \"orderId\");\n+        this.customerName = Objects.requireNonNull(customerName, \"customerName\");\n+        this.product = Objects.requireNonNull(product, \"product\");\n+        if (quantity \u003c= 0) {\n+            throw new IllegalArgumentException(\"quantity must be positive\");\n+        }\n+        this.quantity = quantity;\n+        this.totalPrice = Objects.requireNonNull(totalPrice, \"totalPrice\");\n+        this.status = Objects.requireNonNull(status, \"status\");\n+    }\n+\n+    public String getOrderId() {\n+        return orderId;\n+    }\n+\n+    public String getCustomerName() {\n+        return customerName;\n+    }\n+\n+    public String getProduct() {\n+        return product;\n+    }\n+\n+    public int getQuantity() {\n+        return quantity;\n+    }\n+\n+    public BigDecimal getTotalPrice() {\n+        return totalPrice;\n+    }\n+\n+    public Status getStatus() {\n+        return status;\n+    }\n+\n+    public void setStatus(Status status) {\n+        this.status = Objects.requireNonNull(status, \"status\");\n+    }\n+\n+    @Override\n+    public String toString() {\n+        return \"Order{orderId='%s', customerName='%s', product='%s', quantity=%d, totalPrice=%s, status=%s}\"\n+                .formatted(orderId, customerName, product, quantity, totalPrice, status);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderMessageFactory.java\n+package com.example.orders;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+\n+final class OrderMessageFactory {\n+    static final String CONTENT_TYPE = \"application/json\";\n+    static final String PRIORITY_PROPERTY = \"priority\";\n+    static final String HIGH_PRIORITY = \"high\";\n+    static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);\n+\n+    private final ObjectMapper objectMapper;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    OrderMessageFactory(ObjectMapper objectMapper, BigDecimal highPriorityThreshold) {\n+        this.objectMapper = objectMapper;\n+        this.highPriorityThreshold = highPriorityThreshold;\n+    }\n+\n+    ServiceBusMessage create(Order order) {\n+        try {\n+            ServiceBusMessage message = new ServiceBusMessage(\n+                    BinaryData.fromString(objectMapper.writeValueAsString(order)))\n+                    .setContentType(CONTENT_TYPE)\n+                    .setMessageId(order.getOrderId())\n+                    .setCorrelationId(order.getOrderId())\n+                    .setSessionId(order.getCustomerName());\n+\n+            if (order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0) {\n+                message.getApplicationProperties().put(PRIORITY_PROPERTY, HIGH_PRIORITY);\n+                message.setScheduledEnqueueTime(OffsetDateTime.now().plus(FRAUD_REVIEW_DELAY));\n+            }\n+            return message;\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Could not serialize order \" + order.getOrderId(), exception);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.stream.Collectors;\n+\n+public final class SyncOrderSender implements AutoCloseable {\n+    private final ServiceBusSenderClient sender;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public SyncOrderSender(\n+            ServiceBusSenderClient sender,\n+            ObjectMapper objectMapper,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = Objects.requireNonNull(sender, \"sender\");\n+        this.messageFactory = new OrderMessageFactory(objectMapper, highPriorityThreshold);\n+    }\n+\n+    public void send(Order order) {\n+        sender.sendMessage(messageFactory.create(order));\n+    }\n+\n+    public void sendBatch(List\u003cOrder\u003e orders) {\n+        Map\u003cString, List\u003cOrder\u003e\u003e ordersByCustomer = orders.stream()\n+                .collect(Collectors.groupingBy(\n+                        Order::getCustomerName,\n+                        java.util.LinkedHashMap::new,\n+                        Collectors.toList()));\n+\n+        for (List\u003cOrder\u003e customerOrders : ordersByCustomer.values()) {\n+            sendCustomerBatches(customerOrders);\n+        }\n+    }\n+\n+    private void sendCustomerBatches(List\u003cOrder\u003e orders) {\n+        ServiceBusMessageBatch batch = sender.createMessageBatch();\n+        for (Order order : orders) {\n+            ServiceBusMessage message = messageFactory.create(order);\n+            if (batch.tryAddMessage(message)) {\n+                continue;\n+            }\n+            sender.sendMessages(batch);\n+            batch = sender.createMessageBatch();\n+            if (!batch.tryAddMessage(message)) {\n+                throw new IllegalArgumentException(\n+                        \"Order %s exceeds the maximum Service Bus message size\".formatted(order.getOrderId()));\n+            }\n+        }\n+        if (batch.getCount() \u003e 0) {\n+            sender.sendMessages(batch);\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.stream.Collectors;\n+\n+public final class AsyncOrderSender implements AutoCloseable {\n+    private final ServiceBusSenderAsyncClient sender;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public AsyncOrderSender(\n+            ServiceBusSenderAsyncClient sender,\n+            ObjectMapper objectMapper,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = Objects.requireNonNull(sender, \"sender\");\n+        this.messageFactory = new OrderMessageFactory(objectMapper, highPriorityThreshold);\n+    }\n+\n+    public Mono\u003cVoid\u003e send(Order order) {\n+        return sender.sendMessage(messageFactory.create(order));\n+    }\n+\n+    public Mono\u003cVoid\u003e sendBatch(List\u003cOrder\u003e orders) {\n+        Map\u003cString, List\u003cOrder\u003e\u003e ordersByCustomer = orders.stream()\n+                .collect(Collectors.groupingBy(\n+                        Order::getCustomerName,\n+                        java.util.LinkedHashMap::new,\n+                        Collectors.toList()));\n+\n+        return Flux.fromIterable(ordersByCustomer.values())\n+                .concatMap(this::sendCustomerBatches)\n+                .then();\n+    }\n+\n+    private Mono\u003cVoid\u003e sendCustomerBatches(List\u003cOrder\u003e orders) {\n+        return sender.createMessageBatch().flatMap(batch -\u003e fillAndSend(batch, orders, 0));\n+    }\n+\n+    private Mono\u003cVoid\u003e fillAndSend(ServiceBusMessageBatch batch, List\u003cOrder\u003e orders, int index) {\n+        int nextIndex = index;\n+        while (nextIndex \u003c orders.size()) {\n+            Order order = orders.get(nextIndex);\n+            ServiceBusMessage message = messageFactory.create(order);\n+            if (!batch.tryAddMessage(message)) {\n+                if (batch.getCount() == 0) {\n+                    return Mono.error(new IllegalArgumentException(\n+                            \"Order %s exceeds the maximum Service Bus message size\".formatted(order.getOrderId())));\n+                }\n+                int remainingIndex = nextIndex;\n+                return sender.sendMessages(batch)\n+                        .then(sender.createMessageBatch())\n+                        .flatMap(nextBatch -\u003e fillAndSend(nextBatch, orders, remainingIndex));\n+            }\n+            nextIndex++;\n+        }\n+        return batch.getCount() == 0 ? Mono.empty() : sender.sendMessages(batch);\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.time.Duration;\n+import java.util.Objects;\n+\n+public final class SyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);\n+\n+    private final ServiceBusProcessorClient processor;\n+    private final ServiceBusReceiverClient deadLetterReceiver;\n+    private final ServiceBusSenderClient reprocessSender;\n+    private final ObjectMapper objectMapper;\n+\n+    public SyncOrderProcessor(\n+            ServiceBusProcessorClient processor,\n+            ServiceBusReceiverClient deadLetterReceiver,\n+            ServiceBusSenderClient reprocessSender,\n+            ObjectMapper objectMapper) {\n+        this.processor = Objects.requireNonNull(processor, \"processor\");\n+        this.deadLetterReceiver = Objects.requireNonNull(deadLetterReceiver, \"deadLetterReceiver\");\n+        this.reprocessSender = Objects.requireNonNull(reprocessSender, \"reprocessSender\");\n+        this.objectMapper = Objects.requireNonNull(objectMapper, \"objectMapper\");\n+    }\n+\n+    public void runFor(Duration duration) throws InterruptedException {\n+        processor.start();\n+        try {\n+            Thread.sleep(duration.toMillis());\n+        } finally {\n+            processor.stop();\n+        }\n+    }\n+\n+    public void inspectAndReprocessDeadLetters(int maxMessages, Duration waitTime) {\n+        for (ServiceBusReceivedMessage message : deadLetterReceiver.receiveMessages(maxMessages, waitTime)) {\n+            LOGGER.warn(\n+                    \"DLQ message id={}, reason={}, description={}, body={}\",\n+                    message.getMessageId(),\n+                    message.getDeadLetterReason(),\n+                    message.getDeadLetterErrorDescription(),\n+                    message.getBody());\n+            try {\n+                Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);\n+                ServiceBusMessageReprocessor.send(reprocessSender, order, objectMapper);\n+                deadLetterReceiver.complete(message);\n+                LOGGER.info(\"Reprocessed dead-lettered order {}\", order.getOrderId());\n+            } catch (Exception exception) {\n+                LOGGER.warn(\"DLQ message {} is not a valid order and was left in the DLQ\", message.getMessageId());\n+            }\n+        }\n+    }\n+\n+    public static void processMessage(\n+            com.azure.messaging.servicebus.ServiceBusReceivedMessageContext context,\n+            ObjectMapper objectMapper) {\n+        ServiceBusReceivedMessage message = context.getMessage();\n+        try {\n+            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);\n+            order.setStatus(Order.Status.PROCESSING);\n+            LOGGER.info(\"Processing {}\", order);\n+            order.setStatus(Order.Status.COMPLETED);\n+            LOGGER.info(\"Completed {}\", order);\n+            context.complete();\n+        } catch (Exception exception) {\n+            String reason = \"ORDER_DESERIALIZATION_OR_PROCESSING_FAILED\";\n+            LOGGER.error(\"Dead-lettering message {}: {}\", message.getMessageId(), exception.getMessage());\n+            context.deadLetter(new DeadLetterOptions()\n+                    .setDeadLetterReason(reason)\n+                    .setDeadLetterErrorDescription(truncate(exception.getMessage(), 4096)));\n+        }\n+    }\n+\n+    private static String truncate(String value, int maxLength) {\n+        String safeValue = value == null ? \"Unknown processing error\" : value;\n+        return safeValue.length() \u003c= maxLength ? safeValue : safeValue.substring(0, maxLength);\n+    }\n+\n+    @Override\n+    public void close() {\n+        processor.close();\n+        deadLetterReceiver.close();\n+        reprocessSender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/ServiceBusMessageReprocessor.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+\n+final class ServiceBusMessageReprocessor {\n+    private static final BigDecimal NO_DELAY_THRESHOLD = new BigDecimal(\"999999999999\");\n+\n+    private ServiceBusMessageReprocessor() {\n+    }\n+\n+    static void send(ServiceBusSenderClient sender, Order order, ObjectMapper objectMapper) {\n+        sender.sendMessage(new OrderMessageFactory(objectMapper, NO_DELAY_THRESHOLD).create(order));\n+    }\n+\n+    static Mono\u003cVoid\u003e send(ServiceBusSenderAsyncClient sender, Order order, ObjectMapper objectMapper) {\n+        return sender.sendMessage(new OrderMessageFactory(objectMapper, NO_DELAY_THRESHOLD).create(order));\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.Objects;\n+import java.util.concurrent.TimeoutException;\n+\n+public final class AsyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);\n+\n+    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;\n+    private final ServiceBusReceiverAsyncClient deadLetterReceiver;\n+    private final ServiceBusSenderAsyncClient reprocessSender;\n+    private final ObjectMapper objectMapper;\n+\n+    public AsyncOrderProcessor(\n+            ServiceBusSessionReceiverAsyncClient sessionReceiver,\n+            ServiceBusReceiverAsyncClient deadLetterReceiver,\n+            ServiceBusSenderAsyncClient reprocessSender,\n+            ObjectMapper objectMapper) {\n+        this.sessionReceiver = Objects.requireNonNull(sessionReceiver, \"sessionReceiver\");\n+        this.deadLetterReceiver = Objects.requireNonNull(deadLetterReceiver, \"deadLetterReceiver\");\n+        this.reprocessSender = Objects.requireNonNull(reprocessSender, \"reprocessSender\");\n+        this.objectMapper = Objects.requireNonNull(objectMapper, \"objectMapper\");\n+    }\n+\n+    public Mono\u003cVoid\u003e processFor(Duration duration) {\n+        long deadline = System.nanoTime() + duration.toNanos();\n+        return processNextSession(deadline);\n+    }\n+\n+    private Mono\u003cVoid\u003e processNextSession(long deadline) {\n+        Duration remaining = Duration.ofNanos(Math.max(0, deadline - System.nanoTime()));\n+        if (remaining.isZero()) {\n+            return Mono.empty();\n+        }\n+\n+        return sessionReceiver.acceptNextSession()\n+                .timeout(remaining)\n+                .flatMap(receiver -\u003e processSession(receiver, remaining)\n+                        .doFinally(signal -\u003e receiver.close()))\n+                .then(Mono.defer(() -\u003e processNextSession(deadline)))\n+                .onErrorResume(TimeoutException.class, exception -\u003e Mono.empty());\n+    }\n+\n+    private Mono\u003cVoid\u003e processSession(ServiceBusReceiverAsyncClient receiver, Duration remaining) {\n+        Duration idleTimeout = remaining.compareTo(Duration.ofSeconds(5)) \u003c 0\n+                ? remaining\n+                : Duration.ofSeconds(5);\n+        return receiver.receiveMessages()\n+                .concatMap(message -\u003e processMessage(receiver, message))\n+                .timeout(idleTimeout)\n+                .onErrorResume(TimeoutException.class, exception -\u003e Flux.empty())\n+                .then();\n+    }\n+\n+    private Mono\u003cVoid\u003e processMessage(ServiceBusReceiverAsyncClient receiver, ServiceBusReceivedMessage message) {\n+        try {\n+            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);\n+            order.setStatus(Order.Status.PROCESSING);\n+            LOGGER.info(\"Processing {}\", order);\n+            order.setStatus(Order.Status.COMPLETED);\n+            LOGGER.info(\"Completed {}\", order);\n+            return receiver.complete(message);\n+        } catch (Exception exception) {\n+            LOGGER.error(\"Dead-lettering message {}: {}\", message.getMessageId(), exception.getMessage());\n+            return receiver.deadLetter(message, new DeadLetterOptions()\n+                    .setDeadLetterReason(\"ORDER_DESERIALIZATION_OR_PROCESSING_FAILED\")\n+                    .setDeadLetterErrorDescription(truncate(exception.getMessage(), 4096)));\n+        }\n+    }\n+\n+    public Mono\u003cVoid\u003e inspectAndReprocessDeadLetters(int maxMessages, Duration waitTime) {\n+        return deadLetterReceiver.receiveMessages()\n+                .take(maxMessages)\n+                .concatMap(message -\u003e inspectAndReprocess(message)\n+                        .onErrorResume(exception -\u003e {\n+                            LOGGER.warn(\n+                                    \"DLQ message {} could not be reprocessed and was left in the DLQ: {}\",\n+                                    message.getMessageId(),\n+                                    exception.getMessage());\n+                            return Mono.empty();\n+                        }))\n+                .timeout(waitTime)\n+                .onErrorResume(TimeoutException.class, exception -\u003e Flux.empty())\n+                .then();\n+    }\n+\n+    private Mono\u003cVoid\u003e inspectAndReprocess(ServiceBusReceivedMessage message) {\n+        LOGGER.warn(\n+                \"DLQ message id={}, reason={}, description={}, body={}\",\n+                message.getMessageId(),\n+                message.getDeadLetterReason(),\n+                message.getDeadLetterErrorDescription(),\n+                message.getBody());\n+        try {\n+            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);\n+            return ServiceBusMessageReprocessor.send(reprocessSender, order, objectMapper)\n+                    .then(deadLetterReceiver.complete(message))\n+                    .doOnSuccess(ignored -\u003e LOGGER.info(\"Reprocessed dead-lettered order {}\", order.getOrderId()));\n+        } catch (Exception exception) {\n+            return Mono.error(exception);\n+        }\n+    }\n+\n+    private static String truncate(String value, int maxLength) {\n+        String safeValue = value == null ? \"Unknown processing error\" : value;\n+        return safeValue.length() \u003c= maxLength ? safeValue : safeValue.substring(0, maxLength);\n+    }\n+\n+    @Override\n+    public void close() {\n+        sessionReceiver.close();\n+        deadLetterReceiver.close();\n+        reprocessSender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Main.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.BinaryData;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import com.fasterxml.jackson.databind.json.JsonMapper;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.UUID;\n+\n+public final class Main {\n+    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal(\"1000.00\");\n+    private static final Duration PROCESSING_WINDOW = Duration.ofSeconds(40);\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String namespace = requiredEnvironmentVariable(\"AZURE_SERVICE_BUS_NAMESPACE\");\n+        String queueName = System.getenv().getOrDefault(\"AZURE_SERVICE_BUS_QUEUE\", \"orders\");\n+        ObjectMapper objectMapper = JsonMapper.builder().findAndAddModules().build();\n+        TokenCredential credential = new DefaultAzureCredentialBuilder().build();\n+\n+        runSyncDemo(namespace, queueName, credential, objectMapper);\n+        runAsyncDemo(namespace, queueName, credential, objectMapper);\n+    }\n+\n+    private static void runSyncDemo(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            ObjectMapper objectMapper) throws InterruptedException {\n+        ServiceBusSenderClient senderClient = builder(namespace, credential).sender()\n+                .queueName(queueName)\n+                .buildClient();\n+        ServiceBusProcessorClient processorClient = builder(namespace, credential).sessionProcessor()\n+                .queueName(queueName)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .maxConcurrentSessions(1)\n+                .maxConcurrentCalls(1)\n+                .processMessage(context -\u003e SyncOrderProcessor.processMessage(context, objectMapper))\n+                .processError(context -\u003e System.err.println(\"Sync processor error: \" + context.getException()))\n+                .buildProcessorClient();\n+        ServiceBusReceiverClient deadLetterReceiver = builder(namespace, credential).receiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildClient();\n+        ServiceBusSenderClient reprocessSender = builder(namespace, credential).sender()\n+                .queueName(queueName)\n+                .buildClient();\n+\n+        try (SyncOrderSender sender =\n+                     new SyncOrderSender(senderClient, objectMapper, HIGH_PRIORITY_THRESHOLD);\n+             SyncOrderProcessor processor =\n+                     new SyncOrderProcessor(processorClient, deadLetterReceiver, reprocessSender, objectMapper)) {\n+            sender.sendBatch(sampleOrders(\"sync\"));\n+            sendMalformed(senderClient, \"sync-bad\");\n+            processor.runFor(PROCESSING_WINDOW);\n+            processor.inspectAndReprocessDeadLetters(10, Duration.ofSeconds(5));\n+        }\n+    }\n+\n+    private static void runAsyncDemo(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            ObjectMapper objectMapper) {\n+        ServiceBusSenderAsyncClient senderClient = builder(namespace, credential).sender()\n+                .queueName(queueName)\n+                .buildAsyncClient();\n+        ServiceBusSessionReceiverAsyncClient sessionReceiver = builder(namespace, credential).sessionReceiver()\n+                .queueName(queueName)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+        ServiceBusReceiverAsyncClient deadLetterReceiver = builder(namespace, credential).receiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+        ServiceBusSenderAsyncClient reprocessSender = builder(namespace, credential).sender()\n+                .queueName(queueName)\n+                .buildAsyncClient();\n+\n+        try (AsyncOrderSender sender =\n+                     new AsyncOrderSender(senderClient, objectMapper, HIGH_PRIORITY_THRESHOLD);\n+             AsyncOrderProcessor processor =\n+                     new AsyncOrderProcessor(sessionReceiver, deadLetterReceiver, reprocessSender, objectMapper)) {\n+            sender.sendBatch(sampleOrders(\"async\"))\n+                    .then(senderClient.sendMessage(malformedMessage(\"async-bad\")))\n+                    .then(processor.processFor(PROCESSING_WINDOW))\n+                    .then(processor.inspectAndReprocessDeadLetters(10, Duration.ofSeconds(5)))\n+                    .block();\n+        }\n+    }\n+\n+    private static ServiceBusClientBuilder builder(String namespace, TokenCredential credential) {\n+        return new ServiceBusClientBuilder().credential(namespace, credential);\n+    }\n+\n+    private static List\u003cOrder\u003e sampleOrders(String prefix) {\n+        return List.of(\n+                new Order(prefix + \"-001\", \"Ada\", \"Keyboard\", 1, new BigDecimal(\"149.99\"), Order.Status.PENDING),\n+                new Order(prefix + \"-002\", \"Ada\", \"Monitor\", 2, new BigDecimal(\"799.98\"), Order.Status.PENDING),\n+                new Order(prefix + \"-003\", \"Grace\", \"Server\", 1, new BigDecimal(\"4200.00\"), Order.Status.PENDING));\n+    }\n+\n+    private static void sendMalformed(ServiceBusSenderClient sender, String id) {\n+        sender.sendMessage(malformedMessage(id));\n+    }\n+\n+    private static ServiceBusMessage malformedMessage(String id) {\n+        return new ServiceBusMessage(BinaryData.fromString(\"{not-valid-json\"))\n+                .setMessageId(id + \"-\" + UUID.randomUUID())\n+                .setCorrelationId(id)\n+                .setSessionId(\"Malformed demo\");\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(name + \" must be set, for example contoso.servicebus.windows.net\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Azure Service Bus order processor\n+\n+A Java 17 Maven sample with synchronous and reactive senders/processors, JSON orders, size-aware batching,\n+scheduled high-value orders, explicit dead-lettering, and DLQ inspection/reprocessing.\n+\n+## Prerequisites\n+\n+- Java 17 and Maven 3.9+\n+- An existing Azure Service Bus queue with **sessions enabled**\n+- A managed identity (in Azure) or developer credential (locally) with the Azure Service Bus Data Sender and\n+  Azure Service Bus Data Receiver roles\n+\n+No resources are provisioned by this project.\n+\n+## Run\n+\n+Set the fully qualified namespace, such as `contoso.servicebus.windows.net`. The queue defaults to `orders`.\n+\n+```powershell\n+$env:AZURE_SERVICE_BUS_NAMESPACE = \"contoso.servicebus.windows.net\"\n+$env:AZURE_SERVICE_BUS_QUEUE = \"orders\"\n+mvn compile exec:java\n+```\n+\n+`DefaultAzureCredential` uses managed identity when hosted in Azure and supported developer credentials locally.\n+Orders are assigned a session ID equal to the customer name. Batches are split by customer and by the\n+service-reported maximum batch size. Orders over `$1,000` get a `priority=high` application property and are\n+scheduled 30 seconds into the future for fraud review.\n+\n+The processors use peek-lock mode and explicit settlement. Processing failures are dead-lettered with a reason and\n+description. DLQ readers log every inspected message and re-enqueue messages that still contain a valid `Order`.\n+\n+## References\n+\n+- [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme)\n+- [Service Bus message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions)\n+- [Passwordless authentication for Service Bus](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-java-how-to-use-queues)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml
@@ -1,0 +1,63 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>service-bus-order-processor</artifactId>
+    <version>1.0.0</version>
+    <name>Service Bus Order Processor</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.servicebus.version>7.17.20</azure.servicebus.version>
+        <azure.identity.version>1.18.1</azure.identity.version>
+        <jackson.version>2.22.2</jackson.version>
+        <slf4j.version>2.0.17</slf4j.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-servicebus</artifactId>
+            <version>${azure.servicebus.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.fasterxml.jackson.core</groupId>
+            <artifactId>jackson-databind</artifactId>
+    
... (truncated)
```

### ✅ 23. powershell — 4882ms

**Input:**

```json
{"command":"mvn -q -DskipTests compile","description":"Compile the Java Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 2 completed with exit code 0>
```

### ✅ 24. apply_patch — 755ms

**Input:**

```json
"*** Begin Patch\n*** Update File: pom.xml\n@@\n         \u003cjackson.version\u003e2.22.2\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n         \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n@@\n         \u003cdependency\u003e\n             \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n             \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n             \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n             \u003cscope\u003eruntime\u003c/scope\u003e\n         \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n@@\n             \u003cplugin\u003e\n                 \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n                 \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n                 \u003cversion\u003e3.5.1\u003c/version\u003e\n@@\n                     \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n                 \u003c/configuration\u003e\n             \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n         \u003c/plugins\u003e\n*** Update File: src/main/java/com/example/orders/Order.java\n@@\n import com.fasterxml.jackson.annotation.JsonCreator;\n import com.fasterxml.jackson.annotation.JsonProperty;\n+import com.fasterxml.jackson.annotation.JsonValue;\n \n import java.math.BigDecimal;\n+import java.util.Locale;\n import java.util.Objects;\n@@\n         PROCESSING,\n         COMPLETED,\n-        FAILED\n+        FAILED;\n+\n+        @JsonCreator\n+        public static Status fromJson(String value) {\n+            return Status.valueOf(value.toUpperCase(Locale.ROOT));\n+        }\n+\n+        @JsonValue\n+        public String toJson() {\n+            return name().toLowerCase(Locale.ROOT);\n+        }\n*** Add File: src/test/java/com/example/orders/OrderTest.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import org.junit.jupiter.api.Test;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertNotNull;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class OrderTest {\n+    private final ObjectMapper objectMapper = new ObjectMapper();\n+\n+    @Test\n+    void serializesAndDeserializesOrder() throws Exception {\n+        Order original = order(\"order-1\", new BigDecimal(\"42.50\"));\n+\n+        String json = objectMapper.writeValueAsString(original);\n+        Order restored = objectMapper.readValue(json, Order.class);\n+\n+        assertTrue(json.contains(\"\\\"status\\\":\\\"pending\\\"\"));\n+        assertEquals(original.getOrderId(), restored.getOrderId());\n+        assertEquals(original.getCustomerName(), restored.getCustomerName());\n+        assertEquals(original.getProduct(), restored.getProduct());\n+        assertEquals(original.getQuantity(), restored.getQuantity());\n+        assertEquals(original.getTotalPrice(), restored.getTotalPrice());\n+        assertEquals(original.getStatus(), restored.getStatus());\n+    }\n+\n+    @Test\n+    void highValueOrderHasSessionCorrelationPriorityAndDelay() {\n+        OffsetDateTime earliestExpected = OffsetDateTime.now().plusSeconds(29);\n+        ServiceBusMessage message = new OrderMessageFactory(objectMapper, new BigDecimal(\"1000\"))\n+                .create(order(\"order-2\", new BigDecimal(\"1000.01\")));\n+\n+        assertEquals(\"order-2\", message.getCorrelationId());\n+        assertEquals(\"order-2\", message.getMessageId());\n+        assertEquals(\"Ada\", message.getSessionId());\n+        assertEquals(\"high\", message.getApplicationProperties().get(\"priority\"));\n+        assertNotNull(message.getScheduledEnqueueTime());\n+        assertTrue(message.getScheduledEnqueueTime().isAfter(earliestExpected));\n+    }\n+\n+    private static Order order(String id, BigDecimal total) {\n+        return new Order(id, \"Ada\", \"Keyboard\", 1, total, Order.Status.PENDING);\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/pom.xml
@@ -15,6 +15,7 @@
         <azure.servicebus.version>7.17.20</azure.servicebus.version>
         <azure.identity.version>1.18.1</azure.identity.version>
         <jackson.version>2.22.2</jackson.version>
+        <junit.version>5.13.4</junit.version>
         <slf4j.version>2.0.17</slf4j.version>
     </properties>
 
@@ -40,6 +41,12 @@
             <version>${slf4j.version}</version>
             <scope>runtime</scope>
         </dependency>
+        <dependency>
+            <groupId>org.junit.jupiter</groupId>
+            <artifactId>junit-jupiter</artifactId>
+            <version>${junit.version}</version>
+            <scope>test</scope>
+        </dependency>
     </dependencies>
 
     <build>
@@ -57,6 +64,11 @@
                     <mainClass>com.example.orders.Main</mainClass>
                 </configuration>
             </plugin>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-surefire-plugin</artifactId>
+                <version>3.5.3</version>
+            </plugin>
         </plugins>
     </build>
 </project>


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-3710544099/src/main/java/com/example/orders/Order.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-servi
... (truncated)
```

### ✅ 25. powershell — 6957ms

**Input:**

```json
{"command":"mvn -q test","description":"Run Maven tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\orders\AsyncOrderProcessor.java`
- `src\main\java\com\example\orders\AsyncOrderSender.java`
- `src\main\java\com\example\orders\Main.java`
- `src\main\java\com\example\orders\Order.java`
- `src\main\java\com\example\orders\OrderMessageFactory.java`
- `src\main\java\com\example\orders\ServiceBusMessageReprocessor.java`
- `src\main\java\com\example\orders\SyncOrderProcessor.java`
- `src\main\java\com\example\orders\SyncOrderSender.java`
- `src\test\java\com\example\orders\OrderTest.java`

## Copilot Response

I’m grounding the implementation in the current Java SDK patterns, then I’ll scaffold and verify the local Maven project.

The key design constraint is Service Bus sessions: messages are batched per customer session, and processors use one concurrent session/call so customer streams remain ordered without cross-customer interleaving.

The project is scaffolded; I’m compiling against the pinned SDK now and will correct any API mismatches before handing it off.

Created a complete Java 17 Maven project with:

- JSON-serializable `Order` model and lowercase statuses.
- Synchronous and reactive asynchronous senders.
- Size-aware batching grouped by customer session.
- Correlation IDs, high-priority metadata, and 30-second scheduling for high-value orders.
- Ordered session-based processors with explicit completion and dead-lettering.
- DLQ inspection and valid-message reprocessing.
- Managed identity authentication through `DefaultAzureCredential`.
- Sync-first, async-second demonstration in `Main`.
- Unit tests and setup documentation.

Run against an existing **session-enabled** queue:

```powershell
$env:AZURE_SERVICE_BUS_NAMESPACE = "contoso.servicebus.windows.net"
$env:AZURE_SERVICE_BUS_QUEUE = "orders"
mvn compile exec:java
```

SDK patterns follow the official [Azure Service Bus Java documentation](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme) and [Service Bus sessions guidance](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions).

## Reviewer Notes (LLM-as-Judge)

### Summary

All Azure SDK client instances have lifecycle management through try-with-resources-backed wrappers.

### Strengths

- Both synchronous and asynchronous client sets are closed through AutoCloseable abstractions used in try-with-resources.

## Grader Results

- order-processor.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (9/12)
      - Sender uses `.sender().queueName().buildClient()` chain (or async equivalent): Pass
      - Processor uses `.processor().queueName().processMessage().processError()` chain: Fail
      - Batch sending: creates `ServiceBusMessageBatch`, checks `tryAddMessage()` return value: Pass
      - Handles the case where a message doesn't fit in the current batch: Pass
      - Scheduled delivery: uses `scheduleMessage()` or `setScheduledEnqueueTime()` (~30s delay): Pass
      - Correlation: sets order ID as correlation property via `setCorrelationId()` or application properties: Pass
      - Dead-letter: explicitly dead-letters failed messages with `deadLetter()` and a reason string: Pass
      - Dead-letter queue reading: uses `SubQueue.DEAD_LETTER_QUEUE` or `$deadletterqueue` path: Pass
      - Session-aware processing: uses `.sessionProcessor()` or session-enabled receiver: Pass
      - Session ID keyed by customer name for ordered processing: Pass
      - Error handler in processor logs entity path and error source: Fail
      - Distinguishes transient vs non-transient errors via `isTransient()` or `getReason()`: Fail
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
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Fail
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 7.7% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 10.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 9.7500** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-java-order-processor --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

# Evaluation Report: service-bus-dp-java-order-processor

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 494.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-order-processor` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 19/24 |
| Duration | 494.6s |
| Timestamp | 2026-08-29T03:06:18Z |
| Files Generated | 12 |
| Event Count | 16088 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 245.1s |
| Review | 246.7s |
| **Total** | **494.6s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-skills/azure-skill-mcp

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=780406 out=13331 |
| Turn Count | 14 |

## Tool Calls

### ✅ 1. skill — 2882ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 7322ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Service Bus queues, sessions for per-customer ordering, managed identity, sync and async senders/processors, batching, scheduling, and dead-letter handling. Do not deploy or modify Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. glob — 8711ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. glob — 5989ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 7381ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Java Azure Service Bus code-generation best practices for a local sample only.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 8444ms

**Input:**

```json
{"intent":"Find current official Java Azure Service Bus SDK guidance and examples for DefaultAzureCredential managed identity, queue sessions, session processors, sync and async clients, message batching with tryAddMessage, scheduled messages, correlation ID/application properties, dead-lettering, and receiving from the dead-letter subqueue.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. view — 8080ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
@@ -1,0 +1,1 @@
 SKILL.md

```

### ✅ 8. view — 7731ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449 b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449
@@ -1,0 +1,0 @@


```

### ✅ 9. azure-documentation (via azure) — 13234ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official current Java Azure Service Bus docs.","parameters":{"query":"Azure Service Bus Java sessions ServiceBusSessionProcessorClient ServiceBusSessionReceiverAsyncClient acceptNextSession dead letter subqueue batch scheduled messages correlation id"}}
```

**Output:**

```
{"results":[{"title":"Azure Service Bus client library for Java - version 7.17.17","content":"# Azure Service Bus client library for Java - version 7.17.17\n## Examples\nAzure Service Bus queues and topic subscriptions provide a secondary sub-queue, called a dead-letter queue (DLQ). The dead-letter queue doesn\u0027t need to be explicitly created and can\u0027t be deleted or otherwise managed independent of the main entity. For session enabled or non-session queue or topic subscriptions, the dead-letter receiver can be created the same way as shown below. Learn more about dead-letter queue [here](https://learn.microsoft.com/en-us/azure/service-bus-messaging/service-bus-dead-letter-queues).\n### When to use \u0060ServiceBusProcessorClient\u0060\nWhen to use \u0060ServiceBusProcessorClient\u0060, \u0060ServiceBusReceiverClient\u0060 or \u0060ServiceBusReceiverAsyncClient\u0060?\nThe best option for receiving messages in most common cases is \u0060ServiceBusProcessorClient\u0060. The processor can handle errors automatically and is designed to receive messages continuously. The processor has a simple API to set concurrency for processing messages in parallel.\nThe low-level client, \u0060ServiceBusReceiverAsyncClient\u0060, is for advanced users who want more control and flexibility over their Reactive application at the expense of more complexity in the application. Unlike the processor, the low-level \u0060ServiceBusReceiverAsyncClient\u0060 does not have automatic recovery built into it. The reactive application using this client needs to deal with terminal events and choose the operators (to add recovery, manage backpressure, threading) in the Reactor chain. As stated earlier, for common asynchronous receive situations, the \u0060ServiceBusProcessorClient\u0060 should be the first choice.\n\u0060ServiceBusReceiverClient\u0060 usage is discouraged unless you want to have pull semantics or want to migrate existing code base that uses synchronous receiver approach. Fo
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 12826ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official current Java code patterns for Service Bus APIs.","parameters":{"language":"java","query":"Azure Service Bus Java ServiceBusMessageBatch tryAddMessage scheduleMessage correlationId sessionId deadLetterMessage SubQueue.DEAD_LETTER_QUEUE ServiceBusSessionReceiverAsyncClient acceptNextSession"}}
```

**Output:**

```
{"results":[{"description":"description: Defines a method to send a batch of messages to an Azure Service Bus queue using connection string authentication.\npackage: com.azure.messaging.servicebus\nlanguage: java\n","codeSnippet":"static void sendMessageBatch()\r\n    {\r\n        // create a Service Bus Sender client for the queue\r\n        ServiceBusSenderClient senderClient = new ServiceBusClientBuilder()\r\n                .connectionString(connectionString)\r\n                .sender()\r\n                .queueName(queueName)\r\n                .buildClient();\r\n    \r\n        // Creates an ServiceBusMessageBatch where the ServiceBus.\r\n        ServiceBusMessageBatch messageBatch = senderClient.createMessageBatch();\r\n    \r\n        // create a list of messages\r\n        List\u003CServiceBusMessage\u003E listOfMessages = createMessages();\r\n    \r\n        // We try to add as many messages as a batch can fit based on the maximum size and send to Service Bus when\r\n        // the batch can hold no more messages. Create a new batch for next set of messages and repeat until all\r\n        // messages are sent.\r\n        for (ServiceBusMessage message : listOfMessages) {\r\n            if (messageBatch.tryAddMessage(message)) {\r\n                continue;\r\n            }\r\n    \r\n            // The batch is full, so we create a new batch and send the batch.\r\n            senderClient.sendMessages(messageBatch);\r\n            System.out.println(\u0022Sent a batch of messages to the queue: \u0022 \u002B queueName);\r\n    \r\n            // create a new batch\r\n            messageBatch = senderClient.createMessageBatch();\r\n    \r\n            // Add that message that we couldn\u0027t before.\r\n            if (!messageBatch.tryAddMessage(message)) {\r\n                System.err.printf(\u0022Message is too large for an empty batch. Skipping. Max size: %s.\u0022, messageBatch.getMaxSizeInBytes());\r\n            }\r\n        }\r\n    \r\n        if 
... (truncated)
```

### ✅ 11. web_search (via github-mcp-server) — 13630ms

**Input:**

```json
{"query":"What are the latest stable Maven Central versions as of August 2026 for com.azure:azure-messaging-servicebus, com.azure:azure-identity, and com.fasterxml.jackson.core:jackson-databind? Cite Maven Central or official artifact pages."}
```

**Output:**

```
{"type":"output_text","text":{"value":"As of August 2026, the latest stable Maven Central versions for your requested artifacts are:\n\n- **com.azure:azure-messaging-servicebus**: The latest release is **7.18.0-beta.3**, but the most recent stable (non-beta) version is **7.17.2**. Use stable versions for production unless beta features are specifically required【3:2†source】.\n- **com.azure:azure-identity**: The latest release series includes betas (e.g., 1.19.0-beta.2), but the latest stable version available is **1.18.4** (released June 2026)【3:3†source】.\n- **com.fasterxml.jackson.core:jackson-databind**: While the precise August 2026 release isn't listed above, the most recent stable releases can always be checked directly on [Maven Central Jackson Databind](https://search.maven.org/artifact/com.fasterxml.jackson.core/jackson-databind) for daily updates. As of June 2026, **2.16.x** is the latest stable series.\n\n**Citations:**\n- [azure-messaging-servicebus Maven Central](https://search.maven.org/artifact/com.azure/azure-messaging-servicebus)\n- [azure-identity Maven Central](https://search.maven.org/artifact/com.azure/azure-identity)\n- [jackson-databind Maven Central](https://search.maven.org/artifact/com.fasterxml.jackson.core/jackson-databind) \n\nYou can always verify latest versions by searching those artifact coordinates on Maven Central for real-time accuracy.","annotations":[{"text":"","start_index":322,"end_index":334,"url_citation":{"title":"Maven Central: com.azure:azure-messaging-servicebus","url":"https://central.sonatype.com/artifact/com.azure/azure-messaging-servicebus"}},{"text":"","start_index":506,"end_index":518,"url_citation":{"title":"Maven Central: com.azure:azure-identity","url":"https://central.sonatype.com/artifact/com.azure/azure-identity/versions"}}]},"bing_searches":[{"text":"latest stable Maven Central versions August 2026 for com.azure:azure-messaging-servicebus, com.azure:azure-identity, com.fasterxml.jackson.core:jacks
... (truncated)
```

### ✅ 12. rg — 2430ms

**Input:**

```json
{"-C":2,"head_limit":120,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787972842059-copilot-tool-output-3f7f50139c084348961a283dae43e625.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787972842539-copilot-tool-output-8627e76f25f2420e96b14985b0f8ee3f.txt"],"pattern":"ServiceBusSessionReceiverAsyncClient|acceptNextSession|ServiceBusSessionProcessorClient|deadLetterMessage|scheduleMessage|tryAddMessage|setCorrelationId|setSessionId"}
```

**Output:**

```
Output too large to read at once (109.5 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787972860835-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787972842059-copilot-tool-output-3f7f50139c084348961a283dae43e625.txt:{"results":[{"description":"description: Defines a method to send a batch of messages to an Azure Service Bus queue using connection string authentication.\npackage: com.azure.messaging.servicebus\nlanguage: java\n","codeSnippet":"static void sendMessageBatch()\r\n    {\r\n        // create a Service Bus Sender client for the queue\r\n        ServiceBusSenderClient senderClient = new Service
```

### ✅ 13. web_fetch — 3453ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-messaging-servicebus/maven-metadata.xml"}
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

### ✅ 14. web_fetch — 3101ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 15. web_fetch — 2760ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-databind/maven-metadata.xml"}
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
      <version>2.4.5.1</version>
      <version>2.4.5</version>
      <version>2.4.6.1</version>
      <version>2.4.6</version>
      <version>2.5.0-rc1</version>
      <version>2.5.0</version>
      <version>2.5.1</version>
      <version>2.5.2</version>
      <version>2.5.3</version>
      <version>2.5.4</version>
      <version>2.
... (truncated)
```

### ✅ 16. powershell — 3401ms

**Input:**

```json
{"command":"mvn -version","description":"Check Maven and Java versions","initial_wait":30}
```

**Output:**

```
Apache Maven 3.9.11 (3e54c93a704957b63ee3494413a2b544fd3d825b)
Maven home: C:\apache-maven-3.9.11
Java version: 25.0.1, vendor: Microsoft, runtime: C:\Program Files\Microsoft\jdk-25.0.1.8-hotspot
Default locale: en_US, platform encoding: UTF-8
OS name: "windows 11", version: "10.0", arch: "amd64", family: "windows"
<shellId: 0 completed with exit code 0>
```

### ✅ 17. web_fetch — 4480ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverasyncclient?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceiverasyncclient?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusReceiverAsyncClient Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceiverasyncclient?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>An <strong>asynchronous</strong> receiver responsible for receiving <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceivedmessage'>ServiceBusReceivedMessage</a> from an Azure Service Bus queue or topic/subscription.</p>

  <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication, which is appropriate for most scenarios, including local development and production environments. Additionally, we recommend using <a href="/azure/active-directory/managed-identities-azure-resources/">managed identity</a> for authentication in production environments. You can find more information on different ways of authenticating and their corresponding credential types in the <a href="/java/api/overview/azure/identity-readme">Azure Identity documentation&quot;</a>.</p>

  <p><strong>Sample: Creating a <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceiverasyncclient'>ServiceBusReceiverAsyncClient</a></strong></p>

  <p>The following code sample demonstrates the creation of the asynchronous client <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceiverasyncclient'>ServiceBusReceiverAsyncClient</a>. The <code>fullyQualifiedNamespace</code> is the Service Bus namespace's host name. It is listed under the &quot;Essentials&quot; panel aft
... (truncated)
```

### ✅ 18. web_fetch — 4138ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceiverasyncclient?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverasyncclient?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusSessionReceiverAsyncClient Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverasyncclient?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>This <strong>asynchronous</strong> session receiver client is used to acquire session locks from a queue or topic and create <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceiverasyncclient'>ServiceBusReceiverAsyncClient</a> instances that are tied to the locked sessions. Sessions can be used as a first in, first out (FIFO) processing of messages. Queues and topics/subscriptions support Service Bus sessions, however, it must be <a href="/azure/service-bus-messaging/enable-message-sessions"> enabled at the time of entity creation</a>.</p>

  <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication, which is appropriate for most scenarios, including local development and production environments. Additionally, we recommend using <a href="/azure/active-directory/managed-identities-azure-resources/">managed identity</a> for authentication in production environments. You can find more information on different ways of authenticating and their corresponding credential types in the <a href="/java/api/overview/azure/identity-readme">Azure Identity documentation&quot;</a>.</p>

  <p><strong>Sample: Receive messages from a specific session</strong></p>

  <p>Use <a class=no-loc href='com.azure.messaging.servicebus.servicebussess
... (truncated)
```

### ✅ 19. web_fetch — 3786ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussenderasyncclient?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussenderasyncclient?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusSenderAsyncClient Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussenderasyncclient?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusSenderAsyncClient
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>An <strong>asynchronous</strong> client to send messages to a Service Bus resource.</p>

  <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication, which is appropriate for most scenarios, including local development and production environments. Additionally, we recommend using <a href="/azure/active-directory/managed-identities-azure-resources/">managed identity</a> for authentication in production environments. You can find more information on different ways of authenticating and their corresponding credential types in the <a href="/java/api/overview/azure/identity-readme">Azure Identity documentation&quot;</a>.</p>

  <p><strong>Sample: Create an instance of sender</strong></p>

  <pre><code class="lang-java">TokenCredential credential = new DefaultAzureCredentialBuilder().build();

   // 'fullyQualifiedNamespace' will look similar to &quot;{your-namespace}.servicebus.windows.net&quot;
   ServiceBusSenderAsyncClient asyncSender = new ServiceBusClientBuilder()
       .credential(fullyQualifiedNamespace, credential)
       .sender()
       .queueName(queueName)
       .buildAsyncClient();

   // When users are done with the sender, they should dispose of it.
   // Clients should be long-lived objects as they require resources
   // and time to establish a connection t
... (truncated)
```

### ✅ 20. apply_patch — 3125ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eservice-bus-order-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eService Bus Order Processor\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.messaging.servicebus.version\u003e7.17.20\u003c/azure.messaging.servicebus.version\u003e\n+        \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+        \u003cjackson.version\u003e2.22.2\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.14.0\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.messaging.servicebus.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.4\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.6.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/orders/OrderStatus.java\n+package com.example.orders;\n+\n+import com.fasterxml.jackson.annotation.JsonCreator;\n+import com.fasterxml.jackson.annotation.JsonValue;\n+\n+import java.util.Locale;\n+\n+public enum OrderStatus {\n+    PENDING,\n+    PROCESSING,\n+    COMPLETED,\n+    FAILED;\n+\n+    @JsonValue\n+    public String toJson() {\n+        return name().toLowerCase(Locale.ROOT);\n+    }\n+\n+    @JsonCreator\n+    public static OrderStatus fromJson(String value) {\n+        return valueOf(value.toUpperCase(Locale.ROOT));\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Order.java\n+package com.example.orders;\n+\n+import java.math.BigDecimal;\n+import java.util.Objects;\n+\n+public record Order(\n+        String orderId,\n+        String customerName,\n+        String product,\n+        int quantity,\n+        BigDecimal totalPrice,\n+        OrderStatus status) {\n+\n+    public Order {\n+        Objects.requireNonNull(orderId, \"orderId\");\n+        Objects.requireNonNull(customerName, \"customerName\");\n+        Objects.requireNonNull(product, \"product\");\n+        Objects.requireNonNull(totalPrice, \"totalPrice\");\n+        Objects.requireNonNull(status, \"status\");\n+        if (orderId.isBlank() || customerName.isBlank() || product.isBlank()) {\n+            throw new IllegalArgumentException(\"Order ID, customer name, and product must not be blank\");\n+        }\n+        if (quantity \u003c= 0) {\n+            throw new IllegalArgumentException(\"Quantity must be greater than zero\");\n+        }\n+        if (totalPrice.signum() \u003c 0) {\n+            throw new IllegalArgumentException(\"Total price must not be negative\");\n+        }\n+    }\n+\n+    public Order withStatus(OrderStatus newStatus) {\n+        return new Order(orderId, customerName, product, quantity, totalPrice, newStatus);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderJson.java\n+package com.example.orders;\n+\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.DeserializationFeature;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+public final class OrderJson {\n+    private static final ObjectMapper MAPPER = new ObjectMapper()\n+            .enable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);\n+\n+    private OrderJson() {\n+    }\n+\n+    public static String serialize(Order order) {\n+        try {\n+            return MAPPER.writeValueAsString(order);\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Could not serialize order \" + order.orderId(), exception);\n+        }\n+    }\n+\n+    public static Order deserialize(String json) {\n+        try {\n+            return MAPPER.readValue(json, Order.class);\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Could not deserialize order JSON\", exception);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderMessageFactory.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+import java.time.ZoneOffset;\n+\n+final class OrderMessageFactory {\n+    static final String PRIORITY_PROPERTY = \"priority\";\n+    static final String HIGH_PRIORITY = \"high\";\n+    static final String NORMAL_PRIORITY = \"normal\";\n+    static final int FRAUD_REVIEW_DELAY_SECONDS = 30;\n+\n+    private OrderMessageFactory() {\n+    }\n+\n+    static ServiceBusMessage create(Order order, BigDecimal highPriorityThreshold) {\n+        boolean highPriority = order.totalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+        ServiceBusMessage message = new ServiceBusMessage(OrderJson.serialize(order))\n+                .setContentType(\"application/json\")\n+                .setMessageId(order.orderId())\n+                .setCorrelationId(order.orderId())\n+                .setSessionId(order.customerName());\n+        message.getApplicationProperties().put(\n+                PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY);\n+\n+        if (highPriority) {\n+            message.setScheduledEnqueueTime(\n+                    OffsetDateTime.now(ZoneOffset.UTC).plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));\n+        }\n+        return message;\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class SyncOrderSender implements AutoCloseable {\n+    private final ServiceBusSenderClient sender;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public SyncOrderSender(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildClient();\n+        this.highPriorityThreshold = Objects.requireNonNull(\n+                highPriorityThreshold, \"highPriorityThreshold\");\n+    }\n+\n+    public void sendOrder(Order order) {\n+        sender.sendMessage(OrderMessageFactory.create(order, highPriorityThreshold));\n+    }\n+\n+    public void sendOrders(List\u003cOrder\u003e orders) {\n+        Objects.requireNonNull(orders, \"orders\");\n+        ServiceBusMessageBatch batch = sender.createMessageBatch();\n+\n+        for (Order order : orders) {\n+            ServiceBusMessage message = OrderMessageFactory.create(order, highPriorityThreshold);\n+            if (!batch.tryAddMessage(message)) {\n+                if (batch.getCount() == 0) {\n+                    throw new IllegalArgumentException(\n+                            \"Order message is too large for an empty Service Bus batch: \" + order.orderId());\n+                }\n+                sender.sendMessages(batch);\n+                batch = sender.createMessageBatch();\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalArgumentException(\n+                            \"Order message is too large for an empty Service Bus batch: \" + order.orderId());\n+                }\n+            }\n+        }\n+\n+        if (batch.getCount() \u003e 0) {\n+            sender.sendMessages(batch);\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class AsyncOrderSender implements AutoCloseable {\n+    private final ServiceBusSenderAsyncClient sender;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public AsyncOrderSender(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildAsyncClient();\n+        this.highPriorityThreshold = Objects.requireNonNull(\n+                highPriorityThreshold, \"highPriorityThreshold\");\n+    }\n+\n+    public Mono\u003cVoid\u003e sendOrder(Order order) {\n+        return sender.sendMessage(OrderMessageFactory.create(order, highPriorityThreshold));\n+    }\n+\n+    public Mono\u003cVoid\u003e sendOrders(List\u003cOrder\u003e orders) {\n+        Objects.requireNonNull(orders, \"orders\");\n+        List\u003cServiceBusMessage\u003e messages = orders.stream()\n+                .map(order -\u003e OrderMessageFactory.create(order, highPriorityThreshold))\n+                .toList();\n+        return sendBatchFrom(messages, 0);\n+    }\n+\n+    private Mono\u003cVoid\u003e sendBatchFrom(List\u003cServiceBusMessage\u003e messages, int startIndex) {\n+        if (startIndex \u003e= messages.size()) {\n+            return Mono.empty();\n+        }\n+\n+        return sender.createMessageBatch().flatMap(batch -\u003e {\n+            int nextIndex = fillBatch(batch, messages, startIndex);\n+            return sender.sendMessages(batch).then(sendBatchFrom(messages, nextIndex));\n+        });\n+    }\n+\n+    private static int fillBatch(\n+            ServiceBusMessageBatch batch,\n+            List\u003cServiceBusMessage\u003e messages,\n+            int startIndex) {\n+        int index = startIndex;\n+        while (index \u003c messages.size() \u0026\u0026 batch.tryAddMessage(messages.get(index))) {\n+            index++;\n+        }\n+        if (index == startIndex) {\n+            throw new IllegalArgumentException(\n+                    \"Message is too large for an empty Service Bus batch: \"\n+                            + messages.get(index).getMessageId());\n+        }\n+        return index;\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionProcessorClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.azure.messaging.servicebus.models.ProcessMessageContext;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+\n+import java.time.Duration;\n+import java.util.Optional;\n+import java.util.function.Function;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+\n+public final class SyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = Logger.getLogger(SyncOrderProcessor.class.getName());\n+    private final String fullyQualifiedNamespace;\n+    private final String queueName;\n+    private final TokenCredential credential;\n+    private final ServiceBusSessionProcessorClient processor;\n+\n+    public SyncOrderProcessor(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential) {\n+        this.fullyQualifiedNamespace = fullyQualifiedNamespace;\n+        this.queueName = queueName;\n+        this.credential = credential;\n+        this.processor = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sessionProcessor()\n+                .queueName(queueName)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .maxConcurrentSessions(4)\n+                .maxConcurrentCalls(1)\n+                .processMessage(this::processMessage)\n+                .processError(context -\u003e LOGGER.log(\n+                        Level.SEVERE,\n+                        \"Service Bus processor error for \" + context.getEntityPath(),\n+                        context.getException()))\n+                .buildProcessorClient();\n+    }\n+\n+    public void start() {\n+        processor.start();\n+    }\n+\n+    public void stop() {\n+        processor.stop();\n+    }\n+\n+    private void processMessage(ProcessMessageContext context) {\n+        ServiceBusReceivedMessage message = context.getMessage();\n+        try {\n+            Order pending = OrderJson.deserialize(message.getBody().toString());\n+            Order processing = pending.withStatus(OrderStatus.PROCESSING);\n+            LOGGER.info(() -\u003e \"Processing sync order: \" + OrderJson.serialize(processing));\n+            Order completed = processing.withStatus(OrderStatus.COMPLETED);\n+            LOGGER.info(() -\u003e \"Completed sync order: \" + OrderJson.serialize(completed));\n+            context.complete();\n+        } catch (IllegalArgumentException exception) {\n+            String reason = \"ORDER_DESERIALIZATION_FAILED\";\n+            LOGGER.log(Level.WARNING, \"Dead-lettering order message \" + message.getMessageId(), exception);\n+            context.deadLetter(new DeadLetterOptions()\n+                    .setDeadLetterReason(reason)\n+                    .setDeadLetterErrorDescription(exception.getMessage()));\n+        }\n+    }\n+\n+    public void reprocessDeadLetters(\n+            int maximumMessages,\n+            Duration receiveWindow,\n+            Function\u003cServiceBusReceivedMessage, Optional\u003cOrder\u003e\u003e recovery,\n+            SyncOrderSender sender) {\n+        try (ServiceBusSessionReceiverClient sessionReceiver = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sessionReceiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildClient();\n+             ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {\n+            for (ServiceBusReceivedMessage message\n+                    : receiver.receiveMessages(maximumMessages, receiveWindow)) {\n+                LOGGER.warning(() -\u003e \"Inspecting dead letter: id=\" + message.getMessageId()\n+                        + \", reason=\" + message.getDeadLetterReason()\n+                        + \", description=\" + message.getDeadLetterErrorDescription());\n+                Optional\u003cOrder\u003e recovered = recovery.apply(message);\n+                if (recovered.isPresent()) {\n+                    sender.sendOrder(recovered.get());\n+                    receiver.complete(message);\n+                    LOGGER.info(() -\u003e \"Requeued recovered order \" + recovered.get().orderId());\n+                } else {\n+                    receiver.abandon(message);\n+                }\n+            }\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        processor.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.Optional;\n+import java.util.function.Function;\n+import java.util.logging.Logger;\n+\n+public final class AsyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = Logger.getLogger(AsyncOrderProcessor.class.getName());\n+    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;\n+    private final ServiceBusSessionReceiverAsyncClient deadLetterSessionReceiver;\n+\n+    public AsyncOrderProcessor(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential) {\n+        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential);\n+        this.sessionReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+        this.deadLetterSessionReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+    }\n+\n+    public Mono\u003cVoid\u003e processAvailableSessions(int sessionCount, Duration processingWindow) {\n+        return Flux.range(0, sessionCount)\n+                .flatMap(ignored -\u003e Mono.usingWhen(\n+                        sessionReceiver.acceptNextSession(),\n+                        receiver -\u003e receiver.receiveMessages()\n+                                .concatMap(message -\u003e processMessage(receiver, message))\n+                                .take(processingWindow)\n+                                .then(),\n+                        receiver -\u003e Mono.fromRunnable(receiver::close)), sessionCount)\n+                .then();\n+    }\n+\n+    private Mono\u003cVoid\u003e processMessage(\n+            ServiceBusReceiverAsyncClient receiver,\n+            ServiceBusReceivedMessage message) {\n+        try {\n+            Order pending = OrderJson.deserialize(message.getBody().toString());\n+            Order processing = pending.withStatus(OrderStatus.PROCESSING);\n+            LOGGER.info(() -\u003e \"Processing async order: \" + OrderJson.serialize(processing));\n+            Order completed = processing.withStatus(OrderStatus.COMPLETED);\n+            LOGGER.info(() -\u003e \"Completed async order: \" + OrderJson.serialize(completed));\n+            return receiver.complete(message);\n+        } catch (IllegalArgumentException exception) {\n+            LOGGER.warning(() -\u003e \"Dead-lettering order message \" + message.getMessageId()\n+                    + \": \" + exception.getMessage());\n+            return receiver.deadLetter(message, new DeadLetterOptions()\n+                    .setDeadLetterReason(\"ORDER_DESERIALIZATION_FAILED\")\n+                    .setDeadLetterErrorDescription(exception.getMessage()));\n+        }\n+    }\n+\n+    public Mono\u003cVoid\u003e reprocessDeadLetters(\n+            int maximumMessages,\n+            Duration receiveWindow,\n+            Function\u003cServiceBusReceivedMessage, Optional\u003cOrder\u003e\u003e recovery,\n+            AsyncOrderSender sender) {\n+        return Mono.usingWhen(\n+                deadLetterSessionReceiver.acceptNextSession(),\n+                receiver -\u003e receiver.receiveMessages()\n+                        .take(receiveWindow)\n+                        .take(maximumMessages)\n+                        .concatMap(message -\u003e {\n+                            LOGGER.warning(() -\u003e \"Inspecting dead letter: id=\" + message.getMessageId()\n+                                    + \", reason=\" + message.getDeadLetterReason()\n+                                    + \", description=\" + message.getDeadLetterErrorDescription());\n+                            Optional\u003cOrder\u003e recovered = recovery.apply(message);\n+                            if (recovered.isEmpty()) {\n+                                return receiver.abandon(message);\n+                            }\n+                            return sender.sendOrder(recovered.get())\n+                                    .then(receiver.complete(message))\n+                                    .doOnSuccess(ignored -\u003e LOGGER.info(\n+                                            () -\u003e \"Requeued recovered order \" + recovered.get().orderId()));\n+                        })\n+                        .then(),\n+                receiver -\u003e Mono.fromRunnable(receiver::close));\n+    }\n+\n+    @Override\n+    public void close() {\n+        sessionReceiver.close();\n+        deadLetterSessionReceiver.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Main.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Optional;\n+\n+public final class Main {\n+    private static final Duration PROCESSING_WINDOW = Duration.ofSeconds(35);\n+    private static final String SYNC_BAD_ORDER_ID = \"sync-invalid\";\n+    private static final String ASYNC_BAD_ORDER_ID = \"async-invalid\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String namespace = requiredEnvironment(\"SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE\");\n+        String queueName = requiredEnvironment(\"SERVICE_BUS_QUEUE_NAME\");\n+        BigDecimal priorityThreshold = new BigDecimal(\n+                System.getenv().getOrDefault(\"HIGH_PRIORITY_THRESHOLD\", \"1000.00\"));\n+        TokenCredential credential = new DefaultAzureCredentialBuilder().build();\n+\n+        runSyncDemo(namespace, queueName, credential, priorityThreshold);\n+        runAsyncDemo(namespace, queueName, credential, priorityThreshold).block();\n+    }\n+\n+    private static void runSyncDemo(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal threshold) throws InterruptedException {\n+        try (SyncOrderSender sender = new SyncOrderSender(namespace, queueName, credential, threshold);\n+             SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential)) {\n+            sender.sendOrders(sampleOrders(\"sync\"));\n+            sendMalformedSync(namespace, queueName, credential, SYNC_BAD_ORDER_ID);\n+            processor.start();\n+            Thread.sleep(PROCESSING_WINDOW.toMillis());\n+            processor.stop();\n+            processor.reprocessDeadLetters(\n+                    10,\n+                    Duration.ofSeconds(5),\n+                    message -\u003e recoverDemoOrder(message.getCorrelationId(), \"sync\"),\n+                    sender);\n+        }\n+    }\n+\n+    private static Mono\u003cVoid\u003e runAsyncDemo(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal threshold) {\n+        return Mono.using(\n+                () -\u003e new AsyncOrderSender(namespace, queueName, credential, threshold),\n+                sender -\u003e Mono.using(\n+                        () -\u003e new AsyncOrderProcessor(namespace, queueName, credential),\n+                        processor -\u003e sender.sendOrders(sampleOrders(\"async\"))\n+                                .then(sendMalformedAsync(\n+                                        namespace, queueName, credential, ASYNC_BAD_ORDER_ID))\n+                                .then(processor.processAvailableSessions(3, PROCESSING_WINDOW))\n+                                .then(processor.reprocessDeadLetters(\n+                                        10,\n+                                        Duration.ofSeconds(5),\n+                                        message -\u003e recoverDemoOrder(message.getCorrelationId(), \"async\"),\n+                                        sender)),\n+                        AsyncOrderProcessor::close),\n+                AsyncOrderSender::close);\n+    }\n+\n+    private static List\u003cOrder\u003e sampleOrders(String prefix) {\n+        return List.of(\n+                new Order(prefix + \"-001\", \"Alice\", \"Laptop\", 1,\n+                        new BigDecimal(\"1499.00\"), OrderStatus.PENDING),\n+                new Order(prefix + \"-002\", \"Alice\", \"Mouse\", 2,\n+                        new BigDecimal(\"79.98\"), OrderStatus.PENDING),\n+                new Order(prefix + \"-003\", \"Bob\", \"Monitor\", 1,\n+                        new BigDecimal(\"399.00\"), OrderStatus.PENDING));\n+    }\n+\n+    private static void sendMalformedSync(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            String orderId) {\n+        try (ServiceBusSenderClient sender = new ServiceBusClientBuilder()\n+                .credential(namespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildClient()) {\n+            sender.sendMessage(malformedMessage(orderId));\n+        }\n+    }\n+\n+    private static Mono\u003cVoid\u003e sendMalformedAsync(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            String orderId) {\n+        return Mono.using(\n+                () -\u003e {\n+                    ServiceBusSenderAsyncClient sender = new ServiceBusClientBuilder()\n+                            .credential(namespace, credential)\n+                            .sender()\n+                            .queueName(queueName)\n+                            .buildAsyncClient();\n+                    return sender;\n+                },\n+                sender -\u003e sender.sendMessage(malformedMessage(orderId)),\n+                ServiceBusSenderAsyncClient::close);\n+    }\n+\n+    private static ServiceBusMessage malformedMessage(String orderId) {\n+        return new ServiceBusMessage(\"{not-valid-json\")\n+                .setContentType(\"application/json\")\n+                .setMessageId(orderId)\n+                .setCorrelationId(orderId)\n+                .setSessionId(\"DemoFailure\");\n+    }\n+\n+    private static Optional\u003cOrder\u003e recoverDemoOrder(String correlationId, String prefix) {\n+        if (!(prefix + \"-invalid\").equals(correlationId)) {\n+            return Optional.empty();\n+        }\n+        return Optional.of(new Order(\n+                prefix + \"-recovered\",\n+                \"DemoFailure\",\n+                \"Recovered Product\",\n+                1,\n+                new BigDecimal(\"25.00\"),\n+                OrderStatus.PENDING));\n+    }\n+\n+    private static String requiredEnvironment(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/test/java/com/example/orders/OrderJsonTest.java\n+package com.example.orders;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.math.BigDecimal;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class OrderJsonTest {\n+    @Test\n+    void roundTripsOrderWithLowerCaseStatus() {\n+        Order order = new Order(\n+                \"order-1\",\n+                \"Ada\",\n+                \"Keyboard\",\n+                2,\n+                new BigDecimal(\"199.98\"),\n+                OrderStatus.PENDING);\n+\n+        String json = OrderJson.serialize(order);\n+\n+        assertTrue(json.contains(\"\\\"status\\\":\\\"pending\\\"\"));\n+        assertEquals(order, OrderJson.deserialize(json));\n+    }\n+\n+    @Test\n+    void rejectsMalformedJson() {\n+        assertThrows(IllegalArgumentException.class, () -\u003e OrderJson.deserialize(\"{bad-json\"));\n+    }\n+}\n*** Add File: README.md\n+# Service Bus Order Processor\n+\n+Java 17 example with synchronous and asynchronous Azure Service Bus senders and\n+session-aware processors. The queue must have **sessions enabled when it is\n+created**. Each message uses the customer name as its session ID, which preserves\n+FIFO processing for that customer while allowing different customer sessions to\n+run concurrently.\n+\n+## Configuration\n+\n+The authenticated identity needs the **Azure Service Bus Data Sender** and\n+**Azure Service Bus Data Receiver** roles for the queue.\n+\n+```powershell\n+$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = \"your-namespace.servicebus.windows.net\"\n+$env:SERVICE_BUS_QUEUE_NAME = \"orders\"\n+$env:HIGH_PRIORITY_THRESHOLD = \"1000.00\" # optional\n+mvn clean test\n+mvn exec:java\n+```\n+\n+`DefaultAzureCredential` uses managed identity in Azure. For local development,\n+it can use a supported developer credential. The application never accepts a\n+connection string or access key.\n+\n+Orders above the configured threshold receive a `priority=high` application\n+property and a 30-second scheduled enqueue time for fraud review. Batch senders\n+use `ServiceBusMessageBatch.tryAddMessage` and split full batches before sending.\n+Malformed demo messages are explicitly dead-lettered, inspected, corrected, and\n+requeued.\n+\n+## References\n+\n+- [Azure Service Bus Java client library](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme)\n+- [Service Bus message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions)\n+- [Service Bus dead-letter queues](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-dead-letter-queues)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/pom.xml
@@ -1,0 +1,68 @@
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
+        <azure.messaging.servicebus.version>7.17.20</azure.messaging.servicebus.version>
+        <azure.identity.version>1.18.5</azure.identity.version>
+        <jackson.version>2.22.2</jackson.version>
+        <junit.version>5.14.0</junit.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-servicebus</artifactId>
+            <version>${azure.messaging.servicebus.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.fasterxml.jackson.core</groupId>
+            <artifactId>jackson-databind</artifactId>
+            <version>$
... (truncated)
```

### ✅ 21. powershell — 15074ms

**Input:**

```json
{"command":"mvn --no-transfer-progress test","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example:service-bus-order-processor >---------------
[INFO] Building Service Bus Order Processor 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ service-bus-order-processor ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ service-bus-order-processor ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/SyncOrderProcessor.java:[7,38] cannot find symbol
  symbol:   class ServiceBusSessionProcessorClient
  location: package com.azure.messaging.servicebus
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/SyncOrderProcessor.java:[10,45] cannot find symbol
  symbol:   class ProcessMessageContext
  location: package com.azure
... (truncated)
```

### ✅ 22. glob — 3258ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.m2\\repository","pattern":"**/azure-messaging-servicebus-7.17.20.jar"}
```

**Output:**

```
C:\Users\weidxu\.m2\repository\com\azure\azure-messaging-servicebus\7.17.20\azure-messaging-servicebus-7.17.20.jar
```

### ✅ 23. web_fetch — 2879ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusclientbuilder.servicebussessionprocessorclientbuilder?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceivedmessagecontext?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusReceivedMessageContext Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceivedmessagecontext?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusReceivedMessageContext
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>The Service Bus processor message context that holds a received message and additional methods to settle the message.</p>
package: com.azure.messaging.servicebus
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: generated-reference
ms.devlang: java
ms.date: 2018-07-17T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: routlaw
author: rloutlaw
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.service: azure
langs:
- java
locale: en-us
document_id: bbe38e02-50f9-9a4a-1b21-a7ed2912183a
document_version_independent_id: abf04ef9-e64f-3297-e613-042cac0ae5e2
updated_at: 2026-01-30T12:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-autogen/com.azure.messaging.servicebus.ServiceBusReceivedMessageContext.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/283a699cbf1b12060a29093053119fed1a7754be/docs-ref-autogen/com.azure.messaging.servicebus.ServiceBusReceivedMessageContext.yml
git_commit_id: 283a699cbf1b12060a29093053119fed1a7754be
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: java
page_kind: class
m
... (truncated)
```

### ✅ 24. web_fetch — 2470ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceivedmessagecontext?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusclientbuilder.servicebussessionprocessorclientbuilder?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusClientBuilder.ServiceBusSessionProcessorClientBuilder Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusclientbuilder.servicebussessionprocessorclientbuilder?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusClientBuilder.ServiceBusSessionProcessorClientBuilder
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>Builder for creating <a class=no-loc href='com.azure.messaging.servicebus.servicebusprocessorclient'>ServiceBusProcessorClient</a> to consume messages from a session-based Service Bus entity. <a class=no-loc href='com.azure.messaging.servicebus.servicebusprocessorclient'>ServiceBusProcessorClient</a> processes messages and errors via <a class=no-loc href='com.azure.messaging.servicebus.servicebusclientbuilder.servicebussessionprocessorclientbuilder#com-azure-messaging-servicebus-servicebusclientbuilder-servicebussessionprocessorclientbuilder-processmessage(java-util-function-consumer(com-azure-messaging-servicebus-servicebusreceivedmessagecontext))'>processMessage(Consumer&lt;ServiceBusReceivedMessageContext&gt; processMessage)</a> and <a class=no-loc href='com.azure.messaging.servicebus.servicebusclientbuilder.servicebussessionprocessorclientbuilder#com-azure-messaging-servicebus-servicebusclientbuilder-servicebussessionprocessorclientbuilder-processerror(java-util-function-consumer(com-azure-messaging-servicebus-servicebuserrorcontext))'>processError(Consumer&lt;ServiceBusErrorContext&gt; processError)</a>. When the processor finishes processing a session, it tries to fetch the n
... (truncated)
```

### ✅ 25. apply_patch — 826ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/SyncOrderProcessor.java\n@@\n import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;\n import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n-import com.azure.messaging.servicebus.ServiceBusSessionProcessorClient;\n import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n import com.azure.messaging.servicebus.models.DeadLetterOptions;\n-import com.azure.messaging.servicebus.models.ProcessMessageContext;\n import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n@@\n-    private final ServiceBusSessionProcessorClient processor;\n+    private final ServiceBusProcessorClient processor;\n@@\n-    private void processMessage(ProcessMessageContext context) {\n+    private void processMessage(ServiceBusReceivedMessageContext context) {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/SyncOrderProcessor.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/SyncOrderProcessor.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/SyncOrderProcessor.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/SyncOrderProcessor.java
@@ -2,12 +2,12 @@
 
 import com.azure.core.credential.TokenCredential;
 import com.azure.messaging.servicebus.ServiceBusClientBuilder;
+import com.azure.messaging.servicebus.ServiceBusProcessorClient;
 import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
+import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;
 import com.azure.messaging.servicebus.ServiceBusReceiverClient;
-import com.azure.messaging.servicebus.ServiceBusSessionProcessorClient;
 import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
 import com.azure.messaging.servicebus.models.DeadLetterOptions;
-import com.azure.messaging.servicebus.models.ProcessMessageContext;
 import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
 import com.azure.messaging.servicebus.models.SubQueue;
 
@@ -22,7 +22,7 @@
     private final String fullyQualifiedNamespace;
     private final String queueName;
     private final TokenCredential credential;
-    private final ServiceBusSessionProcessorClient processor;
+    private final ServiceBusProcessorClient processor;
 
     public SyncOrderProcessor(
             String fullyQualifiedNamespace,
@@ -55,7 +55,7 @@
         processor.stop();
     }
 
-    priva
... (truncated)
```

### ✅ 26. powershell — 8130ms

**Input:**

```json
{"command":"mvn --no-transfer-progress test","description":"Recompile project and run tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example:service-bus-order-processor >---------------
[INFO] Building Service Bus Order Processor 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ service-bus-order-processor ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ service-bus-order-processor ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ service-bus-order-processor ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ service-bus-order-processor ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 1 source file with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.4:test (default-test) @ service-bus-order-processor ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider
[INFO] 
[INFO] ---------------------------
... (truncated)
```

### ✅ 27. apply_patch — 4312ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n             processor.reprocessDeadLetters(\n                     10,\n                     Duration.ofSeconds(5),\n                     message -\u003e recoverDemoOrder(message.getCorrelationId(), \"sync\"),\n                     sender);\n+            processor.start();\n+            Thread.sleep(Duration.ofSeconds(5).toMillis());\n+            processor.stop();\n@@\n                                 .then(processor.reprocessDeadLetters(\n                                         10,\n                                         Duration.ofSeconds(5),\n                                         message -\u003e recoverDemoOrder(message.getCorrelationId(), \"async\"),\n-                                        sender)),\n+                                        sender))\n+                                .then(processor.processAvailableSessions(\n+                                        1, Duration.ofSeconds(5))),\n*** Update File: src/test/java/com/example/orders/OrderJsonTest.java\n@@\n package com.example.orders;\n \n+import com.azure.messaging.servicebus.ServiceBusMessage;\n import org.junit.jupiter.api.Test;\n \n import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n@@\n     void rejectsMalformedJson() {\n         assertThrows(IllegalArgumentException.class, () -\u003e OrderJson.deserialize(\"{bad-json\"));\n     }\n+\n+    @Test\n+    void createsSessionAwareHighPriorityScheduledMessage() {\n+        Order order = new Order(\n+                \"order-2\",\n+                \"Grace\",\n+                \"Workstation\",\n+                1,\n+                new BigDecimal(\"2500.00\"),\n+                OrderStatus.PENDING);\n+        OffsetDateTime before = OffsetDateTime.now().plusSeconds(29);\n+\n+        ServiceBusMessage message = OrderMessageFactory.create(order, new BigDecimal(\"1000.00\"));\n+\n+        assertEquals(order.orderId(), message.getCorrelationId());\n+        assertEquals(order.customerName(), message.getSessionId());\n+        assertEquals(OrderMessageFactory.HIGH_PRIORITY,\n+                message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));\n+        assertTrue(message.getScheduledEnqueueTime().isAfter(before));\n+        assertTrue(Duration.between(before, message.getScheduledEnqueueTime()).toSeconds() \u003c= 2);\n+    }\n }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/Main.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/Main.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/Main.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/main/java/com/example/orders/Main.java
@@ -49,6 +49,9 @@
                     Duration.ofSeconds(5),
                     message -> recoverDemoOrder(message.getCorrelationId(), "sync"),
                     sender);
+            processor.start();
+            Thread.sleep(Duration.ofSeconds(5).toMillis());
+            processor.stop();
         }
     }
 
@@ -69,7 +72,9 @@
                                         10,
                                         Duration.ofSeconds(5),
                                         message -> recoverDemoOrder(message.getCorrelationId(), "async"),
-                                        sender)),
+                                        sender))
+                                .then(processor.processAvailableSessions(
+                                        1, Duration.ofSeconds(5))),
                         AsyncOrderProcessor::close),
                 AsyncOrderSender::close);
     }


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/test/java/com/example/orders/OrderJsonTest.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449/src/test/java/com/example/orders/OrderJsonTest.java
index 000000
... (truncated)
```

### ✅ 28. powershell — 8415ms

**Input:**

```json
{"command":"mvn --no-transfer-progress clean test","description":"Run clean build and tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example:service-bus-order-processor >---------------
[INFO] Building Service Bus Order Processor 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ service-bus-order-processor ---
[INFO] Deleting C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449\target
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ service-bus-order-processor ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ service-bus-order-processor ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ service-bus-order-processor ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-3326867449\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ service-bus-order-processor ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 1 source file with javac [debug release 17] to target\test-clas
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\orders\AsyncOrderProcessor.java`
- `src\main\java\com\example\orders\AsyncOrderSender.java`
- `src\main\java\com\example\orders\Main.java`
- `src\main\java\com\example\orders\Order.java`
- `src\main\java\com\example\orders\OrderJson.java`
- `src\main\java\com\example\orders\OrderMessageFactory.java`
- `src\main\java\com\example\orders\OrderStatus.java`
- `src\main\java\com\example\orders\SyncOrderProcessor.java`
- `src\main\java\com\example\orders\SyncOrderSender.java`
- `src\test\java\com\example\orders\OrderJsonTest.java`

## Copilot Response

I’m grounding the implementation in current Azure SDK patterns, then I’ll scaffold the project and validate it locally without contacting Azure.

The key design constraint is Service Bus sessions: every message will use the customer name as its session ID, and both processors will accept one session at a time so each customer’s orders remain sequential.

I’ll keep the demo self-contained but non-provisioning: it expects an existing **session-enabled** queue and only uses `DefaultAzureCredential`; no connection strings, keys, or Azure resource changes are included.

The current SDK uses the shared `ServiceBusProcessorClient` and `ServiceBusReceivedMessageContext` types for session processors; I’m aligning the callback implementation to those concrete APIs.

The project now compiles against the actual SDK. I’m tightening the demo so a recovered DLQ order is not only requeued but consumed again, completing the requested end-to-end cycle.

Created the Java 17 Maven project with:

- JSON-serializable `Order` model and lowercase status enum.
- Synchronous and reactive asynchronous senders with size-safe batching.
- Correlation IDs, customer session IDs, priority properties, and 30-second fraud-review scheduling.
- Session-aware sync/async processors with explicit settlement and dead-letter handling.
- DLQ inspection, recovery, requeueing, and subsequent processing.
- Managed-identity-compatible `DefaultAzureCredential` authentication.
- Executable `Main` demo, tests, configuration instructions, and complete `pom.xml`.

Run against an existing **session-enabled** queue:

```powershell
$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "orders"
$env:HIGH_PRIORITY_THRESHOLD = "1000.00"
mvn exec:java
```

References: [Azure Service Bus Java SDK](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme), [message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions), and [dead-letter queues](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-dead-letter-queues).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Comprehensive lifecycle management covers senders, processors, session receivers, dead-letter receivers, and temporary malformed-message clients in both synchronous and asynchronous flows.

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
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 7.0% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.7500** | **82.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-java-order-processor --config java-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)

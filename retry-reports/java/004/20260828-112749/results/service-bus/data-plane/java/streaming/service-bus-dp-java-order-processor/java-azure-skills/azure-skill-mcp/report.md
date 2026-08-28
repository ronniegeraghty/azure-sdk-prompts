# Evaluation Report: service-bus-dp-java-order-processor

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 812.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-order-processor` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 19/24 |
| Duration | 812.8s |
| Timestamp | 2026-08-28T03:27:49Z |
| Files Generated | 12 |
| Event Count | 19393 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 558.3s |
| Review | 253.3s |
| **Total** | **812.8s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=651204 out=14012 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. skill — 754ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 181787ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Service Bus SDK and managed identity, with sync/async senders and processors, sessions for per-customer ordering, batching, scheduling, dead-letter handling, and DLQ reprocessing.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. glob — 181782ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. glob — 181427ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 183699ms

**Input:**

```json
{"intent":"Get Azure SDK Java best practices for a local sample using Service Bus queues, sessions, managed identity, batch sends, scheduled messages, explicit settlement, dead-lettering, and async Reactor clients.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 183701ms

**Input:**

```json
{"intent":"Find current official Java Azure Service Bus SDK examples and API guidance for ServiceBusSenderClient/AsyncClient batching, scheduling messages, session receivers/processors, dead-lettering, and dead-letter subqueue access.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. web_search (via github-mcp-server) — 183695ms

**Input:**

```json
{"query":"site:learn.microsoft.com java azure service bus sessions ServiceBusSessionProcessorClient dead letter queue scheduled messages batch current SDK"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The latest versions of the requested Azure SDK packages on Maven Central as of August 2026 are:\n\n- **com.azure:azure-messaging-servicebus:** `7.18.0-beta.3`\n- **com.azure:azure-identity:** `1.18.5` (released on August 24, 2026)\n\nThese are confirmed by both Maven Central listings and the Azure SDK for Java GitHub release history【3:1†source】【3:8†source】.","annotations":[{"text":"","start_index":329,"end_index":341,"url_citation":{"title":"Maven Central: com.azure:azure-messaging-servicebus","url":"https://central.sonatype.com/artifact/com.azure/azure-messaging-servicebus"}},{"text":"","start_index":341,"end_index":353,"url_citation":{"title":"Releases: Azure/azure-sdk-for-java - GitHub","url":"https://github.com/Azure/azure-sdk-for-java/releases"}}]},"bing_searches":[{"text":"Maven Central latest versions of com.azure:azure-messaging-servicebus and com.azure:azure-identity as of August 2026","url":"https://www.bing.com/search?q=Maven+Central+latest+versions+of+com.azure%3Aazure-messaging-servicebus+and+com.azure%3Aazure-identity+as+of+August+2026"}],"annotations":null}
```

### ✅ 8. web_search (via github-mcp-server) — 183347ms

**Input:**

```json
{"query":"Maven Central latest com.azure azure-messaging-servicebus azure-identity versions August 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To address your query about handling sessions, ServiceBusSessionProcessorClient, dead-letter queue (DLQ), scheduled messages, and batch processing with the current Azure Service Bus Java SDK, here’s a summary based on the official Microsoft documentation:\n\n### 1. Creating and Configuring ServiceBusSessionProcessorClient\n\n- Use `ServiceBusClientBuilder.ServiceBusSessionProcessorClientBuilder` to build a session processor that listens to session-aware queues.\n- You can configure:\n  - **Session concurrency** (`maxConcurrentSessions`)\n  - **Message handlers** (`processMessage`, `processError`)\n  - **Manual or auto message completion** (`disableAutoComplete`)\n  - **How many concurrent calls per session** (`maxConcurrentCalls`)\n- Example creation:\n  ```java\n  ServiceBusSessionProcessorClient sessionProcessor = new ServiceBusClientBuilder()\n      .credential(fullyQualifiedNamespace, tokenCredential)\n      .sessionProcessor()\n      .queueName(sessionEnabledQueueName)\n      .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n      .disableAutoComplete()\n      .maxConcurrentSessions(2)\n      .processMessage(onMessage)\n      .processError(onError)\n      .buildProcessorClient();\n\n  sessionProcessor.start();\n  ```【3:0†source】\n\n### 2. Receiving from the Dead Letter Queue\n\nEach queue/subscription has a dead-letter subqueue (DLQ). In the Java SDK, use this as the subqueue:\n- Set the subqueue type to **dead-letter** when building the processor.\n- Example:\n  ```java\n  .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n  ```\n- Messages dead-letter automatically if they are explicitly dead-lettered or exceed the max delivery or TTL settings. The reason and error description are available on the message context【3:8†source】.\n\n### 3. Scheduled Messages\n\n- You can send messages scheduled for future delivery using the `scheduleMessage` method.\n- Example:\n  ```java\n  OffsetDateTime scheduledEnqueueTime = OffsetDateTim
... (truncated)
```

### ✅ 9. web_fetch — 5397ms

**Input:**

```json
{"max_length":8000,"url":"https://repo1.maven.org/maven2/com/azure/azure-messaging-servicebus/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceiverclient?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusReceiverClient Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceiverclient?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusReceiverClient
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>A <strong>synchronous</strong> receiver responsible for receiving <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceivedmessage'>ServiceBusReceivedMessage</a> from a queue or topic/subscription on Azure Service Bus.</p>

  <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication, which is appropriate for most scenarios, including local development and production environments. Additionally, we recommend using <a href="/azure/active-directory/managed-identities-azure-resources/">managed identity</a> for authentication in production environments. You can find more information on different ways of authenticating and their corresponding credential types in the <a href="/java/api/overview/azure/identity-readme">Azure Identity documentation&quot;</a>.</p>

  <p><strong>Sample: Create a receiver and receive messages</strong></p>

  <p>The following code sample demonstrates the creation and use of the synchronous client <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceiverclient'>ServiceBusReceiverClient</a> to receive messages from a Service Bus subscription. The receive operation returns when either 10 messages are received or 30 seconds has elapsed. By default, messages are received using <a class=no-loc href='com.azure.messaging.servicebus.models.servi
... (truncated)
```

### ✅ 10. web_fetch — 5038ms

**Input:**

```json
{"max_length":8000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 11. web_fetch — 4671ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebusreceiverclient?view=azure-java-stable"}
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

### ✅ 12. web_fetch — 4265ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverclient?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverclient?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: ServiceBusSessionReceiverClient Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverclient?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.messaging.servicebus.ServiceBusSessionReceiverClient
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>This <strong>synchronous</strong> session receiver client is used to acquire session locks from a queue or topic and create <a class=no-loc href='com.azure.messaging.servicebus.servicebusreceiverclient'>ServiceBusReceiverClient</a> instances that are tied to the locked sessions. Sessions can be used as a first in first out (FIFO) processing of messages. Queues and topics/subscriptions support Service Bus sessions, however, it must be <a href="/azure/service-bus-messaging/enable-message-sessions">enabled at the time of entity creation</a>.</p>

  <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication, which is appropriate for most scenarios, including local development and production environments. Additionally, we recommend using <a href="/azure/active-directory/managed-identities-azure-resources/">managed identity</a> for authentication in production environments. You can find more information on different ways of authenticating and their corresponding credential types in the <a href="/java/api/overview/azure/identity-readme">Azure Identity documentation&quot;</a>.</p>

  <p><strong>Sample: Receive messages from a specific session</strong></p>

  <p>Use <a class=no-loc href='com.azure.messaging.servicebus.servicebussessionreceiverclient#com-azure-messa
... (truncated)
```

### ✅ 13. web_fetch — 3884ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussessionreceiverasyncclient?view=azure-java-stable"}
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

### ✅ 14. web_fetch — 3513ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.messaging.servicebus.servicebussenderasyncclient?view=azure-java-stable"}
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

### ✅ 15. apply_patch — 711ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eservice-bus-order-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eService Bus Order Processor\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+        \u003cazure.servicebus.version\u003e7.17.20\u003c/azure.servicebus.version\u003e\n+        \u003cjackson.version\u003e2.20.0\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.servicebus.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.17\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: README.md\n+# Azure Service Bus order processor\n+\n+Java 17 sample with synchronous and Reactor-based asynchronous send/receive\n+implementations. The queue **must have sessions enabled** because each message\n+uses the customer name as its session ID. Processors consume one session at a\n+time, preserving per-customer FIFO order without interleaving customers.\n+\n+## Configuration\n+\n+The managed identity needs `Azure Service Bus Data Sender` and\n+`Azure Service Bus Data Receiver` on the queue or namespace.\n+\n+```powershell\n+$env:SERVICE_BUS_NAMESPACE = \"your-namespace.servicebus.windows.net\"\n+$env:SERVICE_BUS_QUEUE = \"orders\"\n+# Optional user-assigned managed identity:\n+$env:AZURE_CLIENT_ID = \"00000000-0000-0000-0000-000000000000\"\n+# Optional; defaults to 1000.00:\n+$env:HIGH_PRIORITY_THRESHOLD = \"1000.00\"\n+\n+mvn test\n+mvn exec:java\n+```\n+\n+The demo deliberately sends one malformed message in each cycle to demonstrate\n+explicit dead-lettering and inspection. High-value orders are marked with the\n+`priority=high` application property and scheduled 30 seconds in the future.\n+The demo waits for that delay before its second processing pass.\n*** Add File: src/main/java/com/example/orders/OrderStatus.java\n+package com.example.orders;\n+\n+import com.fasterxml.jackson.annotation.JsonCreator;\n+import com.fasterxml.jackson.annotation.JsonValue;\n+\n+import java.util.Locale;\n+\n+public enum OrderStatus {\n+    PENDING,\n+    PROCESSING,\n+    COMPLETED,\n+    FAILED;\n+\n+    @JsonCreator\n+    public static OrderStatus fromJson(String value) {\n+        return valueOf(value.toUpperCase(Locale.ROOT));\n+    }\n+\n+    @JsonValue\n+    public String toJson() {\n+        return name().toLowerCase(Locale.ROOT);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Order.java\n+package com.example.orders;\n+\n+import java.math.BigDecimal;\n+import java.util.Objects;\n+\n+public record Order(\n+        String orderId,\n+        String customerName,\n+        String product,\n+        int quantity,\n+        BigDecimal totalPrice,\n+        OrderStatus status) {\n+\n+    public Order {\n+        if (orderId == null || orderId.isBlank()) {\n+            throw new IllegalArgumentException(\"orderId is required\");\n+        }\n+        if (customerName == null || customerName.isBlank()) {\n+            throw new IllegalArgumentException(\"customerName is required\");\n+        }\n+        if (product == null || product.isBlank()) {\n+            throw new IllegalArgumentException(\"product is required\");\n+        }\n+        if (quantity \u003c= 0) {\n+            throw new IllegalArgumentException(\"quantity must be positive\");\n+        }\n+        Objects.requireNonNull(totalPrice, \"totalPrice is required\");\n+        if (totalPrice.signum() \u003c 0) {\n+            throw new IllegalArgumentException(\"totalPrice cannot be negative\");\n+        }\n+        Objects.requireNonNull(status, \"status is required\");\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderMessageMapper.java\n+package com.example.orders;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+final class OrderMessageMapper {\n+    private static final ObjectMapper JSON = new ObjectMapper();\n+\n+    private OrderMessageMapper() {\n+    }\n+\n+    static ServiceBusMessage toMessage(Order order, boolean highPriority) {\n+        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(toJson(order)))\n+                .setContentType(\"application/json\")\n+                .setMessageId(order.orderId())\n+                .setCorrelationId(order.orderId())\n+                .setSessionId(order.customerName());\n+        message.getApplicationProperties().put(\"priority\", highPriority ? \"high\" : \"normal\");\n+        return message;\n+    }\n+\n+    static Order fromMessage(ServiceBusReceivedMessage message) throws JsonProcessingException {\n+        return JSON.readValue(message.getBody().toString(), Order.class);\n+    }\n+\n+    static String toJson(Order order) {\n+        try {\n+            return JSON.writeValueAsString(order);\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Could not serialize order \" + order.orderId(), exception);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/DeadLetteredOrder.java\n+package com.example.orders;\n+\n+public record DeadLetteredOrder(\n+        String correlationId,\n+        String sessionId,\n+        String body,\n+        String reason,\n+        String errorDescription) {\n+}\n*** Add File: src/main/java/com/example/orders/OrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+import java.util.ArrayList;\n+import java.util.LinkedHashMap;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class OrderSender implements AutoCloseable {\n+    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;\n+\n+    private final ServiceBusSenderClient sender;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public OrderSender(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildClient();\n+        this.highPriorityThreshold = highPriorityThreshold;\n+    }\n+\n+    public void send(Order order) {\n+        ServiceBusMessage message = messageFor(order);\n+        if (isHighPriority(order)) {\n+            sender.scheduleMessage(message, OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));\n+        } else {\n+            sender.sendMessage(message);\n+        }\n+    }\n+\n+    public void sendBatch(List\u003cOrder\u003e orders) {\n+        Map\u003cString, List\u003cOrder\u003e\u003e immediateByCustomer = new LinkedHashMap\u003c\u003e();\n+        for (Order order : orders) {\n+            if (isHighPriority(order)) {\n+                sender.scheduleMessage(\n+                        messageFor(order),\n+                        OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));\n+            } else {\n+                immediateByCustomer\n+                        .computeIfAbsent(order.customerName(), ignored -\u003e new ArrayList\u003c\u003e())\n+                        .add(order);\n+            }\n+        }\n+        immediateByCustomer.values().forEach(this::sendImmediateBatches);\n+    }\n+\n+    public void sendMalformedForDemo(String correlationId, String customerName) {\n+        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(\"{not-valid-json\"))\n+                .setContentType(\"application/json\")\n+                .setMessageId(correlationId)\n+                .setCorrelationId(correlationId)\n+                .setSessionId(customerName);\n+        sender.sendMessage(message);\n+    }\n+\n+    private void sendImmediateBatches(List\u003cOrder\u003e orders) {\n+        ServiceBusMessageBatch batch = sender.createMessageBatch();\n+        for (Order order : orders) {\n+            ServiceBusMessage message = messageFor(order);\n+            if (!batch.tryAddMessage(message)) {\n+                if (batch.getCount() == 0) {\n+                    throw new IllegalArgumentException(\"Order is too large for an empty Service Bus batch: \"\n+                            + order.orderId());\n+                }\n+                sender.sendMessages(batch);\n+                batch = sender.createMessageBatch();\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalArgumentException(\"Order is too large for an empty Service Bus batch: \"\n+                            + order.orderId());\n+                }\n+            }\n+        }\n+        if (batch.getCount() \u003e 0) {\n+            sender.sendMessages(batch);\n+        }\n+    }\n+\n+    private ServiceBusMessage messageFor(Order order) {\n+        return OrderMessageMapper.toMessage(order, isHighPriority(order));\n+    }\n+\n+    private boolean isHighPriority(Order order) {\n+        return order.totalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+import java.util.ArrayList;\n+import java.util.LinkedHashMap;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class AsyncOrderSender implements AutoCloseable {\n+    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;\n+\n+    private final ServiceBusSenderAsyncClient sender;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public AsyncOrderSender(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildAsyncClient();\n+        this.highPriorityThreshold = highPriorityThreshold;\n+    }\n+\n+    public Mono\u003cVoid\u003e send(Order order) {\n+        ServiceBusMessage message = messageFor(order);\n+        if (isHighPriority(order)) {\n+            return sender.scheduleMessage(\n+                            message,\n+                            OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS))\n+                    .then();\n+        }\n+        return sender.sendMessage(message);\n+    }\n+\n+    public Mono\u003cVoid\u003e sendBatch(List\u003cOrder\u003e orders) {\n+        Map\u003cString, List\u003cOrder\u003e\u003e immediateByCustomer = new LinkedHashMap\u003c\u003e();\n+        List\u003cOrder\u003e scheduled = new ArrayList\u003c\u003e();\n+        for (Order order : orders) {\n+            if (isHighPriority(order)) {\n+                scheduled.add(order);\n+            } else {\n+                immediateByCustomer\n+                        .computeIfAbsent(order.customerName(), ignored -\u003e new ArrayList\u003c\u003e())\n+                        .add(order);\n+            }\n+        }\n+\n+        Mono\u003cVoid\u003e immediate = Flux.fromIterable(immediateByCustomer.values())\n+                .concatMap(this::sendImmediateBatches)\n+                .then();\n+        Mono\u003cVoid\u003e delayed = Flux.fromIterable(scheduled)\n+                .concatMap(this::send)\n+                .then();\n+        return immediate.then(delayed);\n+    }\n+\n+    public Mono\u003cVoid\u003e sendMalformedForDemo(String correlationId, String customerName) {\n+        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(\"{not-valid-json\"))\n+                .setContentType(\"application/json\")\n+                .setMessageId(correlationId)\n+                .setCorrelationId(correlationId)\n+                .setSessionId(customerName);\n+        return sender.sendMessage(message);\n+    }\n+\n+    private Mono\u003cVoid\u003e sendImmediateBatches(List\u003cOrder\u003e orders) {\n+        return sender.createMessageBatch().flatMap(batch -\u003e addAndSend(orders, 0, batch));\n+    }\n+\n+    private Mono\u003cVoid\u003e addAndSend(List\u003cOrder\u003e orders, int index, ServiceBusMessageBatch batch) {\n+        int next = index;\n+        while (next \u003c orders.size()) {\n+            Order order = orders.get(next);\n+            if (!batch.tryAddMessage(messageFor(order))) {\n+                if (batch.getCount() == 0) {\n+                    return Mono.error(new IllegalArgumentException(\n+                            \"Order is too large for an empty Service Bus batch: \" + order.orderId()));\n+                }\n+                int resumeAt = next;\n+                return sender.sendMessages(batch)\n+                        .then(sender.createMessageBatch())\n+                        .flatMap(newBatch -\u003e addAndSend(orders, resumeAt, newBatch));\n+            }\n+            next++;\n+        }\n+        return batch.getCount() == 0 ? Mono.empty() : sender.sendMessages(batch);\n+    }\n+\n+    private ServiceBusMessage messageFor(Order order) {\n+        return OrderMessageMapper.toMessage(order, isHighPriority(order));\n+    }\n+\n+    private boolean isHighPriority(Order order) {\n+        return order.totalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+\n+import java.time.Duration;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Set;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+\n+public final class OrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = Logger.getLogger(OrderProcessor.class.getName());\n+    private static final String DEAD_LETTER_REASON = \"ORDER_DESERIALIZATION_FAILED\";\n+\n+    private final ServiceBusSessionReceiverClient activeReceiver;\n+    private final ServiceBusSessionReceiverClient deadLetterReceiver;\n+\n+    public OrderProcessor(String fullyQualifiedNamespace, String queueName, TokenCredential credential) {\n+        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential);\n+        this.activeReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildClient();\n+        this.deadLetterReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildClient();\n+    }\n+\n+    public void processSessions(Set\u003cString\u003e customerNames, Duration waitPerReceive) {\n+        for (String customerName : customerNames) {\n+            try (ServiceBusReceiverClient receiver = activeReceiver.acceptSession(customerName)) {\n+                boolean received;\n+                do {\n+                    List\u003cServiceBusReceivedMessage\u003e messages =\n+                            receiver.receiveMessages(100, waitPerReceive).stream().toList();\n+                    received = !messages.isEmpty();\n+                    messages.forEach(message -\u003e process(receiver, message));\n+                } while (received);\n+            }\n+        }\n+    }\n+\n+    public List\u003cDeadLetteredOrder\u003e inspectDeadLetters(Set\u003cString\u003e customerNames, Duration waitPerReceive) {\n+        List\u003cDeadLetteredOrder\u003e deadLetters = new ArrayList\u003c\u003e();\n+        for (String customerName : customerNames) {\n+            try (ServiceBusReceiverClient receiver = deadLetterReceiver.acceptSession(customerName)) {\n+                for (ServiceBusReceivedMessage message : receiver.receiveMessages(100, waitPerReceive)) {\n+                    DeadLetteredOrder deadLetter = toDeadLetteredOrder(message);\n+                    deadLetters.add(deadLetter);\n+                    LOGGER.warning(() -\u003e \"Dead letter: \" + deadLetter);\n+                    receiver.abandon(message);\n+                }\n+            }\n+        }\n+        return deadLetters;\n+    }\n+\n+    private void process(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {\n+        try {\n+            Order order = OrderMessageMapper.fromMessage(message);\n+            LOGGER.info(() -\u003e \"Processing order \" + order.orderId()\n+                    + \" for \" + order.customerName() + \": \" + order);\n+            LOGGER.info(() -\u003e \"Completed order \" + order.orderId());\n+            receiver.complete(message);\n+        } catch (RuntimeException exception) {\n+            LOGGER.log(Level.WARNING,\n+                    \"Dead-lettering order \" + message.getCorrelationId() + \" because it cannot be deserialized\",\n+                    exception);\n+            receiver.deadLetter(message, new DeadLetterOptions()\n+                    .setDeadLetterReason(DEAD_LETTER_REASON)\n+                    .setDeadLetterErrorDescription(safeDescription(exception)));\n+        }\n+    }\n+\n+    private static DeadLetteredOrder toDeadLetteredOrder(ServiceBusReceivedMessage message) {\n+        return new DeadLetteredOrder(\n+                message.getCorrelationId(),\n+                message.getSessionId(),\n+                message.getBody().toString(),\n+                message.getDeadLetterReason(),\n+                message.getDeadLetterErrorDescription());\n+    }\n+\n+    private static String safeDescription(RuntimeException exception) {\n+        String message = exception.getMessage();\n+        return message == null ? exception.getClass().getSimpleName() : message.substring(0, Math.min(4096, message.length()));\n+    }\n+\n+    @Override\n+    public void close() {\n+        activeReceiver.close();\n+        deadLetterReceiver.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Set;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+\n+public final class AsyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = Logger.getLogger(AsyncOrderProcessor.class.getName());\n+    private static final String DEAD_LETTER_REASON = \"ORDER_DESERIALIZATION_FAILED\";\n+\n+    private final ServiceBusSessionReceiverAsyncClient activeReceiver;\n+    private final ServiceBusSessionReceiverAsyncClient deadLetterReceiver;\n+\n+    public AsyncOrderProcessor(String fullyQualifiedNamespace, String queueName, TokenCredential credential) {\n+        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential);\n+        this.activeReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+        this.deadLetterReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+    }\n+\n+    public Mono\u003cVoid\u003e processSessions(Set\u003cString\u003e customerNames, Duration receiveWindow) {\n+        return Flux.fromIterable(customerNames)\n+                .concatMap(customer -\u003e Flux.usingWhen(\n+                        activeReceiver.acceptSession(customer),\n+                        receiver -\u003e receiver.receiveMessages()\n+                                .take(receiveWindow)\n+                                .concatMap(message -\u003e process(receiver, message)),\n+                        receiver -\u003e Mono.fromRunnable(receiver::close)))\n+                .then();\n+    }\n+\n+    public Mono\u003cList\u003cDeadLetteredOrder\u003e\u003e inspectDeadLetters(\n+            Set\u003cString\u003e customerNames,\n+            Duration receiveWindow) {\n+        return Flux.fromIterable(customerNames)\n+                .concatMap(customer -\u003e Flux.usingWhen(\n+                        deadLetterReceiver.acceptSession(customer),\n+                        receiver -\u003e receiver.receiveMessages()\n+                                .take(receiveWindow)\n+                                .concatMap(message -\u003e {\n+                                    DeadLetteredOrder deadLetter = toDeadLetteredOrder(message);\n+                                    LOGGER.warning(() -\u003e \"Dead letter: \" + deadLetter);\n+                                    return receiver.abandon(message).thenReturn(deadLetter);\n+                                }),\n+                        receiver -\u003e Mono.fromRunnable(receiver::close)))\n+                .collectList();\n+    }\n+\n+    private Mono\u003cVoid\u003e process(ServiceBusReceiverAsyncClient receiver, ServiceBusReceivedMessage message) {\n+        try {\n+            Order order = OrderMessageMapper.fromMessage(message);\n+            LOGGER.info(() -\u003e \"Processing order \" + order.orderId()\n+                    + \" for \" + order.customerName() + \": \" + order);\n+            return receiver.complete(message)\n+                    .doOnSuccess(ignored -\u003e LOGGER.info(() -\u003e \"Completed order \" + order.orderId()));\n+        } catch (RuntimeException exception) {\n+            LOGGER.log(Level.WARNING,\n+                    \"Dead-lettering order \" + message.getCorrelationId() + \" because it cannot be deserialized\",\n+                    exception);\n+            return receiver.deadLetter(message, new DeadLetterOptions()\n+                    .setDeadLetterReason(DEAD_LETTER_REASON)\n+                    .setDeadLetterErrorDescription(safeDescription(exception)));\n+        }\n+    }\n+\n+    private static DeadLetteredOrder toDeadLetteredOrder(ServiceBusReceivedMessage message) {\n+        return new DeadLetteredOrder(\n+                message.getCorrelationId(),\n+                message.getSessionId(),\n+                message.getBody().toString(),\n+                message.getDeadLetterReason(),\n+                message.getDeadLetterErrorDescription());\n+    }\n+\n+    private static String safeDescription(RuntimeException exception) {\n+        String message = exception.getMessage();\n+        return message == null ? exception.getClass().getSimpleName() : message.substring(0, Math.min(4096, message.length()));\n+    }\n+\n+    @Override\n+    public void close() {\n+        activeReceiver.close();\n+        deadLetterReceiver.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Main.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Set;\n+\n+public final class Main {\n+    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(3);\n+    private static final Duration ASYNC_RECEIVE_WINDOW = Duration.ofSeconds(5);\n+    private static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(31);\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String namespace = requiredEnvironmentVariable(\"SERVICE_BUS_NAMESPACE\");\n+        String queueName = requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE\");\n+        BigDecimal threshold = new BigDecimal(\n+                System.getenv().getOrDefault(\"HIGH_PRIORITY_THRESHOLD\", \"1000.00\"));\n+        TokenCredential credential = managedIdentityCredential();\n+\n+        runSynchronousDemo(namespace, queueName, credential, threshold);\n+        runAsynchronousDemo(namespace, queueName, credential, threshold);\n+    }\n+\n+    private static void runSynchronousDemo(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal threshold) throws InterruptedException {\n+        List\u003cOrder\u003e orders = demoOrders(\"sync\");\n+        Set\u003cString\u003e customers = Set.of(\"Ada\", \"Grace\", \"Demo\");\n+\n+        try (OrderSender sender = new OrderSender(namespace, queueName, credential, threshold);\n+             OrderProcessor processor = new OrderProcessor(namespace, queueName, credential)) {\n+            sender.send(orders.get(0));\n+            sender.sendBatch(orders.subList(1, orders.size()));\n+            sender.sendMalformedForDemo(\"sync-malformed\", \"Demo\");\n+\n+            processor.processSessions(customers, RECEIVE_WAIT);\n+            processor.inspectDeadLetters(customers, RECEIVE_WAIT);\n+\n+            Thread.sleep(FRAUD_REVIEW_DELAY.toMillis());\n+            processor.processSessions(customers, RECEIVE_WAIT);\n+        }\n+    }\n+\n+    private static void runAsynchronousDemo(\n+            String namespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal threshold) {\n+        List\u003cOrder\u003e orders = demoOrders(\"async\");\n+        Set\u003cString\u003e customers = Set.of(\"Ada\", \"Grace\", \"Demo\");\n+\n+        try (AsyncOrderSender sender = new AsyncOrderSender(namespace, queueName, credential, threshold);\n+             AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential)) {\n+            sender.send(orders.get(0))\n+                    .then(sender.sendBatch(orders.subList(1, orders.size())))\n+                    .then(sender.sendMalformedForDemo(\"async-malformed\", \"Demo\"))\n+                    .then(processor.processSessions(customers, ASYNC_RECEIVE_WINDOW))\n+                    .then(processor.inspectDeadLetters(customers, ASYNC_RECEIVE_WINDOW))\n+                    .doOnNext(deadLetters -\u003e\n+                            System.out.println(\"Async dead-letter count: \" + deadLetters.size()))\n+                    .then(MonoDelay.delay(FRAUD_REVIEW_DELAY))\n+                    .then(processor.processSessions(customers, ASYNC_RECEIVE_WINDOW))\n+                    .block();\n+        }\n+    }\n+\n+    private static List\u003cOrder\u003e demoOrders(String prefix) {\n+        return List.of(\n+                new Order(prefix + \"-001\", \"Ada\", \"Keyboard\", 1,\n+                        new BigDecimal(\"129.99\"), OrderStatus.PENDING),\n+                new Order(prefix + \"-002\", \"Ada\", \"Monitor\", 2,\n+                        new BigDecimal(\"1499.98\"), OrderStatus.PENDING),\n+                new Order(prefix + \"-003\", \"Grace\", \"Dock\", 1,\n+                        new BigDecimal(\"249.00\"), OrderStatus.PENDING));\n+    }\n+\n+    private static TokenCredential managedIdentityCredential() {\n+        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();\n+        String clientId = System.getenv(\"AZURE_CLIENT_ID\");\n+        if (clientId != null \u0026\u0026 !clientId.isBlank()) {\n+            builder.clientId(clientId);\n+        }\n+        return builder.build();\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(name + \" must be set\");\n+        }\n+        return value;\n+    }\n+\n+    private static final class MonoDelay {\n+        private MonoDelay() {\n+        }\n+\n+        static reactor.core.publisher.Mono\u003cLong\u003e delay(Duration duration) {\n+            return reactor.core.publisher.Mono.delay(duration);\n+        }\n+    }\n+}\n*** Add File: src/test/java/com/example/orders/OrderJsonTest.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import org.junit.jupiter.api.Test;\n+\n+import java.math.BigDecimal;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class OrderJsonTest {\n+    @Test\n+    void serializesStatusAsLowercaseJson() {\n+        Order order = new Order(\n+                \"order-1\",\n+                \"Ada\",\n+                \"Keyboard\",\n+                2,\n+                new BigDecimal(\"199.98\"),\n+                OrderStatus.PROCESSING);\n+\n+        ServiceBusMessage message = OrderMessageMapper.toMessage(order, false);\n+\n+        assertEquals(\"order-1\", message.getCorrelationId());\n+        assertEquals(\"Ada\", message.getSessionId());\n+        assertEquals(\"normal\", message.getApplicationProperties().get(\"priority\"));\n+        assertEquals(\n+                \"{\\\"orderId\\\":\\\"order-1\\\",\\\"customerName\\\":\\\"Ada\\\",\\\"product\\\":\\\"Keyboard\\\",\"\n+                        + \"\\\"quantity\\\":2,\\\"totalPrice\\\":199.98,\\\"status\\\":\\\"processing\\\"}\",\n+                message.getBody().toString());\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/pom.xml
@@ -1,0 +1,73 @@
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
+        <azure.identity.version>1.18.5</azure.identity.version>
+        <azure.servicebus.version>7.17.20</azure.servicebus.version>
+        <jackson.version>2.20.0</jackson.version>
+        <junit.version>5.13.4</junit.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-servicebus</artifactId>
+            <version>${azure.servicebus.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.fasterxml.jackson.core</groupId>
+            <artifactId>jackson-databind</artifactId>
+            <version>${jackson.version}</version>
+ 
... (truncated)
```

### ✅ 16. powershell — 5619ms

**Input:**

```json
{"command":"mvn test --quiet","description":"Compile project and run Maven tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[ERROR] COMPILATION ERROR : 
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/AsyncOrderProcessor.java:[73,57] unreported exception com.fasterxml.jackson.core.JsonProcessingException; must be caught or declared to be thrown
[ERROR] Failed to execute goal org.apache.maven.plugins:maven-compiler-plugin:3.14.0:compile (default-compile) on project service-bus-order-processor: Compilation failure
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/AsyncOrderProcessor.java:[73,57] unreported exception com.fasterxml.jackson.core.JsonProcessingException; must be caught or declared to be thrown
[ERROR] 
[ERROR] -> [Help 1]
[ERROR] 
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR] 
[ERROR] For more information about the errors and possible solutions, please read the following articles:
[ERROR] [Help 1] http://cwiki.apache.org/confluence/display/MAVEN/MojoFailureException
<shellId: 0 completed with exit code 1>
```

### ✅ 17. apply_patch — 681ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/OrderProcessor.java\n@@\n import java.util.ArrayList;\n import java.util.List;\n+import java.util.Optional;\n import java.util.Set;\n+import java.util.function.Function;\n import java.util.logging.Level;\n import java.util.logging.Logger;\n@@\n     public List\u003cDeadLetteredOrder\u003e inspectDeadLetters(Set\u003cString\u003e customerNames, Duration waitPerReceive) {\n@@\n         return deadLetters;\n     }\n \n+    public int reprocessDeadLetters(\n+            Set\u003cString\u003e customerNames,\n+            Duration waitPerReceive,\n+            Function\u003cDeadLetteredOrder, Optional\u003cOrder\u003e\u003e recovery,\n+            OrderSender sender) {\n+        int reprocessed = 0;\n+        for (String customerName : customerNames) {\n+            try (ServiceBusReceiverClient receiver = deadLetterReceiver.acceptSession(customerName)) {\n+                for (ServiceBusReceivedMessage message : receiver.receiveMessages(100, waitPerReceive)) {\n+                    Optional\u003cOrder\u003e recovered = recovery.apply(toDeadLetteredOrder(message));\n+                    if (recovered.isPresent()) {\n+                        sender.send(recovered.get());\n+                        receiver.complete(message);\n+                        reprocessed++;\n+                    } else {\n+                        receiver.abandon(message);\n+                    }\n+                }\n+            }\n+        }\n+        return reprocessed;\n+    }\n+\n     private void process(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {\n         try {\n             Order order = OrderMessageMapper.fromMessage(message);\n@@\n-        } catch (RuntimeException exception) {\n+        } catch (Exception exception) {\n             LOGGER.log(Level.WARNING,\n                     \"Dead-lettering order \" + message.getCorrelationId() + \" because it cannot be deserialized\",\n                     exception);\n@@\n-    private static String safeDescription(RuntimeException exception) {\n+    private static String safeDescription(Exception exception) {\n*** Update File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n@@\n import java.time.Duration;\n import java.util.List;\n+import java.util.Optional;\n import java.util.Set;\n+import java.util.function.Function;\n import java.util.logging.Level;\n import java.util.logging.Logger;\n@@\n     public Mono\u003cList\u003cDeadLetteredOrder\u003e\u003e inspectDeadLetters(\n@@\n                 .collectList();\n     }\n \n+    public Mono\u003cLong\u003e reprocessDeadLetters(\n+            Set\u003cString\u003e customerNames,\n+            Duration receiveWindow,\n+            Function\u003cDeadLetteredOrder, Optional\u003cOrder\u003e\u003e recovery,\n+            AsyncOrderSender sender) {\n+        return Flux.fromIterable(customerNames)\n+                .concatMap(customer -\u003e Flux.usingWhen(\n+                        deadLetterReceiver.acceptSession(customer),\n+                        receiver -\u003e receiver.receiveMessages()\n+                                .take(receiveWindow)\n+                                .concatMap(message -\u003e {\n+                                    Optional\u003cOrder\u003e recovered = recovery.apply(toDeadLetteredOrder(message));\n+                                    if (recovered.isEmpty()) {\n+                                        return receiver.abandon(message).then(Mono.just(false));\n+                                    }\n+                                    return sender.send(recovered.get())\n+                                            .then(receiver.complete(message))\n+                                            .thenReturn(true);\n+                                }),\n+                        receiver -\u003e Mono.fromRunnable(receiver::close)))\n+                .filter(Boolean::booleanValue)\n+                .count();\n+    }\n+\n     private Mono\u003cVoid\u003e process(ServiceBusReceiverAsyncClient receiver, ServiceBusReceivedMessage message) {\n@@\n-        } catch (RuntimeException exception) {\n+        } catch (Exception exception) {\n@@\n-    private static String safeDescription(RuntimeException exception) {\n+    private static String safeDescription(Exception exception) {\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n import java.time.Duration;\n import java.util.List;\n+import java.util.Optional;\n import java.util.Set;\n@@\n             processor.processSessions(customers, RECEIVE_WAIT);\n             processor.inspectDeadLetters(customers, RECEIVE_WAIT);\n+            int reprocessed = processor.reprocessDeadLetters(\n+                    customers,\n+                    RECEIVE_WAIT,\n+                    deadLetter -\u003e recoverDemoOrder(deadLetter, \"sync\"),\n+                    sender);\n+            System.out.println(\"Sync reprocessed dead-letter count: \" + reprocessed);\n \n             Thread.sleep(FRAUD_REVIEW_DELAY.toMillis());\n             processor.processSessions(customers, RECEIVE_WAIT);\n@@\n                     .doOnNext(deadLetters -\u003e\n                             System.out.println(\"Async dead-letter count: \" + deadLetters.size()))\n-                    .then(MonoDelay.delay(FRAUD_REVIEW_DELAY))\n+                    .then(processor.reprocessDeadLetters(\n+                            customers,\n+                            ASYNC_RECEIVE_WINDOW,\n+                            deadLetter -\u003e recoverDemoOrder(deadLetter, \"async\"),\n+                            sender))\n+                    .doOnNext(count -\u003e\n+                            System.out.println(\"Async reprocessed dead-letter count: \" + count))\n+                    .then(reactor.core.publisher.Mono.delay(FRAUD_REVIEW_DELAY))\n                     .then(processor.processSessions(customers, ASYNC_RECEIVE_WINDOW))\n                     .block();\n         }\n@@\n     private static TokenCredential managedIdentityCredential() {\n@@\n         return builder.build();\n     }\n \n+    private static Optional\u003cOrder\u003e recoverDemoOrder(DeadLetteredOrder deadLetter, String prefix) {\n+        if (!(prefix + \"-malformed\").equals(deadLetter.correlationId())) {\n+            return Optional.empty();\n+        }\n+        return Optional.of(new Order(\n+                prefix + \"-recovered\",\n+                deadLetter.sessionId(),\n+                \"Recovered demo order\",\n+                1,\n+                new BigDecimal(\"1.00\"),\n+                OrderStatus.PENDING));\n+    }\n+\n@@\n-    private static final class MonoDelay {\n-        private MonoDelay() {\n-        }\n-\n-        static reactor.core.publisher.Mono\u003cLong\u003e delay(Duration duration) {\n-            return reactor.core.publisher.Mono.delay(duration);\n-        }\n-    }\n }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java
@@ -12,7 +12,9 @@
 import java.time.Duration;
 import java.util.ArrayList;
 import java.util.List;
+import java.util.Optional;
 import java.util.Set;
+import java.util.function.Function;
 import java.util.logging.Level;
 import java.util.logging.Logger;
 
@@ -68,6 +70,29 @@
         return deadLetters;
     }
 
+    public int reprocessDeadLetters(
+            Set<String> customerNames,
+            Duration waitPerReceive,
+            Function<DeadLetteredOrder, Optional<Order>> recovery,
+            OrderSender sender) {
+        int reprocessed = 0;
+        for (String customerName : customerNames) {
+            try (ServiceBusReceiverClient receiver = deadLetterReceiver.acceptSession(customerName)) {
+                for (ServiceBusReceivedMessage message : receiver.receiveMessages(100, waitPerReceive)) {
+                    Optional<Order> recovered = recovery.apply(toDeadLetteredOrder(message));
+                    if (recovered.isPresent()) {
+                        sender.send(recovered.get());
+                        receiver.complete(message);
+                        reprocessed++;
+                    } else {
+                        receiver.abandon(message);
+                    }
+                }
... (truncated)
```

### ✅ 18. apply_patch — 689ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/OrderProcessor.java\n@@\n import com.azure.messaging.servicebus.models.DeadLetterOptions;\n import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n import com.azure.messaging.servicebus.models.SubQueue;\n+import com.fasterxml.jackson.core.JsonProcessingException;\n@@\n-        } catch (Exception exception) {\n+        } catch (JsonProcessingException exception) {\n@@\n-    private static String safeDescription(Exception exception) {\n+    private static String safeDescription(JsonProcessingException exception) {\n*** Update File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n@@\n import com.azure.messaging.servicebus.models.DeadLetterOptions;\n import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n import com.azure.messaging.servicebus.models.SubQueue;\n+import com.fasterxml.jackson.core.JsonProcessingException;\n@@\n-        } catch (Exception exception) {\n+        } catch (JsonProcessingException exception) {\n@@\n-    private static String safeDescription(Exception exception) {\n+    private static String safeDescription(JsonProcessingException exception) {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/OrderProcessor.java
@@ -8,6 +8,7 @@
 import com.azure.messaging.servicebus.models.DeadLetterOptions;
 import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
 import com.azure.messaging.servicebus.models.SubQueue;
+import com.fasterxml.jackson.core.JsonProcessingException;
 
 import java.time.Duration;
 import java.util.ArrayList;
@@ -100,7 +101,7 @@
                     + " for " + order.customerName() + ": " + order);
             LOGGER.info(() -> "Completed order " + order.orderId());
             receiver.complete(message);
-        } catch (Exception exception) {
+        } catch (JsonProcessingException exception) {
             LOGGER.log(Level.WARNING,
                     "Dead-lettering order " + message.getCorrelationId() + " because it cannot be deserialized",
                     exception);
@@ -119,7 +120,7 @@
                 message.getDeadLetterErrorDescription());
     }
 
-    private static String safeDescription(Exception exception) {
+    private static String safeDescription(JsonProcessingException exception) {
         String message = exception.getMessage();
         return message == null ? exception.getClass().getSimpleName() : message.substring(0, Math.min(4096, message.length()));
     }


diff --git 
... (truncated)
```

### ✅ 19. powershell — 6260ms

**Input:**

```json
{"command":"mvn test --quiet","description":"Compile project and run Maven tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 1 completed with exit code 0>
```

### ✅ 20. apply_patch — 629ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n         List\u003cOrder\u003e orders = demoOrders(\"sync\");\n-        Set\u003cString\u003e customers = Set.of(\"Ada\", \"Grace\", \"Demo\");\n+        Set\u003cString\u003e initialCustomers = Set.of(\"Ada\", \"Grace\", \"Demo\");\n+        Set\u003cString\u003e failedCustomers = Set.of(\"Demo\");\n+        Set\u003cString\u003e delayedCustomers = Set.of(\"Ada\", \"Demo\");\n@@\n-            processor.processSessions(customers, RECEIVE_WAIT);\n-            processor.inspectDeadLetters(customers, RECEIVE_WAIT);\n+            processor.processSessions(initialCustomers, RECEIVE_WAIT);\n+            processor.inspectDeadLetters(failedCustomers, RECEIVE_WAIT);\n             int reprocessed = processor.reprocessDeadLetters(\n-                    customers,\n+                    failedCustomers,\n                     RECEIVE_WAIT,\n                     deadLetter -\u003e recoverDemoOrder(deadLetter, \"sync\"),\n                     sender);\n@@\n             System.out.println(\"Sync reprocessed dead-letter count: \" + reprocessed);\n \n             Thread.sleep(FRAUD_REVIEW_DELAY.toMillis());\n-            processor.processSessions(customers, RECEIVE_WAIT);\n+            processor.processSessions(delayedCustomers, RECEIVE_WAIT);\n@@\n         List\u003cOrder\u003e orders = demoOrders(\"async\");\n-        Set\u003cString\u003e customers = Set.of(\"Ada\", \"Grace\", \"Demo\");\n+        Set\u003cString\u003e initialCustomers = Set.of(\"Ada\", \"Grace\", \"Demo\");\n+        Set\u003cString\u003e failedCustomers = Set.of(\"Demo\");\n+        Set\u003cString\u003e delayedCustomers = Set.of(\"Ada\", \"Demo\");\n@@\n                     .then(sender.sendBatch(orders.subList(1, orders.size())))\n                     .then(sender.sendMalformedForDemo(\"async-malformed\", \"Demo\"))\n-                    .then(processor.processSessions(customers, ASYNC_RECEIVE_WINDOW))\n-                    .then(processor.inspectDeadLetters(customers, ASYNC_RECEIVE_WINDOW))\n+                    .then(processor.processSessions(initialCustomers, ASYNC_RECEIVE_WINDOW))\n+                    .then(processor.inspectDeadLetters(failedCustomers, ASYNC_RECEIVE_WINDOW))\n@@\n                     .then(processor.reprocessDeadLetters(\n-                            customers,\n+                            failedCustomers,\n                             ASYNC_RECEIVE_WINDOW,\n                             deadLetter -\u003e recoverDemoOrder(deadLetter, \"async\"),\n                             sender))\n@@\n                     .doOnNext(count -\u003e\n                             System.out.println(\"Async reprocessed dead-letter count: \" + count))\n                     .then(reactor.core.publisher.Mono.delay(FRAUD_REVIEW_DELAY))\n-                    .then(processor.processSessions(customers, ASYNC_RECEIVE_WINDOW))\n+                    .then(processor.processSessions(delayedCustomers, ASYNC_RECEIVE_WINDOW))\n                     .block();\n*** Update File: src/test/java/com/example/orders/OrderJsonTest.java\n@@\n import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n import org.junit.jupiter.api.Test;\n@@\n     void serializesStatusAsLowercaseJson() {\n@@\n                 message.getBody().toString());\n     }\n+\n+    @Test\n+    void roundTripsThroughJson() throws Exception {\n+        Order expected = new Order(\n+                \"order-2\",\n+                \"Grace\",\n+                \"Monitor\",\n+                1,\n+                new BigDecimal(\"749.99\"),\n+                OrderStatus.COMPLETED);\n+        ObjectMapper mapper = new ObjectMapper();\n+\n+        Order actual = mapper.readValue(mapper.writeValueAsString(expected), Order.class);\n+\n+        assertEquals(expected, actual);\n+    }\n }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/Main.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/Main.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/Main.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-2138395605/src/main/java/com/example/orders/Main.java
@@ -34,7 +34,9 @@
             TokenCredential credential,
             BigDecimal threshold) throws InterruptedException {
         List<Order> orders = demoOrders("sync");
-        Set<String> customers = Set.of("Ada", "Grace", "Demo");
+        Set<String> initialCustomers = Set.of("Ada", "Grace", "Demo");
+        Set<String> failedCustomers = Set.of("Demo");
+        Set<String> delayedCustomers = Set.of("Ada", "Demo");
 
         try (OrderSender sender = new OrderSender(namespace, queueName, credential, threshold);
              OrderProcessor processor = new OrderProcessor(namespace, queueName, credential)) {
@@ -42,10 +44,10 @@
             sender.sendBatch(orders.subList(1, orders.size()));
             sender.sendMalformedForDemo("sync-malformed", "Demo");
 
-            processor.processSessions(customers, RECEIVE_WAIT);
-            processor.inspectDeadLetters(customers, RECEIVE_WAIT);
+            processor.processSessions(initialCustomers, RECEIVE_WAIT);
+            processor.inspectDeadLetters(failedCustomers, RECEIVE_WAIT);
             int reprocessed = processor.reprocessDeadLetters(
-                    customers,
+                    failedCustomers,
                     RECEIVE_WAIT,
                     deadLetter -> recoverDemoOrder(dea
... (truncated)
```

### ✅ 21. powershell — 6650ms

**Input:**

```json
{"command":"mvn package --quiet","description":"Build and package the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\orders\AsyncOrderProcessor.java`
- `src\main\java\com\example\orders\AsyncOrderSender.java`
- `src\main\java\com\example\orders\DeadLetteredOrder.java`
- `src\main\java\com\example\orders\Main.java`
- `src\main\java\com\example\orders\Order.java`
- `src\main\java\com\example\orders\OrderMessageMapper.java`
- `src\main\java\com\example\orders\OrderProcessor.java`
- `src\main\java\com\example\orders\OrderSender.java`
- `src\main\java\com\example\orders\OrderStatus.java`
- `src\test\java\com\example\orders\OrderJsonTest.java`

## Copilot Response

I’m grounding the implementation in the current Java SDK patterns, especially sessions, scheduled messages, batching, and dead-letter settlement.

The implementation will use Service Bus sessions keyed by customer and process one session at a time, which enforces FIFO per customer and prevents cross-customer interleaving. Immediate messages are batched per session; high-value orders are scheduled individually because scheduled sends have a different delivery path.

Compilation exposed the mapper’s checked JSON exception. I’m tightening that handling and adding the missing explicit DLQ reprocessing path so the demo does more than inspect failed messages.

The core build now passes. I’m refining the demo’s session targeting so it doesn’t wait on nonexistent DLQ sessions, and adding a true JSON round-trip assertion for the model.

Created the complete Java 17 Maven project with:

- JSON-serializable `Order` model and status enum
- Synchronous and Reactor-based asynchronous senders
- Size-aware batching grouped by customer session
- Correlation IDs, high-priority properties, and 30-second scheduled delivery
- Session-based sequential customer processing
- Explicit dead-lettering with reason and error description
- DLQ inspection and reprocessing workflows
- Managed identity authentication
- Sync and async end-to-end demos
- JSON serialization tests and usage documentation

The packaged artifact is available at `target\service-bus-order-processor-1.0.0.jar`. Configure the session-enabled queue as documented in `README.md`, then run with `mvn exec:java`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Consistent lifecycle management for sync and async Azure Service Bus clients, including per-session receiver clients.

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
hyoka run --prompt-id service-bus-dp-java-order-processor --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)

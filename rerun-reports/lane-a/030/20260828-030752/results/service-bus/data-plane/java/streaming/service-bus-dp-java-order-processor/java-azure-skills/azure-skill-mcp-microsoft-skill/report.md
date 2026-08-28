# Evaluation Report: service-bus-dp-java-order-processor

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 843.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-order-processor` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 20/24 |
| Duration | 843.2s |
| Timestamp | 2026-08-27T19:07:52Z |
| Files Generated | 9 |
| Event Count | 16021 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 596.0s |
| Review | 246.0s |
| **Total** | **843.2s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=703510 out=12025 |
| Turn Count | 14 |

## Tool Calls

### ✅ 1. skill — 2320ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 182031ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Service Bus with managed identity, synchronous and asynchronous senders/processors, batching, scheduling, sessions, and dead-letter handling. Do not deploy or modify Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. glob — 187533ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 187536ms

**Input:**

```json
{"intent":"Get Java Azure Service Bus SDK code-generation best practices for a local-only Maven sample using managed identity, sessions, batching, scheduled messages, async processing, and dead-letter handling. No Azure deployment or resource modification.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 187538ms

**Input:**

```json
{"intent":"Find current official Java Azure Service Bus documentation for ServiceBusClientBuilder, DefaultAzureCredential, sessions, message batching, scheduled messages, dead-lettering, and dead-letter subqueue receivers.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. view — 9862ms

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

### ✅ 7. glob — 9389ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. github-mcp-server-search_code (via github-mcp-server) — 10790ms

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":10,"query":"\"setScheduledEnqueueTime\" \"ServiceBusMessageBatch\" language:Java"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ServiceBusReceiverAsyncClient.java","path":"sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusReceiverAsyncClient.java","repository":"Azure/azure-sdk-for-java","text_matches":[{"fragment":"     * @see <a href=\"https://docs.microsoft.com/azure/service-bus-messaging/service-bus-dead-letter-queues\">Dead letter\n     *     queues</a>\n     */\n    public Mono<Void> deadLetter(ServiceBusReceivedMessage message) {\n        return deadLetter(message, DEFAULT_DEAD_LETTER_OPTIONS);\n    }\n","matches":[{"indices":[172,182],"text":"deadLetter"},{"indices":[235,245],"text":"deadLetter"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusReceiverAsyncClient.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"}]},{"name":"ServiceBusReceiverAsyncClient.java","path":"src/main/java/com/azure/messaging/servicebus/ServiceBusReceiverAsyncClient.java","repository":"rajlakkad90/SP_Azure_ServiceBus_POC","text_matches":[{"fragment":"     * @see <a href=\"https://docs.microsoft.com/azure/service-bus-messaging/service-bus-dead-letter-queues\">Dead letter\n     *     queues</a>\n     */\n    public Mono<Void> deadLetter(ServiceBusReceivedMessage message) {\n        return deadLetter(message, DEFAULT_DEAD_LETTER_OPTIONS);\n    }\n","matches":[{"indices":[172,182],"text":"deadLetter"},{"indices":[235,245],"text":"deadLetter"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/472490919/contents/src/main/java/com/azure/messaging/servicebus/ServiceBusReceiverAsyncClient.java?ref=b064fd8cf55a9ffe89f97efde4d53fcaa9ba5a10","property":"content"}]},{"name":"ServiceBusReceiverAsyncClient.java","path":"sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusReceiverAsyncClient.java",
... (truncated)
```

### ✅ 9. github-mcp-server-search_code (via github-mcp-server) — 10346ms

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":10,"query":"\"ServiceBusSessionReceiverAsyncClient\" \"deadLetter\" language:Java"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ServiceBusMessage.java","path":"sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusMessage.java","repository":"Azure/azure-sdk-for-java","text_matches":[{"fragment":"     * @return The updated {@link ServiceBusMessage}.\n     * @see #getScheduledEnqueueTime()\n     */\n    public ServiceBusMessage setScheduledEnqueueTime(OffsetDateTime scheduledEnqueueTime) {\n        if (scheduledEnqueueTime != null) {\n            amqpAnnotatedMessage.getMessageAnnotations()\n                .put(SCHEDULED_ENQUEUE_UTC_TIME_NAME.getValue(), scheduledEnqueueTime);","matches":[{"indices":[130,153],"text":"setScheduledEnqueueTime"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusMessage.java?ref=2633d1b825689ea803102622c39305e5794e88e6","property":"content"}]},{"name":"ServiceBusMessage.java","path":"src/main/java/com/azure/messaging/servicebus/ServiceBusMessage.java","repository":"rajlakkad90/SP_Azure_ServiceBus_POC","text_matches":[{"fragment":"     * @return The updated {@link ServiceBusMessage}.\n     * @see #getScheduledEnqueueTime()\n     */\n    public ServiceBusMessage setScheduledEnqueueTime(OffsetDateTime scheduledEnqueueTime) {\n        if (scheduledEnqueueTime != null) {\n            amqpAnnotatedMessage.getMessageAnnotations().put(SCHEDULED_ENQUEUE_UTC_TIME_NAME.getValue(),\n                scheduledEnqueueTime);","matches":[{"indices":[130,153],"text":"setScheduledEnqueueTime"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/472490919/contents/src/main/java/com/azure/messaging/servicebus/ServiceBusMessage.java?ref=b064fd8cf55a9ffe89f97efde4d53fcaa9ba5a10","property":"content"}]},{"name":"ServiceBusMessage.java","path":"sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/Ser
... (truncated)
```

### ✅ 10. web_fetch — 7433ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
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

### ✅ 11. github-mcp-server-get_file_contents (via github-mcp-server) — 8958ms

**Input:**

```json
{"owner":"Azure","path":"sdk/servicebus/azure-messaging-servicebus/src/samples/java/com/azure/messaging/servicebus/ServiceBusReceiverClientJavaDocCodeSamples.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: f3eb733eafbe9b11cd1c83312277e5256ae40c42)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.messaging.servicebus;

import com.azure.core.amqp.AmqpRetryOptions;
import com.azure.core.amqp.exception.AmqpException;
import com.azure.core.amqp.implementation.MessageSerializer;
import com.azure.core.annotation.ReturnType;
import com.azure.core.annotation.ServiceClient;
import com.azure.core.annotation.ServiceMethod;
import com.azure.core.http.HttpHeaders;
import com.azure.core.http.rest.PagedFlux;
import com.azure.core.http.rest.PagedResponse;
import com.azure.core.http.rest.PagedResponseBase;
import com.azure.core.util.CoreUtils;
import com.azure.core.util.logging.ClientLogger;
import com.azure.messaging.servicebus.implementation.ManagementConstants;
import com.azure.messaging.servicebus.implementation.MessagingEntityType;
import com.azure.messaging.servicebus.implementation.ServiceBusConstants;
import com.azure.messaging.servicebus.implementation.instrumentation.ServiceBusReceiverInstrumentation;
import com.azure.messaging.servicebus.implementation.instrumentation.ServiceBusTracer;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import reactor.core.publisher.Mono;

import java.nio.ByteBuffer;
import java.nio.charset.CharacterCodingException;
import java.nio.charset.CodingErrorAction;
import java.nio.charset.StandardCharsets;
import java.time.OffsetDateTime;
import java.util.Base64;
import java.util.Objects;

import static com.azure.core.util.FluxUtil.monoError;
import static com.azure.core.util.FluxUtil.pagedFluxError;
import static com.azure.messaging.servicebus.ReceiverOptions.createNamedSessionOptions;

/**
 * This <b>asynchronous</b> session receiver client is used to acquire session locks from a queue or topic and create
 * {@link ServiceBusReceiverAsyncClient} instances that are tied to the locked sessions.  Sessions can be used as a
 * first
... (truncated)
```

### ✅ 12. github-mcp-server-get_file_contents (via github-mcp-server) — 8397ms

**Input:**

```json
{"owner":"Azure","path":"sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusSessionReceiverClient.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 603025d39b904b9efad90197278d28eb47abff9d)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.messaging.servicebus;

import com.azure.core.amqp.AmqpRetryOptions;
import com.azure.core.amqp.exception.AmqpException;
import com.azure.core.annotation.ReturnType;
import com.azure.core.annotation.ServiceClient;
import com.azure.core.annotation.ServiceMethod;
import com.azure.core.http.rest.PagedIterable;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

/**
 * This <b>synchronous</b> session receiver client is used to acquire session locks from a queue or topic and create
 * {@link ServiceBusReceiverClient} instances that are tied to the locked sessions.  Sessions can be used as a first in
 * first out (FIFO) processing of messages.  Queues and topics/subscriptions support Service Bus sessions, however, it
 * must be <a href="https://learn.microsoft.com/azure/service-bus-messaging/enable-message-sessions">enabled at the time
 *      of entity creation</a>.
 *
 * <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication,
 * which is appropriate for most scenarios, including local development and production environments. Additionally, we
 * recommend using
 * <a href="https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/">managed identity</a>
 * for authentication in production environments. You can find more information on different ways of authenticating and
 * their corresponding credential types in the
 * <a href="https://learn.microsoft.com/java/api/overview/azure/identity-readme">Azure Identity documentation"</a>.
 * </p>
 *
 * <p><strong>Sample: Receive message
... (truncated)
```

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server) — 7934ms

**Input:**

```json
{"owner":"Azure","path":"sdk/servicebus/azure-messaging-servicebus/src/main/java/com/azure/messaging/servicebus/ServiceBusSessionReceiverAsyncClient.java","ref":"refs/heads/main","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: d5a1a2ee300f9efc06ce36afdae16d1212c38cbd)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.messaging.servicebus;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.core.util.IterableStream;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.servicebus.models.AbandonOptions;
import com.azure.messaging.servicebus.models.CompleteOptions;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import org.junit.jupiter.api.Test;
import org.reactivestreams.Subscription;
import reactor.core.Disposable;
import reactor.core.publisher.BaseSubscriber;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.time.Duration;
import java.util.Iterator;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * Code snippets demonstrating various {@link ServiceBusReceiverClient} and {@link ServiceBusReceiverAsyncClient}.
 */
public class ServiceBusReceiverClientJavaDocCodeSamples {
    private static final SecureRandom RANDOM = new SecureRandom();

    /**
     * Fully qualified namespace is the host name of the Service Bus resource.  It can be found by navigating to the
     * Service Bus namespace and looking in the "Essentials" panel.
     */
    private final String fullyQualifiedNamespace = System.getenv("AZURE_SERVICEBUS_FULLY_QUALIFIED_DOMAIN_NAME");
    /**
     * Name of a queue inside the Service Bus namespace.
     */
    private final String queueName = System.getenv("AZURE_SERVICEBUS_SAMPLE_QUEUE_NAME");
    /**
     * Name of a topic inside the Service Bus namespace.
     */
    private final String topicName = System.getenv("AZURE_SERVICEBUS_SAMPLE_TOPIC_NAME");
    /**
     * Name of a subscription associated with the {@link #topicName}.
     */
   
... (truncated)
```

### ✅ 14. apply_patch — 697ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eservice-bus-order-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eService Bus Order Processor\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/orders/Order.java\n+package com.example.orders;\n+\n+import java.math.BigDecimal;\n+import java.util.Objects;\n+\n+public final class Order {\n+    public enum Status {\n+        pending,\n+        processing,\n+        completed,\n+        failed\n+    }\n+\n+    private String orderId;\n+    private String customerName;\n+    private String product;\n+    private int quantity;\n+    private BigDecimal totalPrice;\n+    private Status status;\n+\n+    public Order() {\n+    }\n+\n+    public Order(String orderId, String customerName, String product, int quantity,\n+                 BigDecimal totalPrice, Status status) {\n+        this.orderId = requireText(orderId, \"orderId\");\n+        this.customerName = requireText(customerName, \"customerName\");\n+        this.product = requireText(product, \"product\");\n+        if (quantity \u003c= 0) {\n+            throw new IllegalArgumentException(\"quantity must be positive\");\n+        }\n+        this.quantity = quantity;\n+        this.totalPrice = Objects.requireNonNull(totalPrice, \"totalPrice\");\n+        if (totalPrice.signum() \u003c 0) {\n+            throw new IllegalArgumentException(\"totalPrice cannot be negative\");\n+        }\n+        this.status = Objects.requireNonNull(status, \"status\");\n+    }\n+\n+    public String getOrderId() {\n+        return orderId;\n+    }\n+\n+    public void setOrderId(String orderId) {\n+        this.orderId = orderId;\n+    }\n+\n+    public String getCustomerName() {\n+        return customerName;\n+    }\n+\n+    public void setCustomerName(String customerName) {\n+        this.customerName = customerName;\n+    }\n+\n+    public String getProduct() {\n+        return product;\n+    }\n+\n+    public void setProduct(String product) {\n+        this.product = product;\n+    }\n+\n+    public int getQuantity() {\n+        return quantity;\n+    }\n+\n+    public void setQuantity(int quantity) {\n+        this.quantity = quantity;\n+    }\n+\n+    public BigDecimal getTotalPrice() {\n+        return totalPrice;\n+    }\n+\n+    public void setTotalPrice(BigDecimal totalPrice) {\n+        this.totalPrice = totalPrice;\n+    }\n+\n+    public Status getStatus() {\n+        return status;\n+    }\n+\n+    public void setStatus(Status status) {\n+        this.status = status;\n+    }\n+\n+    @Override\n+    public String toString() {\n+        return \"Order{\" +\n+            \"orderId='\" + orderId + '\\'' +\n+            \", customerName='\" + customerName + '\\'' +\n+            \", product='\" + product + '\\'' +\n+            \", quantity=\" + quantity +\n+            \", totalPrice=\" + totalPrice +\n+            \", status=\" + status +\n+            '}';\n+    }\n+\n+    private static String requireText(String value, String name) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(name + \" cannot be blank\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderMessageFactory.java\n+package com.example.orders;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+final class OrderMessageFactory {\n+    static final String PRIORITY_PROPERTY = \"priority\";\n+    static final String HIGH_PRIORITY = \"high\";\n+    static final String NORMAL_PRIORITY = \"normal\";\n+    private static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);\n+\n+    private final BigDecimal highPriorityThreshold;\n+\n+    OrderMessageFactory(BigDecimal highPriorityThreshold) {\n+        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, \"highPriorityThreshold\");\n+    }\n+\n+    ServiceBusMessage create(Order order) {\n+        Objects.requireNonNull(order, \"order\");\n+        boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+\n+        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromObject(order))\n+            .setContentType(\"application/json\")\n+            .setMessageId(order.getOrderId())\n+            .setCorrelationId(order.getOrderId())\n+            .setSessionId(order.getCustomerName());\n+        message.getApplicationProperties().put(\n+            PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY);\n+\n+        if (highPriority) {\n+            message.setScheduledEnqueueTime(OffsetDateTime.now().plus(FRAUD_REVIEW_DELAY));\n+        }\n+        return message;\n+    }\n+\n+    Order deserialize(BinaryData body) {\n+        return body.toObject(Order.class);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class SyncOrderSender implements AutoCloseable {\n+    private final ServiceBusSenderClient sender;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public SyncOrderSender(String namespace, String queueName, TokenCredential credential,\n+                           BigDecimal highPriorityThreshold) {\n+        sender = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sender()\n+            .queueName(queueName)\n+            .buildClient();\n+        messageFactory = new OrderMessageFactory(highPriorityThreshold);\n+    }\n+\n+    public void sendOrder(Order order) {\n+        sender.sendMessage(messageFactory.create(order));\n+    }\n+\n+    public void sendOrders(List\u003cOrder\u003e orders) {\n+        Objects.requireNonNull(orders, \"orders\");\n+        if (orders.isEmpty()) {\n+            return;\n+        }\n+\n+        ServiceBusMessageBatch batch = sender.createMessageBatch();\n+        for (Order order : orders) {\n+            ServiceBusMessage message = messageFactory.create(order);\n+            if (!batch.tryAddMessage(message)) {\n+                if (batch.getCount() == 0) {\n+                    throw new IllegalArgumentException(\"Order \" + order.getOrderId()\n+                        + \" is too large for a Service Bus batch\");\n+                }\n+                sender.sendMessages(batch);\n+                batch = sender.createMessageBatch();\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalArgumentException(\"Order \" + order.getOrderId()\n+                        + \" is too large for a Service Bus batch\");\n+                }\n+            }\n+        }\n+        if (batch.getCount() \u003e 0) {\n+            sender.sendMessages(batch);\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class AsyncOrderSender implements AutoCloseable {\n+    private final ServiceBusSenderAsyncClient sender;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public AsyncOrderSender(String namespace, String queueName, TokenCredential credential,\n+                            BigDecimal highPriorityThreshold) {\n+        sender = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sender()\n+            .queueName(queueName)\n+            .buildAsyncClient();\n+        messageFactory = new OrderMessageFactory(highPriorityThreshold);\n+    }\n+\n+    public Mono\u003cVoid\u003e sendOrder(Order order) {\n+        return sender.sendMessage(messageFactory.create(order));\n+    }\n+\n+    public Mono\u003cVoid\u003e sendOrders(List\u003cOrder\u003e orders) {\n+        Objects.requireNonNull(orders, \"orders\");\n+        return sendBatch(orders, 0);\n+    }\n+\n+    private Mono\u003cVoid\u003e sendBatch(List\u003cOrder\u003e orders, int startIndex) {\n+        if (startIndex \u003e= orders.size()) {\n+            return Mono.empty();\n+        }\n+\n+        return sender.createMessageBatch().flatMap(batch -\u003e {\n+            int nextIndex = fillBatch(batch, orders, startIndex);\n+            return sender.sendMessages(batch)\n+                .then(Mono.defer(() -\u003e sendBatch(orders, nextIndex)));\n+        });\n+    }\n+\n+    private int fillBatch(ServiceBusMessageBatch batch, List\u003cOrder\u003e orders, int startIndex) {\n+        int index = startIndex;\n+        while (index \u003c orders.size()) {\n+            Order order = orders.get(index);\n+            ServiceBusMessage message = messageFactory.create(order);\n+            if (!batch.tryAddMessage(message)) {\n+                if (index == startIndex) {\n+                    throw new IllegalArgumentException(\"Order \" + order.getOrderId()\n+                        + \" is too large for a Service Bus batch\");\n+                }\n+                break;\n+            }\n+            index++;\n+        }\n+        return index;\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.IterableStream;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.time.Duration;\n+import java.util.function.Function;\n+\n+public final class SyncOrderProcessor {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);\n+    private static final String DESERIALIZATION_FAILURE = \"OrderDeserializationFailed\";\n+\n+    private final String namespace;\n+    private final String queueName;\n+    private final TokenCredential credential;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public SyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {\n+        this.namespace = namespace;\n+        this.queueName = queueName;\n+        this.credential = credential;\n+        this.messageFactory = new OrderMessageFactory(java.math.BigDecimal.ZERO);\n+    }\n+\n+    public void processCustomer(String customerName, int maxMessages, Duration waitTime) {\n+        try (ServiceBusSessionReceiverClient sessions = sessionReceiver(false);\n+             ServiceBusReceiverClient receiver = sessions.acceptSession(customerName)) {\n+            IterableStream\u003cServiceBusReceivedMessage\u003e messages =\n+                receiver.receiveMessages(maxMessages, waitTime);\n+            for (ServiceBusReceivedMessage message : messages) {\n+                try {\n+                    Order order = messageFactory.deserialize(message.getBody());\n+                    order.setStatus(Order.Status.processing);\n+                    LOGGER.info(\"Synchronously processed order: {}\", order);\n+                    order.setStatus(Order.Status.completed);\n+                    receiver.complete(message);\n+                } catch (RuntimeException exception) {\n+                    LOGGER.error(\"Dead-lettering message {}: {}\",\n+                        message.getMessageId(), exception.getMessage());\n+                    receiver.deadLetter(message, new DeadLetterOptions()\n+                        .setDeadLetterReason(DESERIALIZATION_FAILURE)\n+                        .setDeadLetterErrorDescription(exception.getMessage()));\n+                }\n+            }\n+        }\n+    }\n+\n+    public void inspectDeadLetters(String customerName, int maxMessages, Duration waitTime) {\n+        try (ServiceBusSessionReceiverClient sessions = sessionReceiver(true);\n+             ServiceBusReceiverClient receiver = sessions.acceptSession(customerName)) {\n+            for (ServiceBusReceivedMessage message : receiver.receiveMessages(maxMessages, waitTime)) {\n+                LOGGER.warn(\"DLQ message id={}, reason={}, description={}, body={}\",\n+                    message.getMessageId(), message.getDeadLetterReason(),\n+                    message.getDeadLetterErrorDescription(), message.getBody());\n+                receiver.abandon(message);\n+            }\n+        }\n+    }\n+\n+    public void reprocessDeadLetters(String customerName, int maxMessages, Duration waitTime,\n+                                     Function\u003cServiceBusReceivedMessage, Order\u003e repair,\n+                                     SyncOrderSender sender) {\n+        try (ServiceBusSessionReceiverClient sessions = sessionReceiver(true);\n+             ServiceBusReceiverClient receiver = sessions.acceptSession(customerName)) {\n+            for (ServiceBusReceivedMessage message : receiver.receiveMessages(maxMessages, waitTime)) {\n+                try {\n+                    sender.sendOrder(repair.apply(message));\n+                    receiver.complete(message);\n+                    LOGGER.info(\"Requeued dead-letter message {}\", message.getMessageId());\n+                } catch (RuntimeException exception) {\n+                    LOGGER.error(\"Could not reprocess dead-letter message {}\",\n+                        message.getMessageId(), exception);\n+                    receiver.abandon(message);\n+                }\n+            }\n+        }\n+    }\n+\n+    private ServiceBusSessionReceiverClient sessionReceiver(boolean deadLetterQueue) {\n+        ServiceBusClientBuilder.ServiceBusSessionReceiverClientBuilder builder =\n+            new ServiceBusClientBuilder()\n+                .credential(namespace, credential)\n+                .sessionReceiver()\n+                .queueName(queueName)\n+                .disableAutoComplete();\n+        if (deadLetterQueue) {\n+            builder.subQueue(SubQueue.DEAD_LETTER_QUEUE);\n+        }\n+        return builder.buildClient();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.function.Function;\n+\n+public final class AsyncOrderProcessor {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);\n+    private static final String DESERIALIZATION_FAILURE = \"OrderDeserializationFailed\";\n+\n+    private final String namespace;\n+    private final String queueName;\n+    private final TokenCredential credential;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public AsyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {\n+        this.namespace = namespace;\n+        this.queueName = queueName;\n+        this.credential = credential;\n+        this.messageFactory = new OrderMessageFactory(java.math.BigDecimal.ZERO);\n+    }\n+\n+    public Mono\u003cVoid\u003e processCustomer(String customerName, int maxMessages, Duration waitTime) {\n+        ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(false);\n+        return sessions.acceptSession(customerName)\n+            .flatMap(receiver -\u003e receiver.receiveMessages(maxMessages, waitTime)\n+                .concatMap(message -\u003e processMessage(receiver, message))\n+                .then()\n+                .doFinally(signal -\u003e receiver.close()))\n+            .doFinally(signal -\u003e sessions.close());\n+    }\n+\n+    public Mono\u003cVoid\u003e inspectDeadLetters(String customerName, int maxMessages, Duration waitTime) {\n+        ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);\n+        return sessions.acceptSession(customerName)\n+            .flatMap(receiver -\u003e receiver.receiveMessages(maxMessages, waitTime)\n+                .concatMap(message -\u003e {\n+                    LOGGER.warn(\"DLQ message id={}, reason={}, description={}, body={}\",\n+                        message.getMessageId(), message.getDeadLetterReason(),\n+                        message.getDeadLetterErrorDescription(), message.getBody());\n+                    return receiver.abandon(message);\n+                })\n+                .then()\n+                .doFinally(signal -\u003e receiver.close()))\n+            .doFinally(signal -\u003e sessions.close());\n+    }\n+\n+    public Mono\u003cVoid\u003e reprocessDeadLetters(String customerName, int maxMessages, Duration waitTime,\n+                                           Function\u003cServiceBusReceivedMessage, Order\u003e repair,\n+                                           AsyncOrderSender sender) {\n+        ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);\n+        return sessions.acceptSession(customerName)\n+            .flatMap(receiver -\u003e receiver.receiveMessages(maxMessages, waitTime)\n+                .concatMap(message -\u003e Mono.fromCallable(() -\u003e repair.apply(message))\n+                    .flatMap(sender::sendOrder)\n+                    .then(receiver.complete(message))\n+                    .doOnSuccess(ignored -\u003e\n+                        LOGGER.info(\"Requeued dead-letter message {}\", message.getMessageId()))\n+                    .onErrorResume(error -\u003e {\n+                        LOGGER.error(\"Could not reprocess dead-letter message {}\",\n+                            message.getMessageId(), error);\n+                        return receiver.abandon(message);\n+                    }))\n+                .then()\n+                .doFinally(signal -\u003e receiver.close()))\n+            .doFinally(signal -\u003e sessions.close());\n+    }\n+\n+    private Mono\u003cVoid\u003e processMessage(ServiceBusReceiverAsyncClient receiver,\n+                                      ServiceBusReceivedMessage message) {\n+        return Mono.fromCallable(() -\u003e messageFactory.deserialize(message.getBody()))\n+            .doOnNext(order -\u003e {\n+                order.setStatus(Order.Status.processing);\n+                LOGGER.info(\"Asynchronously processed order: {}\", order);\n+                order.setStatus(Order.Status.completed);\n+            })\n+            .then(receiver.complete(message))\n+            .onErrorResume(error -\u003e {\n+                LOGGER.error(\"Dead-lettering message {}: {}\", message.getMessageId(), error.getMessage());\n+                return receiver.deadLetter(message, new DeadLetterOptions()\n+                    .setDeadLetterReason(DESERIALIZATION_FAILURE)\n+                    .setDeadLetterErrorDescription(error.getMessage()));\n+            });\n+    }\n+\n+    private ServiceBusSessionReceiverAsyncClient sessionReceiver(boolean deadLetterQueue) {\n+        ServiceBusClientBuilder.ServiceBusSessionReceiverClientBuilder builder =\n+            new ServiceBusClientBuilder()\n+                .credential(namespace, credential)\n+                .sessionReceiver()\n+                .queueName(queueName)\n+                .disableAutoComplete();\n+        if (deadLetterQueue) {\n+            builder.subQueue(SubQueue.DEAD_LETTER_QUEUE);\n+        }\n+        return builder.buildAsyncClient();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Main.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.BinaryData;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.util.List;\n+\n+public final class Main {\n+    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal(\"1000.00\");\n+    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(5);\n+    private static final String DEMO_CUSTOMER = \"Contoso\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String namespace = requiredEnvironmentVariable(\"SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE\");\n+        String queueName = requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+        String managedIdentityClientId = System.getenv(\"AZURE_CLIENT_ID\");\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            credentialBuilder.clientId(managedIdentityClientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+\n+        runSyncDemo(namespace, queueName, credential);\n+        runAsyncDemo(namespace, queueName, credential);\n+    }\n+\n+    private static void runSyncDemo(String namespace, String queueName, TokenCredential credential)\n+        throws InterruptedException {\n+        SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential);\n+        try (SyncOrderSender sender =\n+                 new SyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {\n+            sender.sendOrders(demoOrders(\"sync\"));\n+            sendMalformedMessage(namespace, queueName, credential, \"sync-invalid\");\n+\n+            processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);\n+            processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT);\n+            processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,\n+                ignored -\u003e repairedOrder(\"sync-repaired\"), sender);\n+            processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);\n+\n+            Thread.sleep(Duration.ofSeconds(31));\n+            processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);\n+        }\n+    }\n+\n+    private static void runAsyncDemo(String namespace, String queueName, TokenCredential credential) {\n+        AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential);\n+        try (AsyncOrderSender sender =\n+                 new AsyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {\n+            sender.sendOrders(demoOrders(\"async\"))\n+                .then(MonoBridge.fromRunnable(() -\u003e\n+                    sendMalformedMessage(namespace, queueName, credential, \"async-invalid\")))\n+                .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))\n+                .then(processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT))\n+                .then(processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,\n+                    ignored -\u003e repairedOrder(\"async-repaired\"), sender))\n+                .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))\n+                .then(reactor.core.publisher.Mono.delay(Duration.ofSeconds(31)))\n+                .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))\n+                .block();\n+        }\n+    }\n+\n+    private static List\u003cOrder\u003e demoOrders(String prefix) {\n+        return List.of(\n+            new Order(prefix + \"-001\", DEMO_CUSTOMER, \"Keyboard\", 2,\n+                new BigDecimal(\"179.98\"), Order.Status.pending),\n+            new Order(prefix + \"-002\", DEMO_CUSTOMER, \"Server\", 1,\n+                new BigDecimal(\"5500.00\"), Order.Status.pending));\n+    }\n+\n+    private static Order repairedOrder(String orderId) {\n+        return new Order(orderId, DEMO_CUSTOMER, \"Replacement item\", 1,\n+            new BigDecimal(\"25.00\"), Order.Status.pending);\n+    }\n+\n+    private static void sendMalformedMessage(String namespace, String queueName,\n+                                             TokenCredential credential, String messageId) {\n+        try (ServiceBusSenderClient sender = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sender()\n+            .queueName(queueName)\n+            .buildClient()) {\n+            ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(\"{not-json\"))\n+                .setContentType(\"application/json\")\n+                .setMessageId(messageId)\n+                .setCorrelationId(messageId)\n+                .setSessionId(DEMO_CUSTOMER);\n+            sender.sendMessage(message);\n+        }\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(name + \" must be set\");\n+        }\n+        return value;\n+    }\n+\n+    private static final class MonoBridge {\n+        private MonoBridge() {\n+        }\n+\n+        static reactor.core.publisher.Mono\u003cVoid\u003e fromRunnable(Runnable runnable) {\n+            return reactor.core.publisher.Mono.fromRunnable(runnable);\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Service Bus Order Processor\n+\n+A Java 17 Maven sample with synchronous and asynchronous Azure Service Bus senders and session receivers.\n+\n+## Prerequisites\n+\n+- Java 17 and Maven 3.9+\n+- An Azure Service Bus queue created with **sessions enabled**\n+- A system-assigned or user-assigned managed identity with the **Azure Service Bus Data Sender** and\n+  **Azure Service Bus Data Receiver** roles\n+\n+Set these environment variables:\n+\n+```powershell\n+$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = \"your-namespace.servicebus.windows.net\"\n+$env:SERVICE_BUS_QUEUE_NAME = \"orders\"\n+# Only for a user-assigned managed identity:\n+$env:AZURE_CLIENT_ID = \"00000000-0000-0000-0000-000000000000\"\n+```\n+\n+Build and run:\n+\n+```powershell\n+mvn clean package\n+mvn exec:java\n+```\n+\n+The demo intentionally sends malformed JSON to exercise dead-letter inspection and repair. Orders over\n+`$1,000` are marked high priority and scheduled 30 seconds into the future. The customer name is used as\n+the Service Bus session ID, and each processor consumes one session serially to preserve per-customer order.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/pom.xml
@@ -1,0 +1,66 @@
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
+        <azure.sdk.bom.version>1.3.8</azure.sdk.bom.version>
+        <slf4j.version>2.0.17</slf4j.version>
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
+            <artifactId>azure-messaging-servicebus</artifactId>
+        
... (truncated)
```

### ✅ 15. powershell — 6503ms

**Input:**

```json
{"command":"mvn -q -DskipTests package","description":"Compile the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[ERROR] COMPILATION ERROR : 
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java:[36,42] method receiveMessages in class com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient cannot be applied to given types;
  required: no arguments
  found:    int,java.time.Duration
  reason: actual and formal argument lists differ in length
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java:[40,23] incompatible types: reactor.core.publisher.Mono<java.lang.Object> cannot be converted to reactor.core.publisher.Mono<java.lang.Void>
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java:[46,42] method receiveMessages in class com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient cannot be applied to given types;
  required: no arguments
  found:    int,java.time.Duration
  reason: actual and formal argument lists differ in length
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java:[55,23] incompatible types: reactor.core.publisher.Mono<jav
... (truncated)
```

### ✅ 16. apply_patch — 680ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n@@\n         ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(false);\n         return sessions.acceptSession(customerName)\n-            .flatMap(receiver -\u003e receiver.receiveMessages(maxMessages, waitTime)\n+            .flatMap(receiver -\u003e receiver.receiveMessages()\n+                .take(maxMessages)\n+                .take(waitTime)\n                 .concatMap(message -\u003e processMessage(receiver, message))\n                 .then()\n                 .doFinally(signal -\u003e receiver.close()))\n@@\n         ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);\n         return sessions.acceptSession(customerName)\n-            .flatMap(receiver -\u003e receiver.receiveMessages(maxMessages, waitTime)\n+            .flatMap(receiver -\u003e receiver.receiveMessages()\n+                .take(maxMessages)\n+                .take(waitTime)\n                 .concatMap(message -\u003e {\n                     LOGGER.warn(\"DLQ message id={}, reason={}, description={}, body={}\",\n@@\n         ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);\n         return sessions.acceptSession(customerName)\n-            .flatMap(receiver -\u003e receiver.receiveMessages(maxMessages, waitTime)\n-                .concatMap(message -\u003e Mono.fromCallable(() -\u003e repair.apply(message))\n+            .flatMap(receiver -\u003e receiver.receiveMessages()\n+                .take(maxMessages)\n+                .take(waitTime)\n+                .concatMap(message -\u003e Mono.fromCallable(() -\u003e (Order) repair.apply(message))\n                     .flatMap(sender::sendOrder)\n@@\n     private Mono\u003cVoid\u003e processMessage(ServiceBusReceiverAsyncClient receiver,\n                                       ServiceBusReceivedMessage message) {\n-        return Mono.fromCallable(() -\u003e messageFactory.deserialize(message.getBody()))\n-            .doOnNext(order -\u003e {\n+        return Mono.defer(() -\u003e {\n+            try {\n+                Order order = messageFactory.deserialize(message.getBody());\n                 order.setStatus(Order.Status.processing);\n                 LOGGER.info(\"Asynchronously processed order: {}\", order);\n                 order.setStatus(Order.Status.completed);\n-            })\n-            .then(receiver.complete(message))\n-            .onErrorResume(error -\u003e {\n-                LOGGER.error(\"Dead-lettering message {}: {}\", message.getMessageId(), error.getMessage());\n+                return receiver.complete(message);\n+            } catch (RuntimeException error) {\n+                LOGGER.error(\"Dead-lettering message {}: {}\",\n+                    message.getMessageId(), error.getMessage());\n                 return receiver.deadLetter(message, new DeadLetterOptions()\n                     .setDeadLetterReason(DESERIALIZATION_FAILURE)\n                     .setDeadLetterErrorDescription(error.getMessage()));\n-            });\n+            }\n+        });\n     }\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n-            Thread.sleep(Duration.ofSeconds(31));\n+            Thread.sleep(Duration.ofSeconds(31).toMillis());\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/AsyncOrderProcessor.java
@@ -33,7 +33,9 @@
     public Mono<Void> processCustomer(String customerName, int maxMessages, Duration waitTime) {
         ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(false);
         return sessions.acceptSession(customerName)
-            .flatMap(receiver -> receiver.receiveMessages(maxMessages, waitTime)
+            .flatMap(receiver -> receiver.receiveMessages()
+                .take(maxMessages)
+                .take(waitTime)
                 .concatMap(message -> processMessage(receiver, message))
                 .then()
                 .doFinally(signal -> receiver.close()))
@@ -43,7 +45,9 @@
     public Mono<Void> inspectDeadLetters(String customerName, int maxMessages, Duration waitTime) {
         ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);
         return sessions.acceptSession(customerName)
-            .flatMap(receiver -> receiver.receiveMessages(maxMessages, waitTime)
+            .flatMap(receiver -> receiver.receiveMessages()
+                .take(maxMessages)
+                .take(waitTime)
                 .concatMap(message -> {
                     LOGGER.warn("D
... (truncated)
```

### ✅ 17. powershell — 5870ms

**Input:**

```json
{"command":"mvn -q -DskipTests package","description":"Recompile the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 1 completed with exit code 0>
```

### ✅ 18. apply_patch — 693ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import reactor.core.publisher.Mono;\n@@\n             sender.sendOrders(demoOrders(\"async\"))\n-                .then(MonoBridge.fromRunnable(() -\u003e\n-                    sendMalformedMessage(namespace, queueName, credential, \"async-invalid\")))\n+                .then(sendMalformedMessageAsync(namespace, queueName, credential, \"async-invalid\"))\n                 .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))\n@@\n     private static void sendMalformedMessage(String namespace, String queueName,\n                                              TokenCredential credential, String messageId) {\n@@\n         }\n     }\n \n+    private static Mono\u003cVoid\u003e sendMalformedMessageAsync(String namespace, String queueName,\n+                                                        TokenCredential credential, String messageId) {\n+        return Mono.using(\n+            () -\u003e new ServiceBusClientBuilder()\n+                .credential(namespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildAsyncClient(),\n+            sender -\u003e sender.sendMessage(malformedMessage(messageId)),\n+            ServiceBusSenderAsyncClient::close);\n+    }\n+\n+    private static ServiceBusMessage malformedMessage(String messageId) {\n+        return new ServiceBusMessage(BinaryData.fromString(\"{not-json\"))\n+            .setContentType(\"application/json\")\n+            .setMessageId(messageId)\n+            .setCorrelationId(messageId)\n+            .setSessionId(DEMO_CUSTOMER);\n+    }\n+\n@@\n-    private static final class MonoBridge {\n-        private MonoBridge() {\n-        }\n-\n-        static reactor.core.publisher.Mono\u003cVoid\u003e fromRunnable(Runnable runnable) {\n-            return reactor.core.publisher.Mono.fromRunnable(runnable);\n-        }\n-    }\n }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java
@@ -5,7 +5,9 @@
 import com.azure.identity.ManagedIdentityCredentialBuilder;
 import com.azure.messaging.servicebus.ServiceBusClientBuilder;
 import com.azure.messaging.servicebus.ServiceBusMessage;
+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
 import com.azure.messaging.servicebus.ServiceBusSenderClient;
+import reactor.core.publisher.Mono;
 
 import java.math.BigDecimal;
 import java.time.Duration;
@@ -58,8 +60,7 @@
         try (AsyncOrderSender sender =
                  new AsyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {
             sender.sendOrders(demoOrders("async"))
-                .then(MonoBridge.fromRunnable(() ->
-                    sendMalformedMessage(namespace, queueName, credential, "async-invalid")))
+                .then(sendMalformedMessageAsync(namespace, queueName, credential, "async-invalid"))
                 .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
                 .then(processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
                 .then(processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,
@@ -100,6 +101,26 @@
         }
     }
 
+    private static Mono<Void> sen
... (truncated)
```

### ✅ 19. rg — 752ms

**Input:**

```json
{"-C":4,"-n":true,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778\\src\\main\\java\\com\\example\\orders\\Main.java","pattern":"sendMalformedMessage|ServiceBusMessage message"}
```

**Output:**

```
[grep content: 43 matches across 1 file(s) under C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders]

Main.java (43 match(es)):
  41-         SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential);
  42-         try (SyncOrderSender sender =
  43-                  new SyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {
  44-             sender.sendOrders(demoOrders("sync"));
  45:             sendMalformedMessage(namespace, queueName, credential, "sync-invalid");
  46- 
  47-             processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);
  48-             processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT);
  49-             processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,
  59-         AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential);
  60-         try (AsyncOrderSender sender =
  61-                  new AsyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {
  62-             sender.sendOrders(demoOrders("async"))
  63:                 .then(sendMalformedMessageAsync(namespace, queueName, credential, "async-invalid"))
  64-                 .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
  65-                 .then(processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
  66-                 .then(processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,
  67-                     ignored -> repairedOrder("async-repaired"), sender))
  84-         return new Order(orderId, DEMO_CUSTOMER, "Replacement item", 1,
  85-             new BigDecimal("25.00"), Order.Status.pending);
  86-     }
  87- 
  88:     private static void sendMalformedMessage(String namespace, String queueName,
  89-                                              TokenCredential credential, String 
... (truncated)
```

### ✅ 20. apply_patch — 793ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n             .sender()\n             .queueName(queueName)\n             .buildClient()) {\n-            ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(\"{not-json\"))\n-                .setContentType(\"application/json\")\n-                .setMessageId(messageId)\n-                .setCorrelationId(messageId)\n-                .setSessionId(DEMO_CUSTOMER);\n-            sender.sendMessage(message);\n+            sender.sendMessage(malformedMessage(messageId));\n         }\n     }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-order-processor-java-azure-skills-azure-skill-mcp-microsoft-skill-2998465778/src/main/java/com/example/orders/Main.java
@@ -92,12 +92,7 @@
             .sender()
             .queueName(queueName)
             .buildClient()) {
-            ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString("{not-json"))
-                .setContentType("application/json")
-                .setMessageId(messageId)
-                .setCorrelationId(messageId)
-                .setSessionId(DEMO_CUSTOMER);
-            sender.sendMessage(message);
+            sender.sendMessage(malformedMessage(messageId));
         }
     }
 

```

### ✅ 21. powershell — 6752ms

**Input:**

```json
{"command":"mvn -q -DskipTests clean package","description":"Build the completed Java project","initial_wait":180,"mode":"sync"}
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
- `src\main\java\com\example\orders\Main.java`
- `src\main\java\com\example\orders\Order.java`
- `src\main\java\com\example\orders\OrderMessageFactory.java`
- `src\main\java\com\example\orders\SyncOrderProcessor.java`
- `src\main\java\com\example\orders\SyncOrderSender.java`

## Copilot Response

I’m grounding the implementation in the Java SDK’s session, batching, scheduling, and dead-letter APIs before scaffolding the project.

The design will use Service Bus sessions keyed by customer, one active session/call at a time for strict sequencing, explicit settlement, and size-aware batches containing both immediate and 30-second scheduled messages.

Compilation exposed an SDK distinction: asynchronous receivers expose a streaming `Flux` rather than bounded receive overloads. I’m adapting the async path to time- and count-bounded streaming while preserving serial settlement.

Created a complete Java 17 Maven project with:

- JSON-serializable `Order` model and status enum
- Synchronous and asynchronous Service Bus senders
- Size-aware message batching
- Correlation IDs, customer-based session IDs, priority metadata, and 30-second scheduling for high-value orders
- Sync/async session processors with explicit completion and dead-letter handling
- DLQ inspection and repaired-message reprocessing
- Managed identity authentication
- A `Main` demo covering the full sync and async lifecycle
- Setup and execution instructions in `README.md`

The target Service Bus queue must have **sessions enabled**. Configure `SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE`, `SERVICE_BUS_QUEUE_NAME`, and optionally `AZURE_CLIENT_ID`, then run:

```powershell
mvn clean package
mvn exec:java
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Consistent lifecycle management across synchronous, asynchronous, main-demo, and temporary malformed-message clients.

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
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 6.4% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.7500** | **90.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-java-order-processor --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)

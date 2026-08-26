# Evaluation Report: service-bus-dp-java-order-processor

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 602.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-order-processor` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 19/24 |
| Duration | 602.2s |
| Timestamp | 2026-08-26T13:38:32Z |
| Files Generated | 15 |
| Event Count | 19198 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 302.7s |
| Review | 298.2s |
| **Total** | **602.2s** |

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
| Token Usage | in=797809 out=16272 |
| Turn Count | 14 |

## Prompt Sent

```
Create a small Java 17 Maven project that implements an order processing system using Azure Service Bus.

The project needs:

- A **model class** for an Order with fields for order ID, customer name, product, quantity, total price, and status (pending/processing/completed/failed). It should be serializable to and from JSON.

- A **sender class** (both sync and async versions) that publishes order messages to a Service Bus queue. It should support sending individual orders and sending a batch of orders efficiently (respecting the maximum batch size so messages aren't rejected). Each message should carry the order ID as a correlation property, and orders above a certain dollar threshold should be sent as high-priority with a scheduled delivery delay of 30 seconds (to allow for fraud review before processing).

- A **processor class** (both sync and async versions) that receives and processes orders from the queue. It should handle messages as they arrive, deserialize them, and log the results. If processing fails (e.g., a deserialization error), the message should be sent to the dead-letter queue with a reason string rather than being silently abandoned. The processor should also be able to read from the dead-letter queue so failed messages can be inspected and reprocessed. It should guarantee that orders from the same customer are processed in sequence, not interleaved with other customers' orders.

- A **Main class** that demos both implementations: connects to the Service Bus namespace (from an environment variable) with managed identity, runs the full send/receive/dead-letter cycle using the sync implementation first, then repeats with the async implementation.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 2. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-order-processor-java-azure-tools-with-azure-tools-789081597","pattern":"**/*"}
```

### 🔧 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Get Java 17 Azure Service Bus and managed identity code-generation best practices for a local Maven sample. Do not provision or modify Azure resources.","learn":true}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Java Azure Service Bus SDK documentation and examples for sync/async senders, size-aware message batching, scheduling messages, processors with manual settlement and dead-lettering, dead-letter subqueue receivers, sessions for ordered processing, and managed identity authentication.","learn":true}
```

### 🔧 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code generation best practices for a Java 17 Maven Azure Service Bus sample using managed identity, local-only validation, and no resource provisioning.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current Java Azure Service Bus SDK documentation for senders, processors, sessions, batching, scheduling, dead-letter queues, manual settlement, and managed identity.","parameters":{"query":"Azure Service Bus Java SDK ServiceBusSenderClient ServiceBusSenderAsyncClient createMessageBatch scheduleMessage processor manual settlement deadLetterMessage sessions ordered processing dead letter subqueue managed identity"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java code samples for Azure Service Bus size-aware batch sending, scheduled messages, session processors, dead-lettering, and dead-letter subqueue receivers.","parameters":{"language":"java","query":"Azure Service Bus Java ServiceBusMessageBatch tryAddMessage scheduleMessage ServiceBusSessionProcessorClient deadLetterMessage SubQueue.DEAD_LETTER_QUEUE"}}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751567531-copilot-tool-output-4356eed9108b40c8a76fc12bf2fd73b7.txt' | ConvertFrom-Json; $d.results | Select-Object title,contentUrl | Format-Table -AutoSize | Out-String -Width 240","description":"List relevant documentation results"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751567360-copilot-tool-output-6faa2b755f60408e88af19a11aed0eda.txt' | ConvertFrom-Json; $d.results | ForEach-Object { \"TITLE: $($_.description)`nURL: $($_.url)`nCODE:`n$($_.codeSnippet)`n---\" } | Select-String -Pattern 'createMessageBatch|tryAddMessage|scheduleMessage|sessionProcessor|deadLetter|SubQueue' -Context 1,12 | Select-Object -First 12 | Out-String -Width 240","description":"Extract official Service Bus code samples"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"mvn -version; java -version","description":"Check Java and Maven availability","initial_wait":30}
```

### 🔧 12. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751567531-copilot-tool-output-4356eed9108b40c8a76fc12bf2fd73b7.txt","view_range":[1,8]}
```

### 🔧 13. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751567360-copilot-tool-output-6faa2b755f60408e88af19a11aed0eda.txt","view_range":[1,8]}
```

### 🔧 14. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Confirm Java SDK APIs for session-enabled queue processing and asynchronous session receivers, including dead-letter subqueue and settlement.","parameters":{"query":"site:learn.microsoft.com/java/api com.azure.messaging.servicebus ServiceBusSessionReceiverAsyncClient acceptNextSession ServiceBusSessionProcessorClientBuilder subQueue deadLetterMessage DeadLetterOptions"}}
```

### 🔧 15. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Confirm whether ServiceBusMessage supports scheduled enqueue time and how scheduled messages interact with batches in Java.","parameters":{"query":"site:learn.microsoft.com/java/api com.azure.messaging.servicebus ServiceBusMessage setScheduledEnqueueTime ServiceBusMessageBatch scheduled message Java"}}
```

### 🔧 16. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the full current Azure Service Bus Java SDK overview identified by documentation search.","parameters":{"url":"https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable"}}
```

### 🔧 17. rg

**Input:**

```json
{"-C":2,"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751607286-copilot-tool-output-4718f4c84ca74a30803ff57b6c0641d3.txt","pattern":"setScheduledEnqueueTime|scheduleMessage|scheduleMessages"}
```

### 🔧 18. rg

**Input:**

```json
{"-C":2,"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751607043-copilot-tool-output-32827728209f4c959a2124401ab42f7d.txt","pattern":"acceptNextSession|buildAsyncClient|subQueue|deadLetterMessage|DeadLetterOptions"}
```

### 🔧 19. rg

**Input:**

```json
{"-C":3,"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751607720-copilot-tool-output-863f690a2a6d4a659f503da74c8df224.txt","pattern":"sessionProcessor\\(|maxConcurrentSessions|SubQueue.DEAD_LETTER_QUEUE|deadLetter"}
```

### 🔧 20. powershell

**Input:**

```json
{"command":"Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751607286-copilot-tool-output-4718f4c84ca74a30803ff57b6c0641d3.txt' | Where-Object { $_.StartsWith('{') } | ForEach-Object { $j = $_ | ConvertFrom-Json; foreach ($r in $j.results) { if ($r.content -match 'setScheduledEnqueueTime') { $m = [regex]::Match($r.content, '.{0,500}setScheduledEnqueueTime.{0,900}', 'Singleline'); \"TITLE: $($r.title)`nURL: $($r.contentUrl)`n$($m.Value)`n---\" } } } | Select-Object -First 8 | Out-String -Width 220","description":"Extract scheduled message API details"}
```

### 🔧 21. powershell

**Input:**

```json
{"command":"Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787751607043-copilot-tool-output-32827728209f4c959a2124401ab42f7d.txt' | Where-Object { $_.StartsWith('{') } | ForEach-Object { $j = $_ | ConvertFrom-Json; foreach ($r in $j.results) { if ($r.content -match 'acceptNextSession|deadLetterMessage|subQueue') { $m = [regex]::Matches($r.content, '.{0,350}(acceptNextSession|deadLetterMessage|subQueue).{0,1000}', 'Singleline'); foreach ($x in $m) { \"TITLE: $($r.title)`nURL: $($r.contentUrl)`n$($x.Value)`n---\" } } } } | Select-Object -First 12 | Out-String -Width 220","description":"Extract session and dead-letter API details"}
```

### 🔧 22. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-service-bus-order-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure Service Bus Order Processor\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cazure.servicebus.version\u003e7.17.17\u003c/azure.servicebus.version\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+        \u003cjackson.version\u003e2.18.3\u003c/jackson.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+        \u003cjunit.version\u003e5.12.2\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.servicebus.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-api\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** End Patch\n"
```

### 🔧 23. powershell

**Input:**

```json
{"command":"mvn -q dependency:build-classpath \"-Dmdep.outputFile=target\\classpath.txt\"; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $cp = Get-Content -Raw 'target\\classpath.txt'; javap -classpath $cp com.azure.messaging.servicebus.ServiceBusSessionReceiverClient com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient com.azure.messaging.servicebus.ServiceBusClientBuilder com.azure.messaging.servicebus.ServiceBusReceivedMessageContext com.azure.messaging.servicebus.models.DeadLetterOptions","description":"Inspect Service Bus Java APIs","initial_wait":180,"mode":"sync"}
```

### 🔧 24. powershell

**Input:**

```json
{"command":"$cp = Get-Content -Raw 'target\\classpath.txt'; javap -classpath $cp 'com.azure.messaging.servicebus.ServiceBusClientBuilder$ServiceBusSessionReceiverClientBuilder' 'com.azure.messaging.servicebus.ServiceBusClientBuilder$ServiceBusSessionProcessorClientBuilder' com.azure.messaging.servicebus.ServiceBusReceiverClient com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient com.azure.messaging.servicebus.ServiceBusSenderAsyncClient","description":"Inspect receiver and processor APIs","initial_wait":60,"mode":"sync"}
```

### 🔧 25. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: src/main/java/com/example/orders/model/OrderStatus.java\n+package com.example.orders.model;\n+\n+import com.fasterxml.jackson.annotation.JsonCreator;\n+import com.fasterxml.jackson.annotation.JsonValue;\n+\n+import java.util.Locale;\n+\n+public enum OrderStatus {\n+    PENDING,\n+    PROCESSING,\n+    COMPLETED,\n+    FAILED;\n+\n+    @JsonCreator\n+    public static OrderStatus fromJson(String value) {\n+        return OrderStatus.valueOf(value.toUpperCase(Locale.ROOT));\n+    }\n+\n+    @JsonValue\n+    public String toJson() {\n+        return name().toLowerCase(Locale.ROOT);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/model/Order.java\n+package com.example.orders.model;\n+\n+import java.math.BigDecimal;\n+import java.util.Objects;\n+\n+public class Order {\n+    private String orderId;\n+    private String customerName;\n+    private String product;\n+    private int quantity;\n+    private BigDecimal totalPrice;\n+    private OrderStatus status;\n+\n+    public Order() {\n+    }\n+\n+    public Order(\n+        String orderId,\n+        String customerName,\n+        String product,\n+        int quantity,\n+        BigDecimal totalPrice,\n+        OrderStatus status\n+    ) {\n+        this.orderId = requireText(orderId, \"orderId\");\n+        this.customerName = requireText(customerName, \"customerName\");\n+        this.product = requireText(product, \"product\");\n+        if (quantity \u003c= 0) {\n+            throw new IllegalArgumentException(\"quantity must be positive\");\n+        }\n+        this.quantity = quantity;\n+        this.totalPrice = Objects.requireNonNull(totalPrice, \"totalPrice\");\n+        if (totalPrice.signum() \u003c 0) {\n+            throw new IllegalArgumentException(\"totalPrice must not be negative\");\n+        }\n+        this.status = Objects.requireNonNull(status, \"status\");\n+    }\n+\n+    private static String requireText(String value, String fieldName) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(fieldName + \" must not be blank\");\n+        }\n+        return value;\n+    }\n+\n+    public String getOrderId() {\n+        return orderId;\n+    }\n+\n+    public void setOrderId(String orderId) {\n+        this.orderId = orderId;\n+    }\n+\n+    public String getCustomerName() {\n+        return customerName;\n+    }\n+\n+    public void setCustomerName(String customerName) {\n+        this.customerName = customerName;\n+    }\n+\n+    public String getProduct() {\n+        return product;\n+    }\n+\n+    public void setProduct(String product) {\n+        this.product = product;\n+    }\n+\n+    public int getQuantity() {\n+        return quantity;\n+    }\n+\n+    public void setQuantity(int quantity) {\n+        this.quantity = quantity;\n+    }\n+\n+    public BigDecimal getTotalPrice() {\n+        return totalPrice;\n+    }\n+\n+    public void setTotalPrice(BigDecimal totalPrice) {\n+        this.totalPrice = totalPrice;\n+    }\n+\n+    public OrderStatus getStatus() {\n+        return status;\n+    }\n+\n+    public void setStatus(OrderStatus status) {\n+        this.status = status;\n+    }\n+\n+    @Override\n+    public String toString() {\n+        return \"Order{\"\n+            + \"orderId='\" + orderId + '\\''\n+            + \", customerName='\" + customerName + '\\''\n+            + \", product='\" + product + '\\''\n+            + \", quantity=\" + quantity\n+            + \", totalPrice=\" + totalPrice\n+            + \", status=\" + status\n+            + '}';\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/codec/OrderJsonCodec.java\n+package com.example.orders.codec;\n+\n+import com.example.orders.model.Order;\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.DeserializationFeature;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+public final class OrderJsonCodec {\n+    private final ObjectMapper objectMapper;\n+\n+    public OrderJsonCodec() {\n+        this.objectMapper = new ObjectMapper()\n+            .enable(DeserializationFeature.FAIL_ON_NULL_CREATOR_PROPERTIES);\n+    }\n+\n+    public String serialize(Order order) {\n+        try {\n+            return objectMapper.writeValueAsString(order);\n+        } catch (JsonProcessingException exception) {\n+            throw new OrderSerializationException(\"Could not serialize order\", exception);\n+        }\n+    }\n+\n+    public Order deserialize(String json) {\n+        try {\n+            return objectMapper.readValue(json, Order.class);\n+        } catch (JsonProcessingException exception) {\n+            throw new OrderSerializationException(\"Could not deserialize order\", exception);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/codec/OrderSerializationException.java\n+package com.example.orders.codec;\n+\n+public final class OrderSerializationException extends RuntimeException {\n+    public OrderSerializationException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/messaging/OrderMessageFactory.java\n+package com.example.orders.messaging;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.example.orders.codec.OrderJsonCodec;\n+import com.example.orders.model.Order;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.time.ZoneOffset;\n+import java.util.Objects;\n+\n+public final class OrderMessageFactory {\n+    public static final String PRIORITY_PROPERTY = \"priority\";\n+    public static final String HIGH_PRIORITY = \"high\";\n+    public static final String NORMAL_PRIORITY = \"normal\";\n+    public static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);\n+\n+    private final OrderJsonCodec codec;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public OrderMessageFactory(OrderJsonCodec codec, BigDecimal highPriorityThreshold) {\n+        this.codec = Objects.requireNonNull(codec, \"codec\");\n+        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, \"highPriorityThreshold\");\n+    }\n+\n+    public ServiceBusMessage createMessage(Order order) {\n+        Objects.requireNonNull(order, \"order\");\n+        boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+\n+        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(codec.serialize(order)))\n+            .setMessageId(order.getOrderId())\n+            .setCorrelationId(order.getOrderId())\n+            .setSessionId(order.getCustomerName())\n+            .setContentType(\"application/json\");\n+        message.getApplicationProperties().put(\n+            PRIORITY_PROPERTY,\n+            highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY\n+        );\n+\n+        if (highPriority) {\n+            message.setScheduledEnqueueTime(OffsetDateTime.now(ZoneOffset.UTC).plus(FRAUD_REVIEW_DELAY));\n+        }\n+        return message;\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/messaging/SyncOrderSender.java\n+package com.example.orders.messaging;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.example.orders.model.Order;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class SyncOrderSender implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderSender.class);\n+\n+    private final ServiceBusSenderClient sender;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public SyncOrderSender(\n+        String fullyQualifiedNamespace,\n+        String queueName,\n+        TokenCredential credential,\n+        BigDecimal highPriorityThreshold\n+    ) {\n+        this.sender = new ServiceBusClientBuilder()\n+            .credential(fullyQualifiedNamespace, credential)\n+            .sender()\n+            .queueName(queueName)\n+            .buildClient();\n+        this.messageFactory = new OrderMessageFactory(\n+            new com.example.orders.codec.OrderJsonCodec(),\n+            highPriorityThreshold\n+        );\n+    }\n+\n+    public void sendOrder(Order order) {\n+        ServiceBusMessage message = messageFactory.createMessage(order);\n+        sender.sendMessage(message);\n+        LOGGER.info(\"Sent order {} with priority {}\", order.getOrderId(),\n+            message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));\n+    }\n+\n+    public void sendOrders(List\u003cOrder\u003e orders) {\n+        Objects.requireNonNull(orders, \"orders\");\n+        if (orders.isEmpty()) {\n+            return;\n+        }\n+\n+        ServiceBusMessageBatch batch = sender.createMessageBatch();\n+        for (Order order : orders) {\n+            ServiceBusMessage message = messageFactory.createMessage(order);\n+            if (batch.tryAddMessage(message)) {\n+                continue;\n+            }\n+\n+            sendBatch(batch);\n+            batch = sender.createMessageBatch();\n+            if (!batch.tryAddMessage(message)) {\n+                throw new IllegalArgumentException(\n+                    \"Order \" + order.getOrderId() + \" exceeds the maximum Service Bus message size\"\n+                );\n+            }\n+        }\n+\n+        if (batch.getCount() \u003e 0) {\n+            sendBatch(batch);\n+        }\n+    }\n+\n+    private void sendBatch(ServiceBusMessageBatch batch) {\n+        sender.sendMessages(batch);\n+        LOGGER.info(\"Sent batch containing {} order(s) and {} byte(s)\",\n+            batch.getCount(), batch.getSizeInBytes());\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/messaging/AsyncOrderSender.java\n+package com.example.orders.messaging;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import com.example.orders.codec.OrderJsonCodec;\n+import com.example.orders.model.Order;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class AsyncOrderSender implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderSender.class);\n+\n+    private final ServiceBusSenderAsyncClient sender;\n+    private final OrderMessageFactory messageFactory;\n+\n+    public AsyncOrderSender(\n+        String fullyQualifiedNamespace,\n+        String queueName,\n+        TokenCredential credential,\n+        BigDecimal highPriorityThreshold\n+    ) {\n+        this.sender = new ServiceBusClientBuilder()\n+            .credential(fullyQualifiedNamespace, credential)\n+            .sender()\n+            .queueName(queueName)\n+            .buildAsyncClient();\n+        this.messageFactory = new OrderMessageFactory(new OrderJsonCodec(), highPriorityThreshold);\n+    }\n+\n+    public Mono\u003cVoid\u003e sendOrder(Order order) {\n+        ServiceBusMessage message = messageFactory.createMessage(order);\n+        return sender.sendMessage(message)\n+            .doOnSuccess(ignored -\u003e LOGGER.info(\"Sent order {} with priority {}\", order.getOrderId(),\n+                message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY)));\n+    }\n+\n+    public Mono\u003cVoid\u003e sendOrders(List\u003cOrder\u003e orders) {\n+        Objects.requireNonNull(orders, \"orders\");\n+        if (orders.isEmpty()) {\n+            return Mono.empty();\n+        }\n+\n+        List\u003cServiceBusMessage\u003e messages = orders.stream()\n+            .map(messageFactory::createMessage)\n+            .toList();\n+        return sender.createMessageBatch()\n+            .flatMap(batch -\u003e fillAndSend(messages, 0, batch));\n+    }\n+\n+    private Mono\u003cVoid\u003e fillAndSend(\n+        List\u003cServiceBusMessage\u003e messages,\n+        int startIndex,\n+        ServiceBusMessageBatch batch\n+    ) {\n+        int index = startIndex;\n+        while (index \u003c messages.size() \u0026\u0026 batch.tryAddMessage(messages.get(index))) {\n+            index++;\n+        }\n+\n+        if (index == startIndex) {\n+            return Mono.error(new IllegalArgumentException(\n+                \"Order message exceeds the maximum Service Bus message size\"\n+            ));\n+        }\n+\n+        int nextIndex = index;\n+        Mono\u003cVoid\u003e sendCurrentBatch = sender.sendMessages(batch)\n+            .doOnSuccess(ignored -\u003e LOGGER.info(\"Sent batch containing {} order(s) and {} byte(s)\",\n+                batch.getCount(), batch.getSizeInBytes()));\n+\n+        if (nextIndex == messages.size()) {\n+            return sendCurrentBatch;\n+        }\n+\n+        return sendCurrentBatch\n+            .then(sender.createMessageBatch())\n+            .flatMap(nextBatch -\u003e fillAndSend(messages, nextIndex, nextBatch));\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/processing/OrderProcessingSupport.java\n+package com.example.orders.processing;\n+\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n+import com.example.orders.codec.OrderJsonCodec;\n+import com.example.orders.model.Order;\n+import com.example.orders.model.OrderStatus;\n+import org.slf4j.Logger;\n+\n+final class OrderProcessingSupport {\n+    private static final int MAX_DEAD_LETTER_DESCRIPTION_LENGTH = 4_096;\n+\n+    private OrderProcessingSupport() {\n+    }\n+\n+    static Order process(ServiceBusReceivedMessage message, OrderJsonCodec codec, Logger logger) {\n+        Order order = codec.deserialize(message.getBody().toString());\n+        if (order.getStatus() == OrderStatus.FAILED) {\n+            throw new IllegalStateException(\"Order is marked failed\");\n+        }\n+\n+        order.setStatus(OrderStatus.PROCESSING);\n+        logger.info(\"Processing order {} for customer {}\", order.getOrderId(), order.getCustomerName());\n+        order.setStatus(OrderStatus.COMPLETED);\n+        logger.info(\"Completed order {}\", order.getOrderId());\n+        return order;\n+    }\n+\n+    static DeadLetterOptions deadLetterOptions(RuntimeException exception) {\n+        String description = exception.getMessage() == null\n+            ? exception.getClass().getName()\n+            : exception.getMessage();\n+        if (description.length() \u003e MAX_DEAD_LETTER_DESCRIPTION_LENGTH) {\n+            description = description.substring(0, MAX_DEAD_LETTER_DESCRIPTION_LENGTH);\n+        }\n+        return new DeadLetterOptions()\n+            .setDeadLetterReason(exception.getClass().getSimpleName())\n+            .setDeadLetterErrorDescription(description);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/processing/SyncOrderProcessor.java\n+package com.example.orders.processing;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import com.example.orders.codec.OrderJsonCodec;\n+import com.example.orders.model.Order;\n+import com.example.orders.model.OrderStatus;\n+import com.example.orders.messaging.SyncOrderSender;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.time.Duration;\n+\n+public final class SyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);\n+    private static final int MAX_MESSAGES_PER_SESSION = 10_000;\n+\n+    private final ServiceBusSessionReceiverClient sessionReceiver;\n+    private final ServiceBusSessionReceiverClient deadLetterSessionReceiver;\n+    private final OrderJsonCodec codec = new OrderJsonCodec();\n+\n+    public SyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {\n+        this.sessionReceiver = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sessionReceiver()\n+            .queueName(queueName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .buildClient();\n+        this.deadLetterSessionReceiver = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sessionReceiver()\n+            .queueName(queueName)\n+            .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .buildClient();\n+    }\n+\n+    public void processSessions(int sessionCount, Duration idleTimeout) {\n+        for (int index = 0; index \u003c sessionCount; index++) {\n+            try (ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {\n+                LOGGER.info(\"Accepted customer session {}\", receiver.getSessionId());\n+                for (ServiceBusReceivedMessage message\n+                    : receiver.receiveMessages(MAX_MESSAGES_PER_SESSION, idleTimeout)) {\n+                    processMessage(receiver, message);\n+                }\n+            }\n+        }\n+    }\n+\n+    public void inspectAndReprocessDeadLetters(\n+        int sessionCount,\n+        Duration idleTimeout,\n+        SyncOrderSender sender\n+    ) {\n+        for (int index = 0; index \u003c sessionCount; index++) {\n+            try (ServiceBusReceiverClient receiver = deadLetterSessionReceiver.acceptNextSession()) {\n+                for (ServiceBusReceivedMessage message\n+                    : receiver.receiveMessages(MAX_MESSAGES_PER_SESSION, idleTimeout)) {\n+                    reprocessDeadLetter(receiver, message, sender);\n+                }\n+            }\n+        }\n+    }\n+\n+    private void processMessage(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {\n+        try {\n+            OrderProcessingSupport.process(message, codec, LOGGER);\n+            receiver.complete(message);\n+        } catch (RuntimeException exception) {\n+            LOGGER.error(\"Dead-lettering message {}: {}\", message.getMessageId(), exception.getMessage());\n+            receiver.deadLetter(message, OrderProcessingSupport.deadLetterOptions(exception));\n+        }\n+    }\n+\n+    private void reprocessDeadLetter(\n+        ServiceBusReceiverClient receiver,\n+        ServiceBusReceivedMessage message,\n+        SyncOrderSender sender\n+    ) {\n+        LOGGER.warn(\"Inspecting dead-lettered message {}: reason={}, description={}\",\n+            message.getMessageId(), message.getDeadLetterReason(), message.getDeadLetterErrorDescription());\n+        try {\n+            Order order = codec.deserialize(message.getBody().toString());\n+            order.setStatus(OrderStatus.PENDING);\n+            sender.sendOrder(order);\n+            receiver.complete(message);\n+            LOGGER.info(\"Re-enqueued dead-lettered order {}\", order.getOrderId());\n+        } catch (RuntimeException exception) {\n+            LOGGER.error(\"Message {} could not be reprocessed and remains in the dead-letter queue\",\n+                message.getMessageId(), exception);\n+            receiver.abandon(message);\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        deadLetterSessionReceiver.close();\n+        sessionReceiver.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/processing/AsyncOrderProcessor.java\n+package com.example.orders.processing;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import com.example.orders.codec.OrderJsonCodec;\n+import com.example.orders.messaging.AsyncOrderSender;\n+import com.example.orders.model.Order;\n+import com.example.orders.model.OrderStatus;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.concurrent.TimeoutException;\n+\n+public final class AsyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);\n+\n+    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;\n+    private final ServiceBusSessionReceiverAsyncClient deadLetterSessionReceiver;\n+    private final OrderJsonCodec codec = new OrderJsonCodec();\n+\n+    public AsyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {\n+        this.sessionReceiver = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sessionReceiver()\n+            .queueName(queueName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .buildAsyncClient();\n+        this.deadLetterSessionReceiver = new ServiceBusClientBuilder()\n+            .credential(namespace, credential)\n+            .sessionReceiver()\n+            .queueName(queueName)\n+            .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .buildAsyncClient();\n+    }\n+\n+    public Mono\u003cVoid\u003e processSessions(int sessionCount, Duration idleTimeout) {\n+        return Flux.range(0, sessionCount)\n+            .concatMap(ignored -\u003e sessionReceiver.acceptNextSession()\n+                .flatMap(receiver -\u003e processSession(receiver, idleTimeout)))\n+            .then();\n+    }\n+\n+    public Mono\u003cVoid\u003e inspectAndReprocessDeadLetters(\n+        int sessionCount,\n+        Duration idleTimeout,\n+        AsyncOrderSender sender\n+    ) {\n+        return Flux.range(0, sessionCount)\n+            .concatMap(ignored -\u003e deadLetterSessionReceiver.acceptNextSession()\n+                .flatMap(receiver -\u003e inspectSession(receiver, idleTimeout, sender)))\n+            .then();\n+    }\n+\n+    private Mono\u003cVoid\u003e processSession(\n+        ServiceBusReceiverAsyncClient receiver,\n+        Duration idleTimeout\n+    ) {\n+        LOGGER.info(\"Accepted customer session {}\", receiver.getSessionId());\n+        return withIdleTimeout(\n+            receiver.receiveMessages().concatMap(message -\u003e processMessage(receiver, message)),\n+            idleTimeout\n+        ).then().doFinally(ignored -\u003e receiver.close());\n+    }\n+\n+    private Mono\u003cVoid\u003e processMessage(\n+        ServiceBusReceiverAsyncClient receiver,\n+        ServiceBusReceivedMessage message\n+    ) {\n+        try {\n+            OrderProcessingSupport.process(message, codec, LOGGER);\n+            return receiver.complete(message);\n+        } catch (RuntimeException exception) {\n+            LOGGER.error(\"Dead-lettering message {}: {}\", message.getMessageId(), exception.getMessage());\n+            return receiver.deadLetter(message, OrderProcessingSupport.deadLetterOptions(exception));\n+        }\n+    }\n+\n+    private Mono\u003cVoid\u003e inspectSession(\n+        ServiceBusReceiverAsyncClient receiver,\n+        Duration idleTimeout,\n+        AsyncOrderSender sender\n+    ) {\n+        return withIdleTimeout(\n+            receiver.receiveMessages().concatMap(message -\u003e reprocessDeadLetter(receiver, message, sender)),\n+            idleTimeout\n+        ).then().doFinally(ignored -\u003e receiver.close());\n+    }\n+\n+    private Mono\u003cVoid\u003e reprocessDeadLetter(\n+        ServiceBusReceiverAsyncClient receiver,\n+        ServiceBusReceivedMessage message,\n+        AsyncOrderSender sender\n+    ) {\n+        LOGGER.warn(\"Inspecting dead-lettered message {}: reason={}, description={}\",\n+            message.getMessageId(), message.getDeadLetterReason(), message.getDeadLetterErrorDescription());\n+\n+        final Order order;\n+        try {\n+            order = codec.deserialize(message.getBody().toString());\n+            order.setStatus(OrderStatus.PENDING);\n+        } catch (RuntimeException exception) {\n+            LOGGER.error(\"Message {} cannot be deserialized and remains in the dead-letter queue\",\n+                message.getMessageId(), exception);\n+            return receiver.abandon(message);\n+        }\n+\n+        return sender.sendOrder(order)\n+            .then(receiver.complete(message))\n+            .doOnSuccess(ignored -\u003e LOGGER.info(\"Re-enqueued dead-lettered order {}\", order.getOrderId()))\n+            .onErrorResume(exception -\u003e receiver.abandon(message)\n+                .then(Mono.error(new IllegalStateException(\n+                    \"Could not reprocess dead-lettered order \" + order.getOrderId(),\n+                    exception\n+                ))));\n+    }\n+\n+    private Flux\u003cVoid\u003e withIdleTimeout(Flux\u003cVoid\u003e messages, Duration idleTimeout) {\n+        return messages.timeout(idleTimeout)\n+            .onErrorResume(TimeoutException.class, ignored -\u003e Flux.empty());\n+    }\n+\n+    @Override\n+    public void close() {\n+        deadLetterSessionReceiver.close();\n+        sessionReceiver.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Main.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.example.orders.messaging.AsyncOrderSender;\n+import com.example.orders.messaging.SyncOrderSender;\n+import com.example.orders.model.Order;\n+import com.example.orders.model.OrderStatus;\n+import com.example.orders.processing.AsyncOrderProcessor;\n+import com.example.orders.processing.SyncOrderProcessor;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.util.List;\n+\n+public final class Main {\n+    private static final Duration SESSION_IDLE_TIMEOUT = Duration.ofSeconds(5);\n+    private static final int CUSTOMER_SESSION_COUNT = 3;\n+    private static final int DEAD_LETTER_SESSION_COUNT = 1;\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String namespace = requiredEnvironmentVariable(\"SERVICE_BUS_NAMESPACE\");\n+        String queueName = requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+        String fullyQualifiedNamespace = namespace.contains(\".\")\n+            ? namespace\n+            : namespace + \".servicebus.windows.net\";\n+        BigDecimal highPriorityThreshold = new BigDecimal(\n+            System.getenv().getOrDefault(\"ORDER_HIGH_PRIORITY_THRESHOLD\", \"1000.00\")\n+        );\n+        TokenCredential credential = managedIdentityCredential();\n+\n+        runSyncDemo(fullyQualifiedNamespace, queueName, credential, highPriorityThreshold);\n+        runAsyncDemo(fullyQualifiedNamespace, queueName, credential, highPriorityThreshold);\n+    }\n+\n+    private static void runSyncDemo(\n+        String namespace,\n+        String queueName,\n+        TokenCredential credential,\n+        BigDecimal highPriorityThreshold\n+    ) {\n+        try (\n+            SyncOrderSender sender = new SyncOrderSender(\n+                namespace, queueName, credential, highPriorityThreshold\n+            );\n+            SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential)\n+        ) {\n+            sender.sendOrders(demoOrders(\"sync\"));\n+            processor.processSessions(CUSTOMER_SESSION_COUNT, SESSION_IDLE_TIMEOUT);\n+            processor.inspectAndReprocessDeadLetters(\n+                DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT, sender\n+            );\n+            processor.processSessions(DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT);\n+        }\n+    }\n+\n+    private static void runAsyncDemo(\n+        String namespace,\n+        String queueName,\n+        TokenCredential credential,\n+        BigDecimal highPriorityThreshold\n+    ) {\n+        try (\n+            AsyncOrderSender sender = new AsyncOrderSender(\n+                namespace, queueName, credential, highPriorityThreshold\n+            );\n+            AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential)\n+        ) {\n+            sender.sendOrders(demoOrders(\"async\"))\n+                .then(processor.processSessions(CUSTOMER_SESSION_COUNT, SESSION_IDLE_TIMEOUT))\n+                .then(processor.inspectAndReprocessDeadLetters(\n+                    DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT, sender\n+                ))\n+                .then(processor.processSessions(DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT))\n+                .block();\n+        }\n+    }\n+\n+    private static List\u003cOrder\u003e demoOrders(String prefix) {\n+        return List.of(\n+            new Order(prefix + \"-001\", prefix + \"-alice\", \"Keyboard\", 1,\n+                new BigDecimal(\"89.99\"), OrderStatus.PENDING),\n+            new Order(prefix + \"-002\", prefix + \"-alice\", \"Mouse\", 2,\n+                new BigDecimal(\"119.98\"), OrderStatus.PENDING),\n+            new Order(prefix + \"-003\", prefix + \"-bob\", \"Monitor\", 1,\n+                new BigDecimal(\"399.00\"), OrderStatus.FAILED),\n+            new Order(prefix + \"-004\", prefix + \"-carol\", \"Workstation\", 1,\n+                new BigDecimal(\"2499.00\"), OrderStatus.PENDING)\n+        );\n+    }\n+\n+    private static TokenCredential managedIdentityCredential() {\n+        String clientId = System.getenv(\"AZURE_CLIENT_ID\");\n+        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();\n+        return clientId == null || clientId.isBlank()\n+            ? builder.build()\n+            : builder.clientId(clientId).build();\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Environment variable \" + name + \" is required\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/test/java/com/example/orders/codec/OrderJsonCodecTest.java\n+package com.example.orders.codec;\n+\n+import com.example.orders.model.Order;\n+import com.example.orders.model.OrderStatus;\n+import org.junit.jupiter.api.Test;\n+\n+import java.math.BigDecimal;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class OrderJsonCodecTest {\n+    private final OrderJsonCodec codec = new OrderJsonCodec();\n+\n+    @Test\n+    void roundTripsOrderAsJson() {\n+        Order original = new Order(\n+            \"order-1\",\n+            \"Ada\",\n+            \"Laptop\",\n+            2,\n+            new BigDecimal(\"2400.50\"),\n+            OrderStatus.PENDING\n+        );\n+\n+        Order restored = codec.deserialize(codec.serialize(original));\n+\n+        assertEquals(original.getOrderId(), restored.getOrderId());\n+        assertEquals(original.getCustomerName(), restored.getCustomerName());\n+        assertEquals(original.getProduct(), restored.getProduct());\n+        assertEquals(original.getQuantity(), restored.getQuantity());\n+        assertEquals(original.getTotalPrice(), restored.getTotalPrice());\n+        assertEquals(original.getStatus(), restored.getStatus());\n+    }\n+}\n*** Add File: src/test/java/com/example/orders/messaging/OrderMessageFactoryTest.java\n+package com.example.orders.messaging;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.example.orders.codec.OrderJsonCodec;\n+import com.example.orders.model.Order;\n+import com.example.orders.model.OrderStatus;\n+import org.junit.jupiter.api.Test;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+import java.time.ZoneOffset;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertNotNull;\n+import static org.junit.jupiter.api.Assertions.assertNull;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class OrderMessageFactoryTest {\n+    private final OrderMessageFactory factory =\n+        new OrderMessageFactory(new OrderJsonCodec(), new BigDecimal(\"1000.00\"));\n+\n+    @Test\n+    void mapsOrderMetadataToServiceBusMessage() {\n+        Order order = order(\"order-1\", \"Ada\", \"999.99\");\n+\n+        ServiceBusMessage message = factory.createMessage(order);\n+\n+        assertEquals(\"order-1\", message.getMessageId());\n+        assertEquals(\"order-1\", message.getCorrelationId());\n+        assertEquals(\"Ada\", message.getSessionId());\n+        assertEquals(\"application/json\", message.getContentType());\n+        assertEquals(OrderMessageFactory.NORMAL_PRIORITY,\n+            message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));\n+        assertNull(message.getScheduledEnqueueTime());\n+    }\n+\n+    @Test\n+    void schedulesHighPriorityOrderForFraudReview() {\n+        OffsetDateTime before = OffsetDateTime.now(ZoneOffset.UTC)\n+            .plus(OrderMessageFactory.FRAUD_REVIEW_DELAY);\n+\n+        ServiceBusMessage message = factory.createMessage(order(\"order-2\", \"Grace\", \"1000.01\"));\n+\n+        OffsetDateTime after = OffsetDateTime.now(ZoneOffset.UTC)\n+            .plus(OrderMessageFactory.FRAUD_REVIEW_DELAY);\n+        assertEquals(OrderMessageFactory.HIGH_PRIORITY,\n+            message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));\n+        assertNotNull(message.getScheduledEnqueueTime());\n+        assertTrue(!message.getScheduledEnqueueTime().isBefore(before));\n+        assertTrue(!message.getScheduledEnqueueTime().isAfter(after));\n+    }\n+\n+    private static Order order(String orderId, String customer, String price) {\n+        return new Order(\n+            orderId,\n+            customer,\n+            \"Product\",\n+            1,\n+            new BigDecimal(price),\n+            OrderStatus.PENDING\n+        );\n+    }\n+}\n*** Add File: README.md\n+# Azure Service Bus Order Processor\n+\n+Small Java 17 Maven sample with synchronous and Reactor-based asynchronous senders and\n+processors for an Azure Service Bus queue.\n+\n+## Prerequisites\n+\n+- Java 17 and Maven 3.9+\n+- An Azure Service Bus queue with **sessions enabled**\n+- A managed identity with the `Azure Service Bus Data Sender` and\n+  `Azure Service Bus Data Receiver` roles scoped to the namespace or queue\n+\n+The customer name is used as the Service Bus session ID. A receiver holds an exclusive\n+session lock and processes all messages in that session serially, so one customer's orders\n+cannot be interleaved with another customer's orders.\n+\n+## Configuration\n+\n+Set these environment variables:\n+\n+```text\n+SERVICE_BUS_NAMESPACE=\u003cnamespace-name-or-fully-qualified-namespace\u003e\n+SERVICE_BUS_QUEUE_NAME=\u003csession-enabled-queue-name\u003e\n+ORDER_HIGH_PRIORITY_THRESHOLD=1000.00\n+AZURE_CLIENT_ID=\u003coptional-user-assigned-managed-identity-client-id\u003e\n+```\n+\n+If `AZURE_CLIENT_ID` is omitted, the application uses the system-assigned managed identity.\n+No connection string or account key is used.\n+\n+Orders above `ORDER_HIGH_PRIORITY_THRESHOLD` receive a `priority=high` application property\n+and a scheduled enqueue time 30 seconds in the future for fraud review. Other orders receive\n+`priority=normal`. All messages use the order ID as both the message ID and correlation ID.\n+\n+## Build and run\n+\n+```text\n+mvn clean verify\n+mvn exec:java\n+```\n+\n+`Main` runs the synchronous cycle first and then the asynchronous cycle. Each cycle sends a\n+deliberately failed order, dead-letters it with a reason and description, reads it from the\n+dead-letter subqueue, resets it to `pending`, re-enqueues it, and processes it successfully.\n+\n+Because the demo includes a scheduled high-value order, each cycle can wait up to 30 seconds\n+for that order to become available.\n+\n+## References\n+\n+- [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable)\n+- [Service Bus message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions)\n+- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)\n*** End Patch\n"
```

### 🔧 26. powershell

**Input:**

```json
{"command":"mvn clean verify","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

### 🔧 27. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/processing/AsyncOrderProcessor.java\n@@\n         Duration idleTimeout\n     ) {\n         LOGGER.info(\"Accepted customer session {}\", receiver.getSessionId());\n         return withIdleTimeout(\n-            receiver.receiveMessages().concatMap(message -\u003e processMessage(receiver, message)),\n+            receiver.receiveMessages()\n+                .concatMap(message -\u003e processMessage(receiver, message).thenReturn(message)),\n             idleTimeout\n         ).then().doFinally(ignored -\u003e receiver.close());\n@@\n     ) {\n         return withIdleTimeout(\n-            receiver.receiveMessages().concatMap(message -\u003e reprocessDeadLetter(receiver, message, sender)),\n+            receiver.receiveMessages()\n+                .concatMap(message -\u003e reprocessDeadLetter(receiver, message, sender).thenReturn(message)),\n             idleTimeout\n         ).then().doFinally(ignored -\u003e receiver.close());\n@@\n-    private Flux\u003cVoid\u003e withIdleTimeout(Flux\u003cVoid\u003e messages, Duration idleTimeout) {\n+    private \u003cT\u003e Flux\u003cT\u003e withIdleTimeout(Flux\u003cT\u003e messages, Duration idleTimeout) {\n         return messages.timeout(idleTimeout)\n             .onErrorResume(TimeoutException.class, ignored -\u003e Flux.empty());\n     }\n*** End Patch\n"
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\orders\Main.java`
- `src\main\java\com\example\orders\codec\OrderJsonCodec.java`
- `src\main\java\com\example\orders\codec\OrderSerializationException.java`
- `src\main\java\com\example\orders\messaging\AsyncOrderSender.java`
- `src\main\java\com\example\orders\messaging\OrderMessageFactory.java`
- `src\main\java\com\example\orders\messaging\SyncOrderSender.java`
- `src\main\java\com\example\orders\model\Order.java`
- `src\main\java\com\example\orders\model\OrderStatus.java`
- `src\main\java\com\example\orders\processing\AsyncOrderProcessor.java`
- `src\main\java\com\example\orders\processing\OrderProcessingSupport.java`
- `src\main\java\com\example\orders\processing\SyncOrderProcessor.java`
- `src\test\java\com\example\orders\codec\OrderJsonCodecTest.java`
- `src\test\java\com\example\orders\messaging\OrderMessageFactoryTest.java`

## Copilot Response

I’m grounding the implementation in the current Java SDK APIs, then I’ll build the project and validate it locally without contacting Azure.

The core design will use Service Bus sessions keyed by customer name; that is the broker feature that enforces per-customer ordering while allowing different customers to be processed concurrently. I’m checking the exact processor and settlement APIs before writing the classes.

I’ve confirmed scheduled enqueue time can be set directly on each message, so high-priority orders can still participate in size-aware batches. I’m creating the Maven skeleton next, then I’ll compile against the SDK to catch any API mismatches.

The full project is in place. I’m at the compile-and-test pass now; any SDK type or Reactor lifecycle issues will be corrected before handoff.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Consistent AutoCloseable ownership for sender and processor wrappers.
- Both synchronous and asynchronous per-session receivers are closed after processing.

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
hyoka run --prompt-id service-bus-dp-java-order-processor --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)

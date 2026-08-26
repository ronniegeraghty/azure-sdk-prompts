# Evaluation Report: service-bus-dp-java-order-processor

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 498.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-order-processor` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 20/24 |
| Duration | 498.9s |
| Timestamp | 2026-08-26T13:30:13Z |
| Files Generated | 10 |
| Event Count | 16431 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 222.1s |
| Review | 275.5s |
| **Total** | **498.9s** |

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
| Token Usage | in=345264 out=11539 |
| Turn Count | 11 |

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

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"pattern":"*"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eservice-bus-order-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.2.35\u003c/azure.sdk.bom.version\u003e\n+        \u003cjackson.version\u003e2.17.2\u003c/jackson.version\u003e\n+        \u003cslf4j.version\u003e2.0.16\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.13.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.4.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.orders.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/orders/Order.java\n+package com.example.orders;\n+\n+import com.fasterxml.jackson.annotation.JsonCreator;\n+import com.fasterxml.jackson.annotation.JsonProperty;\n+\n+import java.math.BigDecimal;\n+import java.util.Objects;\n+\n+public final class Order {\n+    public enum Status {\n+        PENDING,\n+        PROCESSING,\n+        COMPLETED,\n+        FAILED\n+    }\n+\n+    private final String orderId;\n+    private final String customerName;\n+    private final String product;\n+    private final int quantity;\n+    private final BigDecimal totalPrice;\n+    private final Status status;\n+\n+    @JsonCreator\n+    public Order(\n+            @JsonProperty(\"orderId\") String orderId,\n+            @JsonProperty(\"customerName\") String customerName,\n+            @JsonProperty(\"product\") String product,\n+            @JsonProperty(\"quantity\") int quantity,\n+            @JsonProperty(\"totalPrice\") BigDecimal totalPrice,\n+            @JsonProperty(\"status\") Status status) {\n+        this.orderId = Objects.requireNonNull(orderId, \"orderId\");\n+        this.customerName = Objects.requireNonNull(customerName, \"customerName\");\n+        this.product = Objects.requireNonNull(product, \"product\");\n+        this.quantity = quantity;\n+        this.totalPrice = Objects.requireNonNull(totalPrice, \"totalPrice\");\n+        this.status = Objects.requireNonNull(status, \"status\");\n+    }\n+\n+    public String getOrderId() {\n+        return orderId;\n+    }\n+\n+    public String getCustomerName() {\n+        return customerName;\n+    }\n+\n+    public String getProduct() {\n+        return product;\n+    }\n+\n+    public int getQuantity() {\n+        return quantity;\n+    }\n+\n+    public BigDecimal getTotalPrice() {\n+        return totalPrice;\n+    }\n+\n+    public Status getStatus() {\n+        return status;\n+    }\n+\n+    public Order withStatus(Status newStatus) {\n+        return new Order(orderId, customerName, product, quantity, totalPrice, newStatus);\n+    }\n+\n+    @Override\n+    public String toString() {\n+        return \"Order{\" +\n+                \"orderId='\" + orderId + '\\'' +\n+                \", customerName='\" + customerName + '\\'' +\n+                \", product='\" + product + '\\'' +\n+                \", quantity=\" + quantity +\n+                \", totalPrice=\" + totalPrice +\n+                \", status=\" + status +\n+                '}';\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderJson.java\n+package com.example.orders;\n+\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+final class OrderJson {\n+    private static final ObjectMapper MAPPER = new ObjectMapper();\n+\n+    private OrderJson() {\n+    }\n+\n+    static String serialize(Order order) {\n+        try {\n+            return MAPPER.writeValueAsString(order);\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Could not serialize order \" + order.getOrderId(), exception);\n+        }\n+    }\n+\n+    static Order deserialize(String json) throws JsonProcessingException {\n+        return MAPPER.readValue(json, Order.class);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/OrderMessageFactory.java\n+package com.example.orders;\n+\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+\n+import java.math.BigDecimal;\n+\n+final class OrderMessageFactory {\n+    static final String PRIORITY_PROPERTY = \"priority\";\n+    static final String HIGH_PRIORITY = \"high\";\n+    static final String NORMAL_PRIORITY = \"normal\";\n+\n+    private OrderMessageFactory() {\n+    }\n+\n+    static ServiceBusMessage create(Order order, BigDecimal highPriorityThreshold) {\n+        boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+        return new ServiceBusMessage(OrderJson.serialize(order))\n+                .setContentType(\"application/json\")\n+                .setMessageId(order.getOrderId())\n+                .setCorrelationId(order.getOrderId())\n+                .setSessionId(order.getCustomerName())\n+                .setSubject(\"order\")\n+                .setApplicationProperties(\n+                        java.util.Map.of(PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY));\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class SyncOrderSender implements AutoCloseable {\n+    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;\n+\n+    private final ServiceBusSenderClient sender;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public SyncOrderSender(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildClient();\n+        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, \"highPriorityThreshold\");\n+    }\n+\n+    public void send(Order order) {\n+        ServiceBusMessage message = toMessage(order);\n+        if (isHighPriority(order)) {\n+            sender.scheduleMessage(message, OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));\n+        } else {\n+            sender.sendMessage(message);\n+        }\n+    }\n+\n+    public void sendBatch(List\u003cOrder\u003e orders) {\n+        List\u003cServiceBusMessage\u003e immediateMessages = new ArrayList\u003c\u003e();\n+        OffsetDateTime scheduledTime = OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS);\n+\n+        for (Order order : orders) {\n+            ServiceBusMessage message = toMessage(order);\n+            if (isHighPriority(order)) {\n+                sender.scheduleMessage(message, scheduledTime);\n+            } else {\n+                immediateMessages.add(message);\n+            }\n+        }\n+\n+        sendSizeAwareBatches(immediateMessages);\n+    }\n+\n+    private void sendSizeAwareBatches(List\u003cServiceBusMessage\u003e messages) {\n+        ServiceBusMessageBatch batch = sender.createMessageBatch();\n+        for (ServiceBusMessage message : messages) {\n+            if (!batch.tryAddMessage(message)) {\n+                if (batch.getCount() == 0) {\n+                    throw new IllegalArgumentException(\"Order message exceeds the Service Bus maximum message size\");\n+                }\n+                sender.sendMessages(batch);\n+                batch = sender.createMessageBatch();\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalArgumentException(\"Order message exceeds the Service Bus maximum message size\");\n+                }\n+            }\n+        }\n+        if (batch.getCount() \u003e 0) {\n+            sender.sendMessages(batch);\n+        }\n+    }\n+\n+    private ServiceBusMessage toMessage(Order order) {\n+        return OrderMessageFactory.create(order, highPriorityThreshold);\n+    }\n+\n+    private boolean isHighPriority(Order order) {\n+        return order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderSender.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.time.OffsetDateTime;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class AsyncOrderSender implements AutoCloseable {\n+    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;\n+\n+    private final ServiceBusSenderAsyncClient sender;\n+    private final BigDecimal highPriorityThreshold;\n+\n+    public AsyncOrderSender(\n+            String fullyQualifiedNamespace,\n+            String queueName,\n+            TokenCredential credential,\n+            BigDecimal highPriorityThreshold) {\n+        this.sender = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential)\n+                .sender()\n+                .queueName(queueName)\n+                .buildAsyncClient();\n+        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, \"highPriorityThreshold\");\n+    }\n+\n+    public Mono\u003cVoid\u003e send(Order order) {\n+        ServiceBusMessage message = toMessage(order);\n+        if (isHighPriority(order)) {\n+            return sender.scheduleMessage(\n+                            message, OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS))\n+                    .then();\n+        }\n+        return sender.sendMessage(message);\n+    }\n+\n+    public Mono\u003cVoid\u003e sendBatch(List\u003cOrder\u003e orders) {\n+        List\u003cServiceBusMessage\u003e immediateMessages = new ArrayList\u003c\u003e();\n+        List\u003cServiceBusMessage\u003e scheduledMessages = new ArrayList\u003c\u003e();\n+        for (Order order : orders) {\n+            (isHighPriority(order) ? scheduledMessages : immediateMessages).add(toMessage(order));\n+        }\n+\n+        OffsetDateTime scheduledTime = OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS);\n+        Mono\u003cVoid\u003e schedule = Flux.fromIterable(scheduledMessages)\n+                .concatMap(message -\u003e sender.scheduleMessage(message, scheduledTime))\n+                .then();\n+\n+        return schedule.then(sendSizeAwareBatches(immediateMessages, 0));\n+    }\n+\n+    private Mono\u003cVoid\u003e sendSizeAwareBatches(List\u003cServiceBusMessage\u003e messages, int startIndex) {\n+        if (startIndex \u003e= messages.size()) {\n+            return Mono.empty();\n+        }\n+\n+        return sender.createMessageBatch().flatMap(batch -\u003e {\n+            int nextIndex = fillBatch(batch, messages, startIndex);\n+            return sender.sendMessages(batch).then(sendSizeAwareBatches(messages, nextIndex));\n+        });\n+    }\n+\n+    private int fillBatch(ServiceBusMessageBatch batch, List\u003cServiceBusMessage\u003e messages, int startIndex) {\n+        int index = startIndex;\n+        while (index \u003c messages.size() \u0026\u0026 batch.tryAddMessage(messages.get(index))) {\n+            index++;\n+        }\n+        if (index == startIndex) {\n+            throw new IllegalArgumentException(\"Order message exceeds the Service Bus maximum message size\");\n+        }\n+        return index;\n+    }\n+\n+    private ServiceBusMessage toMessage(Order order) {\n+        return OrderMessageFactory.create(order, highPriorityThreshold);\n+    }\n+\n+    private boolean isHighPriority(Order order) {\n+        return order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0;\n+    }\n+\n+    @Override\n+    public void close() {\n+        sender.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/SyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.DeadLetterOptions;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.time.Duration;\n+import java.util.function.Consumer;\n+\n+public final class SyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);\n+\n+    private final ServiceBusSessionReceiverClient receiver;\n+    private final ServiceBusSessionReceiverClient deadLetterReceiver;\n+\n+    public SyncOrderProcessor(\n+            String fullyQualifiedNamespace, String queueName, TokenCredential credential) {\n+        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential);\n+        this.receiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .disableAutoComplete()\n+                .buildClient();\n+        this.deadLetterReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .disableAutoComplete()\n+                .buildClient();\n+    }\n+\n+    public void processAvailableSessions(int maximumSessions, Duration sessionWaitTime) {\n+        processSessions(receiver, maximumSessions, sessionWaitTime, this::processMessage);\n+    }\n+\n+    public void inspectAndReprocessDeadLetters(\n+            int maximumSessions, Duration sessionWaitTime, Consumer\u003cOrder\u003e resubmitter) {\n+        processSessions(deadLetterReceiver, maximumSessions, sessionWaitTime, (session, message) -\u003e {\n+            LOGGER.warn(\"Dead letter: messageId={}, reason={}, description={}\",\n+                    message.getMessageId(),\n+                    message.getDeadLetterReason(),\n+                    message.getDeadLetterErrorDescription());\n+            try {\n+                Order order = OrderJson.deserialize(message.getBody().toString());\n+                resubmitter.accept(order);\n+                session.complete(message);\n+                LOGGER.info(\"Requeued order {}\", order.getOrderId());\n+            } catch (Exception exception) {\n+                LOGGER.error(\"Dead-letter message {} could not be reprocessed\",\n+                        message.getMessageId(), exception);\n+                session.abandon(message);\n+            }\n+        });\n+    }\n+\n+    private void processMessage(ServiceBusReceiverClient session, ServiceBusReceivedMessage message) {\n+        try {\n+            Order order = OrderJson.deserialize(message.getBody().toString());\n+            if (order.getStatus() == Order.Status.FAILED) {\n+                throw new IllegalStateException(\"Order arrived with FAILED status\");\n+            }\n+            LOGGER.info(\"Processing {}\", order.withStatus(Order.Status.PROCESSING));\n+            LOGGER.info(\"Completed {}\", order.withStatus(Order.Status.COMPLETED));\n+            session.complete(message);\n+        } catch (Exception exception) {\n+            DeadLetterOptions options = new DeadLetterOptions()\n+                    .setDeadLetterReason(\"ORDER_PROCESSING_FAILED\")\n+                    .setDeadLetterErrorDescription(exception.getMessage());\n+            session.deadLetter(message, options);\n+            LOGGER.error(\"Dead-lettered message {}\", message.getMessageId(), exception);\n+        }\n+    }\n+\n+    private void processSessions(\n+            ServiceBusSessionReceiverClient sessionReceiver,\n+            int maximumSessions,\n+            Duration sessionWaitTime,\n+            SessionMessageHandler handler) {\n+        for (int i = 0; i \u003c maximumSessions; i++) {\n+            try (ServiceBusReceiverClient session = sessionReceiver.acceptNextSession(sessionWaitTime)) {\n+                for (ServiceBusReceivedMessage message : session.receiveMessages(100, sessionWaitTime)) {\n+                    handler.handle(session, message);\n+                }\n+            } catch (RuntimeException exception) {\n+                LOGGER.debug(\"No additional session was available: {}\", exception.getMessage());\n+                return;\n+            }\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        receiver.close();\n+        deadLetterReceiver.close();\n+    }\n+\n+    @FunctionalInterface\n+    private interface SessionMessageHandler {\n+        void handle(ServiceBusReceiverClient session, ServiceBusReceivedMessage message);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.servicebus.DeadLetterOptions;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n+import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.SubQueue;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.concurrent.TimeoutException;\n+import java.util.function.Function;\n+\n+public final class AsyncOrderProcessor implements AutoCloseable {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);\n+\n+    private final ServiceBusSessionReceiverAsyncClient receiver;\n+    private final ServiceBusSessionReceiverAsyncClient deadLetterReceiver;\n+\n+    public AsyncOrderProcessor(\n+            String fullyQualifiedNamespace, String queueName, TokenCredential credential) {\n+        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()\n+                .credential(fullyQualifiedNamespace, credential);\n+        this.receiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+        this.deadLetterReceiver = builder.sessionReceiver()\n+                .queueName(queueName)\n+                .subQueue(SubQueue.DEAD_LETTER_QUEUE)\n+                .disableAutoComplete()\n+                .buildAsyncClient();\n+    }\n+\n+    public Mono\u003cVoid\u003e processAvailableSessions(int maximumSessions, Duration sessionWaitTime) {\n+        return processSessions(receiver, maximumSessions, sessionWaitTime, this::processMessage);\n+    }\n+\n+    public Mono\u003cVoid\u003e inspectAndReprocessDeadLetters(\n+            int maximumSessions,\n+            Duration sessionWaitTime,\n+            Function\u003cOrder, Mono\u003cVoid\u003e\u003e resubmitter) {\n+        return processSessions(deadLetterReceiver, maximumSessions, sessionWaitTime, (session, message) -\u003e {\n+            LOGGER.warn(\"Dead letter: messageId={}, reason={}, description={}\",\n+                    message.getMessageId(),\n+                    message.getDeadLetterReason(),\n+                    message.getDeadLetterErrorDescription());\n+            final Order order;\n+            try {\n+                order = OrderJson.deserialize(message.getBody().toString());\n+            } catch (Exception exception) {\n+                LOGGER.error(\"Dead-letter message {} could not be deserialized\",\n+                        message.getMessageId(), exception);\n+                return session.abandon(message);\n+            }\n+            return resubmitter.apply(order)\n+                    .then(session.complete(message))\n+                    .doOnSuccess(ignored -\u003e LOGGER.info(\"Requeued order {}\", order.getOrderId()));\n+        });\n+    }\n+\n+    private Mono\u003cVoid\u003e processMessage(\n+            ServiceBusReceiverAsyncClient session, ServiceBusReceivedMessage message) {\n+        try {\n+            Order order = OrderJson.deserialize(message.getBody().toString());\n+            if (order.getStatus() == Order.Status.FAILED) {\n+                throw new IllegalStateException(\"Order arrived with FAILED status\");\n+            }\n+            LOGGER.info(\"Processing {}\", order.withStatus(Order.Status.PROCESSING));\n+            LOGGER.info(\"Completed {}\", order.withStatus(Order.Status.COMPLETED));\n+            return session.complete(message);\n+        } catch (Exception exception) {\n+            DeadLetterOptions options = new DeadLetterOptions()\n+                    .setDeadLetterReason(\"ORDER_PROCESSING_FAILED\")\n+                    .setDeadLetterErrorDescription(exception.getMessage());\n+            LOGGER.error(\"Dead-lettering message {}\", message.getMessageId(), exception);\n+            return session.deadLetter(message, options);\n+        }\n+    }\n+\n+    private Mono\u003cVoid\u003e processSessions(\n+            ServiceBusSessionReceiverAsyncClient sessionReceiver,\n+            int remainingSessions,\n+            Duration sessionWaitTime,\n+            AsyncSessionMessageHandler handler) {\n+        if (remainingSessions == 0) {\n+            return Mono.empty();\n+        }\n+        return sessionReceiver.acceptNextSession(sessionWaitTime)\n+                .flatMap(session -\u003e session.receiveMessages()\n+                        .concatMap(message -\u003e handler.handle(session, message))\n+                        .timeout(sessionWaitTime)\n+                        .onErrorResume(TimeoutException.class, ignored -\u003e Mono.empty())\n+                        .then()\n+                        .doFinally(signal -\u003e session.close()))\n+                .then(processSessions(\n+                        sessionReceiver, remainingSessions - 1, sessionWaitTime, handler))\n+                .onErrorResume(exception -\u003e {\n+                    LOGGER.debug(\"No additional session was available: {}\", exception.getMessage());\n+                    return Mono.empty();\n+                });\n+    }\n+\n+    @Override\n+    public void close() {\n+        receiver.close();\n+        deadLetterReceiver.close();\n+    }\n+\n+    @FunctionalInterface\n+    private interface AsyncSessionMessageHandler {\n+        Mono\u003cVoid\u003e handle(ServiceBusReceiverAsyncClient session, ServiceBusReceivedMessage message);\n+    }\n+}\n*** Add File: src/main/java/com/example/orders/Main.java\n+package com.example.orders;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import reactor.core.publisher.Mono;\n+\n+import java.math.BigDecimal;\n+import java.time.Duration;\n+import java.util.List;\n+\n+public final class Main {\n+    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal(\"1000.00\");\n+    private static final Duration SESSION_WAIT = Duration.ofSeconds(5);\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String namespace = requiredEnvironmentVariable(\"SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE\");\n+        String queueName = requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+        TokenCredential credential = new DefaultAzureCredentialBuilder().build();\n+\n+        runSyncDemo(namespace, queueName, credential);\n+        runAsyncDemo(namespace, queueName, credential).block();\n+    }\n+\n+    private static void runSyncDemo(\n+            String namespace, String queueName, TokenCredential credential) {\n+        List\u003cOrder\u003e orders = demoOrders(\"sync\");\n+        try (SyncOrderSender sender =\n+                     new SyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD);\n+             SyncOrderProcessor processor =\n+                     new SyncOrderProcessor(namespace, queueName, credential)) {\n+            sender.send(orders.get(0));\n+            sender.sendBatch(orders.subList(1, orders.size()));\n+            processor.processAvailableSessions(10, SESSION_WAIT);\n+            processor.inspectAndReprocessDeadLetters(\n+                    10,\n+                    SESSION_WAIT,\n+                    order -\u003e sender.send(order.withStatus(Order.Status.PENDING)));\n+            processor.processAvailableSessions(10, SESSION_WAIT);\n+        }\n+    }\n+\n+    private static Mono\u003cVoid\u003e runAsyncDemo(\n+            String namespace, String queueName, TokenCredential credential) {\n+        List\u003cOrder\u003e orders = demoOrders(\"async\");\n+        AsyncOrderSender sender =\n+                new AsyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD);\n+        AsyncOrderProcessor processor =\n+                new AsyncOrderProcessor(namespace, queueName, credential);\n+\n+        return sender.send(orders.get(0))\n+                .then(sender.sendBatch(orders.subList(1, orders.size())))\n+                .then(processor.processAvailableSessions(10, SESSION_WAIT))\n+                .then(processor.inspectAndReprocessDeadLetters(\n+                        10,\n+                        SESSION_WAIT,\n+                        order -\u003e sender.send(order.withStatus(Order.Status.PENDING))))\n+                .then(processor.processAvailableSessions(10, SESSION_WAIT))\n+                .doFinally(signal -\u003e {\n+                    processor.close();\n+                    sender.close();\n+                });\n+    }\n+\n+    private static List\u003cOrder\u003e demoOrders(String prefix) {\n+        return List.of(\n+                new Order(prefix + \"-001\", \"Ada\", \"Keyboard\", 1,\n+                        new BigDecimal(\"89.99\"), Order.Status.PENDING),\n+                new Order(prefix + \"-002\", \"Ada\", \"Monitor\", 2,\n+                        new BigDecimal(\"699.98\"), Order.Status.PENDING),\n+                new Order(prefix + \"-003\", \"Grace\", \"Workstation\", 1,\n+                        new BigDecimal(\"2499.00\"), Order.Status.PENDING),\n+                new Order(prefix + \"-004\", \"Linus\", \"Cable\", 3,\n+                        new BigDecimal(\"29.97\"), Order.Status.FAILED));\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(name + \" must be set\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Azure Service Bus Order Processor\n+\n+Java 17 Maven sample with synchronous and asynchronous order senders and\n+session-aware processors. The Service Bus queue must be created with sessions\n+enabled because each message uses the customer name as its session ID.\n+\n+Set these environment variables:\n+\n+```text\n+SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE=\u003cnamespace\u003e.servicebus.windows.net\n+SERVICE_BUS_QUEUE_NAME=orders\n+```\n+\n+Authenticate locally with any credential supported by `DefaultAzureCredential`,\n+or run in Azure with a managed identity that has the **Azure Service Bus Data\n+Sender** and **Azure Service Bus Data Receiver** roles. Then run:\n+\n+```text\n+mvn compile exec:java\n+```\n+\n+Orders above `$1,000` are marked high priority and scheduled 30 seconds into\n+the future. The demo includes an order with `FAILED` status to exercise\n+dead-letter inspection and reprocessing.\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile the Maven project","initial_wait":180,"mode":"sync"}
```

### 🔧 5. powershell

**Input:**

```json
{"command":"$jar = Get-ChildItem \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-messaging-servicebus\\*\\azure-messaging-servicebus-*.jar\" | Sort-Object FullName -Descending | Select-Object -First 1 -ExpandProperty FullName; javap -classpath $jar com.azure.messaging.servicebus.ServiceBusMessage; javap -classpath $jar com.azure.messaging.servicebus.ServiceBusSessionReceiverClient; javap -classpath $jar com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient","description":"Inspect Service Bus SDK APIs","initial_wait":30,"mode":"sync"}
```

### 🔧 6. powershell

**Input:**

```json
{"command":"$jar = Get-ChildItem \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-messaging-servicebus\\*\\azure-messaging-servicebus-*.jar\" | Sort-Object FullName -Descending | Select-Object -First 1 -ExpandProperty FullName; jar tf $jar | Select-String 'DeadLetter|SessionReceiverClientBuilder'","description":"Locate dead-letter SDK classes","initial_wait":30,"mode":"sync"}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"$jar = Get-ChildItem \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-messaging-servicebus\\*\\azure-messaging-servicebus-*.jar\" | Sort-Object FullName -Descending | Select-Object -First 1 -ExpandProperty FullName; javap -classpath $jar 'com.azure.messaging.servicebus.ServiceBusClientBuilder$ServiceBusSessionReceiverClientBuilder'; javap -classpath $jar com.azure.messaging.servicebus.models.DeadLetterOptions","description":"Inspect receiver builder options","initial_wait":30,"mode":"sync"}
```

### 🔧 8. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/OrderMessageFactory.java\n@@\n     static ServiceBusMessage create(Order order, BigDecimal highPriorityThreshold) {\n         boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) \u003e 0;\n-        return new ServiceBusMessage(OrderJson.serialize(order))\n+        ServiceBusMessage message = new ServiceBusMessage(OrderJson.serialize(order))\n                 .setContentType(\"application/json\")\n                 .setMessageId(order.getOrderId())\n                 .setCorrelationId(order.getOrderId())\n                 .setSessionId(order.getCustomerName())\n-                .setSubject(\"order\")\n-                .setApplicationProperties(\n-                        java.util.Map.of(PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY));\n+                .setSubject(\"order\");\n+        message.getApplicationProperties().put(\n+                PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY);\n+        return message;\n     }\n }\n*** Update File: src/main/java/com/example/orders/SyncOrderProcessor.java\n@@\n-import com.azure.messaging.servicebus.DeadLetterOptions;\n import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n@@\n         for (int i = 0; i \u003c maximumSessions; i++) {\n-            try (ServiceBusReceiverClient session = sessionReceiver.acceptNextSession(sessionWaitTime)) {\n+            try (ServiceBusReceiverClient session = sessionReceiver.acceptNextSession()) {\n                 for (ServiceBusReceivedMessage message : session.receiveMessages(100, sessionWaitTime)) {\n                     handler.handle(session, message);\n                 }\n-            } catch (RuntimeException exception) {\n-                LOGGER.debug(\"No additional session was available: {}\", exception.getMessage());\n-                return;\n             }\n         }\n*** Update File: src/main/java/com/example/orders/AsyncOrderProcessor.java\n@@\n-import com.azure.messaging.servicebus.DeadLetterOptions;\n import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;\n import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;\n+import com.azure.messaging.servicebus.models.DeadLetterOptions;\n@@\n-        return sessionReceiver.acceptNextSession(sessionWaitTime)\n+        return sessionReceiver.acceptNextSession()\n+                .timeout(sessionWaitTime)\n+                .onErrorResume(TimeoutException.class, ignored -\u003e Mono.empty())\n                 .flatMap(session -\u003e session.receiveMessages()\n                         .concatMap(message -\u003e handler.handle(session, message))\n                         .timeout(sessionWaitTime)\n@@\n                         .then()\n                         .doFinally(signal -\u003e session.close()))\n-                .then(processSessions(\n-                        sessionReceiver, remainingSessions - 1, sessionWaitTime, handler))\n-                .onErrorResume(exception -\u003e {\n-                    LOGGER.debug(\"No additional session was available: {}\", exception.getMessage());\n-                    return Mono.empty();\n-                });\n+                .then(Mono.defer(() -\u003e processSessions(\n+                        sessionReceiver, remainingSessions - 1, sessionWaitTime, handler)));\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n             sender.send(orders.get(0));\n             sender.sendBatch(orders.subList(1, orders.size()));\n-            processor.processAvailableSessions(10, SESSION_WAIT);\n+            processor.processAvailableSessions(2, SESSION_WAIT);\n             processor.inspectAndReprocessDeadLetters(\n-                    10,\n+                    1,\n                     SESSION_WAIT,\n                     order -\u003e sender.send(order.withStatus(Order.Status.PENDING)));\n-            processor.processAvailableSessions(10, SESSION_WAIT);\n+            processor.processAvailableSessions(1, SESSION_WAIT);\n+            waitForScheduledOrder();\n+            processor.processAvailableSessions(1, SESSION_WAIT);\n@@\n                 .then(sender.sendBatch(orders.subList(1, orders.size())))\n-                .then(processor.processAvailableSessions(10, SESSION_WAIT))\n+                .then(processor.processAvailableSessions(2, SESSION_WAIT))\n                 .then(processor.inspectAndReprocessDeadLetters(\n-                        10,\n+                        1,\n                         SESSION_WAIT,\n                         order -\u003e sender.send(order.withStatus(Order.Status.PENDING))))\n-                .then(processor.processAvailableSessions(10, SESSION_WAIT))\n+                .then(processor.processAvailableSessions(1, SESSION_WAIT))\n+                .then(Mono.delay(Duration.ofSeconds(31)))\n+                .then(processor.processAvailableSessions(1, SESSION_WAIT))\n                 .doFinally(signal -\u003e {\n@@\n                 new Order(prefix + \"-004\", \"Linus\", \"Cable\", 3,\n                         new BigDecimal(\"29.97\"), Order.Status.FAILED));\n     }\n+\n+    private static void waitForScheduledOrder() {\n+        try {\n+            Thread.sleep(Duration.ofSeconds(31));\n+        } catch (InterruptedException exception) {\n+            Thread.currentThread().interrupt();\n+            throw new IllegalStateException(\"Interrupted while waiting for scheduled order\", exception);\n+        }\n+    }\n*** End Patch\n"
```

### 🔧 9. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/Main.java\n@@\n         try {\n-            Thread.sleep(Duration.ofSeconds(31));\n+            Thread.sleep(Duration.ofSeconds(31).toMillis());\n*** End Patch\n"
```

### 🔧 10. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Recompile corrected Maven project","initial_wait":180,"mode":"sync"}
```

### 🔧 11. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/orders/Order.java\n@@\n import com.fasterxml.jackson.annotation.JsonCreator;\n import com.fasterxml.jackson.annotation.JsonProperty;\n+import com.fasterxml.jackson.annotation.JsonValue;\n@@\n     public enum Status {\n         PENDING,\n         PROCESSING,\n         COMPLETED,\n-        FAILED\n+        FAILED;\n+\n+        @JsonCreator\n+        public static Status fromJson(String value) {\n+            return Status.valueOf(value.toUpperCase(java.util.Locale.ROOT));\n+        }\n+\n+        @JsonValue\n+        public String toJson() {\n+            return name().toLowerCase(java.util.Locale.ROOT);\n+        }\n     }\n*** Update File: README.md\n@@\n-Authenticate locally with any credential supported by `DefaultAzureCredential`,\n-or run in Azure with a managed identity that has the **Azure Service Bus Data\n-Sender** and **Azure Service Bus Data Receiver** roles. Then run:\n+`DefaultAzureCredential` uses the workload's managed identity when the sample\n+runs in Azure. Assign that identity the **Azure Service Bus Data Sender** and\n+**Azure Service Bus Data Receiver** roles. For local development it can use\n+Azure CLI or another credential in its default chain. Then run:\n*** End Patch\n"
```

### 🔧 12. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress package","description":"Package the Java project","initial_wait":180,"mode":"sync"}
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
- `src\main\java\com\example\orders\SyncOrderProcessor.java`
- `src\main\java\com\example\orders\SyncOrderSender.java`

## Copilot Response

I’ll inspect the workspace, then create the Maven structure and wire the Azure Service Bus session-based sender/processor flows.

The implementation will use Service Bus sessions keyed by customer, size-aware SDK message batches, explicit settlement, and a repair callback for dead-letter reprocessing.

The project is in place; I’m compiling against the actual Azure SDK APIs now and will correct any signature or lifecycle issues surfaced by Maven.

Compilation is clean. I’m tightening the JSON contract so status values match the requested lowercase wire format, then packaging the final artifact.

Created the Java 17 Maven project with:

- JSON-serializable `Order` model and lowercase status values
- Synchronous and reactive asynchronous senders
- Size-aware message batching using `tryAddMessage`
- Correlation IDs and customer-based Service Bus sessions
- High-priority marking and 30-second scheduled delivery
- Explicit completion/dead-letter settlement
- Synchronous and asynchronous dead-letter inspection/reprocessing
- Managed identity support through `DefaultAzureCredential`
- Demo flow in `Main` and setup instructions in `README.md`

The packaged artifact is at `target\service-bus-order-processor-1.0.0.jar`. The queue must have **sessions enabled**.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Consistent lifecycle management covers sender, processor, session, and dead-letter clients in both synchronous and asynchronous flows.

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
hyoka run --prompt-id service-bus-dp-java-order-processor --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)

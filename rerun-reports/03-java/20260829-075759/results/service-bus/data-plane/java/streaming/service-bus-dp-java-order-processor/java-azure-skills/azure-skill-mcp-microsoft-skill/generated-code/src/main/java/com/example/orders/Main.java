package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.json.JsonMapper;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;
import java.util.UUID;

public final class Main {
    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal("1000.00");
    private static final Duration PROCESSING_WINDOW = Duration.ofSeconds(40);

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String namespace = requiredEnvironmentVariable("AZURE_SERVICE_BUS_NAMESPACE");
        String queueName = System.getenv().getOrDefault("AZURE_SERVICE_BUS_QUEUE", "orders");
        ObjectMapper objectMapper = JsonMapper.builder().findAndAddModules().build();
        TokenCredential credential = new DefaultAzureCredentialBuilder().build();

        runSyncDemo(namespace, queueName, credential, objectMapper);
        runAsyncDemo(namespace, queueName, credential, objectMapper);
    }

    private static void runSyncDemo(
            String namespace,
            String queueName,
            TokenCredential credential,
            ObjectMapper objectMapper) throws InterruptedException {
        ServiceBusSenderClient senderClient = builder(namespace, credential).sender()
                .queueName(queueName)
                .buildClient();
        ServiceBusProcessorClient processorClient = builder(namespace, credential).sessionProcessor()
                .queueName(queueName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .maxConcurrentSessions(1)
                .maxConcurrentCalls(1)
                .processMessage(context -> SyncOrderProcessor.processMessage(context, objectMapper))
                .processError(context -> System.err.println("Sync processor error: " + context.getException()))
                .buildProcessorClient();
        ServiceBusReceiverClient deadLetterReceiver = builder(namespace, credential).receiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildClient();
        ServiceBusSenderClient reprocessSender = builder(namespace, credential).sender()
                .queueName(queueName)
                .buildClient();

        try (SyncOrderSender sender =
                     new SyncOrderSender(senderClient, objectMapper, HIGH_PRIORITY_THRESHOLD);
             SyncOrderProcessor processor =
                     new SyncOrderProcessor(processorClient, deadLetterReceiver, reprocessSender, objectMapper)) {
            sender.sendBatch(sampleOrders("sync"));
            sendMalformed(senderClient, "sync-bad");
            processor.runFor(PROCESSING_WINDOW);
            processor.inspectAndReprocessDeadLetters(10, Duration.ofSeconds(5));
        }
    }

    private static void runAsyncDemo(
            String namespace,
            String queueName,
            TokenCredential credential,
            ObjectMapper objectMapper) {
        ServiceBusSenderAsyncClient senderClient = builder(namespace, credential).sender()
                .queueName(queueName)
                .buildAsyncClient();
        ServiceBusSessionReceiverAsyncClient sessionReceiver = builder(namespace, credential).sessionReceiver()
                .queueName(queueName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildAsyncClient();
        ServiceBusReceiverAsyncClient deadLetterReceiver = builder(namespace, credential).receiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildAsyncClient();
        ServiceBusSenderAsyncClient reprocessSender = builder(namespace, credential).sender()
                .queueName(queueName)
                .buildAsyncClient();

        try (AsyncOrderSender sender =
                     new AsyncOrderSender(senderClient, objectMapper, HIGH_PRIORITY_THRESHOLD);
             AsyncOrderProcessor processor =
                     new AsyncOrderProcessor(sessionReceiver, deadLetterReceiver, reprocessSender, objectMapper)) {
            sender.sendBatch(sampleOrders("async"))
                    .then(senderClient.sendMessage(malformedMessage("async-bad")))
                    .then(processor.processFor(PROCESSING_WINDOW))
                    .then(processor.inspectAndReprocessDeadLetters(10, Duration.ofSeconds(5)))
                    .block();
        }
    }

    private static ServiceBusClientBuilder builder(String namespace, TokenCredential credential) {
        return new ServiceBusClientBuilder().credential(namespace, credential);
    }

    private static List<Order> sampleOrders(String prefix) {
        return List.of(
                new Order(prefix + "-001", "Ada", "Keyboard", 1, new BigDecimal("149.99"), Order.Status.PENDING),
                new Order(prefix + "-002", "Ada", "Monitor", 2, new BigDecimal("799.98"), Order.Status.PENDING),
                new Order(prefix + "-003", "Grace", "Server", 1, new BigDecimal("4200.00"), Order.Status.PENDING));
    }

    private static void sendMalformed(ServiceBusSenderClient sender, String id) {
        sender.sendMessage(malformedMessage(id));
    }

    private static ServiceBusMessage malformedMessage(String id) {
        return new ServiceBusMessage(BinaryData.fromString("{not-valid-json"))
                .setMessageId(id + "-" + UUID.randomUUID())
                .setCorrelationId(id)
                .setSessionId("Malformed demo");
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set, for example contoso.servicebus.windows.net");
        }
        return value;
    }
}

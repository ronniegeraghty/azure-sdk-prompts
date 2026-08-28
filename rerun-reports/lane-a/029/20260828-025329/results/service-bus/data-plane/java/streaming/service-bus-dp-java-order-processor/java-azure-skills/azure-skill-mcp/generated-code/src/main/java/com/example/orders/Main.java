package com.example.orders;

import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import com.fasterxml.jackson.databind.ObjectMapper;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;

public final class Main {
    private static final String NAMESPACE_ENV = "SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE";
    private static final String QUEUE_ENV = "SERVICE_BUS_QUEUE_NAME";
    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal("1000.00");
    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(45);

    private Main() {
    }

    public static void main(String[] args) {
        String namespace = requiredEnvironmentVariable(NAMESPACE_ENV);
        String queueName = requiredEnvironmentVariable(QUEUE_ENV);
        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
        ObjectMapper objectMapper = new ObjectMapper();

        runSyncDemo(namespace, queueName, credential, objectMapper);
        runAsyncDemo(namespace, queueName, credential, objectMapper);
    }

    private static void runSyncDemo(
            String namespace,
            String queueName,
            DefaultAzureCredential credential,
            ObjectMapper objectMapper) {
        ServiceBusClientBuilder builder = clientBuilder(namespace, credential);
        try (ServiceBusSenderClient senderClient = builder.sender().queueName(queueName).buildClient();
             SyncOrderSender sender = new SyncOrderSender(
                     senderClient,
                     objectMapper,
                     HIGH_PRIORITY_THRESHOLD);
             SyncOrderProcessor processor = new SyncOrderProcessor(
                     sessionReceiver(builder, queueName, null),
                     objectMapper)) {
            List<Order> orders = demoOrders("sync");
            sender.send(orders.get(0));
            sender.sendBatch(orders.subList(1, orders.size()));
            sendMalformedMessage(senderClient, "Sync-DLQ");

            for (int i = 0; i < 3; i++) {
                processor.processNextCustomer(10, RECEIVE_WAIT);
            }
        }

        try (SyncOrderProcessor deadLetterProcessor = new SyncOrderProcessor(
                sessionReceiver(clientBuilder(namespace, credential), queueName, SubQueue.DEAD_LETTER_QUEUE),
                objectMapper)) {
            deadLetterProcessor.inspectAndReprocessDeadLetters(
                    10,
                    RECEIVE_WAIT,
                    order -> System.out.println("Sync reprocessed: " + order));
        }
    }

    private static void runAsyncDemo(
            String namespace,
            String queueName,
            DefaultAzureCredential credential,
            ObjectMapper objectMapper) {
        ServiceBusClientBuilder builder = clientBuilder(namespace, credential);
        ServiceBusSenderAsyncClient rawSender = builder.sender().queueName(queueName).buildAsyncClient();
        try (AsyncOrderSender sender = new AsyncOrderSender(
                rawSender,
                objectMapper,
                HIGH_PRIORITY_THRESHOLD);
             AsyncOrderProcessor processor = new AsyncOrderProcessor(
                     asyncSessionReceiver(builder, queueName, null),
                     objectMapper)) {
            List<Order> orders = demoOrders("async");
            Mono<Void> demo = sender.send(orders.get(0))
                    .then(sender.sendBatch(orders.subList(1, orders.size())))
                    .then(rawSender.sendMessage(malformedMessage("Async-DLQ")))
                    .thenMany(FluxSupport.repeatSequentially(
                            3,
                            () -> processor.processNextCustomer(10, RECEIVE_WAIT)))
                    .then();
            demo.block();
        }

        try (AsyncOrderProcessor deadLetterProcessor = new AsyncOrderProcessor(
                asyncSessionReceiver(
                        clientBuilder(namespace, credential),
                        queueName,
                        SubQueue.DEAD_LETTER_QUEUE),
                objectMapper)) {
            deadLetterProcessor.inspectAndReprocessDeadLetters(
                            10,
                            RECEIVE_WAIT,
                            order -> Mono.fromRunnable(
                                    () -> System.out.println("Async reprocessed: " + order)))
                    .block();
        }
    }

    private static ServiceBusClientBuilder clientBuilder(
            String namespace,
            DefaultAzureCredential credential) {
        return new ServiceBusClientBuilder()
                .fullyQualifiedNamespace(namespace)
                .credential(credential);
    }

    private static ServiceBusSessionReceiverClient sessionReceiver(
            ServiceBusClientBuilder builder,
            String queueName,
            SubQueue subQueue) {
        ServiceBusClientBuilder.ServiceBusSessionReceiverClientBuilder receiverBuilder =
                builder.sessionReceiver()
                        .queueName(queueName)
                        .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                        .disableAutoComplete();
        if (subQueue != null) {
            receiverBuilder.subQueue(subQueue);
        }
        return receiverBuilder.buildClient();
    }

    private static ServiceBusSessionReceiverAsyncClient asyncSessionReceiver(
            ServiceBusClientBuilder builder,
            String queueName,
            SubQueue subQueue) {
        ServiceBusClientBuilder.ServiceBusSessionReceiverClientBuilder receiverBuilder =
                builder.sessionReceiver()
                        .queueName(queueName)
                        .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                        .disableAutoComplete();
        if (subQueue != null) {
            receiverBuilder.subQueue(subQueue);
        }
        return receiverBuilder.buildAsyncClient();
    }

    private static void sendMalformedMessage(ServiceBusSenderClient sender, String sessionId) {
        sender.sendMessage(malformedMessage(sessionId));
    }

    private static ServiceBusMessage malformedMessage(String sessionId) {
        return new ServiceBusMessage("{not-valid-json")
                .setContentType(OrderMessageFactory.CONTENT_TYPE)
                .setMessageId(sessionId + "-malformed")
                .setCorrelationId(sessionId + "-malformed")
                .setSessionId(sessionId);
    }

    private static List<Order> demoOrders(String prefix) {
        return List.of(
                new Order(prefix + "-001", "Ada", "Keyboard", 1,
                        new BigDecimal("99.95"), Order.Status.PENDING),
                new Order(prefix + "-002", "Ada", "Monitor", 2,
                        new BigDecimal("649.90"), Order.Status.PENDING),
                new Order(prefix + "-003", "Grace", "Workstation", 1,
                        new BigDecimal("2499.00"), Order.Status.PENDING));
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set the " + name + " environment variable");
        }
        return value;
    }
}

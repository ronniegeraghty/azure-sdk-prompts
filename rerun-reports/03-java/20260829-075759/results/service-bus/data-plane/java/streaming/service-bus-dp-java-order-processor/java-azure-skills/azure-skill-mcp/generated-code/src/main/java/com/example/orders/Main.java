package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;
import java.util.Optional;

public final class Main {
    private static final Duration PROCESSING_WINDOW = Duration.ofSeconds(35);
    private static final String SYNC_BAD_ORDER_ID = "sync-invalid";
    private static final String ASYNC_BAD_ORDER_ID = "async-invalid";

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String namespace = requiredEnvironment("SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE");
        String queueName = requiredEnvironment("SERVICE_BUS_QUEUE_NAME");
        BigDecimal priorityThreshold = new BigDecimal(
                System.getenv().getOrDefault("HIGH_PRIORITY_THRESHOLD", "1000.00"));
        TokenCredential credential = new DefaultAzureCredentialBuilder().build();

        runSyncDemo(namespace, queueName, credential, priorityThreshold);
        runAsyncDemo(namespace, queueName, credential, priorityThreshold).block();
    }

    private static void runSyncDemo(
            String namespace,
            String queueName,
            TokenCredential credential,
            BigDecimal threshold) throws InterruptedException {
        try (SyncOrderSender sender = new SyncOrderSender(namespace, queueName, credential, threshold);
             SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential)) {
            sender.sendOrders(sampleOrders("sync"));
            sendMalformedSync(namespace, queueName, credential, SYNC_BAD_ORDER_ID);
            processor.start();
            Thread.sleep(PROCESSING_WINDOW.toMillis());
            processor.stop();
            processor.reprocessDeadLetters(
                    10,
                    Duration.ofSeconds(5),
                    message -> recoverDemoOrder(message.getCorrelationId(), "sync"),
                    sender);
            processor.start();
            Thread.sleep(Duration.ofSeconds(5).toMillis());
            processor.stop();
        }
    }

    private static Mono<Void> runAsyncDemo(
            String namespace,
            String queueName,
            TokenCredential credential,
            BigDecimal threshold) {
        return Mono.using(
                () -> new AsyncOrderSender(namespace, queueName, credential, threshold),
                sender -> Mono.using(
                        () -> new AsyncOrderProcessor(namespace, queueName, credential),
                        processor -> sender.sendOrders(sampleOrders("async"))
                                .then(sendMalformedAsync(
                                        namespace, queueName, credential, ASYNC_BAD_ORDER_ID))
                                .then(processor.processAvailableSessions(3, PROCESSING_WINDOW))
                                .then(processor.reprocessDeadLetters(
                                        10,
                                        Duration.ofSeconds(5),
                                        message -> recoverDemoOrder(message.getCorrelationId(), "async"),
                                        sender))
                                .then(processor.processAvailableSessions(
                                        1, Duration.ofSeconds(5))),
                        AsyncOrderProcessor::close),
                AsyncOrderSender::close);
    }

    private static List<Order> sampleOrders(String prefix) {
        return List.of(
                new Order(prefix + "-001", "Alice", "Laptop", 1,
                        new BigDecimal("1499.00"), OrderStatus.PENDING),
                new Order(prefix + "-002", "Alice", "Mouse", 2,
                        new BigDecimal("79.98"), OrderStatus.PENDING),
                new Order(prefix + "-003", "Bob", "Monitor", 1,
                        new BigDecimal("399.00"), OrderStatus.PENDING));
    }

    private static void sendMalformedSync(
            String namespace,
            String queueName,
            TokenCredential credential,
            String orderId) {
        try (ServiceBusSenderClient sender = new ServiceBusClientBuilder()
                .credential(namespace, credential)
                .sender()
                .queueName(queueName)
                .buildClient()) {
            sender.sendMessage(malformedMessage(orderId));
        }
    }

    private static Mono<Void> sendMalformedAsync(
            String namespace,
            String queueName,
            TokenCredential credential,
            String orderId) {
        return Mono.using(
                () -> {
                    ServiceBusSenderAsyncClient sender = new ServiceBusClientBuilder()
                            .credential(namespace, credential)
                            .sender()
                            .queueName(queueName)
                            .buildAsyncClient();
                    return sender;
                },
                sender -> sender.sendMessage(malformedMessage(orderId)),
                ServiceBusSenderAsyncClient::close);
    }

    private static ServiceBusMessage malformedMessage(String orderId) {
        return new ServiceBusMessage("{not-valid-json")
                .setContentType("application/json")
                .setMessageId(orderId)
                .setCorrelationId(orderId)
                .setSessionId("DemoFailure");
    }

    private static Optional<Order> recoverDemoOrder(String correlationId, String prefix) {
        if (!(prefix + "-invalid").equals(correlationId)) {
            return Optional.empty();
        }
        return Optional.of(new Order(
                prefix + "-recovered",
                "DemoFailure",
                "Recovered Product",
                1,
                new BigDecimal("25.00"),
                OrderStatus.PENDING));
    }

    private static String requiredEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}

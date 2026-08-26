package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;

public final class Main {
    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal("1000.00");
    private static final Duration SESSION_WAIT = Duration.ofSeconds(5);

    private Main() {
    }

    public static void main(String[] args) {
        String namespace = requiredEnvironmentVariable("SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        TokenCredential credential = new DefaultAzureCredentialBuilder().build();

        runSyncDemo(namespace, queueName, credential);
        runAsyncDemo(namespace, queueName, credential).block();
    }

    private static void runSyncDemo(
            String namespace, String queueName, TokenCredential credential) {
        List<Order> orders = demoOrders("sync");
        try (SyncOrderSender sender =
                     new SyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD);
             SyncOrderProcessor processor =
                     new SyncOrderProcessor(namespace, queueName, credential)) {
            sender.send(orders.get(0));
            sender.sendBatch(orders.subList(1, orders.size()));
            processor.processAvailableSessions(2, SESSION_WAIT);
            processor.inspectAndReprocessDeadLetters(
                    1,
                    SESSION_WAIT,
                    order -> sender.send(order.withStatus(Order.Status.PENDING)));
            processor.processAvailableSessions(1, SESSION_WAIT);
            waitForScheduledOrder();
            processor.processAvailableSessions(1, SESSION_WAIT);
        }
    }

    private static Mono<Void> runAsyncDemo(
            String namespace, String queueName, TokenCredential credential) {
        List<Order> orders = demoOrders("async");
        AsyncOrderSender sender =
                new AsyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD);
        AsyncOrderProcessor processor =
                new AsyncOrderProcessor(namespace, queueName, credential);

        return sender.send(orders.get(0))
                .then(sender.sendBatch(orders.subList(1, orders.size())))
                .then(processor.processAvailableSessions(2, SESSION_WAIT))
                .then(processor.inspectAndReprocessDeadLetters(
                        1,
                        SESSION_WAIT,
                        order -> sender.send(order.withStatus(Order.Status.PENDING))))
                .then(processor.processAvailableSessions(1, SESSION_WAIT))
                .then(Mono.delay(Duration.ofSeconds(31)))
                .then(processor.processAvailableSessions(1, SESSION_WAIT))
                .doFinally(signal -> {
                    processor.close();
                    sender.close();
                });
    }

    private static List<Order> demoOrders(String prefix) {
        return List.of(
                new Order(prefix + "-001", "Ada", "Keyboard", 1,
                        new BigDecimal("89.99"), Order.Status.PENDING),
                new Order(prefix + "-002", "Ada", "Monitor", 2,
                        new BigDecimal("699.98"), Order.Status.PENDING),
                new Order(prefix + "-003", "Grace", "Workstation", 1,
                        new BigDecimal("2499.00"), Order.Status.PENDING),
                new Order(prefix + "-004", "Linus", "Cable", 3,
                        new BigDecimal("29.97"), Order.Status.FAILED));
    }

    private static void waitForScheduledOrder() {
        try {
            Thread.sleep(Duration.ofSeconds(31).toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for scheduled order", exception);
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }
}

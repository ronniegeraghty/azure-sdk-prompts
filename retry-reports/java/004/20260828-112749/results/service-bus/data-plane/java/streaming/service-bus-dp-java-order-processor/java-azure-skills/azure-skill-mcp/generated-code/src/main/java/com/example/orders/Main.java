package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;
import java.util.Optional;
import java.util.Set;

public final class Main {
    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(3);
    private static final Duration ASYNC_RECEIVE_WINDOW = Duration.ofSeconds(5);
    private static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(31);

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String namespace = requiredEnvironmentVariable("SERVICE_BUS_NAMESPACE");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE");
        BigDecimal threshold = new BigDecimal(
                System.getenv().getOrDefault("HIGH_PRIORITY_THRESHOLD", "1000.00"));
        TokenCredential credential = managedIdentityCredential();

        runSynchronousDemo(namespace, queueName, credential, threshold);
        runAsynchronousDemo(namespace, queueName, credential, threshold);
    }

    private static void runSynchronousDemo(
            String namespace,
            String queueName,
            TokenCredential credential,
            BigDecimal threshold) throws InterruptedException {
        List<Order> orders = demoOrders("sync");
        Set<String> initialCustomers = Set.of("Ada", "Grace", "Demo");
        Set<String> failedCustomers = Set.of("Demo");
        Set<String> delayedCustomers = Set.of("Ada", "Demo");

        try (OrderSender sender = new OrderSender(namespace, queueName, credential, threshold);
             OrderProcessor processor = new OrderProcessor(namespace, queueName, credential)) {
            sender.send(orders.get(0));
            sender.sendBatch(orders.subList(1, orders.size()));
            sender.sendMalformedForDemo("sync-malformed", "Demo");

            processor.processSessions(initialCustomers, RECEIVE_WAIT);
            processor.inspectDeadLetters(failedCustomers, RECEIVE_WAIT);
            int reprocessed = processor.reprocessDeadLetters(
                    failedCustomers,
                    RECEIVE_WAIT,
                    deadLetter -> recoverDemoOrder(deadLetter, "sync"),
                    sender);
            System.out.println("Sync reprocessed dead-letter count: " + reprocessed);

            Thread.sleep(FRAUD_REVIEW_DELAY.toMillis());
            processor.processSessions(delayedCustomers, RECEIVE_WAIT);
        }
    }

    private static void runAsynchronousDemo(
            String namespace,
            String queueName,
            TokenCredential credential,
            BigDecimal threshold) {
        List<Order> orders = demoOrders("async");
        Set<String> initialCustomers = Set.of("Ada", "Grace", "Demo");
        Set<String> failedCustomers = Set.of("Demo");
        Set<String> delayedCustomers = Set.of("Ada", "Demo");

        try (AsyncOrderSender sender = new AsyncOrderSender(namespace, queueName, credential, threshold);
             AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential)) {
            sender.send(orders.get(0))
                    .then(sender.sendBatch(orders.subList(1, orders.size())))
                    .then(sender.sendMalformedForDemo("async-malformed", "Demo"))
                    .then(processor.processSessions(initialCustomers, ASYNC_RECEIVE_WINDOW))
                    .then(processor.inspectDeadLetters(failedCustomers, ASYNC_RECEIVE_WINDOW))
                    .doOnNext(deadLetters ->
                            System.out.println("Async dead-letter count: " + deadLetters.size()))
                    .then(processor.reprocessDeadLetters(
                            failedCustomers,
                            ASYNC_RECEIVE_WINDOW,
                            deadLetter -> recoverDemoOrder(deadLetter, "async"),
                            sender))
                    .doOnNext(count ->
                            System.out.println("Async reprocessed dead-letter count: " + count))
                    .then(reactor.core.publisher.Mono.delay(FRAUD_REVIEW_DELAY))
                    .then(processor.processSessions(delayedCustomers, ASYNC_RECEIVE_WINDOW))
                    .block();
        }
    }

    private static List<Order> demoOrders(String prefix) {
        return List.of(
                new Order(prefix + "-001", "Ada", "Keyboard", 1,
                        new BigDecimal("129.99"), OrderStatus.PENDING),
                new Order(prefix + "-002", "Ada", "Monitor", 2,
                        new BigDecimal("1499.98"), OrderStatus.PENDING),
                new Order(prefix + "-003", "Grace", "Dock", 1,
                        new BigDecimal("249.00"), OrderStatus.PENDING));
    }

    private static TokenCredential managedIdentityCredential() {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            builder.clientId(clientId);
        }
        return builder.build();
    }

    private static Optional<Order> recoverDemoOrder(DeadLetteredOrder deadLetter, String prefix) {
        if (!(prefix + "-malformed").equals(deadLetter.correlationId())) {
            return Optional.empty();
        }
        return Optional.of(new Order(
                prefix + "-recovered",
                deadLetter.sessionId(),
                "Recovered demo order",
                1,
                new BigDecimal("1.00"),
                OrderStatus.PENDING));
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }

}

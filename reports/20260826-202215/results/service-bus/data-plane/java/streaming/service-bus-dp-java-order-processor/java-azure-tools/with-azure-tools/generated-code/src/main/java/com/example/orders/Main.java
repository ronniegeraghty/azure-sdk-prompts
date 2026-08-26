package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.example.orders.messaging.AsyncOrderSender;
import com.example.orders.messaging.SyncOrderSender;
import com.example.orders.model.Order;
import com.example.orders.model.OrderStatus;
import com.example.orders.processing.AsyncOrderProcessor;
import com.example.orders.processing.SyncOrderProcessor;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;

public final class Main {
    private static final Duration SESSION_IDLE_TIMEOUT = Duration.ofSeconds(5);
    private static final int CUSTOMER_SESSION_COUNT = 3;
    private static final int DEAD_LETTER_SESSION_COUNT = 1;

    private Main() {
    }

    public static void main(String[] args) {
        String namespace = requiredEnvironmentVariable("SERVICE_BUS_NAMESPACE");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        String fullyQualifiedNamespace = namespace.contains(".")
            ? namespace
            : namespace + ".servicebus.windows.net";
        BigDecimal highPriorityThreshold = new BigDecimal(
            System.getenv().getOrDefault("ORDER_HIGH_PRIORITY_THRESHOLD", "1000.00")
        );
        TokenCredential credential = managedIdentityCredential();

        runSyncDemo(fullyQualifiedNamespace, queueName, credential, highPriorityThreshold);
        runAsyncDemo(fullyQualifiedNamespace, queueName, credential, highPriorityThreshold);
    }

    private static void runSyncDemo(
        String namespace,
        String queueName,
        TokenCredential credential,
        BigDecimal highPriorityThreshold
    ) {
        try (
            SyncOrderSender sender = new SyncOrderSender(
                namespace, queueName, credential, highPriorityThreshold
            );
            SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential)
        ) {
            sender.sendOrders(demoOrders("sync"));
            processor.processSessions(CUSTOMER_SESSION_COUNT, SESSION_IDLE_TIMEOUT);
            processor.inspectAndReprocessDeadLetters(
                DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT, sender
            );
            processor.processSessions(DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT);
        }
    }

    private static void runAsyncDemo(
        String namespace,
        String queueName,
        TokenCredential credential,
        BigDecimal highPriorityThreshold
    ) {
        try (
            AsyncOrderSender sender = new AsyncOrderSender(
                namespace, queueName, credential, highPriorityThreshold
            );
            AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential)
        ) {
            sender.sendOrders(demoOrders("async"))
                .then(processor.processSessions(CUSTOMER_SESSION_COUNT, SESSION_IDLE_TIMEOUT))
                .then(processor.inspectAndReprocessDeadLetters(
                    DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT, sender
                ))
                .then(processor.processSessions(DEAD_LETTER_SESSION_COUNT, SESSION_IDLE_TIMEOUT))
                .block();
        }
    }

    private static List<Order> demoOrders(String prefix) {
        return List.of(
            new Order(prefix + "-001", prefix + "-alice", "Keyboard", 1,
                new BigDecimal("89.99"), OrderStatus.PENDING),
            new Order(prefix + "-002", prefix + "-alice", "Mouse", 2,
                new BigDecimal("119.98"), OrderStatus.PENDING),
            new Order(prefix + "-003", prefix + "-bob", "Monitor", 1,
                new BigDecimal("399.00"), OrderStatus.FAILED),
            new Order(prefix + "-004", prefix + "-carol", "Workstation", 1,
                new BigDecimal("2499.00"), OrderStatus.PENDING)
        );
    }

    private static TokenCredential managedIdentityCredential() {
        String clientId = System.getenv("AZURE_CLIENT_ID");
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        return clientId == null || clientId.isBlank()
            ? builder.build()
            : builder.clientId(clientId).build();
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Environment variable " + name + " is required");
        }
        return value;
    }
}

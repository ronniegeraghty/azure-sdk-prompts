package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;

public final class Main {
    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal("1000.00");
    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(5);
    private static final String DEMO_CUSTOMER = "Contoso";

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String namespace = requiredEnvironmentVariable("SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();

        runSyncDemo(namespace, queueName, credential);
        runAsyncDemo(namespace, queueName, credential);
    }

    private static void runSyncDemo(String namespace, String queueName, TokenCredential credential)
        throws InterruptedException {
        SyncOrderProcessor processor = new SyncOrderProcessor(namespace, queueName, credential);
        try (SyncOrderSender sender =
                 new SyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {
            sender.sendOrders(demoOrders("sync"));
            sendMalformedMessage(namespace, queueName, credential, "sync-invalid");

            processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);
            processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT);
            processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,
                ignored -> repairedOrder("sync-repaired"), sender);
            processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);

            Thread.sleep(Duration.ofSeconds(31).toMillis());
            processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT);
        }
    }

    private static void runAsyncDemo(String namespace, String queueName, TokenCredential credential) {
        AsyncOrderProcessor processor = new AsyncOrderProcessor(namespace, queueName, credential);
        try (AsyncOrderSender sender =
                 new AsyncOrderSender(namespace, queueName, credential, HIGH_PRIORITY_THRESHOLD)) {
            sender.sendOrders(demoOrders("async"))
                .then(sendMalformedMessageAsync(namespace, queueName, credential, "async-invalid"))
                .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
                .then(processor.inspectDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
                .then(processor.reprocessDeadLetters(DEMO_CUSTOMER, 10, RECEIVE_WAIT,
                    ignored -> repairedOrder("async-repaired"), sender))
                .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
                .then(reactor.core.publisher.Mono.delay(Duration.ofSeconds(31)))
                .then(processor.processCustomer(DEMO_CUSTOMER, 10, RECEIVE_WAIT))
                .block();
        }
    }

    private static List<Order> demoOrders(String prefix) {
        return List.of(
            new Order(prefix + "-001", DEMO_CUSTOMER, "Keyboard", 2,
                new BigDecimal("179.98"), Order.Status.pending),
            new Order(prefix + "-002", DEMO_CUSTOMER, "Server", 1,
                new BigDecimal("5500.00"), Order.Status.pending));
    }

    private static Order repairedOrder(String orderId) {
        return new Order(orderId, DEMO_CUSTOMER, "Replacement item", 1,
            new BigDecimal("25.00"), Order.Status.pending);
    }

    private static void sendMalformedMessage(String namespace, String queueName,
                                             TokenCredential credential, String messageId) {
        try (ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sender()
            .queueName(queueName)
            .buildClient()) {
            sender.sendMessage(malformedMessage(messageId));
        }
    }

    private static Mono<Void> sendMalformedMessageAsync(String namespace, String queueName,
                                                        TokenCredential credential, String messageId) {
        return Mono.using(
            () -> new ServiceBusClientBuilder()
                .credential(namespace, credential)
                .sender()
                .queueName(queueName)
                .buildAsyncClient(),
            sender -> sender.sendMessage(malformedMessage(messageId)),
            ServiceBusSenderAsyncClient::close);
    }

    private static ServiceBusMessage malformedMessage(String messageId) {
        return new ServiceBusMessage(BinaryData.fromString("{not-json"))
            .setContentType("application/json")
            .setMessageId(messageId)
            .setCorrelationId(messageId)
            .setSessionId(DEMO_CUSTOMER);
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }

}

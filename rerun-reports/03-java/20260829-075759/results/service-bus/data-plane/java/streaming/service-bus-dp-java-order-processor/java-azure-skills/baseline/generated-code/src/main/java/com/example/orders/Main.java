package com.example.orders;

import com.azure.core.amqp.AmqpRetryOptions;
import com.azure.identity.ManagedIdentityCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.json.JsonMapper;

import java.math.BigDecimal;
import java.time.Duration;
import java.util.List;

public final class Main {
    private static final BigDecimal HIGH_PRIORITY_THRESHOLD = new BigDecimal("1000.00");
    private static final Duration PROCESSING_WINDOW = Duration.ofSeconds(40);
    private static final Duration DEAD_LETTER_WINDOW = Duration.ofSeconds(5);

    private Main() {
    }

    public static void main(String[] args) {
        String namespace = requiredEnvironment("SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE");
        String queueName = requiredEnvironment("SERVICE_BUS_QUEUE_NAME");

        ManagedIdentityCredential credential = managedIdentityCredential();
        ServiceBusClientBuilder clientBuilder = new ServiceBusClientBuilder()
                .fullyQualifiedNamespace(namespace)
                .credential(credential)
                .retryOptions(new AmqpRetryOptions().setTryTimeout(Duration.ofSeconds(5)));
        ObjectMapper objectMapper = JsonMapper.builder().findAndAddModules().build();
        OrderMessageFactory messageFactory = new OrderMessageFactory(objectMapper, HIGH_PRIORITY_THRESHOLD);

        runSynchronousDemo(clientBuilder, queueName, objectMapper, messageFactory);
        runAsynchronousDemo(clientBuilder, queueName, objectMapper, messageFactory);
    }

    private static void runSynchronousDemo(
            ServiceBusClientBuilder clientBuilder,
            String queueName,
            ObjectMapper objectMapper,
            OrderMessageFactory messageFactory) {
        try (OrderSender sender = new OrderSender(
                clientBuilder.sender().queueName(queueName).buildClient(), messageFactory)) {
            sender.sendBatch(List.of(
                    order("sync-1", "Ada", "Keyboard", 1, "120.00"),
                    order("sync-2", "Ada", "Mouse", 2, "80.00"),
                    order("sync-3", "Grace", "Monitor", 1, "450.00")));
            sender.send(order("sync-high-1", "Linus", "Server", 1, "2500.00"));
            sender.sendMalformedDemoMessage("sync-invalid-1", "DeadLetterDemo");

            try (OrderProcessor processor = new OrderProcessor(clientBuilder, queueName, objectMapper)) {
                processor.processFor(PROCESSING_WINDOW);
            }

            try (DeadLetterQueueProcessor deadLetters =
                         new DeadLetterQueueProcessor(clientBuilder, queueName, objectMapper)) {
                deadLetters.inspectFor(DEAD_LETTER_WINDOW);
            }
        }
    }

    private static void runAsynchronousDemo(
            ServiceBusClientBuilder clientBuilder,
            String queueName,
            ObjectMapper objectMapper,
            OrderMessageFactory messageFactory) {
        try (AsyncOrderSender sender = new AsyncOrderSender(
                clientBuilder.sender().queueName(queueName).buildAsyncClient(), messageFactory)) {
            sender.sendBatch(List.of(
                    order("async-1", "Katherine", "Dock", 1, "180.00"),
                    order("async-2", "Katherine", "Headset", 1, "90.00"),
                    order("async-3", "Margaret", "Laptop", 1, "900.00")))
                    .then(sender.send(order("async-high-1", "James", "GPU cluster", 1, "5000.00")))
                    .then(sender.sendMalformedDemoMessage("async-invalid-1", "AsyncDeadLetterDemo"))
                    .block();

            try (AsyncOrderProcessor processor =
                         new AsyncOrderProcessor(clientBuilder, queueName, objectMapper)) {
                processor.processFor(PROCESSING_WINDOW);
            }

            try (AsyncDeadLetterQueueProcessor deadLetters =
                         new AsyncDeadLetterQueueProcessor(clientBuilder, queueName, objectMapper, null)) {
                deadLetters.processFor(DEAD_LETTER_WINDOW);
            }
        }
    }

    private static Order order(
            String id, String customer, String product, int quantity, String totalPrice) {
        return new Order(
                id, customer, product, quantity, new BigDecimal(totalPrice), Order.Status.PENDING);
    }

    private static String requiredEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Environment variable " + name + " is required");
        }
        return value;
    }

    private static ManagedIdentityCredential managedIdentityCredential() {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        return clientId == null || clientId.isBlank()
                ? builder.build()
                : builder.clientId(clientId).build();
    }
}

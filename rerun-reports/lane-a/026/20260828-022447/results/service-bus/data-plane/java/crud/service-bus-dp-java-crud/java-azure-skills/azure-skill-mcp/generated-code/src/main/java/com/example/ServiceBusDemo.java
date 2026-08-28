package com.example;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;

import java.time.Duration;

public final class ServiceBusDemo {
    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(10);
    private static final Duration PROCESSOR_RUN_TIME = Duration.ofSeconds(30);

    private ServiceBusDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String connectionString = requiredEnvironmentVariable("SERVICE_BUS_CONNECTION_STRING");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        String topicName = requiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
        String subscriptionName = requiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME");

        demonstrateQueueMessaging(connectionString, queueName);
        demonstrateProcessor(connectionString, queueName);
        demonstrateTopicAndSubscription(
            connectionString, topicName, subscriptionName);
    }

    private static void demonstrateQueueMessaging(
        String connectionString, String queueName) {

        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .queueName(queueName)
            .buildClient();

        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .receiver()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Single queue message"));

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int i = 1; i <= 5; i++) {
                ServiceBusMessage message =
                    new ServiceBusMessage("Batch queue message " + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalStateException(
                        "Message " + i + " did not fit in the Service Bus batch.");
                }
            }
            sender.sendMessages(batch);

            for (ServiceBusReceivedMessage message
                : receiver.receiveMessages(6, RECEIVE_WAIT)) {

                System.out.printf(
                    "Queue received: id=%s, body=%s%n",
                    message.getMessageId(),
                    message.getBody().toString());

                // Complete only after successful processing.
                receiver.complete(message);
            }
        } finally {
            receiver.close();
            sender.close();
        }
    }

    private static void demonstrateProcessor(
        String connectionString, String queueName) throws InterruptedException {

        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .queueName(queueName)
            .buildClient();

        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .processor()
            .queueName(queueName)
            .disableAutoComplete()
            .processMessage(context -> {
                ServiceBusReceivedMessage message = context.getMessage();
                System.out.printf(
                    "Processor received: id=%s, body=%s%n",
                    message.getMessageId(),
                    message.getBody().toString());
                context.complete();
            })
            .processError(context -> System.err.printf(
                "Processor error: source=%s, entity=%s, namespace=%s, error=%s%n",
                context.getErrorSource(),
                context.getEntityPath(),
                context.getFullyQualifiedNamespace(),
                context.getException()))
            .buildProcessorClient();

        try {
            processor.start();
            sender.sendMessage(
                new ServiceBusMessage("Message for the continuous processor"));
            Thread.sleep(PROCESSOR_RUN_TIME.toMillis());
        } finally {
            processor.close();
            sender.close();
        }
    }

    private static void demonstrateTopicAndSubscription(
        String connectionString,
        String topicName,
        String subscriptionName) {

        ServiceBusSenderClient topicSender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .topicName(topicName)
            .buildClient();

        ServiceBusReceiverClient subscriptionReceiver =
            new ServiceBusClientBuilder()
                .connectionString(connectionString)
                .receiver()
                .topicName(topicName)
                .subscriptionName(subscriptionName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .buildClient();

        try {
            topicSender.sendMessage(
                new ServiceBusMessage("Message sent through a topic"));

            for (ServiceBusReceivedMessage message
                : subscriptionReceiver.receiveMessages(1, RECEIVE_WAIT)) {

                System.out.printf(
                    "Subscription received: id=%s, body=%s%n",
                    message.getMessageId(),
                    message.getBody().toString());
                subscriptionReceiver.complete(message);
            }
        } finally {
            subscriptionReceiver.close();
            topicSender.close();
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Set the required environment variable " + name + ".");
        }
        return value;
    }
}

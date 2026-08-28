package com.example;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusSenderClient;

import java.time.Duration;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class ServiceBusDemo {
    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(10);

    private ServiceBusDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String connectionString = requiredEnvironmentVariable("SERVICE_BUS_CONNECTION_STRING");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        String topicName = requiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
        String subscriptionName = requiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME");

        demonstrateQueueMessaging(connectionString, queueName);
        demonstrateProcessor(connectionString, queueName);
        demonstrateTopicAndSubscription(connectionString, topicName, subscriptionName);
    }

    private static void demonstrateQueueMessaging(String connectionString, String queueName) {
        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .queueName(queueName)
            .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Single queue message"));
            System.out.println("Sent one queue message.");

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int i = 1; i <= 5; i++) {
                ServiceBusMessage message = new ServiceBusMessage("Batch message " + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalStateException("The batch cannot hold all five demo messages.");
                }
            }
            sender.sendMessages(batch);
            System.out.println("Sent a batch of five queue messages.");
        } finally {
            sender.close();
        }

        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .receiver()
            .queueName(queueName)
            .disableAutoComplete()
            .buildClient();

        try {
            for (ServiceBusReceivedMessage message : receiver.receiveMessages(6, RECEIVE_TIMEOUT)) {
                System.out.printf(
                    "Processing queue message %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString());

                receiver.complete(message);
            }
        } finally {
            receiver.close();
        }
    }

    private static void demonstrateProcessor(String connectionString, String queueName)
        throws InterruptedException {

        CountDownLatch processed = new CountDownLatch(1);
        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .processor()
            .queueName(queueName)
            .disableAutoComplete()
            .processMessage(context -> {
                ServiceBusReceivedMessage message = context.getMessage();
                System.out.printf(
                    "Processor handled message %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString());
                context.complete();
                processed.countDown();
            })
            .processError(context -> System.err.printf(
                "Processor error from %s: %s%n",
                context.getErrorSource(),
                context.getException().getMessage()))
            .buildProcessorClient();

        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .queueName(queueName)
            .buildClient();

        try {
            processor.start();
            sender.sendMessage(new ServiceBusMessage("Message for the continuous processor"));

            if (!processed.await(30, TimeUnit.SECONDS)) {
                throw new IllegalStateException("The processor did not receive the demo message in time.");
            }
        } finally {
            sender.close();
            processor.close();
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

        try {
            topicSender.sendMessage(new ServiceBusMessage("Topic message"));
            System.out.println("Sent one topic message.");
        } finally {
            topicSender.close();
        }

        ServiceBusReceiverClient subscriptionReceiver = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .receiver()
            .topicName(topicName)
            .subscriptionName(subscriptionName)
            .disableAutoComplete()
            .buildClient();

        try {
            for (ServiceBusReceivedMessage message
                : subscriptionReceiver.receiveMessages(1, RECEIVE_TIMEOUT)) {

                System.out.printf(
                    "Processing subscription message %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString());
                subscriptionReceiver.complete(message);
            }
        } finally {
            subscriptionReceiver.close();
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set the " + name + " environment variable.");
        }
        return value;
    }
}

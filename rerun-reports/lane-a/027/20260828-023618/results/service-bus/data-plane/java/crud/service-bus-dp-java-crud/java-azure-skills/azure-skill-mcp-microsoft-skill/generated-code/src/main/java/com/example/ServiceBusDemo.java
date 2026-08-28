package com.example;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusErrorContext;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;

import java.time.Duration;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class ServiceBusDemo {
    private static final String CONNECTION_STRING_ENV = "AZURE_SERVICE_BUS_CONNECTION_STRING";
    private static final String QUEUE_NAME_ENV = "AZURE_SERVICE_BUS_QUEUE_NAME";
    private static final String TOPIC_NAME_ENV = "AZURE_SERVICE_BUS_TOPIC_NAME";
    private static final String SUBSCRIPTION_NAME_ENV = "AZURE_SERVICE_BUS_SUBSCRIPTION_NAME";

    private ServiceBusDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String connectionString = requiredEnvironmentVariable(CONNECTION_STRING_ENV);
        String queueName = requiredEnvironmentVariable(QUEUE_NAME_ENV);
        String topicName = requiredEnvironmentVariable(TOPIC_NAME_ENV);
        String subscriptionName = requiredEnvironmentVariable(SUBSCRIPTION_NAME_ENV);

        sendToQueue(connectionString, queueName);
        receiveFromQueue(connectionString, queueName);
        processQueueContinuously(connectionString, queueName);
        demonstrateTopicSubscription(connectionString, topicName, subscriptionName);
    }

    private static void sendToQueue(String connectionString, String queueName) {
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
                ServiceBusMessage message = new ServiceBusMessage("Batch queue message " + i)
                    .setMessageId("batch-" + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException(
                        "Batch message " + i + " is too large for the Service Bus batch.");
                }
            }

            sender.sendMessages(batch);
            System.out.printf("Sent a batch of %d queue messages.%n", batch.getCount());
        } finally {
            sender.close();
        }
    }

    private static void receiveFromQueue(String connectionString, String queueName) {
        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .receiver()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .buildClient();

        try {
            for (ServiceBusReceivedMessage message
                : receiver.receiveMessages(6, Duration.ofSeconds(30))) {
                System.out.printf("Queue receiver processed: %s%n", message.getBody());
                receiver.complete(message);
            }
        } finally {
            receiver.close();
        }
    }

    private static void processQueueContinuously(String connectionString, String queueName)
        throws InterruptedException {
        CountDownLatch messageProcessed = new CountDownLatch(1);

        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .processor()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .processMessage(context -> {
                ServiceBusReceivedMessage message = context.getMessage();
                System.out.printf("Processor handled: %s%n", message.getBody());
                context.complete();
                messageProcessed.countDown();
            })
            .processError(ServiceBusDemo::processError)
            .buildProcessorClient();

        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .queueName(queueName)
            .buildClient();

        try {
            processor.start();
            sender.sendMessage(new ServiceBusMessage("Message for the continuous processor"));

            if (!messageProcessed.await(30, TimeUnit.SECONDS)) {
                System.out.println("The processor did not receive a message within 30 seconds.");
            }
        } finally {
            sender.close();
            processor.stop();
            processor.close();
        }
    }

    private static void demonstrateTopicSubscription(
        String connectionString,
        String topicName,
        String subscriptionName) {

        ServiceBusSenderClient topicSender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .topicName(topicName)
            .buildClient();

        ServiceBusReceiverClient subscriptionReceiver = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .receiver()
            .topicName(topicName)
            .subscriptionName(subscriptionName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .buildClient();

        try {
            topicSender.sendMessage(new ServiceBusMessage("Topic message"));

            for (ServiceBusReceivedMessage message
                : subscriptionReceiver.receiveMessages(1, Duration.ofSeconds(30))) {
                System.out.printf("Subscription received: %s%n", message.getBody());
                subscriptionReceiver.complete(message);
            }
        } finally {
            subscriptionReceiver.close();
            topicSender.close();
        }
    }

    private static void processError(ServiceBusErrorContext context) {
        System.err.printf(
            "Processor error from %s in namespace %s: %s%n",
            context.getErrorSource(),
            context.getFullyQualifiedNamespace(),
            context.getException().getMessage());
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set the " + name + " environment variable.");
        }
        return value;
    }
}

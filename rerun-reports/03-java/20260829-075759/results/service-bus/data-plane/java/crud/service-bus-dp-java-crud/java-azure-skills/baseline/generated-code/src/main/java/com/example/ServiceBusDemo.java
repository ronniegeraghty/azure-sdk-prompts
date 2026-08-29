package com.example;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusErrorContext;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;

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

        sendToQueue(connectionString, queueName);
        receiveFromQueue(connectionString, queueName);
        processQueueContinuously(connectionString, queueName);
        sendToTopicAndReceiveFromSubscription(
            connectionString, topicName, subscriptionName);
    }

    private static void sendToQueue(String connectionString, String queueName) {
        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .sender()
            .queueName(queueName)
            .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Single queue message"));

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int i = 1; i <= 5; i++) {
                ServiceBusMessage message = new ServiceBusMessage("Batch message " + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalStateException(
                        "The five demo messages do not fit in one Service Bus batch.");
                }
            }
            sender.sendMessages(batch);
            System.out.println("Sent one message and a batch of five messages to " + queueName);
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
                : receiver.receiveMessages(6, RECEIVE_TIMEOUT)) {
                System.out.printf(
                    "Queue receiver processed %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString());
                receiver.complete(message);
            }
        } finally {
            receiver.close();
        }
    }

    private static void processQueueContinuously(
        String connectionString, String queueName) throws InterruptedException {

        CountDownLatch processed = new CountDownLatch(1);
        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()
            .connectionString(connectionString)
            .processor()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .processMessage(context -> processMessage(context, processed))
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

            if (!processed.await(30, TimeUnit.SECONDS)) {
                throw new IllegalStateException(
                    "The processor did not receive its demo message within 30 seconds.");
            }
        } finally {
            sender.close();
            processor.close();
        }
    }

    private static void processMessage(
        ServiceBusReceivedMessageContext context, CountDownLatch processed) {

        ServiceBusReceivedMessage message = context.getMessage();
        System.out.printf(
            "Processor handled %s: %s%n",
            message.getMessageId(),
            message.getBody().toString());
        context.complete();
        processed.countDown();
    }

    private static void processError(ServiceBusErrorContext context) {
        System.err.printf(
            "Processor error in %s for entity %s: %s%n",
            context.getErrorSource(),
            context.getEntityPath(),
            context.getException().getMessage());
    }

    private static void sendToTopicAndReceiveFromSubscription(
        String connectionString, String topicName, String subscriptionName) {

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
                : subscriptionReceiver.receiveMessages(1, RECEIVE_TIMEOUT)) {
                System.out.printf(
                    "Subscription receiver processed %s: %s%n",
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
            throw new IllegalStateException("Set the " + name + " environment variable.");
        }
        return value;
    }
}

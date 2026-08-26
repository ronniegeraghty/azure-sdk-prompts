package com.example;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;

import java.time.Duration;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class ServiceBusDemo {
    private static final String CONNECTION_STRING =
            requiredEnvironmentVariable("SERVICE_BUS_CONNECTION_STRING");
    private static final String QUEUE_NAME =
            requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
    private static final String TOPIC_NAME =
            requiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
    private static final String SUBSCRIPTION_NAME =
            requiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME");

    private ServiceBusDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        sendQueueMessages();
        receiveAndCompleteQueueMessages();
        processQueueContinuously();
        sendToTopicAndReceiveFromSubscription();
    }

    private static void sendQueueMessages() {
        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
                .connectionString(CONNECTION_STRING)
                .sender()
                .queueName(QUEUE_NAME)
                .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Single queue message"));
            System.out.println("Sent one message to queue " + QUEUE_NAME);

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int i = 1; i <= 5; i++) {
                ServiceBusMessage message = new ServiceBusMessage("Batch message " + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalStateException(
                            "The batch cannot fit all five messages. Send multiple batches instead.");
                }
            }

            sender.sendMessages(batch);
            System.out.println("Sent a batch of five messages to queue " + QUEUE_NAME);
        } finally {
            sender.close();
        }
    }

    private static void receiveAndCompleteQueueMessages() {
        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()
                .connectionString(CONNECTION_STRING)
                .receiver()
                .queueName(QUEUE_NAME)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .buildClient();

        try {
            receiver.receiveMessages(6, Duration.ofSeconds(20)).forEach(message -> {
                System.out.printf("Received queue message: %s%n", message.getBody());

                // PEEK_LOCK messages remain on the queue until explicitly settled.
                receiver.complete(message);
                System.out.printf("Completed message %s%n", message.getMessageId());
            });
        } finally {
            receiver.close();
        }
    }

    private static void processQueueContinuously() throws InterruptedException {
        CountDownLatch processed = new CountDownLatch(1);

        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()
                .connectionString(CONNECTION_STRING)
                .processor()
                .queueName(QUEUE_NAME)
                .disableAutoComplete()
                .processMessage(context -> {
                    System.out.printf("Processor received: %s%n",
                            context.getMessage().getBody());
                    context.complete();
                    processed.countDown();
                })
                .processError(context -> {
                    System.err.printf("Processor error from %s: %s%n",
                            context.getErrorSource(), context.getException().getMessage());
                })
                .buildProcessorClient();

        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
                .connectionString(CONNECTION_STRING)
                .sender()
                .queueName(QUEUE_NAME)
                .buildClient();

        try {
            processor.start();
            sender.sendMessage(new ServiceBusMessage("Message for continuous processor"));

            if (!processed.await(30, TimeUnit.SECONDS)) {
                throw new IllegalStateException(
                        "The processor did not receive its demonstration message within 30 seconds.");
            }
        } finally {
            sender.close();
            processor.close();
        }
    }

    private static void sendToTopicAndReceiveFromSubscription() {
        ServiceBusSenderClient topicSender = new ServiceBusClientBuilder()
                .connectionString(CONNECTION_STRING)
                .sender()
                .topicName(TOPIC_NAME)
                .buildClient();

        ServiceBusReceiverClient subscriptionReceiver = new ServiceBusClientBuilder()
                .connectionString(CONNECTION_STRING)
                .receiver()
                .topicName(TOPIC_NAME)
                .subscriptionName(SUBSCRIPTION_NAME)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .buildClient();

        try {
            topicSender.sendMessage(new ServiceBusMessage("Topic message"));
            System.out.printf("Sent one message to topic %s%n", TOPIC_NAME);

            subscriptionReceiver.receiveMessages(1, Duration.ofSeconds(20))
                    .forEach(message -> {
                        System.out.printf("Subscription %s received: %s%n",
                                SUBSCRIPTION_NAME, message.getBody());
                        subscriptionReceiver.complete(message);
                    });
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

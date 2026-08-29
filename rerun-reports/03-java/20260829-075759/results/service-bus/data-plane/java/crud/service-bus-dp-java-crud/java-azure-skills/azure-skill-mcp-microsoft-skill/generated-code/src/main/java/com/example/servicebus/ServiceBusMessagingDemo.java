package com.example.servicebus;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
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

public final class ServiceBusMessagingDemo {
    private static final int BATCH_MESSAGE_COUNT = 5;
    private static final Duration RECEIVE_WAIT_TIME = Duration.ofSeconds(10);
    private static final long PROCESSOR_WAIT_SECONDS = 30;

    private ServiceBusMessagingDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        Config config = Config.fromEnvironment();
        TokenCredential credential = new DefaultAzureCredentialBuilder().build();

        sendToQueue(config, credential);
        receiveFromQueue(config, credential);
        processQueueContinuously(config, credential);
        sendToTopicAndReceiveFromSubscription(config, credential);
    }

    private static void sendToQueue(Config config, TokenCredential credential) {
        ServiceBusSenderClient sender = clientBuilder(config, credential)
            .sender()
            .queueName(config.queueName())
            .buildClient();

        try {
            ServiceBusMessage singleMessage = new ServiceBusMessage("Single queue message");
            sender.sendMessage(singleMessage);
            System.out.println("Sent one message to queue " + config.queueName());

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int i = 1; i <= BATCH_MESSAGE_COUNT; i++) {
                ServiceBusMessage message = new ServiceBusMessage("Queue batch message " + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException(
                        "Batch message " + i + " is too large for the Service Bus batch.");
                }
            }

            sender.sendMessages(batch);
            System.out.println("Sent a batch of " + batch.getCount() + " messages");
        } finally {
            sender.close();
        }
    }

    private static void receiveFromQueue(Config config, TokenCredential credential) {
        ServiceBusReceiverClient receiver = clientBuilder(config, credential)
            .receiver()
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .queueName(config.queueName())
            .buildClient();

        try {
            receiver.receiveMessages(1 + BATCH_MESSAGE_COUNT, RECEIVE_WAIT_TIME)
                .forEach(message -> {
                    processMessage("Queue receiver", message);
                    receiver.complete(message);
                    System.out.println("Completed queue message " + message.getMessageId());
                });
        } finally {
            receiver.close();
        }
    }

    private static void processQueueContinuously(Config config, TokenCredential credential)
        throws InterruptedException {

        CountDownLatch processedMessage = new CountDownLatch(1);
        ServiceBusProcessorClient processor = clientBuilder(config, credential)
            .processor()
            .queueName(config.queueName())
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .processMessage(context -> processProcessorMessage(context, processedMessage))
            .processError(ServiceBusMessagingDemo::processError)
            .buildProcessorClient();

        try {
            processor.start();
            sendProcessorDemoMessage(config, credential);

            if (!processedMessage.await(PROCESSOR_WAIT_SECONDS, TimeUnit.SECONDS)) {
                System.err.println("Processor did not receive a message before the timeout.");
            }
        } finally {
            processor.close();
        }
    }

    private static void sendProcessorDemoMessage(Config config, TokenCredential credential) {
        ServiceBusSenderClient sender = clientBuilder(config, credential)
            .sender()
            .queueName(config.queueName())
            .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Message for the continuous processor"));
        } finally {
            sender.close();
        }
    }

    private static void processProcessorMessage(
        ServiceBusReceivedMessageContext context,
        CountDownLatch processedMessage
    ) {
        ServiceBusReceivedMessage message = context.getMessage();
        processMessage("Queue processor", message);
        context.complete();
        processedMessage.countDown();
    }

    private static void processError(ServiceBusErrorContext context) {
        System.err.printf(
            "Processor error from %s: %s%n",
            context.getErrorSource(),
            context.getException().getMessage());
    }

    private static void sendToTopicAndReceiveFromSubscription(
        Config config,
        TokenCredential credential
    ) {
        ServiceBusSenderClient topicSender = clientBuilder(config, credential)
            .sender()
            .topicName(config.topicName())
            .buildClient();
        ServiceBusReceiverClient subscriptionReceiver = clientBuilder(config, credential)
            .receiver()
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .topicName(config.topicName())
            .subscriptionName(config.subscriptionName())
            .buildClient();

        try {
            topicSender.sendMessage(new ServiceBusMessage("Message sent through a topic"));
            System.out.println("Sent one message to topic " + config.topicName());

            subscriptionReceiver.receiveMessages(1, RECEIVE_WAIT_TIME)
                .forEach(message -> {
                    processMessage("Topic subscription", message);
                    subscriptionReceiver.complete(message);
                    System.out.println("Completed subscription message " + message.getMessageId());
                });
        } finally {
            subscriptionReceiver.close();
            topicSender.close();
        }
    }

    private static ServiceBusClientBuilder clientBuilder(
        Config config,
        TokenCredential credential
    ) {
        return new ServiceBusClientBuilder()
            .credential(config.fullyQualifiedNamespace(), credential);
    }

    private static void processMessage(String consumer, ServiceBusReceivedMessage message) {
        System.out.printf(
            "%s received sequence number %d: %s%n",
            consumer,
            message.getSequenceNumber(),
            message.getBody());
    }

    private record Config(
        String fullyQualifiedNamespace,
        String queueName,
        String topicName,
        String subscriptionName
    ) {
        private static Config fromEnvironment() {
            return new Config(
                requiredEnvironmentVariable("SERVICE_BUS_FQDN"),
                requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME"),
                requiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME"),
                requiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME"));
        }

        private static String requiredEnvironmentVariable(String name) {
            String value = System.getenv(name);
            if (value == null || value.isBlank()) {
                throw new IllegalStateException("Missing required environment variable: " + name);
            }
            return value;
        }
    }
}

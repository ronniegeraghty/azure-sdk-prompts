package com.example.servicebus;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusErrorContext;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;

import java.time.Duration;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class ServiceBusMessagingDemo {
    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(10);
    private static final Duration PROCESSOR_WAIT = Duration.ofSeconds(30);

    private ServiceBusMessagingDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String namespace = requiredEnvironmentVariable("SERVICE_BUS_NAMESPACE");
        String queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        String topicName = requiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
        String subscriptionName = requiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME");

        TokenCredential credential = new DefaultAzureCredentialBuilder().build();
        ServiceBusClientBuilder clientBuilder = new ServiceBusClientBuilder()
            .credential(namespace, credential);

        demonstrateQueueMessaging(clientBuilder, queueName);
        demonstrateProcessor(clientBuilder, queueName);
        demonstrateTopicAndSubscription(clientBuilder, topicName, subscriptionName);
    }

    private static void demonstrateQueueMessaging(
        ServiceBusClientBuilder clientBuilder,
        String queueName
    ) {
        ServiceBusSenderClient sender = clientBuilder.sender()
            .queueName(queueName)
            .buildClient();

        ServiceBusReceiverClient receiver = clientBuilder.receiver()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Single queue message"));

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int i = 1; i <= 5; i++) {
                ServiceBusMessage message = new ServiceBusMessage("Batch queue message " + i);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalStateException("The batch cannot hold all 5 demo messages.");
                }
            }
            sender.sendMessages(batch);

            for (ServiceBusReceivedMessage message : receiver.receiveMessages(6, RECEIVE_WAIT)) {
                System.out.printf(
                    "Queue message %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString()
                );
                receiver.complete(message);
            }
        } finally {
            receiver.close();
            sender.close();
        }
    }

    private static void demonstrateProcessor(
        ServiceBusClientBuilder clientBuilder,
        String queueName
    ) throws InterruptedException {
        CountDownLatch processed = new CountDownLatch(1);

        ServiceBusProcessorClient processor = clientBuilder.processor()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .processMessage(context -> {
                ServiceBusReceivedMessage message = context.getMessage();
                System.out.printf(
                    "Processor received %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString()
                );
                context.complete();
                processed.countDown();
            })
            .processError(ServiceBusMessagingDemo::processError)
            .buildProcessorClient();

        ServiceBusSenderClient sender = clientBuilder.sender()
            .queueName(queueName)
            .buildClient();

        try {
            processor.start();
            sender.sendMessage(new ServiceBusMessage("Message for continuous processor"));

            if (!processed.await(PROCESSOR_WAIT.toSeconds(), TimeUnit.SECONDS)) {
                throw new IllegalStateException("The processor did not receive the demo message in time.");
            }
        } finally {
            processor.stop();
            processor.close();
            sender.close();
        }
    }

    private static void demonstrateTopicAndSubscription(
        ServiceBusClientBuilder clientBuilder,
        String topicName,
        String subscriptionName
    ) {
        ServiceBusSenderClient topicSender = clientBuilder.sender()
            .topicName(topicName)
            .buildClient();

        ServiceBusReceiverClient subscriptionReceiver = clientBuilder.receiver()
            .topicName(topicName)
            .subscriptionName(subscriptionName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .buildClient();

        try {
            topicSender.sendMessage(new ServiceBusMessage("Topic message"));

            for (ServiceBusReceivedMessage message
                : subscriptionReceiver.receiveMessages(1, RECEIVE_WAIT)) {
                System.out.printf(
                    "Subscription message %s: %s%n",
                    message.getMessageId(),
                    message.getBody().toString()
                );
                subscriptionReceiver.complete(message);
            }
        } finally {
            subscriptionReceiver.close();
            topicSender.close();
        }
    }

    private static void processError(ServiceBusErrorContext context) {
        System.err.printf(
            "Service Bus processor error. source=%s, namespace=%s, entity=%s, error=%s%n",
            context.getErrorSource(),
            context.getFullyQualifiedNamespace(),
            context.getEntityPath(),
            context.getException()
        );
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set the " + name + " environment variable.");
        }
        return value;
    }
}

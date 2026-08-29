package com.example;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;

import java.time.Duration;
import java.util.Map;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class ServiceBusMessagingDemo {
    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(20);
    private static final Duration PROCESSOR_WAIT = Duration.ofSeconds(30);

    private final String fullyQualifiedNamespace;
    private final String queueName;
    private final String topicName;
    private final String subscriptionName;
    private final TokenCredential credential;

    private ServiceBusMessagingDemo(
        String fullyQualifiedNamespace,
        String queueName,
        String topicName,
        String subscriptionName
    ) {
        this.fullyQualifiedNamespace = fullyQualifiedNamespace;
        this.queueName = queueName;
        this.topicName = topicName;
        this.subscriptionName = subscriptionName;
        this.credential = new DefaultAzureCredentialBuilder().build();
    }

    public static void main(String[] args) throws InterruptedException {
        Map<String, String> environment = System.getenv();
        String namespace = environment.get("SERVICE_BUS_FQ_NAMESPACE");
        String queue = environment.get("SERVICE_BUS_QUEUE_NAME");
        String topic = environment.get("SERVICE_BUS_TOPIC_NAME");
        String subscription = environment.get("SERVICE_BUS_SUBSCRIPTION_NAME");

        if (isBlank(namespace) || isBlank(queue) || isBlank(topic) || isBlank(subscription)) {
            System.out.println("""
                No Azure connection was attempted. Set these environment variables to run the demo:
                  SERVICE_BUS_FQ_NAMESPACE=<namespace>.servicebus.windows.net
                  SERVICE_BUS_QUEUE_NAME=<queue>
                  SERVICE_BUS_TOPIC_NAME=<topic>
                  SERVICE_BUS_SUBSCRIPTION_NAME=<subscription>
                """);
            return;
        }

        ServiceBusMessagingDemo demo =
            new ServiceBusMessagingDemo(namespace, queue, topic, subscription);
        demo.demonstrateQueueMessaging();
        demo.demonstrateTopicAndSubscription();
    }

    private void demonstrateQueueMessaging() throws InterruptedException {
        sendQueueMessages();
        receiveAndCompleteQueueMessages();
        processQueueContinuously();
    }

    private void sendQueueMessages() {
        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .sender()
            .queueName(queueName)
            .buildClient();

        try {
            sender.sendMessage(new ServiceBusMessage("Single queue message"));
            System.out.println("Sent one queue message.");

            ServiceBusMessageBatch batch = sender.createMessageBatch();
            for (int index = 1; index <= 5; index++) {
                ServiceBusMessage message = new ServiceBusMessage("Batch message " + index);
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalStateException(
                        "The five-message demo batch exceeded the Service Bus size limit.");
                }
            }

            sender.sendMessages(batch);
            System.out.println("Sent a batch of five queue messages.");
        } finally {
            sender.close();
        }
    }

    private void receiveAndCompleteQueueMessages() {
        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .receiver()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .buildClient();

        try {
            for (ServiceBusReceivedMessage message : receiver.receiveMessages(6, RECEIVE_WAIT)) {
                System.out.printf(
                    "Received queue message %s: %s%n",
                    message.getMessageId(),
                    message.getBody());

                // Complete only after successful processing so Service Bus removes the message.
                receiver.complete(message);
            }
        } finally {
            receiver.close();
        }
    }

    private void processQueueContinuously() throws InterruptedException {
        CountDownLatch processed = new CountDownLatch(1);

        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .processor()
            .queueName(queueName)
            .disableAutoComplete()
            .processMessage(context -> {
                ServiceBusReceivedMessage message = context.getMessage();
                System.out.printf("Processor received: %s%n", message.getBody());
                context.complete();
                processed.countDown();
            })
            .processError(context -> System.err.printf(
                "Processor error from %s: %s%n",
                context.getErrorSource(),
                context.getException().getMessage()))
            .buildProcessorClient();

        ServiceBusSenderClient sender = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .sender()
            .queueName(queueName)
            .buildClient();

        try {
            processor.start();
            sender.sendMessage(new ServiceBusMessage("Message for the continuous processor"));

            if (!processed.await(PROCESSOR_WAIT.toSeconds(), TimeUnit.SECONDS)) {
                throw new IllegalStateException("The processor did not receive a message in time.");
            }
        } finally {
            processor.close();
            sender.close();
        }
    }

    private void demonstrateTopicAndSubscription() {
        ServiceBusSenderClient topicSender = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .sender()
            .topicName(topicName)
            .buildClient();

        ServiceBusReceiverClient subscriptionReceiver = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .receiver()
            .topicName(topicName)
            .subscriptionName(subscriptionName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .buildClient();

        try {
            topicSender.sendMessage(new ServiceBusMessage("Message sent through a topic"));

            for (ServiceBusReceivedMessage message
                : subscriptionReceiver.receiveMessages(1, RECEIVE_WAIT)) {
                System.out.printf("Subscription received: %s%n", message.getBody());
                subscriptionReceiver.complete(message);
            }
        } finally {
            subscriptionReceiver.close();
            topicSender.close();
        }
    }

    private static boolean isBlank(String value) {
        return value == null || value.isBlank();
    }
}

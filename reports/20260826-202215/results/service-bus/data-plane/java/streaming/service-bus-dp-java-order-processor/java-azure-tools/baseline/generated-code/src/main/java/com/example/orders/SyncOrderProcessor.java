package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.SubQueue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.function.Consumer;

public final class SyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);

    private final ServiceBusSessionReceiverClient receiver;
    private final ServiceBusSessionReceiverClient deadLetterReceiver;

    public SyncOrderProcessor(
            String fullyQualifiedNamespace, String queueName, TokenCredential credential) {
        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential);
        this.receiver = builder.sessionReceiver()
                .queueName(queueName)
                .disableAutoComplete()
                .buildClient();
        this.deadLetterReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .disableAutoComplete()
                .buildClient();
    }

    public void processAvailableSessions(int maximumSessions, Duration sessionWaitTime) {
        processSessions(receiver, maximumSessions, sessionWaitTime, this::processMessage);
    }

    public void inspectAndReprocessDeadLetters(
            int maximumSessions, Duration sessionWaitTime, Consumer<Order> resubmitter) {
        processSessions(deadLetterReceiver, maximumSessions, sessionWaitTime, (session, message) -> {
            LOGGER.warn("Dead letter: messageId={}, reason={}, description={}",
                    message.getMessageId(),
                    message.getDeadLetterReason(),
                    message.getDeadLetterErrorDescription());
            try {
                Order order = OrderJson.deserialize(message.getBody().toString());
                resubmitter.accept(order);
                session.complete(message);
                LOGGER.info("Requeued order {}", order.getOrderId());
            } catch (Exception exception) {
                LOGGER.error("Dead-letter message {} could not be reprocessed",
                        message.getMessageId(), exception);
                session.abandon(message);
            }
        });
    }

    private void processMessage(ServiceBusReceiverClient session, ServiceBusReceivedMessage message) {
        try {
            Order order = OrderJson.deserialize(message.getBody().toString());
            if (order.getStatus() == Order.Status.FAILED) {
                throw new IllegalStateException("Order arrived with FAILED status");
            }
            LOGGER.info("Processing {}", order.withStatus(Order.Status.PROCESSING));
            LOGGER.info("Completed {}", order.withStatus(Order.Status.COMPLETED));
            session.complete(message);
        } catch (Exception exception) {
            DeadLetterOptions options = new DeadLetterOptions()
                    .setDeadLetterReason("ORDER_PROCESSING_FAILED")
                    .setDeadLetterErrorDescription(exception.getMessage());
            session.deadLetter(message, options);
            LOGGER.error("Dead-lettered message {}", message.getMessageId(), exception);
        }
    }

    private void processSessions(
            ServiceBusSessionReceiverClient sessionReceiver,
            int maximumSessions,
            Duration sessionWaitTime,
            SessionMessageHandler handler) {
        for (int i = 0; i < maximumSessions; i++) {
            try (ServiceBusReceiverClient session = sessionReceiver.acceptNextSession()) {
                for (ServiceBusReceivedMessage message : session.receiveMessages(100, sessionWaitTime)) {
                    handler.handle(session, message);
                }
            }
        }
    }

    @Override
    public void close() {
        receiver.close();
        deadLetterReceiver.close();
    }

    @FunctionalInterface
    private interface SessionMessageHandler {
        void handle(ServiceBusReceiverClient session, ServiceBusReceivedMessage message);
    }
}

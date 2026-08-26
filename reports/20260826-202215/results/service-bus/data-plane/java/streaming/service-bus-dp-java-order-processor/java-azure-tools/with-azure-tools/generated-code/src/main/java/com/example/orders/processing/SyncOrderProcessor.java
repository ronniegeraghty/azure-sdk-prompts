package com.example.orders.processing;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import com.example.orders.codec.OrderJsonCodec;
import com.example.orders.model.Order;
import com.example.orders.model.OrderStatus;
import com.example.orders.messaging.SyncOrderSender;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;

public final class SyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);
    private static final int MAX_MESSAGES_PER_SESSION = 10_000;

    private final ServiceBusSessionReceiverClient sessionReceiver;
    private final ServiceBusSessionReceiverClient deadLetterSessionReceiver;
    private final OrderJsonCodec codec = new OrderJsonCodec();

    public SyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {
        this.sessionReceiver = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sessionReceiver()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .buildClient();
        this.deadLetterSessionReceiver = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sessionReceiver()
            .queueName(queueName)
            .subQueue(SubQueue.DEAD_LETTER_QUEUE)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .buildClient();
    }

    public void processSessions(int sessionCount, Duration idleTimeout) {
        for (int index = 0; index < sessionCount; index++) {
            try (ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {
                LOGGER.info("Accepted customer session {}", receiver.getSessionId());
                for (ServiceBusReceivedMessage message
                    : receiver.receiveMessages(MAX_MESSAGES_PER_SESSION, idleTimeout)) {
                    processMessage(receiver, message);
                }
            }
        }
    }

    public void inspectAndReprocessDeadLetters(
        int sessionCount,
        Duration idleTimeout,
        SyncOrderSender sender
    ) {
        for (int index = 0; index < sessionCount; index++) {
            try (ServiceBusReceiverClient receiver = deadLetterSessionReceiver.acceptNextSession()) {
                for (ServiceBusReceivedMessage message
                    : receiver.receiveMessages(MAX_MESSAGES_PER_SESSION, idleTimeout)) {
                    reprocessDeadLetter(receiver, message, sender);
                }
            }
        }
    }

    private void processMessage(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {
        try {
            OrderProcessingSupport.process(message, codec, LOGGER);
            receiver.complete(message);
        } catch (RuntimeException exception) {
            LOGGER.error("Dead-lettering message {}: {}", message.getMessageId(), exception.getMessage());
            receiver.deadLetter(message, OrderProcessingSupport.deadLetterOptions(exception));
        }
    }

    private void reprocessDeadLetter(
        ServiceBusReceiverClient receiver,
        ServiceBusReceivedMessage message,
        SyncOrderSender sender
    ) {
        LOGGER.warn("Inspecting dead-lettered message {}: reason={}, description={}",
            message.getMessageId(), message.getDeadLetterReason(), message.getDeadLetterErrorDescription());
        try {
            Order order = codec.deserialize(message.getBody().toString());
            order.setStatus(OrderStatus.PENDING);
            sender.sendOrder(order);
            receiver.complete(message);
            LOGGER.info("Re-enqueued dead-lettered order {}", order.getOrderId());
        } catch (RuntimeException exception) {
            LOGGER.error("Message {} could not be reprocessed and remains in the dead-letter queue",
                message.getMessageId(), exception);
            receiver.abandon(message);
        }
    }

    @Override
    public void close() {
        deadLetterSessionReceiver.close();
        sessionReceiver.close();
    }
}

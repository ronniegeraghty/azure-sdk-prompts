package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.IterableStream;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.SubQueue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.function.Function;

public final class SyncOrderProcessor {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);
    private static final String DESERIALIZATION_FAILURE = "OrderDeserializationFailed";

    private final String namespace;
    private final String queueName;
    private final TokenCredential credential;
    private final OrderMessageFactory messageFactory;

    public SyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {
        this.namespace = namespace;
        this.queueName = queueName;
        this.credential = credential;
        this.messageFactory = new OrderMessageFactory(java.math.BigDecimal.ZERO);
    }

    public void processCustomer(String customerName, int maxMessages, Duration waitTime) {
        try (ServiceBusSessionReceiverClient sessions = sessionReceiver(false);
             ServiceBusReceiverClient receiver = sessions.acceptSession(customerName)) {
            IterableStream<ServiceBusReceivedMessage> messages =
                receiver.receiveMessages(maxMessages, waitTime);
            for (ServiceBusReceivedMessage message : messages) {
                try {
                    Order order = messageFactory.deserialize(message.getBody());
                    order.setStatus(Order.Status.processing);
                    LOGGER.info("Synchronously processed order: {}", order);
                    order.setStatus(Order.Status.completed);
                    receiver.complete(message);
                } catch (RuntimeException exception) {
                    LOGGER.error("Dead-lettering message {}: {}",
                        message.getMessageId(), exception.getMessage());
                    receiver.deadLetter(message, new DeadLetterOptions()
                        .setDeadLetterReason(DESERIALIZATION_FAILURE)
                        .setDeadLetterErrorDescription(exception.getMessage()));
                }
            }
        }
    }

    public void inspectDeadLetters(String customerName, int maxMessages, Duration waitTime) {
        try (ServiceBusSessionReceiverClient sessions = sessionReceiver(true);
             ServiceBusReceiverClient receiver = sessions.acceptSession(customerName)) {
            for (ServiceBusReceivedMessage message : receiver.receiveMessages(maxMessages, waitTime)) {
                LOGGER.warn("DLQ message id={}, reason={}, description={}, body={}",
                    message.getMessageId(), message.getDeadLetterReason(),
                    message.getDeadLetterErrorDescription(), message.getBody());
                receiver.abandon(message);
            }
        }
    }

    public void reprocessDeadLetters(String customerName, int maxMessages, Duration waitTime,
                                     Function<ServiceBusReceivedMessage, Order> repair,
                                     SyncOrderSender sender) {
        try (ServiceBusSessionReceiverClient sessions = sessionReceiver(true);
             ServiceBusReceiverClient receiver = sessions.acceptSession(customerName)) {
            for (ServiceBusReceivedMessage message : receiver.receiveMessages(maxMessages, waitTime)) {
                try {
                    sender.sendOrder(repair.apply(message));
                    receiver.complete(message);
                    LOGGER.info("Requeued dead-letter message {}", message.getMessageId());
                } catch (RuntimeException exception) {
                    LOGGER.error("Could not reprocess dead-letter message {}",
                        message.getMessageId(), exception);
                    receiver.abandon(message);
                }
            }
        }
    }

    private ServiceBusSessionReceiverClient sessionReceiver(boolean deadLetterQueue) {
        ServiceBusClientBuilder.ServiceBusSessionReceiverClientBuilder builder =
            new ServiceBusClientBuilder()
                .credential(namespace, credential)
                .sessionReceiver()
                .queueName(queueName)
                .disableAutoComplete();
        if (deadLetterQueue) {
            builder.subQueue(SubQueue.DEAD_LETTER_QUEUE);
        }
        return builder.buildClient();
    }
}

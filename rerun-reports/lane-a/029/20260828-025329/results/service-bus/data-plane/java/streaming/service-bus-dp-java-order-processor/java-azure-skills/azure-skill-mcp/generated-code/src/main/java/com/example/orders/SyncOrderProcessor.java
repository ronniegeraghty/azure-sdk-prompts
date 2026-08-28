package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.time.Duration;
import java.util.Objects;
import java.util.function.Consumer;

public final class SyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);

    private final ServiceBusSessionReceiverClient sessionReceiver;
    private final ObjectMapper objectMapper;

    public SyncOrderProcessor(
            ServiceBusSessionReceiverClient sessionReceiver,
            ObjectMapper objectMapper) {
        this.sessionReceiver = Objects.requireNonNull(sessionReceiver, "sessionReceiver");
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
    }

    public int processNextCustomer(int maxMessages, Duration maxWait) {
        try (ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {
            int processed = 0;
            for (ServiceBusReceivedMessage message : receiver.receiveMessages(maxMessages, maxWait)) {
                processMessage(receiver, message);
                processed++;
            }
            return processed;
        }
    }

    public int inspectAndReprocessDeadLetters(
            int maxMessages,
            Duration maxWait,
            Consumer<Order> reprocessor) {
        Objects.requireNonNull(reprocessor, "reprocessor");
        try (ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {
            int processed = 0;
            for (ServiceBusReceivedMessage message : receiver.receiveMessages(maxMessages, maxWait)) {
                try {
                    Order order = deserialize(message);
                    LOGGER.warn(
                            "DLQ order={} reason={} description={}",
                            order.getOrderId(),
                            message.getDeadLetterReason(),
                            message.getDeadLetterErrorDescription());
                    reprocessor.accept(order);
                    receiver.complete(message);
                } catch (Exception e) {
                    LOGGER.error("Unable to reprocess dead-letter message {}", message.getMessageId(), e);
                    receiver.abandon(message);
                }
                processed++;
            }
            return processed;
        }
    }

    private void processMessage(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {
        try {
            Order order = deserialize(message);
            LOGGER.info("Sync processed customer={} order={}", order.getCustomerName(), order);
            receiver.complete(message);
        } catch (Exception e) {
            String reason = deadLetterReason(e);
            LOGGER.error("Sync processing failed for message {}; dead-lettering", message.getMessageId(), e);
            receiver.deadLetter(
                    message,
                    new DeadLetterOptions()
                            .setDeadLetterReason(reason)
                            .setDeadLetterErrorDescription(e.getMessage()));
        }
    }

    private Order deserialize(ServiceBusReceivedMessage message) throws IOException {
        return objectMapper.readValue(message.getBody().toBytes(), Order.class);
    }

    private static String deadLetterReason(Exception error) {
        return error instanceof IOException ? "OrderDeserializationFailed" : "OrderProcessingFailed";
    }

    @Override
    public void close() {
        sessionReceiver.close();
    }
}

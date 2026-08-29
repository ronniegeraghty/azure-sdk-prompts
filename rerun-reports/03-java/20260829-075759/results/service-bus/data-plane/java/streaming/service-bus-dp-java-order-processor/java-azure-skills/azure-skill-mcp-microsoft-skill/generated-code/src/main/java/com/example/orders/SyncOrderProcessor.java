package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Objects;

public final class SyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderProcessor.class);

    private final ServiceBusProcessorClient processor;
    private final ServiceBusReceiverClient deadLetterReceiver;
    private final ServiceBusSenderClient reprocessSender;
    private final ObjectMapper objectMapper;

    public SyncOrderProcessor(
            ServiceBusProcessorClient processor,
            ServiceBusReceiverClient deadLetterReceiver,
            ServiceBusSenderClient reprocessSender,
            ObjectMapper objectMapper) {
        this.processor = Objects.requireNonNull(processor, "processor");
        this.deadLetterReceiver = Objects.requireNonNull(deadLetterReceiver, "deadLetterReceiver");
        this.reprocessSender = Objects.requireNonNull(reprocessSender, "reprocessSender");
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
    }

    public void runFor(Duration duration) throws InterruptedException {
        processor.start();
        try {
            Thread.sleep(duration.toMillis());
        } finally {
            processor.stop();
        }
    }

    public void inspectAndReprocessDeadLetters(int maxMessages, Duration waitTime) {
        for (ServiceBusReceivedMessage message : deadLetterReceiver.receiveMessages(maxMessages, waitTime)) {
            LOGGER.warn(
                    "DLQ message id={}, reason={}, description={}, body={}",
                    message.getMessageId(),
                    message.getDeadLetterReason(),
                    message.getDeadLetterErrorDescription(),
                    message.getBody());
            try {
                Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);
                ServiceBusMessageReprocessor.send(reprocessSender, order, objectMapper);
                deadLetterReceiver.complete(message);
                LOGGER.info("Reprocessed dead-lettered order {}", order.getOrderId());
            } catch (Exception exception) {
                LOGGER.warn("DLQ message {} is not a valid order and was left in the DLQ", message.getMessageId());
            }
        }
    }

    public static void processMessage(
            com.azure.messaging.servicebus.ServiceBusReceivedMessageContext context,
            ObjectMapper objectMapper) {
        ServiceBusReceivedMessage message = context.getMessage();
        try {
            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);
            order.setStatus(Order.Status.PROCESSING);
            LOGGER.info("Processing {}", order);
            order.setStatus(Order.Status.COMPLETED);
            LOGGER.info("Completed {}", order);
            context.complete();
        } catch (Exception exception) {
            String reason = "ORDER_DESERIALIZATION_OR_PROCESSING_FAILED";
            LOGGER.error("Dead-lettering message {}: {}", message.getMessageId(), exception.getMessage());
            context.deadLetter(new DeadLetterOptions()
                    .setDeadLetterReason(reason)
                    .setDeadLetterErrorDescription(truncate(exception.getMessage(), 4096)));
        }
    }

    private static String truncate(String value, int maxLength) {
        String safeValue = value == null ? "Unknown processing error" : value;
        return safeValue.length() <= maxLength ? safeValue : safeValue.substring(0, maxLength);
    }

    @Override
    public void close() {
        processor.close();
        deadLetterReceiver.close();
        reprocessSender.close();
    }
}

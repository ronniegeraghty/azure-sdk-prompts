package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusException;
import com.azure.messaging.servicebus.ServiceBusFailureReason;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.SubQueue;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.time.Instant;
import java.util.Objects;
import java.util.function.Consumer;

public final class DeadLetterQueueProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(DeadLetterQueueProcessor.class);
    private final ServiceBusSessionReceiverClient sessionReceiver;
    private final ObjectMapper objectMapper;

    public DeadLetterQueueProcessor(
            ServiceBusClientBuilder clientBuilder, String queueName, ObjectMapper objectMapper) {
        this.sessionReceiver = Objects.requireNonNull(clientBuilder, "clientBuilder")
                .sessionReceiver()
                .queueName(Objects.requireNonNull(queueName, "queueName"))
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .disableAutoComplete()
                .buildClient();
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
    }

    public void inspectFor(Duration duration) {
        readFor(duration, false, message ->
                LOGGER.info("Dead-letter message id={}, reason={}, description={}, body={}",
                        message.getMessageId(),
                        message.getDeadLetterReason(),
                        message.getDeadLetterErrorDescription(),
                        message.getBody()));
    }

    public void reprocessFor(Duration duration, Consumer<Order> resend) {
        Objects.requireNonNull(resend, "resend");
        readFor(duration, true, message -> {
            try {
                Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);
                resend.accept(order);
                LOGGER.info("Re-enqueued dead-lettered order {}", order.getOrderId());
            } catch (Exception exception) {
                throw new IllegalArgumentException(
                        "Dead-letter message " + message.getMessageId() + " cannot be reprocessed", exception);
            }
        });
    }

    private void readFor(
            Duration duration, boolean completeOnSuccess, Consumer<ServiceBusReceivedMessage> action) {
        Instant deadline = Instant.now().plus(duration);
        while (Instant.now().isBefore(deadline)) {
            try (ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {
                receiver.receiveMessages(100, remaining(deadline)).forEach(message -> {
                    try {
                        action.accept(message);
                        if (completeOnSuccess) {
                            receiver.complete(message);
                        } else {
                            receiver.abandon(message);
                        }
                    } catch (Exception exception) {
                        LOGGER.error("Could not handle dead-letter message {}: {}",
                                message.getMessageId(), OrderProcessor.rootMessage(exception));
                        receiver.abandon(message);
                    }
                });
            } catch (ServiceBusException exception) {
                if (exception.getReason() != ServiceBusFailureReason.SERVICE_TIMEOUT) {
                    throw exception;
                }
            }
        }
    }

    private static Duration remaining(Instant deadline) {
        Duration remaining = Duration.between(Instant.now(), deadline);
        return remaining.isNegative() || remaining.isZero() ? Duration.ofMillis(1) : remaining;
    }

    @Override
    public void close() {
        sessionReceiver.close();
    }
}

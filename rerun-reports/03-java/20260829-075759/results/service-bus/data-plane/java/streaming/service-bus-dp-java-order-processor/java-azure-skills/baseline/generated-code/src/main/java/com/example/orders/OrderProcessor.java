package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusException;
import com.azure.messaging.servicebus.ServiceBusFailureReason;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.time.Instant;
import java.util.Objects;

public final class OrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(OrderProcessor.class);
    private final ServiceBusSessionReceiverClient sessionReceiver;
    private final ObjectMapper objectMapper;

    public OrderProcessor(ServiceBusClientBuilder clientBuilder, String queueName, ObjectMapper objectMapper) {
        this.sessionReceiver = Objects.requireNonNull(clientBuilder, "clientBuilder")
                .sessionReceiver()
                .queueName(Objects.requireNonNull(queueName, "queueName"))
                .disableAutoComplete()
                .buildClient();
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
    }

    public void processFor(Duration duration) {
        Instant deadline = Instant.now().plus(duration);
        while (Instant.now().isBefore(deadline)) {
            try (ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {
                receiver.receiveMessages(100, remaining(deadline)).forEach(message -> process(receiver, message));
            } catch (ServiceBusException exception) {
                if (exception.getReason() != ServiceBusFailureReason.SERVICE_TIMEOUT) {
                    throw exception;
                }
            }
        }
    }

    private void process(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {
        try {
            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);
            LOGGER.info("Synchronously processed {}", order);
            receiver.complete(message);
        } catch (Exception exception) {
            String description = rootMessage(exception);
            LOGGER.error("Dead-lettering message {}: {}", message.getMessageId(), description);
            receiver.deadLetter(message, new DeadLetterOptions()
                    .setDeadLetterReason("OrderProcessingFailed")
                    .setDeadLetterErrorDescription(description));
        }
    }

    private static Duration remaining(Instant deadline) {
        Duration remaining = Duration.between(Instant.now(), deadline);
        return remaining.isNegative() || remaining.isZero() ? Duration.ofMillis(1) : remaining;
    }

    static String rootMessage(Throwable throwable) {
        Throwable current = throwable;
        while (current.getCause() != null) {
            current = current.getCause();
        }
        String message = current.getMessage();
        return message == null || message.isBlank() ? current.getClass().getSimpleName() : message;
    }

    @Override
    public void close() {
        sessionReceiver.close();
    }
}

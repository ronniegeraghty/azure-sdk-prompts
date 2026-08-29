package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.Objects;
import java.util.concurrent.TimeoutException;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);

    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;
    private final ServiceBusReceiverAsyncClient deadLetterReceiver;
    private final ServiceBusSenderAsyncClient reprocessSender;
    private final ObjectMapper objectMapper;

    public AsyncOrderProcessor(
            ServiceBusSessionReceiverAsyncClient sessionReceiver,
            ServiceBusReceiverAsyncClient deadLetterReceiver,
            ServiceBusSenderAsyncClient reprocessSender,
            ObjectMapper objectMapper) {
        this.sessionReceiver = Objects.requireNonNull(sessionReceiver, "sessionReceiver");
        this.deadLetterReceiver = Objects.requireNonNull(deadLetterReceiver, "deadLetterReceiver");
        this.reprocessSender = Objects.requireNonNull(reprocessSender, "reprocessSender");
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
    }

    public Mono<Void> processFor(Duration duration) {
        long deadline = System.nanoTime() + duration.toNanos();
        return processNextSession(deadline);
    }

    private Mono<Void> processNextSession(long deadline) {
        Duration remaining = Duration.ofNanos(Math.max(0, deadline - System.nanoTime()));
        if (remaining.isZero()) {
            return Mono.empty();
        }

        return sessionReceiver.acceptNextSession()
                .timeout(remaining)
                .flatMap(receiver -> processSession(receiver, remaining)
                        .doFinally(signal -> receiver.close()))
                .then(Mono.defer(() -> processNextSession(deadline)))
                .onErrorResume(TimeoutException.class, exception -> Mono.empty());
    }

    private Mono<Void> processSession(ServiceBusReceiverAsyncClient receiver, Duration remaining) {
        Duration idleTimeout = remaining.compareTo(Duration.ofSeconds(5)) < 0
                ? remaining
                : Duration.ofSeconds(5);
        return receiver.receiveMessages()
                .concatMap(message -> processMessage(receiver, message))
                .timeout(idleTimeout)
                .onErrorResume(TimeoutException.class, exception -> Flux.empty())
                .then();
    }

    private Mono<Void> processMessage(ServiceBusReceiverAsyncClient receiver, ServiceBusReceivedMessage message) {
        try {
            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);
            order.setStatus(Order.Status.PROCESSING);
            LOGGER.info("Processing {}", order);
            order.setStatus(Order.Status.COMPLETED);
            LOGGER.info("Completed {}", order);
            return receiver.complete(message);
        } catch (Exception exception) {
            LOGGER.error("Dead-lettering message {}: {}", message.getMessageId(), exception.getMessage());
            return receiver.deadLetter(message, new DeadLetterOptions()
                    .setDeadLetterReason("ORDER_DESERIALIZATION_OR_PROCESSING_FAILED")
                    .setDeadLetterErrorDescription(truncate(exception.getMessage(), 4096)));
        }
    }

    public Mono<Void> inspectAndReprocessDeadLetters(int maxMessages, Duration waitTime) {
        return deadLetterReceiver.receiveMessages()
                .take(maxMessages)
                .concatMap(message -> inspectAndReprocess(message)
                        .onErrorResume(exception -> {
                            LOGGER.warn(
                                    "DLQ message {} could not be reprocessed and was left in the DLQ: {}",
                                    message.getMessageId(),
                                    exception.getMessage());
                            return Mono.empty();
                        }))
                .timeout(waitTime)
                .onErrorResume(TimeoutException.class, exception -> Flux.empty())
                .then();
    }

    private Mono<Void> inspectAndReprocess(ServiceBusReceivedMessage message) {
        LOGGER.warn(
                "DLQ message id={}, reason={}, description={}, body={}",
                message.getMessageId(),
                message.getDeadLetterReason(),
                message.getDeadLetterErrorDescription(),
                message.getBody());
        try {
            Order order = objectMapper.readValue(message.getBody().toBytes(), Order.class);
            return ServiceBusMessageReprocessor.send(reprocessSender, order, objectMapper)
                    .then(deadLetterReceiver.complete(message))
                    .doOnSuccess(ignored -> LOGGER.info("Reprocessed dead-lettered order {}", order.getOrderId()));
        } catch (Exception exception) {
            return Mono.error(exception);
        }
    }

    private static String truncate(String value, int maxLength) {
        String safeValue = value == null ? "Unknown processing error" : value;
        return safeValue.length() <= maxLength ? safeValue : safeValue.substring(0, maxLength);
    }

    @Override
    public void close() {
        sessionReceiver.close();
        deadLetterReceiver.close();
        reprocessSender.close();
    }
}

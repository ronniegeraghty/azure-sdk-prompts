package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.io.IOException;
import java.time.Duration;
import java.util.Objects;
import java.util.function.Function;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);

    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;
    private final ObjectMapper objectMapper;

    public AsyncOrderProcessor(
            ServiceBusSessionReceiverAsyncClient sessionReceiver,
            ObjectMapper objectMapper) {
        this.sessionReceiver = Objects.requireNonNull(sessionReceiver, "sessionReceiver");
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
    }

    public Mono<Long> processNextCustomer(int maxMessages, Duration maxWait) {
        return sessionReceiver.acceptNextSession()
                .timeout(maxWait)
                .flatMap(receiver -> receiver.receiveMessages()
                        .take(maxMessages)
                        .timeout(maxWait)
                        .concatMap(message -> processMessage(receiver, message))
                        .count()
                        .doFinally(signal -> receiver.close()));
    }

    public Mono<Long> inspectAndReprocessDeadLetters(
            int maxMessages,
            Duration maxWait,
            Function<Order, Mono<Void>> reprocessor) {
        Objects.requireNonNull(reprocessor, "reprocessor");
        return sessionReceiver.acceptNextSession()
                .timeout(maxWait)
                .flatMap(receiver -> receiver.receiveMessages()
                        .take(maxMessages)
                        .timeout(maxWait)
                        .concatMap(message -> reprocess(receiver, message, reprocessor))
                        .count()
                        .doFinally(signal -> receiver.close()));
    }

    private Mono<Void> processMessage(
            ServiceBusReceiverAsyncClient receiver,
            ServiceBusReceivedMessage message) {
        return Mono.fromCallable(() -> deserialize(message))
                .doOnNext(order -> LOGGER.info(
                        "Async processed customer={} order={}",
                        order.getCustomerName(),
                        order))
                .flatMap(order -> receiver.complete(message))
                .onErrorResume(error -> {
                    String reason = error instanceof IOException
                            ? "OrderDeserializationFailed"
                            : "OrderProcessingFailed";
                    LOGGER.error(
                            "Async processing failed for message {}; dead-lettering",
                            message.getMessageId(),
                            error);
                    return receiver.deadLetter(
                            message,
                            new DeadLetterOptions()
                                    .setDeadLetterReason(reason)
                                    .setDeadLetterErrorDescription(error.getMessage()));
                });
    }

    private Mono<Void> reprocess(
            ServiceBusReceiverAsyncClient receiver,
            ServiceBusReceivedMessage message,
            Function<Order, Mono<Void>> reprocessor) {
        return Mono.fromCallable(() -> deserialize(message))
                .doOnNext(order -> LOGGER.warn(
                        "DLQ order={} reason={} description={}",
                        order.getOrderId(),
                        message.getDeadLetterReason(),
                        message.getDeadLetterErrorDescription()))
                .flatMap(reprocessor)
                .then(receiver.complete(message))
                .onErrorResume(error -> {
                    LOGGER.error("Unable to reprocess dead-letter message {}", message.getMessageId(), error);
                    return receiver.abandon(message);
                });
    }

    private Order deserialize(ServiceBusReceivedMessage message) throws IOException {
        return objectMapper.readValue(message.getBody().toBytes(), Order.class);
    }

    @Override
    public void close() {
        sessionReceiver.close();
    }
}

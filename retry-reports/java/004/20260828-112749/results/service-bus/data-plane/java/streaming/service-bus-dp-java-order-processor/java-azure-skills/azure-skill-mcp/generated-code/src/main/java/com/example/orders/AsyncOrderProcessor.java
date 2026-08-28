package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import com.fasterxml.jackson.core.JsonProcessingException;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.List;
import java.util.Optional;
import java.util.Set;
import java.util.function.Function;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(AsyncOrderProcessor.class.getName());
    private static final String DEAD_LETTER_REASON = "ORDER_DESERIALIZATION_FAILED";

    private final ServiceBusSessionReceiverAsyncClient activeReceiver;
    private final ServiceBusSessionReceiverAsyncClient deadLetterReceiver;

    public AsyncOrderProcessor(String fullyQualifiedNamespace, String queueName, TokenCredential credential) {
        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential);
        this.activeReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildAsyncClient();
        this.deadLetterReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildAsyncClient();
    }

    public Mono<Void> processSessions(Set<String> customerNames, Duration receiveWindow) {
        return Flux.fromIterable(customerNames)
                .concatMap(customer -> Flux.usingWhen(
                        activeReceiver.acceptSession(customer),
                        receiver -> receiver.receiveMessages()
                                .take(receiveWindow)
                                .concatMap(message -> process(receiver, message)),
                        receiver -> Mono.fromRunnable(receiver::close)))
                .then();
    }

    public Mono<List<DeadLetteredOrder>> inspectDeadLetters(
            Set<String> customerNames,
            Duration receiveWindow) {
        return Flux.fromIterable(customerNames)
                .concatMap(customer -> Flux.usingWhen(
                        deadLetterReceiver.acceptSession(customer),
                        receiver -> receiver.receiveMessages()
                                .take(receiveWindow)
                                .concatMap(message -> {
                                    DeadLetteredOrder deadLetter = toDeadLetteredOrder(message);
                                    LOGGER.warning(() -> "Dead letter: " + deadLetter);
                                    return receiver.abandon(message).thenReturn(deadLetter);
                                }),
                        receiver -> Mono.fromRunnable(receiver::close)))
                .collectList();
    }

    public Mono<Long> reprocessDeadLetters(
            Set<String> customerNames,
            Duration receiveWindow,
            Function<DeadLetteredOrder, Optional<Order>> recovery,
            AsyncOrderSender sender) {
        return Flux.fromIterable(customerNames)
                .concatMap(customer -> Flux.usingWhen(
                        deadLetterReceiver.acceptSession(customer),
                        receiver -> receiver.receiveMessages()
                                .take(receiveWindow)
                                .concatMap(message -> {
                                    Optional<Order> recovered = recovery.apply(toDeadLetteredOrder(message));
                                    if (recovered.isEmpty()) {
                                        return receiver.abandon(message).then(Mono.just(false));
                                    }
                                    return sender.send(recovered.get())
                                            .then(receiver.complete(message))
                                            .thenReturn(true);
                                }),
                        receiver -> Mono.fromRunnable(receiver::close)))
                .filter(Boolean::booleanValue)
                .count();
    }

    private Mono<Void> process(ServiceBusReceiverAsyncClient receiver, ServiceBusReceivedMessage message) {
        try {
            Order order = OrderMessageMapper.fromMessage(message);
            LOGGER.info(() -> "Processing order " + order.orderId()
                    + " for " + order.customerName() + ": " + order);
            return receiver.complete(message)
                    .doOnSuccess(ignored -> LOGGER.info(() -> "Completed order " + order.orderId()));
        } catch (JsonProcessingException exception) {
            LOGGER.log(Level.WARNING,
                    "Dead-lettering order " + message.getCorrelationId() + " because it cannot be deserialized",
                    exception);
            return receiver.deadLetter(message, new DeadLetterOptions()
                    .setDeadLetterReason(DEAD_LETTER_REASON)
                    .setDeadLetterErrorDescription(safeDescription(exception)));
        }
    }

    private static DeadLetteredOrder toDeadLetteredOrder(ServiceBusReceivedMessage message) {
        return new DeadLetteredOrder(
                message.getCorrelationId(),
                message.getSessionId(),
                message.getBody().toString(),
                message.getDeadLetterReason(),
                message.getDeadLetterErrorDescription());
    }

    private static String safeDescription(JsonProcessingException exception) {
        String message = exception.getMessage();
        return message == null ? exception.getClass().getSimpleName() : message.substring(0, Math.min(4096, message.length()));
    }

    @Override
    public void close() {
        activeReceiver.close();
        deadLetterReceiver.close();
    }
}

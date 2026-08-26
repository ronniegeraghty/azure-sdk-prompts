package com.example.orders.processing;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import com.example.orders.codec.OrderJsonCodec;
import com.example.orders.messaging.AsyncOrderSender;
import com.example.orders.model.Order;
import com.example.orders.model.OrderStatus;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.concurrent.TimeoutException;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);

    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;
    private final ServiceBusSessionReceiverAsyncClient deadLetterSessionReceiver;
    private final OrderJsonCodec codec = new OrderJsonCodec();

    public AsyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {
        this.sessionReceiver = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sessionReceiver()
            .queueName(queueName)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .buildAsyncClient();
        this.deadLetterSessionReceiver = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sessionReceiver()
            .queueName(queueName)
            .subQueue(SubQueue.DEAD_LETTER_QUEUE)
            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
            .disableAutoComplete()
            .buildAsyncClient();
    }

    public Mono<Void> processSessions(int sessionCount, Duration idleTimeout) {
        return Flux.range(0, sessionCount)
            .concatMap(ignored -> sessionReceiver.acceptNextSession()
                .flatMap(receiver -> processSession(receiver, idleTimeout)))
            .then();
    }

    public Mono<Void> inspectAndReprocessDeadLetters(
        int sessionCount,
        Duration idleTimeout,
        AsyncOrderSender sender
    ) {
        return Flux.range(0, sessionCount)
            .concatMap(ignored -> deadLetterSessionReceiver.acceptNextSession()
                .flatMap(receiver -> inspectSession(receiver, idleTimeout, sender)))
            .then();
    }

    private Mono<Void> processSession(
        ServiceBusReceiverAsyncClient receiver,
        Duration idleTimeout
    ) {
        LOGGER.info("Accepted customer session {}", receiver.getSessionId());
        return withIdleTimeout(
            receiver.receiveMessages().concatMap(message -> processMessage(receiver, message)),
            idleTimeout
        ).then().doFinally(ignored -> receiver.close());
    }

    private Mono<Void> processMessage(
        ServiceBusReceiverAsyncClient receiver,
        ServiceBusReceivedMessage message
    ) {
        try {
            OrderProcessingSupport.process(message, codec, LOGGER);
            return receiver.complete(message);
        } catch (RuntimeException exception) {
            LOGGER.error("Dead-lettering message {}: {}", message.getMessageId(), exception.getMessage());
            return receiver.deadLetter(message, OrderProcessingSupport.deadLetterOptions(exception));
        }
    }

    private Mono<Void> inspectSession(
        ServiceBusReceiverAsyncClient receiver,
        Duration idleTimeout,
        AsyncOrderSender sender
    ) {
        return withIdleTimeout(
            receiver.receiveMessages().concatMap(message -> reprocessDeadLetter(receiver, message, sender)),
            idleTimeout
        ).then().doFinally(ignored -> receiver.close());
    }

    private Mono<Void> reprocessDeadLetter(
        ServiceBusReceiverAsyncClient receiver,
        ServiceBusReceivedMessage message,
        AsyncOrderSender sender
    ) {
        LOGGER.warn("Inspecting dead-lettered message {}: reason={}, description={}",
            message.getMessageId(), message.getDeadLetterReason(), message.getDeadLetterErrorDescription());

        final Order order;
        try {
            order = codec.deserialize(message.getBody().toString());
            order.setStatus(OrderStatus.PENDING);
        } catch (RuntimeException exception) {
            LOGGER.error("Message {} cannot be deserialized and remains in the dead-letter queue",
                message.getMessageId(), exception);
            return receiver.abandon(message);
        }

        return sender.sendOrder(order)
            .then(receiver.complete(message))
            .doOnSuccess(ignored -> LOGGER.info("Re-enqueued dead-lettered order {}", order.getOrderId()))
            .onErrorResume(exception -> receiver.abandon(message)
                .then(Mono.error(new IllegalStateException(
                    "Could not reprocess dead-lettered order " + order.getOrderId(),
                    exception
                ))));
    }

    private Flux<Void> withIdleTimeout(Flux<Void> messages, Duration idleTimeout) {
        return messages.timeout(idleTimeout)
            .onErrorResume(TimeoutException.class, ignored -> Flux.empty());
    }

    @Override
    public void close() {
        deadLetterSessionReceiver.close();
        sessionReceiver.close();
    }
}

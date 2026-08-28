package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.SubQueue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.function.Function;

public final class AsyncOrderProcessor {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);
    private static final String DESERIALIZATION_FAILURE = "OrderDeserializationFailed";

    private final String namespace;
    private final String queueName;
    private final TokenCredential credential;
    private final OrderMessageFactory messageFactory;

    public AsyncOrderProcessor(String namespace, String queueName, TokenCredential credential) {
        this.namespace = namespace;
        this.queueName = queueName;
        this.credential = credential;
        this.messageFactory = new OrderMessageFactory(java.math.BigDecimal.ZERO);
    }

    public Mono<Void> processCustomer(String customerName, int maxMessages, Duration waitTime) {
        ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(false);
        return sessions.acceptSession(customerName)
            .flatMap(receiver -> receiver.receiveMessages()
                .take(maxMessages)
                .take(waitTime)
                .concatMap(message -> processMessage(receiver, message))
                .then()
                .doFinally(signal -> receiver.close()))
            .doFinally(signal -> sessions.close());
    }

    public Mono<Void> inspectDeadLetters(String customerName, int maxMessages, Duration waitTime) {
        ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);
        return sessions.acceptSession(customerName)
            .flatMap(receiver -> receiver.receiveMessages()
                .take(maxMessages)
                .take(waitTime)
                .concatMap(message -> {
                    LOGGER.warn("DLQ message id={}, reason={}, description={}, body={}",
                        message.getMessageId(), message.getDeadLetterReason(),
                        message.getDeadLetterErrorDescription(), message.getBody());
                    return receiver.abandon(message);
                })
                .then()
                .doFinally(signal -> receiver.close()))
            .doFinally(signal -> sessions.close());
    }

    public Mono<Void> reprocessDeadLetters(String customerName, int maxMessages, Duration waitTime,
                                           Function<ServiceBusReceivedMessage, Order> repair,
                                           AsyncOrderSender sender) {
        ServiceBusSessionReceiverAsyncClient sessions = sessionReceiver(true);
        return sessions.acceptSession(customerName)
            .flatMap(receiver -> receiver.receiveMessages()
                .take(maxMessages)
                .take(waitTime)
                .concatMap(message -> Mono.fromCallable(() -> (Order) repair.apply(message))
                    .flatMap(sender::sendOrder)
                    .then(receiver.complete(message))
                    .doOnSuccess(ignored ->
                        LOGGER.info("Requeued dead-letter message {}", message.getMessageId()))
                    .onErrorResume(error -> {
                        LOGGER.error("Could not reprocess dead-letter message {}",
                            message.getMessageId(), error);
                        return receiver.abandon(message);
                    }))
                .then()
                .doFinally(signal -> receiver.close()))
            .doFinally(signal -> sessions.close());
    }

    private Mono<Void> processMessage(ServiceBusReceiverAsyncClient receiver,
                                      ServiceBusReceivedMessage message) {
        return Mono.defer(() -> {
            try {
                Order order = messageFactory.deserialize(message.getBody());
                order.setStatus(Order.Status.processing);
                LOGGER.info("Asynchronously processed order: {}", order);
                order.setStatus(Order.Status.completed);
                return receiver.complete(message);
            } catch (RuntimeException error) {
                LOGGER.error("Dead-lettering message {}: {}",
                    message.getMessageId(), error.getMessage());
                return receiver.deadLetter(message, new DeadLetterOptions()
                    .setDeadLetterReason(DESERIALIZATION_FAILURE)
                    .setDeadLetterErrorDescription(error.getMessage()));
            }
        });
    }

    private ServiceBusSessionReceiverAsyncClient sessionReceiver(boolean deadLetterQueue) {
        ServiceBusClientBuilder.ServiceBusSessionReceiverClientBuilder builder =
            new ServiceBusClientBuilder()
                .credential(namespace, credential)
                .sessionReceiver()
                .queueName(queueName)
                .disableAutoComplete();
        if (deadLetterQueue) {
            builder.subQueue(SubQueue.DEAD_LETTER_QUEUE);
        }
        return builder.buildAsyncClient();
    }
}

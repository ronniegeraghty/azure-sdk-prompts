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
import java.util.concurrent.TimeoutException;
import java.util.function.Function;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);

    private final ServiceBusSessionReceiverAsyncClient receiver;
    private final ServiceBusSessionReceiverAsyncClient deadLetterReceiver;

    public AsyncOrderProcessor(
            String fullyQualifiedNamespace, String queueName, TokenCredential credential) {
        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential);
        this.receiver = builder.sessionReceiver()
                .queueName(queueName)
                .disableAutoComplete()
                .buildAsyncClient();
        this.deadLetterReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .disableAutoComplete()
                .buildAsyncClient();
    }

    public Mono<Void> processAvailableSessions(int maximumSessions, Duration sessionWaitTime) {
        return processSessions(receiver, maximumSessions, sessionWaitTime, this::processMessage);
    }

    public Mono<Void> inspectAndReprocessDeadLetters(
            int maximumSessions,
            Duration sessionWaitTime,
            Function<Order, Mono<Void>> resubmitter) {
        return processSessions(deadLetterReceiver, maximumSessions, sessionWaitTime, (session, message) -> {
            LOGGER.warn("Dead letter: messageId={}, reason={}, description={}",
                    message.getMessageId(),
                    message.getDeadLetterReason(),
                    message.getDeadLetterErrorDescription());
            final Order order;
            try {
                order = OrderJson.deserialize(message.getBody().toString());
            } catch (Exception exception) {
                LOGGER.error("Dead-letter message {} could not be deserialized",
                        message.getMessageId(), exception);
                return session.abandon(message);
            }
            return resubmitter.apply(order)
                    .then(session.complete(message))
                    .doOnSuccess(ignored -> LOGGER.info("Requeued order {}", order.getOrderId()));
        });
    }

    private Mono<Void> processMessage(
            ServiceBusReceiverAsyncClient session, ServiceBusReceivedMessage message) {
        try {
            Order order = OrderJson.deserialize(message.getBody().toString());
            if (order.getStatus() == Order.Status.FAILED) {
                throw new IllegalStateException("Order arrived with FAILED status");
            }
            LOGGER.info("Processing {}", order.withStatus(Order.Status.PROCESSING));
            LOGGER.info("Completed {}", order.withStatus(Order.Status.COMPLETED));
            return session.complete(message);
        } catch (Exception exception) {
            DeadLetterOptions options = new DeadLetterOptions()
                    .setDeadLetterReason("ORDER_PROCESSING_FAILED")
                    .setDeadLetterErrorDescription(exception.getMessage());
            LOGGER.error("Dead-lettering message {}", message.getMessageId(), exception);
            return session.deadLetter(message, options);
        }
    }

    private Mono<Void> processSessions(
            ServiceBusSessionReceiverAsyncClient sessionReceiver,
            int remainingSessions,
            Duration sessionWaitTime,
            AsyncSessionMessageHandler handler) {
        if (remainingSessions == 0) {
            return Mono.empty();
        }
        return sessionReceiver.acceptNextSession()
                .timeout(sessionWaitTime)
                .onErrorResume(TimeoutException.class, ignored -> Mono.empty())
                .flatMap(session -> session.receiveMessages()
                        .concatMap(message -> handler.handle(session, message))
                        .timeout(sessionWaitTime)
                        .onErrorResume(TimeoutException.class, ignored -> Mono.empty())
                        .then()
                        .doFinally(signal -> session.close()))
                .then(Mono.defer(() -> processSessions(
                        sessionReceiver, remainingSessions - 1, sessionWaitTime, handler)));
    }

    @Override
    public void close() {
        receiver.close();
        deadLetterReceiver.close();
    }

    @FunctionalInterface
    private interface AsyncSessionMessageHandler {
        Mono<Void> handle(ServiceBusReceiverAsyncClient session, ServiceBusReceivedMessage message);
    }
}

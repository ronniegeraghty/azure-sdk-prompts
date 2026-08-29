package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverAsyncClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.Optional;
import java.util.function.Function;
import java.util.logging.Logger;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(AsyncOrderProcessor.class.getName());
    private final ServiceBusSessionReceiverAsyncClient sessionReceiver;
    private final ServiceBusSessionReceiverAsyncClient deadLetterSessionReceiver;

    public AsyncOrderProcessor(
            String fullyQualifiedNamespace,
            String queueName,
            TokenCredential credential) {
        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential);
        this.sessionReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildAsyncClient();
        this.deadLetterSessionReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildAsyncClient();
    }

    public Mono<Void> processAvailableSessions(int sessionCount, Duration processingWindow) {
        return Flux.range(0, sessionCount)
                .flatMap(ignored -> Mono.usingWhen(
                        sessionReceiver.acceptNextSession(),
                        receiver -> receiver.receiveMessages()
                                .concatMap(message -> processMessage(receiver, message))
                                .take(processingWindow)
                                .then(),
                        receiver -> Mono.fromRunnable(receiver::close)), sessionCount)
                .then();
    }

    private Mono<Void> processMessage(
            ServiceBusReceiverAsyncClient receiver,
            ServiceBusReceivedMessage message) {
        try {
            Order pending = OrderJson.deserialize(message.getBody().toString());
            Order processing = pending.withStatus(OrderStatus.PROCESSING);
            LOGGER.info(() -> "Processing async order: " + OrderJson.serialize(processing));
            Order completed = processing.withStatus(OrderStatus.COMPLETED);
            LOGGER.info(() -> "Completed async order: " + OrderJson.serialize(completed));
            return receiver.complete(message);
        } catch (IllegalArgumentException exception) {
            LOGGER.warning(() -> "Dead-lettering order message " + message.getMessageId()
                    + ": " + exception.getMessage());
            return receiver.deadLetter(message, new DeadLetterOptions()
                    .setDeadLetterReason("ORDER_DESERIALIZATION_FAILED")
                    .setDeadLetterErrorDescription(exception.getMessage()));
        }
    }

    public Mono<Void> reprocessDeadLetters(
            int maximumMessages,
            Duration receiveWindow,
            Function<ServiceBusReceivedMessage, Optional<Order>> recovery,
            AsyncOrderSender sender) {
        return Mono.usingWhen(
                deadLetterSessionReceiver.acceptNextSession(),
                receiver -> receiver.receiveMessages()
                        .take(receiveWindow)
                        .take(maximumMessages)
                        .concatMap(message -> {
                            LOGGER.warning(() -> "Inspecting dead letter: id=" + message.getMessageId()
                                    + ", reason=" + message.getDeadLetterReason()
                                    + ", description=" + message.getDeadLetterErrorDescription());
                            Optional<Order> recovered = recovery.apply(message);
                            if (recovered.isEmpty()) {
                                return receiver.abandon(message);
                            }
                            return sender.sendOrder(recovered.get())
                                    .then(receiver.complete(message))
                                    .doOnSuccess(ignored -> LOGGER.info(
                                            () -> "Requeued recovered order " + recovered.get().orderId()));
                        })
                        .then(),
                receiver -> Mono.fromRunnable(receiver::close));
    }

    @Override
    public void close() {
        sessionReceiver.close();
        deadLetterSessionReceiver.close();
    }
}

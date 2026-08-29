package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;

import java.time.Duration;
import java.util.Optional;
import java.util.function.Function;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class SyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(SyncOrderProcessor.class.getName());
    private final String fullyQualifiedNamespace;
    private final String queueName;
    private final TokenCredential credential;
    private final ServiceBusProcessorClient processor;

    public SyncOrderProcessor(
            String fullyQualifiedNamespace,
            String queueName,
            TokenCredential credential) {
        this.fullyQualifiedNamespace = fullyQualifiedNamespace;
        this.queueName = queueName;
        this.credential = credential;
        this.processor = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential)
                .sessionProcessor()
                .queueName(queueName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .maxConcurrentSessions(4)
                .maxConcurrentCalls(1)
                .processMessage(this::processMessage)
                .processError(context -> LOGGER.log(
                        Level.SEVERE,
                        "Service Bus processor error for " + context.getEntityPath(),
                        context.getException()))
                .buildProcessorClient();
    }

    public void start() {
        processor.start();
    }

    public void stop() {
        processor.stop();
    }

    private void processMessage(ServiceBusReceivedMessageContext context) {
        ServiceBusReceivedMessage message = context.getMessage();
        try {
            Order pending = OrderJson.deserialize(message.getBody().toString());
            Order processing = pending.withStatus(OrderStatus.PROCESSING);
            LOGGER.info(() -> "Processing sync order: " + OrderJson.serialize(processing));
            Order completed = processing.withStatus(OrderStatus.COMPLETED);
            LOGGER.info(() -> "Completed sync order: " + OrderJson.serialize(completed));
            context.complete();
        } catch (IllegalArgumentException exception) {
            String reason = "ORDER_DESERIALIZATION_FAILED";
            LOGGER.log(Level.WARNING, "Dead-lettering order message " + message.getMessageId(), exception);
            context.deadLetter(new DeadLetterOptions()
                    .setDeadLetterReason(reason)
                    .setDeadLetterErrorDescription(exception.getMessage()));
        }
    }

    public void reprocessDeadLetters(
            int maximumMessages,
            Duration receiveWindow,
            Function<ServiceBusReceivedMessage, Optional<Order>> recovery,
            SyncOrderSender sender) {
        try (ServiceBusSessionReceiverClient sessionReceiver = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential)
                .sessionReceiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildClient();
             ServiceBusReceiverClient receiver = sessionReceiver.acceptNextSession()) {
            for (ServiceBusReceivedMessage message
                    : receiver.receiveMessages(maximumMessages, receiveWindow)) {
                LOGGER.warning(() -> "Inspecting dead letter: id=" + message.getMessageId()
                        + ", reason=" + message.getDeadLetterReason()
                        + ", description=" + message.getDeadLetterErrorDescription());
                Optional<Order> recovered = recovery.apply(message);
                if (recovered.isPresent()) {
                    sender.sendOrder(recovered.get());
                    receiver.complete(message);
                    LOGGER.info(() -> "Requeued recovered order " + recovered.get().orderId());
                } else {
                    receiver.abandon(message);
                }
            }
        }
    }

    @Override
    public void close() {
        processor.close();
    }
}

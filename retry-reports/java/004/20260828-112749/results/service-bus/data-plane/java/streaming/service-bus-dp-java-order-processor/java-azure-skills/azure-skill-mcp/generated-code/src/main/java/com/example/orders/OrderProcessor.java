package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.ServiceBusReceiverClient;
import com.azure.messaging.servicebus.ServiceBusSessionReceiverClient;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
import com.azure.messaging.servicebus.models.SubQueue;
import com.fasterxml.jackson.core.JsonProcessingException;

import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.Set;
import java.util.function.Function;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class OrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(OrderProcessor.class.getName());
    private static final String DEAD_LETTER_REASON = "ORDER_DESERIALIZATION_FAILED";

    private final ServiceBusSessionReceiverClient activeReceiver;
    private final ServiceBusSessionReceiverClient deadLetterReceiver;

    public OrderProcessor(String fullyQualifiedNamespace, String queueName, TokenCredential credential) {
        ServiceBusClientBuilder builder = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential);
        this.activeReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildClient();
        this.deadLetterReceiver = builder.sessionReceiver()
                .queueName(queueName)
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
                .disableAutoComplete()
                .buildClient();
    }

    public void processSessions(Set<String> customerNames, Duration waitPerReceive) {
        for (String customerName : customerNames) {
            try (ServiceBusReceiverClient receiver = activeReceiver.acceptSession(customerName)) {
                boolean received;
                do {
                    List<ServiceBusReceivedMessage> messages =
                            receiver.receiveMessages(100, waitPerReceive).stream().toList();
                    received = !messages.isEmpty();
                    messages.forEach(message -> process(receiver, message));
                } while (received);
            }
        }
    }

    public List<DeadLetteredOrder> inspectDeadLetters(Set<String> customerNames, Duration waitPerReceive) {
        List<DeadLetteredOrder> deadLetters = new ArrayList<>();
        for (String customerName : customerNames) {
            try (ServiceBusReceiverClient receiver = deadLetterReceiver.acceptSession(customerName)) {
                for (ServiceBusReceivedMessage message : receiver.receiveMessages(100, waitPerReceive)) {
                    DeadLetteredOrder deadLetter = toDeadLetteredOrder(message);
                    deadLetters.add(deadLetter);
                    LOGGER.warning(() -> "Dead letter: " + deadLetter);
                    receiver.abandon(message);
                }
            }
        }
        return deadLetters;
    }

    public int reprocessDeadLetters(
            Set<String> customerNames,
            Duration waitPerReceive,
            Function<DeadLetteredOrder, Optional<Order>> recovery,
            OrderSender sender) {
        int reprocessed = 0;
        for (String customerName : customerNames) {
            try (ServiceBusReceiverClient receiver = deadLetterReceiver.acceptSession(customerName)) {
                for (ServiceBusReceivedMessage message : receiver.receiveMessages(100, waitPerReceive)) {
                    Optional<Order> recovered = recovery.apply(toDeadLetteredOrder(message));
                    if (recovered.isPresent()) {
                        sender.send(recovered.get());
                        receiver.complete(message);
                        reprocessed++;
                    } else {
                        receiver.abandon(message);
                    }
                }
            }
        }
        return reprocessed;
    }

    private void process(ServiceBusReceiverClient receiver, ServiceBusReceivedMessage message) {
        try {
            Order order = OrderMessageMapper.fromMessage(message);
            LOGGER.info(() -> "Processing order " + order.orderId()
                    + " for " + order.customerName() + ": " + order);
            LOGGER.info(() -> "Completed order " + order.orderId());
            receiver.complete(message);
        } catch (JsonProcessingException exception) {
            LOGGER.log(Level.WARNING,
                    "Dead-lettering order " + message.getCorrelationId() + " because it cannot be deserialized",
                    exception);
            receiver.deadLetter(message, new DeadLetterOptions()
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

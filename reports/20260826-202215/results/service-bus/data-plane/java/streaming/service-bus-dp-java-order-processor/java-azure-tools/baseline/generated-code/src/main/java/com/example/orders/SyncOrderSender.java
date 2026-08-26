package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public final class SyncOrderSender implements AutoCloseable {
    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;

    private final ServiceBusSenderClient sender;
    private final BigDecimal highPriorityThreshold;

    public SyncOrderSender(
            String fullyQualifiedNamespace,
            String queueName,
            TokenCredential credential,
            BigDecimal highPriorityThreshold) {
        this.sender = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential)
                .sender()
                .queueName(queueName)
                .buildClient();
        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, "highPriorityThreshold");
    }

    public void send(Order order) {
        ServiceBusMessage message = toMessage(order);
        if (isHighPriority(order)) {
            sender.scheduleMessage(message, OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));
        } else {
            sender.sendMessage(message);
        }
    }

    public void sendBatch(List<Order> orders) {
        List<ServiceBusMessage> immediateMessages = new ArrayList<>();
        OffsetDateTime scheduledTime = OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS);

        for (Order order : orders) {
            ServiceBusMessage message = toMessage(order);
            if (isHighPriority(order)) {
                sender.scheduleMessage(message, scheduledTime);
            } else {
                immediateMessages.add(message);
            }
        }

        sendSizeAwareBatches(immediateMessages);
    }

    private void sendSizeAwareBatches(List<ServiceBusMessage> messages) {
        ServiceBusMessageBatch batch = sender.createMessageBatch();
        for (ServiceBusMessage message : messages) {
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    throw new IllegalArgumentException("Order message exceeds the Service Bus maximum message size");
                }
                sender.sendMessages(batch);
                batch = sender.createMessageBatch();
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException("Order message exceeds the Service Bus maximum message size");
                }
            }
        }
        if (batch.getCount() > 0) {
            sender.sendMessages(batch);
        }
    }

    private ServiceBusMessage toMessage(Order order) {
        return OrderMessageFactory.create(order, highPriorityThreshold);
    }

    private boolean isHighPriority(Order order) {
        return order.getTotalPrice().compareTo(highPriorityThreshold) > 0;
    }

    @Override
    public void close() {
        sender.close();
    }
}

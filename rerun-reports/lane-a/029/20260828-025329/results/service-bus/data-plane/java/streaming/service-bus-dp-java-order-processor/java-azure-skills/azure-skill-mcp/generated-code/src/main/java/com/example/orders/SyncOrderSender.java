package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Objects;

public final class SyncOrderSender implements AutoCloseable {
    private final ServiceBusSenderClient sender;
    private final OrderMessageFactory messageFactory;

    public SyncOrderSender(
            ServiceBusSenderClient sender,
            ObjectMapper objectMapper,
            BigDecimal highPriorityThreshold) {
        this.sender = Objects.requireNonNull(sender, "sender");
        this.messageFactory = new OrderMessageFactory(objectMapper, highPriorityThreshold);
    }

    public void send(Order order) {
        ServiceBusMessage message = messageFactory.create(order, false);
        if (messageFactory.isHighPriority(order)) {
            sender.scheduleMessage(
                    message,
                    OffsetDateTime.now().plus(OrderMessageFactory.FRAUD_REVIEW_DELAY));
        } else {
            sender.sendMessage(message);
        }
    }

    public void sendBatch(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        if (orders.isEmpty()) {
            return;
        }

        ServiceBusMessageBatch batch = sender.createMessageBatch();
        for (Order order : orders) {
            ServiceBusMessage message = messageFactory.create(order, true);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    throw new IllegalArgumentException(
                            "Order " + order.getOrderId() + " exceeds the Service Bus maximum message size");
                }
                sender.sendMessages(batch);
                batch = sender.createMessageBatch();
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException(
                            "Order " + order.getOrderId() + " exceeds the Service Bus maximum message size");
                }
            }
        }
        if (batch.getCount() > 0) {
            sender.sendMessages(batch);
        }
    }

    @Override
    public void close() {
        sender.close();
    }
}

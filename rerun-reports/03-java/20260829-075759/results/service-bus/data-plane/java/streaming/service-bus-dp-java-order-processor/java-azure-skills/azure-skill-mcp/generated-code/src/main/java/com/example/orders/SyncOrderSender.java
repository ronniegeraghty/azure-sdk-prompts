package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;

import java.math.BigDecimal;
import java.util.List;
import java.util.Objects;

public final class SyncOrderSender implements AutoCloseable {
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
        this.highPriorityThreshold = Objects.requireNonNull(
                highPriorityThreshold, "highPriorityThreshold");
    }

    public void sendOrder(Order order) {
        sender.sendMessage(OrderMessageFactory.create(order, highPriorityThreshold));
    }

    public void sendOrders(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        ServiceBusMessageBatch batch = sender.createMessageBatch();

        for (Order order : orders) {
            ServiceBusMessage message = OrderMessageFactory.create(order, highPriorityThreshold);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    throw new IllegalArgumentException(
                            "Order message is too large for an empty Service Bus batch: " + order.orderId());
                }
                sender.sendMessages(batch);
                batch = sender.createMessageBatch();
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException(
                            "Order message is too large for an empty Service Bus batch: " + order.orderId());
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

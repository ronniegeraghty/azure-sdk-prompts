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
    private final OrderMessageFactory messageFactory;

    public SyncOrderSender(String namespace, String queueName, TokenCredential credential,
                           BigDecimal highPriorityThreshold) {
        sender = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sender()
            .queueName(queueName)
            .buildClient();
        messageFactory = new OrderMessageFactory(highPriorityThreshold);
    }

    public void sendOrder(Order order) {
        sender.sendMessage(messageFactory.create(order));
    }

    public void sendOrders(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        if (orders.isEmpty()) {
            return;
        }

        ServiceBusMessageBatch batch = sender.createMessageBatch();
        for (Order order : orders) {
            ServiceBusMessage message = messageFactory.create(order);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    throw new IllegalArgumentException("Order " + order.getOrderId()
                        + " is too large for a Service Bus batch");
                }
                sender.sendMessages(batch);
                batch = sender.createMessageBatch();
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException("Order " + order.getOrderId()
                        + " is too large for a Service Bus batch");
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

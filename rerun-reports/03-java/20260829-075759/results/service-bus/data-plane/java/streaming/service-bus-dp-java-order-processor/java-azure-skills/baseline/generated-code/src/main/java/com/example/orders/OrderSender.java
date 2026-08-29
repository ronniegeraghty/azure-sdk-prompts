package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;

import java.util.List;
import java.util.Objects;

public final class OrderSender implements AutoCloseable {
    private final ServiceBusSenderClient client;
    private final OrderMessageFactory messageFactory;

    public OrderSender(ServiceBusSenderClient client, OrderMessageFactory messageFactory) {
        this.client = Objects.requireNonNull(client, "client");
        this.messageFactory = Objects.requireNonNull(messageFactory, "messageFactory");
    }

    public void send(Order order) {
        client.sendMessage(messageFactory.create(order));
    }

    public void sendBatch(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        ServiceBusMessageBatch batch = client.createMessageBatch();

        for (Order order : orders) {
            ServiceBusMessage message = messageFactory.create(order);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    throw new IllegalArgumentException("Order " + order.getOrderId()
                            + " exceeds the maximum Service Bus message size");
                }
                client.sendMessages(batch);
                batch = client.createMessageBatch();
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException("Order " + order.getOrderId()
                            + " exceeds the maximum Service Bus message size");
                }
            }
        }

        if (batch.getCount() > 0) {
            client.sendMessages(batch);
        }
    }

    void sendMalformedDemoMessage(String messageId, String customerName) {
        client.sendMessage(OrderMessageFactory.malformedDemoMessage(messageId, customerName));
    }

    @Override
    public void close() {
        client.close();
    }
}

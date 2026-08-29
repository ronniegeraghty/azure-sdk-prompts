package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

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
        sender.sendMessage(messageFactory.create(order));
    }

    public void sendBatch(List<Order> orders) {
        Map<String, List<Order>> ordersByCustomer = orders.stream()
                .collect(Collectors.groupingBy(
                        Order::getCustomerName,
                        java.util.LinkedHashMap::new,
                        Collectors.toList()));

        for (List<Order> customerOrders : ordersByCustomer.values()) {
            sendCustomerBatches(customerOrders);
        }
    }

    private void sendCustomerBatches(List<Order> orders) {
        ServiceBusMessageBatch batch = sender.createMessageBatch();
        for (Order order : orders) {
            ServiceBusMessage message = messageFactory.create(order);
            if (batch.tryAddMessage(message)) {
                continue;
            }
            sender.sendMessages(batch);
            batch = sender.createMessageBatch();
            if (!batch.tryAddMessage(message)) {
                throw new IllegalArgumentException(
                        "Order %s exceeds the maximum Service Bus message size".formatted(order.getOrderId()));
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

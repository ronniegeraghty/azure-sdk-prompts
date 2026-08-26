package com.example.orders.messaging;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.example.orders.model.Order;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.math.BigDecimal;
import java.util.List;
import java.util.Objects;

public final class SyncOrderSender implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncOrderSender.class);

    private final ServiceBusSenderClient sender;
    private final OrderMessageFactory messageFactory;

    public SyncOrderSender(
        String fullyQualifiedNamespace,
        String queueName,
        TokenCredential credential,
        BigDecimal highPriorityThreshold
    ) {
        this.sender = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .sender()
            .queueName(queueName)
            .buildClient();
        this.messageFactory = new OrderMessageFactory(
            new com.example.orders.codec.OrderJsonCodec(),
            highPriorityThreshold
        );
    }

    public void sendOrder(Order order) {
        ServiceBusMessage message = messageFactory.createMessage(order);
        sender.sendMessage(message);
        LOGGER.info("Sent order {} with priority {}", order.getOrderId(),
            message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));
    }

    public void sendOrders(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        if (orders.isEmpty()) {
            return;
        }

        ServiceBusMessageBatch batch = sender.createMessageBatch();
        for (Order order : orders) {
            ServiceBusMessage message = messageFactory.createMessage(order);
            if (batch.tryAddMessage(message)) {
                continue;
            }

            sendBatch(batch);
            batch = sender.createMessageBatch();
            if (!batch.tryAddMessage(message)) {
                throw new IllegalArgumentException(
                    "Order " + order.getOrderId() + " exceeds the maximum Service Bus message size"
                );
            }
        }

        if (batch.getCount() > 0) {
            sendBatch(batch);
        }
    }

    private void sendBatch(ServiceBusMessageBatch batch) {
        sender.sendMessages(batch);
        LOGGER.info("Sent batch containing {} order(s) and {} byte(s)",
            batch.getCount(), batch.getSizeInBytes());
    }

    @Override
    public void close() {
        sender.close();
    }
}

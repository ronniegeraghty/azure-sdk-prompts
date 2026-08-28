package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderClient;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public final class OrderSender implements AutoCloseable {
    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;

    private final ServiceBusSenderClient sender;
    private final BigDecimal highPriorityThreshold;

    public OrderSender(
            String fullyQualifiedNamespace,
            String queueName,
            TokenCredential credential,
            BigDecimal highPriorityThreshold) {
        this.sender = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential)
                .sender()
                .queueName(queueName)
                .buildClient();
        this.highPriorityThreshold = highPriorityThreshold;
    }

    public void send(Order order) {
        ServiceBusMessage message = messageFor(order);
        if (isHighPriority(order)) {
            sender.scheduleMessage(message, OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));
        } else {
            sender.sendMessage(message);
        }
    }

    public void sendBatch(List<Order> orders) {
        Map<String, List<Order>> immediateByCustomer = new LinkedHashMap<>();
        for (Order order : orders) {
            if (isHighPriority(order)) {
                sender.scheduleMessage(
                        messageFor(order),
                        OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));
            } else {
                immediateByCustomer
                        .computeIfAbsent(order.customerName(), ignored -> new ArrayList<>())
                        .add(order);
            }
        }
        immediateByCustomer.values().forEach(this::sendImmediateBatches);
    }

    public void sendMalformedForDemo(String correlationId, String customerName) {
        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString("{not-valid-json"))
                .setContentType("application/json")
                .setMessageId(correlationId)
                .setCorrelationId(correlationId)
                .setSessionId(customerName);
        sender.sendMessage(message);
    }

    private void sendImmediateBatches(List<Order> orders) {
        ServiceBusMessageBatch batch = sender.createMessageBatch();
        for (Order order : orders) {
            ServiceBusMessage message = messageFor(order);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    throw new IllegalArgumentException("Order is too large for an empty Service Bus batch: "
                            + order.orderId());
                }
                sender.sendMessages(batch);
                batch = sender.createMessageBatch();
                if (!batch.tryAddMessage(message)) {
                    throw new IllegalArgumentException("Order is too large for an empty Service Bus batch: "
                            + order.orderId());
                }
            }
        }
        if (batch.getCount() > 0) {
            sender.sendMessages(batch);
        }
    }

    private ServiceBusMessage messageFor(Order order) {
        return OrderMessageMapper.toMessage(order, isHighPriority(order));
    }

    private boolean isHighPriority(Order order) {
        return order.totalPrice().compareTo(highPriorityThreshold) > 0;
    }

    @Override
    public void close() {
        sender.close();
    }
}

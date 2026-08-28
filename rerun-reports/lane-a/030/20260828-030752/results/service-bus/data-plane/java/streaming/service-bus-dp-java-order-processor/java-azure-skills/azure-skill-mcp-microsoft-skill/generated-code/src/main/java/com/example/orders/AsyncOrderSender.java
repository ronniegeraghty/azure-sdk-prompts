package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.util.List;
import java.util.Objects;

public final class AsyncOrderSender implements AutoCloseable {
    private final ServiceBusSenderAsyncClient sender;
    private final OrderMessageFactory messageFactory;

    public AsyncOrderSender(String namespace, String queueName, TokenCredential credential,
                            BigDecimal highPriorityThreshold) {
        sender = new ServiceBusClientBuilder()
            .credential(namespace, credential)
            .sender()
            .queueName(queueName)
            .buildAsyncClient();
        messageFactory = new OrderMessageFactory(highPriorityThreshold);
    }

    public Mono<Void> sendOrder(Order order) {
        return sender.sendMessage(messageFactory.create(order));
    }

    public Mono<Void> sendOrders(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        return sendBatch(orders, 0);
    }

    private Mono<Void> sendBatch(List<Order> orders, int startIndex) {
        if (startIndex >= orders.size()) {
            return Mono.empty();
        }

        return sender.createMessageBatch().flatMap(batch -> {
            int nextIndex = fillBatch(batch, orders, startIndex);
            return sender.sendMessages(batch)
                .then(Mono.defer(() -> sendBatch(orders, nextIndex)));
        });
    }

    private int fillBatch(ServiceBusMessageBatch batch, List<Order> orders, int startIndex) {
        int index = startIndex;
        while (index < orders.size()) {
            Order order = orders.get(index);
            ServiceBusMessage message = messageFactory.create(order);
            if (!batch.tryAddMessage(message)) {
                if (index == startIndex) {
                    throw new IllegalArgumentException("Order " + order.getOrderId()
                        + " is too large for a Service Bus batch");
                }
                break;
            }
            index++;
        }
        return index;
    }

    @Override
    public void close() {
        sender.close();
    }
}

package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Objects;

public final class AsyncOrderSender implements AutoCloseable {
    private final ServiceBusSenderAsyncClient client;
    private final OrderMessageFactory messageFactory;

    public AsyncOrderSender(ServiceBusSenderAsyncClient client, OrderMessageFactory messageFactory) {
        this.client = Objects.requireNonNull(client, "client");
        this.messageFactory = Objects.requireNonNull(messageFactory, "messageFactory");
    }

    public Mono<Void> send(Order order) {
        return client.sendMessage(messageFactory.create(order));
    }

    public Mono<Void> sendBatch(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        return client.createMessageBatch()
                .flatMap(batch -> fillAndSend(orders, 0, batch));
    }

    private Mono<Void> fillAndSend(
            List<Order> orders, int startIndex, ServiceBusMessageBatch batch) {
        int index = startIndex;
        while (index < orders.size()) {
            Order order = orders.get(index);
            ServiceBusMessage message = messageFactory.create(order);
            if (batch.tryAddMessage(message)) {
                index++;
                continue;
            }
            if (batch.getCount() == 0) {
                return Mono.error(oversized(order));
            }
            int nextIndex = index;
            return client.sendMessages(batch)
                    .then(client.createMessageBatch())
                    .flatMap(nextBatch -> fillAndSend(orders, nextIndex, nextBatch));
        }
        return batch.getCount() == 0 ? Mono.empty() : client.sendMessages(batch);
    }

    private static IllegalArgumentException oversized(Order order) {
        return new IllegalArgumentException("Order " + order.getOrderId()
                + " exceeds the maximum Service Bus message size");
    }

    Mono<Void> sendMalformedDemoMessage(String messageId, String customerName) {
        return client.sendMessage(OrderMessageFactory.malformedDemoMessage(messageId, customerName));
    }

    @Override
    public void close() {
        client.close();
    }
}

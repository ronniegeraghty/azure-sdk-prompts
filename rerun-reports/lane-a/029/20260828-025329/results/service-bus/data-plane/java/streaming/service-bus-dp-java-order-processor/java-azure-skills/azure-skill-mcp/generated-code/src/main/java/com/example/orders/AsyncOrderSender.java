package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.fasterxml.jackson.databind.ObjectMapper;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Objects;

public final class AsyncOrderSender implements AutoCloseable {
    private final ServiceBusSenderAsyncClient sender;
    private final OrderMessageFactory messageFactory;

    public AsyncOrderSender(
            ServiceBusSenderAsyncClient sender,
            ObjectMapper objectMapper,
            BigDecimal highPriorityThreshold) {
        this.sender = Objects.requireNonNull(sender, "sender");
        this.messageFactory = new OrderMessageFactory(objectMapper, highPriorityThreshold);
    }

    public Mono<Void> send(Order order) {
        ServiceBusMessage message = messageFactory.create(order, false);
        if (messageFactory.isHighPriority(order)) {
            return sender.scheduleMessage(
                            message,
                            OffsetDateTime.now().plus(OrderMessageFactory.FRAUD_REVIEW_DELAY))
                    .then();
        }
        return sender.sendMessage(message);
    }

    public Mono<Void> sendBatch(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        if (orders.isEmpty()) {
            return Mono.empty();
        }

        return sender.createMessageBatch()
                .flatMap(batch -> addAndSend(orders, 0, batch));
    }

    private Mono<Void> addAndSend(List<Order> orders, int index, ServiceBusMessageBatch batch) {
        for (int current = index; current < orders.size(); current++) {
            Order order = orders.get(current);
            ServiceBusMessage message = messageFactory.create(order, true);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    return Mono.error(new IllegalArgumentException(
                            "Order " + order.getOrderId() + " exceeds the Service Bus maximum message size"));
                }
                int nextIndex = current;
                return sender.sendMessages(batch)
                        .then(sender.createMessageBatch())
                        .flatMap(nextBatch -> addAndSend(orders, nextIndex, nextBatch));
            }
        }
        return batch.getCount() == 0 ? Mono.empty() : sender.sendMessages(batch);
    }

    @Override
    public void close() {
        sender.close();
    }
}

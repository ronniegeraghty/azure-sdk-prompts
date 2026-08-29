package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.fasterxml.jackson.databind.ObjectMapper;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

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
        return sender.sendMessage(messageFactory.create(order));
    }

    public Mono<Void> sendBatch(List<Order> orders) {
        Map<String, List<Order>> ordersByCustomer = orders.stream()
                .collect(Collectors.groupingBy(
                        Order::getCustomerName,
                        java.util.LinkedHashMap::new,
                        Collectors.toList()));

        return Flux.fromIterable(ordersByCustomer.values())
                .concatMap(this::sendCustomerBatches)
                .then();
    }

    private Mono<Void> sendCustomerBatches(List<Order> orders) {
        return sender.createMessageBatch().flatMap(batch -> fillAndSend(batch, orders, 0));
    }

    private Mono<Void> fillAndSend(ServiceBusMessageBatch batch, List<Order> orders, int index) {
        int nextIndex = index;
        while (nextIndex < orders.size()) {
            Order order = orders.get(nextIndex);
            ServiceBusMessage message = messageFactory.create(order);
            if (!batch.tryAddMessage(message)) {
                if (batch.getCount() == 0) {
                    return Mono.error(new IllegalArgumentException(
                            "Order %s exceeds the maximum Service Bus message size".formatted(order.getOrderId())));
                }
                int remainingIndex = nextIndex;
                return sender.sendMessages(batch)
                        .then(sender.createMessageBatch())
                        .flatMap(nextBatch -> fillAndSend(nextBatch, orders, remainingIndex));
            }
            nextIndex++;
        }
        return batch.getCount() == 0 ? Mono.empty() : sender.sendMessages(batch);
    }

    @Override
    public void close() {
        sender.close();
    }
}

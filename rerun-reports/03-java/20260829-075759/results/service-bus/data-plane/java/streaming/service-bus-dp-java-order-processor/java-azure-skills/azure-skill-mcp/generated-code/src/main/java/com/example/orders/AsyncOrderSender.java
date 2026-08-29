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
    private final BigDecimal highPriorityThreshold;

    public AsyncOrderSender(
            String fullyQualifiedNamespace,
            String queueName,
            TokenCredential credential,
            BigDecimal highPriorityThreshold) {
        this.sender = new ServiceBusClientBuilder()
                .credential(fullyQualifiedNamespace, credential)
                .sender()
                .queueName(queueName)
                .buildAsyncClient();
        this.highPriorityThreshold = Objects.requireNonNull(
                highPriorityThreshold, "highPriorityThreshold");
    }

    public Mono<Void> sendOrder(Order order) {
        return sender.sendMessage(OrderMessageFactory.create(order, highPriorityThreshold));
    }

    public Mono<Void> sendOrders(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        List<ServiceBusMessage> messages = orders.stream()
                .map(order -> OrderMessageFactory.create(order, highPriorityThreshold))
                .toList();
        return sendBatchFrom(messages, 0);
    }

    private Mono<Void> sendBatchFrom(List<ServiceBusMessage> messages, int startIndex) {
        if (startIndex >= messages.size()) {
            return Mono.empty();
        }

        return sender.createMessageBatch().flatMap(batch -> {
            int nextIndex = fillBatch(batch, messages, startIndex);
            return sender.sendMessages(batch).then(sendBatchFrom(messages, nextIndex));
        });
    }

    private static int fillBatch(
            ServiceBusMessageBatch batch,
            List<ServiceBusMessage> messages,
            int startIndex) {
        int index = startIndex;
        while (index < messages.size() && batch.tryAddMessage(messages.get(index))) {
            index++;
        }
        if (index == startIndex) {
            throw new IllegalArgumentException(
                    "Message is too large for an empty Service Bus batch: "
                            + messages.get(index).getMessageId());
        }
        return index;
    }

    @Override
    public void close() {
        sender.close();
    }
}

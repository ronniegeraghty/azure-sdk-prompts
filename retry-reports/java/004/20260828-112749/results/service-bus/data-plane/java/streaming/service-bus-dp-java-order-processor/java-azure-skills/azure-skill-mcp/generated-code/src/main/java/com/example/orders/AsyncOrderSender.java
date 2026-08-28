package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public final class AsyncOrderSender implements AutoCloseable {
    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;

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
        this.highPriorityThreshold = highPriorityThreshold;
    }

    public Mono<Void> send(Order order) {
        ServiceBusMessage message = messageFor(order);
        if (isHighPriority(order)) {
            return sender.scheduleMessage(
                            message,
                            OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS))
                    .then();
        }
        return sender.sendMessage(message);
    }

    public Mono<Void> sendBatch(List<Order> orders) {
        Map<String, List<Order>> immediateByCustomer = new LinkedHashMap<>();
        List<Order> scheduled = new ArrayList<>();
        for (Order order : orders) {
            if (isHighPriority(order)) {
                scheduled.add(order);
            } else {
                immediateByCustomer
                        .computeIfAbsent(order.customerName(), ignored -> new ArrayList<>())
                        .add(order);
            }
        }

        Mono<Void> immediate = Flux.fromIterable(immediateByCustomer.values())
                .concatMap(this::sendImmediateBatches)
                .then();
        Mono<Void> delayed = Flux.fromIterable(scheduled)
                .concatMap(this::send)
                .then();
        return immediate.then(delayed);
    }

    public Mono<Void> sendMalformedForDemo(String correlationId, String customerName) {
        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString("{not-valid-json"))
                .setContentType("application/json")
                .setMessageId(correlationId)
                .setCorrelationId(correlationId)
                .setSessionId(customerName);
        return sender.sendMessage(message);
    }

    private Mono<Void> sendImmediateBatches(List<Order> orders) {
        return sender.createMessageBatch().flatMap(batch -> addAndSend(orders, 0, batch));
    }

    private Mono<Void> addAndSend(List<Order> orders, int index, ServiceBusMessageBatch batch) {
        int next = index;
        while (next < orders.size()) {
            Order order = orders.get(next);
            if (!batch.tryAddMessage(messageFor(order))) {
                if (batch.getCount() == 0) {
                    return Mono.error(new IllegalArgumentException(
                            "Order is too large for an empty Service Bus batch: " + order.orderId()));
                }
                int resumeAt = next;
                return sender.sendMessages(batch)
                        .then(sender.createMessageBatch())
                        .flatMap(newBatch -> addAndSend(orders, resumeAt, newBatch));
            }
            next++;
        }
        return batch.getCount() == 0 ? Mono.empty() : sender.sendMessages(batch);
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

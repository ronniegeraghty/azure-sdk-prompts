package com.example.orders;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

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
        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, "highPriorityThreshold");
    }

    public Mono<Void> send(Order order) {
        ServiceBusMessage message = toMessage(order);
        if (isHighPriority(order)) {
            return sender.scheduleMessage(
                            message, OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS))
                    .then();
        }
        return sender.sendMessage(message);
    }

    public Mono<Void> sendBatch(List<Order> orders) {
        List<ServiceBusMessage> immediateMessages = new ArrayList<>();
        List<ServiceBusMessage> scheduledMessages = new ArrayList<>();
        for (Order order : orders) {
            (isHighPriority(order) ? scheduledMessages : immediateMessages).add(toMessage(order));
        }

        OffsetDateTime scheduledTime = OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS);
        Mono<Void> schedule = Flux.fromIterable(scheduledMessages)
                .concatMap(message -> sender.scheduleMessage(message, scheduledTime))
                .then();

        return schedule.then(sendSizeAwareBatches(immediateMessages, 0));
    }

    private Mono<Void> sendSizeAwareBatches(List<ServiceBusMessage> messages, int startIndex) {
        if (startIndex >= messages.size()) {
            return Mono.empty();
        }

        return sender.createMessageBatch().flatMap(batch -> {
            int nextIndex = fillBatch(batch, messages, startIndex);
            return sender.sendMessages(batch).then(sendSizeAwareBatches(messages, nextIndex));
        });
    }

    private int fillBatch(ServiceBusMessageBatch batch, List<ServiceBusMessage> messages, int startIndex) {
        int index = startIndex;
        while (index < messages.size() && batch.tryAddMessage(messages.get(index))) {
            index++;
        }
        if (index == startIndex) {
            throw new IllegalArgumentException("Order message exceeds the Service Bus maximum message size");
        }
        return index;
    }

    private ServiceBusMessage toMessage(Order order) {
        return OrderMessageFactory.create(order, highPriorityThreshold);
    }

    private boolean isHighPriority(Order order) {
        return order.getTotalPrice().compareTo(highPriorityThreshold) > 0;
    }

    @Override
    public void close() {
        sender.close();
    }
}

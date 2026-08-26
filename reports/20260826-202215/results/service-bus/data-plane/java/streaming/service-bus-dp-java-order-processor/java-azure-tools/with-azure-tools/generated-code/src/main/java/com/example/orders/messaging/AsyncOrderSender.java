package com.example.orders.messaging;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusMessageBatch;
import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.example.orders.codec.OrderJsonCodec;
import com.example.orders.model.Order;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;
import java.util.List;
import java.util.Objects;

public final class AsyncOrderSender implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderSender.class);

    private final ServiceBusSenderAsyncClient sender;
    private final OrderMessageFactory messageFactory;

    public AsyncOrderSender(
        String fullyQualifiedNamespace,
        String queueName,
        TokenCredential credential,
        BigDecimal highPriorityThreshold
    ) {
        this.sender = new ServiceBusClientBuilder()
            .credential(fullyQualifiedNamespace, credential)
            .sender()
            .queueName(queueName)
            .buildAsyncClient();
        this.messageFactory = new OrderMessageFactory(new OrderJsonCodec(), highPriorityThreshold);
    }

    public Mono<Void> sendOrder(Order order) {
        ServiceBusMessage message = messageFactory.createMessage(order);
        return sender.sendMessage(message)
            .doOnSuccess(ignored -> LOGGER.info("Sent order {} with priority {}", order.getOrderId(),
                message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY)));
    }

    public Mono<Void> sendOrders(List<Order> orders) {
        Objects.requireNonNull(orders, "orders");
        if (orders.isEmpty()) {
            return Mono.empty();
        }

        List<ServiceBusMessage> messages = orders.stream()
            .map(messageFactory::createMessage)
            .toList();
        return sender.createMessageBatch()
            .flatMap(batch -> fillAndSend(messages, 0, batch));
    }

    private Mono<Void> fillAndSend(
        List<ServiceBusMessage> messages,
        int startIndex,
        ServiceBusMessageBatch batch
    ) {
        int index = startIndex;
        while (index < messages.size() && batch.tryAddMessage(messages.get(index))) {
            index++;
        }

        if (index == startIndex) {
            return Mono.error(new IllegalArgumentException(
                "Order message exceeds the maximum Service Bus message size"
            ));
        }

        int nextIndex = index;
        Mono<Void> sendCurrentBatch = sender.sendMessages(batch)
            .doOnSuccess(ignored -> LOGGER.info("Sent batch containing {} order(s) and {} byte(s)",
                batch.getCount(), batch.getSizeInBytes()));

        if (nextIndex == messages.size()) {
            return sendCurrentBatch;
        }

        return sendCurrentBatch
            .then(sender.createMessageBatch())
            .flatMap(nextBatch -> fillAndSend(messages, nextIndex, nextBatch));
    }

    @Override
    public void close() {
        sender.close();
    }
}

package com.example.orders.processing;

import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.example.orders.codec.OrderJsonCodec;
import com.example.orders.model.Order;
import com.example.orders.model.OrderStatus;
import org.slf4j.Logger;

final class OrderProcessingSupport {
    private static final int MAX_DEAD_LETTER_DESCRIPTION_LENGTH = 4_096;

    private OrderProcessingSupport() {
    }

    static Order process(ServiceBusReceivedMessage message, OrderJsonCodec codec, Logger logger) {
        Order order = codec.deserialize(message.getBody().toString());
        if (order.getStatus() == OrderStatus.FAILED) {
            throw new IllegalStateException("Order is marked failed");
        }

        order.setStatus(OrderStatus.PROCESSING);
        logger.info("Processing order {} for customer {}", order.getOrderId(), order.getCustomerName());
        order.setStatus(OrderStatus.COMPLETED);
        logger.info("Completed order {}", order.getOrderId());
        return order;
    }

    static DeadLetterOptions deadLetterOptions(RuntimeException exception) {
        String description = exception.getMessage() == null
            ? exception.getClass().getName()
            : exception.getMessage();
        if (description.length() > MAX_DEAD_LETTER_DESCRIPTION_LENGTH) {
            description = description.substring(0, MAX_DEAD_LETTER_DESCRIPTION_LENGTH);
        }
        return new DeadLetterOptions()
            .setDeadLetterReason(exception.getClass().getSimpleName())
            .setDeadLetterErrorDescription(description);
    }
}

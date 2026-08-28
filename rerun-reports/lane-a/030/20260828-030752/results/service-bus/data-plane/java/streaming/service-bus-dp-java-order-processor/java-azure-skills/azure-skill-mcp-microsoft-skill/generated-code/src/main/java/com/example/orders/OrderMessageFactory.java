package com.example.orders;

import com.azure.core.util.BinaryData;
import com.azure.messaging.servicebus.ServiceBusMessage;

import java.math.BigDecimal;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

final class OrderMessageFactory {
    static final String PRIORITY_PROPERTY = "priority";
    static final String HIGH_PRIORITY = "high";
    static final String NORMAL_PRIORITY = "normal";
    private static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);

    private final BigDecimal highPriorityThreshold;

    OrderMessageFactory(BigDecimal highPriorityThreshold) {
        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, "highPriorityThreshold");
    }

    ServiceBusMessage create(Order order) {
        Objects.requireNonNull(order, "order");
        boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) > 0;

        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromObject(order))
            .setContentType("application/json")
            .setMessageId(order.getOrderId())
            .setCorrelationId(order.getOrderId())
            .setSessionId(order.getCustomerName());
        message.getApplicationProperties().put(
            PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY);

        if (highPriority) {
            message.setScheduledEnqueueTime(OffsetDateTime.now().plus(FRAUD_REVIEW_DELAY));
        }
        return message;
    }

    Order deserialize(BinaryData body) {
        return body.toObject(Order.class);
    }
}

package com.example.orders.messaging;

import com.azure.core.util.BinaryData;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.example.orders.codec.OrderJsonCodec;
import com.example.orders.model.Order;

import java.math.BigDecimal;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.Objects;

public final class OrderMessageFactory {
    public static final String PRIORITY_PROPERTY = "priority";
    public static final String HIGH_PRIORITY = "high";
    public static final String NORMAL_PRIORITY = "normal";
    public static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);

    private final OrderJsonCodec codec;
    private final BigDecimal highPriorityThreshold;

    public OrderMessageFactory(OrderJsonCodec codec, BigDecimal highPriorityThreshold) {
        this.codec = Objects.requireNonNull(codec, "codec");
        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, "highPriorityThreshold");
    }

    public ServiceBusMessage createMessage(Order order) {
        Objects.requireNonNull(order, "order");
        boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) > 0;

        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(codec.serialize(order)))
            .setMessageId(order.getOrderId())
            .setCorrelationId(order.getOrderId())
            .setSessionId(order.getCustomerName())
            .setContentType("application/json");
        message.getApplicationProperties().put(
            PRIORITY_PROPERTY,
            highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY
        );

        if (highPriority) {
            message.setScheduledEnqueueTime(OffsetDateTime.now(ZoneOffset.UTC).plus(FRAUD_REVIEW_DELAY));
        }
        return message;
    }
}

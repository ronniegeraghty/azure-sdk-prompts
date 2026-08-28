package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigDecimal;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

final class OrderMessageFactory {
    static final String CONTENT_TYPE = "application/json";
    static final String PRIORITY_PROPERTY = "priority";
    static final String HIGH_PRIORITY = "high";
    static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);

    private final ObjectMapper objectMapper;
    private final BigDecimal highPriorityThreshold;

    OrderMessageFactory(ObjectMapper objectMapper, BigDecimal highPriorityThreshold) {
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, "highPriorityThreshold");
    }

    ServiceBusMessage create(Order order, boolean setScheduledTime) {
        Objects.requireNonNull(order, "order");
        try {
            ServiceBusMessage message = new ServiceBusMessage(objectMapper.writeValueAsBytes(order))
                    .setContentType(CONTENT_TYPE)
                    .setMessageId(order.getOrderId())
                    .setCorrelationId(order.getOrderId())
                    .setSessionId(order.getCustomerName());

            if (isHighPriority(order)) {
                message.getApplicationProperties().put(PRIORITY_PROPERTY, HIGH_PRIORITY);
                if (setScheduledTime) {
                    message.setScheduledEnqueueTime(OffsetDateTime.now().plus(FRAUD_REVIEW_DELAY));
                }
            }
            return message;
        } catch (JsonProcessingException e) {
            throw new IllegalArgumentException("Unable to serialize order " + order.getOrderId(), e);
        }
    }

    boolean isHighPriority(Order order) {
        return order.getTotalPrice().compareTo(highPriorityThreshold) > 0;
    }
}

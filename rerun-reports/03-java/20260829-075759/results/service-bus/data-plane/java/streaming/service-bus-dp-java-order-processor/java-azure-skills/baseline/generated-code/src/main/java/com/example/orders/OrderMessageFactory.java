package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.Objects;

final class OrderMessageFactory {
    static final String PRIORITY_PROPERTY = "priority";
    static final String HIGH_PRIORITY = "high";
    private static final int FRAUD_REVIEW_DELAY_SECONDS = 30;

    private final ObjectMapper objectMapper;
    private final BigDecimal highPriorityThreshold;

    OrderMessageFactory(ObjectMapper objectMapper, BigDecimal highPriorityThreshold) {
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
        this.highPriorityThreshold = Objects.requireNonNull(highPriorityThreshold, "highPriorityThreshold");
    }

    ServiceBusMessage create(Order order) {
        Objects.requireNonNull(order, "order");
        try {
            ServiceBusMessage message = new ServiceBusMessage(objectMapper.writeValueAsBytes(order))
                    .setContentType("application/json")
                    .setMessageId(order.getOrderId())
                    .setCorrelationId(order.getOrderId())
                    .setSessionId(order.getCustomerName());

            if (order.getTotalPrice().compareTo(highPriorityThreshold) > 0) {
                message.getApplicationProperties().put(PRIORITY_PROPERTY, HIGH_PRIORITY);
                message.setScheduledEnqueueTime(OffsetDateTime.now().plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));
            }
            return message;
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Could not serialize order " + order.getOrderId(), exception);
        }
    }

    static ServiceBusMessage malformedDemoMessage(String messageId, String customerName) {
        return new ServiceBusMessage("{not-valid-json")
                .setContentType("application/json")
                .setMessageId(messageId)
                .setCorrelationId(messageId)
                .setSessionId(customerName);
    }
}

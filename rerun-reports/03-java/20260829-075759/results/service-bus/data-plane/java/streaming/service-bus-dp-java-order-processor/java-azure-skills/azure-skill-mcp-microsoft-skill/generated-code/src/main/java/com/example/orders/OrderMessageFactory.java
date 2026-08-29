package com.example.orders;

import com.azure.core.util.BinaryData;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigDecimal;
import java.time.Duration;
import java.time.OffsetDateTime;

final class OrderMessageFactory {
    static final String CONTENT_TYPE = "application/json";
    static final String PRIORITY_PROPERTY = "priority";
    static final String HIGH_PRIORITY = "high";
    static final Duration FRAUD_REVIEW_DELAY = Duration.ofSeconds(30);

    private final ObjectMapper objectMapper;
    private final BigDecimal highPriorityThreshold;

    OrderMessageFactory(ObjectMapper objectMapper, BigDecimal highPriorityThreshold) {
        this.objectMapper = objectMapper;
        this.highPriorityThreshold = highPriorityThreshold;
    }

    ServiceBusMessage create(Order order) {
        try {
            ServiceBusMessage message = new ServiceBusMessage(
                    BinaryData.fromString(objectMapper.writeValueAsString(order)))
                    .setContentType(CONTENT_TYPE)
                    .setMessageId(order.getOrderId())
                    .setCorrelationId(order.getOrderId())
                    .setSessionId(order.getCustomerName());

            if (order.getTotalPrice().compareTo(highPriorityThreshold) > 0) {
                message.getApplicationProperties().put(PRIORITY_PROPERTY, HIGH_PRIORITY);
                message.setScheduledEnqueueTime(OffsetDateTime.now().plus(FRAUD_REVIEW_DELAY));
            }
            return message;
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Could not serialize order " + order.getOrderId(), exception);
        }
    }
}

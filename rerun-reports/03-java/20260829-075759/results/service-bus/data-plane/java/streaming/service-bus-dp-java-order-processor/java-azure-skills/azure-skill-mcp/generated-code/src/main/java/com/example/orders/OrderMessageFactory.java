package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;

final class OrderMessageFactory {
    static final String PRIORITY_PROPERTY = "priority";
    static final String HIGH_PRIORITY = "high";
    static final String NORMAL_PRIORITY = "normal";
    static final int FRAUD_REVIEW_DELAY_SECONDS = 30;

    private OrderMessageFactory() {
    }

    static ServiceBusMessage create(Order order, BigDecimal highPriorityThreshold) {
        boolean highPriority = order.totalPrice().compareTo(highPriorityThreshold) > 0;
        ServiceBusMessage message = new ServiceBusMessage(OrderJson.serialize(order))
                .setContentType("application/json")
                .setMessageId(order.orderId())
                .setCorrelationId(order.orderId())
                .setSessionId(order.customerName());
        message.getApplicationProperties().put(
                PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY);

        if (highPriority) {
            message.setScheduledEnqueueTime(
                    OffsetDateTime.now(ZoneOffset.UTC).plusSeconds(FRAUD_REVIEW_DELAY_SECONDS));
        }
        return message;
    }
}

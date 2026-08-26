package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;

import java.math.BigDecimal;

final class OrderMessageFactory {
    static final String PRIORITY_PROPERTY = "priority";
    static final String HIGH_PRIORITY = "high";
    static final String NORMAL_PRIORITY = "normal";

    private OrderMessageFactory() {
    }

    static ServiceBusMessage create(Order order, BigDecimal highPriorityThreshold) {
        boolean highPriority = order.getTotalPrice().compareTo(highPriorityThreshold) > 0;
        ServiceBusMessage message = new ServiceBusMessage(OrderJson.serialize(order))
                .setContentType("application/json")
                .setMessageId(order.getOrderId())
                .setCorrelationId(order.getOrderId())
                .setSessionId(order.getCustomerName())
                .setSubject("order");
        message.getApplicationProperties().put(
                PRIORITY_PROPERTY, highPriority ? HIGH_PRIORITY : NORMAL_PRIORITY);
        return message;
    }
}

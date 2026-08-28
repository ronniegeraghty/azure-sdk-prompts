package com.example.orders;

import com.azure.core.util.BinaryData;
import com.azure.messaging.servicebus.ServiceBusMessage;
import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

final class OrderMessageMapper {
    private static final ObjectMapper JSON = new ObjectMapper();

    private OrderMessageMapper() {
    }

    static ServiceBusMessage toMessage(Order order, boolean highPriority) {
        ServiceBusMessage message = new ServiceBusMessage(BinaryData.fromString(toJson(order)))
                .setContentType("application/json")
                .setMessageId(order.orderId())
                .setCorrelationId(order.orderId())
                .setSessionId(order.customerName());
        message.getApplicationProperties().put("priority", highPriority ? "high" : "normal");
        return message;
    }

    static Order fromMessage(ServiceBusReceivedMessage message) throws JsonProcessingException {
        return JSON.readValue(message.getBody().toString(), Order.class);
    }

    static String toJson(Order order) {
        try {
            return JSON.writeValueAsString(order);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Could not serialize order " + order.orderId(), exception);
        }
    }
}

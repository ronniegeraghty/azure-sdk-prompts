package com.example.orders;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

final class OrderJson {
    private static final ObjectMapper MAPPER = new ObjectMapper();

    private OrderJson() {
    }

    static String serialize(Order order) {
        try {
            return MAPPER.writeValueAsString(order);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Could not serialize order " + order.getOrderId(), exception);
        }
    }

    static Order deserialize(String json) throws JsonProcessingException {
        return MAPPER.readValue(json, Order.class);
    }
}

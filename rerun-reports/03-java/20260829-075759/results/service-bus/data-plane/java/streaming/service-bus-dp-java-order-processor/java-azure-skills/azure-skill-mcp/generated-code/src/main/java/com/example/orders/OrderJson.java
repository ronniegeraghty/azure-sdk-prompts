package com.example.orders;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;

public final class OrderJson {
    private static final ObjectMapper MAPPER = new ObjectMapper()
            .enable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);

    private OrderJson() {
    }

    public static String serialize(Order order) {
        try {
            return MAPPER.writeValueAsString(order);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Could not serialize order " + order.orderId(), exception);
        }
    }

    public static Order deserialize(String json) {
        try {
            return MAPPER.readValue(json, Order.class);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Could not deserialize order JSON", exception);
        }
    }
}

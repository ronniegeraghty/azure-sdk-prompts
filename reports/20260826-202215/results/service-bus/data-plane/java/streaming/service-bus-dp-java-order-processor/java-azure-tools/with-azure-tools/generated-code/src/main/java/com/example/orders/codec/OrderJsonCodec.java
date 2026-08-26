package com.example.orders.codec;

import com.example.orders.model.Order;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;

public final class OrderJsonCodec {
    private final ObjectMapper objectMapper;

    public OrderJsonCodec() {
        this.objectMapper = new ObjectMapper()
            .enable(DeserializationFeature.FAIL_ON_NULL_CREATOR_PROPERTIES);
    }

    public String serialize(Order order) {
        try {
            return objectMapper.writeValueAsString(order);
        } catch (JsonProcessingException exception) {
            throw new OrderSerializationException("Could not serialize order", exception);
        }
    }

    public Order deserialize(String json) {
        try {
            return objectMapper.readValue(json, Order.class);
        } catch (JsonProcessingException exception) {
            throw new OrderSerializationException("Could not deserialize order", exception);
        }
    }
}

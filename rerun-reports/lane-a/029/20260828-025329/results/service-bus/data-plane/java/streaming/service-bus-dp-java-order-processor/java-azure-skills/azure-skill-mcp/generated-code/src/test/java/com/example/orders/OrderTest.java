package com.example.orders;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;

import static org.junit.jupiter.api.Assertions.assertEquals;

class OrderTest {
    private final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    void roundTripsAsJson() throws Exception {
        Order original = new Order(
                "order-42",
                "Ada",
                "Keyboard",
                2,
                new BigDecimal("199.90"),
                Order.Status.PROCESSING);

        Order decoded = objectMapper.readValue(objectMapper.writeValueAsBytes(original), Order.class);

        assertEquals(original.getOrderId(), decoded.getOrderId());
        assertEquals(original.getCustomerName(), decoded.getCustomerName());
        assertEquals(original.getProduct(), decoded.getProduct());
        assertEquals(original.getQuantity(), decoded.getQuantity());
        assertEquals(original.getTotalPrice(), decoded.getTotalPrice());
        assertEquals(original.getStatus(), decoded.getStatus());
    }
}

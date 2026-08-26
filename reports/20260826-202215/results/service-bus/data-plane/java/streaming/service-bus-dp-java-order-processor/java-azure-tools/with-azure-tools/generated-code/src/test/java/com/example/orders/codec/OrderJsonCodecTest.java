package com.example.orders.codec;

import com.example.orders.model.Order;
import com.example.orders.model.OrderStatus;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;

import static org.junit.jupiter.api.Assertions.assertEquals;

class OrderJsonCodecTest {
    private final OrderJsonCodec codec = new OrderJsonCodec();

    @Test
    void roundTripsOrderAsJson() {
        Order original = new Order(
            "order-1",
            "Ada",
            "Laptop",
            2,
            new BigDecimal("2400.50"),
            OrderStatus.PENDING
        );

        Order restored = codec.deserialize(codec.serialize(original));

        assertEquals(original.getOrderId(), restored.getOrderId());
        assertEquals(original.getCustomerName(), restored.getCustomerName());
        assertEquals(original.getProduct(), restored.getProduct());
        assertEquals(original.getQuantity(), restored.getQuantity());
        assertEquals(original.getTotalPrice(), restored.getTotalPrice());
        assertEquals(original.getStatus(), restored.getStatus());
    }
}

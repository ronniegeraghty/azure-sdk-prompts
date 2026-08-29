package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;
import java.time.Duration;
import java.time.OffsetDateTime;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class OrderJsonTest {
    @Test
    void roundTripsOrderWithLowerCaseStatus() {
        Order order = new Order(
                "order-1",
                "Ada",
                "Keyboard",
                2,
                new BigDecimal("199.98"),
                OrderStatus.PENDING);

        String json = OrderJson.serialize(order);

        assertTrue(json.contains("\"status\":\"pending\""));
        assertEquals(order, OrderJson.deserialize(json));
    }

    @Test
    void rejectsMalformedJson() {
        assertThrows(IllegalArgumentException.class, () -> OrderJson.deserialize("{bad-json"));
    }

    @Test
    void createsSessionAwareHighPriorityScheduledMessage() {
        Order order = new Order(
                "order-2",
                "Grace",
                "Workstation",
                1,
                new BigDecimal("2500.00"),
                OrderStatus.PENDING);
        OffsetDateTime before = OffsetDateTime.now().plusSeconds(29);

        ServiceBusMessage message = OrderMessageFactory.create(order, new BigDecimal("1000.00"));

        assertEquals(order.orderId(), message.getCorrelationId());
        assertEquals(order.customerName(), message.getSessionId());
        assertEquals(OrderMessageFactory.HIGH_PRIORITY,
                message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));
        assertTrue(message.getScheduledEnqueueTime().isAfter(before));
        assertTrue(Duration.between(before, message.getScheduledEnqueueTime()).toSeconds() <= 2);
    }
}

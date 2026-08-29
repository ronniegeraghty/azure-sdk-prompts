package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;
import java.time.OffsetDateTime;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class OrderTest {
    private final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    void serializesAndDeserializesOrder() throws Exception {
        Order original = order("order-1", new BigDecimal("42.50"));

        String json = objectMapper.writeValueAsString(original);
        Order restored = objectMapper.readValue(json, Order.class);

        assertTrue(json.contains("\"status\":\"pending\""));
        assertEquals(original.getOrderId(), restored.getOrderId());
        assertEquals(original.getCustomerName(), restored.getCustomerName());
        assertEquals(original.getProduct(), restored.getProduct());
        assertEquals(original.getQuantity(), restored.getQuantity());
        assertEquals(original.getTotalPrice(), restored.getTotalPrice());
        assertEquals(original.getStatus(), restored.getStatus());
    }

    @Test
    void highValueOrderHasSessionCorrelationPriorityAndDelay() {
        OffsetDateTime earliestExpected = OffsetDateTime.now().plusSeconds(29);
        ServiceBusMessage message = new OrderMessageFactory(objectMapper, new BigDecimal("1000"))
                .create(order("order-2", new BigDecimal("1000.01")));

        assertEquals("order-2", message.getCorrelationId());
        assertEquals("order-2", message.getMessageId());
        assertEquals("Ada", message.getSessionId());
        assertEquals("high", message.getApplicationProperties().get("priority"));
        assertNotNull(message.getScheduledEnqueueTime());
        assertTrue(message.getScheduledEnqueueTime().isAfter(earliestExpected));
    }

    private static Order order(String id, BigDecimal total) {
        return new Order(id, "Ada", "Keyboard", 1, total, Order.Status.PENDING);
    }
}

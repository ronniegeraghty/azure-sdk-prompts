package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;

import static org.junit.jupiter.api.Assertions.assertEquals;

class OrderJsonTest {
    @Test
    void serializesStatusAsLowercaseJson() {
        Order order = new Order(
                "order-1",
                "Ada",
                "Keyboard",
                2,
                new BigDecimal("199.98"),
                OrderStatus.PROCESSING);

        ServiceBusMessage message = OrderMessageMapper.toMessage(order, false);

        assertEquals("order-1", message.getCorrelationId());
        assertEquals("Ada", message.getSessionId());
        assertEquals("normal", message.getApplicationProperties().get("priority"));
        assertEquals(
                "{\"orderId\":\"order-1\",\"customerName\":\"Ada\",\"product\":\"Keyboard\","
                        + "\"quantity\":2,\"totalPrice\":199.98,\"status\":\"processing\"}",
                message.getBody().toString());
    }

    @Test
    void roundTripsThroughJson() throws Exception {
        Order expected = new Order(
                "order-2",
                "Grace",
                "Monitor",
                1,
                new BigDecimal("749.99"),
                OrderStatus.COMPLETED);
        ObjectMapper mapper = new ObjectMapper();

        Order actual = mapper.readValue(mapper.writeValueAsString(expected), Order.class);

        assertEquals(expected, actual);
    }
}

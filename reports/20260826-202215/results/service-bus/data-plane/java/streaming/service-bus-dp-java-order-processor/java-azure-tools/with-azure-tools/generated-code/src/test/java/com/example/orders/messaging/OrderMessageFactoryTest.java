package com.example.orders.messaging;

import com.azure.messaging.servicebus.ServiceBusMessage;
import com.example.orders.codec.OrderJsonCodec;
import com.example.orders.model.Order;
import com.example.orders.model.OrderStatus;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class OrderMessageFactoryTest {
    private final OrderMessageFactory factory =
        new OrderMessageFactory(new OrderJsonCodec(), new BigDecimal("1000.00"));

    @Test
    void mapsOrderMetadataToServiceBusMessage() {
        Order order = order("order-1", "Ada", "999.99");

        ServiceBusMessage message = factory.createMessage(order);

        assertEquals("order-1", message.getMessageId());
        assertEquals("order-1", message.getCorrelationId());
        assertEquals("Ada", message.getSessionId());
        assertEquals("application/json", message.getContentType());
        assertEquals(OrderMessageFactory.NORMAL_PRIORITY,
            message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));
        assertNull(message.getScheduledEnqueueTime());
    }

    @Test
    void schedulesHighPriorityOrderForFraudReview() {
        OffsetDateTime before = OffsetDateTime.now(ZoneOffset.UTC)
            .plus(OrderMessageFactory.FRAUD_REVIEW_DELAY);

        ServiceBusMessage message = factory.createMessage(order("order-2", "Grace", "1000.01"));

        OffsetDateTime after = OffsetDateTime.now(ZoneOffset.UTC)
            .plus(OrderMessageFactory.FRAUD_REVIEW_DELAY);
        assertEquals(OrderMessageFactory.HIGH_PRIORITY,
            message.getApplicationProperties().get(OrderMessageFactory.PRIORITY_PROPERTY));
        assertNotNull(message.getScheduledEnqueueTime());
        assertTrue(!message.getScheduledEnqueueTime().isBefore(before));
        assertTrue(!message.getScheduledEnqueueTime().isAfter(after));
    }

    private static Order order(String orderId, String customer, String price) {
        return new Order(
            orderId,
            customer,
            "Product",
            1,
            new BigDecimal(price),
            OrderStatus.PENDING
        );
    }
}

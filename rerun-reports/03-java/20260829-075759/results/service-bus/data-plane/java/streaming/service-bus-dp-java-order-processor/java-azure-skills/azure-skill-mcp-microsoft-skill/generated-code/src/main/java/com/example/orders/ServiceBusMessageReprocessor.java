package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusSenderAsyncClient;
import com.azure.messaging.servicebus.ServiceBusSenderClient;
import com.fasterxml.jackson.databind.ObjectMapper;
import reactor.core.publisher.Mono;

import java.math.BigDecimal;

final class ServiceBusMessageReprocessor {
    private static final BigDecimal NO_DELAY_THRESHOLD = new BigDecimal("999999999999");

    private ServiceBusMessageReprocessor() {
    }

    static void send(ServiceBusSenderClient sender, Order order, ObjectMapper objectMapper) {
        sender.sendMessage(new OrderMessageFactory(objectMapper, NO_DELAY_THRESHOLD).create(order));
    }

    static Mono<Void> send(ServiceBusSenderAsyncClient sender, Order order, ObjectMapper objectMapper) {
        return sender.sendMessage(new OrderMessageFactory(objectMapper, NO_DELAY_THRESHOLD).create(order));
    }
}

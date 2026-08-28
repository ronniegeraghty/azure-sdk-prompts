package com.example.orders;

public record DeadLetteredOrder(
        String correlationId,
        String sessionId,
        String body,
        String reason,
        String errorDescription) {
}

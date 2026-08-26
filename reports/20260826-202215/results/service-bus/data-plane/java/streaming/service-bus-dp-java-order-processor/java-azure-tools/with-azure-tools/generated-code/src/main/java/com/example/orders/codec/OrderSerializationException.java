package com.example.orders.codec;

public final class OrderSerializationException extends RuntimeException {
    public OrderSerializationException(String message, Throwable cause) {
        super(message, cause);
    }
}

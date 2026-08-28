package com.example.todo;

public class OptimisticConcurrencyException extends RuntimeException {
    public OptimisticConcurrencyException(String message, Throwable cause) {
        super(message, cause);
    }
}

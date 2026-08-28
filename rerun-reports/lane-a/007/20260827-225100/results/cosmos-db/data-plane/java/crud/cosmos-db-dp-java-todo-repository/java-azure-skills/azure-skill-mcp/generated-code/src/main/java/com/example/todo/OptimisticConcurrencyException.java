package com.example.todo;

public class OptimisticConcurrencyException extends RuntimeException {
    public OptimisticConcurrencyException(String id, Throwable cause) {
        super("ToDo item '" + id
                + "' was modified by another process. Read it again and retry with the new ETag.", cause);
    }
}

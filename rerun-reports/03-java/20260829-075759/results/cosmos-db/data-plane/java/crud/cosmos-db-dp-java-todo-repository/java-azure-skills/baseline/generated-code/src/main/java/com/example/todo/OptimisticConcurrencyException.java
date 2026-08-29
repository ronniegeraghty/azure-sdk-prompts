package com.example.todo;

public class OptimisticConcurrencyException extends RuntimeException {
    public OptimisticConcurrencyException(String itemId, Throwable cause) {
        super("ToDo item '" + itemId
                + "' was modified after it was read; reload it before retrying the update.", cause);
    }
}

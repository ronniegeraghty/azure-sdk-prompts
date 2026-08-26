package com.example.todo;

public class OptimisticConcurrencyException extends RuntimeException {
    public OptimisticConcurrencyException(String itemId, Throwable cause) {
        super("ToDo item '" + itemId
            + "' was changed by another process. Read the latest version before retrying the update.", cause);
    }
}

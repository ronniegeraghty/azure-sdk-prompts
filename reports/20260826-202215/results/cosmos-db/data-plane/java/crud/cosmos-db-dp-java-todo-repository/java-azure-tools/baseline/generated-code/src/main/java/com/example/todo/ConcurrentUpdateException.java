package com.example.todo;

public class ConcurrentUpdateException extends RuntimeException {
    public ConcurrentUpdateException(String itemId, Throwable cause) {
        super("ToDo item '" + itemId
                + "' was modified after it was read; reload it before updating.", cause);
    }
}

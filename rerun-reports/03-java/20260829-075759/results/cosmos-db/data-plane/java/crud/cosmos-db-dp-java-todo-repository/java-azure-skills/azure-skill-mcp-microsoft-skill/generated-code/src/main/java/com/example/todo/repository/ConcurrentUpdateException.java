package com.example.todo.repository;

public class ConcurrentUpdateException extends RuntimeException {
    public ConcurrentUpdateException(String itemId, Throwable cause) {
        super("ToDo item '" + itemId
                + "' was modified by another process; read it again before retrying the update.", cause);
    }
}

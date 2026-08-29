package com.example.todo.repository;

public record OperationResult<T>(T item, String etag, double requestCharge) {
}

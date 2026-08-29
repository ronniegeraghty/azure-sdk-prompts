package com.example.todo;

public record OperationResult<T>(T value, double requestCharge) {
}

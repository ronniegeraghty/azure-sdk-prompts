package com.example.todo;

public record CosmosOperationResult<T>(T value, double requestCharge) {
}

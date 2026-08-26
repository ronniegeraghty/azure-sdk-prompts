package com.example.todo;

public record RepositoryResponse<T>(T value, double requestCharge) {
}

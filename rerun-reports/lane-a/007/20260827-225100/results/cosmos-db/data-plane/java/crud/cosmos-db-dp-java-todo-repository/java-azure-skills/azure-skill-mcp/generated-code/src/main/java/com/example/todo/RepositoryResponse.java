package com.example.todo;

public record RepositoryResponse<T>(T value, String etag, double requestCharge) {
}

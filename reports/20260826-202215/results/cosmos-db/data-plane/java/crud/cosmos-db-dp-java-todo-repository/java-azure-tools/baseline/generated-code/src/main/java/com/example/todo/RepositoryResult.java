package com.example.todo;

public record RepositoryResult<T>(T value, double requestCharge) {
}

package com.example.todo;

import java.util.List;

public record RepositoryPage<T>(
        List<T> items,
        double requestCharge,
        String continuationToken) {

    public RepositoryPage {
        items = List.copyOf(items);
    }
}

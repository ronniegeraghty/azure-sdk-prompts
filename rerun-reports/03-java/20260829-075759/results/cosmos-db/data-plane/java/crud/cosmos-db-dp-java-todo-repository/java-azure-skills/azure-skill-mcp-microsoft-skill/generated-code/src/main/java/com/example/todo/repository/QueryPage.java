package com.example.todo.repository;

import java.util.List;

public record QueryPage<T>(
        List<T> items,
        String continuationToken,
        double requestCharge,
        int pageNumber) {

    public QueryPage {
        items = List.copyOf(items);
    }
}

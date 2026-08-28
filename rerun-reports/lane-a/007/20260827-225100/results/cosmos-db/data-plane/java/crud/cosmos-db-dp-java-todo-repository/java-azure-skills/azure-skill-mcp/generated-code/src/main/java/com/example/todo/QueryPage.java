package com.example.todo;

import java.util.List;

public record QueryPage<T>(
        List<T> results,
        String continuationToken,
        double requestCharge,
        int pageNumber) {

    public QueryPage {
        results = List.copyOf(results);
    }
}

package com.example.todo;

import java.util.List;

public record QueryPage<T>(List<T> items, double requestCharge, String continuationToken) {
    public QueryPage {
        items = List.copyOf(items);
    }
}

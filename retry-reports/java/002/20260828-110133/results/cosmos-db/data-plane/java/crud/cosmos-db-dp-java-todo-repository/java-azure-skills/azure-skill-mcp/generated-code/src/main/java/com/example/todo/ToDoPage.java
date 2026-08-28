package com.example.todo;

import java.util.List;

public record ToDoPage(
        List<ToDoItem> items,
        double requestCharge,
        String continuationToken) {

    public ToDoPage {
        items = List.copyOf(items);
    }
}

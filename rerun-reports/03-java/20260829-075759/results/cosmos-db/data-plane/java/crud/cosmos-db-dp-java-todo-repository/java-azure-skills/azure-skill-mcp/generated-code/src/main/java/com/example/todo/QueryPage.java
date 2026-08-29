package com.example.todo;

import java.util.List;

public record QueryPage(List<ToDoItem> items, double requestCharge) {
    public QueryPage {
        items = List.copyOf(items);
    }
}

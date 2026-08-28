package com.example.todo;

import java.time.Instant;
import java.util.Objects;

public class ToDoItem {
    private String id;
    private String title;
    private String description;
    private boolean completed;
    private Instant createdTimestamp;
    private String category;

    public ToDoItem() {
    }

    public ToDoItem(
            String id,
            String title,
            String description,
            boolean completed,
            Instant createdTimestamp,
            String category) {
        this.id = Objects.requireNonNull(id, "id");
        this.title = Objects.requireNonNull(title, "title");
        this.description = description;
        this.completed = completed;
        this.createdTimestamp = Objects.requireNonNull(createdTimestamp, "createdTimestamp");
        this.category = Objects.requireNonNull(category, "category");
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public boolean isCompleted() {
        return completed;
    }

    public void setCompleted(boolean completed) {
        this.completed = completed;
    }

    public Instant getCreatedTimestamp() {
        return createdTimestamp;
    }

    public void setCreatedTimestamp(Instant createdTimestamp) {
        this.createdTimestamp = createdTimestamp;
    }

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    @Override
    public String toString() {
        return "ToDoItem{" +
                "id='" + id + '\'' +
                ", title='" + title + '\'' +
                ", completed=" + completed +
                ", createdTimestamp=" + createdTimestamp +
                ", category='" + category + '\'' +
                '}';
    }
}

package com.example.todo;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.Instant;
import java.util.Objects;

public class ToDoItem {
    private String id;
    private String title;
    private String description;
    private boolean completed;
    private Instant createdAt;
    private String category;

    @JsonProperty("_etag")
    private String etag;

    public ToDoItem() {
    }

    public ToDoItem(
        String id,
        String title,
        String description,
        boolean completed,
        Instant createdAt,
        String category
    ) {
        this.id = Objects.requireNonNull(id, "id");
        this.title = Objects.requireNonNull(title, "title");
        this.description = description;
        this.completed = completed;
        this.createdAt = Objects.requireNonNull(createdAt, "createdAt");
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

    public Instant getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Instant createdAt) {
        this.createdAt = createdAt;
    }

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public String getEtag() {
        return etag;
    }

    public void setEtag(String etag) {
        this.etag = etag;
    }

    @Override
    public String toString() {
        return "ToDoItem{"
            + "id='" + id + '\''
            + ", title='" + title + '\''
            + ", completed=" + completed
            + ", createdAt=" + createdAt
            + ", category='" + category + '\''
            + '}';
    }
}

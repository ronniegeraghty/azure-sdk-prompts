package com.example.blobevents.model;

import java.util.Map;
import java.util.Objects;

public record CustomEvent(String subject, String type, Map<String, Object> data, String dataVersion) {
    public CustomEvent {
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(type, "type");
        data = Map.copyOf(Objects.requireNonNull(data, "data"));
        Objects.requireNonNull(dataVersion, "dataVersion");
        if (!subject.startsWith("/")) {
            throw new IllegalArgumentException("Subject must be an absolute hierarchy beginning with '/'");
        }
    }
}

package com.example.blobevents;

import java.util.Map;
import java.util.Objects;

public record CustomEvent(String eventType, String subject, Map<String, Object> data, String dataVersion) {
    public CustomEvent {
        Objects.requireNonNull(eventType, "eventType");
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(data, "data");
        Objects.requireNonNull(dataVersion, "dataVersion");
        if (!subject.startsWith("/")) {
            throw new IllegalArgumentException("subject must be an absolute hierarchy starting with '/'");
        }
        data = Map.copyOf(data);
    }
}

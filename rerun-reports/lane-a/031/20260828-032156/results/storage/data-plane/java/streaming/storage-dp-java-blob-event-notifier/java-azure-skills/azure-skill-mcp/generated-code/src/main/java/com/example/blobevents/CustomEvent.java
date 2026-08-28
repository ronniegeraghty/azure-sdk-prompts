package com.example.blobevents;

import java.util.Objects;

public record CustomEvent(String subject, String eventType, String dataVersion, Object data) {
    public CustomEvent {
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(eventType, "eventType");
        Objects.requireNonNull(dataVersion, "dataVersion");
        Objects.requireNonNull(data, "data");
        if (!subject.startsWith("/")) {
            throw new IllegalArgumentException("subject must be an absolute hierarchy beginning with '/'");
        }
    }
}

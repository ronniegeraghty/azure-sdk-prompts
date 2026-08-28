package com.example.blobevents;

import java.util.Objects;

public record CustomEvent(String eventType, String subject, Object data) {
    public CustomEvent {
        Objects.requireNonNull(eventType, "eventType");
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(data, "data");
        if (!subject.startsWith("/")) {
            throw new IllegalArgumentException("subject must be an absolute hierarchy beginning with '/'");
        }
    }
}

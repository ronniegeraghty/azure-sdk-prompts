package com.example.blobevents.model;

import java.util.Objects;

public record CustomEvent(String type, String subject, String dataVersion, Object data) {
    public CustomEvent {
        Objects.requireNonNull(type, "type");
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(dataVersion, "dataVersion");
        Objects.requireNonNull(data, "data");
        if (!subject.startsWith("/")) {
            throw new IllegalArgumentException("subject must be a hierarchy beginning with '/'");
        }
    }
}

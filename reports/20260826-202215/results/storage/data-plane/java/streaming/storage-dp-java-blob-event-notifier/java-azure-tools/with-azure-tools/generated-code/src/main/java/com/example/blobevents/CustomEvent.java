package com.example.blobevents;

import java.util.Objects;

public record CustomEvent(
    String type,
    String subject,
    Object data,
    String dataVersion
) {
    public CustomEvent {
        Objects.requireNonNull(type, "type");
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(data, "data");
        Objects.requireNonNull(dataVersion, "dataVersion");
        if (!subject.startsWith("/")) {
            throw new IllegalArgumentException("subject must start with '/'");
        }
    }
}

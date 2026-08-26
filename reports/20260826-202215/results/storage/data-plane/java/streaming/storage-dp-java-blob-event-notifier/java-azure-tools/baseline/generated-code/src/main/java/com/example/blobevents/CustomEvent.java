package com.example.blobevents;

import java.util.Objects;

public record CustomEvent(String subject, String eventType, String dataVersion, Object data) {
    public CustomEvent {
        requireHierarchy(subject);
        Objects.requireNonNull(eventType, "eventType");
        Objects.requireNonNull(dataVersion, "dataVersion");
        Objects.requireNonNull(data, "data");
    }

    private static void requireHierarchy(String subject) {
        if (subject == null || !subject.startsWith("/") || subject.length() == 1) {
            throw new IllegalArgumentException("subject must be a hierarchy beginning with '/', for example "
                    + "'/documents/invoices/processed'");
        }
    }
}

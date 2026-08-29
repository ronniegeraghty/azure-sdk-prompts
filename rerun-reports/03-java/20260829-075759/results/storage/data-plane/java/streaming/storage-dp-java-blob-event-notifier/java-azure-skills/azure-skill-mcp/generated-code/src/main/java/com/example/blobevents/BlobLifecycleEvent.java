package com.example.blobevents;

import com.azure.core.util.BinaryData;

import java.time.OffsetDateTime;
import java.util.Objects;

public record BlobLifecycleEvent(
    String id,
    String type,
    String subject,
    OffsetDateTime time,
    BinaryData data,
    EventSchema schema
) {
    public BlobLifecycleEvent {
        Objects.requireNonNull(id, "id");
        Objects.requireNonNull(type, "type");
        Objects.requireNonNull(subject, "subject");
        Objects.requireNonNull(data, "data");
        Objects.requireNonNull(schema, "schema");
    }
}

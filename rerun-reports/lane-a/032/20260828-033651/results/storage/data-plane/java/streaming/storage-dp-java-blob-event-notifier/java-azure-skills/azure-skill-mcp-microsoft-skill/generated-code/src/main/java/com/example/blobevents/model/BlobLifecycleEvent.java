package com.example.blobevents.model;

import com.azure.core.util.BinaryData;

import java.time.OffsetDateTime;

public record BlobLifecycleEvent(
    String id,
    String type,
    String subject,
    OffsetDateTime time,
    BinaryData data,
    EventSchema schema
) {
    public enum EventSchema {
        EVENT_GRID,
        CLOUD_EVENTS
    }
}

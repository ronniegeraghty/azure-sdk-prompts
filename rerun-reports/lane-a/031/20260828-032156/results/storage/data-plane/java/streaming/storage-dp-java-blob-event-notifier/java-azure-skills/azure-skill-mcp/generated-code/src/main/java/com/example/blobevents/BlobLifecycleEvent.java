package com.example.blobevents;

import com.fasterxml.jackson.databind.JsonNode;
import java.time.OffsetDateTime;

public record BlobLifecycleEvent(
    String id,
    String eventType,
    String subject,
    OffsetDateTime eventTime,
    JsonNode data,
    EventSchema schema
) {
    public enum EventSchema {
        EVENT_GRID,
        CLOUD_EVENTS_1_0
    }
}

package com.example.blobevents;

import com.fasterxml.jackson.databind.JsonNode;

import java.time.OffsetDateTime;

public record BlobLifecycleEvent(
        EventSchema schema,
        String id,
        String type,
        String subject,
        OffsetDateTime time,
        JsonNode data) {
}

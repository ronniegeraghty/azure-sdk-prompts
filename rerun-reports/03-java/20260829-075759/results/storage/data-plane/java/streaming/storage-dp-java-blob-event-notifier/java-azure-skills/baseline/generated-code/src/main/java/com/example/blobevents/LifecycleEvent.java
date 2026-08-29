package com.example.blobevents;

import com.fasterxml.jackson.databind.JsonNode;

import java.time.OffsetDateTime;

public record LifecycleEvent(
        String id,
        String type,
        String subject,
        OffsetDateTime time,
        String dataVersion,
        JsonNode data,
        Schema schema) {

    public enum Schema {
        EVENT_GRID,
        CLOUD_EVENTS
    }
}

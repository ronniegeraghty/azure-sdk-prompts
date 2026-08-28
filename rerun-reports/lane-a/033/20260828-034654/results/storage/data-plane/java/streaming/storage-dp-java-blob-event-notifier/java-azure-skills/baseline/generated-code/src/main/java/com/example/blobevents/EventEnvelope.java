package com.example.blobevents;

import com.azure.core.util.BinaryData;

import java.time.OffsetDateTime;

public record EventEnvelope(
        String id,
        String type,
        String subject,
        OffsetDateTime time,
        BinaryData data,
        Schema schema) {

    public enum Schema {
        EVENT_GRID,
        CLOUD_EVENTS
    }
}

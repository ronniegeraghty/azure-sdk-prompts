package com.example.blobevents.model;

import com.azure.core.util.BinaryData;
import java.time.OffsetDateTime;

public record IncomingEvent(
    Schema schema,
    String id,
    String type,
    String subject,
    OffsetDateTime time,
    BinaryData data
) {
    public enum Schema {
        EVENT_GRID,
        CLOUD_EVENT
    }
}

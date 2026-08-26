package com.example.blobevents;

import com.azure.core.util.BinaryData;

import java.time.OffsetDateTime;

public record BlobLifecycleEvent(
    String id,
    String type,
    String subject,
    OffsetDateTime time,
    BinaryData data
) {
}

package com.example.blobevents;

public record BlobSummary(
    String name,
    long size,
    String contentType,
    String accessTier
) {
}

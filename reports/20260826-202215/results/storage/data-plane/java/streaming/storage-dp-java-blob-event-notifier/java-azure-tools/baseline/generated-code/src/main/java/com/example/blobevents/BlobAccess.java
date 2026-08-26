package com.example.blobevents;

@FunctionalInterface
public interface BlobAccess {
    BlobSummary download(String container, String blobName);
}

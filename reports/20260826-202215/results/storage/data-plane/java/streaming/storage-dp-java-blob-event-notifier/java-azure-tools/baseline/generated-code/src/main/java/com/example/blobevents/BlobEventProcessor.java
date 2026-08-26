package com.example.blobevents;

public interface BlobEventProcessor {
    void onBlobCreated(BlobLifecycleEvent event);

    void onBlobDeleted(BlobLifecycleEvent event);
}

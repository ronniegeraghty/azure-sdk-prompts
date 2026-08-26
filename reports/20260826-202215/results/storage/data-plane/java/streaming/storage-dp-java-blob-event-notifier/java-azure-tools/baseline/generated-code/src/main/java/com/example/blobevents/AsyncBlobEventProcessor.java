package com.example.blobevents;

import reactor.core.publisher.Mono;

public interface AsyncBlobEventProcessor {
    Mono<Void> onBlobCreated(BlobLifecycleEvent event);

    Mono<Void> onBlobDeleted(BlobLifecycleEvent event);
}

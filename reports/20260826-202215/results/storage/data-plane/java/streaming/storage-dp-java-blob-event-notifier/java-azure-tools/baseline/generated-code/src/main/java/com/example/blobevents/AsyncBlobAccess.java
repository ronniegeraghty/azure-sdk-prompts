package com.example.blobevents;

import reactor.core.publisher.Mono;

@FunctionalInterface
public interface AsyncBlobAccess {
    Mono<BlobSummary> download(String container, String blobName);
}

package com.example.blobevents.storage;

import com.example.blobevents.model.BlobSummary;
import reactor.core.publisher.Mono;

public interface BlobStore {
    BlobSummary download(String container, String blobName);

    Mono<BlobSummary> downloadAsync(String container, String blobName);
}

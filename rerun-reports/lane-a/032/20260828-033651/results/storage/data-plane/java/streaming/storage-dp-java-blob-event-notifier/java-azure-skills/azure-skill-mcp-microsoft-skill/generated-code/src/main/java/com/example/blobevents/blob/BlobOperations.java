package com.example.blobevents.blob;

import com.azure.core.util.BinaryData;
import reactor.core.publisher.Mono;

public interface BlobOperations {
    DownloadedBlob download(String container, String name);

    Mono<DownloadedBlob> downloadAsync(String container, String name);

    record DownloadedBlob(BinaryData content, BlobSummary summary) {
    }
}

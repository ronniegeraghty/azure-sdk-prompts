package com.example.blobevents.storage;

import com.azure.storage.blob.models.BlobStorageException;
import com.example.blobevents.model.BlobSummary;
import java.util.Map;
import reactor.core.publisher.Mono;

public final class InMemoryBlobStore implements BlobStore {
    private final Map<String, BlobSummary> blobs;

    public InMemoryBlobStore(Map<String, BlobSummary> blobs) {
        this.blobs = Map.copyOf(blobs);
    }

    @Override
    public BlobSummary download(String container, String blobName) {
        BlobSummary summary = blobs.get(container + "/" + blobName);
        if (summary == null) {
            throw new BlobStorageException("Mock blob was deleted before it could be read", null, null);
        }
        return summary;
    }

    @Override
    public Mono<BlobSummary> downloadAsync(String container, String blobName) {
        return Mono.fromCallable(() -> download(container, blobName));
    }
}

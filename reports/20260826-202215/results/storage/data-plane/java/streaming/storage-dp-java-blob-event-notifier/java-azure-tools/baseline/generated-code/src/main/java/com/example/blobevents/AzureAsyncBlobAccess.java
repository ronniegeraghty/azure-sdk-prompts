package com.example.blobevents;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobProperties;
import reactor.core.publisher.Mono;

public final class AzureAsyncBlobAccess implements AsyncBlobAccess {
    private final BlobServiceAsyncClient blobServiceClient;

    public AzureAsyncBlobAccess(BlobServiceAsyncClient blobServiceClient) {
        this.blobServiceClient = blobServiceClient;
    }

    @Override
    public Mono<BlobSummary> download(String container, String blobName) {
        BlobAsyncClient client = blobServiceClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(blobName);
        return client.downloadContent()
                .then(client.getProperties())
                .map(properties -> toSummary(blobName, properties));
    }

    private static BlobSummary toSummary(String blobName, BlobProperties properties) {
        return new BlobSummary(
                blobName,
                properties.getBlobSize(),
                properties.getContentType(),
                properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString());
    }
}

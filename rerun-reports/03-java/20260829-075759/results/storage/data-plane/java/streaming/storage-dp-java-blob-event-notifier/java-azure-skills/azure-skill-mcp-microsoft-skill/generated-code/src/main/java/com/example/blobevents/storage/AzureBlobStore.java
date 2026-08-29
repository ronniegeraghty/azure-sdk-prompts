package com.example.blobevents.storage;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;
import com.example.blobevents.model.BlobSummary;
import reactor.core.publisher.Mono;

public final class AzureBlobStore implements BlobStore {
    private final BlobServiceClient syncClient;
    private final BlobServiceAsyncClient asyncClient;

    public AzureBlobStore(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
    }

    @Override
    public BlobSummary download(String container, String blobName) {
        BlobClient blob = syncClient.getBlobContainerClient(container).getBlobClient(blobName);
        BlobProperties properties = blob.getProperties();
        blob.downloadContent();
        return toSummary(blobName, properties);
    }

    @Override
    public Mono<BlobSummary> downloadAsync(String container, String blobName) {
        BlobAsyncClient blob = asyncClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(blobName);
        return blob.getProperties()
            .flatMap(properties -> blob.downloadContent().thenReturn(toSummary(blobName, properties)));
    }

    private static BlobSummary toSummary(String name, BlobProperties properties) {
        String tier = properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString();
        return new BlobSummary(name, properties.getBlobSize(), properties.getContentType(), tier);
    }
}

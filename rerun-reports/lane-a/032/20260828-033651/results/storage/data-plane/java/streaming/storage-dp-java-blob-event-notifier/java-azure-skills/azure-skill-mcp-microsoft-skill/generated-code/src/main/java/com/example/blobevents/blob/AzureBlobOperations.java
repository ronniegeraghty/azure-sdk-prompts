package com.example.blobevents.blob;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;
import reactor.core.publisher.Mono;

public final class AzureBlobOperations implements BlobOperations {
    private final BlobServiceClient syncClient;
    private final BlobServiceAsyncClient asyncClient;

    public AzureBlobOperations(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
    }

    @Override
    public DownloadedBlob download(String container, String name) {
        BlobClient blob = syncClient.getBlobContainerClient(container).getBlobClient(name);
        BlobProperties properties = blob.getProperties();
        return new DownloadedBlob(blob.downloadContent(), summary(name, properties));
    }

    @Override
    public Mono<DownloadedBlob> downloadAsync(String container, String name) {
        BlobAsyncClient blob = asyncClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(name);
        return blob.getProperties()
            .flatMap(properties -> blob.downloadContent()
                .map(content -> new DownloadedBlob(content, summary(name, properties))));
    }

    private static BlobSummary summary(String name, BlobProperties properties) {
        String tier = properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString();
        return new BlobSummary(name, properties.getBlobSize(), properties.getContentType(), tier);
    }
}

package com.example.blobmanager;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.nio.file.Path;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageAsyncService {
    private final BlobContainerAsyncClient containerClient;

    public BlobStorageAsyncService(BlobServiceAsyncClient serviceClient, String containerName) {
        Objects.requireNonNull(serviceClient, "serviceClient");
        this.containerClient = serviceClient.getBlobContainerAsyncClient(
                Objects.requireNonNull(containerName, "containerName"));
    }

    public Mono<Void> ensureContainerExists() {
        return containerClient.createIfNotExists().then();
    }

    public Mono<BlobProperties> upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags
    ) {
        return upload(blobName, source, metadata, tags, null, null);
    }

    public Mono<BlobProperties> upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String expectedETag,
            String leaseId
    ) {
        BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
        BlobParallelUploadOptions options = BlobUploadOptionsFactory.create(
                source, metadata, tags, expectedETag, leaseId);

        return blobClient.uploadWithResponse(options).then(blobClient.getProperties());
    }

    public Mono<Void> download(String blobName, Path destination, boolean overwrite) {
        return containerClient.getBlobAsyncClient(blobName)
                .downloadToFile(destination.toString(), overwrite)
                .then();
    }

    public Flux<BlobItem> listBlobs() {
        return containerClient.listBlobs();
    }

    public Mono<Boolean> delete(String blobName) {
        return containerClient.getBlobAsyncClient(blobName).deleteIfExists();
    }

    public Mono<String> acquireLease(String blobName, int durationSeconds) {
        return leaseClient(blobName, null).acquireLease(durationSeconds);
    }

    public Mono<Void> releaseLease(String blobName, String leaseId) {
        return leaseClient(blobName, leaseId).releaseLease();
    }

    private BlobLeaseAsyncClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder()
                .blobAsyncClient(containerClient.getBlobAsyncClient(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildAsyncClient();
    }

}

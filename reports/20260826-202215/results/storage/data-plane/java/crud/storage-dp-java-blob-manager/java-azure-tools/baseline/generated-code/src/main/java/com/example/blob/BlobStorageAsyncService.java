package com.example.blob;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobListDetails;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.ListBlobsOptions;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import reactor.core.publisher.Mono;

import java.nio.file.Path;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageAsyncService {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;
    private static final Duration LEASE_DURATION = Duration.ofSeconds(60);

    private final BlobContainerAsyncClient containerClient;

    public BlobStorageAsyncService(BlobContainerAsyncClient containerClient) {
        this.containerClient = Objects.requireNonNull(containerClient, "containerClient");
    }

    public Mono<Void> upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        return upload(blobName, source, metadata, tags, null);
    }

    public Mono<Void> upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String leaseId) {
        BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
        ParallelTransferOptions transfer = new ParallelTransferOptions()
                .setBlockSizeLong(BLOCK_SIZE)
                .setMaxConcurrency(MAX_CONCURRENCY);

        return concurrencyConditions(blobClient, leaseId)
                .flatMap(conditions -> {
                    BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                            .setParallelTransferOptions(transfer)
                            .setMetadata(copyOrNull(metadata))
                            .setTags(copyOrNull(tags))
                            .setRequestConditions(conditions);
                    return blobClient.uploadFromFileWithResponse(options);
                })
                .then();
    }

    public Mono<Path> download(String blobName, Path destination) {
        return containerClient.getBlobAsyncClient(blobName)
                .downloadToFile(destination.toString(), true)
                .thenReturn(destination);
    }

    public Mono<List<BlobItem>> list() {
        ListBlobsOptions options = new ListBlobsOptions()
                .setDetails(new BlobListDetails().setRetrieveMetadata(true).setRetrieveTags(true));
        return containerClient.listBlobs(options).collectList();
    }

    public Mono<Boolean> delete(String blobName) {
        return containerClient.getBlobAsyncClient(blobName).deleteIfExists();
    }

    public Mono<String> acquireLease(String blobName) {
        return leaseClient(blobName, null).acquireLease((int) LEASE_DURATION.toSeconds());
    }

    public Mono<Void> releaseLease(String blobName, String leaseId) {
        return leaseClient(blobName, leaseId).releaseLease();
    }

    private Mono<BlobRequestConditions> concurrencyConditions(BlobAsyncClient blobClient, String leaseId) {
        return blobClient.getProperties()
                .map(properties -> new BlobRequestConditions().setIfMatch(properties.getETag()))
                .onErrorResume(BlobStorageException.class, e -> {
                    if (e.getStatusCode() == 404) {
                        return Mono.just(new BlobRequestConditions().setIfNoneMatch("*"));
                    }
                    return Mono.error(e);
                })
                .map(conditions -> leaseId == null ? conditions : conditions.setLeaseId(leaseId));
    }

    private BlobLeaseAsyncClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder()
                .blobAsyncClient(containerClient.getBlobAsyncClient(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildAsyncClient();
    }

    private static Map<String, String> copyOrNull(Map<String, String> values) {
        return values == null || values.isEmpty() ? null : Map.copyOf(values);
    }
}

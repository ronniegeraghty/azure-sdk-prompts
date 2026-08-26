package com.example.blobmanager;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import reactor.core.publisher.Mono;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Map;
import java.util.Objects;

public final class AsyncBlobStorageService {
    private final BlobContainerAsyncClient containerClient;
    private final ParallelTransferOptions transferOptions;

    public AsyncBlobStorageService(
            BlobServiceAsyncClient serviceClient,
            String containerName,
            ParallelTransferOptions transferOptions) {
        this.containerClient = Objects.requireNonNull(serviceClient, "serviceClient")
                .getBlobContainerAsyncClient(requireName(containerName, "containerName"));
        this.transferOptions = Objects.requireNonNull(transferOptions, "transferOptions");
    }

    public Mono<Void> ensureContainerExists() {
        return containerClient.createIfNotExists().then();
    }

    public Mono<BlobUploadResult> upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        requireReadableFile(source);
        BlobAsyncClient blobClient = blob(blobName);
        return optimisticCondition(blobClient)
                .flatMap(conditions -> upload(blobClient, source, metadata, tags, conditions));
    }

    public Mono<BlobUploadResult> uploadWithLease(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String leaseId) {
        requireReadableFile(source);
        if (leaseId == null || leaseId.isBlank()) {
            return Mono.error(new IllegalArgumentException("leaseId must not be blank"));
        }
        return upload(
                blob(blobName),
                source,
                metadata,
                tags,
                new BlobRequestConditions().setLeaseId(leaseId));
    }

    public Mono<Void> download(String blobName, Path destination) {
        Objects.requireNonNull(destination, "destination");
        return Mono.fromRunnable(() -> createParentDirectories(destination))
                .then(blob(blobName).downloadToFile(destination.toString(), true))
                .then();
    }

    public Mono<List<BlobItem>> listBlobs() {
        return containerClient.listBlobs().collectList();
    }

    public Mono<Boolean> delete(String blobName) {
        return blob(blobName).deleteIfExists();
    }

    public Mono<String> acquireLease(String blobName, int leaseDurationSeconds) {
        if (leaseDurationSeconds < 15 || leaseDurationSeconds > 60) {
            return Mono.error(
                    new IllegalArgumentException("leaseDurationSeconds must be between 15 and 60"));
        }
        return leaseClient(blobName).acquireLease(leaseDurationSeconds);
    }

    public Mono<Void> releaseLease(String blobName, String leaseId) {
        return leaseClient(blobName, leaseId).releaseLease();
    }

    private Mono<BlobUploadResult> upload(
            BlobAsyncClient blobClient,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            BlobRequestConditions conditions) {
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setParallelTransferOptions(transferOptions)
                .setMetadata(emptyIfNull(metadata))
                .setTags(emptyIfNull(tags))
                .setRequestConditions(conditions);

        return blobClient.uploadFromFileWithResponse(options)
                .map(response -> new BlobUploadResult(
                        blobClient.getBlobName(),
                        response.getValue().getETag(),
                        response.getValue().getVersionId()));
    }

    private Mono<BlobRequestConditions> optimisticCondition(BlobAsyncClient blobClient) {
        return blobClient.getProperties()
                .map(properties -> new BlobRequestConditions().setIfMatch(properties.getETag()))
                .onErrorResume(
                        BlobStorageException.class,
                        exception -> exception.getStatusCode() == 404
                                ? Mono.just(new BlobRequestConditions().setIfNoneMatch("*"))
                                : Mono.error(exception));
    }

    private BlobAsyncClient blob(String blobName) {
        return containerClient.getBlobAsyncClient(requireName(blobName, "blobName"));
    }

    private BlobLeaseAsyncClient leaseClient(String blobName) {
        return new BlobLeaseClientBuilder()
                .blobAsyncClient(blob(blobName))
                .buildAsyncClient();
    }

    private BlobLeaseAsyncClient leaseClient(String blobName, String leaseId) {
        return new BlobLeaseClientBuilder()
                .blobAsyncClient(blob(blobName))
                .leaseId(leaseId)
                .buildAsyncClient();
    }

    private static void requireReadableFile(Path source) {
        Objects.requireNonNull(source, "source");
        if (!Files.isRegularFile(source) || !Files.isReadable(source)) {
            throw new IllegalArgumentException("Source must be a readable regular file: " + source);
        }
    }

    private static void createParentDirectories(Path destination) {
        Path parent = destination.toAbsolutePath().getParent();
        if (parent == null) {
            return;
        }
        try {
            Files.createDirectories(parent);
        } catch (java.io.IOException exception) {
            throw new IllegalStateException("Could not create destination directory: " + parent, exception);
        }
    }

    private static String requireName(String value, String parameter) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(parameter + " must not be blank");
        }
        return value;
    }

    private static Map<String, String> emptyIfNull(Map<String, String> values) {
        return values == null ? Map.of() : Map.copyOf(values);
    }
}

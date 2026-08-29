package com.example.blobmanager;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobErrorCode;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.BlockBlobItem;
import com.azure.storage.blob.models.DeleteSnapshotsOptionType;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobDownloadToFileOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.util.Map;
import java.util.Objects;
import java.util.Set;

public final class AsyncBlobStorageService {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private final BlobContainerAsyncClient containerClient;

    public AsyncBlobStorageService(BlobContainerAsyncClient containerClient) {
        this.containerClient = Objects.requireNonNull(containerClient, "containerClient");
    }

    public Mono<BlockBlobItem> upload(
        Path source,
        String blobName,
        Map<String, String> metadata,
        Map<String, String> tags
    ) {
        return upload(source, blobName, metadata, tags, null);
    }

    public Mono<BlockBlobItem> upload(
        Path source,
        String blobName,
        Map<String, String> metadata,
        Map<String, String> tags,
        String leaseId
    ) {
        Objects.requireNonNull(source, "source");
        BlobAsyncClient blobClient = blob(blobName);
        return writeConditions(blobClient, leaseId)
            .flatMap(conditions -> {
                BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                    .setMetadata(metadata)
                    .setTags(tags)
                    .setParallelTransferOptions(transferOptions())
                    .setRequestConditions(conditions);
                return blobClient.uploadFromFileWithResponse(options);
            })
            .map(response -> response.getValue());
    }

    public Mono<Path> download(String blobName, Path destination) {
        Objects.requireNonNull(destination, "destination");
        BlobDownloadToFileOptions options = new BlobDownloadToFileOptions(destination.toString())
            .setOpenOptions(Set.of(
                StandardOpenOption.CREATE,
                StandardOpenOption.TRUNCATE_EXISTING,
                StandardOpenOption.WRITE));
        return blob(blobName).downloadToFileWithResponse(options)
            .thenReturn(destination);
    }

    public Flux<BlobItem> listBlobs() {
        return containerClient.listBlobs();
    }

    public Mono<Boolean> delete(String blobName) {
        return blob(blobName)
            .deleteIfExistsWithResponse(DeleteSnapshotsOptionType.INCLUDE, null)
            .map(response -> response.getValue());
    }

    public Mono<String> acquireLease(String blobName, int leaseDurationSeconds) {
        return leaseClient(blobName, null).acquireLease(leaseDurationSeconds);
    }

    public Mono<Void> releaseLease(String blobName, String leaseId) {
        return leaseClient(blobName, leaseId).releaseLease();
    }

    private BlobLeaseAsyncClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder().blobAsyncClient(blob(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildAsyncClient();
    }

    private Mono<BlobRequestConditions> writeConditions(BlobAsyncClient blobClient, String leaseId) {
        BlobRequestConditions conditions = new BlobRequestConditions();
        if (leaseId != null) {
            conditions.setLeaseId(leaseId);
        }
        return blobClient.getProperties()
            .map(properties -> conditions.setIfMatch(properties.getETag()))
            .onErrorResume(BlobStorageException.class, exception -> {
                if (isNotFound(exception)) {
                    return Mono.just(conditions.setIfNoneMatch("*"));
                }
                return Mono.error(exception);
            });
    }

    private BlobAsyncClient blob(String blobName) {
        if (blobName == null || blobName.isBlank()) {
            throw new IllegalArgumentException("blobName must not be blank");
        }
        return containerClient.getBlobAsyncClient(blobName);
    }

    private static ParallelTransferOptions transferOptions() {
        return new ParallelTransferOptions()
            .setBlockSizeLong(BLOCK_SIZE)
            .setMaxConcurrency(MAX_CONCURRENCY);
    }

    private static boolean isNotFound(BlobStorageException exception) {
        return exception.getStatusCode() == 404
            || exception.getErrorCode() == BlobErrorCode.BLOB_NOT_FOUND
            || exception.getErrorCode() == BlobErrorCode.RESOURCE_NOT_FOUND;
    }
}

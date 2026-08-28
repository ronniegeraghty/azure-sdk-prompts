package com.example.blobmanager;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.nio.file.Path;
import java.util.Map;
import java.util.Objects;

public final class AzureBlobManagerAsync {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final long MAX_SINGLE_UPLOAD_SIZE = 32L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private final BlobContainerAsyncClient containerClient;

    public AzureBlobManagerAsync(BlobServiceAsyncClient serviceClient, String containerName) {
        this.containerClient = Objects.requireNonNull(serviceClient, "serviceClient")
                .getBlobContainerAsyncClient(requireText(containerName, "containerName"));
    }

    public Mono<String> upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            BlobWriteCondition writeCondition) {
        ParallelTransferOptions transferOptions = new ParallelTransferOptions()
                .setBlockSizeLong(BLOCK_SIZE)
                .setMaxSingleUploadSizeLong(MAX_SINGLE_UPLOAD_SIZE)
                .setMaxConcurrency(MAX_CONCURRENCY);
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setParallelTransferOptions(transferOptions)
                .setMetadata(metadata)
                .setTags(tags)
                .setRequestConditions(Objects.requireNonNull(writeCondition, "writeCondition")
                        .toRequestConditions());

        return blob(blobName)
                .uploadFromFileWithResponse(options)
                .map(response -> response.getValue().getETag());
    }

    public Mono<Void> download(String blobName, Path destination, boolean overwrite) {
        return blob(blobName)
                .downloadToFile(destination.toString(), overwrite)
                .then();
    }

    public Flux<BlobItem> list() {
        return containerClient.listBlobs();
    }

    public Mono<Boolean> delete(String blobName) {
        return blob(blobName).deleteIfExists();
    }

    public Mono<String> acquireLease(String blobName, int durationSeconds) {
        return leaseClient(blobName, null)
                .acquireLease(durationSeconds);
    }

    public Mono<Void> releaseLease(String blobName, String leaseId) {
        return leaseClient(blobName, requireText(leaseId, "leaseId"))
                .releaseLease();
    }

    private BlobAsyncClient blob(String blobName) {
        return containerClient.getBlobAsyncClient(requireText(blobName, "blobName"));
    }

    private BlobLeaseAsyncClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder().blobAsyncClient(blob(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildAsyncClient();
    }

    private static String requireText(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }
}

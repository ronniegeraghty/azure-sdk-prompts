package com.example.blobmanager;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.DownloadRetryOptions;
import com.azure.storage.blob.options.BlobDownloadToFileOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import com.azure.storage.common.ParallelTransferOptions;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.time.Duration;
import java.util.Map;

public final class AsyncBlobStorageService {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private final BlobServiceAsyncClient serviceClient;
    private final Duration requestTimeout;

    public AsyncBlobStorageService(BlobServiceAsyncClient serviceClient, Duration requestTimeout) {
        this.serviceClient = serviceClient;
        this.requestTimeout = requestTimeout;
    }

    public Mono<Void> upload(
            String containerName,
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        BlobAsyncClient blobClient = blobClient(containerName, blobName);
        return concurrencyConditions(blobClient)
                .flatMap(conditions -> upload(blobClient, source, metadata, tags, conditions));
    }

    public Mono<Void> download(String containerName, String blobName, Path destination) {
        BlobDownloadToFileOptions options = new BlobDownloadToFileOptions(destination.toString())
                .setParallelTransferOptions(downloadTransferOptions())
                .setDownloadRetryOptions(new DownloadRetryOptions().setMaxRetryRequests(3))
                .setOpenOptions(java.util.Set.of(
                        StandardOpenOption.CREATE,
                        StandardOpenOption.WRITE,
                        StandardOpenOption.TRUNCATE_EXISTING));
        return blobClient(containerName, blobName)
                .downloadToFileWithResponse(options)
                .then();
    }

    public Flux<BlobItem> list(String containerName) {
        return serviceClient.getBlobContainerAsyncClient(containerName).listBlobs();
    }

    public Mono<Boolean> delete(String containerName, String blobName) {
        return blobClient(containerName, blobName).deleteIfExists().timeout(requestTimeout);
    }

    public Mono<Void> overwriteWithLease(
            String containerName,
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        BlobAsyncClient blobClient = blobClient(containerName, blobName);
        BlobLeaseAsyncClient leaseClient = new BlobLeaseClientBuilder()
                .blobAsyncClient(blobClient)
                .buildAsyncClient();

        return Mono.usingWhen(
                leaseClient.acquireLease(60).timeout(requestTimeout),
                leaseId -> blobClient.getProperties()
                        .timeout(requestTimeout)
                        .flatMap(properties -> {
                            BlobRequestConditions conditions = new BlobRequestConditions()
                                    .setLeaseId(leaseId)
                                    .setIfMatch(properties.getETag());
                            return upload(blobClient, source, metadata, tags, conditions);
                        }),
                ignored -> leaseClient.releaseLease().timeout(requestTimeout),
                (ignored, error) -> leaseClient.releaseLease().timeout(requestTimeout),
                ignored -> leaseClient.releaseLease().timeout(requestTimeout));
    }

    private Mono<Void> upload(
            BlobAsyncClient blobClient,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            BlobRequestConditions conditions) {
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setParallelTransferOptions(uploadTransferOptions())
                .setMetadata(metadata == null ? Map.of() : metadata)
                .setTags(tags == null ? Map.of() : tags)
                .setRequestConditions(conditions);
        return blobClient.uploadFromFileWithResponse(options)
                .then();
    }

    private Mono<BlobRequestConditions> concurrencyConditions(BlobAsyncClient blobClient) {
        return blobClient.exists()
                .timeout(requestTimeout)
                .flatMap(exists -> exists
                        ? blobClient.getProperties()
                                .timeout(requestTimeout)
                                .map(properties -> new BlobRequestConditions()
                                        .setIfMatch(properties.getETag()))
                        : Mono.just(new BlobRequestConditions().setIfNoneMatch("*")));
    }

    private BlobAsyncClient blobClient(String containerName, String blobName) {
        return serviceClient.getBlobContainerAsyncClient(containerName).getBlobAsyncClient(blobName);
    }

    private static ParallelTransferOptions downloadTransferOptions() {
        return new ParallelTransferOptions()
                .setBlockSizeLong(BLOCK_SIZE)
                .setMaxSingleUploadSizeLong(BLOCK_SIZE)
                .setMaxConcurrency(MAX_CONCURRENCY);
    }

    private static com.azure.storage.blob.models.ParallelTransferOptions uploadTransferOptions() {
        return new com.azure.storage.blob.models.ParallelTransferOptions()
                .setBlockSizeLong(BLOCK_SIZE)
                .setMaxSingleUploadSizeLong(BLOCK_SIZE)
                .setMaxConcurrency(MAX_CONCURRENCY);
    }
}

package com.example.blobmanager;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseAsyncClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.nio.file.Path;
import java.time.Duration;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageAsyncService {
    private final BlobContainerAsyncClient containerClient;

    public BlobStorageAsyncService(BlobServiceAsyncClient serviceClient, String containerName) {
        Objects.requireNonNull(serviceClient, "serviceClient");
        this.containerClient = serviceClient.getBlobContainerAsyncClient(requireName(containerName, "containerName"));
    }

    public Mono<BlobUploadResult> upload(
            Path source,
            String blobName,
            Map<String, String> metadata,
            Map<String, String> tags) {
        return upload(source, blobName, metadata, tags, null, null);
    }

    public Mono<BlobUploadResult> upload(
            Path source,
            String blobName,
            Map<String, String> metadata,
            Map<String, String> tags,
            String expectedETag,
            String leaseId) {
        Objects.requireNonNull(source, "source");
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setMetadata(metadata)
                .setTags(tags)
                .setParallelTransferOptions(BlobStorageService.transferOptions())
                .setRequestConditions(BlobStorageService.writeConditions(expectedETag, leaseId));

        return blobClient(blobName)
                .uploadFromFileWithResponse(options)
                .map(response -> new BlobUploadResult(blobName, response.getValue().getETag()));
    }

    public Mono<Void> download(String blobName, Path destination, boolean overwrite) {
        Objects.requireNonNull(destination, "destination");
        return blobClient(blobName).downloadToFile(destination.toString(), overwrite).then();
    }

    public Flux<BlobItem> listBlobs() {
        return containerClient.listBlobs();
    }

    public Mono<Boolean> delete(String blobName) {
        return blobClient(blobName).deleteIfExists().defaultIfEmpty(false);
    }

    public Mono<String> getETag(String blobName) {
        return blobClient(blobName).getProperties().map(properties -> properties.getETag());
    }

    public Mono<String> acquireLease(String blobName, Duration duration) {
        int seconds = Math.toIntExact(Objects.requireNonNull(duration, "duration").toSeconds());
        if (seconds < 15 || seconds > 60) {
            return Mono.error(new IllegalArgumentException("A finite blob lease must be between 15 and 60 seconds"));
        }
        return leaseClient(blobName, null).acquireLease(seconds);
    }

    public Mono<Void> releaseLease(String blobName, String leaseId) {
        return leaseClient(blobName, requireName(leaseId, "leaseId")).releaseLease();
    }

    private BlobAsyncClient blobClient(String blobName) {
        return containerClient.getBlobAsyncClient(requireName(blobName, "blobName"));
    }

    private BlobLeaseAsyncClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder().blobAsyncClient(blobClient(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildAsyncClient();
    }

    private static String requireName(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }
}

package com.example.blobmanager;

import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobErrorCode;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.BlockBlobItem;
import com.azure.storage.blob.models.DeleteSnapshotsOptionType;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobDownloadToFileOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;

import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Set;

public final class BlobStorageService {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private final BlobContainerClient containerClient;
    private final Duration requestTimeout;

    public BlobStorageService(BlobContainerClient containerClient, Duration requestTimeout) {
        this.containerClient = Objects.requireNonNull(containerClient, "containerClient");
        this.requestTimeout = Objects.requireNonNull(requestTimeout, "requestTimeout");
    }

    public BlockBlobItem upload(
        Path source,
        String blobName,
        Map<String, String> metadata,
        Map<String, String> tags
    ) {
        return upload(source, blobName, metadata, tags, null);
    }

    public BlockBlobItem upload(
        Path source,
        String blobName,
        Map<String, String> metadata,
        Map<String, String> tags,
        String leaseId
    ) {
        Objects.requireNonNull(source, "source");
        BlobClient blobClient = blob(blobName);
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
            .setMetadata(metadata)
            .setTags(tags)
            .setParallelTransferOptions(transferOptions())
            .setRequestConditions(writeConditions(blobClient, leaseId));

        return blobClient.uploadFromFileWithResponse(options, requestTimeout, Context.NONE).getValue();
    }

    public Path download(String blobName, Path destination) {
        Objects.requireNonNull(destination, "destination");
        BlobDownloadToFileOptions options = new BlobDownloadToFileOptions(destination.toString())
            .setOpenOptions(Set.of(
                StandardOpenOption.CREATE,
                StandardOpenOption.TRUNCATE_EXISTING,
                StandardOpenOption.WRITE));
        blob(blobName).downloadToFileWithResponse(options, requestTimeout, Context.NONE);
        return destination;
    }

    public List<BlobItem> listBlobs() {
        return containerClient.listBlobs().stream().toList();
    }

    public boolean delete(String blobName) {
        Response<Boolean> response = blob(blobName).deleteIfExistsWithResponse(
            DeleteSnapshotsOptionType.INCLUDE, null, requestTimeout, Context.NONE);
        return response.getValue();
    }

    public String acquireLease(String blobName, int leaseDurationSeconds) {
        return leaseClient(blobName, null).acquireLease(leaseDurationSeconds);
    }

    public void releaseLease(String blobName, String leaseId) {
        leaseClient(blobName, leaseId).releaseLease();
    }

    private BlobLeaseClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder().blobClient(blob(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildClient();
    }

    private BlobRequestConditions writeConditions(BlobClient blobClient, String leaseId) {
        BlobRequestConditions conditions = new BlobRequestConditions();
        if (leaseId != null) {
            conditions.setLeaseId(leaseId);
        }
        try {
            return conditions.setIfMatch(blobClient.getProperties().getETag());
        } catch (BlobStorageException exception) {
            if (isNotFound(exception)) {
                return conditions.setIfNoneMatch("*");
            }
            throw exception;
        }
    }

    private BlobClient blob(String blobName) {
        if (blobName == null || blobName.isBlank()) {
            throw new IllegalArgumentException("blobName must not be blank");
        }
        return containerClient.getBlobClient(blobName);
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

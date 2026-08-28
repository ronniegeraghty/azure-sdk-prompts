package com.example.blobmanager;

import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlockBlobItem;
import com.azure.storage.blob.models.DeleteSnapshotsOptionType;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;

import java.nio.file.Path;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageService {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final long SINGLE_UPLOAD_THRESHOLD = 32L * 1024 * 1024;
    private static final int TRANSFER_CONCURRENCY = 4;

    private final BlobContainerClient containerClient;

    public BlobStorageService(BlobServiceClient serviceClient, String containerName) {
        Objects.requireNonNull(serviceClient, "serviceClient");
        this.containerClient = serviceClient.getBlobContainerClient(requireName(containerName, "containerName"));
    }

    public BlobUploadResult upload(Path source, String blobName, Map<String, String> metadata, Map<String, String> tags) {
        return upload(source, blobName, metadata, tags, null, null);
    }

    public BlobUploadResult upload(
            Path source,
            String blobName,
            Map<String, String> metadata,
            Map<String, String> tags,
            String expectedETag,
            String leaseId) {
        Objects.requireNonNull(source, "source");
        BlobClient blobClient = blobClient(blobName);
        BlobRequestConditions conditions = writeConditions(expectedETag, leaseId);
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setMetadata(metadata)
                .setTags(tags)
                .setParallelTransferOptions(transferOptions())
                .setRequestConditions(conditions);

        Response<BlockBlobItem> response = blobClient.uploadFromFileWithResponse(options, null, Context.NONE);
        return new BlobUploadResult(blobName, response.getValue().getETag());
    }

    public void download(String blobName, Path destination, boolean overwrite) {
        Objects.requireNonNull(destination, "destination");
        blobClient(blobName).downloadToFile(destination.toString(), overwrite);
    }

    public List<BlobItem> listBlobs() {
        return containerClient.listBlobs().stream().toList();
    }

    public boolean delete(String blobName) {
        return blobClient(blobName)
                .deleteIfExistsWithResponse(DeleteSnapshotsOptionType.INCLUDE, null, null, Context.NONE)
                .getValue();
    }

    public String getETag(String blobName) {
        return blobClient(blobName).getProperties().getETag();
    }

    public String acquireLease(String blobName, Duration duration) {
        int seconds = Math.toIntExact(Objects.requireNonNull(duration, "duration").toSeconds());
        if (seconds < 15 || seconds > 60) {
            throw new IllegalArgumentException("A finite blob lease must be between 15 and 60 seconds");
        }
        return leaseClient(blobName).acquireLease(seconds);
    }

    public void releaseLease(String blobName, String leaseId) {
        leaseClient(blobName, leaseId).releaseLease();
    }

    private BlobClient blobClient(String blobName) {
        return containerClient.getBlobClient(requireName(blobName, "blobName"));
    }

    private BlobLeaseClient leaseClient(String blobName) {
        return new BlobLeaseClientBuilder().blobClient(blobClient(blobName)).buildClient();
    }

    private BlobLeaseClient leaseClient(String blobName, String leaseId) {
        return new BlobLeaseClientBuilder()
                .blobClient(blobClient(blobName))
                .leaseId(requireName(leaseId, "leaseId"))
                .buildClient();
    }

    static ParallelTransferOptions transferOptions() {
        return new ParallelTransferOptions()
                .setBlockSizeLong(BLOCK_SIZE)
                .setMaxSingleUploadSizeLong(SINGLE_UPLOAD_THRESHOLD)
                .setMaxConcurrency(TRANSFER_CONCURRENCY);
    }

    static BlobRequestConditions writeConditions(String expectedETag, String leaseId) {
        BlobRequestConditions conditions = new BlobRequestConditions();
        if (expectedETag == null || expectedETag.isBlank()) {
            conditions.setIfNoneMatch("*");
        } else {
            conditions.setIfMatch(expectedETag);
        }
        if (leaseId != null && !leaseId.isBlank()) {
            conditions.setLeaseId(leaseId);
        }
        return conditions;
    }

    private static String requireName(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }
}

package com.example.blobmanager;

import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;

import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;

public final class AzureBlobManager {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final long MAX_SINGLE_UPLOAD_SIZE = 32L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private final BlobContainerClient containerClient;

    public AzureBlobManager(BlobServiceClient serviceClient, String containerName) {
        this.containerClient = Objects.requireNonNull(serviceClient, "serviceClient")
                .getBlobContainerClient(requireText(containerName, "containerName"));
    }

    public String upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            BlobWriteCondition writeCondition) {
        BlobClient blob = blob(blobName);
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

        return blob.uploadFromFileWithResponse(options, null, Context.NONE)
                .getValue()
                .getETag();
    }

    public void download(String blobName, Path destination, boolean overwrite) {
        blob(blobName).downloadToFile(destination.toString(), overwrite);
    }

    public List<BlobItem> list() {
        List<BlobItem> blobs = new ArrayList<>();
        containerClient.listBlobs().forEach(blobs::add);
        return List.copyOf(blobs);
    }

    public boolean delete(String blobName) {
        return blob(blobName).deleteIfExists();
    }

    public String acquireLease(String blobName, int durationSeconds) {
        BlobLeaseClient leaseClient = new BlobLeaseClientBuilder()
                .blobClient(blob(blobName))
                .buildClient();
        return leaseClient.acquireLease(durationSeconds);
    }

    public void releaseLease(String blobName, String leaseId) {
        new BlobLeaseClientBuilder()
                .blobClient(blob(blobName))
                .leaseId(requireText(leaseId, "leaseId"))
                .buildClient()
                .releaseLease();
    }

    private BlobClient blob(String blobName) {
        return containerClient.getBlobClient(requireText(blobName, "blobName"));
    }

    private static String requireText(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }
}

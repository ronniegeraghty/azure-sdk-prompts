package com.example.blob;

import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobListDetails;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.ListBlobsOptions;
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
    private static final int MAX_CONCURRENCY = 4;
    private static final Duration LEASE_DURATION = Duration.ofSeconds(60);

    private final BlobContainerClient containerClient;

    public BlobStorageService(BlobContainerClient containerClient) {
        this.containerClient = Objects.requireNonNull(containerClient, "containerClient");
    }

    public void upload(String blobName, Path source, Map<String, String> metadata, Map<String, String> tags) {
        upload(blobName, source, metadata, tags, null);
    }

    public void upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String leaseId) {
        BlobClient blobClient = containerClient.getBlobClient(blobName);
        BlobRequestConditions conditions = concurrencyConditions(blobClient, leaseId);
        ParallelTransferOptions transfer = new ParallelTransferOptions()
                .setBlockSizeLong(BLOCK_SIZE)
                .setMaxConcurrency(MAX_CONCURRENCY);

        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setParallelTransferOptions(transfer)
                .setMetadata(copyOrNull(metadata))
                .setTags(copyOrNull(tags))
                .setRequestConditions(conditions);
        blobClient.uploadFromFileWithResponse(options, null, Context.NONE);
    }

    public Path download(String blobName, Path destination) {
        containerClient.getBlobClient(blobName).downloadToFile(destination.toString(), true);
        return destination;
    }

    public List<BlobItem> list() {
        ListBlobsOptions options = new ListBlobsOptions()
                .setDetails(new BlobListDetails().setRetrieveMetadata(true).setRetrieveTags(true));
        return containerClient.listBlobs(options, null).stream().toList();
    }

    public boolean delete(String blobName) {
        return containerClient.getBlobClient(blobName).deleteIfExists();
    }

    public String acquireLease(String blobName) {
        return leaseClient(blobName, null).acquireLease((int) LEASE_DURATION.toSeconds());
    }

    public void releaseLease(String blobName, String leaseId) {
        leaseClient(blobName, leaseId).releaseLease();
    }

    private BlobRequestConditions concurrencyConditions(BlobClient blobClient, String leaseId) {
        BlobRequestConditions conditions = new BlobRequestConditions();
        try {
            conditions.setIfMatch(blobClient.getProperties().getETag());
        } catch (BlobStorageException e) {
            if (e.getStatusCode() != 404) {
                throw e;
            }
            conditions.setIfNoneMatch("*");
        }
        if (leaseId != null) {
            conditions.setLeaseId(leaseId);
        }
        return conditions;
    }

    private BlobLeaseClient leaseClient(String blobName, String leaseId) {
        BlobLeaseClientBuilder builder = new BlobLeaseClientBuilder()
                .blobClient(containerClient.getBlobClient(blobName));
        if (leaseId != null) {
            builder.leaseId(leaseId);
        }
        return builder.buildClient();
    }

    private static Map<String, String> copyOrNull(Map<String, String> values) {
        return values == null || values.isEmpty() ? null : Map.copyOf(values);
    }
}

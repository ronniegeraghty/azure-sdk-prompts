package com.example.blobmanager;

import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.BlockBlobItem;
import com.azure.storage.blob.models.ListBlobsOptions;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageService {
    private final BlobContainerClient containerClient;
    private final ParallelTransferOptions transferOptions;

    public BlobStorageService(
            BlobServiceClient serviceClient,
            String containerName,
            ParallelTransferOptions transferOptions) {
        this.containerClient = Objects.requireNonNull(serviceClient, "serviceClient")
                .getBlobContainerClient(requireName(containerName, "containerName"));
        this.transferOptions = Objects.requireNonNull(transferOptions, "transferOptions");
    }

    public void ensureContainerExists() {
        containerClient.createIfNotExists();
    }

    public BlobUploadResult upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        return upload(blobName, source, metadata, tags, optimisticCondition(blob(blobName)));
    }

    public BlobUploadResult uploadWithLease(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String leaseId) {
        if (leaseId == null || leaseId.isBlank()) {
            throw new IllegalArgumentException("leaseId must not be blank");
        }
        return upload(
                blobName,
                source,
                metadata,
                tags,
                new BlobRequestConditions().setLeaseId(leaseId));
    }

    public void download(String blobName, Path destination) {
        Objects.requireNonNull(destination, "destination");
        createParentDirectories(destination);
        blob(blobName).downloadToFile(destination.toString(), true);
    }

    public List<BlobItem> listBlobs() {
        return containerClient.listBlobs(new ListBlobsOptions(), null).stream().toList();
    }

    public boolean delete(String blobName) {
        return blob(blobName).deleteIfExists();
    }

    public String acquireLease(String blobName, int leaseDurationSeconds) {
        if (leaseDurationSeconds < 15 || leaseDurationSeconds > 60) {
            throw new IllegalArgumentException("leaseDurationSeconds must be between 15 and 60");
        }
        return leaseClient(blobName).acquireLease(leaseDurationSeconds);
    }

    public void releaseLease(String blobName, String leaseId) {
        leaseClient(blobName, leaseId).releaseLease();
    }

    private BlobUploadResult upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            BlobRequestConditions conditions) {
        requireReadableFile(source);
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setParallelTransferOptions(transferOptions)
                .setMetadata(emptyIfNull(metadata))
                .setTags(emptyIfNull(tags))
                .setRequestConditions(conditions);

        Response<BlockBlobItem> response = blob(blobName)
                .uploadFromFileWithResponse(options, null, Context.NONE);
        BlockBlobItem value = response.getValue();
        return new BlobUploadResult(blobName, value.getETag(), value.getVersionId());
    }

    private BlobRequestConditions optimisticCondition(BlobClient blobClient) {
        try {
            String eTag = blobClient.getProperties().getETag();
            return new BlobRequestConditions().setIfMatch(eTag);
        } catch (BlobStorageException exception) {
            if (exception.getStatusCode() == 404) {
                return new BlobRequestConditions().setIfNoneMatch("*");
            }
            throw exception;
        }
    }

    private BlobClient blob(String blobName) {
        return containerClient.getBlobClient(requireName(blobName, "blobName"));
    }

    private BlobLeaseClient leaseClient(String blobName) {
        return new BlobLeaseClientBuilder().blobClient(blob(blobName)).buildClient();
    }

    private BlobLeaseClient leaseClient(String blobName, String leaseId) {
        return new BlobLeaseClientBuilder()
                .blobClient(blob(blobName))
                .leaseId(leaseId)
                .buildClient();
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

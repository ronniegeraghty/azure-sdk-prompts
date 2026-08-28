package com.example.blobmanager;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import com.azure.storage.blob.specialized.BlobLeaseClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;

import java.nio.file.Path;
import java.util.List;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageService {
    private final BlobContainerClient containerClient;

    public BlobStorageService(BlobServiceClient serviceClient, String containerName) {
        Objects.requireNonNull(serviceClient, "serviceClient");
        this.containerClient = serviceClient.getBlobContainerClient(
                Objects.requireNonNull(containerName, "containerName"));
    }

    public void ensureContainerExists() {
        containerClient.createIfNotExists();
    }

    public BlobProperties upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags
    ) {
        return upload(blobName, source, metadata, tags, null, null);
    }

    public BlobProperties upload(
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String expectedETag,
            String leaseId
    ) {
        BlobClient blobClient = containerClient.getBlobClient(blobName);
        BlobParallelUploadOptions options = BlobUploadOptionsFactory.create(
                source, metadata, tags, expectedETag, leaseId);

        blobClient.uploadWithResponse(options, null, null);
        return blobClient.getProperties();
    }

    public void download(String blobName, Path destination, boolean overwrite) {
        containerClient.getBlobClient(blobName).downloadToFile(destination.toString(), overwrite);
    }

    public List<BlobItem> listBlobs() {
        return containerClient.listBlobs().stream().toList();
    }

    public boolean delete(String blobName) {
        return containerClient.getBlobClient(blobName).deleteIfExists();
    }

    public Lease acquireLease(String blobName, int durationSeconds) {
        BlobLeaseClient leaseClient = new BlobLeaseClientBuilder()
                .blobClient(containerClient.getBlobClient(blobName))
                .buildClient();
        return new Lease(leaseClient.acquireLease(durationSeconds), leaseClient);
    }

    public static final class Lease implements AutoCloseable {
        private final String leaseId;
        private final BlobLeaseClient leaseClient;
        private boolean released;

        private Lease(String leaseId, BlobLeaseClient leaseClient) {
            this.leaseId = leaseId;
            this.leaseClient = leaseClient;
        }

        public String leaseId() {
            return leaseId;
        }

        @Override
        public void close() {
            if (!released) {
                leaseClient.releaseLease();
                released = true;
            }
        }
    }
}

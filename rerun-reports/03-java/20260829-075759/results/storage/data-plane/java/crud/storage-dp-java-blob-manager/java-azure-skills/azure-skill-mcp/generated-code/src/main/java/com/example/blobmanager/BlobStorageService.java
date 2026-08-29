package com.example.blobmanager;

import com.azure.core.http.rest.PagedIterable;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.DownloadRetryOptions;
import com.azure.storage.blob.options.BlobDownloadToFileOptions;
import com.azure.storage.blob.options.BlobUploadFromFileOptions;
import com.azure.storage.blob.specialized.BlobLeaseClient;
import com.azure.storage.blob.specialized.BlobLeaseClientBuilder;
import com.azure.storage.common.ParallelTransferOptions;

import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.util.List;
import java.util.Map;

public final class BlobStorageService {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private final BlobServiceClient serviceClient;

    public BlobStorageService(BlobServiceClient serviceClient) {
        this.serviceClient = serviceClient;
    }

    public void upload(
            String containerName,
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        BlobClient blobClient = blobClient(containerName, blobName);
        upload(blobClient, source, metadata, tags, concurrencyConditions(blobClient));
    }

    public void download(String containerName, String blobName, Path destination) {
        BlobDownloadToFileOptions options = new BlobDownloadToFileOptions(destination.toString())
                .setParallelTransferOptions(downloadTransferOptions())
                .setDownloadRetryOptions(new DownloadRetryOptions().setMaxRetryRequests(3))
                .setOpenOptions(java.util.Set.of(
                        StandardOpenOption.CREATE,
                        StandardOpenOption.WRITE,
                        StandardOpenOption.TRUNCATE_EXISTING));
        blobClient(containerName, blobName)
                .downloadToFileWithResponse(options, null, Context.NONE);
    }

    public List<BlobItem> list(String containerName) {
        PagedIterable<BlobItem> blobs = serviceClient.getBlobContainerClient(containerName).listBlobs();
        return blobs.stream().toList();
    }

    public boolean delete(String containerName, String blobName) {
        return blobClient(containerName, blobName).deleteIfExists();
    }

    public void overwriteWithLease(
            String containerName,
            String blobName,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags) {
        BlobClient blobClient = blobClient(containerName, blobName);
        BlobLeaseClient leaseClient = new BlobLeaseClientBuilder().blobClient(blobClient).buildClient();
        String leaseId = leaseClient.acquireLease(60);
        try {
            BlobProperties properties = blobClient.getProperties();
            BlobRequestConditions conditions = new BlobRequestConditions()
                    .setLeaseId(leaseId)
                    .setIfMatch(properties.getETag());
            upload(blobClient, source, metadata, tags, conditions);
        } finally {
            leaseClient.releaseLease();
        }
    }

    private void upload(
            BlobClient blobClient,
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            BlobRequestConditions conditions) {
        BlobUploadFromFileOptions options = new BlobUploadFromFileOptions(source.toString())
                .setParallelTransferOptions(uploadTransferOptions())
                .setMetadata(metadata == null ? Map.of() : metadata)
                .setTags(tags == null ? Map.of() : tags)
                .setRequestConditions(conditions);
        blobClient.uploadFromFileWithResponse(options, null, Context.NONE);
    }

    private BlobRequestConditions concurrencyConditions(BlobClient blobClient) {
        if (!blobClient.exists()) {
            return new BlobRequestConditions().setIfNoneMatch("*");
        }
        return new BlobRequestConditions().setIfMatch(blobClient.getProperties().getETag());
    }

    private BlobClient blobClient(String containerName, String blobName) {
        return serviceClient.getBlobContainerClient(containerName).getBlobClient(blobName);
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

package com.example.blobevents;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;

import java.io.OutputStream;

public final class AzureBlobAccess implements BlobAccess {
    private final BlobServiceClient blobServiceClient;

    public AzureBlobAccess(BlobServiceClient blobServiceClient) {
        this.blobServiceClient = blobServiceClient;
    }

    @Override
    public BlobSummary download(String container, String blobName) {
        BlobClient client = blobServiceClient.getBlobContainerClient(container).getBlobClient(blobName);
        client.downloadStream(OutputStream.nullOutputStream());
        BlobProperties properties = client.getProperties();
        return new BlobSummary(
                blobName,
                properties.getBlobSize(),
                properties.getContentType(),
                properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString());
    }
}

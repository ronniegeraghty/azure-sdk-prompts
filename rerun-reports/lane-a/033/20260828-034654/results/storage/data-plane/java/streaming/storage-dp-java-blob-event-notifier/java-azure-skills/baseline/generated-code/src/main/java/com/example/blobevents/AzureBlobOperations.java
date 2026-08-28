package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import reactor.core.publisher.Mono;

import static com.example.blobevents.BlobEventHandler.BlobAddress;
import static com.example.blobevents.BlobEventHandler.BlobSummary;
import static com.example.blobevents.BlobEventHandler.BlobUnavailableException;

public final class AzureBlobOperations
        implements BlobEventHandler.BlobOperations, BlobEventHandler.AsyncBlobOperations {

    private final BlobServiceClient syncClient;
    private final BlobServiceAsyncClient asyncClient;

    public AzureBlobOperations(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
    }

    @Override
    public BlobSummary download(BlobAddress address) {
        BlobClient blob = syncClient.getBlobContainerClient(address.container()).getBlobClient(address.name());
        try {
            BlobProperties properties = blob.getProperties();
            BinaryData content = blob.downloadContent();
            return toSummary(address, properties, content.getLength());
        } catch (BlobStorageException exception) {
            throw translate(exception);
        }
    }

    @Override
    public Mono<BlobSummary> downloadAsync(BlobAddress address) {
        BlobAsyncClient blob = asyncClient.getBlobContainerAsyncClient(address.container())
                .getBlobAsyncClient(address.name());
        return blob.getProperties()
                .flatMap(properties -> blob.downloadContent()
                        .map(content -> toSummary(address, properties, content.getLength())))
                .onErrorMap(BlobStorageException.class, this::translate);
    }

    private BlobSummary toSummary(BlobAddress address, BlobProperties properties, long downloadedSize) {
        String tier = properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString();
        return new BlobSummary(address.name(), downloadedSize, properties.getContentType(), tier);
    }

    private BlobUnavailableException translate(BlobStorageException exception) {
        String errorCode = exception.getErrorCode() == null ? "unknown" : exception.getErrorCode().toString();
        if (exception.getStatusCode() == 404) {
            return new BlobUnavailableException("the blob was deleted or moved (HTTP 404)", exception);
        }
        if ("BlobArchived".equalsIgnoreCase(errorCode)) {
            return new BlobUnavailableException("the blob is in the archive tier", exception);
        }
        throw exception;
    }
}

package com.example.blobevents;

import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public class AsyncBlobEventHandler {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncBlobEventHandler.class);
    private final BlobServiceAsyncClient blobServiceClient;

    public AsyncBlobEventHandler(BlobServiceAsyncClient blobServiceClient) {
        this.blobServiceClient = blobServiceClient;
    }

    public Mono<Void> onBlobCreated(LifecycleEvent event) {
        BlobEventHandler.BlobAddress address = BlobEventHandler.BlobAddress.fromSubject(event.subject());
        var client = blobServiceClient
                .getBlobContainerAsyncClient(address.container())
                .getBlobAsyncClient(address.name());

        Mono<BlobProperties> properties = client.getProperties();
        return properties.flatMap(blobProperties -> client.downloadContent()
                        .doOnNext(ignored -> LOGGER.info(
                                "Blob created: name='{}', size={}, contentType='{}', accessTier='{}'",
                                address.name(), blobProperties.getBlobSize(), blobProperties.getContentType(),
                                blobProperties.getAccessTier()))
                        .then())
                .onErrorResume(BlobStorageException.class, exception -> handleReadRace(address, exception));
    }

    public Mono<Void> onBlobDeleted(LifecycleEvent event) {
        BlobEventHandler.BlobAddress address = BlobEventHandler.BlobAddress.fromSubject(event.subject());
        LOGGER.info("Blob deleted: container='{}', name='{}'", address.container(), address.name());
        return Mono.empty();
    }

    private static Mono<Void> handleReadRace(
            BlobEventHandler.BlobAddress address, BlobStorageException exception) {
        int status = exception.getStatusCode();
        String errorCode = exception.getErrorCode() == null ? "" : exception.getErrorCode().toString();
        if (status == 404) {
            LOGGER.warn("Blob '{}/{}' no longer exists; it was likely deleted or moved before processing",
                    address.container(), address.name());
            return Mono.empty();
        }
        if (status == 409 && ("BlobArchived".equals(errorCode) || "BlobBeingRehydrated".equals(errorCode))) {
            LOGGER.warn("Blob '{}/{}' cannot currently be downloaded because its access tier is {}",
                    address.container(), address.name(), errorCode);
            return Mono.empty();
        }
        return Mono.error(exception);
    }
}

package com.example.blobevents;

import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.models.BlobStorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public final class AsyncBlobEventHandler {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncBlobEventHandler.class);
    private final AsyncBlobDownloader downloader;

    public AsyncBlobEventHandler(BlobServiceAsyncClient blobServiceClient) {
        this((container, blobName) -> {
            BlobAsyncClient blob = blobServiceClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(blobName);
            return blob.getProperties()
                .flatMap(properties -> blob.downloadContent()
                    .thenReturn(new BlobSummary(
                        blobName,
                        properties.getBlobSize(),
                        properties.getContentType(),
                        properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString())));
        });
    }

    public AsyncBlobEventHandler(AsyncBlobDownloader downloader) {
        this.downloader = downloader;
    }

    public Mono<Void> handle(BlobLifecycleEvent event) {
        if (BlobEventHandler.BLOB_DELETED.equals(event.type())) {
            return Mono.fromRunnable(() ->
                LOGGER.info("Blob deleted: subject={}, eventId={}", event.subject(), event.id()));
        }
        if (!BlobEventHandler.BLOB_CREATED.equals(event.type())) {
            return Mono.error(new IllegalArgumentException("Unsupported event type: " + event.type()));
        }

        BlobSubject subject = BlobSubject.parse(event.subject());
        return downloader.download(subject.container(), subject.blobName())
            .doOnNext(summary -> LOGGER.info(
                "Blob created and downloaded: name={}, size={} bytes, contentType={}, accessTier={}",
                summary.name(), summary.size(), summary.contentType(), summary.accessTier()))
            .onErrorResume(BlobStorageException.class, exception -> {
                if (exception.getStatusCode() == 404 || exception.getStatusCode() == 409) {
                    LOGGER.warn(
                        "Blob is no longer readable after event delivery: container={}, blob={}, status={}, errorCode={}",
                        subject.container(), subject.blobName(), exception.getStatusCode(), exception.getErrorCode());
                    return Mono.empty();
                }
                return Mono.error(exception);
            })
            .then();
    }

    @FunctionalInterface
    public interface AsyncBlobDownloader {
        Mono<BlobSummary> download(String container, String blobName);
    }
}

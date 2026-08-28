package com.example.blobevents.blob;

import com.azure.storage.blob.models.BlobStorageException;
import com.example.blobevents.model.BlobLifecycleEvent;
import reactor.core.publisher.Mono;

import java.util.logging.Logger;

public final class AsyncBlobEventHandler {
    private static final Logger LOGGER = Logger.getLogger(AsyncBlobEventHandler.class.getName());

    private final BlobOperations blobs;

    public AsyncBlobEventHandler(BlobOperations blobs) {
        this.blobs = blobs;
    }

    public Mono<Void> handleCreated(BlobLifecycleEvent event) {
        BlobEventHandler.BlobLocation location = BlobEventHandler.parseSubject(event.subject());
        return blobs.downloadAsync(location.container(), location.name())
            .doOnNext(download -> {
                BlobSummary summary = download.summary();
                LOGGER.info(() -> "Blob created (async): name=%s, size=%d, contentType=%s, accessTier=%s"
                    .formatted(summary.name(), summary.size(), summary.contentType(), summary.accessTier()));
            })
            .then()
            .onErrorResume(BlobStorageException.class, exception -> {
                if (!BlobEventHandler.isLifecycleRace(exception)) {
                    return Mono.error(exception);
                }
                LOGGER.warning(() -> "Blob is no longer readable after creation event: "
                    + location.container() + "/" + location.name() + " (" + exception.getStatusCode() + ")");
                return Mono.empty();
            });
    }

    public Mono<Void> handleDeleted(BlobLifecycleEvent event) {
        BlobEventHandler.BlobLocation location = BlobEventHandler.parseSubject(event.subject());
        return Mono.fromRunnable(() -> LOGGER.info(
            () -> "Blob deleted (async): " + location.container() + "/" + location.name()));
    }
}

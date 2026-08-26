package com.example.blobevents;

import com.azure.storage.blob.models.BlobStorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public final class AsyncBlobEventHandler implements AsyncBlobEventProcessor {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncBlobEventHandler.class);

    private final AsyncBlobAccess blobAccess;

    public AsyncBlobEventHandler(AsyncBlobAccess blobAccess) {
        this.blobAccess = blobAccess;
    }

    @Override
    public Mono<Void> onBlobCreated(BlobLifecycleEvent event) {
        return Mono.fromCallable(() -> BlobLocation.fromSubject(event.subject()))
                .flatMap(location -> blobAccess.download(location.container(), location.blobName())
                        .doOnNext(summary -> LOGGER.info(
                                "Blob downloaded: name={}, size={} bytes, contentType={}, accessTier={}",
                                summary.name(), summary.size(), summary.contentType(), summary.accessTier()))
                        .then()
                        .onErrorResume(BlobStorageException.class, error -> {
                            if (!BlobEventHandler.isExpectedRace(error)) {
                                return Mono.error(error);
                            }
                            LOGGER.warn("Blob '{}' is no longer readable (status {}, error {}); it may have been "
                                            + "deleted, renamed, or moved to an offline tier",
                                    location.blobName(), error.getStatusCode(), error.getErrorCode());
                            return Mono.empty();
                        }));
    }

    @Override
    public Mono<Void> onBlobDeleted(BlobLifecycleEvent event) {
        return Mono.fromRunnable(() -> {
            BlobLocation location = BlobLocation.fromSubject(event.subject());
            LOGGER.info("Blob deleted: container={}, name={}", location.container(), location.blobName());
        });
    }
}

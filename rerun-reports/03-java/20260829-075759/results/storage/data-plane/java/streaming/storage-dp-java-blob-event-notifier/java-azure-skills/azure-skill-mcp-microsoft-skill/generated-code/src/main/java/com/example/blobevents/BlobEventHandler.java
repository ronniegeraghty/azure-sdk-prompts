package com.example.blobevents;

import com.azure.storage.blob.models.BlobStorageException;
import com.example.blobevents.model.BlobLocation;
import com.example.blobevents.model.BlobSummary;
import com.example.blobevents.model.IncomingEvent;
import com.example.blobevents.storage.BlobStore;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public final class BlobEventHandler {
    public static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    public static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);
    private final BlobStore blobStore;

    public BlobEventHandler(BlobStore blobStore) {
        this.blobStore = blobStore;
    }

    public void handleCreated(IncomingEvent event) {
        BlobLocation location = BlobLocation.fromSubject(event.subject());
        try {
            printSummary(location, blobStore.download(location.container(), location.blobName()));
        } catch (BlobStorageException exception) {
            if (isLifecycleRace(exception)) {
                LOGGER.warn("Blob {}/{} is no longer readable; lifecycle processing likely changed it",
                    location.container(), location.blobName());
                return;
            }
            throw exception;
        }
    }

    public Mono<Void> handleCreatedAsync(IncomingEvent event) {
        BlobLocation location = BlobLocation.fromSubject(event.subject());
        return blobStore.downloadAsync(location.container(), location.blobName())
            .doOnNext(summary -> printSummary(location, summary))
            .then()
            .onErrorResume(BlobStorageException.class, exception -> {
                if (isLifecycleRace(exception)) {
                    LOGGER.warn("Blob {}/{} is no longer readable; lifecycle processing likely changed it",
                        location.container(), location.blobName());
                    return Mono.empty();
                }
                return Mono.error(exception);
            });
    }

    public void handleDeleted(IncomingEvent event) {
        BlobLocation location = BlobLocation.fromSubject(event.subject());
        LOGGER.info("Blob deleted: container={}, name={}", location.container(), location.blobName());
    }

    public Mono<Void> handleDeletedAsync(IncomingEvent event) {
        return Mono.fromRunnable(() -> handleDeleted(event));
    }

    private static boolean isLifecycleRace(BlobStorageException exception) {
        return exception.getStatusCode() == 404 || exception.getStatusCode() == 409
            || exception.getResponse() == null;
    }

    private static void printSummary(BlobLocation location, BlobSummary summary) {
        LOGGER.info("Blob created: container={}, name={}, size={} bytes, contentType={}, accessTier={}",
            location.container(), summary.name(), summary.size(), summary.contentType(), summary.accessTier());
    }
}

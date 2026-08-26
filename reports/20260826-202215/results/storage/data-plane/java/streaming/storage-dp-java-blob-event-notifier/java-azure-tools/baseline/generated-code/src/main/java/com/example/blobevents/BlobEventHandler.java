package com.example.blobevents;

import com.azure.storage.blob.models.BlobStorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public final class BlobEventHandler implements BlobEventProcessor {
    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);

    private final BlobAccess blobAccess;

    public BlobEventHandler(BlobAccess blobAccess) {
        this.blobAccess = blobAccess;
    }

    @Override
    public void onBlobCreated(BlobLifecycleEvent event) {
        BlobLocation location = BlobLocation.fromSubject(event.subject());
        try {
            BlobSummary summary = blobAccess.download(location.container(), location.blobName());
            LOGGER.info("Blob downloaded: name={}, size={} bytes, contentType={}, accessTier={}",
                    summary.name(), summary.size(), summary.contentType(), summary.accessTier());
        } catch (BlobStorageException e) {
            if (isExpectedRace(e)) {
                LOGGER.warn("Blob '{}' is no longer readable (status {}, error {}); it may have been deleted, "
                                + "renamed, or moved to an offline tier",
                        location.blobName(), e.getStatusCode(), e.getErrorCode());
                return;
            }
            throw e;
        }
    }

    @Override
    public void onBlobDeleted(BlobLifecycleEvent event) {
        BlobLocation location = BlobLocation.fromSubject(event.subject());
        LOGGER.info("Blob deleted: container={}, name={}", location.container(), location.blobName());
    }

    static boolean isExpectedRace(BlobStorageException error) {
        return error.getStatusCode() == 404
                || error.getStatusCode() == 409
                || error.getStatusCode() == 410;
    }
}

package com.example.blobevents;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.OutputStream;

public final class BlobEventHandler {
    public static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    public static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);
    private final BlobDownloader downloader;

    public BlobEventHandler(BlobServiceClient blobServiceClient) {
        this((container, blobName) -> {
            BlobClient blob = blobServiceClient.getBlobContainerClient(container).getBlobClient(blobName);
            BlobProperties properties = blob.getProperties();
            blob.downloadStream(OutputStream.nullOutputStream());
            return new BlobSummary(
                blobName,
                properties.getBlobSize(),
                properties.getContentType(),
                properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString());
        });
    }

    public BlobEventHandler(BlobDownloader downloader) {
        this.downloader = downloader;
    }

    public void handle(BlobLifecycleEvent event) {
        if (BLOB_CREATED.equals(event.type())) {
            handleCreated(event);
        } else if (BLOB_DELETED.equals(event.type())) {
            LOGGER.info("Blob deleted: subject={}, eventId={}", event.subject(), event.id());
        } else {
            throw new IllegalArgumentException("Unsupported event type: " + event.type());
        }
    }

    private void handleCreated(BlobLifecycleEvent event) {
        BlobSubject subject = BlobSubject.parse(event.subject());
        try {
            BlobSummary summary = downloader.download(subject.container(), subject.blobName());
            LOGGER.info(
                "Blob created and downloaded: name={}, size={} bytes, contentType={}, accessTier={}",
                summary.name(), summary.size(), summary.contentType(), summary.accessTier());
        } catch (BlobStorageException exception) {
            if (exception.getStatusCode() == 404 || exception.getStatusCode() == 409) {
                LOGGER.warn(
                    "Blob is no longer readable after event delivery: container={}, blob={}, status={}, errorCode={}",
                    subject.container(), subject.blobName(), exception.getStatusCode(), exception.getErrorCode());
                return;
            }
            throw exception;
        }
    }

    @FunctionalInterface
    public interface BlobDownloader {
        BlobSummary download(String container, String blobName);
    }
}

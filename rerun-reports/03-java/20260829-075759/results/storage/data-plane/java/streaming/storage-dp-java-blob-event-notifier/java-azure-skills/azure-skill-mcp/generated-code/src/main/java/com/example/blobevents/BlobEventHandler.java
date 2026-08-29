package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import reactor.core.publisher.Mono;

import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;
import java.util.Objects;
import java.util.logging.Logger;

public final class BlobEventHandler {
    private static final Logger LOGGER = Logger.getLogger(BlobEventHandler.class.getName());
    private static final String CONTAINERS_SEGMENT = "/containers/";
    private static final String BLOBS_SEGMENT = "/blobs/";

    private final SyncBlobReader syncReader;
    private final AsyncBlobReader asyncReader;

    public BlobEventHandler(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {
        this(
            location -> readBlob(syncClient, location),
            location -> readBlobAsync(asyncClient, location));
    }

    BlobEventHandler(SyncBlobReader syncReader, AsyncBlobReader asyncReader) {
        this.syncReader = Objects.requireNonNull(syncReader, "syncReader");
        this.asyncReader = Objects.requireNonNull(asyncReader, "asyncReader");
    }

    public void handleCreated(BlobLifecycleEvent event) {
        BlobLocation location = parseSubject(event.subject());
        try {
            logSummary(syncReader.read(location));
        } catch (BlobStorageException exception) {
            if (!handleExpectedRace(location, exception)) {
                throw exception;
            }
        }
    }

    public Mono<Void> handleCreatedAsync(BlobLifecycleEvent event) {
        BlobLocation location = parseSubject(event.subject());
        return asyncReader.read(location)
            .doOnNext(BlobEventHandler::logSummary)
            .then()
            .onErrorResume(BlobStorageException.class, exception -> {
                if (handleExpectedRace(location, exception)) {
                    return Mono.empty();
                }
                return Mono.error(exception);
            });
    }

    public void handleDeleted(BlobLifecycleEvent event) {
        BlobLocation location = parseSubject(event.subject());
        LOGGER.info(() -> "Blob deleted: container=" + location.container() + ", name=" + location.name());
    }

    public Mono<Void> handleDeletedAsync(BlobLifecycleEvent event) {
        handleDeleted(event);
        return Mono.empty();
    }

    static BlobLocation parseSubject(String subject) {
        int containerStart = subject.indexOf(CONTAINERS_SEGMENT);
        int blobSeparator = subject.indexOf(BLOBS_SEGMENT, containerStart + CONTAINERS_SEGMENT.length());
        if (containerStart < 0 || blobSeparator < 0) {
            throw new IllegalArgumentException("Unexpected blob event subject: " + subject);
        }

        String container = subject.substring(containerStart + CONTAINERS_SEGMENT.length(), blobSeparator);
        String name = subject.substring(blobSeparator + BLOBS_SEGMENT.length());
        if (container.isBlank() || name.isBlank()) {
            throw new IllegalArgumentException("Blob event subject must contain a container and blob name: " + subject);
        }
        return new BlobLocation(decodePathPart(container), decodePathPart(name));
    }

    private static BlobSummary readBlob(BlobServiceClient serviceClient, BlobLocation location) {
        BlobClient blob = serviceClient.getBlobContainerClient(location.container()).getBlobClient(location.name());
        blob.downloadContent();
        BlobProperties properties = blob.getProperties();
        return toSummary(location, properties);
    }

    private static Mono<BlobSummary> readBlobAsync(
        BlobServiceAsyncClient serviceClient,
        BlobLocation location
    ) {
        BlobAsyncClient blob = serviceClient.getBlobContainerAsyncClient(location.container())
            .getBlobAsyncClient(location.name());
        return blob.downloadContent()
            .then(blob.getProperties())
            .map(properties -> toSummary(location, properties));
    }

    private static BlobSummary toSummary(BlobLocation location, BlobProperties properties) {
        String accessTier = properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString();
        return new BlobSummary(
            location.name(),
            properties.getBlobSize(),
            properties.getContentType(),
            accessTier);
    }

    private static void logSummary(BlobSummary summary) {
        LOGGER.info(() -> "Blob created: name=" + summary.name()
            + ", size=" + summary.size()
            + ", contentType=" + summary.contentType()
            + ", accessTier=" + summary.accessTier());
    }

    private static boolean handleExpectedRace(BlobLocation location, BlobStorageException exception) {
        String errorCode = exception.getErrorCode() == null ? "" : exception.getErrorCode().toString();
        if (exception.getStatusCode() == 404) {
            LOGGER.warning(() -> "Blob disappeared before it could be read: " + location);
            return true;
        }
        if (exception.getStatusCode() == 409
            && (errorCode.contains("Archive") || errorCode.contains("Rehydr"))) {
            LOGGER.warning(() -> "Blob is temporarily unreadable in its current access tier: " + location
                + " (" + errorCode + ")");
            return true;
        }
        return false;
    }

    private static String decodePathPart(String value) {
        return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
    }

    record BlobLocation(String container, String name) {
    }

    record BlobSummary(String name, long size, String contentType, String accessTier) {
    }

    @FunctionalInterface
    interface SyncBlobReader {
        BlobSummary read(BlobLocation location);
    }

    @FunctionalInterface
    interface AsyncBlobReader {
        Mono<BlobSummary> read(BlobLocation location);
    }
}

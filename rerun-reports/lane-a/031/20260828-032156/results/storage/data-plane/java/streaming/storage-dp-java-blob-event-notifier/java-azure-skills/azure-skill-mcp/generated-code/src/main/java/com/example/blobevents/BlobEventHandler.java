package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobErrorCode;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import java.net.URI;
import java.util.Map;
import java.util.Objects;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public final class BlobEventHandler {
    public static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    public static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);
    private static final String CONTAINER_MARKER = "/containers/";
    private static final String BLOB_MARKER = "/blobs/";

    private final BlobServiceClient blobServiceClient;
    private final BlobServiceAsyncClient blobServiceAsyncClient;
    private final Map<BlobAddress, BlobSummary> demoBlobs;

    public BlobEventHandler(
        BlobServiceClient blobServiceClient,
        BlobServiceAsyncClient blobServiceAsyncClient
    ) {
        this.blobServiceClient = Objects.requireNonNull(blobServiceClient, "blobServiceClient");
        this.blobServiceAsyncClient = Objects.requireNonNull(blobServiceAsyncClient, "blobServiceAsyncClient");
        this.demoBlobs = null;
    }

    public static BlobEventHandler forDemo(Map<BlobAddress, BlobSummary> demoBlobs) {
        return new BlobEventHandler(Map.copyOf(demoBlobs));
    }

    private BlobEventHandler(Map<BlobAddress, BlobSummary> demoBlobs) {
        this.blobServiceClient = null;
        this.blobServiceAsyncClient = null;
        this.demoBlobs = demoBlobs;
    }

    public void handle(BlobLifecycleEvent event) {
        switch (event.eventType()) {
            case BLOB_CREATED -> handleCreated(event);
            case BLOB_DELETED -> logDeletion(event);
            default -> LOGGER.warn("Ignoring unrecognized event type '{}' for event {}", event.eventType(), event.id());
        }
    }

    public Mono<Void> handleAsync(BlobLifecycleEvent event) {
        return switch (event.eventType()) {
            case BLOB_CREATED -> handleCreatedAsync(event);
            case BLOB_DELETED -> Mono.fromRunnable(() -> logDeletion(event));
            default -> Mono.fromRunnable(
                () -> LOGGER.warn("Ignoring unrecognized event type '{}' for event {}", event.eventType(), event.id())
            );
        };
    }

    private void handleCreated(BlobLifecycleEvent event) {
        BlobAddress address = parseAddress(event.subject());
        if (demoBlobs != null) {
            printSummary(address, requireDemoBlob(address));
            return;
        }

        BlobClient blobClient = blobServiceClient
            .getBlobContainerClient(address.container())
            .getBlobClient(address.blobName());
        try {
            BlobProperties properties = blobClient.getProperties();
            BinaryData ignoredContent = blobClient.downloadContent();
            printSummary(address, fromProperties(properties));
        } catch (BlobStorageException exception) {
            handleStorageRace(address, exception);
        }
    }

    private Mono<Void> handleCreatedAsync(BlobLifecycleEvent event) {
        BlobAddress address = parseAddress(event.subject());
        if (demoBlobs != null) {
            return Mono.fromRunnable(() -> printSummary(address, requireDemoBlob(address)));
        }

        return blobServiceAsyncClient
            .getBlobContainerAsyncClient(address.container())
            .getBlobAsyncClient(address.blobName())
            .getProperties()
            .flatMap(properties -> blobServiceAsyncClient
                .getBlobContainerAsyncClient(address.container())
                .getBlobAsyncClient(address.blobName())
                .downloadContent()
                .doOnNext(ignored -> printSummary(address, fromProperties(properties))))
            .onErrorResume(BlobStorageException.class, exception -> {
                handleStorageRace(address, exception);
                return Mono.empty();
            })
            .then();
    }

    private static void logDeletion(BlobLifecycleEvent event) {
        BlobAddress address = parseAddress(event.subject());
        LOGGER.info("Blob deleted: container='{}', name='{}'", address.container(), address.blobName());
    }

    private static BlobSummary fromProperties(BlobProperties properties) {
        String contentType = properties.getContentType() == null ? "application/octet-stream" : properties.getContentType();
        String accessTier = properties.getAccessTier() == null ? "unknown" : properties.getAccessTier().toString();
        return new BlobSummary(properties.getBlobSize(), contentType, accessTier);
    }

    private static void printSummary(BlobAddress address, BlobSummary summary) {
        LOGGER.info(
            "Blob created: name='{}', container='{}', size={} bytes, contentType='{}', accessTier='{}'",
            address.blobName(),
            address.container(),
            summary.size(),
            summary.contentType(),
            summary.accessTier()
        );
    }

    private BlobSummary requireDemoBlob(BlobAddress address) {
        BlobSummary summary = demoBlobs.get(address);
        if (summary == null) {
            LOGGER.warn("Blob disappeared before it could be downloaded: {}/{}", address.container(), address.blobName());
            return new BlobSummary(0, "unknown", "unknown");
        }
        return summary;
    }

    private static void handleStorageRace(BlobAddress address, BlobStorageException exception) {
        if (exception.getStatusCode() == 404) {
            LOGGER.warn("Blob disappeared before it could be downloaded: {}/{}", address.container(), address.blobName());
            return;
        }
        if (BlobErrorCode.BLOB_ARCHIVED.equals(exception.getErrorCode())
            || BlobErrorCode.BLOB_BEING_REHYDRATED.equals(exception.getErrorCode())) {
            LOGGER.warn(
                "Blob is currently unavailable, possibly because its access tier changed: {}/{} ({})",
                address.container(),
                address.blobName(),
                exception.getErrorCode()
            );
            return;
        }
        throw exception;
    }

    public static BlobAddress parseAddress(String subject) {
        int containerStart = subject.indexOf(CONTAINER_MARKER);
        int blobStart = subject.indexOf(BLOB_MARKER, containerStart + CONTAINER_MARKER.length());
        if (containerStart < 0 || blobStart < 0) {
            throw new IllegalArgumentException("Blob event subject has an unexpected format: " + subject);
        }
        String container = subject.substring(containerStart + CONTAINER_MARKER.length(), blobStart);
        String encodedBlobName = subject.substring(blobStart + BLOB_MARKER.length());
        String blobName = URI.create(encodedBlobName).getPath();
        if (container.isBlank() || blobName.isBlank()) {
            throw new IllegalArgumentException("Blob event subject must identify a container and blob: " + subject);
        }
        return new BlobAddress(container, blobName);
    }

    public record BlobAddress(String container, String blobName) {
    }

    public record BlobSummary(long size, String contentType, String accessTier) {
    }
}

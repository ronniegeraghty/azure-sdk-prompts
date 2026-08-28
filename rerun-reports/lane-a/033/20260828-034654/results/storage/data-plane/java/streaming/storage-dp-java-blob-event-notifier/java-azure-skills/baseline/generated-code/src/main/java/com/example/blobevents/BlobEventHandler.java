package com.example.blobevents;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;

public final class BlobEventHandler {
    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);
    private static final String SUBJECT_PREFIX = "/blobServices/default/containers/";
    private static final String BLOB_SEPARATOR = "/blobs/";

    private final BlobOperations blobOperations;
    private final AsyncBlobOperations asyncBlobOperations;

    public BlobEventHandler(BlobOperations blobOperations, AsyncBlobOperations asyncBlobOperations) {
        this.blobOperations = blobOperations;
        this.asyncBlobOperations = asyncBlobOperations;
    }

    public void handleCreated(EventEnvelope event) {
        BlobAddress address = parseSubject(event.subject());
        try {
            logSummary(blobOperations.download(address));
        } catch (BlobUnavailableException exception) {
            LOGGER.warn("Blob {} in container {} is no longer readable: {}",
                    address.name(), address.container(), exception.getMessage());
        }
    }

    public Mono<Void> handleCreatedAsync(EventEnvelope event) {
        BlobAddress address = parseSubject(event.subject());
        return asyncBlobOperations.downloadAsync(address)
                .doOnNext(this::logSummary)
                .onErrorResume(BlobUnavailableException.class, exception -> {
                    LOGGER.warn("Blob {} in container {} is no longer readable: {}",
                            address.name(), address.container(), exception.getMessage());
                    return Mono.empty();
                })
                .then();
    }

    public void handleDeleted(EventEnvelope event) {
        BlobAddress address = parseSubject(event.subject());
        LOGGER.info("Blob deleted: container={}, name={}", address.container(), address.name());
    }

    public Mono<Void> handleDeletedAsync(EventEnvelope event) {
        handleDeleted(event);
        return Mono.empty();
    }

    static BlobAddress parseSubject(String subject) {
        if (subject == null || !subject.startsWith(SUBJECT_PREFIX)) {
            throw new IllegalArgumentException("Unexpected blob event subject: " + subject);
        }

        int blobSeparator = subject.indexOf(BLOB_SEPARATOR, SUBJECT_PREFIX.length());
        if (blobSeparator < 0) {
            throw new IllegalArgumentException("Blob event subject has no blob name: " + subject);
        }

        String container = subject.substring(SUBJECT_PREFIX.length(), blobSeparator);
        String name = subject.substring(blobSeparator + BLOB_SEPARATOR.length());
        if (container.isBlank() || name.isBlank()) {
            throw new IllegalArgumentException("Blob event subject has an empty container or blob name: " + subject);
        }
        return new BlobAddress(decode(container), decode(name));
    }

    private static String decode(String value) {
        return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
    }

    private void logSummary(BlobSummary summary) {
        LOGGER.info("Blob created: name={}, size={} bytes, contentType={}, accessTier={}",
                summary.name(), summary.size(), summary.contentType(), summary.accessTier());
    }

    @FunctionalInterface
    public interface BlobOperations {
        BlobSummary download(BlobAddress address);
    }

    @FunctionalInterface
    public interface AsyncBlobOperations {
        Mono<BlobSummary> downloadAsync(BlobAddress address);
    }

    public record BlobAddress(String container, String name) {
    }

    public record BlobSummary(String name, long size, String contentType, String accessTier) {
    }

    public static final class BlobUnavailableException extends RuntimeException {
        public BlobUnavailableException(String message, Throwable cause) {
            super(message, cause);
        }
    }
}

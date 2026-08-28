package com.example.blobevents.blob;

import com.azure.storage.blob.models.BlobStorageException;
import com.example.blobevents.model.BlobLifecycleEvent;

import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public final class BlobEventHandler {
    private static final Logger LOGGER = Logger.getLogger(BlobEventHandler.class.getName());
    private static final Pattern SUBJECT_PATTERN =
        Pattern.compile("^/blobServices/default/containers/([^/]+)/blobs/(.+)$");

    private final BlobOperations blobs;

    public BlobEventHandler(BlobOperations blobs) {
        this.blobs = blobs;
    }

    public void handleCreated(BlobLifecycleEvent event) {
        BlobLocation location = parseSubject(event.subject());
        try {
            BlobSummary summary = blobs.download(location.container(), location.name()).summary();
            LOGGER.info(() -> "Blob created: name=%s, size=%d, contentType=%s, accessTier=%s"
                .formatted(summary.name(), summary.size(), summary.contentType(), summary.accessTier()));
        } catch (BlobStorageException exception) {
            if (isLifecycleRace(exception)) {
                LOGGER.warning(() -> "Blob is no longer readable after creation event: "
                    + location.container() + "/" + location.name() + " (" + exception.getStatusCode() + ")");
                return;
            }
            throw exception;
        }
    }

    public void handleDeleted(BlobLifecycleEvent event) {
        BlobLocation location = parseSubject(event.subject());
        LOGGER.info(() -> "Blob deleted: " + location.container() + "/" + location.name());
    }

    static BlobLocation parseSubject(String subject) {
        Matcher matcher = SUBJECT_PATTERN.matcher(subject);
        if (!matcher.matches()) {
            throw new IllegalArgumentException("Unexpected blob event subject: " + subject);
        }
        return new BlobLocation(decode(matcher.group(1)), decode(matcher.group(2)));
    }

    private static String decode(String value) {
        return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
    }

    static boolean isLifecycleRace(BlobStorageException exception) {
        return exception.getStatusCode() == 404 || exception.getStatusCode() == 409;
    }

    record BlobLocation(String container, String name) {
    }
}

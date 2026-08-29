package com.example.blobevents;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.ByteArrayOutputStream;
import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;

public class BlobEventHandler {
    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);
    private final BlobServiceClient blobServiceClient;

    public BlobEventHandler(BlobServiceClient blobServiceClient) {
        this.blobServiceClient = blobServiceClient;
    }

    public void onBlobCreated(LifecycleEvent event) {
        BlobAddress address = BlobAddress.fromSubject(event.subject());
        BlobClient client = blobServiceClient
                .getBlobContainerClient(address.container())
                .getBlobClient(address.name());

        try {
            BlobProperties properties = client.getProperties();
            try (ByteArrayOutputStream content = new ByteArrayOutputStream()) {
                client.downloadStream(content);
                LOGGER.info("Blob created: name='{}', size={}, contentType='{}', accessTier='{}'",
                        address.name(), properties.getBlobSize(), properties.getContentType(),
                        properties.getAccessTier());
            } catch (BlobStorageException exception) {
                handleReadRace(address, exception);
            } catch (java.io.IOException exception) {
                throw new IllegalStateException("Unable to close blob download stream", exception);
            }
        } catch (BlobStorageException exception) {
            handleReadRace(address, exception);
        }
    }

    public void onBlobDeleted(LifecycleEvent event) {
        BlobAddress address = BlobAddress.fromSubject(event.subject());
        LOGGER.info("Blob deleted: container='{}', name='{}'", address.container(), address.name());
    }

    private static void handleReadRace(BlobAddress address, BlobStorageException exception) {
        int status = exception.getStatusCode();
        String errorCode = exception.getErrorCode() == null ? "" : exception.getErrorCode().toString();
        if (status == 404) {
            LOGGER.warn("Blob '{}/{}' no longer exists; it was likely deleted or moved before processing",
                    address.container(), address.name());
        } else if (status == 409 && ("BlobArchived".equals(errorCode) || "BlobBeingRehydrated".equals(errorCode))) {
            LOGGER.warn("Blob '{}/{}' cannot currently be downloaded because its access tier is {}",
                    address.container(), address.name(), errorCode);
        } else {
            throw exception;
        }
    }

    protected record BlobAddress(String container, String name) {
        private static final String CONTAINERS = "/containers/";
        private static final String BLOBS = "/blobs/";

        static BlobAddress fromSubject(String subject) {
            int containerStart = subject.indexOf(CONTAINERS);
            int blobStart = subject.indexOf(BLOBS, containerStart + CONTAINERS.length());
            if (containerStart < 0 || blobStart < 0) {
                throw new IllegalArgumentException("Unexpected blob event subject: " + subject);
            }

            String container = subject.substring(containerStart + CONTAINERS.length(), blobStart);
            String name = subject.substring(blobStart + BLOBS.length());
            if (container.isBlank() || name.isBlank()) {
                throw new IllegalArgumentException("Blob event subject has an empty container or blob name: " + subject);
            }
            return new BlobAddress(decode(container), decode(name));
        }

        private static String decode(String value) {
            return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
        }
    }
}

package com.example.blobevents;

import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;

record BlobSubject(String container, String blobName) {
    private static final String CONTAINER_MARKER = "/containers/";
    private static final String BLOB_MARKER = "/blobs/";

    static BlobSubject parse(String subject) {
        int containerStart = subject.indexOf(CONTAINER_MARKER);
        int blobStart = subject.indexOf(BLOB_MARKER, containerStart + CONTAINER_MARKER.length());
        if (containerStart < 0 || blobStart < 0) {
            throw new IllegalArgumentException("Unsupported blob event subject: " + subject);
        }

        String container = subject.substring(containerStart + CONTAINER_MARKER.length(), blobStart);
        String blobName = subject.substring(blobStart + BLOB_MARKER.length());
        if (container.isBlank() || blobName.isBlank()) {
            throw new IllegalArgumentException("Blob event subject is missing a container or blob name: " + subject);
        }
        return new BlobSubject(decodePath(container), decodePath(blobName));
    }

    private static String decodePath(String value) {
        return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
    }
}

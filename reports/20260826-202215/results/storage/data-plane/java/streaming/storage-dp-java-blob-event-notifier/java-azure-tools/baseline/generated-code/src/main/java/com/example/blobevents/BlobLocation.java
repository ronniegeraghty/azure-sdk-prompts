package com.example.blobevents;

import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;

public record BlobLocation(String container, String blobName) {
    private static final String CONTAINERS_MARKER = "/containers/";
    private static final String BLOBS_MARKER = "/blobs/";

    public static BlobLocation fromSubject(String subject) {
        int containerStart = subject.indexOf(CONTAINERS_MARKER);
        int blobMarker = subject.indexOf(BLOBS_MARKER, containerStart + CONTAINERS_MARKER.length());
        if (containerStart < 0 || blobMarker < 0) {
            throw new IllegalArgumentException("Invalid blob event subject: " + subject);
        }

        String container = subject.substring(containerStart + CONTAINERS_MARKER.length(), blobMarker);
        String blobName = subject.substring(blobMarker + BLOBS_MARKER.length());
        if (container.isBlank() || blobName.isBlank()) {
            throw new IllegalArgumentException("Blob event subject has an empty container or blob name: " + subject);
        }
        return new BlobLocation(decode(container), decode(blobName));
    }

    private static String decode(String value) {
        return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
    }
}

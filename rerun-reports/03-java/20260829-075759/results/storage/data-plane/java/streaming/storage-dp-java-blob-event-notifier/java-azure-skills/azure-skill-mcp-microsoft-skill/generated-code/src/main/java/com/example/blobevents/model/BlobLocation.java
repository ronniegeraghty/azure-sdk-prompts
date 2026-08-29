package com.example.blobevents.model;

import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;

public record BlobLocation(String container, String blobName) {
    private static final String CONTAINERS = "/containers/";
    private static final String BLOBS = "/blobs/";

    public static BlobLocation fromSubject(String subject) {
        if (subject == null) {
            throw new IllegalArgumentException("Event subject is required");
        }

        int containerStart = subject.indexOf(CONTAINERS);
        int blobMarker = subject.indexOf(BLOBS, containerStart + CONTAINERS.length());
        if (containerStart < 0 || blobMarker < 0) {
            throw new IllegalArgumentException("Unsupported blob event subject: " + subject);
        }

        String container = subject.substring(containerStart + CONTAINERS.length(), blobMarker);
        String blobName = subject.substring(blobMarker + BLOBS.length());
        if (container.isBlank() || blobName.isBlank()) {
            throw new IllegalArgumentException("Blob event subject has an empty container or blob name: " + subject);
        }
        return new BlobLocation(decodePathValue(container), decodePathValue(blobName));
    }

    private static String decodePathValue(String value) {
        return URLDecoder.decode(value.replace("+", "%2B"), StandardCharsets.UTF_8);
    }
}

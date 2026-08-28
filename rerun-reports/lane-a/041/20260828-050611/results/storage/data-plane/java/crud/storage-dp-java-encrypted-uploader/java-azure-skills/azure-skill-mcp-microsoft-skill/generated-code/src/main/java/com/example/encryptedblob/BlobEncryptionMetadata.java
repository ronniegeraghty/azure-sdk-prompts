package com.example.encryptedblob;

import java.util.Base64;
import java.util.HashMap;
import java.util.Map;

record BlobEncryptionMetadata(ProtectedDataKey protectedKey, byte[] iv) {
    private static final String VERSION = "1";
    private static final String CONTENT_ALGORITHM = "AES256GCM";
    private static final String WRAP_ALGORITHM = "RSAOAEP256";

    static Map<String, String> create(ProtectedDataKey protectedKey, byte[] iv) {
        Map<String, String> metadata = new HashMap<>();
        metadata.put("encryptionversion", VERSION);
        metadata.put("contentalgorithm", CONTENT_ALGORITHM);
        metadata.put("wrapalgorithm", WRAP_ALGORITHM);
        metadata.put("keyid", protectedKey.keyId());
        metadata.put("wrappedkey", Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));
        metadata.put("iv", Base64.getEncoder().encodeToString(iv));
        return metadata;
    }

    static BlobEncryptionMetadata parse(Map<String, String> metadata) {
        requireValue(metadata, "encryptionversion", VERSION);
        requireValue(metadata, "contentalgorithm", CONTENT_ALGORITHM);
        requireValue(metadata, "wrapalgorithm", WRAP_ALGORITHM);

        String keyId = require(metadata, "keyid");
        try {
            byte[] wrappedKey = Base64.getDecoder().decode(require(metadata, "wrappedkey"));
            byte[] iv = Base64.getDecoder().decode(require(metadata, "iv"));
            if (iv.length != 12) {
                throw new EnvelopeEncryptionException("Invalid AES-GCM IV length in blob metadata.");
            }
            return new BlobEncryptionMetadata(new ProtectedDataKey(keyId, wrappedKey), iv);
        } catch (IllegalArgumentException exception) {
            throw new EnvelopeEncryptionException("Blob encryption metadata contains invalid Base64.", exception);
        }
    }

    private static String require(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new EnvelopeEncryptionException("Blob is missing required encryption metadata: " + name);
        }
        return value;
    }

    private static void requireValue(Map<String, String> metadata, String name, String expected) {
        String actual = require(metadata, name);
        if (!expected.equals(actual)) {
            throw new EnvelopeEncryptionException(
                "Unsupported blob encryption metadata " + name + ": " + actual);
        }
    }
}

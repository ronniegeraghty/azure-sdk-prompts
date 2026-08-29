package com.example.encryptedblob;

import java.util.Arrays;
import java.util.Base64;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

public record BlobEncryptionMetadata(
    String keyId,
    String wrapAlgorithm,
    byte[] wrappedDataKey,
    byte[] initializationVector
) {
    static final String ENCRYPTION_VERSION = "1";
    static final String CONTENT_ALGORITHM = "AES-256-GCM";

    private static final String VERSION = "encversion";
    private static final String CONTENT_ALG = "contentalg";
    private static final String WRAP_ALG = "wrapalg";
    private static final String KEY_ID = "keyid";
    private static final String WRAPPED_KEY = "wrappedkey";
    private static final String IV = "iv";

    public BlobEncryptionMetadata {
        Objects.requireNonNull(keyId, "keyId");
        Objects.requireNonNull(wrapAlgorithm, "wrapAlgorithm");
        wrappedDataKey = Arrays.copyOf(
            Objects.requireNonNull(wrappedDataKey, "wrappedDataKey"),
            wrappedDataKey.length);
        initializationVector = Arrays.copyOf(
            Objects.requireNonNull(initializationVector, "initializationVector"),
            initializationVector.length);
    }

    static BlobEncryptionMetadata fromMap(Map<String, String> metadata) {
        requireValue(metadata, VERSION, ENCRYPTION_VERSION);
        requireValue(metadata, CONTENT_ALG, CONTENT_ALGORITHM);
        try {
            return new BlobEncryptionMetadata(
                required(metadata, KEY_ID),
                required(metadata, WRAP_ALG),
                Base64.getDecoder().decode(required(metadata, WRAPPED_KEY)),
                Base64.getDecoder().decode(required(metadata, IV)));
        } catch (IllegalArgumentException exception) {
            throw new IllegalArgumentException("Blob encryption metadata is malformed", exception);
        }
    }

    Map<String, String> toMap() {
        Map<String, String> metadata = new HashMap<>();
        metadata.put(VERSION, ENCRYPTION_VERSION);
        metadata.put(CONTENT_ALG, CONTENT_ALGORITHM);
        metadata.put(WRAP_ALG, wrapAlgorithm);
        metadata.put(KEY_ID, keyId);
        metadata.put(WRAPPED_KEY, wrappedDataKeyBase64());
        metadata.put(IV, Base64.getEncoder().encodeToString(initializationVector));
        return metadata;
    }

    public String wrappedDataKeyBase64() {
        return Base64.getEncoder().encodeToString(wrappedDataKey);
    }

    @Override
    public byte[] wrappedDataKey() {
        return Arrays.copyOf(wrappedDataKey, wrappedDataKey.length);
    }

    @Override
    public byte[] initializationVector() {
        return Arrays.copyOf(initializationVector, initializationVector.length);
    }

    private static String required(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("Blob is missing encryption metadata: " + name);
        }
        return value;
    }

    private static void requireValue(Map<String, String> metadata, String name, String expected) {
        String actual = required(metadata, name);
        if (!expected.equals(actual)) {
            throw new IllegalArgumentException("Unsupported " + name + ": " + actual);
        }
    }
}

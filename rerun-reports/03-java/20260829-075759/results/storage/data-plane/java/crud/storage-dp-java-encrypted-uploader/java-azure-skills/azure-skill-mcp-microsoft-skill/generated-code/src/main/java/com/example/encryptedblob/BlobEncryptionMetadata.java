package com.example.encryptedblob;

import java.util.Base64;
import java.util.Map;

final class BlobEncryptionMetadata {
    private static final String CONTENT_ALGORITHM = "encryption-algorithm";
    private static final String WRAP_ALGORITHM = "key-wrap-algorithm";
    private static final String KEY_ID = "key-id";
    private static final String WRAPPED_KEY = "wrapped-key";
    private static final String IV = "iv";

    private final String keyId;
    private final byte[] wrappedKey;
    private final byte[] iv;

    private BlobEncryptionMetadata(String keyId, byte[] wrappedKey, byte[] iv) {
        this.keyId = keyId;
        this.wrappedKey = wrappedKey;
        this.iv = iv;
    }

    static BlobEncryptionMetadata create(String keyId, byte[] wrappedKey, byte[] iv) {
        return new BlobEncryptionMetadata(keyId, wrappedKey, iv);
    }

    static BlobEncryptionMetadata parse(Map<String, String> metadata) {
        String contentAlgorithm = required(metadata, CONTENT_ALGORITHM);
        String wrapAlgorithm = required(metadata, WRAP_ALGORITHM);
        if (!EnvelopeCrypto.CONTENT_ALGORITHM.equals(contentAlgorithm)) {
            throw new EncryptionStorageException("Unsupported content encryption algorithm: " + contentAlgorithm);
        }
        if (!KeyManagementService.WRAP_ALGORITHM_NAME.equals(wrapAlgorithm)) {
            throw new EncryptionStorageException("Unsupported key wrap algorithm: " + wrapAlgorithm);
        }

        try {
            return new BlobEncryptionMetadata(
                    required(metadata, KEY_ID),
                    Base64.getDecoder().decode(required(metadata, WRAPPED_KEY)),
                    Base64.getDecoder().decode(required(metadata, IV)));
        } catch (IllegalArgumentException exception) {
            throw new EncryptionStorageException("Blob encryption metadata contains invalid base64", exception);
        }
    }

    Map<String, String> toMap() {
        Base64.Encoder encoder = Base64.getEncoder();
        return Map.of(
                CONTENT_ALGORITHM, EnvelopeCrypto.CONTENT_ALGORITHM,
                WRAP_ALGORITHM, KeyManagementService.WRAP_ALGORITHM_NAME,
                KEY_ID, keyId,
                WRAPPED_KEY, encoder.encodeToString(wrappedKey),
                IV, encoder.encodeToString(iv));
    }

    String keyId() {
        return keyId;
    }

    byte[] wrappedKey() {
        return wrappedKey;
    }

    byte[] iv() {
        return iv;
    }

    byte[] authenticatedMetadata() {
        return EnvelopeCrypto.authenticatedMetadata(keyId, wrappedKey);
    }

    private static String required(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new EncryptionStorageException("Blob is missing encryption metadata field: " + name);
        }
        return value;
    }
}

package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.Map;
import java.util.Objects;

final class EncryptionMetadata {
    private static final String FORMAT_VERSION = "1";
    private static final String VERSION = "encversion";
    private static final String CONTENT_ALGORITHM = "encalgorithm";
    private static final String IV = "enciv";
    private static final String KEY_ID = "enckeyid";
    private static final String WRAP_ALGORITHM = "encwrapalgorithm";
    private static final String WRAPPED_KEY = "encwrappedkey";

    private final ProtectedDataKey protectedDataKey;
    private final byte[] iv;

    private EncryptionMetadata(ProtectedDataKey protectedDataKey, byte[] iv) {
        this.protectedDataKey = Objects.requireNonNull(protectedDataKey, "protectedDataKey");
        this.iv = Objects.requireNonNull(iv, "iv").clone();
    }

    static EncryptionMetadata create(ProtectedDataKey protectedDataKey, byte[] iv) {
        return new EncryptionMetadata(protectedDataKey, iv);
    }

    static EncryptionMetadata parse(Map<String, String> metadata) {
        try {
            String version = required(metadata, VERSION);
            if (!FORMAT_VERSION.equals(version)) {
                throw new BlobEncryptionException(
                    "Unsupported encryption metadata version: " + version
                );
            }

            String contentAlgorithm = required(metadata, CONTENT_ALGORITHM);
            if (!LocalAesGcm.ALGORITHM.equals(contentAlgorithm)) {
                throw new BlobEncryptionException(
                    "Unsupported content encryption algorithm: " + contentAlgorithm
                );
            }

            ProtectedDataKey protectedKey = new ProtectedDataKey(
                required(metadata, KEY_ID),
                required(metadata, WRAP_ALGORITHM),
                Base64.getDecoder().decode(required(metadata, WRAPPED_KEY))
            );
            byte[] iv = Base64.getDecoder().decode(required(metadata, IV));
            return new EncryptionMetadata(protectedKey, iv);
        } catch (IllegalArgumentException e) {
            throw new BlobEncryptionException("Blob encryption metadata is invalid", e);
        }
    }

    Map<String, String> toMap() {
        return Map.of(
            VERSION, FORMAT_VERSION,
            CONTENT_ALGORITHM, LocalAesGcm.ALGORITHM,
            IV, Base64.getEncoder().encodeToString(iv),
            KEY_ID, protectedDataKey.keyId(),
            WRAP_ALGORITHM, protectedDataKey.algorithm(),
            WRAPPED_KEY, Base64.getEncoder().encodeToString(protectedDataKey.wrappedKey())
        );
    }

    ProtectedDataKey protectedDataKey() {
        return protectedDataKey;
    }

    byte[] iv() {
        return iv.clone();
    }

    byte[] authenticatedData() {
        return authenticatedData(protectedDataKey);
    }

    static byte[] authenticatedData(ProtectedDataKey protectedDataKey) {
        String canonicalMetadata = String.join(
            "\n",
            FORMAT_VERSION,
            LocalAesGcm.ALGORITHM,
            protectedDataKey.keyId(),
            protectedDataKey.algorithm(),
            Base64.getEncoder().encodeToString(protectedDataKey.wrappedKey())
        );
        return canonicalMetadata.getBytes(StandardCharsets.UTF_8);
    }

    private static String required(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new BlobEncryptionException("Missing blob encryption metadata: " + name);
        }
        return value;
    }
}

package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.LinkedHashMap;
import java.util.Map;

final class EnvelopeMetadata {
    static final String WRAP_ALGORITHM = "RSA-OAEP-256";
    private static final String FORMAT_VERSION = "1";
    private static final String VERSION = "enc_version";
    private static final String CONTENT_ALG = "enc_content_alg";
    private static final String WRAP_ALG = "enc_wrap_alg";
    private static final String KEY_ID = "enc_key_id";
    private static final String WRAPPED_KEY = "enc_wrapped_key";
    private static final String IV = "enc_iv";

    private EnvelopeMetadata() {
    }

    static Map<String, String> create(String keyId, byte[] wrappedKey, byte[] iv) {
        Map<String, String> metadata = new LinkedHashMap<>();
        metadata.put(VERSION, FORMAT_VERSION);
        metadata.put(CONTENT_ALG, CipherSupport.CONTENT_ALGORITHM);
        metadata.put(WRAP_ALG, WRAP_ALGORITHM);
        metadata.put(KEY_ID, keyId);
        metadata.put(WRAPPED_KEY, encode(wrappedKey));
        metadata.put(IV, encode(iv));
        return metadata;
    }

    static Parsed parse(Map<String, String> metadata) {
        String version = required(metadata, VERSION);
        String contentAlgorithm = required(metadata, CONTENT_ALG);
        String wrapAlgorithm = required(metadata, WRAP_ALG);
        if (!FORMAT_VERSION.equals(version)
                || !CipherSupport.CONTENT_ALGORITHM.equals(contentAlgorithm)
                || !WRAP_ALGORITHM.equals(wrapAlgorithm)) {
            throw new EnvelopeEncryptionException("Blob uses an unsupported encryption format");
        }

        byte[] wrappedKey = decode(required(metadata, WRAPPED_KEY), WRAPPED_KEY);
        byte[] iv = decode(required(metadata, IV), IV);
        if (wrappedKey.length == 0 || iv.length != CipherSupport.IV_BYTES) {
            throw new EnvelopeEncryptionException("Blob encryption metadata is malformed");
        }
        return new Parsed(required(metadata, KEY_ID), wrappedKey, iv);
    }

    static byte[] authenticatedData(String blobName, String keyId) {
        return ("azure-envelope-v1\n" + keyId + "\n" + blobName)
                .getBytes(StandardCharsets.UTF_8);
    }

    private static String required(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new EnvelopeEncryptionException(
                    "Blob is missing required encryption metadata: " + name);
        }
        return value;
    }

    private static String encode(byte[] value) {
        return Base64.getEncoder().encodeToString(value);
    }

    private static byte[] decode(String value, String name) {
        try {
            return Base64.getDecoder().decode(value);
        } catch (IllegalArgumentException e) {
            throw new EnvelopeEncryptionException(
                    "Blob encryption metadata is not valid base64: " + name, e);
        }
    }

    record Parsed(String keyId, byte[] wrappedKey, byte[] iv) {
    }
}

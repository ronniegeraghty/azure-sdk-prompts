package com.example.encryptedblob;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import java.nio.charset.StandardCharsets;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Base64;
import java.util.HashMap;
import java.util.Map;

final class EnvelopeEncryption {
    static final String VERSION = "1";
    static final String CONTENT_ALGORITHM = "A256GCM";
    static final String META_VERSION = "ce_version";
    static final String META_CONTENT_ALGORITHM = "ce_algorithm";
    static final String META_IV = "ce_iv";
    static final String META_KEY_ID = "ce_key_id";
    static final String META_WRAP_ALGORITHM = "ce_wrap_algorithm";
    static final String META_WRAPPED_KEY = "ce_wrapped_key";

    private static final int IV_BYTES = 12;
    private static final int TAG_BITS = 128;

    private final SecureRandom secureRandom;

    EnvelopeEncryption() {
        this(new SecureRandom());
    }

    EnvelopeEncryption(SecureRandom secureRandom) {
        this.secureRandom = secureRandom;
    }

    EncryptedPayload encrypt(byte[] plaintext, DataKeyMaterial dataKey, String blobName) {
        byte[] iv = new byte[IV_BYTES];
        secureRandom.nextBytes(iv);
        try {
            Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
            cipher.init(Cipher.ENCRYPT_MODE, dataKey.asAesKey(), new GCMParameterSpec(TAG_BITS, iv));
            cipher.updateAAD(aad(blobName));
            return new EncryptedPayload(cipher.doFinal(plaintext), iv);
        } catch (GeneralSecurityException exception) {
            throw new EncryptedBlobException("Local encryption failed", exception);
        }
    }

    byte[] decrypt(
            byte[] ciphertext,
            byte[] iv,
            DataKeyMaterial dataKey,
            String blobName) {
        try {
            Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
            cipher.init(Cipher.DECRYPT_MODE, dataKey.asAesKey(), new GCMParameterSpec(TAG_BITS, iv));
            cipher.updateAAD(aad(blobName));
            return cipher.doFinal(ciphertext);
        } catch (AEADBadTagException exception) {
            throw new EncryptedBlobException(
                    "Ciphertext or encryption metadata failed authentication",
                    exception);
        } catch (GeneralSecurityException exception) {
            throw new EncryptedBlobException("Local decryption failed", exception);
        }
    }

    Map<String, String> metadata(EncryptedPayload payload, ProtectedDataKey protectedKey) {
        Map<String, String> metadata = new HashMap<>();
        metadata.put(META_VERSION, VERSION);
        metadata.put(META_CONTENT_ALGORITHM, CONTENT_ALGORITHM);
        metadata.put(META_IV, Base64.getEncoder().encodeToString(payload.iv()));
        metadata.put(META_KEY_ID, protectedKey.keyId());
        metadata.put(META_WRAP_ALGORITHM, protectedKey.wrapAlgorithm());
        metadata.put(META_WRAPPED_KEY, Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));
        return metadata;
    }

    EnvelopeMetadata parseMetadata(Map<String, String> metadata) {
        String version = required(metadata, META_VERSION);
        if (!VERSION.equals(version)) {
            throw new EncryptedBlobException("Unsupported envelope metadata version: " + version);
        }
        String contentAlgorithm = required(metadata, META_CONTENT_ALGORITHM);
        if (!CONTENT_ALGORITHM.equals(contentAlgorithm)) {
            throw new EncryptedBlobException(
                    "Unsupported content encryption algorithm: " + contentAlgorithm);
        }

        try {
            byte[] iv = Base64.getDecoder().decode(required(metadata, META_IV));
            if (iv.length != IV_BYTES) {
                throw new EncryptedBlobException("Invalid AES-GCM IV length: " + iv.length);
            }
            ProtectedDataKey protectedKey = new ProtectedDataKey(
                    required(metadata, META_KEY_ID),
                    required(metadata, META_WRAP_ALGORITHM),
                    Base64.getDecoder().decode(required(metadata, META_WRAPPED_KEY)));
            return new EnvelopeMetadata(iv, protectedKey);
        } catch (IllegalArgumentException exception) {
            throw new EncryptedBlobException("Encryption metadata contains invalid Base64", exception);
        }
    }

    private static String required(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new EncryptedBlobException("Blob is missing encryption metadata: " + name);
        }
        return value;
    }

    private static byte[] aad(String blobName) {
        return ("encrypted-blob:v1:" + blobName).getBytes(StandardCharsets.UTF_8);
    }

    record EncryptedPayload(byte[] ciphertext, byte[] iv) {
    }

    record EnvelopeMetadata(byte[] iv, ProtectedDataKey protectedKey) {
    }
}

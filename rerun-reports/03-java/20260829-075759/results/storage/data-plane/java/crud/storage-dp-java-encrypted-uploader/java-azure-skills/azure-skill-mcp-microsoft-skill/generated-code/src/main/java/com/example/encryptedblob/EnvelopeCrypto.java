package com.example.encryptedblob;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;

final class EnvelopeCrypto {
    static final String CONTENT_ALGORITHM = "AES-256-GCM";
    private static final String CIPHER_TRANSFORMATION = "AES/GCM/NoPadding";
    private static final int IV_BYTES = 12;
    private static final int GCM_TAG_BITS = 128;

    private EnvelopeCrypto() {
    }

    static byte[] generateIv(SecureRandom secureRandom) {
        byte[] iv = new byte[IV_BYTES];
        secureRandom.nextBytes(iv);
        return iv;
    }

    static byte[] encrypt(byte[] plaintext, byte[] dataKey, byte[] iv, byte[] authenticatedMetadata) {
        return applyCipher(Cipher.ENCRYPT_MODE, plaintext, dataKey, iv, authenticatedMetadata);
    }

    static byte[] decrypt(byte[] ciphertext, byte[] dataKey, byte[] iv, byte[] authenticatedMetadata) {
        try {
            return applyCipher(Cipher.DECRYPT_MODE, ciphertext, dataKey, iv, authenticatedMetadata);
        } catch (EncryptionStorageException exception) {
            if (exception.getCause() instanceof AEADBadTagException) {
                throw new EncryptionStorageException(
                        "Ciphertext or encryption metadata failed authentication", exception.getCause());
            }
            throw exception;
        }
    }

    static byte[] authenticatedMetadata(String keyId, byte[] wrappedKey) {
        byte[] keyIdBytes = keyId.getBytes(StandardCharsets.UTF_8);
        byte[] contentAlgorithm = CONTENT_ALGORITHM.getBytes(StandardCharsets.US_ASCII);
        byte[] wrapAlgorithm = KeyManagementService.WRAP_ALGORITHM_NAME.getBytes(StandardCharsets.US_ASCII);
        return ByteBuffer.allocate(
                        Integer.BYTES * 4
                                + keyIdBytes.length
                                + wrappedKey.length
                                + contentAlgorithm.length
                                + wrapAlgorithm.length)
                .putInt(keyIdBytes.length).put(keyIdBytes)
                .putInt(wrappedKey.length).put(wrappedKey)
                .putInt(contentAlgorithm.length).put(contentAlgorithm)
                .putInt(wrapAlgorithm.length).put(wrapAlgorithm)
                .array();
    }

    private static byte[] applyCipher(
            int mode,
            byte[] input,
            byte[] dataKey,
            byte[] iv,
            byte[] authenticatedMetadata) {
        try {
            Cipher cipher = Cipher.getInstance(CIPHER_TRANSFORMATION);
            cipher.init(mode, new SecretKeySpec(dataKey, "AES"), new GCMParameterSpec(GCM_TAG_BITS, iv));
            cipher.updateAAD(authenticatedMetadata);
            return cipher.doFinal(input);
        } catch (GeneralSecurityException exception) {
            throw new EncryptionStorageException("Local AES-GCM operation failed", exception);
        }
    }
}

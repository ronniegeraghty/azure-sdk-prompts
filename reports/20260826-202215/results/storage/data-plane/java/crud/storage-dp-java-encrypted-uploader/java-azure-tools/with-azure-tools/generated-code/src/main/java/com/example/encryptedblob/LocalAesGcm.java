package com.example.encryptedblob;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Objects;

final class LocalAesGcm {
    static final String ALGORITHM = "AES/GCM/NoPadding";
    private static final int IV_BYTES = 12;
    private static final int TAG_BITS = 128;
    private static final SecureRandom SECURE_RANDOM = new SecureRandom();

    private LocalAesGcm() {
    }

    static EncryptedPayload encrypt(byte[] dataKey, byte[] plaintext, byte[] authenticatedMetadata) {
        Objects.requireNonNull(dataKey, "dataKey");
        Objects.requireNonNull(plaintext, "plaintext");
        Objects.requireNonNull(authenticatedMetadata, "authenticatedMetadata");

        byte[] iv = new byte[IV_BYTES];
        SECURE_RANDOM.nextBytes(iv);
        try {
            Cipher cipher = Cipher.getInstance(ALGORITHM);
            cipher.init(
                Cipher.ENCRYPT_MODE,
                new SecretKeySpec(dataKey, "AES"),
                new GCMParameterSpec(TAG_BITS, iv)
            );
            cipher.updateAAD(authenticatedMetadata);
            return new EncryptedPayload(iv, cipher.doFinal(plaintext));
        } catch (GeneralSecurityException e) {
            throw new BlobEncryptionException("Local AES-GCM encryption failed", e);
        }
    }

    static byte[] decrypt(
        byte[] dataKey,
        byte[] ciphertext,
        byte[] iv,
        byte[] authenticatedMetadata
    ) {
        Objects.requireNonNull(dataKey, "dataKey");
        Objects.requireNonNull(ciphertext, "ciphertext");
        Objects.requireNonNull(iv, "iv");
        Objects.requireNonNull(authenticatedMetadata, "authenticatedMetadata");
        if (iv.length != IV_BYTES) {
            throw new BlobEncryptionException("Invalid AES-GCM IV length: " + iv.length);
        }

        try {
            Cipher cipher = Cipher.getInstance(ALGORITHM);
            cipher.init(
                Cipher.DECRYPT_MODE,
                new SecretKeySpec(dataKey, "AES"),
                new GCMParameterSpec(TAG_BITS, iv)
            );
            cipher.updateAAD(authenticatedMetadata);
            return cipher.doFinal(ciphertext);
        } catch (AEADBadTagException e) {
            throw new BlobEncryptionException(
                "AES-GCM authentication failed; ciphertext or encryption metadata was modified",
                e
            );
        } catch (GeneralSecurityException e) {
            throw new BlobEncryptionException("Local AES-GCM decryption failed", e);
        }
    }

    record EncryptedPayload(byte[] iv, byte[] ciphertext) {
        EncryptedPayload {
            iv = iv.clone();
            ciphertext = ciphertext.clone();
        }

        @Override
        public byte[] iv() {
            return iv.clone();
        }

        @Override
        public byte[] ciphertext() {
            return ciphertext.clone();
        }
    }
}

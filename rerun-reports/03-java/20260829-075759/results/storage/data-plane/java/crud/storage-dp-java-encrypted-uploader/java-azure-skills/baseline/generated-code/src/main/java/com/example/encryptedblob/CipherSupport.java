package com.example.encryptedblob;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;

final class CipherSupport {
    static final String CONTENT_ALGORITHM = "AES/GCM/NoPadding";
    static final int DATA_KEY_BYTES = 32;
    static final int IV_BYTES = 12;
    private static final int GCM_TAG_BITS = 128;
    private static final SecureRandom RANDOM = new SecureRandom();

    private CipherSupport() {
    }

    static byte[] generateDataKey() {
        byte[] key = new byte[DATA_KEY_BYTES];
        RANDOM.nextBytes(key);
        return key;
    }

    static EncryptedData encrypt(byte[] plaintext, byte[] key, byte[] authenticatedData) {
        byte[] iv = new byte[IV_BYTES];
        RANDOM.nextBytes(iv);
        try {
            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);
            cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(key, "AES"),
                    new GCMParameterSpec(GCM_TAG_BITS, iv));
            cipher.updateAAD(authenticatedData);
            return new EncryptedData(iv, cipher.doFinal(plaintext));
        } catch (GeneralSecurityException e) {
            throw new EnvelopeEncryptionException("Local encryption failed", e);
        }
    }

    static byte[] decrypt(byte[] ciphertext, byte[] key, byte[] iv, byte[] authenticatedData) {
        try {
            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);
            cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(key, "AES"),
                    new GCMParameterSpec(GCM_TAG_BITS, iv));
            cipher.updateAAD(authenticatedData);
            return cipher.doFinal(ciphertext);
        } catch (AEADBadTagException e) {
            throw new EnvelopeEncryptionException(
                    "Ciphertext or encryption metadata failed authentication", e);
        } catch (GeneralSecurityException e) {
            throw new EnvelopeEncryptionException("Local decryption failed", e);
        }
    }

    record EncryptedData(byte[] iv, byte[] ciphertext) {
    }
}

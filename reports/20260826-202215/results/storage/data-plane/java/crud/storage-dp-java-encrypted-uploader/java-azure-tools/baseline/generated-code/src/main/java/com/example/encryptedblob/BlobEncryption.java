package com.example.encryptedblob;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import javax.crypto.spec.GCMParameterSpec;
import java.security.GeneralSecurityException;
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;

final class BlobEncryption {
    static final String ALGORITHM = "A256GCM";
    static final int IV_SIZE_BYTES = 12;
    private static final int GCM_TAG_SIZE_BITS = 128;

    private BlobEncryption() {
    }

    static byte[] newIv(SecureRandom secureRandom) {
        byte[] iv = new byte[IV_SIZE_BYTES];
        secureRandom.nextBytes(iv);
        return iv;
    }

    static byte[] encrypt(byte[] plaintext, DataEncryptionKey key, byte[] iv) {
        try {
            Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
            cipher.init(Cipher.ENCRYPT_MODE, key.asSecretKey(), new GCMParameterSpec(GCM_TAG_SIZE_BITS, iv));
            return cipher.doFinal(plaintext);
        } catch (GeneralSecurityException e) {
            throw new EncryptedBlobException("Local AES-GCM encryption failed", e);
        }
    }

    static byte[] decrypt(byte[] ciphertext, DataEncryptionKey key, byte[] iv) {
        try {
            Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
            cipher.init(Cipher.DECRYPT_MODE, key.asSecretKey(), new GCMParameterSpec(GCM_TAG_SIZE_BITS, iv));
            return cipher.doFinal(ciphertext);
        } catch (AEADBadTagException e) {
            throw new EncryptedBlobException(
                    "Encrypted blob authentication failed; ciphertext or metadata may have been modified", e);
        } catch (GeneralSecurityException e) {
            throw new EncryptedBlobException("Local AES-GCM decryption failed", e);
        }
    }
}

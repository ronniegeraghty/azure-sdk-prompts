package com.example.encryptedblob;

import org.junit.jupiter.api.Test;

import java.nio.charset.StandardCharsets;
import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class CipherSupportTest {
    @Test
    void encryptsAndDecryptsWithAuthenticatedMetadata() {
        byte[] key = CipherSupport.generateDataKey();
        byte[] plaintext = "test data".getBytes(StandardCharsets.UTF_8);
        byte[] aad = EnvelopeMetadata.authenticatedData(
                "example.bin", "https://example.vault.azure.net/keys/test/version");

        try {
            CipherSupport.EncryptedData encrypted =
                    CipherSupport.encrypt(plaintext, key, aad);

            assertArrayEquals(plaintext, CipherSupport.decrypt(
                    encrypted.ciphertext(), key, encrypted.iv(), aad));
        } finally {
            Arrays.fill(key, (byte) 0);
        }
    }

    @Test
    void rejectsChangedAuthenticatedMetadata() {
        byte[] key = CipherSupport.generateDataKey();
        byte[] plaintext = "test data".getBytes(StandardCharsets.UTF_8);
        byte[] aad = EnvelopeMetadata.authenticatedData("example.bin", "key-id");

        try {
            CipherSupport.EncryptedData encrypted =
                    CipherSupport.encrypt(plaintext, key, aad);

            assertThrows(EnvelopeEncryptionException.class, () -> CipherSupport.decrypt(
                    encrypted.ciphertext(),
                    key,
                    encrypted.iv(),
                    EnvelopeMetadata.authenticatedData("other.bin", "key-id")));
        } finally {
            Arrays.fill(key, (byte) 0);
        }
    }
}

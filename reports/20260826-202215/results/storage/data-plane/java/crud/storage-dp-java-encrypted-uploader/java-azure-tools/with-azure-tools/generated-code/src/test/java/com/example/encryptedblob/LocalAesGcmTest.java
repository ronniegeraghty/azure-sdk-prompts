package com.example.encryptedblob;

import org.junit.jupiter.api.Test;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class LocalAesGcmTest {
    @Test
    void roundTripsAuthenticatedCiphertext() {
        byte[] key = randomBytes(32);
        byte[] plaintext = "client-side encrypted".getBytes(StandardCharsets.UTF_8);
        byte[] metadata = "authenticated metadata".getBytes(StandardCharsets.UTF_8);

        LocalAesGcm.EncryptedPayload encrypted =
            LocalAesGcm.encrypt(key, plaintext, metadata);

        assertArrayEquals(
            plaintext,
            LocalAesGcm.decrypt(key, encrypted.ciphertext(), encrypted.iv(), metadata)
        );
    }

    @Test
    void rejectsModifiedAuthenticatedMetadata() {
        byte[] key = randomBytes(32);
        byte[] plaintext = "client-side encrypted".getBytes(StandardCharsets.UTF_8);
        byte[] metadata = "original metadata".getBytes(StandardCharsets.UTF_8);
        LocalAesGcm.EncryptedPayload encrypted =
            LocalAesGcm.encrypt(key, plaintext, metadata);

        assertThrows(
            BlobEncryptionException.class,
            () -> LocalAesGcm.decrypt(
                key,
                encrypted.ciphertext(),
                encrypted.iv(),
                "modified metadata".getBytes(StandardCharsets.UTF_8)
            )
        );
    }

    private static byte[] randomBytes(int size) {
        byte[] bytes = new byte[size];
        new SecureRandom().nextBytes(bytes);
        return bytes;
    }
}

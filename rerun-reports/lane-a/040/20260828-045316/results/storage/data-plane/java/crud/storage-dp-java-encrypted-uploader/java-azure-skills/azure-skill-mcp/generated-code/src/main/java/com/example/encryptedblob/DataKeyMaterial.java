package com.example.encryptedblob;

import javax.crypto.spec.SecretKeySpec;
import java.util.Arrays;

final class DataKeyMaterial implements AutoCloseable {
    private final byte[] keyBytes;

    DataKeyMaterial(byte[] keyBytes) {
        this.keyBytes = keyBytes;
    }

    SecretKeySpec asAesKey() {
        return new SecretKeySpec(keyBytes, "AES");
    }

    byte[] bytesForWrapping() {
        return keyBytes;
    }

    @Override
    public void close() {
        Arrays.fill(keyBytes, (byte) 0);
    }
}

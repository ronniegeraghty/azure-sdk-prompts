package com.example.encryptedblob;

import javax.crypto.SecretKey;
import javax.crypto.spec.SecretKeySpec;
import java.util.Arrays;

final class DataEncryptionKey implements AutoCloseable {
    private final byte[] bytes;

    DataEncryptionKey(byte[] bytes) {
        this.bytes = bytes;
    }

    byte[] bytes() {
        return bytes;
    }

    SecretKey asSecretKey() {
        return new SecretKeySpec(bytes, "AES");
    }

    @Override
    public void close() {
        Arrays.fill(bytes, (byte) 0);
    }
}

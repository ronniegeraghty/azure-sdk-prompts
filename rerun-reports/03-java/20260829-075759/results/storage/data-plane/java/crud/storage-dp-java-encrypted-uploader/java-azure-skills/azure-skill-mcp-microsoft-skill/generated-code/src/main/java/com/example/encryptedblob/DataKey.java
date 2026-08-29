package com.example.encryptedblob;

import java.util.Arrays;

final class DataKey implements AutoCloseable {
    private static final int AES_256_KEY_BYTES = 32;

    private final byte[] bytes;
    private boolean destroyed;

    DataKey(byte[] bytes) {
        if (bytes.length != AES_256_KEY_BYTES) {
            Arrays.fill(bytes, (byte) 0);
            throw new EncryptionStorageException(
                    "Unwrapped data key is not a 256-bit AES key");
        }
        this.bytes = bytes;
    }

    byte[] bytes() {
        if (destroyed) {
            throw new IllegalStateException("Data encryption key has already been destroyed");
        }
        return bytes;
    }

    @Override
    public void close() {
        Arrays.fill(bytes, (byte) 0);
        destroyed = true;
    }
}

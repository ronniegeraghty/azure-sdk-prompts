package com.example.encryptedblob;

import java.util.Arrays;

record ProtectedDataKey(String keyId, byte[] wrappedKey) {
    ProtectedDataKey {
        if (keyId == null || keyId.isBlank()) {
            throw new IllegalArgumentException("keyId must not be blank");
        }
        wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);
    }

    @Override
    public byte[] wrappedKey() {
        return Arrays.copyOf(wrappedKey, wrappedKey.length);
    }
}

final class DataKeyEnvelope implements AutoCloseable {
    private final byte[] plaintextKey;
    private final ProtectedDataKey protectedKey;

    DataKeyEnvelope(byte[] plaintextKey, ProtectedDataKey protectedKey) {
        this.plaintextKey = plaintextKey;
        this.protectedKey = protectedKey;
    }

    byte[] plaintextKey() {
        return plaintextKey;
    }

    ProtectedDataKey protectedKey() {
        return protectedKey;
    }

    @Override
    public void close() {
        Arrays.fill(plaintextKey, (byte) 0);
    }
}

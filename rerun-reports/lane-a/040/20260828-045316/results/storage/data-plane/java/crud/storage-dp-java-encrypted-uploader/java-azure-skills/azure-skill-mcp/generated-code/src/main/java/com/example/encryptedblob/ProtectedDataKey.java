package com.example.encryptedblob;

import java.util.Arrays;

public record ProtectedDataKey(String keyId, String wrapAlgorithm, byte[] wrappedKey) {
    public ProtectedDataKey {
        wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);
    }

    @Override
    public byte[] wrappedKey() {
        return Arrays.copyOf(wrappedKey, wrappedKey.length);
    }
}

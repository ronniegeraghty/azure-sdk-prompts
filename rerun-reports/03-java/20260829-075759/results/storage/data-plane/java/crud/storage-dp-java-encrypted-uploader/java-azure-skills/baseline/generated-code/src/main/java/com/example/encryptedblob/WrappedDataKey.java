package com.example.encryptedblob;

import java.util.Arrays;

public record WrappedDataKey(byte[] plaintextKey, byte[] wrappedKey, String keyId) {
    public WrappedDataKey {
        plaintextKey = Arrays.copyOf(plaintextKey, plaintextKey.length);
        wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);
    }

    @Override
    public byte[] plaintextKey() {
        return Arrays.copyOf(plaintextKey, plaintextKey.length);
    }

    @Override
    public byte[] wrappedKey() {
        return Arrays.copyOf(wrappedKey, wrappedKey.length);
    }

    public void destroy() {
        Arrays.fill(plaintextKey, (byte) 0);
    }
}

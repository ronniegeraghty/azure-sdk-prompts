package com.example.encryptedblob;

import java.util.Objects;

public record ProtectedDataKey(String keyId, String algorithm, byte[] wrappedKey) {
    public ProtectedDataKey {
        Objects.requireNonNull(keyId, "keyId");
        Objects.requireNonNull(algorithm, "algorithm");
        wrappedKey = Objects.requireNonNull(wrappedKey, "wrappedKey").clone();
    }

    @Override
    public byte[] wrappedKey() {
        return wrappedKey.clone();
    }
}

package com.example.encryptedblob;

import java.util.Arrays;
import java.util.Objects;

public record WrappedDataKey(String keyId, String algorithm, byte[] bytes) {
    public WrappedDataKey {
        Objects.requireNonNull(keyId, "keyId");
        Objects.requireNonNull(algorithm, "algorithm");
        bytes = Arrays.copyOf(Objects.requireNonNull(bytes, "bytes"), bytes.length);
    }

    @Override
    public byte[] bytes() {
        return Arrays.copyOf(bytes, bytes.length);
    }
}

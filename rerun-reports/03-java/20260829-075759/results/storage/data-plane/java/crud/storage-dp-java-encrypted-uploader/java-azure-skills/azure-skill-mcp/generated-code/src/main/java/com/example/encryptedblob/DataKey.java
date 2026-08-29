package com.example.encryptedblob;

import java.util.Arrays;
import java.util.Objects;

final class DataKey implements AutoCloseable {
    private final byte[] bytes;

    DataKey(byte[] bytes) {
        this.bytes = Objects.requireNonNull(bytes, "bytes");
    }

    byte[] bytes() {
        return bytes;
    }

    @Override
    public void close() {
        Arrays.fill(bytes, (byte) 0);
    }
}

package com.example.encryptedblob;

import java.util.Arrays;
import java.util.Objects;

final class EnvelopeKey implements AutoCloseable {
    private final ProtectedDataKey protectedDataKey;
    private final byte[] dataKey;
    private boolean closed;

    EnvelopeKey(ProtectedDataKey protectedDataKey, byte[] dataKey) {
        this.protectedDataKey = Objects.requireNonNull(protectedDataKey, "protectedDataKey");
        this.dataKey = Objects.requireNonNull(dataKey, "dataKey");
    }

    ProtectedDataKey protectedDataKey() {
        ensureOpen();
        return protectedDataKey;
    }

    byte[] dataKey() {
        ensureOpen();
        return dataKey;
    }

    @Override
    public void close() {
        if (!closed) {
            Arrays.fill(dataKey, (byte) 0);
            closed = true;
        }
    }

    private void ensureOpen() {
        if (closed) {
            throw new IllegalStateException("The data encryption key has already been destroyed");
        }
    }
}

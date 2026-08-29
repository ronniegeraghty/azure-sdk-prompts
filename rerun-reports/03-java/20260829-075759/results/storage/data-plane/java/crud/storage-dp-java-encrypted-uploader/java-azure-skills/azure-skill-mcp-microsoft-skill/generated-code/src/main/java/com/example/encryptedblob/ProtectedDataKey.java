package com.example.encryptedblob;

final class ProtectedDataKey implements AutoCloseable {
    private final DataKey dataKey;
    private final byte[] wrappedKey;
    private final String keyId;

    ProtectedDataKey(DataKey dataKey, byte[] wrappedKey, String keyId) {
        this.dataKey = dataKey;
        this.wrappedKey = wrappedKey;
        this.keyId = keyId;
    }

    DataKey dataKey() {
        return dataKey;
    }

    byte[] wrappedKey() {
        return wrappedKey;
    }

    String keyId() {
        return keyId;
    }

    @Override
    public void close() {
        dataKey.close();
    }
}

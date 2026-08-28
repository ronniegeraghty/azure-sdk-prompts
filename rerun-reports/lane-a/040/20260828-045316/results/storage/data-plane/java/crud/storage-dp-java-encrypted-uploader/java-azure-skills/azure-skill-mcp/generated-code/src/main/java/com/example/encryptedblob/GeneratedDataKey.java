package com.example.encryptedblob;

final class GeneratedDataKey implements AutoCloseable {
    private final DataKeyMaterial plaintextKey;
    private final ProtectedDataKey protectedKey;

    GeneratedDataKey(DataKeyMaterial plaintextKey, ProtectedDataKey protectedKey) {
        this.plaintextKey = plaintextKey;
        this.protectedKey = protectedKey;
    }

    DataKeyMaterial plaintextKey() {
        return plaintextKey;
    }

    ProtectedDataKey protectedKey() {
        return protectedKey;
    }

    @Override
    public void close() {
        plaintextKey.close();
    }
}

package com.example.encryptedblob;

record GeneratedDataKey(DataKey plaintextKey, WrappedDataKey wrappedKey) implements AutoCloseable {
    @Override
    public void close() {
        plaintextKey.close();
    }
}

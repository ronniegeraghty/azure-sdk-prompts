package com.example.encryptedblob;

import java.util.Base64;
import java.util.Objects;

public record UploadResult(String keyId, byte[] wrappedDataKey) {
    public UploadResult {
        Objects.requireNonNull(keyId, "keyId");
        wrappedDataKey = Objects.requireNonNull(wrappedDataKey, "wrappedDataKey").clone();
    }

    @Override
    public byte[] wrappedDataKey() {
        return wrappedDataKey.clone();
    }

    public String wrappedDataKeyBase64() {
        return Base64.getEncoder().encodeToString(wrappedDataKey);
    }
}

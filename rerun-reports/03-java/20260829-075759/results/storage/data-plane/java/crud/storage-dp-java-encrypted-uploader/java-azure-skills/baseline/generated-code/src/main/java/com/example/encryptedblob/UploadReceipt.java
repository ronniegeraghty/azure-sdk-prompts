package com.example.encryptedblob;

import java.util.Arrays;
import java.util.Base64;

public record UploadReceipt(String keyId, byte[] wrappedDataKey) {
    public UploadReceipt {
        wrappedDataKey = Arrays.copyOf(wrappedDataKey, wrappedDataKey.length);
    }

    @Override
    public byte[] wrappedDataKey() {
        return Arrays.copyOf(wrappedDataKey, wrappedDataKey.length);
    }

    public String wrappedDataKeyBase64() {
        return Base64.getEncoder().encodeToString(wrappedDataKey);
    }
}

package com.example.encryptedblob;

public record UploadResult(String keyId, String wrappedDataKeyBase64) {
}

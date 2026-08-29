package com.example.encryptedblob;

public record EncryptedBlobInfo(String keyId, String wrappedDataKeyBase64) {
}
